// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package straddle

import (
	"context"
	"encoding/json"
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

// Customers represent the end users who send or receive payments through your
// integration. Each customer undergoes automatic identity verification and fraud
// screening upon creation. Use customers to track payment history, manage bank
// account connections, and maintain a secure record of all transactions associated
// with a user. Customers can be either individuals or businesses with appropriate
// compliance checks for each type.
//
// CustomerService contains methods and other services that help with interacting
// with the straddle API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCustomerService] method instead.
type CustomerService struct {
	options []option.RequestOption
	// Customers represent the end users who send or receive payments through your
	// integration. Each customer undergoes automatic identity verification and fraud
	// screening upon creation. Use customers to track payment history, manage bank
	// account connections, and maintain a secure record of all transactions associated
	// with a user. Customers can be either individuals or businesses with appropriate
	// compliance checks for each type.
	Review CustomerReviewService
}

// NewCustomerService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCustomerService(opts ...option.RequestOption) (r CustomerService) {
	r = CustomerService{}
	r.options = opts
	r.Review = NewCustomerReviewService(opts...)
	return
}

// Creates a new customer record and automatically initiates identity, fraud, and
// risk assessment scores. This endpoint allows you to create a customer profile
// and associate it with paykeys and payments.
func (r *CustomerService) New(ctx context.Context, params CustomerNewParams, opts ...option.RequestOption) (res *CustomerV1, err error) {
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
	path := "v1/customers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Updates an existing customer's information. This endpoint allows you to modify
// the customer's contact details, PII, and metadata.
func (r *CustomerService) Update(ctx context.Context, id string, params CustomerUpdateParams, opts ...option.RequestOption) (res *CustomerV1, err error) {
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
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/customers/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return res, err
}

// Lists or searches customers connected to your account. All supported query
// parameters are optional. If none are provided, the response will include all
// customers connected to your account. This endpoint supports advanced sorting and
// filtering options.
func (r *CustomerService) List(ctx context.Context, params CustomerListParams, opts ...option.RequestOption) (res *pagination.PageNumberSchema[CustomerSummaryPagedV1Data], err error) {
	var raw *http.Response
	if !param.IsOmitted(params.CorrelationID) {
		opts = append(opts, option.WithHeader("Correlation-Id", fmt.Sprintf("%v", params.CorrelationID.Value)))
	}
	if !param.IsOmitted(params.RequestID) {
		opts = append(opts, option.WithHeader("Request-Id", fmt.Sprintf("%v", params.RequestID.Value)))
	}
	if !param.IsOmitted(params.StraddleAccountID) {
		opts = append(opts, option.WithHeader("Straddle-Account-Id", fmt.Sprintf("%v", params.StraddleAccountID.Value)))
	}
	opts = slices.Concat(r.options, opts)
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
func (r *CustomerService) Delete(ctx context.Context, id string, body CustomerDeleteParams, opts ...option.RequestOption) (res *CustomerV1, err error) {
	if !param.IsOmitted(body.CorrelationID) {
		opts = append(opts, option.WithHeader("Correlation-Id", fmt.Sprintf("%v", body.CorrelationID.Value)))
	}
	if !param.IsOmitted(body.IdempotencyKey) {
		opts = append(opts, option.WithHeader("Idempotency-Key", fmt.Sprintf("%v", body.IdempotencyKey.Value)))
	}
	if !param.IsOmitted(body.RequestID) {
		opts = append(opts, option.WithHeader("Request-Id", fmt.Sprintf("%v", body.RequestID.Value)))
	}
	if !param.IsOmitted(body.StraddleAccountID) {
		opts = append(opts, option.WithHeader("Straddle-Account-Id", fmt.Sprintf("%v", body.StraddleAccountID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/customers/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Retrieves the details of an existing customer. Supply the unique customer ID
// that was returned from your 'create customer' request, and Straddle will return
// the corresponding customer information.
func (r *CustomerService) Get(ctx context.Context, id string, query CustomerGetParams, opts ...option.RequestOption) (res *CustomerV1, err error) {
	if !param.IsOmitted(query.CorrelationID) {
		opts = append(opts, option.WithHeader("Correlation-Id", fmt.Sprintf("%v", query.CorrelationID.Value)))
	}
	if !param.IsOmitted(query.RequestID) {
		opts = append(opts, option.WithHeader("Request-Id", fmt.Sprintf("%v", query.RequestID.Value)))
	}
	if !param.IsOmitted(query.StraddleAccountID) {
		opts = append(opts, option.WithHeader("Straddle-Account-Id", fmt.Sprintf("%v", query.StraddleAccountID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/customers/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Retrieves the unmasked details, including PII, of an existing customer. Supply
// the unique customer ID that was returned from your 'create customer' request,
// and Straddle will return the corresponding customer information. This endpoint
// needs to be enabled by Straddle and should only be used when absolutely
// necessary.
func (r *CustomerService) Unmasked(ctx context.Context, id string, query CustomerUnmaskedParams, opts ...option.RequestOption) (res *CustomerUnmaskedV1, err error) {
	if !param.IsOmitted(query.CorrelationID) {
		opts = append(opts, option.WithHeader("Correlation-Id", fmt.Sprintf("%v", query.CorrelationID.Value)))
	}
	if !param.IsOmitted(query.RequestID) {
		opts = append(opts, option.WithHeader("Request-Id", fmt.Sprintf("%v", query.RequestID.Value)))
	}
	if !param.IsOmitted(query.StraddleAccountID) {
		opts = append(opts, option.WithHeader("Straddle-Account-Id", fmt.Sprintf("%v", query.StraddleAccountID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/customers/%s/unmasked", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// An object containing the customer's address. This is optional, but if provided,
// all required fields must be present.
type CustomerAddressV1 struct {
	// Primary address line (e.g., street, PO Box).
	Address1 string `json:"address1" api:"required"`
	// City, district, suburb, town, or village.
	City string `json:"city" api:"required"`
	// Two-letter state code.
	State string `json:"state" api:"required"`
	// Zip or postal code.
	Zip string `json:"zip" api:"required"`
	// Secondary address line (e.g., apartment, suite, unit, or building).
	Address2 string `json:"address2" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Address1    respjson.Field
		City        respjson.Field
		State       respjson.Field
		Zip         respjson.Field
		Address2    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CustomerAddressV1) RawJSON() string { return r.JSON.raw }
func (r *CustomerAddressV1) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this CustomerAddressV1 to a CustomerAddressV1Param.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// CustomerAddressV1Param.Overrides()
func (r CustomerAddressV1) ToParam() CustomerAddressV1Param {
	return param.Override[CustomerAddressV1Param](json.RawMessage(r.RawJSON()))
}

// An object containing the customer's address. This is optional, but if provided,
// all required fields must be present.
//
// The properties Address1, City, State, Zip are required.
type CustomerAddressV1Param struct {
	// Primary address line (e.g., street, PO Box).
	Address1 string `json:"address1" api:"required"`
	// City, district, suburb, town, or village.
	City string `json:"city" api:"required"`
	// Two-letter state code.
	State string `json:"state" api:"required"`
	// Zip or postal code.
	Zip string `json:"zip" api:"required"`
	// Secondary address line (e.g., apartment, suite, unit, or building).
	Address2 param.Opt[string] `json:"address2,omitzero"`
	paramObj
}

func (r CustomerAddressV1Param) MarshalJSON() (data []byte, err error) {
	type shadow CustomerAddressV1Param
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CustomerAddressV1Param) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerSummaryPagedV1 struct {
	Data []CustomerSummaryPagedV1Data `json:"data" api:"required"`
	Meta CustomerSummaryPagedV1Meta   `json:"meta" api:"required"`
	// Indicates the structure of the returned content.
	//
	//   - "object" means the `data` field contains a single JSON object.
	//   - "array" means the `data` field contains an array of objects.
	//   - "error" means the `data` field contains an error object with details of the
	//     issue.
	//   - "none" means no data is returned.
	//
	// Any of "object", "array", "error", "none".
	ResponseType CustomerSummaryPagedV1ResponseType `json:"response_type" api:"required"`
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
func (r CustomerSummaryPagedV1) RawJSON() string { return r.JSON.raw }
func (r *CustomerSummaryPagedV1) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerSummaryPagedV1Data struct {
	// Unique identifier for the customer.
	ID string `json:"id" api:"required" format:"uuid"`
	// Timestamp of when the customer record was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// The customer's email address.
	Email string `json:"email" api:"required" format:"email"`
	// Full name of the individual or business name.
	Name string `json:"name" api:"required"`
	// The customer's phone number in E.164 format.
	Phone string `json:"phone" api:"required"`
	// Any of "pending", "review", "verified", "inactive", "rejected".
	Status string `json:"status" api:"required"`
	// Any of "individual", "business".
	Type string `json:"type" api:"required"`
	// Timestamp of the most recent update to the customer record.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// Unique identifier for the customer in your database, used for cross-referencing
	// between Straddle and your systems.
	ExternalID string `json:"external_id" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Email       respjson.Field
		Name        respjson.Field
		Phone       respjson.Field
		Status      respjson.Field
		Type        respjson.Field
		UpdatedAt   respjson.Field
		ExternalID  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CustomerSummaryPagedV1Data) RawJSON() string { return r.JSON.raw }
func (r *CustomerSummaryPagedV1Data) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerSummaryPagedV1Meta struct {
	// Unique identifier for this API request, useful for troubleshooting.
	APIRequestID string `json:"api_request_id" api:"required" format:"uuid"`
	// Timestamp for this API request, useful for troubleshooting.
	APIRequestTimestamp time.Time `json:"api_request_timestamp" api:"required" format:"date-time"`
	// Maximum allowed page size for this endpoint.
	MaxPageSize int64 `json:"max_page_size" api:"required"`
	// Page number for paginated results.
	PageNumber int64 `json:"page_number" api:"required"`
	// Number of items per page in this response.
	PageSize int64 `json:"page_size" api:"required"`
	// The field that the results were sorted by.
	SortBy string `json:"sort_by" api:"required"`
	// Any of "asc", "desc".
	SortOrder  string `json:"sort_order" api:"required"`
	TotalItems int64  `json:"total_items" api:"required"`
	// The number of pages available.
	TotalPages int64 `json:"total_pages" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		APIRequestID        respjson.Field
		APIRequestTimestamp respjson.Field
		MaxPageSize         respjson.Field
		PageNumber          respjson.Field
		PageSize            respjson.Field
		SortBy              respjson.Field
		SortOrder           respjson.Field
		TotalItems          respjson.Field
		TotalPages          respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CustomerSummaryPagedV1Meta) RawJSON() string { return r.JSON.raw }
func (r *CustomerSummaryPagedV1Meta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indicates the structure of the returned content.
//
//   - "object" means the `data` field contains a single JSON object.
//   - "array" means the `data` field contains an array of objects.
//   - "error" means the `data` field contains an error object with details of the
//     issue.
//   - "none" means no data is returned.
type CustomerSummaryPagedV1ResponseType string

const (
	CustomerSummaryPagedV1ResponseTypeObject CustomerSummaryPagedV1ResponseType = "object"
	CustomerSummaryPagedV1ResponseTypeArray  CustomerSummaryPagedV1ResponseType = "array"
	CustomerSummaryPagedV1ResponseTypeError  CustomerSummaryPagedV1ResponseType = "error"
	CustomerSummaryPagedV1ResponseTypeNone   CustomerSummaryPagedV1ResponseType = "none"
)

type CustomerUnmaskedV1 struct {
	Data CustomerUnmaskedV1Data `json:"data" api:"required"`
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
	ResponseType CustomerUnmaskedV1ResponseType `json:"response_type" api:"required"`
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
func (r CustomerUnmaskedV1) RawJSON() string { return r.JSON.raw }
func (r *CustomerUnmaskedV1) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerUnmaskedV1Data struct {
	// Unique identifier for the customer.
	ID string `json:"id" api:"required" format:"uuid"`
	// Timestamp of when the customer record was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// The customer's email address.
	Email string `json:"email" api:"required" format:"email"`
	// Full name of the individual or business name.
	Name string `json:"name" api:"required"`
	// The customer's phone number in E.164 format.
	Phone string `json:"phone" api:"required"`
	// Any of "pending", "review", "verified", "inactive", "rejected".
	Status string `json:"status" api:"required"`
	// Any of "individual", "business".
	Type string `json:"type" api:"required"`
	// Timestamp of the most recent update to the customer record.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// An object containing the customer's address. This is optional, but if provided,
	// all required fields must be present.
	Address CustomerAddressV1 `json:"address" api:"nullable"`
	// Individual PII data required to trigger Patriot Act compliant KYC verification.
	ComplianceProfile CustomerUnmaskedV1DataComplianceProfileUnion `json:"compliance_profile" api:"nullable"`
	Config            CustomerUnmaskedV1DataConfig                 `json:"config"`
	Device            DeviceUnmaskedV1                             `json:"device"`
	// Unique identifier for the customer in your database, used for cross-referencing
	// between Straddle and your systems.
	ExternalID string `json:"external_id" api:"nullable"`
	// Up to 20 additional user-defined key-value pairs. Useful for storing additional
	// information about the customer in a structured format.
	Metadata map[string]string `json:"metadata" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		CreatedAt         respjson.Field
		Email             respjson.Field
		Name              respjson.Field
		Phone             respjson.Field
		Status            respjson.Field
		Type              respjson.Field
		UpdatedAt         respjson.Field
		Address           respjson.Field
		ComplianceProfile respjson.Field
		Config            respjson.Field
		Device            respjson.Field
		ExternalID        respjson.Field
		Metadata          respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CustomerUnmaskedV1Data) RawJSON() string { return r.JSON.raw }
func (r *CustomerUnmaskedV1Data) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// CustomerUnmaskedV1DataComplianceProfileUnion contains all possible properties
// and values from
// [CustomerUnmaskedV1DataComplianceProfileIndividualComplianceProfile],
// [CustomerUnmaskedV1DataComplianceProfileBusinessComplianceProfile].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type CustomerUnmaskedV1DataComplianceProfileUnion struct {
	// This field is from variant
	// [CustomerUnmaskedV1DataComplianceProfileIndividualComplianceProfile].
	Dob time.Time `json:"dob"`
	// This field is from variant
	// [CustomerUnmaskedV1DataComplianceProfileIndividualComplianceProfile].
	Ssn string `json:"ssn"`
	// This field is from variant
	// [CustomerUnmaskedV1DataComplianceProfileBusinessComplianceProfile].
	Ein string `json:"ein"`
	// This field is from variant
	// [CustomerUnmaskedV1DataComplianceProfileBusinessComplianceProfile].
	LegalBusinessName string `json:"legal_business_name"`
	// This field is from variant
	// [CustomerUnmaskedV1DataComplianceProfileBusinessComplianceProfile].
	Representatives []CustomerUnmaskedV1DataComplianceProfileBusinessComplianceProfileRepresentative `json:"representatives"`
	// This field is from variant
	// [CustomerUnmaskedV1DataComplianceProfileBusinessComplianceProfile].
	Website string `json:"website"`
	JSON    struct {
		Dob               respjson.Field
		Ssn               respjson.Field
		Ein               respjson.Field
		LegalBusinessName respjson.Field
		Representatives   respjson.Field
		Website           respjson.Field
		raw               string
	} `json:"-"`
}

func (u CustomerUnmaskedV1DataComplianceProfileUnion) AsIndividualComplianceProfile() (v CustomerUnmaskedV1DataComplianceProfileIndividualComplianceProfile) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CustomerUnmaskedV1DataComplianceProfileUnion) AsBusinessComplianceProfile() (v CustomerUnmaskedV1DataComplianceProfileBusinessComplianceProfile) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u CustomerUnmaskedV1DataComplianceProfileUnion) RawJSON() string { return u.JSON.raw }

func (r *CustomerUnmaskedV1DataComplianceProfileUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Individual PII data required to trigger Patriot Act compliant KYC verification.
type CustomerUnmaskedV1DataComplianceProfileIndividualComplianceProfile struct {
	// Date of birth (YYYY-MM-DD). Required for Patriot Act-compliant KYC verification.
	Dob time.Time `json:"dob" api:"required" format:"date"`
	// Social Security Number (format XXX-XX-XXXX). Required for Patriot Act-compliant
	// KYC verification.
	Ssn string `json:"ssn" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Dob         respjson.Field
		Ssn         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CustomerUnmaskedV1DataComplianceProfileIndividualComplianceProfile) RawJSON() string {
	return r.JSON.raw
}
func (r *CustomerUnmaskedV1DataComplianceProfileIndividualComplianceProfile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Business registration data required to trigger Patriot Act compliant KYB
// verification.
type CustomerUnmaskedV1DataComplianceProfileBusinessComplianceProfile struct {
	// Employer Identification Number (format XX-XXXXXXX). Required for Patriot
	// Act-compliant KYB verification.
	Ein string `json:"ein" api:"required"`
	// Official registered business name as listed with the IRS. This value will be
	// matched against the 'legal_business name'.
	LegalBusinessName string `json:"legal_business_name" api:"required"`
	// A list of people related to the company. Only valid where customer type is
	// 'business'.
	Representatives []CustomerUnmaskedV1DataComplianceProfileBusinessComplianceProfileRepresentative `json:"representatives" api:"nullable"`
	// Official business website URL. Optional but recommended for enhanced KYB.
	Website string `json:"website" api:"nullable" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Ein               respjson.Field
		LegalBusinessName respjson.Field
		Representatives   respjson.Field
		Website           respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CustomerUnmaskedV1DataComplianceProfileBusinessComplianceProfile) RawJSON() string {
	return r.JSON.raw
}
func (r *CustomerUnmaskedV1DataComplianceProfileBusinessComplianceProfile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerUnmaskedV1DataComplianceProfileBusinessComplianceProfileRepresentative struct {
	Name  string `json:"name" api:"required"`
	Email string `json:"email" api:"nullable"`
	Phone string `json:"phone" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Email       respjson.Field
		Phone       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CustomerUnmaskedV1DataComplianceProfileBusinessComplianceProfileRepresentative) RawJSON() string {
	return r.JSON.raw
}
func (r *CustomerUnmaskedV1DataComplianceProfileBusinessComplianceProfileRepresentative) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerUnmaskedV1DataConfig struct {
	// Any of "inline", "background", "skip".
	ProcessingMethod string `json:"processing_method"`
	// Any of "standard", "verified", "rejected", "review".
	SandboxOutcome string `json:"sandbox_outcome"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ProcessingMethod respjson.Field
		SandboxOutcome   respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CustomerUnmaskedV1DataConfig) RawJSON() string { return r.JSON.raw }
func (r *CustomerUnmaskedV1DataConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indicates the structure of the returned content.
//
//   - "object" means the `data` field contains a single JSON object.
//   - "array" means the `data` field contains an array of objects.
//   - "error" means the `data` field contains an error object with details of the
//     issue.
//   - "none" means no data is returned.
type CustomerUnmaskedV1ResponseType string

const (
	CustomerUnmaskedV1ResponseTypeObject CustomerUnmaskedV1ResponseType = "object"
	CustomerUnmaskedV1ResponseTypeArray  CustomerUnmaskedV1ResponseType = "array"
	CustomerUnmaskedV1ResponseTypeError  CustomerUnmaskedV1ResponseType = "error"
	CustomerUnmaskedV1ResponseTypeNone   CustomerUnmaskedV1ResponseType = "none"
)

type CustomerV1 struct {
	Data CustomerV1Data `json:"data" api:"required"`
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
	ResponseType CustomerV1ResponseType `json:"response_type" api:"required"`
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
func (r CustomerV1) RawJSON() string { return r.JSON.raw }
func (r *CustomerV1) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerV1Data struct {
	// Unique identifier for the customer.
	ID string `json:"id" api:"required" format:"uuid"`
	// Timestamp of when the customer record was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// The customer's email address.
	Email string `json:"email" api:"required" format:"email"`
	// Full name of the individual or business name.
	Name string `json:"name" api:"required"`
	// The customer's phone number in E.164 format.
	Phone string `json:"phone" api:"required"`
	// Any of "pending", "review", "verified", "inactive", "rejected".
	Status string `json:"status" api:"required"`
	// Any of "individual", "business".
	Type string `json:"type" api:"required"`
	// Timestamp of the most recent update to the customer record.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// An object containing the customer's address. This is optional, but if provided,
	// all required fields must be present.
	Address CustomerAddressV1 `json:"address" api:"nullable"`
	// PII required to trigger Patriot Act compliant KYC verification.
	ComplianceProfile CustomerV1DataComplianceProfileUnion `json:"compliance_profile" api:"nullable"`
	Config            CustomerV1DataConfig                 `json:"config"`
	Device            CustomerV1DataDevice                 `json:"device"`
	// Unique identifier for the customer in your database, used for cross-referencing
	// between Straddle and your systems.
	ExternalID string `json:"external_id" api:"nullable"`
	// Up to 20 additional user-defined key-value pairs. Useful for storing additional
	// information about the customer in a structured format.
	Metadata map[string]string `json:"metadata" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		CreatedAt         respjson.Field
		Email             respjson.Field
		Name              respjson.Field
		Phone             respjson.Field
		Status            respjson.Field
		Type              respjson.Field
		UpdatedAt         respjson.Field
		Address           respjson.Field
		ComplianceProfile respjson.Field
		Config            respjson.Field
		Device            respjson.Field
		ExternalID        respjson.Field
		Metadata          respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CustomerV1Data) RawJSON() string { return r.JSON.raw }
func (r *CustomerV1Data) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// CustomerV1DataComplianceProfileUnion contains all possible properties and values
// from [CustomerV1DataComplianceProfileIndividualComplianceProfile],
// [CustomerV1DataComplianceProfileBusinessComplianceProfile].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type CustomerV1DataComplianceProfileUnion struct {
	// This field is from variant
	// [CustomerV1DataComplianceProfileIndividualComplianceProfile].
	Dob time.Time `json:"dob"`
	// This field is from variant
	// [CustomerV1DataComplianceProfileIndividualComplianceProfile].
	Ssn string `json:"ssn"`
	// This field is from variant
	// [CustomerV1DataComplianceProfileBusinessComplianceProfile].
	Ein string `json:"ein"`
	// This field is from variant
	// [CustomerV1DataComplianceProfileBusinessComplianceProfile].
	LegalBusinessName string `json:"legal_business_name"`
	// This field is from variant
	// [CustomerV1DataComplianceProfileBusinessComplianceProfile].
	Representatives []CustomerV1DataComplianceProfileBusinessComplianceProfileRepresentative `json:"representatives"`
	// This field is from variant
	// [CustomerV1DataComplianceProfileBusinessComplianceProfile].
	Website string `json:"website"`
	JSON    struct {
		Dob               respjson.Field
		Ssn               respjson.Field
		Ein               respjson.Field
		LegalBusinessName respjson.Field
		Representatives   respjson.Field
		Website           respjson.Field
		raw               string
	} `json:"-"`
}

func (u CustomerV1DataComplianceProfileUnion) AsIndividualComplianceProfile() (v CustomerV1DataComplianceProfileIndividualComplianceProfile) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CustomerV1DataComplianceProfileUnion) AsBusinessComplianceProfile() (v CustomerV1DataComplianceProfileBusinessComplianceProfile) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u CustomerV1DataComplianceProfileUnion) RawJSON() string { return u.JSON.raw }

func (r *CustomerV1DataComplianceProfileUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PII required to trigger Patriot Act compliant KYC verification.
type CustomerV1DataComplianceProfileIndividualComplianceProfile struct {
	// Masked date of birth in \***\*-**-\*\* format.
	Dob time.Time `json:"dob" api:"required" format:"date"`
	// Masked Social Security Number in the format **\*-**-\*\*\*\*.
	Ssn string `json:"ssn" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Dob         respjson.Field
		Ssn         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CustomerV1DataComplianceProfileIndividualComplianceProfile) RawJSON() string {
	return r.JSON.raw
}
func (r *CustomerV1DataComplianceProfileIndividualComplianceProfile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Business registration data required to trigger Patriot Act compliant KYB
// verification.
type CustomerV1DataComplianceProfileBusinessComplianceProfile struct {
	// Masked Employer Identification Number in the format **-**\*****
	Ein string `json:"ein" api:"required"`
	// The official registered name of the business. This name should be correlated
	// with the `ein` value.
	LegalBusinessName string `json:"legal_business_name" api:"required"`
	// A list of people related to the company. Only valid where customer type is
	// 'business'.
	Representatives []CustomerV1DataComplianceProfileBusinessComplianceProfileRepresentative `json:"representatives" api:"nullable"`
	// Official business website URL. Optional but recommended for enhanced KYB.
	Website string `json:"website" api:"nullable" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Ein               respjson.Field
		LegalBusinessName respjson.Field
		Representatives   respjson.Field
		Website           respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CustomerV1DataComplianceProfileBusinessComplianceProfile) RawJSON() string { return r.JSON.raw }
func (r *CustomerV1DataComplianceProfileBusinessComplianceProfile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerV1DataComplianceProfileBusinessComplianceProfileRepresentative struct {
	Name  string `json:"name" api:"required"`
	Email string `json:"email" api:"nullable"`
	Phone string `json:"phone" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Email       respjson.Field
		Phone       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CustomerV1DataComplianceProfileBusinessComplianceProfileRepresentative) RawJSON() string {
	return r.JSON.raw
}
func (r *CustomerV1DataComplianceProfileBusinessComplianceProfileRepresentative) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerV1DataConfig struct {
	// Any of "inline", "background", "skip".
	ProcessingMethod string `json:"processing_method"`
	// Any of "standard", "verified", "rejected", "review".
	SandboxOutcome string `json:"sandbox_outcome"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ProcessingMethod respjson.Field
		SandboxOutcome   respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CustomerV1DataConfig) RawJSON() string { return r.JSON.raw }
func (r *CustomerV1DataConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerV1DataDevice struct {
	// The customer's IP address at the time of profile creation. Use `0.0.0.0` to
	// represent an offline customer registration.
	IPAddress string `json:"ip_address" api:"required" format:"ipv4"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IPAddress   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CustomerV1DataDevice) RawJSON() string { return r.JSON.raw }
func (r *CustomerV1DataDevice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indicates the structure of the returned content.
//
//   - "object" means the `data` field contains a single JSON object.
//   - "array" means the `data` field contains an array of objects.
//   - "error" means the `data` field contains an error object with details of the
//     issue.
//   - "none" means no data is returned.
type CustomerV1ResponseType string

const (
	CustomerV1ResponseTypeObject CustomerV1ResponseType = "object"
	CustomerV1ResponseTypeArray  CustomerV1ResponseType = "array"
	CustomerV1ResponseTypeError  CustomerV1ResponseType = "error"
	CustomerV1ResponseTypeNone   CustomerV1ResponseType = "none"
)

type DeviceUnmaskedV1 struct {
	// The customer's IP address at the time of profile creation. Use `0.0.0.0` to
	// represent an offline customer registration.
	IPAddress string `json:"ip_address" api:"required" format:"ipv4"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IPAddress   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DeviceUnmaskedV1) RawJSON() string { return r.JSON.raw }
func (r *DeviceUnmaskedV1) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this DeviceUnmaskedV1 to a DeviceUnmaskedV1Param.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// DeviceUnmaskedV1Param.Overrides()
func (r DeviceUnmaskedV1) ToParam() DeviceUnmaskedV1Param {
	return param.Override[DeviceUnmaskedV1Param](json.RawMessage(r.RawJSON()))
}

// The property IPAddress is required.
type DeviceUnmaskedV1Param struct {
	// The customer's IP address at the time of profile creation. Use `0.0.0.0` to
	// represent an offline customer registration.
	IPAddress string `json:"ip_address" api:"required" format:"ipv4"`
	paramObj
}

func (r DeviceUnmaskedV1Param) MarshalJSON() (data []byte, err error) {
	type shadow DeviceUnmaskedV1Param
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DeviceUnmaskedV1Param) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerNewParams struct {
	Device DeviceUnmaskedV1Param `json:"device,omitzero" api:"required"`
	// The customer's email address.
	Email string `json:"email" api:"required" format:"email"`
	// Full name of the individual or business name.
	Name string `json:"name" api:"required"`
	// The customer's phone number in E.164 format. Mobile number is preferred.
	Phone string `json:"phone" api:"required"`
	// Any of "individual", "business".
	Type CustomerNewParamsType `json:"type,omitzero" api:"required"`
	// Unique identifier for the customer in your database, used for cross-referencing
	// between Straddle and your systems.
	ExternalID        param.Opt[string] `json:"external_id,omitzero"`
	CorrelationID     param.Opt[string] `header:"Correlation-Id,omitzero" json:"-"`
	IdempotencyKey    param.Opt[string] `header:"Idempotency-Key,omitzero" json:"-"`
	RequestID         param.Opt[string] `header:"Request-Id,omitzero" json:"-"`
	StraddleAccountID param.Opt[string] `header:"Straddle-Account-Id,omitzero" format:"uuid" json:"-"`
	// An object containing the customer's address. **This is optional.** If used, all
	// required fields must be present.
	Address CustomerAddressV1Param `json:"address,omitzero"`
	// An object containing the customer's compliance profile. **This is optional.** If
	// all required fields must be present for the appropriate customer type.
	ComplianceProfile CustomerNewParamsComplianceProfileUnion `json:"compliance_profile,omitzero"`
	// Up to 20 additional user-defined key-value pairs. Useful for storing additional
	// information about the customer in a structured format.
	Metadata map[string]string       `json:"metadata,omitzero"`
	Config   CustomerNewParamsConfig `json:"config,omitzero"`
	paramObj
}

func (r CustomerNewParams) MarshalJSON() (data []byte, err error) {
	type shadow CustomerNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CustomerNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerNewParamsType string

const (
	CustomerNewParamsTypeIndividual CustomerNewParamsType = "individual"
	CustomerNewParamsTypeBusiness   CustomerNewParamsType = "business"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type CustomerNewParamsComplianceProfileUnion struct {
	OfIndividualComplianceProfile *CustomerNewParamsComplianceProfileIndividualComplianceProfile `json:",omitzero,inline"`
	OfBusinessComplianceProfile   *CustomerNewParamsComplianceProfileBusinessComplianceProfile   `json:",omitzero,inline"`
	paramUnion
}

func (u CustomerNewParamsComplianceProfileUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfIndividualComplianceProfile, u.OfBusinessComplianceProfile)
}
func (u *CustomerNewParamsComplianceProfileUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Individual PII data required to trigger Patriot Act compliant KYC verification.
//
// The properties Dob, Ssn are required.
type CustomerNewParamsComplianceProfileIndividualComplianceProfile struct {
	// Date of birth (YYYY-MM-DD). Required for Patriot Act-compliant KYC verification.
	Dob param.Opt[time.Time] `json:"dob,omitzero" api:"required" format:"date"`
	// Social Security Number (format XXX-XX-XXXX). Required for Patriot Act-compliant
	// KYC verification.
	Ssn param.Opt[string] `json:"ssn,omitzero" api:"required"`
	paramObj
}

func (r CustomerNewParamsComplianceProfileIndividualComplianceProfile) MarshalJSON() (data []byte, err error) {
	type shadow CustomerNewParamsComplianceProfileIndividualComplianceProfile
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CustomerNewParamsComplianceProfileIndividualComplianceProfile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Business registration data required to trigger Patriot Act compliant KYB
// verification.
//
// The properties Ein, LegalBusinessName are required.
type CustomerNewParamsComplianceProfileBusinessComplianceProfile struct {
	// Employer Identification Number (format XX-XXXXXXX). Required for Patriot
	// Act-compliant KYB verification.
	Ein param.Opt[string] `json:"ein,omitzero" api:"required"`
	// Official registered business name as listed with the IRS. This value will be
	// matched against the 'legal_business name'.
	LegalBusinessName param.Opt[string] `json:"legal_business_name,omitzero" api:"required"`
	// Official business website URL. Optional but recommended for enhanced KYB.
	Website param.Opt[string] `json:"website,omitzero" format:"uri"`
	// A list of people related to the company. Only valid where customer type is
	// 'business'.
	Representatives []CustomerNewParamsComplianceProfileBusinessComplianceProfileRepresentative `json:"representatives,omitzero"`
	paramObj
}

func (r CustomerNewParamsComplianceProfileBusinessComplianceProfile) MarshalJSON() (data []byte, err error) {
	type shadow CustomerNewParamsComplianceProfileBusinessComplianceProfile
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CustomerNewParamsComplianceProfileBusinessComplianceProfile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Name is required.
type CustomerNewParamsComplianceProfileBusinessComplianceProfileRepresentative struct {
	Name  string            `json:"name" api:"required"`
	Email param.Opt[string] `json:"email,omitzero"`
	Phone param.Opt[string] `json:"phone,omitzero"`
	paramObj
}

func (r CustomerNewParamsComplianceProfileBusinessComplianceProfileRepresentative) MarshalJSON() (data []byte, err error) {
	type shadow CustomerNewParamsComplianceProfileBusinessComplianceProfileRepresentative
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CustomerNewParamsComplianceProfileBusinessComplianceProfileRepresentative) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerNewParamsConfig struct {
	// Any of "inline", "background", "skip".
	ProcessingMethod string `json:"processing_method,omitzero"`
	// Any of "standard", "verified", "rejected", "review".
	SandboxOutcome string `json:"sandbox_outcome,omitzero"`
	paramObj
}

func (r CustomerNewParamsConfig) MarshalJSON() (data []byte, err error) {
	type shadow CustomerNewParamsConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CustomerNewParamsConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CustomerNewParamsConfig](
		"processing_method", "inline", "background", "skip",
	)
	apijson.RegisterFieldValidator[CustomerNewParamsConfig](
		"sandbox_outcome", "standard", "verified", "rejected", "review",
	)
}

type CustomerUpdateParams struct {
	Device DeviceUnmaskedV1Param `json:"device,omitzero" api:"required"`
	// The customer's email address.
	Email string `json:"email" api:"required" format:"email"`
	// The customer's full name or business name.
	Name string `json:"name" api:"required"`
	// The customer's phone number in E.164 format.
	Phone string `json:"phone" api:"required"`
	// Any of "pending", "review", "verified", "inactive", "rejected".
	Status CustomerUpdateParamsStatus `json:"status,omitzero" api:"required"`
	// Unique identifier for the customer in your database, used for cross-referencing
	// between Straddle and your systems.
	ExternalID        param.Opt[string] `json:"external_id,omitzero"`
	CorrelationID     param.Opt[string] `header:"Correlation-Id,omitzero" json:"-"`
	IdempotencyKey    param.Opt[string] `header:"Idempotency-Key,omitzero" json:"-"`
	RequestID         param.Opt[string] `header:"Request-Id,omitzero" json:"-"`
	StraddleAccountID param.Opt[string] `header:"Straddle-Account-Id,omitzero" format:"uuid" json:"-"`
	// An object containing the customer's address. This is optional, but if provided,
	// all required fields must be present.
	Address CustomerAddressV1Param `json:"address,omitzero"`
	// Individual PII data required to trigger Patriot Act compliant KYC verification.
	ComplianceProfile CustomerUpdateParamsComplianceProfileUnion `json:"compliance_profile,omitzero"`
	// Up to 20 additional user-defined key-value pairs. Useful for storing additional
	// information about the customer in a structured format.
	Metadata map[string]string `json:"metadata,omitzero"`
	paramObj
}

func (r CustomerUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow CustomerUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CustomerUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerUpdateParamsStatus string

const (
	CustomerUpdateParamsStatusPending  CustomerUpdateParamsStatus = "pending"
	CustomerUpdateParamsStatusReview   CustomerUpdateParamsStatus = "review"
	CustomerUpdateParamsStatusVerified CustomerUpdateParamsStatus = "verified"
	CustomerUpdateParamsStatusInactive CustomerUpdateParamsStatus = "inactive"
	CustomerUpdateParamsStatusRejected CustomerUpdateParamsStatus = "rejected"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type CustomerUpdateParamsComplianceProfileUnion struct {
	OfIndividualComplianceProfile *CustomerUpdateParamsComplianceProfileIndividualComplianceProfile `json:",omitzero,inline"`
	OfBusinessComplianceProfile   *CustomerUpdateParamsComplianceProfileBusinessComplianceProfile   `json:",omitzero,inline"`
	paramUnion
}

func (u CustomerUpdateParamsComplianceProfileUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfIndividualComplianceProfile, u.OfBusinessComplianceProfile)
}
func (u *CustomerUpdateParamsComplianceProfileUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Individual PII data required to trigger Patriot Act compliant KYC verification.
//
// The properties Dob, Ssn are required.
type CustomerUpdateParamsComplianceProfileIndividualComplianceProfile struct {
	// Date of birth (YYYY-MM-DD). Required for Patriot Act-compliant KYC verification.
	Dob param.Opt[time.Time] `json:"dob,omitzero" api:"required" format:"date"`
	// Social Security Number (format XXX-XX-XXXX). Required for Patriot Act-compliant
	// KYC verification.
	Ssn param.Opt[string] `json:"ssn,omitzero" api:"required"`
	paramObj
}

func (r CustomerUpdateParamsComplianceProfileIndividualComplianceProfile) MarshalJSON() (data []byte, err error) {
	type shadow CustomerUpdateParamsComplianceProfileIndividualComplianceProfile
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CustomerUpdateParamsComplianceProfileIndividualComplianceProfile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Business registration data required to trigger Patriot Act compliant KYB
// verification.
//
// The properties Ein, LegalBusinessName are required.
type CustomerUpdateParamsComplianceProfileBusinessComplianceProfile struct {
	// Employer Identification Number (format XX-XXXXXXX). Required for Patriot
	// Act-compliant KYB verification.
	Ein param.Opt[string] `json:"ein,omitzero" api:"required"`
	// Official registered business name as listed with the IRS. This value will be
	// matched against the 'legal_business name'.
	LegalBusinessName param.Opt[string] `json:"legal_business_name,omitzero" api:"required"`
	// Official business website URL. Optional but recommended for enhanced KYB.
	Website param.Opt[string] `json:"website,omitzero" format:"uri"`
	// A list of people related to the company. Only valid where customer type is
	// 'business'.
	Representatives []CustomerUpdateParamsComplianceProfileBusinessComplianceProfileRepresentative `json:"representatives,omitzero"`
	paramObj
}

func (r CustomerUpdateParamsComplianceProfileBusinessComplianceProfile) MarshalJSON() (data []byte, err error) {
	type shadow CustomerUpdateParamsComplianceProfileBusinessComplianceProfile
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CustomerUpdateParamsComplianceProfileBusinessComplianceProfile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Name is required.
type CustomerUpdateParamsComplianceProfileBusinessComplianceProfileRepresentative struct {
	Name  string            `json:"name" api:"required"`
	Email param.Opt[string] `json:"email,omitzero"`
	Phone param.Opt[string] `json:"phone,omitzero"`
	paramObj
}

func (r CustomerUpdateParamsComplianceProfileBusinessComplianceProfileRepresentative) MarshalJSON() (data []byte, err error) {
	type shadow CustomerUpdateParamsComplianceProfileBusinessComplianceProfileRepresentative
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CustomerUpdateParamsComplianceProfileBusinessComplianceProfileRepresentative) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerListParams struct {
	// Start date for filtering by `created_at` date.
	CreatedFrom param.Opt[time.Time] `query:"created_from,omitzero" format:"date-time" json:"-"`
	// End date for filtering by `created_at` date.
	CreatedTo param.Opt[time.Time] `query:"created_to,omitzero" format:"date-time" json:"-"`
	// Filter customers by `email` address.
	Email param.Opt[string] `query:"email,omitzero" json:"-"`
	// Filter by your system's `external_id`.
	ExternalID param.Opt[string] `query:"external_id,omitzero" json:"-"`
	// Filter customers by `name` (partial match).
	Name param.Opt[string] `query:"name,omitzero" json:"-"`
	// Page number for paginated results. Starts at 1.
	PageNumber param.Opt[int64] `query:"page_number,omitzero" json:"-"`
	// Number of results per page. Maximum: 1000.
	PageSize param.Opt[int64] `query:"page_size,omitzero" json:"-"`
	// General search term to filter customers.
	SearchText        param.Opt[string] `query:"search_text,omitzero" json:"-"`
	CorrelationID     param.Opt[string] `header:"Correlation-Id,omitzero" json:"-"`
	RequestID         param.Opt[string] `header:"Request-Id,omitzero" json:"-"`
	StraddleAccountID param.Opt[string] `header:"Straddle-Account-Id,omitzero" format:"uuid" json:"-"`
	// Any of "name", "created_at".
	SortBy CustomerListParamsSortBy `query:"sort_by,omitzero" json:"-"`
	// Any of "asc", "desc".
	SortOrder CustomerListParamsSortOrder `query:"sort_order,omitzero" json:"-"`
	// Filter customers by their current `status`.
	//
	// Any of "pending", "review", "verified", "inactive", "rejected".
	Status []string `query:"status,omitzero" json:"-"`
	// Filter by customer type `individual` or `business`.
	//
	// Any of "individual", "business".
	Types []string `query:"types,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CustomerListParams]'s query parameters as `url.Values`.
func (r CustomerListParams) URLQuery() (v url.Values, err error) {
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

type CustomerListParamsSortOrder string

const (
	CustomerListParamsSortOrderAsc  CustomerListParamsSortOrder = "asc"
	CustomerListParamsSortOrderDesc CustomerListParamsSortOrder = "desc"
)

type CustomerDeleteParams struct {
	CorrelationID     param.Opt[string] `header:"Correlation-Id,omitzero" json:"-"`
	IdempotencyKey    param.Opt[string] `header:"Idempotency-Key,omitzero" json:"-"`
	RequestID         param.Opt[string] `header:"Request-Id,omitzero" json:"-"`
	StraddleAccountID param.Opt[string] `header:"Straddle-Account-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

type CustomerGetParams struct {
	CorrelationID     param.Opt[string] `header:"Correlation-Id,omitzero" json:"-"`
	RequestID         param.Opt[string] `header:"Request-Id,omitzero" json:"-"`
	StraddleAccountID param.Opt[string] `header:"Straddle-Account-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

type CustomerUnmaskedParams struct {
	CorrelationID     param.Opt[string] `header:"Correlation-Id,omitzero" json:"-"`
	RequestID         param.Opt[string] `header:"Request-Id,omitzero" json:"-"`
	StraddleAccountID param.Opt[string] `header:"Straddle-Account-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}
