package amazonspapi

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/amazinsellers/amazon-sp-api-sdk-go/resources"
	xj "github.com/basgys/goxml2json"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"regexp"
	"time"
)

type AccessTokenResponse struct {
	AccessToken      string `json:"access_token"`
	RefreshToken     string `json:"refresh_token"`
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
}

type SellingPartner struct {
	Config          *SellingPartnerConfig
	Credentials     *CredentialsConfig
	AccessToken     string
	RoleCredentials *RoleCredentials
}

func NewSellingPartner(config *SellingPartnerConfig) (*SellingPartner, error) {
	if isValid, err := config.IsValid(); !isValid {
		return nil, err
	}

	if config.Options == nil {
		config.Options = NewOptionsConfig()
	}

	sp := &SellingPartner{}
	sp.Config = config

	if credentials, err := NewCredentialsConfig(); err != nil {
		return nil, err
	} else {
		sp.Credentials = credentials
	}

	return sp, nil
}

func (o *SellingPartner) RefreshToken() error {
	reqBody, _ := json.Marshal(map[string]string{
		"grant_type": "refresh_token",
		"refresh_token": o.Config.RefreshToken,
		"client_id": o.Credentials.AppClient.Id,
		"client_secret": o.Credentials.AppClient.Secret,
	})

	resp, err := http.Post(
		"https://api.amazon.com/auth/o2/token",
		"application/json",
		bytes.NewBuffer(reqBody))

	if err != nil {
		return errors.New("RefreshToken call failed with error " + err.Error())
	}

	defer resp.Body.Close()

	respBodyBytes, _ := ioutil.ReadAll(resp.Body)
	theResp := &AccessTokenResponse{}

	if err = json.Unmarshal(respBodyBytes, theResp); err != nil {
		return errors.New("RefreshToken response parse failed. Body: " + string(respBodyBytes))
	}

	if theResp.AccessToken != "" {
		o.AccessToken = theResp.AccessToken
	} else if theResp.Error != "" {
		return errors.New(fmt.Sprintf(
			"RefreshToken failed with code %s, description %s", theResp.Error, theResp.ErrorDescription))
	} else {
		return errors.New(fmt.Sprintf(
			"RefreshToken failed with unknown reason. Body: %s", string(respBodyBytes)))
	}

	return nil
}

type RefreshRoleResponse struct {
	AssumeRoleResponse AssumeRoleResponse
	ErrorResponse      RefreshRoleErrorResponse
}

type AssumeRoleResponse struct {
	AssumeRoleResult AssumeRoleResult
}

type AssumeRoleResult struct {
	Credentials RoleCredentials
}

type RoleCredentials struct {
	AccessKeyId string
	SecretAccessKey string
	SessionToken string
}

type RefreshRoleErrorResponse struct {
	Error RefreshRoleError
}

type RefreshRoleError struct {
	Code string
	Message string
}

type APIResponse struct {
	Errors  []APIResponseError      `json:"errors"`
	Payload *map[string]interface{} `json:"payload"`
}

type APIResponseError struct {
	Code    string `json:"code"`
	Details string `json:"details"`
	Message string `json:"message"`
}

func (o *SellingPartner) RefreshRoleCredentials() error {
	signedRequest, err := NewSigner(o.Config.Region).SignRoleCredentialsRequest(*o.Credentials.AWSUser)

	if err != nil {
		return err
	}

	if o.Config.Options.Debug {
		requestDump, err := httputil.DumpRequest(signedRequest, true)
		if err == nil {
			fmt.Println("==========================================")
			fmt.Print("RefreshRoleCredentials request dump:\n\n")
			fmt.Println(string(requestDump))
		}
	}

	response, err := http.DefaultTransport.RoundTrip(signedRequest)

	if err != nil {
		return err
	}

	if o.Config.Options.Debug {
		dumpResponse, err := httputil.DumpResponse(response, true)
		if err == nil {
			fmt.Println("==========================================")
			fmt.Print("RefreshRoleCredentials response dump:\n\n")
			fmt.Println(string(dumpResponse))
		}
	}

	respBodyBytes, _ := ioutil.ReadAll(response.Body)
	respBody, err := xj.Convert(bytes.NewReader(respBodyBytes))
	if err != nil {
		return fmt.Errorf("failed to parse xml: %s", response.Body)
	}

	jsonRes := respBody.String()

	if o.Config.Options.Debug {
		fmt.Println("==========================================")
		fmt.Print("Response converted to json:\n\n")
		fmt.Println(jsonRes)
	}

	refreshRoleResponse := &RefreshRoleResponse{}
	err = json.Unmarshal([]byte(jsonRes), refreshRoleResponse)

	if err != nil {
		return fmt.Errorf("failed to parse unmarshal json: %s", jsonRes)
	}

	if refreshRoleResponse.ErrorResponse.Error.Code != "" {
		return fmt.Errorf("failed refresh role. Code: %s. Description: %s",
			refreshRoleResponse.ErrorResponse.Error.Code, refreshRoleResponse.ErrorResponse.Error.Message)
	}

	if refreshRoleResponse.AssumeRoleResponse.AssumeRoleResult.Credentials.AccessKeyId != "" {
		resCredentials := refreshRoleResponse.AssumeRoleResponse.AssumeRoleResult.Credentials
		o.Config.RoleCredentials = &RoleCredentialsConfig{
			Id: resCredentials.AccessKeyId,
			Secret: resCredentials.SecretAccessKey,
			SecurityToken: resCredentials.SessionToken,
		}
		return nil
	}

	return fmt.Errorf("no role Credentials received. Body: %s", respBody)
}

func (o *SellingPartner) CallAPI(params resources.SellingPartnerParams) (*string, error) {
	if params.Operation == "" {
		return nil, fmt.Errorf("operation is a required parameter")
	}

	applyParams, present := AvailableOperations[params.Operation]
	if !present {
		return nil, fmt.Errorf(`operation "%s" not found`, params.Operation)
	}

	if o.Config.Options.AutoRequestTokens {
		if o.AccessToken == "" {
			if err := o.RefreshToken(); err != nil {
				return nil, fmt.Errorf("cannot refresh token. Error: %s", err.Error())
			}
		}

		if o.RoleCredentials == nil {
			if err := o.RefreshRoleCredentials(); err != nil {
				return nil, fmt.Errorf("cannot refresh rrole credentials. Error: %s", err.Error())
			}
		}
	}

	if o.AccessToken == "" || o.Config.RoleCredentials == nil {
		return nil, fmt.Errorf("no access token or role credentials found")
	}

	if err := applyParams(&params); err != nil {
		return nil, fmt.Errorf("cannot apply operation params. Error: " + err.Error())
	}

	signer := NewSigner(o.Config.Region)

	signedRequest, err := signer.SignAPIRequest(
							o.AccessToken, *o.Config.RoleCredentials, params)

	if err != nil {
		return nil, fmt.Errorf("cannot sign api request. Error: %s", err.Error())
	}

	if o.Config.Options.Debug {
		requestDump, err := httputil.DumpRequest(signedRequest, true)
		if err == nil {
			fmt.Println("==========================================")
			fmt.Print("CallAPI request dump:\n\n")
			fmt.Println(string(requestDump))
		}
	}

	response, err := http.DefaultTransport.RoundTrip(signedRequest)

	if err != nil {
		return nil, err
	}

	if o.Config.Options.Debug {
		dumpResponse, err := httputil.DumpResponse(response, true)
		if err == nil {
			fmt.Println("==========================================")
			fmt.Print("CallAPI response dump:\n\n")
			fmt.Println(string(dumpResponse))
		}
	}

	if response.StatusCode == 204 && params.Method == "DELETE" {
		successRes := `{"success": "true"}`
		return &successRes, nil
	}

	apiResponse := &APIResponse{}
	apiResponseBody, _ := ioutil.ReadAll(response.Body)
	if err = json.Unmarshal(apiResponseBody, apiResponse); err != nil {
		return nil, fmt.Errorf("cannot unmarshal api response body. Err: %s", err.Error())
	}

	if len(apiResponse.Errors) == 0 {
		if apiResponse.Payload != nil {
			if payload, err := json.Marshal(apiResponse.Payload); err == nil {
				payloadStr := string(payload)
				return &payloadStr, nil
			} else {
				return nil, fmt.Errorf("cannot convert payload to string")
			}
		}
		apiResponseBodyStr := string(apiResponseBody)
		return &apiResponseBodyStr, nil
	}

	theError := apiResponse.Errors[0]

	if response.StatusCode == 403 && theError.Code == "Unauthorized" {
		if o.Config.Options.AutoRequestTokens {
			if doesMatch, _ := regexp.MatchString("access token.*expired", theError.Details); doesMatch {
				if err := o.RefreshToken(); err == nil {
					return o.CallAPI(params)
				} else {
					return nil, err
				}
			} else if doesMatch, _ := regexp.MatchString("security token.*expired", theError.Message); doesMatch {
				if err := o.RefreshRoleCredentials(); err == nil {
					return o.CallAPI(params)
				} else {
					return nil, err
				}
			}
		}
	} else if response.StatusCode == 429 && theError.Code == "QuotaExceeded" && o.Config.Options.AutoRequestThrottled {
		time.Sleep(params.RestoreRate)
		return o.CallAPI(params)
	}

	errJson, _ := json.Marshal(theError)
	errStr := string(errJson)
	return &errStr, fmt.Errorf("unknown error: %s", errStr)
}
