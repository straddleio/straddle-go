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

// CustomerService contains methods and other services that help with interacting
// with the straddle API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCustomerService] method instead.
type CustomerService struct {
	Options []option.RequestOption
	Review  *CustomerReviewService
}

// NewCustomerService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCustomerService(opts ...option.RequestOption) (r *CustomerService) {
	r = &CustomerService{}
	r.Options = opts
	r.Review = NewCustomerReviewService(opts...)
	return
}

// Creates a new customer record and automatically initiates identity, fraud, and
// risk assessment scores. This endpoint allows you to create a customer profile
// and associate it with paykeys and payments.
func (r *CustomerService) New(ctx context.Context, params CustomerNewParams, opts ...option.RequestOption) (res *shared.CustomerV1ItemResponse, err error) {
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
	path := "v1/customers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Updates an existing customer's information. This endpoint allows you to modify
// the customer's contact details, PII, and metadata.
func (r *CustomerService) Update(ctx context.Context, id string, params CustomerUpdateParams, opts ...option.RequestOption) (res *shared.CustomerV1ItemResponse, err error) {
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
	path := fmt.Sprintf("v1/customers/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

// Lists or searches customers connected to your account. All supported query
// parameters are optional. If none are provided, the response will include all
// customers connected to your account. This endpoint supports advanced sorting and
// filtering options.
func (r *CustomerService) List(ctx context.Context, params CustomerListParams, opts ...option.RequestOption) (res *pagination.PageNumberSchema[CustomerSummaryPagedV1Data], err error) {
	var raw *http.Response
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
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/customers"
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

// Lists or searches customers connected to your account. All supported query
// parameters are optional. If none are provided, the response will include all
// customers connected to your account. This endpoint supports advanced sorting and
// filtering options.
func (r *CustomerService) ListAutoPaging(ctx context.Context, params CustomerListParams, opts ...option.RequestOption) *pagination.PageNumberSchemaAutoPager[CustomerSummaryPagedV1Data] {
	return pagination.NewPageNumberSchemaAutoPager(r.List(ctx, params, opts...))
}

// Permanently removes a customer record from Straddle. This action cannot be
// undone and should only be used to satisfy regulatory requirements or for privacy
// compliance.
func (r *CustomerService) Delete(ctx context.Context, id string, body CustomerDeleteParams, opts ...option.RequestOption) (res *shared.CustomerV1ItemResponse, err error) {
	if body.CorrelationID.Present {
		opts = append(opts, option.WithHeader("Correlation-Id", fmt.Sprintf("%s", body.CorrelationID)))
	}
	if body.RequestID.Present {
		opts = append(opts, option.WithHeader("Request-Id", fmt.Sprintf("%s", body.RequestID)))
	}
	if body.StraddleAccountID.Present {
		opts = append(opts, option.WithHeader("Straddle-Account-Id", fmt.Sprintf("%s", body.StraddleAccountID)))
	}
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/customers/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Retrieves the details of an existing customer. Supply the unique customer ID
// that was returned from your 'create customer' request, and Straddle will return
// the corresponding customer information.
func (r *CustomerService) Get(ctx context.Context, id string, query CustomerGetParams, opts ...option.RequestOption) (res *shared.CustomerV1ItemResponse, err error) {
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
	path := fmt.Sprintf("v1/customers/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Retrieves the unmasked details, including PII, of an existing customer. Supply
// the unique customer ID that was returned from your 'create customer' request,
// and Straddle will return the corresponding customer information. This endpoint
// needs to be enabled by Straddle and should only be used when absolutely
// necessary.
func (r *CustomerService) Unmasked(ctx context.Context, id string, query CustomerUnmaskedParams, opts ...option.RequestOption) (res *CustomerUnmaskedV1, err error) {
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
	path := fmt.Sprintf("v1/customers/%s/unmasked", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type CustomerSummaryPagedV1 struct {
	Data []CustomerSummaryPagedV1Data  `json:"data,required"`
	Meta shared.PagedResponseMetadata1 `json:"meta,required"`
	// Indicates the structure of the returned content.
	//
	//   - "object" means the `data` field contains a single JSON object.
	//   - "array" means the `data` field contains an array of objects.
	//   - "error" means the `data` field contains an error object with details of the
	//     issue.
	//   - "none" means no data is returned.
	ResponseType shared.ResponseTypeEnum    `json:"response_type,required"`
	JSON         customerSummaryPagedV1JSON `json:"-"`
}

// customerSummaryPagedV1JSON contains the JSON metadata for the struct
// [CustomerSummaryPagedV1]
type customerSummaryPagedV1JSON struct {
	Data         apijson.Field
	Meta         apijson.Field
	ResponseType apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *CustomerSummaryPagedV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerSummaryPagedV1JSON) RawJSON() string {
	return r.raw
}

type CustomerSummaryPagedV1Data struct {
	// Unique identifier for the customer.
	ID string `json:"id,required" format:"uuid"`
	// Timestamp of when the customer record was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The customer's email address.
	Email string `json:"email,required" format:"email"`
	// Full name of the individual or business name.
	Name string `json:"name,required"`
	// The customer's phone number in E.164 format.
	Phone  string                  `json:"phone,required"`
	Status shared.CustomerStatusV1 `json:"status,required"`
	Type   shared.CustomerTypeV1   `json:"type,required"`
	// Timestamp of the most recent update to the customer record.
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// Unique identifier for the customer in your database, used for cross-referencing
	// between Straddle and your systems.
	ExternalID string                         `json:"external_id,nullable"`
	JSON       customerSummaryPagedV1DataJSON `json:"-"`
}

// customerSummaryPagedV1DataJSON contains the JSON metadata for the struct
// [CustomerSummaryPagedV1Data]
type customerSummaryPagedV1DataJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Email       apijson.Field
	Name        apijson.Field
	Phone       apijson.Field
	Status      apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	ExternalID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerSummaryPagedV1Data) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerSummaryPagedV1DataJSON) RawJSON() string {
	return r.raw
}

type CustomerUnmaskedV1 struct {
	Data CustomerUnmaskedV1Data `json:"data,required"`
	// Metadata about the API request, including an identifier and timestamp.
	Meta shared.ResponseMetadata `json:"meta,required"`
	// Indicates the structure of the returned content.
	//
	//   - "object" means the `data` field contains a single JSON object.
	//   - "array" means the `data` field contains an array of objects.
	//   - "error" means the `data` field contains an error object with details of the
	//     issue.
	//   - "none" means no data is returned.
	ResponseType shared.ResponseTypeEnum `json:"response_type,required"`
	JSON         customerUnmaskedV1JSON  `json:"-"`
}

// customerUnmaskedV1JSON contains the JSON metadata for the struct
// [CustomerUnmaskedV1]
type customerUnmaskedV1JSON struct {
	Data         apijson.Field
	Meta         apijson.Field
	ResponseType apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *CustomerUnmaskedV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerUnmaskedV1JSON) RawJSON() string {
	return r.raw
}

type CustomerUnmaskedV1Data struct {
	// Unique identifier for the customer.
	ID string `json:"id,required" format:"uuid"`
	// Timestamp of when the customer record was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The customer's email address.
	Email string `json:"email,required" format:"email"`
	// Full name of the individual or business name.
	Name string `json:"name,required"`
	// The customer's phone number in E.164 format.
	Phone  string                  `json:"phone,required"`
	Status shared.CustomerStatusV1 `json:"status,required"`
	Type   shared.CustomerTypeV1   `json:"type,required"`
	// Timestamp of the most recent update to the customer record.
	UpdatedAt time.Time         `json:"updated_at,required" format:"date-time"`
	Address   shared.AddressV11 `json:"address,nullable"`
	// Compliance profile for individual customers
	ComplianceProfile shared.ComplianceProfileUnmaskedV1 `json:"compliance_profile"`
	Device            shared.DeviceUnmaskedV1            `json:"device"`
	// Unique identifier for the customer in your database, used for cross-referencing
	// between Straddle and your systems.
	ExternalID string `json:"external_id,nullable"`
	// Up to 20 additional user-defined key-value pairs. Useful for storing additional
	// information about the customer in a structured format.
	Metadata map[string]string          `json:"metadata,nullable"`
	JSON     customerUnmaskedV1DataJSON `json:"-"`
}

// customerUnmaskedV1DataJSON contains the JSON metadata for the struct
// [CustomerUnmaskedV1Data]
type customerUnmaskedV1DataJSON struct {
	ID                apijson.Field
	CreatedAt         apijson.Field
	Email             apijson.Field
	Name              apijson.Field
	Phone             apijson.Field
	Status            apijson.Field
	Type              apijson.Field
	UpdatedAt         apijson.Field
	Address           apijson.Field
	ComplianceProfile apijson.Field
	Device            apijson.Field
	ExternalID        apijson.Field
	Metadata          apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *CustomerUnmaskedV1Data) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerUnmaskedV1DataJSON) RawJSON() string {
	return r.raw
}

type CustomerNewParams struct {
	Device param.Field[shared.DeviceUnmaskedV1Param] `json:"device,required"`
	// The customer's email address.
	Email param.Field[string] `json:"email,required" format:"email"`
	// Full name of the individual or business name.
	Name param.Field[string] `json:"name,required"`
	// The customer's phone number in E.164 format. Mobile number is preferred.
	Phone param.Field[string]                `json:"phone,required"`
	Type  param.Field[shared.CustomerTypeV1] `json:"type,required"`
	// An object containing the customer's address. This is optional, but if provided,
	// all required fields must be present.
	Address param.Field[shared.AddressV11Param] `json:"address"`
	// An object containing the customer's compliance profile. This is optional, but if
	// provided, all required fields must be present for the appropriate customer type.
	ComplianceProfile param.Field[shared.ComplianceProfileUnmaskedV1UnionParam] `json:"compliance_profile"`
	// Unique identifier for the customer in your database, used for cross-referencing
	// between Straddle and your systems.
	ExternalID param.Field[string] `json:"external_id"`
	// Up to 20 additional user-defined key-value pairs. Useful for storing additional
	// information about the customer in a structured format.
	Metadata          param.Field[map[string]string] `json:"metadata"`
	CorrelationID     param.Field[string]            `header:"Correlation-Id"`
	RequestID         param.Field[string]            `header:"Request-Id"`
	StraddleAccountID param.Field[string]            `header:"Straddle-Account-Id" format:"uuid"`
}

func (r CustomerNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CustomerUpdateParams struct {
	Device param.Field[shared.DeviceUnmaskedV1Param] `json:"device,required"`
	// The customer's email address.
	Email param.Field[string] `json:"email,required" format:"email"`
	// The customer's full name or business name.
	Name param.Field[string] `json:"name,required"`
	// The customer's phone number in E.164 format.
	Phone  param.Field[string]                     `json:"phone,required"`
	Status param.Field[CustomerUpdateParamsStatus] `json:"status,required"`
	// An object containing the customer's address. This is optional, but if provided,
	// all required fields must be present.
	Address param.Field[shared.AddressV11Param] `json:"address"`
	// Compliance profile for individual customers
	ComplianceProfile param.Field[shared.ComplianceProfileUnmaskedV1UnionParam] `json:"compliance_profile"`
	// Unique identifier for the customer in your database, used for cross-referencing
	// between Straddle and your systems.
	ExternalID param.Field[string] `json:"external_id"`
	// Up to 20 additional user-defined key-value pairs. Useful for storing additional
	// information about the customer in a structured format.
	Metadata          param.Field[map[string]string] `json:"metadata"`
	CorrelationID     param.Field[string]            `header:"Correlation-Id"`
	RequestID         param.Field[string]            `header:"Request-Id"`
	StraddleAccountID param.Field[string]            `header:"Straddle-Account-Id" format:"uuid"`
}

func (r CustomerUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CustomerUpdateParamsStatus string

const (
	CustomerUpdateParamsStatusPending  CustomerUpdateParamsStatus = "pending"
	CustomerUpdateParamsStatusReview   CustomerUpdateParamsStatus = "review"
	CustomerUpdateParamsStatusVerified CustomerUpdateParamsStatus = "verified"
	CustomerUpdateParamsStatusInactive CustomerUpdateParamsStatus = "inactive"
	CustomerUpdateParamsStatusRejected CustomerUpdateParamsStatus = "rejected"
)

func (r CustomerUpdateParamsStatus) IsKnown() bool {
	switch r {
	case CustomerUpdateParamsStatusPending, CustomerUpdateParamsStatusReview, CustomerUpdateParamsStatusVerified, CustomerUpdateParamsStatusInactive, CustomerUpdateParamsStatusRejected:
		return true
	}
	return false
}

type CustomerListParams struct {
	// Start date for filtering by `created_at` date.
	CreatedFrom param.Field[time.Time] `query:"created_from" format:"date-time"`
	// End date for filtering by `created_at` date.
	CreatedTo param.Field[time.Time] `query:"created_to" format:"date-time"`
	// Filter customers by `email` address.
	Email param.Field[string] `query:"email"`
	// Filter by your system's `external_id`.
	ExternalID param.Field[string] `query:"external_id"`
	// Filter customers by `name` (partial match).
	Name param.Field[string] `query:"name"`
	// Page number for paginated results. Starts at 1.
	PageNumber param.Field[int64] `query:"page_number"`
	// Number of results per page. Maximum: 1000.
	PageSize param.Field[int64] `query:"page_size"`
	// General search term to filter customers.
	SearchText param.Field[string]                   `query:"search_text"`
	SortBy     param.Field[CustomerListParamsSortBy] `query:"sort_by"`
	SortOrder  param.Field[shared.SortOrder]         `query:"sort_order"`
	// Filter customers by their current `status`.
	Status param.Field[[]shared.CustomerStatusV1] `query:"status"`
	// Filter by customer type `individual` or `business`.
	Types             param.Field[[]shared.CustomerTypeV1] `query:"types"`
	CorrelationID     param.Field[string]                  `header:"Correlation-Id"`
	RequestID         param.Field[string]                  `header:"Request-Id"`
	StraddleAccountID param.Field[string]                  `header:"Straddle-Account-Id" format:"uuid"`
}

// URLQuery serializes [CustomerListParams]'s query parameters as `url.Values`.
func (r CustomerListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CustomerListParamsSortBy string

const (
	CustomerListParamsSortByName      CustomerListParamsSortBy = "name"
	CustomerListParamsSortByCreatedAt CustomerListParamsSortBy = "created_at"
)

func (r CustomerListParamsSortBy) IsKnown() bool {
	switch r {
	case CustomerListParamsSortByName, CustomerListParamsSortByCreatedAt:
		return true
	}
	return false
}

type CustomerDeleteParams struct {
	CorrelationID     param.Field[string] `header:"Correlation-Id"`
	RequestID         param.Field[string] `header:"Request-Id"`
	StraddleAccountID param.Field[string] `header:"Straddle-Account-Id" format:"uuid"`
}

type CustomerGetParams struct {
	CorrelationID     param.Field[string] `header:"Correlation-Id"`
	RequestID         param.Field[string] `header:"Request-Id"`
	StraddleAccountID param.Field[string] `header:"Straddle-Account-Id" format:"uuid"`
}

type CustomerUnmaskedParams struct {
	CorrelationID     param.Field[string] `header:"Correlation-Id"`
	RequestID         param.Field[string] `header:"Request-Id"`
	StraddleAccountID param.Field[string] `header:"Straddle-Account-Id" format:"uuid"`
}
