// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package straddle

import (
	"context"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/straddle-go/internal/apijson"
	"github.com/stainless-sdks/straddle-go/internal/param"
	"github.com/stainless-sdks/straddle-go/internal/requestconfig"
	"github.com/stainless-sdks/straddle-go/option"
	"github.com/stainless-sdks/straddle-go/shared"
)

// BridgeService contains methods and other services that help with interacting
// with the straddle API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBridgeService] method instead.
type BridgeService struct {
	Options []option.RequestOption
	Link    *BridgeLinkService
}

// NewBridgeService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewBridgeService(opts ...option.RequestOption) (r *BridgeService) {
	r = &BridgeService{}
	r.Options = opts
	r.Link = NewBridgeLinkService(opts...)
	return
}

// Use this endpoint to generate a session token for use in the Bridge widget.
func (r *BridgeService) Initialize(ctx context.Context, params BridgeInitializeParams, opts ...option.RequestOption) (res *BridgeTokenV1, err error) {
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
	path := "v1/bridge/initialize"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type BridgeTokenV1 struct {
	Data BridgeTokenV1Data `json:"data,required"`
	// Metadata about the API request, including an identifier and timestamp.
	Meta shared.ResponseMetadata `json:"meta,required"`
	// Indicates the structure of the returned content.
	//
	//   - "object" means the `data` field contains a single JSON object.
	//   - "array" means the `data` field contains an array of objects.
	//   - "error" means the `data` field contains an error object with details of the
	//     issue.
	//   - "none" means no data is returned.
	ResponseType BridgeTokenV1ResponseType `json:"response_type,required"`
	JSON         bridgeTokenV1JSON         `json:"-"`
}

// bridgeTokenV1JSON contains the JSON metadata for the struct [BridgeTokenV1]
type bridgeTokenV1JSON struct {
	Data         apijson.Field
	Meta         apijson.Field
	ResponseType apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *BridgeTokenV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bridgeTokenV1JSON) RawJSON() string {
	return r.raw
}

type BridgeTokenV1Data struct {
	// JWT Token to use in the bridge widget.
	BridgeToken string                `json:"bridge_token,required"`
	JSON        bridgeTokenV1DataJSON `json:"-"`
}

// bridgeTokenV1DataJSON contains the JSON metadata for the struct
// [BridgeTokenV1Data]
type bridgeTokenV1DataJSON struct {
	BridgeToken apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BridgeTokenV1Data) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bridgeTokenV1DataJSON) RawJSON() string {
	return r.raw
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

func (r BridgeTokenV1ResponseType) IsKnown() bool {
	switch r {
	case BridgeTokenV1ResponseTypeObject, BridgeTokenV1ResponseTypeArray, BridgeTokenV1ResponseTypeError, BridgeTokenV1ResponseTypeNone:
		return true
	}
	return false
}

type BridgeInitializeParams struct {
	// The Straddle generated unique identifier of the `customer` to create a bridge
	// token for.
	CustomerID        param.Field[string] `json:"customer_id,required" format:"uuid"`
	CorrelationID     param.Field[string] `header:"Correlation-Id"`
	RequestID         param.Field[string] `header:"Request-Id"`
	StraddleAccountID param.Field[string] `header:"Straddle-Account-Id" format:"uuid"`
}

func (r BridgeInitializeParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
