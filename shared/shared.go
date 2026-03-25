// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

import (
	"encoding/json"
	"time"

	"github.com/straddleio/straddle-go/internal/apijson"
	"github.com/straddleio/straddle-go/packages/param"
	"github.com/straddleio/straddle-go/packages/respjson"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

// Information about the customer associated with the charge or payout.
type CustomerDetailsV1 struct {
	// Unique identifier for the customer
	ID string `json:"id" api:"required" format:"uuid"`
	// The type of customer
	//
	// Any of "individual", "business".
	CustomerType CustomerDetailsV1CustomerType `json:"customer_type" api:"required"`
	// The customer's email address
	Email string `json:"email" api:"required"`
	// The name of the customer
	Name string `json:"name" api:"required"`
	// The customer's phone number in E.164 format
	Phone string `json:"phone" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		CustomerType respjson.Field
		Email        respjson.Field
		Name         respjson.Field
		Phone        respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CustomerDetailsV1) RawJSON() string { return r.JSON.raw }
func (r *CustomerDetailsV1) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of customer
type CustomerDetailsV1CustomerType string

const (
	CustomerDetailsV1CustomerTypeIndividual CustomerDetailsV1CustomerType = "individual"
	CustomerDetailsV1CustomerTypeBusiness   CustomerDetailsV1CustomerType = "business"
)

type DeviceInfoV1 struct {
	// The IP address of the device used when the customer authorized the charge or
	// payout. Use `0.0.0.0` to represent an offline consent interaction.
	IPAddress string `json:"ip_address" api:"required" format:"ipv4"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IPAddress   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DeviceInfoV1) RawJSON() string { return r.JSON.raw }
func (r *DeviceInfoV1) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this DeviceInfoV1 to a DeviceInfoV1Param.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// DeviceInfoV1Param.Overrides()
func (r DeviceInfoV1) ToParam() DeviceInfoV1Param {
	return param.Override[DeviceInfoV1Param](json.RawMessage(r.RawJSON()))
}

// The property IPAddress is required.
type DeviceInfoV1Param struct {
	// The IP address of the device used when the customer authorized the charge or
	// payout. Use `0.0.0.0` to represent an offline consent interaction.
	IPAddress string `json:"ip_address" api:"required" format:"ipv4"`
	paramObj
}

func (r DeviceInfoV1Param) MarshalJSON() (data []byte, err error) {
	type shadow DeviceInfoV1Param
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DeviceInfoV1Param) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Metadata about the API request, including an identifier, timestamp, and
// pagination details.
type PagedResponseMetadata struct {
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
	SortOrder PagedResponseMetadataSortOrder `json:"sort_order" api:"required"`
	// Total number of items returned in this response.
	TotalItems int64 `json:"total_items" api:"required"`
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
func (r PagedResponseMetadata) RawJSON() string { return r.JSON.raw }
func (r *PagedResponseMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PagedResponseMetadataSortOrder string

const (
	PagedResponseMetadataSortOrderAsc  PagedResponseMetadataSortOrder = "asc"
	PagedResponseMetadataSortOrderDesc PagedResponseMetadataSortOrder = "desc"
)

type PaykeyDetailsV1 struct {
	// Unique identifier for the paykey.
	ID string `json:"id" api:"required" format:"uuid"`
	// Unique identifier for the customer associated with the paykey.
	CustomerID string `json:"customer_id" api:"required" format:"uuid"`
	// Human-readable label that combines the bank name and masked account number to
	// help easility represent this paykey in a UI
	Label string `json:"label" api:"required"`
	// The most recent balance of the bank account associated with the paykey in
	// dollars.
	Balance int64 `json:"balance" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CustomerID  respjson.Field
		Label       respjson.Field
		Balance     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaykeyDetailsV1) RawJSON() string { return r.JSON.raw }
func (r *PaykeyDetailsV1) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Metadata about the API request, including an identifier and timestamp.
type ResponseMetadata struct {
	// Unique identifier for this API request, useful for troubleshooting.
	APIRequestID string `json:"api_request_id" api:"required" format:"uuid"`
	// Timestamp for this API request, useful for troubleshooting.
	APIRequestTimestamp time.Time `json:"api_request_timestamp" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		APIRequestID        respjson.Field
		APIRequestTimestamp respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseMetadata) RawJSON() string { return r.JSON.raw }
func (r *ResponseMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StatusDetailsV1 struct {
	// The time the status change occurred.
	ChangedAt time.Time `json:"changed_at" api:"required" format:"date-time"`
	// A human-readable description of the current status.
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
	Reason StatusDetailsV1Reason `json:"reason" api:"required"`
	// Identifies the origin of the status change (e.g., `bank_decline`, `watchtower`).
	// This helps in tracking the cause of status updates.
	//
	// Any of "watchtower", "bank_decline", "customer_dispute", "user_action",
	// "system".
	Source StatusDetailsV1Source `json:"source" api:"required"`
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
func (r StatusDetailsV1) RawJSON() string { return r.JSON.raw }
func (r *StatusDetailsV1) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A machine-readable identifier for the specific status, useful for programmatic
// handling.
type StatusDetailsV1Reason string

const (
	StatusDetailsV1ReasonInsufficientFunds   StatusDetailsV1Reason = "insufficient_funds"
	StatusDetailsV1ReasonClosedBankAccount   StatusDetailsV1Reason = "closed_bank_account"
	StatusDetailsV1ReasonInvalidBankAccount  StatusDetailsV1Reason = "invalid_bank_account"
	StatusDetailsV1ReasonInvalidRouting      StatusDetailsV1Reason = "invalid_routing"
	StatusDetailsV1ReasonDisputed            StatusDetailsV1Reason = "disputed"
	StatusDetailsV1ReasonPaymentStopped      StatusDetailsV1Reason = "payment_stopped"
	StatusDetailsV1ReasonOwnerDeceased       StatusDetailsV1Reason = "owner_deceased"
	StatusDetailsV1ReasonFrozenBankAccount   StatusDetailsV1Reason = "frozen_bank_account"
	StatusDetailsV1ReasonRiskReview          StatusDetailsV1Reason = "risk_review"
	StatusDetailsV1ReasonFraudulent          StatusDetailsV1Reason = "fraudulent"
	StatusDetailsV1ReasonDuplicateEntry      StatusDetailsV1Reason = "duplicate_entry"
	StatusDetailsV1ReasonInvalidPaykey       StatusDetailsV1Reason = "invalid_paykey"
	StatusDetailsV1ReasonPaymentBlocked      StatusDetailsV1Reason = "payment_blocked"
	StatusDetailsV1ReasonAmountTooLarge      StatusDetailsV1Reason = "amount_too_large"
	StatusDetailsV1ReasonTooManyAttempts     StatusDetailsV1Reason = "too_many_attempts"
	StatusDetailsV1ReasonInternalSystemError StatusDetailsV1Reason = "internal_system_error"
	StatusDetailsV1ReasonUserRequest         StatusDetailsV1Reason = "user_request"
	StatusDetailsV1ReasonOk                  StatusDetailsV1Reason = "ok"
	StatusDetailsV1ReasonOtherNetworkReturn  StatusDetailsV1Reason = "other_network_return"
	StatusDetailsV1ReasonPayoutRefused       StatusDetailsV1Reason = "payout_refused"
	StatusDetailsV1ReasonCancelRequest       StatusDetailsV1Reason = "cancel_request"
	StatusDetailsV1ReasonFailedVerification  StatusDetailsV1Reason = "failed_verification"
	StatusDetailsV1ReasonRequireReview       StatusDetailsV1Reason = "require_review"
	StatusDetailsV1ReasonBlockedBySystem     StatusDetailsV1Reason = "blocked_by_system"
	StatusDetailsV1ReasonWatchtowerReview    StatusDetailsV1Reason = "watchtower_review"
	StatusDetailsV1ReasonValidating          StatusDetailsV1Reason = "validating"
	StatusDetailsV1ReasonAutoHold            StatusDetailsV1Reason = "auto_hold"
)

// Identifies the origin of the status change (e.g., `bank_decline`, `watchtower`).
// This helps in tracking the cause of status updates.
type StatusDetailsV1Source string

const (
	StatusDetailsV1SourceWatchtower      StatusDetailsV1Source = "watchtower"
	StatusDetailsV1SourceBankDecline     StatusDetailsV1Source = "bank_decline"
	StatusDetailsV1SourceCustomerDispute StatusDetailsV1Source = "customer_dispute"
	StatusDetailsV1SourceUserAction      StatusDetailsV1Source = "user_action"
	StatusDetailsV1SourceSystem          StatusDetailsV1Source = "system"
)
