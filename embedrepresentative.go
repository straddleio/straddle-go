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

// EmbedRepresentativeService contains methods and other services that help with
// interacting with the straddle API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEmbedRepresentativeService] method instead.
type EmbedRepresentativeService struct {
	Options []option.RequestOption
}

// NewEmbedRepresentativeService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewEmbedRepresentativeService(opts ...option.RequestOption) (r *EmbedRepresentativeService) {
	r = &EmbedRepresentativeService{}
	r.Options = opts
	return
}

// Creates a new representative associated with an account. Representatives are
// individuals who have legal authority or significant responsibility within the
// business.
func (r *EmbedRepresentativeService) New(ctx context.Context, params EmbedRepresentativeNewParams, opts ...option.RequestOption) (res *shared.ItemResponseOfRepresentativeV1, err error) {
	if params.CorrelationID.Present {
		opts = append(opts, option.WithHeader("correlation-id", fmt.Sprintf("%s", params.CorrelationID)))
	}
	if params.RequestID.Present {
		opts = append(opts, option.WithHeader("request-id", fmt.Sprintf("%s", params.RequestID)))
	}
	opts = append(r.Options[:], opts...)
	path := "v1/representatives"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Updates an existing representative's information. This can be used to update
// personal details, contact information, or the relationship to the account or
// organization.
func (r *EmbedRepresentativeService) Update(ctx context.Context, representativeID string, params EmbedRepresentativeUpdateParams, opts ...option.RequestOption) (res *shared.ItemResponseOfRepresentativeV1, err error) {
	if params.CorrelationID.Present {
		opts = append(opts, option.WithHeader("correlation-id", fmt.Sprintf("%s", params.CorrelationID)))
	}
	if params.RequestID.Present {
		opts = append(opts, option.WithHeader("request-id", fmt.Sprintf("%s", params.RequestID)))
	}
	opts = append(r.Options[:], opts...)
	if representativeID == "" {
		err = errors.New("missing required representative_id parameter")
		return
	}
	path := fmt.Sprintf("v1/representatives/%s", representativeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

// Returns a list of representatives associated with a specific account or
// organization. The representatives are returned sorted by creation date, with the
// most recently created representatives appearing first. This endpoint supports
// advanced sorting and filtering options.
func (r *EmbedRepresentativeService) List(ctx context.Context, params EmbedRepresentativeListParams, opts ...option.RequestOption) (res *pagination.PageNumberSchema[shared.RepresentativeV1], err error) {
	var raw *http.Response
	if params.CorrelationID.Present {
		opts = append(opts, option.WithHeader("correlation-id", fmt.Sprintf("%s", params.CorrelationID)))
	}
	if params.RequestID.Present {
		opts = append(opts, option.WithHeader("request-id", fmt.Sprintf("%s", params.RequestID)))
	}
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/representatives"
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

// Returns a list of representatives associated with a specific account or
// organization. The representatives are returned sorted by creation date, with the
// most recently created representatives appearing first. This endpoint supports
// advanced sorting and filtering options.
func (r *EmbedRepresentativeService) ListAutoPaging(ctx context.Context, params EmbedRepresentativeListParams, opts ...option.RequestOption) *pagination.PageNumberSchemaAutoPager[shared.RepresentativeV1] {
	return pagination.NewPageNumberSchemaAutoPager(r.List(ctx, params, opts...))
}

// Retrieves the details of an existing representative. Supply the unique
// representative ID, and Straddle will return the corresponding representative
// information.
func (r *EmbedRepresentativeService) Get(ctx context.Context, representativeID string, query EmbedRepresentativeGetParams, opts ...option.RequestOption) (res *shared.ItemResponseOfRepresentativeV1, err error) {
	if query.CorrelationID.Present {
		opts = append(opts, option.WithHeader("correlation-id", fmt.Sprintf("%s", query.CorrelationID)))
	}
	if query.RequestID.Present {
		opts = append(opts, option.WithHeader("request-id", fmt.Sprintf("%s", query.RequestID)))
	}
	opts = append(r.Options[:], opts...)
	if representativeID == "" {
		err = errors.New("missing required representative_id parameter")
		return
	}
	path := fmt.Sprintf("v1/representatives/%s", representativeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Retrieves the unmasked details of a representative that has previously been
// created. Supply the unique representative ID, and Straddle will return the
// corresponding representative information, including sensitive details. This
// endpoint requires additional authentication and should be used with caution.
func (r *EmbedRepresentativeService) Unmask(ctx context.Context, representativeID string, query EmbedRepresentativeUnmaskParams, opts ...option.RequestOption) (res *shared.ItemResponseOfRepresentativeV1, err error) {
	if query.CorrelationID.Present {
		opts = append(opts, option.WithHeader("correlation-id", fmt.Sprintf("%s", query.CorrelationID)))
	}
	if query.RequestID.Present {
		opts = append(opts, option.WithHeader("request-id", fmt.Sprintf("%s", query.RequestID)))
	}
	opts = append(r.Options[:], opts...)
	if representativeID == "" {
		err = errors.New("missing required representative_id parameter")
		return
	}
	path := fmt.Sprintf("v1/representatives/%s/unmask", representativeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type RepresentativePaged struct {
	Data []shared.RepresentativeV1 `json:"data,required"`
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
	ResponseType RepresentativePagedResponseType `json:"response_type,required"`
	JSON         representativePagedJSON         `json:"-"`
}

// representativePagedJSON contains the JSON metadata for the struct
// [RepresentativePaged]
type representativePagedJSON struct {
	Data         apijson.Field
	Meta         apijson.Field
	ResponseType apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RepresentativePaged) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r representativePagedJSON) RawJSON() string {
	return r.raw
}

// Indicates the structure of the returned content.
//
//   - "object" means the `data` field contains a single JSON object.
//   - "array" means the `data` field contains an array of objects.
//   - "error" means the `data` field contains an error object with details of the
//     issue.
//   - "none" means no data is returned.
type RepresentativePagedResponseType string

const (
	RepresentativePagedResponseTypeObject RepresentativePagedResponseType = "object"
	RepresentativePagedResponseTypeArray  RepresentativePagedResponseType = "array"
	RepresentativePagedResponseTypeError  RepresentativePagedResponseType = "error"
	RepresentativePagedResponseTypeNone   RepresentativePagedResponseType = "none"
)

func (r RepresentativePagedResponseType) IsKnown() bool {
	switch r {
	case RepresentativePagedResponseTypeObject, RepresentativePagedResponseTypeArray, RepresentativePagedResponseTypeError, RepresentativePagedResponseTypeNone:
		return true
	}
	return false
}

type EmbedRepresentativeNewParams struct {
	// The unique identifier of the account this representative is associated with.
	AccountID param.Field[string] `json:"account_id,required" format:"uuid"`
	// Date of birth for the representative in ISO 8601 format (YYYY-MM-DD).
	Dob param.Field[time.Time] `json:"dob,required" format:"date"`
	// The company email address of the representative.
	Email param.Field[string] `json:"email,required" format:"email"`
	// The first name of the representative.
	FirstName param.Field[string] `json:"first_name,required"`
	// The last name of the representative.
	LastName param.Field[string] `json:"last_name,required"`
	// The mobile phone number of the representative.
	MobileNumber param.Field[string]                     `json:"mobile_number,required"`
	Relationship param.Field[shared.RelationshipV1Param] `json:"relationship,required"`
	// The last 4 digits of the representative's Social Security Number.
	SsnLast4 param.Field[string] `json:"ssn_last4,required"`
	// Unique identifier for the representative in your database, used for
	// cross-referencing between Straddle and your systems.
	ExternalID    param.Field[string] `json:"external_id"`
	CorrelationID param.Field[string] `header:"correlation-id"`
	RequestID     param.Field[string] `header:"request-id"`
}

func (r EmbedRepresentativeNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EmbedRepresentativeUpdateParams struct {
	// The date of birth of the representative, in ISO 8601 format (YYYY-MM-DD).
	Dob param.Field[time.Time] `json:"dob,required" format:"date"`
	// The email address of the representative.
	Email param.Field[string] `json:"email,required" format:"email"`
	// The first name of the representative.
	FirstName param.Field[string] `json:"first_name,required"`
	// The last name of the representative.
	LastName param.Field[string] `json:"last_name,required"`
	// The mobile phone number of the representative.
	MobileNumber param.Field[string]                     `json:"mobile_number,required"`
	Relationship param.Field[shared.RelationshipV1Param] `json:"relationship,required"`
	// The last 4 digits of the representative's Social Security Number.
	SsnLast4 param.Field[string] `json:"ssn_last4,required"`
	// Unique identifier for the representative in your database, used for
	// cross-referencing between Straddle and your systems.
	ExternalID    param.Field[string] `json:"external_id"`
	CorrelationID param.Field[string] `header:"correlation-id"`
	RequestID     param.Field[string] `header:"request-id"`
}

func (r EmbedRepresentativeUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EmbedRepresentativeListParams struct {
	// The unique identifier of the account to list representatives for.
	AccountID      param.Field[string] `query:"account_id" format:"uuid"`
	OrganizationID param.Field[string] `query:"organization_id" format:"uuid"`
	// Results page number. Starts at page 1.
	PageNumber param.Field[int64] `query:"page_number"`
	// Page size. Max value: 1000
	PageSize   param.Field[int64]  `query:"page_size"`
	PlatformID param.Field[string] `query:"platform_id" format:"uuid"`
	// Sort By.
	SortBy param.Field[string] `query:"sort_by"`
	// Sort Order.
	SortOrder     param.Field[EmbedRepresentativeListParamsSortOrder] `query:"sort_order"`
	CorrelationID param.Field[string]                                 `header:"correlation-id"`
	RequestID     param.Field[string]                                 `header:"request-id"`
}

// URLQuery serializes [EmbedRepresentativeListParams]'s query parameters as
// `url.Values`.
func (r EmbedRepresentativeListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Sort Order.
type EmbedRepresentativeListParamsSortOrder string

const (
	EmbedRepresentativeListParamsSortOrderAsc  EmbedRepresentativeListParamsSortOrder = "asc"
	EmbedRepresentativeListParamsSortOrderDesc EmbedRepresentativeListParamsSortOrder = "desc"
)

func (r EmbedRepresentativeListParamsSortOrder) IsKnown() bool {
	switch r {
	case EmbedRepresentativeListParamsSortOrderAsc, EmbedRepresentativeListParamsSortOrderDesc:
		return true
	}
	return false
}

type EmbedRepresentativeGetParams struct {
	CorrelationID param.Field[string] `header:"correlation-id"`
	RequestID     param.Field[string] `header:"request-id"`
}

type EmbedRepresentativeUnmaskParams struct {
	CorrelationID param.Field[string] `header:"correlation-id"`
	RequestID     param.Field[string] `header:"request-id"`
}
