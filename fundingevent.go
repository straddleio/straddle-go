// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package straddle

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/straddleio/straddle-go/internal/apijson"
	"github.com/straddleio/straddle-go/internal/apiquery"
	"github.com/straddleio/straddle-go/internal/requestconfig"
	"github.com/straddleio/straddle-go/option"
	"github.com/straddleio/straddle-go/packages/pagination"
	"github.com/straddleio/straddle-go/packages/param"
	"github.com/straddleio/straddle-go/packages/respjson"
	"github.com/straddleio/straddle-go/shared"
)

// Funding events represent all money movement between Straddle and an Account's
// external bank accounts. They are automatically generated when charges settle or
// payouts are initiated. Each event provides detailed tracking of settlement
// status, fee breakdowns, and reconciliation data across both incoming and
// outgoing transfers. Use funding events to monitor your platform's entire money
// movement lifecycle.
//
// FundingEventService contains methods and other services that help with
// interacting with the straddle API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFundingEventService] method instead.
type FundingEventService struct {
	options []option.RequestOption
}

// NewFundingEventService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewFundingEventService(opts ...option.RequestOption) (r FundingEventService) {
	r = FundingEventService{}
	r.options = opts
	return
}

// Retrieves a list of funding events for your account. This endpoint supports
// advanced sorting and filtering options.
func (r *FundingEventService) List(ctx context.Context, params FundingEventListParams, opts ...option.RequestOption) (res *pagination.PageNumberSchema[FundingEventSummaryPagedV1Data], err error) {
	var raw *http.Response
	if !param.IsOmitted(params.CorrelationID) {
		opts = append(opts, option.WithHeader("Correlation-Id", fmt.Sprintf("%v", params.CorrelationID.Value)))
	}
	if !param.IsOmitted(params.RequestID) {
		opts = append(opts, option.WithHeader("Request-Id", fmt.Sprintf("%v", params.RequestID.Value)))
	}
	if !param.IsOmitted(params.StraddleAccountID) {
		opts = append(opts, option.WithHeader("Straddle-Account-Id", fmt.Sprintf("%v", params.StraddleAccountID.Value)))
	}
	opts = slices.Concat(r.options, opts)
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
	if !param.IsOmitted(query.CorrelationID) {
		opts = append(opts, option.WithHeader("Correlation-Id", fmt.Sprintf("%v", query.CorrelationID.Value)))
	}
	if !param.IsOmitted(query.RequestID) {
		opts = append(opts, option.WithHeader("Request-Id", fmt.Sprintf("%v", query.RequestID.Value)))
	}
	if !param.IsOmitted(query.StraddleAccountID) {
		opts = append(opts, option.WithHeader("Straddle-Account-Id", fmt.Sprintf("%v", query.StraddleAccountID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/funding_events/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type FundingEventSummaryItemV1 struct {
	Data FundingEventSummaryItemV1Data `json:"data" api:"required"`
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
	ResponseType FundingEventSummaryItemV1ResponseType `json:"response_type" api:"required"`
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
func (r FundingEventSummaryItemV1) RawJSON() string { return r.JSON.raw }
func (r *FundingEventSummaryItemV1) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FundingEventSummaryItemV1Data struct {
	// Unique identifier for the funding event.
	ID string `json:"id" api:"required" format:"uuid"`
	// The amount of the funding event in cents.
	Amount int64 `json:"amount" api:"required"`
	// Created at.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Describes the direction of the funding event from the perspective of the
	// `linked_bank_account`.
	//
	// Any of "deposit", "withdrawal".
	Direction string `json:"direction" api:"required"`
	// The funding event types describes the direction and reason for the funding
	// event.
	//
	// Any of "charge_deposit", "charge_reversal", "payout_return",
	// "payout_withdrawal".
	EventType string `json:"event_type" api:"required"`
	// The number of payments associated with the funding event.
	PaymentCount int64 `json:"payment_count" api:"required"`
	// Trace Ids.
	TraceIDs map[string]string `json:"trace_ids" api:"required"`
	// Trace number.
	TraceNumbers []string `json:"trace_numbers" api:"required"`
	// The date on which the funding event occurred. For `deposits` and `returns`, this
	// is the date the funds were credited to your bank account. For `withdrawals` and
	// `reversals`, this is the date the funds were debited from your bank account.
	TransferDate time.Time `json:"transfer_date" api:"required" format:"date"`
	// Updated at.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// The current status of the `charge` or `payout`.
	//
	// Any of "created", "scheduled", "failed", "cancelled", "on_hold", "pending",
	// "paid", "reversed", "validating".
	Status        string                                     `json:"status"`
	StatusDetails FundingEventSummaryItemV1DataStatusDetails `json:"status_details"`
	// The trace number of the funding event.
	TraceNumber string `json:"trace_number" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		Amount        respjson.Field
		CreatedAt     respjson.Field
		Direction     respjson.Field
		EventType     respjson.Field
		PaymentCount  respjson.Field
		TraceIDs      respjson.Field
		TraceNumbers  respjson.Field
		TransferDate  respjson.Field
		UpdatedAt     respjson.Field
		Status        respjson.Field
		StatusDetails respjson.Field
		TraceNumber   respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FundingEventSummaryItemV1Data) RawJSON() string { return r.JSON.raw }
func (r *FundingEventSummaryItemV1Data) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FundingEventSummaryItemV1DataStatusDetails struct {
	// The time the status change occurred.
	ChangedAt time.Time `json:"changed_at" api:"required" format:"date-time"`
	// A human-readable description of the current status.
	Message string `json:"message" api:"required"`
	// Any of "insufficient_funds", "closed_bank_account", "invalid_bank_account",
	// "invalid_routing", "disputed", "payment_stopped", "owner_deceased",
	// "frozen_bank_account", "risk_review", "fraudulent", "duplicate_entry",
	// "invalid_paykey", "payment_blocked", "amount_too_large", "too_many_attempts",
	// "internal_system_error", "user_request", "ok", "other_network_return",
	// "payout_refused", "cancel_request", "failed_verification", "require_review",
	// "blocked_by_system", "watchtower_review", "validating", "auto_hold".
	Reason string `json:"reason" api:"required"`
	// Any of "watchtower", "bank_decline", "customer_dispute", "user_action",
	// "system".
	Source string `json:"source" api:"required"`
	// The status code if applicable.
	Code string `json:"code" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChangedAt   respjson.Field
		Message     respjson.Field
		Reason      respjson.Field
		Source      respjson.Field
		Code        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FundingEventSummaryItemV1DataStatusDetails) RawJSON() string { return r.JSON.raw }
func (r *FundingEventSummaryItemV1DataStatusDetails) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
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

type FundingEventSummaryPagedV1 struct {
	Data []FundingEventSummaryPagedV1Data `json:"data" api:"required"`
	Meta FundingEventSummaryPagedV1Meta   `json:"meta" api:"required"`
	// Indicates the structure of the returned content.
	//
	//   - "object" means the `data` field contains a single JSON object.
	//   - "array" means the `data` field contains an array of objects.
	//   - "error" means the `data` field contains an error object with details of the
	//     issue.
	//   - "none" means no data is returned.
	//
	// Any of "object", "array", "error", "none".
	ResponseType FundingEventSummaryPagedV1ResponseType `json:"response_type" api:"required"`
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
func (r FundingEventSummaryPagedV1) RawJSON() string { return r.JSON.raw }
func (r *FundingEventSummaryPagedV1) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FundingEventSummaryPagedV1Data struct {
	// Unique identifier for the funding event.
	ID string `json:"id" api:"required" format:"uuid"`
	// The amount of the funding event in cents.
	Amount int64 `json:"amount" api:"required"`
	// Created at.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Describes the direction of the funding event from the perspective of the
	// `linked_bank_account`.
	//
	// Any of "deposit", "withdrawal".
	Direction string `json:"direction" api:"required"`
	// The funding event types describes the direction and reason for the funding
	// event.
	//
	// Any of "charge_deposit", "charge_reversal", "payout_return",
	// "payout_withdrawal".
	EventType string `json:"event_type" api:"required"`
	// The number of payments associated with the funding event.
	PaymentCount int64 `json:"payment_count" api:"required"`
	// Trace Ids.
	TraceIDs map[string]string `json:"trace_ids" api:"required"`
	// Trace number.
	TraceNumbers []string `json:"trace_numbers" api:"required"`
	// The date on which the funding event occurred. For `deposits` and `returns`, this
	// is the date the funds were credited to your bank account. For `withdrawals` and
	// `reversals`, this is the date the funds were debited from your bank account.
	TransferDate time.Time `json:"transfer_date" api:"required" format:"date"`
	// Updated at.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// The current status of the `charge` or `payout`.
	//
	// Any of "created", "scheduled", "failed", "cancelled", "on_hold", "pending",
	// "paid", "reversed", "validating".
	Status        string                                      `json:"status"`
	StatusDetails FundingEventSummaryPagedV1DataStatusDetails `json:"status_details"`
	// The trace number of the funding event.
	TraceNumber string `json:"trace_number" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		Amount        respjson.Field
		CreatedAt     respjson.Field
		Direction     respjson.Field
		EventType     respjson.Field
		PaymentCount  respjson.Field
		TraceIDs      respjson.Field
		TraceNumbers  respjson.Field
		TransferDate  respjson.Field
		UpdatedAt     respjson.Field
		Status        respjson.Field
		StatusDetails respjson.Field
		TraceNumber   respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FundingEventSummaryPagedV1Data) RawJSON() string { return r.JSON.raw }
func (r *FundingEventSummaryPagedV1Data) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FundingEventSummaryPagedV1DataStatusDetails struct {
	// The time the status change occurred.
	ChangedAt time.Time `json:"changed_at" api:"required" format:"date-time"`
	// A human-readable description of the current status.
	Message string `json:"message" api:"required"`
	// Any of "insufficient_funds", "closed_bank_account", "invalid_bank_account",
	// "invalid_routing", "disputed", "payment_stopped", "owner_deceased",
	// "frozen_bank_account", "risk_review", "fraudulent", "duplicate_entry",
	// "invalid_paykey", "payment_blocked", "amount_too_large", "too_many_attempts",
	// "internal_system_error", "user_request", "ok", "other_network_return",
	// "payout_refused", "cancel_request", "failed_verification", "require_review",
	// "blocked_by_system", "watchtower_review", "validating", "auto_hold".
	Reason string `json:"reason" api:"required"`
	// Any of "watchtower", "bank_decline", "customer_dispute", "user_action",
	// "system".
	Source string `json:"source" api:"required"`
	// The status code if applicable.
	Code string `json:"code" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChangedAt   respjson.Field
		Message     respjson.Field
		Reason      respjson.Field
		Source      respjson.Field
		Code        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FundingEventSummaryPagedV1DataStatusDetails) RawJSON() string { return r.JSON.raw }
func (r *FundingEventSummaryPagedV1DataStatusDetails) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FundingEventSummaryPagedV1Meta struct {
	// Unique identifier for this API request, useful for troubleshooting.
	APIRequestID string `json:"api_request_id" api:"required" format:"uuid"`
	// Timestamp for this API request, useful for troubleshooting.
	APIRequestTimestamp time.Time `json:"api_request_timestamp" api:"required" format:"date-time"`
	// Maximum allowed page size for this endpoint.
	MaxPageSize int64 `json:"max_page_size" api:"required"`
	// Page number for paginated results.
	PageNumber int64 `json:"page_number" api:"required"`
	// Number of items per page in this response.
	PageSize int64 `json:"page_size" api:"required"`
	// The field that the results were sorted by.
	SortBy string `json:"sort_by" api:"required"`
	// Any of "asc", "desc".
	SortOrder  string `json:"sort_order" api:"required"`
	TotalItems int64  `json:"total_items" api:"required"`
	// The number of pages available.
	TotalPages int64 `json:"total_pages" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		APIRequestID        respjson.Field
		APIRequestTimestamp respjson.Field
		MaxPageSize         respjson.Field
		PageNumber          respjson.Field
		PageSize            respjson.Field
		SortBy              respjson.Field
		SortOrder           respjson.Field
		TotalItems          respjson.Field
		TotalPages          respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FundingEventSummaryPagedV1Meta) RawJSON() string { return r.JSON.raw }
func (r *FundingEventSummaryPagedV1Meta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
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

type FundingEventListParams struct {
	// The start date of the range to filter by using the `YYYY-MM-DD` format.
	CreatedFrom param.Opt[time.Time] `query:"created_from,omitzero" format:"date" json:"-"`
	// The end date of the range to filter by using the `YYYY-MM-DD` format.
	CreatedTo param.Opt[time.Time] `query:"created_to,omitzero" format:"date" json:"-"`
	// Search text.
	SearchText param.Opt[string] `query:"search_text,omitzero" json:"-"`
	// Trace Id.
	TraceID param.Opt[string] `query:"trace_id,omitzero" json:"-"`
	// Trace number.
	TraceNumber param.Opt[string] `query:"trace_number,omitzero" json:"-"`
	// Results page number. Starts at page 1.
	PageNumber param.Opt[int64] `query:"page_number,omitzero" json:"-"`
	// Results page size. Max value: 1000
	PageSize          param.Opt[int64]  `query:"page_size,omitzero" json:"-"`
	CorrelationID     param.Opt[string] `header:"Correlation-Id,omitzero" json:"-"`
	RequestID         param.Opt[string] `header:"Request-Id,omitzero" json:"-"`
	StraddleAccountID param.Opt[string] `header:"Straddle-Account-Id,omitzero" format:"uuid" json:"-"`
	// Funding Event status.
	//
	// Any of "created", "scheduled", "failed", "cancelled", "on_hold", "pending",
	// "paid", "reversed", "validating".
	Status []string `query:"status,omitzero" json:"-"`
	// Reason for latest payment status change.
	//
	// Any of "insufficient_funds", "closed_bank_account", "invalid_bank_account",
	// "invalid_routing", "disputed", "payment_stopped", "owner_deceased",
	// "frozen_bank_account", "risk_review", "fraudulent", "duplicate_entry",
	// "invalid_paykey", "payment_blocked", "amount_too_large", "too_many_attempts",
	// "internal_system_error", "user_request", "ok", "other_network_return",
	// "payout_refused", "cancel_request", "failed_verification", "require_review",
	// "blocked_by_system", "watchtower_review", "validating", "auto_hold".
	StatusReason []string `query:"status_reason,omitzero" json:"-"`
	// Source of latest payment status change.
	//
	// Any of "watchtower", "bank_decline", "customer_dispute", "user_action",
	// "system".
	StatusSource []string `query:"status_source,omitzero" json:"-"`
	// Describes the direction of the funding event from the perspective of the
	// `linked_bank_account`.
	//
	// Any of "deposit", "withdrawal".
	Direction FundingEventListParamsDirection `query:"direction,omitzero" json:"-"`
	// The funding event types describes the direction and reason for the funding
	// event.
	//
	// Any of "charge_deposit", "charge_reversal", "payout_return",
	// "payout_withdrawal".
	EventType FundingEventListParamsEventType `query:"event_type,omitzero" json:"-"`
	// The field to sort the results by.
	//
	// Any of "transfer_date", "id", "amount".
	SortBy FundingEventListParamsSortBy `query:"sort_by,omitzero" json:"-"`
	// The order in which to sort the results.
	//
	// Any of "asc", "desc".
	SortOrder FundingEventListParamsSortOrder `query:"sort_order,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [FundingEventListParams]'s query parameters as `url.Values`.
func (r FundingEventListParams) URLQuery() (v url.Values, err error) {
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

// The funding event types describes the direction and reason for the funding
// event.
type FundingEventListParamsEventType string

const (
	FundingEventListParamsEventTypeChargeDeposit    FundingEventListParamsEventType = "charge_deposit"
	FundingEventListParamsEventTypeChargeReversal   FundingEventListParamsEventType = "charge_reversal"
	FundingEventListParamsEventTypePayoutReturn     FundingEventListParamsEventType = "payout_return"
	FundingEventListParamsEventTypePayoutWithdrawal FundingEventListParamsEventType = "payout_withdrawal"
)

// The field to sort the results by.
type FundingEventListParamsSortBy string

const (
	FundingEventListParamsSortByTransferDate FundingEventListParamsSortBy = "transfer_date"
	FundingEventListParamsSortByID           FundingEventListParamsSortBy = "id"
	FundingEventListParamsSortByAmount       FundingEventListParamsSortBy = "amount"
)

// The order in which to sort the results.
type FundingEventListParamsSortOrder string

const (
	FundingEventListParamsSortOrderAsc  FundingEventListParamsSortOrder = "asc"
	FundingEventListParamsSortOrderDesc FundingEventListParamsSortOrder = "desc"
)

type FundingEventGetParams struct {
	CorrelationID     param.Opt[string] `header:"Correlation-Id,omitzero" json:"-"`
	RequestID         param.Opt[string] `header:"Request-Id,omitzero" json:"-"`
	StraddleAccountID param.Opt[string] `header:"Straddle-Account-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}
