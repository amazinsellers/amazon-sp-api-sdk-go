package resources

import (
	"fmt"
	"time"
)

func GetEligibleShipmentServicesOld(params *SellingPartnerParams) error {
	params.Method = "POST"
	params.APIPath = "/mfn/v0/eligibleServices"
	params.RestoreRate = 1 * time.Second
	return nil
}

func GetEligibleShipmentServices(params *SellingPartnerParams) error {
	params.Method = "POST"
	params.APIPath = "/mfn/v0/eligibleShippingServices"
	params.RestoreRate = 1 * time.Second
	return nil
}

func GetMfnShipment(params *SellingPartnerParams) error {
	if _, present := params.PathParams["shipmentId"]; !present {
		return fmt.Errorf("path param 'shipmentId' not present")
	}

	params.Method = "GET"
	params.APIPath = "/mfn/v0/shipments/" + params.PathParams["shipmentId"]
	params.RestoreRate = 1 * time.Second
	return nil
}

func CancelMfnShipment(params *SellingPartnerParams) error {
	if _, present := params.PathParams["shipmentId"]; !present {
		return fmt.Errorf("path param 'shipmentId' not present")
	}

	params.Method = "DELETE"
	params.APIPath = "/mfn/v0/shipments/" + params.PathParams["shipmentId"]
	params.RestoreRate = 1 * time.Second
	return nil
}

func CancelShipmentOld(params *SellingPartnerParams) error {
	if _, present := params.PathParams["shipmentId"]; !present {
		return fmt.Errorf("path param 'shipmentId' not present")
	}

	params.Method = "PUT"
	params.APIPath = "/mfn/v0/shipments/" + params.PathParams["shipmentId"] + "/cancel"
	params.RestoreRate = 1 * time.Second
	return nil
}

func CreateShipment(params *SellingPartnerParams) error {
	params.Method = "POST"
	params.APIPath = "/mfn/v0/shipments"
	params.RestoreRate = 1 * time.Second
	return nil
}

func GetAdditionalSellerInputsOld(params *SellingPartnerParams) error {
	params.Method = "POST"
	params.APIPath = "/mfn/v0/sellerInputs"
	params.RestoreRate = 1 * time.Second
	return nil
}

func GetAdditionalSellerInputs(params *SellingPartnerParams) error {
	params.Method = "POST"
	params.APIPath = "/mfn/v0/additionalSellerInputs"
	params.RestoreRate = 1 * time.Second
	return nil
}
