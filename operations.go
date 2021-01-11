package amazonspapi

import "github.com/amazinsellers/amazon-sp-api-sdk-go/resources"

type Operations map[string]func(params *resources.SellingPartnerParams) error

var AvailableOperations = Operations{
	// Orders
	"getOrders": resources.GetOrders,

	// Authorization
	"getAuthorizationCode": resources.GetAuthorizationCode,

	// Catalog
	"listCatalogItems": resources.ListCatalogItems,
	"getCatalogItem":   resources.GetCatalogItem,

	// FBA Inbound eligibility
	"getItemEligibilityPreview": resources.GetItemEligibilityPreview,

	// FBA Inventory
	"getInventorySummaries": resources.GetInventorySummaries,

	// FBA Small & Light
	"getSmallAndLightEnrollmentBySellerSKU":    resources.GetSmallAndLightEnrollmentBySellerSKU,
	"putSmallAndLightEnrollmentBySellerSKU":    resources.PutSmallAndLightEnrollmentBySellerSKU,
	"deleteSmallAndLightEnrollmentBySellerSKU": resources.DeleteSmallAndLightEnrollmentBySellerSKU,
	"getSmallAndLightEligibilityBySellerSKU":   resources.GetSmallAndLightEligibilityBySellerSKU,
	"getSmallAndLightFeePreview":               resources.GetSmallAndLightFeePreview,

	// Feeds
	"getFeeds":           resources.GetFeeds,
	"createFeed":         resources.CreateFeed,
	"getFeed":            resources.GetFeed,
	"cancelFeed":         resources.CancelFeed,
	"createFeedDocument": resources.CreateFeedDocument,
	"getFeedDocument":    resources.GetFeedDocument,


	// Finance
	"listFinancialEventGroups":     resources.ListFinancialEventGroups,
	"listFinancialEventsByGroupId": resources.ListFinancialEventsByGroupId,
	"listFinancialEventsByOrderId": resources.ListFinancialEventsByOrderId,
	"listFinancialEvents":          resources.ListFinancialEvents,

	// Fulfilment Inbound
	"getInboundGuidance":           resources.GetInboundGuidance,
	"createInboundShipmentPlan":    resources.CreateInboundShipmentPlan,
	"updateInboundShipment":        resources.UpdateInboundShipment,
	"createInboundShipment":        resources.CreateInboundShipment,
	"getPreorderInfo":              resources.GetPreorderInfo,
	"confirmPreorder":              resources.ConfirmPreorder,
	"getPrepInstructions":          resources.GetPrepInstructions,
	"getTransportDetails":          resources.GetTransportDetails,
	"putTransportDetails":          resources.PutTransportDetails,
	"voidTransport":                resources.VoidTransport,
	"estimateTransport":            resources.EstimateTransport,
	"confirmTransport":             resources.ConfirmTransport,
	"getLabels":                    resources.GetLabels,
	"getBillOfLading":              resources.GetBillOfLading,
	"getShipments":                 resources.GetShipments,
	"getShipmentItemsByShipmentId": resources.GetShipmentItemsByShipmentId,
	"getShipmentItems":             resources.GetShipmentItems,

	// Fulfilment Outbound
	"getFulfillmentPreview":     resources.GetFulfillmentPreview,
	"listAllFulfillmentOrders":  resources.ListAllFulfillmentOrders,
	"createFulfillmentOrder":    resources.CreateFulfillmentOrder,
	"getPackageTrackingDetails": resources.GetPackageTrackingDetails,
	"listReturnReasonCodes":     resources.ListReturnReasonCodes,
	"createFulfillmentReturn":   resources.CreateFulfillmentReturn,
	"getFulfillmentOrder":       resources.GetFulfillmentOrder,
	"updateFulfillmentOrder":    resources.UpdateFulfillmentOrder,
	"cancelFulfillmentOrder":    resources.CancelFulfillmentOrder,
	"getFeatures":               resources.GetFeatures,
	"getFeatureInventory":       resources.GetFeatureInventory,
	"getFeatureSKU":             resources.GetFeatureSKU,

	//"": resources.,
	//"": resources.,
	//"": resources.,
	//"": resources.,
	//"": resources.,
	//"": resources.,
	//"": resources.,
	//"": resources.,
	//"": resources.,
	//"": resources.,
	//"": resources.,
	//"": resources.,
}
