// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package straddle

import (
	"context"
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

// PaymentService contains methods and other services that help with interacting
// with the straddle API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPaymentService] method instead.
type PaymentService struct {
	Options []option.RequestOption
}

// NewPaymentService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewPaymentService(opts ...option.RequestOption) (r *PaymentService) {
	r = &PaymentService{}
	r.Options = opts
	return
}

// Search for payments, including `charges` and `payouts`, using a variety of
// criteria. This endpoint supports advanced sorting and filtering options.
func (r *PaymentService) List(ctx context.Context, params PaymentListParams, opts ...option.RequestOption) (res *pagination.PageNumberSchema[PaymentSummaryPagedV1Data], err error) {
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
	Data []PaymentSummaryPagedV1Data `json:"data,required"`
	Meta PaymentSummaryPagedV1Meta   `json:"meta,required"`
	// Indicates the structure of the returned content.
	//
	//   - "object" means the `data` field contains a single JSON object.
	//   - "array" means the `data` field contains an array of objects.
	//   - "error" means the `data` field contains an error object with details of the
	//     issue.
	//   - "none" means no data is returned.
	ResponseType PaymentSummaryPagedV1ResponseType `json:"response_type,required"`
	JSON         paymentSummaryPagedV1JSON         `json:"-"`
}

// paymentSummaryPagedV1JSON contains the JSON metadata for the struct
// [PaymentSummaryPagedV1]
type paymentSummaryPagedV1JSON struct {
	Data         apijson.Field
	Meta         apijson.Field
	ResponseType apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *PaymentSummaryPagedV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r paymentSummaryPagedV1JSON) RawJSON() string {
	return r.raw
}

type PaymentSummaryPagedV1Data struct {
	// Unique identifier for the `charge` or `payout`.
	ID string `json:"id,required" format:"uuid"`
	// The amount of the `charge` or `payout` in cents.
	Amount int64 `json:"amount,required"`
	// The time the `charge` or `payout` was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The currency of the `charge` or `payout`. Only USD is supported.
	Currency string `json:"currency,required"`
	// An arbitrary description for the `charge` or `payout`.
	Description string `json:"description,required"`
	// Unique identifier for the `charge` or `payout` in your database. This value must
	// be unique across all charges or payouts.
	ExternalID string `json:"external_id,required"`
	// Funding ids.
	FundingIDs []string `json:"funding_ids,required" format:"uuid"`
	// Value of the `paykey` used for the `charge` or `payout`.
	Paykey string `json:"paykey,required"`
	// The desired date on which the payment should be occur. For charges, this means
	// the date you want the customer to be debited on. For payouts, this means the
	// date you want the funds to be sent from your bank account.
	PaymentDate time.Time `json:"payment_date,required" format:"date"`
	// The type of payment. Valid values are `charge` or `payout`.
	PaymentType PaymentSummaryPagedV1DataPaymentType `json:"payment_type,required"`
	// The current status of the `charge` or `payout`.
	Status PaymentSummaryPagedV1DataStatus `json:"status,required"`
	// Details about the current status of the `charge` or `payout`.
	StatusDetails shared.StatusDetailsV1 `json:"status_details,required"`
	// The time the `charge` or `payout` was last updated.
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// Information about the customer associated with the charge or payout.
	CustomerDetails shared.CustomerDetailsV1 `json:"customer_details"`
	// The actual date on which the payment occurred. For charges, this is the date the
	// customer was debited. For payouts, this is the date the funds were sent from
	// your bank account.
	EffectiveAt time.Time `json:"effective_at,nullable" format:"date-time"`
	// Unique identifier for the funding event associated with the `charge` or
	// `payout`.
	FundingID string `json:"funding_id,nullable" format:"uuid"`
	// Information about the paykey used for the `charge` or `payout`.
	PaykeyDetails shared.PaykeyDetailsV1        `json:"paykey_details"`
	JSON          paymentSummaryPagedV1DataJSON `json:"-"`
}

// paymentSummaryPagedV1DataJSON contains the JSON metadata for the struct
// [PaymentSummaryPagedV1Data]
type paymentSummaryPagedV1DataJSON struct {
	ID              apijson.Field
	Amount          apijson.Field
	CreatedAt       apijson.Field
	Currency        apijson.Field
	Description     apijson.Field
	ExternalID      apijson.Field
	FundingIDs      apijson.Field
	Paykey          apijson.Field
	PaymentDate     apijson.Field
	PaymentType     apijson.Field
	Status          apijson.Field
	StatusDetails   apijson.Field
	UpdatedAt       apijson.Field
	CustomerDetails apijson.Field
	EffectiveAt     apijson.Field
	FundingID       apijson.Field
	PaykeyDetails   apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *PaymentSummaryPagedV1Data) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r paymentSummaryPagedV1DataJSON) RawJSON() string {
	return r.raw
}

// The type of payment. Valid values are `charge` or `payout`.
type PaymentSummaryPagedV1DataPaymentType string

const (
	PaymentSummaryPagedV1DataPaymentTypeCharge PaymentSummaryPagedV1DataPaymentType = "charge"
	PaymentSummaryPagedV1DataPaymentTypePayout PaymentSummaryPagedV1DataPaymentType = "payout"
)

func (r PaymentSummaryPagedV1DataPaymentType) IsKnown() bool {
	switch r {
	case PaymentSummaryPagedV1DataPaymentTypeCharge, PaymentSummaryPagedV1DataPaymentTypePayout:
		return true
	}
	return false
}

// The current status of the `charge` or `payout`.
type PaymentSummaryPagedV1DataStatus string

const (
	PaymentSummaryPagedV1DataStatusCreated   PaymentSummaryPagedV1DataStatus = "created"
	PaymentSummaryPagedV1DataStatusScheduled PaymentSummaryPagedV1DataStatus = "scheduled"
	PaymentSummaryPagedV1DataStatusFailed    PaymentSummaryPagedV1DataStatus = "failed"
	PaymentSummaryPagedV1DataStatusCancelled PaymentSummaryPagedV1DataStatus = "cancelled"
	PaymentSummaryPagedV1DataStatusOnHold    PaymentSummaryPagedV1DataStatus = "on_hold"
	PaymentSummaryPagedV1DataStatusPending   PaymentSummaryPagedV1DataStatus = "pending"
	PaymentSummaryPagedV1DataStatusPaid      PaymentSummaryPagedV1DataStatus = "paid"
	PaymentSummaryPagedV1DataStatusReversed  PaymentSummaryPagedV1DataStatus = "reversed"
)

func (r PaymentSummaryPagedV1DataStatus) IsKnown() bool {
	switch r {
	case PaymentSummaryPagedV1DataStatusCreated, PaymentSummaryPagedV1DataStatusScheduled, PaymentSummaryPagedV1DataStatusFailed, PaymentSummaryPagedV1DataStatusCancelled, PaymentSummaryPagedV1DataStatusOnHold, PaymentSummaryPagedV1DataStatusPending, PaymentSummaryPagedV1DataStatusPaid, PaymentSummaryPagedV1DataStatusReversed:
		return true
	}
	return false
}

type PaymentSummaryPagedV1Meta struct {
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
	SortBy     string                             `json:"sort_by,required"`
	SortOrder  PaymentSummaryPagedV1MetaSortOrder `json:"sort_order,required"`
	TotalItems int64                              `json:"total_items,required"`
	// The number of pages available.
	TotalPages int64                         `json:"total_pages,required"`
	JSON       paymentSummaryPagedV1MetaJSON `json:"-"`
}

// paymentSummaryPagedV1MetaJSON contains the JSON metadata for the struct
// [PaymentSummaryPagedV1Meta]
type paymentSummaryPagedV1MetaJSON struct {
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

func (r *PaymentSummaryPagedV1Meta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r paymentSummaryPagedV1MetaJSON) RawJSON() string {
	return r.raw
}

type PaymentSummaryPagedV1MetaSortOrder string

const (
	PaymentSummaryPagedV1MetaSortOrderAsc  PaymentSummaryPagedV1MetaSortOrder = "asc"
	PaymentSummaryPagedV1MetaSortOrderDesc PaymentSummaryPagedV1MetaSortOrder = "desc"
)

func (r PaymentSummaryPagedV1MetaSortOrder) IsKnown() bool {
	switch r {
	case PaymentSummaryPagedV1MetaSortOrderAsc, PaymentSummaryPagedV1MetaSortOrderDesc:
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
type PaymentSummaryPagedV1ResponseType string

const (
	PaymentSummaryPagedV1ResponseTypeObject PaymentSummaryPagedV1ResponseType = "object"
	PaymentSummaryPagedV1ResponseTypeArray  PaymentSummaryPagedV1ResponseType = "array"
	PaymentSummaryPagedV1ResponseTypeError  PaymentSummaryPagedV1ResponseType = "error"
	PaymentSummaryPagedV1ResponseTypeNone   PaymentSummaryPagedV1ResponseType = "none"
)

func (r PaymentSummaryPagedV1ResponseType) IsKnown() bool {
	switch r {
	case PaymentSummaryPagedV1ResponseTypeObject, PaymentSummaryPagedV1ResponseTypeArray, PaymentSummaryPagedV1ResponseTypeError, PaymentSummaryPagedV1ResponseTypeNone:
		return true
	}
	return false
}

type PaymentListParams struct {
	// Search using the `customer_id` of a `charge` or `payout`.
	CustomerID      param.Field[string] `query:"customer_id" format:"uuid"`
	DefaultPageSize param.Field[int64]  `query:"default_page_size"`
	// The field to sort the results by.
	DefaultSort      param.Field[PaymentListParamsDefaultSort]      `query:"default_sort"`
	DefaultSortOrder param.Field[PaymentListParamsDefaultSortOrder] `query:"default_sort_order"`
	// Search using the `external_id` of a `charge` or `payout`.
	ExternalID param.Field[string] `query:"external_id"`
	// Search using the `funding_id` of a `charge` or `payout`.
	FundingID param.Field[string] `query:"funding_id" format:"uuid"`
	// Search using a maximum `amount` of a `charge` or `payout`.
	MaxAmount param.Field[int64] `query:"max_amount"`
	// Search using the latest `created_at` date of a `charge` or `payout`.
	MaxCreatedAt param.Field[time.Time] `query:"max_created_at" format:"date-time"`
	// Search using the latest `effective_date` of a `charge` or `payout`.
	MaxEffectiveAt param.Field[time.Time] `query:"max_effective_at" format:"date-time"`
	// Search using the latest `payment_date` of a `charge` or `payout`.
	MaxPaymentDate param.Field[time.Time] `query:"max_payment_date" format:"date"`
	// Search using the minimum `amount of a `charge`or`payout`.
	MinAmount param.Field[int64] `query:"min_amount"`
	// Search using the earliest `created_at` date of a `charge` or `payout`.
	MinCreatedAt param.Field[time.Time] `query:"min_created_at" format:"date-time"`
	// Search using the earliest `effective_date` of a `charge` or `payout`.
	MinEffectiveAt param.Field[time.Time] `query:"min_effective_at" format:"date-time"`
	// Search using the earliest ` `of a `charge` or `payout`.
	MinPaymentDate param.Field[time.Time] `query:"min_payment_date" format:"date"`
	// Results page number. Starts at page 1.
	PageNumber param.Field[int64] `query:"page_number"`
	// Results page size. Max value: 1000
	PageSize param.Field[int64] `query:"page_size"`
	// Search using the `paykey` of a `charge` or `payout`.
	Paykey param.Field[string] `query:"paykey"`
	// Search using the `paykey_id` of a `charge` or `payout`.
	PaykeyID param.Field[string] `query:"paykey_id" format:"uuid"`
	// Search using the `id` of a `charge` or `payout`.
	PaymentID param.Field[string] `query:"payment_id" format:"uuid"`
	// Search by the status of a `charge` or `payout`.
	PaymentStatus param.Field[[]PaymentListParamsPaymentStatus] `query:"payment_status"`
	// Search by the type of a `charge` or `payout`.
	PaymentType param.Field[[]PaymentListParamsPaymentType] `query:"payment_type"`
	// Search using a text string associated with a `charge` or `payout`.
	SearchText param.Field[string] `query:"search_text"`
	// The field to sort the results by.
	SortBy            param.Field[PaymentListParamsSortBy]    `query:"sort_by"`
	SortOrder         param.Field[PaymentListParamsSortOrder] `query:"sort_order"`
	CorrelationID     param.Field[string]                     `header:"Correlation-Id"`
	RequestID         param.Field[string]                     `header:"Request-Id"`
	StraddleAccountID param.Field[string]                     `header:"Straddle-Account-Id" format:"uuid"`
}

// URLQuery serializes [PaymentListParams]'s query parameters as `url.Values`.
func (r PaymentListParams) URLQuery() (v url.Values) {
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

func (r PaymentListParamsDefaultSort) IsKnown() bool {
	switch r {
	case PaymentListParamsDefaultSortCreatedAt, PaymentListParamsDefaultSortPaymentDate, PaymentListParamsDefaultSortEffectiveAt, PaymentListParamsDefaultSortID, PaymentListParamsDefaultSortAmount:
		return true
	}
	return false
}

type PaymentListParamsDefaultSortOrder string

const (
	PaymentListParamsDefaultSortOrderAsc  PaymentListParamsDefaultSortOrder = "asc"
	PaymentListParamsDefaultSortOrderDesc PaymentListParamsDefaultSortOrder = "desc"
)

func (r PaymentListParamsDefaultSortOrder) IsKnown() bool {
	switch r {
	case PaymentListParamsDefaultSortOrderAsc, PaymentListParamsDefaultSortOrderDesc:
		return true
	}
	return false
}

// The current status of the `charge` or `payout`.
type PaymentListParamsPaymentStatus string

const (
	PaymentListParamsPaymentStatusCreated   PaymentListParamsPaymentStatus = "created"
	PaymentListParamsPaymentStatusScheduled PaymentListParamsPaymentStatus = "scheduled"
	PaymentListParamsPaymentStatusFailed    PaymentListParamsPaymentStatus = "failed"
	PaymentListParamsPaymentStatusCancelled PaymentListParamsPaymentStatus = "cancelled"
	PaymentListParamsPaymentStatusOnHold    PaymentListParamsPaymentStatus = "on_hold"
	PaymentListParamsPaymentStatusPending   PaymentListParamsPaymentStatus = "pending"
	PaymentListParamsPaymentStatusPaid      PaymentListParamsPaymentStatus = "paid"
	PaymentListParamsPaymentStatusReversed  PaymentListParamsPaymentStatus = "reversed"
)

func (r PaymentListParamsPaymentStatus) IsKnown() bool {
	switch r {
	case PaymentListParamsPaymentStatusCreated, PaymentListParamsPaymentStatusScheduled, PaymentListParamsPaymentStatusFailed, PaymentListParamsPaymentStatusCancelled, PaymentListParamsPaymentStatusOnHold, PaymentListParamsPaymentStatusPending, PaymentListParamsPaymentStatusPaid, PaymentListParamsPaymentStatusReversed:
		return true
	}
	return false
}

// The type of payment.
type PaymentListParamsPaymentType string

const (
	PaymentListParamsPaymentTypeCharge PaymentListParamsPaymentType = "charge"
	PaymentListParamsPaymentTypePayout PaymentListParamsPaymentType = "payout"
)

func (r PaymentListParamsPaymentType) IsKnown() bool {
	switch r {
	case PaymentListParamsPaymentTypeCharge, PaymentListParamsPaymentTypePayout:
		return true
	}
	return false
}

// The field to sort the results by.
type PaymentListParamsSortBy string

const (
	PaymentListParamsSortByCreatedAt   PaymentListParamsSortBy = "created_at"
	PaymentListParamsSortByPaymentDate PaymentListParamsSortBy = "payment_date"
	PaymentListParamsSortByEffectiveAt PaymentListParamsSortBy = "effective_at"
	PaymentListParamsSortByID          PaymentListParamsSortBy = "id"
	PaymentListParamsSortByAmount      PaymentListParamsSortBy = "amount"
)

func (r PaymentListParamsSortBy) IsKnown() bool {
	switch r {
	case PaymentListParamsSortByCreatedAt, PaymentListParamsSortByPaymentDate, PaymentListParamsSortByEffectiveAt, PaymentListParamsSortByID, PaymentListParamsSortByAmount:
		return true
	}
	return false
}

type PaymentListParamsSortOrder string

const (
	PaymentListParamsSortOrderAsc  PaymentListParamsSortOrder = "asc"
	PaymentListParamsSortOrderDesc PaymentListParamsSortOrder = "desc"
)

func (r PaymentListParamsSortOrder) IsKnown() bool {
	switch r {
	case PaymentListParamsSortOrderAsc, PaymentListParamsSortOrderDesc:
		return true
	}
	return false
}
