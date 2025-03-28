// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package straddle

import (
	"github.com/stainless-sdks/straddle-go/internal/apierror"
	"github.com/stainless-sdks/straddle-go/shared"
)

type Error = apierror.Error

// Information about the customer associated with the charge or payout.
//
// This is an alias to an internal type.
type CustomerDetailsV1 = shared.CustomerDetailsV1

// The type of customer
//
// This is an alias to an internal type.
type CustomerDetailsV1CustomerType = shared.CustomerDetailsV1CustomerType

// This is an alias to an internal value.
const CustomerDetailsV1CustomerTypeIndividual = shared.CustomerDetailsV1CustomerTypeIndividual

// This is an alias to an internal value.
const CustomerDetailsV1CustomerTypeBusiness = shared.CustomerDetailsV1CustomerTypeBusiness

// This is an alias to an internal value.
const CustomerDetailsV1CustomerTypeUnknown = shared.CustomerDetailsV1CustomerTypeUnknown

// This is an alias to an internal type.
type DeviceInfoV1 = shared.DeviceInfoV1

// This is an alias to an internal type.
type DeviceInfoV1Param = shared.DeviceInfoV1Param

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

// This is an alias to an internal value.
const StatusDetailsV1ReasonInsufficientFunds = shared.StatusDetailsV1ReasonInsufficientFunds

// This is an alias to an internal value.
const StatusDetailsV1ReasonClosedBankAccount = shared.StatusDetailsV1ReasonClosedBankAccount

// This is an alias to an internal value.
const StatusDetailsV1ReasonInvalidBankAccount = shared.StatusDetailsV1ReasonInvalidBankAccount

// This is an alias to an internal value.
const StatusDetailsV1ReasonInvalidRouting = shared.StatusDetailsV1ReasonInvalidRouting

// This is an alias to an internal value.
const StatusDetailsV1ReasonDisputed = shared.StatusDetailsV1ReasonDisputed

// This is an alias to an internal value.
const StatusDetailsV1ReasonPaymentStopped = shared.StatusDetailsV1ReasonPaymentStopped

// This is an alias to an internal value.
const StatusDetailsV1ReasonOwnerDeceased = shared.StatusDetailsV1ReasonOwnerDeceased

// This is an alias to an internal value.
const StatusDetailsV1ReasonFrozenBankAccount = shared.StatusDetailsV1ReasonFrozenBankAccount

// This is an alias to an internal value.
const StatusDetailsV1ReasonRiskReview = shared.StatusDetailsV1ReasonRiskReview

// This is an alias to an internal value.
const StatusDetailsV1ReasonFraudulent = shared.StatusDetailsV1ReasonFraudulent

// This is an alias to an internal value.
const StatusDetailsV1ReasonDuplicateEntry = shared.StatusDetailsV1ReasonDuplicateEntry

// This is an alias to an internal value.
const StatusDetailsV1ReasonInvalidPaykey = shared.StatusDetailsV1ReasonInvalidPaykey

// This is an alias to an internal value.
const StatusDetailsV1ReasonPaymentBlocked = shared.StatusDetailsV1ReasonPaymentBlocked

// This is an alias to an internal value.
const StatusDetailsV1ReasonAmountTooLarge = shared.StatusDetailsV1ReasonAmountTooLarge

// This is an alias to an internal value.
const StatusDetailsV1ReasonTooManyAttempts = shared.StatusDetailsV1ReasonTooManyAttempts

// This is an alias to an internal value.
const StatusDetailsV1ReasonInternalSystemError = shared.StatusDetailsV1ReasonInternalSystemError

// This is an alias to an internal value.
const StatusDetailsV1ReasonUserRequest = shared.StatusDetailsV1ReasonUserRequest

// This is an alias to an internal value.
const StatusDetailsV1ReasonOk = shared.StatusDetailsV1ReasonOk

// This is an alias to an internal value.
const StatusDetailsV1ReasonOtherNetworkReturn = shared.StatusDetailsV1ReasonOtherNetworkReturn

// This is an alias to an internal value.
const StatusDetailsV1ReasonPayoutRefused = shared.StatusDetailsV1ReasonPayoutRefused

// Identifies the origin of the status change (e.g., `bank_decline`, `watchtower`).
// This helps in tracking the cause of status updates.
//
// This is an alias to an internal type.
type StatusDetailsV1Source = shared.StatusDetailsV1Source

// This is an alias to an internal value.
const StatusDetailsV1SourceWatchtower = shared.StatusDetailsV1SourceWatchtower

// This is an alias to an internal value.
const StatusDetailsV1SourceBankDecline = shared.StatusDetailsV1SourceBankDecline

// This is an alias to an internal value.
const StatusDetailsV1SourceCustomerDispute = shared.StatusDetailsV1SourceCustomerDispute

// This is an alias to an internal value.
const StatusDetailsV1SourceUserAction = shared.StatusDetailsV1SourceUserAction

// This is an alias to an internal value.
const StatusDetailsV1SourceSystem = shared.StatusDetailsV1SourceSystem
