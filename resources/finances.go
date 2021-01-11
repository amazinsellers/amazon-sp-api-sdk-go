package resources

import (
	"fmt"
	"time"
)

func ListFinancialEventGroups(params *SellingPartnerParams) error {
	params.Method = "GET"
	params.APIPath = "/finances/v0/financialEventGroups"
	params.RestoreRate = 2 * time.Second

	return nil
}

func ListFinancialEventsByGroupId(params *SellingPartnerParams) error {
	if _,present := params.PathParams["eventGroupId"]; !present {
		return fmt.Errorf("path param 'eventGroupId' not present")
	}

	params.Method = "GET"
	params.APIPath = fmt.Sprintf("/finances/v0/financialEventGroups/%s/financialEvents", params.PathParams["eventGroupId"])
	params.RestoreRate = 2 * time.Second

	return nil
}

func ListFinancialEventsByOrderId(params *SellingPartnerParams) error {
	if _,present := params.PathParams["orderId"]; !present {
		return fmt.Errorf("path param 'orderId' not present")
	}

	params.Method = "GET"
	params.APIPath = fmt.Sprintf("/finances/v0/orders/%s/financialEvents", params.PathParams["orderId"])
	params.RestoreRate = 2 * time.Second

	return nil
}

func ListFinancialEvents(params *SellingPartnerParams) error {
	params.Method = "GET"
	params.APIPath = "/finances/v0/financialEvents"
	params.RestoreRate = 2 * time.Second

	return nil
}
