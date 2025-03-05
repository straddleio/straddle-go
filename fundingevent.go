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
func (r *FundingEventService) List(ctx context.Context, params FundingEventListParams, opts ...option.RequestOption) (res *pagination.PageNumberSchema[FundingEventSummaryPagedV1Data], err error) {
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
func (r *FundingEventService) ListAutoPaging(ctx context.Context, params FundingEventListParams, opts ...option.RequestOption) *pagination.PageNumberSchemaAutoPager[FundingEventSummaryPagedV1Data] {
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
	Data FundingEventSummaryItemV1Data `json:"data,required"`
	// Metadata about the API request, including an identifier and timestamp.
	Meta shared.ResponseMetadata `json:"meta,required"`
	// Indicates the structure of the returned content.
	//
	//   - "object" means the `data` field contains a single JSON object.
	//   - "array" means the `data` field contains an array of objects.
	//   - "error" means the `data` field contains an error object with details of the
	//     issue.
	//   - "none" means no data is returned.
	ResponseType FundingEventSummaryItemV1ResponseType `json:"response_type,required"`
	JSON         fundingEventSummaryItemV1JSON         `json:"-"`
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

type FundingEventSummaryItemV1Data struct {
	// Unique identifier for the funding event.
	ID string `json:"id,required" format:"uuid"`
	// The amount of the funding event in cents.
	Amount int64 `json:"amount,required"`
	// Describes the direction of the funding event from the perspective of the
	// `linked_bank_account`.
	Direction FundingEventSummaryItemV1DataDirection `json:"direction,required"`
	// The funding event types describes the direction and reason for the funding
	// event.
	EventType FundingEventSummaryItemV1DataEventType `json:"event_type,required"`
	// The number of payments associated with the funding event.
	PaymentCount int64 `json:"payment_count,required"`
	// Trace number.
	TraceNumbers []string `json:"trace_numbers,required"`
	// The date on which the funding event occurred. For `deposits` and `returns`, this
	// is the date the funds were credited to your bank account. For `withdrawals` and
	// `reversals`, this is the date the funds were debited from your bank account.
	TransferDate time.Time `json:"transfer_date,required" format:"date"`
	// The trace number of the funding event.
	TraceNumber string                            `json:"trace_number,nullable"`
	JSON        fundingEventSummaryItemV1DataJSON `json:"-"`
}

// fundingEventSummaryItemV1DataJSON contains the JSON metadata for the struct
// [FundingEventSummaryItemV1Data]
type fundingEventSummaryItemV1DataJSON struct {
	ID           apijson.Field
	Amount       apijson.Field
	Direction    apijson.Field
	EventType    apijson.Field
	PaymentCount apijson.Field
	TraceNumbers apijson.Field
	TransferDate apijson.Field
	TraceNumber  apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *FundingEventSummaryItemV1Data) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fundingEventSummaryItemV1DataJSON) RawJSON() string {
	return r.raw
}

// Describes the direction of the funding event from the perspective of the
// `linked_bank_account`.
type FundingEventSummaryItemV1DataDirection string

const (
	FundingEventSummaryItemV1DataDirectionDeposit    FundingEventSummaryItemV1DataDirection = "deposit"
	FundingEventSummaryItemV1DataDirectionWithdrawal FundingEventSummaryItemV1DataDirection = "withdrawal"
)

func (r FundingEventSummaryItemV1DataDirection) IsKnown() bool {
	switch r {
	case FundingEventSummaryItemV1DataDirectionDeposit, FundingEventSummaryItemV1DataDirectionWithdrawal:
		return true
	}
	return false
}

// The funding event types describes the direction and reason for the funding
// event.
type FundingEventSummaryItemV1DataEventType string

const (
	FundingEventSummaryItemV1DataEventTypeChargeDeposit    FundingEventSummaryItemV1DataEventType = "charge_deposit"
	FundingEventSummaryItemV1DataEventTypeChargeReversal   FundingEventSummaryItemV1DataEventType = "charge_reversal"
	FundingEventSummaryItemV1DataEventTypePayoutReturn     FundingEventSummaryItemV1DataEventType = "payout_return"
	FundingEventSummaryItemV1DataEventTypePayoutWithdrawal FundingEventSummaryItemV1DataEventType = "payout_withdrawal"
)

func (r FundingEventSummaryItemV1DataEventType) IsKnown() bool {
	switch r {
	case FundingEventSummaryItemV1DataEventTypeChargeDeposit, FundingEventSummaryItemV1DataEventTypeChargeReversal, FundingEventSummaryItemV1DataEventTypePayoutReturn, FundingEventSummaryItemV1DataEventTypePayoutWithdrawal:
		return true
	}
	return false
}

// Indicates the structure of the returned content.
//
//   - "object" means the `data` field contains a single JSON object.
//   - "array" means the `data` field contains an array of objects.
//   - "error" means the `data` field contains an error object with details of the
//     issue.
//   - "none" means no data is returned.
type FundingEventSummaryItemV1ResponseType string

const (
	FundingEventSummaryItemV1ResponseTypeObject FundingEventSummaryItemV1ResponseType = "object"
	FundingEventSummaryItemV1ResponseTypeArray  FundingEventSummaryItemV1ResponseType = "array"
	FundingEventSummaryItemV1ResponseTypeError  FundingEventSummaryItemV1ResponseType = "error"
	FundingEventSummaryItemV1ResponseTypeNone   FundingEventSummaryItemV1ResponseType = "none"
)

func (r FundingEventSummaryItemV1ResponseType) IsKnown() bool {
	switch r {
	case FundingEventSummaryItemV1ResponseTypeObject, FundingEventSummaryItemV1ResponseTypeArray, FundingEventSummaryItemV1ResponseTypeError, FundingEventSummaryItemV1ResponseTypeNone:
		return true
	}
	return false
}

type FundingEventSummaryPagedV1 struct {
	Data []FundingEventSummaryPagedV1Data `json:"data,required"`
	Meta FundingEventSummaryPagedV1Meta   `json:"meta,required"`
	// Indicates the structure of the returned content.
	//
	//   - "object" means the `data` field contains a single JSON object.
	//   - "array" means the `data` field contains an array of objects.
	//   - "error" means the `data` field contains an error object with details of the
	//     issue.
	//   - "none" means no data is returned.
	ResponseType FundingEventSummaryPagedV1ResponseType `json:"response_type,required"`
	JSON         fundingEventSummaryPagedV1JSON         `json:"-"`
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

type FundingEventSummaryPagedV1Data struct {
	// Unique identifier for the funding event.
	ID string `json:"id,required" format:"uuid"`
	// The amount of the funding event in cents.
	Amount int64 `json:"amount,required"`
	// Describes the direction of the funding event from the perspective of the
	// `linked_bank_account`.
	Direction FundingEventSummaryPagedV1DataDirection `json:"direction,required"`
	// The funding event types describes the direction and reason for the funding
	// event.
	EventType FundingEventSummaryPagedV1DataEventType `json:"event_type,required"`
	// The number of payments associated with the funding event.
	PaymentCount int64 `json:"payment_count,required"`
	// Trace number.
	TraceNumbers []string `json:"trace_numbers,required"`
	// The date on which the funding event occurred. For `deposits` and `returns`, this
	// is the date the funds were credited to your bank account. For `withdrawals` and
	// `reversals`, this is the date the funds were debited from your bank account.
	TransferDate time.Time `json:"transfer_date,required" format:"date"`
	// The trace number of the funding event.
	TraceNumber string                             `json:"trace_number,nullable"`
	JSON        fundingEventSummaryPagedV1DataJSON `json:"-"`
}

// fundingEventSummaryPagedV1DataJSON contains the JSON metadata for the struct
// [FundingEventSummaryPagedV1Data]
type fundingEventSummaryPagedV1DataJSON struct {
	ID           apijson.Field
	Amount       apijson.Field
	Direction    apijson.Field
	EventType    apijson.Field
	PaymentCount apijson.Field
	TraceNumbers apijson.Field
	TransferDate apijson.Field
	TraceNumber  apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *FundingEventSummaryPagedV1Data) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fundingEventSummaryPagedV1DataJSON) RawJSON() string {
	return r.raw
}

// Describes the direction of the funding event from the perspective of the
// `linked_bank_account`.
type FundingEventSummaryPagedV1DataDirection string

const (
	FundingEventSummaryPagedV1DataDirectionDeposit    FundingEventSummaryPagedV1DataDirection = "deposit"
	FundingEventSummaryPagedV1DataDirectionWithdrawal FundingEventSummaryPagedV1DataDirection = "withdrawal"
)

func (r FundingEventSummaryPagedV1DataDirection) IsKnown() bool {
	switch r {
	case FundingEventSummaryPagedV1DataDirectionDeposit, FundingEventSummaryPagedV1DataDirectionWithdrawal:
		return true
	}
	return false
}

// The funding event types describes the direction and reason for the funding
// event.
type FundingEventSummaryPagedV1DataEventType string

const (
	FundingEventSummaryPagedV1DataEventTypeChargeDeposit    FundingEventSummaryPagedV1DataEventType = "charge_deposit"
	FundingEventSummaryPagedV1DataEventTypeChargeReversal   FundingEventSummaryPagedV1DataEventType = "charge_reversal"
	FundingEventSummaryPagedV1DataEventTypePayoutReturn     FundingEventSummaryPagedV1DataEventType = "payout_return"
	FundingEventSummaryPagedV1DataEventTypePayoutWithdrawal FundingEventSummaryPagedV1DataEventType = "payout_withdrawal"
)

func (r FundingEventSummaryPagedV1DataEventType) IsKnown() bool {
	switch r {
	case FundingEventSummaryPagedV1DataEventTypeChargeDeposit, FundingEventSummaryPagedV1DataEventTypeChargeReversal, FundingEventSummaryPagedV1DataEventTypePayoutReturn, FundingEventSummaryPagedV1DataEventTypePayoutWithdrawal:
		return true
	}
	return false
}

type FundingEventSummaryPagedV1Meta struct {
	// Unique identifier for this API request, useful for troubleshooting.
	APIRequestID string `json:"api_request_id,required" format:"uuid"`
	// Timestamp for this API request, useful for troubleshooting.
	APIRequestTimestamp time.Time `json:"api_request_timestamp,required" format:"date-time"`
	// Maximum allowed page size for this endpoint.
	MaxPageSize int64 `json:"max_page_size,required"`
	// Page number for paginated results.
	PageNumber int64 `json:"page_number,required"`
	// Number of items per page in this response.
	PageSize int64 `json:"page_size,required"`
	// The field that the results were sorted by.
	SortBy     string                                  `json:"sort_by,required"`
	SortOrder  FundingEventSummaryPagedV1MetaSortOrder `json:"sort_order,required"`
	TotalItems int64                                   `json:"total_items,required"`
	// The number of pages available.
	TotalPages int64                              `json:"total_pages,required"`
	JSON       fundingEventSummaryPagedV1MetaJSON `json:"-"`
}

// fundingEventSummaryPagedV1MetaJSON contains the JSON metadata for the struct
// [FundingEventSummaryPagedV1Meta]
type fundingEventSummaryPagedV1MetaJSON struct {
	APIRequestID        apijson.Field
	APIRequestTimestamp apijson.Field
	MaxPageSize         apijson.Field
	PageNumber          apijson.Field
	PageSize            apijson.Field
	SortBy              apijson.Field
	SortOrder           apijson.Field
	TotalItems          apijson.Field
	TotalPages          apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *FundingEventSummaryPagedV1Meta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fundingEventSummaryPagedV1MetaJSON) RawJSON() string {
	return r.raw
}

type FundingEventSummaryPagedV1MetaSortOrder string

const (
	FundingEventSummaryPagedV1MetaSortOrderAsc  FundingEventSummaryPagedV1MetaSortOrder = "asc"
	FundingEventSummaryPagedV1MetaSortOrderDesc FundingEventSummaryPagedV1MetaSortOrder = "desc"
)

func (r FundingEventSummaryPagedV1MetaSortOrder) IsKnown() bool {
	switch r {
	case FundingEventSummaryPagedV1MetaSortOrderAsc, FundingEventSummaryPagedV1MetaSortOrderDesc:
		return true
	}
	return false
}

// Indicates the structure of the returned content.
//
//   - "object" means the `data` field contains a single JSON object.
//   - "array" means the `data` field contains an array of objects.
//   - "error" means the `data` field contains an error object with details of the
//     issue.
//   - "none" means no data is returned.
type FundingEventSummaryPagedV1ResponseType string

const (
	FundingEventSummaryPagedV1ResponseTypeObject FundingEventSummaryPagedV1ResponseType = "object"
	FundingEventSummaryPagedV1ResponseTypeArray  FundingEventSummaryPagedV1ResponseType = "array"
	FundingEventSummaryPagedV1ResponseTypeError  FundingEventSummaryPagedV1ResponseType = "error"
	FundingEventSummaryPagedV1ResponseTypeNone   FundingEventSummaryPagedV1ResponseType = "none"
)

func (r FundingEventSummaryPagedV1ResponseType) IsKnown() bool {
	switch r {
	case FundingEventSummaryPagedV1ResponseTypeObject, FundingEventSummaryPagedV1ResponseTypeArray, FundingEventSummaryPagedV1ResponseTypeError, FundingEventSummaryPagedV1ResponseTypeNone:
		return true
	}
	return false
}

type FundingEventListParams struct {
	// The start date of the range to filter by using the `YYYY-MM-DD` format.
	CreatedFrom param.Field[time.Time] `query:"created_from" format:"date"`
	// The end date of the range to filter by using the `YYYY-MM-DD` format.
	CreatedTo param.Field[time.Time] `query:"created_to" format:"date"`
	// Describes the direction of the funding event from the perspective of the
	// `linked_bank_account`.
	Direction param.Field[FundingEventListParamsDirection] `query:"direction"`
	// The funding event types describes the direction and reason for the funding
	// event.
	EventType param.Field[FundingEventListParamsEventType] `query:"event_type"`
	// Results page number. Starts at page 1.
	PageNumber param.Field[int64] `query:"page_number"`
	// Results page size. Max value: 1000
	PageSize param.Field[int64] `query:"page_size"`
	// The field to sort the results by.
	SortBy param.Field[FundingEventListParamsSortBy] `query:"sort_by"`
	// The order in which to sort the results.
	SortOrder param.Field[FundingEventListParamsSortOrder] `query:"sort_order"`
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

// Describes the direction of the funding event from the perspective of the
// `linked_bank_account`.
type FundingEventListParamsDirection string

const (
	FundingEventListParamsDirectionDeposit    FundingEventListParamsDirection = "deposit"
	FundingEventListParamsDirectionWithdrawal FundingEventListParamsDirection = "withdrawal"
)

func (r FundingEventListParamsDirection) IsKnown() bool {
	switch r {
	case FundingEventListParamsDirectionDeposit, FundingEventListParamsDirectionWithdrawal:
		return true
	}
	return false
}

// The funding event types describes the direction and reason for the funding
// event.
type FundingEventListParamsEventType string

const (
	FundingEventListParamsEventTypeChargeDeposit    FundingEventListParamsEventType = "charge_deposit"
	FundingEventListParamsEventTypeChargeReversal   FundingEventListParamsEventType = "charge_reversal"
	FundingEventListParamsEventTypePayoutReturn     FundingEventListParamsEventType = "payout_return"
	FundingEventListParamsEventTypePayoutWithdrawal FundingEventListParamsEventType = "payout_withdrawal"
)

func (r FundingEventListParamsEventType) IsKnown() bool {
	switch r {
	case FundingEventListParamsEventTypeChargeDeposit, FundingEventListParamsEventTypeChargeReversal, FundingEventListParamsEventTypePayoutReturn, FundingEventListParamsEventTypePayoutWithdrawal:
		return true
	}
	return false
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

// The order in which to sort the results.
type FundingEventListParamsSortOrder string

const (
	FundingEventListParamsSortOrderAsc  FundingEventListParamsSortOrder = "asc"
	FundingEventListParamsSortOrderDesc FundingEventListParamsSortOrder = "desc"
)

func (r FundingEventListParamsSortOrder) IsKnown() bool {
	switch r {
	case FundingEventListParamsSortOrderAsc, FundingEventListParamsSortOrderDesc:
		return true
	}
	return false
}

type FundingEventGetParams struct {
	CorrelationID     param.Field[string] `header:"Correlation-Id"`
	RequestID         param.Field[string] `header:"Request-Id"`
	StraddleAccountID param.Field[string] `header:"Straddle-Account-Id" format:"uuid"`
}
