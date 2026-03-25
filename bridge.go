// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package straddle

import (
	"context"
	"fmt"
	"net/http"
	"slices"

	"github.com/stainless-sdks/straddle-go/internal/apijson"
	"github.com/stainless-sdks/straddle-go/internal/requestconfig"
	"github.com/stainless-sdks/straddle-go/option"
	"github.com/stainless-sdks/straddle-go/packages/param"
	"github.com/stainless-sdks/straddle-go/packages/respjson"
	"github.com/stainless-sdks/straddle-go/shared"
)

// Bridge provides a comprehensive suite of tools for connecting customer bank
// accounts. Use it to generate secure widget sessions for instant account
// verification, accept tokens from major providers like Plaid and Finicity, or
// verify accounts directly via our API. Bridge handles all sensitive banking
// credentials and ensures secure, compliant connections with support for 90% of US
// bank accounts.
//
// BridgeService contains methods and other services that help with interacting
// with the straddle API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBridgeService] method instead.
type BridgeService struct {
	options []option.RequestOption
	// Bridge provides a comprehensive suite of tools for connecting customer bank
	// accounts. Use it to generate secure widget sessions for instant account
	// verification, accept tokens from major providers like Plaid and Finicity, or
	// verify accounts directly via our API. Bridge handles all sensitive banking
	// credentials and ensures secure, compliant connections with support for 90% of US
	// bank accounts.
	Link BridgeLinkService
}

// NewBridgeService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewBridgeService(opts ...option.RequestOption) (r BridgeService) {
	r = BridgeService{}
	r.options = opts
	r.Link = NewBridgeLinkService(opts...)
	return
}

// Use this endpoint to generate a session token for use in the Bridge widget.
func (r *BridgeService) Initialize(ctx context.Context, params BridgeInitializeParams, opts ...option.RequestOption) (res *BridgeTokenV1, err error) {
	if !param.IsOmitted(params.CorrelationID) {
		opts = append(opts, option.WithHeader("Correlation-Id", fmt.Sprintf("%v", params.CorrelationID.Value)))
	}
	if !param.IsOmitted(params.IdempotencyKey) {
		opts = append(opts, option.WithHeader("Idempotency-Key", fmt.Sprintf("%v", params.IdempotencyKey.Value)))
	}
	if !param.IsOmitted(params.RequestID) {
		opts = append(opts, option.WithHeader("Request-Id", fmt.Sprintf("%v", params.RequestID.Value)))
	}
	if !param.IsOmitted(params.StraddleAccountID) {
		opts = append(opts, option.WithHeader("Straddle-Account-Id", fmt.Sprintf("%v", params.StraddleAccountID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	path := "v1/bridge/initialize"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

type BridgeTokenV1 struct {
	Data BridgeTokenV1Data `json:"data" api:"required"`
	// Metadata about the API request, including an identifier and timestamp.
	Meta shared.ResponseMetadata `json:"meta" api:"required"`
	// Indicates the structure of the returned content.
	//
	//   - "object" means the `data` field contains a single JSON object.
	//   - "array" means the `data` field contains an array of objects.
	//   - "error" means the `data` field contains an error object with details of the
	//     issue.
	//   - "none" means no data is returned.
	//
	// Any of "object", "array", "error", "none".
	ResponseType BridgeTokenV1ResponseType `json:"response_type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data         respjson.Field
		Meta         respjson.Field
		ResponseType respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BridgeTokenV1) RawJSON() string { return r.JSON.raw }
func (r *BridgeTokenV1) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BridgeTokenV1Data struct {
	// JWT Token to use in the bridge widget.
	BridgeToken string `json:"bridge_token" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BridgeToken respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BridgeTokenV1Data) RawJSON() string { return r.JSON.raw }
func (r *BridgeTokenV1Data) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indicates the structure of the returned content.
//
//   - "object" means the `data` field contains a single JSON object.
//   - "array" means the `data` field contains an array of objects.
//   - "error" means the `data` field contains an error object with details of the
//     issue.
//   - "none" means no data is returned.
type BridgeTokenV1ResponseType string

const (
	BridgeTokenV1ResponseTypeObject BridgeTokenV1ResponseType = "object"
	BridgeTokenV1ResponseTypeArray  BridgeTokenV1ResponseType = "array"
	BridgeTokenV1ResponseTypeError  BridgeTokenV1ResponseType = "error"
	BridgeTokenV1ResponseTypeNone   BridgeTokenV1ResponseType = "none"
)

type BridgeInitializeParams struct {
	// The Straddle generated unique identifier of the `customer` to create a bridge
	// token for.
	CustomerID string `json:"customer_id" api:"required" format:"uuid"`
	// Unique identifier for the paykey in your database, used for cross-referencing
	// between Straddle and your systems.
	ExternalID        param.Opt[string]            `json:"external_id,omitzero"`
	CorrelationID     param.Opt[string]            `header:"Correlation-Id,omitzero" json:"-"`
	IdempotencyKey    param.Opt[string]            `header:"Idempotency-Key,omitzero" json:"-"`
	RequestID         param.Opt[string]            `header:"Request-Id,omitzero" json:"-"`
	StraddleAccountID param.Opt[string]            `header:"Straddle-Account-Id,omitzero" format:"uuid" json:"-"`
	Config            BridgeInitializeParamsConfig `json:"config,omitzero"`
	paramObj
}

func (r BridgeInitializeParams) MarshalJSON() (data []byte, err error) {
	type shadow BridgeInitializeParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BridgeInitializeParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BridgeInitializeParamsConfig struct {
	// Any of "inline", "background", "skip".
	ProcessingMethod string `json:"processing_method,omitzero"`
	// Any of "standard", "active", "rejected", "review".
	SandboxOutcome string `json:"sandbox_outcome,omitzero"`
	paramObj
}

func (r BridgeInitializeParamsConfig) MarshalJSON() (data []byte, err error) {
	type shadow BridgeInitializeParamsConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BridgeInitializeParamsConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[BridgeInitializeParamsConfig](
		"processing_method", "inline", "background", "skip",
	)
	apijson.RegisterFieldValidator[BridgeInitializeParamsConfig](
		"sandbox_outcome", "standard", "active", "rejected", "review",
	)
}
