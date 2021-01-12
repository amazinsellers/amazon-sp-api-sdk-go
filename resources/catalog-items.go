package resources

import (
	"fmt"
	"time"
)

func ListCatalogItems(params *SellingPartnerParams) error {
	params.Method = "GET"
	params.APIPath = "/catalog/v0/items"
	params.RestoreRate = 1 * time.Second

	return nil
}

func GetCatalogItem(params *SellingPartnerParams) error {
	if _,present := params.PathParams["asin"]; !present {
		return fmt.Errorf("path param 'asin' not present")
	}
	params.Method = "GET"
	params.APIPath = "/catalog/v0/items/" + params.PathParams["asin"]
	params.RestoreRate = 1 * time.Second

	return nil
}

func ListCatalogCategories(params *SellingPartnerParams) error {
	params.Method = "GET"
	params.APIPath = "/catalog/v0/categories"
	params.RestoreRate = 1 * time.Second
	return nil
}
