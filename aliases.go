// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package straddle

import (
	"github.com/straddleio/straddle-go/internal/apierror"
	"github.com/straddleio/straddle-go/packages/param"
	"github.com/straddleio/straddle-go/shared"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

type Error = apierror.Error

// Information about the customer associated with the charge or payout.
//
// This is an alias to an internal type.
type CustomerDetailsV1 = shared.CustomerDetailsV1

// The type of customer
//
// This is an alias to an internal type.
type CustomerDetailsV1CustomerType = shared.CustomerDetailsV1CustomerType

// Equals "individual"
const CustomerDetailsV1CustomerTypeIndividual = shared.CustomerDetailsV1CustomerTypeIndividual

// Equals "business"
const CustomerDetailsV1CustomerTypeBusiness = shared.CustomerDetailsV1CustomerTypeBusiness

// This is an alias to an internal type.
type DeviceInfoV1 = shared.DeviceInfoV1

// This is an alias to an internal type.
type DeviceInfoV1Param = shared.DeviceInfoV1Param

// Metadata about the API request, including an identifier, timestamp, and
// pagination details.
//
// This is an alias to an internal type.
type PagedResponseMetadata = shared.PagedResponseMetadata

// This is an alias to an internal type.
type PagedResponseMetadataSortOrder = shared.PagedResponseMetadataSortOrder

// Equals "asc"
const PagedResponseMetadataSortOrderAsc = shared.PagedResponseMetadataSortOrderAsc

// Equals "desc"
const PagedResponseMetadataSortOrderDesc = shared.PagedResponseMetadataSortOrderDesc

// This is an alias to an internal type.
type PaykeyDetailsV1 = shared.PaykeyDetailsV1

// Metadata about the API request, including an identifier and timestamp.
//
// This is an alias to an internal type.
type ResponseMetadata = shared.ResponseMetadata

// This is an alias to an internal type.
type StatusDetailsV1 = shared.StatusDetailsV1

// A machine-readable identifier for the specific status, useful for programmatic
// handling.
//
// This is an alias to an internal type.
type StatusDetailsV1Reason = shared.StatusDetailsV1Reason

// Equals "insufficient_funds"
const StatusDetailsV1ReasonInsufficientFunds = shared.StatusDetailsV1ReasonInsufficientFunds

// Equals "closed_bank_account"
const StatusDetailsV1ReasonClosedBankAccount = shared.StatusDetailsV1ReasonClosedBankAccount

// Equals "invalid_bank_account"
const StatusDetailsV1ReasonInvalidBankAccount = shared.StatusDetailsV1ReasonInvalidBankAccount

// Equals "invalid_routing"
const StatusDetailsV1ReasonInvalidRouting = shared.StatusDetailsV1ReasonInvalidRouting

// Equals "disputed"
const StatusDetailsV1ReasonDisputed = shared.StatusDetailsV1ReasonDisputed

// Equals "payment_stopped"
const StatusDetailsV1ReasonPaymentStopped = shared.StatusDetailsV1ReasonPaymentStopped

// Equals "owner_deceased"
const StatusDetailsV1ReasonOwnerDeceased = shared.StatusDetailsV1ReasonOwnerDeceased

// Equals "frozen_bank_account"
const StatusDetailsV1ReasonFrozenBankAccount = shared.StatusDetailsV1ReasonFrozenBankAccount

// Equals "risk_review"
const StatusDetailsV1ReasonRiskReview = shared.StatusDetailsV1ReasonRiskReview

// Equals "fraudulent"
const StatusDetailsV1ReasonFraudulent = shared.StatusDetailsV1ReasonFraudulent

// Equals "duplicate_entry"
const StatusDetailsV1ReasonDuplicateEntry = shared.StatusDetailsV1ReasonDuplicateEntry

// Equals "invalid_paykey"
const StatusDetailsV1ReasonInvalidPaykey = shared.StatusDetailsV1ReasonInvalidPaykey

// Equals "payment_blocked"
const StatusDetailsV1ReasonPaymentBlocked = shared.StatusDetailsV1ReasonPaymentBlocked

// Equals "amount_too_large"
const StatusDetailsV1ReasonAmountTooLarge = shared.StatusDetailsV1ReasonAmountTooLarge

// Equals "too_many_attempts"
const StatusDetailsV1ReasonTooManyAttempts = shared.StatusDetailsV1ReasonTooManyAttempts

// Equals "internal_system_error"
const StatusDetailsV1ReasonInternalSystemError = shared.StatusDetailsV1ReasonInternalSystemError

// Equals "user_request"
const StatusDetailsV1ReasonUserRequest = shared.StatusDetailsV1ReasonUserRequest

// Equals "ok"
const StatusDetailsV1ReasonOk = shared.StatusDetailsV1ReasonOk

// Equals "other_network_return"
const StatusDetailsV1ReasonOtherNetworkReturn = shared.StatusDetailsV1ReasonOtherNetworkReturn

// Equals "payout_refused"
const StatusDetailsV1ReasonPayoutRefused = shared.StatusDetailsV1ReasonPayoutRefused

// Equals "cancel_request"
const StatusDetailsV1ReasonCancelRequest = shared.StatusDetailsV1ReasonCancelRequest

// Equals "failed_verification"
const StatusDetailsV1ReasonFailedVerification = shared.StatusDetailsV1ReasonFailedVerification

// Equals "require_review"
const StatusDetailsV1ReasonRequireReview = shared.StatusDetailsV1ReasonRequireReview

// Equals "blocked_by_system"
const StatusDetailsV1ReasonBlockedBySystem = shared.StatusDetailsV1ReasonBlockedBySystem

// Equals "watchtower_review"
const StatusDetailsV1ReasonWatchtowerReview = shared.StatusDetailsV1ReasonWatchtowerReview

// Equals "validating"
const StatusDetailsV1ReasonValidating = shared.StatusDetailsV1ReasonValidating

// Equals "auto_hold"
const StatusDetailsV1ReasonAutoHold = shared.StatusDetailsV1ReasonAutoHold

// Identifies the origin of the status change (e.g., `bank_decline`, `watchtower`).
// This helps in tracking the cause of status updates.
//
// This is an alias to an internal type.
type StatusDetailsV1Source = shared.StatusDetailsV1Source

// Equals "watchtower"
const StatusDetailsV1SourceWatchtower = shared.StatusDetailsV1SourceWatchtower

// Equals "bank_decline"
const StatusDetailsV1SourceBankDecline = shared.StatusDetailsV1SourceBankDecline

// Equals "customer_dispute"
const StatusDetailsV1SourceCustomerDispute = shared.StatusDetailsV1SourceCustomerDispute

// Equals "user_action"
const StatusDetailsV1SourceUserAction = shared.StatusDetailsV1SourceUserAction

// Equals "system"
const StatusDetailsV1SourceSystem = shared.StatusDetailsV1SourceSystem
