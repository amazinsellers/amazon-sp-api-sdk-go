package resources

import (
	"fmt"
	"time"
)

func GetMyFeesEstimateForSKU(params *SellingPartnerParams) error {
	if _, present := params.PathParams["SellerSKU"]; !present {
		return fmt.Errorf("path param 'SellerSKU' not present")
	}

	params.Method = "POST"
	params.APIPath = "/products/fees/v0/listings/" + params.PathParams["SellerSKU"] + "/feesEstimate"
	params.RestoreRate = 1 * time.Second
	return nil
}

func GetMyFeesEstimateForASIN(params *SellingPartnerParams) error {
	if _, present := params.PathParams["Asin"]; !present {
		return fmt.Errorf("path param 'Asin' not present")
	}

	params.Method = "POST"
	params.APIPath = "/products/fees/v0/items/" + params.PathParams["Asin"] + "/feesEstimate"
	params.RestoreRate = 1 * time.Second
	return nil
}
