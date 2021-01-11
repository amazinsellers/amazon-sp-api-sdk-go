package resources

import (
	"fmt"
	"time"
)

func GetFeeds(params *SellingPartnerParams) error {
	params.Method = "GET"
	params.APIPath = "/feeds/2020-09-04/feeds"
	params.RestoreRate = 45 * time.Second

	return nil
}

func CreateFeed(params *SellingPartnerParams) error {
	params.Method = "POST"
	params.APIPath = "/feeds/2020-09-04/feeds"
	params.RestoreRate = 120 * time.Second

	return nil
}

func GetFeed(params *SellingPartnerParams) error {
	if _,present := params.PathParams["feedId"]; !present {
		return fmt.Errorf("path param 'feedId' not present")
	}

	params.Method = "GET"
	params.APIPath = "/feeds/2020-09-04/feeds/" + params.PathParams["feedId"]
	params.RestoreRate = 500 * time.Millisecond

	return nil
}

func CancelFeed(params *SellingPartnerParams) error {
	if _,present := params.PathParams["feedId"]; !present {
		return fmt.Errorf("path param 'feedId' not present")
	}

	params.Method = "DELETE"
	params.APIPath = "/feeds/2020-09-04/feeds/" + params.PathParams["feedId"]
	params.RestoreRate = 45 * time.Second

	return nil
}

func CreateFeedDocument(params *SellingPartnerParams) error {
	params.Method = "POST"
	params.APIPath = "/feeds/2020-09-04/documents"
	params.RestoreRate = 120 * time.Second

	return nil
}

func GetFeedDocument(params *SellingPartnerParams) error {
	if _,present := params.PathParams["feedDocumentId"]; !present {
		return fmt.Errorf("path param 'feedDocumentId' not present")
	}

	params.Method = "GET"
	params.APIPath = "/feeds/2020-09-04/documents/" + params.PathParams["feedDocumentId"]
	params.RestoreRate = 45 * time.Second

	return nil
}
