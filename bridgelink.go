// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package straddle

import (
	"context"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/stainless-sdks/straddle-go/internal/apijson"
	"github.com/stainless-sdks/straddle-go/internal/requestconfig"
	"github.com/stainless-sdks/straddle-go/option"
	"github.com/stainless-sdks/straddle-go/packages/param"
	"github.com/stainless-sdks/straddle-go/packages/respjson"
	"github.com/stainless-sdks/straddle-go/shared"
)

// Bridge provides a comprehensive suite of tools for connecting customer bank
// accounts. Use it to generate secure widget sessions for instant account
// verification, accept tokens from major providers like Plaid and Finicity, or
// verify accounts directly via our API. Bridge handles all sensitive banking
// credentials and ensures secure, compliant connections with support for 90% of US
// bank accounts.
//
// BridgeLinkService contains methods and other services that help with interacting
// with the straddle API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBridgeLinkService] method instead.
type BridgeLinkService struct {
	options []option.RequestOption
}

// NewBridgeLinkService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewBridgeLinkService(opts ...option.RequestOption) (r BridgeLinkService) {
	r = BridgeLinkService{}
	r.options = opts
	return
}

// Use Bridge to create a new paykey using a bank routing and account number as the
// source. This endpoint allows you to create a secure payment token linked to a
// specific bank account.
func (r *BridgeLinkService) BankAccount(ctx context.Context, params BridgeLinkBankAccountParams, opts ...option.RequestOption) (res *PaykeyV1, err error) {
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
	path := "v1/bridge/bank_account"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Creates a new paykey using a Quiltt token as the source. This endpoint allows
// you to create a secure payment token linked to a bank account authenticated
// through Quiltt.
func (r *BridgeLinkService) NewPaykey(ctx context.Context, params BridgeLinkNewPaykeyParams, opts ...option.RequestOption) (res *BridgeLinkNewPaykeyResponse, err error) {
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
	path := "v1/bridge/quiltt"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

func (r *BridgeLinkService) NewTan(ctx context.Context, params BridgeLinkNewTanParams, opts ...option.RequestOption) (res *BridgeLinkNewTanResponse, err error) {
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
	path := "v1/bridge/tan"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Use Bridge to create a new paykey using a Plaid token as the source. This
// endpoint allows you to create a secure payment token linked to a bank account
// authenticated through Plaid.
func (r *BridgeLinkService) Plaid(ctx context.Context, params BridgeLinkPlaidParams, opts ...option.RequestOption) (res *PaykeyV1, err error) {
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
	path := "v1/bridge/plaid"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

type BridgeLinkNewPaykeyResponse struct {
	Data BridgeLinkNewPaykeyResponseData `json:"data" api:"required"`
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
	ResponseType BridgeLinkNewPaykeyResponseResponseType `json:"response_type" api:"required"`
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
func (r BridgeLinkNewPaykeyResponse) RawJSON() string { return r.JSON.raw }
func (r *BridgeLinkNewPaykeyResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BridgeLinkNewPaykeyResponseData struct {
	// Unique identifier for the paykey.
	ID     string                                `json:"id" api:"required" format:"uuid"`
	Config BridgeLinkNewPaykeyResponseDataConfig `json:"config" api:"required"`
	// Timestamp of when the paykey was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Human-readable label that combines the bank name and masked account number to
	// help easility represent this paykey in a UI
	Label string `json:"label" api:"required"`
	// The tokenized paykey value. This token is used to create payments and should be
	// stored securely.
	Paykey string `json:"paykey" api:"required"`
	// Any of "bank_account", "straddle", "mx", "plaid", "tan", "quiltt".
	Source string `json:"source" api:"required"`
	// Any of "pending", "active", "inactive", "rejected", "review", "blocked".
	Status string `json:"status" api:"required"`
	// Timestamp of the most recent update to the paykey.
	UpdatedAt time.Time                               `json:"updated_at" api:"required" format:"date-time"`
	Balance   BridgeLinkNewPaykeyResponseDataBalance  `json:"balance"`
	BankData  BridgeLinkNewPaykeyResponseDataBankData `json:"bank_data"`
	// Unique identifier of the related customer object.
	CustomerID string `json:"customer_id" api:"nullable" format:"uuid"`
	// Expiration date and time of the paykey, if applicable.
	ExpiresAt time.Time `json:"expires_at" api:"nullable" format:"date-time"`
	// Unique identifier for the paykey in your database, used for cross-referencing
	// between Straddle and your systems.
	ExternalID string `json:"external_id" api:"nullable"`
	// Name of the financial institution.
	InstitutionName string `json:"institution_name" api:"nullable"`
	// Up to 20 additional user-defined key-value pairs. Useful for storing additional
	// information about the paykey in a structured format.
	Metadata      map[string]string                            `json:"metadata" api:"nullable"`
	StatusDetails BridgeLinkNewPaykeyResponseDataStatusDetails `json:"status_details"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		Config          respjson.Field
		CreatedAt       respjson.Field
		Label           respjson.Field
		Paykey          respjson.Field
		Source          respjson.Field
		Status          respjson.Field
		UpdatedAt       respjson.Field
		Balance         respjson.Field
		BankData        respjson.Field
		CustomerID      respjson.Field
		ExpiresAt       respjson.Field
		ExternalID      respjson.Field
		InstitutionName respjson.Field
		Metadata        respjson.Field
		StatusDetails   respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BridgeLinkNewPaykeyResponseData) RawJSON() string { return r.JSON.raw }
func (r *BridgeLinkNewPaykeyResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BridgeLinkNewPaykeyResponseDataConfig struct {
	// Any of "inline", "background", "skip".
	ProcessingMethod string `json:"processing_method"`
	// Any of "standard", "active", "rejected", "review".
	SandboxOutcome string `json:"sandbox_outcome"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ProcessingMethod respjson.Field
		SandboxOutcome   respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BridgeLinkNewPaykeyResponseDataConfig) RawJSON() string { return r.JSON.raw }
func (r *BridgeLinkNewPaykeyResponseDataConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BridgeLinkNewPaykeyResponseDataBalance struct {
	// Any of "pending", "completed", "failed".
	Status string `json:"status" api:"required"`
	// Account Balance when last retrieved
	AccountBalance int64 `json:"account_balance" api:"nullable"`
	// Last time account balance was updated.
	UpdatedAt time.Time `json:"updated_at" api:"nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Status         respjson.Field
		AccountBalance respjson.Field
		UpdatedAt      respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BridgeLinkNewPaykeyResponseDataBalance) RawJSON() string { return r.JSON.raw }
func (r *BridgeLinkNewPaykeyResponseDataBalance) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BridgeLinkNewPaykeyResponseDataBankData struct {
	// Bank account number. This value is masked by default for security reasons. Use
	// the /unmask endpoint to access the unmasked value.
	AccountNumber string `json:"account_number" api:"required"`
	// Any of "checking", "savings".
	AccountType string `json:"account_type" api:"required"`
	// The routing number of the bank account.
	RoutingNumber string `json:"routing_number" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccountNumber respjson.Field
		AccountType   respjson.Field
		RoutingNumber respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BridgeLinkNewPaykeyResponseDataBankData) RawJSON() string { return r.JSON.raw }
func (r *BridgeLinkNewPaykeyResponseDataBankData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BridgeLinkNewPaykeyResponseDataStatusDetails struct {
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
func (r BridgeLinkNewPaykeyResponseDataStatusDetails) RawJSON() string { return r.JSON.raw }
func (r *BridgeLinkNewPaykeyResponseDataStatusDetails) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indicates the structure of the returned content.
//
//   - "object" means the `data` field contains a single JSON object.
//   - "array" means the `data` field contains an array of objects.
//   - "error" means the `data` field contains an error object with details of the
//     issue.
//   - "none" means no data is returned.
type BridgeLinkNewPaykeyResponseResponseType string

const (
	BridgeLinkNewPaykeyResponseResponseTypeObject BridgeLinkNewPaykeyResponseResponseType = "object"
	BridgeLinkNewPaykeyResponseResponseTypeArray  BridgeLinkNewPaykeyResponseResponseType = "array"
	BridgeLinkNewPaykeyResponseResponseTypeError  BridgeLinkNewPaykeyResponseResponseType = "error"
	BridgeLinkNewPaykeyResponseResponseTypeNone   BridgeLinkNewPaykeyResponseResponseType = "none"
)

type BridgeLinkNewTanResponse struct {
	Data BridgeLinkNewTanResponseData `json:"data" api:"required"`
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
	ResponseType BridgeLinkNewTanResponseResponseType `json:"response_type" api:"required"`
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
func (r BridgeLinkNewTanResponse) RawJSON() string { return r.JSON.raw }
func (r *BridgeLinkNewTanResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BridgeLinkNewTanResponseData struct {
	// Unique identifier for the paykey.
	ID     string                             `json:"id" api:"required" format:"uuid"`
	Config BridgeLinkNewTanResponseDataConfig `json:"config" api:"required"`
	// Timestamp of when the paykey was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Human-readable label that combines the bank name and masked account number to
	// help easility represent this paykey in a UI
	Label string `json:"label" api:"required"`
	// The tokenized paykey value. This token is used to create payments and should be
	// stored securely.
	Paykey string `json:"paykey" api:"required"`
	// Any of "bank_account", "straddle", "mx", "plaid", "tan", "quiltt".
	Source string `json:"source" api:"required"`
	// Any of "pending", "active", "inactive", "rejected", "review", "blocked".
	Status string `json:"status" api:"required"`
	// Timestamp of the most recent update to the paykey.
	UpdatedAt time.Time                            `json:"updated_at" api:"required" format:"date-time"`
	Balance   BridgeLinkNewTanResponseDataBalance  `json:"balance"`
	BankData  BridgeLinkNewTanResponseDataBankData `json:"bank_data"`
	// Unique identifier of the related customer object.
	CustomerID string `json:"customer_id" api:"nullable" format:"uuid"`
	// Expiration date and time of the paykey, if applicable.
	ExpiresAt time.Time `json:"expires_at" api:"nullable" format:"date-time"`
	// Unique identifier for the paykey in your database, used for cross-referencing
	// between Straddle and your systems.
	ExternalID string `json:"external_id" api:"nullable"`
	// Name of the financial institution.
	InstitutionName string `json:"institution_name" api:"nullable"`
	// Up to 20 additional user-defined key-value pairs. Useful for storing additional
	// information about the paykey in a structured format.
	Metadata      map[string]string                         `json:"metadata" api:"nullable"`
	StatusDetails BridgeLinkNewTanResponseDataStatusDetails `json:"status_details"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		Config          respjson.Field
		CreatedAt       respjson.Field
		Label           respjson.Field
		Paykey          respjson.Field
		Source          respjson.Field
		Status          respjson.Field
		UpdatedAt       respjson.Field
		Balance         respjson.Field
		BankData        respjson.Field
		CustomerID      respjson.Field
		ExpiresAt       respjson.Field
		ExternalID      respjson.Field
		InstitutionName respjson.Field
		Metadata        respjson.Field
		StatusDetails   respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BridgeLinkNewTanResponseData) RawJSON() string { return r.JSON.raw }
func (r *BridgeLinkNewTanResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BridgeLinkNewTanResponseDataConfig struct {
	// Any of "inline", "background", "skip".
	ProcessingMethod string `json:"processing_method"`
	// Any of "standard", "active", "rejected", "review".
	SandboxOutcome string `json:"sandbox_outcome"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ProcessingMethod respjson.Field
		SandboxOutcome   respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BridgeLinkNewTanResponseDataConfig) RawJSON() string { return r.JSON.raw }
func (r *BridgeLinkNewTanResponseDataConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BridgeLinkNewTanResponseDataBalance struct {
	// Any of "pending", "completed", "failed".
	Status string `json:"status" api:"required"`
	// Account Balance when last retrieved
	AccountBalance int64 `json:"account_balance" api:"nullable"`
	// Last time account balance was updated.
	UpdatedAt time.Time `json:"updated_at" api:"nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Status         respjson.Field
		AccountBalance respjson.Field
		UpdatedAt      respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BridgeLinkNewTanResponseDataBalance) RawJSON() string { return r.JSON.raw }
func (r *BridgeLinkNewTanResponseDataBalance) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BridgeLinkNewTanResponseDataBankData struct {
	// Bank account number. This value is masked by default for security reasons. Use
	// the /unmask endpoint to access the unmasked value.
	AccountNumber string `json:"account_number" api:"required"`
	// Any of "checking", "savings".
	AccountType string `json:"account_type" api:"required"`
	// The routing number of the bank account.
	RoutingNumber string `json:"routing_number" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccountNumber respjson.Field
		AccountType   respjson.Field
		RoutingNumber respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BridgeLinkNewTanResponseDataBankData) RawJSON() string { return r.JSON.raw }
func (r *BridgeLinkNewTanResponseDataBankData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BridgeLinkNewTanResponseDataStatusDetails struct {
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
func (r BridgeLinkNewTanResponseDataStatusDetails) RawJSON() string { return r.JSON.raw }
func (r *BridgeLinkNewTanResponseDataStatusDetails) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indicates the structure of the returned content.
//
//   - "object" means the `data` field contains a single JSON object.
//   - "array" means the `data` field contains an array of objects.
//   - "error" means the `data` field contains an error object with details of the
//     issue.
//   - "none" means no data is returned.
type BridgeLinkNewTanResponseResponseType string

const (
	BridgeLinkNewTanResponseResponseTypeObject BridgeLinkNewTanResponseResponseType = "object"
	BridgeLinkNewTanResponseResponseTypeArray  BridgeLinkNewTanResponseResponseType = "array"
	BridgeLinkNewTanResponseResponseTypeError  BridgeLinkNewTanResponseResponseType = "error"
	BridgeLinkNewTanResponseResponseTypeNone   BridgeLinkNewTanResponseResponseType = "none"
)

type BridgeLinkBankAccountParams struct {
	// The bank account number.
	AccountNumber string `json:"account_number" api:"required"`
	// Any of "checking", "savings".
	AccountType BridgeLinkBankAccountParamsAccountType `json:"account_type,omitzero" api:"required"`
	// Unique identifier of the related customer object.
	CustomerID string `json:"customer_id" api:"required" format:"uuid"`
	// The routing number of the bank account.
	RoutingNumber string `json:"routing_number" api:"required"`
	// Unique identifier for the paykey in your database, used for cross-referencing
	// between Straddle and your systems.
	ExternalID        param.Opt[string] `json:"external_id,omitzero"`
	CorrelationID     param.Opt[string] `header:"Correlation-Id,omitzero" json:"-"`
	IdempotencyKey    param.Opt[string] `header:"Idempotency-Key,omitzero" json:"-"`
	RequestID         param.Opt[string] `header:"Request-Id,omitzero" json:"-"`
	StraddleAccountID param.Opt[string] `header:"Straddle-Account-Id,omitzero" format:"uuid" json:"-"`
	// Up to 20 additional user-defined key-value pairs. Useful for storing additional
	// information about the paykey in a structured format.
	Metadata map[string]string                 `json:"metadata,omitzero"`
	Config   BridgeLinkBankAccountParamsConfig `json:"config,omitzero"`
	paramObj
}

func (r BridgeLinkBankAccountParams) MarshalJSON() (data []byte, err error) {
	type shadow BridgeLinkBankAccountParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BridgeLinkBankAccountParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BridgeLinkBankAccountParamsAccountType string

const (
	BridgeLinkBankAccountParamsAccountTypeChecking BridgeLinkBankAccountParamsAccountType = "checking"
	BridgeLinkBankAccountParamsAccountTypeSavings  BridgeLinkBankAccountParamsAccountType = "savings"
)

type BridgeLinkBankAccountParamsConfig struct {
	// Any of "inline", "background", "skip".
	ProcessingMethod string `json:"processing_method,omitzero"`
	// Any of "standard", "active", "rejected", "review".
	SandboxOutcome string `json:"sandbox_outcome,omitzero"`
	paramObj
}

func (r BridgeLinkBankAccountParamsConfig) MarshalJSON() (data []byte, err error) {
	type shadow BridgeLinkBankAccountParamsConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BridgeLinkBankAccountParamsConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[BridgeLinkBankAccountParamsConfig](
		"processing_method", "inline", "background", "skip",
	)
	apijson.RegisterFieldValidator[BridgeLinkBankAccountParamsConfig](
		"sandbox_outcome", "standard", "active", "rejected", "review",
	)
}

type BridgeLinkNewPaykeyParams struct {
	// Unique identifier of the related customer object.
	CustomerID string `json:"customer_id" api:"required" format:"uuid"`
	// Quiltt processor token generated by your application for use with the Straddle
	// API.
	QuilttToken string `json:"quiltt_token" api:"required"`
	// Unique identifier for the paykey in your database, used for cross-referencing
	// between Straddle and your systems.
	ExternalID        param.Opt[string] `json:"external_id,omitzero"`
	CorrelationID     param.Opt[string] `header:"Correlation-Id,omitzero" json:"-"`
	IdempotencyKey    param.Opt[string] `header:"Idempotency-Key,omitzero" json:"-"`
	RequestID         param.Opt[string] `header:"Request-Id,omitzero" json:"-"`
	StraddleAccountID param.Opt[string] `header:"Straddle-Account-Id,omitzero" format:"uuid" json:"-"`
	// Up to 20 additional user-defined key-value pairs. Useful for storing additional
	// information about the paykey in a structured format.
	Metadata map[string]string               `json:"metadata,omitzero"`
	Config   BridgeLinkNewPaykeyParamsConfig `json:"config,omitzero"`
	paramObj
}

func (r BridgeLinkNewPaykeyParams) MarshalJSON() (data []byte, err error) {
	type shadow BridgeLinkNewPaykeyParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BridgeLinkNewPaykeyParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BridgeLinkNewPaykeyParamsConfig struct {
	// Any of "inline", "background", "skip".
	ProcessingMethod string `json:"processing_method,omitzero"`
	// Any of "standard", "active", "rejected", "review".
	SandboxOutcome string `json:"sandbox_outcome,omitzero"`
	paramObj
}

func (r BridgeLinkNewPaykeyParamsConfig) MarshalJSON() (data []byte, err error) {
	type shadow BridgeLinkNewPaykeyParamsConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BridgeLinkNewPaykeyParamsConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[BridgeLinkNewPaykeyParamsConfig](
		"processing_method", "inline", "background", "skip",
	)
	apijson.RegisterFieldValidator[BridgeLinkNewPaykeyParamsConfig](
		"sandbox_outcome", "standard", "active", "rejected", "review",
	)
}

type BridgeLinkNewTanParams struct {
	// Any of "checking", "savings".
	AccountType BridgeLinkNewTanParamsAccountType `json:"account_type,omitzero" api:"required"`
	// Unique identifier of the related customer object.
	CustomerID string `json:"customer_id" api:"required" format:"uuid"`
	// Bank routing number.
	RoutingNumber string `json:"routing_number" api:"required"`
	// Tokenized account number.
	Tan string `json:"tan" api:"required"`
	// Unique identifier for the paykey in your database, used for cross-referencing
	// between Straddle and your systems.
	ExternalID        param.Opt[string] `json:"external_id,omitzero"`
	CorrelationID     param.Opt[string] `header:"Correlation-Id,omitzero" json:"-"`
	IdempotencyKey    param.Opt[string] `header:"Idempotency-Key,omitzero" json:"-"`
	RequestID         param.Opt[string] `header:"Request-Id,omitzero" json:"-"`
	StraddleAccountID param.Opt[string] `header:"Straddle-Account-Id,omitzero" format:"uuid" json:"-"`
	// Up to 20 additional user-defined key-value pairs. Useful for storing additional
	// information about the paykey in a structured format.
	Metadata map[string]string            `json:"metadata,omitzero"`
	Config   BridgeLinkNewTanParamsConfig `json:"config,omitzero"`
	paramObj
}

func (r BridgeLinkNewTanParams) MarshalJSON() (data []byte, err error) {
	type shadow BridgeLinkNewTanParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BridgeLinkNewTanParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BridgeLinkNewTanParamsAccountType string

const (
	BridgeLinkNewTanParamsAccountTypeChecking BridgeLinkNewTanParamsAccountType = "checking"
	BridgeLinkNewTanParamsAccountTypeSavings  BridgeLinkNewTanParamsAccountType = "savings"
)

type BridgeLinkNewTanParamsConfig struct {
	// Any of "inline", "background", "skip".
	ProcessingMethod string `json:"processing_method,omitzero"`
	// Any of "standard", "active", "rejected", "review".
	SandboxOutcome string `json:"sandbox_outcome,omitzero"`
	paramObj
}

func (r BridgeLinkNewTanParamsConfig) MarshalJSON() (data []byte, err error) {
	type shadow BridgeLinkNewTanParamsConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BridgeLinkNewTanParamsConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[BridgeLinkNewTanParamsConfig](
		"processing_method", "inline", "background", "skip",
	)
	apijson.RegisterFieldValidator[BridgeLinkNewTanParamsConfig](
		"sandbox_outcome", "standard", "active", "rejected", "review",
	)
}

type BridgeLinkPlaidParams struct {
	// Unique identifier of the related customer object.
	CustomerID string `json:"customer_id" api:"required" format:"uuid"`
	// Plaid processor token generated by your application for use with the Straddle
	// API.
	PlaidToken string `json:"plaid_token" api:"required"`
	// Unique identifier for the paykey in your database, used for cross-referencing
	// between Straddle and your systems.
	ExternalID        param.Opt[string] `json:"external_id,omitzero"`
	CorrelationID     param.Opt[string] `header:"Correlation-Id,omitzero" json:"-"`
	IdempotencyKey    param.Opt[string] `header:"Idempotency-Key,omitzero" json:"-"`
	RequestID         param.Opt[string] `header:"Request-Id,omitzero" json:"-"`
	StraddleAccountID param.Opt[string] `header:"Straddle-Account-Id,omitzero" format:"uuid" json:"-"`
	// Up to 20 additional user-defined key-value pairs. Useful for storing additional
	// information about the paykey in a structured format.
	Metadata map[string]string           `json:"metadata,omitzero"`
	Config   BridgeLinkPlaidParamsConfig `json:"config,omitzero"`
	paramObj
}

func (r BridgeLinkPlaidParams) MarshalJSON() (data []byte, err error) {
	type shadow BridgeLinkPlaidParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BridgeLinkPlaidParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BridgeLinkPlaidParamsConfig struct {
	// Any of "inline", "background", "skip".
	ProcessingMethod string `json:"processing_method,omitzero"`
	// Any of "standard", "active", "rejected", "review".
	SandboxOutcome string `json:"sandbox_outcome,omitzero"`
	paramObj
}

func (r BridgeLinkPlaidParamsConfig) MarshalJSON() (data []byte, err error) {
	type shadow BridgeLinkPlaidParamsConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BridgeLinkPlaidParamsConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[BridgeLinkPlaidParamsConfig](
		"processing_method", "inline", "background", "skip",
	)
	apijson.RegisterFieldValidator[BridgeLinkPlaidParamsConfig](
		"sandbox_outcome", "standard", "active", "rejected", "review",
	)
}
