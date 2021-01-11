package resources

import "time"

func GetInventorySummaries(params *SellingPartnerParams) error {
	params.Method = "GET"
	params.APIPath = "/fba/inventory/v1/summaries"
	params.RestoreRate = 12 * time.Millisecond

	return nil
}
