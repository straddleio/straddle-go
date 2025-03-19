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

// PaykeyService contains methods and other services that help with interacting
// with the straddle API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPaykeyService] method instead.
type PaykeyService struct {
	Options []option.RequestOption
}

// NewPaykeyService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewPaykeyService(opts ...option.RequestOption) (r *PaykeyService) {
	r = &PaykeyService{}
	r.Options = opts
	return
}

// Returns a list of paykeys associated with a Straddle account. This endpoint
// supports advanced sorting and filtering options.
func (r *PaykeyService) List(ctx context.Context, params PaykeyListParams, opts ...option.RequestOption) (res *pagination.PageNumberSchema[PaykeySummaryPagedV1Data], err error) {
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
	path := "v1/paykeys"
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

// Returns a list of paykeys associated with a Straddle account. This endpoint
// supports advanced sorting and filtering options.
func (r *PaykeyService) ListAutoPaging(ctx context.Context, params PaykeyListParams, opts ...option.RequestOption) *pagination.PageNumberSchemaAutoPager[PaykeySummaryPagedV1Data] {
	return pagination.NewPageNumberSchemaAutoPager(r.List(ctx, params, opts...))
}

// Retrieves the details of an existing paykey. Supply the unique paykey `id` and
// Straddle will return the corresponding paykey record , including the `paykey`
// token value and masked bank account details.
func (r *PaykeyService) Get(ctx context.Context, id string, query PaykeyGetParams, opts ...option.RequestOption) (res *PaykeyV1, err error) {
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
	path := fmt.Sprintf("v1/paykeys/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Retrieves the details of a paykey that has previously been created, including
// unmasked bank account fields. Supply the unique paykey ID that was returned from
// your previous request, and Straddle will return the corresponding paykey
// information.
func (r *PaykeyService) Reveal(ctx context.Context, id string, query PaykeyRevealParams, opts ...option.RequestOption) (res *PaykeyRevealResponse, err error) {
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
	path := fmt.Sprintf("v1/paykeys/%s/reveal", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Retrieves the unmasked details of an existing paykey. Supply the unique paykey
// `id` and Straddle will return the corresponding paykey record, including the
// unmasked bank account details. This endpoint needs to be enabled by Straddle for
// your account and should only be used when absolutely necessary.
func (r *PaykeyService) Unmasked(ctx context.Context, id string, query PaykeyUnmaskedParams, opts ...option.RequestOption) (res *PaykeyUnmaskedV1, err error) {
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
	path := fmt.Sprintf("v1/paykeys/%s/unmasked", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type PaykeySummaryPagedV1 struct {
	Data []PaykeySummaryPagedV1Data `json:"data,required"`
	Meta PaykeySummaryPagedV1Meta   `json:"meta,required"`
	// Indicates the structure of the returned content.
	//
	//   - "object" means the `data` field contains a single JSON object.
	//   - "array" means the `data` field contains an array of objects.
	//   - "error" means the `data` field contains an error object with details of the
	//     issue.
	//   - "none" means no data is returned.
	ResponseType PaykeySummaryPagedV1ResponseType `json:"response_type,required"`
	JSON         paykeySummaryPagedV1JSON         `json:"-"`
}

// paykeySummaryPagedV1JSON contains the JSON metadata for the struct
// [PaykeySummaryPagedV1]
type paykeySummaryPagedV1JSON struct {
	Data         apijson.Field
	Meta         apijson.Field
	ResponseType apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *PaykeySummaryPagedV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r paykeySummaryPagedV1JSON) RawJSON() string {
	return r.raw
}

type PaykeySummaryPagedV1Data struct {
	// Unique identifier for the paykey.
	ID string `json:"id,required" format:"uuid"`
	// Timestamp of when the paykey was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Human-readable label used to represent this paykey in a UI.
	Label string `json:"label,required"`
	// The tokenized paykey value. This value is used to create payments and should be
	// stored securely.
	Paykey string                         `json:"paykey,required"`
	Source PaykeySummaryPagedV1DataSource `json:"source,required"`
	Status PaykeySummaryPagedV1DataStatus `json:"status,required"`
	// Timestamp of the most recent update to the paykey.
	UpdatedAt time.Time                        `json:"updated_at,required" format:"date-time"`
	BankData  PaykeySummaryPagedV1DataBankData `json:"bank_data"`
	// Unique identifier of the related customer object.
	CustomerID string `json:"customer_id,nullable" format:"uuid"`
	// Expiration date and time of the paykey, if applicable.
	ExpiresAt time.Time `json:"expires_at,nullable" format:"date-time"`
	// Name of the financial institution.
	InstitutionName string                                `json:"institution_name,nullable"`
	StatusDetails   PaykeySummaryPagedV1DataStatusDetails `json:"status_details"`
	JSON            paykeySummaryPagedV1DataJSON          `json:"-"`
}

// paykeySummaryPagedV1DataJSON contains the JSON metadata for the struct
// [PaykeySummaryPagedV1Data]
type paykeySummaryPagedV1DataJSON struct {
	ID              apijson.Field
	CreatedAt       apijson.Field
	Label           apijson.Field
	Paykey          apijson.Field
	Source          apijson.Field
	Status          apijson.Field
	UpdatedAt       apijson.Field
	BankData        apijson.Field
	CustomerID      apijson.Field
	ExpiresAt       apijson.Field
	InstitutionName apijson.Field
	StatusDetails   apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *PaykeySummaryPagedV1Data) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r paykeySummaryPagedV1DataJSON) RawJSON() string {
	return r.raw
}

type PaykeySummaryPagedV1DataSource string

const (
	PaykeySummaryPagedV1DataSourceBankAccount PaykeySummaryPagedV1DataSource = "bank_account"
	PaykeySummaryPagedV1DataSourceStraddle    PaykeySummaryPagedV1DataSource = "straddle"
	PaykeySummaryPagedV1DataSourceMx          PaykeySummaryPagedV1DataSource = "mx"
	PaykeySummaryPagedV1DataSourcePlaid       PaykeySummaryPagedV1DataSource = "plaid"
)

func (r PaykeySummaryPagedV1DataSource) IsKnown() bool {
	switch r {
	case PaykeySummaryPagedV1DataSourceBankAccount, PaykeySummaryPagedV1DataSourceStraddle, PaykeySummaryPagedV1DataSourceMx, PaykeySummaryPagedV1DataSourcePlaid:
		return true
	}
	return false
}

type PaykeySummaryPagedV1DataStatus string

const (
	PaykeySummaryPagedV1DataStatusPending  PaykeySummaryPagedV1DataStatus = "pending"
	PaykeySummaryPagedV1DataStatusActive   PaykeySummaryPagedV1DataStatus = "active"
	PaykeySummaryPagedV1DataStatusInactive PaykeySummaryPagedV1DataStatus = "inactive"
	PaykeySummaryPagedV1DataStatusRejected PaykeySummaryPagedV1DataStatus = "rejected"
)

func (r PaykeySummaryPagedV1DataStatus) IsKnown() bool {
	switch r {
	case PaykeySummaryPagedV1DataStatusPending, PaykeySummaryPagedV1DataStatusActive, PaykeySummaryPagedV1DataStatusInactive, PaykeySummaryPagedV1DataStatusRejected:
		return true
	}
	return false
}

type PaykeySummaryPagedV1DataBankData struct {
	// Bank account number. This value is masked by default for security reasons. Use
	// the /unmask endpoint to access the unmasked value.
	AccountNumber string                                      `json:"account_number,required"`
	AccountType   PaykeySummaryPagedV1DataBankDataAccountType `json:"account_type,required"`
	// The routing number of the bank account.
	RoutingNumber string                               `json:"routing_number,required"`
	JSON          paykeySummaryPagedV1DataBankDataJSON `json:"-"`
}

// paykeySummaryPagedV1DataBankDataJSON contains the JSON metadata for the struct
// [PaykeySummaryPagedV1DataBankData]
type paykeySummaryPagedV1DataBankDataJSON struct {
	AccountNumber apijson.Field
	AccountType   apijson.Field
	RoutingNumber apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *PaykeySummaryPagedV1DataBankData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r paykeySummaryPagedV1DataBankDataJSON) RawJSON() string {
	return r.raw
}

type PaykeySummaryPagedV1DataBankDataAccountType string

const (
	PaykeySummaryPagedV1DataBankDataAccountTypeChecking PaykeySummaryPagedV1DataBankDataAccountType = "checking"
	PaykeySummaryPagedV1DataBankDataAccountTypeSavings  PaykeySummaryPagedV1DataBankDataAccountType = "savings"
)

func (r PaykeySummaryPagedV1DataBankDataAccountType) IsKnown() bool {
	switch r {
	case PaykeySummaryPagedV1DataBankDataAccountTypeChecking, PaykeySummaryPagedV1DataBankDataAccountTypeSavings:
		return true
	}
	return false
}

type PaykeySummaryPagedV1DataStatusDetails struct {
	// A human-readable description of the current status.
	Message string `json:"message,required"`
	// A machine-readable identifier for the specific status, useful for programmatic
	// handling.
	Reason string `json:"reason,required"`
	// Identifies the origin of the status change (e.g., `bank_decline`, `watchtower`).
	// This helps in tracking the cause of status updates.
	Source string                                    `json:"source,required"`
	JSON   paykeySummaryPagedV1DataStatusDetailsJSON `json:"-"`
}

// paykeySummaryPagedV1DataStatusDetailsJSON contains the JSON metadata for the
// struct [PaykeySummaryPagedV1DataStatusDetails]
type paykeySummaryPagedV1DataStatusDetailsJSON struct {
	Message     apijson.Field
	Reason      apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PaykeySummaryPagedV1DataStatusDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r paykeySummaryPagedV1DataStatusDetailsJSON) RawJSON() string {
	return r.raw
}

type PaykeySummaryPagedV1Meta struct {
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
	SortBy     string                            `json:"sort_by,required"`
	SortOrder  PaykeySummaryPagedV1MetaSortOrder `json:"sort_order,required"`
	TotalItems int64                             `json:"total_items,required"`
	// The number of pages available.
	TotalPages int64                        `json:"total_pages,required"`
	JSON       paykeySummaryPagedV1MetaJSON `json:"-"`
}

// paykeySummaryPagedV1MetaJSON contains the JSON metadata for the struct
// [PaykeySummaryPagedV1Meta]
type paykeySummaryPagedV1MetaJSON struct {
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

func (r *PaykeySummaryPagedV1Meta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r paykeySummaryPagedV1MetaJSON) RawJSON() string {
	return r.raw
}

type PaykeySummaryPagedV1MetaSortOrder string

const (
	PaykeySummaryPagedV1MetaSortOrderAsc  PaykeySummaryPagedV1MetaSortOrder = "asc"
	PaykeySummaryPagedV1MetaSortOrderDesc PaykeySummaryPagedV1MetaSortOrder = "desc"
)

func (r PaykeySummaryPagedV1MetaSortOrder) IsKnown() bool {
	switch r {
	case PaykeySummaryPagedV1MetaSortOrderAsc, PaykeySummaryPagedV1MetaSortOrderDesc:
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
type PaykeySummaryPagedV1ResponseType string

const (
	PaykeySummaryPagedV1ResponseTypeObject PaykeySummaryPagedV1ResponseType = "object"
	PaykeySummaryPagedV1ResponseTypeArray  PaykeySummaryPagedV1ResponseType = "array"
	PaykeySummaryPagedV1ResponseTypeError  PaykeySummaryPagedV1ResponseType = "error"
	PaykeySummaryPagedV1ResponseTypeNone   PaykeySummaryPagedV1ResponseType = "none"
)

func (r PaykeySummaryPagedV1ResponseType) IsKnown() bool {
	switch r {
	case PaykeySummaryPagedV1ResponseTypeObject, PaykeySummaryPagedV1ResponseTypeArray, PaykeySummaryPagedV1ResponseTypeError, PaykeySummaryPagedV1ResponseTypeNone:
		return true
	}
	return false
}

type PaykeyUnmaskedV1 struct {
	Data PaykeyUnmaskedV1Data `json:"data,required"`
	// Metadata about the API request, including an identifier and timestamp.
	Meta shared.ResponseMetadata `json:"meta,required"`
	// Indicates the structure of the returned content.
	//
	//   - "object" means the `data` field contains a single JSON object.
	//   - "array" means the `data` field contains an array of objects.
	//   - "error" means the `data` field contains an error object with details of the
	//     issue.
	//   - "none" means no data is returned.
	ResponseType PaykeyUnmaskedV1ResponseType `json:"response_type,required"`
	JSON         paykeyUnmaskedV1JSON         `json:"-"`
}

// paykeyUnmaskedV1JSON contains the JSON metadata for the struct
// [PaykeyUnmaskedV1]
type paykeyUnmaskedV1JSON struct {
	Data         apijson.Field
	Meta         apijson.Field
	ResponseType apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *PaykeyUnmaskedV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r paykeyUnmaskedV1JSON) RawJSON() string {
	return r.raw
}

type PaykeyUnmaskedV1Data struct {
	// Unique identifier for the paykey.
	ID string `json:"id,required" format:"uuid"`
	// Timestamp of when the paykey was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Human-readable label used to represent this paykey in a UI.
	Label string `json:"label,required"`
	// The tokenized paykey value. This value is used to create payments and should be
	// stored securely.
	Paykey string                     `json:"paykey,required"`
	Source PaykeyUnmaskedV1DataSource `json:"source,required"`
	Status PaykeyUnmaskedV1DataStatus `json:"status,required"`
	// Timestamp of the most recent update to the paykey.
	UpdatedAt time.Time                    `json:"updated_at,required" format:"date-time"`
	BankData  PaykeyUnmaskedV1DataBankData `json:"bank_data"`
	// Unique identifier of the related customer object.
	CustomerID string `json:"customer_id,nullable" format:"uuid"`
	// Expiration date and time of the paykey, if applicable.
	ExpiresAt time.Time `json:"expires_at,nullable" format:"date-time"`
	// Name of the financial institution.
	InstitutionName string `json:"institution_name,nullable"`
	// Up to 20 additional user-defined key-value pairs. Useful for storing additional
	// information about the paykey in a structured format.
	Metadata      map[string]string                 `json:"metadata,nullable"`
	StatusDetails PaykeyUnmaskedV1DataStatusDetails `json:"status_details"`
	JSON          paykeyUnmaskedV1DataJSON          `json:"-"`
}

// paykeyUnmaskedV1DataJSON contains the JSON metadata for the struct
// [PaykeyUnmaskedV1Data]
type paykeyUnmaskedV1DataJSON struct {
	ID              apijson.Field
	CreatedAt       apijson.Field
	Label           apijson.Field
	Paykey          apijson.Field
	Source          apijson.Field
	Status          apijson.Field
	UpdatedAt       apijson.Field
	BankData        apijson.Field
	CustomerID      apijson.Field
	ExpiresAt       apijson.Field
	InstitutionName apijson.Field
	Metadata        apijson.Field
	StatusDetails   apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *PaykeyUnmaskedV1Data) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r paykeyUnmaskedV1DataJSON) RawJSON() string {
	return r.raw
}

type PaykeyUnmaskedV1DataSource string

const (
	PaykeyUnmaskedV1DataSourceBankAccount PaykeyUnmaskedV1DataSource = "bank_account"
	PaykeyUnmaskedV1DataSourceStraddle    PaykeyUnmaskedV1DataSource = "straddle"
	PaykeyUnmaskedV1DataSourceMx          PaykeyUnmaskedV1DataSource = "mx"
	PaykeyUnmaskedV1DataSourcePlaid       PaykeyUnmaskedV1DataSource = "plaid"
)

func (r PaykeyUnmaskedV1DataSource) IsKnown() bool {
	switch r {
	case PaykeyUnmaskedV1DataSourceBankAccount, PaykeyUnmaskedV1DataSourceStraddle, PaykeyUnmaskedV1DataSourceMx, PaykeyUnmaskedV1DataSourcePlaid:
		return true
	}
	return false
}

type PaykeyUnmaskedV1DataStatus string

const (
	PaykeyUnmaskedV1DataStatusPending  PaykeyUnmaskedV1DataStatus = "pending"
	PaykeyUnmaskedV1DataStatusActive   PaykeyUnmaskedV1DataStatus = "active"
	PaykeyUnmaskedV1DataStatusInactive PaykeyUnmaskedV1DataStatus = "inactive"
	PaykeyUnmaskedV1DataStatusRejected PaykeyUnmaskedV1DataStatus = "rejected"
)

func (r PaykeyUnmaskedV1DataStatus) IsKnown() bool {
	switch r {
	case PaykeyUnmaskedV1DataStatusPending, PaykeyUnmaskedV1DataStatusActive, PaykeyUnmaskedV1DataStatusInactive, PaykeyUnmaskedV1DataStatusRejected:
		return true
	}
	return false
}

type PaykeyUnmaskedV1DataBankData struct {
	// The bank account number
	AccountNumber string                                  `json:"account_number,required"`
	AccountType   PaykeyUnmaskedV1DataBankDataAccountType `json:"account_type,required"`
	// The routing number of the bank account.
	RoutingNumber string                           `json:"routing_number,required"`
	JSON          paykeyUnmaskedV1DataBankDataJSON `json:"-"`
}

// paykeyUnmaskedV1DataBankDataJSON contains the JSON metadata for the struct
// [PaykeyUnmaskedV1DataBankData]
type paykeyUnmaskedV1DataBankDataJSON struct {
	AccountNumber apijson.Field
	AccountType   apijson.Field
	RoutingNumber apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *PaykeyUnmaskedV1DataBankData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r paykeyUnmaskedV1DataBankDataJSON) RawJSON() string {
	return r.raw
}

type PaykeyUnmaskedV1DataBankDataAccountType string

const (
	PaykeyUnmaskedV1DataBankDataAccountTypeChecking PaykeyUnmaskedV1DataBankDataAccountType = "checking"
	PaykeyUnmaskedV1DataBankDataAccountTypeSavings  PaykeyUnmaskedV1DataBankDataAccountType = "savings"
)

func (r PaykeyUnmaskedV1DataBankDataAccountType) IsKnown() bool {
	switch r {
	case PaykeyUnmaskedV1DataBankDataAccountTypeChecking, PaykeyUnmaskedV1DataBankDataAccountTypeSavings:
		return true
	}
	return false
}

type PaykeyUnmaskedV1DataStatusDetails struct {
	// A human-readable description of the current status.
	Message string `json:"message,required"`
	// A machine-readable identifier for the specific status, useful for programmatic
	// handling.
	Reason string `json:"reason,required"`
	// Identifies the origin of the status change (e.g., `bank_decline`, `watchtower`).
	// This helps in tracking the cause of status updates.
	Source string                                `json:"source,required"`
	JSON   paykeyUnmaskedV1DataStatusDetailsJSON `json:"-"`
}

// paykeyUnmaskedV1DataStatusDetailsJSON contains the JSON metadata for the struct
// [PaykeyUnmaskedV1DataStatusDetails]
type paykeyUnmaskedV1DataStatusDetailsJSON struct {
	Message     apijson.Field
	Reason      apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PaykeyUnmaskedV1DataStatusDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r paykeyUnmaskedV1DataStatusDetailsJSON) RawJSON() string {
	return r.raw
}

// Indicates the structure of the returned content.
//
//   - "object" means the `data` field contains a single JSON object.
//   - "array" means the `data` field contains an array of objects.
//   - "error" means the `data` field contains an error object with details of the
//     issue.
//   - "none" means no data is returned.
type PaykeyUnmaskedV1ResponseType string

const (
	PaykeyUnmaskedV1ResponseTypeObject PaykeyUnmaskedV1ResponseType = "object"
	PaykeyUnmaskedV1ResponseTypeArray  PaykeyUnmaskedV1ResponseType = "array"
	PaykeyUnmaskedV1ResponseTypeError  PaykeyUnmaskedV1ResponseType = "error"
	PaykeyUnmaskedV1ResponseTypeNone   PaykeyUnmaskedV1ResponseType = "none"
)

func (r PaykeyUnmaskedV1ResponseType) IsKnown() bool {
	switch r {
	case PaykeyUnmaskedV1ResponseTypeObject, PaykeyUnmaskedV1ResponseTypeArray, PaykeyUnmaskedV1ResponseTypeError, PaykeyUnmaskedV1ResponseTypeNone:
		return true
	}
	return false
}

type PaykeyV1 struct {
	Data PaykeyV1Data `json:"data,required"`
	// Metadata about the API request, including an identifier and timestamp.
	Meta shared.ResponseMetadata `json:"meta,required"`
	// Indicates the structure of the returned content.
	//
	//   - "object" means the `data` field contains a single JSON object.
	//   - "array" means the `data` field contains an array of objects.
	//   - "error" means the `data` field contains an error object with details of the
	//     issue.
	//   - "none" means no data is returned.
	ResponseType PaykeyV1ResponseType `json:"response_type,required"`
	JSON         paykeyV1JSON         `json:"-"`
}

// paykeyV1JSON contains the JSON metadata for the struct [PaykeyV1]
type paykeyV1JSON struct {
	Data         apijson.Field
	Meta         apijson.Field
	ResponseType apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *PaykeyV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r paykeyV1JSON) RawJSON() string {
	return r.raw
}

type PaykeyV1Data struct {
	// Unique identifier for the paykey.
	ID string `json:"id,required" format:"uuid"`
	// Timestamp of when the paykey was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Human-readable label used to represent this paykey in a UI.
	Label string `json:"label,required"`
	// The tokenized paykey value. This value is used to create payments and should be
	// stored securely.
	Paykey string             `json:"paykey,required"`
	Source PaykeyV1DataSource `json:"source,required"`
	Status PaykeyV1DataStatus `json:"status,required"`
	// Timestamp of the most recent update to the paykey.
	UpdatedAt time.Time            `json:"updated_at,required" format:"date-time"`
	BankData  PaykeyV1DataBankData `json:"bank_data"`
	// Unique identifier of the related customer object.
	CustomerID string `json:"customer_id,nullable" format:"uuid"`
	// Expiration date and time of the paykey, if applicable.
	ExpiresAt time.Time `json:"expires_at,nullable" format:"date-time"`
	// Name of the financial institution.
	InstitutionName string `json:"institution_name,nullable"`
	// Up to 20 additional user-defined key-value pairs. Useful for storing additional
	// information about the paykey in a structured format.
	Metadata      map[string]string         `json:"metadata,nullable"`
	StatusDetails PaykeyV1DataStatusDetails `json:"status_details"`
	JSON          paykeyV1DataJSON          `json:"-"`
}

// paykeyV1DataJSON contains the JSON metadata for the struct [PaykeyV1Data]
type paykeyV1DataJSON struct {
	ID              apijson.Field
	CreatedAt       apijson.Field
	Label           apijson.Field
	Paykey          apijson.Field
	Source          apijson.Field
	Status          apijson.Field
	UpdatedAt       apijson.Field
	BankData        apijson.Field
	CustomerID      apijson.Field
	ExpiresAt       apijson.Field
	InstitutionName apijson.Field
	Metadata        apijson.Field
	StatusDetails   apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *PaykeyV1Data) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r paykeyV1DataJSON) RawJSON() string {
	return r.raw
}

type PaykeyV1DataSource string

const (
	PaykeyV1DataSourceBankAccount PaykeyV1DataSource = "bank_account"
	PaykeyV1DataSourceStraddle    PaykeyV1DataSource = "straddle"
	PaykeyV1DataSourceMx          PaykeyV1DataSource = "mx"
	PaykeyV1DataSourcePlaid       PaykeyV1DataSource = "plaid"
)

func (r PaykeyV1DataSource) IsKnown() bool {
	switch r {
	case PaykeyV1DataSourceBankAccount, PaykeyV1DataSourceStraddle, PaykeyV1DataSourceMx, PaykeyV1DataSourcePlaid:
		return true
	}
	return false
}

type PaykeyV1DataStatus string

const (
	PaykeyV1DataStatusPending  PaykeyV1DataStatus = "pending"
	PaykeyV1DataStatusActive   PaykeyV1DataStatus = "active"
	PaykeyV1DataStatusInactive PaykeyV1DataStatus = "inactive"
	PaykeyV1DataStatusRejected PaykeyV1DataStatus = "rejected"
)

func (r PaykeyV1DataStatus) IsKnown() bool {
	switch r {
	case PaykeyV1DataStatusPending, PaykeyV1DataStatusActive, PaykeyV1DataStatusInactive, PaykeyV1DataStatusRejected:
		return true
	}
	return false
}

type PaykeyV1DataBankData struct {
	// Bank account number. This value is masked by default for security reasons. Use
	// the /unmask endpoint to access the unmasked value.
	AccountNumber string                          `json:"account_number,required"`
	AccountType   PaykeyV1DataBankDataAccountType `json:"account_type,required"`
	// The routing number of the bank account.
	RoutingNumber string                   `json:"routing_number,required"`
	JSON          paykeyV1DataBankDataJSON `json:"-"`
}

// paykeyV1DataBankDataJSON contains the JSON metadata for the struct
// [PaykeyV1DataBankData]
type paykeyV1DataBankDataJSON struct {
	AccountNumber apijson.Field
	AccountType   apijson.Field
	RoutingNumber apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *PaykeyV1DataBankData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r paykeyV1DataBankDataJSON) RawJSON() string {
	return r.raw
}

type PaykeyV1DataBankDataAccountType string

const (
	PaykeyV1DataBankDataAccountTypeChecking PaykeyV1DataBankDataAccountType = "checking"
	PaykeyV1DataBankDataAccountTypeSavings  PaykeyV1DataBankDataAccountType = "savings"
)

func (r PaykeyV1DataBankDataAccountType) IsKnown() bool {
	switch r {
	case PaykeyV1DataBankDataAccountTypeChecking, PaykeyV1DataBankDataAccountTypeSavings:
		return true
	}
	return false
}

type PaykeyV1DataStatusDetails struct {
	// A human-readable description of the current status.
	Message string `json:"message,required"`
	// A machine-readable identifier for the specific status, useful for programmatic
	// handling.
	Reason string `json:"reason,required"`
	// Identifies the origin of the status change (e.g., `bank_decline`, `watchtower`).
	// This helps in tracking the cause of status updates.
	Source string                        `json:"source,required"`
	JSON   paykeyV1DataStatusDetailsJSON `json:"-"`
}

// paykeyV1DataStatusDetailsJSON contains the JSON metadata for the struct
// [PaykeyV1DataStatusDetails]
type paykeyV1DataStatusDetailsJSON struct {
	Message     apijson.Field
	Reason      apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PaykeyV1DataStatusDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r paykeyV1DataStatusDetailsJSON) RawJSON() string {
	return r.raw
}

// Indicates the structure of the returned content.
//
//   - "object" means the `data` field contains a single JSON object.
//   - "array" means the `data` field contains an array of objects.
//   - "error" means the `data` field contains an error object with details of the
//     issue.
//   - "none" means no data is returned.
type PaykeyV1ResponseType string

const (
	PaykeyV1ResponseTypeObject PaykeyV1ResponseType = "object"
	PaykeyV1ResponseTypeArray  PaykeyV1ResponseType = "array"
	PaykeyV1ResponseTypeError  PaykeyV1ResponseType = "error"
	PaykeyV1ResponseTypeNone   PaykeyV1ResponseType = "none"
)

func (r PaykeyV1ResponseType) IsKnown() bool {
	switch r {
	case PaykeyV1ResponseTypeObject, PaykeyV1ResponseTypeArray, PaykeyV1ResponseTypeError, PaykeyV1ResponseTypeNone:
		return true
	}
	return false
}

type PaykeyRevealResponse struct {
	Data PaykeyRevealResponseData `json:"data,required"`
	// Metadata about the API request, including an identifier and timestamp.
	Meta shared.ResponseMetadata `json:"meta,required"`
	// Indicates the structure of the returned content.
	//
	//   - "object" means the `data` field contains a single JSON object.
	//   - "array" means the `data` field contains an array of objects.
	//   - "error" means the `data` field contains an error object with details of the
	//     issue.
	//   - "none" means no data is returned.
	ResponseType PaykeyRevealResponseResponseType `json:"response_type,required"`
	JSON         paykeyRevealResponseJSON         `json:"-"`
}

// paykeyRevealResponseJSON contains the JSON metadata for the struct
// [PaykeyRevealResponse]
type paykeyRevealResponseJSON struct {
	Data         apijson.Field
	Meta         apijson.Field
	ResponseType apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *PaykeyRevealResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r paykeyRevealResponseJSON) RawJSON() string {
	return r.raw
}

type PaykeyRevealResponseData struct {
	// Unique identifier for the paykey.
	ID string `json:"id,required" format:"uuid"`
	// Timestamp of when the paykey was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Human-readable label used to represent this paykey in a UI.
	Label string `json:"label,required"`
	// The tokenized paykey value. This value is used to create payments and should be
	// stored securely.
	Paykey string                         `json:"paykey,required"`
	Source PaykeyRevealResponseDataSource `json:"source,required"`
	Status PaykeyRevealResponseDataStatus `json:"status,required"`
	// Timestamp of the most recent update to the paykey.
	UpdatedAt time.Time                        `json:"updated_at,required" format:"date-time"`
	BankData  PaykeyRevealResponseDataBankData `json:"bank_data"`
	// Unique identifier of the related customer object.
	CustomerID string `json:"customer_id,nullable" format:"uuid"`
	// Expiration date and time of the paykey, if applicable.
	ExpiresAt time.Time `json:"expires_at,nullable" format:"date-time"`
	// Name of the financial institution.
	InstitutionName string `json:"institution_name,nullable"`
	// Up to 20 additional user-defined key-value pairs. Useful for storing additional
	// information about the paykey in a structured format.
	Metadata      map[string]string                     `json:"metadata,nullable"`
	StatusDetails PaykeyRevealResponseDataStatusDetails `json:"status_details"`
	JSON          paykeyRevealResponseDataJSON          `json:"-"`
}

// paykeyRevealResponseDataJSON contains the JSON metadata for the struct
// [PaykeyRevealResponseData]
type paykeyRevealResponseDataJSON struct {
	ID              apijson.Field
	CreatedAt       apijson.Field
	Label           apijson.Field
	Paykey          apijson.Field
	Source          apijson.Field
	Status          apijson.Field
	UpdatedAt       apijson.Field
	BankData        apijson.Field
	CustomerID      apijson.Field
	ExpiresAt       apijson.Field
	InstitutionName apijson.Field
	Metadata        apijson.Field
	StatusDetails   apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *PaykeyRevealResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r paykeyRevealResponseDataJSON) RawJSON() string {
	return r.raw
}

type PaykeyRevealResponseDataSource string

const (
	PaykeyRevealResponseDataSourceBankAccount PaykeyRevealResponseDataSource = "bank_account"
	PaykeyRevealResponseDataSourceStraddle    PaykeyRevealResponseDataSource = "straddle"
	PaykeyRevealResponseDataSourceMx          PaykeyRevealResponseDataSource = "mx"
	PaykeyRevealResponseDataSourcePlaid       PaykeyRevealResponseDataSource = "plaid"
)

func (r PaykeyRevealResponseDataSource) IsKnown() bool {
	switch r {
	case PaykeyRevealResponseDataSourceBankAccount, PaykeyRevealResponseDataSourceStraddle, PaykeyRevealResponseDataSourceMx, PaykeyRevealResponseDataSourcePlaid:
		return true
	}
	return false
}

type PaykeyRevealResponseDataStatus string

const (
	PaykeyRevealResponseDataStatusPending  PaykeyRevealResponseDataStatus = "pending"
	PaykeyRevealResponseDataStatusActive   PaykeyRevealResponseDataStatus = "active"
	PaykeyRevealResponseDataStatusInactive PaykeyRevealResponseDataStatus = "inactive"
	PaykeyRevealResponseDataStatusRejected PaykeyRevealResponseDataStatus = "rejected"
)

func (r PaykeyRevealResponseDataStatus) IsKnown() bool {
	switch r {
	case PaykeyRevealResponseDataStatusPending, PaykeyRevealResponseDataStatusActive, PaykeyRevealResponseDataStatusInactive, PaykeyRevealResponseDataStatusRejected:
		return true
	}
	return false
}

type PaykeyRevealResponseDataBankData struct {
	// Bank account number. This value is masked by default for security reasons. Use
	// the /unmask endpoint to access the unmasked value.
	AccountNumber string                                      `json:"account_number,required"`
	AccountType   PaykeyRevealResponseDataBankDataAccountType `json:"account_type,required"`
	// The routing number of the bank account.
	RoutingNumber string                               `json:"routing_number,required"`
	JSON          paykeyRevealResponseDataBankDataJSON `json:"-"`
}

// paykeyRevealResponseDataBankDataJSON contains the JSON metadata for the struct
// [PaykeyRevealResponseDataBankData]
type paykeyRevealResponseDataBankDataJSON struct {
	AccountNumber apijson.Field
	AccountType   apijson.Field
	RoutingNumber apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *PaykeyRevealResponseDataBankData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r paykeyRevealResponseDataBankDataJSON) RawJSON() string {
	return r.raw
}

type PaykeyRevealResponseDataBankDataAccountType string

const (
	PaykeyRevealResponseDataBankDataAccountTypeChecking PaykeyRevealResponseDataBankDataAccountType = "checking"
	PaykeyRevealResponseDataBankDataAccountTypeSavings  PaykeyRevealResponseDataBankDataAccountType = "savings"
)

func (r PaykeyRevealResponseDataBankDataAccountType) IsKnown() bool {
	switch r {
	case PaykeyRevealResponseDataBankDataAccountTypeChecking, PaykeyRevealResponseDataBankDataAccountTypeSavings:
		return true
	}
	return false
}

type PaykeyRevealResponseDataStatusDetails struct {
	// A human-readable description of the current status.
	Message string `json:"message,required"`
	// A machine-readable identifier for the specific status, useful for programmatic
	// handling.
	Reason string `json:"reason,required"`
	// Identifies the origin of the status change (e.g., `bank_decline`, `watchtower`).
	// This helps in tracking the cause of status updates.
	Source string                                    `json:"source,required"`
	JSON   paykeyRevealResponseDataStatusDetailsJSON `json:"-"`
}

// paykeyRevealResponseDataStatusDetailsJSON contains the JSON metadata for the
// struct [PaykeyRevealResponseDataStatusDetails]
type paykeyRevealResponseDataStatusDetailsJSON struct {
	Message     apijson.Field
	Reason      apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PaykeyRevealResponseDataStatusDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r paykeyRevealResponseDataStatusDetailsJSON) RawJSON() string {
	return r.raw
}

// Indicates the structure of the returned content.
//
//   - "object" means the `data` field contains a single JSON object.
//   - "array" means the `data` field contains an array of objects.
//   - "error" means the `data` field contains an error object with details of the
//     issue.
//   - "none" means no data is returned.
type PaykeyRevealResponseResponseType string

const (
	PaykeyRevealResponseResponseTypeObject PaykeyRevealResponseResponseType = "object"
	PaykeyRevealResponseResponseTypeArray  PaykeyRevealResponseResponseType = "array"
	PaykeyRevealResponseResponseTypeError  PaykeyRevealResponseResponseType = "error"
	PaykeyRevealResponseResponseTypeNone   PaykeyRevealResponseResponseType = "none"
)

func (r PaykeyRevealResponseResponseType) IsKnown() bool {
	switch r {
	case PaykeyRevealResponseResponseTypeObject, PaykeyRevealResponseResponseTypeArray, PaykeyRevealResponseResponseTypeError, PaykeyRevealResponseResponseTypeNone:
		return true
	}
	return false
}

type PaykeyListParams struct {
	// Filter paykeys by related customer ID.
	CustomerID param.Field[string] `query:"customer_id" format:"uuid"`
	// Page number for paginated results. Starts at 1.
	PageNumber param.Field[int64] `query:"page_number"`
	// Number of results per page. Maximum: 1000.
	PageSize  param.Field[int64]                     `query:"page_size"`
	SortBy    param.Field[PaykeyListParamsSortBy]    `query:"sort_by"`
	SortOrder param.Field[PaykeyListParamsSortOrder] `query:"sort_order"`
	// Filter paykeys by their current status.
	Status            param.Field[[]PaykeyListParamsStatus] `query:"status"`
	CorrelationID     param.Field[string]                   `header:"Correlation-Id"`
	RequestID         param.Field[string]                   `header:"Request-Id"`
	StraddleAccountID param.Field[string]                   `header:"Straddle-Account-Id" format:"uuid"`
}

// URLQuery serializes [PaykeyListParams]'s query parameters as `url.Values`.
func (r PaykeyListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PaykeyListParamsSortBy string

const (
	PaykeyListParamsSortByInstitutionName PaykeyListParamsSortBy = "institution_name"
	PaykeyListParamsSortByExpiresAt       PaykeyListParamsSortBy = "expires_at"
	PaykeyListParamsSortByCreatedAt       PaykeyListParamsSortBy = "created_at"
)

func (r PaykeyListParamsSortBy) IsKnown() bool {
	switch r {
	case PaykeyListParamsSortByInstitutionName, PaykeyListParamsSortByExpiresAt, PaykeyListParamsSortByCreatedAt:
		return true
	}
	return false
}

type PaykeyListParamsSortOrder string

const (
	PaykeyListParamsSortOrderAsc  PaykeyListParamsSortOrder = "asc"
	PaykeyListParamsSortOrderDesc PaykeyListParamsSortOrder = "desc"
)

func (r PaykeyListParamsSortOrder) IsKnown() bool {
	switch r {
	case PaykeyListParamsSortOrderAsc, PaykeyListParamsSortOrderDesc:
		return true
	}
	return false
}

type PaykeyListParamsStatus string

const (
	PaykeyListParamsStatusPending  PaykeyListParamsStatus = "pending"
	PaykeyListParamsStatusActive   PaykeyListParamsStatus = "active"
	PaykeyListParamsStatusInactive PaykeyListParamsStatus = "inactive"
	PaykeyListParamsStatusRejected PaykeyListParamsStatus = "rejected"
)

func (r PaykeyListParamsStatus) IsKnown() bool {
	switch r {
	case PaykeyListParamsStatusPending, PaykeyListParamsStatusActive, PaykeyListParamsStatusInactive, PaykeyListParamsStatusRejected:
		return true
	}
	return false
}

type PaykeyGetParams struct {
	CorrelationID     param.Field[string] `header:"Correlation-Id"`
	RequestID         param.Field[string] `header:"Request-Id"`
	StraddleAccountID param.Field[string] `header:"Straddle-Account-Id" format:"uuid"`
}

type PaykeyRevealParams struct {
	CorrelationID     param.Field[string] `header:"Correlation-Id"`
	RequestID         param.Field[string] `header:"Request-Id"`
	StraddleAccountID param.Field[string] `header:"Straddle-Account-Id" format:"uuid"`
}

type PaykeyUnmaskedParams struct {
	CorrelationID     param.Field[string] `header:"Correlation-Id"`
	RequestID         param.Field[string] `header:"Request-Id"`
	StraddleAccountID param.Field[string] `header:"Straddle-Account-Id" format:"uuid"`
}
