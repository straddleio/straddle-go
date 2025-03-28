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
func (r *EmbedRepresentativeService) New(ctx context.Context, params EmbedRepresentativeNewParams, opts ...option.RequestOption) (res *Representative, err error) {
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
func (r *EmbedRepresentativeService) Update(ctx context.Context, representativeID string, params EmbedRepresentativeUpdateParams, opts ...option.RequestOption) (res *Representative, err error) {
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
func (r *EmbedRepresentativeService) List(ctx context.Context, params EmbedRepresentativeListParams, opts ...option.RequestOption) (res *pagination.PageNumberSchema[RepresentativePagedData], err error) {
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
func (r *EmbedRepresentativeService) ListAutoPaging(ctx context.Context, params EmbedRepresentativeListParams, opts ...option.RequestOption) *pagination.PageNumberSchemaAutoPager[RepresentativePagedData] {
	return pagination.NewPageNumberSchemaAutoPager(r.List(ctx, params, opts...))
}

// Retrieves the details of an existing representative. Supply the unique
// representative ID, and Straddle will return the corresponding representative
// information.
func (r *EmbedRepresentativeService) Get(ctx context.Context, representativeID string, query EmbedRepresentativeGetParams, opts ...option.RequestOption) (res *Representative, err error) {
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
func (r *EmbedRepresentativeService) Unmask(ctx context.Context, representativeID string, query EmbedRepresentativeUnmaskParams, opts ...option.RequestOption) (res *Representative, err error) {
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

type Representative struct {
	Data RepresentativeData `json:"data,required"`
	// Metadata about the API request, including an identifier and timestamp.
	Meta shared.ResponseMetadata `json:"meta,required"`
	// Indicates the structure of the returned content.
	//
	//   - "object" means the `data` field contains a single JSON object.
	//   - "array" means the `data` field contains an array of objects.
	//   - "error" means the `data` field contains an error object with details of the
	//     issue.
	//   - "none" means no data is returned.
	ResponseType RepresentativeResponseType `json:"response_type,required"`
	JSON         representativeJSON         `json:"-"`
}

// representativeJSON contains the JSON metadata for the struct [Representative]
type representativeJSON struct {
	Data         apijson.Field
	Meta         apijson.Field
	ResponseType apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *Representative) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r representativeJSON) RawJSON() string {
	return r.raw
}

type RepresentativeData struct {
	// Unique identifier for the representative.
	ID string `json:"id,required" format:"uuid"`
	// The unique identifier of the account this representative is associated with.
	AccountID string `json:"account_id,required" format:"uuid"`
	// Timestamp of when the representative was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The date of birth of the representative, in ISO 8601 format (YYYY-MM-DD).
	Dob time.Time `json:"dob,required" format:"date"`
	// The email address of the representative.
	Email string `json:"email,required" format:"email"`
	// The first name of the representative.
	FirstName string `json:"first_name,required"`
	// The last name of the representative.
	LastName string `json:"last_name,required"`
	// The mobile phone number of the representative.
	MobileNumber string                         `json:"mobile_number,required"`
	Relationship RepresentativeDataRelationship `json:"relationship,required"`
	// The last 4 digits of the representative's Social Security Number.
	SsnLast4 string `json:"ssn_last4,required"`
	// The current status of the representative.
	Status       RepresentativeDataStatus       `json:"status,required"`
	StatusDetail RepresentativeDataStatusDetail `json:"status_detail,required"`
	// Timestamp of the most recent update to the representative.
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// Unique identifier for the representative in your database, used for
	// cross-referencing between Straddle and your systems.
	ExternalID string `json:"external_id,nullable"`
	// The unique identifier of the user account associated with this representative,
	// if applicable.
	UserID string                 `json:"user_id,nullable" format:"uuid"`
	JSON   representativeDataJSON `json:"-"`
}

// representativeDataJSON contains the JSON metadata for the struct
// [RepresentativeData]
type representativeDataJSON struct {
	ID           apijson.Field
	AccountID    apijson.Field
	CreatedAt    apijson.Field
	Dob          apijson.Field
	Email        apijson.Field
	FirstName    apijson.Field
	LastName     apijson.Field
	MobileNumber apijson.Field
	Relationship apijson.Field
	SsnLast4     apijson.Field
	Status       apijson.Field
	StatusDetail apijson.Field
	UpdatedAt    apijson.Field
	ExternalID   apijson.Field
	UserID       apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RepresentativeData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r representativeDataJSON) RawJSON() string {
	return r.raw
}

type RepresentativeDataRelationship struct {
	// Whether the representative has significant responsibility to control, manage, or
	// direct the organization. One representative must be identified under the control
	// prong for each legal entity.
	Control bool `json:"control,required"`
	// Whether the representative owns any percentage of of the equity interests of the
	// legal entity.
	Owner bool `json:"owner,required"`
	// Whether the person is authorized as the primary representative of the account.
	// This is the person chosen by the business to provide information about
	// themselves, general information about the account, and who consented to the
	// services agreement.
	//
	// There can be only one primary representative for an account at a time.
	Primary bool `json:"primary,required"`
	// The percentage of ownership the representative has. Required if 'Owner' is true.
	PercentOwnership float64 `json:"percent_ownership,nullable"`
	// The job title of the representative.
	Title string                             `json:"title,nullable"`
	JSON  representativeDataRelationshipJSON `json:"-"`
}

// representativeDataRelationshipJSON contains the JSON metadata for the struct
// [RepresentativeDataRelationship]
type representativeDataRelationshipJSON struct {
	Control          apijson.Field
	Owner            apijson.Field
	Primary          apijson.Field
	PercentOwnership apijson.Field
	Title            apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RepresentativeDataRelationship) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r representativeDataRelationshipJSON) RawJSON() string {
	return r.raw
}

// The current status of the representative.
type RepresentativeDataStatus string

const (
	RepresentativeDataStatusCreated    RepresentativeDataStatus = "created"
	RepresentativeDataStatusOnboarding RepresentativeDataStatus = "onboarding"
	RepresentativeDataStatusActive     RepresentativeDataStatus = "active"
	RepresentativeDataStatusRejected   RepresentativeDataStatus = "rejected"
	RepresentativeDataStatusInactive   RepresentativeDataStatus = "inactive"
)

func (r RepresentativeDataStatus) IsKnown() bool {
	switch r {
	case RepresentativeDataStatusCreated, RepresentativeDataStatusOnboarding, RepresentativeDataStatusActive, RepresentativeDataStatusRejected, RepresentativeDataStatusInactive:
		return true
	}
	return false
}

type RepresentativeDataStatusDetail struct {
	// A machine-readable code for the specific status, useful for programmatic
	// handling.
	Code string `json:"code,required"`
	// A human-readable message describing the current status.
	Message string `json:"message,required"`
	// A machine-readable identifier for the specific status, useful for programmatic
	// handling.
	Reason RepresentativeDataStatusDetailReason `json:"reason,required"`
	// Identifies the origin of the status change (e.g., `watchtower`). This helps in
	// tracking the cause of status updates.
	Source RepresentativeDataStatusDetailSource `json:"source,required"`
	JSON   representativeDataStatusDetailJSON   `json:"-"`
}

// representativeDataStatusDetailJSON contains the JSON metadata for the struct
// [RepresentativeDataStatusDetail]
type representativeDataStatusDetailJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	Reason      apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RepresentativeDataStatusDetail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r representativeDataStatusDetailJSON) RawJSON() string {
	return r.raw
}

// A machine-readable identifier for the specific status, useful for programmatic
// handling.
type RepresentativeDataStatusDetailReason string

const (
	RepresentativeDataStatusDetailReasonUnverified         RepresentativeDataStatusDetailReason = "unverified"
	RepresentativeDataStatusDetailReasonInReview           RepresentativeDataStatusDetailReason = "in_review"
	RepresentativeDataStatusDetailReasonPending            RepresentativeDataStatusDetailReason = "pending"
	RepresentativeDataStatusDetailReasonStuck              RepresentativeDataStatusDetailReason = "stuck"
	RepresentativeDataStatusDetailReasonVerified           RepresentativeDataStatusDetailReason = "verified"
	RepresentativeDataStatusDetailReasonFailedVerification RepresentativeDataStatusDetailReason = "failed_verification"
	RepresentativeDataStatusDetailReasonDisabled           RepresentativeDataStatusDetailReason = "disabled"
	RepresentativeDataStatusDetailReasonNew                RepresentativeDataStatusDetailReason = "new"
)

func (r RepresentativeDataStatusDetailReason) IsKnown() bool {
	switch r {
	case RepresentativeDataStatusDetailReasonUnverified, RepresentativeDataStatusDetailReasonInReview, RepresentativeDataStatusDetailReasonPending, RepresentativeDataStatusDetailReasonStuck, RepresentativeDataStatusDetailReasonVerified, RepresentativeDataStatusDetailReasonFailedVerification, RepresentativeDataStatusDetailReasonDisabled, RepresentativeDataStatusDetailReasonNew:
		return true
	}
	return false
}

// Identifies the origin of the status change (e.g., `watchtower`). This helps in
// tracking the cause of status updates.
type RepresentativeDataStatusDetailSource string

const (
	RepresentativeDataStatusDetailSourceWatchtower RepresentativeDataStatusDetailSource = "watchtower"
)

func (r RepresentativeDataStatusDetailSource) IsKnown() bool {
	switch r {
	case RepresentativeDataStatusDetailSourceWatchtower:
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
type RepresentativeResponseType string

const (
	RepresentativeResponseTypeObject RepresentativeResponseType = "object"
	RepresentativeResponseTypeArray  RepresentativeResponseType = "array"
	RepresentativeResponseTypeError  RepresentativeResponseType = "error"
	RepresentativeResponseTypeNone   RepresentativeResponseType = "none"
)

func (r RepresentativeResponseType) IsKnown() bool {
	switch r {
	case RepresentativeResponseTypeObject, RepresentativeResponseTypeArray, RepresentativeResponseTypeError, RepresentativeResponseTypeNone:
		return true
	}
	return false
}

type RepresentativePaged struct {
	Data []RepresentativePagedData `json:"data,required"`
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

type RepresentativePagedData struct {
	// Unique identifier for the representative.
	ID string `json:"id,required" format:"uuid"`
	// The unique identifier of the account this representative is associated with.
	AccountID string `json:"account_id,required" format:"uuid"`
	// Timestamp of when the representative was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The date of birth of the representative, in ISO 8601 format (YYYY-MM-DD).
	Dob time.Time `json:"dob,required" format:"date"`
	// The email address of the representative.
	Email string `json:"email,required" format:"email"`
	// The first name of the representative.
	FirstName string `json:"first_name,required"`
	// The last name of the representative.
	LastName string `json:"last_name,required"`
	// The mobile phone number of the representative.
	MobileNumber string                              `json:"mobile_number,required"`
	Relationship RepresentativePagedDataRelationship `json:"relationship,required"`
	// The last 4 digits of the representative's Social Security Number.
	SsnLast4 string `json:"ssn_last4,required"`
	// The current status of the representative.
	Status       RepresentativePagedDataStatus       `json:"status,required"`
	StatusDetail RepresentativePagedDataStatusDetail `json:"status_detail,required"`
	// Timestamp of the most recent update to the representative.
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// Unique identifier for the representative in your database, used for
	// cross-referencing between Straddle and your systems.
	ExternalID string `json:"external_id,nullable"`
	// The unique identifier of the user account associated with this representative,
	// if applicable.
	UserID string                      `json:"user_id,nullable" format:"uuid"`
	JSON   representativePagedDataJSON `json:"-"`
}

// representativePagedDataJSON contains the JSON metadata for the struct
// [RepresentativePagedData]
type representativePagedDataJSON struct {
	ID           apijson.Field
	AccountID    apijson.Field
	CreatedAt    apijson.Field
	Dob          apijson.Field
	Email        apijson.Field
	FirstName    apijson.Field
	LastName     apijson.Field
	MobileNumber apijson.Field
	Relationship apijson.Field
	SsnLast4     apijson.Field
	Status       apijson.Field
	StatusDetail apijson.Field
	UpdatedAt    apijson.Field
	ExternalID   apijson.Field
	UserID       apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RepresentativePagedData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r representativePagedDataJSON) RawJSON() string {
	return r.raw
}

type RepresentativePagedDataRelationship struct {
	// Whether the representative has significant responsibility to control, manage, or
	// direct the organization. One representative must be identified under the control
	// prong for each legal entity.
	Control bool `json:"control,required"`
	// Whether the representative owns any percentage of of the equity interests of the
	// legal entity.
	Owner bool `json:"owner,required"`
	// Whether the person is authorized as the primary representative of the account.
	// This is the person chosen by the business to provide information about
	// themselves, general information about the account, and who consented to the
	// services agreement.
	//
	// There can be only one primary representative for an account at a time.
	Primary bool `json:"primary,required"`
	// The percentage of ownership the representative has. Required if 'Owner' is true.
	PercentOwnership float64 `json:"percent_ownership,nullable"`
	// The job title of the representative.
	Title string                                  `json:"title,nullable"`
	JSON  representativePagedDataRelationshipJSON `json:"-"`
}

// representativePagedDataRelationshipJSON contains the JSON metadata for the
// struct [RepresentativePagedDataRelationship]
type representativePagedDataRelationshipJSON struct {
	Control          apijson.Field
	Owner            apijson.Field
	Primary          apijson.Field
	PercentOwnership apijson.Field
	Title            apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RepresentativePagedDataRelationship) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r representativePagedDataRelationshipJSON) RawJSON() string {
	return r.raw
}

// The current status of the representative.
type RepresentativePagedDataStatus string

const (
	RepresentativePagedDataStatusCreated    RepresentativePagedDataStatus = "created"
	RepresentativePagedDataStatusOnboarding RepresentativePagedDataStatus = "onboarding"
	RepresentativePagedDataStatusActive     RepresentativePagedDataStatus = "active"
	RepresentativePagedDataStatusRejected   RepresentativePagedDataStatus = "rejected"
	RepresentativePagedDataStatusInactive   RepresentativePagedDataStatus = "inactive"
)

func (r RepresentativePagedDataStatus) IsKnown() bool {
	switch r {
	case RepresentativePagedDataStatusCreated, RepresentativePagedDataStatusOnboarding, RepresentativePagedDataStatusActive, RepresentativePagedDataStatusRejected, RepresentativePagedDataStatusInactive:
		return true
	}
	return false
}

type RepresentativePagedDataStatusDetail struct {
	// A machine-readable code for the specific status, useful for programmatic
	// handling.
	Code string `json:"code,required"`
	// A human-readable message describing the current status.
	Message string `json:"message,required"`
	// A machine-readable identifier for the specific status, useful for programmatic
	// handling.
	Reason RepresentativePagedDataStatusDetailReason `json:"reason,required"`
	// Identifies the origin of the status change (e.g., `watchtower`). This helps in
	// tracking the cause of status updates.
	Source RepresentativePagedDataStatusDetailSource `json:"source,required"`
	JSON   representativePagedDataStatusDetailJSON   `json:"-"`
}

// representativePagedDataStatusDetailJSON contains the JSON metadata for the
// struct [RepresentativePagedDataStatusDetail]
type representativePagedDataStatusDetailJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	Reason      apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RepresentativePagedDataStatusDetail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r representativePagedDataStatusDetailJSON) RawJSON() string {
	return r.raw
}

// A machine-readable identifier for the specific status, useful for programmatic
// handling.
type RepresentativePagedDataStatusDetailReason string

const (
	RepresentativePagedDataStatusDetailReasonUnverified         RepresentativePagedDataStatusDetailReason = "unverified"
	RepresentativePagedDataStatusDetailReasonInReview           RepresentativePagedDataStatusDetailReason = "in_review"
	RepresentativePagedDataStatusDetailReasonPending            RepresentativePagedDataStatusDetailReason = "pending"
	RepresentativePagedDataStatusDetailReasonStuck              RepresentativePagedDataStatusDetailReason = "stuck"
	RepresentativePagedDataStatusDetailReasonVerified           RepresentativePagedDataStatusDetailReason = "verified"
	RepresentativePagedDataStatusDetailReasonFailedVerification RepresentativePagedDataStatusDetailReason = "failed_verification"
	RepresentativePagedDataStatusDetailReasonDisabled           RepresentativePagedDataStatusDetailReason = "disabled"
	RepresentativePagedDataStatusDetailReasonNew                RepresentativePagedDataStatusDetailReason = "new"
)

func (r RepresentativePagedDataStatusDetailReason) IsKnown() bool {
	switch r {
	case RepresentativePagedDataStatusDetailReasonUnverified, RepresentativePagedDataStatusDetailReasonInReview, RepresentativePagedDataStatusDetailReasonPending, RepresentativePagedDataStatusDetailReasonStuck, RepresentativePagedDataStatusDetailReasonVerified, RepresentativePagedDataStatusDetailReasonFailedVerification, RepresentativePagedDataStatusDetailReasonDisabled, RepresentativePagedDataStatusDetailReasonNew:
		return true
	}
	return false
}

// Identifies the origin of the status change (e.g., `watchtower`). This helps in
// tracking the cause of status updates.
type RepresentativePagedDataStatusDetailSource string

const (
	RepresentativePagedDataStatusDetailSourceWatchtower RepresentativePagedDataStatusDetailSource = "watchtower"
)

func (r RepresentativePagedDataStatusDetailSource) IsKnown() bool {
	switch r {
	case RepresentativePagedDataStatusDetailSourceWatchtower:
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
	MobileNumber param.Field[string]                                   `json:"mobile_number,required"`
	Relationship param.Field[EmbedRepresentativeNewParamsRelationship] `json:"relationship,required"`
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

type EmbedRepresentativeNewParamsRelationship struct {
	// Whether the representative has significant responsibility to control, manage, or
	// direct the organization. One representative must be identified under the control
	// prong for each legal entity.
	Control param.Field[bool] `json:"control,required"`
	// Whether the representative owns any percentage of of the equity interests of the
	// legal entity.
	Owner param.Field[bool] `json:"owner,required"`
	// Whether the person is authorized as the primary representative of the account.
	// This is the person chosen by the business to provide information about
	// themselves, general information about the account, and who consented to the
	// services agreement.
	//
	// There can be only one primary representative for an account at a time.
	Primary param.Field[bool] `json:"primary,required"`
	// The percentage of ownership the representative has. Required if 'Owner' is true.
	PercentOwnership param.Field[float64] `json:"percent_ownership"`
	// The job title of the representative.
	Title param.Field[string] `json:"title"`
}

func (r EmbedRepresentativeNewParamsRelationship) MarshalJSON() (data []byte, err error) {
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
	MobileNumber param.Field[string]                                      `json:"mobile_number,required"`
	Relationship param.Field[EmbedRepresentativeUpdateParamsRelationship] `json:"relationship,required"`
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

type EmbedRepresentativeUpdateParamsRelationship struct {
	// Whether the representative has significant responsibility to control, manage, or
	// direct the organization. One representative must be identified under the control
	// prong for each legal entity.
	Control param.Field[bool] `json:"control,required"`
	// Whether the representative owns any percentage of of the equity interests of the
	// legal entity.
	Owner param.Field[bool] `json:"owner,required"`
	// Whether the person is authorized as the primary representative of the account.
	// This is the person chosen by the business to provide information about
	// themselves, general information about the account, and who consented to the
	// services agreement.
	//
	// There can be only one primary representative for an account at a time.
	Primary param.Field[bool] `json:"primary,required"`
	// The percentage of ownership the representative has. Required if 'Owner' is true.
	PercentOwnership param.Field[float64] `json:"percent_ownership"`
	// The job title of the representative.
	Title param.Field[string] `json:"title"`
}

func (r EmbedRepresentativeUpdateParamsRelationship) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EmbedRepresentativeListParams struct {
	// The unique identifier of the account to list representatives for.
	AccountID      param.Field[string]                             `query:"account_id" format:"uuid"`
	Level          param.Field[EmbedRepresentativeListParamsLevel] `query:"level"`
	OrganizationID param.Field[string]                             `query:"organization_id" format:"uuid"`
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

type EmbedRepresentativeListParamsLevel string

const (
	EmbedRepresentativeListParamsLevelAccount  EmbedRepresentativeListParamsLevel = "account"
	EmbedRepresentativeListParamsLevelPlatform EmbedRepresentativeListParamsLevel = "platform"
)

func (r EmbedRepresentativeListParamsLevel) IsKnown() bool {
	switch r {
	case EmbedRepresentativeListParamsLevelAccount, EmbedRepresentativeListParamsLevelPlatform:
		return true
	}
	return false
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
