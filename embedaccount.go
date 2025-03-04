// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package straddle

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/stainless-sdks/straddle-go/internal/apijson"
	"github.com/stainless-sdks/straddle-go/internal/apiquery"
	"github.com/stainless-sdks/straddle-go/internal/param"
	"github.com/stainless-sdks/straddle-go/internal/requestconfig"
	"github.com/stainless-sdks/straddle-go/option"
	"github.com/stainless-sdks/straddle-go/packages/pagination"
	"github.com/stainless-sdks/straddle-go/shared"
)

// EmbedAccountService contains methods and other services that help with
// interacting with the straddle API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEmbedAccountService] method instead.
type EmbedAccountService struct {
	Options            []option.RequestOption
	CapabilityRequests *EmbedAccountCapabilityRequestService
}

// NewEmbedAccountService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewEmbedAccountService(opts ...option.RequestOption) (r *EmbedAccountService) {
	r = &EmbedAccountService{}
	r.Options = opts
	r.CapabilityRequests = NewEmbedAccountCapabilityRequestService(opts...)
	return
}

// Creates a new account associated with your Straddle platform integration. This
// endpoint allows you to set up an account with specified details, including
// business information and access levels.
func (r *EmbedAccountService) New(ctx context.Context, params EmbedAccountNewParams, opts ...option.RequestOption) (res *shared.ItemResponseOfAccountV1, err error) {
	if params.CorrelationID.Present {
		opts = append(opts, option.WithHeader("correlation-id", fmt.Sprintf("%s", params.CorrelationID)))
	}
	if params.RequestID.Present {
		opts = append(opts, option.WithHeader("request-id", fmt.Sprintf("%s", params.RequestID)))
	}
	opts = append(r.Options[:], opts...)
	path := "v1/accounts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Updates an existing account's information. This endpoint allows you to update
// various account details during onboarding or after the account has been created.
func (r *EmbedAccountService) Update(ctx context.Context, accountID string, params EmbedAccountUpdateParams, opts ...option.RequestOption) (res *shared.ItemResponseOfAccountV1, err error) {
	if params.CorrelationID.Present {
		opts = append(opts, option.WithHeader("correlation-id", fmt.Sprintf("%s", params.CorrelationID)))
	}
	if params.RequestID.Present {
		opts = append(opts, option.WithHeader("request-id", fmt.Sprintf("%s", params.RequestID)))
	}
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("v1/accounts/%s", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

// Returns a list of accounts associated with your Straddle platform integration.
// The accounts are returned sorted by creation date, with the most recently
// created accounts appearing first. This endpoint supports advanced sorting and
// filtering options.
func (r *EmbedAccountService) List(ctx context.Context, params EmbedAccountListParams, opts ...option.RequestOption) (res *pagination.PageNumberSchema[shared.AccountV1], err error) {
	var raw *http.Response
	if params.CorrelationID.Present {
		opts = append(opts, option.WithHeader("correlation-id", fmt.Sprintf("%s", params.CorrelationID)))
	}
	if params.RequestID.Present {
		opts = append(opts, option.WithHeader("request-id", fmt.Sprintf("%s", params.RequestID)))
	}
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/accounts"
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

// Returns a list of accounts associated with your Straddle platform integration.
// The accounts are returned sorted by creation date, with the most recently
// created accounts appearing first. This endpoint supports advanced sorting and
// filtering options.
func (r *EmbedAccountService) ListAutoPaging(ctx context.Context, params EmbedAccountListParams, opts ...option.RequestOption) *pagination.PageNumberSchemaAutoPager[shared.AccountV1] {
	return pagination.NewPageNumberSchemaAutoPager(r.List(ctx, params, opts...))
}

// Retrieves the details of an account that has previously been created. Supply the
// unique account ID that was returned from your previous request, and Straddle
// will return the corresponding account information.
func (r *EmbedAccountService) Get(ctx context.Context, accountID string, query EmbedAccountGetParams, opts ...option.RequestOption) (res *shared.ItemResponseOfAccountV1, err error) {
	if query.CorrelationID.Present {
		opts = append(opts, option.WithHeader("correlation-id", fmt.Sprintf("%s", query.CorrelationID)))
	}
	if query.RequestID.Present {
		opts = append(opts, option.WithHeader("request-id", fmt.Sprintf("%s", query.RequestID)))
	}
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("v1/accounts/%s", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Initiates the onboarding process for a new account. This endpoint can only be
// used for accounts where at least one representative and one bank account have
// already been created.
func (r *EmbedAccountService) Onboard(ctx context.Context, accountID string, params EmbedAccountOnboardParams, opts ...option.RequestOption) (res *shared.ItemResponseOfAccountV1, err error) {
	if params.CorrelationID.Present {
		opts = append(opts, option.WithHeader("correlation-id", fmt.Sprintf("%s", params.CorrelationID)))
	}
	if params.RequestID.Present {
		opts = append(opts, option.WithHeader("request-id", fmt.Sprintf("%s", params.RequestID)))
	}
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("v1/accounts/%s/onboard", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Simulte the status transitions for sandbox accounts. This endpoint can only be
// used for sandbox accounts.
func (r *EmbedAccountService) Simulate(ctx context.Context, accountID string, params EmbedAccountSimulateParams, opts ...option.RequestOption) (res *shared.ItemResponseOfAccountV1, err error) {
	if params.CorrelationID.Present {
		opts = append(opts, option.WithHeader("correlation-id", fmt.Sprintf("%s", params.CorrelationID)))
	}
	if params.RequestID.Present {
		opts = append(opts, option.WithHeader("request-id", fmt.Sprintf("%s", params.RequestID)))
	}
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("v1/accounts/%s/simulate", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type AccountPagedV1 struct {
	Data []shared.AccountV1 `json:"data,required"`
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
	ResponseType AccountPagedV1ResponseType `json:"response_type,required"`
	JSON         accountPagedV1JSON         `json:"-"`
}

// accountPagedV1JSON contains the JSON metadata for the struct [AccountPagedV1]
type accountPagedV1JSON struct {
	Data         apijson.Field
	Meta         apijson.Field
	ResponseType apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountPagedV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountPagedV1JSON) RawJSON() string {
	return r.raw
}

// Indicates the structure of the returned content.
//
//   - "object" means the `data` field contains a single JSON object.
//   - "array" means the `data` field contains an array of objects.
//   - "error" means the `data` field contains an error object with details of the
//     issue.
//   - "none" means no data is returned.
type AccountPagedV1ResponseType string

const (
	AccountPagedV1ResponseTypeObject AccountPagedV1ResponseType = "object"
	AccountPagedV1ResponseTypeArray  AccountPagedV1ResponseType = "array"
	AccountPagedV1ResponseTypeError  AccountPagedV1ResponseType = "error"
	AccountPagedV1ResponseTypeNone   AccountPagedV1ResponseType = "none"
)

func (r AccountPagedV1ResponseType) IsKnown() bool {
	switch r {
	case AccountPagedV1ResponseTypeObject, AccountPagedV1ResponseTypeArray, AccountPagedV1ResponseTypeError, AccountPagedV1ResponseTypeNone:
		return true
	}
	return false
}

// The address object is optional. If provided, it must be a valid address.
type AddressV1 struct {
	// City, district, suburb, town, or village.
	City string `json:"city,nullable"`
	// The country of the address, in ISO 3166-1 alpha-2 format.
	Country string `json:"country,nullable"`
	// Primary address line (e.g., street, PO Box).
	Line1 string `json:"line1,nullable"`
	// Secondary address line (e.g., apartment, suite, unit, or building).
	Line2 string `json:"line2,nullable"`
	// Postal or ZIP code.
	PostalCode string `json:"postal_code,nullable"`
	// Two-letter state code.
	State string        `json:"state,nullable"`
	JSON  addressV1JSON `json:"-"`
}

// addressV1JSON contains the JSON metadata for the struct [AddressV1]
type addressV1JSON struct {
	City        apijson.Field
	Country     apijson.Field
	Line1       apijson.Field
	Line2       apijson.Field
	PostalCode  apijson.Field
	State       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressV1JSON) RawJSON() string {
	return r.raw
}

// The address object is optional. If provided, it must be a valid address.
type AddressV1Param struct {
	// City, district, suburb, town, or village.
	City param.Field[string] `json:"city"`
	// The country of the address, in ISO 3166-1 alpha-2 format.
	Country param.Field[string] `json:"country"`
	// Primary address line (e.g., street, PO Box).
	Line1 param.Field[string] `json:"line1"`
	// Secondary address line (e.g., apartment, suite, unit, or building).
	Line2 param.Field[string] `json:"line2"`
	// Postal or ZIP code.
	PostalCode param.Field[string] `json:"postal_code"`
	// Two-letter state code.
	State param.Field[string] `json:"state"`
}

func (r AddressV1Param) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type IndustryV1 struct {
	// The general category of the industry. Required if not providing MCC.
	Category string `json:"category,nullable"`
	// The Merchant Category Code (MCC) that best describes the business. Optional.
	Mcc string `json:"mcc,nullable"`
	// The specific sector within the industry category. Required if not providing MCC.
	Sector string         `json:"sector,nullable"`
	JSON   industryV1JSON `json:"-"`
}

// industryV1JSON contains the JSON metadata for the struct [IndustryV1]
type industryV1JSON struct {
	Category    apijson.Field
	Mcc         apijson.Field
	Sector      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IndustryV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r industryV1JSON) RawJSON() string {
	return r.raw
}

type IndustryV1Param struct {
	// The general category of the industry. Required if not providing MCC.
	Category param.Field[string] `json:"category"`
	// The Merchant Category Code (MCC) that best describes the business. Optional.
	Mcc param.Field[string] `json:"mcc"`
	// The specific sector within the industry category. Required if not providing MCC.
	Sector param.Field[string] `json:"sector"`
}

func (r IndustryV1Param) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SupportChannelsV1 struct {
	// The email address for customer support inquiries.
	Email string `json:"email,nullable" format:"email"`
	// The phone number for customer support.
	Phone string `json:"phone,nullable"`
	// The URL of the business's customer support page or contact form.
	URL  string                `json:"url,nullable" format:"uri"`
	JSON supportChannelsV1JSON `json:"-"`
}

// supportChannelsV1JSON contains the JSON metadata for the struct
// [SupportChannelsV1]
type supportChannelsV1JSON struct {
	Email       apijson.Field
	Phone       apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SupportChannelsV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r supportChannelsV1JSON) RawJSON() string {
	return r.raw
}

type SupportChannelsV1Param struct {
	// The email address for customer support inquiries.
	Email param.Field[string] `json:"email" format:"email"`
	// The phone number for customer support.
	Phone param.Field[string] `json:"phone"`
	// The URL of the business's customer support page or contact form.
	URL param.Field[string] `json:"url" format:"uri"`
}

func (r SupportChannelsV1Param) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EmbedAccountNewParams struct {
	// The access level granted to the account. This is determined by your platform
	// configuration. Use `standard` unless instructed otherwise by Straddle.
	AccessLevel param.Field[EmbedAccountNewParamsAccessLevel] `json:"access_level,required"`
	// The type of account to be created. Currently, only `business` is supported.
	AccountType     param.Field[EmbedAccountNewParamsAccountType] `json:"account_type,required"`
	BusinessProfile param.Field[shared.BusinessProfileV1Param]    `json:"business_profile,required"`
	// The unique identifier of the organization related to this account.
	OrganizationID param.Field[string] `json:"organization_id,required" format:"uuid"`
	// Unique identifier for the account in your database, used for cross-referencing
	// between Straddle and your systems.
	ExternalID param.Field[string] `json:"external_id"`
	// Up to 20 additional user-defined key-value pairs. Useful for storing additional
	// information about the account in a structured format.
	Metadata      param.Field[map[string]string] `json:"metadata"`
	CorrelationID param.Field[string]            `header:"correlation-id"`
	RequestID     param.Field[string]            `header:"request-id"`
}

func (r EmbedAccountNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The access level granted to the account. This is determined by your platform
// configuration. Use `standard` unless instructed otherwise by Straddle.
type EmbedAccountNewParamsAccessLevel string

const (
	EmbedAccountNewParamsAccessLevelStandard EmbedAccountNewParamsAccessLevel = "standard"
	EmbedAccountNewParamsAccessLevelManaged  EmbedAccountNewParamsAccessLevel = "managed"
)

func (r EmbedAccountNewParamsAccessLevel) IsKnown() bool {
	switch r {
	case EmbedAccountNewParamsAccessLevelStandard, EmbedAccountNewParamsAccessLevelManaged:
		return true
	}
	return false
}

// The type of account to be created. Currently, only `business` is supported.
type EmbedAccountNewParamsAccountType string

const (
	EmbedAccountNewParamsAccountTypeBusiness EmbedAccountNewParamsAccountType = "business"
)

func (r EmbedAccountNewParamsAccountType) IsKnown() bool {
	switch r {
	case EmbedAccountNewParamsAccountTypeBusiness:
		return true
	}
	return false
}

type EmbedAccountUpdateParams struct {
	BusinessProfile param.Field[shared.BusinessProfileV1Param] `json:"business_profile,required"`
	// Unique identifier for the account in your database, used for cross-referencing
	// between Straddle and your systems.
	ExternalID param.Field[string] `json:"external_id"`
	// Up to 20 additional user-defined key-value pairs. Useful for storing additional
	// information about the account in a structured format.
	Metadata      param.Field[map[string]string] `json:"metadata"`
	CorrelationID param.Field[string]            `header:"correlation-id"`
	RequestID     param.Field[string]            `header:"request-id"`
}

func (r EmbedAccountUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EmbedAccountListParams struct {
	// Results page number. Starts at page 1. Default value: 1
	PageNumber param.Field[int64] `query:"page_number"`
	// Page size. Default value: 100. Max value: 1000
	PageSize   param.Field[int64]  `query:"page_size"`
	SearchText param.Field[string] `query:"search_text"`
	// Sort By. Default value: 'id'.
	SortBy param.Field[string] `query:"sort_by"`
	// Sort Order. Default value: 'asc'.
	SortOrder     param.Field[EmbedAccountListParamsSortOrder] `query:"sort_order"`
	CorrelationID param.Field[string]                          `header:"correlation-id"`
	RequestID     param.Field[string]                          `header:"request-id"`
}

// URLQuery serializes [EmbedAccountListParams]'s query parameters as `url.Values`.
func (r EmbedAccountListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Sort Order. Default value: 'asc'.
type EmbedAccountListParamsSortOrder string

const (
	EmbedAccountListParamsSortOrderAsc  EmbedAccountListParamsSortOrder = "asc"
	EmbedAccountListParamsSortOrderDesc EmbedAccountListParamsSortOrder = "desc"
)

func (r EmbedAccountListParamsSortOrder) IsKnown() bool {
	switch r {
	case EmbedAccountListParamsSortOrderAsc, EmbedAccountListParamsSortOrderDesc:
		return true
	}
	return false
}

type EmbedAccountGetParams struct {
	CorrelationID param.Field[string] `header:"correlation-id"`
	RequestID     param.Field[string] `header:"request-id"`
}

type EmbedAccountOnboardParams struct {
	TermsOfService param.Field[shared.TermsOfServiceV1Param] `json:"terms_of_service,required"`
	CorrelationID  param.Field[string]                       `header:"correlation-id"`
	RequestID      param.Field[string]                       `header:"request-id"`
}

func (r EmbedAccountOnboardParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EmbedAccountSimulateParams struct {
	FinalStatus   param.Field[EmbedAccountSimulateParamsFinalStatus] `query:"final_status"`
	CorrelationID param.Field[string]                                `header:"correlation-id"`
	RequestID     param.Field[string]                                `header:"request-id"`
}

// URLQuery serializes [EmbedAccountSimulateParams]'s query parameters as
// `url.Values`.
func (r EmbedAccountSimulateParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type EmbedAccountSimulateParamsFinalStatus string

const (
	EmbedAccountSimulateParamsFinalStatusOnboarding EmbedAccountSimulateParamsFinalStatus = "onboarding"
	EmbedAccountSimulateParamsFinalStatusActive     EmbedAccountSimulateParamsFinalStatus = "active"
)

func (r EmbedAccountSimulateParamsFinalStatus) IsKnown() bool {
	switch r {
	case EmbedAccountSimulateParamsFinalStatusOnboarding, EmbedAccountSimulateParamsFinalStatusActive:
		return true
	}
	return false
}
