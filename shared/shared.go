// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

import (
	"reflect"
	"time"

	"github.com/stainless-sdks/straddle-go"
	"github.com/stainless-sdks/straddle-go/internal/apijson"
	"github.com/stainless-sdks/straddle-go/internal/param"
	"github.com/tidwall/gjson"
)

type AccountTypeV1 string

const (
	AccountTypeV1Checking AccountTypeV1 = "checking"
	AccountTypeV1Savings  AccountTypeV1 = "savings"
)

func (r AccountTypeV1) IsKnown() bool {
	switch r {
	case AccountTypeV1Checking, AccountTypeV1Savings:
		return true
	}
	return false
}

type AccountV1 struct {
	// Unique identifier for the account.
	ID string `json:"id,required" format:"uuid"`
	// The access level granted to the account. This is determined by your platform
	// configuration. Use `standard` unless instructed otherwise by Straddle.
	AccessLevel AccountV1AccessLevel `json:"access_level,required"`
	// The unique identifier of the organization this account belongs to.
	OrganizationID string `json:"organization_id,required" format:"uuid"`
	// The current status of the account (e.g., 'active', 'inactive', 'pending').
	Status       AccountV1Status       `json:"status,required"`
	StatusDetail AccountV1StatusDetail `json:"status_detail,required"`
	// The type of account (e.g., 'individual', 'business').
	Type            AccountV1Type         `json:"type,required"`
	BusinessProfile BusinessProfileV1     `json:"business_profile"`
	Capabilities    AccountV1Capabilities `json:"capabilities"`
	// Timestamp of when the account was created.
	CreatedAt time.Time `json:"created_at,nullable" format:"date-time"`
	// Unique identifier for the account in your database, used for cross-referencing
	// between Straddle and your systems.
	ExternalID string `json:"external_id,nullable"`
	// Up to 20 additional user-defined key-value pairs. Useful for storing additional
	// information about the account in a structured format.
	Metadata       map[string]string `json:"metadata,nullable"`
	Settings       AccountV1Settings `json:"settings"`
	TermsOfService TermsOfServiceV1  `json:"terms_of_service"`
	// Timestamp of the most recent update to the account.
	UpdatedAt time.Time     `json:"updated_at,nullable" format:"date-time"`
	JSON      accountV1JSON `json:"-"`
}

// accountV1JSON contains the JSON metadata for the struct [AccountV1]
type accountV1JSON struct {
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

func (r *AccountV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountV1JSON) RawJSON() string {
	return r.raw
}

// The access level granted to the account. This is determined by your platform
// configuration. Use `standard` unless instructed otherwise by Straddle.
type AccountV1AccessLevel string

const (
	AccountV1AccessLevelStandard AccountV1AccessLevel = "standard"
	AccountV1AccessLevelManaged  AccountV1AccessLevel = "managed"
)

func (r AccountV1AccessLevel) IsKnown() bool {
	switch r {
	case AccountV1AccessLevelStandard, AccountV1AccessLevelManaged:
		return true
	}
	return false
}

// The current status of the account (e.g., 'active', 'inactive', 'pending').
type AccountV1Status string

const (
	AccountV1StatusCreated    AccountV1Status = "created"
	AccountV1StatusOnboarding AccountV1Status = "onboarding"
	AccountV1StatusActive     AccountV1Status = "active"
	AccountV1StatusRejected   AccountV1Status = "rejected"
	AccountV1StatusInactive   AccountV1Status = "inactive"
)

func (r AccountV1Status) IsKnown() bool {
	switch r {
	case AccountV1StatusCreated, AccountV1StatusOnboarding, AccountV1StatusActive, AccountV1StatusRejected, AccountV1StatusInactive:
		return true
	}
	return false
}

type AccountV1StatusDetail struct {
	// A machine-readable code for the specific status, useful for programmatic
	// handling.
	Code string `json:"code,required"`
	// A human-readable message describing the current status.
	Message string `json:"message,required"`
	// A machine-readable identifier for the specific status, useful for programmatic
	// handling.
	Reason AccountV1StatusDetailReason `json:"reason,required"`
	// Identifies the origin of the status change (e.g., `bank_decline`, `watchtower`).
	// This helps in tracking the cause of status updates.
	Source AccountV1StatusDetailSource `json:"source,required"`
	JSON   accountV1StatusDetailJSON   `json:"-"`
}

// accountV1StatusDetailJSON contains the JSON metadata for the struct
// [AccountV1StatusDetail]
type accountV1StatusDetailJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	Reason      apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountV1StatusDetail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountV1StatusDetailJSON) RawJSON() string {
	return r.raw
}

// A machine-readable identifier for the specific status, useful for programmatic
// handling.
type AccountV1StatusDetailReason string

const (
	AccountV1StatusDetailReasonUnverified         AccountV1StatusDetailReason = "unverified"
	AccountV1StatusDetailReasonInReview           AccountV1StatusDetailReason = "in_review"
	AccountV1StatusDetailReasonPending            AccountV1StatusDetailReason = "pending"
	AccountV1StatusDetailReasonStuck              AccountV1StatusDetailReason = "stuck"
	AccountV1StatusDetailReasonVerified           AccountV1StatusDetailReason = "verified"
	AccountV1StatusDetailReasonFailedVerification AccountV1StatusDetailReason = "failed_verification"
	AccountV1StatusDetailReasonDisabled           AccountV1StatusDetailReason = "disabled"
	AccountV1StatusDetailReasonTerminated         AccountV1StatusDetailReason = "terminated"
	AccountV1StatusDetailReasonNew                AccountV1StatusDetailReason = "new"
)

func (r AccountV1StatusDetailReason) IsKnown() bool {
	switch r {
	case AccountV1StatusDetailReasonUnverified, AccountV1StatusDetailReasonInReview, AccountV1StatusDetailReasonPending, AccountV1StatusDetailReasonStuck, AccountV1StatusDetailReasonVerified, AccountV1StatusDetailReasonFailedVerification, AccountV1StatusDetailReasonDisabled, AccountV1StatusDetailReasonTerminated, AccountV1StatusDetailReasonNew:
		return true
	}
	return false
}

// Identifies the origin of the status change (e.g., `bank_decline`, `watchtower`).
// This helps in tracking the cause of status updates.
type AccountV1StatusDetailSource string

const (
	AccountV1StatusDetailSourceWatchtower AccountV1StatusDetailSource = "watchtower"
)

func (r AccountV1StatusDetailSource) IsKnown() bool {
	switch r {
	case AccountV1StatusDetailSourceWatchtower:
		return true
	}
	return false
}

// The type of account (e.g., 'individual', 'business').
type AccountV1Type string

const (
	AccountV1TypeBusiness AccountV1Type = "business"
)

func (r AccountV1Type) IsKnown() bool {
	switch r {
	case AccountV1TypeBusiness:
		return true
	}
	return false
}

type AccountV1Capabilities struct {
	ConsentTypes  AccountV1CapabilitiesConsentTypes  `json:"consent_types,required"`
	CustomerTypes AccountV1CapabilitiesCustomerTypes `json:"customer_types,required"`
	PaymentTypes  AccountV1CapabilitiesPaymentTypes  `json:"payment_types,required"`
	JSON          accountV1CapabilitiesJSON          `json:"-"`
}

// accountV1CapabilitiesJSON contains the JSON metadata for the struct
// [AccountV1Capabilities]
type accountV1CapabilitiesJSON struct {
	ConsentTypes  apijson.Field
	CustomerTypes apijson.Field
	PaymentTypes  apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccountV1Capabilities) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountV1CapabilitiesJSON) RawJSON() string {
	return r.raw
}

type AccountV1CapabilitiesConsentTypes struct {
	// Whether the internet payment authorization capability is enabled for the
	// account.
	Internet Capability `json:"internet,required"`
	// Whether the signed agreement payment authorization capability is enabled for the
	// account.
	SignedAgreement Capability                            `json:"signed_agreement,required"`
	JSON            accountV1CapabilitiesConsentTypesJSON `json:"-"`
}

// accountV1CapabilitiesConsentTypesJSON contains the JSON metadata for the struct
// [AccountV1CapabilitiesConsentTypes]
type accountV1CapabilitiesConsentTypesJSON struct {
	Internet        apijson.Field
	SignedAgreement apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccountV1CapabilitiesConsentTypes) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountV1CapabilitiesConsentTypesJSON) RawJSON() string {
	return r.raw
}

type AccountV1CapabilitiesCustomerTypes struct {
	Businesses  Capability                             `json:"businesses,required"`
	Individuals Capability                             `json:"individuals,required"`
	JSON        accountV1CapabilitiesCustomerTypesJSON `json:"-"`
}

// accountV1CapabilitiesCustomerTypesJSON contains the JSON metadata for the struct
// [AccountV1CapabilitiesCustomerTypes]
type accountV1CapabilitiesCustomerTypesJSON struct {
	Businesses  apijson.Field
	Individuals apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountV1CapabilitiesCustomerTypes) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountV1CapabilitiesCustomerTypesJSON) RawJSON() string {
	return r.raw
}

type AccountV1CapabilitiesPaymentTypes struct {
	Charges Capability                            `json:"charges,required"`
	Payouts Capability                            `json:"payouts,required"`
	JSON    accountV1CapabilitiesPaymentTypesJSON `json:"-"`
}

// accountV1CapabilitiesPaymentTypesJSON contains the JSON metadata for the struct
// [AccountV1CapabilitiesPaymentTypes]
type accountV1CapabilitiesPaymentTypesJSON struct {
	Charges     apijson.Field
	Payouts     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountV1CapabilitiesPaymentTypes) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountV1CapabilitiesPaymentTypesJSON) RawJSON() string {
	return r.raw
}

type AccountV1Settings struct {
	Charges AccountV1SettingsCharges `json:"charges,required"`
	Payouts AccountV1SettingsPayouts `json:"payouts,required"`
	JSON    accountV1SettingsJSON    `json:"-"`
}

// accountV1SettingsJSON contains the JSON metadata for the struct
// [AccountV1Settings]
type accountV1SettingsJSON struct {
	Charges     apijson.Field
	Payouts     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountV1Settings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountV1SettingsJSON) RawJSON() string {
	return r.raw
}

type AccountV1SettingsCharges struct {
	// The maximum dollar amount of charges in a calendar day.
	DailyAmount int64 `json:"daily_amount,required"`
	// The amount of time it takes for a charge to be funded. This value is defined by
	// Straddle.
	FundingTime AccountV1SettingsChargesFundingTime `json:"funding_time,required"`
	// The unique identifier of the linked bank account associated with charges. This
	// value is defined by Straddle.
	LinkedBankAccountID string `json:"linked_bank_account_id,required" format:"uuid"`
	// The maximum amount of a single charge.
	MaxAmount int64 `json:"max_amount,required"`
	// The maximum dollar amount of charges in a calendar month.
	MonthlyAmount int64 `json:"monthly_amount,required"`
	// The maximum number of charges in a calendar month.
	MonthlyCount int64                        `json:"monthly_count,required"`
	JSON         accountV1SettingsChargesJSON `json:"-"`
}

// accountV1SettingsChargesJSON contains the JSON metadata for the struct
// [AccountV1SettingsCharges]
type accountV1SettingsChargesJSON struct {
	DailyAmount         apijson.Field
	FundingTime         apijson.Field
	LinkedBankAccountID apijson.Field
	MaxAmount           apijson.Field
	MonthlyAmount       apijson.Field
	MonthlyCount        apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *AccountV1SettingsCharges) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountV1SettingsChargesJSON) RawJSON() string {
	return r.raw
}

// The amount of time it takes for a charge to be funded. This value is defined by
// Straddle.
type AccountV1SettingsChargesFundingTime string

const (
	AccountV1SettingsChargesFundingTimeImmediate AccountV1SettingsChargesFundingTime = "immediate"
	AccountV1SettingsChargesFundingTimeNextDay   AccountV1SettingsChargesFundingTime = "next_day"
	AccountV1SettingsChargesFundingTimeOneDay    AccountV1SettingsChargesFundingTime = "one_day"
	AccountV1SettingsChargesFundingTimeTwoDay    AccountV1SettingsChargesFundingTime = "two_day"
	AccountV1SettingsChargesFundingTimeThreeDay  AccountV1SettingsChargesFundingTime = "three_day"
)

func (r AccountV1SettingsChargesFundingTime) IsKnown() bool {
	switch r {
	case AccountV1SettingsChargesFundingTimeImmediate, AccountV1SettingsChargesFundingTimeNextDay, AccountV1SettingsChargesFundingTimeOneDay, AccountV1SettingsChargesFundingTimeTwoDay, AccountV1SettingsChargesFundingTimeThreeDay:
		return true
	}
	return false
}

type AccountV1SettingsPayouts struct {
	// The maximum dollar amount of payouts in a day.
	DailyAmount int64 `json:"daily_amount,required"`
	// The amount of time it takes for a payout to be funded. This value is defined by
	// Straddle.
	FundingTime AccountV1SettingsPayoutsFundingTime `json:"funding_time,required"`
	// The unique identifier of the linked bank account to use for payouts.
	LinkedBankAccountID string `json:"linked_bank_account_id,required" format:"uuid"`
	// The maximum amount of a single payout.
	MaxAmount int64 `json:"max_amount,required"`
	// The maximum dollar amount of payouts in a month.
	MonthlyAmount int64 `json:"monthly_amount,required"`
	// The maximum number of payouts in a month.
	MonthlyCount int64                        `json:"monthly_count,required"`
	JSON         accountV1SettingsPayoutsJSON `json:"-"`
}

// accountV1SettingsPayoutsJSON contains the JSON metadata for the struct
// [AccountV1SettingsPayouts]
type accountV1SettingsPayoutsJSON struct {
	DailyAmount         apijson.Field
	FundingTime         apijson.Field
	LinkedBankAccountID apijson.Field
	MaxAmount           apijson.Field
	MonthlyAmount       apijson.Field
	MonthlyCount        apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *AccountV1SettingsPayouts) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountV1SettingsPayoutsJSON) RawJSON() string {
	return r.raw
}

// The amount of time it takes for a payout to be funded. This value is defined by
// Straddle.
type AccountV1SettingsPayoutsFundingTime string

const (
	AccountV1SettingsPayoutsFundingTimeImmediate AccountV1SettingsPayoutsFundingTime = "immediate"
	AccountV1SettingsPayoutsFundingTimeNextDay   AccountV1SettingsPayoutsFundingTime = "next_day"
	AccountV1SettingsPayoutsFundingTimeOneDay    AccountV1SettingsPayoutsFundingTime = "one_day"
	AccountV1SettingsPayoutsFundingTimeTwoDay    AccountV1SettingsPayoutsFundingTime = "two_day"
	AccountV1SettingsPayoutsFundingTimeThreeDay  AccountV1SettingsPayoutsFundingTime = "three_day"
)

func (r AccountV1SettingsPayoutsFundingTime) IsKnown() bool {
	switch r {
	case AccountV1SettingsPayoutsFundingTimeImmediate, AccountV1SettingsPayoutsFundingTimeNextDay, AccountV1SettingsPayoutsFundingTimeOneDay, AccountV1SettingsPayoutsFundingTimeTwoDay, AccountV1SettingsPayoutsFundingTimeThreeDay:
		return true
	}
	return false
}

type AddressV11 struct {
	// Primary address line (e.g., street, PO Box).
	Address1 string `json:"address1,required"`
	// City, district, suburb, town, or village.
	City string `json:"city,required"`
	// Two-letter state code.
	State string `json:"state,required"`
	// Zip or postal code.
	Zip string `json:"zip,required"`
	// Secondary address line (e.g., apartment, suite, unit, or building).
	Address2 string         `json:"address2,nullable"`
	JSON     addressV11JSON `json:"-"`
}

// addressV11JSON contains the JSON metadata for the struct [AddressV11]
type addressV11JSON struct {
	Address1    apijson.Field
	City        apijson.Field
	State       apijson.Field
	Zip         apijson.Field
	Address2    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressV11) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressV11JSON) RawJSON() string {
	return r.raw
}

type AddressV11Param struct {
	// Primary address line (e.g., street, PO Box).
	Address1 param.Field[string] `json:"address1,required"`
	// City, district, suburb, town, or village.
	City param.Field[string] `json:"city,required"`
	// Two-letter state code.
	State param.Field[string] `json:"state,required"`
	// Zip or postal code.
	Zip param.Field[string] `json:"zip,required"`
	// Secondary address line (e.g., apartment, suite, unit, or building).
	Address2 param.Field[string] `json:"address2"`
}

func (r AddressV11Param) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BankAccountV1RequestParam struct {
	// The name of the account holder as it appears on the bank account. Typically,
	// this is the legal name of the business associated with the account.
	AccountHolder param.Field[string] `json:"account_holder,required"`
	// The bank account number.
	AccountNumber param.Field[string] `json:"account_number,required"`
	// The routing number of the bank account.
	RoutingNumber param.Field[string] `json:"routing_number,required"`
}

func (r BankAccountV1RequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BusinessProfileV1 struct {
	// The operating or trade name of the business.
	Name string `json:"name,required"`
	// URL of the business's primary marketing website.
	Website string `json:"website,required" format:"uri"`
	// The address object is optional. If provided, it must be a valid address.
	Address straddle.AddressV1 `json:"address,nullable"`
	// A brief description of the business and its products or services.
	Description string              `json:"description,nullable"`
	Industry    straddle.IndustryV1 `json:"industry"`
	// The official registered name of the business.
	LegalName string `json:"legal_name,nullable"`
	// The primary contact phone number for the business.
	Phone           string                     `json:"phone,nullable"`
	SupportChannels straddle.SupportChannelsV1 `json:"support_channels"`
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
	Address param.Field[straddle.AddressV1Param] `json:"address"`
	// A brief description of the business and its products or services.
	Description param.Field[string]                   `json:"description"`
	Industry    param.Field[straddle.IndustryV1Param] `json:"industry"`
	// The official registered name of the business.
	LegalName param.Field[string] `json:"legal_name"`
	// The primary contact phone number for the business.
	Phone           param.Field[string]                          `json:"phone"`
	SupportChannels param.Field[straddle.SupportChannelsV1Param] `json:"support_channels"`
	// The business's tax identification number (e.g., EIN in the US).
	TaxID param.Field[string] `json:"tax_id"`
	// A description of how the business intends to use Straddle's services.
	UseCase param.Field[string] `json:"use_case"`
}

func (r BusinessProfileV1Param) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type Capability struct {
	CapabilityStatus CapabilityCapabilityStatus `json:"capability_status,required"`
	JSON             capabilityJSON             `json:"-"`
}

// capabilityJSON contains the JSON metadata for the struct [Capability]
type capabilityJSON struct {
	CapabilityStatus apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *Capability) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r capabilityJSON) RawJSON() string {
	return r.raw
}

type CapabilityCapabilityStatus string

const (
	CapabilityCapabilityStatusActive   CapabilityCapabilityStatus = "active"
	CapabilityCapabilityStatusInactive CapabilityCapabilityStatus = "inactive"
)

func (r CapabilityCapabilityStatus) IsKnown() bool {
	switch r {
	case CapabilityCapabilityStatusActive, CapabilityCapabilityStatusInactive:
		return true
	}
	return false
}

type ChargeConfigurationV1 struct {
	// Defines whether to check the customer's balance before processing the charge.
	BalanceCheck ChargeConfigurationV1BalanceCheck `json:"balance_check,required"`
	JSON         chargeConfigurationV1JSON         `json:"-"`
}

// chargeConfigurationV1JSON contains the JSON metadata for the struct
// [ChargeConfigurationV1]
type chargeConfigurationV1JSON struct {
	BalanceCheck apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ChargeConfigurationV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r chargeConfigurationV1JSON) RawJSON() string {
	return r.raw
}

// Defines whether to check the customer's balance before processing the charge.
type ChargeConfigurationV1BalanceCheck string

const (
	ChargeConfigurationV1BalanceCheckRequired ChargeConfigurationV1BalanceCheck = "required"
	ChargeConfigurationV1BalanceCheckEnabled  ChargeConfigurationV1BalanceCheck = "enabled"
	ChargeConfigurationV1BalanceCheckDisabled ChargeConfigurationV1BalanceCheck = "disabled"
)

func (r ChargeConfigurationV1BalanceCheck) IsKnown() bool {
	switch r {
	case ChargeConfigurationV1BalanceCheckRequired, ChargeConfigurationV1BalanceCheckEnabled, ChargeConfigurationV1BalanceCheckDisabled:
		return true
	}
	return false
}

type ChargeConfigurationV1Param struct {
	// Defines whether to check the customer's balance before processing the charge.
	BalanceCheck param.Field[ChargeConfigurationV1BalanceCheck] `json:"balance_check,required"`
}

func (r ChargeConfigurationV1Param) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ChargeV1ItemResponse struct {
	Data ChargeV1ItemResponseData `json:"data,required"`
	// Metadata about the API request, including an identifier and timestamp.
	Meta ResponseMetadata `json:"meta,required"`
	// Indicates the structure of the returned content.
	//
	//   - "object" means the `data` field contains a single JSON object.
	//   - "array" means the `data` field contains an array of objects.
	//   - "error" means the `data` field contains an error object with details of the
	//     issue.
	//   - "none" means no data is returned.
	ResponseType ResponseTypeEnum         `json:"response_type,required"`
	JSON         chargeV1ItemResponseJSON `json:"-"`
}

// chargeV1ItemResponseJSON contains the JSON metadata for the struct
// [ChargeV1ItemResponse]
type chargeV1ItemResponseJSON struct {
	Data         apijson.Field
	Meta         apijson.Field
	ResponseType apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ChargeV1ItemResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r chargeV1ItemResponseJSON) RawJSON() string {
	return r.raw
}

type ChargeV1ItemResponseData struct {
	// Unique identifier for the charge.
	ID string `json:"id,required" format:"uuid"`
	// The amount of the charge in cents.
	Amount int64 `json:"amount,required"`
	// Configuration options for the charge.
	Config ChargeConfigurationV1 `json:"config,required"`
	// The channel or mechanism through which the payment was authorized. Use
	// `internet` for payments made online or through a mobile app and `signed` for
	// signed agreements where there is a consent form or contract. Use `signed` for
	// PDF signatures.
	ConsentType ConsentTypeV1 `json:"consent_type,required"`
	// Timestamp of when the charge was created.
	CreatedAt time.Time `json:"created_at,required,nullable" format:"date-time"`
	// The currency of the charge. Only USD is supported.
	Currency string `json:"currency,required"`
	// An arbitrary description for the charge.
	Description string `json:"description,required"`
	// Information about the device used when the customer authorized the payment.
	Device DeviceInfoV1 `json:"device,required"`
	// Unique identifier for the charge in your database. This value must be unique
	// across all charges.
	ExternalID string `json:"external_id,required"`
	// Value of the `paykey` used for the charge.
	Paykey string `json:"paykey,required"`
	// The desired date on which the payment should be occur. For charges, this means
	// the date you want the customer to be debited on.
	PaymentDate time.Time `json:"payment_date,required" format:"date"`
	// The current status of the charge.
	Status PaymentStatusV1 `json:"status,required"`
	// Additional details about the current status of the charge.
	StatusDetails StatusDetailsV1 `json:"status_details,required"`
	// Status history.
	StatusHistory []StatusHistoryV1 `json:"status_history,required"`
	// Timestamp of when the charge was last updated.
	UpdatedAt time.Time `json:"updated_at,required,nullable" format:"date-time"`
	// Information about the customer associated with the charge.
	CustomerDetails CustomerDetailsV1 `json:"customer_details"`
	// Timestamp of when the charge was effective in the customer's bank account,
	// otherwise known as the date on which the customer is debited.
	EffectiveAt time.Time `json:"effective_at,nullable" format:"date-time"`
	// Up to 20 additional user-defined key-value pairs. Useful for storing additional
	// information about the charge in a structured format.
	Metadata map[string]string `json:"metadata,nullable"`
	// Information about the paykey used for the charge.
	PaykeyDetails PaykeyDetailsV1 `json:"paykey_details"`
	// The payment rail that the charge will be processed through.
	PaymentRail PaymentRailV1 `json:"payment_rail"`
	// Timestamp of when the charge was processed by Straddle and originated to the
	// payment rail.
	ProcessedAt time.Time                    `json:"processed_at,nullable" format:"date-time"`
	JSON        chargeV1ItemResponseDataJSON `json:"-"`
}

// chargeV1ItemResponseDataJSON contains the JSON metadata for the struct
// [ChargeV1ItemResponseData]
type chargeV1ItemResponseDataJSON struct {
	ID              apijson.Field
	Amount          apijson.Field
	Config          apijson.Field
	ConsentType     apijson.Field
	CreatedAt       apijson.Field
	Currency        apijson.Field
	Description     apijson.Field
	Device          apijson.Field
	ExternalID      apijson.Field
	Paykey          apijson.Field
	PaymentDate     apijson.Field
	Status          apijson.Field
	StatusDetails   apijson.Field
	StatusHistory   apijson.Field
	UpdatedAt       apijson.Field
	CustomerDetails apijson.Field
	EffectiveAt     apijson.Field
	Metadata        apijson.Field
	PaykeyDetails   apijson.Field
	PaymentRail     apijson.Field
	ProcessedAt     apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ChargeV1ItemResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r chargeV1ItemResponseDataJSON) RawJSON() string {
	return r.raw
}

// Compliance profile for individual customers
type ComplianceProfileUnmaskedV1 struct {
	// This field can have the runtime type of [time.Time], [string].
	Dob interface{} `json:"dob"`
	// Full 9-digit Employer Identification Number for businesses. This data is
	// required to trigger Patriot Act compliant KYB verification. Only valid where
	// customer type is 'business'.
	Ein string `json:"ein,nullable" format:"**-*******"`
	// The official name of the business. This name should be correlated with the ein
	// value. Only valid where customer type is 'business'.
	LegalBusinessName string `json:"legal_business_name,nullable"`
	// Social Security Number in the format XXX-XX-XXXX.
	Ssn string `json:"ssn,nullable"`
	// URL of the company's official website.
	Website string                          `json:"website,nullable"`
	JSON    complianceProfileUnmaskedV1JSON `json:"-"`
	union   ComplianceProfileUnmaskedV1Union
}

// complianceProfileUnmaskedV1JSON contains the JSON metadata for the struct
// [ComplianceProfileUnmaskedV1]
type complianceProfileUnmaskedV1JSON struct {
	Dob               apijson.Field
	Ein               apijson.Field
	LegalBusinessName apijson.Field
	Ssn               apijson.Field
	Website           apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r complianceProfileUnmaskedV1JSON) RawJSON() string {
	return r.raw
}

func (r *ComplianceProfileUnmaskedV1) UnmarshalJSON(data []byte) (err error) {
	*r = ComplianceProfileUnmaskedV1{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ComplianceProfileUnmaskedV1Union] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [shared.ComplianceProfileUnmaskedV1IndividualComplianceProfile],
// [shared.ComplianceProfileUnmaskedV1BusinessComplianceProfile].
func (r ComplianceProfileUnmaskedV1) AsUnion() ComplianceProfileUnmaskedV1Union {
	return r.union
}

// Compliance profile for individual customers
//
// Union satisfied by
// [shared.ComplianceProfileUnmaskedV1IndividualComplianceProfile] or
// [shared.ComplianceProfileUnmaskedV1BusinessComplianceProfile].
type ComplianceProfileUnmaskedV1Union interface {
	implementsComplianceProfileUnmaskedV1()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ComplianceProfileUnmaskedV1Union)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ComplianceProfileUnmaskedV1IndividualComplianceProfile{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ComplianceProfileUnmaskedV1BusinessComplianceProfile{}),
		},
	)
}

// Compliance profile for individual customers
type ComplianceProfileUnmaskedV1IndividualComplianceProfile struct {
	// Date of birth in YYYY-MM-DD format.
	Dob time.Time `json:"dob,required" format:"date"`
	// Social Security Number in the format XXX-XX-XXXX.
	Ssn string `json:"ssn,required"`
	// Full 9-digit Employer Identification Number for businesses. This data is
	// required to trigger Patriot Act compliant KYB verification. Only valid where
	// customer type is 'business'.
	Ein string `json:"ein,nullable" format:"**-*******"`
	// The official name of the business. This name should be correlated with the ein
	// value. Only valid where customer type is 'business'.
	LegalBusinessName string `json:"legal_business_name,nullable"`
	// URL of the company's official website.
	Website string                                                     `json:"website,nullable"`
	JSON    complianceProfileUnmaskedV1IndividualComplianceProfileJSON `json:"-"`
}

// complianceProfileUnmaskedV1IndividualComplianceProfileJSON contains the JSON
// metadata for the struct [ComplianceProfileUnmaskedV1IndividualComplianceProfile]
type complianceProfileUnmaskedV1IndividualComplianceProfileJSON struct {
	Dob               apijson.Field
	Ssn               apijson.Field
	Ein               apijson.Field
	LegalBusinessName apijson.Field
	Website           apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ComplianceProfileUnmaskedV1IndividualComplianceProfile) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r complianceProfileUnmaskedV1IndividualComplianceProfileJSON) RawJSON() string {
	return r.raw
}

func (r ComplianceProfileUnmaskedV1IndividualComplianceProfile) implementsComplianceProfileUnmaskedV1() {
}

// Compliance profile for business customers
type ComplianceProfileUnmaskedV1BusinessComplianceProfile struct {
	// Employer Identification Number in the format XX-XXXXXXX.
	Ein string `json:"ein,required"`
	// The official registered name of the business. This name should be correlated
	// with the `ein` value.
	LegalBusinessName string `json:"legal_business_name,required"`
	// Date of birth for individual customers in ISO 8601 format (YYYY-MM-DD). This
	// data is required to trigger Patriot Act compliant KYC verification. Required if
	// SSN is provided. Only valid where customer type is 'individual'.
	Dob string `json:"dob,nullable" format:"****-**-**"`
	// Full 9-digit Social Security Number or government identifier for individuals.
	// This data is required to trigger Patriot Act compliant KYC verification.
	// Required if DOB is provided. Only valid where customer type is 'individual'.
	Ssn string `json:"ssn,nullable" format:"***-**-****"`
	// Business website URL.
	Website string                                                   `json:"website" format:"uri"`
	JSON    complianceProfileUnmaskedV1BusinessComplianceProfileJSON `json:"-"`
}

// complianceProfileUnmaskedV1BusinessComplianceProfileJSON contains the JSON
// metadata for the struct [ComplianceProfileUnmaskedV1BusinessComplianceProfile]
type complianceProfileUnmaskedV1BusinessComplianceProfileJSON struct {
	Ein               apijson.Field
	LegalBusinessName apijson.Field
	Dob               apijson.Field
	Ssn               apijson.Field
	Website           apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ComplianceProfileUnmaskedV1BusinessComplianceProfile) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r complianceProfileUnmaskedV1BusinessComplianceProfileJSON) RawJSON() string {
	return r.raw
}

func (r ComplianceProfileUnmaskedV1BusinessComplianceProfile) implementsComplianceProfileUnmaskedV1() {
}

// Compliance profile for individual customers
type ComplianceProfileUnmaskedV1Param struct {
	Dob param.Field[interface{}] `json:"dob"`
	// Full 9-digit Employer Identification Number for businesses. This data is
	// required to trigger Patriot Act compliant KYB verification. Only valid where
	// customer type is 'business'.
	Ein param.Field[string] `json:"ein" format:"**-*******"`
	// The official name of the business. This name should be correlated with the ein
	// value. Only valid where customer type is 'business'.
	LegalBusinessName param.Field[string] `json:"legal_business_name"`
	// Social Security Number in the format XXX-XX-XXXX.
	Ssn param.Field[string] `json:"ssn"`
	// URL of the company's official website.
	Website param.Field[string] `json:"website"`
}

func (r ComplianceProfileUnmaskedV1Param) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ComplianceProfileUnmaskedV1Param) implementsComplianceProfileUnmaskedV1UnionParam() {}

// Compliance profile for individual customers
//
// Satisfied by
// [shared.ComplianceProfileUnmaskedV1IndividualComplianceProfileParam],
// [shared.ComplianceProfileUnmaskedV1BusinessComplianceProfileParam],
// [ComplianceProfileUnmaskedV1Param].
type ComplianceProfileUnmaskedV1UnionParam interface {
	implementsComplianceProfileUnmaskedV1UnionParam()
}

// Compliance profile for individual customers
type ComplianceProfileUnmaskedV1IndividualComplianceProfileParam struct {
	// Date of birth in YYYY-MM-DD format.
	Dob param.Field[time.Time] `json:"dob,required" format:"date"`
	// Social Security Number in the format XXX-XX-XXXX.
	Ssn param.Field[string] `json:"ssn,required"`
	// Full 9-digit Employer Identification Number for businesses. This data is
	// required to trigger Patriot Act compliant KYB verification. Only valid where
	// customer type is 'business'.
	Ein param.Field[string] `json:"ein" format:"**-*******"`
	// The official name of the business. This name should be correlated with the ein
	// value. Only valid where customer type is 'business'.
	LegalBusinessName param.Field[string] `json:"legal_business_name"`
	// URL of the company's official website.
	Website param.Field[string] `json:"website"`
}

func (r ComplianceProfileUnmaskedV1IndividualComplianceProfileParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ComplianceProfileUnmaskedV1IndividualComplianceProfileParam) implementsComplianceProfileUnmaskedV1UnionParam() {
}

// Compliance profile for business customers
type ComplianceProfileUnmaskedV1BusinessComplianceProfileParam struct {
	// Employer Identification Number in the format XX-XXXXXXX.
	Ein param.Field[string] `json:"ein,required"`
	// The official registered name of the business. This name should be correlated
	// with the `ein` value.
	LegalBusinessName param.Field[string] `json:"legal_business_name,required"`
	// Date of birth for individual customers in ISO 8601 format (YYYY-MM-DD). This
	// data is required to trigger Patriot Act compliant KYC verification. Required if
	// SSN is provided. Only valid where customer type is 'individual'.
	Dob param.Field[string] `json:"dob" format:"****-**-**"`
	// Full 9-digit Social Security Number or government identifier for individuals.
	// This data is required to trigger Patriot Act compliant KYC verification.
	// Required if DOB is provided. Only valid where customer type is 'individual'.
	Ssn param.Field[string] `json:"ssn" format:"***-**-****"`
	// Business website URL.
	Website param.Field[string] `json:"website" format:"uri"`
}

func (r ComplianceProfileUnmaskedV1BusinessComplianceProfileParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ComplianceProfileUnmaskedV1BusinessComplianceProfileParam) implementsComplianceProfileUnmaskedV1UnionParam() {
}

// The channel or mechanism through which the payment was authorized. Use
// `internet` for payments made online or through a mobile app and `signed` for
// signed agreements where there is a consent form or contract. Use `signed` for
// PDF signatures.
type ConsentTypeV1 string

const (
	ConsentTypeV1Internet ConsentTypeV1 = "internet"
	ConsentTypeV1Signed   ConsentTypeV1 = "signed"
)

func (r ConsentTypeV1) IsKnown() bool {
	switch r {
	case ConsentTypeV1Internet, ConsentTypeV1Signed:
		return true
	}
	return false
}

type CustomerStatusV1 string

const (
	CustomerStatusV1Pending  CustomerStatusV1 = "pending"
	CustomerStatusV1Review   CustomerStatusV1 = "review"
	CustomerStatusV1Verified CustomerStatusV1 = "verified"
	CustomerStatusV1Inactive CustomerStatusV1 = "inactive"
	CustomerStatusV1Rejected CustomerStatusV1 = "rejected"
)

func (r CustomerStatusV1) IsKnown() bool {
	switch r {
	case CustomerStatusV1Pending, CustomerStatusV1Review, CustomerStatusV1Verified, CustomerStatusV1Inactive, CustomerStatusV1Rejected:
		return true
	}
	return false
}

type CustomerTypeV1 string

const (
	CustomerTypeV1Individual CustomerTypeV1 = "individual"
	CustomerTypeV1Business   CustomerTypeV1 = "business"
)

func (r CustomerTypeV1) IsKnown() bool {
	switch r {
	case CustomerTypeV1Individual, CustomerTypeV1Business:
		return true
	}
	return false
}

type CustomerV1 struct {
	// Unique identifier for the customer.
	ID string `json:"id,required" format:"uuid"`
	// Timestamp of when the customer record was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The customer's email address.
	Email string `json:"email,required" format:"email"`
	// Full name of the individual or business name.
	Name string `json:"name,required"`
	// The customer's phone number in E.164 format.
	Phone  string           `json:"phone,required"`
	Status CustomerStatusV1 `json:"status,required"`
	Type   CustomerTypeV1   `json:"type,required"`
	// Timestamp of the most recent update to the customer record.
	UpdatedAt time.Time  `json:"updated_at,required" format:"date-time"`
	Address   AddressV11 `json:"address,nullable"`
	// Compliance profile for individual customers
	ComplianceProfile CustomerV1ComplianceProfile `json:"compliance_profile"`
	Device            CustomerV1Device            `json:"device"`
	// Unique identifier for the customer in your database, used for cross-referencing
	// between Straddle and your systems.
	ExternalID string `json:"external_id,nullable"`
	// Up to 20 additional user-defined key-value pairs. Useful for storing additional
	// information about the customer in a structured format.
	Metadata map[string]string `json:"metadata,nullable"`
	JSON     customerV1JSON    `json:"-"`
}

// customerV1JSON contains the JSON metadata for the struct [CustomerV1]
type customerV1JSON struct {
	ID                apijson.Field
	CreatedAt         apijson.Field
	Email             apijson.Field
	Name              apijson.Field
	Phone             apijson.Field
	Status            apijson.Field
	Type              apijson.Field
	UpdatedAt         apijson.Field
	Address           apijson.Field
	ComplianceProfile apijson.Field
	Device            apijson.Field
	ExternalID        apijson.Field
	Metadata          apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *CustomerV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerV1JSON) RawJSON() string {
	return r.raw
}

// Compliance profile for individual customers
type CustomerV1ComplianceProfile struct {
	// This field can have the runtime type of [time.Time], [string].
	Dob interface{} `json:"dob"`
	// Full 9-digit Employer Identification Number for businesses. This data is
	// required to trigger Patriot Act compliant Know Your Business (KYB) verification.
	// Only valid where customer type is 'business'.
	Ein string `json:"ein,nullable" format:"**-*******"`
	// The official name of the business. This name should be correlated with the ein
	// value. Only valid where customer type is 'business'.
	LegalBusinessName string `json:"legal_business_name,nullable"`
	// Social Security Number in the format XXX-XX-XXXX.
	Ssn string `json:"ssn,nullable"`
	// URL of the company's official website. Only valid where customer type is
	// 'business'.
	Website string                          `json:"website,nullable"`
	JSON    customerV1ComplianceProfileJSON `json:"-"`
	union   CustomerV1ComplianceProfileUnion
}

// customerV1ComplianceProfileJSON contains the JSON metadata for the struct
// [CustomerV1ComplianceProfile]
type customerV1ComplianceProfileJSON struct {
	Dob               apijson.Field
	Ein               apijson.Field
	LegalBusinessName apijson.Field
	Ssn               apijson.Field
	Website           apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r customerV1ComplianceProfileJSON) RawJSON() string {
	return r.raw
}

func (r *CustomerV1ComplianceProfile) UnmarshalJSON(data []byte) (err error) {
	*r = CustomerV1ComplianceProfile{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [CustomerV1ComplianceProfileUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [shared.CustomerV1ComplianceProfileIndividualComplianceProfile],
// [shared.CustomerV1ComplianceProfileBusinessComplianceProfile].
func (r CustomerV1ComplianceProfile) AsUnion() CustomerV1ComplianceProfileUnion {
	return r.union
}

// Compliance profile for individual customers
//
// Union satisfied by
// [shared.CustomerV1ComplianceProfileIndividualComplianceProfile] or
// [shared.CustomerV1ComplianceProfileBusinessComplianceProfile].
type CustomerV1ComplianceProfileUnion interface {
	implementsCustomerV1ComplianceProfile()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CustomerV1ComplianceProfileUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CustomerV1ComplianceProfileIndividualComplianceProfile{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CustomerV1ComplianceProfileBusinessComplianceProfile{}),
		},
	)
}

// Compliance profile for individual customers
type CustomerV1ComplianceProfileIndividualComplianceProfile struct {
	// Date of birth in YYYY-MM-DD format.
	Dob time.Time `json:"dob,required" format:"date"`
	// Social Security Number in the format XXX-XX-XXXX.
	Ssn string `json:"ssn,required"`
	// Full 9-digit Employer Identification Number for businesses. This data is
	// required to trigger Patriot Act compliant Know Your Business (KYB) verification.
	// Only valid where customer type is 'business'.
	Ein string `json:"ein,nullable" format:"**-*******"`
	// The official name of the business. This name should be correlated with the ein
	// value. Only valid where customer type is 'business'.
	LegalBusinessName string `json:"legal_business_name,nullable"`
	// URL of the company's official website. Only valid where customer type is
	// 'business'.
	Website string                                                     `json:"website,nullable"`
	JSON    customerV1ComplianceProfileIndividualComplianceProfileJSON `json:"-"`
}

// customerV1ComplianceProfileIndividualComplianceProfileJSON contains the JSON
// metadata for the struct [CustomerV1ComplianceProfileIndividualComplianceProfile]
type customerV1ComplianceProfileIndividualComplianceProfileJSON struct {
	Dob               apijson.Field
	Ssn               apijson.Field
	Ein               apijson.Field
	LegalBusinessName apijson.Field
	Website           apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *CustomerV1ComplianceProfileIndividualComplianceProfile) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerV1ComplianceProfileIndividualComplianceProfileJSON) RawJSON() string {
	return r.raw
}

func (r CustomerV1ComplianceProfileIndividualComplianceProfile) implementsCustomerV1ComplianceProfile() {
}

// Compliance profile for business customers
type CustomerV1ComplianceProfileBusinessComplianceProfile struct {
	// Employer Identification Number in the format XX-XXXXXXX.
	Ein string `json:"ein,required"`
	// The official registered name of the business. This name should be correlated
	// with the `ein` value.
	LegalBusinessName string `json:"legal_business_name,required"`
	// Date of birth for individual customers in ISO 8601 format (YYYY-MM-DD). This
	// data is required to trigger Patriot Act compliant Know Your Customer (KYC)
	// verification. Required if SSN is provided. Only valid where customer type is
	// 'individual'.
	Dob string `json:"dob,nullable" format:"****-**-**"`
	// Full 9-digit Social Security Number or government identifier for individuals.
	// This data is required to trigger Patriot Act compliant KYC verification.
	// Required if DOB is provided. Only valid where customer type is 'individual'.
	Ssn string `json:"ssn,nullable" format:"***-**-****"`
	// Business website URL.
	Website string                                                   `json:"website" format:"uri"`
	JSON    customerV1ComplianceProfileBusinessComplianceProfileJSON `json:"-"`
}

// customerV1ComplianceProfileBusinessComplianceProfileJSON contains the JSON
// metadata for the struct [CustomerV1ComplianceProfileBusinessComplianceProfile]
type customerV1ComplianceProfileBusinessComplianceProfileJSON struct {
	Ein               apijson.Field
	LegalBusinessName apijson.Field
	Dob               apijson.Field
	Ssn               apijson.Field
	Website           apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *CustomerV1ComplianceProfileBusinessComplianceProfile) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerV1ComplianceProfileBusinessComplianceProfileJSON) RawJSON() string {
	return r.raw
}

func (r CustomerV1ComplianceProfileBusinessComplianceProfile) implementsCustomerV1ComplianceProfile() {
}

type CustomerV1Device struct {
	// The customer's IP address at the time of profile creation. Use `0.0.0.0` to
	// represent an offline customer registration.
	IPAddress string               `json:"ip_address,required" format:"ipv4"`
	JSON      customerV1DeviceJSON `json:"-"`
}

// customerV1DeviceJSON contains the JSON metadata for the struct
// [CustomerV1Device]
type customerV1DeviceJSON struct {
	IPAddress   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerV1Device) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerV1DeviceJSON) RawJSON() string {
	return r.raw
}

type CustomerV1ItemResponse struct {
	Data CustomerV1 `json:"data,required"`
	// Metadata about the API request, including an identifier and timestamp.
	Meta ResponseMetadata `json:"meta,required"`
	// Indicates the structure of the returned content.
	//
	//   - "object" means the `data` field contains a single JSON object.
	//   - "array" means the `data` field contains an array of objects.
	//   - "error" means the `data` field contains an error object with details of the
	//     issue.
	//   - "none" means no data is returned.
	ResponseType ResponseTypeEnum           `json:"response_type,required"`
	JSON         customerV1ItemResponseJSON `json:"-"`
}

// customerV1ItemResponseJSON contains the JSON metadata for the struct
// [CustomerV1ItemResponse]
type customerV1ItemResponseJSON struct {
	Data         apijson.Field
	Meta         apijson.Field
	ResponseType apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *CustomerV1ItemResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerV1ItemResponseJSON) RawJSON() string {
	return r.raw
}

// Information about the customer associated with the charge or payout.
type CustomerDetailsV1 struct {
	// Unique identifier for the customer.
	ID string `json:"id,required" format:"uuid"`
	// The type of customer.
	CustomerType CustomerDetailsV1CustomerType `json:"customer_type,required"`
	// Email.
	Email string `json:"email,required"`
	// The name of the customer.
	Name string `json:"name,required"`
	// Phone.
	Phone string                `json:"phone,required"`
	JSON  customerDetailsV1JSON `json:"-"`
}

// customerDetailsV1JSON contains the JSON metadata for the struct
// [CustomerDetailsV1]
type customerDetailsV1JSON struct {
	ID           apijson.Field
	CustomerType apijson.Field
	Email        apijson.Field
	Name         apijson.Field
	Phone        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *CustomerDetailsV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerDetailsV1JSON) RawJSON() string {
	return r.raw
}

// The type of customer.
type CustomerDetailsV1CustomerType string

const (
	CustomerDetailsV1CustomerTypeIndividual CustomerDetailsV1CustomerType = "individual"
	CustomerDetailsV1CustomerTypeBusiness   CustomerDetailsV1CustomerType = "business"
)

func (r CustomerDetailsV1CustomerType) IsKnown() bool {
	switch r {
	case CustomerDetailsV1CustomerTypeIndividual, CustomerDetailsV1CustomerTypeBusiness:
		return true
	}
	return false
}

type DeviceUnmaskedV1 struct {
	// The customer's IP address at the time of profile creation. Use `0.0.0.0` to
	// represent an offline customer registration.
	IPAddress string               `json:"ip_address,required" format:"ipv4"`
	JSON      deviceUnmaskedV1JSON `json:"-"`
}

// deviceUnmaskedV1JSON contains the JSON metadata for the struct
// [DeviceUnmaskedV1]
type deviceUnmaskedV1JSON struct {
	IPAddress   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceUnmaskedV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceUnmaskedV1JSON) RawJSON() string {
	return r.raw
}

type DeviceUnmaskedV1Param struct {
	// The customer's IP address at the time of profile creation. Use `0.0.0.0` to
	// represent an offline customer registration.
	IPAddress param.Field[string] `json:"ip_address,required" format:"ipv4"`
}

func (r DeviceUnmaskedV1Param) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DeviceInfoV1 struct {
	// The IP address of the device used when the customer authorized the charge or
	// payout. Use `0.0.0.0` to represent an offline consent interaction.
	IPAddress string           `json:"ip_address,required" format:"ipv4"`
	JSON      deviceInfoV1JSON `json:"-"`
}

// deviceInfoV1JSON contains the JSON metadata for the struct [DeviceInfoV1]
type deviceInfoV1JSON struct {
	IPAddress   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceInfoV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceInfoV1JSON) RawJSON() string {
	return r.raw
}

type DeviceInfoV1Param struct {
	// The IP address of the device used when the customer authorized the charge or
	// payout. Use `0.0.0.0` to represent an offline consent interaction.
	IPAddress param.Field[string] `json:"ip_address,required" format:"ipv4"`
}

func (r DeviceInfoV1Param) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FundingEventSummaryV1 struct {
	// Unique identifier for the funding event.
	ID string `json:"id,required" format:"uuid"`
	// The amount of the funding event in cents.
	Amount int64 `json:"amount,required"`
	// Describes the direction of the funding event from the perspective of the
	// `linked_bank_account`.
	Direction TransferDirectionV1 `json:"direction,required"`
	// The funding event types describes the direction and reason for the funding
	// event.
	EventType FundingEventTypeV1 `json:"event_type,required"`
	// The number of payments associated with the funding event.
	PaymentCount int64 `json:"payment_count,required"`
	// Trace number.
	TraceNumbers []string `json:"trace_numbers,required"`
	// The date on which the funding event occurred. For `deposits` and `returns`, this
	// is the date the funds were credited to your bank account. For `withdrawals` and
	// `reversals`, this is the date the funds were debited from your bank account.
	TransferDate time.Time `json:"transfer_date,required" format:"date"`
	// The trace number of the funding event.
	TraceNumber string                    `json:"trace_number,nullable"`
	JSON        fundingEventSummaryV1JSON `json:"-"`
}

// fundingEventSummaryV1JSON contains the JSON metadata for the struct
// [FundingEventSummaryV1]
type fundingEventSummaryV1JSON struct {
	ID           apijson.Field
	Amount       apijson.Field
	Direction    apijson.Field
	EventType    apijson.Field
	PaymentCount apijson.Field
	TraceNumbers apijson.Field
	TransferDate apijson.Field
	TraceNumber  apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *FundingEventSummaryV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fundingEventSummaryV1JSON) RawJSON() string {
	return r.raw
}

// The funding event types describes the direction and reason for the funding
// event.
type FundingEventTypeV1 string

const (
	FundingEventTypeV1ChargeDeposit    FundingEventTypeV1 = "charge_deposit"
	FundingEventTypeV1ChargeReversal   FundingEventTypeV1 = "charge_reversal"
	FundingEventTypeV1PayoutReturn     FundingEventTypeV1 = "payout_return"
	FundingEventTypeV1PayoutWithdrawal FundingEventTypeV1 = "payout_withdrawal"
)

func (r FundingEventTypeV1) IsKnown() bool {
	switch r {
	case FundingEventTypeV1ChargeDeposit, FundingEventTypeV1ChargeReversal, FundingEventTypeV1PayoutReturn, FundingEventTypeV1PayoutWithdrawal:
		return true
	}
	return false
}

type IdentityDecisionV1 string

const (
	IdentityDecisionV1Accept IdentityDecisionV1 = "accept"
	IdentityDecisionV1Reject IdentityDecisionV1 = "reject"
	IdentityDecisionV1Review IdentityDecisionV1 = "review"
)

func (r IdentityDecisionV1) IsKnown() bool {
	switch r {
	case IdentityDecisionV1Accept, IdentityDecisionV1Reject, IdentityDecisionV1Review:
		return true
	}
	return false
}

type IdentityVerificationBreakdownV1 struct {
	// List of specific result codes from the fraud and risk screening.
	Codes []string `json:"codes,nullable"`
	// Represents the strength of the correlation between provided and known
	// information. A higher score indicates a stronger correlation.
	CorrelationScore float64            `json:"correlation_score,nullable"`
	Decision         IdentityDecisionV1 `json:"decision"`
	// Predicts the inherent risk associated with the customer for a given module. A
	// higher score indicates a greater likelihood of fraud.
	RiskScore float64                             `json:"risk_score,nullable"`
	JSON      identityVerificationBreakdownV1JSON `json:"-"`
}

// identityVerificationBreakdownV1JSON contains the JSON metadata for the struct
// [IdentityVerificationBreakdownV1]
type identityVerificationBreakdownV1JSON struct {
	Codes            apijson.Field
	CorrelationScore apijson.Field
	Decision         apijson.Field
	RiskScore        apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *IdentityVerificationBreakdownV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityVerificationBreakdownV1JSON) RawJSON() string {
	return r.raw
}

type ItemResponseOfAccountV1 struct {
	Data AccountV1 `json:"data,required"`
	// Metadata about the API request, including an identifier and timestamp.
	Meta ResponseMetadata `json:"meta,required"`
	// Indicates the structure of the returned content.
	//
	//   - "object" means the `data` field contains a single JSON object.
	//   - "array" means the `data` field contains an array of objects.
	//   - "error" means the `data` field contains an error object with details of the
	//     issue.
	//   - "none" means no data is returned.
	ResponseType ItemResponseOfAccountV1ResponseType `json:"response_type,required"`
	JSON         itemResponseOfAccountV1JSON         `json:"-"`
}

// itemResponseOfAccountV1JSON contains the JSON metadata for the struct
// [ItemResponseOfAccountV1]
type itemResponseOfAccountV1JSON struct {
	Data         apijson.Field
	Meta         apijson.Field
	ResponseType apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ItemResponseOfAccountV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r itemResponseOfAccountV1JSON) RawJSON() string {
	return r.raw
}

// Indicates the structure of the returned content.
//
//   - "object" means the `data` field contains a single JSON object.
//   - "array" means the `data` field contains an array of objects.
//   - "error" means the `data` field contains an error object with details of the
//     issue.
//   - "none" means no data is returned.
type ItemResponseOfAccountV1ResponseType string

const (
	ItemResponseOfAccountV1ResponseTypeObject ItemResponseOfAccountV1ResponseType = "object"
	ItemResponseOfAccountV1ResponseTypeArray  ItemResponseOfAccountV1ResponseType = "array"
	ItemResponseOfAccountV1ResponseTypeError  ItemResponseOfAccountV1ResponseType = "error"
	ItemResponseOfAccountV1ResponseTypeNone   ItemResponseOfAccountV1ResponseType = "none"
)

func (r ItemResponseOfAccountV1ResponseType) IsKnown() bool {
	switch r {
	case ItemResponseOfAccountV1ResponseTypeObject, ItemResponseOfAccountV1ResponseTypeArray, ItemResponseOfAccountV1ResponseTypeError, ItemResponseOfAccountV1ResponseTypeNone:
		return true
	}
	return false
}

type ItemResponseOfLinkedBankAccountV1 struct {
	Data LinkedBankAccountV1 `json:"data,required"`
	// Metadata about the API request, including an identifier and timestamp.
	Meta ResponseMetadata `json:"meta,required"`
	// Indicates the structure of the returned content.
	//
	//   - "object" means the `data` field contains a single JSON object.
	//   - "array" means the `data` field contains an array of objects.
	//   - "error" means the `data` field contains an error object with details of the
	//     issue.
	//   - "none" means no data is returned.
	ResponseType ItemResponseOfLinkedBankAccountV1ResponseType `json:"response_type,required"`
	JSON         itemResponseOfLinkedBankAccountV1JSON         `json:"-"`
}

// itemResponseOfLinkedBankAccountV1JSON contains the JSON metadata for the struct
// [ItemResponseOfLinkedBankAccountV1]
type itemResponseOfLinkedBankAccountV1JSON struct {
	Data         apijson.Field
	Meta         apijson.Field
	ResponseType apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ItemResponseOfLinkedBankAccountV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r itemResponseOfLinkedBankAccountV1JSON) RawJSON() string {
	return r.raw
}

// Indicates the structure of the returned content.
//
//   - "object" means the `data` field contains a single JSON object.
//   - "array" means the `data` field contains an array of objects.
//   - "error" means the `data` field contains an error object with details of the
//     issue.
//   - "none" means no data is returned.
type ItemResponseOfLinkedBankAccountV1ResponseType string

const (
	ItemResponseOfLinkedBankAccountV1ResponseTypeObject ItemResponseOfLinkedBankAccountV1ResponseType = "object"
	ItemResponseOfLinkedBankAccountV1ResponseTypeArray  ItemResponseOfLinkedBankAccountV1ResponseType = "array"
	ItemResponseOfLinkedBankAccountV1ResponseTypeError  ItemResponseOfLinkedBankAccountV1ResponseType = "error"
	ItemResponseOfLinkedBankAccountV1ResponseTypeNone   ItemResponseOfLinkedBankAccountV1ResponseType = "none"
)

func (r ItemResponseOfLinkedBankAccountV1ResponseType) IsKnown() bool {
	switch r {
	case ItemResponseOfLinkedBankAccountV1ResponseTypeObject, ItemResponseOfLinkedBankAccountV1ResponseTypeArray, ItemResponseOfLinkedBankAccountV1ResponseTypeError, ItemResponseOfLinkedBankAccountV1ResponseTypeNone:
		return true
	}
	return false
}

type ItemResponseOfOrganizationV1 struct {
	Data OrganizationV1 `json:"data,required"`
	// Metadata about the API request, including an identifier and timestamp.
	Meta ResponseMetadata `json:"meta,required"`
	// Indicates the structure of the returned content.
	//
	//   - "object" means the `data` field contains a single JSON object.
	//   - "array" means the `data` field contains an array of objects.
	//   - "error" means the `data` field contains an error object with details of the
	//     issue.
	//   - "none" means no data is returned.
	ResponseType ItemResponseOfOrganizationV1ResponseType `json:"response_type,required"`
	JSON         itemResponseOfOrganizationV1JSON         `json:"-"`
}

// itemResponseOfOrganizationV1JSON contains the JSON metadata for the struct
// [ItemResponseOfOrganizationV1]
type itemResponseOfOrganizationV1JSON struct {
	Data         apijson.Field
	Meta         apijson.Field
	ResponseType apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ItemResponseOfOrganizationV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r itemResponseOfOrganizationV1JSON) RawJSON() string {
	return r.raw
}

// Indicates the structure of the returned content.
//
//   - "object" means the `data` field contains a single JSON object.
//   - "array" means the `data` field contains an array of objects.
//   - "error" means the `data` field contains an error object with details of the
//     issue.
//   - "none" means no data is returned.
type ItemResponseOfOrganizationV1ResponseType string

const (
	ItemResponseOfOrganizationV1ResponseTypeObject ItemResponseOfOrganizationV1ResponseType = "object"
	ItemResponseOfOrganizationV1ResponseTypeArray  ItemResponseOfOrganizationV1ResponseType = "array"
	ItemResponseOfOrganizationV1ResponseTypeError  ItemResponseOfOrganizationV1ResponseType = "error"
	ItemResponseOfOrganizationV1ResponseTypeNone   ItemResponseOfOrganizationV1ResponseType = "none"
)

func (r ItemResponseOfOrganizationV1ResponseType) IsKnown() bool {
	switch r {
	case ItemResponseOfOrganizationV1ResponseTypeObject, ItemResponseOfOrganizationV1ResponseTypeArray, ItemResponseOfOrganizationV1ResponseTypeError, ItemResponseOfOrganizationV1ResponseTypeNone:
		return true
	}
	return false
}

type ItemResponseOfRepresentativeV1 struct {
	Data RepresentativeV1 `json:"data,required"`
	// Metadata about the API request, including an identifier and timestamp.
	Meta ResponseMetadata `json:"meta,required"`
	// Indicates the structure of the returned content.
	//
	//   - "object" means the `data` field contains a single JSON object.
	//   - "array" means the `data` field contains an array of objects.
	//   - "error" means the `data` field contains an error object with details of the
	//     issue.
	//   - "none" means no data is returned.
	ResponseType ItemResponseOfRepresentativeV1ResponseType `json:"response_type,required"`
	JSON         itemResponseOfRepresentativeV1JSON         `json:"-"`
}

// itemResponseOfRepresentativeV1JSON contains the JSON metadata for the struct
// [ItemResponseOfRepresentativeV1]
type itemResponseOfRepresentativeV1JSON struct {
	Data         apijson.Field
	Meta         apijson.Field
	ResponseType apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ItemResponseOfRepresentativeV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r itemResponseOfRepresentativeV1JSON) RawJSON() string {
	return r.raw
}

// Indicates the structure of the returned content.
//
//   - "object" means the `data` field contains a single JSON object.
//   - "array" means the `data` field contains an array of objects.
//   - "error" means the `data` field contains an error object with details of the
//     issue.
//   - "none" means no data is returned.
type ItemResponseOfRepresentativeV1ResponseType string

const (
	ItemResponseOfRepresentativeV1ResponseTypeObject ItemResponseOfRepresentativeV1ResponseType = "object"
	ItemResponseOfRepresentativeV1ResponseTypeArray  ItemResponseOfRepresentativeV1ResponseType = "array"
	ItemResponseOfRepresentativeV1ResponseTypeError  ItemResponseOfRepresentativeV1ResponseType = "error"
	ItemResponseOfRepresentativeV1ResponseTypeNone   ItemResponseOfRepresentativeV1ResponseType = "none"
)

func (r ItemResponseOfRepresentativeV1ResponseType) IsKnown() bool {
	switch r {
	case ItemResponseOfRepresentativeV1ResponseTypeObject, ItemResponseOfRepresentativeV1ResponseTypeArray, ItemResponseOfRepresentativeV1ResponseTypeError, ItemResponseOfRepresentativeV1ResponseTypeNone:
		return true
	}
	return false
}

type LinkedBankAccountV1 struct {
	// Unique identifier for the linked bank account.
	ID string `json:"id,required" format:"uuid"`
	// The unique identifier of the Straddle account related to this bank account.
	AccountID   string                         `json:"account_id,required" format:"uuid"`
	BankAccount LinkedBankAccountV1BankAccount `json:"bank_account,required"`
	// Timestamp of when the bank account object was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The current status of the linked bank account.
	Status       LinkedBankAccountV1Status                       `json:"status,required"`
	StatusDetail StatusDetailOfLinkedBankAccountStatusDetailEnum `json:"status_detail,required"`
	// Timestamp of the most recent update to the linked bank account.
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// Up to 20 additional user-defined key-value pairs. Useful for storing additional
	// information about the linked bank account in a structured format.
	Metadata map[string]string       `json:"metadata,nullable"`
	JSON     linkedBankAccountV1JSON `json:"-"`
}

// linkedBankAccountV1JSON contains the JSON metadata for the struct
// [LinkedBankAccountV1]
type linkedBankAccountV1JSON struct {
	ID           apijson.Field
	AccountID    apijson.Field
	BankAccount  apijson.Field
	CreatedAt    apijson.Field
	Status       apijson.Field
	StatusDetail apijson.Field
	UpdatedAt    apijson.Field
	Metadata     apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *LinkedBankAccountV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r linkedBankAccountV1JSON) RawJSON() string {
	return r.raw
}

type LinkedBankAccountV1BankAccount struct {
	AccountHolder   string                             `json:"account_holder,required"`
	AccountMask     string                             `json:"account_mask,required"`
	InstitutionName string                             `json:"institution_name,required"`
	RoutingNumber   string                             `json:"routing_number,required"`
	JSON            linkedBankAccountV1BankAccountJSON `json:"-"`
}

// linkedBankAccountV1BankAccountJSON contains the JSON metadata for the struct
// [LinkedBankAccountV1BankAccount]
type linkedBankAccountV1BankAccountJSON struct {
	AccountHolder   apijson.Field
	AccountMask     apijson.Field
	InstitutionName apijson.Field
	RoutingNumber   apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *LinkedBankAccountV1BankAccount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r linkedBankAccountV1BankAccountJSON) RawJSON() string {
	return r.raw
}

// The current status of the linked bank account.
type LinkedBankAccountV1Status string

const (
	LinkedBankAccountV1StatusCreated    LinkedBankAccountV1Status = "created"
	LinkedBankAccountV1StatusOnboarding LinkedBankAccountV1Status = "onboarding"
	LinkedBankAccountV1StatusActive     LinkedBankAccountV1Status = "active"
	LinkedBankAccountV1StatusRejected   LinkedBankAccountV1Status = "rejected"
	LinkedBankAccountV1StatusInactive   LinkedBankAccountV1Status = "inactive"
)

func (r LinkedBankAccountV1Status) IsKnown() bool {
	switch r {
	case LinkedBankAccountV1StatusCreated, LinkedBankAccountV1StatusOnboarding, LinkedBankAccountV1StatusActive, LinkedBankAccountV1StatusRejected, LinkedBankAccountV1StatusInactive:
		return true
	}
	return false
}

type OrganizationV1 struct {
	// Straddle's unique identifier for the organization.
	ID string `json:"id,required" format:"uuid"`
	// Timestamp of when the organization was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The name of the organization.
	Name string `json:"name,required"`
	// Timestamp of the most recent update to the organization.
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// Unique identifier for the organization in your database, used for
	// cross-referencing between Straddle and your systems.
	ExternalID string `json:"external_id,nullable"`
	// Up to 20 additional user-defined key-value pairs. Useful for storing additional
	// information about the organization in a structured format.
	Metadata map[string]string  `json:"metadata,nullable"`
	JSON     organizationV1JSON `json:"-"`
}

// organizationV1JSON contains the JSON metadata for the struct [OrganizationV1]
type organizationV1JSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Name        apijson.Field
	UpdatedAt   apijson.Field
	ExternalID  apijson.Field
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationV1JSON) RawJSON() string {
	return r.raw
}

type PagedResponseMetadata1 struct {
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
	SortBy     string    `json:"sort_by,required"`
	SortOrder  SortOrder `json:"sort_order,required"`
	TotalItems int64     `json:"total_items,required"`
	// The number of pages available.
	TotalPages int64                      `json:"total_pages,required"`
	JSON       pagedResponseMetadata1JSON `json:"-"`
}

// pagedResponseMetadata1JSON contains the JSON metadata for the struct
// [PagedResponseMetadata1]
type pagedResponseMetadata1JSON struct {
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

func (r *PagedResponseMetadata1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pagedResponseMetadata1JSON) RawJSON() string {
	return r.raw
}

type PagedResponseMetadata2 struct {
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
	SortBy     string    `json:"sort_by,required"`
	SortOrder  SortOrder `json:"sort_order,required"`
	TotalItems int64     `json:"total_items,required"`
	// The number of pages available.
	TotalPages int64                      `json:"total_pages,required"`
	JSON       pagedResponseMetadata2JSON `json:"-"`
}

// pagedResponseMetadata2JSON contains the JSON metadata for the struct
// [PagedResponseMetadata2]
type pagedResponseMetadata2JSON struct {
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

func (r *PagedResponseMetadata2) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pagedResponseMetadata2JSON) RawJSON() string {
	return r.raw
}

type PagedResponseOfCapabilityRequestV1 struct {
	Data []PagedResponseOfCapabilityRequestV1Data `json:"data,required"`
	// Metadata about the API request, including an identifier, timestamp, and
	// pagination details.
	Meta PagedResponseMetadata `json:"meta,required"`
	// Indicates the structure of the returned content.
	//
	//   - "object" means the `data` field contains a single JSON object.
	//   - "array" means the `data` field contains an array of objects.
	//   - "error" means the `data` field contains an error object with details of the
	//     issue.
	//   - "none" means no data is returned.
	ResponseType PagedResponseOfCapabilityRequestV1ResponseType `json:"response_type,required"`
	JSON         pagedResponseOfCapabilityRequestV1JSON         `json:"-"`
}

// pagedResponseOfCapabilityRequestV1JSON contains the JSON metadata for the struct
// [PagedResponseOfCapabilityRequestV1]
type pagedResponseOfCapabilityRequestV1JSON struct {
	Data         apijson.Field
	Meta         apijson.Field
	ResponseType apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *PagedResponseOfCapabilityRequestV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pagedResponseOfCapabilityRequestV1JSON) RawJSON() string {
	return r.raw
}

type PagedResponseOfCapabilityRequestV1Data struct {
	// Unique identifier for the capability request.
	ID string `json:"id,required" format:"uuid"`
	// The unique identifier of the account associated with this capability request.
	AccountID string `json:"account_id,required" format:"uuid"`
	// The category of the requested capability. Use `payment_type` for charges and
	// payouts, `customer_type` to define `individuals` or `businesses`, and
	// `consent_type` for `signed_agreement` or `internet` payment authorization.
	Category PagedResponseOfCapabilityRequestV1DataCategory `json:"category,required"`
	// Timestamp of when the capability request was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The current status of the capability request.
	Status PagedResponseOfCapabilityRequestV1DataStatus `json:"status,required"`
	// The specific type of capability being requested within the category.
	Type PagedResponseOfCapabilityRequestV1DataType `json:"type,required"`
	// Timestamp of the most recent update to the capability request.
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// Any specific settings or configurations related to the requested capability.
	Settings map[string]interface{}                     `json:"settings,nullable"`
	JSON     pagedResponseOfCapabilityRequestV1DataJSON `json:"-"`
}

// pagedResponseOfCapabilityRequestV1DataJSON contains the JSON metadata for the
// struct [PagedResponseOfCapabilityRequestV1Data]
type pagedResponseOfCapabilityRequestV1DataJSON struct {
	ID          apijson.Field
	AccountID   apijson.Field
	Category    apijson.Field
	CreatedAt   apijson.Field
	Status      apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	Settings    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PagedResponseOfCapabilityRequestV1Data) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pagedResponseOfCapabilityRequestV1DataJSON) RawJSON() string {
	return r.raw
}

// The category of the requested capability. Use `payment_type` for charges and
// payouts, `customer_type` to define `individuals` or `businesses`, and
// `consent_type` for `signed_agreement` or `internet` payment authorization.
type PagedResponseOfCapabilityRequestV1DataCategory string

const (
	PagedResponseOfCapabilityRequestV1DataCategoryPaymentType  PagedResponseOfCapabilityRequestV1DataCategory = "payment_type"
	PagedResponseOfCapabilityRequestV1DataCategoryCustomerType PagedResponseOfCapabilityRequestV1DataCategory = "customer_type"
	PagedResponseOfCapabilityRequestV1DataCategoryConsentType  PagedResponseOfCapabilityRequestV1DataCategory = "consent_type"
)

func (r PagedResponseOfCapabilityRequestV1DataCategory) IsKnown() bool {
	switch r {
	case PagedResponseOfCapabilityRequestV1DataCategoryPaymentType, PagedResponseOfCapabilityRequestV1DataCategoryCustomerType, PagedResponseOfCapabilityRequestV1DataCategoryConsentType:
		return true
	}
	return false
}

// The current status of the capability request.
type PagedResponseOfCapabilityRequestV1DataStatus string

const (
	PagedResponseOfCapabilityRequestV1DataStatusActive    PagedResponseOfCapabilityRequestV1DataStatus = "active"
	PagedResponseOfCapabilityRequestV1DataStatusInactive  PagedResponseOfCapabilityRequestV1DataStatus = "inactive"
	PagedResponseOfCapabilityRequestV1DataStatusInReview  PagedResponseOfCapabilityRequestV1DataStatus = "in_review"
	PagedResponseOfCapabilityRequestV1DataStatusRejected  PagedResponseOfCapabilityRequestV1DataStatus = "rejected"
	PagedResponseOfCapabilityRequestV1DataStatusApproved  PagedResponseOfCapabilityRequestV1DataStatus = "approved"
	PagedResponseOfCapabilityRequestV1DataStatusReviewing PagedResponseOfCapabilityRequestV1DataStatus = "reviewing"
)

func (r PagedResponseOfCapabilityRequestV1DataStatus) IsKnown() bool {
	switch r {
	case PagedResponseOfCapabilityRequestV1DataStatusActive, PagedResponseOfCapabilityRequestV1DataStatusInactive, PagedResponseOfCapabilityRequestV1DataStatusInReview, PagedResponseOfCapabilityRequestV1DataStatusRejected, PagedResponseOfCapabilityRequestV1DataStatusApproved, PagedResponseOfCapabilityRequestV1DataStatusReviewing:
		return true
	}
	return false
}

// The specific type of capability being requested within the category.
type PagedResponseOfCapabilityRequestV1DataType string

const (
	PagedResponseOfCapabilityRequestV1DataTypeCharges         PagedResponseOfCapabilityRequestV1DataType = "charges"
	PagedResponseOfCapabilityRequestV1DataTypePayouts         PagedResponseOfCapabilityRequestV1DataType = "payouts"
	PagedResponseOfCapabilityRequestV1DataTypeIndividuals     PagedResponseOfCapabilityRequestV1DataType = "individuals"
	PagedResponseOfCapabilityRequestV1DataTypeBusinesses      PagedResponseOfCapabilityRequestV1DataType = "businesses"
	PagedResponseOfCapabilityRequestV1DataTypeSignedAgreement PagedResponseOfCapabilityRequestV1DataType = "signed_agreement"
	PagedResponseOfCapabilityRequestV1DataTypeInternet        PagedResponseOfCapabilityRequestV1DataType = "internet"
)

func (r PagedResponseOfCapabilityRequestV1DataType) IsKnown() bool {
	switch r {
	case PagedResponseOfCapabilityRequestV1DataTypeCharges, PagedResponseOfCapabilityRequestV1DataTypePayouts, PagedResponseOfCapabilityRequestV1DataTypeIndividuals, PagedResponseOfCapabilityRequestV1DataTypeBusinesses, PagedResponseOfCapabilityRequestV1DataTypeSignedAgreement, PagedResponseOfCapabilityRequestV1DataTypeInternet:
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
type PagedResponseOfCapabilityRequestV1ResponseType string

const (
	PagedResponseOfCapabilityRequestV1ResponseTypeObject PagedResponseOfCapabilityRequestV1ResponseType = "object"
	PagedResponseOfCapabilityRequestV1ResponseTypeArray  PagedResponseOfCapabilityRequestV1ResponseType = "array"
	PagedResponseOfCapabilityRequestV1ResponseTypeError  PagedResponseOfCapabilityRequestV1ResponseType = "error"
	PagedResponseOfCapabilityRequestV1ResponseTypeNone   PagedResponseOfCapabilityRequestV1ResponseType = "none"
)

func (r PagedResponseOfCapabilityRequestV1ResponseType) IsKnown() bool {
	switch r {
	case PagedResponseOfCapabilityRequestV1ResponseTypeObject, PagedResponseOfCapabilityRequestV1ResponseTypeArray, PagedResponseOfCapabilityRequestV1ResponseTypeError, PagedResponseOfCapabilityRequestV1ResponseTypeNone:
		return true
	}
	return false
}

// Metadata about the API request, including an identifier, timestamp, and
// pagination details.
type PagedResponseMetadata struct {
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
	SortBy string `json:"sort_by,required"`
	// The order that the results were sorted by.
	SortOrder PagedResponseMetadataSortOrder `json:"sort_order,required"`
	// Total number of items returned in this response.
	TotalItems int64 `json:"total_items,required"`
	// The number of pages available.
	TotalPages int64                     `json:"total_pages,required"`
	JSON       pagedResponseMetadataJSON `json:"-"`
}

// pagedResponseMetadataJSON contains the JSON metadata for the struct
// [PagedResponseMetadata]
type pagedResponseMetadataJSON struct {
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

func (r *PagedResponseMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pagedResponseMetadataJSON) RawJSON() string {
	return r.raw
}

// The order that the results were sorted by.
type PagedResponseMetadataSortOrder string

const (
	PagedResponseMetadataSortOrderAsc  PagedResponseMetadataSortOrder = "asc"
	PagedResponseMetadataSortOrderDesc PagedResponseMetadataSortOrder = "desc"
)

func (r PagedResponseMetadataSortOrder) IsKnown() bool {
	switch r {
	case PagedResponseMetadataSortOrderAsc, PagedResponseMetadataSortOrderDesc:
		return true
	}
	return false
}

type PaykeyBankDetailsV1 struct {
	// Bank account number. This value is masked by default for security reasons. Use
	// the /unmask endpoint to access the unmasked value.
	AccountNumber string        `json:"account_number,required"`
	AccountType   AccountTypeV1 `json:"account_type,required"`
	// The routing number of the bank account.
	RoutingNumber string                  `json:"routing_number,required"`
	JSON          paykeyBankDetailsV1JSON `json:"-"`
}

// paykeyBankDetailsV1JSON contains the JSON metadata for the struct
// [PaykeyBankDetailsV1]
type paykeyBankDetailsV1JSON struct {
	AccountNumber apijson.Field
	AccountType   apijson.Field
	RoutingNumber apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *PaykeyBankDetailsV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r paykeyBankDetailsV1JSON) RawJSON() string {
	return r.raw
}

type PaykeySourceV1 string

const (
	PaykeySourceV1BankAccount PaykeySourceV1 = "bank_account"
	PaykeySourceV1Straddle    PaykeySourceV1 = "straddle"
	PaykeySourceV1Mx          PaykeySourceV1 = "mx"
	PaykeySourceV1Plaid       PaykeySourceV1 = "plaid"
)

func (r PaykeySourceV1) IsKnown() bool {
	switch r {
	case PaykeySourceV1BankAccount, PaykeySourceV1Straddle, PaykeySourceV1Mx, PaykeySourceV1Plaid:
		return true
	}
	return false
}

type PaykeyStatusV1 string

const (
	PaykeyStatusV1Pending  PaykeyStatusV1 = "pending"
	PaykeyStatusV1Active   PaykeyStatusV1 = "active"
	PaykeyStatusV1Inactive PaykeyStatusV1 = "inactive"
	PaykeyStatusV1Rejected PaykeyStatusV1 = "rejected"
)

func (r PaykeyStatusV1) IsKnown() bool {
	switch r {
	case PaykeyStatusV1Pending, PaykeyStatusV1Active, PaykeyStatusV1Inactive, PaykeyStatusV1Rejected:
		return true
	}
	return false
}

type PaykeyV1ItemResponse struct {
	Data PaykeyV1ItemResponseData `json:"data,required"`
	// Metadata about the API request, including an identifier and timestamp.
	Meta ResponseMetadata `json:"meta,required"`
	// Indicates the structure of the returned content.
	//
	//   - "object" means the `data` field contains a single JSON object.
	//   - "array" means the `data` field contains an array of objects.
	//   - "error" means the `data` field contains an error object with details of the
	//     issue.
	//   - "none" means no data is returned.
	ResponseType ResponseTypeEnum         `json:"response_type,required"`
	JSON         paykeyV1ItemResponseJSON `json:"-"`
}

// paykeyV1ItemResponseJSON contains the JSON metadata for the struct
// [PaykeyV1ItemResponse]
type paykeyV1ItemResponseJSON struct {
	Data         apijson.Field
	Meta         apijson.Field
	ResponseType apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *PaykeyV1ItemResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r paykeyV1ItemResponseJSON) RawJSON() string {
	return r.raw
}

type PaykeyV1ItemResponseData struct {
	// Unique identifier for the paykey.
	ID string `json:"id,required" format:"uuid"`
	// Timestamp of when the paykey was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Human-readable label used to represent this paykey in a UI.
	Label string `json:"label,required"`
	// The tokenized paykey value. This value is used to create payments and should be
	// stored securely.
	Paykey string         `json:"paykey,required"`
	Source PaykeySourceV1 `json:"source,required"`
	Status PaykeyStatusV1 `json:"status,required"`
	// Timestamp of the most recent update to the paykey.
	UpdatedAt time.Time           `json:"updated_at,required" format:"date-time"`
	BankData  PaykeyBankDetailsV1 `json:"bank_data"`
	// Unique identifier of the related customer object.
	CustomerID string `json:"customer_id,nullable" format:"uuid"`
	// Expiration date and time of the paykey, if applicable.
	ExpiresAt time.Time `json:"expires_at,nullable" format:"date-time"`
	// Name of the financial institution.
	InstitutionName string `json:"institution_name,nullable"`
	// Up to 20 additional user-defined key-value pairs. Useful for storing additional
	// information about the paykey in a structured format.
	Metadata      map[string]string            `json:"metadata,nullable"`
	StatusDetails StatusDetailsV1              `json:"status_details"`
	JSON          paykeyV1ItemResponseDataJSON `json:"-"`
}

// paykeyV1ItemResponseDataJSON contains the JSON metadata for the struct
// [PaykeyV1ItemResponseData]
type paykeyV1ItemResponseDataJSON struct {
	ID              apijson.Field
	CreatedAt       apijson.Field
	Label           apijson.Field
	Paykey          apijson.Field
	Source          apijson.Field
	Status          apijson.Field
	UpdatedAt       apijson.Field
	BankData        apijson.Field
	CustomerID      apijson.Field
	ExpiresAt       apijson.Field
	InstitutionName apijson.Field
	Metadata        apijson.Field
	StatusDetails   apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *PaykeyV1ItemResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r paykeyV1ItemResponseDataJSON) RawJSON() string {
	return r.raw
}

type PaykeyDetailsV1 struct {
	// Unique identifier for the paykey.
	ID string `json:"id,required" format:"uuid"`
	// Unique identifier for the customer associated with the paykey.
	CustomerID string `json:"customer_id,required" format:"uuid"`
	// Human-readable label used to represent this paykey in a UI.
	Label string `json:"label,required"`
	// The most recent balance of the bank account associated with the paykey in
	// dollars.
	Balance int64               `json:"balance,nullable"`
	JSON    paykeyDetailsV1JSON `json:"-"`
}

// paykeyDetailsV1JSON contains the JSON metadata for the struct [PaykeyDetailsV1]
type paykeyDetailsV1JSON struct {
	ID          apijson.Field
	CustomerID  apijson.Field
	Label       apijson.Field
	Balance     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PaykeyDetailsV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r paykeyDetailsV1JSON) RawJSON() string {
	return r.raw
}

// The payment rail used for the charge or payout.
type PaymentRailV1 string

const (
	PaymentRailV1ACH PaymentRailV1 = "ach"
)

func (r PaymentRailV1) IsKnown() bool {
	switch r {
	case PaymentRailV1ACH:
		return true
	}
	return false
}

// The field to sort the results by.
type PaymentSortByV1 string

const (
	PaymentSortByV1CreatedAt   PaymentSortByV1 = "created_at"
	PaymentSortByV1PaymentDate PaymentSortByV1 = "payment_date"
	PaymentSortByV1EffectiveAt PaymentSortByV1 = "effective_at"
	PaymentSortByV1ID          PaymentSortByV1 = "id"
	PaymentSortByV1Amount      PaymentSortByV1 = "amount"
)

func (r PaymentSortByV1) IsKnown() bool {
	switch r {
	case PaymentSortByV1CreatedAt, PaymentSortByV1PaymentDate, PaymentSortByV1EffectiveAt, PaymentSortByV1ID, PaymentSortByV1Amount:
		return true
	}
	return false
}

// The current status of the `charge` or `payout`.
type PaymentStatusV1 string

const (
	PaymentStatusV1Created   PaymentStatusV1 = "created"
	PaymentStatusV1Scheduled PaymentStatusV1 = "scheduled"
	PaymentStatusV1Failed    PaymentStatusV1 = "failed"
	PaymentStatusV1Cancelled PaymentStatusV1 = "cancelled"
	PaymentStatusV1OnHold    PaymentStatusV1 = "on_hold"
	PaymentStatusV1Pending   PaymentStatusV1 = "pending"
	PaymentStatusV1Paid      PaymentStatusV1 = "paid"
	PaymentStatusV1Reversed  PaymentStatusV1 = "reversed"
)

func (r PaymentStatusV1) IsKnown() bool {
	switch r {
	case PaymentStatusV1Created, PaymentStatusV1Scheduled, PaymentStatusV1Failed, PaymentStatusV1Cancelled, PaymentStatusV1OnHold, PaymentStatusV1Pending, PaymentStatusV1Paid, PaymentStatusV1Reversed:
		return true
	}
	return false
}

// The type of payment.
type PaymentTypeV1 string

const (
	PaymentTypeV1Charge PaymentTypeV1 = "charge"
	PaymentTypeV1Payout PaymentTypeV1 = "payout"
)

func (r PaymentTypeV1) IsKnown() bool {
	switch r {
	case PaymentTypeV1Charge, PaymentTypeV1Payout:
		return true
	}
	return false
}

type PayoutV1ItemResponse struct {
	Data PayoutV1ItemResponseData `json:"data,required"`
	// Metadata about the API request, including an identifier and timestamp.
	Meta ResponseMetadata `json:"meta,required"`
	// Indicates the structure of the returned content.
	//
	//   - "object" means the `data` field contains a single JSON object.
	//   - "array" means the `data` field contains an array of objects.
	//   - "error" means the `data` field contains an error object with details of the
	//     issue.
	//   - "none" means no data is returned.
	ResponseType ResponseTypeEnum         `json:"response_type,required"`
	JSON         payoutV1ItemResponseJSON `json:"-"`
}

// payoutV1ItemResponseJSON contains the JSON metadata for the struct
// [PayoutV1ItemResponse]
type payoutV1ItemResponseJSON struct {
	Data         apijson.Field
	Meta         apijson.Field
	ResponseType apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *PayoutV1ItemResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r payoutV1ItemResponseJSON) RawJSON() string {
	return r.raw
}

type PayoutV1ItemResponseData struct {
	// Unique identifier for the payout.
	ID string `json:"id,required" format:"uuid"`
	// The amount of the payout in cents.
	Amount int64 `json:"amount,required"`
	// Configuration for the payout.
	Config interface{} `json:"config,required"`
	// The currency of the payout. Only USD is supported.
	Currency string `json:"currency,required"`
	// An arbitrary description for the payout.
	Description string `json:"description,required"`
	// Information about the device used when the customer authorized the payout.
	Device DeviceInfoV1 `json:"device,required"`
	// Unique identifier for the payout in your database. This value must be unique
	// across all payouts.
	ExternalID string `json:"external_id,required"`
	// Value of the `paykey` used for the payout.
	Paykey string `json:"paykey,required"`
	// The desired date on which the payment should be occur. For payouts, this means
	// the date you want the funds to be sent from your bank account.
	PaymentDate time.Time `json:"payment_date,required" format:"date"`
	// The current status of the payout.
	Status PaymentStatusV1 `json:"status,required"`
	// Details about the current status of the payout.
	StatusDetails StatusDetailsV1 `json:"status_details,required"`
	// History of the status changes for the payout.
	StatusHistory []StatusHistoryV1 `json:"status_history,required"`
	// The time the payout was created.
	CreatedAt time.Time `json:"created_at,nullable" format:"date-time"`
	// Information about the customer associated with the payout.
	CustomerDetails CustomerDetailsV1 `json:"customer_details"`
	// The actual date on which the payment occurred. For payouts, this is the date the
	// funds were sent from your bank account.
	EffectiveAt time.Time `json:"effective_at,nullable" format:"date-time"`
	// Up to 20 additional user-defined key-value pairs. Useful for storing additional
	// information about the payout in a structured format.
	Metadata map[string]string `json:"metadata,nullable"`
	// Information about the paykey used for the payout.
	PaykeyDetails PaykeyDetailsV1 `json:"paykey_details"`
	// The payment rail used for the payout.
	PaymentRail PaymentRailV1 `json:"payment_rail"`
	// The time the payout was processed by Straddle and originated to the payment
	// rail.
	ProcessedAt time.Time `json:"processed_at,nullable" format:"date-time"`
	// The time the payout was last updated.
	UpdatedAt time.Time                    `json:"updated_at,nullable" format:"date-time"`
	JSON      payoutV1ItemResponseDataJSON `json:"-"`
}

// payoutV1ItemResponseDataJSON contains the JSON metadata for the struct
// [PayoutV1ItemResponseData]
type payoutV1ItemResponseDataJSON struct {
	ID              apijson.Field
	Amount          apijson.Field
	Config          apijson.Field
	Currency        apijson.Field
	Description     apijson.Field
	Device          apijson.Field
	ExternalID      apijson.Field
	Paykey          apijson.Field
	PaymentDate     apijson.Field
	Status          apijson.Field
	StatusDetails   apijson.Field
	StatusHistory   apijson.Field
	CreatedAt       apijson.Field
	CustomerDetails apijson.Field
	EffectiveAt     apijson.Field
	Metadata        apijson.Field
	PaykeyDetails   apijson.Field
	PaymentRail     apijson.Field
	ProcessedAt     apijson.Field
	UpdatedAt       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *PayoutV1ItemResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r payoutV1ItemResponseDataJSON) RawJSON() string {
	return r.raw
}

type RelationshipV1Param struct {
	// Whether the representative has significant responsibility to control, manage, or
	// direct the organization. One representative must be identified under the control
	// prong for each legal entity.
	Control param.Field[bool] `json:"control,required"`
	// Whether the representative owns any percentage of of the equity interests of the
	// legal entity.
	Owner param.Field[bool] `json:"owner,required"`
	// Whether the person is authorized as the primary representative of the account.
	// This is the person chosen by the business to provide information about
	// themselves, general information about the account, and who consented to the
	// services agreement.
	//
	// There can be only one primary representative for an account at a time.
	Primary param.Field[bool] `json:"primary,required"`
	// The percentage of ownership the representative has. Required if 'Owner' is true.
	PercentOwnership param.Field[float64] `json:"percent_ownership"`
	// The job title of the representative.
	Title param.Field[string] `json:"title"`
}

func (r RelationshipV1Param) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RepresentativeV1 struct {
	// Unique identifier for the representative.
	ID string `json:"id,required" format:"uuid"`
	// The unique identifier of the account this representative is associated with.
	AccountID string `json:"account_id,required" format:"uuid"`
	// Timestamp of when the representative was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The date of birth of the representative, in ISO 8601 format (YYYY-MM-DD).
	Dob time.Time `json:"dob,required" format:"date"`
	// The email address of the representative.
	Email string `json:"email,required" format:"email"`
	// The first name of the representative.
	FirstName string `json:"first_name,required"`
	// The last name of the representative.
	LastName string `json:"last_name,required"`
	// The mobile phone number of the representative.
	MobileNumber string                       `json:"mobile_number,required"`
	Relationship RepresentativeV1Relationship `json:"relationship,required"`
	// The last 4 digits of the representative's Social Security Number.
	SsnLast4 string `json:"ssn_last4,required"`
	// The current status of the representative.
	Status       RepresentativeV1Status       `json:"status,required"`
	StatusDetail RepresentativeV1StatusDetail `json:"status_detail,required"`
	// Timestamp of the most recent update to the representative.
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// Unique identifier for the representative in your database, used for
	// cross-referencing between Straddle and your systems.
	ExternalID string `json:"external_id,nullable"`
	// The unique identifier of the user account associated with this representative,
	// if applicable.
	UserID string               `json:"user_id,nullable" format:"uuid"`
	JSON   representativeV1JSON `json:"-"`
}

// representativeV1JSON contains the JSON metadata for the struct
// [RepresentativeV1]
type representativeV1JSON struct {
	ID           apijson.Field
	AccountID    apijson.Field
	CreatedAt    apijson.Field
	Dob          apijson.Field
	Email        apijson.Field
	FirstName    apijson.Field
	LastName     apijson.Field
	MobileNumber apijson.Field
	Relationship apijson.Field
	SsnLast4     apijson.Field
	Status       apijson.Field
	StatusDetail apijson.Field
	UpdatedAt    apijson.Field
	ExternalID   apijson.Field
	UserID       apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RepresentativeV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r representativeV1JSON) RawJSON() string {
	return r.raw
}

type RepresentativeV1Relationship struct {
	// Whether the representative has significant responsibility to control, manage, or
	// direct the organization. One representative must be identified under the control
	// prong for each legal entity.
	Control bool `json:"control,required"`
	// Whether the representative owns any percentage of of the equity interests of the
	// legal entity.
	Owner bool `json:"owner,required"`
	// Whether the person is authorized as the primary representative of the account.
	// This is the person chosen by the business to provide information about
	// themselves, general information about the account, and who consented to the
	// services agreement.
	//
	// There can be only one primary representative for an account at a time.
	Primary bool `json:"primary,required"`
	// The percentage of ownership the representative has. Required if 'Owner' is true.
	PercentOwnership float64 `json:"percent_ownership,nullable"`
	// The job title of the representative.
	Title string                           `json:"title,nullable"`
	JSON  representativeV1RelationshipJSON `json:"-"`
}

// representativeV1RelationshipJSON contains the JSON metadata for the struct
// [RepresentativeV1Relationship]
type representativeV1RelationshipJSON struct {
	Control          apijson.Field
	Owner            apijson.Field
	Primary          apijson.Field
	PercentOwnership apijson.Field
	Title            apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RepresentativeV1Relationship) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r representativeV1RelationshipJSON) RawJSON() string {
	return r.raw
}

// The current status of the representative.
type RepresentativeV1Status string

const (
	RepresentativeV1StatusCreated    RepresentativeV1Status = "created"
	RepresentativeV1StatusOnboarding RepresentativeV1Status = "onboarding"
	RepresentativeV1StatusActive     RepresentativeV1Status = "active"
	RepresentativeV1StatusRejected   RepresentativeV1Status = "rejected"
	RepresentativeV1StatusInactive   RepresentativeV1Status = "inactive"
)

func (r RepresentativeV1Status) IsKnown() bool {
	switch r {
	case RepresentativeV1StatusCreated, RepresentativeV1StatusOnboarding, RepresentativeV1StatusActive, RepresentativeV1StatusRejected, RepresentativeV1StatusInactive:
		return true
	}
	return false
}

type RepresentativeV1StatusDetail struct {
	// A machine-readable code for the specific status, useful for programmatic
	// handling.
	Code string `json:"code,required"`
	// A human-readable message describing the current status.
	Message string `json:"message,required"`
	// A machine-readable identifier for the specific status, useful for programmatic
	// handling.
	Reason RepresentativeV1StatusDetailReason `json:"reason,required"`
	// Identifies the origin of the status change (e.g., `watchtower`). This helps in
	// tracking the cause of status updates.
	Source RepresentativeV1StatusDetailSource `json:"source,required"`
	JSON   representativeV1StatusDetailJSON   `json:"-"`
}

// representativeV1StatusDetailJSON contains the JSON metadata for the struct
// [RepresentativeV1StatusDetail]
type representativeV1StatusDetailJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	Reason      apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RepresentativeV1StatusDetail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r representativeV1StatusDetailJSON) RawJSON() string {
	return r.raw
}

// A machine-readable identifier for the specific status, useful for programmatic
// handling.
type RepresentativeV1StatusDetailReason string

const (
	RepresentativeV1StatusDetailReasonUnverified         RepresentativeV1StatusDetailReason = "unverified"
	RepresentativeV1StatusDetailReasonInReview           RepresentativeV1StatusDetailReason = "in_review"
	RepresentativeV1StatusDetailReasonPending            RepresentativeV1StatusDetailReason = "pending"
	RepresentativeV1StatusDetailReasonStuck              RepresentativeV1StatusDetailReason = "stuck"
	RepresentativeV1StatusDetailReasonVerified           RepresentativeV1StatusDetailReason = "verified"
	RepresentativeV1StatusDetailReasonFailedVerification RepresentativeV1StatusDetailReason = "failed_verification"
	RepresentativeV1StatusDetailReasonDisabled           RepresentativeV1StatusDetailReason = "disabled"
	RepresentativeV1StatusDetailReasonNew                RepresentativeV1StatusDetailReason = "new"
)

func (r RepresentativeV1StatusDetailReason) IsKnown() bool {
	switch r {
	case RepresentativeV1StatusDetailReasonUnverified, RepresentativeV1StatusDetailReasonInReview, RepresentativeV1StatusDetailReasonPending, RepresentativeV1StatusDetailReasonStuck, RepresentativeV1StatusDetailReasonVerified, RepresentativeV1StatusDetailReasonFailedVerification, RepresentativeV1StatusDetailReasonDisabled, RepresentativeV1StatusDetailReasonNew:
		return true
	}
	return false
}

// Identifies the origin of the status change (e.g., `watchtower`). This helps in
// tracking the cause of status updates.
type RepresentativeV1StatusDetailSource string

const (
	RepresentativeV1StatusDetailSourceWatchtower RepresentativeV1StatusDetailSource = "watchtower"
)

func (r RepresentativeV1StatusDetailSource) IsKnown() bool {
	switch r {
	case RepresentativeV1StatusDetailSourceWatchtower:
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
type ResponseTypeEnum string

const (
	ResponseTypeEnumObject ResponseTypeEnum = "object"
	ResponseTypeEnumArray  ResponseTypeEnum = "array"
	ResponseTypeEnumError  ResponseTypeEnum = "error"
	ResponseTypeEnumNone   ResponseTypeEnum = "none"
)

func (r ResponseTypeEnum) IsKnown() bool {
	switch r {
	case ResponseTypeEnumObject, ResponseTypeEnumArray, ResponseTypeEnumError, ResponseTypeEnumNone:
		return true
	}
	return false
}

// Metadata about the API request, including an identifier and timestamp.
type ResponseMetadata struct {
	// Unique identifier for this API request, useful for troubleshooting.
	APIRequestID string `json:"api_request_id,required" format:"uuid"`
	// Timestamp for this API request, useful for troubleshooting.
	APIRequestTimestamp time.Time            `json:"api_request_timestamp,required" format:"date-time"`
	JSON                responseMetadataJSON `json:"-"`
}

// responseMetadataJSON contains the JSON metadata for the struct
// [ResponseMetadata]
type responseMetadataJSON struct {
	APIRequestID        apijson.Field
	APIRequestTimestamp apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *ResponseMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r responseMetadataJSON) RawJSON() string {
	return r.raw
}

type SortOrder string

const (
	SortOrderAsc  SortOrder = "asc"
	SortOrderDesc SortOrder = "desc"
)

func (r SortOrder) IsKnown() bool {
	switch r {
	case SortOrderAsc, SortOrderDesc:
		return true
	}
	return false
}

type StatusDetailOfLinkedBankAccountStatusDetailEnum struct {
	// A machine-readable code for the specific status, useful for programmatic
	// handling.
	Code string `json:"code,required"`
	// A human-readable message describing the current status.
	Message string `json:"message,required"`
	// A machine-readable identifier for the specific status, useful for programmatic
	// handling.
	Reason StatusDetailOfLinkedBankAccountStatusDetailEnumReason `json:"reason,required"`
	// Identifies the origin of the status change (e.g., `watchtower`). This helps in
	// tracking the cause of status updates.
	Source StatusDetailOfLinkedBankAccountStatusDetailEnumSource `json:"source,required"`
	JSON   statusDetailOfLinkedBankAccountStatusDetailEnumJSON   `json:"-"`
}

// statusDetailOfLinkedBankAccountStatusDetailEnumJSON contains the JSON metadata
// for the struct [StatusDetailOfLinkedBankAccountStatusDetailEnum]
type statusDetailOfLinkedBankAccountStatusDetailEnumJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	Reason      apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StatusDetailOfLinkedBankAccountStatusDetailEnum) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r statusDetailOfLinkedBankAccountStatusDetailEnumJSON) RawJSON() string {
	return r.raw
}

// A machine-readable identifier for the specific status, useful for programmatic
// handling.
type StatusDetailOfLinkedBankAccountStatusDetailEnumReason string

const (
	StatusDetailOfLinkedBankAccountStatusDetailEnumReasonUnverified         StatusDetailOfLinkedBankAccountStatusDetailEnumReason = "unverified"
	StatusDetailOfLinkedBankAccountStatusDetailEnumReasonInReview           StatusDetailOfLinkedBankAccountStatusDetailEnumReason = "in_review"
	StatusDetailOfLinkedBankAccountStatusDetailEnumReasonPending            StatusDetailOfLinkedBankAccountStatusDetailEnumReason = "pending"
	StatusDetailOfLinkedBankAccountStatusDetailEnumReasonStuck              StatusDetailOfLinkedBankAccountStatusDetailEnumReason = "stuck"
	StatusDetailOfLinkedBankAccountStatusDetailEnumReasonVerified           StatusDetailOfLinkedBankAccountStatusDetailEnumReason = "verified"
	StatusDetailOfLinkedBankAccountStatusDetailEnumReasonFailedVerification StatusDetailOfLinkedBankAccountStatusDetailEnumReason = "failed_verification"
	StatusDetailOfLinkedBankAccountStatusDetailEnumReasonDisabled           StatusDetailOfLinkedBankAccountStatusDetailEnumReason = "disabled"
	StatusDetailOfLinkedBankAccountStatusDetailEnumReasonNew                StatusDetailOfLinkedBankAccountStatusDetailEnumReason = "new"
)

func (r StatusDetailOfLinkedBankAccountStatusDetailEnumReason) IsKnown() bool {
	switch r {
	case StatusDetailOfLinkedBankAccountStatusDetailEnumReasonUnverified, StatusDetailOfLinkedBankAccountStatusDetailEnumReasonInReview, StatusDetailOfLinkedBankAccountStatusDetailEnumReasonPending, StatusDetailOfLinkedBankAccountStatusDetailEnumReasonStuck, StatusDetailOfLinkedBankAccountStatusDetailEnumReasonVerified, StatusDetailOfLinkedBankAccountStatusDetailEnumReasonFailedVerification, StatusDetailOfLinkedBankAccountStatusDetailEnumReasonDisabled, StatusDetailOfLinkedBankAccountStatusDetailEnumReasonNew:
		return true
	}
	return false
}

// Identifies the origin of the status change (e.g., `watchtower`). This helps in
// tracking the cause of status updates.
type StatusDetailOfLinkedBankAccountStatusDetailEnumSource string

const (
	StatusDetailOfLinkedBankAccountStatusDetailEnumSourceWatchtower StatusDetailOfLinkedBankAccountStatusDetailEnumSource = "watchtower"
)

func (r StatusDetailOfLinkedBankAccountStatusDetailEnumSource) IsKnown() bool {
	switch r {
	case StatusDetailOfLinkedBankAccountStatusDetailEnumSourceWatchtower:
		return true
	}
	return false
}

type StatusDetailsV1 struct {
	// A human-readable description of the current status.
	Message string `json:"message,required"`
	// A machine-readable identifier for the specific status, useful for programmatic
	// handling.
	Reason string `json:"reason,required"`
	// Identifies the origin of the status change (e.g., `bank_decline`, `watchtower`).
	// This helps in tracking the cause of status updates.
	Source string              `json:"source,required"`
	JSON   statusDetailsV1JSON `json:"-"`
}

// statusDetailsV1JSON contains the JSON metadata for the struct [StatusDetailsV1]
type statusDetailsV1JSON struct {
	Message     apijson.Field
	Reason      apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StatusDetailsV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r statusDetailsV1JSON) RawJSON() string {
	return r.raw
}

type StatusHistoryV1 struct {
	// The time the status change occurred.
	ChangedAt time.Time `json:"changed_at,required" format:"date-time"`
	// A human-readable description of the status.
	Message string `json:"message,required"`
	// A machine-readable identifier for the specific status, useful for programmatic
	// handling.
	Reason StatusReasonV1 `json:"reason,required"`
	// Identifies the origin of the status change (e.g., `bank_decline`, `watchtower`).
	// This helps in tracking the cause of status updates.
	Source StatusSourceV1 `json:"source,required"`
	// The current status of the `charge` or `payout`.
	Status PaymentStatusV1 `json:"status,required"`
	// The status code if applicable.
	Code string              `json:"code,nullable"`
	JSON statusHistoryV1JSON `json:"-"`
}

// statusHistoryV1JSON contains the JSON metadata for the struct [StatusHistoryV1]
type statusHistoryV1JSON struct {
	ChangedAt   apijson.Field
	Message     apijson.Field
	Reason      apijson.Field
	Source      apijson.Field
	Status      apijson.Field
	Code        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StatusHistoryV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r statusHistoryV1JSON) RawJSON() string {
	return r.raw
}

type StatusReasonV1 string

const (
	StatusReasonV1InsufficientFunds   StatusReasonV1 = "insufficient_funds"
	StatusReasonV1ClosedBankAccount   StatusReasonV1 = "closed_bank_account"
	StatusReasonV1InvalidBankAccount  StatusReasonV1 = "invalid_bank_account"
	StatusReasonV1InvalidRouting      StatusReasonV1 = "invalid_routing"
	StatusReasonV1Disputed            StatusReasonV1 = "disputed"
	StatusReasonV1PaymentStopped      StatusReasonV1 = "payment_stopped"
	StatusReasonV1OwnerDeceased       StatusReasonV1 = "owner_deceased"
	StatusReasonV1FrozenBankAccount   StatusReasonV1 = "frozen_bank_account"
	StatusReasonV1RiskReview          StatusReasonV1 = "risk_review"
	StatusReasonV1Fraudulent          StatusReasonV1 = "fraudulent"
	StatusReasonV1DuplicateEntry      StatusReasonV1 = "duplicate_entry"
	StatusReasonV1InvalidPaykey       StatusReasonV1 = "invalid_paykey"
	StatusReasonV1PaymentBlocked      StatusReasonV1 = "payment_blocked"
	StatusReasonV1AmountTooLarge      StatusReasonV1 = "amount_too_large"
	StatusReasonV1TooManyAttempts     StatusReasonV1 = "too_many_attempts"
	StatusReasonV1InternalSystemError StatusReasonV1 = "internal_system_error"
	StatusReasonV1UserRequest         StatusReasonV1 = "user_request"
	StatusReasonV1Ok                  StatusReasonV1 = "ok"
	StatusReasonV1OtherNetworkReturn  StatusReasonV1 = "other_network_return"
	StatusReasonV1PayoutRefused       StatusReasonV1 = "payout_refused"
)

func (r StatusReasonV1) IsKnown() bool {
	switch r {
	case StatusReasonV1InsufficientFunds, StatusReasonV1ClosedBankAccount, StatusReasonV1InvalidBankAccount, StatusReasonV1InvalidRouting, StatusReasonV1Disputed, StatusReasonV1PaymentStopped, StatusReasonV1OwnerDeceased, StatusReasonV1FrozenBankAccount, StatusReasonV1RiskReview, StatusReasonV1Fraudulent, StatusReasonV1DuplicateEntry, StatusReasonV1InvalidPaykey, StatusReasonV1PaymentBlocked, StatusReasonV1AmountTooLarge, StatusReasonV1TooManyAttempts, StatusReasonV1InternalSystemError, StatusReasonV1UserRequest, StatusReasonV1Ok, StatusReasonV1OtherNetworkReturn, StatusReasonV1PayoutRefused:
		return true
	}
	return false
}

type StatusSourceV1 string

const (
	StatusSourceV1Watchtower      StatusSourceV1 = "watchtower"
	StatusSourceV1BankDecline     StatusSourceV1 = "bank_decline"
	StatusSourceV1CustomerDispute StatusSourceV1 = "customer_dispute"
	StatusSourceV1UserAction      StatusSourceV1 = "user_action"
	StatusSourceV1System          StatusSourceV1 = "system"
)

func (r StatusSourceV1) IsKnown() bool {
	switch r {
	case StatusSourceV1Watchtower, StatusSourceV1BankDecline, StatusSourceV1CustomerDispute, StatusSourceV1UserAction, StatusSourceV1System:
		return true
	}
	return false
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

// Describes the direction of the funding event from the perspective of the
// `linked_bank_account`.
type TransferDirectionV1 string

const (
	TransferDirectionV1Deposit    TransferDirectionV1 = "deposit"
	TransferDirectionV1Withdrawal TransferDirectionV1 = "withdrawal"
)

func (r TransferDirectionV1) IsKnown() bool {
	switch r {
	case TransferDirectionV1Deposit, TransferDirectionV1Withdrawal:
		return true
	}
	return false
}

type UpdateChargeStatusV1RequestParam struct {
	// Details about why the charge status was updated.
	Reason param.Field[string] `json:"reason"`
}

func (r UpdateChargeStatusV1RequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UpdatePayoutStatusV1RequestParam struct {
	// Details about why the payout status was updated.
	Reason param.Field[string] `json:"reason,required"`
}

func (r UpdatePayoutStatusV1RequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
