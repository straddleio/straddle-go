// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package straddle

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/stainless-sdks/straddle-go/internal/apijson"
	"github.com/stainless-sdks/straddle-go/internal/apiquery"
	"github.com/stainless-sdks/straddle-go/internal/param"
	"github.com/stainless-sdks/straddle-go/internal/requestconfig"
	"github.com/stainless-sdks/straddle-go/option"
	"github.com/stainless-sdks/straddle-go/packages/pagination"
	"github.com/stainless-sdks/straddle-go/shared"
)

// EmbedOrganizationService contains methods and other services that help with
// interacting with the straddle API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEmbedOrganizationService] method instead.
type EmbedOrganizationService struct {
	Options []option.RequestOption
}

// NewEmbedOrganizationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewEmbedOrganizationService(opts ...option.RequestOption) (r *EmbedOrganizationService) {
	r = &EmbedOrganizationService{}
	r.Options = opts
	return
}

// Creates a new organization related to your Straddle integration. Organizations
// can be used to group related accounts and manage permissions across multiple
// users.
func (r *EmbedOrganizationService) New(ctx context.Context, params EmbedOrganizationNewParams, opts ...option.RequestOption) (res *OrganizationV1, err error) {
	if params.CorrelationID.Present {
		opts = append(opts, option.WithHeader("correlation-id", fmt.Sprintf("%s", params.CorrelationID)))
	}
	if params.RequestID.Present {
		opts = append(opts, option.WithHeader("request-id", fmt.Sprintf("%s", params.RequestID)))
	}
	opts = append(r.Options[:], opts...)
	path := "v1/organizations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Retrieves a list of organizations associated with your Straddle integration. The
// organizations are returned sorted by creation date, with the most recently
// created organizations appearing first. This endpoint supports advanced sorting
// and filtering options to help you find specific organizations.
func (r *EmbedOrganizationService) List(ctx context.Context, params EmbedOrganizationListParams, opts ...option.RequestOption) (res *pagination.PageNumberSchema[OrganizationPagedV1Data], err error) {
	var raw *http.Response
	if params.CorrelationID.Present {
		opts = append(opts, option.WithHeader("correlation-id", fmt.Sprintf("%s", params.CorrelationID)))
	}
	if params.RequestID.Present {
		opts = append(opts, option.WithHeader("request-id", fmt.Sprintf("%s", params.RequestID)))
	}
	opts = append(r.Options[:], opts...)
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
	if query.CorrelationID.Present {
		opts = append(opts, option.WithHeader("correlation-id", fmt.Sprintf("%s", query.CorrelationID)))
	}
	if query.RequestID.Present {
		opts = append(opts, option.WithHeader("request-id", fmt.Sprintf("%s", query.RequestID)))
	}
	opts = append(r.Options[:], opts...)
	if organizationID == "" {
		err = errors.New("missing required organization_id parameter")
		return
	}
	path := fmt.Sprintf("v1/organizations/%s", organizationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type OrganizationPagedV1 struct {
	Data []OrganizationPagedV1Data `json:"data,required"`
	// Metadata about the API request, including an identifier, timestamp, and
	// pagination details.
	Meta shared.PagedResponseMetadata `json:"meta,required"`
	// Indicates the structure of the returned content.
	//
	//   - "object" means the `data` field contains a single JSON object.
	//   - "array" means the `data` field contains an array of objects.
	//   - "error" means the `data` field contains an error object with details of the
	//     issue.
	//   - "none" means no data is returned.
	ResponseType OrganizationPagedV1ResponseType `json:"response_type,required"`
	JSON         organizationPagedV1JSON         `json:"-"`
}

// organizationPagedV1JSON contains the JSON metadata for the struct
// [OrganizationPagedV1]
type organizationPagedV1JSON struct {
	Data         apijson.Field
	Meta         apijson.Field
	ResponseType apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *OrganizationPagedV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationPagedV1JSON) RawJSON() string {
	return r.raw
}

type OrganizationPagedV1Data struct {
	// Straddle's unique identifier for the organization.
	ID string `json:"id,required" format:"uuid"`
	// Timestamp of when the organization was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The name of the organization.
	Name string `json:"name,required"`
	// Timestamp of the most recent update to the organization.
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// Unique identifier for the organization in your database, used for
	// cross-referencing between Straddle and your systems.
	ExternalID string `json:"external_id,nullable"`
	// Up to 20 additional user-defined key-value pairs. Useful for storing additional
	// information about the organization in a structured format.
	Metadata map[string]string           `json:"metadata,nullable"`
	JSON     organizationPagedV1DataJSON `json:"-"`
}

// organizationPagedV1DataJSON contains the JSON metadata for the struct
// [OrganizationPagedV1Data]
type organizationPagedV1DataJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Name        apijson.Field
	UpdatedAt   apijson.Field
	ExternalID  apijson.Field
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationPagedV1Data) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationPagedV1DataJSON) RawJSON() string {
	return r.raw
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

func (r OrganizationPagedV1ResponseType) IsKnown() bool {
	switch r {
	case OrganizationPagedV1ResponseTypeObject, OrganizationPagedV1ResponseTypeArray, OrganizationPagedV1ResponseTypeError, OrganizationPagedV1ResponseTypeNone:
		return true
	}
	return false
}

type OrganizationV1 struct {
	Data OrganizationV1Data `json:"data,required"`
	// Metadata about the API request, including an identifier and timestamp.
	Meta shared.ResponseMetadata `json:"meta,required"`
	// Indicates the structure of the returned content.
	//
	//   - "object" means the `data` field contains a single JSON object.
	//   - "array" means the `data` field contains an array of objects.
	//   - "error" means the `data` field contains an error object with details of the
	//     issue.
	//   - "none" means no data is returned.
	ResponseType OrganizationV1ResponseType `json:"response_type,required"`
	JSON         organizationV1JSON         `json:"-"`
}

// organizationV1JSON contains the JSON metadata for the struct [OrganizationV1]
type organizationV1JSON struct {
	Data         apijson.Field
	Meta         apijson.Field
	ResponseType apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *OrganizationV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationV1JSON) RawJSON() string {
	return r.raw
}

type OrganizationV1Data struct {
	// Straddle's unique identifier for the organization.
	ID string `json:"id,required" format:"uuid"`
	// Timestamp of when the organization was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The name of the organization.
	Name string `json:"name,required"`
	// Timestamp of the most recent update to the organization.
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// Unique identifier for the organization in your database, used for
	// cross-referencing between Straddle and your systems.
	ExternalID string `json:"external_id,nullable"`
	// Up to 20 additional user-defined key-value pairs. Useful for storing additional
	// information about the organization in a structured format.
	Metadata map[string]string      `json:"metadata,nullable"`
	JSON     organizationV1DataJSON `json:"-"`
}

// organizationV1DataJSON contains the JSON metadata for the struct
// [OrganizationV1Data]
type organizationV1DataJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Name        apijson.Field
	UpdatedAt   apijson.Field
	ExternalID  apijson.Field
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationV1Data) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationV1DataJSON) RawJSON() string {
	return r.raw
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

func (r OrganizationV1ResponseType) IsKnown() bool {
	switch r {
	case OrganizationV1ResponseTypeObject, OrganizationV1ResponseTypeArray, OrganizationV1ResponseTypeError, OrganizationV1ResponseTypeNone:
		return true
	}
	return false
}

type EmbedOrganizationNewParams struct {
	// The name of the organization.
	Name param.Field[string] `json:"name,required"`
	// Unique identifier for the organization in your database, used for
	// cross-referencing between Straddle and your systems.
	ExternalID param.Field[string] `json:"external_id"`
	// Up to 20 additional user-defined key-value pairs. Useful for storing additional
	// information about the organization in a structured format.
	Metadata      param.Field[map[string]string] `json:"metadata"`
	CorrelationID param.Field[string]            `header:"correlation-id"`
	RequestID     param.Field[string]            `header:"request-id"`
}

func (r EmbedOrganizationNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EmbedOrganizationListParams struct {
	// List organizations by their external ID.
	ExternalID param.Field[string] `query:"external_id"`
	// List organizations by name (partial match supported).
	Name param.Field[string] `query:"name"`
	// Results page number. Starts at page 1.
	PageNumber param.Field[int64] `query:"page_number"`
	// Page size. Max value: 1000
	PageSize param.Field[int64] `query:"page_size"`
	// Sort By.
	SortBy param.Field[string] `query:"sort_by"`
	// Sort Order.
	SortOrder     param.Field[EmbedOrganizationListParamsSortOrder] `query:"sort_order"`
	CorrelationID param.Field[string]                               `header:"correlation-id"`
	RequestID     param.Field[string]                               `header:"request-id"`
}

// URLQuery serializes [EmbedOrganizationListParams]'s query parameters as
// `url.Values`.
func (r EmbedOrganizationListParams) URLQuery() (v url.Values) {
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

func (r EmbedOrganizationListParamsSortOrder) IsKnown() bool {
	switch r {
	case EmbedOrganizationListParamsSortOrderAsc, EmbedOrganizationListParamsSortOrderDesc:
		return true
	}
	return false
}

type EmbedOrganizationGetParams struct {
	CorrelationID param.Field[string] `header:"correlation-id"`
	RequestID     param.Field[string] `header:"request-id"`
}
