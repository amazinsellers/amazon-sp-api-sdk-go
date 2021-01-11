package resources

import (
	"fmt"
	"time"
)

func GetFulfillmentPreview(params *SellingPartnerParams) error {
	params.Method = "POST"
	params.APIPath = "/fba/outbound/2020-07-01/fulfillmentOrders/preview"
	params.RestoreRate = 500 * time.Millisecond
	return nil
}

func ListAllFulfillmentOrders(params *SellingPartnerParams) error {
	params.Method = "GET"
	params.APIPath = "/fba/outbound/2020-07-01/fulfillmentOrders"
	params.RestoreRate = 500 * time.Millisecond
	return nil
}

func CreateFulfillmentOrder(params *SellingPartnerParams) error {
	params.Method = "POST"
	params.APIPath = "/fba/outbound/2020-07-01/fulfillmentOrders"
	params.RestoreRate = 500 * time.Millisecond
	return nil
}

func GetPackageTrackingDetails(params *SellingPartnerParams) error {
	params.Method = "GET"
	params.APIPath = "/fba/outbound/2020-07-01/tracking"
	params.RestoreRate = 500 * time.Millisecond
	return nil
}

func ListReturnReasonCodes(params *SellingPartnerParams) error {
	params.Method = "GET"
	params.APIPath = "/fba/outbound/2020-07-01/returnReasonCodes"
	params.RestoreRate = 500 * time.Millisecond
	return nil
}

func CreateFulfillmentReturn(params *SellingPartnerParams) error {
	if _, present := params.PathParams["sellerFulfillmentOrderId"]; !present {
		return fmt.Errorf("path param 'sellerFulfillmentOrderId' not present")
	}

	params.Method = "PUT"
	params.APIPath = "/fba/outbound/2020-07-01/fulfillmentOrders/" + params.PathParams["sellerFulfillmentOrderId"] + "/return"
	params.RestoreRate = 500 * time.Millisecond
	return nil
}

func GetFulfillmentOrder(params *SellingPartnerParams) error {
	if _, present := params.PathParams["sellerFulfillmentOrderId"]; !present {
		return fmt.Errorf("path param 'sellerFulfillmentOrderId' not present")
	}

	params.Method = "GET"
	params.APIPath = "/fba/outbound/2020-07-01/fulfillmentOrders/" + params.PathParams["sellerFulfillmentOrderId"]
	params.RestoreRate = 500 * time.Millisecond
	return nil
}

func UpdateFulfillmentOrder(params *SellingPartnerParams) error {
	if _, present := params.PathParams["sellerFulfillmentOrderId"]; !present {
		return fmt.Errorf("path param 'sellerFulfillmentOrderId' not present")
	}

	params.Method = "GET"
	params.APIPath = "/fba/outbound/2020-07-01/fulfillmentOrders/" + params.PathParams["sellerFulfillmentOrderId"]
	params.RestoreRate = 500 * time.Millisecond
	return nil
}

func CancelFulfillmentOrder(params *SellingPartnerParams) error {
	if _, present := params.PathParams["sellerFulfillmentOrderId"]; !present {
		return fmt.Errorf("path param 'sellerFulfillmentOrderId' not present")
	}

	params.Method = "PUT"
	params.APIPath = "/fba/outbound/2020-07-01/fulfillmentOrders/" + params.PathParams["sellerFulfillmentOrderId"] + "/cancel"
	params.RestoreRate = 500 * time.Millisecond
	return nil
}

func GetFeatures(params *SellingPartnerParams) error {
	params.Method = "GET"
	params.APIPath = "/fba/outbound/2020-07-01/features"
	params.RestoreRate = 500 * time.Millisecond
	return nil
}

func GetFeatureInventory(params *SellingPartnerParams) error {
	if _, present := params.PathParams["featureName"]; !present {
		return fmt.Errorf("path param 'featureName' not present")
	}

	params.Method = "GET"
	params.APIPath = "/fba/outbound/2020-07-01/features/inventory/" + params.PathParams["featureName"]
	params.RestoreRate = 500 * time.Millisecond
	return nil
}

func GetFeatureSKU(params *SellingPartnerParams) error {
	if _, present := params.PathParams["featureName"]; !present {
		return fmt.Errorf("path param 'featureName' not present")
	}
	if _, present := params.PathParams["sellerSku"]; !present {
		return fmt.Errorf("path param 'sellerSku' not present")
	}
	params.Method = "GET"
	params.APIPath = "/fba/outbound/2020-07-01/features/inventory/" + params.PathParams["featureName"] + "/" + params.PathParams["sellerSku"]
	params.RestoreRate = 500 * time.Millisecond
	return nil
}
