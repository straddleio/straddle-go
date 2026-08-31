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

// Paykeys are secure tokens that link verified customer identities to their bank
// accounts. Each Paykey includes built-in balance checking, fraud detection
// through LSTM machine learning models, and can be reused for subscriptions and
// recurring payments without storing sensitive data. Paykeys eliminate fraud by
// ensuring the person initiating payment owns the funding account.
//
// PaykeyService contains methods and other services that help with interacting
// with the straddle API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPaykeyService] method instead.
type PaykeyService struct {
	options []option.RequestOption
	// Paykeys are secure tokens that link verified customer identities to their bank
	// accounts. Each Paykey includes built-in balance checking, fraud detection
	// through LSTM machine learning models, and can be reused for subscriptions and
	// recurring payments without storing sensitive data. Paykeys eliminate fraud by
	// ensuring the person initiating payment owns the funding account.
	Review PaykeyReviewService
}

// NewPaykeyService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewPaykeyService(opts ...option.RequestOption) (r PaykeyService) {
	r = PaykeyService{}
	r.options = opts
	r.Review = NewPaykeyReviewService(opts...)
	return
}

// Returns a list of paykeys associated with a Straddle account. This endpoint
// supports advanced sorting and filtering options.
func (r *PaykeyService) List(ctx context.Context, params PaykeyListParams, opts ...option.RequestOption) (res *pagination.PageNumberSchema[PaykeySummaryPagedV1Data], err error) {
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
	path := "v1/paykeys"
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

// Returns a list of paykeys associated with a Straddle account. This endpoint
// supports advanced sorting and filtering options.
func (r *PaykeyService) ListAutoPaging(ctx context.Context, params PaykeyListParams, opts ...option.RequestOption) *pagination.PageNumberSchemaAutoPager[PaykeySummaryPagedV1Data] {
	return pagination.NewPageNumberSchemaAutoPager(r.List(ctx, params, opts...))
}

func (r *PaykeyService) Cancel(ctx context.Context, id string, params PaykeyCancelParams, opts ...option.RequestOption) (res *PaykeyV1, err error) {
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
	path := fmt.Sprintf("v1/paykeys/%s/cancel", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return res, err
}

// Retrieves the details of an existing paykey. Supply the unique paykey `id` and
// Straddle will return the corresponding paykey record , including the `paykey`
// token value and masked bank account details.
func (r *PaykeyService) Get(ctx context.Context, id string, query PaykeyGetParams, opts ...option.RequestOption) (res *PaykeyV1, err error) {
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
	path := fmt.Sprintf("v1/paykeys/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Retrieves the details of a paykey that has previously been created. Supply the
// unique paykey ID that was returned from your previous request, and Straddle will
// return the corresponding paykey information including the unmasked token.
func (r *PaykeyService) Reveal(ctx context.Context, id string, query PaykeyRevealParams, opts ...option.RequestOption) (res *PaykeyRevealResponse, err error) {
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
	path := fmt.Sprintf("v1/paykeys/%s/reveal", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Retrieves the unmasked details of an existing paykey. Supply the unique paykey
// `id` and Straddle will return the corresponding paykey record, including the
// unmasked bank account details. This endpoint needs to be enabled by Straddle for
// your account and should only be used when absolutely necessary.
func (r *PaykeyService) Unmasked(ctx context.Context, id string, query PaykeyUnmaskedParams, opts ...option.RequestOption) (res *PaykeyUnmaskedV1, err error) {
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
	path := fmt.Sprintf("v1/paykeys/%s/unmasked", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Updates the balance of a paykey. This endpoint allows you to refresh the balance
// of a paykey.
func (r *PaykeyService) UpdateBalance(ctx context.Context, id string, body PaykeyUpdateBalanceParams, opts ...option.RequestOption) (res *PaykeyV1, err error) {
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
	path := fmt.Sprintf("v1/paykeys/%s/refresh_balance", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, nil, &res, opts...)
	return res, err
}

type PaykeySummaryPagedV1 struct {
	Data []PaykeySummaryPagedV1Data `json:"data" api:"required"`
	Meta PaykeySummaryPagedV1Meta   `json:"meta" api:"required"`
	// Indicates the structure of the returned content.
	//
	//   - "object" means the `data` field contains a single JSON object.
	//   - "array" means the `data` field contains an array of objects.
	//   - "error" means the `data` field contains an error object with details of the
	//     issue.
	//   - "none" means no data is returned.
	//
	// Any of "object", "array", "error", "none".
	ResponseType PaykeySummaryPagedV1ResponseType `json:"response_type" api:"required"`
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
func (r PaykeySummaryPagedV1) RawJSON() string { return r.JSON.raw }
func (r *PaykeySummaryPagedV1) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaykeySummaryPagedV1Data struct {
	// Unique identifier for the paykey.
	ID     string                         `json:"id" api:"required" format:"uuid"`
	Config PaykeySummaryPagedV1DataConfig `json:"config" api:"required"`
	// Timestamp of when the paykey was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Human-readable label that combines the bank name and masked account number to
	// help easility represent this paykey in a UI
	Label string `json:"label" api:"required"`
	// The tokenized paykey value. This value is used to create payments and should be
	// stored securely.
	Paykey string `json:"paykey" api:"required"`
	// Any of "bank_account", "straddle", "mx", "plaid", "tan", "quiltt", "mastercard".
	Source string `json:"source" api:"required"`
	// Any of "pending", "active", "inactive", "rejected", "review", "blocked".
	Status string `json:"status" api:"required"`
	// Timestamp of the most recent update to the paykey.
	UpdatedAt time.Time                        `json:"updated_at" api:"required" format:"date-time"`
	BankData  PaykeySummaryPagedV1DataBankData `json:"bank_data"`
	// Unique identifier of the related customer object.
	CustomerID string `json:"customer_id" api:"nullable" format:"uuid"`
	// Expiration date and time of the paykey, if applicable.
	ExpiresAt time.Time `json:"expires_at" api:"nullable" format:"date-time"`
	// Unique identifier for the paykey in your database, used for cross-referencing
	// between Straddle and your systems.
	ExternalID string `json:"external_id" api:"nullable"`
	// Name of the financial institution.
	InstitutionName string                                `json:"institution_name" api:"nullable"`
	StatusDetails   PaykeySummaryPagedV1DataStatusDetails `json:"status_details"`
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
		BankData        respjson.Field
		CustomerID      respjson.Field
		ExpiresAt       respjson.Field
		ExternalID      respjson.Field
		InstitutionName respjson.Field
		StatusDetails   respjson.Field
		UnblockEligible respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaykeySummaryPagedV1Data) RawJSON() string { return r.JSON.raw }
func (r *PaykeySummaryPagedV1Data) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaykeySummaryPagedV1DataConfig struct {
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
func (r PaykeySummaryPagedV1DataConfig) RawJSON() string { return r.JSON.raw }
func (r *PaykeySummaryPagedV1DataConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaykeySummaryPagedV1DataBankData struct {
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
func (r PaykeySummaryPagedV1DataBankData) RawJSON() string { return r.JSON.raw }
func (r *PaykeySummaryPagedV1DataBankData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaykeySummaryPagedV1DataStatusDetails struct {
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
func (r PaykeySummaryPagedV1DataStatusDetails) RawJSON() string { return r.JSON.raw }
func (r *PaykeySummaryPagedV1DataStatusDetails) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaykeySummaryPagedV1Meta struct {
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
func (r PaykeySummaryPagedV1Meta) RawJSON() string { return r.JSON.raw }
func (r *PaykeySummaryPagedV1Meta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indicates the structure of the returned content.
//
//   - "object" means the `data` field contains a single JSON object.
//   - "array" means the `data` field contains an array of objects.
//   - "error" means the `data` field contains an error object with details of the
//     issue.
//   - "none" means no data is returned.
type PaykeySummaryPagedV1ResponseType string

const (
	PaykeySummaryPagedV1ResponseTypeObject PaykeySummaryPagedV1ResponseType = "object"
	PaykeySummaryPagedV1ResponseTypeArray  PaykeySummaryPagedV1ResponseType = "array"
	PaykeySummaryPagedV1ResponseTypeError  PaykeySummaryPagedV1ResponseType = "error"
	PaykeySummaryPagedV1ResponseTypeNone   PaykeySummaryPagedV1ResponseType = "none"
)

type PaykeyUnmaskedV1 struct {
	Data PaykeyUnmaskedV1Data `json:"data" api:"required"`
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
	ResponseType PaykeyUnmaskedV1ResponseType `json:"response_type" api:"required"`
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
func (r PaykeyUnmaskedV1) RawJSON() string { return r.JSON.raw }
func (r *PaykeyUnmaskedV1) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaykeyUnmaskedV1Data struct {
	// Unique identifier for the paykey.
	ID     string                     `json:"id" api:"required" format:"uuid"`
	Config PaykeyUnmaskedV1DataConfig `json:"config" api:"required"`
	// Timestamp of when the paykey was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Human-readable label used to represent this paykey in a UI.
	Label string `json:"label" api:"required"`
	// The tokenized paykey value. This value is used to create payments and should be
	// stored securely.
	Paykey string `json:"paykey" api:"required"`
	// Any of "bank_account", "straddle", "mx", "plaid", "tan", "quiltt", "mastercard".
	Source string `json:"source" api:"required"`
	// Any of "pending", "active", "inactive", "rejected", "review", "blocked".
	Status string `json:"status" api:"required"`
	// Timestamp of the most recent update to the paykey.
	UpdatedAt time.Time                    `json:"updated_at" api:"required" format:"date-time"`
	Balance   PaykeyUnmaskedV1DataBalance  `json:"balance"`
	BankData  PaykeyUnmaskedV1DataBankData `json:"bank_data"`
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
	Metadata      map[string]string                 `json:"metadata" api:"nullable"`
	StatusDetails PaykeyUnmaskedV1DataStatusDetails `json:"status_details"`
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
func (r PaykeyUnmaskedV1Data) RawJSON() string { return r.JSON.raw }
func (r *PaykeyUnmaskedV1Data) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaykeyUnmaskedV1DataConfig struct {
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
func (r PaykeyUnmaskedV1DataConfig) RawJSON() string { return r.JSON.raw }
func (r *PaykeyUnmaskedV1DataConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaykeyUnmaskedV1DataBalance struct {
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
func (r PaykeyUnmaskedV1DataBalance) RawJSON() string { return r.JSON.raw }
func (r *PaykeyUnmaskedV1DataBalance) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaykeyUnmaskedV1DataBankData struct {
	// The bank account number
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
func (r PaykeyUnmaskedV1DataBankData) RawJSON() string { return r.JSON.raw }
func (r *PaykeyUnmaskedV1DataBankData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaykeyUnmaskedV1DataStatusDetails struct {
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
func (r PaykeyUnmaskedV1DataStatusDetails) RawJSON() string { return r.JSON.raw }
func (r *PaykeyUnmaskedV1DataStatusDetails) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indicates the structure of the returned content.
//
//   - "object" means the `data` field contains a single JSON object.
//   - "array" means the `data` field contains an array of objects.
//   - "error" means the `data` field contains an error object with details of the
//     issue.
//   - "none" means no data is returned.
type PaykeyUnmaskedV1ResponseType string

const (
	PaykeyUnmaskedV1ResponseTypeObject PaykeyUnmaskedV1ResponseType = "object"
	PaykeyUnmaskedV1ResponseTypeArray  PaykeyUnmaskedV1ResponseType = "array"
	PaykeyUnmaskedV1ResponseTypeError  PaykeyUnmaskedV1ResponseType = "error"
	PaykeyUnmaskedV1ResponseTypeNone   PaykeyUnmaskedV1ResponseType = "none"
)

type PaykeyV1 struct {
	Data PaykeyV1Data `json:"data" api:"required"`
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
	ResponseType PaykeyV1ResponseType `json:"response_type" api:"required"`
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
func (r PaykeyV1) RawJSON() string { return r.JSON.raw }
func (r *PaykeyV1) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaykeyV1Data struct {
	// Unique identifier for the paykey.
	ID     string             `json:"id" api:"required" format:"uuid"`
	Config PaykeyV1DataConfig `json:"config" api:"required"`
	// Timestamp of when the paykey was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Human-readable label used to represent this paykey in a UI.
	Label string `json:"label" api:"required"`
	// The tokenized paykey value. This value is used to create payments and should be
	// stored securely.
	Paykey string `json:"paykey" api:"required"`
	// Any of "bank_account", "straddle", "mx", "plaid", "tan", "quiltt", "mastercard".
	Source string `json:"source" api:"required"`
	// Any of "pending", "active", "inactive", "rejected", "review", "blocked".
	Status string `json:"status" api:"required"`
	// Timestamp of the most recent update to the paykey.
	UpdatedAt time.Time            `json:"updated_at" api:"required" format:"date-time"`
	Balance   PaykeyV1DataBalance  `json:"balance"`
	BankData  PaykeyV1DataBankData `json:"bank_data"`
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
	Metadata      map[string]string         `json:"metadata" api:"nullable"`
	StatusDetails PaykeyV1DataStatusDetails `json:"status_details"`
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
func (r PaykeyV1Data) RawJSON() string { return r.JSON.raw }
func (r *PaykeyV1Data) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaykeyV1DataConfig struct {
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
func (r PaykeyV1DataConfig) RawJSON() string { return r.JSON.raw }
func (r *PaykeyV1DataConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaykeyV1DataBalance struct {
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
func (r PaykeyV1DataBalance) RawJSON() string { return r.JSON.raw }
func (r *PaykeyV1DataBalance) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaykeyV1DataBankData struct {
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
func (r PaykeyV1DataBankData) RawJSON() string { return r.JSON.raw }
func (r *PaykeyV1DataBankData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaykeyV1DataStatusDetails struct {
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
func (r PaykeyV1DataStatusDetails) RawJSON() string { return r.JSON.raw }
func (r *PaykeyV1DataStatusDetails) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indicates the structure of the returned content.
//
//   - "object" means the `data` field contains a single JSON object.
//   - "array" means the `data` field contains an array of objects.
//   - "error" means the `data` field contains an error object with details of the
//     issue.
//   - "none" means no data is returned.
type PaykeyV1ResponseType string

const (
	PaykeyV1ResponseTypeObject PaykeyV1ResponseType = "object"
	PaykeyV1ResponseTypeArray  PaykeyV1ResponseType = "array"
	PaykeyV1ResponseTypeError  PaykeyV1ResponseType = "error"
	PaykeyV1ResponseTypeNone   PaykeyV1ResponseType = "none"
)

type PaykeyRevealResponse struct {
	Data PaykeyRevealResponseData `json:"data" api:"required"`
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
	ResponseType PaykeyRevealResponseResponseType `json:"response_type" api:"required"`
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
func (r PaykeyRevealResponse) RawJSON() string { return r.JSON.raw }
func (r *PaykeyRevealResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaykeyRevealResponseData struct {
	// Unique identifier for the paykey.
	ID     string                         `json:"id" api:"required" format:"uuid"`
	Config PaykeyRevealResponseDataConfig `json:"config" api:"required"`
	// Timestamp of when the paykey was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Human-readable label that combines the bank name and masked account number to
	// help easility represent this paykey in a UI
	Label string `json:"label" api:"required"`
	// The tokenized paykey value. This token is used to create payments and should be
	// stored securely.
	Paykey string `json:"paykey" api:"required"`
	// Any of "bank_account", "straddle", "mx", "plaid", "tan", "quiltt", "mastercard".
	Source string `json:"source" api:"required"`
	// Any of "pending", "active", "inactive", "rejected", "review", "blocked".
	Status string `json:"status" api:"required"`
	// Timestamp of the most recent update to the paykey.
	UpdatedAt time.Time                        `json:"updated_at" api:"required" format:"date-time"`
	Balance   PaykeyRevealResponseDataBalance  `json:"balance"`
	BankData  PaykeyRevealResponseDataBankData `json:"bank_data"`
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
	Metadata      map[string]string                     `json:"metadata" api:"nullable"`
	StatusDetails PaykeyRevealResponseDataStatusDetails `json:"status_details"`
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
func (r PaykeyRevealResponseData) RawJSON() string { return r.JSON.raw }
func (r *PaykeyRevealResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaykeyRevealResponseDataConfig struct {
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
func (r PaykeyRevealResponseDataConfig) RawJSON() string { return r.JSON.raw }
func (r *PaykeyRevealResponseDataConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaykeyRevealResponseDataBalance struct {
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
func (r PaykeyRevealResponseDataBalance) RawJSON() string { return r.JSON.raw }
func (r *PaykeyRevealResponseDataBalance) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaykeyRevealResponseDataBankData struct {
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
func (r PaykeyRevealResponseDataBankData) RawJSON() string { return r.JSON.raw }
func (r *PaykeyRevealResponseDataBankData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaykeyRevealResponseDataStatusDetails struct {
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
func (r PaykeyRevealResponseDataStatusDetails) RawJSON() string { return r.JSON.raw }
func (r *PaykeyRevealResponseDataStatusDetails) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indicates the structure of the returned content.
//
//   - "object" means the `data` field contains a single JSON object.
//   - "array" means the `data` field contains an array of objects.
//   - "error" means the `data` field contains an error object with details of the
//     issue.
//   - "none" means no data is returned.
type PaykeyRevealResponseResponseType string

const (
	PaykeyRevealResponseResponseTypeObject PaykeyRevealResponseResponseType = "object"
	PaykeyRevealResponseResponseTypeArray  PaykeyRevealResponseResponseType = "array"
	PaykeyRevealResponseResponseTypeError  PaykeyRevealResponseResponseType = "error"
	PaykeyRevealResponseResponseTypeNone   PaykeyRevealResponseResponseType = "none"
)

type PaykeyListParams struct {
	// Start date for filtering by creation date.
	CreatedFrom param.Opt[time.Time] `query:"created_from,omitzero" format:"date-time" json:"-"`
	// End date for filtering by creation date.
	CreatedTo param.Opt[time.Time] `query:"created_to,omitzero" format:"date-time" json:"-"`
	// Filter paykeys by related customer ID.
	CustomerID param.Opt[string] `query:"customer_id,omitzero" format:"uuid" json:"-"`
	// Page number for paginated results. Starts at 1.
	PageNumber param.Opt[int64] `query:"page_number,omitzero" json:"-"`
	// Number of results per page. Maximum: 1000.
	PageSize param.Opt[int64] `query:"page_size,omitzero" json:"-"`
	// General search term to filter paykeys.
	SearchText param.Opt[string] `query:"search_text,omitzero" json:"-"`
	// Filter paykeys by unblock eligibility. When true, returns only blocked paykeys
	// eligible for client-initiated unblocking (blocked due to R29 returns and not
	// previously unblocked). When false, returns only blocked paykeys that are not
	// eligible for unblocking.
	UnblockEligible   param.Opt[bool]   `query:"unblock_eligible,omitzero" json:"-"`
	CorrelationID     param.Opt[string] `header:"Correlation-Id,omitzero" json:"-"`
	RequestID         param.Opt[string] `header:"Request-Id,omitzero" json:"-"`
	StraddleAccountID param.Opt[string] `header:"Straddle-Account-Id,omitzero" format:"uuid" json:"-"`
	// Any of "institution_name", "expires_at", "created_at".
	SortBy PaykeyListParamsSortBy `query:"sort_by,omitzero" json:"-"`
	// Any of "asc", "desc".
	SortOrder PaykeyListParamsSortOrder `query:"sort_order,omitzero" json:"-"`
	// Filter paykeys by their source.
	//
	// Any of "bank_account", "straddle", "mx", "plaid", "tan", "quiltt", "mastercard".
	Source []string `query:"source,omitzero" json:"-"`
	// Filter paykeys by their current status.
	//
	// Any of "pending", "active", "inactive", "rejected", "review", "blocked".
	Status []string `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PaykeyListParams]'s query parameters as `url.Values`.
func (r PaykeyListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PaykeyListParamsSortBy string

const (
	PaykeyListParamsSortByInstitutionName PaykeyListParamsSortBy = "institution_name"
	PaykeyListParamsSortByExpiresAt       PaykeyListParamsSortBy = "expires_at"
	PaykeyListParamsSortByCreatedAt       PaykeyListParamsSortBy = "created_at"
)

type PaykeyListParamsSortOrder string

const (
	PaykeyListParamsSortOrderAsc  PaykeyListParamsSortOrder = "asc"
	PaykeyListParamsSortOrderDesc PaykeyListParamsSortOrder = "desc"
)

type PaykeyCancelParams struct {
	Reason            param.Opt[string] `json:"reason,omitzero"`
	CorrelationID     param.Opt[string] `header:"Correlation-Id,omitzero" json:"-"`
	IdempotencyKey    param.Opt[string] `header:"Idempotency-Key,omitzero" json:"-"`
	RequestID         param.Opt[string] `header:"Request-Id,omitzero" json:"-"`
	StraddleAccountID param.Opt[string] `header:"Straddle-Account-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

func (r PaykeyCancelParams) MarshalJSON() (data []byte, err error) {
	type shadow PaykeyCancelParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PaykeyCancelParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaykeyGetParams struct {
	CorrelationID     param.Opt[string] `header:"Correlation-Id,omitzero" json:"-"`
	RequestID         param.Opt[string] `header:"Request-Id,omitzero" json:"-"`
	StraddleAccountID param.Opt[string] `header:"Straddle-Account-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

type PaykeyRevealParams struct {
	CorrelationID     param.Opt[string] `header:"Correlation-Id,omitzero" json:"-"`
	RequestID         param.Opt[string] `header:"Request-Id,omitzero" json:"-"`
	StraddleAccountID param.Opt[string] `header:"Straddle-Account-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

type PaykeyUnmaskedParams struct {
	CorrelationID     param.Opt[string] `header:"Correlation-Id,omitzero" json:"-"`
	RequestID         param.Opt[string] `header:"Request-Id,omitzero" json:"-"`
	StraddleAccountID param.Opt[string] `header:"Straddle-Account-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

type PaykeyUpdateBalanceParams struct {
	CorrelationID     param.Opt[string] `header:"Correlation-Id,omitzero" json:"-"`
	IdempotencyKey    param.Opt[string] `header:"Idempotency-Key,omitzero" json:"-"`
	RequestID         param.Opt[string] `header:"Request-Id,omitzero" json:"-"`
	StraddleAccountID param.Opt[string] `header:"Straddle-Account-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}
