// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package straddle

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/stainless-sdks/straddle-go/internal/apijson"
	"github.com/stainless-sdks/straddle-go/internal/apiquery"
	"github.com/stainless-sdks/straddle-go/internal/param"
	"github.com/stainless-sdks/straddle-go/internal/requestconfig"
	"github.com/stainless-sdks/straddle-go/option"
	"github.com/stainless-sdks/straddle-go/packages/pagination"
	"github.com/stainless-sdks/straddle-go/shared"
)

// FundingEventService contains methods and other services that help with
// interacting with the straddle API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFundingEventService] method instead.
type FundingEventService struct {
	Options []option.RequestOption
}

// NewFundingEventService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewFundingEventService(opts ...option.RequestOption) (r *FundingEventService) {
	r = &FundingEventService{}
	r.Options = opts
	return
}

// Retrieves a list of funding events for your account. This endpoint supports
// advanced sorting and filtering options.
func (r *FundingEventService) List(ctx context.Context, params FundingEventListParams, opts ...option.RequestOption) (res *pagination.PageNumberSchema[shared.FundingEventSummaryV1], err error) {
	var raw *http.Response
	if params.CorrelationID.Present {
		opts = append(opts, option.WithHeader("Correlation-Id", fmt.Sprintf("%s", params.CorrelationID)))
	}
	if params.RequestID.Present {
		opts = append(opts, option.WithHeader("Request-Id", fmt.Sprintf("%s", params.RequestID)))
	}
	if params.StraddleAccountID.Present {
		opts = append(opts, option.WithHeader("Straddle-Account-Id", fmt.Sprintf("%s", params.StraddleAccountID)))
	}
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/funding_events"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, params, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Retrieves a list of funding events for your account. This endpoint supports
// advanced sorting and filtering options.
func (r *FundingEventService) ListAutoPaging(ctx context.Context, params FundingEventListParams, opts ...option.RequestOption) *pagination.PageNumberSchemaAutoPager[shared.FundingEventSummaryV1] {
	return pagination.NewPageNumberSchemaAutoPager(r.List(ctx, params, opts...))
}

// Retrieves the details of an existing funding event. Supply the unique funding
// event `id`, and Straddle will return the individual transaction items that make
// up the funding event.
func (r *FundingEventService) Get(ctx context.Context, id string, query FundingEventGetParams, opts ...option.RequestOption) (res *FundingEventSummaryItemV1, err error) {
	if query.CorrelationID.Present {
		opts = append(opts, option.WithHeader("Correlation-Id", fmt.Sprintf("%s", query.CorrelationID)))
	}
	if query.RequestID.Present {
		opts = append(opts, option.WithHeader("Request-Id", fmt.Sprintf("%s", query.RequestID)))
	}
	if query.StraddleAccountID.Present {
		opts = append(opts, option.WithHeader("Straddle-Account-Id", fmt.Sprintf("%s", query.StraddleAccountID)))
	}
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/funding_events/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type FundingEventSummaryItemV1 struct {
	Data shared.FundingEventSummaryV1 `json:"data,required"`
	// Metadata about the API request, including an identifier and timestamp.
	Meta shared.ResponseMetadata `json:"meta,required"`
	// Indicates the structure of the returned content.
	//
	//   - "object" means the `data` field contains a single JSON object.
	//   - "array" means the `data` field contains an array of objects.
	//   - "error" means the `data` field contains an error object with details of the
	//     issue.
	//   - "none" means no data is returned.
	ResponseType shared.ResponseTypeEnum       `json:"response_type,required"`
	JSON         fundingEventSummaryItemV1JSON `json:"-"`
}

// fundingEventSummaryItemV1JSON contains the JSON metadata for the struct
// [FundingEventSummaryItemV1]
type fundingEventSummaryItemV1JSON struct {
	Data         apijson.Field
	Meta         apijson.Field
	ResponseType apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *FundingEventSummaryItemV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fundingEventSummaryItemV1JSON) RawJSON() string {
	return r.raw
}

type FundingEventSummaryPagedV1 struct {
	Data []shared.FundingEventSummaryV1 `json:"data,required"`
	Meta shared.PagedResponseMetadata2  `json:"meta,required"`
	// Indicates the structure of the returned content.
	//
	//   - "object" means the `data` field contains a single JSON object.
	//   - "array" means the `data` field contains an array of objects.
	//   - "error" means the `data` field contains an error object with details of the
	//     issue.
	//   - "none" means no data is returned.
	ResponseType shared.ResponseTypeEnum        `json:"response_type,required"`
	JSON         fundingEventSummaryPagedV1JSON `json:"-"`
}

// fundingEventSummaryPagedV1JSON contains the JSON metadata for the struct
// [FundingEventSummaryPagedV1]
type fundingEventSummaryPagedV1JSON struct {
	Data         apijson.Field
	Meta         apijson.Field
	ResponseType apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *FundingEventSummaryPagedV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fundingEventSummaryPagedV1JSON) RawJSON() string {
	return r.raw
}

type FundingEventListParams struct {
	// The start date of the range to filter by using the `YYYY-MM-DD` format.
	CreatedFrom param.Field[time.Time] `query:"created_from" format:"date"`
	// The end date of the range to filter by using the `YYYY-MM-DD` format.
	CreatedTo param.Field[time.Time] `query:"created_to" format:"date"`
	// Describes the direction of the funding event from the perspective of the
	// `linked_bank_account`.
	Direction param.Field[shared.TransferDirectionV1] `query:"direction"`
	// The funding event types describes the direction and reason for the funding
	// event.
	EventType param.Field[shared.FundingEventTypeV1] `query:"event_type"`
	// Results page number. Starts at page 1.
	PageNumber param.Field[int64] `query:"page_number"`
	// Results page size. Max value: 1000
	PageSize param.Field[int64] `query:"page_size"`
	// The field to sort the results by.
	SortBy param.Field[FundingEventListParamsSortBy] `query:"sort_by"`
	// The order in which to sort the results.
	SortOrder param.Field[shared.SortOrder] `query:"sort_order"`
	// Trace number.
	TraceNumber       param.Field[string] `query:"trace_number"`
	CorrelationID     param.Field[string] `header:"Correlation-Id"`
	RequestID         param.Field[string] `header:"Request-Id"`
	StraddleAccountID param.Field[string] `header:"Straddle-Account-Id" format:"uuid"`
}

// URLQuery serializes [FundingEventListParams]'s query parameters as `url.Values`.
func (r FundingEventListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The field to sort the results by.
type FundingEventListParamsSortBy string

const (
	FundingEventListParamsSortByTransferDate FundingEventListParamsSortBy = "transfer_date"
	FundingEventListParamsSortByID           FundingEventListParamsSortBy = "id"
	FundingEventListParamsSortByAmount       FundingEventListParamsSortBy = "amount"
)

func (r FundingEventListParamsSortBy) IsKnown() bool {
	switch r {
	case FundingEventListParamsSortByTransferDate, FundingEventListParamsSortByID, FundingEventListParamsSortByAmount:
		return true
	}
	return false
}

type FundingEventGetParams struct {
	CorrelationID     param.Field[string] `header:"Correlation-Id"`
	RequestID         param.Field[string] `header:"Request-Id"`
	StraddleAccountID param.Field[string] `header:"Straddle-Account-Id" format:"uuid"`
}
