package resources

import (
	"fmt"
	"time"
)

func GetSolicitationActionsForOrder(params *SellingPartnerParams) error {
	if _, present := params.PathParams["amazonOrderId"]; !present {
		return fmt.Errorf("path param 'amazonOrderId' not present")
	}

	params.Method = "GET"
	params.APIPath = "/solicitations/v1/orders/" + params.PathParams["amazonOrderId"]
	params.RestoreRate = 1 * time.Second
	return nil
}

func CreateProductReviewAndSellerFeedbackSolicitation(params *SellingPartnerParams) error {
	if _, present := params.PathParams["amazonOrderId"]; !present {
		return fmt.Errorf("path param 'amazonOrderId' not present")
	}

	params.Method = "POST"
	params.APIPath = "/solicitations/v1/orders/" + params.PathParams["amazonOrderId"] + "/solicitations/productReviewAndSellerFeedback"
	params.RestoreRate = 60 * time.Second
	return nil
}
