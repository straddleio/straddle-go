// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package straddle

import (
	"context"
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

// Payments provide endpoints to filter both Charges and Payouts with multiple
// different parameters.
//
// PaymentService contains methods and other services that help with interacting
// with the straddle API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPaymentService] method instead.
type PaymentService struct {
	options []option.RequestOption
}

// NewPaymentService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewPaymentService(opts ...option.RequestOption) (r PaymentService) {
	r = PaymentService{}
	r.options = opts
	return
}

// Search for payments, including `charges` and `payouts`, using a variety of
// criteria. This endpoint supports advanced sorting and filtering options.
func (r *PaymentService) List(ctx context.Context, params PaymentListParams, opts ...option.RequestOption) (res *pagination.PageNumberSchema[PaymentSummaryPagedV1Data], err error) {
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
	path := "v1/payments"
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

// Search for payments, including `charges` and `payouts`, using a variety of
// criteria. This endpoint supports advanced sorting and filtering options.
func (r *PaymentService) ListAutoPaging(ctx context.Context, params PaymentListParams, opts ...option.RequestOption) *pagination.PageNumberSchemaAutoPager[PaymentSummaryPagedV1Data] {
	return pagination.NewPageNumberSchemaAutoPager(r.List(ctx, params, opts...))
}

type PaymentSummaryPagedV1 struct {
	Data []PaymentSummaryPagedV1Data `json:"data" api:"required"`
	Meta PaymentSummaryPagedV1Meta   `json:"meta" api:"required"`
	// Indicates the structure of the returned content.
	//
	//   - "object" means the `data` field contains a single JSON object.
	//   - "array" means the `data` field contains an array of objects.
	//   - "error" means the `data` field contains an error object with details of the
	//     issue.
	//   - "none" means no data is returned.
	//
	// Any of "object", "array", "error", "none".
	ResponseType PaymentSummaryPagedV1ResponseType `json:"response_type" api:"required"`
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
func (r PaymentSummaryPagedV1) RawJSON() string { return r.JSON.raw }
func (r *PaymentSummaryPagedV1) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaymentSummaryPagedV1Data struct {
	// Unique identifier for the `charge` or `payout`.
	ID string `json:"id" api:"required" format:"uuid"`
	// The amount of the `charge` or `payout` in cents.
	Amount int64 `json:"amount" api:"required"`
	// The time the `charge` or `payout` was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// The currency of the `charge` or `payout`. Only USD is supported.
	Currency string `json:"currency" api:"required"`
	// An arbitrary description for the `charge` or `payout`.
	Description string `json:"description" api:"required"`
	// Unique identifier for the `charge` or `payout` in your database. This value must
	// be unique across all charges or payouts.
	ExternalID string `json:"external_id" api:"required"`
	// Funding ids.
	FundingIDs []string `json:"funding_ids" api:"required" format:"uuid"`
	// Value of the `paykey` used for the `charge` or `payout`.
	Paykey string `json:"paykey" api:"required"`
	// The desired date on which the payment should be occur. For charges, this means
	// the date you want the customer to be debited on. For payouts, this means the
	// date you want the funds to be sent from your bank account.
	PaymentDate time.Time `json:"payment_date" api:"required" format:"date"`
	// The type of payment. Valid values are `charge` or `payout`.
	//
	// Any of "charge", "payout".
	PaymentType string `json:"payment_type" api:"required"`
	// The current status of the `charge` or `payout`.
	//
	// Any of "created", "scheduled", "failed", "cancelled", "on_hold", "pending",
	// "paid", "reversed", "validating".
	Status string `json:"status" api:"required"`
	// Details about the current status of the `charge` or `payout`.
	StatusDetails shared.StatusDetailsV1 `json:"status_details" api:"required"`
	// Trace ids.
	TraceIDs map[string]string `json:"trace_ids" api:"required"`
	// The time the `charge` or `payout` was last updated.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// Information about the customer associated with the charge or payout.
	CustomerDetails shared.CustomerDetailsV1 `json:"customer_details"`
	// The actual date on which the payment occurred. For charges, this is the date the
	// customer was debited. For payouts, this is the date the funds were sent from
	// your bank account.
	EffectiveAt time.Time `json:"effective_at" api:"nullable" format:"date-time"`
	// Unique identifier for the funding event associated with the `charge` or
	// `payout`.
	FundingID string `json:"funding_id" api:"nullable" format:"uuid"`
	// Metadata for payment - only included if requested.
	Metadata map[string]string `json:"metadata" api:"nullable"`
	// Information about the paykey used for the `charge` or `payout`.
	PaykeyDetails shared.PaykeyDetailsV1 `json:"paykey_details"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		Amount          respjson.Field
		CreatedAt       respjson.Field
		Currency        respjson.Field
		Description     respjson.Field
		ExternalID      respjson.Field
		FundingIDs      respjson.Field
		Paykey          respjson.Field
		PaymentDate     respjson.Field
		PaymentType     respjson.Field
		Status          respjson.Field
		StatusDetails   respjson.Field
		TraceIDs        respjson.Field
		UpdatedAt       respjson.Field
		CustomerDetails respjson.Field
		EffectiveAt     respjson.Field
		FundingID       respjson.Field
		Metadata        respjson.Field
		PaykeyDetails   respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaymentSummaryPagedV1Data) RawJSON() string { return r.JSON.raw }
func (r *PaymentSummaryPagedV1Data) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaymentSummaryPagedV1Meta struct {
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
func (r PaymentSummaryPagedV1Meta) RawJSON() string { return r.JSON.raw }
func (r *PaymentSummaryPagedV1Meta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indicates the structure of the returned content.
//
//   - "object" means the `data` field contains a single JSON object.
//   - "array" means the `data` field contains an array of objects.
//   - "error" means the `data` field contains an error object with details of the
//     issue.
//   - "none" means no data is returned.
type PaymentSummaryPagedV1ResponseType string

const (
	PaymentSummaryPagedV1ResponseTypeObject PaymentSummaryPagedV1ResponseType = "object"
	PaymentSummaryPagedV1ResponseTypeArray  PaymentSummaryPagedV1ResponseType = "array"
	PaymentSummaryPagedV1ResponseTypeError  PaymentSummaryPagedV1ResponseType = "error"
	PaymentSummaryPagedV1ResponseTypeNone   PaymentSummaryPagedV1ResponseType = "none"
)

type PaymentListParams struct {
	// Search using the `customer_id` of a `charge` or `payout`.
	CustomerID      param.Opt[string] `query:"customer_id,omitzero" format:"uuid" json:"-"`
	DefaultPageSize param.Opt[int64]  `query:"default_page_size,omitzero" json:"-"`
	// Search using the `external_id` of a `charge` or `payout`.
	ExternalID param.Opt[string] `query:"external_id,omitzero" json:"-"`
	// Search using the `funding_id` of a `charge` or `payout`.
	FundingID param.Opt[string] `query:"funding_id,omitzero" format:"uuid" json:"-"`
	// Include the metadata for payments in the returned data.
	IncludeMetadata param.Opt[bool] `query:"include_metadata,omitzero" json:"-"`
	// Search using a maximum `amount` of a `charge` or `payout`.
	MaxAmount param.Opt[int64] `query:"max_amount,omitzero" json:"-"`
	// Search using the latest `created_at` date of a `charge` or `payout`.
	MaxCreatedAt param.Opt[time.Time] `query:"max_created_at,omitzero" format:"date-time" json:"-"`
	// Search using the latest `effective_date` of a `charge` or `payout`.
	MaxEffectiveAt param.Opt[time.Time] `query:"max_effective_at,omitzero" format:"date-time" json:"-"`
	// Search using the latest `payment_date` of a `charge` or `payout`.
	MaxPaymentDate param.Opt[time.Time] `query:"max_payment_date,omitzero" format:"date" json:"-"`
	// Search using the minimum `amount of a `charge`or`payout`.
	MinAmount param.Opt[int64] `query:"min_amount,omitzero" json:"-"`
	// Search using the earliest `created_at` date of a `charge` or `payout`.
	MinCreatedAt param.Opt[time.Time] `query:"min_created_at,omitzero" format:"date-time" json:"-"`
	// Search using the earliest `effective_date` of a `charge` or `payout`.
	MinEffectiveAt param.Opt[time.Time] `query:"min_effective_at,omitzero" format:"date-time" json:"-"`
	// Search using the earliest ` `of a `charge` or `payout`.
	MinPaymentDate param.Opt[time.Time] `query:"min_payment_date,omitzero" format:"date" json:"-"`
	// Results page number. Starts at page 1.
	PageNumber param.Opt[int64] `query:"page_number,omitzero" json:"-"`
	// Results page size. Max value: 1000
	PageSize param.Opt[int64] `query:"page_size,omitzero" json:"-"`
	// Search using the `paykey` of a `charge` or `payout`.
	Paykey param.Opt[string] `query:"paykey,omitzero" json:"-"`
	// Search using the `paykey_id` of a `charge` or `payout`.
	PaykeyID param.Opt[string] `query:"paykey_id,omitzero" format:"uuid" json:"-"`
	// Search using the `id` of a `charge` or `payout`.
	PaymentID param.Opt[string] `query:"payment_id,omitzero" format:"uuid" json:"-"`
	// Search using a text string associated with a `charge` or `payout`.
	SearchText        param.Opt[string] `query:"search_text,omitzero" json:"-"`
	CorrelationID     param.Opt[string] `header:"Correlation-Id,omitzero" json:"-"`
	RequestID         param.Opt[string] `header:"Request-Id,omitzero" json:"-"`
	StraddleAccountID param.Opt[string] `header:"Straddle-Account-Id,omitzero" format:"uuid" json:"-"`
	// The field to sort the results by.
	//
	// Any of "created_at", "payment_date", "effective_at", "id", "amount".
	DefaultSort PaymentListParamsDefaultSort `query:"default_sort,omitzero" json:"-"`
	// Any of "asc", "desc".
	DefaultSortOrder PaymentListParamsDefaultSortOrder `query:"default_sort_order,omitzero" json:"-"`
	// Search by the status of a `charge` or `payout`.
	//
	// Any of "created", "scheduled", "failed", "cancelled", "on_hold", "pending",
	// "paid", "reversed", "validating".
	PaymentStatus []string `query:"payment_status,omitzero" json:"-"`
	// Search by the type of a `charge` or `payout`.
	//
	// Any of "charge", "payout".
	PaymentType []string `query:"payment_type,omitzero" json:"-"`
	// The field to sort the results by.
	//
	// Any of "created_at", "payment_date", "effective_at", "id", "amount".
	SortBy PaymentListParamsSortBy `query:"sort_by,omitzero" json:"-"`
	// Any of "asc", "desc".
	SortOrder PaymentListParamsSortOrder `query:"sort_order,omitzero" json:"-"`
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
	paramObj
}

// URLQuery serializes [PaymentListParams]'s query parameters as `url.Values`.
func (r PaymentListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The field to sort the results by.
type PaymentListParamsDefaultSort string

const (
	PaymentListParamsDefaultSortCreatedAt   PaymentListParamsDefaultSort = "created_at"
	PaymentListParamsDefaultSortPaymentDate PaymentListParamsDefaultSort = "payment_date"
	PaymentListParamsDefaultSortEffectiveAt PaymentListParamsDefaultSort = "effective_at"
	PaymentListParamsDefaultSortID          PaymentListParamsDefaultSort = "id"
	PaymentListParamsDefaultSortAmount      PaymentListParamsDefaultSort = "amount"
)

type PaymentListParamsDefaultSortOrder string

const (
	PaymentListParamsDefaultSortOrderAsc  PaymentListParamsDefaultSortOrder = "asc"
	PaymentListParamsDefaultSortOrderDesc PaymentListParamsDefaultSortOrder = "desc"
)

// The field to sort the results by.
type PaymentListParamsSortBy string

const (
	PaymentListParamsSortByCreatedAt   PaymentListParamsSortBy = "created_at"
	PaymentListParamsSortByPaymentDate PaymentListParamsSortBy = "payment_date"
	PaymentListParamsSortByEffectiveAt PaymentListParamsSortBy = "effective_at"
	PaymentListParamsSortByID          PaymentListParamsSortBy = "id"
	PaymentListParamsSortByAmount      PaymentListParamsSortBy = "amount"
)

type PaymentListParamsSortOrder string

const (
	PaymentListParamsSortOrderAsc  PaymentListParamsSortOrder = "asc"
	PaymentListParamsSortOrderDesc PaymentListParamsSortOrder = "desc"
)
