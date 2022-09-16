package resources

import (
	"fmt"
	"time"
)

func GetReports(params *SellingPartnerParams) error {
	params.Method = "GET"
	params.APIPath = "/reports/2021-06-30/reports"
	params.RestoreRate = 45 * time.Second
	return nil
}

func CreateReport(params *SellingPartnerParams) error {
	params.Method = "POST"
	params.APIPath = "/reports/2021-06-30/reports"
	params.RestoreRate = 60 * time.Second
	return nil
}

func GetReport(params *SellingPartnerParams) error {
	if _, present := params.PathParams["reportId"]; !present {
		return fmt.Errorf("path param 'reportId' not present")
	}

	params.Method = "GET"
	params.APIPath = "/reports/2021-06-30/reports/" + params.PathParams["reportId"]
	params.RestoreRate = 500 * time.Millisecond
	return nil
}

func CancelReport(params *SellingPartnerParams) error {
	if _, present := params.PathParams["reportId"]; !present {
		return fmt.Errorf("path param 'reportId' not present")
	}

	params.Method = "DELETE"
	params.APIPath = "/reports/2021-06-30/reports/" + params.PathParams["reportId"]
	params.RestoreRate = 45 * time.Second
	return nil
}

func GetReportSchedules(params *SellingPartnerParams) error {
	params.Method = "GET"
	params.APIPath = "/reports/2021-06-30/schedules"
	params.RestoreRate = 500 * time.Millisecond
	return nil
}

func CreateReportSchedule(params *SellingPartnerParams) error {
	params.Method = "POST"
	params.APIPath = "/reports/2021-06-30/schedules"
	params.RestoreRate = 45 * time.Second
	return nil
}

func GetReportSchedule(params *SellingPartnerParams) error {
	if _, present := params.PathParams["reportScheduleId"]; !present {
		return fmt.Errorf("path param 'reportScheduleId' not present")
	}

	params.Method = "GET"
	params.APIPath = "/reports/2021-06-30/schedules/" + params.PathParams["reportScheduleId"]
	params.RestoreRate = 45 * time.Second
	return nil
}

func CancelReportSchedule(params *SellingPartnerParams) error {
	if _, present := params.PathParams["reportScheduleId"]; !present {
		return fmt.Errorf("path param 'reportScheduleId' not present")
	}

	params.Method = "DELETE"
	params.APIPath = "/reports/2021-06-30/schedules/" + params.PathParams["reportScheduleId"]
	params.RestoreRate = 45 * time.Second
	return nil
}

func GetReportDocument(params *SellingPartnerParams) error {
	if _, present := params.PathParams["reportDocumentId"]; !present {
		return fmt.Errorf("path param 'reportDocumentId' not present")
	}

	params.Method = "GET"
	params.APIPath = "/reports/2021-06-30/documents/" + params.PathParams["reportDocumentId"]
	params.RestoreRate = 45 * time.Second
	return nil
}
