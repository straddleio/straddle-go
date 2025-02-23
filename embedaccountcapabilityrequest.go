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

// EmbedAccountCapabilityRequestService contains methods and other services that
// help with interacting with the straddle API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEmbedAccountCapabilityRequestService] method instead.
type EmbedAccountCapabilityRequestService struct {
	Options []option.RequestOption
}

// NewEmbedAccountCapabilityRequestService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewEmbedAccountCapabilityRequestService(opts ...option.RequestOption) (r *EmbedAccountCapabilityRequestService) {
	r = &EmbedAccountCapabilityRequestService{}
	r.Options = opts
	return
}

// Submits a request to enable a specific capability for an account. Use this
// endpoint to request additional features or services for an account.
func (r *EmbedAccountCapabilityRequestService) New(ctx context.Context, accountID string, params EmbedAccountCapabilityRequestNewParams, opts ...option.RequestOption) (res *CapabilityRequestPagedV1, err error) {
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
	path := fmt.Sprintf("v1/accounts/%s/capability_requests", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Retrieves a list of capability requests associated with an account. The requests
// are returned sorted by creation date, with the most recent requests appearing
// first. This endpoint supports advanced sorting and filtering options.
func (r *EmbedAccountCapabilityRequestService) List(ctx context.Context, accountID string, params EmbedAccountCapabilityRequestListParams, opts ...option.RequestOption) (res *pagination.PageNumberSchema[CapabilityRequestPagedV1Data], err error) {
	var raw *http.Response
	if params.CorrelationID.Present {
		opts = append(opts, option.WithHeader("correlation-id", fmt.Sprintf("%s", params.CorrelationID)))
	}
	if params.RequestID.Present {
		opts = append(opts, option.WithHeader("request-id", fmt.Sprintf("%s", params.RequestID)))
	}
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("v1/accounts/%s/capability_requests", accountID)
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

// Retrieves a list of capability requests associated with an account. The requests
// are returned sorted by creation date, with the most recent requests appearing
// first. This endpoint supports advanced sorting and filtering options.
func (r *EmbedAccountCapabilityRequestService) ListAutoPaging(ctx context.Context, accountID string, params EmbedAccountCapabilityRequestListParams, opts ...option.RequestOption) *pagination.PageNumberSchemaAutoPager[CapabilityRequestPagedV1Data] {
	return pagination.NewPageNumberSchemaAutoPager(r.List(ctx, accountID, params, opts...))
}

type CapabilityRequestPagedV1 struct {
	Data []CapabilityRequestPagedV1Data `json:"data,required"`
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
	ResponseType CapabilityRequestPagedV1ResponseType `json:"response_type,required"`
	JSON         capabilityRequestPagedV1JSON         `json:"-"`
}

// capabilityRequestPagedV1JSON contains the JSON metadata for the struct
// [CapabilityRequestPagedV1]
type capabilityRequestPagedV1JSON struct {
	Data         apijson.Field
	Meta         apijson.Field
	ResponseType apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *CapabilityRequestPagedV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r capabilityRequestPagedV1JSON) RawJSON() string {
	return r.raw
}

type CapabilityRequestPagedV1Data struct {
	// Unique identifier for the capability request.
	ID string `json:"id,required" format:"uuid"`
	// The unique identifier of the account associated with this capability request.
	AccountID string `json:"account_id,required" format:"uuid"`
	// The category of the requested capability. Use `payment_type` for charges and
	// payouts, `customer_type` to define `individuals` or `businesses`, and
	// `consent_type` for `signed_agreement` or `internet` payment authorization.
	Category CapabilityRequestPagedV1DataCategory `json:"category,required"`
	// Timestamp of when the capability request was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The current status of the capability request.
	Status CapabilityRequestPagedV1DataStatus `json:"status,required"`
	// The specific type of capability being requested within the category.
	Type CapabilityRequestPagedV1DataType `json:"type,required"`
	// Timestamp of the most recent update to the capability request.
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// Any specific settings or configurations related to the requested capability.
	Settings map[string]interface{}           `json:"settings,nullable"`
	JSON     capabilityRequestPagedV1DataJSON `json:"-"`
}

// capabilityRequestPagedV1DataJSON contains the JSON metadata for the struct
// [CapabilityRequestPagedV1Data]
type capabilityRequestPagedV1DataJSON struct {
	ID          apijson.Field
	AccountID   apijson.Field
	Category    apijson.Field
	CreatedAt   apijson.Field
	Status      apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	Settings    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CapabilityRequestPagedV1Data) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r capabilityRequestPagedV1DataJSON) RawJSON() string {
	return r.raw
}

// The category of the requested capability. Use `payment_type` for charges and
// payouts, `customer_type` to define `individuals` or `businesses`, and
// `consent_type` for `signed_agreement` or `internet` payment authorization.
type CapabilityRequestPagedV1DataCategory string

const (
	CapabilityRequestPagedV1DataCategoryPaymentType  CapabilityRequestPagedV1DataCategory = "payment_type"
	CapabilityRequestPagedV1DataCategoryCustomerType CapabilityRequestPagedV1DataCategory = "customer_type"
	CapabilityRequestPagedV1DataCategoryConsentType  CapabilityRequestPagedV1DataCategory = "consent_type"
)

func (r CapabilityRequestPagedV1DataCategory) IsKnown() bool {
	switch r {
	case CapabilityRequestPagedV1DataCategoryPaymentType, CapabilityRequestPagedV1DataCategoryCustomerType, CapabilityRequestPagedV1DataCategoryConsentType:
		return true
	}
	return false
}

// The current status of the capability request.
type CapabilityRequestPagedV1DataStatus string

const (
	CapabilityRequestPagedV1DataStatusActive    CapabilityRequestPagedV1DataStatus = "active"
	CapabilityRequestPagedV1DataStatusInactive  CapabilityRequestPagedV1DataStatus = "inactive"
	CapabilityRequestPagedV1DataStatusInReview  CapabilityRequestPagedV1DataStatus = "in_review"
	CapabilityRequestPagedV1DataStatusRejected  CapabilityRequestPagedV1DataStatus = "rejected"
	CapabilityRequestPagedV1DataStatusApproved  CapabilityRequestPagedV1DataStatus = "approved"
	CapabilityRequestPagedV1DataStatusReviewing CapabilityRequestPagedV1DataStatus = "reviewing"
)

func (r CapabilityRequestPagedV1DataStatus) IsKnown() bool {
	switch r {
	case CapabilityRequestPagedV1DataStatusActive, CapabilityRequestPagedV1DataStatusInactive, CapabilityRequestPagedV1DataStatusInReview, CapabilityRequestPagedV1DataStatusRejected, CapabilityRequestPagedV1DataStatusApproved, CapabilityRequestPagedV1DataStatusReviewing:
		return true
	}
	return false
}

// The specific type of capability being requested within the category.
type CapabilityRequestPagedV1DataType string

const (
	CapabilityRequestPagedV1DataTypeCharges         CapabilityRequestPagedV1DataType = "charges"
	CapabilityRequestPagedV1DataTypePayouts         CapabilityRequestPagedV1DataType = "payouts"
	CapabilityRequestPagedV1DataTypeIndividuals     CapabilityRequestPagedV1DataType = "individuals"
	CapabilityRequestPagedV1DataTypeBusinesses      CapabilityRequestPagedV1DataType = "businesses"
	CapabilityRequestPagedV1DataTypeSignedAgreement CapabilityRequestPagedV1DataType = "signed_agreement"
	CapabilityRequestPagedV1DataTypeInternet        CapabilityRequestPagedV1DataType = "internet"
)

func (r CapabilityRequestPagedV1DataType) IsKnown() bool {
	switch r {
	case CapabilityRequestPagedV1DataTypeCharges, CapabilityRequestPagedV1DataTypePayouts, CapabilityRequestPagedV1DataTypeIndividuals, CapabilityRequestPagedV1DataTypeBusinesses, CapabilityRequestPagedV1DataTypeSignedAgreement, CapabilityRequestPagedV1DataTypeInternet:
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
type CapabilityRequestPagedV1ResponseType string

const (
	CapabilityRequestPagedV1ResponseTypeObject CapabilityRequestPagedV1ResponseType = "object"
	CapabilityRequestPagedV1ResponseTypeArray  CapabilityRequestPagedV1ResponseType = "array"
	CapabilityRequestPagedV1ResponseTypeError  CapabilityRequestPagedV1ResponseType = "error"
	CapabilityRequestPagedV1ResponseTypeNone   CapabilityRequestPagedV1ResponseType = "none"
)

func (r CapabilityRequestPagedV1ResponseType) IsKnown() bool {
	switch r {
	case CapabilityRequestPagedV1ResponseTypeObject, CapabilityRequestPagedV1ResponseTypeArray, CapabilityRequestPagedV1ResponseTypeError, CapabilityRequestPagedV1ResponseTypeNone:
		return true
	}
	return false
}

type EmbedAccountCapabilityRequestNewParams struct {
	// Allows the account to accept payments from businesses.
	Businesses param.Field[EmbedAccountCapabilityRequestNewParamsBusinesses] `json:"businesses"`
	// The charges capability settings for the account.
	Charges param.Field[EmbedAccountCapabilityRequestNewParamsCharges] `json:"charges"`
	// Allows the account to accept payments from individuals.
	Individuals param.Field[EmbedAccountCapabilityRequestNewParamsIndividuals] `json:"individuals"`
	// Allows the account to accept payments authorized via the internet or mobile
	// applications.
	Internet param.Field[EmbedAccountCapabilityRequestNewParamsInternet] `json:"internet"`
	// The payouts capability settings for the account.
	Payouts param.Field[EmbedAccountCapabilityRequestNewParamsPayouts] `json:"payouts"`
	// Allows the account to accept payments authorized by signed agreements or
	// contracts.
	SignedAgreement param.Field[EmbedAccountCapabilityRequestNewParamsSignedAgreement] `json:"signed_agreement"`
	CorrelationID   param.Field[string]                                                `header:"correlation-id"`
	RequestID       param.Field[string]                                                `header:"request-id"`
}

func (r EmbedAccountCapabilityRequestNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Allows the account to accept payments from businesses.
type EmbedAccountCapabilityRequestNewParamsBusinesses struct {
	Enable param.Field[bool] `json:"enable,required"`
}

func (r EmbedAccountCapabilityRequestNewParamsBusinesses) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The charges capability settings for the account.
type EmbedAccountCapabilityRequestNewParamsCharges struct {
	// The maximum dollar amount of charges in a calendar day.
	DailyAmount param.Field[float64] `json:"daily_amount,required"`
	// Determines whether `charges` are enabled for the account.
	Enable param.Field[bool] `json:"enable,required"`
	// The maximum amount of a single charge.
	MaxAmount param.Field[float64] `json:"max_amount,required"`
	// The maximum dollar amount of charges in a calendar month.
	MonthlyAmount param.Field[float64] `json:"monthly_amount,required"`
	// The maximum number of charges in a calendar month.
	MonthlyCount param.Field[int64] `json:"monthly_count,required"`
}

func (r EmbedAccountCapabilityRequestNewParamsCharges) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Allows the account to accept payments from individuals.
type EmbedAccountCapabilityRequestNewParamsIndividuals struct {
	Enable param.Field[bool] `json:"enable,required"`
}

func (r EmbedAccountCapabilityRequestNewParamsIndividuals) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Allows the account to accept payments authorized via the internet or mobile
// applications.
type EmbedAccountCapabilityRequestNewParamsInternet struct {
	Enable param.Field[bool] `json:"enable,required"`
}

func (r EmbedAccountCapabilityRequestNewParamsInternet) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The payouts capability settings for the account.
type EmbedAccountCapabilityRequestNewParamsPayouts struct {
	// The maximum dollar amount of payouts in a day.
	DailyAmount param.Field[float64] `json:"daily_amount,required"`
	// Determines whether `payouts` are enabled for the account.
	Enable param.Field[bool] `json:"enable,required"`
	// The maximum amount of a single payout.
	MaxAmount param.Field[float64] `json:"max_amount,required"`
	// The maximum dollar amount of payouts in a month.
	MonthlyAmount param.Field[float64] `json:"monthly_amount,required"`
	// The maximum number of payouts in a month.
	MonthlyCount param.Field[int64] `json:"monthly_count,required"`
}

func (r EmbedAccountCapabilityRequestNewParamsPayouts) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Allows the account to accept payments authorized by signed agreements or
// contracts.
type EmbedAccountCapabilityRequestNewParamsSignedAgreement struct {
	Enable param.Field[bool] `json:"enable,required"`
}

func (r EmbedAccountCapabilityRequestNewParamsSignedAgreement) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EmbedAccountCapabilityRequestListParams struct {
	// Filter capability requests by category.
	Category param.Field[EmbedAccountCapabilityRequestListParamsCategory] `query:"category"`
	// Results page number. Starts at page 1.
	PageNumber param.Field[int64] `query:"page_number"`
	// Page size.Max value: 1000
	PageSize param.Field[int64] `query:"page_size"`
	// Sort By.
	SortBy param.Field[string] `query:"sort_by"`
	// Sort Order.
	SortOrder param.Field[EmbedAccountCapabilityRequestListParamsSortOrder] `query:"sort_order"`
	// Filter capability requests by their current status.
	Status param.Field[EmbedAccountCapabilityRequestListParamsStatus] `query:"status"`
	// Filter capability requests by the specific type of capability.
	Type          param.Field[EmbedAccountCapabilityRequestListParamsType] `query:"type"`
	CorrelationID param.Field[string]                                      `header:"correlation-id"`
	RequestID     param.Field[string]                                      `header:"request-id"`
}

// URLQuery serializes [EmbedAccountCapabilityRequestListParams]'s query parameters
// as `url.Values`.
func (r EmbedAccountCapabilityRequestListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter capability requests by category.
type EmbedAccountCapabilityRequestListParamsCategory string

const (
	EmbedAccountCapabilityRequestListParamsCategoryPaymentType  EmbedAccountCapabilityRequestListParamsCategory = "payment_type"
	EmbedAccountCapabilityRequestListParamsCategoryCustomerType EmbedAccountCapabilityRequestListParamsCategory = "customer_type"
	EmbedAccountCapabilityRequestListParamsCategoryConsentType  EmbedAccountCapabilityRequestListParamsCategory = "consent_type"
)

func (r EmbedAccountCapabilityRequestListParamsCategory) IsKnown() bool {
	switch r {
	case EmbedAccountCapabilityRequestListParamsCategoryPaymentType, EmbedAccountCapabilityRequestListParamsCategoryCustomerType, EmbedAccountCapabilityRequestListParamsCategoryConsentType:
		return true
	}
	return false
}

// Sort Order.
type EmbedAccountCapabilityRequestListParamsSortOrder string

const (
	EmbedAccountCapabilityRequestListParamsSortOrderAsc  EmbedAccountCapabilityRequestListParamsSortOrder = "asc"
	EmbedAccountCapabilityRequestListParamsSortOrderDesc EmbedAccountCapabilityRequestListParamsSortOrder = "desc"
)

func (r EmbedAccountCapabilityRequestListParamsSortOrder) IsKnown() bool {
	switch r {
	case EmbedAccountCapabilityRequestListParamsSortOrderAsc, EmbedAccountCapabilityRequestListParamsSortOrderDesc:
		return true
	}
	return false
}

// Filter capability requests by their current status.
type EmbedAccountCapabilityRequestListParamsStatus string

const (
	EmbedAccountCapabilityRequestListParamsStatusActive   EmbedAccountCapabilityRequestListParamsStatus = "active"
	EmbedAccountCapabilityRequestListParamsStatusInactive EmbedAccountCapabilityRequestListParamsStatus = "inactive"
	EmbedAccountCapabilityRequestListParamsStatusInReview EmbedAccountCapabilityRequestListParamsStatus = "in_review"
	EmbedAccountCapabilityRequestListParamsStatusRejected EmbedAccountCapabilityRequestListParamsStatus = "rejected"
)

func (r EmbedAccountCapabilityRequestListParamsStatus) IsKnown() bool {
	switch r {
	case EmbedAccountCapabilityRequestListParamsStatusActive, EmbedAccountCapabilityRequestListParamsStatusInactive, EmbedAccountCapabilityRequestListParamsStatusInReview, EmbedAccountCapabilityRequestListParamsStatusRejected:
		return true
	}
	return false
}

// Filter capability requests by the specific type of capability.
type EmbedAccountCapabilityRequestListParamsType string

const (
	EmbedAccountCapabilityRequestListParamsTypeCharges         EmbedAccountCapabilityRequestListParamsType = "charges"
	EmbedAccountCapabilityRequestListParamsTypePayouts         EmbedAccountCapabilityRequestListParamsType = "payouts"
	EmbedAccountCapabilityRequestListParamsTypeIndividuals     EmbedAccountCapabilityRequestListParamsType = "individuals"
	EmbedAccountCapabilityRequestListParamsTypeBusinesses      EmbedAccountCapabilityRequestListParamsType = "businesses"
	EmbedAccountCapabilityRequestListParamsTypeSignedAgreement EmbedAccountCapabilityRequestListParamsType = "signed_agreement"
	EmbedAccountCapabilityRequestListParamsTypeInternet        EmbedAccountCapabilityRequestListParamsType = "internet"
)

func (r EmbedAccountCapabilityRequestListParamsType) IsKnown() bool {
	switch r {
	case EmbedAccountCapabilityRequestListParamsTypeCharges, EmbedAccountCapabilityRequestListParamsTypePayouts, EmbedAccountCapabilityRequestListParamsTypeIndividuals, EmbedAccountCapabilityRequestListParamsTypeBusinesses, EmbedAccountCapabilityRequestListParamsTypeSignedAgreement, EmbedAccountCapabilityRequestListParamsTypeInternet:
		return true
	}
	return false
}
