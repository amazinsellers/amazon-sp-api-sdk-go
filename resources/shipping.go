package resources

import (
	"fmt"
	"time"
)

func GetShipment(params *SellingPartnerParams) error {
	if _, present := params.PathParams["shipmentId"]; !present {
		return fmt.Errorf("path param 'shipmentId' not present")
	}

	params.Method = "GET"
	params.APIPath = "/shipping/v1/shipments/" + params.PathParams["shipmentId"]
	params.RestoreRate = 200 * time.Millisecond
	return nil
}

func CancelShipment(params *SellingPartnerParams) error {
	if _, present := params.PathParams["shipmentId"]; !present {
		return fmt.Errorf("path param 'shipmentId' not present")
	}

	params.Method = "POST"
	params.APIPath = "/shipping/v1/shipments/" + params.PathParams["shipmentId"] + "/cancel"
	params.RestoreRate = 200 * time.Millisecond
	return nil
}

func PurchaseLabels(params *SellingPartnerParams) error {
	if _, present := params.PathParams["shipmentId"]; !present {
		return fmt.Errorf("path param 'shipmentId' not present")
	}

	params.Method = "POST"
	params.APIPath = "/shipping/v1/shipments/" + params.PathParams["shipmentId"] + "/purchaseLabels"
	params.RestoreRate = 200 * time.Millisecond
	return nil
}

func RetrieveShippingLabel(params *SellingPartnerParams) error {
	if _, present := params.PathParams["shipmentId"]; !present {
		return fmt.Errorf("path param 'shipmentId' not present")
	}

	if _, present := params.PathParams["trackingId"]; !present {
		return fmt.Errorf("path param 'trackingId' not present")
	}

	params.Method = "POST"
	params.APIPath = "/shipping/v1/shipments/" + params.PathParams["shipmentId"] + "/containers/" + params.PathParams["trackingId"] + "/label"
	params.RestoreRate = 200 * time.Millisecond
	return nil
}

func PurchaseShipment(params *SellingPartnerParams) error {
	params.Method = "POST"
	params.APIPath = "/shipping/v1/purchaseShipment"
	params.RestoreRate = 200 * time.Millisecond
	return nil
}

func GetRates(params *SellingPartnerParams) error {

	params.Method = "POST"
	params.APIPath = "/shipping/v1/rates"
	params.RestoreRate = 200 * time.Millisecond
	return nil
}

func GetAccount(params *SellingPartnerParams) error {

	params.Method = "GET"
	params.APIPath = "/shipping/v1/account"
	params.RestoreRate = 200 * time.Millisecond
	return nil
}

func GetTrackingInformation(params *SellingPartnerParams) error {
	if _, present := params.PathParams["trackingId"]; !present {
		return fmt.Errorf("path param 'trackingId' not present")
	}

	params.Method = "GET"
	params.APIPath = "/shipping/v1/tracking/" + params.PathParams["trackingId"]
	params.RestoreRate = 1 * time.Second
	return nil
}
