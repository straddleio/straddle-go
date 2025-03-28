// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

import (
	"time"

	"github.com/stainless-sdks/straddle-go/internal/apijson"
	"github.com/stainless-sdks/straddle-go/internal/param"
)

// Information about the customer associated with the charge or payout.
type CustomerDetailsV1 struct {
	// Unique identifier for the customer
	ID string `json:"id,required" format:"uuid"`
	// The type of customer
	CustomerType CustomerDetailsV1CustomerType `json:"customer_type,required"`
	// The customer's email address
	Email string `json:"email,required"`
	// The name of the customer
	Name string `json:"name,required"`
	// The customer's phone number in E.164 format
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

// The type of customer
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

type PaykeyDetailsV1 struct {
	// Unique identifier for the paykey.
	ID string `json:"id,required" format:"uuid"`
	// Unique identifier for the customer associated with the paykey.
	CustomerID string `json:"customer_id,required" format:"uuid"`
	// Human-readable label that combines the bank name and masked account number to
	// help easility represent this paykey in a UI
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

type StatusDetailsV1 struct {
	// The time the status change occurred.
	ChangedAt time.Time `json:"changed_at,required" format:"date-time"`
	// A human-readable description of the current status.
	Message string `json:"message,required"`
	// A machine-readable identifier for the specific status, useful for programmatic
	// handling.
	Reason StatusDetailsV1Reason `json:"reason,required"`
	// Identifies the origin of the status change (e.g., `bank_decline`, `watchtower`).
	// This helps in tracking the cause of status updates.
	Source StatusDetailsV1Source `json:"source,required"`
	// The status code if applicable.
	Code string              `json:"code,nullable"`
	JSON statusDetailsV1JSON `json:"-"`
}

// statusDetailsV1JSON contains the JSON metadata for the struct [StatusDetailsV1]
type statusDetailsV1JSON struct {
	ChangedAt   apijson.Field
	Message     apijson.Field
	Reason      apijson.Field
	Source      apijson.Field
	Code        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StatusDetailsV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r statusDetailsV1JSON) RawJSON() string {
	return r.raw
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
)

func (r StatusDetailsV1Reason) IsKnown() bool {
	switch r {
	case StatusDetailsV1ReasonInsufficientFunds, StatusDetailsV1ReasonClosedBankAccount, StatusDetailsV1ReasonInvalidBankAccount, StatusDetailsV1ReasonInvalidRouting, StatusDetailsV1ReasonDisputed, StatusDetailsV1ReasonPaymentStopped, StatusDetailsV1ReasonOwnerDeceased, StatusDetailsV1ReasonFrozenBankAccount, StatusDetailsV1ReasonRiskReview, StatusDetailsV1ReasonFraudulent, StatusDetailsV1ReasonDuplicateEntry, StatusDetailsV1ReasonInvalidPaykey, StatusDetailsV1ReasonPaymentBlocked, StatusDetailsV1ReasonAmountTooLarge, StatusDetailsV1ReasonTooManyAttempts, StatusDetailsV1ReasonInternalSystemError, StatusDetailsV1ReasonUserRequest, StatusDetailsV1ReasonOk, StatusDetailsV1ReasonOtherNetworkReturn, StatusDetailsV1ReasonPayoutRefused:
		return true
	}
	return false
}

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

func (r StatusDetailsV1Source) IsKnown() bool {
	switch r {
	case StatusDetailsV1SourceWatchtower, StatusDetailsV1SourceBankDecline, StatusDetailsV1SourceCustomerDispute, StatusDetailsV1SourceUserAction, StatusDetailsV1SourceSystem:
		return true
	}
	return false
}
