// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package straddle

import (
	"context"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/straddle-go/internal/apijson"
	"github.com/stainless-sdks/straddle-go/internal/param"
	"github.com/stainless-sdks/straddle-go/internal/requestconfig"
	"github.com/stainless-sdks/straddle-go/option"
	"github.com/stainless-sdks/straddle-go/shared"
)

// ReportService contains methods and other services that help with interacting
// with the straddle API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewReportService] method instead.
type ReportService struct {
	Options []option.RequestOption
}

// NewReportService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewReportService(opts ...option.RequestOption) (r *ReportService) {
	r = &ReportService{}
	r.Options = opts
	return
}

func (r *ReportService) NewTotalCustomersByStatus(ctx context.Context, body ReportNewTotalCustomersByStatusParams, opts ...option.RequestOption) (res *ReportNewTotalCustomersByStatusResponse, err error) {
	if body.CorrelationID.Present {
		opts = append(opts, option.WithHeader("Correlation-Id", fmt.Sprintf("%s", body.CorrelationID)))
	}
	if body.RequestID.Present {
		opts = append(opts, option.WithHeader("Request-Id", fmt.Sprintf("%s", body.RequestID)))
	}
	if body.StraddleAccountID.Present {
		opts = append(opts, option.WithHeader("Straddle-Account-Id", fmt.Sprintf("%s", body.StraddleAccountID)))
	}
	opts = append(r.Options[:], opts...)
	path := "v1/reports/total_customers_by_status"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

type ReportNewTotalCustomersByStatusResponse struct {
	Data ReportNewTotalCustomersByStatusResponseData `json:"data,required"`
	// Metadata about the API request, including an identifier and timestamp.
	Meta shared.ResponseMetadata `json:"meta,required"`
	// Indicates the structure of the returned content.
	//
	//   - "object" means the `data` field contains a single JSON object.
	//   - "array" means the `data` field contains an array of objects.
	//   - "error" means the `data` field contains an error object with details of the
	//     issue.
	//   - "none" means no data is returned.
	ResponseType ReportNewTotalCustomersByStatusResponseResponseType `json:"response_type,required"`
	JSON         reportNewTotalCustomersByStatusResponseJSON         `json:"-"`
}

// reportNewTotalCustomersByStatusResponseJSON contains the JSON metadata for the
// struct [ReportNewTotalCustomersByStatusResponse]
type reportNewTotalCustomersByStatusResponseJSON struct {
	Data         apijson.Field
	Meta         apijson.Field
	ResponseType apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ReportNewTotalCustomersByStatusResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r reportNewTotalCustomersByStatusResponseJSON) RawJSON() string {
	return r.raw
}

type ReportNewTotalCustomersByStatusResponseData struct {
	Inactive int64                                           `json:"inactive,required"`
	Pending  int64                                           `json:"pending,required"`
	Rejected int64                                           `json:"rejected,required"`
	Review   int64                                           `json:"review,required"`
	Verified int64                                           `json:"verified,required"`
	JSON     reportNewTotalCustomersByStatusResponseDataJSON `json:"-"`
}

// reportNewTotalCustomersByStatusResponseDataJSON contains the JSON metadata for
// the struct [ReportNewTotalCustomersByStatusResponseData]
type reportNewTotalCustomersByStatusResponseDataJSON struct {
	Inactive    apijson.Field
	Pending     apijson.Field
	Rejected    apijson.Field
	Review      apijson.Field
	Verified    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ReportNewTotalCustomersByStatusResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r reportNewTotalCustomersByStatusResponseDataJSON) RawJSON() string {
	return r.raw
}

// Indicates the structure of the returned content.
//
//   - "object" means the `data` field contains a single JSON object.
//   - "array" means the `data` field contains an array of objects.
//   - "error" means the `data` field contains an error object with details of the
//     issue.
//   - "none" means no data is returned.
type ReportNewTotalCustomersByStatusResponseResponseType string

const (
	ReportNewTotalCustomersByStatusResponseResponseTypeObject ReportNewTotalCustomersByStatusResponseResponseType = "object"
	ReportNewTotalCustomersByStatusResponseResponseTypeArray  ReportNewTotalCustomersByStatusResponseResponseType = "array"
	ReportNewTotalCustomersByStatusResponseResponseTypeError  ReportNewTotalCustomersByStatusResponseResponseType = "error"
	ReportNewTotalCustomersByStatusResponseResponseTypeNone   ReportNewTotalCustomersByStatusResponseResponseType = "none"
)

func (r ReportNewTotalCustomersByStatusResponseResponseType) IsKnown() bool {
	switch r {
	case ReportNewTotalCustomersByStatusResponseResponseTypeObject, ReportNewTotalCustomersByStatusResponseResponseTypeArray, ReportNewTotalCustomersByStatusResponseResponseTypeError, ReportNewTotalCustomersByStatusResponseResponseTypeNone:
		return true
	}
	return false
}

type ReportNewTotalCustomersByStatusParams struct {
	CorrelationID     param.Field[string] `header:"Correlation-Id"`
	RequestID         param.Field[string] `header:"Request-Id"`
	StraddleAccountID param.Field[string] `header:"Straddle-Account-Id" format:"uuid"`
}
