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
func (r *PaykeyService) Get(ctx context.Context, id string, query PaykeyGetParams, opts ...option.RequestOption) (res *shared.PaykeyV1ItemResponse, err error) {
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
	Data []PaykeySummaryPagedV1Data    `json:"data,required"`
	Meta shared.PagedResponseMetadata1 `json:"meta,required"`
	// Indicates the structure of the returned content.
	//
	//   - "object" means the `data` field contains a single JSON object.
	//   - "array" means the `data` field contains an array of objects.
	//   - "error" means the `data` field contains an error object with details of the
	//     issue.
	//   - "none" means no data is returned.
	ResponseType shared.ResponseTypeEnum  `json:"response_type,required"`
	JSON         paykeySummaryPagedV1JSON `json:"-"`
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
	Paykey string                `json:"paykey,required"`
	Source shared.PaykeySourceV1 `json:"source,required"`
	Status shared.PaykeyStatusV1 `json:"status,required"`
	// Timestamp of the most recent update to the paykey.
	UpdatedAt time.Time                  `json:"updated_at,required" format:"date-time"`
	BankData  shared.PaykeyBankDetailsV1 `json:"bank_data"`
	// Unique identifier of the related customer object.
	CustomerID string `json:"customer_id,nullable" format:"uuid"`
	// Expiration date and time of the paykey, if applicable.
	ExpiresAt time.Time `json:"expires_at,nullable" format:"date-time"`
	// Name of the financial institution.
	InstitutionName string                       `json:"institution_name,nullable"`
	StatusDetails   shared.StatusDetailsV1       `json:"status_details"`
	JSON            paykeySummaryPagedV1DataJSON `json:"-"`
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
	ResponseType shared.ResponseTypeEnum `json:"response_type,required"`
	JSON         paykeyUnmaskedV1JSON    `json:"-"`
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
	Paykey string                `json:"paykey,required"`
	Source shared.PaykeySourceV1 `json:"source,required"`
	Status shared.PaykeyStatusV1 `json:"status,required"`
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
	Metadata      map[string]string        `json:"metadata,nullable"`
	StatusDetails shared.StatusDetailsV1   `json:"status_details"`
	JSON          paykeyUnmaskedV1DataJSON `json:"-"`
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

type PaykeyUnmaskedV1DataBankData struct {
	// The bank account number. This value is masked by default for security reasons.
	// Use the /unmask endpoint to access the unmasked value.
	AccountNumber string               `json:"account_number,required"`
	AccountType   shared.AccountTypeV1 `json:"account_type,required"`
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
	ResponseType shared.ResponseTypeEnum  `json:"response_type,required"`
	JSON         paykeyRevealResponseJSON `json:"-"`
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
	Paykey string                `json:"paykey,required"`
	Source shared.PaykeySourceV1 `json:"source,required"`
	Status shared.PaykeyStatusV1 `json:"status,required"`
	// Timestamp of the most recent update to the paykey.
	UpdatedAt time.Time                  `json:"updated_at,required" format:"date-time"`
	BankData  shared.PaykeyBankDetailsV1 `json:"bank_data"`
	// Unique identifier of the related customer object.
	CustomerID string `json:"customer_id,nullable" format:"uuid"`
	// Expiration date and time of the paykey, if applicable.
	ExpiresAt time.Time `json:"expires_at,nullable" format:"date-time"`
	// Name of the financial institution.
	InstitutionName string `json:"institution_name,nullable"`
	// Up to 20 additional user-defined key-value pairs. Useful for storing additional
	// information about the paykey in a structured format.
	Metadata      map[string]string            `json:"metadata,nullable"`
	StatusDetails shared.StatusDetailsV1       `json:"status_details"`
	JSON          paykeyRevealResponseDataJSON `json:"-"`
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

type PaykeyListParams struct {
	// Filter paykeys by related customer ID.
	CustomerID param.Field[string] `query:"customer_id" format:"uuid"`
	// Page number for paginated results. Starts at 1.
	PageNumber param.Field[int64] `query:"page_number"`
	// Number of results per page. Maximum: 1000.
	PageSize  param.Field[int64]                  `query:"page_size"`
	SortBy    param.Field[PaykeyListParamsSortBy] `query:"sort_by"`
	SortOrder param.Field[shared.SortOrder]       `query:"sort_order"`
	// Filter paykeys by their current status.
	Status            param.Field[[]shared.PaykeyStatusV1] `query:"status"`
	CorrelationID     param.Field[string]                  `header:"Correlation-Id"`
	RequestID         param.Field[string]                  `header:"Request-Id"`
	StraddleAccountID param.Field[string]                  `header:"Straddle-Account-Id" format:"uuid"`
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
