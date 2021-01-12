package resources

import (
	"fmt"
	"time"
)

func GetPricing(params *SellingPartnerParams) error {

	params.Method = "GET"
	params.APIPath = "/products/pricing/v0/price"
	params.RestoreRate = 1 * time.Second
	return nil
}

func GetCompetitivePricing(params *SellingPartnerParams) error {

	params.Method = "GET"
	params.APIPath = "/products/pricing/v0/competitivePrice"
	params.RestoreRate = 1 * time.Second
	return nil
}

func GetListingOffers(params *SellingPartnerParams) error {
	if _, present := params.PathParams["SellerSKU"]; !present {
		return fmt.Errorf("path param 'SellerSKU' not present")
	}

	params.Method = "GET"
	params.APIPath = "/products/pricing/v0/listings/" + params.PathParams["SellerSKU"] + "/offers"
	params.RestoreRate = 1 * time.Second
	return nil
}

func GetItemOffers(params *SellingPartnerParams) error {
	if _, present := params.PathParams["Asin"]; !present {
		return fmt.Errorf("path param 'Asin' not present")
	}

	params.Method = "GET"
	params.APIPath = "/products/pricing/v0/items/" + params.PathParams["Asin"] + "/offers"
	params.RestoreRate = 1 * time.Second
	return nil
}
