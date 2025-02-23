// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package straddle

import (
	"context"
	"errors"
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

// EmbedAccountService contains methods and other services that help with
// interacting with the straddle API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEmbedAccountService] method instead.
type EmbedAccountService struct {
	Options            []option.RequestOption
	CapabilityRequests *EmbedAccountCapabilityRequestService
}

// NewEmbedAccountService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewEmbedAccountService(opts ...option.RequestOption) (r *EmbedAccountService) {
	r = &EmbedAccountService{}
	r.Options = opts
	r.CapabilityRequests = NewEmbedAccountCapabilityRequestService(opts...)
	return
}

// Creates a new account associated with your Straddle platform integration. This
// endpoint allows you to set up an account with specified details, including
// business information and access levels.
func (r *EmbedAccountService) New(ctx context.Context, params EmbedAccountNewParams, opts ...option.RequestOption) (res *AccountV1, err error) {
	if params.CorrelationID.Present {
		opts = append(opts, option.WithHeader("correlation-id", fmt.Sprintf("%s", params.CorrelationID)))
	}
	if params.RequestID.Present {
		opts = append(opts, option.WithHeader("request-id", fmt.Sprintf("%s", params.RequestID)))
	}
	opts = append(r.Options[:], opts...)
	path := "v1/accounts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Updates an existing account's information. This endpoint allows you to update
// various account details during onboarding or after the account has been created.
func (r *EmbedAccountService) Update(ctx context.Context, accountID string, params EmbedAccountUpdateParams, opts ...option.RequestOption) (res *AccountV1, err error) {
	if params.CorrelationID.Present {
		opts = append(opts, option.WithHeader("correlation-id", fmt.Sprintf("%s", params.CorrelationID)))
	}
	if params.RequestID.Present {
		opts = append(opts, option.WithHeader("request-id", fmt.Sprintf("%s", params.RequestID)))
	}
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("v1/accounts/%s", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

// Returns a list of accounts associated with your Straddle platform integration.
// The accounts are returned sorted by creation date, with the most recently
// created accounts appearing first. This endpoint supports advanced sorting and
// filtering options.
func (r *EmbedAccountService) List(ctx context.Context, params EmbedAccountListParams, opts ...option.RequestOption) (res *pagination.PageNumberSchema[AccountPagedV1Data], err error) {
	var raw *http.Response
	if params.CorrelationID.Present {
		opts = append(opts, option.WithHeader("correlation-id", fmt.Sprintf("%s", params.CorrelationID)))
	}
	if params.RequestID.Present {
		opts = append(opts, option.WithHeader("request-id", fmt.Sprintf("%s", params.RequestID)))
	}
	opts = append(r.Options[:], opts...)
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
	if query.CorrelationID.Present {
		opts = append(opts, option.WithHeader("correlation-id", fmt.Sprintf("%s", query.CorrelationID)))
	}
	if query.RequestID.Present {
		opts = append(opts, option.WithHeader("request-id", fmt.Sprintf("%s", query.RequestID)))
	}
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("v1/accounts/%s", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Initiates the onboarding process for a new account. This endpoint can only be
// used for accounts where at least one representative and one bank account have
// already been created.
func (r *EmbedAccountService) Onboard(ctx context.Context, accountID string, params EmbedAccountOnboardParams, opts ...option.RequestOption) (res *AccountV1, err error) {
	if params.CorrelationID.Present {
		opts = append(opts, option.WithHeader("correlation-id", fmt.Sprintf("%s", params.CorrelationID)))
	}
	if params.RequestID.Present {
		opts = append(opts, option.WithHeader("request-id", fmt.Sprintf("%s", params.RequestID)))
	}
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("v1/accounts/%s/onboard", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Simulte the status transitions for sandbox accounts. This endpoint can only be
// used for sandbox accounts.
func (r *EmbedAccountService) Simulate(ctx context.Context, accountID string, params EmbedAccountSimulateParams, opts ...option.RequestOption) (res *AccountV1, err error) {
	if params.CorrelationID.Present {
		opts = append(opts, option.WithHeader("correlation-id", fmt.Sprintf("%s", params.CorrelationID)))
	}
	if params.RequestID.Present {
		opts = append(opts, option.WithHeader("request-id", fmt.Sprintf("%s", params.RequestID)))
	}
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("v1/accounts/%s/simulate", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type AccountPagedV1 struct {
	Data []AccountPagedV1Data `json:"data,required"`
	// Metadata about the API request, including an identifier, timestamp, and
	// pagination details.
	Meta shared.PagedResponseMetadata `json:"meta,required"`
	// Indicates the structure of the returned content.
	//
	//   - "object" means the `data` field contains a single JSON object.
	//   - "array" means the `data` field contains an array of objects.
	//   - "error" means the `data` field contains an error object with details of the
	//     issue.
	//   - "none" means no data is returned.
	ResponseType AccountPagedV1ResponseType `json:"response_type,required"`
	JSON         accountPagedV1JSON         `json:"-"`
}

// accountPagedV1JSON contains the JSON metadata for the struct [AccountPagedV1]
type accountPagedV1JSON struct {
	Data         apijson.Field
	Meta         apijson.Field
	ResponseType apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountPagedV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountPagedV1JSON) RawJSON() string {
	return r.raw
}

type AccountPagedV1Data struct {
	// Unique identifier for the account.
	ID string `json:"id,required" format:"uuid"`
	// The access level granted to the account. This is determined by your platform
	// configuration. Use `standard` unless instructed otherwise by Straddle.
	AccessLevel AccountPagedV1DataAccessLevel `json:"access_level,required"`
	// The unique identifier of the organization this account belongs to.
	OrganizationID string `json:"organization_id,required" format:"uuid"`
	// The current status of the account (e.g., 'active', 'inactive', 'pending').
	Status       AccountPagedV1DataStatus       `json:"status,required"`
	StatusDetail AccountPagedV1DataStatusDetail `json:"status_detail,required"`
	// The type of account (e.g., 'individual', 'business').
	Type            AccountPagedV1DataType         `json:"type,required"`
	BusinessProfile BusinessProfileV1              `json:"business_profile"`
	Capabilities    AccountPagedV1DataCapabilities `json:"capabilities"`
	// Timestamp of when the account was created.
	CreatedAt time.Time `json:"created_at,nullable" format:"date-time"`
	// Unique identifier for the account in your database, used for cross-referencing
	// between Straddle and your systems.
	ExternalID string `json:"external_id,nullable"`
	// Up to 20 additional user-defined key-value pairs. Useful for storing additional
	// information about the account in a structured format.
	Metadata       map[string]string          `json:"metadata,nullable"`
	Settings       AccountPagedV1DataSettings `json:"settings"`
	TermsOfService TermsOfServiceV1           `json:"terms_of_service"`
	// Timestamp of the most recent update to the account.
	UpdatedAt time.Time              `json:"updated_at,nullable" format:"date-time"`
	JSON      accountPagedV1DataJSON `json:"-"`
}

// accountPagedV1DataJSON contains the JSON metadata for the struct
// [AccountPagedV1Data]
type accountPagedV1DataJSON struct {
	ID              apijson.Field
	AccessLevel     apijson.Field
	OrganizationID  apijson.Field
	Status          apijson.Field
	StatusDetail    apijson.Field
	Type            apijson.Field
	BusinessProfile apijson.Field
	Capabilities    apijson.Field
	CreatedAt       apijson.Field
	ExternalID      apijson.Field
	Metadata        apijson.Field
	Settings        apijson.Field
	TermsOfService  apijson.Field
	UpdatedAt       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccountPagedV1Data) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountPagedV1DataJSON) RawJSON() string {
	return r.raw
}

// The access level granted to the account. This is determined by your platform
// configuration. Use `standard` unless instructed otherwise by Straddle.
type AccountPagedV1DataAccessLevel string

const (
	AccountPagedV1DataAccessLevelStandard AccountPagedV1DataAccessLevel = "standard"
	AccountPagedV1DataAccessLevelManaged  AccountPagedV1DataAccessLevel = "managed"
)

func (r AccountPagedV1DataAccessLevel) IsKnown() bool {
	switch r {
	case AccountPagedV1DataAccessLevelStandard, AccountPagedV1DataAccessLevelManaged:
		return true
	}
	return false
}

// The current status of the account (e.g., 'active', 'inactive', 'pending').
type AccountPagedV1DataStatus string

const (
	AccountPagedV1DataStatusCreated    AccountPagedV1DataStatus = "created"
	AccountPagedV1DataStatusOnboarding AccountPagedV1DataStatus = "onboarding"
	AccountPagedV1DataStatusActive     AccountPagedV1DataStatus = "active"
	AccountPagedV1DataStatusRejected   AccountPagedV1DataStatus = "rejected"
	AccountPagedV1DataStatusInactive   AccountPagedV1DataStatus = "inactive"
)

func (r AccountPagedV1DataStatus) IsKnown() bool {
	switch r {
	case AccountPagedV1DataStatusCreated, AccountPagedV1DataStatusOnboarding, AccountPagedV1DataStatusActive, AccountPagedV1DataStatusRejected, AccountPagedV1DataStatusInactive:
		return true
	}
	return false
}

type AccountPagedV1DataStatusDetail struct {
	// A machine-readable code for the specific status, useful for programmatic
	// handling.
	Code string `json:"code,required"`
	// A human-readable message describing the current status.
	Message string `json:"message,required"`
	// A machine-readable identifier for the specific status, useful for programmatic
	// handling.
	Reason AccountPagedV1DataStatusDetailReason `json:"reason,required"`
	// Identifies the origin of the status change (e.g., `bank_decline`, `watchtower`).
	// This helps in tracking the cause of status updates.
	Source AccountPagedV1DataStatusDetailSource `json:"source,required"`
	JSON   accountPagedV1DataStatusDetailJSON   `json:"-"`
}

// accountPagedV1DataStatusDetailJSON contains the JSON metadata for the struct
// [AccountPagedV1DataStatusDetail]
type accountPagedV1DataStatusDetailJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	Reason      apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountPagedV1DataStatusDetail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountPagedV1DataStatusDetailJSON) RawJSON() string {
	return r.raw
}

// A machine-readable identifier for the specific status, useful for programmatic
// handling.
type AccountPagedV1DataStatusDetailReason string

const (
	AccountPagedV1DataStatusDetailReasonUnverified         AccountPagedV1DataStatusDetailReason = "unverified"
	AccountPagedV1DataStatusDetailReasonInReview           AccountPagedV1DataStatusDetailReason = "in_review"
	AccountPagedV1DataStatusDetailReasonPending            AccountPagedV1DataStatusDetailReason = "pending"
	AccountPagedV1DataStatusDetailReasonStuck              AccountPagedV1DataStatusDetailReason = "stuck"
	AccountPagedV1DataStatusDetailReasonVerified           AccountPagedV1DataStatusDetailReason = "verified"
	AccountPagedV1DataStatusDetailReasonFailedVerification AccountPagedV1DataStatusDetailReason = "failed_verification"
	AccountPagedV1DataStatusDetailReasonDisabled           AccountPagedV1DataStatusDetailReason = "disabled"
	AccountPagedV1DataStatusDetailReasonTerminated         AccountPagedV1DataStatusDetailReason = "terminated"
	AccountPagedV1DataStatusDetailReasonNew                AccountPagedV1DataStatusDetailReason = "new"
)

func (r AccountPagedV1DataStatusDetailReason) IsKnown() bool {
	switch r {
	case AccountPagedV1DataStatusDetailReasonUnverified, AccountPagedV1DataStatusDetailReasonInReview, AccountPagedV1DataStatusDetailReasonPending, AccountPagedV1DataStatusDetailReasonStuck, AccountPagedV1DataStatusDetailReasonVerified, AccountPagedV1DataStatusDetailReasonFailedVerification, AccountPagedV1DataStatusDetailReasonDisabled, AccountPagedV1DataStatusDetailReasonTerminated, AccountPagedV1DataStatusDetailReasonNew:
		return true
	}
	return false
}

// Identifies the origin of the status change (e.g., `bank_decline`, `watchtower`).
// This helps in tracking the cause of status updates.
type AccountPagedV1DataStatusDetailSource string

const (
	AccountPagedV1DataStatusDetailSourceWatchtower AccountPagedV1DataStatusDetailSource = "watchtower"
)

func (r AccountPagedV1DataStatusDetailSource) IsKnown() bool {
	switch r {
	case AccountPagedV1DataStatusDetailSourceWatchtower:
		return true
	}
	return false
}

// The type of account (e.g., 'individual', 'business').
type AccountPagedV1DataType string

const (
	AccountPagedV1DataTypeBusiness AccountPagedV1DataType = "business"
)

func (r AccountPagedV1DataType) IsKnown() bool {
	switch r {
	case AccountPagedV1DataTypeBusiness:
		return true
	}
	return false
}

type AccountPagedV1DataCapabilities struct {
	ConsentTypes  AccountPagedV1DataCapabilitiesConsentTypes  `json:"consent_types,required"`
	CustomerTypes AccountPagedV1DataCapabilitiesCustomerTypes `json:"customer_types,required"`
	PaymentTypes  AccountPagedV1DataCapabilitiesPaymentTypes  `json:"payment_types,required"`
	JSON          accountPagedV1DataCapabilitiesJSON          `json:"-"`
}

// accountPagedV1DataCapabilitiesJSON contains the JSON metadata for the struct
// [AccountPagedV1DataCapabilities]
type accountPagedV1DataCapabilitiesJSON struct {
	ConsentTypes  apijson.Field
	CustomerTypes apijson.Field
	PaymentTypes  apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccountPagedV1DataCapabilities) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountPagedV1DataCapabilitiesJSON) RawJSON() string {
	return r.raw
}

type AccountPagedV1DataCapabilitiesConsentTypes struct {
	// Whether the internet payment authorization capability is enabled for the
	// account.
	Internet CapabilityV1 `json:"internet,required"`
	// Whether the signed agreement payment authorization capability is enabled for the
	// account.
	SignedAgreement CapabilityV1                                   `json:"signed_agreement,required"`
	JSON            accountPagedV1DataCapabilitiesConsentTypesJSON `json:"-"`
}

// accountPagedV1DataCapabilitiesConsentTypesJSON contains the JSON metadata for
// the struct [AccountPagedV1DataCapabilitiesConsentTypes]
type accountPagedV1DataCapabilitiesConsentTypesJSON struct {
	Internet        apijson.Field
	SignedAgreement apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccountPagedV1DataCapabilitiesConsentTypes) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountPagedV1DataCapabilitiesConsentTypesJSON) RawJSON() string {
	return r.raw
}

type AccountPagedV1DataCapabilitiesCustomerTypes struct {
	Businesses  CapabilityV1                                    `json:"businesses,required"`
	Individuals CapabilityV1                                    `json:"individuals,required"`
	JSON        accountPagedV1DataCapabilitiesCustomerTypesJSON `json:"-"`
}

// accountPagedV1DataCapabilitiesCustomerTypesJSON contains the JSON metadata for
// the struct [AccountPagedV1DataCapabilitiesCustomerTypes]
type accountPagedV1DataCapabilitiesCustomerTypesJSON struct {
	Businesses  apijson.Field
	Individuals apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountPagedV1DataCapabilitiesCustomerTypes) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountPagedV1DataCapabilitiesCustomerTypesJSON) RawJSON() string {
	return r.raw
}

type AccountPagedV1DataCapabilitiesPaymentTypes struct {
	Charges CapabilityV1                                   `json:"charges,required"`
	Payouts CapabilityV1                                   `json:"payouts,required"`
	JSON    accountPagedV1DataCapabilitiesPaymentTypesJSON `json:"-"`
}

// accountPagedV1DataCapabilitiesPaymentTypesJSON contains the JSON metadata for
// the struct [AccountPagedV1DataCapabilitiesPaymentTypes]
type accountPagedV1DataCapabilitiesPaymentTypesJSON struct {
	Charges     apijson.Field
	Payouts     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountPagedV1DataCapabilitiesPaymentTypes) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountPagedV1DataCapabilitiesPaymentTypesJSON) RawJSON() string {
	return r.raw
}

type AccountPagedV1DataSettings struct {
	Charges AccountPagedV1DataSettingsCharges `json:"charges,required"`
	Payouts AccountPagedV1DataSettingsPayouts `json:"payouts,required"`
	JSON    accountPagedV1DataSettingsJSON    `json:"-"`
}

// accountPagedV1DataSettingsJSON contains the JSON metadata for the struct
// [AccountPagedV1DataSettings]
type accountPagedV1DataSettingsJSON struct {
	Charges     apijson.Field
	Payouts     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountPagedV1DataSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountPagedV1DataSettingsJSON) RawJSON() string {
	return r.raw
}

type AccountPagedV1DataSettingsCharges struct {
	// The maximum dollar amount of charges in a calendar day.
	DailyAmount int64 `json:"daily_amount,required"`
	// The amount of time it takes for a charge to be funded. This value is defined by
	// Straddle.
	FundingTime AccountPagedV1DataSettingsChargesFundingTime `json:"funding_time,required"`
	// The unique identifier of the linked bank account associated with charges. This
	// value is defined by Straddle.
	LinkedBankAccountID string `json:"linked_bank_account_id,required" format:"uuid"`
	// The maximum amount of a single charge.
	MaxAmount int64 `json:"max_amount,required"`
	// The maximum dollar amount of charges in a calendar month.
	MonthlyAmount int64 `json:"monthly_amount,required"`
	// The maximum number of charges in a calendar month.
	MonthlyCount int64                                 `json:"monthly_count,required"`
	JSON         accountPagedV1DataSettingsChargesJSON `json:"-"`
}

// accountPagedV1DataSettingsChargesJSON contains the JSON metadata for the struct
// [AccountPagedV1DataSettingsCharges]
type accountPagedV1DataSettingsChargesJSON struct {
	DailyAmount         apijson.Field
	FundingTime         apijson.Field
	LinkedBankAccountID apijson.Field
	MaxAmount           apijson.Field
	MonthlyAmount       apijson.Field
	MonthlyCount        apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *AccountPagedV1DataSettingsCharges) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountPagedV1DataSettingsChargesJSON) RawJSON() string {
	return r.raw
}

// The amount of time it takes for a charge to be funded. This value is defined by
// Straddle.
type AccountPagedV1DataSettingsChargesFundingTime string

const (
	AccountPagedV1DataSettingsChargesFundingTimeImmediate AccountPagedV1DataSettingsChargesFundingTime = "immediate"
	AccountPagedV1DataSettingsChargesFundingTimeNextDay   AccountPagedV1DataSettingsChargesFundingTime = "next_day"
	AccountPagedV1DataSettingsChargesFundingTimeOneDay    AccountPagedV1DataSettingsChargesFundingTime = "one_day"
	AccountPagedV1DataSettingsChargesFundingTimeTwoDay    AccountPagedV1DataSettingsChargesFundingTime = "two_day"
	AccountPagedV1DataSettingsChargesFundingTimeThreeDay  AccountPagedV1DataSettingsChargesFundingTime = "three_day"
)

func (r AccountPagedV1DataSettingsChargesFundingTime) IsKnown() bool {
	switch r {
	case AccountPagedV1DataSettingsChargesFundingTimeImmediate, AccountPagedV1DataSettingsChargesFundingTimeNextDay, AccountPagedV1DataSettingsChargesFundingTimeOneDay, AccountPagedV1DataSettingsChargesFundingTimeTwoDay, AccountPagedV1DataSettingsChargesFundingTimeThreeDay:
		return true
	}
	return false
}

type AccountPagedV1DataSettingsPayouts struct {
	// The maximum dollar amount of payouts in a day.
	DailyAmount int64 `json:"daily_amount,required"`
	// The amount of time it takes for a payout to be funded. This value is defined by
	// Straddle.
	FundingTime AccountPagedV1DataSettingsPayoutsFundingTime `json:"funding_time,required"`
	// The unique identifier of the linked bank account to use for payouts.
	LinkedBankAccountID string `json:"linked_bank_account_id,required" format:"uuid"`
	// The maximum amount of a single payout.
	MaxAmount int64 `json:"max_amount,required"`
	// The maximum dollar amount of payouts in a month.
	MonthlyAmount int64 `json:"monthly_amount,required"`
	// The maximum number of payouts in a month.
	MonthlyCount int64                                 `json:"monthly_count,required"`
	JSON         accountPagedV1DataSettingsPayoutsJSON `json:"-"`
}

// accountPagedV1DataSettingsPayoutsJSON contains the JSON metadata for the struct
// [AccountPagedV1DataSettingsPayouts]
type accountPagedV1DataSettingsPayoutsJSON struct {
	DailyAmount         apijson.Field
	FundingTime         apijson.Field
	LinkedBankAccountID apijson.Field
	MaxAmount           apijson.Field
	MonthlyAmount       apijson.Field
	MonthlyCount        apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *AccountPagedV1DataSettingsPayouts) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountPagedV1DataSettingsPayoutsJSON) RawJSON() string {
	return r.raw
}

// The amount of time it takes for a payout to be funded. This value is defined by
// Straddle.
type AccountPagedV1DataSettingsPayoutsFundingTime string

const (
	AccountPagedV1DataSettingsPayoutsFundingTimeImmediate AccountPagedV1DataSettingsPayoutsFundingTime = "immediate"
	AccountPagedV1DataSettingsPayoutsFundingTimeNextDay   AccountPagedV1DataSettingsPayoutsFundingTime = "next_day"
	AccountPagedV1DataSettingsPayoutsFundingTimeOneDay    AccountPagedV1DataSettingsPayoutsFundingTime = "one_day"
	AccountPagedV1DataSettingsPayoutsFundingTimeTwoDay    AccountPagedV1DataSettingsPayoutsFundingTime = "two_day"
	AccountPagedV1DataSettingsPayoutsFundingTimeThreeDay  AccountPagedV1DataSettingsPayoutsFundingTime = "three_day"
)

func (r AccountPagedV1DataSettingsPayoutsFundingTime) IsKnown() bool {
	switch r {
	case AccountPagedV1DataSettingsPayoutsFundingTimeImmediate, AccountPagedV1DataSettingsPayoutsFundingTimeNextDay, AccountPagedV1DataSettingsPayoutsFundingTimeOneDay, AccountPagedV1DataSettingsPayoutsFundingTimeTwoDay, AccountPagedV1DataSettingsPayoutsFundingTimeThreeDay:
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
type AccountPagedV1ResponseType string

const (
	AccountPagedV1ResponseTypeObject AccountPagedV1ResponseType = "object"
	AccountPagedV1ResponseTypeArray  AccountPagedV1ResponseType = "array"
	AccountPagedV1ResponseTypeError  AccountPagedV1ResponseType = "error"
	AccountPagedV1ResponseTypeNone   AccountPagedV1ResponseType = "none"
)

func (r AccountPagedV1ResponseType) IsKnown() bool {
	switch r {
	case AccountPagedV1ResponseTypeObject, AccountPagedV1ResponseTypeArray, AccountPagedV1ResponseTypeError, AccountPagedV1ResponseTypeNone:
		return true
	}
	return false
}

type AccountV1 struct {
	Data AccountV1Data `json:"data,required"`
	// Metadata about the API request, including an identifier and timestamp.
	Meta shared.ResponseMetadata `json:"meta,required"`
	// Indicates the structure of the returned content.
	//
	//   - "object" means the `data` field contains a single JSON object.
	//   - "array" means the `data` field contains an array of objects.
	//   - "error" means the `data` field contains an error object with details of the
	//     issue.
	//   - "none" means no data is returned.
	ResponseType AccountV1ResponseType `json:"response_type,required"`
	JSON         accountV1JSON         `json:"-"`
}

// accountV1JSON contains the JSON metadata for the struct [AccountV1]
type accountV1JSON struct {
	Data         apijson.Field
	Meta         apijson.Field
	ResponseType apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountV1JSON) RawJSON() string {
	return r.raw
}

type AccountV1Data struct {
	// Unique identifier for the account.
	ID string `json:"id,required" format:"uuid"`
	// The access level granted to the account. This is determined by your platform
	// configuration. Use `standard` unless instructed otherwise by Straddle.
	AccessLevel AccountV1DataAccessLevel `json:"access_level,required"`
	// The unique identifier of the organization this account belongs to.
	OrganizationID string `json:"organization_id,required" format:"uuid"`
	// The current status of the account (e.g., 'active', 'inactive', 'pending').
	Status       AccountV1DataStatus       `json:"status,required"`
	StatusDetail AccountV1DataStatusDetail `json:"status_detail,required"`
	// The type of account (e.g., 'individual', 'business').
	Type            AccountV1DataType         `json:"type,required"`
	BusinessProfile BusinessProfileV1         `json:"business_profile"`
	Capabilities    AccountV1DataCapabilities `json:"capabilities"`
	// Timestamp of when the account was created.
	CreatedAt time.Time `json:"created_at,nullable" format:"date-time"`
	// Unique identifier for the account in your database, used for cross-referencing
	// between Straddle and your systems.
	ExternalID string `json:"external_id,nullable"`
	// Up to 20 additional user-defined key-value pairs. Useful for storing additional
	// information about the account in a structured format.
	Metadata       map[string]string     `json:"metadata,nullable"`
	Settings       AccountV1DataSettings `json:"settings"`
	TermsOfService TermsOfServiceV1      `json:"terms_of_service"`
	// Timestamp of the most recent update to the account.
	UpdatedAt time.Time         `json:"updated_at,nullable" format:"date-time"`
	JSON      accountV1DataJSON `json:"-"`
}

// accountV1DataJSON contains the JSON metadata for the struct [AccountV1Data]
type accountV1DataJSON struct {
	ID              apijson.Field
	AccessLevel     apijson.Field
	OrganizationID  apijson.Field
	Status          apijson.Field
	StatusDetail    apijson.Field
	Type            apijson.Field
	BusinessProfile apijson.Field
	Capabilities    apijson.Field
	CreatedAt       apijson.Field
	ExternalID      apijson.Field
	Metadata        apijson.Field
	Settings        apijson.Field
	TermsOfService  apijson.Field
	UpdatedAt       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccountV1Data) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountV1DataJSON) RawJSON() string {
	return r.raw
}

// The access level granted to the account. This is determined by your platform
// configuration. Use `standard` unless instructed otherwise by Straddle.
type AccountV1DataAccessLevel string

const (
	AccountV1DataAccessLevelStandard AccountV1DataAccessLevel = "standard"
	AccountV1DataAccessLevelManaged  AccountV1DataAccessLevel = "managed"
)

func (r AccountV1DataAccessLevel) IsKnown() bool {
	switch r {
	case AccountV1DataAccessLevelStandard, AccountV1DataAccessLevelManaged:
		return true
	}
	return false
}

// The current status of the account (e.g., 'active', 'inactive', 'pending').
type AccountV1DataStatus string

const (
	AccountV1DataStatusCreated    AccountV1DataStatus = "created"
	AccountV1DataStatusOnboarding AccountV1DataStatus = "onboarding"
	AccountV1DataStatusActive     AccountV1DataStatus = "active"
	AccountV1DataStatusRejected   AccountV1DataStatus = "rejected"
	AccountV1DataStatusInactive   AccountV1DataStatus = "inactive"
)

func (r AccountV1DataStatus) IsKnown() bool {
	switch r {
	case AccountV1DataStatusCreated, AccountV1DataStatusOnboarding, AccountV1DataStatusActive, AccountV1DataStatusRejected, AccountV1DataStatusInactive:
		return true
	}
	return false
}

type AccountV1DataStatusDetail struct {
	// A machine-readable code for the specific status, useful for programmatic
	// handling.
	Code string `json:"code,required"`
	// A human-readable message describing the current status.
	Message string `json:"message,required"`
	// A machine-readable identifier for the specific status, useful for programmatic
	// handling.
	Reason AccountV1DataStatusDetailReason `json:"reason,required"`
	// Identifies the origin of the status change (e.g., `bank_decline`, `watchtower`).
	// This helps in tracking the cause of status updates.
	Source AccountV1DataStatusDetailSource `json:"source,required"`
	JSON   accountV1DataStatusDetailJSON   `json:"-"`
}

// accountV1DataStatusDetailJSON contains the JSON metadata for the struct
// [AccountV1DataStatusDetail]
type accountV1DataStatusDetailJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	Reason      apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountV1DataStatusDetail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountV1DataStatusDetailJSON) RawJSON() string {
	return r.raw
}

// A machine-readable identifier for the specific status, useful for programmatic
// handling.
type AccountV1DataStatusDetailReason string

const (
	AccountV1DataStatusDetailReasonUnverified         AccountV1DataStatusDetailReason = "unverified"
	AccountV1DataStatusDetailReasonInReview           AccountV1DataStatusDetailReason = "in_review"
	AccountV1DataStatusDetailReasonPending            AccountV1DataStatusDetailReason = "pending"
	AccountV1DataStatusDetailReasonStuck              AccountV1DataStatusDetailReason = "stuck"
	AccountV1DataStatusDetailReasonVerified           AccountV1DataStatusDetailReason = "verified"
	AccountV1DataStatusDetailReasonFailedVerification AccountV1DataStatusDetailReason = "failed_verification"
	AccountV1DataStatusDetailReasonDisabled           AccountV1DataStatusDetailReason = "disabled"
	AccountV1DataStatusDetailReasonTerminated         AccountV1DataStatusDetailReason = "terminated"
	AccountV1DataStatusDetailReasonNew                AccountV1DataStatusDetailReason = "new"
)

func (r AccountV1DataStatusDetailReason) IsKnown() bool {
	switch r {
	case AccountV1DataStatusDetailReasonUnverified, AccountV1DataStatusDetailReasonInReview, AccountV1DataStatusDetailReasonPending, AccountV1DataStatusDetailReasonStuck, AccountV1DataStatusDetailReasonVerified, AccountV1DataStatusDetailReasonFailedVerification, AccountV1DataStatusDetailReasonDisabled, AccountV1DataStatusDetailReasonTerminated, AccountV1DataStatusDetailReasonNew:
		return true
	}
	return false
}

// Identifies the origin of the status change (e.g., `bank_decline`, `watchtower`).
// This helps in tracking the cause of status updates.
type AccountV1DataStatusDetailSource string

const (
	AccountV1DataStatusDetailSourceWatchtower AccountV1DataStatusDetailSource = "watchtower"
)

func (r AccountV1DataStatusDetailSource) IsKnown() bool {
	switch r {
	case AccountV1DataStatusDetailSourceWatchtower:
		return true
	}
	return false
}

// The type of account (e.g., 'individual', 'business').
type AccountV1DataType string

const (
	AccountV1DataTypeBusiness AccountV1DataType = "business"
)

func (r AccountV1DataType) IsKnown() bool {
	switch r {
	case AccountV1DataTypeBusiness:
		return true
	}
	return false
}

type AccountV1DataCapabilities struct {
	ConsentTypes  AccountV1DataCapabilitiesConsentTypes  `json:"consent_types,required"`
	CustomerTypes AccountV1DataCapabilitiesCustomerTypes `json:"customer_types,required"`
	PaymentTypes  AccountV1DataCapabilitiesPaymentTypes  `json:"payment_types,required"`
	JSON          accountV1DataCapabilitiesJSON          `json:"-"`
}

// accountV1DataCapabilitiesJSON contains the JSON metadata for the struct
// [AccountV1DataCapabilities]
type accountV1DataCapabilitiesJSON struct {
	ConsentTypes  apijson.Field
	CustomerTypes apijson.Field
	PaymentTypes  apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccountV1DataCapabilities) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountV1DataCapabilitiesJSON) RawJSON() string {
	return r.raw
}

type AccountV1DataCapabilitiesConsentTypes struct {
	// Whether the internet payment authorization capability is enabled for the
	// account.
	Internet CapabilityV1 `json:"internet,required"`
	// Whether the signed agreement payment authorization capability is enabled for the
	// account.
	SignedAgreement CapabilityV1                              `json:"signed_agreement,required"`
	JSON            accountV1DataCapabilitiesConsentTypesJSON `json:"-"`
}

// accountV1DataCapabilitiesConsentTypesJSON contains the JSON metadata for the
// struct [AccountV1DataCapabilitiesConsentTypes]
type accountV1DataCapabilitiesConsentTypesJSON struct {
	Internet        apijson.Field
	SignedAgreement apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccountV1DataCapabilitiesConsentTypes) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountV1DataCapabilitiesConsentTypesJSON) RawJSON() string {
	return r.raw
}

type AccountV1DataCapabilitiesCustomerTypes struct {
	Businesses  CapabilityV1                               `json:"businesses,required"`
	Individuals CapabilityV1                               `json:"individuals,required"`
	JSON        accountV1DataCapabilitiesCustomerTypesJSON `json:"-"`
}

// accountV1DataCapabilitiesCustomerTypesJSON contains the JSON metadata for the
// struct [AccountV1DataCapabilitiesCustomerTypes]
type accountV1DataCapabilitiesCustomerTypesJSON struct {
	Businesses  apijson.Field
	Individuals apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountV1DataCapabilitiesCustomerTypes) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountV1DataCapabilitiesCustomerTypesJSON) RawJSON() string {
	return r.raw
}

type AccountV1DataCapabilitiesPaymentTypes struct {
	Charges CapabilityV1                              `json:"charges,required"`
	Payouts CapabilityV1                              `json:"payouts,required"`
	JSON    accountV1DataCapabilitiesPaymentTypesJSON `json:"-"`
}

// accountV1DataCapabilitiesPaymentTypesJSON contains the JSON metadata for the
// struct [AccountV1DataCapabilitiesPaymentTypes]
type accountV1DataCapabilitiesPaymentTypesJSON struct {
	Charges     apijson.Field
	Payouts     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountV1DataCapabilitiesPaymentTypes) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountV1DataCapabilitiesPaymentTypesJSON) RawJSON() string {
	return r.raw
}

type AccountV1DataSettings struct {
	Charges AccountV1DataSettingsCharges `json:"charges,required"`
	Payouts AccountV1DataSettingsPayouts `json:"payouts,required"`
	JSON    accountV1DataSettingsJSON    `json:"-"`
}

// accountV1DataSettingsJSON contains the JSON metadata for the struct
// [AccountV1DataSettings]
type accountV1DataSettingsJSON struct {
	Charges     apijson.Field
	Payouts     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountV1DataSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountV1DataSettingsJSON) RawJSON() string {
	return r.raw
}

type AccountV1DataSettingsCharges struct {
	// The maximum dollar amount of charges in a calendar day.
	DailyAmount int64 `json:"daily_amount,required"`
	// The amount of time it takes for a charge to be funded. This value is defined by
	// Straddle.
	FundingTime AccountV1DataSettingsChargesFundingTime `json:"funding_time,required"`
	// The unique identifier of the linked bank account associated with charges. This
	// value is defined by Straddle.
	LinkedBankAccountID string `json:"linked_bank_account_id,required" format:"uuid"`
	// The maximum amount of a single charge.
	MaxAmount int64 `json:"max_amount,required"`
	// The maximum dollar amount of charges in a calendar month.
	MonthlyAmount int64 `json:"monthly_amount,required"`
	// The maximum number of charges in a calendar month.
	MonthlyCount int64                            `json:"monthly_count,required"`
	JSON         accountV1DataSettingsChargesJSON `json:"-"`
}

// accountV1DataSettingsChargesJSON contains the JSON metadata for the struct
// [AccountV1DataSettingsCharges]
type accountV1DataSettingsChargesJSON struct {
	DailyAmount         apijson.Field
	FundingTime         apijson.Field
	LinkedBankAccountID apijson.Field
	MaxAmount           apijson.Field
	MonthlyAmount       apijson.Field
	MonthlyCount        apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *AccountV1DataSettingsCharges) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountV1DataSettingsChargesJSON) RawJSON() string {
	return r.raw
}

// The amount of time it takes for a charge to be funded. This value is defined by
// Straddle.
type AccountV1DataSettingsChargesFundingTime string

const (
	AccountV1DataSettingsChargesFundingTimeImmediate AccountV1DataSettingsChargesFundingTime = "immediate"
	AccountV1DataSettingsChargesFundingTimeNextDay   AccountV1DataSettingsChargesFundingTime = "next_day"
	AccountV1DataSettingsChargesFundingTimeOneDay    AccountV1DataSettingsChargesFundingTime = "one_day"
	AccountV1DataSettingsChargesFundingTimeTwoDay    AccountV1DataSettingsChargesFundingTime = "two_day"
	AccountV1DataSettingsChargesFundingTimeThreeDay  AccountV1DataSettingsChargesFundingTime = "three_day"
)

func (r AccountV1DataSettingsChargesFundingTime) IsKnown() bool {
	switch r {
	case AccountV1DataSettingsChargesFundingTimeImmediate, AccountV1DataSettingsChargesFundingTimeNextDay, AccountV1DataSettingsChargesFundingTimeOneDay, AccountV1DataSettingsChargesFundingTimeTwoDay, AccountV1DataSettingsChargesFundingTimeThreeDay:
		return true
	}
	return false
}

type AccountV1DataSettingsPayouts struct {
	// The maximum dollar amount of payouts in a day.
	DailyAmount int64 `json:"daily_amount,required"`
	// The amount of time it takes for a payout to be funded. This value is defined by
	// Straddle.
	FundingTime AccountV1DataSettingsPayoutsFundingTime `json:"funding_time,required"`
	// The unique identifier of the linked bank account to use for payouts.
	LinkedBankAccountID string `json:"linked_bank_account_id,required" format:"uuid"`
	// The maximum amount of a single payout.
	MaxAmount int64 `json:"max_amount,required"`
	// The maximum dollar amount of payouts in a month.
	MonthlyAmount int64 `json:"monthly_amount,required"`
	// The maximum number of payouts in a month.
	MonthlyCount int64                            `json:"monthly_count,required"`
	JSON         accountV1DataSettingsPayoutsJSON `json:"-"`
}

// accountV1DataSettingsPayoutsJSON contains the JSON metadata for the struct
// [AccountV1DataSettingsPayouts]
type accountV1DataSettingsPayoutsJSON struct {
	DailyAmount         apijson.Field
	FundingTime         apijson.Field
	LinkedBankAccountID apijson.Field
	MaxAmount           apijson.Field
	MonthlyAmount       apijson.Field
	MonthlyCount        apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *AccountV1DataSettingsPayouts) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountV1DataSettingsPayoutsJSON) RawJSON() string {
	return r.raw
}

// The amount of time it takes for a payout to be funded. This value is defined by
// Straddle.
type AccountV1DataSettingsPayoutsFundingTime string

const (
	AccountV1DataSettingsPayoutsFundingTimeImmediate AccountV1DataSettingsPayoutsFundingTime = "immediate"
	AccountV1DataSettingsPayoutsFundingTimeNextDay   AccountV1DataSettingsPayoutsFundingTime = "next_day"
	AccountV1DataSettingsPayoutsFundingTimeOneDay    AccountV1DataSettingsPayoutsFundingTime = "one_day"
	AccountV1DataSettingsPayoutsFundingTimeTwoDay    AccountV1DataSettingsPayoutsFundingTime = "two_day"
	AccountV1DataSettingsPayoutsFundingTimeThreeDay  AccountV1DataSettingsPayoutsFundingTime = "three_day"
)

func (r AccountV1DataSettingsPayoutsFundingTime) IsKnown() bool {
	switch r {
	case AccountV1DataSettingsPayoutsFundingTimeImmediate, AccountV1DataSettingsPayoutsFundingTimeNextDay, AccountV1DataSettingsPayoutsFundingTimeOneDay, AccountV1DataSettingsPayoutsFundingTimeTwoDay, AccountV1DataSettingsPayoutsFundingTimeThreeDay:
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
type AccountV1ResponseType string

const (
	AccountV1ResponseTypeObject AccountV1ResponseType = "object"
	AccountV1ResponseTypeArray  AccountV1ResponseType = "array"
	AccountV1ResponseTypeError  AccountV1ResponseType = "error"
	AccountV1ResponseTypeNone   AccountV1ResponseType = "none"
)

func (r AccountV1ResponseType) IsKnown() bool {
	switch r {
	case AccountV1ResponseTypeObject, AccountV1ResponseTypeArray, AccountV1ResponseTypeError, AccountV1ResponseTypeNone:
		return true
	}
	return false
}

// The address object is optional. If provided, it must be a valid address.
type AddressV1 struct {
	// City, district, suburb, town, or village.
	City string `json:"city,nullable"`
	// The country of the address, in ISO 3166-1 alpha-2 format.
	Country string `json:"country,nullable"`
	// Primary address line (e.g., street, PO Box).
	Line1 string `json:"line1,nullable"`
	// Secondary address line (e.g., apartment, suite, unit, or building).
	Line2 string `json:"line2,nullable"`
	// Postal or ZIP code.
	PostalCode string `json:"postal_code,nullable"`
	// Two-letter state code.
	State string        `json:"state,nullable"`
	JSON  addressV1JSON `json:"-"`
}

// addressV1JSON contains the JSON metadata for the struct [AddressV1]
type addressV1JSON struct {
	City        apijson.Field
	Country     apijson.Field
	Line1       apijson.Field
	Line2       apijson.Field
	PostalCode  apijson.Field
	State       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressV1JSON) RawJSON() string {
	return r.raw
}

// The address object is optional. If provided, it must be a valid address.
type AddressV1Param struct {
	// City, district, suburb, town, or village.
	City param.Field[string] `json:"city"`
	// The country of the address, in ISO 3166-1 alpha-2 format.
	Country param.Field[string] `json:"country"`
	// Primary address line (e.g., street, PO Box).
	Line1 param.Field[string] `json:"line1"`
	// Secondary address line (e.g., apartment, suite, unit, or building).
	Line2 param.Field[string] `json:"line2"`
	// Postal or ZIP code.
	PostalCode param.Field[string] `json:"postal_code"`
	// Two-letter state code.
	State param.Field[string] `json:"state"`
}

func (r AddressV1Param) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BusinessProfileV1 struct {
	// The operating or trade name of the business.
	Name string `json:"name,required"`
	// URL of the business's primary marketing website.
	Website string `json:"website,required" format:"uri"`
	// The address object is optional. If provided, it must be a valid address.
	Address AddressV1 `json:"address,nullable"`
	// A brief description of the business and its products or services.
	Description string     `json:"description,nullable"`
	Industry    IndustryV1 `json:"industry"`
	// The official registered name of the business.
	LegalName string `json:"legal_name,nullable"`
	// The primary contact phone number for the business.
	Phone           string            `json:"phone,nullable"`
	SupportChannels SupportChannelsV1 `json:"support_channels"`
	// The business's tax identification number (e.g., EIN in the US).
	TaxID string `json:"tax_id,nullable"`
	// A description of how the business intends to use Straddle's services.
	UseCase string                `json:"use_case,nullable"`
	JSON    businessProfileV1JSON `json:"-"`
}

// businessProfileV1JSON contains the JSON metadata for the struct
// [BusinessProfileV1]
type businessProfileV1JSON struct {
	Name            apijson.Field
	Website         apijson.Field
	Address         apijson.Field
	Description     apijson.Field
	Industry        apijson.Field
	LegalName       apijson.Field
	Phone           apijson.Field
	SupportChannels apijson.Field
	TaxID           apijson.Field
	UseCase         apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *BusinessProfileV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r businessProfileV1JSON) RawJSON() string {
	return r.raw
}

type BusinessProfileV1Param struct {
	// The operating or trade name of the business.
	Name param.Field[string] `json:"name,required"`
	// URL of the business's primary marketing website.
	Website param.Field[string] `json:"website,required" format:"uri"`
	// The address object is optional. If provided, it must be a valid address.
	Address param.Field[AddressV1Param] `json:"address"`
	// A brief description of the business and its products or services.
	Description param.Field[string]          `json:"description"`
	Industry    param.Field[IndustryV1Param] `json:"industry"`
	// The official registered name of the business.
	LegalName param.Field[string] `json:"legal_name"`
	// The primary contact phone number for the business.
	Phone           param.Field[string]                 `json:"phone"`
	SupportChannels param.Field[SupportChannelsV1Param] `json:"support_channels"`
	// The business's tax identification number (e.g., EIN in the US).
	TaxID param.Field[string] `json:"tax_id"`
	// A description of how the business intends to use Straddle's services.
	UseCase param.Field[string] `json:"use_case"`
}

func (r BusinessProfileV1Param) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CapabilityV1 struct {
	CapabilityStatus CapabilityV1CapabilityStatus `json:"capability_status,required"`
	JSON             capabilityV1JSON             `json:"-"`
}

// capabilityV1JSON contains the JSON metadata for the struct [CapabilityV1]
type capabilityV1JSON struct {
	CapabilityStatus apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CapabilityV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r capabilityV1JSON) RawJSON() string {
	return r.raw
}

type CapabilityV1CapabilityStatus string

const (
	CapabilityV1CapabilityStatusActive   CapabilityV1CapabilityStatus = "active"
	CapabilityV1CapabilityStatusInactive CapabilityV1CapabilityStatus = "inactive"
)

func (r CapabilityV1CapabilityStatus) IsKnown() bool {
	switch r {
	case CapabilityV1CapabilityStatusActive, CapabilityV1CapabilityStatusInactive:
		return true
	}
	return false
}

type IndustryV1 struct {
	// The general category of the industry. Required if not providing MCC.
	Category string `json:"category,nullable"`
	// The Merchant Category Code (MCC) that best describes the business. Optional.
	Mcc string `json:"mcc,nullable"`
	// The specific sector within the industry category. Required if not providing MCC.
	Sector string         `json:"sector,nullable"`
	JSON   industryV1JSON `json:"-"`
}

// industryV1JSON contains the JSON metadata for the struct [IndustryV1]
type industryV1JSON struct {
	Category    apijson.Field
	Mcc         apijson.Field
	Sector      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IndustryV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r industryV1JSON) RawJSON() string {
	return r.raw
}

type IndustryV1Param struct {
	// The general category of the industry. Required if not providing MCC.
	Category param.Field[string] `json:"category"`
	// The Merchant Category Code (MCC) that best describes the business. Optional.
	Mcc param.Field[string] `json:"mcc"`
	// The specific sector within the industry category. Required if not providing MCC.
	Sector param.Field[string] `json:"sector"`
}

func (r IndustryV1Param) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SupportChannelsV1 struct {
	// The email address for customer support inquiries.
	Email string `json:"email,nullable" format:"email"`
	// The phone number for customer support.
	Phone string `json:"phone,nullable"`
	// The URL of the business's customer support page or contact form.
	URL  string                `json:"url,nullable" format:"uri"`
	JSON supportChannelsV1JSON `json:"-"`
}

// supportChannelsV1JSON contains the JSON metadata for the struct
// [SupportChannelsV1]
type supportChannelsV1JSON struct {
	Email       apijson.Field
	Phone       apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SupportChannelsV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r supportChannelsV1JSON) RawJSON() string {
	return r.raw
}

type SupportChannelsV1Param struct {
	// The email address for customer support inquiries.
	Email param.Field[string] `json:"email" format:"email"`
	// The phone number for customer support.
	Phone param.Field[string] `json:"phone"`
	// The URL of the business's customer support page or contact form.
	URL param.Field[string] `json:"url" format:"uri"`
}

func (r SupportChannelsV1Param) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TermsOfServiceV1 struct {
	// The datetime of when the terms of service were accepted, in ISO 8601 format.
	AcceptedDate time.Time `json:"accepted_date,required" format:"date-time"`
	// The type or version of the agreement accepted. Use `embedded` unless your
	// platform was specifically enabled for `direct` agreements.
	AgreementType TermsOfServiceV1AgreementType `json:"agreement_type,required"`
	// The IP address from which the terms of service were accepted.
	AcceptedIP string `json:"accepted_ip,nullable"`
	// The user agent string of the browser or application used to accept the terms.
	AcceptedUserAgent string `json:"accepted_user_agent,nullable"`
	// The URL where the full text of the accepted agreement can be found.
	AgreementURL string               `json:"agreement_url,nullable"`
	JSON         termsOfServiceV1JSON `json:"-"`
}

// termsOfServiceV1JSON contains the JSON metadata for the struct
// [TermsOfServiceV1]
type termsOfServiceV1JSON struct {
	AcceptedDate      apijson.Field
	AgreementType     apijson.Field
	AcceptedIP        apijson.Field
	AcceptedUserAgent apijson.Field
	AgreementURL      apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *TermsOfServiceV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r termsOfServiceV1JSON) RawJSON() string {
	return r.raw
}

// The type or version of the agreement accepted. Use `embedded` unless your
// platform was specifically enabled for `direct` agreements.
type TermsOfServiceV1AgreementType string

const (
	TermsOfServiceV1AgreementTypeEmbedded TermsOfServiceV1AgreementType = "embedded"
	TermsOfServiceV1AgreementTypeDirect   TermsOfServiceV1AgreementType = "direct"
)

func (r TermsOfServiceV1AgreementType) IsKnown() bool {
	switch r {
	case TermsOfServiceV1AgreementTypeEmbedded, TermsOfServiceV1AgreementTypeDirect:
		return true
	}
	return false
}

type TermsOfServiceV1Param struct {
	// The datetime of when the terms of service were accepted, in ISO 8601 format.
	AcceptedDate param.Field[time.Time] `json:"accepted_date,required" format:"date-time"`
	// The type or version of the agreement accepted. Use `embedded` unless your
	// platform was specifically enabled for `direct` agreements.
	AgreementType param.Field[TermsOfServiceV1AgreementType] `json:"agreement_type,required"`
	// The IP address from which the terms of service were accepted.
	AcceptedIP param.Field[string] `json:"accepted_ip"`
	// The user agent string of the browser or application used to accept the terms.
	AcceptedUserAgent param.Field[string] `json:"accepted_user_agent"`
	// The URL where the full text of the accepted agreement can be found.
	AgreementURL param.Field[string] `json:"agreement_url"`
}

func (r TermsOfServiceV1Param) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EmbedAccountNewParams struct {
	// The access level granted to the account. This is determined by your platform
	// configuration. Use `standard` unless instructed otherwise by Straddle.
	AccessLevel param.Field[EmbedAccountNewParamsAccessLevel] `json:"access_level,required"`
	// The type of account to be created. Currently, only `business` is supported.
	AccountType     param.Field[EmbedAccountNewParamsAccountType] `json:"account_type,required"`
	BusinessProfile param.Field[BusinessProfileV1Param]           `json:"business_profile,required"`
	// The unique identifier of the organization related to this account.
	OrganizationID param.Field[string] `json:"organization_id,required" format:"uuid"`
	// Unique identifier for the account in your database, used for cross-referencing
	// between Straddle and your systems.
	ExternalID param.Field[string] `json:"external_id"`
	// Up to 20 additional user-defined key-value pairs. Useful for storing additional
	// information about the account in a structured format.
	Metadata      param.Field[map[string]string] `json:"metadata"`
	CorrelationID param.Field[string]            `header:"correlation-id"`
	RequestID     param.Field[string]            `header:"request-id"`
}

func (r EmbedAccountNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The access level granted to the account. This is determined by your platform
// configuration. Use `standard` unless instructed otherwise by Straddle.
type EmbedAccountNewParamsAccessLevel string

const (
	EmbedAccountNewParamsAccessLevelStandard EmbedAccountNewParamsAccessLevel = "standard"
	EmbedAccountNewParamsAccessLevelManaged  EmbedAccountNewParamsAccessLevel = "managed"
)

func (r EmbedAccountNewParamsAccessLevel) IsKnown() bool {
	switch r {
	case EmbedAccountNewParamsAccessLevelStandard, EmbedAccountNewParamsAccessLevelManaged:
		return true
	}
	return false
}

// The type of account to be created. Currently, only `business` is supported.
type EmbedAccountNewParamsAccountType string

const (
	EmbedAccountNewParamsAccountTypeBusiness EmbedAccountNewParamsAccountType = "business"
)

func (r EmbedAccountNewParamsAccountType) IsKnown() bool {
	switch r {
	case EmbedAccountNewParamsAccountTypeBusiness:
		return true
	}
	return false
}

type EmbedAccountUpdateParams struct {
	BusinessProfile param.Field[BusinessProfileV1Param] `json:"business_profile,required"`
	// Unique identifier for the account in your database, used for cross-referencing
	// between Straddle and your systems.
	ExternalID param.Field[string] `json:"external_id"`
	// Up to 20 additional user-defined key-value pairs. Useful for storing additional
	// information about the account in a structured format.
	Metadata      param.Field[map[string]string] `json:"metadata"`
	CorrelationID param.Field[string]            `header:"correlation-id"`
	RequestID     param.Field[string]            `header:"request-id"`
}

func (r EmbedAccountUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EmbedAccountListParams struct {
	// Results page number. Starts at page 1. Default value: 1
	PageNumber param.Field[int64] `query:"page_number"`
	// Page size. Default value: 100. Max value: 1000
	PageSize   param.Field[int64]  `query:"page_size"`
	SearchText param.Field[string] `query:"search_text"`
	// Sort By. Default value: 'id'.
	SortBy param.Field[string] `query:"sort_by"`
	// Sort Order. Default value: 'asc'.
	SortOrder     param.Field[EmbedAccountListParamsSortOrder] `query:"sort_order"`
	CorrelationID param.Field[string]                          `header:"correlation-id"`
	RequestID     param.Field[string]                          `header:"request-id"`
}

// URLQuery serializes [EmbedAccountListParams]'s query parameters as `url.Values`.
func (r EmbedAccountListParams) URLQuery() (v url.Values) {
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

func (r EmbedAccountListParamsSortOrder) IsKnown() bool {
	switch r {
	case EmbedAccountListParamsSortOrderAsc, EmbedAccountListParamsSortOrderDesc:
		return true
	}
	return false
}

type EmbedAccountGetParams struct {
	CorrelationID param.Field[string] `header:"correlation-id"`
	RequestID     param.Field[string] `header:"request-id"`
}

type EmbedAccountOnboardParams struct {
	TermsOfService param.Field[TermsOfServiceV1Param] `json:"terms_of_service,required"`
	CorrelationID  param.Field[string]                `header:"correlation-id"`
	RequestID      param.Field[string]                `header:"request-id"`
}

func (r EmbedAccountOnboardParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EmbedAccountSimulateParams struct {
	FinalStatus   param.Field[EmbedAccountSimulateParamsFinalStatus] `query:"final_status"`
	CorrelationID param.Field[string]                                `header:"correlation-id"`
	RequestID     param.Field[string]                                `header:"request-id"`
}

// URLQuery serializes [EmbedAccountSimulateParams]'s query parameters as
// `url.Values`.
func (r EmbedAccountSimulateParams) URLQuery() (v url.Values) {
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

func (r EmbedAccountSimulateParamsFinalStatus) IsKnown() bool {
	switch r {
	case EmbedAccountSimulateParamsFinalStatusOnboarding, EmbedAccountSimulateParamsFinalStatusActive:
		return true
	}
	return false
}
