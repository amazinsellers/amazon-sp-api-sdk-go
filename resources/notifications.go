package resources

import (
	"fmt"
	"time"
)

func GetSubscription(params *SellingPartnerParams) error {
	if _, present := params.PathParams["notificationType"]; !present {
		return fmt.Errorf("path param 'notificationType' not present")
	}

	params.Method = "GET"
	params.APIPath = "/notifications/v1/subscriptions/" + params.PathParams["notificationType"]
	params.RestoreRate = 1 * time.Second
	return nil
}

func CreateSubscription(params *SellingPartnerParams) error {
	if _, present := params.PathParams["notificationType"]; !present {
		return fmt.Errorf("path param 'notificationType' not present")
	}

	params.Method = "POST"
	params.APIPath = "/notifications/v1/subscriptions/" + params.PathParams["notificationType"]
	params.RestoreRate = 1 * time.Second
	return nil
}

func GetSubscriptionById(params *SellingPartnerParams) error {
	if _, present := params.PathParams["notificationType"]; !present {
		return fmt.Errorf("path param 'notificationType' not present")
	}

	if _, present := params.PathParams["subscriptionId"]; !present {
		return fmt.Errorf("path param 'subscriptionId' not present")
	}

	params.Method = "GET"
	params.APIPath = "/notifications/v1/subscriptions/" + params.PathParams["notificationType"] + "/" + params.PathParams["subscriptionId"]
	params.RestoreRate = 1 * time.Second
	return nil
}

func DeleteSubscriptionById(params *SellingPartnerParams) error {
	if _, present := params.PathParams["notificationType"]; !present {
		return fmt.Errorf("path param 'notificationType' not present")
	}
	if _, present := params.PathParams["subscriptionId"]; !present {
		return fmt.Errorf("path param 'subscriptionId' not present")
	}

	params.Method = "DELETE"
	params.APIPath = "/notifications/v1/subscriptions/" + params.PathParams["notificationType"] + "/" + params.PathParams["subscriptionId"]
	params.RestoreRate = 1 * time.Second
	return nil
}

func GetDestinations(params *SellingPartnerParams) error {
	params.Method = "GET"
	params.APIPath = "/notifications/v1/destinations"
	params.RestoreRate = 1 * time.Second
	return nil
}
func CreateDestination(params *SellingPartnerParams) error {
	params.Method = "POST"
	params.APIPath = "/notifications/v1/destinations"
	params.RestoreRate = 1 * time.Second
	return nil
}

func GetDestination(params *SellingPartnerParams) error {
	if _, present := params.PathParams["destinationId"]; !present {
		return fmt.Errorf("path param 'destinationId' not present")
	}

	params.Method = "GET"
	params.APIPath = "/notifications/v1/destinations/" + params.PathParams["destinationId"]
	params.RestoreRate = 1 * time.Second
	return nil
}

func DeleteDestination(params *SellingPartnerParams) error {
	if _, present := params.PathParams["destinationId"]; !present {
		return fmt.Errorf("path param 'destinationId' not present")
	}

	params.Method = "DELETE"
	params.APIPath = "/notifications/v1/destinations/" + params.PathParams["destinationId"]
	params.RestoreRate = 1 * time.Second
	return nil
}
