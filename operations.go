package amazonspapi

import "github.com/amazinsellers/amazon-sp-api-sdk-go/resources"

type Operations map[string]func(params *resources.SellingPartnerParams) error

var AvailableOperations = Operations {
	// Orders
	"getOrders": resources.GetOrders,

	// Authorization
	"getAuthorizationCode": resources.GetAuthorizationCode,

	//Catalog
	"listCatalogItems": resources.ListCatalogItems,
	"getCatalogItem": resources.GetCatalogItem,
}
