// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package straddle

import (
	"context"
	"encoding/json"
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

// Accounts represent businesses using Straddle through your platform. Each account
// must complete automated verification before processing payments. Use accounts to
// manage your users' payment capabilities, track verification status, and control
// access to features. Accounts can be instantly created in sandbox and require
// additional verification for production access.
//
// EmbedAccountService contains methods and other services that help with
// interacting with the straddle API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEmbedAccountService] method instead.
type EmbedAccountService struct {
	options []option.RequestOption
	// Capabilities enable specific features and services for an Account. Use
	// capability requests to unlock higher processing limits, new payment types, or
	// additional platform features as your users' businesses grow. Track approval
	// status and manage documentation requirements through a single interface.
	CapabilityRequests EmbedAccountCapabilityRequestService
}

// NewEmbedAccountService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewEmbedAccountService(opts ...option.RequestOption) (r EmbedAccountService) {
	r = EmbedAccountService{}
	r.options = opts
	r.CapabilityRequests = NewEmbedAccountCapabilityRequestService(opts...)
	return
}

// Creates a new account associated with your Straddle platform integration. This
// endpoint allows you to set up an account with specified details, including
// business information and access levels.
func (r *EmbedAccountService) New(ctx context.Context, params EmbedAccountNewParams, opts ...option.RequestOption) (res *AccountV1, err error) {
	if !param.IsOmitted(params.CorrelationID) {
		opts = append(opts, option.WithHeader("correlation-id", fmt.Sprintf("%v", params.CorrelationID.Value)))
	}
	if !param.IsOmitted(params.IdempotencyKey) {
		opts = append(opts, option.WithHeader("idempotency-key", fmt.Sprintf("%v", params.IdempotencyKey.Value)))
	}
	if !param.IsOmitted(params.RequestID) {
		opts = append(opts, option.WithHeader("request-id", fmt.Sprintf("%v", params.RequestID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	path := "v1/accounts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Updates an existing account's information. This endpoint allows you to update
// various account details during onboarding or after the account has been created.
func (r *EmbedAccountService) Update(ctx context.Context, accountID string, params EmbedAccountUpdateParams, opts ...option.RequestOption) (res *AccountV1, err error) {
	if !param.IsOmitted(params.CorrelationID) {
		opts = append(opts, option.WithHeader("correlation-id", fmt.Sprintf("%v", params.CorrelationID.Value)))
	}
	if !param.IsOmitted(params.IdempotencyKey) {
		opts = append(opts, option.WithHeader("idempotency-key", fmt.Sprintf("%v", params.IdempotencyKey.Value)))
	}
	if !param.IsOmitted(params.RequestID) {
		opts = append(opts, option.WithHeader("request-id", fmt.Sprintf("%v", params.RequestID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/accounts/%s", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return res, err
}

// Returns a list of accounts associated with your Straddle platform integration.
// The accounts are returned sorted by creation date, with the most recently
// created accounts appearing first. This endpoint supports advanced sorting and
// filtering options.
func (r *EmbedAccountService) List(ctx context.Context, params EmbedAccountListParams, opts ...option.RequestOption) (res *pagination.PageNumberSchema[AccountPagedV1Data], err error) {
	var raw *http.Response
	if !param.IsOmitted(params.CorrelationID) {
		opts = append(opts, option.WithHeader("correlation-id", fmt.Sprintf("%v", params.CorrelationID.Value)))
	}
	if !param.IsOmitted(params.RequestID) {
		opts = append(opts, option.WithHeader("request-id", fmt.Sprintf("%v", params.RequestID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/accounts"
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

// Returns a list of accounts associated with your Straddle platform integration.
// The accounts are returned sorted by creation date, with the most recently
// created accounts appearing first. This endpoint supports advanced sorting and
// filtering options.
func (r *EmbedAccountService) ListAutoPaging(ctx context.Context, params EmbedAccountListParams, opts ...option.RequestOption) *pagination.PageNumberSchemaAutoPager[AccountPagedV1Data] {
	return pagination.NewPageNumberSchemaAutoPager(r.List(ctx, params, opts...))
}

// Retrieves the details of an account that has previously been created. Supply the
// unique account ID that was returned from your previous request, and Straddle
// will return the corresponding account information.
func (r *EmbedAccountService) Get(ctx context.Context, accountID string, query EmbedAccountGetParams, opts ...option.RequestOption) (res *AccountV1, err error) {
	if !param.IsOmitted(query.CorrelationID) {
		opts = append(opts, option.WithHeader("correlation-id", fmt.Sprintf("%v", query.CorrelationID.Value)))
	}
	if !param.IsOmitted(query.RequestID) {
		opts = append(opts, option.WithHeader("request-id", fmt.Sprintf("%v", query.RequestID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/accounts/%s", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Initiates the onboarding process for a new account. This endpoint can only be
// used for accounts where at least one representative and one bank account have
// already been created.
func (r *EmbedAccountService) Onboard(ctx context.Context, accountID string, params EmbedAccountOnboardParams, opts ...option.RequestOption) (res *AccountV1, err error) {
	if !param.IsOmitted(params.CorrelationID) {
		opts = append(opts, option.WithHeader("correlation-id", fmt.Sprintf("%v", params.CorrelationID.Value)))
	}
	if !param.IsOmitted(params.IdempotencyKey) {
		opts = append(opts, option.WithHeader("idempotency-key", fmt.Sprintf("%v", params.IdempotencyKey.Value)))
	}
	if !param.IsOmitted(params.RequestID) {
		opts = append(opts, option.WithHeader("request-id", fmt.Sprintf("%v", params.RequestID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/accounts/%s/onboard", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Simulate the status transitions for sandbox accounts. This endpoint can only be
// used for sandbox accounts.
func (r *EmbedAccountService) Simulate(ctx context.Context, accountID string, params EmbedAccountSimulateParams, opts ...option.RequestOption) (res *AccountV1, err error) {
	if !param.IsOmitted(params.CorrelationID) {
		opts = append(opts, option.WithHeader("correlation-id", fmt.Sprintf("%v", params.CorrelationID.Value)))
	}
	if !param.IsOmitted(params.IdempotencyKey) {
		opts = append(opts, option.WithHeader("idempotency-key", fmt.Sprintf("%v", params.IdempotencyKey.Value)))
	}
	if !param.IsOmitted(params.RequestID) {
		opts = append(opts, option.WithHeader("request-id", fmt.Sprintf("%v", params.RequestID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/accounts/%s/simulate", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

type AccountPagedV1 struct {
	Data []AccountPagedV1Data `json:"data" api:"required"`
	// Metadata about the API request, including an identifier, timestamp, and
	// pagination details.
	Meta shared.PagedResponseMetadata `json:"meta" api:"required"`
	// Indicates the structure of the returned content.
	//
	//   - "object" means the `data` field contains a single JSON object.
	//   - "array" means the `data` field contains an array of objects.
	//   - "error" means the `data` field contains an error object with details of the
	//     issue.
	//   - "none" means no data is returned.
	//
	// Any of "object", "array", "error", "none".
	ResponseType AccountPagedV1ResponseType `json:"response_type" api:"required"`
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
func (r AccountPagedV1) RawJSON() string { return r.JSON.raw }
func (r *AccountPagedV1) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountPagedV1Data struct {
	// Unique identifier for the account.
	ID string `json:"id" api:"required" format:"uuid"`
	// The access level granted to the account. This is determined by your platform
	// configuration. Use `standard` unless instructed otherwise by Straddle.
	//
	// Any of "standard", "managed".
	AccessLevel string `json:"access_level" api:"required"`
	// The unique identifier of the organization this account belongs to.
	OrganizationID string `json:"organization_id" api:"required" format:"uuid"`
	// The current status of the account (e.g., 'active', 'inactive', 'pending').
	//
	// Any of "created", "onboarding", "active", "rejected", "inactive".
	Status       string                         `json:"status" api:"required"`
	StatusDetail AccountPagedV1DataStatusDetail `json:"status_detail" api:"required"`
	// The type of account (e.g., 'individual', 'business').
	//
	// Any of "business".
	Type            string                         `json:"type" api:"required"`
	BusinessProfile BusinessProfileV1              `json:"business_profile"`
	Capabilities    AccountPagedV1DataCapabilities `json:"capabilities"`
	// Timestamp of when the account was created.
	CreatedAt time.Time `json:"created_at" api:"nullable" format:"date-time"`
	// Unique identifier for the account in your database, used for cross-referencing
	// between Straddle and your systems.
	ExternalID string `json:"external_id" api:"nullable"`
	// Up to 20 additional user-defined key-value pairs. Useful for storing additional
	// information about the account in a structured format.
	Metadata       map[string]string          `json:"metadata" api:"nullable"`
	Settings       AccountPagedV1DataSettings `json:"settings"`
	TermsOfService TermsOfServiceV1           `json:"terms_of_service"`
	// Timestamp of the most recent update to the account.
	UpdatedAt time.Time `json:"updated_at" api:"nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		AccessLevel     respjson.Field
		OrganizationID  respjson.Field
		Status          respjson.Field
		StatusDetail    respjson.Field
		Type            respjson.Field
		BusinessProfile respjson.Field
		Capabilities    respjson.Field
		CreatedAt       respjson.Field
		ExternalID      respjson.Field
		Metadata        respjson.Field
		Settings        respjson.Field
		TermsOfService  respjson.Field
		UpdatedAt       respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccountPagedV1Data) RawJSON() string { return r.JSON.raw }
func (r *AccountPagedV1Data) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountPagedV1DataStatusDetail struct {
	// A machine-readable code for the specific status, useful for programmatic
	// handling.
	Code string `json:"code" api:"required"`
	// A human-readable message describing the current status.
	Message string `json:"message" api:"required"`
	// A machine-readable identifier for the specific status, useful for programmatic
	// handling.
	//
	// Any of "unverified", "in_review", "pending", "stuck", "verified",
	// "failed_verification", "disabled", "terminated", "new".
	Reason string `json:"reason" api:"required"`
	// Identifies the origin of the status change (e.g., `bank_decline`, `watchtower`).
	// This helps in tracking the cause of status updates.
	//
	// Any of "watchtower".
	Source string `json:"source" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Message     respjson.Field
		Reason      respjson.Field
		Source      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccountPagedV1DataStatusDetail) RawJSON() string { return r.JSON.raw }
func (r *AccountPagedV1DataStatusDetail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountPagedV1DataCapabilities struct {
	ConsentTypes  AccountPagedV1DataCapabilitiesConsentTypes  `json:"consent_types" api:"required"`
	CustomerTypes AccountPagedV1DataCapabilitiesCustomerTypes `json:"customer_types" api:"required"`
	PaymentTypes  AccountPagedV1DataCapabilitiesPaymentTypes  `json:"payment_types" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ConsentTypes  respjson.Field
		CustomerTypes respjson.Field
		PaymentTypes  respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccountPagedV1DataCapabilities) RawJSON() string { return r.JSON.raw }
func (r *AccountPagedV1DataCapabilities) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountPagedV1DataCapabilitiesConsentTypes struct {
	// Whether the internet payment authorization capability is enabled for the
	// account.
	Internet CapabilityV1 `json:"internet" api:"required"`
	// Whether the signed agreement payment authorization capability is enabled for the
	// account.
	SignedAgreement CapabilityV1 `json:"signed_agreement" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Internet        respjson.Field
		SignedAgreement respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccountPagedV1DataCapabilitiesConsentTypes) RawJSON() string { return r.JSON.raw }
func (r *AccountPagedV1DataCapabilitiesConsentTypes) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountPagedV1DataCapabilitiesCustomerTypes struct {
	Businesses  CapabilityV1 `json:"businesses" api:"required"`
	Individuals CapabilityV1 `json:"individuals" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Businesses  respjson.Field
		Individuals respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccountPagedV1DataCapabilitiesCustomerTypes) RawJSON() string { return r.JSON.raw }
func (r *AccountPagedV1DataCapabilitiesCustomerTypes) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountPagedV1DataCapabilitiesPaymentTypes struct {
	Charges CapabilityV1 `json:"charges" api:"required"`
	Payouts CapabilityV1 `json:"payouts" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Charges     respjson.Field
		Payouts     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccountPagedV1DataCapabilitiesPaymentTypes) RawJSON() string { return r.JSON.raw }
func (r *AccountPagedV1DataCapabilitiesPaymentTypes) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountPagedV1DataSettings struct {
	Charges AccountPagedV1DataSettingsCharges `json:"charges" api:"required"`
	Payouts AccountPagedV1DataSettingsPayouts `json:"payouts" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Charges     respjson.Field
		Payouts     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccountPagedV1DataSettings) RawJSON() string { return r.JSON.raw }
func (r *AccountPagedV1DataSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountPagedV1DataSettingsCharges struct {
	// The maximum dollar amount of charges in a calendar day.
	DailyAmount int64 `json:"daily_amount" api:"required"`
	// The amount of time it takes for a charge to be funded. This value is defined by
	// Straddle.
	//
	// Any of "immediate", "next_day", "one_day", "two_day", "three_day", "four_day",
	// "five_day".
	FundingTime string `json:"funding_time" api:"required"`
	// The unique identifier of the linked bank account associated with charges. This
	// value is defined by Straddle.
	LinkedBankAccountID string `json:"linked_bank_account_id" api:"required" format:"uuid"`
	// The maximum amount of a single charge.
	MaxAmount int64 `json:"max_amount" api:"required"`
	// The maximum dollar amount of charges in a calendar month.
	MonthlyAmount int64 `json:"monthly_amount" api:"required"`
	// The maximum number of charges in a calendar month.
	MonthlyCount int64 `json:"monthly_count" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DailyAmount         respjson.Field
		FundingTime         respjson.Field
		LinkedBankAccountID respjson.Field
		MaxAmount           respjson.Field
		MonthlyAmount       respjson.Field
		MonthlyCount        respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccountPagedV1DataSettingsCharges) RawJSON() string { return r.JSON.raw }
func (r *AccountPagedV1DataSettingsCharges) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountPagedV1DataSettingsPayouts struct {
	// The maximum dollar amount of payouts in a day.
	DailyAmount int64 `json:"daily_amount" api:"required"`
	// The amount of time it takes for a payout to be funded. This value is defined by
	// Straddle.
	//
	// Any of "immediate", "next_day", "one_day", "two_day", "three_day", "four_day",
	// "five_day".
	FundingTime string `json:"funding_time" api:"required"`
	// The unique identifier of the linked bank account to use for payouts.
	LinkedBankAccountID string `json:"linked_bank_account_id" api:"required" format:"uuid"`
	// The maximum amount of a single payout.
	MaxAmount int64 `json:"max_amount" api:"required"`
	// The maximum dollar amount of payouts in a month.
	MonthlyAmount int64 `json:"monthly_amount" api:"required"`
	// The maximum number of payouts in a month.
	MonthlyCount int64 `json:"monthly_count" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DailyAmount         respjson.Field
		FundingTime         respjson.Field
		LinkedBankAccountID respjson.Field
		MaxAmount           respjson.Field
		MonthlyAmount       respjson.Field
		MonthlyCount        respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccountPagedV1DataSettingsPayouts) RawJSON() string { return r.JSON.raw }
func (r *AccountPagedV1DataSettingsPayouts) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indicates the structure of the returned content.
//
//   - "object" means the `data` field contains a single JSON object.
//   - "array" means the `data` field contains an array of objects.
//   - "error" means the `data` field contains an error object with details of the
//     issue.
//   - "none" means no data is returned.
type AccountPagedV1ResponseType string

const (
	AccountPagedV1ResponseTypeObject AccountPagedV1ResponseType = "object"
	AccountPagedV1ResponseTypeArray  AccountPagedV1ResponseType = "array"
	AccountPagedV1ResponseTypeError  AccountPagedV1ResponseType = "error"
	AccountPagedV1ResponseTypeNone   AccountPagedV1ResponseType = "none"
)

type AccountV1 struct {
	Data AccountV1Data `json:"data" api:"required"`
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
	ResponseType AccountV1ResponseType `json:"response_type" api:"required"`
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
func (r AccountV1) RawJSON() string { return r.JSON.raw }
func (r *AccountV1) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountV1Data struct {
	// Unique identifier for the account.
	ID string `json:"id" api:"required" format:"uuid"`
	// The access level granted to the account. This is determined by your platform
	// configuration. Use `standard` unless instructed otherwise by Straddle.
	//
	// Any of "standard", "managed".
	AccessLevel string `json:"access_level" api:"required"`
	// The unique identifier of the organization this account belongs to.
	OrganizationID string `json:"organization_id" api:"required" format:"uuid"`
	// The current status of the account (e.g., 'active', 'inactive', 'pending').
	//
	// Any of "created", "onboarding", "active", "rejected", "inactive".
	Status       string                    `json:"status" api:"required"`
	StatusDetail AccountV1DataStatusDetail `json:"status_detail" api:"required"`
	// The type of account (e.g., 'individual', 'business').
	//
	// Any of "business".
	Type            string                    `json:"type" api:"required"`
	BusinessProfile BusinessProfileV1         `json:"business_profile"`
	Capabilities    AccountV1DataCapabilities `json:"capabilities"`
	// Timestamp of when the account was created.
	CreatedAt time.Time `json:"created_at" api:"nullable" format:"date-time"`
	// Unique identifier for the account in your database, used for cross-referencing
	// between Straddle and your systems.
	ExternalID string `json:"external_id" api:"nullable"`
	// Up to 20 additional user-defined key-value pairs. Useful for storing additional
	// information about the account in a structured format.
	Metadata       map[string]string     `json:"metadata" api:"nullable"`
	Settings       AccountV1DataSettings `json:"settings"`
	TermsOfService TermsOfServiceV1      `json:"terms_of_service"`
	// Timestamp of the most recent update to the account.
	UpdatedAt time.Time `json:"updated_at" api:"nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		AccessLevel     respjson.Field
		OrganizationID  respjson.Field
		Status          respjson.Field
		StatusDetail    respjson.Field
		Type            respjson.Field
		BusinessProfile respjson.Field
		Capabilities    respjson.Field
		CreatedAt       respjson.Field
		ExternalID      respjson.Field
		Metadata        respjson.Field
		Settings        respjson.Field
		TermsOfService  respjson.Field
		UpdatedAt       respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccountV1Data) RawJSON() string { return r.JSON.raw }
func (r *AccountV1Data) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountV1DataStatusDetail struct {
	// A machine-readable code for the specific status, useful for programmatic
	// handling.
	Code string `json:"code" api:"required"`
	// A human-readable message describing the current status.
	Message string `json:"message" api:"required"`
	// A machine-readable identifier for the specific status, useful for programmatic
	// handling.
	//
	// Any of "unverified", "in_review", "pending", "stuck", "verified",
	// "failed_verification", "disabled", "terminated", "new".
	Reason string `json:"reason" api:"required"`
	// Identifies the origin of the status change (e.g., `bank_decline`, `watchtower`).
	// This helps in tracking the cause of status updates.
	//
	// Any of "watchtower".
	Source string `json:"source" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Message     respjson.Field
		Reason      respjson.Field
		Source      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccountV1DataStatusDetail) RawJSON() string { return r.JSON.raw }
func (r *AccountV1DataStatusDetail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountV1DataCapabilities struct {
	ConsentTypes  AccountV1DataCapabilitiesConsentTypes  `json:"consent_types" api:"required"`
	CustomerTypes AccountV1DataCapabilitiesCustomerTypes `json:"customer_types" api:"required"`
	PaymentTypes  AccountV1DataCapabilitiesPaymentTypes  `json:"payment_types" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ConsentTypes  respjson.Field
		CustomerTypes respjson.Field
		PaymentTypes  respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccountV1DataCapabilities) RawJSON() string { return r.JSON.raw }
func (r *AccountV1DataCapabilities) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountV1DataCapabilitiesConsentTypes struct {
	// Whether the internet payment authorization capability is enabled for the
	// account.
	Internet CapabilityV1 `json:"internet" api:"required"`
	// Whether the signed agreement payment authorization capability is enabled for the
	// account.
	SignedAgreement CapabilityV1 `json:"signed_agreement" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Internet        respjson.Field
		SignedAgreement respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccountV1DataCapabilitiesConsentTypes) RawJSON() string { return r.JSON.raw }
func (r *AccountV1DataCapabilitiesConsentTypes) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountV1DataCapabilitiesCustomerTypes struct {
	Businesses  CapabilityV1 `json:"businesses" api:"required"`
	Individuals CapabilityV1 `json:"individuals" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Businesses  respjson.Field
		Individuals respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccountV1DataCapabilitiesCustomerTypes) RawJSON() string { return r.JSON.raw }
func (r *AccountV1DataCapabilitiesCustomerTypes) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountV1DataCapabilitiesPaymentTypes struct {
	Charges CapabilityV1 `json:"charges" api:"required"`
	Payouts CapabilityV1 `json:"payouts" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Charges     respjson.Field
		Payouts     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccountV1DataCapabilitiesPaymentTypes) RawJSON() string { return r.JSON.raw }
func (r *AccountV1DataCapabilitiesPaymentTypes) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountV1DataSettings struct {
	Charges AccountV1DataSettingsCharges `json:"charges" api:"required"`
	Payouts AccountV1DataSettingsPayouts `json:"payouts" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Charges     respjson.Field
		Payouts     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccountV1DataSettings) RawJSON() string { return r.JSON.raw }
func (r *AccountV1DataSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountV1DataSettingsCharges struct {
	// The maximum dollar amount of charges in a calendar day.
	DailyAmount int64 `json:"daily_amount" api:"required"`
	// The amount of time it takes for a charge to be funded. This value is defined by
	// Straddle.
	//
	// Any of "immediate", "next_day", "one_day", "two_day", "three_day", "four_day",
	// "five_day".
	FundingTime string `json:"funding_time" api:"required"`
	// The unique identifier of the linked bank account associated with charges. This
	// value is defined by Straddle.
	LinkedBankAccountID string `json:"linked_bank_account_id" api:"required" format:"uuid"`
	// The maximum amount of a single charge.
	MaxAmount int64 `json:"max_amount" api:"required"`
	// The maximum dollar amount of charges in a calendar month.
	MonthlyAmount int64 `json:"monthly_amount" api:"required"`
	// The maximum number of charges in a calendar month.
	MonthlyCount int64 `json:"monthly_count" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DailyAmount         respjson.Field
		FundingTime         respjson.Field
		LinkedBankAccountID respjson.Field
		MaxAmount           respjson.Field
		MonthlyAmount       respjson.Field
		MonthlyCount        respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccountV1DataSettingsCharges) RawJSON() string { return r.JSON.raw }
func (r *AccountV1DataSettingsCharges) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountV1DataSettingsPayouts struct {
	// The maximum dollar amount of payouts in a day.
	DailyAmount int64 `json:"daily_amount" api:"required"`
	// The amount of time it takes for a payout to be funded. This value is defined by
	// Straddle.
	//
	// Any of "immediate", "next_day", "one_day", "two_day", "three_day", "four_day",
	// "five_day".
	FundingTime string `json:"funding_time" api:"required"`
	// The unique identifier of the linked bank account to use for payouts.
	LinkedBankAccountID string `json:"linked_bank_account_id" api:"required" format:"uuid"`
	// The maximum amount of a single payout.
	MaxAmount int64 `json:"max_amount" api:"required"`
	// The maximum dollar amount of payouts in a month.
	MonthlyAmount int64 `json:"monthly_amount" api:"required"`
	// The maximum number of payouts in a month.
	MonthlyCount int64 `json:"monthly_count" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DailyAmount         respjson.Field
		FundingTime         respjson.Field
		LinkedBankAccountID respjson.Field
		MaxAmount           respjson.Field
		MonthlyAmount       respjson.Field
		MonthlyCount        respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccountV1DataSettingsPayouts) RawJSON() string { return r.JSON.raw }
func (r *AccountV1DataSettingsPayouts) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indicates the structure of the returned content.
//
//   - "object" means the `data` field contains a single JSON object.
//   - "array" means the `data` field contains an array of objects.
//   - "error" means the `data` field contains an error object with details of the
//     issue.
//   - "none" means no data is returned.
type AccountV1ResponseType string

const (
	AccountV1ResponseTypeObject AccountV1ResponseType = "object"
	AccountV1ResponseTypeArray  AccountV1ResponseType = "array"
	AccountV1ResponseTypeError  AccountV1ResponseType = "error"
	AccountV1ResponseTypeNone   AccountV1ResponseType = "none"
)

// The address object is optional. If provided, it must be a valid address.
type AddressV1 struct {
	// Primary address line (e.g., street, PO Box).
	Address1 string `json:"address1" api:"required"`
	// City, district, suburb, town, or village.
	City string `json:"city" api:"required"`
	// Primary address line (e.g., street, PO Box).
	Line1 string `json:"line1" api:"required"`
	// Postal or ZIP code.
	PostalCode string `json:"postal_code" api:"required"`
	// Two-letter state code.
	State string `json:"state" api:"required"`
	// Zip or postal code.
	Zip string `json:"zip" api:"required"`
	// Secondary address line (e.g., apartment, suite, unit, or building).
	Address2 string `json:"address2" api:"nullable"`
	// The country of the address, in ISO 3166-1 alpha-2 format.
	Country string `json:"country" api:"nullable"`
	// Secondary address line (e.g., apartment, suite, unit, or building).
	Line2 string `json:"line2" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Address1    respjson.Field
		City        respjson.Field
		Line1       respjson.Field
		PostalCode  respjson.Field
		State       respjson.Field
		Zip         respjson.Field
		Address2    respjson.Field
		Country     respjson.Field
		Line2       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AddressV1) RawJSON() string { return r.JSON.raw }
func (r *AddressV1) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this AddressV1 to a AddressV1Param.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// AddressV1Param.Overrides()
func (r AddressV1) ToParam() AddressV1Param {
	return param.Override[AddressV1Param](json.RawMessage(r.RawJSON()))
}

// The address object is optional. If provided, it must be a valid address.
//
// The properties Address1, City, Line1, PostalCode, State, Zip are required.
type AddressV1Param struct {
	// City, district, suburb, town, or village.
	City param.Opt[string] `json:"city,omitzero" api:"required"`
	// Primary address line (e.g., street, PO Box).
	Line1 param.Opt[string] `json:"line1,omitzero" api:"required"`
	// Postal or ZIP code.
	PostalCode param.Opt[string] `json:"postal_code,omitzero" api:"required"`
	// Two-letter state code.
	State param.Opt[string] `json:"state,omitzero" api:"required"`
	// Primary address line (e.g., street, PO Box).
	Address1 string `json:"address1" api:"required"`
	// Zip or postal code.
	Zip string `json:"zip" api:"required"`
	// Secondary address line (e.g., apartment, suite, unit, or building).
	Address2 param.Opt[string] `json:"address2,omitzero"`
	// The country of the address, in ISO 3166-1 alpha-2 format.
	Country param.Opt[string] `json:"country,omitzero"`
	// Secondary address line (e.g., apartment, suite, unit, or building).
	Line2 param.Opt[string] `json:"line2,omitzero"`
	paramObj
}

func (r AddressV1Param) MarshalJSON() (data []byte, err error) {
	type shadow AddressV1Param
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AddressV1Param) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BusinessProfileV1 struct {
	// The operating or trade name of the business.
	Name string `json:"name" api:"required"`
	// URL of the business's primary marketing website.
	Website string `json:"website" api:"required" format:"uri"`
	// The address object is optional. If provided, it must be a valid address.
	Address AddressV1 `json:"address" api:"nullable"`
	// A brief description of the business and its products or services.
	Description string     `json:"description" api:"nullable"`
	Industry    IndustryV1 `json:"industry"`
	// The official registered name of the business.
	LegalName string `json:"legal_name" api:"nullable"`
	// The primary contact phone number for the business.
	Phone           string            `json:"phone" api:"nullable"`
	SupportChannels SupportChannelsV1 `json:"support_channels"`
	// The business's tax identification number (e.g., EIN in the US).
	TaxID string `json:"tax_id" api:"nullable"`
	// A description of how the business intends to use Straddle's services.
	UseCase string `json:"use_case" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name            respjson.Field
		Website         respjson.Field
		Address         respjson.Field
		Description     respjson.Field
		Industry        respjson.Field
		LegalName       respjson.Field
		Phone           respjson.Field
		SupportChannels respjson.Field
		TaxID           respjson.Field
		UseCase         respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BusinessProfileV1) RawJSON() string { return r.JSON.raw }
func (r *BusinessProfileV1) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this BusinessProfileV1 to a BusinessProfileV1Param.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// BusinessProfileV1Param.Overrides()
func (r BusinessProfileV1) ToParam() BusinessProfileV1Param {
	return param.Override[BusinessProfileV1Param](json.RawMessage(r.RawJSON()))
}

// The properties Name, Website are required.
type BusinessProfileV1Param struct {
	// The operating or trade name of the business.
	Name string `json:"name" api:"required"`
	// URL of the business's primary marketing website.
	Website string `json:"website" api:"required" format:"uri"`
	// A brief description of the business and its products or services.
	Description param.Opt[string] `json:"description,omitzero"`
	// The official registered name of the business.
	LegalName param.Opt[string] `json:"legal_name,omitzero"`
	// The primary contact phone number for the business.
	Phone param.Opt[string] `json:"phone,omitzero"`
	// The business's tax identification number (e.g., EIN in the US).
	TaxID param.Opt[string] `json:"tax_id,omitzero"`
	// A description of how the business intends to use Straddle's services.
	UseCase param.Opt[string] `json:"use_case,omitzero"`
	// The address object is optional. If provided, it must be a valid address.
	Address         AddressV1Param         `json:"address,omitzero"`
	Industry        IndustryV1Param        `json:"industry,omitzero"`
	SupportChannels SupportChannelsV1Param `json:"support_channels,omitzero"`
	paramObj
}

func (r BusinessProfileV1Param) MarshalJSON() (data []byte, err error) {
	type shadow BusinessProfileV1Param
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BusinessProfileV1Param) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CapabilityV1 struct {
	// Any of "active", "inactive".
	CapabilityStatus CapabilityV1CapabilityStatus `json:"capability_status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CapabilityStatus respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CapabilityV1) RawJSON() string { return r.JSON.raw }
func (r *CapabilityV1) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CapabilityV1CapabilityStatus string

const (
	CapabilityV1CapabilityStatusActive   CapabilityV1CapabilityStatus = "active"
	CapabilityV1CapabilityStatusInactive CapabilityV1CapabilityStatus = "inactive"
)

type IndustryV1 struct {
	// The general category of the industry. Required if not providing MCC.
	Category string `json:"category" api:"nullable"`
	// The Merchant Category Code (MCC) that best describes the business. Optional.
	Mcc string `json:"mcc" api:"nullable"`
	// The specific sector within the industry category. Required if not providing MCC.
	Sector string `json:"sector" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Category    respjson.Field
		Mcc         respjson.Field
		Sector      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IndustryV1) RawJSON() string { return r.JSON.raw }
func (r *IndustryV1) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this IndustryV1 to a IndustryV1Param.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// IndustryV1Param.Overrides()
func (r IndustryV1) ToParam() IndustryV1Param {
	return param.Override[IndustryV1Param](json.RawMessage(r.RawJSON()))
}

type IndustryV1Param struct {
	// The general category of the industry. Required if not providing MCC.
	Category param.Opt[string] `json:"category,omitzero"`
	// The Merchant Category Code (MCC) that best describes the business. Optional.
	Mcc param.Opt[string] `json:"mcc,omitzero"`
	// The specific sector within the industry category. Required if not providing MCC.
	Sector param.Opt[string] `json:"sector,omitzero"`
	paramObj
}

func (r IndustryV1Param) MarshalJSON() (data []byte, err error) {
	type shadow IndustryV1Param
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *IndustryV1Param) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SupportChannelsV1 struct {
	// The email address for customer support inquiries.
	Email string `json:"email" api:"nullable" format:"email"`
	// The phone number for customer support.
	Phone string `json:"phone" api:"nullable"`
	// The URL of the business's customer support page or contact form.
	URL string `json:"url" api:"nullable" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Email       respjson.Field
		Phone       respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SupportChannelsV1) RawJSON() string { return r.JSON.raw }
func (r *SupportChannelsV1) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this SupportChannelsV1 to a SupportChannelsV1Param.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// SupportChannelsV1Param.Overrides()
func (r SupportChannelsV1) ToParam() SupportChannelsV1Param {
	return param.Override[SupportChannelsV1Param](json.RawMessage(r.RawJSON()))
}

type SupportChannelsV1Param struct {
	// The email address for customer support inquiries.
	Email param.Opt[string] `json:"email,omitzero" format:"email"`
	// The phone number for customer support.
	Phone param.Opt[string] `json:"phone,omitzero"`
	// The URL of the business's customer support page or contact form.
	URL param.Opt[string] `json:"url,omitzero" format:"uri"`
	paramObj
}

func (r SupportChannelsV1Param) MarshalJSON() (data []byte, err error) {
	type shadow SupportChannelsV1Param
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SupportChannelsV1Param) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TermsOfServiceV1 struct {
	// The datetime of when the terms of service were accepted, in ISO 8601 format.
	AcceptedDate time.Time `json:"accepted_date" api:"required" format:"date-time"`
	// The type or version of the agreement accepted. Use `embedded` unless your
	// platform was specifically enabled for `direct` agreements.
	//
	// Any of "embedded", "direct".
	AgreementType TermsOfServiceV1AgreementType `json:"agreement_type" api:"required"`
	// The URL where the full text of the accepted agreement can be found.
	AgreementURL string `json:"agreement_url" api:"required"`
	// The IP address from which the terms of service were accepted.
	AcceptedIP string `json:"accepted_ip" api:"nullable"`
	// The user agent string of the browser or application used to accept the terms.
	AcceptedUserAgent string `json:"accepted_user_agent" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AcceptedDate      respjson.Field
		AgreementType     respjson.Field
		AgreementURL      respjson.Field
		AcceptedIP        respjson.Field
		AcceptedUserAgent respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TermsOfServiceV1) RawJSON() string { return r.JSON.raw }
func (r *TermsOfServiceV1) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this TermsOfServiceV1 to a TermsOfServiceV1Param.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// TermsOfServiceV1Param.Overrides()
func (r TermsOfServiceV1) ToParam() TermsOfServiceV1Param {
	return param.Override[TermsOfServiceV1Param](json.RawMessage(r.RawJSON()))
}

// The type or version of the agreement accepted. Use `embedded` unless your
// platform was specifically enabled for `direct` agreements.
type TermsOfServiceV1AgreementType string

const (
	TermsOfServiceV1AgreementTypeEmbedded TermsOfServiceV1AgreementType = "embedded"
	TermsOfServiceV1AgreementTypeDirect   TermsOfServiceV1AgreementType = "direct"
)

// The properties AcceptedDate, AgreementType, AgreementURL are required.
type TermsOfServiceV1Param struct {
	// The URL where the full text of the accepted agreement can be found.
	AgreementURL param.Opt[string] `json:"agreement_url,omitzero" api:"required"`
	// The datetime of when the terms of service were accepted, in ISO 8601 format.
	AcceptedDate time.Time `json:"accepted_date" api:"required" format:"date-time"`
	// The type or version of the agreement accepted. Use `embedded` unless your
	// platform was specifically enabled for `direct` agreements.
	//
	// Any of "embedded", "direct".
	AgreementType TermsOfServiceV1AgreementType `json:"agreement_type,omitzero" api:"required"`
	// The IP address from which the terms of service were accepted.
	AcceptedIP param.Opt[string] `json:"accepted_ip,omitzero"`
	// The user agent string of the browser or application used to accept the terms.
	AcceptedUserAgent param.Opt[string] `json:"accepted_user_agent,omitzero"`
	paramObj
}

func (r TermsOfServiceV1Param) MarshalJSON() (data []byte, err error) {
	type shadow TermsOfServiceV1Param
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TermsOfServiceV1Param) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbedAccountNewParams struct {
	// The access level granted to the account. This is determined by your platform
	// configuration. Use `standard` unless instructed otherwise by Straddle.
	//
	// Any of "standard", "managed".
	AccessLevel EmbedAccountNewParamsAccessLevel `json:"access_level,omitzero" api:"required"`
	// The type of account to be created. Currently, only `business` is supported.
	//
	// Any of "business".
	AccountType     EmbedAccountNewParamsAccountType `json:"account_type,omitzero" api:"required"`
	BusinessProfile BusinessProfileV1Param           `json:"business_profile,omitzero" api:"required"`
	// The unique identifier of the organization related to this account.
	OrganizationID string `json:"organization_id" api:"required" format:"uuid"`
	// Unique identifier for the account in your database, used for cross-referencing
	// between Straddle and your systems.
	ExternalID     param.Opt[string] `json:"external_id,omitzero"`
	CorrelationID  param.Opt[string] `header:"correlation-id,omitzero" json:"-"`
	IdempotencyKey param.Opt[string] `header:"idempotency-key,omitzero" json:"-"`
	RequestID      param.Opt[string] `header:"request-id,omitzero" json:"-"`
	// Up to 20 additional user-defined key-value pairs. Useful for storing additional
	// information about the account in a structured format.
	Metadata map[string]string `json:"metadata,omitzero"`
	paramObj
}

func (r EmbedAccountNewParams) MarshalJSON() (data []byte, err error) {
	type shadow EmbedAccountNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EmbedAccountNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The access level granted to the account. This is determined by your platform
// configuration. Use `standard` unless instructed otherwise by Straddle.
type EmbedAccountNewParamsAccessLevel string

const (
	EmbedAccountNewParamsAccessLevelStandard EmbedAccountNewParamsAccessLevel = "standard"
	EmbedAccountNewParamsAccessLevelManaged  EmbedAccountNewParamsAccessLevel = "managed"
)

// The type of account to be created. Currently, only `business` is supported.
type EmbedAccountNewParamsAccountType string

const (
	EmbedAccountNewParamsAccountTypeBusiness EmbedAccountNewParamsAccountType = "business"
)

type EmbedAccountUpdateParams struct {
	BusinessProfile BusinessProfileV1Param `json:"business_profile,omitzero" api:"required"`
	// Unique identifier for the account in your database, used for cross-referencing
	// between Straddle and your systems.
	ExternalID     param.Opt[string] `json:"external_id,omitzero"`
	CorrelationID  param.Opt[string] `header:"correlation-id,omitzero" json:"-"`
	IdempotencyKey param.Opt[string] `header:"idempotency-key,omitzero" json:"-"`
	RequestID      param.Opt[string] `header:"request-id,omitzero" json:"-"`
	// Up to 20 additional user-defined key-value pairs. Useful for storing additional
	// information about the account in a structured format.
	Metadata map[string]string `json:"metadata,omitzero"`
	paramObj
}

func (r EmbedAccountUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow EmbedAccountUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EmbedAccountUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbedAccountListParams struct {
	ExternalID param.Opt[string] `query:"external_id,omitzero" json:"-"`
	// Results page number. Starts at page 1. Default value: 1
	PageNumber param.Opt[int64] `query:"page_number,omitzero" json:"-"`
	// Page size. Default value: 100. Max value: 1000
	PageSize   param.Opt[int64]  `query:"page_size,omitzero" json:"-"`
	SearchText param.Opt[string] `query:"search_text,omitzero" json:"-"`
	// Sort By. Default value: 'id'.
	SortBy        param.Opt[string] `query:"sort_by,omitzero" json:"-"`
	CorrelationID param.Opt[string] `header:"correlation-id,omitzero" json:"-"`
	RequestID     param.Opt[string] `header:"request-id,omitzero" json:"-"`
	// Sort Order. Default value: 'asc'.
	//
	// Any of "asc", "desc".
	SortOrder EmbedAccountListParamsSortOrder `query:"sort_order,omitzero" json:"-"`
	// Any of "created", "onboarding", "active", "rejected", "inactive".
	Status EmbedAccountListParamsStatus `query:"status,omitzero" json:"-"`
	// Any of "business".
	Type EmbedAccountListParamsType `query:"type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [EmbedAccountListParams]'s query parameters as `url.Values`.
func (r EmbedAccountListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Sort Order. Default value: 'asc'.
type EmbedAccountListParamsSortOrder string

const (
	EmbedAccountListParamsSortOrderAsc  EmbedAccountListParamsSortOrder = "asc"
	EmbedAccountListParamsSortOrderDesc EmbedAccountListParamsSortOrder = "desc"
)

type EmbedAccountListParamsStatus string

const (
	EmbedAccountListParamsStatusCreated    EmbedAccountListParamsStatus = "created"
	EmbedAccountListParamsStatusOnboarding EmbedAccountListParamsStatus = "onboarding"
	EmbedAccountListParamsStatusActive     EmbedAccountListParamsStatus = "active"
	EmbedAccountListParamsStatusRejected   EmbedAccountListParamsStatus = "rejected"
	EmbedAccountListParamsStatusInactive   EmbedAccountListParamsStatus = "inactive"
)

type EmbedAccountListParamsType string

const (
	EmbedAccountListParamsTypeBusiness EmbedAccountListParamsType = "business"
)

type EmbedAccountGetParams struct {
	CorrelationID param.Opt[string] `header:"correlation-id,omitzero" json:"-"`
	RequestID     param.Opt[string] `header:"request-id,omitzero" json:"-"`
	paramObj
}

type EmbedAccountOnboardParams struct {
	TermsOfService TermsOfServiceV1Param `json:"terms_of_service,omitzero" api:"required"`
	CorrelationID  param.Opt[string]     `header:"correlation-id,omitzero" json:"-"`
	IdempotencyKey param.Opt[string]     `header:"idempotency-key,omitzero" json:"-"`
	RequestID      param.Opt[string]     `header:"request-id,omitzero" json:"-"`
	paramObj
}

func (r EmbedAccountOnboardParams) MarshalJSON() (data []byte, err error) {
	type shadow EmbedAccountOnboardParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EmbedAccountOnboardParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbedAccountSimulateParams struct {
	CorrelationID  param.Opt[string] `header:"correlation-id,omitzero" json:"-"`
	IdempotencyKey param.Opt[string] `header:"idempotency-key,omitzero" json:"-"`
	RequestID      param.Opt[string] `header:"request-id,omitzero" json:"-"`
	// Any of "onboarding", "active".
	FinalStatus EmbedAccountSimulateParamsFinalStatus `query:"final_status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [EmbedAccountSimulateParams]'s query parameters as
// `url.Values`.
func (r EmbedAccountSimulateParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type EmbedAccountSimulateParamsFinalStatus string

const (
	EmbedAccountSimulateParamsFinalStatusOnboarding EmbedAccountSimulateParamsFinalStatus = "onboarding"
	EmbedAccountSimulateParamsFinalStatusActive     EmbedAccountSimulateParamsFinalStatus = "active"
)
