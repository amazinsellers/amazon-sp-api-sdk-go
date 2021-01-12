package resources

import (
	"net/url"
	"time"
)

type SellingPartnerParams struct {
	Operation   string
	Method      string
	APIPath     string
	Body        string
	Query       url.Values
	RestoreRate time.Duration
	PathParams  map[string]string
}
