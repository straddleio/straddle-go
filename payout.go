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
func (r *PayoutService) New(ctx context.Context, params PayoutNewParams, opts ...option.RequestOption) (res *shared.PayoutV1ItemResponse, err error) {
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
func (r *PayoutService) Update(ctx context.Context, id string, params PayoutUpdateParams, opts ...option.RequestOption) (res *shared.PayoutV1ItemResponse, err error) {
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
func (r *PayoutService) Cancel(ctx context.Context, id string, params PayoutCancelParams, opts ...option.RequestOption) (res *shared.PayoutV1ItemResponse, err error) {
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
func (r *PayoutService) Get(ctx context.Context, id string, query PayoutGetParams, opts ...option.RequestOption) (res *shared.PayoutV1ItemResponse, err error) {
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
func (r *PayoutService) Hold(ctx context.Context, id string, params PayoutHoldParams, opts ...option.RequestOption) (res *shared.PayoutV1ItemResponse, err error) {
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
func (r *PayoutService) Release(ctx context.Context, id string, params PayoutReleaseParams, opts ...option.RequestOption) (res *shared.PayoutV1ItemResponse, err error) {
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
	UpdatePayoutStatusV1Request shared.UpdatePayoutStatusV1RequestParam `json:"update_payout_status_v1_request,required"`
	CorrelationID               param.Field[string]                     `header:"Correlation-Id"`
	RequestID                   param.Field[string]                     `header:"Request-Id"`
	StraddleAccountID           param.Field[string]                     `header:"Straddle-Account-Id" format:"uuid"`
}

func (r PayoutCancelParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.UpdatePayoutStatusV1Request)
}

type PayoutGetParams struct {
	CorrelationID     param.Field[string] `header:"Correlation-Id"`
	RequestID         param.Field[string] `header:"Request-Id"`
	StraddleAccountID param.Field[string] `header:"Straddle-Account-Id" format:"uuid"`
}

type PayoutHoldParams struct {
	UpdatePayoutStatusV1Request shared.UpdatePayoutStatusV1RequestParam `json:"update_payout_status_v1_request,required"`
	CorrelationID               param.Field[string]                     `header:"Correlation-Id"`
	RequestID                   param.Field[string]                     `header:"Request-Id"`
	StraddleAccountID           param.Field[string]                     `header:"Straddle-Account-Id" format:"uuid"`
}

func (r PayoutHoldParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.UpdatePayoutStatusV1Request)
}

type PayoutReleaseParams struct {
	UpdatePayoutStatusV1Request shared.UpdatePayoutStatusV1RequestParam `json:"update_payout_status_v1_request,required"`
	CorrelationID               param.Field[string]                     `header:"Correlation-Id"`
	RequestID                   param.Field[string]                     `header:"Request-Id"`
	StraddleAccountID           param.Field[string]                     `header:"Straddle-Account-Id" format:"uuid"`
}

func (r PayoutReleaseParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.UpdatePayoutStatusV1Request)
}
