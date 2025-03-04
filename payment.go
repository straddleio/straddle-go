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
	Data []PaymentSummaryPagedV1Data   `json:"data,required"`
	Meta shared.PagedResponseMetadata2 `json:"meta,required"`
	// Indicates the structure of the returned content.
	//
	//   - "object" means the `data` field contains a single JSON object.
	//   - "array" means the `data` field contains an array of objects.
	//   - "error" means the `data` field contains an error object with details of the
	//     issue.
	//   - "none" means no data is returned.
	ResponseType shared.ResponseTypeEnum   `json:"response_type,required"`
	JSON         paymentSummaryPagedV1JSON `json:"-"`
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
	// Value of the `paykey` used for the `charge` or `payout`.
	Paykey string `json:"paykey,required"`
	// The desired date on which the payment should be occur. For charges, this means
	// the date you want the customer to be debited on. For payouts, this means the
	// date you want the funds to be sent from your bank account.
	PaymentDate time.Time `json:"payment_date,required" format:"date"`
	// The type of payment. Valid values are `charge` or `payout`.
	PaymentType shared.PaymentTypeV1 `json:"payment_type,required"`
	// The current status of the `charge` or `payout`.
	Status shared.PaymentStatusV1 `json:"status,required"`
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

type PaymentListParams struct {
	// Search using the `customer_id` of a `charge` or `payout`.
	CustomerID      param.Field[string] `query:"customer_id" format:"uuid"`
	DefaultPageSize param.Field[int64]  `query:"default_page_size"`
	// The field to sort the results by.
	DefaultSort      param.Field[shared.PaymentSortByV1] `query:"default_sort"`
	DefaultSortOrder param.Field[shared.SortOrder]       `query:"default_sort_order"`
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
	PaymentStatus param.Field[[]shared.PaymentStatusV1] `query:"payment_status"`
	// Search by the type of a `charge` or `payout`.
	PaymentType param.Field[[]shared.PaymentTypeV1] `query:"payment_type"`
	// Search using a text string associated with a `charge` or `payout`.
	SearchText param.Field[string] `query:"search_text"`
	// The field to sort the results by.
	SortBy            param.Field[shared.PaymentSortByV1] `query:"sort_by"`
	SortOrder         param.Field[shared.SortOrder]       `query:"sort_order"`
	CorrelationID     param.Field[string]                 `header:"Correlation-Id"`
	RequestID         param.Field[string]                 `header:"Request-Id"`
	StraddleAccountID param.Field[string]                 `header:"Straddle-Account-Id" format:"uuid"`
}

// URLQuery serializes [PaymentListParams]'s query parameters as `url.Values`.
func (r PaymentListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
