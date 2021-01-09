package resources

import "time"

func GetOrders(params *SellingPartnerParams) error {
	params.Method = "GET"
	params.APIPath = "/orders/v0/orders"
	params.RestoreRate = 1 * time.Second

	return nil
}
