package resources

import (
	"fmt"
	"time"
)

func GetSmallAndLightEnrollmentBySellerSKU(params *SellingPartnerParams) error {
	if _,present := params.PathParams["sellerSKU"]; !present {
		return fmt.Errorf("path param 'sellerSKU' not present")
	}

	params.Method = "GET"
	params.APIPath = "/fba/smallAndLight/v1/enrollments/" + params.PathParams["sellerSKU"]
	params.RestoreRate = 500 * time.Millisecond

	return nil
}

func PutSmallAndLightEnrollmentBySellerSKU(params *SellingPartnerParams) error {
	if _,present := params.PathParams["sellerSKU"]; !present {
		return fmt.Errorf("path param 'sellerSKU' not present")
	}

	params.Method = "PUT"
	params.APIPath = "/fba/smallAndLight/v1/enrollments/" + params.PathParams["sellerSKU"]
	params.RestoreRate = 500 * time.Millisecond

	return nil
}

func DeleteSmallAndLightEnrollmentBySellerSKU(params *SellingPartnerParams) error {
	if _,present := params.PathParams["sellerSKU"]; !present {
		return fmt.Errorf("path param 'sellerSKU' not present")
	}

	params.Method = "DELETE"
	params.APIPath = "/fba/smallAndLight/v1/enrollments/" + params.PathParams["sellerSKU"]
	params.RestoreRate = 500 * time.Millisecond

	return nil
}

func GetSmallAndLightEligibilityBySellerSKU(params *SellingPartnerParams) error {
	if _,present := params.PathParams["sellerSKU"]; !present {
		return fmt.Errorf("path param 'sellerSKU' not present")
	}

	params.Method = "GET"
	params.APIPath = "/fba/smallAndLight/v1/eligibilities/" + params.PathParams["sellerSKU"]
	params.RestoreRate = 500 * time.Millisecond

	return nil
}

func GetSmallAndLightFeePreview(params *SellingPartnerParams) error {
	params.Method = "POST"
	params.APIPath = "/fba/smallAndLight/v1/feePreviews"
	params.RestoreRate = 1 * time.Second

	return nil
}
