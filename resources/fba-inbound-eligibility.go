package resources

import "time"

func GetItemEligibilityPreview(params *SellingPartnerParams) error {
	params.Method = "GET"
	params.APIPath = "/fba/inbound/v1/eligibility/itemPreview"
	params.RestoreRate = 1 * time.Second

	return nil
}
