package resources

import "time"

func GetAuthorizationCode(params *SellingPartnerParams) error {
	params.Method = "GET"
	params.APIPath = "/authorization/v1/authorizationCode"
	params.RestoreRate = 1 * time.Second

	return nil
}
