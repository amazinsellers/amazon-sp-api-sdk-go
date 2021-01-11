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

/*
listCatalogItems:(req_params) => {
    return Object.assign(req_params, {
      method:'GET',
      api_path:'/catalog/v0/items',
      restore_rate:1
    });
  },
  getCatalogItem:(req_params) => {
  	utils.checkParams(req_params, {
  	  path:{
        asin:{
          type:'string'
        }
      }
	  });
    return Object.assign(req_params, {
      method:'GET',
      api_path:'/catalog/v0/items/' + req_params.path.asin,
      restore_rate:1
    });
  },
  listCatalogCategories:(req_params) => {
    return Object.assign(req_params, {
      method:'GET',
      api_path:'/catalog/v0/categories',
      restore_rate:1
    });
  }
 */
