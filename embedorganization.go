// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package straddle

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/straddleio/straddle-go/internal/apijson"
	"github.com/straddleio/straddle-go/internal/apiquery"
	"github.com/straddleio/straddle-go/internal/requestconfig"
	"github.com/straddleio/straddle-go/option"
	"github.com/straddleio/straddle-go/packages/pagination"
	"github.com/straddleio/straddle-go/packages/param"
	"github.com/straddleio/straddle-go/packages/respjson"
	"github.com/straddleio/straddle-go/shared"
)

// Organizations are a powerful feature in Straddle that allow you to manage
// multiple accounts under a single umbrella. This hierarchical structure is
// particularly useful for businesses with complex operations, multiple
// departments, or legally related entities.
//
// EmbedOrganizationService contains methods and other services that help with
// interacting with the straddle API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEmbedOrganizationService] method instead.
type EmbedOrganizationService struct {
	options []option.RequestOption
}

// NewEmbedOrganizationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewEmbedOrganizationService(opts ...option.RequestOption) (r EmbedOrganizationService) {
	r = EmbedOrganizationService{}
	r.options = opts
	return
}

// Creates a new organization related to your Straddle integration. Organizations
// can be used to group related accounts and manage permissions across multiple
// users.
func (r *EmbedOrganizationService) New(ctx context.Context, params EmbedOrganizationNewParams, opts ...option.RequestOption) (res *OrganizationV1, err error) {
	if !param.IsOmitted(params.CorrelationID) {
		opts = append(opts, option.WithHeader("correlation-id", fmt.Sprintf("%v", params.CorrelationID.Value)))
	}
	if !param.IsOmitted(params.IdempotencyKey) {
		opts = append(opts, option.WithHeader("idempotency-key", fmt.Sprintf("%v", params.IdempotencyKey.Value)))
	}
	if !param.IsOmitted(params.RequestID) {
		opts = append(opts, option.WithHeader("request-id", fmt.Sprintf("%v", params.RequestID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	path := "v1/organizations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Retrieves a list of organizations associated with your Straddle integration. The
// organizations are returned sorted by creation date, with the most recently
// created organizations appearing first. This endpoint supports advanced sorting
// and filtering options to help you find specific organizations.
func (r *EmbedOrganizationService) List(ctx context.Context, params EmbedOrganizationListParams, opts ...option.RequestOption) (res *pagination.PageNumberSchema[OrganizationPagedV1Data], err error) {
	var raw *http.Response
	if !param.IsOmitted(params.CorrelationID) {
		opts = append(opts, option.WithHeader("correlation-id", fmt.Sprintf("%v", params.CorrelationID.Value)))
	}
	if !param.IsOmitted(params.RequestID) {
		opts = append(opts, option.WithHeader("request-id", fmt.Sprintf("%v", params.RequestID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/organizations"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, params, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Retrieves a list of organizations associated with your Straddle integration. The
// organizations are returned sorted by creation date, with the most recently
// created organizations appearing first. This endpoint supports advanced sorting
// and filtering options to help you find specific organizations.
func (r *EmbedOrganizationService) ListAutoPaging(ctx context.Context, params EmbedOrganizationListParams, opts ...option.RequestOption) *pagination.PageNumberSchemaAutoPager[OrganizationPagedV1Data] {
	return pagination.NewPageNumberSchemaAutoPager(r.List(ctx, params, opts...))
}

// Retrieves the details of an Organization that has previously been created.
// Supply the unique organization ID that was returned from your previous request,
// and Straddle will return the corresponding organization information.
func (r *EmbedOrganizationService) Get(ctx context.Context, organizationID string, query EmbedOrganizationGetParams, opts ...option.RequestOption) (res *OrganizationV1, err error) {
	if !param.IsOmitted(query.CorrelationID) {
		opts = append(opts, option.WithHeader("correlation-id", fmt.Sprintf("%v", query.CorrelationID.Value)))
	}
	if !param.IsOmitted(query.RequestID) {
		opts = append(opts, option.WithHeader("request-id", fmt.Sprintf("%v", query.RequestID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	if organizationID == "" {
		err = errors.New("missing required organization_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/organizations/%s", organizationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type OrganizationPagedV1 struct {
	Data []OrganizationPagedV1Data `json:"data" api:"required"`
	// Metadata about the API request, including an identifier, timestamp, and
	// pagination details.
	Meta shared.PagedResponseMetadata `json:"meta" api:"required"`
	// Indicates the structure of the returned content.
	//
	//   - "object" means the `data` field contains a single JSON object.
	//   - "array" means the `data` field contains an array of objects.
	//   - "error" means the `data` field contains an error object with details of the
	//     issue.
	//   - "none" means no data is returned.
	//
	// Any of "object", "array", "error", "none".
	ResponseType OrganizationPagedV1ResponseType `json:"response_type" api:"required"`
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
func (r OrganizationPagedV1) RawJSON() string { return r.JSON.raw }
func (r *OrganizationPagedV1) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OrganizationPagedV1Data struct {
	// Straddle's unique identifier for the organization.
	ID string `json:"id" api:"required" format:"uuid"`
	// Timestamp of when the organization was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// The name of the organization.
	Name string `json:"name" api:"required"`
	// Timestamp of the most recent update to the organization.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// Unique identifier for the organization in your database, used for
	// cross-referencing between Straddle and your systems.
	ExternalID string `json:"external_id" api:"nullable"`
	// Up to 20 additional user-defined key-value pairs. Useful for storing additional
	// information about the organization in a structured format.
	Metadata map[string]string `json:"metadata" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Name        respjson.Field
		UpdatedAt   respjson.Field
		ExternalID  respjson.Field
		Metadata    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OrganizationPagedV1Data) RawJSON() string { return r.JSON.raw }
func (r *OrganizationPagedV1Data) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indicates the structure of the returned content.
//
//   - "object" means the `data` field contains a single JSON object.
//   - "array" means the `data` field contains an array of objects.
//   - "error" means the `data` field contains an error object with details of the
//     issue.
//   - "none" means no data is returned.
type OrganizationPagedV1ResponseType string

const (
	OrganizationPagedV1ResponseTypeObject OrganizationPagedV1ResponseType = "object"
	OrganizationPagedV1ResponseTypeArray  OrganizationPagedV1ResponseType = "array"
	OrganizationPagedV1ResponseTypeError  OrganizationPagedV1ResponseType = "error"
	OrganizationPagedV1ResponseTypeNone   OrganizationPagedV1ResponseType = "none"
)

type OrganizationV1 struct {
	Data OrganizationV1Data `json:"data" api:"required"`
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
	ResponseType OrganizationV1ResponseType `json:"response_type" api:"required"`
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
func (r OrganizationV1) RawJSON() string { return r.JSON.raw }
func (r *OrganizationV1) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OrganizationV1Data struct {
	// Straddle's unique identifier for the organization.
	ID string `json:"id" api:"required" format:"uuid"`
	// Timestamp of when the organization was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// The name of the organization.
	Name string `json:"name" api:"required"`
	// Timestamp of the most recent update to the organization.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// Unique identifier for the organization in your database, used for
	// cross-referencing between Straddle and your systems.
	ExternalID string `json:"external_id" api:"nullable"`
	// Up to 20 additional user-defined key-value pairs. Useful for storing additional
	// information about the organization in a structured format.
	Metadata map[string]string `json:"metadata" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Name        respjson.Field
		UpdatedAt   respjson.Field
		ExternalID  respjson.Field
		Metadata    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OrganizationV1Data) RawJSON() string { return r.JSON.raw }
func (r *OrganizationV1Data) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indicates the structure of the returned content.
//
//   - "object" means the `data` field contains a single JSON object.
//   - "array" means the `data` field contains an array of objects.
//   - "error" means the `data` field contains an error object with details of the
//     issue.
//   - "none" means no data is returned.
type OrganizationV1ResponseType string

const (
	OrganizationV1ResponseTypeObject OrganizationV1ResponseType = "object"
	OrganizationV1ResponseTypeArray  OrganizationV1ResponseType = "array"
	OrganizationV1ResponseTypeError  OrganizationV1ResponseType = "error"
	OrganizationV1ResponseTypeNone   OrganizationV1ResponseType = "none"
)

type EmbedOrganizationNewParams struct {
	// The name of the organization.
	Name string `json:"name" api:"required"`
	// Unique identifier for the organization in your database, used for
	// cross-referencing between Straddle and your systems.
	ExternalID     param.Opt[string] `json:"external_id,omitzero"`
	CorrelationID  param.Opt[string] `header:"correlation-id,omitzero" json:"-"`
	IdempotencyKey param.Opt[string] `header:"idempotency-key,omitzero" json:"-"`
	RequestID      param.Opt[string] `header:"request-id,omitzero" json:"-"`
	// Up to 20 additional user-defined key-value pairs. Useful for storing additional
	// information about the organization in a structured format.
	Metadata map[string]string `json:"metadata,omitzero"`
	paramObj
}

func (r EmbedOrganizationNewParams) MarshalJSON() (data []byte, err error) {
	type shadow EmbedOrganizationNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EmbedOrganizationNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbedOrganizationListParams struct {
	// List organizations by their external ID.
	ExternalID param.Opt[string] `query:"external_id,omitzero" json:"-"`
	// List organizations by name (partial match supported).
	Name param.Opt[string] `query:"name,omitzero" json:"-"`
	// Results page number. Starts at page 1.
	PageNumber param.Opt[int64] `query:"page_number,omitzero" json:"-"`
	// Page size. Max value: 1000
	PageSize param.Opt[int64] `query:"page_size,omitzero" json:"-"`
	// Sort By.
	SortBy        param.Opt[string] `query:"sort_by,omitzero" json:"-"`
	CorrelationID param.Opt[string] `header:"correlation-id,omitzero" json:"-"`
	RequestID     param.Opt[string] `header:"request-id,omitzero" json:"-"`
	// Sort Order.
	//
	// Any of "asc", "desc".
	SortOrder EmbedOrganizationListParamsSortOrder `query:"sort_order,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [EmbedOrganizationListParams]'s query parameters as
// `url.Values`.
func (r EmbedOrganizationListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Sort Order.
type EmbedOrganizationListParamsSortOrder string

const (
	EmbedOrganizationListParamsSortOrderAsc  EmbedOrganizationListParamsSortOrder = "asc"
	EmbedOrganizationListParamsSortOrderDesc EmbedOrganizationListParamsSortOrder = "desc"
)

type EmbedOrganizationGetParams struct {
	CorrelationID param.Opt[string] `header:"correlation-id,omitzero" json:"-"`
	RequestID     param.Opt[string] `header:"request-id,omitzero" json:"-"`
	paramObj
}
