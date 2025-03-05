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

// EmbedLinkedBankAccountService contains methods and other services that help with
// interacting with the straddle API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEmbedLinkedBankAccountService] method instead.
type EmbedLinkedBankAccountService struct {
	Options []option.RequestOption
}

// NewEmbedLinkedBankAccountService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewEmbedLinkedBankAccountService(opts ...option.RequestOption) (r *EmbedLinkedBankAccountService) {
	r = &EmbedLinkedBankAccountService{}
	r.Options = opts
	return
}

// Creates a new linked bank account associated with a Straddle account. This
// endpoint allows you to associate external bank accounts with a Straddle account
// for various payment operations such as payment deposits, payout withdrawals, and
// more.
func (r *EmbedLinkedBankAccountService) New(ctx context.Context, params EmbedLinkedBankAccountNewParams, opts ...option.RequestOption) (res *LinkedBankAccountV1, err error) {
	if params.CorrelationID.Present {
		opts = append(opts, option.WithHeader("correlation-id", fmt.Sprintf("%s", params.CorrelationID)))
	}
	if params.RequestID.Present {
		opts = append(opts, option.WithHeader("request-id", fmt.Sprintf("%s", params.RequestID)))
	}
	opts = append(r.Options[:], opts...)
	path := "v1/linked_bank_accounts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Updates an existing linked bank account's information. This can be used to
// update account details during onboarding or to update metadata associated with
// the linked account. The linked bank account must be in 'created' or 'onboarding'
// status.
func (r *EmbedLinkedBankAccountService) Update(ctx context.Context, linkedBankAccountID string, params EmbedLinkedBankAccountUpdateParams, opts ...option.RequestOption) (res *LinkedBankAccountV1, err error) {
	if params.CorrelationID.Present {
		opts = append(opts, option.WithHeader("correlation-id", fmt.Sprintf("%s", params.CorrelationID)))
	}
	if params.RequestID.Present {
		opts = append(opts, option.WithHeader("request-id", fmt.Sprintf("%s", params.RequestID)))
	}
	opts = append(r.Options[:], opts...)
	if linkedBankAccountID == "" {
		err = errors.New("missing required linked_bank_account_id parameter")
		return
	}
	path := fmt.Sprintf("v1/linked_bank_accounts/%s", linkedBankAccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

// Returns a list of bank accounts associated with a specific Straddle account. The
// linked bank accounts are returned sorted by creation date, with the most
// recently created appearing first. This endpoint supports pagination to handle
// accounts with multiple linked bank accounts.
func (r *EmbedLinkedBankAccountService) List(ctx context.Context, params EmbedLinkedBankAccountListParams, opts ...option.RequestOption) (res *pagination.PageNumberSchema[LinkedBankAccountPagedV1Data], err error) {
	var raw *http.Response
	if params.CorrelationID.Present {
		opts = append(opts, option.WithHeader("correlation-id", fmt.Sprintf("%s", params.CorrelationID)))
	}
	if params.RequestID.Present {
		opts = append(opts, option.WithHeader("request-id", fmt.Sprintf("%s", params.RequestID)))
	}
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/linked_bank_accounts"
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

// Returns a list of bank accounts associated with a specific Straddle account. The
// linked bank accounts are returned sorted by creation date, with the most
// recently created appearing first. This endpoint supports pagination to handle
// accounts with multiple linked bank accounts.
func (r *EmbedLinkedBankAccountService) ListAutoPaging(ctx context.Context, params EmbedLinkedBankAccountListParams, opts ...option.RequestOption) *pagination.PageNumberSchemaAutoPager[LinkedBankAccountPagedV1Data] {
	return pagination.NewPageNumberSchemaAutoPager(r.List(ctx, params, opts...))
}

// Retrieves the details of a linked bank account that has previously been created.
// Supply the unique linked bank account `id`, and Straddle will return the
// corresponding information. The response includes masked account details for
// security purposes.
func (r *EmbedLinkedBankAccountService) Get(ctx context.Context, linkedBankAccountID string, query EmbedLinkedBankAccountGetParams, opts ...option.RequestOption) (res *LinkedBankAccountV1, err error) {
	if query.CorrelationID.Present {
		opts = append(opts, option.WithHeader("correlation-id", fmt.Sprintf("%s", query.CorrelationID)))
	}
	if query.RequestID.Present {
		opts = append(opts, option.WithHeader("request-id", fmt.Sprintf("%s", query.RequestID)))
	}
	opts = append(r.Options[:], opts...)
	if linkedBankAccountID == "" {
		err = errors.New("missing required linked_bank_account_id parameter")
		return
	}
	path := fmt.Sprintf("v1/linked_bank_accounts/%s", linkedBankAccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Retrieves the unmasked details of a linked bank account that has previously been
// created. Supply the unique linked bank account `id`, and Straddle will return
// the corresponding information, including sensitive details. This endpoint needs
// to be enabled by Straddle for your account and should only be used when
// absolutely necessary.
func (r *EmbedLinkedBankAccountService) Unmask(ctx context.Context, linkedBankAccountID string, query EmbedLinkedBankAccountUnmaskParams, opts ...option.RequestOption) (res *LinkedBankAccountUnmaskV1, err error) {
	if query.CorrelationID.Present {
		opts = append(opts, option.WithHeader("correlation-id", fmt.Sprintf("%s", query.CorrelationID)))
	}
	if query.RequestID.Present {
		opts = append(opts, option.WithHeader("request-id", fmt.Sprintf("%s", query.RequestID)))
	}
	opts = append(r.Options[:], opts...)
	if linkedBankAccountID == "" {
		err = errors.New("missing required linked_bank_account_id parameter")
		return
	}
	path := fmt.Sprintf("v1/linked_bank_accounts/%s/unmask", linkedBankAccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type LinkedBankAccountPagedV1 struct {
	Data []LinkedBankAccountPagedV1Data `json:"data,required"`
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
	ResponseType LinkedBankAccountPagedV1ResponseType `json:"response_type,required"`
	JSON         linkedBankAccountPagedV1JSON         `json:"-"`
}

// linkedBankAccountPagedV1JSON contains the JSON metadata for the struct
// [LinkedBankAccountPagedV1]
type linkedBankAccountPagedV1JSON struct {
	Data         apijson.Field
	Meta         apijson.Field
	ResponseType apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *LinkedBankAccountPagedV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r linkedBankAccountPagedV1JSON) RawJSON() string {
	return r.raw
}

type LinkedBankAccountPagedV1Data struct {
	// Unique identifier for the linked bank account.
	ID string `json:"id,required" format:"uuid"`
	// The unique identifier of the Straddle account related to this bank account.
	AccountID   string                                  `json:"account_id,required" format:"uuid"`
	BankAccount LinkedBankAccountPagedV1DataBankAccount `json:"bank_account,required"`
	// Timestamp of when the bank account object was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The current status of the linked bank account.
	Status       LinkedBankAccountPagedV1DataStatus       `json:"status,required"`
	StatusDetail LinkedBankAccountPagedV1DataStatusDetail `json:"status_detail,required"`
	// Timestamp of the most recent update to the linked bank account.
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// Up to 20 additional user-defined key-value pairs. Useful for storing additional
	// information about the linked bank account in a structured format.
	Metadata map[string]string                `json:"metadata,nullable"`
	JSON     linkedBankAccountPagedV1DataJSON `json:"-"`
}

// linkedBankAccountPagedV1DataJSON contains the JSON metadata for the struct
// [LinkedBankAccountPagedV1Data]
type linkedBankAccountPagedV1DataJSON struct {
	ID           apijson.Field
	AccountID    apijson.Field
	BankAccount  apijson.Field
	CreatedAt    apijson.Field
	Status       apijson.Field
	StatusDetail apijson.Field
	UpdatedAt    apijson.Field
	Metadata     apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *LinkedBankAccountPagedV1Data) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r linkedBankAccountPagedV1DataJSON) RawJSON() string {
	return r.raw
}

type LinkedBankAccountPagedV1DataBankAccount struct {
	AccountHolder   string                                      `json:"account_holder,required"`
	AccountMask     string                                      `json:"account_mask,required"`
	InstitutionName string                                      `json:"institution_name,required"`
	RoutingNumber   string                                      `json:"routing_number,required"`
	JSON            linkedBankAccountPagedV1DataBankAccountJSON `json:"-"`
}

// linkedBankAccountPagedV1DataBankAccountJSON contains the JSON metadata for the
// struct [LinkedBankAccountPagedV1DataBankAccount]
type linkedBankAccountPagedV1DataBankAccountJSON struct {
	AccountHolder   apijson.Field
	AccountMask     apijson.Field
	InstitutionName apijson.Field
	RoutingNumber   apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *LinkedBankAccountPagedV1DataBankAccount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r linkedBankAccountPagedV1DataBankAccountJSON) RawJSON() string {
	return r.raw
}

// The current status of the linked bank account.
type LinkedBankAccountPagedV1DataStatus string

const (
	LinkedBankAccountPagedV1DataStatusCreated    LinkedBankAccountPagedV1DataStatus = "created"
	LinkedBankAccountPagedV1DataStatusOnboarding LinkedBankAccountPagedV1DataStatus = "onboarding"
	LinkedBankAccountPagedV1DataStatusActive     LinkedBankAccountPagedV1DataStatus = "active"
	LinkedBankAccountPagedV1DataStatusRejected   LinkedBankAccountPagedV1DataStatus = "rejected"
	LinkedBankAccountPagedV1DataStatusInactive   LinkedBankAccountPagedV1DataStatus = "inactive"
)

func (r LinkedBankAccountPagedV1DataStatus) IsKnown() bool {
	switch r {
	case LinkedBankAccountPagedV1DataStatusCreated, LinkedBankAccountPagedV1DataStatusOnboarding, LinkedBankAccountPagedV1DataStatusActive, LinkedBankAccountPagedV1DataStatusRejected, LinkedBankAccountPagedV1DataStatusInactive:
		return true
	}
	return false
}

type LinkedBankAccountPagedV1DataStatusDetail struct {
	// A machine-readable code for the specific status, useful for programmatic
	// handling.
	Code string `json:"code,required"`
	// A human-readable message describing the current status.
	Message string `json:"message,required"`
	// A machine-readable identifier for the specific status, useful for programmatic
	// handling.
	Reason LinkedBankAccountPagedV1DataStatusDetailReason `json:"reason,required"`
	// Identifies the origin of the status change (e.g., `watchtower`). This helps in
	// tracking the cause of status updates.
	Source LinkedBankAccountPagedV1DataStatusDetailSource `json:"source,required"`
	JSON   linkedBankAccountPagedV1DataStatusDetailJSON   `json:"-"`
}

// linkedBankAccountPagedV1DataStatusDetailJSON contains the JSON metadata for the
// struct [LinkedBankAccountPagedV1DataStatusDetail]
type linkedBankAccountPagedV1DataStatusDetailJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	Reason      apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LinkedBankAccountPagedV1DataStatusDetail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r linkedBankAccountPagedV1DataStatusDetailJSON) RawJSON() string {
	return r.raw
}

// A machine-readable identifier for the specific status, useful for programmatic
// handling.
type LinkedBankAccountPagedV1DataStatusDetailReason string

const (
	LinkedBankAccountPagedV1DataStatusDetailReasonUnverified         LinkedBankAccountPagedV1DataStatusDetailReason = "unverified"
	LinkedBankAccountPagedV1DataStatusDetailReasonInReview           LinkedBankAccountPagedV1DataStatusDetailReason = "in_review"
	LinkedBankAccountPagedV1DataStatusDetailReasonPending            LinkedBankAccountPagedV1DataStatusDetailReason = "pending"
	LinkedBankAccountPagedV1DataStatusDetailReasonStuck              LinkedBankAccountPagedV1DataStatusDetailReason = "stuck"
	LinkedBankAccountPagedV1DataStatusDetailReasonVerified           LinkedBankAccountPagedV1DataStatusDetailReason = "verified"
	LinkedBankAccountPagedV1DataStatusDetailReasonFailedVerification LinkedBankAccountPagedV1DataStatusDetailReason = "failed_verification"
	LinkedBankAccountPagedV1DataStatusDetailReasonDisabled           LinkedBankAccountPagedV1DataStatusDetailReason = "disabled"
	LinkedBankAccountPagedV1DataStatusDetailReasonNew                LinkedBankAccountPagedV1DataStatusDetailReason = "new"
)

func (r LinkedBankAccountPagedV1DataStatusDetailReason) IsKnown() bool {
	switch r {
	case LinkedBankAccountPagedV1DataStatusDetailReasonUnverified, LinkedBankAccountPagedV1DataStatusDetailReasonInReview, LinkedBankAccountPagedV1DataStatusDetailReasonPending, LinkedBankAccountPagedV1DataStatusDetailReasonStuck, LinkedBankAccountPagedV1DataStatusDetailReasonVerified, LinkedBankAccountPagedV1DataStatusDetailReasonFailedVerification, LinkedBankAccountPagedV1DataStatusDetailReasonDisabled, LinkedBankAccountPagedV1DataStatusDetailReasonNew:
		return true
	}
	return false
}

// Identifies the origin of the status change (e.g., `watchtower`). This helps in
// tracking the cause of status updates.
type LinkedBankAccountPagedV1DataStatusDetailSource string

const (
	LinkedBankAccountPagedV1DataStatusDetailSourceWatchtower LinkedBankAccountPagedV1DataStatusDetailSource = "watchtower"
)

func (r LinkedBankAccountPagedV1DataStatusDetailSource) IsKnown() bool {
	switch r {
	case LinkedBankAccountPagedV1DataStatusDetailSourceWatchtower:
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
type LinkedBankAccountPagedV1ResponseType string

const (
	LinkedBankAccountPagedV1ResponseTypeObject LinkedBankAccountPagedV1ResponseType = "object"
	LinkedBankAccountPagedV1ResponseTypeArray  LinkedBankAccountPagedV1ResponseType = "array"
	LinkedBankAccountPagedV1ResponseTypeError  LinkedBankAccountPagedV1ResponseType = "error"
	LinkedBankAccountPagedV1ResponseTypeNone   LinkedBankAccountPagedV1ResponseType = "none"
)

func (r LinkedBankAccountPagedV1ResponseType) IsKnown() bool {
	switch r {
	case LinkedBankAccountPagedV1ResponseTypeObject, LinkedBankAccountPagedV1ResponseTypeArray, LinkedBankAccountPagedV1ResponseTypeError, LinkedBankAccountPagedV1ResponseTypeNone:
		return true
	}
	return false
}

type LinkedBankAccountUnmaskV1 struct {
	Data LinkedBankAccountUnmaskV1Data `json:"data,required"`
	// Metadata about the API request, including an identifier and timestamp.
	Meta shared.ResponseMetadata `json:"meta,required"`
	// Indicates the structure of the returned content.
	//
	//   - "object" means the `data` field contains a single JSON object.
	//   - "array" means the `data` field contains an array of objects.
	//   - "error" means the `data` field contains an error object with details of the
	//     issue.
	//   - "none" means no data is returned.
	ResponseType LinkedBankAccountUnmaskV1ResponseType `json:"response_type,required"`
	JSON         linkedBankAccountUnmaskV1JSON         `json:"-"`
}

// linkedBankAccountUnmaskV1JSON contains the JSON metadata for the struct
// [LinkedBankAccountUnmaskV1]
type linkedBankAccountUnmaskV1JSON struct {
	Data         apijson.Field
	Meta         apijson.Field
	ResponseType apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *LinkedBankAccountUnmaskV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r linkedBankAccountUnmaskV1JSON) RawJSON() string {
	return r.raw
}

type LinkedBankAccountUnmaskV1Data struct {
	// Unique identifier for the linked bank account.
	ID string `json:"id,required" format:"uuid"`
	// Unique identifier for the Straddle account related to this bank account.
	AccountID string `json:"account_id,required" format:"uuid"`
	// The bank account details associated with the linked bank account.
	BankAccount LinkedBankAccountUnmaskV1DataBankAccount `json:"bank_account,required"`
	// Timestamp of when the linked bank account was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The current status of the linked bank account.
	Status LinkedBankAccountUnmaskV1DataStatus `json:"status,required"`
	// Additional details about the current status of the linked bank account.
	StatusDetail LinkedBankAccountUnmaskV1DataStatusDetail `json:"status_detail,required"`
	// Timestamp of when the linked bank account was last updated.
	UpdatedAt time.Time                         `json:"updated_at,required" format:"date-time"`
	Metadata  map[string]string                 `json:"metadata,nullable"`
	JSON      linkedBankAccountUnmaskV1DataJSON `json:"-"`
}

// linkedBankAccountUnmaskV1DataJSON contains the JSON metadata for the struct
// [LinkedBankAccountUnmaskV1Data]
type linkedBankAccountUnmaskV1DataJSON struct {
	ID           apijson.Field
	AccountID    apijson.Field
	BankAccount  apijson.Field
	CreatedAt    apijson.Field
	Status       apijson.Field
	StatusDetail apijson.Field
	UpdatedAt    apijson.Field
	Metadata     apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *LinkedBankAccountUnmaskV1Data) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r linkedBankAccountUnmaskV1DataJSON) RawJSON() string {
	return r.raw
}

// The bank account details associated with the linked bank account.
type LinkedBankAccountUnmaskV1DataBankAccount struct {
	AccountHolder   string                                       `json:"account_holder,required"`
	AccountNumber   string                                       `json:"account_number,required"`
	InstitutionName string                                       `json:"institution_name,required"`
	RoutingNumber   string                                       `json:"routing_number,required"`
	JSON            linkedBankAccountUnmaskV1DataBankAccountJSON `json:"-"`
}

// linkedBankAccountUnmaskV1DataBankAccountJSON contains the JSON metadata for the
// struct [LinkedBankAccountUnmaskV1DataBankAccount]
type linkedBankAccountUnmaskV1DataBankAccountJSON struct {
	AccountHolder   apijson.Field
	AccountNumber   apijson.Field
	InstitutionName apijson.Field
	RoutingNumber   apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *LinkedBankAccountUnmaskV1DataBankAccount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r linkedBankAccountUnmaskV1DataBankAccountJSON) RawJSON() string {
	return r.raw
}

// The current status of the linked bank account.
type LinkedBankAccountUnmaskV1DataStatus string

const (
	LinkedBankAccountUnmaskV1DataStatusCreated    LinkedBankAccountUnmaskV1DataStatus = "created"
	LinkedBankAccountUnmaskV1DataStatusOnboarding LinkedBankAccountUnmaskV1DataStatus = "onboarding"
	LinkedBankAccountUnmaskV1DataStatusActive     LinkedBankAccountUnmaskV1DataStatus = "active"
	LinkedBankAccountUnmaskV1DataStatusRejected   LinkedBankAccountUnmaskV1DataStatus = "rejected"
	LinkedBankAccountUnmaskV1DataStatusInactive   LinkedBankAccountUnmaskV1DataStatus = "inactive"
)

func (r LinkedBankAccountUnmaskV1DataStatus) IsKnown() bool {
	switch r {
	case LinkedBankAccountUnmaskV1DataStatusCreated, LinkedBankAccountUnmaskV1DataStatusOnboarding, LinkedBankAccountUnmaskV1DataStatusActive, LinkedBankAccountUnmaskV1DataStatusRejected, LinkedBankAccountUnmaskV1DataStatusInactive:
		return true
	}
	return false
}

// Additional details about the current status of the linked bank account.
type LinkedBankAccountUnmaskV1DataStatusDetail struct {
	// A machine-readable code for the specific status, useful for programmatic
	// handling.
	Code string `json:"code,required"`
	// A human-readable message describing the current status.
	Message string `json:"message,required"`
	// A machine-readable identifier for the specific status, useful for programmatic
	// handling.
	Reason LinkedBankAccountUnmaskV1DataStatusDetailReason `json:"reason,required"`
	// Identifies the origin of the status change (e.g., `watchtower`). This helps in
	// tracking the cause of status updates.
	Source LinkedBankAccountUnmaskV1DataStatusDetailSource `json:"source,required"`
	JSON   linkedBankAccountUnmaskV1DataStatusDetailJSON   `json:"-"`
}

// linkedBankAccountUnmaskV1DataStatusDetailJSON contains the JSON metadata for the
// struct [LinkedBankAccountUnmaskV1DataStatusDetail]
type linkedBankAccountUnmaskV1DataStatusDetailJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	Reason      apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LinkedBankAccountUnmaskV1DataStatusDetail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r linkedBankAccountUnmaskV1DataStatusDetailJSON) RawJSON() string {
	return r.raw
}

// A machine-readable identifier for the specific status, useful for programmatic
// handling.
type LinkedBankAccountUnmaskV1DataStatusDetailReason string

const (
	LinkedBankAccountUnmaskV1DataStatusDetailReasonUnverified         LinkedBankAccountUnmaskV1DataStatusDetailReason = "unverified"
	LinkedBankAccountUnmaskV1DataStatusDetailReasonInReview           LinkedBankAccountUnmaskV1DataStatusDetailReason = "in_review"
	LinkedBankAccountUnmaskV1DataStatusDetailReasonPending            LinkedBankAccountUnmaskV1DataStatusDetailReason = "pending"
	LinkedBankAccountUnmaskV1DataStatusDetailReasonStuck              LinkedBankAccountUnmaskV1DataStatusDetailReason = "stuck"
	LinkedBankAccountUnmaskV1DataStatusDetailReasonVerified           LinkedBankAccountUnmaskV1DataStatusDetailReason = "verified"
	LinkedBankAccountUnmaskV1DataStatusDetailReasonFailedVerification LinkedBankAccountUnmaskV1DataStatusDetailReason = "failed_verification"
	LinkedBankAccountUnmaskV1DataStatusDetailReasonDisabled           LinkedBankAccountUnmaskV1DataStatusDetailReason = "disabled"
	LinkedBankAccountUnmaskV1DataStatusDetailReasonNew                LinkedBankAccountUnmaskV1DataStatusDetailReason = "new"
)

func (r LinkedBankAccountUnmaskV1DataStatusDetailReason) IsKnown() bool {
	switch r {
	case LinkedBankAccountUnmaskV1DataStatusDetailReasonUnverified, LinkedBankAccountUnmaskV1DataStatusDetailReasonInReview, LinkedBankAccountUnmaskV1DataStatusDetailReasonPending, LinkedBankAccountUnmaskV1DataStatusDetailReasonStuck, LinkedBankAccountUnmaskV1DataStatusDetailReasonVerified, LinkedBankAccountUnmaskV1DataStatusDetailReasonFailedVerification, LinkedBankAccountUnmaskV1DataStatusDetailReasonDisabled, LinkedBankAccountUnmaskV1DataStatusDetailReasonNew:
		return true
	}
	return false
}

// Identifies the origin of the status change (e.g., `watchtower`). This helps in
// tracking the cause of status updates.
type LinkedBankAccountUnmaskV1DataStatusDetailSource string

const (
	LinkedBankAccountUnmaskV1DataStatusDetailSourceWatchtower LinkedBankAccountUnmaskV1DataStatusDetailSource = "watchtower"
)

func (r LinkedBankAccountUnmaskV1DataStatusDetailSource) IsKnown() bool {
	switch r {
	case LinkedBankAccountUnmaskV1DataStatusDetailSourceWatchtower:
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
type LinkedBankAccountUnmaskV1ResponseType string

const (
	LinkedBankAccountUnmaskV1ResponseTypeObject LinkedBankAccountUnmaskV1ResponseType = "object"
	LinkedBankAccountUnmaskV1ResponseTypeArray  LinkedBankAccountUnmaskV1ResponseType = "array"
	LinkedBankAccountUnmaskV1ResponseTypeError  LinkedBankAccountUnmaskV1ResponseType = "error"
	LinkedBankAccountUnmaskV1ResponseTypeNone   LinkedBankAccountUnmaskV1ResponseType = "none"
)

func (r LinkedBankAccountUnmaskV1ResponseType) IsKnown() bool {
	switch r {
	case LinkedBankAccountUnmaskV1ResponseTypeObject, LinkedBankAccountUnmaskV1ResponseTypeArray, LinkedBankAccountUnmaskV1ResponseTypeError, LinkedBankAccountUnmaskV1ResponseTypeNone:
		return true
	}
	return false
}

type LinkedBankAccountV1 struct {
	Data LinkedBankAccountV1Data `json:"data,required"`
	// Metadata about the API request, including an identifier and timestamp.
	Meta shared.ResponseMetadata `json:"meta,required"`
	// Indicates the structure of the returned content.
	//
	//   - "object" means the `data` field contains a single JSON object.
	//   - "array" means the `data` field contains an array of objects.
	//   - "error" means the `data` field contains an error object with details of the
	//     issue.
	//   - "none" means no data is returned.
	ResponseType LinkedBankAccountV1ResponseType `json:"response_type,required"`
	JSON         linkedBankAccountV1JSON         `json:"-"`
}

// linkedBankAccountV1JSON contains the JSON metadata for the struct
// [LinkedBankAccountV1]
type linkedBankAccountV1JSON struct {
	Data         apijson.Field
	Meta         apijson.Field
	ResponseType apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *LinkedBankAccountV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r linkedBankAccountV1JSON) RawJSON() string {
	return r.raw
}

type LinkedBankAccountV1Data struct {
	// Unique identifier for the linked bank account.
	ID string `json:"id,required" format:"uuid"`
	// The unique identifier of the Straddle account related to this bank account.
	AccountID   string                             `json:"account_id,required" format:"uuid"`
	BankAccount LinkedBankAccountV1DataBankAccount `json:"bank_account,required"`
	// Timestamp of when the bank account object was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The current status of the linked bank account.
	Status       LinkedBankAccountV1DataStatus       `json:"status,required"`
	StatusDetail LinkedBankAccountV1DataStatusDetail `json:"status_detail,required"`
	// Timestamp of the most recent update to the linked bank account.
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// Up to 20 additional user-defined key-value pairs. Useful for storing additional
	// information about the linked bank account in a structured format.
	Metadata map[string]string           `json:"metadata,nullable"`
	JSON     linkedBankAccountV1DataJSON `json:"-"`
}

// linkedBankAccountV1DataJSON contains the JSON metadata for the struct
// [LinkedBankAccountV1Data]
type linkedBankAccountV1DataJSON struct {
	ID           apijson.Field
	AccountID    apijson.Field
	BankAccount  apijson.Field
	CreatedAt    apijson.Field
	Status       apijson.Field
	StatusDetail apijson.Field
	UpdatedAt    apijson.Field
	Metadata     apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *LinkedBankAccountV1Data) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r linkedBankAccountV1DataJSON) RawJSON() string {
	return r.raw
}

type LinkedBankAccountV1DataBankAccount struct {
	AccountHolder   string                                 `json:"account_holder,required"`
	AccountMask     string                                 `json:"account_mask,required"`
	InstitutionName string                                 `json:"institution_name,required"`
	RoutingNumber   string                                 `json:"routing_number,required"`
	JSON            linkedBankAccountV1DataBankAccountJSON `json:"-"`
}

// linkedBankAccountV1DataBankAccountJSON contains the JSON metadata for the struct
// [LinkedBankAccountV1DataBankAccount]
type linkedBankAccountV1DataBankAccountJSON struct {
	AccountHolder   apijson.Field
	AccountMask     apijson.Field
	InstitutionName apijson.Field
	RoutingNumber   apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *LinkedBankAccountV1DataBankAccount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r linkedBankAccountV1DataBankAccountJSON) RawJSON() string {
	return r.raw
}

// The current status of the linked bank account.
type LinkedBankAccountV1DataStatus string

const (
	LinkedBankAccountV1DataStatusCreated    LinkedBankAccountV1DataStatus = "created"
	LinkedBankAccountV1DataStatusOnboarding LinkedBankAccountV1DataStatus = "onboarding"
	LinkedBankAccountV1DataStatusActive     LinkedBankAccountV1DataStatus = "active"
	LinkedBankAccountV1DataStatusRejected   LinkedBankAccountV1DataStatus = "rejected"
	LinkedBankAccountV1DataStatusInactive   LinkedBankAccountV1DataStatus = "inactive"
)

func (r LinkedBankAccountV1DataStatus) IsKnown() bool {
	switch r {
	case LinkedBankAccountV1DataStatusCreated, LinkedBankAccountV1DataStatusOnboarding, LinkedBankAccountV1DataStatusActive, LinkedBankAccountV1DataStatusRejected, LinkedBankAccountV1DataStatusInactive:
		return true
	}
	return false
}

type LinkedBankAccountV1DataStatusDetail struct {
	// A machine-readable code for the specific status, useful for programmatic
	// handling.
	Code string `json:"code,required"`
	// A human-readable message describing the current status.
	Message string `json:"message,required"`
	// A machine-readable identifier for the specific status, useful for programmatic
	// handling.
	Reason LinkedBankAccountV1DataStatusDetailReason `json:"reason,required"`
	// Identifies the origin of the status change (e.g., `watchtower`). This helps in
	// tracking the cause of status updates.
	Source LinkedBankAccountV1DataStatusDetailSource `json:"source,required"`
	JSON   linkedBankAccountV1DataStatusDetailJSON   `json:"-"`
}

// linkedBankAccountV1DataStatusDetailJSON contains the JSON metadata for the
// struct [LinkedBankAccountV1DataStatusDetail]
type linkedBankAccountV1DataStatusDetailJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	Reason      apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LinkedBankAccountV1DataStatusDetail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r linkedBankAccountV1DataStatusDetailJSON) RawJSON() string {
	return r.raw
}

// A machine-readable identifier for the specific status, useful for programmatic
// handling.
type LinkedBankAccountV1DataStatusDetailReason string

const (
	LinkedBankAccountV1DataStatusDetailReasonUnverified         LinkedBankAccountV1DataStatusDetailReason = "unverified"
	LinkedBankAccountV1DataStatusDetailReasonInReview           LinkedBankAccountV1DataStatusDetailReason = "in_review"
	LinkedBankAccountV1DataStatusDetailReasonPending            LinkedBankAccountV1DataStatusDetailReason = "pending"
	LinkedBankAccountV1DataStatusDetailReasonStuck              LinkedBankAccountV1DataStatusDetailReason = "stuck"
	LinkedBankAccountV1DataStatusDetailReasonVerified           LinkedBankAccountV1DataStatusDetailReason = "verified"
	LinkedBankAccountV1DataStatusDetailReasonFailedVerification LinkedBankAccountV1DataStatusDetailReason = "failed_verification"
	LinkedBankAccountV1DataStatusDetailReasonDisabled           LinkedBankAccountV1DataStatusDetailReason = "disabled"
	LinkedBankAccountV1DataStatusDetailReasonNew                LinkedBankAccountV1DataStatusDetailReason = "new"
)

func (r LinkedBankAccountV1DataStatusDetailReason) IsKnown() bool {
	switch r {
	case LinkedBankAccountV1DataStatusDetailReasonUnverified, LinkedBankAccountV1DataStatusDetailReasonInReview, LinkedBankAccountV1DataStatusDetailReasonPending, LinkedBankAccountV1DataStatusDetailReasonStuck, LinkedBankAccountV1DataStatusDetailReasonVerified, LinkedBankAccountV1DataStatusDetailReasonFailedVerification, LinkedBankAccountV1DataStatusDetailReasonDisabled, LinkedBankAccountV1DataStatusDetailReasonNew:
		return true
	}
	return false
}

// Identifies the origin of the status change (e.g., `watchtower`). This helps in
// tracking the cause of status updates.
type LinkedBankAccountV1DataStatusDetailSource string

const (
	LinkedBankAccountV1DataStatusDetailSourceWatchtower LinkedBankAccountV1DataStatusDetailSource = "watchtower"
)

func (r LinkedBankAccountV1DataStatusDetailSource) IsKnown() bool {
	switch r {
	case LinkedBankAccountV1DataStatusDetailSourceWatchtower:
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
type LinkedBankAccountV1ResponseType string

const (
	LinkedBankAccountV1ResponseTypeObject LinkedBankAccountV1ResponseType = "object"
	LinkedBankAccountV1ResponseTypeArray  LinkedBankAccountV1ResponseType = "array"
	LinkedBankAccountV1ResponseTypeError  LinkedBankAccountV1ResponseType = "error"
	LinkedBankAccountV1ResponseTypeNone   LinkedBankAccountV1ResponseType = "none"
)

func (r LinkedBankAccountV1ResponseType) IsKnown() bool {
	switch r {
	case LinkedBankAccountV1ResponseTypeObject, LinkedBankAccountV1ResponseTypeArray, LinkedBankAccountV1ResponseTypeError, LinkedBankAccountV1ResponseTypeNone:
		return true
	}
	return false
}

type EmbedLinkedBankAccountNewParams struct {
	// The unique identifier of the Straddle account to associate this bank account
	// with.
	AccountID   param.Field[string]                                     `json:"account_id,required" format:"uuid"`
	BankAccount param.Field[EmbedLinkedBankAccountNewParamsBankAccount] `json:"bank_account,required"`
	// Up to 20 additional user-defined key-value pairs. Useful for storing additional
	// information about the linked bank account in a structured format.
	Metadata      param.Field[map[string]string] `json:"metadata"`
	CorrelationID param.Field[string]            `header:"correlation-id"`
	RequestID     param.Field[string]            `header:"request-id"`
}

func (r EmbedLinkedBankAccountNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EmbedLinkedBankAccountNewParamsBankAccount struct {
	// The name of the account holder as it appears on the bank account. Typically,
	// this is the legal name of the business associated with the account.
	AccountHolder param.Field[string] `json:"account_holder,required"`
	// The bank account number.
	AccountNumber param.Field[string] `json:"account_number,required"`
	// The routing number of the bank account.
	RoutingNumber param.Field[string] `json:"routing_number,required"`
}

func (r EmbedLinkedBankAccountNewParamsBankAccount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EmbedLinkedBankAccountUpdateParams struct {
	BankAccount param.Field[EmbedLinkedBankAccountUpdateParamsBankAccount] `json:"bank_account,required"`
	// Up to 20 additional user-defined key-value pairs. Useful for storing additional
	// information about the linked bank account in a structured format.
	Metadata      param.Field[map[string]string] `json:"metadata"`
	CorrelationID param.Field[string]            `header:"correlation-id"`
	RequestID     param.Field[string]            `header:"request-id"`
}

func (r EmbedLinkedBankAccountUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EmbedLinkedBankAccountUpdateParamsBankAccount struct {
	// The name of the account holder as it appears on the bank account. Typically,
	// this is the legal name of the business associated with the account.
	AccountHolder param.Field[string] `json:"account_holder,required"`
	// The bank account number.
	AccountNumber param.Field[string] `json:"account_number,required"`
	// The routing number of the bank account.
	RoutingNumber param.Field[string] `json:"routing_number,required"`
}

func (r EmbedLinkedBankAccountUpdateParamsBankAccount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EmbedLinkedBankAccountListParams struct {
	// The unique identifier of the related account.
	AccountID param.Field[string] `query:"account_id" format:"uuid"`
	// Results page number. Starts at page 1.
	PageNumber param.Field[int64] `query:"page_number"`
	// Page size. Max value: 1000
	PageSize param.Field[int64] `query:"page_size"`
	// Sort By.
	SortBy param.Field[string] `query:"sort_by"`
	// Sort Order.
	SortOrder     param.Field[EmbedLinkedBankAccountListParamsSortOrder] `query:"sort_order"`
	CorrelationID param.Field[string]                                    `header:"correlation-id"`
	RequestID     param.Field[string]                                    `header:"request-id"`
}

// URLQuery serializes [EmbedLinkedBankAccountListParams]'s query parameters as
// `url.Values`.
func (r EmbedLinkedBankAccountListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Sort Order.
type EmbedLinkedBankAccountListParamsSortOrder string

const (
	EmbedLinkedBankAccountListParamsSortOrderAsc  EmbedLinkedBankAccountListParamsSortOrder = "asc"
	EmbedLinkedBankAccountListParamsSortOrderDesc EmbedLinkedBankAccountListParamsSortOrder = "desc"
)

func (r EmbedLinkedBankAccountListParamsSortOrder) IsKnown() bool {
	switch r {
	case EmbedLinkedBankAccountListParamsSortOrderAsc, EmbedLinkedBankAccountListParamsSortOrderDesc:
		return true
	}
	return false
}

type EmbedLinkedBankAccountGetParams struct {
	CorrelationID param.Field[string] `header:"correlation-id"`
	RequestID     param.Field[string] `header:"request-id"`
}

type EmbedLinkedBankAccountUnmaskParams struct {
	CorrelationID param.Field[string] `header:"correlation-id"`
	RequestID     param.Field[string] `header:"request-id"`
}
