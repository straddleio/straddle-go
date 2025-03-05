// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package straddle

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/stainless-sdks/straddle-go/internal/apijson"
	"github.com/stainless-sdks/straddle-go/internal/param"
	"github.com/stainless-sdks/straddle-go/internal/requestconfig"
	"github.com/stainless-sdks/straddle-go/option"
	"github.com/stainless-sdks/straddle-go/shared"
)

// ChargeService contains methods and other services that help with interacting
// with the straddle API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewChargeService] method instead.
type ChargeService struct {
	Options []option.RequestOption
}

// NewChargeService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewChargeService(opts ...option.RequestOption) (r *ChargeService) {
	r = &ChargeService{}
	r.Options = opts
	return
}

// Use charges to collect money from a customer for the sale of goods or services.
func (r *ChargeService) New(ctx context.Context, params ChargeNewParams, opts ...option.RequestOption) (res *ChargeV1, err error) {
	if params.CorrelationID.Present {
		opts = append(opts, option.WithHeader("Correlation-Id", fmt.Sprintf("%s", params.CorrelationID)))
	}
	if params.RequestID.Present {
		opts = append(opts, option.WithHeader("Request-Id", fmt.Sprintf("%s", params.RequestID)))
	}
	if params.StraddleAccountID.Present {
		opts = append(opts, option.WithHeader("Straddle-Account-Id", fmt.Sprintf("%s", params.StraddleAccountID)))
	}
	opts = append(r.Options[:], opts...)
	path := "v1/charges"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Change the values of parameters associated with a charge prior to processing.
// The status of the charge must be `created`, `scheduled`, or `on_hold`.
func (r *ChargeService) Update(ctx context.Context, id string, params ChargeUpdateParams, opts ...option.RequestOption) (res *ChargeV1, err error) {
	if params.CorrelationID.Present {
		opts = append(opts, option.WithHeader("Correlation-Id", fmt.Sprintf("%s", params.CorrelationID)))
	}
	if params.RequestID.Present {
		opts = append(opts, option.WithHeader("Request-Id", fmt.Sprintf("%s", params.RequestID)))
	}
	if params.StraddleAccountID.Present {
		opts = append(opts, option.WithHeader("Straddle-Account-Id", fmt.Sprintf("%s", params.StraddleAccountID)))
	}
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/charges/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

// Cancel a charge to prevent it from being originated for processing. The status
// of the charge must be `created`, `scheduled`, or `on_hold`.
func (r *ChargeService) Cancel(ctx context.Context, id string, params ChargeCancelParams, opts ...option.RequestOption) (res *ChargeV1, err error) {
	if params.CorrelationID.Present {
		opts = append(opts, option.WithHeader("Correlation-Id", fmt.Sprintf("%s", params.CorrelationID)))
	}
	if params.RequestID.Present {
		opts = append(opts, option.WithHeader("Request-Id", fmt.Sprintf("%s", params.RequestID)))
	}
	if params.StraddleAccountID.Present {
		opts = append(opts, option.WithHeader("Straddle-Account-Id", fmt.Sprintf("%s", params.StraddleAccountID)))
	}
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/charges/%s/cancel", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

// Retrieves the details of an existing charge. Supply the unique charge `id`, and
// Straddle will return the corresponding charge information.
func (r *ChargeService) Get(ctx context.Context, id string, query ChargeGetParams, opts ...option.RequestOption) (res *ChargeV1, err error) {
	if query.CorrelationID.Present {
		opts = append(opts, option.WithHeader("Correlation-Id", fmt.Sprintf("%s", query.CorrelationID)))
	}
	if query.RequestID.Present {
		opts = append(opts, option.WithHeader("Request-Id", fmt.Sprintf("%s", query.RequestID)))
	}
	if query.StraddleAccountID.Present {
		opts = append(opts, option.WithHeader("Straddle-Account-Id", fmt.Sprintf("%s", query.StraddleAccountID)))
	}
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/charges/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Place a charge on hold to prevent it from being originated for processing. The
// status of the charge must be `created` or `scheduled`.
func (r *ChargeService) Hold(ctx context.Context, id string, params ChargeHoldParams, opts ...option.RequestOption) (res *ChargeV1, err error) {
	if params.CorrelationID.Present {
		opts = append(opts, option.WithHeader("Correlation-Id", fmt.Sprintf("%s", params.CorrelationID)))
	}
	if params.RequestID.Present {
		opts = append(opts, option.WithHeader("Request-Id", fmt.Sprintf("%s", params.RequestID)))
	}
	if params.StraddleAccountID.Present {
		opts = append(opts, option.WithHeader("Straddle-Account-Id", fmt.Sprintf("%s", params.StraddleAccountID)))
	}
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/charges/%s/hold", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

// Release a charge from an `on_hold` status to allow it to be rescheduled for
// processing.
func (r *ChargeService) Release(ctx context.Context, id string, params ChargeReleaseParams, opts ...option.RequestOption) (res *ChargeV1, err error) {
	if params.CorrelationID.Present {
		opts = append(opts, option.WithHeader("Correlation-Id", fmt.Sprintf("%s", params.CorrelationID)))
	}
	if params.RequestID.Present {
		opts = append(opts, option.WithHeader("Request-Id", fmt.Sprintf("%s", params.RequestID)))
	}
	if params.StraddleAccountID.Present {
		opts = append(opts, option.WithHeader("Straddle-Account-Id", fmt.Sprintf("%s", params.StraddleAccountID)))
	}
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/charges/%s/release", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

type ChargeV1 struct {
	Data ChargeV1Data `json:"data,required"`
	// Metadata about the API request, including an identifier and timestamp.
	Meta shared.ResponseMetadata `json:"meta,required"`
	// Indicates the structure of the returned content.
	//
	//   - "object" means the `data` field contains a single JSON object.
	//   - "array" means the `data` field contains an array of objects.
	//   - "error" means the `data` field contains an error object with details of the
	//     issue.
	//   - "none" means no data is returned.
	ResponseType ChargeV1ResponseType `json:"response_type,required"`
	JSON         chargeV1JSON         `json:"-"`
}

// chargeV1JSON contains the JSON metadata for the struct [ChargeV1]
type chargeV1JSON struct {
	Data         apijson.Field
	Meta         apijson.Field
	ResponseType apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ChargeV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r chargeV1JSON) RawJSON() string {
	return r.raw
}

type ChargeV1Data struct {
	// Unique identifier for the charge.
	ID string `json:"id,required" format:"uuid"`
	// The amount of the charge in cents.
	Amount int64 `json:"amount,required"`
	// Configuration options for the charge.
	Config ChargeV1DataConfig `json:"config,required"`
	// The channel or mechanism through which the payment was authorized. Use
	// `internet` for payments made online or through a mobile app and `signed` for
	// signed agreements where there is a consent form or contract. Use `signed` for
	// PDF signatures.
	ConsentType ChargeV1DataConsentType `json:"consent_type,required"`
	// Timestamp of when the charge was created.
	CreatedAt time.Time `json:"created_at,required,nullable" format:"date-time"`
	// The currency of the charge. Only USD is supported.
	Currency string `json:"currency,required"`
	// An arbitrary description for the charge.
	Description string `json:"description,required"`
	// Information about the device used when the customer authorized the payment.
	Device shared.DeviceInfoV1 `json:"device,required"`
	// Unique identifier for the charge in your database. This value must be unique
	// across all charges.
	ExternalID string `json:"external_id,required"`
	// Value of the `paykey` used for the charge.
	Paykey string `json:"paykey,required"`
	// The desired date on which the payment should be occur. For charges, this means
	// the date you want the customer to be debited on.
	PaymentDate time.Time `json:"payment_date,required" format:"date"`
	// The current status of the charge.
	Status ChargeV1DataStatus `json:"status,required"`
	// Additional details about the current status of the charge.
	StatusDetails shared.StatusDetailsV1 `json:"status_details,required"`
	// Status history.
	StatusHistory []ChargeV1DataStatusHistory `json:"status_history,required"`
	// Timestamp of when the charge was last updated.
	UpdatedAt time.Time `json:"updated_at,required,nullable" format:"date-time"`
	// Information about the customer associated with the charge.
	CustomerDetails shared.CustomerDetailsV1 `json:"customer_details"`
	// Timestamp of when the charge was effective in the customer's bank account,
	// otherwise known as the date on which the customer is debited.
	EffectiveAt time.Time `json:"effective_at,nullable" format:"date-time"`
	// Up to 20 additional user-defined key-value pairs. Useful for storing additional
	// information about the charge in a structured format.
	Metadata map[string]string `json:"metadata,nullable"`
	// Information about the paykey used for the charge.
	PaykeyDetails shared.PaykeyDetailsV1 `json:"paykey_details"`
	// The payment rail that the charge will be processed through.
	PaymentRail ChargeV1DataPaymentRail `json:"payment_rail"`
	// Timestamp of when the charge was processed by Straddle and originated to the
	// payment rail.
	ProcessedAt time.Time        `json:"processed_at,nullable" format:"date-time"`
	JSON        chargeV1DataJSON `json:"-"`
}

// chargeV1DataJSON contains the JSON metadata for the struct [ChargeV1Data]
type chargeV1DataJSON struct {
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

func (r *ChargeV1Data) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r chargeV1DataJSON) RawJSON() string {
	return r.raw
}

// Configuration options for the charge.
type ChargeV1DataConfig struct {
	// Defines whether to check the customer's balance before processing the charge.
	BalanceCheck ChargeV1DataConfigBalanceCheck `json:"balance_check,required"`
	JSON         chargeV1DataConfigJSON         `json:"-"`
}

// chargeV1DataConfigJSON contains the JSON metadata for the struct
// [ChargeV1DataConfig]
type chargeV1DataConfigJSON struct {
	BalanceCheck apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ChargeV1DataConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r chargeV1DataConfigJSON) RawJSON() string {
	return r.raw
}

// Defines whether to check the customer's balance before processing the charge.
type ChargeV1DataConfigBalanceCheck string

const (
	ChargeV1DataConfigBalanceCheckRequired ChargeV1DataConfigBalanceCheck = "required"
	ChargeV1DataConfigBalanceCheckEnabled  ChargeV1DataConfigBalanceCheck = "enabled"
	ChargeV1DataConfigBalanceCheckDisabled ChargeV1DataConfigBalanceCheck = "disabled"
)

func (r ChargeV1DataConfigBalanceCheck) IsKnown() bool {
	switch r {
	case ChargeV1DataConfigBalanceCheckRequired, ChargeV1DataConfigBalanceCheckEnabled, ChargeV1DataConfigBalanceCheckDisabled:
		return true
	}
	return false
}

// The channel or mechanism through which the payment was authorized. Use
// `internet` for payments made online or through a mobile app and `signed` for
// signed agreements where there is a consent form or contract. Use `signed` for
// PDF signatures.
type ChargeV1DataConsentType string

const (
	ChargeV1DataConsentTypeInternet ChargeV1DataConsentType = "internet"
	ChargeV1DataConsentTypeSigned   ChargeV1DataConsentType = "signed"
)

func (r ChargeV1DataConsentType) IsKnown() bool {
	switch r {
	case ChargeV1DataConsentTypeInternet, ChargeV1DataConsentTypeSigned:
		return true
	}
	return false
}

// The current status of the charge.
type ChargeV1DataStatus string

const (
	ChargeV1DataStatusCreated   ChargeV1DataStatus = "created"
	ChargeV1DataStatusScheduled ChargeV1DataStatus = "scheduled"
	ChargeV1DataStatusFailed    ChargeV1DataStatus = "failed"
	ChargeV1DataStatusCancelled ChargeV1DataStatus = "cancelled"
	ChargeV1DataStatusOnHold    ChargeV1DataStatus = "on_hold"
	ChargeV1DataStatusPending   ChargeV1DataStatus = "pending"
	ChargeV1DataStatusPaid      ChargeV1DataStatus = "paid"
	ChargeV1DataStatusReversed  ChargeV1DataStatus = "reversed"
)

func (r ChargeV1DataStatus) IsKnown() bool {
	switch r {
	case ChargeV1DataStatusCreated, ChargeV1DataStatusScheduled, ChargeV1DataStatusFailed, ChargeV1DataStatusCancelled, ChargeV1DataStatusOnHold, ChargeV1DataStatusPending, ChargeV1DataStatusPaid, ChargeV1DataStatusReversed:
		return true
	}
	return false
}

// A record of the charge's status changes over time.
type ChargeV1DataStatusHistory struct {
	// The time the status change occurred.
	ChangedAt time.Time `json:"changed_at,required" format:"date-time"`
	// A human-readable description of the status.
	Message string `json:"message,required"`
	// A machine-readable identifier for the specific status, useful for programmatic
	// handling.
	Reason ChargeV1DataStatusHistoryReason `json:"reason,required"`
	// Identifies the origin of the status change (e.g., `bank_decline`, `watchtower`).
	// This helps in tracking the cause of status updates.
	Source ChargeV1DataStatusHistorySource `json:"source,required"`
	// The current status of the `charge` or `payout`.
	Status ChargeV1DataStatusHistoryStatus `json:"status,required"`
	// The status code if applicable.
	Code string                        `json:"code,nullable"`
	JSON chargeV1DataStatusHistoryJSON `json:"-"`
}

// chargeV1DataStatusHistoryJSON contains the JSON metadata for the struct
// [ChargeV1DataStatusHistory]
type chargeV1DataStatusHistoryJSON struct {
	ChangedAt   apijson.Field
	Message     apijson.Field
	Reason      apijson.Field
	Source      apijson.Field
	Status      apijson.Field
	Code        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ChargeV1DataStatusHistory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r chargeV1DataStatusHistoryJSON) RawJSON() string {
	return r.raw
}

// A machine-readable identifier for the specific status, useful for programmatic
// handling.
type ChargeV1DataStatusHistoryReason string

const (
	ChargeV1DataStatusHistoryReasonInsufficientFunds   ChargeV1DataStatusHistoryReason = "insufficient_funds"
	ChargeV1DataStatusHistoryReasonClosedBankAccount   ChargeV1DataStatusHistoryReason = "closed_bank_account"
	ChargeV1DataStatusHistoryReasonInvalidBankAccount  ChargeV1DataStatusHistoryReason = "invalid_bank_account"
	ChargeV1DataStatusHistoryReasonInvalidRouting      ChargeV1DataStatusHistoryReason = "invalid_routing"
	ChargeV1DataStatusHistoryReasonDisputed            ChargeV1DataStatusHistoryReason = "disputed"
	ChargeV1DataStatusHistoryReasonPaymentStopped      ChargeV1DataStatusHistoryReason = "payment_stopped"
	ChargeV1DataStatusHistoryReasonOwnerDeceased       ChargeV1DataStatusHistoryReason = "owner_deceased"
	ChargeV1DataStatusHistoryReasonFrozenBankAccount   ChargeV1DataStatusHistoryReason = "frozen_bank_account"
	ChargeV1DataStatusHistoryReasonRiskReview          ChargeV1DataStatusHistoryReason = "risk_review"
	ChargeV1DataStatusHistoryReasonFraudulent          ChargeV1DataStatusHistoryReason = "fraudulent"
	ChargeV1DataStatusHistoryReasonDuplicateEntry      ChargeV1DataStatusHistoryReason = "duplicate_entry"
	ChargeV1DataStatusHistoryReasonInvalidPaykey       ChargeV1DataStatusHistoryReason = "invalid_paykey"
	ChargeV1DataStatusHistoryReasonPaymentBlocked      ChargeV1DataStatusHistoryReason = "payment_blocked"
	ChargeV1DataStatusHistoryReasonAmountTooLarge      ChargeV1DataStatusHistoryReason = "amount_too_large"
	ChargeV1DataStatusHistoryReasonTooManyAttempts     ChargeV1DataStatusHistoryReason = "too_many_attempts"
	ChargeV1DataStatusHistoryReasonInternalSystemError ChargeV1DataStatusHistoryReason = "internal_system_error"
	ChargeV1DataStatusHistoryReasonUserRequest         ChargeV1DataStatusHistoryReason = "user_request"
	ChargeV1DataStatusHistoryReasonOk                  ChargeV1DataStatusHistoryReason = "ok"
	ChargeV1DataStatusHistoryReasonOtherNetworkReturn  ChargeV1DataStatusHistoryReason = "other_network_return"
	ChargeV1DataStatusHistoryReasonPayoutRefused       ChargeV1DataStatusHistoryReason = "payout_refused"
)

func (r ChargeV1DataStatusHistoryReason) IsKnown() bool {
	switch r {
	case ChargeV1DataStatusHistoryReasonInsufficientFunds, ChargeV1DataStatusHistoryReasonClosedBankAccount, ChargeV1DataStatusHistoryReasonInvalidBankAccount, ChargeV1DataStatusHistoryReasonInvalidRouting, ChargeV1DataStatusHistoryReasonDisputed, ChargeV1DataStatusHistoryReasonPaymentStopped, ChargeV1DataStatusHistoryReasonOwnerDeceased, ChargeV1DataStatusHistoryReasonFrozenBankAccount, ChargeV1DataStatusHistoryReasonRiskReview, ChargeV1DataStatusHistoryReasonFraudulent, ChargeV1DataStatusHistoryReasonDuplicateEntry, ChargeV1DataStatusHistoryReasonInvalidPaykey, ChargeV1DataStatusHistoryReasonPaymentBlocked, ChargeV1DataStatusHistoryReasonAmountTooLarge, ChargeV1DataStatusHistoryReasonTooManyAttempts, ChargeV1DataStatusHistoryReasonInternalSystemError, ChargeV1DataStatusHistoryReasonUserRequest, ChargeV1DataStatusHistoryReasonOk, ChargeV1DataStatusHistoryReasonOtherNetworkReturn, ChargeV1DataStatusHistoryReasonPayoutRefused:
		return true
	}
	return false
}

// Identifies the origin of the status change (e.g., `bank_decline`, `watchtower`).
// This helps in tracking the cause of status updates.
type ChargeV1DataStatusHistorySource string

const (
	ChargeV1DataStatusHistorySourceWatchtower      ChargeV1DataStatusHistorySource = "watchtower"
	ChargeV1DataStatusHistorySourceBankDecline     ChargeV1DataStatusHistorySource = "bank_decline"
	ChargeV1DataStatusHistorySourceCustomerDispute ChargeV1DataStatusHistorySource = "customer_dispute"
	ChargeV1DataStatusHistorySourceUserAction      ChargeV1DataStatusHistorySource = "user_action"
	ChargeV1DataStatusHistorySourceSystem          ChargeV1DataStatusHistorySource = "system"
)

func (r ChargeV1DataStatusHistorySource) IsKnown() bool {
	switch r {
	case ChargeV1DataStatusHistorySourceWatchtower, ChargeV1DataStatusHistorySourceBankDecline, ChargeV1DataStatusHistorySourceCustomerDispute, ChargeV1DataStatusHistorySourceUserAction, ChargeV1DataStatusHistorySourceSystem:
		return true
	}
	return false
}

// The current status of the `charge` or `payout`.
type ChargeV1DataStatusHistoryStatus string

const (
	ChargeV1DataStatusHistoryStatusCreated   ChargeV1DataStatusHistoryStatus = "created"
	ChargeV1DataStatusHistoryStatusScheduled ChargeV1DataStatusHistoryStatus = "scheduled"
	ChargeV1DataStatusHistoryStatusFailed    ChargeV1DataStatusHistoryStatus = "failed"
	ChargeV1DataStatusHistoryStatusCancelled ChargeV1DataStatusHistoryStatus = "cancelled"
	ChargeV1DataStatusHistoryStatusOnHold    ChargeV1DataStatusHistoryStatus = "on_hold"
	ChargeV1DataStatusHistoryStatusPending   ChargeV1DataStatusHistoryStatus = "pending"
	ChargeV1DataStatusHistoryStatusPaid      ChargeV1DataStatusHistoryStatus = "paid"
	ChargeV1DataStatusHistoryStatusReversed  ChargeV1DataStatusHistoryStatus = "reversed"
)

func (r ChargeV1DataStatusHistoryStatus) IsKnown() bool {
	switch r {
	case ChargeV1DataStatusHistoryStatusCreated, ChargeV1DataStatusHistoryStatusScheduled, ChargeV1DataStatusHistoryStatusFailed, ChargeV1DataStatusHistoryStatusCancelled, ChargeV1DataStatusHistoryStatusOnHold, ChargeV1DataStatusHistoryStatusPending, ChargeV1DataStatusHistoryStatusPaid, ChargeV1DataStatusHistoryStatusReversed:
		return true
	}
	return false
}

// The payment rail that the charge will be processed through.
type ChargeV1DataPaymentRail string

const (
	ChargeV1DataPaymentRailACH ChargeV1DataPaymentRail = "ach"
)

func (r ChargeV1DataPaymentRail) IsKnown() bool {
	switch r {
	case ChargeV1DataPaymentRailACH:
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
type ChargeV1ResponseType string

const (
	ChargeV1ResponseTypeObject ChargeV1ResponseType = "object"
	ChargeV1ResponseTypeArray  ChargeV1ResponseType = "array"
	ChargeV1ResponseTypeError  ChargeV1ResponseType = "error"
	ChargeV1ResponseTypeNone   ChargeV1ResponseType = "none"
)

func (r ChargeV1ResponseType) IsKnown() bool {
	switch r {
	case ChargeV1ResponseTypeObject, ChargeV1ResponseTypeArray, ChargeV1ResponseTypeError, ChargeV1ResponseTypeNone:
		return true
	}
	return false
}

type ChargeNewParams struct {
	// The amount of the charge in cents.
	Amount param.Field[int64]                 `json:"amount,required"`
	Config param.Field[ChargeNewParamsConfig] `json:"config,required"`
	// The channel or mechanism through which the payment was authorized. Use
	// `internet` for payments made online or through a mobile app and `signed` for
	// signed agreements where there is a consent form or contract. Use `signed` for
	// PDF signatures.
	ConsentType param.Field[ChargeNewParamsConsentType] `json:"consent_type,required"`
	// The currency of the charge. Only USD is supported.
	Currency param.Field[string] `json:"currency,required"`
	// An arbitrary description for the charge.
	Description param.Field[string]                   `json:"description,required"`
	Device      param.Field[shared.DeviceInfoV1Param] `json:"device,required"`
	// Unique identifier for the charge in your database. This value must be unique
	// across all charges.
	ExternalID param.Field[string] `json:"external_id,required"`
	// Value of the `paykey` used for the charge.
	Paykey param.Field[string] `json:"paykey,required"`
	// The desired date on which the payment should be occur. For charges, this means
	// the date you want the customer to be debited on.
	PaymentDate param.Field[time.Time] `json:"payment_date,required" format:"date"`
	// Up to 20 additional user-defined key-value pairs. Useful for storing additional
	// information about the charge in a structured format.
	Metadata          param.Field[map[string]string] `json:"metadata"`
	CorrelationID     param.Field[string]            `header:"Correlation-Id"`
	RequestID         param.Field[string]            `header:"Request-Id"`
	StraddleAccountID param.Field[string]            `header:"Straddle-Account-Id" format:"uuid"`
}

func (r ChargeNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ChargeNewParamsConfig struct {
	// Defines whether to check the customer's balance before processing the charge.
	BalanceCheck param.Field[ChargeNewParamsConfigBalanceCheck] `json:"balance_check,required"`
}

func (r ChargeNewParamsConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Defines whether to check the customer's balance before processing the charge.
type ChargeNewParamsConfigBalanceCheck string

const (
	ChargeNewParamsConfigBalanceCheckRequired ChargeNewParamsConfigBalanceCheck = "required"
	ChargeNewParamsConfigBalanceCheckEnabled  ChargeNewParamsConfigBalanceCheck = "enabled"
	ChargeNewParamsConfigBalanceCheckDisabled ChargeNewParamsConfigBalanceCheck = "disabled"
)

func (r ChargeNewParamsConfigBalanceCheck) IsKnown() bool {
	switch r {
	case ChargeNewParamsConfigBalanceCheckRequired, ChargeNewParamsConfigBalanceCheckEnabled, ChargeNewParamsConfigBalanceCheckDisabled:
		return true
	}
	return false
}

// The channel or mechanism through which the payment was authorized. Use
// `internet` for payments made online or through a mobile app and `signed` for
// signed agreements where there is a consent form or contract. Use `signed` for
// PDF signatures.
type ChargeNewParamsConsentType string

const (
	ChargeNewParamsConsentTypeInternet ChargeNewParamsConsentType = "internet"
	ChargeNewParamsConsentTypeSigned   ChargeNewParamsConsentType = "signed"
)

func (r ChargeNewParamsConsentType) IsKnown() bool {
	switch r {
	case ChargeNewParamsConsentTypeInternet, ChargeNewParamsConsentTypeSigned:
		return true
	}
	return false
}

type ChargeUpdateParams struct {
	// The amount of the charge in cents.
	Amount param.Field[int64] `json:"amount,required"`
	// An arbitrary description for the charge.
	Description param.Field[string] `json:"description,required"`
	// The desired date on which the payment should be occur. For charges, this means
	// the date you want the customer to be debited on.
	PaymentDate param.Field[time.Time] `json:"payment_date,required" format:"date"`
	// Up to 20 additional user-defined key-value pairs. Useful for storing additional
	// information about the charge in a structured format.
	Metadata          param.Field[map[string]string] `json:"metadata"`
	CorrelationID     param.Field[string]            `header:"Correlation-Id"`
	RequestID         param.Field[string]            `header:"Request-Id"`
	StraddleAccountID param.Field[string]            `header:"Straddle-Account-Id" format:"uuid"`
}

func (r ChargeUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ChargeCancelParams struct {
	// Details about why the charge status was updated.
	Reason            param.Field[string] `json:"reason"`
	CorrelationID     param.Field[string] `header:"Correlation-Id"`
	RequestID         param.Field[string] `header:"Request-Id"`
	StraddleAccountID param.Field[string] `header:"Straddle-Account-Id" format:"uuid"`
}

func (r ChargeCancelParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ChargeGetParams struct {
	CorrelationID     param.Field[string] `header:"Correlation-Id"`
	RequestID         param.Field[string] `header:"Request-Id"`
	StraddleAccountID param.Field[string] `header:"Straddle-Account-Id" format:"uuid"`
}

type ChargeHoldParams struct {
	// Details about why the charge status was updated.
	Reason            param.Field[string] `json:"reason"`
	CorrelationID     param.Field[string] `header:"Correlation-Id"`
	RequestID         param.Field[string] `header:"Request-Id"`
	StraddleAccountID param.Field[string] `header:"Straddle-Account-Id" format:"uuid"`
}

func (r ChargeHoldParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ChargeReleaseParams struct {
	// Details about why the charge status was updated.
	Reason            param.Field[string] `json:"reason"`
	CorrelationID     param.Field[string] `header:"Correlation-Id"`
	RequestID         param.Field[string] `header:"Request-Id"`
	StraddleAccountID param.Field[string] `header:"Straddle-Account-Id" format:"uuid"`
}

func (r ChargeReleaseParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
