// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package straddle

import (
	"context"
	"fmt"
	"net/http"
	"slices"

	"github.com/straddleio/straddle-go/internal/apijson"
	"github.com/straddleio/straddle-go/internal/requestconfig"
	"github.com/straddleio/straddle-go/option"
	"github.com/straddleio/straddle-go/packages/param"
	"github.com/straddleio/straddle-go/packages/respjson"
	"github.com/straddleio/straddle-go/shared"
)

// ReportService contains methods and other services that help with interacting
// with the straddle API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewReportService] method instead.
type ReportService struct {
	options []option.RequestOption
}

// NewReportService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewReportService(opts ...option.RequestOption) (r ReportService) {
	r = ReportService{}
	r.options = opts
	return
}

func (r *ReportService) NewTotalCustomersByStatus(ctx context.Context, body ReportNewTotalCustomersByStatusParams, opts ...option.RequestOption) (res *ReportNewTotalCustomersByStatusResponse, err error) {
	if !param.IsOmitted(body.CorrelationID) {
		opts = append(opts, option.WithHeader("Correlation-Id", fmt.Sprintf("%v", body.CorrelationID.Value)))
	}
	if !param.IsOmitted(body.RequestID) {
		opts = append(opts, option.WithHeader("Request-Id", fmt.Sprintf("%v", body.RequestID.Value)))
	}
	if !param.IsOmitted(body.StraddleAccountID) {
		opts = append(opts, option.WithHeader("Straddle-Account-Id", fmt.Sprintf("%v", body.StraddleAccountID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	path := "v1/reports/total_customers_by_status"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

type ReportNewTotalCustomersByStatusResponse struct {
	Data ReportNewTotalCustomersByStatusResponseData `json:"data" api:"required"`
	// Metadata about the API request, including an identifier and timestamp.
	Meta shared.ResponseMetadata `json:"meta" api:"required"`
	// Indicates the structure of the returned content.
	//
	//   - "object" means the `data` field contains a single JSON object.
	//   - "array" means the `data` field contains an array of objects.
	//   - "error" means the `data` field contains an error object with details of the
	//     issue.
	//   - "none" means no data is returned.
	//
	// Any of "object", "array", "error", "none".
	ResponseType ReportNewTotalCustomersByStatusResponseResponseType `json:"response_type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data         respjson.Field
		Meta         respjson.Field
		ResponseType respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ReportNewTotalCustomersByStatusResponse) RawJSON() string { return r.JSON.raw }
func (r *ReportNewTotalCustomersByStatusResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ReportNewTotalCustomersByStatusResponseData struct {
	Inactive int64 `json:"inactive" api:"required"`
	Pending  int64 `json:"pending" api:"required"`
	Rejected int64 `json:"rejected" api:"required"`
	Review   int64 `json:"review" api:"required"`
	Verified int64 `json:"verified" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Inactive    respjson.Field
		Pending     respjson.Field
		Rejected    respjson.Field
		Review      respjson.Field
		Verified    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ReportNewTotalCustomersByStatusResponseData) RawJSON() string { return r.JSON.raw }
func (r *ReportNewTotalCustomersByStatusResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
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

type ReportNewTotalCustomersByStatusParams struct {
	CorrelationID     param.Opt[string] `header:"Correlation-Id,omitzero" json:"-"`
	RequestID         param.Opt[string] `header:"Request-Id,omitzero" json:"-"`
	StraddleAccountID param.Opt[string] `header:"Straddle-Account-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}
