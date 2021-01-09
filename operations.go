package amazonspapi

import "github.com/amazinsellers/amazon-sp-api-sdk-go/resources"

type Operations map[string]func(params *resources.SellingPartnerParams) error

var AvailableOperations = Operations {
	"getOrders": resources.GetOrders,
}
