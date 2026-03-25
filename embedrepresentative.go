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

// Representatives are individuals who have legal authority or significant
// responsibility within a business entity associated with a Straddle account. Each
// representative undergoes automated verification as part of KYC/KYB compliance.
// Use representatives to collect and verify beneficial owners, control persons,
// and authorized signers required for account onboarding. Representatives also
// determine who can legally operate the account and make important changes.
//
// EmbedRepresentativeService contains methods and other services that help with
// interacting with the straddle API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEmbedRepresentativeService] method instead.
type EmbedRepresentativeService struct {
	options []option.RequestOption
}

// NewEmbedRepresentativeService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewEmbedRepresentativeService(opts ...option.RequestOption) (r EmbedRepresentativeService) {
	r = EmbedRepresentativeService{}
	r.options = opts
	return
}

// Creates a new representative associated with an account. Representatives are
// individuals who have legal authority or significant responsibility within the
// business.
func (r *EmbedRepresentativeService) New(ctx context.Context, params EmbedRepresentativeNewParams, opts ...option.RequestOption) (res *Representative, err error) {
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
	path := "v1/representatives"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Updates an existing representative's information. This can be used to update
// personal details, contact information, or the relationship to the account or
// organization.
func (r *EmbedRepresentativeService) Update(ctx context.Context, representativeID string, params EmbedRepresentativeUpdateParams, opts ...option.RequestOption) (res *Representative, err error) {
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
	if representativeID == "" {
		err = errors.New("missing required representative_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/representatives/%s", representativeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return res, err
}

// Returns a list of representatives associated with a specific account or
// organization. The representatives are returned sorted by creation date, with the
// most recently created representatives appearing first. This endpoint supports
// advanced sorting and filtering options.
func (r *EmbedRepresentativeService) List(ctx context.Context, params EmbedRepresentativeListParams, opts ...option.RequestOption) (res *pagination.PageNumberSchema[RepresentativePagedData], err error) {
	var raw *http.Response
	if !param.IsOmitted(params.CorrelationID) {
		opts = append(opts, option.WithHeader("correlation-id", fmt.Sprintf("%v", params.CorrelationID.Value)))
	}
	if !param.IsOmitted(params.RequestID) {
		opts = append(opts, option.WithHeader("request-id", fmt.Sprintf("%v", params.RequestID.Value)))
	}
	opts = slices.Concat(r.options, opts)
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
func (r *EmbedRepresentativeService) ListAutoPaging(ctx context.Context, params EmbedRepresentativeListParams, opts ...option.RequestOption) *pagination.PageNumberSchemaAutoPager[RepresentativePagedData] {
	return pagination.NewPageNumberSchemaAutoPager(r.List(ctx, params, opts...))
}

// Retrieves the details of an existing representative. Supply the unique
// representative ID, and Straddle will return the corresponding representative
// information.
func (r *EmbedRepresentativeService) Get(ctx context.Context, representativeID string, query EmbedRepresentativeGetParams, opts ...option.RequestOption) (res *Representative, err error) {
	if !param.IsOmitted(query.CorrelationID) {
		opts = append(opts, option.WithHeader("correlation-id", fmt.Sprintf("%v", query.CorrelationID.Value)))
	}
	if !param.IsOmitted(query.RequestID) {
		opts = append(opts, option.WithHeader("request-id", fmt.Sprintf("%v", query.RequestID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	if representativeID == "" {
		err = errors.New("missing required representative_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/representatives/%s", representativeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Retrieves the unmasked details of a representative that has previously been
// created. Supply the unique representative ID, and Straddle will return the
// corresponding representative information, including sensitive details. This
// endpoint requires additional authentication and should be used with caution.
func (r *EmbedRepresentativeService) Unmask(ctx context.Context, representativeID string, query EmbedRepresentativeUnmaskParams, opts ...option.RequestOption) (res *Representative, err error) {
	if !param.IsOmitted(query.CorrelationID) {
		opts = append(opts, option.WithHeader("correlation-id", fmt.Sprintf("%v", query.CorrelationID.Value)))
	}
	if !param.IsOmitted(query.RequestID) {
		opts = append(opts, option.WithHeader("request-id", fmt.Sprintf("%v", query.RequestID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	if representativeID == "" {
		err = errors.New("missing required representative_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/representatives/%s/unmask", representativeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type Representative struct {
	Data RepresentativeData `json:"data" api:"required"`
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
	ResponseType RepresentativeResponseType `json:"response_type" api:"required"`
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
func (r Representative) RawJSON() string { return r.JSON.raw }
func (r *Representative) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RepresentativeData struct {
	// Unique identifier for the representative.
	ID string `json:"id" api:"required" format:"uuid"`
	// The unique identifier of the account this representative is associated with.
	AccountID string `json:"account_id" api:"required" format:"uuid"`
	// Timestamp of when the representative was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// The date of birth of the representative, in ISO 8601 format (YYYY-MM-DD).
	Dob time.Time `json:"dob" api:"required" format:"date"`
	// The email address of the representative.
	Email string `json:"email" api:"required" format:"email"`
	// The first name of the representative.
	FirstName string `json:"first_name" api:"required"`
	// The last name of the representative.
	LastName string `json:"last_name" api:"required"`
	// The mobile phone number of the representative.
	MobileNumber string                         `json:"mobile_number" api:"required"`
	Name         string                         `json:"name" api:"required"`
	Relationship RepresentativeDataRelationship `json:"relationship" api:"required"`
	// The last 4 digits of the representative's Social Security Number.
	SsnLast4 string `json:"ssn_last4" api:"required"`
	// The current status of the representative.
	//
	// Any of "created", "onboarding", "active", "rejected", "inactive".
	Status       string                         `json:"status" api:"required"`
	StatusDetail RepresentativeDataStatusDetail `json:"status_detail" api:"required"`
	// Timestamp of the most recent update to the representative.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// Unique identifier for the representative in your database, used for
	// cross-referencing between Straddle and your systems.
	ExternalID string `json:"external_id" api:"nullable"`
	// Up to 20 additional user-defined key-value pairs. Useful for storing additional
	// information about the represetative in a structured format.
	Metadata map[string]string `json:"metadata" api:"nullable"`
	Phone    string            `json:"phone" api:"nullable"`
	// The unique identifier of the user account associated with this representative,
	// if applicable.
	UserID string `json:"user_id" api:"nullable" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		AccountID    respjson.Field
		CreatedAt    respjson.Field
		Dob          respjson.Field
		Email        respjson.Field
		FirstName    respjson.Field
		LastName     respjson.Field
		MobileNumber respjson.Field
		Name         respjson.Field
		Relationship respjson.Field
		SsnLast4     respjson.Field
		Status       respjson.Field
		StatusDetail respjson.Field
		UpdatedAt    respjson.Field
		ExternalID   respjson.Field
		Metadata     respjson.Field
		Phone        respjson.Field
		UserID       respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RepresentativeData) RawJSON() string { return r.JSON.raw }
func (r *RepresentativeData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RepresentativeDataRelationship struct {
	// Whether the representative has significant responsibility to control, manage, or
	// direct the organization. One representative must be identified under the control
	// prong for each legal entity.
	Control bool `json:"control" api:"required"`
	// Whether the representative owns any percentage of of the equity interests of the
	// legal entity.
	Owner bool `json:"owner" api:"required"`
	// Whether the person is authorized as the primary representative of the account.
	// This is the person chosen by the business to provide information about
	// themselves, general information about the account, and who consented to the
	// services agreement.
	//
	// There can be only one primary representative for an account at a time.
	Primary bool `json:"primary" api:"required"`
	// The percentage of ownership the representative has. Required if 'Owner' is true.
	PercentOwnership float64 `json:"percent_ownership" api:"nullable"`
	// The job title of the representative.
	Title string `json:"title" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Control          respjson.Field
		Owner            respjson.Field
		Primary          respjson.Field
		PercentOwnership respjson.Field
		Title            respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RepresentativeDataRelationship) RawJSON() string { return r.JSON.raw }
func (r *RepresentativeDataRelationship) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RepresentativeDataStatusDetail struct {
	// A machine-readable code for the specific status, useful for programmatic
	// handling.
	Code string `json:"code" api:"required"`
	// A human-readable message describing the current status.
	Message string `json:"message" api:"required"`
	// A machine-readable identifier for the specific status, useful for programmatic
	// handling.
	//
	// Any of "unverified", "in_review", "pending", "stuck", "verified",
	// "failed_verification", "disabled", "new".
	Reason string `json:"reason" api:"required"`
	// Identifies the origin of the status change (e.g., `watchtower`). This helps in
	// tracking the cause of status updates.
	//
	// Any of "watchtower".
	Source string `json:"source" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Message     respjson.Field
		Reason      respjson.Field
		Source      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RepresentativeDataStatusDetail) RawJSON() string { return r.JSON.raw }
func (r *RepresentativeDataStatusDetail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indicates the structure of the returned content.
//
//   - "object" means the `data` field contains a single JSON object.
//   - "array" means the `data` field contains an array of objects.
//   - "error" means the `data` field contains an error object with details of the
//     issue.
//   - "none" means no data is returned.
type RepresentativeResponseType string

const (
	RepresentativeResponseTypeObject RepresentativeResponseType = "object"
	RepresentativeResponseTypeArray  RepresentativeResponseType = "array"
	RepresentativeResponseTypeError  RepresentativeResponseType = "error"
	RepresentativeResponseTypeNone   RepresentativeResponseType = "none"
)

type RepresentativePaged struct {
	Data []RepresentativePagedData `json:"data" api:"required"`
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
	ResponseType RepresentativePagedResponseType `json:"response_type" api:"required"`
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
func (r RepresentativePaged) RawJSON() string { return r.JSON.raw }
func (r *RepresentativePaged) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RepresentativePagedData struct {
	// Unique identifier for the representative.
	ID string `json:"id" api:"required" format:"uuid"`
	// The unique identifier of the account this representative is associated with.
	AccountID string `json:"account_id" api:"required" format:"uuid"`
	// Timestamp of when the representative was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// The date of birth of the representative, in ISO 8601 format (YYYY-MM-DD).
	Dob time.Time `json:"dob" api:"required" format:"date"`
	// The email address of the representative.
	Email string `json:"email" api:"required" format:"email"`
	// The first name of the representative.
	FirstName string `json:"first_name" api:"required"`
	// The last name of the representative.
	LastName string `json:"last_name" api:"required"`
	// The mobile phone number of the representative.
	MobileNumber string                              `json:"mobile_number" api:"required"`
	Name         string                              `json:"name" api:"required"`
	Relationship RepresentativePagedDataRelationship `json:"relationship" api:"required"`
	// The last 4 digits of the representative's Social Security Number.
	SsnLast4 string `json:"ssn_last4" api:"required"`
	// The current status of the representative.
	//
	// Any of "created", "onboarding", "active", "rejected", "inactive".
	Status       string                              `json:"status" api:"required"`
	StatusDetail RepresentativePagedDataStatusDetail `json:"status_detail" api:"required"`
	// Timestamp of the most recent update to the representative.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// Unique identifier for the representative in your database, used for
	// cross-referencing between Straddle and your systems.
	ExternalID string `json:"external_id" api:"nullable"`
	// Up to 20 additional user-defined key-value pairs. Useful for storing additional
	// information about the represetative in a structured format.
	Metadata map[string]string `json:"metadata" api:"nullable"`
	Phone    string            `json:"phone" api:"nullable"`
	// The unique identifier of the user account associated with this representative,
	// if applicable.
	UserID string `json:"user_id" api:"nullable" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		AccountID    respjson.Field
		CreatedAt    respjson.Field
		Dob          respjson.Field
		Email        respjson.Field
		FirstName    respjson.Field
		LastName     respjson.Field
		MobileNumber respjson.Field
		Name         respjson.Field
		Relationship respjson.Field
		SsnLast4     respjson.Field
		Status       respjson.Field
		StatusDetail respjson.Field
		UpdatedAt    respjson.Field
		ExternalID   respjson.Field
		Metadata     respjson.Field
		Phone        respjson.Field
		UserID       respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RepresentativePagedData) RawJSON() string { return r.JSON.raw }
func (r *RepresentativePagedData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RepresentativePagedDataRelationship struct {
	// Whether the representative has significant responsibility to control, manage, or
	// direct the organization. One representative must be identified under the control
	// prong for each legal entity.
	Control bool `json:"control" api:"required"`
	// Whether the representative owns any percentage of of the equity interests of the
	// legal entity.
	Owner bool `json:"owner" api:"required"`
	// Whether the person is authorized as the primary representative of the account.
	// This is the person chosen by the business to provide information about
	// themselves, general information about the account, and who consented to the
	// services agreement.
	//
	// There can be only one primary representative for an account at a time.
	Primary bool `json:"primary" api:"required"`
	// The percentage of ownership the representative has. Required if 'Owner' is true.
	PercentOwnership float64 `json:"percent_ownership" api:"nullable"`
	// The job title of the representative.
	Title string `json:"title" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Control          respjson.Field
		Owner            respjson.Field
		Primary          respjson.Field
		PercentOwnership respjson.Field
		Title            respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RepresentativePagedDataRelationship) RawJSON() string { return r.JSON.raw }
func (r *RepresentativePagedDataRelationship) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RepresentativePagedDataStatusDetail struct {
	// A machine-readable code for the specific status, useful for programmatic
	// handling.
	Code string `json:"code" api:"required"`
	// A human-readable message describing the current status.
	Message string `json:"message" api:"required"`
	// A machine-readable identifier for the specific status, useful for programmatic
	// handling.
	//
	// Any of "unverified", "in_review", "pending", "stuck", "verified",
	// "failed_verification", "disabled", "new".
	Reason string `json:"reason" api:"required"`
	// Identifies the origin of the status change (e.g., `watchtower`). This helps in
	// tracking the cause of status updates.
	//
	// Any of "watchtower".
	Source string `json:"source" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Message     respjson.Field
		Reason      respjson.Field
		Source      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RepresentativePagedDataStatusDetail) RawJSON() string { return r.JSON.raw }
func (r *RepresentativePagedDataStatusDetail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
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

type EmbedRepresentativeNewParams struct {
	// The unique identifier of the account this representative is associated with.
	AccountID string `json:"account_id" api:"required" format:"uuid"`
	// Date of birth for the representative in ISO 8601 format (YYYY-MM-DD).
	Dob time.Time `json:"dob" api:"required" format:"date"`
	// The company email address of the representative.
	Email string `json:"email" api:"required" format:"email"`
	// The first name of the representative.
	FirstName string `json:"first_name" api:"required"`
	// The last name of the representative.
	LastName string `json:"last_name" api:"required"`
	// The mobile phone number of the representative.
	MobileNumber string                                   `json:"mobile_number" api:"required"`
	Relationship EmbedRepresentativeNewParamsRelationship `json:"relationship,omitzero" api:"required"`
	// The last 4 digits of the representative's Social Security Number.
	SsnLast4 string `json:"ssn_last4" api:"required"`
	// Unique identifier for the representative in your database, used for
	// cross-referencing between Straddle and your systems.
	ExternalID     param.Opt[string] `json:"external_id,omitzero"`
	CorrelationID  param.Opt[string] `header:"correlation-id,omitzero" json:"-"`
	IdempotencyKey param.Opt[string] `header:"idempotency-key,omitzero" json:"-"`
	RequestID      param.Opt[string] `header:"request-id,omitzero" json:"-"`
	// Up to 20 additional user-defined key-value pairs. Useful for storing additional
	// information about the represetative in a structured format.
	Metadata map[string]string `json:"metadata,omitzero"`
	paramObj
}

func (r EmbedRepresentativeNewParams) MarshalJSON() (data []byte, err error) {
	type shadow EmbedRepresentativeNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EmbedRepresentativeNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Control, Owner, Primary are required.
type EmbedRepresentativeNewParamsRelationship struct {
	// Whether the representative has significant responsibility to control, manage, or
	// direct the organization. One representative must be identified under the control
	// prong for each legal entity.
	Control bool `json:"control" api:"required"`
	// Whether the representative owns any percentage of of the equity interests of the
	// legal entity.
	Owner bool `json:"owner" api:"required"`
	// Whether the person is authorized as the primary representative of the account.
	// This is the person chosen by the business to provide information about
	// themselves, general information about the account, and who consented to the
	// services agreement.
	//
	// There can be only one primary representative for an account at a time.
	Primary bool `json:"primary" api:"required"`
	// The percentage of ownership the representative has. Required if 'Owner' is true.
	PercentOwnership param.Opt[float64] `json:"percent_ownership,omitzero"`
	// The job title of the representative.
	Title param.Opt[string] `json:"title,omitzero"`
	paramObj
}

func (r EmbedRepresentativeNewParamsRelationship) MarshalJSON() (data []byte, err error) {
	type shadow EmbedRepresentativeNewParamsRelationship
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EmbedRepresentativeNewParamsRelationship) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbedRepresentativeUpdateParams struct {
	// The date of birth of the representative, in ISO 8601 format (YYYY-MM-DD).
	Dob time.Time `json:"dob" api:"required" format:"date"`
	// The email address of the representative.
	Email string `json:"email" api:"required" format:"email"`
	// The first name of the representative.
	FirstName string `json:"first_name" api:"required"`
	// The last name of the representative.
	LastName string `json:"last_name" api:"required"`
	// The mobile phone number of the representative.
	MobileNumber string                                      `json:"mobile_number" api:"required"`
	Relationship EmbedRepresentativeUpdateParamsRelationship `json:"relationship,omitzero" api:"required"`
	// The last 4 digits of the representative's Social Security Number.
	SsnLast4 string `json:"ssn_last4" api:"required"`
	// Unique identifier for the representative in your database, used for
	// cross-referencing between Straddle and your systems.
	ExternalID     param.Opt[string] `json:"external_id,omitzero"`
	CorrelationID  param.Opt[string] `header:"correlation-id,omitzero" json:"-"`
	IdempotencyKey param.Opt[string] `header:"idempotency-key,omitzero" json:"-"`
	RequestID      param.Opt[string] `header:"request-id,omitzero" json:"-"`
	// Up to 20 additional user-defined key-value pairs. Useful for storing additional
	// information about the represetative in a structured format.
	Metadata map[string]string `json:"metadata,omitzero"`
	paramObj
}

func (r EmbedRepresentativeUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow EmbedRepresentativeUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EmbedRepresentativeUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Control, Owner, Primary are required.
type EmbedRepresentativeUpdateParamsRelationship struct {
	// Whether the representative has significant responsibility to control, manage, or
	// direct the organization. One representative must be identified under the control
	// prong for each legal entity.
	Control bool `json:"control" api:"required"`
	// Whether the representative owns any percentage of of the equity interests of the
	// legal entity.
	Owner bool `json:"owner" api:"required"`
	// Whether the person is authorized as the primary representative of the account.
	// This is the person chosen by the business to provide information about
	// themselves, general information about the account, and who consented to the
	// services agreement.
	//
	// There can be only one primary representative for an account at a time.
	Primary bool `json:"primary" api:"required"`
	// The percentage of ownership the representative has. Required if 'Owner' is true.
	PercentOwnership param.Opt[float64] `json:"percent_ownership,omitzero"`
	// The job title of the representative.
	Title param.Opt[string] `json:"title,omitzero"`
	paramObj
}

func (r EmbedRepresentativeUpdateParamsRelationship) MarshalJSON() (data []byte, err error) {
	type shadow EmbedRepresentativeUpdateParamsRelationship
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EmbedRepresentativeUpdateParamsRelationship) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbedRepresentativeListParams struct {
	// The unique identifier of the account to list representatives for.
	AccountID      param.Opt[string] `query:"account_id,omitzero" format:"uuid" json:"-"`
	OrganizationID param.Opt[string] `query:"organization_id,omitzero" format:"uuid" json:"-"`
	// Results page number. Starts at page 1.
	PageNumber param.Opt[int64] `query:"page_number,omitzero" json:"-"`
	// Page size. Max value: 1000
	PageSize   param.Opt[int64]  `query:"page_size,omitzero" json:"-"`
	PlatformID param.Opt[string] `query:"platform_id,omitzero" format:"uuid" json:"-"`
	// Sort By.
	SortBy        param.Opt[string] `query:"sort_by,omitzero" json:"-"`
	CorrelationID param.Opt[string] `header:"correlation-id,omitzero" json:"-"`
	RequestID     param.Opt[string] `header:"request-id,omitzero" json:"-"`
	// Any of "account", "platform".
	Level EmbedRepresentativeListParamsLevel `query:"level,omitzero" json:"-"`
	// Sort Order.
	//
	// Any of "asc", "desc".
	SortOrder EmbedRepresentativeListParamsSortOrder `query:"sort_order,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [EmbedRepresentativeListParams]'s query parameters as
// `url.Values`.
func (r EmbedRepresentativeListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type EmbedRepresentativeListParamsLevel string

const (
	EmbedRepresentativeListParamsLevelAccount  EmbedRepresentativeListParamsLevel = "account"
	EmbedRepresentativeListParamsLevelPlatform EmbedRepresentativeListParamsLevel = "platform"
)

// Sort Order.
type EmbedRepresentativeListParamsSortOrder string

const (
	EmbedRepresentativeListParamsSortOrderAsc  EmbedRepresentativeListParamsSortOrder = "asc"
	EmbedRepresentativeListParamsSortOrderDesc EmbedRepresentativeListParamsSortOrder = "desc"
)

type EmbedRepresentativeGetParams struct {
	CorrelationID param.Opt[string] `header:"correlation-id,omitzero" json:"-"`
	RequestID     param.Opt[string] `header:"request-id,omitzero" json:"-"`
	paramObj
}

type EmbedRepresentativeUnmaskParams struct {
	CorrelationID param.Opt[string] `header:"correlation-id,omitzero" json:"-"`
	RequestID     param.Opt[string] `header:"request-id,omitzero" json:"-"`
	paramObj
}
