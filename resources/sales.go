package resources

import "time"

func GetOrderMetrics(params *SellingPartnerParams) error {
	params.Method = "GET"
	params.APIPath = "/sales/v1/orderMetrics"
	params.RestoreRate = 2 * time.Millisecond
	return nil
}
