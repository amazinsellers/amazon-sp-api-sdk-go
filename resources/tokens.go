package resources

import (
	"time"
)

func CreateRestrictedDataToken(params *SellingPartnerParams) error {
	params.Method = "POST"
	params.APIPath = "/tokens/2021-03-01/restrictedDataToken"
	params.RestoreRate = 60 * time.Second
	return nil
}
