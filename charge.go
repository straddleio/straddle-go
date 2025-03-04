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
func (r *ChargeService) New(ctx context.Context, params ChargeNewParams, opts ...option.RequestOption) (res *shared.ChargeV1ItemResponse, err error) {
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
func (r *ChargeService) Update(ctx context.Context, id string, params ChargeUpdateParams, opts ...option.RequestOption) (res *shared.ChargeV1ItemResponse, err error) {
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
func (r *ChargeService) Cancel(ctx context.Context, id string, params ChargeCancelParams, opts ...option.RequestOption) (res *shared.ChargeV1ItemResponse, err error) {
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
func (r *ChargeService) Get(ctx context.Context, id string, query ChargeGetParams, opts ...option.RequestOption) (res *shared.ChargeV1ItemResponse, err error) {
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
func (r *ChargeService) Hold(ctx context.Context, id string, params ChargeHoldParams, opts ...option.RequestOption) (res *shared.ChargeV1ItemResponse, err error) {
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
func (r *ChargeService) Release(ctx context.Context, id string, params ChargeReleaseParams, opts ...option.RequestOption) (res *shared.ChargeV1ItemResponse, err error) {
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

type ChargeNewParams struct {
	// The amount of the charge in cents.
	Amount param.Field[int64]                             `json:"amount,required"`
	Config param.Field[shared.ChargeConfigurationV1Param] `json:"config,required"`
	// The channel or mechanism through which the payment was authorized. Use
	// `internet` for payments made online or through a mobile app and `signed` for
	// signed agreements where there is a consent form or contract. Use `signed` for
	// PDF signatures.
	ConsentType param.Field[shared.ConsentTypeV1] `json:"consent_type,required"`
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
	UpdateChargeStatusV1Request shared.UpdateChargeStatusV1RequestParam `json:"update_charge_status_v1_request"`
	CorrelationID               param.Field[string]                     `header:"Correlation-Id"`
	RequestID                   param.Field[string]                     `header:"Request-Id"`
	StraddleAccountID           param.Field[string]                     `header:"Straddle-Account-Id" format:"uuid"`
}

func (r ChargeCancelParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.UpdateChargeStatusV1Request)
}

type ChargeGetParams struct {
	CorrelationID     param.Field[string] `header:"Correlation-Id"`
	RequestID         param.Field[string] `header:"Request-Id"`
	StraddleAccountID param.Field[string] `header:"Straddle-Account-Id" format:"uuid"`
}

type ChargeHoldParams struct {
	UpdateChargeStatusV1Request shared.UpdateChargeStatusV1RequestParam `json:"update_charge_status_v1_request"`
	CorrelationID               param.Field[string]                     `header:"Correlation-Id"`
	RequestID                   param.Field[string]                     `header:"Request-Id"`
	StraddleAccountID           param.Field[string]                     `header:"Straddle-Account-Id" format:"uuid"`
}

func (r ChargeHoldParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.UpdateChargeStatusV1Request)
}

type ChargeReleaseParams struct {
	UpdateChargeStatusV1Request shared.UpdateChargeStatusV1RequestParam `json:"update_charge_status_v1_request"`
	CorrelationID               param.Field[string]                     `header:"Correlation-Id"`
	RequestID                   param.Field[string]                     `header:"Request-Id"`
	StraddleAccountID           param.Field[string]                     `header:"Straddle-Account-Id" format:"uuid"`
}

func (r ChargeReleaseParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.UpdateChargeStatusV1Request)
}
