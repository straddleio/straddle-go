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

	"github.com/stainless-sdks/straddle-go/internal/apijson"
	"github.com/stainless-sdks/straddle-go/internal/apiquery"
	"github.com/stainless-sdks/straddle-go/internal/requestconfig"
	"github.com/stainless-sdks/straddle-go/option"
	"github.com/stainless-sdks/straddle-go/packages/pagination"
	"github.com/stainless-sdks/straddle-go/packages/param"
	"github.com/stainless-sdks/straddle-go/packages/respjson"
	"github.com/stainless-sdks/straddle-go/shared"
)

// Capabilities enable specific features and services for an Account. Use
// capability requests to unlock higher processing limits, new payment types, or
// additional platform features as your users' businesses grow. Track approval
// status and manage documentation requirements through a single interface.
//
// EmbedAccountCapabilityRequestService contains methods and other services that
// help with interacting with the straddle API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEmbedAccountCapabilityRequestService] method instead.
type EmbedAccountCapabilityRequestService struct {
	options []option.RequestOption
}

// NewEmbedAccountCapabilityRequestService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewEmbedAccountCapabilityRequestService(opts ...option.RequestOption) (r EmbedAccountCapabilityRequestService) {
	r = EmbedAccountCapabilityRequestService{}
	r.options = opts
	return
}

// Submits a request to enable a specific capability for an account. Use this
// endpoint to request additional features or services for an account.
func (r *EmbedAccountCapabilityRequestService) New(ctx context.Context, accountID string, params EmbedAccountCapabilityRequestNewParams, opts ...option.RequestOption) (res *CapabilityRequestPagedV1, err error) {
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
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/accounts/%s/capability_requests", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Retrieves a list of capability requests associated with an account. The requests
// are returned sorted by creation date, with the most recent requests appearing
// first. This endpoint supports advanced sorting and filtering options.
func (r *EmbedAccountCapabilityRequestService) List(ctx context.Context, accountID string, params EmbedAccountCapabilityRequestListParams, opts ...option.RequestOption) (res *pagination.PageNumberSchema[CapabilityRequestPagedV1Data], err error) {
	var raw *http.Response
	if !param.IsOmitted(params.CorrelationID) {
		opts = append(opts, option.WithHeader("correlation-id", fmt.Sprintf("%v", params.CorrelationID.Value)))
	}
	if !param.IsOmitted(params.RequestID) {
		opts = append(opts, option.WithHeader("request-id", fmt.Sprintf("%v", params.RequestID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
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
	Data []CapabilityRequestPagedV1Data `json:"data" api:"required"`
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
	ResponseType CapabilityRequestPagedV1ResponseType `json:"response_type" api:"required"`
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
func (r CapabilityRequestPagedV1) RawJSON() string { return r.JSON.raw }
func (r *CapabilityRequestPagedV1) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CapabilityRequestPagedV1Data struct {
	// Unique identifier for the capability request.
	ID string `json:"id" api:"required" format:"uuid"`
	// The unique identifier of the account associated with this capability request.
	AccountID string `json:"account_id" api:"required" format:"uuid"`
	// The category of the requested capability. Use `payment_type` for charges and
	// payouts, `customer_type` to define `individuals` or `businesses`, and
	// `consent_type` for `signed_agreement` or `internet` payment authorization.
	//
	// Any of "payment_type", "customer_type", "consent_type".
	Category string `json:"category" api:"required"`
	// Timestamp of when the capability request was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Whether this capability request is to enable or disable the capability.
	Enable bool `json:"enable" api:"required"`
	// The current status of the capability request.
	//
	// Any of "active", "inactive", "in_review", "rejected", "approved", "reviewing".
	Status string `json:"status" api:"required"`
	// The specific type of capability being requested within the category.
	//
	// Any of "charges", "payouts", "individuals", "businesses", "signed_agreement",
	// "internet".
	Type string `json:"type" api:"required"`
	// Timestamp of the most recent update to the capability request.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// Any specific settings or configurations related to the requested capability.
	Settings map[string]any `json:"settings" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		AccountID   respjson.Field
		Category    respjson.Field
		CreatedAt   respjson.Field
		Enable      respjson.Field
		Status      respjson.Field
		Type        respjson.Field
		UpdatedAt   respjson.Field
		Settings    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CapabilityRequestPagedV1Data) RawJSON() string { return r.JSON.raw }
func (r *CapabilityRequestPagedV1Data) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
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

type EmbedAccountCapabilityRequestNewParams struct {
	CorrelationID  param.Opt[string] `header:"correlation-id,omitzero" json:"-"`
	IdempotencyKey param.Opt[string] `header:"idempotency-key,omitzero" json:"-"`
	RequestID      param.Opt[string] `header:"request-id,omitzero" json:"-"`
	// Allows the account to accept payments from businesses.
	Businesses EmbedAccountCapabilityRequestNewParamsBusinesses `json:"businesses,omitzero"`
	// The charges capability settings for the account.
	Charges EmbedAccountCapabilityRequestNewParamsCharges `json:"charges,omitzero"`
	// Allows the account to accept payments from individuals.
	Individuals EmbedAccountCapabilityRequestNewParamsIndividuals `json:"individuals,omitzero"`
	// Allows the account to accept payments authorized via the internet or mobile
	// applications.
	Internet EmbedAccountCapabilityRequestNewParamsInternet `json:"internet,omitzero"`
	// The payouts capability settings for the account.
	Payouts EmbedAccountCapabilityRequestNewParamsPayouts `json:"payouts,omitzero"`
	// Allows the account to accept payments authorized by signed agreements or
	// contracts.
	SignedAgreement EmbedAccountCapabilityRequestNewParamsSignedAgreement `json:"signed_agreement,omitzero"`
	paramObj
}

func (r EmbedAccountCapabilityRequestNewParams) MarshalJSON() (data []byte, err error) {
	type shadow EmbedAccountCapabilityRequestNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EmbedAccountCapabilityRequestNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows the account to accept payments from businesses.
//
// The property Enable is required.
type EmbedAccountCapabilityRequestNewParamsBusinesses struct {
	Enable bool `json:"enable" api:"required"`
	paramObj
}

func (r EmbedAccountCapabilityRequestNewParamsBusinesses) MarshalJSON() (data []byte, err error) {
	type shadow EmbedAccountCapabilityRequestNewParamsBusinesses
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EmbedAccountCapabilityRequestNewParamsBusinesses) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The charges capability settings for the account.
//
// The properties DailyAmount, Enable, MaxAmount, MonthlyAmount, MonthlyCount are
// required.
type EmbedAccountCapabilityRequestNewParamsCharges struct {
	// The maximum dollar amount of charges in a calendar day.
	DailyAmount float64 `json:"daily_amount" api:"required"`
	// Determines whether `charges` are enabled for the account.
	Enable bool `json:"enable" api:"required"`
	// The maximum amount of a single charge.
	MaxAmount float64 `json:"max_amount" api:"required"`
	// The maximum dollar amount of charges in a calendar month.
	MonthlyAmount float64 `json:"monthly_amount" api:"required"`
	// The maximum number of charges in a calendar month.
	MonthlyCount int64 `json:"monthly_count" api:"required"`
	paramObj
}

func (r EmbedAccountCapabilityRequestNewParamsCharges) MarshalJSON() (data []byte, err error) {
	type shadow EmbedAccountCapabilityRequestNewParamsCharges
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EmbedAccountCapabilityRequestNewParamsCharges) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows the account to accept payments from individuals.
//
// The property Enable is required.
type EmbedAccountCapabilityRequestNewParamsIndividuals struct {
	Enable bool `json:"enable" api:"required"`
	paramObj
}

func (r EmbedAccountCapabilityRequestNewParamsIndividuals) MarshalJSON() (data []byte, err error) {
	type shadow EmbedAccountCapabilityRequestNewParamsIndividuals
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EmbedAccountCapabilityRequestNewParamsIndividuals) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows the account to accept payments authorized via the internet or mobile
// applications.
//
// The property Enable is required.
type EmbedAccountCapabilityRequestNewParamsInternet struct {
	Enable bool `json:"enable" api:"required"`
	paramObj
}

func (r EmbedAccountCapabilityRequestNewParamsInternet) MarshalJSON() (data []byte, err error) {
	type shadow EmbedAccountCapabilityRequestNewParamsInternet
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EmbedAccountCapabilityRequestNewParamsInternet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The payouts capability settings for the account.
//
// The properties DailyAmount, Enable, MaxAmount, MonthlyAmount, MonthlyCount are
// required.
type EmbedAccountCapabilityRequestNewParamsPayouts struct {
	// The maximum dollar amount of payouts in a day.
	DailyAmount float64 `json:"daily_amount" api:"required"`
	// Determines whether `payouts` are enabled for the account.
	Enable bool `json:"enable" api:"required"`
	// The maximum amount of a single payout.
	MaxAmount float64 `json:"max_amount" api:"required"`
	// The maximum dollar amount of payouts in a month.
	MonthlyAmount float64 `json:"monthly_amount" api:"required"`
	// The maximum number of payouts in a month.
	MonthlyCount int64 `json:"monthly_count" api:"required"`
	paramObj
}

func (r EmbedAccountCapabilityRequestNewParamsPayouts) MarshalJSON() (data []byte, err error) {
	type shadow EmbedAccountCapabilityRequestNewParamsPayouts
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EmbedAccountCapabilityRequestNewParamsPayouts) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows the account to accept payments authorized by signed agreements or
// contracts.
//
// The property Enable is required.
type EmbedAccountCapabilityRequestNewParamsSignedAgreement struct {
	Enable bool `json:"enable" api:"required"`
	paramObj
}

func (r EmbedAccountCapabilityRequestNewParamsSignedAgreement) MarshalJSON() (data []byte, err error) {
	type shadow EmbedAccountCapabilityRequestNewParamsSignedAgreement
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EmbedAccountCapabilityRequestNewParamsSignedAgreement) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbedAccountCapabilityRequestListParams struct {
	// Results page number. Starts at page 1.
	PageNumber param.Opt[int64] `query:"page_number,omitzero" json:"-"`
	// Page size.Max value: 1000
	PageSize param.Opt[int64] `query:"page_size,omitzero" json:"-"`
	// Sort By.
	SortBy        param.Opt[string] `query:"sort_by,omitzero" json:"-"`
	CorrelationID param.Opt[string] `header:"correlation-id,omitzero" json:"-"`
	RequestID     param.Opt[string] `header:"request-id,omitzero" json:"-"`
	// Filter capability requests by category.
	//
	// Any of "payment_type", "customer_type", "consent_type".
	Category EmbedAccountCapabilityRequestListParamsCategory `query:"category,omitzero" json:"-"`
	// Sort Order.
	//
	// Any of "asc", "desc".
	SortOrder EmbedAccountCapabilityRequestListParamsSortOrder `query:"sort_order,omitzero" json:"-"`
	// Filter capability requests by their current status.
	//
	// Any of "active", "inactive", "in_review", "rejected".
	Status EmbedAccountCapabilityRequestListParamsStatus `query:"status,omitzero" json:"-"`
	// Filter capability requests by the specific type of capability.
	//
	// Any of "charges", "payouts", "individuals", "businesses", "signed_agreement",
	// "internet".
	Type EmbedAccountCapabilityRequestListParamsType `query:"type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [EmbedAccountCapabilityRequestListParams]'s query parameters
// as `url.Values`.
func (r EmbedAccountCapabilityRequestListParams) URLQuery() (v url.Values, err error) {
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

// Sort Order.
type EmbedAccountCapabilityRequestListParamsSortOrder string

const (
	EmbedAccountCapabilityRequestListParamsSortOrderAsc  EmbedAccountCapabilityRequestListParamsSortOrder = "asc"
	EmbedAccountCapabilityRequestListParamsSortOrderDesc EmbedAccountCapabilityRequestListParamsSortOrder = "desc"
)

// Filter capability requests by their current status.
type EmbedAccountCapabilityRequestListParamsStatus string

const (
	EmbedAccountCapabilityRequestListParamsStatusActive   EmbedAccountCapabilityRequestListParamsStatus = "active"
	EmbedAccountCapabilityRequestListParamsStatusInactive EmbedAccountCapabilityRequestListParamsStatus = "inactive"
	EmbedAccountCapabilityRequestListParamsStatusInReview EmbedAccountCapabilityRequestListParamsStatus = "in_review"
	EmbedAccountCapabilityRequestListParamsStatusRejected EmbedAccountCapabilityRequestListParamsStatus = "rejected"
)

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
