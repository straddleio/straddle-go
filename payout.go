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

// PayoutService contains methods and other services that help with interacting
// with the straddle API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPayoutService] method instead.
type PayoutService struct {
	Options []option.RequestOption
}

// NewPayoutService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewPayoutService(opts ...option.RequestOption) (r *PayoutService) {
	r = &PayoutService{}
	r.Options = opts
	return
}

// Use payouts to send money to your customers.
func (r *PayoutService) New(ctx context.Context, params PayoutNewParams, opts ...option.RequestOption) (res *PayoutV1, err error) {
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
	path := "v1/payouts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Update the details of a payout prior to processing. The status of the payout
// must be `created`, `scheduled`, or `on_hold`.
func (r *PayoutService) Update(ctx context.Context, id string, params PayoutUpdateParams, opts ...option.RequestOption) (res *PayoutV1, err error) {
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
	path := fmt.Sprintf("v1/payouts/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

// Cancel a payout to prevent it from being processed. The status of the payout
// must be `created`, `scheduled`, or `on_hold`.
func (r *PayoutService) Cancel(ctx context.Context, id string, params PayoutCancelParams, opts ...option.RequestOption) (res *PayoutV1, err error) {
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
	path := fmt.Sprintf("v1/payouts/%s/cancel", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

// Retrieves the details of an existing payout. Supply the unique payout `id` to
// retrieve the corresponding payout information.
func (r *PayoutService) Get(ctx context.Context, id string, query PayoutGetParams, opts ...option.RequestOption) (res *PayoutV1, err error) {
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
	path := fmt.Sprintf("v1/payouts/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Hold a payout to prevent it from being processed. The status of the payout must
// be `created`, `scheduled`, or `on_hold`.
func (r *PayoutService) Hold(ctx context.Context, id string, params PayoutHoldParams, opts ...option.RequestOption) (res *PayoutV1, err error) {
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
	path := fmt.Sprintf("v1/payouts/%s/hold", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

// Release a payout from a `hold` status to allow it to be rescheduled for
// processing.
func (r *PayoutService) Release(ctx context.Context, id string, params PayoutReleaseParams, opts ...option.RequestOption) (res *PayoutV1, err error) {
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
	path := fmt.Sprintf("v1/payouts/%s/release", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

type PayoutV1 struct {
	Data PayoutV1Data `json:"data,required"`
	// Metadata about the API request, including an identifier and timestamp.
	Meta shared.ResponseMetadata `json:"meta,required"`
	// Indicates the structure of the returned content.
	//
	//   - "object" means the `data` field contains a single JSON object.
	//   - "array" means the `data` field contains an array of objects.
	//   - "error" means the `data` field contains an error object with details of the
	//     issue.
	//   - "none" means no data is returned.
	ResponseType PayoutV1ResponseType `json:"response_type,required"`
	JSON         payoutV1JSON         `json:"-"`
}

// payoutV1JSON contains the JSON metadata for the struct [PayoutV1]
type payoutV1JSON struct {
	Data         apijson.Field
	Meta         apijson.Field
	ResponseType apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *PayoutV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r payoutV1JSON) RawJSON() string {
	return r.raw
}

type PayoutV1Data struct {
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
	Device shared.DeviceInfoV1 `json:"device,required"`
	// Unique identifier for the payout in your database. This value must be unique
	// across all payouts.
	ExternalID string `json:"external_id,required"`
	// Funding Ids
	FundingIDs []string `json:"funding_ids,required" format:"uuid"`
	// Value of the `paykey` used for the payout.
	Paykey string `json:"paykey,required"`
	// The desired date on which the payment should be occur. For payouts, this means
	// the date you want the funds to be sent from your bank account.
	PaymentDate time.Time `json:"payment_date,required" format:"date"`
	// The current status of the payout.
	Status PayoutV1DataStatus `json:"status,required"`
	// Details about the current status of the payout.
	StatusDetails shared.StatusDetailsV1 `json:"status_details,required"`
	// History of the status changes for the payout.
	StatusHistory []PayoutV1DataStatusHistory `json:"status_history,required"`
	// The time the payout was created.
	CreatedAt time.Time `json:"created_at,nullable" format:"date-time"`
	// Information about the customer associated with the payout.
	CustomerDetails shared.CustomerDetailsV1 `json:"customer_details"`
	// The actual date on which the payment occurred. For payouts, this is the date the
	// funds were sent from your bank account.
	EffectiveAt time.Time `json:"effective_at,nullable" format:"date-time"`
	// Up to 20 additional user-defined key-value pairs. Useful for storing additional
	// information about the payout in a structured format.
	Metadata map[string]string `json:"metadata,nullable"`
	// Information about the paykey used for the payout.
	PaykeyDetails shared.PaykeyDetailsV1 `json:"paykey_details"`
	// The payment rail used for the payout.
	PaymentRail PayoutV1DataPaymentRail `json:"payment_rail"`
	// The time the payout was processed by Straddle and originated to the payment
	// rail.
	ProcessedAt time.Time `json:"processed_at,nullable" format:"date-time"`
	// The time the payout was last updated.
	UpdatedAt time.Time        `json:"updated_at,nullable" format:"date-time"`
	JSON      payoutV1DataJSON `json:"-"`
}

// payoutV1DataJSON contains the JSON metadata for the struct [PayoutV1Data]
type payoutV1DataJSON struct {
	ID              apijson.Field
	Amount          apijson.Field
	Config          apijson.Field
	Currency        apijson.Field
	Description     apijson.Field
	Device          apijson.Field
	ExternalID      apijson.Field
	FundingIDs      apijson.Field
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

func (r *PayoutV1Data) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r payoutV1DataJSON) RawJSON() string {
	return r.raw
}

// The current status of the payout.
type PayoutV1DataStatus string

const (
	PayoutV1DataStatusCreated   PayoutV1DataStatus = "created"
	PayoutV1DataStatusScheduled PayoutV1DataStatus = "scheduled"
	PayoutV1DataStatusFailed    PayoutV1DataStatus = "failed"
	PayoutV1DataStatusCancelled PayoutV1DataStatus = "cancelled"
	PayoutV1DataStatusOnHold    PayoutV1DataStatus = "on_hold"
	PayoutV1DataStatusPending   PayoutV1DataStatus = "pending"
	PayoutV1DataStatusPaid      PayoutV1DataStatus = "paid"
	PayoutV1DataStatusReversed  PayoutV1DataStatus = "reversed"
)

func (r PayoutV1DataStatus) IsKnown() bool {
	switch r {
	case PayoutV1DataStatusCreated, PayoutV1DataStatusScheduled, PayoutV1DataStatusFailed, PayoutV1DataStatusCancelled, PayoutV1DataStatusOnHold, PayoutV1DataStatusPending, PayoutV1DataStatusPaid, PayoutV1DataStatusReversed:
		return true
	}
	return false
}

type PayoutV1DataStatusHistory struct {
	// The time the status change occurred.
	ChangedAt time.Time `json:"changed_at,required" format:"date-time"`
	// A human-readable description of the status.
	Message string `json:"message,required"`
	// A machine-readable identifier for the specific status, useful for programmatic
	// handling.
	Reason PayoutV1DataStatusHistoryReason `json:"reason,required"`
	// Identifies the origin of the status change (e.g., `bank_decline`, `watchtower`).
	// This helps in tracking the cause of status updates.
	Source PayoutV1DataStatusHistorySource `json:"source,required"`
	// The current status of the `charge` or `payout`.
	Status PayoutV1DataStatusHistoryStatus `json:"status,required"`
	// The status code if applicable.
	Code string                        `json:"code,nullable"`
	JSON payoutV1DataStatusHistoryJSON `json:"-"`
}

// payoutV1DataStatusHistoryJSON contains the JSON metadata for the struct
// [PayoutV1DataStatusHistory]
type payoutV1DataStatusHistoryJSON struct {
	ChangedAt   apijson.Field
	Message     apijson.Field
	Reason      apijson.Field
	Source      apijson.Field
	Status      apijson.Field
	Code        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PayoutV1DataStatusHistory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r payoutV1DataStatusHistoryJSON) RawJSON() string {
	return r.raw
}

// A machine-readable identifier for the specific status, useful for programmatic
// handling.
type PayoutV1DataStatusHistoryReason string

const (
	PayoutV1DataStatusHistoryReasonInsufficientFunds   PayoutV1DataStatusHistoryReason = "insufficient_funds"
	PayoutV1DataStatusHistoryReasonClosedBankAccount   PayoutV1DataStatusHistoryReason = "closed_bank_account"
	PayoutV1DataStatusHistoryReasonInvalidBankAccount  PayoutV1DataStatusHistoryReason = "invalid_bank_account"
	PayoutV1DataStatusHistoryReasonInvalidRouting      PayoutV1DataStatusHistoryReason = "invalid_routing"
	PayoutV1DataStatusHistoryReasonDisputed            PayoutV1DataStatusHistoryReason = "disputed"
	PayoutV1DataStatusHistoryReasonPaymentStopped      PayoutV1DataStatusHistoryReason = "payment_stopped"
	PayoutV1DataStatusHistoryReasonOwnerDeceased       PayoutV1DataStatusHistoryReason = "owner_deceased"
	PayoutV1DataStatusHistoryReasonFrozenBankAccount   PayoutV1DataStatusHistoryReason = "frozen_bank_account"
	PayoutV1DataStatusHistoryReasonRiskReview          PayoutV1DataStatusHistoryReason = "risk_review"
	PayoutV1DataStatusHistoryReasonFraudulent          PayoutV1DataStatusHistoryReason = "fraudulent"
	PayoutV1DataStatusHistoryReasonDuplicateEntry      PayoutV1DataStatusHistoryReason = "duplicate_entry"
	PayoutV1DataStatusHistoryReasonInvalidPaykey       PayoutV1DataStatusHistoryReason = "invalid_paykey"
	PayoutV1DataStatusHistoryReasonPaymentBlocked      PayoutV1DataStatusHistoryReason = "payment_blocked"
	PayoutV1DataStatusHistoryReasonAmountTooLarge      PayoutV1DataStatusHistoryReason = "amount_too_large"
	PayoutV1DataStatusHistoryReasonTooManyAttempts     PayoutV1DataStatusHistoryReason = "too_many_attempts"
	PayoutV1DataStatusHistoryReasonInternalSystemError PayoutV1DataStatusHistoryReason = "internal_system_error"
	PayoutV1DataStatusHistoryReasonUserRequest         PayoutV1DataStatusHistoryReason = "user_request"
	PayoutV1DataStatusHistoryReasonOk                  PayoutV1DataStatusHistoryReason = "ok"
	PayoutV1DataStatusHistoryReasonOtherNetworkReturn  PayoutV1DataStatusHistoryReason = "other_network_return"
	PayoutV1DataStatusHistoryReasonPayoutRefused       PayoutV1DataStatusHistoryReason = "payout_refused"
)

func (r PayoutV1DataStatusHistoryReason) IsKnown() bool {
	switch r {
	case PayoutV1DataStatusHistoryReasonInsufficientFunds, PayoutV1DataStatusHistoryReasonClosedBankAccount, PayoutV1DataStatusHistoryReasonInvalidBankAccount, PayoutV1DataStatusHistoryReasonInvalidRouting, PayoutV1DataStatusHistoryReasonDisputed, PayoutV1DataStatusHistoryReasonPaymentStopped, PayoutV1DataStatusHistoryReasonOwnerDeceased, PayoutV1DataStatusHistoryReasonFrozenBankAccount, PayoutV1DataStatusHistoryReasonRiskReview, PayoutV1DataStatusHistoryReasonFraudulent, PayoutV1DataStatusHistoryReasonDuplicateEntry, PayoutV1DataStatusHistoryReasonInvalidPaykey, PayoutV1DataStatusHistoryReasonPaymentBlocked, PayoutV1DataStatusHistoryReasonAmountTooLarge, PayoutV1DataStatusHistoryReasonTooManyAttempts, PayoutV1DataStatusHistoryReasonInternalSystemError, PayoutV1DataStatusHistoryReasonUserRequest, PayoutV1DataStatusHistoryReasonOk, PayoutV1DataStatusHistoryReasonOtherNetworkReturn, PayoutV1DataStatusHistoryReasonPayoutRefused:
		return true
	}
	return false
}

// Identifies the origin of the status change (e.g., `bank_decline`, `watchtower`).
// This helps in tracking the cause of status updates.
type PayoutV1DataStatusHistorySource string

const (
	PayoutV1DataStatusHistorySourceWatchtower      PayoutV1DataStatusHistorySource = "watchtower"
	PayoutV1DataStatusHistorySourceBankDecline     PayoutV1DataStatusHistorySource = "bank_decline"
	PayoutV1DataStatusHistorySourceCustomerDispute PayoutV1DataStatusHistorySource = "customer_dispute"
	PayoutV1DataStatusHistorySourceUserAction      PayoutV1DataStatusHistorySource = "user_action"
	PayoutV1DataStatusHistorySourceSystem          PayoutV1DataStatusHistorySource = "system"
)

func (r PayoutV1DataStatusHistorySource) IsKnown() bool {
	switch r {
	case PayoutV1DataStatusHistorySourceWatchtower, PayoutV1DataStatusHistorySourceBankDecline, PayoutV1DataStatusHistorySourceCustomerDispute, PayoutV1DataStatusHistorySourceUserAction, PayoutV1DataStatusHistorySourceSystem:
		return true
	}
	return false
}

// The current status of the `charge` or `payout`.
type PayoutV1DataStatusHistoryStatus string

const (
	PayoutV1DataStatusHistoryStatusCreated   PayoutV1DataStatusHistoryStatus = "created"
	PayoutV1DataStatusHistoryStatusScheduled PayoutV1DataStatusHistoryStatus = "scheduled"
	PayoutV1DataStatusHistoryStatusFailed    PayoutV1DataStatusHistoryStatus = "failed"
	PayoutV1DataStatusHistoryStatusCancelled PayoutV1DataStatusHistoryStatus = "cancelled"
	PayoutV1DataStatusHistoryStatusOnHold    PayoutV1DataStatusHistoryStatus = "on_hold"
	PayoutV1DataStatusHistoryStatusPending   PayoutV1DataStatusHistoryStatus = "pending"
	PayoutV1DataStatusHistoryStatusPaid      PayoutV1DataStatusHistoryStatus = "paid"
	PayoutV1DataStatusHistoryStatusReversed  PayoutV1DataStatusHistoryStatus = "reversed"
)

func (r PayoutV1DataStatusHistoryStatus) IsKnown() bool {
	switch r {
	case PayoutV1DataStatusHistoryStatusCreated, PayoutV1DataStatusHistoryStatusScheduled, PayoutV1DataStatusHistoryStatusFailed, PayoutV1DataStatusHistoryStatusCancelled, PayoutV1DataStatusHistoryStatusOnHold, PayoutV1DataStatusHistoryStatusPending, PayoutV1DataStatusHistoryStatusPaid, PayoutV1DataStatusHistoryStatusReversed:
		return true
	}
	return false
}

// The payment rail used for the payout.
type PayoutV1DataPaymentRail string

const (
	PayoutV1DataPaymentRailACH PayoutV1DataPaymentRail = "ach"
)

func (r PayoutV1DataPaymentRail) IsKnown() bool {
	switch r {
	case PayoutV1DataPaymentRailACH:
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
type PayoutV1ResponseType string

const (
	PayoutV1ResponseTypeObject PayoutV1ResponseType = "object"
	PayoutV1ResponseTypeArray  PayoutV1ResponseType = "array"
	PayoutV1ResponseTypeError  PayoutV1ResponseType = "error"
	PayoutV1ResponseTypeNone   PayoutV1ResponseType = "none"
)

func (r PayoutV1ResponseType) IsKnown() bool {
	switch r {
	case PayoutV1ResponseTypeObject, PayoutV1ResponseTypeArray, PayoutV1ResponseTypeError, PayoutV1ResponseTypeNone:
		return true
	}
	return false
}

type PayoutNewParams struct {
	// The amount of the payout in cents.
	Amount param.Field[int64] `json:"amount,required"`
	// The currency of the payout. Only USD is supported.
	Currency param.Field[string] `json:"currency,required"`
	// An arbitrary description for the payout.
	Description param.Field[string] `json:"description,required"`
	// Information about the device used when the customer authorized the payout.
	Device param.Field[shared.DeviceInfoV1Param] `json:"device,required"`
	// Unique identifier for the payout in your database. This value must be unique
	// across all payouts.
	ExternalID param.Field[string] `json:"external_id,required"`
	// Value of the `paykey` used for the payout.
	Paykey param.Field[string] `json:"paykey,required"`
	// The desired date on which the payout should be occur. For payouts, this means
	// the date you want the funds to be sent from your bank account.
	PaymentDate param.Field[time.Time]   `json:"payment_date,required" format:"date"`
	Config      param.Field[interface{}] `json:"config"`
	// Up to 20 additional user-defined key-value pairs. Useful for storing additional
	// information about the payout in a structured format.
	Metadata          param.Field[map[string]string] `json:"metadata"`
	CorrelationID     param.Field[string]            `header:"Correlation-Id"`
	RequestID         param.Field[string]            `header:"Request-Id"`
	StraddleAccountID param.Field[string]            `header:"Straddle-Account-Id" format:"uuid"`
}

func (r PayoutNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PayoutUpdateParams struct {
	// The amount of the payout in cents.
	Amount param.Field[int64] `json:"amount,required"`
	// An arbitrary description for the payout.
	Description param.Field[string] `json:"description,required"`
	// The desired date on which the payment should be occur. For payouts, this means
	// the date you want the funds to be sent from your bank account.
	PaymentDate param.Field[time.Time] `json:"payment_date,required" format:"date"`
	// Up to 20 additional user-defined key-value pairs. Useful for storing additional
	// information about the payout in a structured format.
	Metadata          param.Field[map[string]string] `json:"metadata"`
	CorrelationID     param.Field[string]            `header:"Correlation-Id"`
	RequestID         param.Field[string]            `header:"Request-Id"`
	StraddleAccountID param.Field[string]            `header:"Straddle-Account-Id" format:"uuid"`
}

func (r PayoutUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PayoutCancelParams struct {
	// Details about why the payout status was updated.
	Reason            param.Field[string] `json:"reason,required"`
	CorrelationID     param.Field[string] `header:"Correlation-Id"`
	RequestID         param.Field[string] `header:"Request-Id"`
	StraddleAccountID param.Field[string] `header:"Straddle-Account-Id" format:"uuid"`
}

func (r PayoutCancelParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PayoutGetParams struct {
	CorrelationID     param.Field[string] `header:"Correlation-Id"`
	RequestID         param.Field[string] `header:"Request-Id"`
	StraddleAccountID param.Field[string] `header:"Straddle-Account-Id" format:"uuid"`
}

type PayoutHoldParams struct {
	// Details about why the payout status was updated.
	Reason            param.Field[string] `json:"reason,required"`
	CorrelationID     param.Field[string] `header:"Correlation-Id"`
	RequestID         param.Field[string] `header:"Request-Id"`
	StraddleAccountID param.Field[string] `header:"Straddle-Account-Id" format:"uuid"`
}

func (r PayoutHoldParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PayoutReleaseParams struct {
	// Details about why the payout status was updated.
	Reason            param.Field[string] `json:"reason,required"`
	CorrelationID     param.Field[string] `header:"Correlation-Id"`
	RequestID         param.Field[string] `header:"Request-Id"`
	StraddleAccountID param.Field[string] `header:"Straddle-Account-Id" format:"uuid"`
}

func (r PayoutReleaseParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
