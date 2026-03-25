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

	"github.com/straddleio/straddle-go/internal/apijson"
	"github.com/straddleio/straddle-go/internal/apiquery"
	"github.com/straddleio/straddle-go/internal/requestconfig"
	"github.com/straddleio/straddle-go/option"
	"github.com/straddleio/straddle-go/packages/pagination"
	"github.com/straddleio/straddle-go/packages/param"
	"github.com/straddleio/straddle-go/packages/respjson"
	"github.com/straddleio/straddle-go/shared"
)

// Linked bank accounts connect your platform users' external bank accounts to
// Straddle for settlements and payment funding. Each linked account undergoes
// automated verification and continuous monitoring. Use linked accounts to manage
// where clients receive deposits, fund payouts, and track settlement preferences.
//
// EmbedLinkedBankAccountService contains methods and other services that help with
// interacting with the straddle API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEmbedLinkedBankAccountService] method instead.
type EmbedLinkedBankAccountService struct {
	options []option.RequestOption
}

// NewEmbedLinkedBankAccountService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewEmbedLinkedBankAccountService(opts ...option.RequestOption) (r EmbedLinkedBankAccountService) {
	r = EmbedLinkedBankAccountService{}
	r.options = opts
	return
}

// Creates a new linked bank account associated with a Straddle account. This
// endpoint allows you to associate external bank accounts with a Straddle account
// for various payment operations such as payment deposits, payout withdrawals, and
// more.
func (r *EmbedLinkedBankAccountService) New(ctx context.Context, params EmbedLinkedBankAccountNewParams, opts ...option.RequestOption) (res *LinkedBankAccountV1, err error) {
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
	path := "v1/linked_bank_accounts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Updates an existing linked bank account's information. This can be used to
// update account details during onboarding or to update metadata associated with
// the linked account. The linked bank account must be in 'created' or 'onboarding'
// status.
func (r *EmbedLinkedBankAccountService) Update(ctx context.Context, linkedBankAccountID string, params EmbedLinkedBankAccountUpdateParams, opts ...option.RequestOption) (res *LinkedBankAccountV1, err error) {
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
	if linkedBankAccountID == "" {
		err = errors.New("missing required linked_bank_account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/linked_bank_accounts/%s", linkedBankAccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return res, err
}

// Returns a list of bank accounts associated with a specific Straddle account. The
// linked bank accounts are returned sorted by creation date, with the most
// recently created appearing first. This endpoint supports pagination to handle
// accounts with multiple linked bank accounts.
func (r *EmbedLinkedBankAccountService) List(ctx context.Context, params EmbedLinkedBankAccountListParams, opts ...option.RequestOption) (res *pagination.PageNumberSchema[LinkedBankAccountPagedV1Data], err error) {
	var raw *http.Response
	if !param.IsOmitted(params.CorrelationID) {
		opts = append(opts, option.WithHeader("correlation-id", fmt.Sprintf("%v", params.CorrelationID.Value)))
	}
	if !param.IsOmitted(params.RequestID) {
		opts = append(opts, option.WithHeader("request-id", fmt.Sprintf("%v", params.RequestID.Value)))
	}
	opts = slices.Concat(r.options, opts)
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

// Cancels an existing linked bank account. This can be used to cancel a linked
// bank account before it has been reviewed. The linked bank account must be in
// 'created' status.
func (r *EmbedLinkedBankAccountService) Cancel(ctx context.Context, linkedBankAccountID string, body EmbedLinkedBankAccountCancelParams, opts ...option.RequestOption) (res *LinkedBankAccountV1, err error) {
	if !param.IsOmitted(body.CorrelationID) {
		opts = append(opts, option.WithHeader("correlation-id", fmt.Sprintf("%v", body.CorrelationID.Value)))
	}
	if !param.IsOmitted(body.IdempotencyKey) {
		opts = append(opts, option.WithHeader("idempotency-key", fmt.Sprintf("%v", body.IdempotencyKey.Value)))
	}
	if !param.IsOmitted(body.RequestID) {
		opts = append(opts, option.WithHeader("request-id", fmt.Sprintf("%v", body.RequestID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	if linkedBankAccountID == "" {
		err = errors.New("missing required linked_bank_account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/linked_bank_accounts/%s/cancel", linkedBankAccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, nil, &res, opts...)
	return res, err
}

// Retrieves the details of a linked bank account that has previously been created.
// Supply the unique linked bank account `id`, and Straddle will return the
// corresponding information. The response includes masked account details for
// security purposes.
func (r *EmbedLinkedBankAccountService) Get(ctx context.Context, linkedBankAccountID string, query EmbedLinkedBankAccountGetParams, opts ...option.RequestOption) (res *LinkedBankAccountV1, err error) {
	if !param.IsOmitted(query.CorrelationID) {
		opts = append(opts, option.WithHeader("correlation-id", fmt.Sprintf("%v", query.CorrelationID.Value)))
	}
	if !param.IsOmitted(query.RequestID) {
		opts = append(opts, option.WithHeader("request-id", fmt.Sprintf("%v", query.RequestID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	if linkedBankAccountID == "" {
		err = errors.New("missing required linked_bank_account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/linked_bank_accounts/%s", linkedBankAccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Retrieves the unmasked details of a linked bank account that has previously been
// created. Supply the unique linked bank account `id`, and Straddle will return
// the corresponding information, including sensitive details. This endpoint needs
// to be enabled by Straddle for your account and should only be used when
// absolutely necessary.
func (r *EmbedLinkedBankAccountService) Unmask(ctx context.Context, linkedBankAccountID string, query EmbedLinkedBankAccountUnmaskParams, opts ...option.RequestOption) (res *LinkedBankAccountUnmaskV1, err error) {
	if !param.IsOmitted(query.CorrelationID) {
		opts = append(opts, option.WithHeader("correlation-id", fmt.Sprintf("%v", query.CorrelationID.Value)))
	}
	if !param.IsOmitted(query.RequestID) {
		opts = append(opts, option.WithHeader("request-id", fmt.Sprintf("%v", query.RequestID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	if linkedBankAccountID == "" {
		err = errors.New("missing required linked_bank_account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/linked_bank_accounts/%s/unmask", linkedBankAccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type LinkedBankAccountPagedV1 struct {
	Data []LinkedBankAccountPagedV1Data `json:"data" api:"required"`
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
	ResponseType LinkedBankAccountPagedV1ResponseType `json:"response_type" api:"required"`
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
func (r LinkedBankAccountPagedV1) RawJSON() string { return r.JSON.raw }
func (r *LinkedBankAccountPagedV1) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LinkedBankAccountPagedV1Data struct {
	// Unique identifier for the linked bank account.
	ID string `json:"id" api:"required" format:"uuid"`
	// The unique identifier of the Straddle account related to this bank account.
	AccountID   string                                  `json:"account_id" api:"required" format:"uuid"`
	BankAccount LinkedBankAccountPagedV1DataBankAccount `json:"bank_account" api:"required"`
	// Timestamp of when the bank account object was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// The purposes for the linked bank account.
	//
	// Any of "charges", "payouts", "billing".
	Purposes []string `json:"purposes" api:"required"`
	// The current status of the linked bank account.
	//
	// Any of "created", "onboarding", "active", "rejected", "inactive", "canceled".
	Status       string                                   `json:"status" api:"required"`
	StatusDetail LinkedBankAccountPagedV1DataStatusDetail `json:"status_detail" api:"required"`
	// Timestamp of the most recent update to the linked bank account.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// Optional description for the bank account.
	Description string `json:"description" api:"nullable"`
	// Up to 20 additional user-defined key-value pairs. Useful for storing additional
	// information about the linked bank account in a structured format.
	Metadata map[string]string `json:"metadata" api:"nullable"`
	// The unique identifier of the Straddle Platform relatd to this bank account.
	PlatformID string `json:"platform_id" api:"nullable" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		AccountID    respjson.Field
		BankAccount  respjson.Field
		CreatedAt    respjson.Field
		Purposes     respjson.Field
		Status       respjson.Field
		StatusDetail respjson.Field
		UpdatedAt    respjson.Field
		Description  respjson.Field
		Metadata     respjson.Field
		PlatformID   respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LinkedBankAccountPagedV1Data) RawJSON() string { return r.JSON.raw }
func (r *LinkedBankAccountPagedV1Data) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LinkedBankAccountPagedV1DataBankAccount struct {
	AccountHolder   string `json:"account_holder" api:"required"`
	AccountMask     string `json:"account_mask" api:"required"`
	InstitutionName string `json:"institution_name" api:"required"`
	RoutingNumber   string `json:"routing_number" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccountHolder   respjson.Field
		AccountMask     respjson.Field
		InstitutionName respjson.Field
		RoutingNumber   respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LinkedBankAccountPagedV1DataBankAccount) RawJSON() string { return r.JSON.raw }
func (r *LinkedBankAccountPagedV1DataBankAccount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LinkedBankAccountPagedV1DataStatusDetail struct {
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
func (r LinkedBankAccountPagedV1DataStatusDetail) RawJSON() string { return r.JSON.raw }
func (r *LinkedBankAccountPagedV1DataStatusDetail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
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

type LinkedBankAccountUnmaskV1 struct {
	Data LinkedBankAccountUnmaskV1Data `json:"data" api:"required"`
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
	ResponseType LinkedBankAccountUnmaskV1ResponseType `json:"response_type" api:"required"`
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
func (r LinkedBankAccountUnmaskV1) RawJSON() string { return r.JSON.raw }
func (r *LinkedBankAccountUnmaskV1) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LinkedBankAccountUnmaskV1Data struct {
	// Unique identifier for the linked bank account.
	ID string `json:"id" api:"required" format:"uuid"`
	// Unique identifier for the Straddle account related to this bank account.
	AccountID string `json:"account_id" api:"required" format:"uuid"`
	// The bank account details associated with the linked bank account.
	BankAccount LinkedBankAccountUnmaskV1DataBankAccount `json:"bank_account" api:"required"`
	// Timestamp of when the linked bank account was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// The current status of the linked bank account.
	//
	// Any of "created", "onboarding", "active", "rejected", "inactive", "canceled".
	Status string `json:"status" api:"required"`
	// Additional details about the current status of the linked bank account.
	StatusDetail LinkedBankAccountUnmaskV1DataStatusDetail `json:"status_detail" api:"required"`
	// Timestamp of when the linked bank account was last updated.
	UpdatedAt time.Time         `json:"updated_at" api:"required" format:"date-time"`
	Metadata  map[string]string `json:"metadata" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		AccountID    respjson.Field
		BankAccount  respjson.Field
		CreatedAt    respjson.Field
		Status       respjson.Field
		StatusDetail respjson.Field
		UpdatedAt    respjson.Field
		Metadata     respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LinkedBankAccountUnmaskV1Data) RawJSON() string { return r.JSON.raw }
func (r *LinkedBankAccountUnmaskV1Data) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The bank account details associated with the linked bank account.
type LinkedBankAccountUnmaskV1DataBankAccount struct {
	AccountHolder   string `json:"account_holder" api:"required"`
	AccountNumber   string `json:"account_number" api:"required"`
	InstitutionName string `json:"institution_name" api:"required"`
	RoutingNumber   string `json:"routing_number" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccountHolder   respjson.Field
		AccountNumber   respjson.Field
		InstitutionName respjson.Field
		RoutingNumber   respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LinkedBankAccountUnmaskV1DataBankAccount) RawJSON() string { return r.JSON.raw }
func (r *LinkedBankAccountUnmaskV1DataBankAccount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Additional details about the current status of the linked bank account.
type LinkedBankAccountUnmaskV1DataStatusDetail struct {
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
func (r LinkedBankAccountUnmaskV1DataStatusDetail) RawJSON() string { return r.JSON.raw }
func (r *LinkedBankAccountUnmaskV1DataStatusDetail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
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

type LinkedBankAccountV1 struct {
	Data LinkedBankAccountV1Data `json:"data" api:"required"`
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
	ResponseType LinkedBankAccountV1ResponseType `json:"response_type" api:"required"`
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
func (r LinkedBankAccountV1) RawJSON() string { return r.JSON.raw }
func (r *LinkedBankAccountV1) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LinkedBankAccountV1Data struct {
	// Unique identifier for the linked bank account.
	ID string `json:"id" api:"required" format:"uuid"`
	// The unique identifier of the Straddle account related to this bank account.
	AccountID   string                             `json:"account_id" api:"required" format:"uuid"`
	BankAccount LinkedBankAccountV1DataBankAccount `json:"bank_account" api:"required"`
	// Timestamp of when the bank account object was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// The purposes for the linked bank account.
	//
	// Any of "charges", "payouts", "billing".
	Purposes []string `json:"purposes" api:"required"`
	// The current status of the linked bank account.
	//
	// Any of "created", "onboarding", "active", "rejected", "inactive", "canceled".
	Status       string                              `json:"status" api:"required"`
	StatusDetail LinkedBankAccountV1DataStatusDetail `json:"status_detail" api:"required"`
	// Timestamp of the most recent update to the linked bank account.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// Optional description for the bank account.
	Description string `json:"description" api:"nullable"`
	// Up to 20 additional user-defined key-value pairs. Useful for storing additional
	// information about the linked bank account in a structured format.
	Metadata map[string]string `json:"metadata" api:"nullable"`
	// The unique identifier of the Straddle Platform relatd to this bank account.
	PlatformID string `json:"platform_id" api:"nullable" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		AccountID    respjson.Field
		BankAccount  respjson.Field
		CreatedAt    respjson.Field
		Purposes     respjson.Field
		Status       respjson.Field
		StatusDetail respjson.Field
		UpdatedAt    respjson.Field
		Description  respjson.Field
		Metadata     respjson.Field
		PlatformID   respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LinkedBankAccountV1Data) RawJSON() string { return r.JSON.raw }
func (r *LinkedBankAccountV1Data) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LinkedBankAccountV1DataBankAccount struct {
	AccountHolder   string `json:"account_holder" api:"required"`
	AccountMask     string `json:"account_mask" api:"required"`
	InstitutionName string `json:"institution_name" api:"required"`
	RoutingNumber   string `json:"routing_number" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccountHolder   respjson.Field
		AccountMask     respjson.Field
		InstitutionName respjson.Field
		RoutingNumber   respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LinkedBankAccountV1DataBankAccount) RawJSON() string { return r.JSON.raw }
func (r *LinkedBankAccountV1DataBankAccount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LinkedBankAccountV1DataStatusDetail struct {
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
func (r LinkedBankAccountV1DataStatusDetail) RawJSON() string { return r.JSON.raw }
func (r *LinkedBankAccountV1DataStatusDetail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
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

type EmbedLinkedBankAccountNewParams struct {
	// The unique identifier of the Straddle account to associate this bank account
	// with.
	AccountID   param.Opt[string]                          `json:"account_id,omitzero" api:"required" format:"uuid"`
	BankAccount EmbedLinkedBankAccountNewParamsBankAccount `json:"bank_account,omitzero" api:"required"`
	// Optional description for the bank account.
	Description param.Opt[string] `json:"description,omitzero"`
	// The unique identifier of the Straddle Platform to associate this bank account
	// with.
	PlatformID     param.Opt[string] `json:"platform_id,omitzero" format:"uuid"`
	CorrelationID  param.Opt[string] `header:"correlation-id,omitzero" json:"-"`
	IdempotencyKey param.Opt[string] `header:"idempotency-key,omitzero" json:"-"`
	RequestID      param.Opt[string] `header:"request-id,omitzero" json:"-"`
	// Up to 20 additional user-defined key-value pairs. Useful for storing additional
	// information about the linked bank account in a structured format.
	Metadata map[string]string `json:"metadata,omitzero"`
	// The purposes for the linked bank account.
	//
	// Any of "charges", "payouts", "billing".
	Purposes []string `json:"purposes,omitzero"`
	paramObj
}

func (r EmbedLinkedBankAccountNewParams) MarshalJSON() (data []byte, err error) {
	type shadow EmbedLinkedBankAccountNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EmbedLinkedBankAccountNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties AccountHolder, AccountNumber, RoutingNumber are required.
type EmbedLinkedBankAccountNewParamsBankAccount struct {
	// The name of the account holder as it appears on the bank account. Typically,
	// this is the legal name of the business associated with the account.
	AccountHolder string `json:"account_holder" api:"required"`
	// The bank account number.
	AccountNumber string `json:"account_number" api:"required"`
	// The routing number of the bank account.
	RoutingNumber string `json:"routing_number" api:"required"`
	paramObj
}

func (r EmbedLinkedBankAccountNewParamsBankAccount) MarshalJSON() (data []byte, err error) {
	type shadow EmbedLinkedBankAccountNewParamsBankAccount
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EmbedLinkedBankAccountNewParamsBankAccount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbedLinkedBankAccountUpdateParams struct {
	BankAccount    EmbedLinkedBankAccountUpdateParamsBankAccount `json:"bank_account,omitzero" api:"required"`
	CorrelationID  param.Opt[string]                             `header:"correlation-id,omitzero" json:"-"`
	IdempotencyKey param.Opt[string]                             `header:"idempotency-key,omitzero" json:"-"`
	RequestID      param.Opt[string]                             `header:"request-id,omitzero" json:"-"`
	// Up to 20 additional user-defined key-value pairs. Useful for storing additional
	// information about the linked bank account in a structured format.
	Metadata map[string]string `json:"metadata,omitzero"`
	paramObj
}

func (r EmbedLinkedBankAccountUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow EmbedLinkedBankAccountUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EmbedLinkedBankAccountUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties AccountHolder, AccountNumber, RoutingNumber are required.
type EmbedLinkedBankAccountUpdateParamsBankAccount struct {
	// The name of the account holder as it appears on the bank account. Typically,
	// this is the legal name of the business associated with the account.
	AccountHolder string `json:"account_holder" api:"required"`
	// The bank account number.
	AccountNumber string `json:"account_number" api:"required"`
	// The routing number of the bank account.
	RoutingNumber string `json:"routing_number" api:"required"`
	paramObj
}

func (r EmbedLinkedBankAccountUpdateParamsBankAccount) MarshalJSON() (data []byte, err error) {
	type shadow EmbedLinkedBankAccountUpdateParamsBankAccount
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EmbedLinkedBankAccountUpdateParamsBankAccount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbedLinkedBankAccountListParams struct {
	// The unique identifier of the related account.
	AccountID param.Opt[string] `query:"account_id,omitzero" format:"uuid" json:"-"`
	// Results page number. Starts at page 1.
	PageNumber param.Opt[int64] `query:"page_number,omitzero" json:"-"`
	// Page size. Max value: 1000
	PageSize param.Opt[int64] `query:"page_size,omitzero" json:"-"`
	// Sort By.
	SortBy        param.Opt[string] `query:"sort_by,omitzero" json:"-"`
	CorrelationID param.Opt[string] `header:"correlation-id,omitzero" json:"-"`
	RequestID     param.Opt[string] `header:"request-id,omitzero" json:"-"`
	// Any of "account", "platform".
	Level EmbedLinkedBankAccountListParamsLevel `query:"level,omitzero" json:"-"`
	// The purpose of the linked bank accounts to return. Possible values: 'charges',
	// 'payouts', 'billing'.
	//
	// Any of "charges", "payouts", "billing".
	Purpose EmbedLinkedBankAccountListParamsPurpose `query:"purpose,omitzero" json:"-"`
	// Sort Order.
	//
	// Any of "asc", "desc".
	SortOrder EmbedLinkedBankAccountListParamsSortOrder `query:"sort_order,omitzero" json:"-"`
	// The status of the linked bank accounts to return. Possible values: 'created',
	// 'onboarding', 'active', 'inactive', 'rejected'.
	//
	// Any of "created", "onboarding", "active", "rejected", "inactive", "canceled".
	Status EmbedLinkedBankAccountListParamsStatus `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [EmbedLinkedBankAccountListParams]'s query parameters as
// `url.Values`.
func (r EmbedLinkedBankAccountListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type EmbedLinkedBankAccountListParamsLevel string

const (
	EmbedLinkedBankAccountListParamsLevelAccount  EmbedLinkedBankAccountListParamsLevel = "account"
	EmbedLinkedBankAccountListParamsLevelPlatform EmbedLinkedBankAccountListParamsLevel = "platform"
)

// The purpose of the linked bank accounts to return. Possible values: 'charges',
// 'payouts', 'billing'.
type EmbedLinkedBankAccountListParamsPurpose string

const (
	EmbedLinkedBankAccountListParamsPurposeCharges EmbedLinkedBankAccountListParamsPurpose = "charges"
	EmbedLinkedBankAccountListParamsPurposePayouts EmbedLinkedBankAccountListParamsPurpose = "payouts"
	EmbedLinkedBankAccountListParamsPurposeBilling EmbedLinkedBankAccountListParamsPurpose = "billing"
)

// Sort Order.
type EmbedLinkedBankAccountListParamsSortOrder string

const (
	EmbedLinkedBankAccountListParamsSortOrderAsc  EmbedLinkedBankAccountListParamsSortOrder = "asc"
	EmbedLinkedBankAccountListParamsSortOrderDesc EmbedLinkedBankAccountListParamsSortOrder = "desc"
)

// The status of the linked bank accounts to return. Possible values: 'created',
// 'onboarding', 'active', 'inactive', 'rejected'.
type EmbedLinkedBankAccountListParamsStatus string

const (
	EmbedLinkedBankAccountListParamsStatusCreated    EmbedLinkedBankAccountListParamsStatus = "created"
	EmbedLinkedBankAccountListParamsStatusOnboarding EmbedLinkedBankAccountListParamsStatus = "onboarding"
	EmbedLinkedBankAccountListParamsStatusActive     EmbedLinkedBankAccountListParamsStatus = "active"
	EmbedLinkedBankAccountListParamsStatusRejected   EmbedLinkedBankAccountListParamsStatus = "rejected"
	EmbedLinkedBankAccountListParamsStatusInactive   EmbedLinkedBankAccountListParamsStatus = "inactive"
	EmbedLinkedBankAccountListParamsStatusCanceled   EmbedLinkedBankAccountListParamsStatus = "canceled"
)

type EmbedLinkedBankAccountCancelParams struct {
	CorrelationID  param.Opt[string] `header:"correlation-id,omitzero" json:"-"`
	IdempotencyKey param.Opt[string] `header:"idempotency-key,omitzero" json:"-"`
	RequestID      param.Opt[string] `header:"request-id,omitzero" json:"-"`
	paramObj
}

type EmbedLinkedBankAccountGetParams struct {
	CorrelationID param.Opt[string] `header:"correlation-id,omitzero" json:"-"`
	RequestID     param.Opt[string] `header:"request-id,omitzero" json:"-"`
	paramObj
}

type EmbedLinkedBankAccountUnmaskParams struct {
	CorrelationID param.Opt[string] `header:"correlation-id,omitzero" json:"-"`
	RequestID     param.Opt[string] `header:"request-id,omitzero" json:"-"`
	paramObj
}
