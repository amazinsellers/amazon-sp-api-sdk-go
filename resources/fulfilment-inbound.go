package resources

import (
	"fmt"
	"time"
)

func GetInboundGuidance(params *SellingPartnerParams) error {
	params.Method = "GET"
	params.APIPath = "/fba/inbound/v0/itemsGuidance"
	params.RestoreRate = 500 * time.Millisecond

	return nil
}

func CreateInboundShipmentPlan(params *SellingPartnerParams) error {
	params.Method = "POST"
	params.APIPath = "/fba/inbound/v0/plans"
	params.RestoreRate = 500 * time.Millisecond

	return nil
}

func UpdateInboundShipment(params *SellingPartnerParams) error {
	if _,present := params.PathParams["shipmentId"]; !present {
		return fmt.Errorf("path param 'shipmentId' not present")
	}

	params.Method = "GET"
	params.APIPath = fmt.Sprintf("/fba/inbound/v0/shipments/%s", params.PathParams["shipmentId"])
	params.RestoreRate = 500 * time.Millisecond

	return nil
}

func CreateInboundShipment(params *SellingPartnerParams) error {
	if _,present := params.PathParams["shipmentId"]; !present {
		return fmt.Errorf("path param 'shipmentId' not present")
	}

	params.Method = "POST"
	params.APIPath = fmt.Sprintf("/fba/inbound/v0/shipments/%s", params.PathParams["shipmentId"])
	params.RestoreRate = 500 * time.Millisecond

	return nil
}

func GetPreorderInfo(params *SellingPartnerParams) error {
	if _,present := params.PathParams["shipmentId"]; !present {
		return fmt.Errorf("path param 'shipmentId' not present")
	}

	params.Method = "GET"
	params.APIPath = fmt.Sprintf("/fba/inbound/v0/shipments/%s/preorder", params.PathParams["shipmentId"])
	params.RestoreRate = 500 * time.Millisecond

	return nil
}

func ConfirmPreorder(params *SellingPartnerParams) error {
	if _,present := params.PathParams["shipmentId"]; !present {
		return fmt.Errorf("path param 'shipmentId' not present")
	}

	params.Method = "PUT"
	params.APIPath = fmt.Sprintf("/fba/inbound/v0/shipments/%s/preorder", params.PathParams["shipmentId"])
	params.RestoreRate = 500 * time.Millisecond

	return nil
}

func GetPrepInstructions(params *SellingPartnerParams) error {
	params.Method = "GET"
	params.APIPath = "/fba/inbound/v0/prepInstructions"
	params.RestoreRate = 500 * time.Millisecond

	return nil
}

func GetTransportDetails(params *SellingPartnerParams) error {
	if _,present := params.PathParams["shipmentId"]; !present {
		return fmt.Errorf("path param 'shipmentId' not present")
	}

	params.Method = "GET"
	params.APIPath = fmt.Sprintf("/fba/inbound/v0/shipments/%s/transport", params.PathParams["shipmentId"])
	params.RestoreRate = 500 * time.Millisecond

	return nil
}

func PutTransportDetails(params *SellingPartnerParams) error {
	if _,present := params.PathParams["shipmentId"]; !present {
		return fmt.Errorf("path param 'shipmentId' not present")
	}

	params.Method = "PUT"
	params.APIPath = fmt.Sprintf("/fba/inbound/v0/shipments/%s/transport", params.PathParams["shipmentId"])
	params.RestoreRate = 500 * time.Millisecond

	return nil
}

func VoidTransport(params *SellingPartnerParams) error {
	if _,present := params.PathParams["shipmentId"]; !present {
		return fmt.Errorf("path param 'shipmentId' not present")
	}

	params.Method = "POST"
	params.APIPath = fmt.Sprintf("/fba/inbound/v0/shipments/%s/transport/void", params.PathParams["shipmentId"])
	params.RestoreRate = 500 * time.Millisecond

	return nil
}

func EstimateTransport(params *SellingPartnerParams) error {
	if _,present := params.PathParams["shipmentId"]; !present {
		return fmt.Errorf("path param 'shipmentId' not present")
	}

	params.Method = "POST"
	params.APIPath = fmt.Sprintf("/fba/inbound/v0/shipments/%s/transport/estimate", params.PathParams["shipmentId"])
	params.RestoreRate = 500 * time.Millisecond

	return nil
}

func ConfirmTransport(params *SellingPartnerParams) error {
	if _,present := params.PathParams["shipmentId"]; !present {
		return fmt.Errorf("path param 'shipmentId' not present")
	}

	params.Method = "POST"
	params.APIPath = fmt.Sprintf("/fba/inbound/v0/shipments/%s/transport/confirm", params.PathParams["shipmentId"])
	params.RestoreRate = 500 * time.Millisecond

	return nil
}

func GetLabels(params *SellingPartnerParams) error {
	if _,present := params.PathParams["shipmentId"]; !present {
		return fmt.Errorf("path param 'shipmentId' not present")
	}

	params.Method = "GET"
	params.APIPath = fmt.Sprintf("/fba/inbound/v0/shipments/%s/labels", params.PathParams["shipmentId"])
	params.RestoreRate = 500 * time.Millisecond

	return nil
}

func GetBillOfLading(params *SellingPartnerParams) error {
	if _,present := params.PathParams["shipmentId"]; !present {
		return fmt.Errorf("path param 'shipmentId' not present")
	}

	params.Method = "GET"
	params.APIPath = fmt.Sprintf("/fba/inbound/v0/shipments/%s/billOfLading", params.PathParams["shipmentId"])
	params.RestoreRate = 500 * time.Millisecond

	return nil
}

func GetShipments(params *SellingPartnerParams) error {
	params.Method = "GET"
	params.APIPath = "/fba/inbound/v0/shipments"
	params.RestoreRate = 500 * time.Millisecond

	return nil
}

func GetShipmentItemsByShipmentId(params *SellingPartnerParams) error {
	if _,present := params.PathParams["shipmentId"]; !present {
		return fmt.Errorf("path param 'shipmentId' not present")
	}

	params.Method = "GET"
	params.APIPath = fmt.Sprintf("/fba/inbound/v0/shipments/%s/items", params.PathParams["shipmentId"])
	params.RestoreRate = 500 * time.Millisecond

	return nil
}

func GetShipmentItems(params *SellingPartnerParams) error {
	params.Method = "GET"
	params.APIPath = "/fba/inbound/v0/shipmentItems"
	params.RestoreRate = 500 * time.Millisecond

	return nil
}
