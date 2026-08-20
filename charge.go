// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package straddle

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/straddleio/straddle-go/internal/apijson"
	"github.com/straddleio/straddle-go/internal/requestconfig"
	"github.com/straddleio/straddle-go/option"
	"github.com/straddleio/straddle-go/packages/param"
	"github.com/straddleio/straddle-go/packages/respjson"
	"github.com/straddleio/straddle-go/shared"
)

// Charges represent attempts to debit money from a customer's bank account using a
// Paykey. Each charge includes automatic balance verification, real-time fraud
// screening, and multi-rail optimization and detailed status tracking throughout
// the payment lifecycle. Use charges to accept bank payments with confidence
// knowing every transaction is protected.
//
// ChargeService contains methods and other services that help with interacting
// with the straddle API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewChargeService] method instead.
type ChargeService struct {
	options []option.RequestOption
}

// NewChargeService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewChargeService(opts ...option.RequestOption) (r ChargeService) {
	r = ChargeService{}
	r.options = opts
	return
}

// Use charges to collect money from a customer for the sale of goods or services.
func (r *ChargeService) New(ctx context.Context, params ChargeNewParams, opts ...option.RequestOption) (res *ChargeV1, err error) {
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
	path := "v1/charges"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Change the values of parameters associated with a charge prior to processing.
// The status of the charge must be `created`, `scheduled`, or `on_hold`.
func (r *ChargeService) Update(ctx context.Context, id string, params ChargeUpdateParams, opts ...option.RequestOption) (res *ChargeV1, err error) {
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
	path := fmt.Sprintf("v1/charges/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return res, err
}

// Cancel a charge to prevent it from being originated for processing. The status
// of the charge must be `created`, `scheduled`, or `on_hold`.
func (r *ChargeService) Cancel(ctx context.Context, id string, params ChargeCancelParams, opts ...option.RequestOption) (res *ChargeV1, err error) {
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
	path := fmt.Sprintf("v1/charges/%s/cancel", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return res, err
}

// Retrieves the details of an existing charge. Supply the unique charge `id`, and
// Straddle will return the corresponding charge information.
func (r *ChargeService) Get(ctx context.Context, id string, query ChargeGetParams, opts ...option.RequestOption) (res *ChargeV1, err error) {
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
	path := fmt.Sprintf("v1/charges/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Place a charge on hold to prevent it from being originated for processing. The
// status of the charge must be `created` or `scheduled`.
func (r *ChargeService) Hold(ctx context.Context, id string, params ChargeHoldParams, opts ...option.RequestOption) (res *ChargeV1, err error) {
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
	path := fmt.Sprintf("v1/charges/%s/hold", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return res, err
}

// Release a charge from an `on_hold` status to allow it to be rescheduled for
// processing.
func (r *ChargeService) Release(ctx context.Context, id string, params ChargeReleaseParams, opts ...option.RequestOption) (res *ChargeV1, err error) {
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
	path := fmt.Sprintf("v1/charges/%s/release", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return res, err
}

// Get a charge by id.
func (r *ChargeService) Unmask(ctx context.Context, id string, query ChargeUnmaskParams, opts ...option.RequestOption) (res *ChargeUnmaskResponse, err error) {
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
	path := fmt.Sprintf("v1/charges/%s/unmask", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type ChargeV1 struct {
	Data ChargeV1Data `json:"data" api:"required"`
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
	ResponseType ChargeV1ResponseType `json:"response_type" api:"required"`
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
func (r ChargeV1) RawJSON() string { return r.JSON.raw }
func (r *ChargeV1) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChargeV1Data struct {
	// Unique identifier for the charge.
	ID string `json:"id" api:"required" format:"uuid"`
	// The amount of the charge in cents.
	Amount int64 `json:"amount" api:"required"`
	// Configuration options for the charge.
	Config ChargeV1DataConfig `json:"config" api:"required"`
	// The channel or mechanism through which the payment was authorized. Use
	// `internet` for payments made online or through a mobile app and `signed` for
	// signed agreements where there is a consent form or contract. Use `signed` for
	// PDF signatures.
	//
	// Any of "internet", "signed".
	ConsentType string `json:"consent_type" api:"required"`
	// Timestamp of when the charge was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// The currency of the charge. Only USD is supported.
	Currency string `json:"currency" api:"required"`
	// An arbitrary description for the charge.
	Description string `json:"description" api:"required"`
	// Information about the device used when the customer authorized the payment.
	Device shared.DeviceInfoV1 `json:"device" api:"required"`
	// Unique identifier for the charge in your database. This value must be unique
	// across all charges.
	ExternalID string `json:"external_id" api:"required"`
	// Funding Ids
	FundingIDs []string `json:"funding_ids" api:"required" format:"uuid"`
	// Has the charge been refunded by an associated payout.
	HasRefund bool `json:"has_refund" api:"required"`
	// Has the charge been resubmitted.
	HasResubmit bool `json:"has_resubmit" api:"required"`
	// Is the charge a resubmit of an original charge.
	IsResubmit bool `json:"is_resubmit" api:"required"`
	// Value of the `paykey` used for the charge.
	Paykey string `json:"paykey" api:"required"`
	// The desired date on which the payment should be occur. For charges, this means
	// the date you want the customer to be debited on.
	PaymentDate time.Time `json:"payment_date" api:"required" format:"date"`
	// The current status of the charge.
	//
	// Any of "created", "scheduled", "failed", "cancelled", "on_hold", "pending",
	// "paid", "reversed", "validating".
	Status string `json:"status" api:"required"`
	// Additional details about the current status of the charge.
	StatusDetails shared.StatusDetailsV1 `json:"status_details" api:"required"`
	// Status history.
	StatusHistory []ChargeV1DataStatusHistory `json:"status_history" api:"required"`
	// Trace Ids.
	TraceIDs map[string]string `json:"trace_ids" api:"required"`
	// Timestamp of when the charge was last updated.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// Information about the customer associated with the charge.
	CustomerDetails shared.CustomerDetailsV1 `json:"customer_details"`
	// Documents uploaded for this charge (e.g. proof of authorization), in the order
	// they were uploaded.
	Documents []ChargeV1DataDocument `json:"documents" api:"nullable"`
	// Timestamp of when the charge was effective in the customer's bank account,
	// otherwise known as the date on which the customer is debited.
	EffectiveAt time.Time `json:"effective_at" api:"nullable" format:"date-time"`
	// Up to 20 additional user-defined key-value pairs. Useful for storing additional
	// information about the charge in a structured format.
	Metadata map[string]string `json:"metadata" api:"nullable"`
	// Information about the paykey used for the charge.
	PaykeyDetails shared.PaykeyDetailsV1 `json:"paykey_details"`
	// The payment rail that the charge will be processed through.
	//
	// Any of "ach".
	PaymentRail string `json:"payment_rail"`
	// Timestamp of when the charge was processed by Straddle and originated to the
	// payment rail.
	ProcessedAt time.Time `json:"processed_at" api:"nullable" format:"date-time"`
	// Related payments.
	RelatedPayments []ChargeV1DataRelatedPayment `json:"related_payments" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		Amount          respjson.Field
		Config          respjson.Field
		ConsentType     respjson.Field
		CreatedAt       respjson.Field
		Currency        respjson.Field
		Description     respjson.Field
		Device          respjson.Field
		ExternalID      respjson.Field
		FundingIDs      respjson.Field
		HasRefund       respjson.Field
		HasResubmit     respjson.Field
		IsResubmit      respjson.Field
		Paykey          respjson.Field
		PaymentDate     respjson.Field
		Status          respjson.Field
		StatusDetails   respjson.Field
		StatusHistory   respjson.Field
		TraceIDs        respjson.Field
		UpdatedAt       respjson.Field
		CustomerDetails respjson.Field
		Documents       respjson.Field
		EffectiveAt     respjson.Field
		Metadata        respjson.Field
		PaykeyDetails   respjson.Field
		PaymentRail     respjson.Field
		ProcessedAt     respjson.Field
		RelatedPayments respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChargeV1Data) RawJSON() string { return r.JSON.raw }
func (r *ChargeV1Data) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configuration options for the charge.
type ChargeV1DataConfig struct {
	// Defines whether to check the customer's balance before processing the charge.
	//
	// Any of "required", "enabled", "disabled".
	BalanceCheck string `json:"balance_check" api:"required"`
	// Defines whether to automatically place this charge on hold after being created.
	AutoHold bool `json:"auto_hold" api:"nullable"`
	// The reason the charge is being automatically held on creation.
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
		BalanceCheck    respjson.Field
		AutoHold        respjson.Field
		AutoHoldMessage respjson.Field
		SandboxOutcome  respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChargeV1DataConfig) RawJSON() string { return r.JSON.raw }
func (r *ChargeV1DataConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A record of the charge's status changes over time.
type ChargeV1DataStatusHistory struct {
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
func (r ChargeV1DataStatusHistory) RawJSON() string { return r.JSON.raw }
func (r *ChargeV1DataStatusHistory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChargeV1DataDocument struct {
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
func (r ChargeV1DataDocument) RawJSON() string { return r.JSON.raw }
func (r *ChargeV1DataDocument) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChargeV1DataRelatedPayment struct {
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
func (r ChargeV1DataRelatedPayment) RawJSON() string { return r.JSON.raw }
func (r *ChargeV1DataRelatedPayment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indicates the structure of the returned content.
//
//   - "object" means the `data` field contains a single JSON object.
//   - "array" means the `data` field contains an array of objects.
//   - "error" means the `data` field contains an error object with details of the
//     issue.
//   - "none" means no data is returned.
type ChargeV1ResponseType string

const (
	ChargeV1ResponseTypeObject ChargeV1ResponseType = "object"
	ChargeV1ResponseTypeArray  ChargeV1ResponseType = "array"
	ChargeV1ResponseTypeError  ChargeV1ResponseType = "error"
	ChargeV1ResponseTypeNone   ChargeV1ResponseType = "none"
)

type ChargeUnmaskResponse struct {
	Data ChargeUnmaskResponseData `json:"data" api:"required"`
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
	ResponseType ChargeUnmaskResponseResponseType `json:"response_type" api:"required"`
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
func (r ChargeUnmaskResponse) RawJSON() string { return r.JSON.raw }
func (r *ChargeUnmaskResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChargeUnmaskResponseData struct {
	// Id.
	ID string `json:"id" api:"required" format:"uuid"`
	// Amount.
	Amount int64                          `json:"amount" api:"required"`
	Config ChargeUnmaskResponseDataConfig `json:"config" api:"required"`
	// The channel or mechanism through which the payment was authorized. Use
	// `internet` for payments made online or through a mobile app and `signed` for
	// signed agreements where there is a consent form or contract. Use `signed` for
	// PDF signatures.
	//
	// Any of "internet", "signed".
	ConsentType string `json:"consent_type" api:"required"`
	// Created at.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Currency.
	Currency string `json:"currency" api:"required"`
	// Description.
	Description string                         `json:"description" api:"required"`
	Device      ChargeUnmaskResponseDataDevice `json:"device" api:"required"`
	// External id.
	ExternalID string `json:"external_id" api:"required"`
	// Funding Ids
	FundingIDs []string `json:"funding_ids" api:"required" format:"uuid"`
	// Has the charge been refunded by an associated payout.
	HasRefund bool `json:"has_refund" api:"required"`
	// Has the charge been resubmitted.
	HasResubmit bool `json:"has_resubmit" api:"required"`
	// Is the charge a resubmit of an original charge.
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
	StatusHistory []ChargeUnmaskResponseDataStatusHistory `json:"status_history" api:"required"`
	// Trace Ids.
	TraceIDs map[string]string `json:"trace_ids" api:"required"`
	// Updated at.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// Information about the customer associated with the charge or payout.
	CustomerDetails shared.CustomerDetailsV1 `json:"customer_details"`
	// Documents uploaded for this charge (e.g. proof of authorization), in the order
	// they were uploaded.
	Documents []ChargeUnmaskResponseDataDocument `json:"documents" api:"nullable"`
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
	RelatedPayments []ChargeUnmaskResponseDataRelatedPayment `json:"related_payments" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		Amount          respjson.Field
		Config          respjson.Field
		ConsentType     respjson.Field
		CreatedAt       respjson.Field
		Currency        respjson.Field
		Description     respjson.Field
		Device          respjson.Field
		ExternalID      respjson.Field
		FundingIDs      respjson.Field
		HasRefund       respjson.Field
		HasResubmit     respjson.Field
		IsResubmit      respjson.Field
		Paykey          respjson.Field
		PaymentDate     respjson.Field
		Status          respjson.Field
		StatusDetails   respjson.Field
		StatusHistory   respjson.Field
		TraceIDs        respjson.Field
		UpdatedAt       respjson.Field
		CustomerDetails respjson.Field
		Documents       respjson.Field
		EffectiveAt     respjson.Field
		Metadata        respjson.Field
		PaykeyDetails   respjson.Field
		PaymentRail     respjson.Field
		ProcessedAt     respjson.Field
		RelatedPayments respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChargeUnmaskResponseData) RawJSON() string { return r.JSON.raw }
func (r *ChargeUnmaskResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChargeUnmaskResponseDataConfig struct {
	// Defines whether to check the customer's balance before processing the charge.
	//
	// Any of "required", "enabled", "disabled".
	BalanceCheck string `json:"balance_check" api:"required"`
	// Defines whether to automatically place this charge on hold after being created.
	AutoHold bool `json:"auto_hold" api:"nullable"`
	// The reason the charge is being automatically held on creation.
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
		BalanceCheck    respjson.Field
		AutoHold        respjson.Field
		AutoHoldMessage respjson.Field
		SandboxOutcome  respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChargeUnmaskResponseDataConfig) RawJSON() string { return r.JSON.raw }
func (r *ChargeUnmaskResponseDataConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChargeUnmaskResponseDataDevice struct {
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
func (r ChargeUnmaskResponseDataDevice) RawJSON() string { return r.JSON.raw }
func (r *ChargeUnmaskResponseDataDevice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChargeUnmaskResponseDataStatusHistory struct {
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
func (r ChargeUnmaskResponseDataStatusHistory) RawJSON() string { return r.JSON.raw }
func (r *ChargeUnmaskResponseDataStatusHistory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChargeUnmaskResponseDataDocument struct {
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
func (r ChargeUnmaskResponseDataDocument) RawJSON() string { return r.JSON.raw }
func (r *ChargeUnmaskResponseDataDocument) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChargeUnmaskResponseDataRelatedPayment struct {
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
func (r ChargeUnmaskResponseDataRelatedPayment) RawJSON() string { return r.JSON.raw }
func (r *ChargeUnmaskResponseDataRelatedPayment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indicates the structure of the returned content.
//
//   - "object" means the `data` field contains a single JSON object.
//   - "array" means the `data` field contains an array of objects.
//   - "error" means the `data` field contains an error object with details of the
//     issue.
//   - "none" means no data is returned.
type ChargeUnmaskResponseResponseType string

const (
	ChargeUnmaskResponseResponseTypeObject ChargeUnmaskResponseResponseType = "object"
	ChargeUnmaskResponseResponseTypeArray  ChargeUnmaskResponseResponseType = "array"
	ChargeUnmaskResponseResponseTypeError  ChargeUnmaskResponseResponseType = "error"
	ChargeUnmaskResponseResponseTypeNone   ChargeUnmaskResponseResponseType = "none"
)

type ChargeNewParams struct {
	// An arbitrary description for the charge.
	Description param.Opt[string] `json:"description,omitzero" api:"required"`
	// The amount of the charge in cents.
	Amount int64                 `json:"amount" api:"required"`
	Config ChargeNewParamsConfig `json:"config,omitzero" api:"required"`
	// The channel or mechanism through which the payment was authorized. Use
	// `internet` for payments made online or through a mobile app and `signed` for
	// signed agreements where there is a consent form or contract. Use `signed` for
	// PDF signatures.
	//
	// Any of "internet", "signed".
	ConsentType ChargeNewParamsConsentType `json:"consent_type,omitzero" api:"required"`
	// The currency of the charge. Only USD is supported.
	Currency string                   `json:"currency" api:"required"`
	Device   shared.DeviceInfoV1Param `json:"device,omitzero" api:"required"`
	// Unique identifier for the charge in your database. This value must be unique
	// across all charges.
	ExternalID string `json:"external_id" api:"required"`
	// Value of the `paykey` used for the charge.
	Paykey string `json:"paykey" api:"required"`
	// The desired date on which the payment should be occur. For charges, this means
	// the date you want the customer to be debited on.
	PaymentDate       time.Time         `json:"payment_date" api:"required" format:"date"`
	CorrelationID     param.Opt[string] `header:"Correlation-Id,omitzero" json:"-"`
	IdempotencyKey    param.Opt[string] `header:"Idempotency-Key,omitzero" json:"-"`
	RequestID         param.Opt[string] `header:"Request-Id,omitzero" json:"-"`
	StraddleAccountID param.Opt[string] `header:"Straddle-Account-Id,omitzero" format:"uuid" json:"-"`
	// Up to 20 additional user-defined key-value pairs. Useful for storing additional
	// information about the charge in a structured format.
	Metadata map[string]string `json:"metadata,omitzero"`
	paramObj
}

func (r ChargeNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ChargeNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChargeNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property BalanceCheck is required.
type ChargeNewParamsConfig struct {
	// Defines whether to check the customer's balance before processing the charge.
	//
	// Any of "required", "enabled", "disabled".
	BalanceCheck string `json:"balance_check,omitzero" api:"required"`
	// Defines whether to automatically place this charge on hold after being created.
	AutoHold param.Opt[bool] `json:"auto_hold,omitzero"`
	// The reason the charge is being automatically held on creation.
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

func (r ChargeNewParamsConfig) MarshalJSON() (data []byte, err error) {
	type shadow ChargeNewParamsConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChargeNewParamsConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ChargeNewParamsConfig](
		"balance_check", "required", "enabled", "disabled",
	)
	apijson.RegisterFieldValidator[ChargeNewParamsConfig](
		"sandbox_outcome", "standard", "paid", "on_hold_daily_limit", "cancelled_for_fraud_risk", "cancelled_for_balance_check", "failed_insufficient_funds", "reversed_insufficient_funds", "failed_customer_dispute", "reversed_customer_dispute", "failed_closed_bank_account", "reversed_closed_bank_account", "failed_not_authorized", "reversed_not_authorized",
	)
}

// The channel or mechanism through which the payment was authorized. Use
// `internet` for payments made online or through a mobile app and `signed` for
// signed agreements where there is a consent form or contract. Use `signed` for
// PDF signatures.
type ChargeNewParamsConsentType string

const (
	ChargeNewParamsConsentTypeInternet ChargeNewParamsConsentType = "internet"
	ChargeNewParamsConsentTypeSigned   ChargeNewParamsConsentType = "signed"
)

type ChargeUpdateParams struct {
	// An arbitrary description for the charge.
	Description param.Opt[string] `json:"description,omitzero" api:"required"`
	// The amount of the charge in cents.
	Amount int64 `json:"amount" api:"required"`
	// The desired date on which the payment should be occur. For charges, this means
	// the date you want the customer to be debited on.
	PaymentDate       time.Time         `json:"payment_date" api:"required" format:"date"`
	CorrelationID     param.Opt[string] `header:"Correlation-Id,omitzero" json:"-"`
	IdempotencyKey    param.Opt[string] `header:"Idempotency-Key,omitzero" json:"-"`
	RequestID         param.Opt[string] `header:"Request-Id,omitzero" json:"-"`
	StraddleAccountID param.Opt[string] `header:"Straddle-Account-Id,omitzero" format:"uuid" json:"-"`
	// Up to 20 additional user-defined key-value pairs. Useful for storing additional
	// information about the charge in a structured format.
	Metadata map[string]string `json:"metadata,omitzero"`
	paramObj
}

func (r ChargeUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow ChargeUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChargeUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChargeCancelParams struct {
	// Details about why the charge status was updated.
	Reason            param.Opt[string] `json:"reason,omitzero"`
	CorrelationID     param.Opt[string] `header:"Correlation-Id,omitzero" json:"-"`
	IdempotencyKey    param.Opt[string] `header:"Idempotency-Key,omitzero" json:"-"`
	RequestID         param.Opt[string] `header:"Request-Id,omitzero" json:"-"`
	StraddleAccountID param.Opt[string] `header:"Straddle-Account-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

func (r ChargeCancelParams) MarshalJSON() (data []byte, err error) {
	type shadow ChargeCancelParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChargeCancelParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChargeGetParams struct {
	CorrelationID     param.Opt[string] `header:"Correlation-Id,omitzero" json:"-"`
	RequestID         param.Opt[string] `header:"Request-Id,omitzero" json:"-"`
	StraddleAccountID param.Opt[string] `header:"Straddle-Account-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

type ChargeHoldParams struct {
	// Details about why the charge status was updated.
	Reason            param.Opt[string] `json:"reason,omitzero"`
	CorrelationID     param.Opt[string] `header:"Correlation-Id,omitzero" json:"-"`
	IdempotencyKey    param.Opt[string] `header:"Idempotency-Key,omitzero" json:"-"`
	RequestID         param.Opt[string] `header:"Request-Id,omitzero" json:"-"`
	StraddleAccountID param.Opt[string] `header:"Straddle-Account-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

func (r ChargeHoldParams) MarshalJSON() (data []byte, err error) {
	type shadow ChargeHoldParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChargeHoldParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChargeReleaseParams struct {
	// Details about why the charge status was updated.
	Reason            param.Opt[string] `json:"reason,omitzero"`
	CorrelationID     param.Opt[string] `header:"Correlation-Id,omitzero" json:"-"`
	IdempotencyKey    param.Opt[string] `header:"Idempotency-Key,omitzero" json:"-"`
	RequestID         param.Opt[string] `header:"Request-Id,omitzero" json:"-"`
	StraddleAccountID param.Opt[string] `header:"Straddle-Account-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

func (r ChargeReleaseParams) MarshalJSON() (data []byte, err error) {
	type shadow ChargeReleaseParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChargeReleaseParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChargeUnmaskParams struct {
	CorrelationID     param.Opt[string] `header:"Correlation-Id,omitzero" json:"-"`
	RequestID         param.Opt[string] `header:"Request-Id,omitzero" json:"-"`
	StraddleAccountID param.Opt[string] `header:"Straddle-Account-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}
