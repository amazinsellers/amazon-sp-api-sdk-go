package resources

import (
	"fmt"
	"time"
)

func GetOrders(params *SellingPartnerParams) error {
	params.Method = "GET"
	params.APIPath = "/orders/v0/orders"
	params.RestoreRate = 1 * time.Second
	return nil
}

func GetOrder(params *SellingPartnerParams) error {
	if _, present := params.PathParams["orderId"]; !present {
		return fmt.Errorf("path param 'orderId' not present")
	}

	params.Method = "GET"
	params.APIPath = "/orders/v0/orders/" + params.PathParams["orderId"]
	params.RestoreRate = 1 * time.Second
	return nil
}

func GetOrderBuyerInfo(params *SellingPartnerParams) error {
	if _, present := params.PathParams["orderId"]; !present {
		return fmt.Errorf("path param 'orderId' not present")
	}

	params.Method = "GET"
	params.APIPath = "/orders/v0/orders/" + params.PathParams["orderId"] + "/buyerInfo"
	params.RestoreRate = 1 * time.Second
	return nil
}

func GetOrderAddress(params *SellingPartnerParams) error {
	if _, present := params.PathParams["orderId"]; !present {
		return fmt.Errorf("path param 'orderId' not present")
	}

	params.Method = "GET"
	params.APIPath = "/orders/v0/orders/" + params.PathParams["orderId"] + "/address"
	params.RestoreRate = 1 * time.Second
	return nil
}

func GetOrderItems(params *SellingPartnerParams) error {
	if _, present := params.PathParams["orderId"]; !present {
		return fmt.Errorf("path param 'orderId' not present")
	}

	params.Method = "GET"
	params.APIPath = "/orders/v0/orders/" + params.PathParams["orderId"] + "/orderItems"
	params.RestoreRate = 1 * time.Second
	return nil
}

func GetOrderItemsBuyerInfo(params *SellingPartnerParams) error {
	if _, present := params.PathParams["orderId"]; !present {
		return fmt.Errorf("path param 'orderId' not present")
	}

	params.Method = "GET"
	params.APIPath = "/orders/v0/orders/" + params.PathParams["orderId"] + "/orderItems/buyerInfo"
	params.RestoreRate = 1 * time.Second
	return nil
}
