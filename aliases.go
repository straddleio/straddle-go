// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package straddle

import (
	"github.com/stainless-sdks/straddle-go/internal/apierror"
	"github.com/stainless-sdks/straddle-go/shared"
)

type Error = apierror.Error

// This is an alias to an internal type.
type AccountTypeV1 = shared.AccountTypeV1

// This is an alias to an internal value.
const AccountTypeV1Checking = shared.AccountTypeV1Checking

// This is an alias to an internal value.
const AccountTypeV1Savings = shared.AccountTypeV1Savings

// This is an alias to an internal type.
type AccountV1 = shared.AccountV1

// The access level granted to the account. This is determined by your platform
// configuration. Use `standard` unless instructed otherwise by Straddle.
//
// This is an alias to an internal type.
type AccountV1AccessLevel = shared.AccountV1AccessLevel

// This is an alias to an internal value.
const AccountV1AccessLevelStandard = shared.AccountV1AccessLevelStandard

// This is an alias to an internal value.
const AccountV1AccessLevelManaged = shared.AccountV1AccessLevelManaged

// The current status of the account (e.g., 'active', 'inactive', 'pending').
//
// This is an alias to an internal type.
type AccountV1Status = shared.AccountV1Status

// This is an alias to an internal value.
const AccountV1StatusCreated = shared.AccountV1StatusCreated

// This is an alias to an internal value.
const AccountV1StatusOnboarding = shared.AccountV1StatusOnboarding

// This is an alias to an internal value.
const AccountV1StatusActive = shared.AccountV1StatusActive

// This is an alias to an internal value.
const AccountV1StatusRejected = shared.AccountV1StatusRejected

// This is an alias to an internal value.
const AccountV1StatusInactive = shared.AccountV1StatusInactive

// This is an alias to an internal type.
type AccountV1StatusDetail = shared.AccountV1StatusDetail

// A machine-readable identifier for the specific status, useful for programmatic
// handling.
//
// This is an alias to an internal type.
type AccountV1StatusDetailReason = shared.AccountV1StatusDetailReason

// This is an alias to an internal value.
const AccountV1StatusDetailReasonUnverified = shared.AccountV1StatusDetailReasonUnverified

// This is an alias to an internal value.
const AccountV1StatusDetailReasonInReview = shared.AccountV1StatusDetailReasonInReview

// This is an alias to an internal value.
const AccountV1StatusDetailReasonPending = shared.AccountV1StatusDetailReasonPending

// This is an alias to an internal value.
const AccountV1StatusDetailReasonStuck = shared.AccountV1StatusDetailReasonStuck

// This is an alias to an internal value.
const AccountV1StatusDetailReasonVerified = shared.AccountV1StatusDetailReasonVerified

// This is an alias to an internal value.
const AccountV1StatusDetailReasonFailedVerification = shared.AccountV1StatusDetailReasonFailedVerification

// This is an alias to an internal value.
const AccountV1StatusDetailReasonDisabled = shared.AccountV1StatusDetailReasonDisabled

// This is an alias to an internal value.
const AccountV1StatusDetailReasonTerminated = shared.AccountV1StatusDetailReasonTerminated

// This is an alias to an internal value.
const AccountV1StatusDetailReasonNew = shared.AccountV1StatusDetailReasonNew

// Identifies the origin of the status change (e.g., `bank_decline`, `watchtower`).
// This helps in tracking the cause of status updates.
//
// This is an alias to an internal type.
type AccountV1StatusDetailSource = shared.AccountV1StatusDetailSource

// This is an alias to an internal value.
const AccountV1StatusDetailSourceWatchtower = shared.AccountV1StatusDetailSourceWatchtower

// The type of account (e.g., 'individual', 'business').
//
// This is an alias to an internal type.
type AccountV1Type = shared.AccountV1Type

// This is an alias to an internal value.
const AccountV1TypeBusiness = shared.AccountV1TypeBusiness

// This is an alias to an internal type.
type AccountV1Capabilities = shared.AccountV1Capabilities

// This is an alias to an internal type.
type AccountV1CapabilitiesConsentTypes = shared.AccountV1CapabilitiesConsentTypes

// This is an alias to an internal type.
type AccountV1CapabilitiesCustomerTypes = shared.AccountV1CapabilitiesCustomerTypes

// This is an alias to an internal type.
type AccountV1CapabilitiesPaymentTypes = shared.AccountV1CapabilitiesPaymentTypes

// This is an alias to an internal type.
type AccountV1Settings = shared.AccountV1Settings

// This is an alias to an internal type.
type AccountV1SettingsCharges = shared.AccountV1SettingsCharges

// The amount of time it takes for a charge to be funded. This value is defined by
// Straddle.
//
// This is an alias to an internal type.
type AccountV1SettingsChargesFundingTime = shared.AccountV1SettingsChargesFundingTime

// This is an alias to an internal value.
const AccountV1SettingsChargesFundingTimeImmediate = shared.AccountV1SettingsChargesFundingTimeImmediate

// This is an alias to an internal value.
const AccountV1SettingsChargesFundingTimeNextDay = shared.AccountV1SettingsChargesFundingTimeNextDay

// This is an alias to an internal value.
const AccountV1SettingsChargesFundingTimeOneDay = shared.AccountV1SettingsChargesFundingTimeOneDay

// This is an alias to an internal value.
const AccountV1SettingsChargesFundingTimeTwoDay = shared.AccountV1SettingsChargesFundingTimeTwoDay

// This is an alias to an internal value.
const AccountV1SettingsChargesFundingTimeThreeDay = shared.AccountV1SettingsChargesFundingTimeThreeDay

// This is an alias to an internal type.
type AccountV1SettingsPayouts = shared.AccountV1SettingsPayouts

// The amount of time it takes for a payout to be funded. This value is defined by
// Straddle.
//
// This is an alias to an internal type.
type AccountV1SettingsPayoutsFundingTime = shared.AccountV1SettingsPayoutsFundingTime

// This is an alias to an internal value.
const AccountV1SettingsPayoutsFundingTimeImmediate = shared.AccountV1SettingsPayoutsFundingTimeImmediate

// This is an alias to an internal value.
const AccountV1SettingsPayoutsFundingTimeNextDay = shared.AccountV1SettingsPayoutsFundingTimeNextDay

// This is an alias to an internal value.
const AccountV1SettingsPayoutsFundingTimeOneDay = shared.AccountV1SettingsPayoutsFundingTimeOneDay

// This is an alias to an internal value.
const AccountV1SettingsPayoutsFundingTimeTwoDay = shared.AccountV1SettingsPayoutsFundingTimeTwoDay

// This is an alias to an internal value.
const AccountV1SettingsPayoutsFundingTimeThreeDay = shared.AccountV1SettingsPayoutsFundingTimeThreeDay

// This is an alias to an internal type.
type AddressV11 = shared.AddressV11

// This is an alias to an internal type.
type AddressV11Param = shared.AddressV11Param

// This is an alias to an internal type.
type BankAccountV1RequestParam = shared.BankAccountV1RequestParam

// This is an alias to an internal type.
type BusinessProfileV1 = shared.BusinessProfileV1

// This is an alias to an internal type.
type BusinessProfileV1Param = shared.BusinessProfileV1Param

// This is an alias to an internal type.
type Capability = shared.Capability

// This is an alias to an internal type.
type CapabilityCapabilityStatus = shared.CapabilityCapabilityStatus

// This is an alias to an internal value.
const CapabilityCapabilityStatusActive = shared.CapabilityCapabilityStatusActive

// This is an alias to an internal value.
const CapabilityCapabilityStatusInactive = shared.CapabilityCapabilityStatusInactive

// This is an alias to an internal type.
type ChargeConfigurationV1 = shared.ChargeConfigurationV1

// Defines whether to check the customer's balance before processing the charge.
//
// This is an alias to an internal type.
type ChargeConfigurationV1BalanceCheck = shared.ChargeConfigurationV1BalanceCheck

// This is an alias to an internal value.
const ChargeConfigurationV1BalanceCheckRequired = shared.ChargeConfigurationV1BalanceCheckRequired

// This is an alias to an internal value.
const ChargeConfigurationV1BalanceCheckEnabled = shared.ChargeConfigurationV1BalanceCheckEnabled

// This is an alias to an internal value.
const ChargeConfigurationV1BalanceCheckDisabled = shared.ChargeConfigurationV1BalanceCheckDisabled

// This is an alias to an internal type.
type ChargeConfigurationV1Param = shared.ChargeConfigurationV1Param

// This is an alias to an internal type.
type ChargeV1ItemResponse = shared.ChargeV1ItemResponse

// This is an alias to an internal type.
type ChargeV1ItemResponseData = shared.ChargeV1ItemResponseData

// Compliance profile for individual customers
//
// This is an alias to an internal type.
type ComplianceProfileUnmaskedV1 = shared.ComplianceProfileUnmaskedV1

// Compliance profile for individual customers
//
// This is an alias to an internal type.
type ComplianceProfileUnmaskedV1IndividualComplianceProfile = shared.ComplianceProfileUnmaskedV1IndividualComplianceProfile

// Compliance profile for business customers
//
// This is an alias to an internal type.
type ComplianceProfileUnmaskedV1BusinessComplianceProfile = shared.ComplianceProfileUnmaskedV1BusinessComplianceProfile

// Compliance profile for individual customers
//
// This is an alias to an internal type.
type ComplianceProfileUnmaskedV1UnionParam = shared.ComplianceProfileUnmaskedV1UnionParam

// Compliance profile for individual customers
//
// This is an alias to an internal type.
type ComplianceProfileUnmaskedV1IndividualComplianceProfileParam = shared.ComplianceProfileUnmaskedV1IndividualComplianceProfileParam

// Compliance profile for business customers
//
// This is an alias to an internal type.
type ComplianceProfileUnmaskedV1BusinessComplianceProfileParam = shared.ComplianceProfileUnmaskedV1BusinessComplianceProfileParam

// The channel or mechanism through which the payment was authorized. Use
// `internet` for payments made online or through a mobile app and `signed` for
// signed agreements where there is a consent form or contract. Use `signed` for
// PDF signatures.
//
// This is an alias to an internal type.
type ConsentTypeV1 = shared.ConsentTypeV1

// This is an alias to an internal value.
const ConsentTypeV1Internet = shared.ConsentTypeV1Internet

// This is an alias to an internal value.
const ConsentTypeV1Signed = shared.ConsentTypeV1Signed

// This is an alias to an internal type.
type CustomerStatusV1 = shared.CustomerStatusV1

// This is an alias to an internal value.
const CustomerStatusV1Pending = shared.CustomerStatusV1Pending

// This is an alias to an internal value.
const CustomerStatusV1Review = shared.CustomerStatusV1Review

// This is an alias to an internal value.
const CustomerStatusV1Verified = shared.CustomerStatusV1Verified

// This is an alias to an internal value.
const CustomerStatusV1Inactive = shared.CustomerStatusV1Inactive

// This is an alias to an internal value.
const CustomerStatusV1Rejected = shared.CustomerStatusV1Rejected

// This is an alias to an internal type.
type CustomerTypeV1 = shared.CustomerTypeV1

// This is an alias to an internal value.
const CustomerTypeV1Individual = shared.CustomerTypeV1Individual

// This is an alias to an internal value.
const CustomerTypeV1Business = shared.CustomerTypeV1Business

// This is an alias to an internal type.
type CustomerV1 = shared.CustomerV1

// Compliance profile for individual customers
//
// This is an alias to an internal type.
type CustomerV1ComplianceProfile = shared.CustomerV1ComplianceProfile

// Compliance profile for individual customers
//
// This is an alias to an internal type.
type CustomerV1ComplianceProfileIndividualComplianceProfile = shared.CustomerV1ComplianceProfileIndividualComplianceProfile

// Compliance profile for business customers
//
// This is an alias to an internal type.
type CustomerV1ComplianceProfileBusinessComplianceProfile = shared.CustomerV1ComplianceProfileBusinessComplianceProfile

// This is an alias to an internal type.
type CustomerV1Device = shared.CustomerV1Device

// This is an alias to an internal type.
type CustomerV1ItemResponse = shared.CustomerV1ItemResponse

// Information about the customer associated with the charge or payout.
//
// This is an alias to an internal type.
type CustomerDetailsV1 = shared.CustomerDetailsV1

// The type of customer.
//
// This is an alias to an internal type.
type CustomerDetailsV1CustomerType = shared.CustomerDetailsV1CustomerType

// This is an alias to an internal value.
const CustomerDetailsV1CustomerTypeIndividual = shared.CustomerDetailsV1CustomerTypeIndividual

// This is an alias to an internal value.
const CustomerDetailsV1CustomerTypeBusiness = shared.CustomerDetailsV1CustomerTypeBusiness

// This is an alias to an internal type.
type DeviceUnmaskedV1 = shared.DeviceUnmaskedV1

// This is an alias to an internal type.
type DeviceUnmaskedV1Param = shared.DeviceUnmaskedV1Param

// This is an alias to an internal type.
type DeviceInfoV1 = shared.DeviceInfoV1

// This is an alias to an internal type.
type DeviceInfoV1Param = shared.DeviceInfoV1Param

// This is an alias to an internal type.
type FundingEventSummaryV1 = shared.FundingEventSummaryV1

// The funding event types describes the direction and reason for the funding
// event.
//
// This is an alias to an internal type.
type FundingEventTypeV1 = shared.FundingEventTypeV1

// This is an alias to an internal value.
const FundingEventTypeV1ChargeDeposit = shared.FundingEventTypeV1ChargeDeposit

// This is an alias to an internal value.
const FundingEventTypeV1ChargeReversal = shared.FundingEventTypeV1ChargeReversal

// This is an alias to an internal value.
const FundingEventTypeV1PayoutReturn = shared.FundingEventTypeV1PayoutReturn

// This is an alias to an internal value.
const FundingEventTypeV1PayoutWithdrawal = shared.FundingEventTypeV1PayoutWithdrawal

// This is an alias to an internal type.
type IdentityDecisionV1 = shared.IdentityDecisionV1

// This is an alias to an internal value.
const IdentityDecisionV1Accept = shared.IdentityDecisionV1Accept

// This is an alias to an internal value.
const IdentityDecisionV1Reject = shared.IdentityDecisionV1Reject

// This is an alias to an internal value.
const IdentityDecisionV1Review = shared.IdentityDecisionV1Review

// This is an alias to an internal type.
type IdentityVerificationBreakdownV1 = shared.IdentityVerificationBreakdownV1

// This is an alias to an internal type.
type ItemResponseOfAccountV1 = shared.ItemResponseOfAccountV1

// Indicates the structure of the returned content.
//
//   - "object" means the `data` field contains a single JSON object.
//   - "array" means the `data` field contains an array of objects.
//   - "error" means the `data` field contains an error object with details of the
//     issue.
//   - "none" means no data is returned.
//
// This is an alias to an internal type.
type ItemResponseOfAccountV1ResponseType = shared.ItemResponseOfAccountV1ResponseType

// This is an alias to an internal value.
const ItemResponseOfAccountV1ResponseTypeObject = shared.ItemResponseOfAccountV1ResponseTypeObject

// This is an alias to an internal value.
const ItemResponseOfAccountV1ResponseTypeArray = shared.ItemResponseOfAccountV1ResponseTypeArray

// This is an alias to an internal value.
const ItemResponseOfAccountV1ResponseTypeError = shared.ItemResponseOfAccountV1ResponseTypeError

// This is an alias to an internal value.
const ItemResponseOfAccountV1ResponseTypeNone = shared.ItemResponseOfAccountV1ResponseTypeNone

// This is an alias to an internal type.
type ItemResponseOfLinkedBankAccountV1 = shared.ItemResponseOfLinkedBankAccountV1

// Indicates the structure of the returned content.
//
//   - "object" means the `data` field contains a single JSON object.
//   - "array" means the `data` field contains an array of objects.
//   - "error" means the `data` field contains an error object with details of the
//     issue.
//   - "none" means no data is returned.
//
// This is an alias to an internal type.
type ItemResponseOfLinkedBankAccountV1ResponseType = shared.ItemResponseOfLinkedBankAccountV1ResponseType

// This is an alias to an internal value.
const ItemResponseOfLinkedBankAccountV1ResponseTypeObject = shared.ItemResponseOfLinkedBankAccountV1ResponseTypeObject

// This is an alias to an internal value.
const ItemResponseOfLinkedBankAccountV1ResponseTypeArray = shared.ItemResponseOfLinkedBankAccountV1ResponseTypeArray

// This is an alias to an internal value.
const ItemResponseOfLinkedBankAccountV1ResponseTypeError = shared.ItemResponseOfLinkedBankAccountV1ResponseTypeError

// This is an alias to an internal value.
const ItemResponseOfLinkedBankAccountV1ResponseTypeNone = shared.ItemResponseOfLinkedBankAccountV1ResponseTypeNone

// This is an alias to an internal type.
type ItemResponseOfOrganizationV1 = shared.ItemResponseOfOrganizationV1

// Indicates the structure of the returned content.
//
//   - "object" means the `data` field contains a single JSON object.
//   - "array" means the `data` field contains an array of objects.
//   - "error" means the `data` field contains an error object with details of the
//     issue.
//   - "none" means no data is returned.
//
// This is an alias to an internal type.
type ItemResponseOfOrganizationV1ResponseType = shared.ItemResponseOfOrganizationV1ResponseType

// This is an alias to an internal value.
const ItemResponseOfOrganizationV1ResponseTypeObject = shared.ItemResponseOfOrganizationV1ResponseTypeObject

// This is an alias to an internal value.
const ItemResponseOfOrganizationV1ResponseTypeArray = shared.ItemResponseOfOrganizationV1ResponseTypeArray

// This is an alias to an internal value.
const ItemResponseOfOrganizationV1ResponseTypeError = shared.ItemResponseOfOrganizationV1ResponseTypeError

// This is an alias to an internal value.
const ItemResponseOfOrganizationV1ResponseTypeNone = shared.ItemResponseOfOrganizationV1ResponseTypeNone

// This is an alias to an internal type.
type ItemResponseOfRepresentativeV1 = shared.ItemResponseOfRepresentativeV1

// Indicates the structure of the returned content.
//
//   - "object" means the `data` field contains a single JSON object.
//   - "array" means the `data` field contains an array of objects.
//   - "error" means the `data` field contains an error object with details of the
//     issue.
//   - "none" means no data is returned.
//
// This is an alias to an internal type.
type ItemResponseOfRepresentativeV1ResponseType = shared.ItemResponseOfRepresentativeV1ResponseType

// This is an alias to an internal value.
const ItemResponseOfRepresentativeV1ResponseTypeObject = shared.ItemResponseOfRepresentativeV1ResponseTypeObject

// This is an alias to an internal value.
const ItemResponseOfRepresentativeV1ResponseTypeArray = shared.ItemResponseOfRepresentativeV1ResponseTypeArray

// This is an alias to an internal value.
const ItemResponseOfRepresentativeV1ResponseTypeError = shared.ItemResponseOfRepresentativeV1ResponseTypeError

// This is an alias to an internal value.
const ItemResponseOfRepresentativeV1ResponseTypeNone = shared.ItemResponseOfRepresentativeV1ResponseTypeNone

// This is an alias to an internal type.
type LinkedBankAccountV1 = shared.LinkedBankAccountV1

// This is an alias to an internal type.
type LinkedBankAccountV1BankAccount = shared.LinkedBankAccountV1BankAccount

// The current status of the linked bank account.
//
// This is an alias to an internal type.
type LinkedBankAccountV1Status = shared.LinkedBankAccountV1Status

// This is an alias to an internal value.
const LinkedBankAccountV1StatusCreated = shared.LinkedBankAccountV1StatusCreated

// This is an alias to an internal value.
const LinkedBankAccountV1StatusOnboarding = shared.LinkedBankAccountV1StatusOnboarding

// This is an alias to an internal value.
const LinkedBankAccountV1StatusActive = shared.LinkedBankAccountV1StatusActive

// This is an alias to an internal value.
const LinkedBankAccountV1StatusRejected = shared.LinkedBankAccountV1StatusRejected

// This is an alias to an internal value.
const LinkedBankAccountV1StatusInactive = shared.LinkedBankAccountV1StatusInactive

// This is an alias to an internal type.
type OrganizationV1 = shared.OrganizationV1

// This is an alias to an internal type.
type PagedResponseMetadata1 = shared.PagedResponseMetadata1

// This is an alias to an internal type.
type PagedResponseMetadata2 = shared.PagedResponseMetadata2

// This is an alias to an internal type.
type PagedResponseOfCapabilityRequestV1 = shared.PagedResponseOfCapabilityRequestV1

// This is an alias to an internal type.
type PagedResponseOfCapabilityRequestV1Data = shared.PagedResponseOfCapabilityRequestV1Data

// The category of the requested capability. Use `payment_type` for charges and
// payouts, `customer_type` to define `individuals` or `businesses`, and
// `consent_type` for `signed_agreement` or `internet` payment authorization.
//
// This is an alias to an internal type.
type PagedResponseOfCapabilityRequestV1DataCategory = shared.PagedResponseOfCapabilityRequestV1DataCategory

// This is an alias to an internal value.
const PagedResponseOfCapabilityRequestV1DataCategoryPaymentType = shared.PagedResponseOfCapabilityRequestV1DataCategoryPaymentType

// This is an alias to an internal value.
const PagedResponseOfCapabilityRequestV1DataCategoryCustomerType = shared.PagedResponseOfCapabilityRequestV1DataCategoryCustomerType

// This is an alias to an internal value.
const PagedResponseOfCapabilityRequestV1DataCategoryConsentType = shared.PagedResponseOfCapabilityRequestV1DataCategoryConsentType

// The current status of the capability request.
//
// This is an alias to an internal type.
type PagedResponseOfCapabilityRequestV1DataStatus = shared.PagedResponseOfCapabilityRequestV1DataStatus

// This is an alias to an internal value.
const PagedResponseOfCapabilityRequestV1DataStatusActive = shared.PagedResponseOfCapabilityRequestV1DataStatusActive

// This is an alias to an internal value.
const PagedResponseOfCapabilityRequestV1DataStatusInactive = shared.PagedResponseOfCapabilityRequestV1DataStatusInactive

// This is an alias to an internal value.
const PagedResponseOfCapabilityRequestV1DataStatusInReview = shared.PagedResponseOfCapabilityRequestV1DataStatusInReview

// This is an alias to an internal value.
const PagedResponseOfCapabilityRequestV1DataStatusRejected = shared.PagedResponseOfCapabilityRequestV1DataStatusRejected

// This is an alias to an internal value.
const PagedResponseOfCapabilityRequestV1DataStatusApproved = shared.PagedResponseOfCapabilityRequestV1DataStatusApproved

// This is an alias to an internal value.
const PagedResponseOfCapabilityRequestV1DataStatusReviewing = shared.PagedResponseOfCapabilityRequestV1DataStatusReviewing

// The specific type of capability being requested within the category.
//
// This is an alias to an internal type.
type PagedResponseOfCapabilityRequestV1DataType = shared.PagedResponseOfCapabilityRequestV1DataType

// This is an alias to an internal value.
const PagedResponseOfCapabilityRequestV1DataTypeCharges = shared.PagedResponseOfCapabilityRequestV1DataTypeCharges

// This is an alias to an internal value.
const PagedResponseOfCapabilityRequestV1DataTypePayouts = shared.PagedResponseOfCapabilityRequestV1DataTypePayouts

// This is an alias to an internal value.
const PagedResponseOfCapabilityRequestV1DataTypeIndividuals = shared.PagedResponseOfCapabilityRequestV1DataTypeIndividuals

// This is an alias to an internal value.
const PagedResponseOfCapabilityRequestV1DataTypeBusinesses = shared.PagedResponseOfCapabilityRequestV1DataTypeBusinesses

// This is an alias to an internal value.
const PagedResponseOfCapabilityRequestV1DataTypeSignedAgreement = shared.PagedResponseOfCapabilityRequestV1DataTypeSignedAgreement

// This is an alias to an internal value.
const PagedResponseOfCapabilityRequestV1DataTypeInternet = shared.PagedResponseOfCapabilityRequestV1DataTypeInternet

// Indicates the structure of the returned content.
//
//   - "object" means the `data` field contains a single JSON object.
//   - "array" means the `data` field contains an array of objects.
//   - "error" means the `data` field contains an error object with details of the
//     issue.
//   - "none" means no data is returned.
//
// This is an alias to an internal type.
type PagedResponseOfCapabilityRequestV1ResponseType = shared.PagedResponseOfCapabilityRequestV1ResponseType

// This is an alias to an internal value.
const PagedResponseOfCapabilityRequestV1ResponseTypeObject = shared.PagedResponseOfCapabilityRequestV1ResponseTypeObject

// This is an alias to an internal value.
const PagedResponseOfCapabilityRequestV1ResponseTypeArray = shared.PagedResponseOfCapabilityRequestV1ResponseTypeArray

// This is an alias to an internal value.
const PagedResponseOfCapabilityRequestV1ResponseTypeError = shared.PagedResponseOfCapabilityRequestV1ResponseTypeError

// This is an alias to an internal value.
const PagedResponseOfCapabilityRequestV1ResponseTypeNone = shared.PagedResponseOfCapabilityRequestV1ResponseTypeNone

// Metadata about the API request, including an identifier, timestamp, and
// pagination details.
//
// This is an alias to an internal type.
type PagedResponseMetadata = shared.PagedResponseMetadata

// The order that the results were sorted by.
//
// This is an alias to an internal type.
type PagedResponseMetadataSortOrder = shared.PagedResponseMetadataSortOrder

// This is an alias to an internal value.
const PagedResponseMetadataSortOrderAsc = shared.PagedResponseMetadataSortOrderAsc

// This is an alias to an internal value.
const PagedResponseMetadataSortOrderDesc = shared.PagedResponseMetadataSortOrderDesc

// This is an alias to an internal type.
type PaykeyBankDetailsV1 = shared.PaykeyBankDetailsV1

// This is an alias to an internal type.
type PaykeySourceV1 = shared.PaykeySourceV1

// This is an alias to an internal value.
const PaykeySourceV1BankAccount = shared.PaykeySourceV1BankAccount

// This is an alias to an internal value.
const PaykeySourceV1Straddle = shared.PaykeySourceV1Straddle

// This is an alias to an internal value.
const PaykeySourceV1Mx = shared.PaykeySourceV1Mx

// This is an alias to an internal value.
const PaykeySourceV1Plaid = shared.PaykeySourceV1Plaid

// This is an alias to an internal type.
type PaykeyStatusV1 = shared.PaykeyStatusV1

// This is an alias to an internal value.
const PaykeyStatusV1Pending = shared.PaykeyStatusV1Pending

// This is an alias to an internal value.
const PaykeyStatusV1Active = shared.PaykeyStatusV1Active

// This is an alias to an internal value.
const PaykeyStatusV1Inactive = shared.PaykeyStatusV1Inactive

// This is an alias to an internal value.
const PaykeyStatusV1Rejected = shared.PaykeyStatusV1Rejected

// This is an alias to an internal type.
type PaykeyV1ItemResponse = shared.PaykeyV1ItemResponse

// This is an alias to an internal type.
type PaykeyV1ItemResponseData = shared.PaykeyV1ItemResponseData

// This is an alias to an internal type.
type PaykeyDetailsV1 = shared.PaykeyDetailsV1

// The payment rail used for the charge or payout.
//
// This is an alias to an internal type.
type PaymentRailV1 = shared.PaymentRailV1

// This is an alias to an internal value.
const PaymentRailV1ACH = shared.PaymentRailV1ACH

// The field to sort the results by.
//
// This is an alias to an internal type.
type PaymentSortByV1 = shared.PaymentSortByV1

// This is an alias to an internal value.
const PaymentSortByV1CreatedAt = shared.PaymentSortByV1CreatedAt

// This is an alias to an internal value.
const PaymentSortByV1PaymentDate = shared.PaymentSortByV1PaymentDate

// This is an alias to an internal value.
const PaymentSortByV1EffectiveAt = shared.PaymentSortByV1EffectiveAt

// This is an alias to an internal value.
const PaymentSortByV1ID = shared.PaymentSortByV1ID

// This is an alias to an internal value.
const PaymentSortByV1Amount = shared.PaymentSortByV1Amount

// The current status of the `charge` or `payout`.
//
// This is an alias to an internal type.
type PaymentStatusV1 = shared.PaymentStatusV1

// This is an alias to an internal value.
const PaymentStatusV1Created = shared.PaymentStatusV1Created

// This is an alias to an internal value.
const PaymentStatusV1Scheduled = shared.PaymentStatusV1Scheduled

// This is an alias to an internal value.
const PaymentStatusV1Failed = shared.PaymentStatusV1Failed

// This is an alias to an internal value.
const PaymentStatusV1Cancelled = shared.PaymentStatusV1Cancelled

// This is an alias to an internal value.
const PaymentStatusV1OnHold = shared.PaymentStatusV1OnHold

// This is an alias to an internal value.
const PaymentStatusV1Pending = shared.PaymentStatusV1Pending

// This is an alias to an internal value.
const PaymentStatusV1Paid = shared.PaymentStatusV1Paid

// This is an alias to an internal value.
const PaymentStatusV1Reversed = shared.PaymentStatusV1Reversed

// The type of payment.
//
// This is an alias to an internal type.
type PaymentTypeV1 = shared.PaymentTypeV1

// This is an alias to an internal value.
const PaymentTypeV1Charge = shared.PaymentTypeV1Charge

// This is an alias to an internal value.
const PaymentTypeV1Payout = shared.PaymentTypeV1Payout

// This is an alias to an internal type.
type PayoutV1ItemResponse = shared.PayoutV1ItemResponse

// This is an alias to an internal type.
type PayoutV1ItemResponseData = shared.PayoutV1ItemResponseData

// This is an alias to an internal type.
type RelationshipV1Param = shared.RelationshipV1Param

// This is an alias to an internal type.
type RepresentativeV1 = shared.RepresentativeV1

// This is an alias to an internal type.
type RepresentativeV1Relationship = shared.RepresentativeV1Relationship

// The current status of the representative.
//
// This is an alias to an internal type.
type RepresentativeV1Status = shared.RepresentativeV1Status

// This is an alias to an internal value.
const RepresentativeV1StatusCreated = shared.RepresentativeV1StatusCreated

// This is an alias to an internal value.
const RepresentativeV1StatusOnboarding = shared.RepresentativeV1StatusOnboarding

// This is an alias to an internal value.
const RepresentativeV1StatusActive = shared.RepresentativeV1StatusActive

// This is an alias to an internal value.
const RepresentativeV1StatusRejected = shared.RepresentativeV1StatusRejected

// This is an alias to an internal value.
const RepresentativeV1StatusInactive = shared.RepresentativeV1StatusInactive

// This is an alias to an internal type.
type RepresentativeV1StatusDetail = shared.RepresentativeV1StatusDetail

// A machine-readable identifier for the specific status, useful for programmatic
// handling.
//
// This is an alias to an internal type.
type RepresentativeV1StatusDetailReason = shared.RepresentativeV1StatusDetailReason

// This is an alias to an internal value.
const RepresentativeV1StatusDetailReasonUnverified = shared.RepresentativeV1StatusDetailReasonUnverified

// This is an alias to an internal value.
const RepresentativeV1StatusDetailReasonInReview = shared.RepresentativeV1StatusDetailReasonInReview

// This is an alias to an internal value.
const RepresentativeV1StatusDetailReasonPending = shared.RepresentativeV1StatusDetailReasonPending

// This is an alias to an internal value.
const RepresentativeV1StatusDetailReasonStuck = shared.RepresentativeV1StatusDetailReasonStuck

// This is an alias to an internal value.
const RepresentativeV1StatusDetailReasonVerified = shared.RepresentativeV1StatusDetailReasonVerified

// This is an alias to an internal value.
const RepresentativeV1StatusDetailReasonFailedVerification = shared.RepresentativeV1StatusDetailReasonFailedVerification

// This is an alias to an internal value.
const RepresentativeV1StatusDetailReasonDisabled = shared.RepresentativeV1StatusDetailReasonDisabled

// This is an alias to an internal value.
const RepresentativeV1StatusDetailReasonNew = shared.RepresentativeV1StatusDetailReasonNew

// Identifies the origin of the status change (e.g., `watchtower`). This helps in
// tracking the cause of status updates.
//
// This is an alias to an internal type.
type RepresentativeV1StatusDetailSource = shared.RepresentativeV1StatusDetailSource

// This is an alias to an internal value.
const RepresentativeV1StatusDetailSourceWatchtower = shared.RepresentativeV1StatusDetailSourceWatchtower

// Indicates the structure of the returned content.
//
//   - "object" means the `data` field contains a single JSON object.
//   - "array" means the `data` field contains an array of objects.
//   - "error" means the `data` field contains an error object with details of the
//     issue.
//   - "none" means no data is returned.
//
// This is an alias to an internal type.
type ResponseTypeEnum = shared.ResponseTypeEnum

// This is an alias to an internal value.
const ResponseTypeEnumObject = shared.ResponseTypeEnumObject

// This is an alias to an internal value.
const ResponseTypeEnumArray = shared.ResponseTypeEnumArray

// This is an alias to an internal value.
const ResponseTypeEnumError = shared.ResponseTypeEnumError

// This is an alias to an internal value.
const ResponseTypeEnumNone = shared.ResponseTypeEnumNone

// Metadata about the API request, including an identifier and timestamp.
//
// This is an alias to an internal type.
type ResponseMetadata = shared.ResponseMetadata

// This is an alias to an internal type.
type SortOrder = shared.SortOrder

// This is an alias to an internal value.
const SortOrderAsc = shared.SortOrderAsc

// This is an alias to an internal value.
const SortOrderDesc = shared.SortOrderDesc

// This is an alias to an internal type.
type StatusDetailOfLinkedBankAccountStatusDetailEnum = shared.StatusDetailOfLinkedBankAccountStatusDetailEnum

// A machine-readable identifier for the specific status, useful for programmatic
// handling.
//
// This is an alias to an internal type.
type StatusDetailOfLinkedBankAccountStatusDetailEnumReason = shared.StatusDetailOfLinkedBankAccountStatusDetailEnumReason

// This is an alias to an internal value.
const StatusDetailOfLinkedBankAccountStatusDetailEnumReasonUnverified = shared.StatusDetailOfLinkedBankAccountStatusDetailEnumReasonUnverified

// This is an alias to an internal value.
const StatusDetailOfLinkedBankAccountStatusDetailEnumReasonInReview = shared.StatusDetailOfLinkedBankAccountStatusDetailEnumReasonInReview

// This is an alias to an internal value.
const StatusDetailOfLinkedBankAccountStatusDetailEnumReasonPending = shared.StatusDetailOfLinkedBankAccountStatusDetailEnumReasonPending

// This is an alias to an internal value.
const StatusDetailOfLinkedBankAccountStatusDetailEnumReasonStuck = shared.StatusDetailOfLinkedBankAccountStatusDetailEnumReasonStuck

// This is an alias to an internal value.
const StatusDetailOfLinkedBankAccountStatusDetailEnumReasonVerified = shared.StatusDetailOfLinkedBankAccountStatusDetailEnumReasonVerified

// This is an alias to an internal value.
const StatusDetailOfLinkedBankAccountStatusDetailEnumReasonFailedVerification = shared.StatusDetailOfLinkedBankAccountStatusDetailEnumReasonFailedVerification

// This is an alias to an internal value.
const StatusDetailOfLinkedBankAccountStatusDetailEnumReasonDisabled = shared.StatusDetailOfLinkedBankAccountStatusDetailEnumReasonDisabled

// This is an alias to an internal value.
const StatusDetailOfLinkedBankAccountStatusDetailEnumReasonNew = shared.StatusDetailOfLinkedBankAccountStatusDetailEnumReasonNew

// Identifies the origin of the status change (e.g., `watchtower`). This helps in
// tracking the cause of status updates.
//
// This is an alias to an internal type.
type StatusDetailOfLinkedBankAccountStatusDetailEnumSource = shared.StatusDetailOfLinkedBankAccountStatusDetailEnumSource

// This is an alias to an internal value.
const StatusDetailOfLinkedBankAccountStatusDetailEnumSourceWatchtower = shared.StatusDetailOfLinkedBankAccountStatusDetailEnumSourceWatchtower

// This is an alias to an internal type.
type StatusDetailsV1 = shared.StatusDetailsV1

// This is an alias to an internal type.
type StatusHistoryV1 = shared.StatusHistoryV1

// This is an alias to an internal type.
type StatusReasonV1 = shared.StatusReasonV1

// This is an alias to an internal value.
const StatusReasonV1InsufficientFunds = shared.StatusReasonV1InsufficientFunds

// This is an alias to an internal value.
const StatusReasonV1ClosedBankAccount = shared.StatusReasonV1ClosedBankAccount

// This is an alias to an internal value.
const StatusReasonV1InvalidBankAccount = shared.StatusReasonV1InvalidBankAccount

// This is an alias to an internal value.
const StatusReasonV1InvalidRouting = shared.StatusReasonV1InvalidRouting

// This is an alias to an internal value.
const StatusReasonV1Disputed = shared.StatusReasonV1Disputed

// This is an alias to an internal value.
const StatusReasonV1PaymentStopped = shared.StatusReasonV1PaymentStopped

// This is an alias to an internal value.
const StatusReasonV1OwnerDeceased = shared.StatusReasonV1OwnerDeceased

// This is an alias to an internal value.
const StatusReasonV1FrozenBankAccount = shared.StatusReasonV1FrozenBankAccount

// This is an alias to an internal value.
const StatusReasonV1RiskReview = shared.StatusReasonV1RiskReview

// This is an alias to an internal value.
const StatusReasonV1Fraudulent = shared.StatusReasonV1Fraudulent

// This is an alias to an internal value.
const StatusReasonV1DuplicateEntry = shared.StatusReasonV1DuplicateEntry

// This is an alias to an internal value.
const StatusReasonV1InvalidPaykey = shared.StatusReasonV1InvalidPaykey

// This is an alias to an internal value.
const StatusReasonV1PaymentBlocked = shared.StatusReasonV1PaymentBlocked

// This is an alias to an internal value.
const StatusReasonV1AmountTooLarge = shared.StatusReasonV1AmountTooLarge

// This is an alias to an internal value.
const StatusReasonV1TooManyAttempts = shared.StatusReasonV1TooManyAttempts

// This is an alias to an internal value.
const StatusReasonV1InternalSystemError = shared.StatusReasonV1InternalSystemError

// This is an alias to an internal value.
const StatusReasonV1UserRequest = shared.StatusReasonV1UserRequest

// This is an alias to an internal value.
const StatusReasonV1Ok = shared.StatusReasonV1Ok

// This is an alias to an internal value.
const StatusReasonV1OtherNetworkReturn = shared.StatusReasonV1OtherNetworkReturn

// This is an alias to an internal value.
const StatusReasonV1PayoutRefused = shared.StatusReasonV1PayoutRefused

// This is an alias to an internal type.
type StatusSourceV1 = shared.StatusSourceV1

// This is an alias to an internal value.
const StatusSourceV1Watchtower = shared.StatusSourceV1Watchtower

// This is an alias to an internal value.
const StatusSourceV1BankDecline = shared.StatusSourceV1BankDecline

// This is an alias to an internal value.
const StatusSourceV1CustomerDispute = shared.StatusSourceV1CustomerDispute

// This is an alias to an internal value.
const StatusSourceV1UserAction = shared.StatusSourceV1UserAction

// This is an alias to an internal value.
const StatusSourceV1System = shared.StatusSourceV1System

// This is an alias to an internal type.
type TermsOfServiceV1 = shared.TermsOfServiceV1

// The type or version of the agreement accepted. Use `embedded` unless your
// platform was specifically enabled for `direct` agreements.
//
// This is an alias to an internal type.
type TermsOfServiceV1AgreementType = shared.TermsOfServiceV1AgreementType

// This is an alias to an internal value.
const TermsOfServiceV1AgreementTypeEmbedded = shared.TermsOfServiceV1AgreementTypeEmbedded

// This is an alias to an internal value.
const TermsOfServiceV1AgreementTypeDirect = shared.TermsOfServiceV1AgreementTypeDirect

// This is an alias to an internal type.
type TermsOfServiceV1Param = shared.TermsOfServiceV1Param

// Describes the direction of the funding event from the perspective of the
// `linked_bank_account`.
//
// This is an alias to an internal type.
type TransferDirectionV1 = shared.TransferDirectionV1

// This is an alias to an internal value.
const TransferDirectionV1Deposit = shared.TransferDirectionV1Deposit

// This is an alias to an internal value.
const TransferDirectionV1Withdrawal = shared.TransferDirectionV1Withdrawal

// This is an alias to an internal type.
type UpdateChargeStatusV1RequestParam = shared.UpdateChargeStatusV1RequestParam

// This is an alias to an internal type.
type UpdatePayoutStatusV1RequestParam = shared.UpdatePayoutStatusV1RequestParam
