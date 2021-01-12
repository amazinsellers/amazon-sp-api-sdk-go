package resources

import (
	"fmt"
	"time"
)

func CreateUploadDestinationForResource(params *SellingPartnerParams) error {
	if _, present := params.PathParams["resource"]; !present {
		return fmt.Errorf("path param 'resource' not present")
	}

	params.Method = "POST"
	params.APIPath = "/uploads/2020-11-01/uploadDestinations/" + params.PathParams["resource"]
	params.RestoreRate = 10 * time.Second
	return nil
}
