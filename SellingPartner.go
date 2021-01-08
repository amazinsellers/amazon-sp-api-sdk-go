package amazonspapi

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	xj "github.com/basgys/goxml2json"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
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
	RoleCredentials RoleCredentials
}

type SellingPartnerParams struct {
	Method string
	APIPath string
	Body string
	Query map[string]interface{}
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
	ErrorResponse ErrorResponse
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

type ErrorResponse struct {
	Error RefreshRoleError
}

type RefreshRoleError struct {
	Code string
	Message string
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
		o.RoleCredentials = refreshRoleResponse.AssumeRoleResponse.AssumeRoleResult.Credentials
		return nil
	}

	return fmt.Errorf("no role Credentials received. Body: %s", respBody)
}
