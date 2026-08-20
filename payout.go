// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package straddle

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"slices"
	"time"

	"github.com/straddleio/straddle-go/internal/apiform"
	"github.com/straddleio/straddle-go/internal/apijson"
	"github.com/straddleio/straddle-go/internal/requestconfig"
	"github.com/straddleio/straddle-go/option"
	"github.com/straddleio/straddle-go/packages/param"
	"github.com/straddleio/straddle-go/packages/respjson"
	"github.com/straddleio/straddle-go/shared"
)

// Payouts represent transfers from Straddle to customer bank accounts. Create
// payouts to handle disbursements, process refunds, or manage marketplace
// settlements. Use payouts to send money quickly and securely with the most
// cost-effective rail automatically selected.
//
// PayoutService contains methods and other services that help with interacting
// with the straddle API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPayoutService] method instead.
type PayoutService struct {
	options []option.RequestOption
}

// NewPayoutService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewPayoutService(opts ...option.RequestOption) (r PayoutService) {
	r = PayoutService{}
	r.options = opts
	return
}

// Use payouts to send money to your customers.
func (r *PayoutService) New(ctx context.Context, params PayoutNewParams, opts ...option.RequestOption) (res *PayoutV1, err error) {
	if !param.IsOmitted(params.CorrelationID) {
		opts = append(opts, option.WithHeader("Correlation-Id", fmt.Sprintf("%v", params.CorrelationID.Value)))
	}
	if !param.IsOmitted(params.IdempotencyKey) {
		opts = append(opts, option.WithHeader("Idempotency-Key", fmt.Sprintf("%v", params.IdempotencyKey.Value)))
	}
	if !param.IsOmitted(params.RequestID) {
		opts = append(opts, option.WithHeader("Request-Id", fmt.Sprintf("%v", params.RequestID.Value)))
	}
	if !param.IsOmitted(params.StraddleAccountID) {
		opts = append(opts, option.WithHeader("Straddle-Account-Id", fmt.Sprintf("%v", params.StraddleAccountID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	path := "v1/payouts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Update the details of a payout prior to processing. The status of the payout
// must be `created`, `scheduled`, or `on_hold`.
func (r *PayoutService) Update(ctx context.Context, id string, params PayoutUpdateParams, opts ...option.RequestOption) (res *PayoutV1, err error) {
	if !param.IsOmitted(params.CorrelationID) {
		opts = append(opts, option.WithHeader("Correlation-Id", fmt.Sprintf("%v", params.CorrelationID.Value)))
	}
	if !param.IsOmitted(params.IdempotencyKey) {
		opts = append(opts, option.WithHeader("Idempotency-Key", fmt.Sprintf("%v", params.IdempotencyKey.Value)))
	}
	if !param.IsOmitted(params.RequestID) {
		opts = append(opts, option.WithHeader("Request-Id", fmt.Sprintf("%v", params.RequestID.Value)))
	}
	if !param.IsOmitted(params.StraddleAccountID) {
		opts = append(opts, option.WithHeader("Straddle-Account-Id", fmt.Sprintf("%v", params.StraddleAccountID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/payouts/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return res, err
}

// Cancel a payout to prevent it from being processed. The status of the payout
// must be `created`, `scheduled`, or `on_hold`.
func (r *PayoutService) Cancel(ctx context.Context, id string, params PayoutCancelParams, opts ...option.RequestOption) (res *PayoutV1, err error) {
	if !param.IsOmitted(params.CorrelationID) {
		opts = append(opts, option.WithHeader("Correlation-Id", fmt.Sprintf("%v", params.CorrelationID.Value)))
	}
	if !param.IsOmitted(params.IdempotencyKey) {
		opts = append(opts, option.WithHeader("Idempotency-Key", fmt.Sprintf("%v", params.IdempotencyKey.Value)))
	}
	if !param.IsOmitted(params.RequestID) {
		opts = append(opts, option.WithHeader("Request-Id", fmt.Sprintf("%v", params.RequestID.Value)))
	}
	if !param.IsOmitted(params.StraddleAccountID) {
		opts = append(opts, option.WithHeader("Straddle-Account-Id", fmt.Sprintf("%v", params.StraddleAccountID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/payouts/%s/cancel", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return res, err
}

// Retrieves the details of an existing payout. Supply the unique payout `id` to
// retrieve the corresponding payout information.
func (r *PayoutService) Get(ctx context.Context, id string, query PayoutGetParams, opts ...option.RequestOption) (res *PayoutV1, err error) {
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
	path := fmt.Sprintf("v1/payouts/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Hold a payout to prevent it from being processed. The status of the payout must
// be `created`, `scheduled`, or `on_hold`.
func (r *PayoutService) Hold(ctx context.Context, id string, params PayoutHoldParams, opts ...option.RequestOption) (res *PayoutV1, err error) {
	if !param.IsOmitted(params.CorrelationID) {
		opts = append(opts, option.WithHeader("Correlation-Id", fmt.Sprintf("%v", params.CorrelationID.Value)))
	}
	if !param.IsOmitted(params.IdempotencyKey) {
		opts = append(opts, option.WithHeader("Idempotency-Key", fmt.Sprintf("%v", params.IdempotencyKey.Value)))
	}
	if !param.IsOmitted(params.RequestID) {
		opts = append(opts, option.WithHeader("Request-Id", fmt.Sprintf("%v", params.RequestID.Value)))
	}
	if !param.IsOmitted(params.StraddleAccountID) {
		opts = append(opts, option.WithHeader("Straddle-Account-Id", fmt.Sprintf("%v", params.StraddleAccountID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/payouts/%s/hold", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return res, err
}

// Release a payout from a `hold` status to allow it to be rescheduled for
// processing.
func (r *PayoutService) Release(ctx context.Context, id string, params PayoutReleaseParams, opts ...option.RequestOption) (res *PayoutV1, err error) {
	if !param.IsOmitted(params.CorrelationID) {
		opts = append(opts, option.WithHeader("Correlation-Id", fmt.Sprintf("%v", params.CorrelationID.Value)))
	}
	if !param.IsOmitted(params.IdempotencyKey) {
		opts = append(opts, option.WithHeader("Idempotency-Key", fmt.Sprintf("%v", params.IdempotencyKey.Value)))
	}
	if !param.IsOmitted(params.RequestID) {
		opts = append(opts, option.WithHeader("Request-Id", fmt.Sprintf("%v", params.RequestID.Value)))
	}
	if !param.IsOmitted(params.StraddleAccountID) {
		opts = append(opts, option.WithHeader("Straddle-Account-Id", fmt.Sprintf("%v", params.StraddleAccountID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/payouts/%s/release", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return res, err
}

// Get a payout by id.
func (r *PayoutService) Unmask(ctx context.Context, id string, query PayoutUnmaskParams, opts ...option.RequestOption) (res *PayoutUnmaskResponse, err error) {
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
	path := fmt.Sprintf("v1/payouts/%s/unmask", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Uploads a document as proof of authorization for a payout. Uploading again adds
// another entry to documents rather than replacing the previous one.
func (r *PayoutService) UploadAuthorizationDocument(ctx context.Context, id string, params PayoutUploadAuthorizationDocumentParams, opts ...option.RequestOption) (res *PayoutV1, err error) {
	if !param.IsOmitted(params.CorrelationID) {
		opts = append(opts, option.WithHeader("Correlation-Id", fmt.Sprintf("%v", params.CorrelationID.Value)))
	}
	if !param.IsOmitted(params.IdempotencyKey) {
		opts = append(opts, option.WithHeader("Idempotency-Key", fmt.Sprintf("%v", params.IdempotencyKey.Value)))
	}
	if !param.IsOmitted(params.RequestID) {
		opts = append(opts, option.WithHeader("Request-Id", fmt.Sprintf("%v", params.RequestID.Value)))
	}
	if !param.IsOmitted(params.StraddleAccountID) {
		opts = append(opts, option.WithHeader("Straddle-Account-Id", fmt.Sprintf("%v", params.StraddleAccountID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/payouts/%s/authorization", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

type PayoutV1 struct {
	Data PayoutV1Data `json:"data" api:"required"`
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
	ResponseType PayoutV1ResponseType `json:"response_type" api:"required"`
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
func (r PayoutV1) RawJSON() string { return r.JSON.raw }
func (r *PayoutV1) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PayoutV1Data struct {
	// Unique identifier for the payout.
	ID string `json:"id" api:"required" format:"uuid"`
	// The amount of the payout in cents.
	Amount int64 `json:"amount" api:"required"`
	// Configuration for the payout.
	Config PayoutV1DataConfig `json:"config" api:"required"`
	// The currency of the payout. Only USD is supported.
	Currency string `json:"currency" api:"required"`
	// An arbitrary description for the payout.
	Description string `json:"description" api:"required"`
	// Information about the device used when the customer authorized the payout.
	Device shared.DeviceInfoV1 `json:"device" api:"required"`
	// Unique identifier for the payout in your database. This value must be unique
	// across all payouts.
	ExternalID string `json:"external_id" api:"required"`
	// Funding Ids
	FundingIDs []string `json:"funding_ids" api:"required" format:"uuid"`
	// Has the payout been resubmitted.
	HasResubmit bool `json:"has_resubmit" api:"required"`
	// Is the payout a refund of an original charge.
	IsRefund bool `json:"is_refund" api:"required"`
	// Is the payout a resubmit of an original payout.
	IsResubmit bool `json:"is_resubmit" api:"required"`
	// Value of the `paykey` used for the payout.
	Paykey string `json:"paykey" api:"required"`
	// The desired date on which the payment should be occur. For payouts, this means
	// the date you want the funds to be sent from your bank account.
	PaymentDate time.Time `json:"payment_date" api:"required" format:"date"`
	// The current status of the payout.
	//
	// Any of "created", "scheduled", "failed", "cancelled", "on_hold", "pending",
	// "paid", "reversed", "validating".
	Status string `json:"status" api:"required"`
	// Details about the current status of the payout.
	StatusDetails shared.StatusDetailsV1 `json:"status_details" api:"required"`
	// History of the status changes for the payout.
	StatusHistory []PayoutV1DataStatusHistory `json:"status_history" api:"required"`
	// Trace Ids.
	TraceIDs map[string]string `json:"trace_ids" api:"required"`
	// The time the payout was created.
	CreatedAt time.Time `json:"created_at" api:"nullable" format:"date-time"`
	// Information about the customer associated with the payout.
	CustomerDetails shared.CustomerDetailsV1 `json:"customer_details"`
	// Documents uploaded for this payout (e.g. proof of authorization), in the order
	// they were uploaded.
	Documents []PayoutV1DataDocument `json:"documents" api:"nullable"`
	// The actual date on which the payment occurred. For payouts, this is the date the
	// funds were sent from your bank account.
	EffectiveAt time.Time `json:"effective_at" api:"nullable" format:"date-time"`
	// Up to 20 additional user-defined key-value pairs. Useful for storing additional
	// information about the payout in a structured format.
	Metadata map[string]string `json:"metadata" api:"nullable"`
	// Information about the paykey used for the payout.
	PaykeyDetails shared.PaykeyDetailsV1 `json:"paykey_details"`
	// The payment rail used for the payout.
	//
	// Any of "ach".
	PaymentRail string `json:"payment_rail"`
	// The time the payout was processed by Straddle and originated to the payment
	// rail.
	ProcessedAt time.Time `json:"processed_at" api:"nullable" format:"date-time"`
	// Related payments.
	RelatedPayments []PayoutV1DataRelatedPayment `json:"related_payments" api:"nullable"`
	// The time the payout was last updated.
	UpdatedAt time.Time `json:"updated_at" api:"nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		Amount          respjson.Field
		Config          respjson.Field
		Currency        respjson.Field
		Description     respjson.Field
		Device          respjson.Field
		ExternalID      respjson.Field
		FundingIDs      respjson.Field
		HasResubmit     respjson.Field
		IsRefund        respjson.Field
		IsResubmit      respjson.Field
		Paykey          respjson.Field
		PaymentDate     respjson.Field
		Status          respjson.Field
		StatusDetails   respjson.Field
		StatusHistory   respjson.Field
		TraceIDs        respjson.Field
		CreatedAt       respjson.Field
		CustomerDetails respjson.Field
		Documents       respjson.Field
		EffectiveAt     respjson.Field
		Metadata        respjson.Field
		PaykeyDetails   respjson.Field
		PaymentRail     respjson.Field
		ProcessedAt     respjson.Field
		RelatedPayments respjson.Field
		UpdatedAt       respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PayoutV1Data) RawJSON() string { return r.JSON.raw }
func (r *PayoutV1Data) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configuration for the payout.
type PayoutV1DataConfig struct {
	// Defines whether to automatically place this charge on hold after being created.
	AutoHold bool `json:"auto_hold" api:"nullable"`
	// The reason the payout is being automatically held on creation.
	AutoHoldMessage string `json:"auto_hold_message" api:"nullable"`
	// Payment will simulate processing if not Standard.
	//
	// Any of "standard", "paid", "on_hold_daily_limit", "cancelled_for_fraud_risk",
	// "cancelled_for_balance_check", "failed_insufficient_funds",
	// "reversed_insufficient_funds", "failed_customer_dispute",
	// "reversed_customer_dispute", "failed_closed_bank_account",
	// "reversed_closed_bank_account", "failed_not_authorized",
	// "reversed_not_authorized".
	SandboxOutcome string `json:"sandbox_outcome"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AutoHold        respjson.Field
		AutoHoldMessage respjson.Field
		SandboxOutcome  respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PayoutV1DataConfig) RawJSON() string { return r.JSON.raw }
func (r *PayoutV1DataConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PayoutV1DataStatusHistory struct {
	// The time the status change occurred.
	ChangedAt time.Time `json:"changed_at" api:"required" format:"date-time"`
	// A human-readable description of the status.
	Message string `json:"message" api:"required"`
	// A machine-readable identifier for the specific status, useful for programmatic
	// handling.
	//
	// Any of "insufficient_funds", "closed_bank_account", "invalid_bank_account",
	// "invalid_routing", "disputed", "payment_stopped", "owner_deceased",
	// "frozen_bank_account", "risk_review", "fraudulent", "duplicate_entry",
	// "invalid_paykey", "payment_blocked", "amount_too_large", "too_many_attempts",
	// "internal_system_error", "user_request", "ok", "other_network_return",
	// "payout_refused", "cancel_request", "failed_verification", "require_review",
	// "blocked_by_system", "watchtower_review", "validating", "auto_hold".
	Reason string `json:"reason" api:"required"`
	// Identifies the origin of the status change (e.g., `bank_decline`, `watchtower`).
	// This helps in tracking the cause of status updates.
	//
	// Any of "watchtower", "bank_decline", "customer_dispute", "user_action",
	// "system".
	Source string `json:"source" api:"required"`
	// The current status of the `charge` or `payout`.
	//
	// Any of "created", "scheduled", "failed", "cancelled", "on_hold", "pending",
	// "paid", "reversed", "validating".
	Status string `json:"status" api:"required"`
	// The status code if applicable.
	Code string `json:"code" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChangedAt   respjson.Field
		Message     respjson.Field
		Reason      respjson.Field
		Source      respjson.Field
		Status      respjson.Field
		Code        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PayoutV1DataStatusHistory) RawJSON() string { return r.JSON.raw }
func (r *PayoutV1DataStatusHistory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PayoutV1DataDocument struct {
	// Unique identifier for this document.
	DocumentID string `json:"document_id" api:"required" format:"uuid"`
	// The file name of this document as uploaded.
	DocumentName string `json:"document_name" api:"required"`
	// The size of this document in bytes.
	DocumentSize int64 `json:"document_size" api:"required"`
	// Any of "payment_authorization".
	DocumentType string `json:"document_type" api:"required"`
	// The UTC timestamp when this document was uploaded.
	UploadedAt time.Time `json:"uploaded_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DocumentID   respjson.Field
		DocumentName respjson.Field
		DocumentSize respjson.Field
		DocumentType respjson.Field
		UploadedAt   respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PayoutV1DataDocument) RawJSON() string { return r.JSON.raw }
func (r *PayoutV1DataDocument) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PayoutV1DataRelatedPayment struct {
	// The ID of the related payment.
	ID string `json:"id" api:"required" format:"uuid"`
	// The type of payment.
	//
	// Any of "charge", "payout".
	PaymentType string `json:"payment_type" api:"required"`
	// Any of "original", "resubmit", "refund".
	Relationship string `json:"relationship" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		PaymentType  respjson.Field
		Relationship respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PayoutV1DataRelatedPayment) RawJSON() string { return r.JSON.raw }
func (r *PayoutV1DataRelatedPayment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indicates the structure of the returned content.
//
//   - "object" means the `data` field contains a single JSON object.
//   - "array" means the `data` field contains an array of objects.
//   - "error" means the `data` field contains an error object with details of the
//     issue.
//   - "none" means no data is returned.
type PayoutV1ResponseType string

const (
	PayoutV1ResponseTypeObject PayoutV1ResponseType = "object"
	PayoutV1ResponseTypeArray  PayoutV1ResponseType = "array"
	PayoutV1ResponseTypeError  PayoutV1ResponseType = "error"
	PayoutV1ResponseTypeNone   PayoutV1ResponseType = "none"
)

type PayoutUnmaskResponse struct {
	Data PayoutUnmaskResponseData `json:"data" api:"required"`
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
	ResponseType PayoutUnmaskResponseResponseType `json:"response_type" api:"required"`
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
func (r PayoutUnmaskResponse) RawJSON() string { return r.JSON.raw }
func (r *PayoutUnmaskResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PayoutUnmaskResponseData struct {
	// Id.
	ID string `json:"id" api:"required" format:"uuid"`
	// Amount.
	Amount int64                          `json:"amount" api:"required"`
	Config PayoutUnmaskResponseDataConfig `json:"config" api:"required"`
	// Currency.
	Currency string `json:"currency" api:"required"`
	// Description.
	Description string                         `json:"description" api:"required"`
	Device      PayoutUnmaskResponseDataDevice `json:"device" api:"required"`
	// External id.
	ExternalID string `json:"external_id" api:"required"`
	// Funding Ids
	FundingIDs []string `json:"funding_ids" api:"required" format:"uuid"`
	// Has the payout been resubmitted.
	HasResubmit bool `json:"has_resubmit" api:"required"`
	// Is the payout a refund of an original charge.
	IsRefund bool `json:"is_refund" api:"required"`
	// Is the payout a resubmit of an original payout.
	IsResubmit bool `json:"is_resubmit" api:"required"`
	// Paykey.
	Paykey string `json:"paykey" api:"required"`
	// Payment date.
	PaymentDate time.Time `json:"payment_date" api:"required" format:"date"`
	// The current status of the `charge` or `payout`.
	//
	// Any of "created", "scheduled", "failed", "cancelled", "on_hold", "pending",
	// "paid", "reversed", "validating".
	Status        string                 `json:"status" api:"required"`
	StatusDetails shared.StatusDetailsV1 `json:"status_details" api:"required"`
	// Status history.
	StatusHistory []PayoutUnmaskResponseDataStatusHistory `json:"status_history" api:"required"`
	// Trace Ids.
	TraceIDs map[string]string `json:"trace_ids" api:"required"`
	// Created at.
	CreatedAt time.Time `json:"created_at" api:"nullable" format:"date-time"`
	// Information about the customer associated with the charge or payout.
	CustomerDetails shared.CustomerDetailsV1 `json:"customer_details"`
	// Documents uploaded for this payout (e.g. proof of authorization), in the order
	// they were uploaded.
	Documents []PayoutUnmaskResponseDataDocument `json:"documents" api:"nullable"`
	// Effective at.
	EffectiveAt time.Time `json:"effective_at" api:"nullable" format:"date-time"`
	// Metadata.
	Metadata      map[string]string      `json:"metadata" api:"nullable"`
	PaykeyDetails shared.PaykeyDetailsV1 `json:"paykey_details"`
	// The payment rail used for the charge or payout.
	//
	// Any of "ach".
	PaymentRail string `json:"payment_rail"`
	// Processed at.
	ProcessedAt time.Time `json:"processed_at" api:"nullable" format:"date-time"`
	// Related payments.
	RelatedPayments []PayoutUnmaskResponseDataRelatedPayment `json:"related_payments" api:"nullable"`
	// Updated at.
	UpdatedAt time.Time `json:"updated_at" api:"nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		Amount          respjson.Field
		Config          respjson.Field
		Currency        respjson.Field
		Description     respjson.Field
		Device          respjson.Field
		ExternalID      respjson.Field
		FundingIDs      respjson.Field
		HasResubmit     respjson.Field
		IsRefund        respjson.Field
		IsResubmit      respjson.Field
		Paykey          respjson.Field
		PaymentDate     respjson.Field
		Status          respjson.Field
		StatusDetails   respjson.Field
		StatusHistory   respjson.Field
		TraceIDs        respjson.Field
		CreatedAt       respjson.Field
		CustomerDetails respjson.Field
		Documents       respjson.Field
		EffectiveAt     respjson.Field
		Metadata        respjson.Field
		PaykeyDetails   respjson.Field
		PaymentRail     respjson.Field
		ProcessedAt     respjson.Field
		RelatedPayments respjson.Field
		UpdatedAt       respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PayoutUnmaskResponseData) RawJSON() string { return r.JSON.raw }
func (r *PayoutUnmaskResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PayoutUnmaskResponseDataConfig struct {
	// Defines whether to automatically place this charge on hold after being created.
	AutoHold bool `json:"auto_hold" api:"nullable"`
	// The reason the payout is being automatically held on creation.
	AutoHoldMessage string `json:"auto_hold_message" api:"nullable"`
	// Payment will simulate processing if not Standard.
	//
	// Any of "standard", "paid", "on_hold_daily_limit", "cancelled_for_fraud_risk",
	// "cancelled_for_balance_check", "failed_insufficient_funds",
	// "reversed_insufficient_funds", "failed_customer_dispute",
	// "reversed_customer_dispute", "failed_closed_bank_account",
	// "reversed_closed_bank_account", "failed_not_authorized",
	// "reversed_not_authorized".
	SandboxOutcome string `json:"sandbox_outcome"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AutoHold        respjson.Field
		AutoHoldMessage respjson.Field
		SandboxOutcome  respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PayoutUnmaskResponseDataConfig) RawJSON() string { return r.JSON.raw }
func (r *PayoutUnmaskResponseDataConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PayoutUnmaskResponseDataDevice struct {
	// Ip address.
	IPAddress string `json:"ip_address" api:"required" format:"**.**.**.**"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IPAddress   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PayoutUnmaskResponseDataDevice) RawJSON() string { return r.JSON.raw }
func (r *PayoutUnmaskResponseDataDevice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PayoutUnmaskResponseDataStatusHistory struct {
	// The time the status change occurred.
	ChangedAt time.Time `json:"changed_at" api:"required" format:"date-time"`
	// A human-readable description of the status.
	Message string `json:"message" api:"required"`
	// A machine-readable identifier for the specific status, useful for programmatic
	// handling.
	//
	// Any of "insufficient_funds", "closed_bank_account", "invalid_bank_account",
	// "invalid_routing", "disputed", "payment_stopped", "owner_deceased",
	// "frozen_bank_account", "risk_review", "fraudulent", "duplicate_entry",
	// "invalid_paykey", "payment_blocked", "amount_too_large", "too_many_attempts",
	// "internal_system_error", "user_request", "ok", "other_network_return",
	// "payout_refused", "cancel_request", "failed_verification", "require_review",
	// "blocked_by_system", "watchtower_review", "validating", "auto_hold".
	Reason string `json:"reason" api:"required"`
	// Identifies the origin of the status change (e.g., `bank_decline`, `watchtower`).
	// This helps in tracking the cause of status updates.
	//
	// Any of "watchtower", "bank_decline", "customer_dispute", "user_action",
	// "system".
	Source string `json:"source" api:"required"`
	// The current status of the `charge` or `payout`.
	//
	// Any of "created", "scheduled", "failed", "cancelled", "on_hold", "pending",
	// "paid", "reversed", "validating".
	Status string `json:"status" api:"required"`
	// The status code if applicable.
	Code string `json:"code" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChangedAt   respjson.Field
		Message     respjson.Field
		Reason      respjson.Field
		Source      respjson.Field
		Status      respjson.Field
		Code        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PayoutUnmaskResponseDataStatusHistory) RawJSON() string { return r.JSON.raw }
func (r *PayoutUnmaskResponseDataStatusHistory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PayoutUnmaskResponseDataDocument struct {
	// Unique identifier for this document.
	DocumentID string `json:"document_id" api:"required" format:"uuid"`
	// The file name of this document as uploaded.
	DocumentName string `json:"document_name" api:"required"`
	// The size of this document in bytes.
	DocumentSize int64 `json:"document_size" api:"required"`
	// Any of "payment_authorization".
	DocumentType string `json:"document_type" api:"required"`
	// The UTC timestamp when this document was uploaded.
	UploadedAt time.Time `json:"uploaded_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DocumentID   respjson.Field
		DocumentName respjson.Field
		DocumentSize respjson.Field
		DocumentType respjson.Field
		UploadedAt   respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PayoutUnmaskResponseDataDocument) RawJSON() string { return r.JSON.raw }
func (r *PayoutUnmaskResponseDataDocument) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PayoutUnmaskResponseDataRelatedPayment struct {
	// The ID of the related payment.
	ID string `json:"id" api:"required" format:"uuid"`
	// The type of payment.
	//
	// Any of "charge", "payout".
	PaymentType string `json:"payment_type" api:"required"`
	// Any of "original", "resubmit", "refund".
	Relationship string `json:"relationship" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		PaymentType  respjson.Field
		Relationship respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PayoutUnmaskResponseDataRelatedPayment) RawJSON() string { return r.JSON.raw }
func (r *PayoutUnmaskResponseDataRelatedPayment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indicates the structure of the returned content.
//
//   - "object" means the `data` field contains a single JSON object.
//   - "array" means the `data` field contains an array of objects.
//   - "error" means the `data` field contains an error object with details of the
//     issue.
//   - "none" means no data is returned.
type PayoutUnmaskResponseResponseType string

const (
	PayoutUnmaskResponseResponseTypeObject PayoutUnmaskResponseResponseType = "object"
	PayoutUnmaskResponseResponseTypeArray  PayoutUnmaskResponseResponseType = "array"
	PayoutUnmaskResponseResponseTypeError  PayoutUnmaskResponseResponseType = "error"
	PayoutUnmaskResponseResponseTypeNone   PayoutUnmaskResponseResponseType = "none"
)

type PayoutNewParams struct {
	// An arbitrary description for the payout.
	Description param.Opt[string] `json:"description,omitzero" api:"required"`
	// The amount of the payout in cents.
	Amount int64 `json:"amount" api:"required"`
	// The currency of the payout. Only USD is supported.
	Currency string `json:"currency" api:"required"`
	// Information about the device used when the customer authorized the payout.
	Device shared.DeviceInfoV1Param `json:"device,omitzero" api:"required"`
	// Unique identifier for the payout in your database. This value must be unique
	// across all payouts.
	ExternalID string `json:"external_id" api:"required"`
	// Value of the `paykey` used for the payout.
	Paykey string `json:"paykey" api:"required"`
	// The desired date on which the payout should be occur. For payouts, this means
	// the date you want the funds to be sent from your bank account.
	PaymentDate       time.Time         `json:"payment_date" api:"required" format:"date"`
	CorrelationID     param.Opt[string] `header:"Correlation-Id,omitzero" json:"-"`
	IdempotencyKey    param.Opt[string] `header:"Idempotency-Key,omitzero" json:"-"`
	RequestID         param.Opt[string] `header:"Request-Id,omitzero" json:"-"`
	StraddleAccountID param.Opt[string] `header:"Straddle-Account-Id,omitzero" format:"uuid" json:"-"`
	// Up to 20 additional user-defined key-value pairs. Useful for storing additional
	// information about the payout in a structured format.
	Metadata map[string]string     `json:"metadata,omitzero"`
	Config   PayoutNewParamsConfig `json:"config,omitzero"`
	paramObj
}

func (r PayoutNewParams) MarshalJSON() (data []byte, err error) {
	type shadow PayoutNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PayoutNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PayoutNewParamsConfig struct {
	// Defines whether to automatically place this charge on hold after being created.
	AutoHold param.Opt[bool] `json:"auto_hold,omitzero"`
	// The reason the payout is being automatically held on creation.
	AutoHoldMessage param.Opt[string] `json:"auto_hold_message,omitzero"`
	// Payment will simulate processing if not Standard.
	//
	// Any of "standard", "paid", "on_hold_daily_limit", "cancelled_for_fraud_risk",
	// "cancelled_for_balance_check", "failed_insufficient_funds",
	// "reversed_insufficient_funds", "failed_customer_dispute",
	// "reversed_customer_dispute", "failed_closed_bank_account",
	// "reversed_closed_bank_account", "failed_not_authorized",
	// "reversed_not_authorized".
	SandboxOutcome string `json:"sandbox_outcome,omitzero"`
	paramObj
}

func (r PayoutNewParamsConfig) MarshalJSON() (data []byte, err error) {
	type shadow PayoutNewParamsConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PayoutNewParamsConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PayoutNewParamsConfig](
		"sandbox_outcome", "standard", "paid", "on_hold_daily_limit", "cancelled_for_fraud_risk", "cancelled_for_balance_check", "failed_insufficient_funds", "reversed_insufficient_funds", "failed_customer_dispute", "reversed_customer_dispute", "failed_closed_bank_account", "reversed_closed_bank_account", "failed_not_authorized", "reversed_not_authorized",
	)
}

type PayoutUpdateParams struct {
	// An arbitrary description for the payout.
	Description param.Opt[string] `json:"description,omitzero" api:"required"`
	// The amount of the payout in cents.
	Amount int64 `json:"amount" api:"required"`
	// The desired date on which the payment should be occur. For payouts, this means
	// the date you want the funds to be sent from your bank account.
	PaymentDate       time.Time         `json:"payment_date" api:"required" format:"date"`
	CorrelationID     param.Opt[string] `header:"Correlation-Id,omitzero" json:"-"`
	IdempotencyKey    param.Opt[string] `header:"Idempotency-Key,omitzero" json:"-"`
	RequestID         param.Opt[string] `header:"Request-Id,omitzero" json:"-"`
	StraddleAccountID param.Opt[string] `header:"Straddle-Account-Id,omitzero" format:"uuid" json:"-"`
	// Up to 20 additional user-defined key-value pairs. Useful for storing additional
	// information about the payout in a structured format.
	Metadata map[string]string `json:"metadata,omitzero"`
	paramObj
}

func (r PayoutUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow PayoutUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PayoutUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PayoutCancelParams struct {
	// Details about why the payout status was updated.
	Reason            string            `json:"reason" api:"required"`
	CorrelationID     param.Opt[string] `header:"Correlation-Id,omitzero" json:"-"`
	IdempotencyKey    param.Opt[string] `header:"Idempotency-Key,omitzero" json:"-"`
	RequestID         param.Opt[string] `header:"Request-Id,omitzero" json:"-"`
	StraddleAccountID param.Opt[string] `header:"Straddle-Account-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

func (r PayoutCancelParams) MarshalJSON() (data []byte, err error) {
	type shadow PayoutCancelParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PayoutCancelParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PayoutGetParams struct {
	CorrelationID     param.Opt[string] `header:"Correlation-Id,omitzero" json:"-"`
	RequestID         param.Opt[string] `header:"Request-Id,omitzero" json:"-"`
	StraddleAccountID param.Opt[string] `header:"Straddle-Account-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

type PayoutHoldParams struct {
	// Details about why the payout status was updated.
	Reason            string            `json:"reason" api:"required"`
	CorrelationID     param.Opt[string] `header:"Correlation-Id,omitzero" json:"-"`
	IdempotencyKey    param.Opt[string] `header:"Idempotency-Key,omitzero" json:"-"`
	RequestID         param.Opt[string] `header:"Request-Id,omitzero" json:"-"`
	StraddleAccountID param.Opt[string] `header:"Straddle-Account-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

func (r PayoutHoldParams) MarshalJSON() (data []byte, err error) {
	type shadow PayoutHoldParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PayoutHoldParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PayoutReleaseParams struct {
	// Details about why the payout status was updated.
	Reason            string            `json:"reason" api:"required"`
	CorrelationID     param.Opt[string] `header:"Correlation-Id,omitzero" json:"-"`
	IdempotencyKey    param.Opt[string] `header:"Idempotency-Key,omitzero" json:"-"`
	RequestID         param.Opt[string] `header:"Request-Id,omitzero" json:"-"`
	StraddleAccountID param.Opt[string] `header:"Straddle-Account-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

func (r PayoutReleaseParams) MarshalJSON() (data []byte, err error) {
	type shadow PayoutReleaseParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PayoutReleaseParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PayoutUnmaskParams struct {
	CorrelationID     param.Opt[string] `header:"Correlation-Id,omitzero" json:"-"`
	RequestID         param.Opt[string] `header:"Request-Id,omitzero" json:"-"`
	StraddleAccountID param.Opt[string] `header:"Straddle-Account-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

type PayoutUploadAuthorizationDocumentParams struct {
	// The document file to upload as proof of authorization for this payout.
	File              io.Reader         `json:"File,omitzero" api:"required" format:"binary"`
	CorrelationID     param.Opt[string] `header:"Correlation-Id,omitzero" json:"-"`
	IdempotencyKey    param.Opt[string] `header:"Idempotency-Key,omitzero" json:"-"`
	RequestID         param.Opt[string] `header:"Request-Id,omitzero" json:"-"`
	StraddleAccountID param.Opt[string] `header:"Straddle-Account-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

func (r PayoutUploadAuthorizationDocumentParams) MarshalMultipart() (data []byte, contentType string, err error) {
	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)
	err = apiform.MarshalRoot(r, writer)
	if err == nil {
		err = apiform.WriteExtras(writer, r.ExtraFields())
	}
	if err != nil {
		writer.Close()
		return nil, "", err
	}
	err = writer.Close()
	if err != nil {
		return nil, "", err
	}
	return buf.Bytes(), writer.FormDataContentType(), nil
}
