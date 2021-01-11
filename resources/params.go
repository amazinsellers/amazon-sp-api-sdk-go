package resources

import "time"

type SellingPartnerParams struct {
	Operation string
	Method string
	APIPath string
	Body string
	Query map[string]interface{}
	RestoreRate time.Duration
	PathParams map[string]string
}
