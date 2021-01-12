package resources

import (
	"fmt"
	"time"
)

func GetServiceJobByServiceJobId(params *SellingPartnerParams) error {
	if _, present := params.PathParams["serviceJobId"]; !present {
		return fmt.Errorf("path param 'serviceJobId' not present")
	}

	params.Method = "GET"
	params.APIPath = "/service/v1/serviceJobs/" + params.PathParams["serviceJobId"]
	params.RestoreRate = 50 * time.Millisecond
	return nil
}

func CancelServiceJobByServiceJobId(params *SellingPartnerParams) error {
	if _, present := params.PathParams["serviceJobId"]; !present {
		return fmt.Errorf("path param 'serviceJobId' not present")
	}

	params.Method = "PUT"
	params.APIPath = "/service/v1/serviceJobs/" + params.PathParams["serviceJobId"] + "/cancellations"
	params.RestoreRate = 200 * time.Millisecond
	return nil
}

func CompleteServiceJobByServiceJobId(params *SellingPartnerParams) error {
	if _, present := params.PathParams["serviceJobId"]; !present {
		return fmt.Errorf("path param 'serviceJobId' not present")
	}

	params.Method = "PUT"
	params.APIPath = "/service/v1/serviceJobs/" + params.PathParams["serviceJobId"] + "/completions"
	params.RestoreRate = 200 * time.Millisecond
	return nil
}

func GetServiceJobs(params *SellingPartnerParams) error {

	params.Method = "GET"
	params.APIPath = "/service/v1/serviceJobs"
	params.RestoreRate = 100 * time.Millisecond
	return nil
}

func AddAppointmentForServiceJobByServiceJobId(params *SellingPartnerParams) error {
	if _, present := params.PathParams["serviceJobId"]; !present {
		return fmt.Errorf("path param 'serviceJobId' not present")
	}

	params.Method = "POST"
	params.APIPath = "/service/v1/serviceJobs/" + params.PathParams["serviceJobId"] + "/appointments"
	params.RestoreRate = 200 * time.Millisecond
	return nil
}

func RescheduleAppointmentForServiceJobByServiceJobId(params *SellingPartnerParams) error {
	if _, present := params.PathParams["serviceJobId"]; !present {
		return fmt.Errorf("path param 'serviceJobId' not present")
	}

	if _, present := params.PathParams["appointmentId"]; !present {
		return fmt.Errorf("path param 'appointmentId' not present")
	}

	params.Method = "POST"
	params.APIPath = "/service/v1/serviceJobs/" + params.PathParams["serviceJobId"] + "/appointments/" + params.PathParams["appointmentId"]
	params.RestoreRate = 200 * time.Millisecond
	return nil
}
