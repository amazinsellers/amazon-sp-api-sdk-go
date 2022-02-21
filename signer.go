package amazonspapi

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/amazinsellers/amazon-sp-api-sdk-go/qs"
	"github.com/amazinsellers/amazon-sp-api-sdk-go/resources"
	"net/http"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
)

type Signer struct {
	Region      string
	AWSRegions  map[string]string
	APIEndpoint string
	ISODate     ISODate
}

type ISODate struct {
	Short string
	Full  string
}

func NewSigner(region string) *Signer {
	return &Signer{
		Region: region,
		AWSRegions: map[string]string {
			"eu": "eu-west-1",
			"na": "us-east-1",
			"fe": "us-west-2",
		},
		APIEndpoint: fmt.Sprintf("sellingpartnerapi-%s.amazon.com", region),
	}
}

func (o *Signer) CreateUTCISODate() {
	var re = regexp.MustCompile(`[:\-]|\.\d{3}`)

	isoDate := re.ReplaceAllString(time.Now().UTC().Format(time.RFC3339), "")

	o.ISODate = ISODate{
		Short: isoDate[:8],
		Full: isoDate,
	}
}

func (o Signer) SortQuery(query url.Values) url.Values {
	if len(query) == 0 {
		return url.Values{}
	}

	mk := make([]string, len(query))
	i := 0
	for k := range query {
		mk[i] = k
		i++
	}
	sort.Strings(mk)

	sortedQuery := url.Values{}

	for _, value := range mk {
		sortedQuery[value] = query[value]
	}

	return sortedQuery
}

func (o Signer) SortQueryString(query map[string]string) string {
	if len(query) == 0 {
		return ""
	}

	mk := make([]string, len(query))
	i := 0
	for k := range query {
		mk[i] = k
		i++
	}
	sort.Strings(mk)

	//sortedQuery := map[string]string{}
	sortedQueryString := make([]string, len(query))

	i = 0
	for _, value := range mk {
		sortedQueryString[i] = value + "=" + query[value]
		i++
	}

	if len(sortedQueryString) != 0 {
		return strings.Join(sortedQueryString, "&")
	}

	return ""
}

func (o Signer) ConstructEncodedQueryString(query url.Values) string {
	if len(query) == 0 {
		return ""
	}

	queryInherent := make(map[string]interface{}, len(query))

	for aKey, aValue := range query {
		queryInherent[aKey] = aValue
	}

	queryString, _ := qs.Marshal(queryInherent)
	queryString = strings.Replace(queryString, "+", "%20", -1)

	println(queryString)

	encodedQuery := map[string]string{}
	queryParams := strings.Split(queryString, "&")

	for _, aParam := range queryParams {
		paramKeyValue := strings.Split(aParam, "=")
		encodedQuery[paramKeyValue[0]] = paramKeyValue[1]
	}

	return o.SortQueryString(encodedQuery)
}

func (o Signer) ConstructCanonicalRequestForRoleCredentials(encodedQueryString string) string {
	h := sha256.New()
	h.Write([]byte(encodedQueryString))
	sha := hex.EncodeToString(h.Sum(nil))

	canonical := []string{
		"POST",
		"/",
		"",
		"host:sts.amazonaws.com",
		"x-amz-content-sha256:" + sha,
		"x-amz-date:" + o.ISODate.Full,
		"",
		"host;x-amz-content-sha256;x-amz-date",
		sha,
	}

	return strings.Join(canonical, "\n")
}

func (o Signer) ConstructCanonicalRequestForAPI(
		accessToken string, params resources.SellingPartnerParams,
		encodedQueryString string) string {
	h := sha256.New()
	h.Write([]byte(params.Body))
	sha := hex.EncodeToString(h.Sum(nil))

	canonical := []string{
		params.Method,
		params.APIPath,
		encodedQueryString,
		"host:" + o.APIEndpoint,
		"x-amz-access-token:" + accessToken,
		"x-amz-date:" + o.ISODate.Full,
		"",
		"host;x-amz-access-token;x-amz-date",
		sha,
	}

	return strings.Join(canonical, "\n")
}

func (o Signer) ConstructStringToSign(region string, actionType string, canonicalRequest string) string {
	h := sha256.New()
	h.Write([]byte(canonicalRequest))
	sha := hex.EncodeToString(h.Sum(nil))

	stringToSign := []string {
		"AWS4-HMAC-SHA256",
		o.ISODate.Full,
		fmt.Sprintf("%s/%s/%s/aws4_request",
			o.ISODate.Short, region, actionType),
        sha,
	}

	return strings.Join(stringToSign, "\n")
}

func (o Signer) ConstructSignature(region string, actionType string, stringToSign string, secret string) string {
	h := hmac.New(sha256.New, []byte("AWS4" + secret))
	h.Write([]byte(o.ISODate.Short))
	sign := h.Sum(nil)

	h = hmac.New(sha256.New, sign)
	h.Write([]byte(region))
	sign = h.Sum(nil)

	h = hmac.New(sha256.New, sign)
	h.Write([]byte(actionType))
	sign = h.Sum(nil)

	h = hmac.New(sha256.New, sign)
	h.Write([]byte("aws4_request"))
	sign = h.Sum(nil)

	h = hmac.New(sha256.New, sign)
	h.Write([]byte(stringToSign))

	return hex.EncodeToString(h.Sum(nil))
}

func (o Signer) ConstructURL(params resources.SellingPartnerParams, encodedQueryString string) string {
	url := fmt.Sprintf("https://%s%s", o.APIEndpoint, params.APIPath)

	if encodedQueryString != "" {
		url += "?" + encodedQueryString
	}

	return url
}

func (o *Signer) SignAPIRequest(accessToken string,
			config RoleCredentialsConfig, params resources.SellingPartnerParams) (*http.Request, error) {
	params.Query = o.SortQuery(params.Query)
	o.CreateUTCISODate()

	encodedQueryString := o.ConstructEncodedQueryString(params.Query)
	canonicalRequest := o.ConstructCanonicalRequestForAPI(accessToken, params, encodedQueryString)
	stringToSign := o.ConstructStringToSign(o.AWSRegions[o.Region], "execute-api", canonicalRequest)
	signature := o.ConstructSignature(o.AWSRegions[o.Region], "execute-api", stringToSign, config.Secret)

	url := o.ConstructURL(params, encodedQueryString)
	req, err := http.NewRequest(params.Method, url, strings.NewReader(params.Body))
	req.Host = o.APIEndpoint

	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization",
		"AWS4-HMAC-SHA256 Credential=" +
		config.Id + "/" + o.ISODate.Short +
		"/" + o.AWSRegions[o.Region] +
		"/execute-api/aws4_request, SignedHeaders=host;x-amz-access-token;x-amz-date, Signature=" +
		signature)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("host", o.APIEndpoint)
	req.Header.Set("x-amz-access-token", accessToken)
	req.Header.Set("x-amz-security-token", config.SecurityToken)
	req.Header.Set("x-amz-date", o.ISODate.Full)

	return req, nil
}

func (o *Signer) SignRoleCredentialsRequest(config AWSUserConfig) (*http.Request, error) {
	query := url.Values {
		"Action": {"AssumeRole"},
		"DurationSeconds": {"3600"},
		"RoleArn": {config.Role},
		"RoleSessionName": {"SPAPISession"},
		"Version": {"2011-06-15"},
    }

	o.CreateUTCISODate()
	encodedQueryString := o.ConstructEncodedQueryString(query)
	canonicalRequest := o.ConstructCanonicalRequestForRoleCredentials(encodedQueryString)
	stringToSign := o.ConstructStringToSign("us-east-1", "sts", canonicalRequest)
	signature := o.ConstructSignature("us-east-1", "sts", stringToSign, config.Secret)

	url := "https://sts.amazonaws.com"
	req, err := http.NewRequest("POST", url, strings.NewReader(encodedQueryString))
	req.Host = "sts.amazonaws.com"

	if err != nil {
		return nil, err
	}

	h := sha256.New()
	h.Write([]byte(encodedQueryString))
	contentSha := hex.EncodeToString(h.Sum(nil))

	req.Header.Set("Authorization",
		"AWS4-HMAC-SHA256 Credential=" +
			config.Id + "/" + o.ISODate.Short +
			"/us-east-1" +
			"/sts/aws4_request, SignedHeaders=host;x-amz-content-sha256;x-amz-date, Signature=" +
			signature)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=utf-8")
	req.Header.Set("Host", "sts.amazonaws.com")
	req.Header.Set("X-Amz-Content-Sha256", contentSha)
	req.Header.Set("X-Amz-Date", o.ISODate.Full)

	return req, nil
}
