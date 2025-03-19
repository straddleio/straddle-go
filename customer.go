// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package straddle

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/stainless-sdks/straddle-go/internal/apijson"
	"github.com/stainless-sdks/straddle-go/internal/apiquery"
	"github.com/stainless-sdks/straddle-go/internal/param"
	"github.com/stainless-sdks/straddle-go/internal/requestconfig"
	"github.com/stainless-sdks/straddle-go/option"
	"github.com/stainless-sdks/straddle-go/packages/pagination"
	"github.com/stainless-sdks/straddle-go/shared"
	"github.com/tidwall/gjson"
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
func (r *CustomerService) New(ctx context.Context, params CustomerNewParams, opts ...option.RequestOption) (res *CustomerV1, err error) {
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
func (r *CustomerService) Update(ctx context.Context, id string, params CustomerUpdateParams, opts ...option.RequestOption) (res *CustomerV1, err error) {
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
func (r *CustomerService) Delete(ctx context.Context, id string, body CustomerDeleteParams, opts ...option.RequestOption) (res *CustomerV1, err error) {
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
func (r *CustomerService) Get(ctx context.Context, id string, query CustomerGetParams, opts ...option.RequestOption) (res *CustomerV1, err error) {
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

// Updates the decision of a customer's identity validation. This endpoint allows
// you to modify the outcome of a customer decision and is useful for correcting or
// updating the status of a customer's verification.
func (r *CustomerService) RefreshReview(ctx context.Context, id string, body CustomerRefreshReviewParams, opts ...option.RequestOption) (res *CustomerV1, err error) {
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
	path := fmt.Sprintf("v1/customers/%s/refresh_review", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, nil, &res, opts...)
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

// An object containing the customer's address. This is optional, but if provided,
// all required fields must be present.
type CustomerAddressV1 struct {
	// Primary address line (e.g., street, PO Box).
	Address1 string `json:"address1,required"`
	// City, district, suburb, town, or village.
	City string `json:"city,required"`
	// Two-letter state code.
	State string `json:"state,required"`
	// Zip or postal code.
	Zip string `json:"zip,required"`
	// Secondary address line (e.g., apartment, suite, unit, or building).
	Address2 string                `json:"address2,nullable"`
	JSON     customerAddressV1JSON `json:"-"`
}

// customerAddressV1JSON contains the JSON metadata for the struct
// [CustomerAddressV1]
type customerAddressV1JSON struct {
	Address1    apijson.Field
	City        apijson.Field
	State       apijson.Field
	Zip         apijson.Field
	Address2    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerAddressV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerAddressV1JSON) RawJSON() string {
	return r.raw
}

// An object containing the customer's address. This is optional, but if provided,
// all required fields must be present.
type CustomerAddressV1Param struct {
	// Primary address line (e.g., street, PO Box).
	Address1 param.Field[string] `json:"address1,required"`
	// City, district, suburb, town, or village.
	City param.Field[string] `json:"city,required"`
	// Two-letter state code.
	State param.Field[string] `json:"state,required"`
	// Zip or postal code.
	Zip param.Field[string] `json:"zip,required"`
	// Secondary address line (e.g., apartment, suite, unit, or building).
	Address2 param.Field[string] `json:"address2"`
}

func (r CustomerAddressV1Param) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CustomerSummaryPagedV1 struct {
	Data []CustomerSummaryPagedV1Data `json:"data,required"`
	Meta CustomerSummaryPagedV1Meta   `json:"meta,required"`
	// Indicates the structure of the returned content.
	//
	//   - "object" means the `data` field contains a single JSON object.
	//   - "array" means the `data` field contains an array of objects.
	//   - "error" means the `data` field contains an error object with details of the
	//     issue.
	//   - "none" means no data is returned.
	ResponseType CustomerSummaryPagedV1ResponseType `json:"response_type,required"`
	JSON         customerSummaryPagedV1JSON         `json:"-"`
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
	Phone  string                           `json:"phone,required"`
	Status CustomerSummaryPagedV1DataStatus `json:"status,required"`
	Type   CustomerSummaryPagedV1DataType   `json:"type,required"`
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

type CustomerSummaryPagedV1DataStatus string

const (
	CustomerSummaryPagedV1DataStatusPending  CustomerSummaryPagedV1DataStatus = "pending"
	CustomerSummaryPagedV1DataStatusReview   CustomerSummaryPagedV1DataStatus = "review"
	CustomerSummaryPagedV1DataStatusVerified CustomerSummaryPagedV1DataStatus = "verified"
	CustomerSummaryPagedV1DataStatusInactive CustomerSummaryPagedV1DataStatus = "inactive"
	CustomerSummaryPagedV1DataStatusRejected CustomerSummaryPagedV1DataStatus = "rejected"
)

func (r CustomerSummaryPagedV1DataStatus) IsKnown() bool {
	switch r {
	case CustomerSummaryPagedV1DataStatusPending, CustomerSummaryPagedV1DataStatusReview, CustomerSummaryPagedV1DataStatusVerified, CustomerSummaryPagedV1DataStatusInactive, CustomerSummaryPagedV1DataStatusRejected:
		return true
	}
	return false
}

type CustomerSummaryPagedV1DataType string

const (
	CustomerSummaryPagedV1DataTypeIndividual CustomerSummaryPagedV1DataType = "individual"
	CustomerSummaryPagedV1DataTypeBusiness   CustomerSummaryPagedV1DataType = "business"
)

func (r CustomerSummaryPagedV1DataType) IsKnown() bool {
	switch r {
	case CustomerSummaryPagedV1DataTypeIndividual, CustomerSummaryPagedV1DataTypeBusiness:
		return true
	}
	return false
}

type CustomerSummaryPagedV1Meta struct {
	// Unique identifier for this API request, useful for troubleshooting.
	APIRequestID string `json:"api_request_id,required" format:"uuid"`
	// Timestamp for this API request, useful for troubleshooting.
	APIRequestTimestamp time.Time `json:"api_request_timestamp,required" format:"date-time"`
	// Maximum allowed page size for this endpoint.
	MaxPageSize int64 `json:"max_page_size,required"`
	// Page number for paginated results.
	PageNumber int64 `json:"page_number,required"`
	// Number of items per page in this response.
	PageSize int64 `json:"page_size,required"`
	// The field that the results were sorted by.
	SortBy     string                              `json:"sort_by,required"`
	SortOrder  CustomerSummaryPagedV1MetaSortOrder `json:"sort_order,required"`
	TotalItems int64                               `json:"total_items,required"`
	// The number of pages available.
	TotalPages int64                          `json:"total_pages,required"`
	JSON       customerSummaryPagedV1MetaJSON `json:"-"`
}

// customerSummaryPagedV1MetaJSON contains the JSON metadata for the struct
// [CustomerSummaryPagedV1Meta]
type customerSummaryPagedV1MetaJSON struct {
	APIRequestID        apijson.Field
	APIRequestTimestamp apijson.Field
	MaxPageSize         apijson.Field
	PageNumber          apijson.Field
	PageSize            apijson.Field
	SortBy              apijson.Field
	SortOrder           apijson.Field
	TotalItems          apijson.Field
	TotalPages          apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *CustomerSummaryPagedV1Meta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerSummaryPagedV1MetaJSON) RawJSON() string {
	return r.raw
}

type CustomerSummaryPagedV1MetaSortOrder string

const (
	CustomerSummaryPagedV1MetaSortOrderAsc  CustomerSummaryPagedV1MetaSortOrder = "asc"
	CustomerSummaryPagedV1MetaSortOrderDesc CustomerSummaryPagedV1MetaSortOrder = "desc"
)

func (r CustomerSummaryPagedV1MetaSortOrder) IsKnown() bool {
	switch r {
	case CustomerSummaryPagedV1MetaSortOrderAsc, CustomerSummaryPagedV1MetaSortOrderDesc:
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
type CustomerSummaryPagedV1ResponseType string

const (
	CustomerSummaryPagedV1ResponseTypeObject CustomerSummaryPagedV1ResponseType = "object"
	CustomerSummaryPagedV1ResponseTypeArray  CustomerSummaryPagedV1ResponseType = "array"
	CustomerSummaryPagedV1ResponseTypeError  CustomerSummaryPagedV1ResponseType = "error"
	CustomerSummaryPagedV1ResponseTypeNone   CustomerSummaryPagedV1ResponseType = "none"
)

func (r CustomerSummaryPagedV1ResponseType) IsKnown() bool {
	switch r {
	case CustomerSummaryPagedV1ResponseTypeObject, CustomerSummaryPagedV1ResponseTypeArray, CustomerSummaryPagedV1ResponseTypeError, CustomerSummaryPagedV1ResponseTypeNone:
		return true
	}
	return false
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
	ResponseType CustomerUnmaskedV1ResponseType `json:"response_type,required"`
	JSON         customerUnmaskedV1JSON         `json:"-"`
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
	Phone  string                       `json:"phone,required"`
	Status CustomerUnmaskedV1DataStatus `json:"status,required"`
	Type   CustomerUnmaskedV1DataType   `json:"type,required"`
	// Timestamp of the most recent update to the customer record.
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// An object containing the customer's address. This is optional, but if provided,
	// all required fields must be present.
	Address CustomerAddressV1 `json:"address,nullable"`
	// Individual PII data required to trigger Patriot Act compliant KYC verification.
	ComplianceProfile CustomerUnmaskedV1DataComplianceProfile `json:"compliance_profile,nullable"`
	Device            DeviceUnmaskedV1                        `json:"device"`
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

type CustomerUnmaskedV1DataStatus string

const (
	CustomerUnmaskedV1DataStatusPending  CustomerUnmaskedV1DataStatus = "pending"
	CustomerUnmaskedV1DataStatusReview   CustomerUnmaskedV1DataStatus = "review"
	CustomerUnmaskedV1DataStatusVerified CustomerUnmaskedV1DataStatus = "verified"
	CustomerUnmaskedV1DataStatusInactive CustomerUnmaskedV1DataStatus = "inactive"
	CustomerUnmaskedV1DataStatusRejected CustomerUnmaskedV1DataStatus = "rejected"
)

func (r CustomerUnmaskedV1DataStatus) IsKnown() bool {
	switch r {
	case CustomerUnmaskedV1DataStatusPending, CustomerUnmaskedV1DataStatusReview, CustomerUnmaskedV1DataStatusVerified, CustomerUnmaskedV1DataStatusInactive, CustomerUnmaskedV1DataStatusRejected:
		return true
	}
	return false
}

type CustomerUnmaskedV1DataType string

const (
	CustomerUnmaskedV1DataTypeIndividual CustomerUnmaskedV1DataType = "individual"
	CustomerUnmaskedV1DataTypeBusiness   CustomerUnmaskedV1DataType = "business"
)

func (r CustomerUnmaskedV1DataType) IsKnown() bool {
	switch r {
	case CustomerUnmaskedV1DataTypeIndividual, CustomerUnmaskedV1DataTypeBusiness:
		return true
	}
	return false
}

// Individual PII data required to trigger Patriot Act compliant KYC verification.
type CustomerUnmaskedV1DataComplianceProfile struct {
	// Date of birth (YYYY-MM-DD). Required for Patriot Act-compliant KYC verification.
	Dob time.Time `json:"dob,nullable" format:"date"`
	// Employer Identification Number (format XX-XXXXXXX). Required for Patriot
	// Act-compliant KYB verification.
	Ein string `json:"ein,nullable"`
	// Official registered business name as listed with the IRS. This value will be
	// matched against the 'legal_business name'.
	LegalBusinessName string `json:"legal_business_name,nullable"`
	// Social Security Number (format XXX-XX-XXXX). Required for Patriot Act-compliant
	// KYC verification.
	Ssn string `json:"ssn,nullable"`
	// Official business website URL. Optional but recommended for enhanced KYB.
	Website string                                      `json:"website,nullable" format:"uri"`
	JSON    customerUnmaskedV1DataComplianceProfileJSON `json:"-"`
	union   CustomerUnmaskedV1DataComplianceProfileUnion
}

// customerUnmaskedV1DataComplianceProfileJSON contains the JSON metadata for the
// struct [CustomerUnmaskedV1DataComplianceProfile]
type customerUnmaskedV1DataComplianceProfileJSON struct {
	Dob               apijson.Field
	Ein               apijson.Field
	LegalBusinessName apijson.Field
	Ssn               apijson.Field
	Website           apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r customerUnmaskedV1DataComplianceProfileJSON) RawJSON() string {
	return r.raw
}

func (r *CustomerUnmaskedV1DataComplianceProfile) UnmarshalJSON(data []byte) (err error) {
	*r = CustomerUnmaskedV1DataComplianceProfile{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [CustomerUnmaskedV1DataComplianceProfileUnion] interface which
// you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [CustomerUnmaskedV1DataComplianceProfileIndividualComplianceProfile],
// [CustomerUnmaskedV1DataComplianceProfileBusinessComplianceProfile].
func (r CustomerUnmaskedV1DataComplianceProfile) AsUnion() CustomerUnmaskedV1DataComplianceProfileUnion {
	return r.union
}

// Individual PII data required to trigger Patriot Act compliant KYC verification.
//
// Union satisfied by
// [CustomerUnmaskedV1DataComplianceProfileIndividualComplianceProfile] or
// [CustomerUnmaskedV1DataComplianceProfileBusinessComplianceProfile].
type CustomerUnmaskedV1DataComplianceProfileUnion interface {
	implementsCustomerUnmaskedV1DataComplianceProfile()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CustomerUnmaskedV1DataComplianceProfileUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CustomerUnmaskedV1DataComplianceProfileIndividualComplianceProfile{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CustomerUnmaskedV1DataComplianceProfileBusinessComplianceProfile{}),
		},
	)
}

// Individual PII data required to trigger Patriot Act compliant KYC verification.
type CustomerUnmaskedV1DataComplianceProfileIndividualComplianceProfile struct {
	// Date of birth (YYYY-MM-DD). Required for Patriot Act-compliant KYC verification.
	Dob time.Time `json:"dob,required,nullable" format:"date"`
	// Social Security Number (format XXX-XX-XXXX). Required for Patriot Act-compliant
	// KYC verification.
	Ssn  string                                                                 `json:"ssn,required,nullable"`
	JSON customerUnmaskedV1DataComplianceProfileIndividualComplianceProfileJSON `json:"-"`
}

// customerUnmaskedV1DataComplianceProfileIndividualComplianceProfileJSON contains
// the JSON metadata for the struct
// [CustomerUnmaskedV1DataComplianceProfileIndividualComplianceProfile]
type customerUnmaskedV1DataComplianceProfileIndividualComplianceProfileJSON struct {
	Dob         apijson.Field
	Ssn         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerUnmaskedV1DataComplianceProfileIndividualComplianceProfile) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerUnmaskedV1DataComplianceProfileIndividualComplianceProfileJSON) RawJSON() string {
	return r.raw
}

func (r CustomerUnmaskedV1DataComplianceProfileIndividualComplianceProfile) implementsCustomerUnmaskedV1DataComplianceProfile() {
}

// Business registration data required to trigger Patriot Act compliant KYB
// verification.
type CustomerUnmaskedV1DataComplianceProfileBusinessComplianceProfile struct {
	// Employer Identification Number (format XX-XXXXXXX). Required for Patriot
	// Act-compliant KYB verification.
	Ein string `json:"ein,required,nullable"`
	// Official registered business name as listed with the IRS. This value will be
	// matched against the 'legal_business name'.
	LegalBusinessName string `json:"legal_business_name,required,nullable"`
	// Official business website URL. Optional but recommended for enhanced KYB.
	Website string                                                               `json:"website,nullable" format:"uri"`
	JSON    customerUnmaskedV1DataComplianceProfileBusinessComplianceProfileJSON `json:"-"`
}

// customerUnmaskedV1DataComplianceProfileBusinessComplianceProfileJSON contains
// the JSON metadata for the struct
// [CustomerUnmaskedV1DataComplianceProfileBusinessComplianceProfile]
type customerUnmaskedV1DataComplianceProfileBusinessComplianceProfileJSON struct {
	Ein               apijson.Field
	LegalBusinessName apijson.Field
	Website           apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *CustomerUnmaskedV1DataComplianceProfileBusinessComplianceProfile) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerUnmaskedV1DataComplianceProfileBusinessComplianceProfileJSON) RawJSON() string {
	return r.raw
}

func (r CustomerUnmaskedV1DataComplianceProfileBusinessComplianceProfile) implementsCustomerUnmaskedV1DataComplianceProfile() {
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

func (r CustomerUnmaskedV1ResponseType) IsKnown() bool {
	switch r {
	case CustomerUnmaskedV1ResponseTypeObject, CustomerUnmaskedV1ResponseTypeArray, CustomerUnmaskedV1ResponseTypeError, CustomerUnmaskedV1ResponseTypeNone:
		return true
	}
	return false
}

type CustomerV1 struct {
	Data CustomerV1Data `json:"data,required"`
	// Metadata about the API request, including an identifier and timestamp.
	Meta shared.ResponseMetadata `json:"meta,required"`
	// Indicates the structure of the returned content.
	//
	//   - "object" means the `data` field contains a single JSON object.
	//   - "array" means the `data` field contains an array of objects.
	//   - "error" means the `data` field contains an error object with details of the
	//     issue.
	//   - "none" means no data is returned.
	ResponseType CustomerV1ResponseType `json:"response_type,required"`
	JSON         customerV1JSON         `json:"-"`
}

// customerV1JSON contains the JSON metadata for the struct [CustomerV1]
type customerV1JSON struct {
	Data         apijson.Field
	Meta         apijson.Field
	ResponseType apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *CustomerV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerV1JSON) RawJSON() string {
	return r.raw
}

type CustomerV1Data struct {
	// Unique identifier for the customer.
	ID string `json:"id,required" format:"uuid"`
	// Timestamp of when the customer record was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The customer's email address.
	Email string `json:"email,required" format:"email"`
	// Full name of the individual or business name.
	Name string `json:"name,required"`
	// The customer's phone number in E.164 format.
	Phone  string               `json:"phone,required"`
	Status CustomerV1DataStatus `json:"status,required"`
	Type   CustomerV1DataType   `json:"type,required"`
	// Timestamp of the most recent update to the customer record.
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// An object containing the customer's address. This is optional, but if provided,
	// all required fields must be present.
	Address CustomerAddressV1 `json:"address,nullable"`
	// PII required to trigger Patriot Act compliant KYC verification.
	ComplianceProfile CustomerV1DataComplianceProfile `json:"compliance_profile,nullable"`
	Device            CustomerV1DataDevice            `json:"device"`
	// Unique identifier for the customer in your database, used for cross-referencing
	// between Straddle and your systems.
	ExternalID string `json:"external_id,nullable"`
	// Up to 20 additional user-defined key-value pairs. Useful for storing additional
	// information about the customer in a structured format.
	Metadata map[string]string  `json:"metadata,nullable"`
	JSON     customerV1DataJSON `json:"-"`
}

// customerV1DataJSON contains the JSON metadata for the struct [CustomerV1Data]
type customerV1DataJSON struct {
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

func (r *CustomerV1Data) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerV1DataJSON) RawJSON() string {
	return r.raw
}

type CustomerV1DataStatus string

const (
	CustomerV1DataStatusPending  CustomerV1DataStatus = "pending"
	CustomerV1DataStatusReview   CustomerV1DataStatus = "review"
	CustomerV1DataStatusVerified CustomerV1DataStatus = "verified"
	CustomerV1DataStatusInactive CustomerV1DataStatus = "inactive"
	CustomerV1DataStatusRejected CustomerV1DataStatus = "rejected"
)

func (r CustomerV1DataStatus) IsKnown() bool {
	switch r {
	case CustomerV1DataStatusPending, CustomerV1DataStatusReview, CustomerV1DataStatusVerified, CustomerV1DataStatusInactive, CustomerV1DataStatusRejected:
		return true
	}
	return false
}

type CustomerV1DataType string

const (
	CustomerV1DataTypeIndividual CustomerV1DataType = "individual"
	CustomerV1DataTypeBusiness   CustomerV1DataType = "business"
)

func (r CustomerV1DataType) IsKnown() bool {
	switch r {
	case CustomerV1DataTypeIndividual, CustomerV1DataTypeBusiness:
		return true
	}
	return false
}

// PII required to trigger Patriot Act compliant KYC verification.
type CustomerV1DataComplianceProfile struct {
	// Masked date of birth in \***\*-**-\*\* format.
	Dob time.Time `json:"dob,nullable" format:"date"`
	// Masked Employer Identification Number in the format **-**\*****
	Ein string `json:"ein,nullable"`
	// The official registered name of the business. This name should be correlated
	// with the `ein` value.
	LegalBusinessName string `json:"legal_business_name,nullable"`
	// Masked Social Security Number in the format **\*-**-\*\*\*\*.
	Ssn string `json:"ssn,nullable"`
	// Official business website URL. Optional but recommended for enhanced KYB.
	Website string                              `json:"website,nullable" format:"uri"`
	JSON    customerV1DataComplianceProfileJSON `json:"-"`
	union   CustomerV1DataComplianceProfileUnion
}

// customerV1DataComplianceProfileJSON contains the JSON metadata for the struct
// [CustomerV1DataComplianceProfile]
type customerV1DataComplianceProfileJSON struct {
	Dob               apijson.Field
	Ein               apijson.Field
	LegalBusinessName apijson.Field
	Ssn               apijson.Field
	Website           apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r customerV1DataComplianceProfileJSON) RawJSON() string {
	return r.raw
}

func (r *CustomerV1DataComplianceProfile) UnmarshalJSON(data []byte) (err error) {
	*r = CustomerV1DataComplianceProfile{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [CustomerV1DataComplianceProfileUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [CustomerV1DataComplianceProfileIndividualCustomerComplianceProfile],
// [CustomerV1DataComplianceProfileBusinessCustomerComplianceProfile].
func (r CustomerV1DataComplianceProfile) AsUnion() CustomerV1DataComplianceProfileUnion {
	return r.union
}

// PII required to trigger Patriot Act compliant KYC verification.
//
// Union satisfied by
// [CustomerV1DataComplianceProfileIndividualCustomerComplianceProfile] or
// [CustomerV1DataComplianceProfileBusinessCustomerComplianceProfile].
type CustomerV1DataComplianceProfileUnion interface {
	implementsCustomerV1DataComplianceProfile()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CustomerV1DataComplianceProfileUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CustomerV1DataComplianceProfileIndividualCustomerComplianceProfile{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CustomerV1DataComplianceProfileBusinessCustomerComplianceProfile{}),
		},
	)
}

// PII required to trigger Patriot Act compliant KYC verification.
type CustomerV1DataComplianceProfileIndividualCustomerComplianceProfile struct {
	// Masked date of birth in \***\*-**-\*\* format.
	Dob time.Time `json:"dob,required,nullable" format:"date"`
	// Masked Social Security Number in the format **\*-**-\*\*\*\*.
	Ssn  string                                                                 `json:"ssn,required,nullable"`
	JSON customerV1DataComplianceProfileIndividualCustomerComplianceProfileJSON `json:"-"`
}

// customerV1DataComplianceProfileIndividualCustomerComplianceProfileJSON contains
// the JSON metadata for the struct
// [CustomerV1DataComplianceProfileIndividualCustomerComplianceProfile]
type customerV1DataComplianceProfileIndividualCustomerComplianceProfileJSON struct {
	Dob         apijson.Field
	Ssn         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerV1DataComplianceProfileIndividualCustomerComplianceProfile) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerV1DataComplianceProfileIndividualCustomerComplianceProfileJSON) RawJSON() string {
	return r.raw
}

func (r CustomerV1DataComplianceProfileIndividualCustomerComplianceProfile) implementsCustomerV1DataComplianceProfile() {
}

// Business registration data required to trigger Patriot Act compliant KYB
// verification.
type CustomerV1DataComplianceProfileBusinessCustomerComplianceProfile struct {
	// Masked Employer Identification Number in the format **-**\*****
	Ein string `json:"ein,required,nullable"`
	// The official registered name of the business. This name should be correlated
	// with the `ein` value.
	LegalBusinessName string `json:"legal_business_name,required,nullable"`
	// Official business website URL. Optional but recommended for enhanced KYB.
	Website string                                                               `json:"website,nullable" format:"uri"`
	JSON    customerV1DataComplianceProfileBusinessCustomerComplianceProfileJSON `json:"-"`
}

// customerV1DataComplianceProfileBusinessCustomerComplianceProfileJSON contains
// the JSON metadata for the struct
// [CustomerV1DataComplianceProfileBusinessCustomerComplianceProfile]
type customerV1DataComplianceProfileBusinessCustomerComplianceProfileJSON struct {
	Ein               apijson.Field
	LegalBusinessName apijson.Field
	Website           apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *CustomerV1DataComplianceProfileBusinessCustomerComplianceProfile) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerV1DataComplianceProfileBusinessCustomerComplianceProfileJSON) RawJSON() string {
	return r.raw
}

func (r CustomerV1DataComplianceProfileBusinessCustomerComplianceProfile) implementsCustomerV1DataComplianceProfile() {
}

type CustomerV1DataDevice struct {
	// The customer's IP address at the time of profile creation. Use `0.0.0.0` to
	// represent an offline customer registration.
	IPAddress string                   `json:"ip_address,required" format:"ipv4"`
	JSON      customerV1DataDeviceJSON `json:"-"`
}

// customerV1DataDeviceJSON contains the JSON metadata for the struct
// [CustomerV1DataDevice]
type customerV1DataDeviceJSON struct {
	IPAddress   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerV1DataDevice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerV1DataDeviceJSON) RawJSON() string {
	return r.raw
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

func (r CustomerV1ResponseType) IsKnown() bool {
	switch r {
	case CustomerV1ResponseTypeObject, CustomerV1ResponseTypeArray, CustomerV1ResponseTypeError, CustomerV1ResponseTypeNone:
		return true
	}
	return false
}

type DeviceUnmaskedV1 struct {
	// The customer's IP address at the time of profile creation. Use `0.0.0.0` to
	// represent an offline customer registration.
	IPAddress string               `json:"ip_address,required" format:"ipv4"`
	JSON      deviceUnmaskedV1JSON `json:"-"`
}

// deviceUnmaskedV1JSON contains the JSON metadata for the struct
// [DeviceUnmaskedV1]
type deviceUnmaskedV1JSON struct {
	IPAddress   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceUnmaskedV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceUnmaskedV1JSON) RawJSON() string {
	return r.raw
}

type DeviceUnmaskedV1Param struct {
	// The customer's IP address at the time of profile creation. Use `0.0.0.0` to
	// represent an offline customer registration.
	IPAddress param.Field[string] `json:"ip_address,required" format:"ipv4"`
}

func (r DeviceUnmaskedV1Param) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CustomerNewParams struct {
	Device param.Field[DeviceUnmaskedV1Param] `json:"device,required"`
	// The customer's email address.
	Email param.Field[string] `json:"email,required" format:"email"`
	// Full name of the individual or business name.
	Name param.Field[string] `json:"name,required"`
	// The customer's phone number in E.164 format. Mobile number is preferred.
	Phone param.Field[string]                `json:"phone,required"`
	Type  param.Field[CustomerNewParamsType] `json:"type,required"`
	// An object containing the customer's address. This is optional, but if provided,
	// all required fields must be present.
	Address param.Field[CustomerAddressV1Param] `json:"address"`
	// An object containing the customer's compliance profile. **This is optional**,
	// but if provided, all required fields must be present for the appropriate
	// customer type.
	ComplianceProfile param.Field[CustomerNewParamsComplianceProfileUnion] `json:"compliance_profile"`
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

type CustomerNewParamsType string

const (
	CustomerNewParamsTypeIndividual CustomerNewParamsType = "individual"
	CustomerNewParamsTypeBusiness   CustomerNewParamsType = "business"
)

func (r CustomerNewParamsType) IsKnown() bool {
	switch r {
	case CustomerNewParamsTypeIndividual, CustomerNewParamsTypeBusiness:
		return true
	}
	return false
}

// An object containing the customer's compliance profile. **This is optional**,
// but if provided, all required fields must be present for the appropriate
// customer type.
type CustomerNewParamsComplianceProfile struct {
	// Date of birth (YYYY-MM-DD). Required for Patriot Act-compliant KYC verification.
	Dob param.Field[time.Time] `json:"dob" format:"date"`
	// Employer Identification Number (format XX-XXXXXXX). Required for Patriot
	// Act-compliant KYB verification.
	Ein param.Field[string] `json:"ein"`
	// Official registered business name as listed with the IRS. This value will be
	// matched against the 'legal_business name'.
	LegalBusinessName param.Field[string] `json:"legal_business_name"`
	// Social Security Number (format XXX-XX-XXXX). Required for Patriot Act-compliant
	// KYC verification.
	Ssn param.Field[string] `json:"ssn"`
	// Official business website URL. Optional but recommended for enhanced KYB.
	Website param.Field[string] `json:"website" format:"uri"`
}

func (r CustomerNewParamsComplianceProfile) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CustomerNewParamsComplianceProfile) implementsCustomerNewParamsComplianceProfileUnion() {}

// An object containing the customer's compliance profile. **This is optional**,
// but if provided, all required fields must be present for the appropriate
// customer type.
//
// Satisfied by [CustomerNewParamsComplianceProfileIndividualComplianceProfile],
// [CustomerNewParamsComplianceProfileBusinessComplianceProfile],
// [CustomerNewParamsComplianceProfile].
type CustomerNewParamsComplianceProfileUnion interface {
	implementsCustomerNewParamsComplianceProfileUnion()
}

// Individual PII data required to trigger Patriot Act compliant KYC verification.
type CustomerNewParamsComplianceProfileIndividualComplianceProfile struct {
	// Date of birth (YYYY-MM-DD). Required for Patriot Act-compliant KYC verification.
	Dob param.Field[time.Time] `json:"dob,required" format:"date"`
	// Social Security Number (format XXX-XX-XXXX). Required for Patriot Act-compliant
	// KYC verification.
	Ssn param.Field[string] `json:"ssn,required"`
}

func (r CustomerNewParamsComplianceProfileIndividualComplianceProfile) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CustomerNewParamsComplianceProfileIndividualComplianceProfile) implementsCustomerNewParamsComplianceProfileUnion() {
}

// Business registration data required to trigger Patriot Act compliant KYB
// verification.
type CustomerNewParamsComplianceProfileBusinessComplianceProfile struct {
	// Employer Identification Number (format XX-XXXXXXX). Required for Patriot
	// Act-compliant KYB verification.
	Ein param.Field[string] `json:"ein,required"`
	// Official registered business name as listed with the IRS. This value will be
	// matched against the 'legal_business name'.
	LegalBusinessName param.Field[string] `json:"legal_business_name,required"`
	// Official business website URL. Optional but recommended for enhanced KYB.
	Website param.Field[string] `json:"website" format:"uri"`
}

func (r CustomerNewParamsComplianceProfileBusinessComplianceProfile) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CustomerNewParamsComplianceProfileBusinessComplianceProfile) implementsCustomerNewParamsComplianceProfileUnion() {
}

type CustomerUpdateParams struct {
	Device param.Field[DeviceUnmaskedV1Param] `json:"device,required"`
	// The customer's email address.
	Email param.Field[string] `json:"email,required" format:"email"`
	// The customer's full name or business name.
	Name param.Field[string] `json:"name,required"`
	// The customer's phone number in E.164 format.
	Phone  param.Field[string]                     `json:"phone,required"`
	Status param.Field[CustomerUpdateParamsStatus] `json:"status,required"`
	// An object containing the customer's address. This is optional, but if provided,
	// all required fields must be present.
	Address param.Field[CustomerAddressV1Param] `json:"address"`
	// Individual PII data required to trigger Patriot Act compliant KYC verification.
	ComplianceProfile param.Field[CustomerUpdateParamsComplianceProfileUnion] `json:"compliance_profile"`
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

// Individual PII data required to trigger Patriot Act compliant KYC verification.
type CustomerUpdateParamsComplianceProfile struct {
	// Date of birth (YYYY-MM-DD). Required for Patriot Act-compliant KYC verification.
	Dob param.Field[time.Time] `json:"dob" format:"date"`
	// Employer Identification Number (format XX-XXXXXXX). Required for Patriot
	// Act-compliant KYB verification.
	Ein param.Field[string] `json:"ein"`
	// Official registered business name as listed with the IRS. This value will be
	// matched against the 'legal_business name'.
	LegalBusinessName param.Field[string] `json:"legal_business_name"`
	// Social Security Number (format XXX-XX-XXXX). Required for Patriot Act-compliant
	// KYC verification.
	Ssn param.Field[string] `json:"ssn"`
	// Official business website URL. Optional but recommended for enhanced KYB.
	Website param.Field[string] `json:"website" format:"uri"`
}

func (r CustomerUpdateParamsComplianceProfile) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CustomerUpdateParamsComplianceProfile) implementsCustomerUpdateParamsComplianceProfileUnion() {
}

// Individual PII data required to trigger Patriot Act compliant KYC verification.
//
// Satisfied by [CustomerUpdateParamsComplianceProfileIndividualComplianceProfile],
// [CustomerUpdateParamsComplianceProfileBusinessComplianceProfile],
// [CustomerUpdateParamsComplianceProfile].
type CustomerUpdateParamsComplianceProfileUnion interface {
	implementsCustomerUpdateParamsComplianceProfileUnion()
}

// Individual PII data required to trigger Patriot Act compliant KYC verification.
type CustomerUpdateParamsComplianceProfileIndividualComplianceProfile struct {
	// Date of birth (YYYY-MM-DD). Required for Patriot Act-compliant KYC verification.
	Dob param.Field[time.Time] `json:"dob,required" format:"date"`
	// Social Security Number (format XXX-XX-XXXX). Required for Patriot Act-compliant
	// KYC verification.
	Ssn param.Field[string] `json:"ssn,required"`
}

func (r CustomerUpdateParamsComplianceProfileIndividualComplianceProfile) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CustomerUpdateParamsComplianceProfileIndividualComplianceProfile) implementsCustomerUpdateParamsComplianceProfileUnion() {
}

// Business registration data required to trigger Patriot Act compliant KYB
// verification.
type CustomerUpdateParamsComplianceProfileBusinessComplianceProfile struct {
	// Employer Identification Number (format XX-XXXXXXX). Required for Patriot
	// Act-compliant KYB verification.
	Ein param.Field[string] `json:"ein,required"`
	// Official registered business name as listed with the IRS. This value will be
	// matched against the 'legal_business name'.
	LegalBusinessName param.Field[string] `json:"legal_business_name,required"`
	// Official business website URL. Optional but recommended for enhanced KYB.
	Website param.Field[string] `json:"website" format:"uri"`
}

func (r CustomerUpdateParamsComplianceProfileBusinessComplianceProfile) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CustomerUpdateParamsComplianceProfileBusinessComplianceProfile) implementsCustomerUpdateParamsComplianceProfileUnion() {
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
	SearchText param.Field[string]                      `query:"search_text"`
	SortBy     param.Field[CustomerListParamsSortBy]    `query:"sort_by"`
	SortOrder  param.Field[CustomerListParamsSortOrder] `query:"sort_order"`
	// Filter customers by their current `status`.
	Status param.Field[[]CustomerListParamsStatus] `query:"status"`
	// Filter by customer type `individual` or `business`.
	Types             param.Field[[]CustomerListParamsType] `query:"types"`
	CorrelationID     param.Field[string]                   `header:"Correlation-Id"`
	RequestID         param.Field[string]                   `header:"Request-Id"`
	StraddleAccountID param.Field[string]                   `header:"Straddle-Account-Id" format:"uuid"`
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

type CustomerListParamsSortOrder string

const (
	CustomerListParamsSortOrderAsc  CustomerListParamsSortOrder = "asc"
	CustomerListParamsSortOrderDesc CustomerListParamsSortOrder = "desc"
)

func (r CustomerListParamsSortOrder) IsKnown() bool {
	switch r {
	case CustomerListParamsSortOrderAsc, CustomerListParamsSortOrderDesc:
		return true
	}
	return false
}

type CustomerListParamsStatus string

const (
	CustomerListParamsStatusPending  CustomerListParamsStatus = "pending"
	CustomerListParamsStatusReview   CustomerListParamsStatus = "review"
	CustomerListParamsStatusVerified CustomerListParamsStatus = "verified"
	CustomerListParamsStatusInactive CustomerListParamsStatus = "inactive"
	CustomerListParamsStatusRejected CustomerListParamsStatus = "rejected"
)

func (r CustomerListParamsStatus) IsKnown() bool {
	switch r {
	case CustomerListParamsStatusPending, CustomerListParamsStatusReview, CustomerListParamsStatusVerified, CustomerListParamsStatusInactive, CustomerListParamsStatusRejected:
		return true
	}
	return false
}

type CustomerListParamsType string

const (
	CustomerListParamsTypeIndividual CustomerListParamsType = "individual"
	CustomerListParamsTypeBusiness   CustomerListParamsType = "business"
)

func (r CustomerListParamsType) IsKnown() bool {
	switch r {
	case CustomerListParamsTypeIndividual, CustomerListParamsTypeBusiness:
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

type CustomerRefreshReviewParams struct {
	CorrelationID     param.Field[string] `header:"Correlation-Id"`
	RequestID         param.Field[string] `header:"Request-Id"`
	StraddleAccountID param.Field[string] `header:"Straddle-Account-Id" format:"uuid"`
}

type CustomerUnmaskedParams struct {
	CorrelationID     param.Field[string] `header:"Correlation-Id"`
	RequestID         param.Field[string] `header:"Request-Id"`
	StraddleAccountID param.Field[string] `header:"Straddle-Account-Id" format:"uuid"`
}
