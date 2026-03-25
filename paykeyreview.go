// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package straddle

import (
	"context"
	"errors"
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

// Paykeys are secure tokens that link verified customer identities to their bank
// accounts. Each Paykey includes built-in balance checking, fraud detection
// through LSTM machine learning models, and can be reused for subscriptions and
// recurring payments without storing sensitive data. Paykeys eliminate fraud by
// ensuring the person initiating payment owns the funding account.
//
// PaykeyReviewService contains methods and other services that help with
// interacting with the straddle API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPaykeyReviewService] method instead.
type PaykeyReviewService struct {
	options []option.RequestOption
}

// NewPaykeyReviewService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPaykeyReviewService(opts ...option.RequestOption) (r PaykeyReviewService) {
	r = PaykeyReviewService{}
	r.options = opts
	return
}

// Update the status of a paykey when in review status
func (r *PaykeyReviewService) Decision(ctx context.Context, id string, params PaykeyReviewDecisionParams, opts ...option.RequestOption) (res *PaykeyV1, err error) {
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
	path := fmt.Sprintf("v1/paykeys/%s/review", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Get additional details about a paykey.
func (r *PaykeyReviewService) Get(ctx context.Context, id string, query PaykeyReviewGetParams, opts ...option.RequestOption) (res *PaykeyReviewGetResponse, err error) {
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
	path := fmt.Sprintf("v1/paykeys/%s/review", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Updates the decision of a paykey's review validation. This endpoint allows you
// to refresh the outcome of a paykey's decision and is useful for correcting or
// updating the status of a paykey's verification.
func (r *PaykeyReviewService) RefreshReview(ctx context.Context, id string, body PaykeyReviewRefreshReviewParams, opts ...option.RequestOption) (res *PaykeyV1, err error) {
	if !param.IsOmitted(body.CorrelationID) {
		opts = append(opts, option.WithHeader("Correlation-Id", fmt.Sprintf("%v", body.CorrelationID.Value)))
	}
	if !param.IsOmitted(body.IdempotencyKey) {
		opts = append(opts, option.WithHeader("Idempotency-Key", fmt.Sprintf("%v", body.IdempotencyKey.Value)))
	}
	if !param.IsOmitted(body.RequestID) {
		opts = append(opts, option.WithHeader("Request-Id", fmt.Sprintf("%v", body.RequestID.Value)))
	}
	if !param.IsOmitted(body.StraddleAccountID) {
		opts = append(opts, option.WithHeader("Straddle-Account-Id", fmt.Sprintf("%v", body.StraddleAccountID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/paykeys/%s/refresh_review", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, nil, &res, opts...)
	return res, err
}

type PaykeyReviewGetResponse struct {
	Data PaykeyReviewGetResponseData `json:"data" api:"required"`
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
	ResponseType PaykeyReviewGetResponseResponseType `json:"response_type" api:"required"`
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
func (r PaykeyReviewGetResponse) RawJSON() string { return r.JSON.raw }
func (r *PaykeyReviewGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaykeyReviewGetResponseData struct {
	PaykeyDetails       PaykeyReviewGetResponseDataPaykeyDetails       `json:"paykey_details" api:"required"`
	VerificationDetails PaykeyReviewGetResponseDataVerificationDetails `json:"verification_details"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PaykeyDetails       respjson.Field
		VerificationDetails respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaykeyReviewGetResponseData) RawJSON() string { return r.JSON.raw }
func (r *PaykeyReviewGetResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaykeyReviewGetResponseDataPaykeyDetails struct {
	// Unique identifier for the paykey.
	ID     string                                         `json:"id" api:"required" format:"uuid"`
	Config PaykeyReviewGetResponseDataPaykeyDetailsConfig `json:"config" api:"required"`
	// Timestamp of when the paykey was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Human-readable label used to represent this paykey in a UI.
	Label string `json:"label" api:"required"`
	// The tokenized paykey value. This value is used to create payments and should be
	// stored securely.
	Paykey string `json:"paykey" api:"required"`
	// Any of "bank_account", "straddle", "mx", "plaid", "tan", "quiltt".
	Source string `json:"source" api:"required"`
	// Any of "pending", "active", "inactive", "rejected", "review", "blocked".
	Status string `json:"status" api:"required"`
	// Timestamp of the most recent update to the paykey.
	UpdatedAt time.Time                                        `json:"updated_at" api:"required" format:"date-time"`
	Balance   PaykeyReviewGetResponseDataPaykeyDetailsBalance  `json:"balance"`
	BankData  PaykeyReviewGetResponseDataPaykeyDetailsBankData `json:"bank_data"`
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
	Metadata      map[string]string                                     `json:"metadata" api:"nullable"`
	StatusDetails PaykeyReviewGetResponseDataPaykeyDetailsStatusDetails `json:"status_details"`
	// Indicates whether this paykey is eligible for client-initiated unblocking. Only
	// present for blocked paykeys. True when blocked due to R29 returns and not
	// previously unblocked, false otherwise. Null when paykey is not blocked.
	UnblockEligible bool `json:"unblock_eligible" api:"nullable"`
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
		UnblockEligible respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaykeyReviewGetResponseDataPaykeyDetails) RawJSON() string { return r.JSON.raw }
func (r *PaykeyReviewGetResponseDataPaykeyDetails) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaykeyReviewGetResponseDataPaykeyDetailsConfig struct {
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
func (r PaykeyReviewGetResponseDataPaykeyDetailsConfig) RawJSON() string { return r.JSON.raw }
func (r *PaykeyReviewGetResponseDataPaykeyDetailsConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaykeyReviewGetResponseDataPaykeyDetailsBalance struct {
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
func (r PaykeyReviewGetResponseDataPaykeyDetailsBalance) RawJSON() string { return r.JSON.raw }
func (r *PaykeyReviewGetResponseDataPaykeyDetailsBalance) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaykeyReviewGetResponseDataPaykeyDetailsBankData struct {
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
func (r PaykeyReviewGetResponseDataPaykeyDetailsBankData) RawJSON() string { return r.JSON.raw }
func (r *PaykeyReviewGetResponseDataPaykeyDetailsBankData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaykeyReviewGetResponseDataPaykeyDetailsStatusDetails struct {
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
func (r PaykeyReviewGetResponseDataPaykeyDetailsStatusDetails) RawJSON() string { return r.JSON.raw }
func (r *PaykeyReviewGetResponseDataPaykeyDetailsStatusDetails) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaykeyReviewGetResponseDataVerificationDetails struct {
	// Unique identifier for the verification details.
	ID        string                                                  `json:"id" api:"required" format:"uuid"`
	Breakdown PaykeyReviewGetResponseDataVerificationDetailsBreakdown `json:"breakdown" api:"required"`
	// Timestamp of when the verification was initiated.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Any of "accept", "reject", "review".
	Decision string `json:"decision" api:"required"`
	// Dictionary of all messages from the paykey verification process.
	Messages map[string]string `json:"messages" api:"required"`
	// Timestamp of the most recent update to the verification details.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Breakdown   respjson.Field
		CreatedAt   respjson.Field
		Decision    respjson.Field
		Messages    respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaykeyReviewGetResponseDataVerificationDetails) RawJSON() string { return r.JSON.raw }
func (r *PaykeyReviewGetResponseDataVerificationDetails) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaykeyReviewGetResponseDataVerificationDetailsBreakdown struct {
	AccountValidation PaykeyReviewGetResponseDataVerificationDetailsBreakdownAccountValidation `json:"account_validation"`
	NameMatch         PaykeyReviewGetResponseDataVerificationDetailsBreakdownNameMatch         `json:"name_match"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccountValidation respjson.Field
		NameMatch         respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaykeyReviewGetResponseDataVerificationDetailsBreakdown) RawJSON() string { return r.JSON.raw }
func (r *PaykeyReviewGetResponseDataVerificationDetailsBreakdown) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaykeyReviewGetResponseDataVerificationDetailsBreakdownAccountValidation struct {
	Codes []string `json:"codes" api:"required"`
	// Any of "accept", "reject", "review".
	Decision string `json:"decision" api:"required"`
	Reason   string `json:"reason" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Codes       respjson.Field
		Decision    respjson.Field
		Reason      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaykeyReviewGetResponseDataVerificationDetailsBreakdownAccountValidation) RawJSON() string {
	return r.JSON.raw
}
func (r *PaykeyReviewGetResponseDataVerificationDetailsBreakdownAccountValidation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaykeyReviewGetResponseDataVerificationDetailsBreakdownNameMatch struct {
	Codes []string `json:"codes" api:"required"`
	// Any of "accept", "reject", "review".
	Decision         string   `json:"decision" api:"required"`
	CorrelationScore float64  `json:"correlation_score" api:"nullable"`
	CustomerName     string   `json:"customer_name" api:"nullable"`
	MatchedName      string   `json:"matched_name" api:"nullable"`
	NamesOnAccount   []string `json:"names_on_account" api:"nullable"`
	Reason           string   `json:"reason" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Codes            respjson.Field
		Decision         respjson.Field
		CorrelationScore respjson.Field
		CustomerName     respjson.Field
		MatchedName      respjson.Field
		NamesOnAccount   respjson.Field
		Reason           respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaykeyReviewGetResponseDataVerificationDetailsBreakdownNameMatch) RawJSON() string {
	return r.JSON.raw
}
func (r *PaykeyReviewGetResponseDataVerificationDetailsBreakdownNameMatch) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indicates the structure of the returned content.
//
//   - "object" means the `data` field contains a single JSON object.
//   - "array" means the `data` field contains an array of objects.
//   - "error" means the `data` field contains an error object with details of the
//     issue.
//   - "none" means no data is returned.
type PaykeyReviewGetResponseResponseType string

const (
	PaykeyReviewGetResponseResponseTypeObject PaykeyReviewGetResponseResponseType = "object"
	PaykeyReviewGetResponseResponseTypeArray  PaykeyReviewGetResponseResponseType = "array"
	PaykeyReviewGetResponseResponseTypeError  PaykeyReviewGetResponseResponseType = "error"
	PaykeyReviewGetResponseResponseTypeNone   PaykeyReviewGetResponseResponseType = "none"
)

type PaykeyReviewDecisionParams struct {
	// Any of "active", "rejected".
	Status            PaykeyReviewDecisionParamsStatus `json:"status,omitzero" api:"required"`
	CorrelationID     param.Opt[string]                `header:"Correlation-Id,omitzero" json:"-"`
	IdempotencyKey    param.Opt[string]                `header:"Idempotency-Key,omitzero" json:"-"`
	RequestID         param.Opt[string]                `header:"Request-Id,omitzero" json:"-"`
	StraddleAccountID param.Opt[string]                `header:"Straddle-Account-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

func (r PaykeyReviewDecisionParams) MarshalJSON() (data []byte, err error) {
	type shadow PaykeyReviewDecisionParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PaykeyReviewDecisionParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaykeyReviewDecisionParamsStatus string

const (
	PaykeyReviewDecisionParamsStatusActive   PaykeyReviewDecisionParamsStatus = "active"
	PaykeyReviewDecisionParamsStatusRejected PaykeyReviewDecisionParamsStatus = "rejected"
)

type PaykeyReviewGetParams struct {
	CorrelationID     param.Opt[string] `header:"Correlation-Id,omitzero" json:"-"`
	RequestID         param.Opt[string] `header:"Request-Id,omitzero" json:"-"`
	StraddleAccountID param.Opt[string] `header:"Straddle-Account-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

type PaykeyReviewRefreshReviewParams struct {
	CorrelationID     param.Opt[string] `header:"Correlation-Id,omitzero" json:"-"`
	IdempotencyKey    param.Opt[string] `header:"Idempotency-Key,omitzero" json:"-"`
	RequestID         param.Opt[string] `header:"Request-Id,omitzero" json:"-"`
	StraddleAccountID param.Opt[string] `header:"Straddle-Account-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}
