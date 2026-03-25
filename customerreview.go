// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package straddle

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/straddleio/straddle-go/internal/apijson"
	"github.com/straddleio/straddle-go/internal/requestconfig"
	"github.com/straddleio/straddle-go/option"
	"github.com/straddleio/straddle-go/packages/param"
	"github.com/straddleio/straddle-go/packages/respjson"
	"github.com/straddleio/straddle-go/shared"
)

// Customers represent the end users who send or receive payments through your
// integration. Each customer undergoes automatic identity verification and fraud
// screening upon creation. Use customers to track payment history, manage bank
// account connections, and maintain a secure record of all transactions associated
// with a user. Customers can be either individuals or businesses with appropriate
// compliance checks for each type.
//
// CustomerReviewService contains methods and other services that help with
// interacting with the straddle API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCustomerReviewService] method instead.
type CustomerReviewService struct {
	options []option.RequestOption
}

// NewCustomerReviewService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCustomerReviewService(opts ...option.RequestOption) (r CustomerReviewService) {
	r = CustomerReviewService{}
	r.options = opts
	return
}

// Updates the status of a customer's identity decision. This endpoint allows you
// to modify the outcome of a customer risk screening and is useful for correcting
// or updating the status of a customer's verification. Note that this endpoint is
// only available for customers with a current status of `review`.
func (r *CustomerReviewService) Decision(ctx context.Context, id string, params CustomerReviewDecisionParams, opts ...option.RequestOption) (res *CustomerV1, err error) {
	if !param.IsOmitted(params.CorrelationID) {
		opts = append(opts, option.WithHeader("Correlation-Id", fmt.Sprintf("%v", params.CorrelationID.Value)))
	}
	if !param.IsOmitted(params.IdempotencyKey) {
		opts = append(opts, option.WithHeader("Idempotency-Key", fmt.Sprintf("%v", params.IdempotencyKey.Value)))
	}
	if !param.IsOmitted(params.RequestID) {
		opts = append(opts, option.WithHeader("Request-Id", fmt.Sprintf("%v", params.RequestID.Value)))
	}
	if !param.IsOmitted(params.StraddleAccountID) {
		opts = append(opts, option.WithHeader("Straddle-Account-Id", fmt.Sprintf("%v", params.StraddleAccountID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/customers/%s/review", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Retrieves and analyzes the results of a customer's identity validation and fraud
// score. This endpoint provides a comprehensive breakdown of the validation
// outcome, including:
//
//   - Risk and correlation scores
//   - Reason codes for the decision
//   - Results of watchlist screening
//   - Any network alerts detected Use this endpoint to gain insights into the
//     verification process and make informed decisions about customer onboarding.
func (r *CustomerReviewService) Get(ctx context.Context, id string, query CustomerReviewGetParams, opts ...option.RequestOption) (res *CustomerReviewV1, err error) {
	if !param.IsOmitted(query.CorrelationID) {
		opts = append(opts, option.WithHeader("Correlation-Id", fmt.Sprintf("%v", query.CorrelationID.Value)))
	}
	if !param.IsOmitted(query.RequestID) {
		opts = append(opts, option.WithHeader("Request-Id", fmt.Sprintf("%v", query.RequestID.Value)))
	}
	if !param.IsOmitted(query.StraddleAccountID) {
		opts = append(opts, option.WithHeader("Straddle-Account-Id", fmt.Sprintf("%v", query.StraddleAccountID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/customers/%s/review", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Updates the decision of a customer's identity validation. This endpoint allows
// you to modify the outcome of a customer decision and is useful for correcting or
// updating the status of a customer's verification.
func (r *CustomerReviewService) RefreshReview(ctx context.Context, id string, body CustomerReviewRefreshReviewParams, opts ...option.RequestOption) (res *CustomerV1, err error) {
	if !param.IsOmitted(body.CorrelationID) {
		opts = append(opts, option.WithHeader("Correlation-Id", fmt.Sprintf("%v", body.CorrelationID.Value)))
	}
	if !param.IsOmitted(body.IdempotencyKey) {
		opts = append(opts, option.WithHeader("Idempotency-Key", fmt.Sprintf("%v", body.IdempotencyKey.Value)))
	}
	if !param.IsOmitted(body.RequestID) {
		opts = append(opts, option.WithHeader("Request-Id", fmt.Sprintf("%v", body.RequestID.Value)))
	}
	if !param.IsOmitted(body.StraddleAccountID) {
		opts = append(opts, option.WithHeader("Straddle-Account-Id", fmt.Sprintf("%v", body.StraddleAccountID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/customers/%s/refresh_review", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, nil, &res, opts...)
	return res, err
}

type CustomerReviewV1 struct {
	Data CustomerReviewV1Data `json:"data" api:"required"`
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
	ResponseType CustomerReviewV1ResponseType `json:"response_type" api:"required"`
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
func (r CustomerReviewV1) RawJSON() string { return r.JSON.raw }
func (r *CustomerReviewV1) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerReviewV1Data struct {
	CustomerDetails CustomerReviewV1DataCustomerDetails `json:"customer_details" api:"required"`
	IdentityDetails CustomerReviewV1DataIdentityDetails `json:"identity_details"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CustomerDetails respjson.Field
		IdentityDetails respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CustomerReviewV1Data) RawJSON() string { return r.JSON.raw }
func (r *CustomerReviewV1Data) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerReviewV1DataCustomerDetails struct {
	// Unique identifier for the customer.
	ID string `json:"id" api:"required" format:"uuid"`
	// Timestamp of when the customer record was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// The customer's email address.
	Email string `json:"email" api:"required" format:"email"`
	// Full name of the individual or business name.
	Name string `json:"name" api:"required"`
	// The customer's phone number in E.164 format.
	Phone string `json:"phone" api:"required"`
	// Any of "pending", "review", "verified", "inactive", "rejected".
	Status string `json:"status" api:"required"`
	// Any of "individual", "business".
	Type string `json:"type" api:"required"`
	// Timestamp of the most recent update to the customer record.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// An object containing the customer's address. This is optional, but if provided,
	// all required fields must be present.
	Address CustomerAddressV1 `json:"address" api:"nullable"`
	// PII required to trigger Patriot Act compliant KYC verification.
	ComplianceProfile CustomerReviewV1DataCustomerDetailsComplianceProfileUnion `json:"compliance_profile" api:"nullable"`
	Config            CustomerReviewV1DataCustomerDetailsConfig                 `json:"config"`
	Device            CustomerReviewV1DataCustomerDetailsDevice                 `json:"device"`
	// Unique identifier for the customer in your database, used for cross-referencing
	// between Straddle and your systems.
	ExternalID string `json:"external_id" api:"nullable"`
	// Up to 20 additional user-defined key-value pairs. Useful for storing additional
	// information about the customer in a structured format.
	Metadata map[string]string `json:"metadata" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		CreatedAt         respjson.Field
		Email             respjson.Field
		Name              respjson.Field
		Phone             respjson.Field
		Status            respjson.Field
		Type              respjson.Field
		UpdatedAt         respjson.Field
		Address           respjson.Field
		ComplianceProfile respjson.Field
		Config            respjson.Field
		Device            respjson.Field
		ExternalID        respjson.Field
		Metadata          respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CustomerReviewV1DataCustomerDetails) RawJSON() string { return r.JSON.raw }
func (r *CustomerReviewV1DataCustomerDetails) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// CustomerReviewV1DataCustomerDetailsComplianceProfileUnion contains all possible
// properties and values from
// [CustomerReviewV1DataCustomerDetailsComplianceProfileIndividualComplianceProfile],
// [CustomerReviewV1DataCustomerDetailsComplianceProfileBusinessComplianceProfile].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type CustomerReviewV1DataCustomerDetailsComplianceProfileUnion struct {
	// This field is from variant
	// [CustomerReviewV1DataCustomerDetailsComplianceProfileIndividualComplianceProfile].
	Dob time.Time `json:"dob"`
	// This field is from variant
	// [CustomerReviewV1DataCustomerDetailsComplianceProfileIndividualComplianceProfile].
	Ssn string `json:"ssn"`
	// This field is from variant
	// [CustomerReviewV1DataCustomerDetailsComplianceProfileBusinessComplianceProfile].
	Ein string `json:"ein"`
	// This field is from variant
	// [CustomerReviewV1DataCustomerDetailsComplianceProfileBusinessComplianceProfile].
	LegalBusinessName string `json:"legal_business_name"`
	// This field is from variant
	// [CustomerReviewV1DataCustomerDetailsComplianceProfileBusinessComplianceProfile].
	Representatives []CustomerReviewV1DataCustomerDetailsComplianceProfileBusinessComplianceProfileRepresentative `json:"representatives"`
	// This field is from variant
	// [CustomerReviewV1DataCustomerDetailsComplianceProfileBusinessComplianceProfile].
	Website string `json:"website"`
	JSON    struct {
		Dob               respjson.Field
		Ssn               respjson.Field
		Ein               respjson.Field
		LegalBusinessName respjson.Field
		Representatives   respjson.Field
		Website           respjson.Field
		raw               string
	} `json:"-"`
}

func (u CustomerReviewV1DataCustomerDetailsComplianceProfileUnion) AsIndividualComplianceProfile() (v CustomerReviewV1DataCustomerDetailsComplianceProfileIndividualComplianceProfile) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CustomerReviewV1DataCustomerDetailsComplianceProfileUnion) AsBusinessComplianceProfile() (v CustomerReviewV1DataCustomerDetailsComplianceProfileBusinessComplianceProfile) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u CustomerReviewV1DataCustomerDetailsComplianceProfileUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *CustomerReviewV1DataCustomerDetailsComplianceProfileUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PII required to trigger Patriot Act compliant KYC verification.
type CustomerReviewV1DataCustomerDetailsComplianceProfileIndividualComplianceProfile struct {
	// Masked date of birth in \***\*-**-\*\* format.
	Dob time.Time `json:"dob" api:"required" format:"date"`
	// Masked Social Security Number in the format **\*-**-\*\*\*\*.
	Ssn string `json:"ssn" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Dob         respjson.Field
		Ssn         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CustomerReviewV1DataCustomerDetailsComplianceProfileIndividualComplianceProfile) RawJSON() string {
	return r.JSON.raw
}
func (r *CustomerReviewV1DataCustomerDetailsComplianceProfileIndividualComplianceProfile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Business registration data required to trigger Patriot Act compliant KYB
// verification.
type CustomerReviewV1DataCustomerDetailsComplianceProfileBusinessComplianceProfile struct {
	// Masked Employer Identification Number in the format **-**\*****
	Ein string `json:"ein" api:"required"`
	// The official registered name of the business. This name should be correlated
	// with the `ein` value.
	LegalBusinessName string `json:"legal_business_name" api:"required"`
	// A list of people related to the company. Only valid where customer type is
	// 'business'.
	Representatives []CustomerReviewV1DataCustomerDetailsComplianceProfileBusinessComplianceProfileRepresentative `json:"representatives" api:"nullable"`
	// Official business website URL. Optional but recommended for enhanced KYB.
	Website string `json:"website" api:"nullable" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Ein               respjson.Field
		LegalBusinessName respjson.Field
		Representatives   respjson.Field
		Website           respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CustomerReviewV1DataCustomerDetailsComplianceProfileBusinessComplianceProfile) RawJSON() string {
	return r.JSON.raw
}
func (r *CustomerReviewV1DataCustomerDetailsComplianceProfileBusinessComplianceProfile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerReviewV1DataCustomerDetailsComplianceProfileBusinessComplianceProfileRepresentative struct {
	Name  string `json:"name" api:"required"`
	Email string `json:"email" api:"nullable"`
	Phone string `json:"phone" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Email       respjson.Field
		Phone       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CustomerReviewV1DataCustomerDetailsComplianceProfileBusinessComplianceProfileRepresentative) RawJSON() string {
	return r.JSON.raw
}
func (r *CustomerReviewV1DataCustomerDetailsComplianceProfileBusinessComplianceProfileRepresentative) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerReviewV1DataCustomerDetailsConfig struct {
	// Any of "inline", "background", "skip".
	ProcessingMethod string `json:"processing_method"`
	// Any of "standard", "verified", "rejected", "review".
	SandboxOutcome string `json:"sandbox_outcome"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ProcessingMethod respjson.Field
		SandboxOutcome   respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CustomerReviewV1DataCustomerDetailsConfig) RawJSON() string { return r.JSON.raw }
func (r *CustomerReviewV1DataCustomerDetailsConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerReviewV1DataCustomerDetailsDevice struct {
	// The customer's IP address at the time of profile creation. Use `0.0.0.0` to
	// represent an offline customer registration.
	IPAddress string `json:"ip_address" api:"required" format:"ipv4"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IPAddress   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CustomerReviewV1DataCustomerDetailsDevice) RawJSON() string { return r.JSON.raw }
func (r *CustomerReviewV1DataCustomerDetailsDevice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerReviewV1DataIdentityDetails struct {
	// Detailed breakdown of the customer verification results, including decisions,
	// risk scores, correlation score, and more.
	Breakdown CustomerReviewV1DataIdentityDetailsBreakdown `json:"breakdown" api:"required"`
	// Timestamp of when the review was initiated.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Any of "accept", "reject", "review".
	Decision string `json:"decision" api:"required"`
	// Unique identifier for the review.
	ReviewID string `json:"review_id" api:"required" format:"uuid"`
	// Timestamp of the most recent update to the review.
	UpdatedAt time.Time                              `json:"updated_at" api:"required" format:"date-time"`
	KYC       CustomerReviewV1DataIdentityDetailsKYC `json:"kyc"`
	// Dictionary of all messages from the customer verification process.
	Messages      map[string]string                                `json:"messages" api:"nullable"`
	NetworkAlerts CustomerReviewV1DataIdentityDetailsNetworkAlerts `json:"network_alerts"`
	Reputation    CustomerReviewV1DataIdentityDetailsReputation    `json:"reputation"`
	WatchList     CustomerReviewV1DataIdentityDetailsWatchList     `json:"watch_list"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Breakdown     respjson.Field
		CreatedAt     respjson.Field
		Decision      respjson.Field
		ReviewID      respjson.Field
		UpdatedAt     respjson.Field
		KYC           respjson.Field
		Messages      respjson.Field
		NetworkAlerts respjson.Field
		Reputation    respjson.Field
		WatchList     respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CustomerReviewV1DataIdentityDetails) RawJSON() string { return r.JSON.raw }
func (r *CustomerReviewV1DataIdentityDetails) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Detailed breakdown of the customer verification results, including decisions,
// risk scores, correlation score, and more.
type CustomerReviewV1DataIdentityDetailsBreakdown struct {
	Address                IdentityVerificationBreakdownV1 `json:"address"`
	BusinessEvaluation     IdentityVerificationBreakdownV1 `json:"business_evaluation"`
	BusinessIdentification IdentityVerificationBreakdownV1 `json:"business_identification"`
	BusinessValidation     IdentityVerificationBreakdownV1 `json:"business_validation"`
	Email                  IdentityVerificationBreakdownV1 `json:"email"`
	Fraud                  IdentityVerificationBreakdownV1 `json:"fraud"`
	Phone                  IdentityVerificationBreakdownV1 `json:"phone"`
	Synthetic              IdentityVerificationBreakdownV1 `json:"synthetic"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Address                respjson.Field
		BusinessEvaluation     respjson.Field
		BusinessIdentification respjson.Field
		BusinessValidation     respjson.Field
		Email                  respjson.Field
		Fraud                  respjson.Field
		Phone                  respjson.Field
		Synthetic              respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CustomerReviewV1DataIdentityDetailsBreakdown) RawJSON() string { return r.JSON.raw }
func (r *CustomerReviewV1DataIdentityDetailsBreakdown) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerReviewV1DataIdentityDetailsKYC struct {
	// Boolean values indicating the result of each validation in the KYC process.
	Validations CustomerReviewV1DataIdentityDetailsKYCValidations `json:"validations" api:"required"`
	// List of specific result codes from the KYC screening process.
	Codes []string `json:"codes" api:"nullable"`
	// Any of "accept", "reject", "review".
	Decision string `json:"decision"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Validations respjson.Field
		Codes       respjson.Field
		Decision    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CustomerReviewV1DataIdentityDetailsKYC) RawJSON() string { return r.JSON.raw }
func (r *CustomerReviewV1DataIdentityDetailsKYC) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Boolean values indicating the result of each validation in the KYC process.
type CustomerReviewV1DataIdentityDetailsKYCValidations struct {
	Address   bool `json:"address"`
	City      bool `json:"city"`
	Dob       bool `json:"dob"`
	Email     bool `json:"email"`
	FirstName bool `json:"first_name"`
	LastName  bool `json:"last_name"`
	Phone     bool `json:"phone"`
	Ssn       bool `json:"ssn"`
	State     bool `json:"state"`
	Zip       bool `json:"zip"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Address     respjson.Field
		City        respjson.Field
		Dob         respjson.Field
		Email       respjson.Field
		FirstName   respjson.Field
		LastName    respjson.Field
		Phone       respjson.Field
		Ssn         respjson.Field
		State       respjson.Field
		Zip         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CustomerReviewV1DataIdentityDetailsKYCValidations) RawJSON() string { return r.JSON.raw }
func (r *CustomerReviewV1DataIdentityDetailsKYCValidations) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerReviewV1DataIdentityDetailsNetworkAlerts struct {
	// Any alerts or flags raised during the consortium alert screening.
	Alerts []string `json:"alerts" api:"nullable"`
	// List of specific result codes from the consortium alert screening.
	Codes []string `json:"codes" api:"nullable"`
	// Any of "accept", "reject", "review".
	Decision string `json:"decision"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Alerts      respjson.Field
		Codes       respjson.Field
		Decision    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CustomerReviewV1DataIdentityDetailsNetworkAlerts) RawJSON() string { return r.JSON.raw }
func (r *CustomerReviewV1DataIdentityDetailsNetworkAlerts) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerReviewV1DataIdentityDetailsReputation struct {
	// Specific codes related to the Straddle reputation screening results.
	Codes []string `json:"codes" api:"nullable"`
	// Any of "accept", "reject", "review".
	Decision  string                                                `json:"decision"`
	Insights  CustomerReviewV1DataIdentityDetailsReputationInsights `json:"insights"`
	RiskScore float64                                               `json:"risk_score" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Codes       respjson.Field
		Decision    respjson.Field
		Insights    respjson.Field
		RiskScore   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CustomerReviewV1DataIdentityDetailsReputation) RawJSON() string { return r.JSON.raw }
func (r *CustomerReviewV1DataIdentityDetailsReputation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerReviewV1DataIdentityDetailsReputationInsights struct {
	AccountsActiveCount                 int64       `json:"accounts_active_count" api:"nullable"`
	AccountsClosedCount                 int64       `json:"accounts_closed_count" api:"nullable"`
	AccountsClosedDates                 []time.Time `json:"accounts_closed_dates" api:"nullable" format:"date"`
	AccountsCount                       int64       `json:"accounts_count" api:"nullable"`
	AccountsFraudCount                  int64       `json:"accounts_fraud_count" api:"nullable"`
	AccountsFraudLabeledDates           []time.Time `json:"accounts_fraud_labeled_dates" api:"nullable" format:"date"`
	AccountsFraudLossTotalAmount        float64     `json:"accounts_fraud_loss_total_amount" api:"nullable"`
	ACHFraudTransactionsCount           int64       `json:"ach_fraud_transactions_count" api:"nullable"`
	ACHFraudTransactionsDates           []time.Time `json:"ach_fraud_transactions_dates" api:"nullable" format:"date"`
	ACHFraudTransactionsTotalAmount     float64     `json:"ach_fraud_transactions_total_amount" api:"nullable"`
	ACHReturnedTransactionsCount        int64       `json:"ach_returned_transactions_count" api:"nullable"`
	ACHReturnedTransactionsDates        []time.Time `json:"ach_returned_transactions_dates" api:"nullable" format:"date"`
	ACHReturnedTransactionsTotalAmount  float64     `json:"ach_returned_transactions_total_amount" api:"nullable"`
	ApplicationsApprovedCount           int64       `json:"applications_approved_count" api:"nullable"`
	ApplicationsCount                   int64       `json:"applications_count" api:"nullable"`
	ApplicationsDates                   []time.Time `json:"applications_dates" api:"nullable" format:"date"`
	ApplicationsDeclinedCount           int64       `json:"applications_declined_count" api:"nullable"`
	ApplicationsFraudCount              int64       `json:"applications_fraud_count" api:"nullable"`
	CardDisputedTransactionsCount       int64       `json:"card_disputed_transactions_count" api:"nullable"`
	CardDisputedTransactionsDates       []time.Time `json:"card_disputed_transactions_dates" api:"nullable" format:"date"`
	CardDisputedTransactionsTotalAmount float64     `json:"card_disputed_transactions_total_amount" api:"nullable"`
	CardFraudTransactionsCount          int64       `json:"card_fraud_transactions_count" api:"nullable"`
	CardFraudTransactionsDates          []time.Time `json:"card_fraud_transactions_dates" api:"nullable" format:"date"`
	CardFraudTransactionsTotalAmount    float64     `json:"card_fraud_transactions_total_amount" api:"nullable"`
	CardStoppedTransactionsCount        int64       `json:"card_stopped_transactions_count" api:"nullable"`
	CardStoppedTransactionsDates        []time.Time `json:"card_stopped_transactions_dates" api:"nullable" format:"date"`
	UserActiveProfileCount              int64       `json:"user_active_profile_count" api:"nullable"`
	UserAddressCount                    int64       `json:"user_address_count" api:"nullable"`
	UserClosedProfileCount              int64       `json:"user_closed_profile_count" api:"nullable"`
	UserDobCount                        int64       `json:"user_dob_count" api:"nullable"`
	UserEmailCount                      int64       `json:"user_email_count" api:"nullable"`
	UserInstitutionCount                int64       `json:"user_institution_count" api:"nullable"`
	UserMobileCount                     int64       `json:"user_mobile_count" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccountsActiveCount                 respjson.Field
		AccountsClosedCount                 respjson.Field
		AccountsClosedDates                 respjson.Field
		AccountsCount                       respjson.Field
		AccountsFraudCount                  respjson.Field
		AccountsFraudLabeledDates           respjson.Field
		AccountsFraudLossTotalAmount        respjson.Field
		ACHFraudTransactionsCount           respjson.Field
		ACHFraudTransactionsDates           respjson.Field
		ACHFraudTransactionsTotalAmount     respjson.Field
		ACHReturnedTransactionsCount        respjson.Field
		ACHReturnedTransactionsDates        respjson.Field
		ACHReturnedTransactionsTotalAmount  respjson.Field
		ApplicationsApprovedCount           respjson.Field
		ApplicationsCount                   respjson.Field
		ApplicationsDates                   respjson.Field
		ApplicationsDeclinedCount           respjson.Field
		ApplicationsFraudCount              respjson.Field
		CardDisputedTransactionsCount       respjson.Field
		CardDisputedTransactionsDates       respjson.Field
		CardDisputedTransactionsTotalAmount respjson.Field
		CardFraudTransactionsCount          respjson.Field
		CardFraudTransactionsDates          respjson.Field
		CardFraudTransactionsTotalAmount    respjson.Field
		CardStoppedTransactionsCount        respjson.Field
		CardStoppedTransactionsDates        respjson.Field
		UserActiveProfileCount              respjson.Field
		UserAddressCount                    respjson.Field
		UserClosedProfileCount              respjson.Field
		UserDobCount                        respjson.Field
		UserEmailCount                      respjson.Field
		UserInstitutionCount                respjson.Field
		UserMobileCount                     respjson.Field
		ExtraFields                         map[string]respjson.Field
		raw                                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CustomerReviewV1DataIdentityDetailsReputationInsights) RawJSON() string { return r.JSON.raw }
func (r *CustomerReviewV1DataIdentityDetailsReputationInsights) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerReviewV1DataIdentityDetailsWatchList struct {
	// Specific codes related to the Straddle watchlist screening results.
	Codes []string `json:"codes" api:"nullable"`
	// Any of "accept", "reject", "review".
	Decision string `json:"decision"`
	// Information about any matches found during screening.
	Matched []string `json:"matched" api:"nullable"`
	// Information about any matches found during screening.
	Matches []CustomerReviewV1DataIdentityDetailsWatchListMatch `json:"matches" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Codes       respjson.Field
		Decision    respjson.Field
		Matched     respjson.Field
		Matches     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CustomerReviewV1DataIdentityDetailsWatchList) RawJSON() string { return r.JSON.raw }
func (r *CustomerReviewV1DataIdentityDetailsWatchList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerReviewV1DataIdentityDetailsWatchListMatch struct {
	// Any of "low_confidence", "potential_match", "likely_match", "high_confidence".
	Correlation string `json:"correlation" api:"required"`
	// The name of the list the match was found.
	ListName string `json:"list_name" api:"required"`
	// Data fields that matched.
	MatchFields []string `json:"match_fields" api:"required"`
	// Relevent Urls to review.
	URLs []string `json:"urls" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Correlation respjson.Field
		ListName    respjson.Field
		MatchFields respjson.Field
		URLs        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CustomerReviewV1DataIdentityDetailsWatchListMatch) RawJSON() string { return r.JSON.raw }
func (r *CustomerReviewV1DataIdentityDetailsWatchListMatch) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indicates the structure of the returned content.
//
//   - "object" means the `data` field contains a single JSON object.
//   - "array" means the `data` field contains an array of objects.
//   - "error" means the `data` field contains an error object with details of the
//     issue.
//   - "none" means no data is returned.
type CustomerReviewV1ResponseType string

const (
	CustomerReviewV1ResponseTypeObject CustomerReviewV1ResponseType = "object"
	CustomerReviewV1ResponseTypeArray  CustomerReviewV1ResponseType = "array"
	CustomerReviewV1ResponseTypeError  CustomerReviewV1ResponseType = "error"
	CustomerReviewV1ResponseTypeNone   CustomerReviewV1ResponseType = "none"
)

type IdentityVerificationBreakdownV1 struct {
	// List of specific result codes from the fraud and risk screening.
	Codes []string `json:"codes" api:"nullable"`
	// Any of "low_confidence", "potential_match", "likely_match", "high_confidence".
	Correlation IdentityVerificationBreakdownV1Correlation `json:"correlation"`
	// Represents the strength of the correlation between provided and known
	// information. A higher score indicates a stronger correlation.
	CorrelationScore float64 `json:"correlation_score" api:"nullable"`
	// Any of "accept", "reject", "review".
	Decision IdentityVerificationBreakdownV1Decision `json:"decision"`
	// Predicts the inherent risk associated with the customer for a given module. A
	// higher score indicates a greater likelihood of fraud.
	RiskScore float64 `json:"risk_score" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Codes            respjson.Field
		Correlation      respjson.Field
		CorrelationScore respjson.Field
		Decision         respjson.Field
		RiskScore        respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IdentityVerificationBreakdownV1) RawJSON() string { return r.JSON.raw }
func (r *IdentityVerificationBreakdownV1) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IdentityVerificationBreakdownV1Correlation string

const (
	IdentityVerificationBreakdownV1CorrelationLowConfidence  IdentityVerificationBreakdownV1Correlation = "low_confidence"
	IdentityVerificationBreakdownV1CorrelationPotentialMatch IdentityVerificationBreakdownV1Correlation = "potential_match"
	IdentityVerificationBreakdownV1CorrelationLikelyMatch    IdentityVerificationBreakdownV1Correlation = "likely_match"
	IdentityVerificationBreakdownV1CorrelationHighConfidence IdentityVerificationBreakdownV1Correlation = "high_confidence"
)

type IdentityVerificationBreakdownV1Decision string

const (
	IdentityVerificationBreakdownV1DecisionAccept IdentityVerificationBreakdownV1Decision = "accept"
	IdentityVerificationBreakdownV1DecisionReject IdentityVerificationBreakdownV1Decision = "reject"
	IdentityVerificationBreakdownV1DecisionReview IdentityVerificationBreakdownV1Decision = "review"
)

type CustomerReviewDecisionParams struct {
	// The final status of the customer review.
	//
	// Any of "verified", "rejected".
	Status            CustomerReviewDecisionParamsStatus `json:"status,omitzero" api:"required"`
	CorrelationID     param.Opt[string]                  `header:"Correlation-Id,omitzero" json:"-"`
	IdempotencyKey    param.Opt[string]                  `header:"Idempotency-Key,omitzero" json:"-"`
	RequestID         param.Opt[string]                  `header:"Request-Id,omitzero" json:"-"`
	StraddleAccountID param.Opt[string]                  `header:"Straddle-Account-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

func (r CustomerReviewDecisionParams) MarshalJSON() (data []byte, err error) {
	type shadow CustomerReviewDecisionParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CustomerReviewDecisionParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The final status of the customer review.
type CustomerReviewDecisionParamsStatus string

const (
	CustomerReviewDecisionParamsStatusVerified CustomerReviewDecisionParamsStatus = "verified"
	CustomerReviewDecisionParamsStatusRejected CustomerReviewDecisionParamsStatus = "rejected"
)

type CustomerReviewGetParams struct {
	CorrelationID     param.Opt[string] `header:"Correlation-Id,omitzero" json:"-"`
	RequestID         param.Opt[string] `header:"Request-Id,omitzero" json:"-"`
	StraddleAccountID param.Opt[string] `header:"Straddle-Account-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

type CustomerReviewRefreshReviewParams struct {
	CorrelationID     param.Opt[string] `header:"Correlation-Id,omitzero" json:"-"`
	IdempotencyKey    param.Opt[string] `header:"Idempotency-Key,omitzero" json:"-"`
	RequestID         param.Opt[string] `header:"Request-Id,omitzero" json:"-"`
	StraddleAccountID param.Opt[string] `header:"Straddle-Account-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}
