package resources

import (
	"fmt"
	"time"
)

func GetMessagingActionsForOrder(params *SellingPartnerParams) error {
	if _, present := params.PathParams["amazonOrderId"]; !present {
		return fmt.Errorf("path param 'amazonOrderId' not present")
	}

	params.Method = "GET"
	params.APIPath = "/messaging/v1/orders/" + params.PathParams["amazonOrderId"]
	params.RestoreRate = 1 * time.Second
	return nil
}

func ConfirmCustomizationDetails(params *SellingPartnerParams) error {
	if _, present := params.PathParams["amazonOrderId"]; !present {
		return fmt.Errorf("path param 'amazonOrderId' not present")
	}

	params.Method = "POST"
	params.APIPath = "/messaging/v1/orders/" + params.PathParams["amazonOrderId"] + "/messages/confirmCustomizationDetails"
	params.RestoreRate = 1 * time.Second
	return nil
}

func CreateConfirmDeliveryDetails(params *SellingPartnerParams) error {
	if _, present := params.PathParams["amazonOrderId"]; !present {
		return fmt.Errorf("path param 'amazonOrderId' not present")
	}

	params.Method = "POST"
	params.APIPath = "/messaging/v1/orders/" + params.PathParams["amazonOrderId"] + "/messages/confirmDeliveryDetails"
	params.RestoreRate = 1 * time.Second
	return nil
}

func CreateLegalDisclosure(params *SellingPartnerParams) error {
	if _, present := params.PathParams["amazonOrderId"]; !present {
		return fmt.Errorf("path param 'amazonOrderId' not present")
	}

	params.Method = "POST"
	params.APIPath = "/messaging/v1/orders/" + params.PathParams["amazonOrderId"] + "/messages/legalDisclosure"
	params.RestoreRate = 1 * time.Second
	return nil
}

func CreateNegativeFeedbackRemoval(params *SellingPartnerParams) error {
	if _, present := params.PathParams["amazonOrderId"]; !present {
		return fmt.Errorf("path param 'amazonOrderId' not present")
	}

	params.Method = "POST"
	params.APIPath = "/messaging/v1/orders/" + params.PathParams["amazonOrderId"] + "/messages/negativeFeedbackRemoval"
	params.RestoreRate = 1 * time.Second
	return nil
}

func CreateConfirmOrderDetails(params *SellingPartnerParams) error {
	if _, present := params.PathParams["amazonOrderId"]; !present {
		return fmt.Errorf("path param 'amazonOrderId' not present")
	}

	params.Method = "POST"
	params.APIPath = "/messaging/v1/orders/" + params.PathParams["amazonOrderId"] + "/messages/confirmOrderDetails"
	params.RestoreRate = 1 * time.Second
	return nil
}

func CreateConfirmServiceDetails(params *SellingPartnerParams) error {
	if _, present := params.PathParams["amazonOrderId"]; !present {
		return fmt.Errorf("path param 'amazonOrderId' not present")
	}

	params.Method = "POST"
	params.APIPath = "/messaging/v1/orders/" + params.PathParams["amazonOrderId"] + "/messages/confirmServiceDetails"
	params.RestoreRate = 1 * time.Second
	return nil
}

func CreateAmazonMotors(params *SellingPartnerParams) error {
	if _, present := params.PathParams["amazonOrderId"]; !present {
		return fmt.Errorf("path param 'amazonOrderId' not present")
	}

	params.Method = "POST"
	params.APIPath = "/messaging/v1/orders/" + params.PathParams["amazonOrderId"] + "/messages/amazonMotors"
	params.RestoreRate = 1 * time.Second
	return nil
}

func CreateWarranty(params *SellingPartnerParams) error {
	if _, present := params.PathParams["amazonOrderId"]; !present {
		return fmt.Errorf("path param 'amazonOrderId' not present")
	}

	params.Method = "POST"
	params.APIPath = "/messaging/v1/orders/" + params.PathParams["amazonOrderId"] + "/messages/warranty"
	params.RestoreRate = 1 * time.Second
	return nil
}

func GetAttributes(params *SellingPartnerParams) error {
	if _, present := params.PathParams["amazonOrderId"]; !present {
		return fmt.Errorf("path param 'amazonOrderId' not present")
	}

	params.Method = "GET"
	params.APIPath = "/messaging/v1/orders/" + params.PathParams["amazonOrderId"] + "/attributes"
	params.RestoreRate = 1 * time.Second
	return nil
}

func CreateDigitalAccessKey(params *SellingPartnerParams) error {
	if _, present := params.PathParams["amazonOrderId"]; !present {
		return fmt.Errorf("path param 'amazonOrderId' not present")
	}

	params.Method = "POST"
	params.APIPath = "/messaging/v1/orders/" + params.PathParams["amazonOrderId"] + "/messages/digitalAccessKey"
	params.RestoreRate = 1 * time.Second
	return nil
}

func CreateUnexpectedProblem(params *SellingPartnerParams) error {
	if _, present := params.PathParams["amazonOrderId"]; !present {
		return fmt.Errorf("path param 'amazonOrderId' not present")
	}

	params.Method = "POST"
	params.APIPath = "/messaging/v1/orders/" + params.PathParams["amazonOrderId"] + "/messages/unexpectedProblem"
	params.RestoreRate = 1 * time.Second
	return nil
}
