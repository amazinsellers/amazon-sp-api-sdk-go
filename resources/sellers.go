package resources

import "time"

func GetMarketplaceParticipations(params *SellingPartnerParams) error {
	params.Method = "GET"
	params.APIPath = "/sellers/v1/marketplaceParticipations"
	params.RestoreRate = 60 * time.Millisecond
	return nil
}
