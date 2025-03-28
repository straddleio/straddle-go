// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package straddle

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/stainless-sdks/straddle-go/internal/apijson"
	"github.com/stainless-sdks/straddle-go/internal/param"
	"github.com/stainless-sdks/straddle-go/internal/requestconfig"
	"github.com/stainless-sdks/straddle-go/option"
	"github.com/stainless-sdks/straddle-go/shared"
	"github.com/tidwall/gjson"
)

// CustomerReviewService contains methods and other services that help with
// interacting with the straddle API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCustomerReviewService] method instead.
type CustomerReviewService struct {
	Options []option.RequestOption
}

// NewCustomerReviewService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCustomerReviewService(opts ...option.RequestOption) (r *CustomerReviewService) {
	r = &CustomerReviewService{}
	r.Options = opts
	return
}

// Updates the status of a customer's identity decision. This endpoint allows you
// to modify the outcome of a customer risk screening and is useful for correcting
// or updating the status of a customer's verification. Note that this endpoint is
// only available for customers with a current status of `review`.
func (r *CustomerReviewService) Decision(ctx context.Context, id string, params CustomerReviewDecisionParams, opts ...option.RequestOption) (res *CustomerV1, err error) {
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
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/customers/%s/review", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
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
	path := fmt.Sprintf("v1/customers/%s/review", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type CustomerReviewV1 struct {
	Data CustomerReviewV1Data `json:"data,required"`
	// Metadata about the API request, including an identifier and timestamp.
	Meta shared.ResponseMetadata `json:"meta,required"`
	// Indicates the structure of the returned content.
	//
	//   - "object" means the `data` field contains a single JSON object.
	//   - "array" means the `data` field contains an array of objects.
	//   - "error" means the `data` field contains an error object with details of the
	//     issue.
	//   - "none" means no data is returned.
	ResponseType CustomerReviewV1ResponseType `json:"response_type,required"`
	JSON         customerReviewV1JSON         `json:"-"`
}

// customerReviewV1JSON contains the JSON metadata for the struct
// [CustomerReviewV1]
type customerReviewV1JSON struct {
	Data         apijson.Field
	Meta         apijson.Field
	ResponseType apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *CustomerReviewV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerReviewV1JSON) RawJSON() string {
	return r.raw
}

type CustomerReviewV1Data struct {
	CustomerDetails CustomerReviewV1DataCustomerDetails `json:"customer_details,required"`
	IdentityDetails CustomerReviewV1DataIdentityDetails `json:"identity_details"`
	JSON            customerReviewV1DataJSON            `json:"-"`
}

// customerReviewV1DataJSON contains the JSON metadata for the struct
// [CustomerReviewV1Data]
type customerReviewV1DataJSON struct {
	CustomerDetails apijson.Field
	IdentityDetails apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *CustomerReviewV1Data) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerReviewV1DataJSON) RawJSON() string {
	return r.raw
}

type CustomerReviewV1DataCustomerDetails struct {
	// Unique identifier for the customer.
	ID string `json:"id,required" format:"uuid"`
	// Timestamp of when the customer record was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The customer's email address.
	Email string `json:"email,required" format:"email"`
	// Full name of the individual or business name.
	Name string `json:"name,required"`
	// The customer's phone number in E.164 format.
	Phone  string                                    `json:"phone,required"`
	Status CustomerReviewV1DataCustomerDetailsStatus `json:"status,required"`
	Type   CustomerReviewV1DataCustomerDetailsType   `json:"type,required"`
	// Timestamp of the most recent update to the customer record.
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// An object containing the customer's address. This is optional, but if provided,
	// all required fields must be present.
	Address CustomerAddressV1 `json:"address,nullable"`
	// PII required to trigger Patriot Act compliant KYC verification.
	ComplianceProfile CustomerReviewV1DataCustomerDetailsComplianceProfile `json:"compliance_profile,nullable"`
	Device            CustomerReviewV1DataCustomerDetailsDevice            `json:"device"`
	// Unique identifier for the customer in your database, used for cross-referencing
	// between Straddle and your systems.
	ExternalID string `json:"external_id,nullable"`
	// Up to 20 additional user-defined key-value pairs. Useful for storing additional
	// information about the customer in a structured format.
	Metadata map[string]string                       `json:"metadata,nullable"`
	JSON     customerReviewV1DataCustomerDetailsJSON `json:"-"`
}

// customerReviewV1DataCustomerDetailsJSON contains the JSON metadata for the
// struct [CustomerReviewV1DataCustomerDetails]
type customerReviewV1DataCustomerDetailsJSON struct {
	ID                apijson.Field
	CreatedAt         apijson.Field
	Email             apijson.Field
	Name              apijson.Field
	Phone             apijson.Field
	Status            apijson.Field
	Type              apijson.Field
	UpdatedAt         apijson.Field
	Address           apijson.Field
	ComplianceProfile apijson.Field
	Device            apijson.Field
	ExternalID        apijson.Field
	Metadata          apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *CustomerReviewV1DataCustomerDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerReviewV1DataCustomerDetailsJSON) RawJSON() string {
	return r.raw
}

type CustomerReviewV1DataCustomerDetailsStatus string

const (
	CustomerReviewV1DataCustomerDetailsStatusPending  CustomerReviewV1DataCustomerDetailsStatus = "pending"
	CustomerReviewV1DataCustomerDetailsStatusReview   CustomerReviewV1DataCustomerDetailsStatus = "review"
	CustomerReviewV1DataCustomerDetailsStatusVerified CustomerReviewV1DataCustomerDetailsStatus = "verified"
	CustomerReviewV1DataCustomerDetailsStatusInactive CustomerReviewV1DataCustomerDetailsStatus = "inactive"
	CustomerReviewV1DataCustomerDetailsStatusRejected CustomerReviewV1DataCustomerDetailsStatus = "rejected"
)

func (r CustomerReviewV1DataCustomerDetailsStatus) IsKnown() bool {
	switch r {
	case CustomerReviewV1DataCustomerDetailsStatusPending, CustomerReviewV1DataCustomerDetailsStatusReview, CustomerReviewV1DataCustomerDetailsStatusVerified, CustomerReviewV1DataCustomerDetailsStatusInactive, CustomerReviewV1DataCustomerDetailsStatusRejected:
		return true
	}
	return false
}

type CustomerReviewV1DataCustomerDetailsType string

const (
	CustomerReviewV1DataCustomerDetailsTypeIndividual CustomerReviewV1DataCustomerDetailsType = "individual"
	CustomerReviewV1DataCustomerDetailsTypeBusiness   CustomerReviewV1DataCustomerDetailsType = "business"
)

func (r CustomerReviewV1DataCustomerDetailsType) IsKnown() bool {
	switch r {
	case CustomerReviewV1DataCustomerDetailsTypeIndividual, CustomerReviewV1DataCustomerDetailsTypeBusiness:
		return true
	}
	return false
}

// PII required to trigger Patriot Act compliant KYC verification.
type CustomerReviewV1DataCustomerDetailsComplianceProfile struct {
	// This field can have the runtime type of [time.Time], [string].
	Dob interface{} `json:"dob"`
	// Full 9-digit Employer Identification Number for businesses. This data is
	// required to trigger Patriot Act compliant Know Your Business (KYB) verification.
	// Only valid where customer type is 'business'.
	Ein string `json:"ein,nullable" format:"**-*******"`
	// The official name of the business. This name should be correlated with the ein
	// value. Only valid where customer type is 'business'.
	LegalBusinessName string `json:"legal_business_name,nullable"`
	// This field can have the runtime type of
	// [[]CustomerReviewV1DataCustomerDetailsComplianceProfileIndividualComplianceProfileRepresentative],
	// [[]CustomerReviewV1DataCustomerDetailsComplianceProfileBusinessComplianceProfileRepresentative].
	Representatives interface{} `json:"representatives"`
	// Masked Social Security Number in the format **\*-**-\*\*\*\*.
	Ssn string `json:"ssn,nullable"`
	// URL of the company's official website. Only valid where customer type is
	// 'business'.
	Website string                                                   `json:"website,nullable"`
	JSON    customerReviewV1DataCustomerDetailsComplianceProfileJSON `json:"-"`
	union   CustomerReviewV1DataCustomerDetailsComplianceProfileUnion
}

// customerReviewV1DataCustomerDetailsComplianceProfileJSON contains the JSON
// metadata for the struct [CustomerReviewV1DataCustomerDetailsComplianceProfile]
type customerReviewV1DataCustomerDetailsComplianceProfileJSON struct {
	Dob               apijson.Field
	Ein               apijson.Field
	LegalBusinessName apijson.Field
	Representatives   apijson.Field
	Ssn               apijson.Field
	Website           apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r customerReviewV1DataCustomerDetailsComplianceProfileJSON) RawJSON() string {
	return r.raw
}

func (r *CustomerReviewV1DataCustomerDetailsComplianceProfile) UnmarshalJSON(data []byte) (err error) {
	*r = CustomerReviewV1DataCustomerDetailsComplianceProfile{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [CustomerReviewV1DataCustomerDetailsComplianceProfileUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [CustomerReviewV1DataCustomerDetailsComplianceProfileIndividualComplianceProfile],
// [CustomerReviewV1DataCustomerDetailsComplianceProfileBusinessComplianceProfile].
func (r CustomerReviewV1DataCustomerDetailsComplianceProfile) AsUnion() CustomerReviewV1DataCustomerDetailsComplianceProfileUnion {
	return r.union
}

// PII required to trigger Patriot Act compliant KYC verification.
//
// Union satisfied by
// [CustomerReviewV1DataCustomerDetailsComplianceProfileIndividualComplianceProfile]
// or
// [CustomerReviewV1DataCustomerDetailsComplianceProfileBusinessComplianceProfile].
type CustomerReviewV1DataCustomerDetailsComplianceProfileUnion interface {
	implementsCustomerReviewV1DataCustomerDetailsComplianceProfile()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CustomerReviewV1DataCustomerDetailsComplianceProfileUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CustomerReviewV1DataCustomerDetailsComplianceProfileIndividualComplianceProfile{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CustomerReviewV1DataCustomerDetailsComplianceProfileBusinessComplianceProfile{}),
		},
	)
}

// PII required to trigger Patriot Act compliant KYC verification.
type CustomerReviewV1DataCustomerDetailsComplianceProfileIndividualComplianceProfile struct {
	// Masked date of birth in \***\*-**-\*\* format.
	Dob time.Time `json:"dob,required,nullable" format:"date"`
	// Masked Social Security Number in the format **\*-**-\*\*\*\*.
	Ssn string `json:"ssn,required,nullable"`
	// Full 9-digit Employer Identification Number for businesses. This data is
	// required to trigger Patriot Act compliant Know Your Business (KYB) verification.
	// Only valid where customer type is 'business'.
	Ein string `json:"ein,nullable" format:"**-*******"`
	// The official name of the business. This name should be correlated with the ein
	// value. Only valid where customer type is 'business'.
	LegalBusinessName string `json:"legal_business_name,nullable"`
	// A list of people related to the company. Only valid where customer type is
	// 'business'.
	Representatives []CustomerReviewV1DataCustomerDetailsComplianceProfileIndividualComplianceProfileRepresentative `json:"representatives,nullable"`
	// URL of the company's official website. Only valid where customer type is
	// 'business'.
	Website string                                                                              `json:"website,nullable"`
	JSON    customerReviewV1DataCustomerDetailsComplianceProfileIndividualComplianceProfileJSON `json:"-"`
}

// customerReviewV1DataCustomerDetailsComplianceProfileIndividualComplianceProfileJSON
// contains the JSON metadata for the struct
// [CustomerReviewV1DataCustomerDetailsComplianceProfileIndividualComplianceProfile]
type customerReviewV1DataCustomerDetailsComplianceProfileIndividualComplianceProfileJSON struct {
	Dob               apijson.Field
	Ssn               apijson.Field
	Ein               apijson.Field
	LegalBusinessName apijson.Field
	Representatives   apijson.Field
	Website           apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *CustomerReviewV1DataCustomerDetailsComplianceProfileIndividualComplianceProfile) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerReviewV1DataCustomerDetailsComplianceProfileIndividualComplianceProfileJSON) RawJSON() string {
	return r.raw
}

func (r CustomerReviewV1DataCustomerDetailsComplianceProfileIndividualComplianceProfile) implementsCustomerReviewV1DataCustomerDetailsComplianceProfile() {
}

type CustomerReviewV1DataCustomerDetailsComplianceProfileIndividualComplianceProfileRepresentative struct {
	Name  string                                                                                            `json:"name,required"`
	Email string                                                                                            `json:"email,nullable"`
	Phone string                                                                                            `json:"phone,nullable"`
	JSON  customerReviewV1DataCustomerDetailsComplianceProfileIndividualComplianceProfileRepresentativeJSON `json:"-"`
}

// customerReviewV1DataCustomerDetailsComplianceProfileIndividualComplianceProfileRepresentativeJSON
// contains the JSON metadata for the struct
// [CustomerReviewV1DataCustomerDetailsComplianceProfileIndividualComplianceProfileRepresentative]
type customerReviewV1DataCustomerDetailsComplianceProfileIndividualComplianceProfileRepresentativeJSON struct {
	Name        apijson.Field
	Email       apijson.Field
	Phone       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerReviewV1DataCustomerDetailsComplianceProfileIndividualComplianceProfileRepresentative) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerReviewV1DataCustomerDetailsComplianceProfileIndividualComplianceProfileRepresentativeJSON) RawJSON() string {
	return r.raw
}

// Business registration data required to trigger Patriot Act compliant KYB
// verification.
type CustomerReviewV1DataCustomerDetailsComplianceProfileBusinessComplianceProfile struct {
	// Masked Employer Identification Number in the format **-**\*****
	Ein string `json:"ein,required,nullable"`
	// The official registered name of the business. This name should be correlated
	// with the `ein` value.
	LegalBusinessName string `json:"legal_business_name,required,nullable"`
	// Date of birth for individual customers in ISO 8601 format (YYYY-MM-DD). This
	// data is required to trigger Patriot Act compliant Know Your Customer (KYC)
	// verification. Required if SSN is provided. Only valid where customer type is
	// 'individual'.
	Dob string `json:"dob,nullable" format:"****-**-**"`
	// A list of people related to the company. Only valid where customer type is
	// 'business'.
	Representatives []CustomerReviewV1DataCustomerDetailsComplianceProfileBusinessComplianceProfileRepresentative `json:"representatives,nullable"`
	// Full 9-digit Social Security Number or government identifier for individuals.
	// This data is required to trigger Patriot Act compliant KYC verification.
	// Required if DOB is provided. Only valid where customer type is 'individual'.
	Ssn string `json:"ssn,nullable" format:"***-**-****"`
	// Official business website URL. Optional but recommended for enhanced KYB.
	Website string                                                                            `json:"website,nullable" format:"uri"`
	JSON    customerReviewV1DataCustomerDetailsComplianceProfileBusinessComplianceProfileJSON `json:"-"`
}

// customerReviewV1DataCustomerDetailsComplianceProfileBusinessComplianceProfileJSON
// contains the JSON metadata for the struct
// [CustomerReviewV1DataCustomerDetailsComplianceProfileBusinessComplianceProfile]
type customerReviewV1DataCustomerDetailsComplianceProfileBusinessComplianceProfileJSON struct {
	Ein               apijson.Field
	LegalBusinessName apijson.Field
	Dob               apijson.Field
	Representatives   apijson.Field
	Ssn               apijson.Field
	Website           apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *CustomerReviewV1DataCustomerDetailsComplianceProfileBusinessComplianceProfile) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerReviewV1DataCustomerDetailsComplianceProfileBusinessComplianceProfileJSON) RawJSON() string {
	return r.raw
}

func (r CustomerReviewV1DataCustomerDetailsComplianceProfileBusinessComplianceProfile) implementsCustomerReviewV1DataCustomerDetailsComplianceProfile() {
}

type CustomerReviewV1DataCustomerDetailsComplianceProfileBusinessComplianceProfileRepresentative struct {
	Name  string                                                                                          `json:"name,required"`
	Email string                                                                                          `json:"email,nullable"`
	Phone string                                                                                          `json:"phone,nullable"`
	JSON  customerReviewV1DataCustomerDetailsComplianceProfileBusinessComplianceProfileRepresentativeJSON `json:"-"`
}

// customerReviewV1DataCustomerDetailsComplianceProfileBusinessComplianceProfileRepresentativeJSON
// contains the JSON metadata for the struct
// [CustomerReviewV1DataCustomerDetailsComplianceProfileBusinessComplianceProfileRepresentative]
type customerReviewV1DataCustomerDetailsComplianceProfileBusinessComplianceProfileRepresentativeJSON struct {
	Name        apijson.Field
	Email       apijson.Field
	Phone       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerReviewV1DataCustomerDetailsComplianceProfileBusinessComplianceProfileRepresentative) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerReviewV1DataCustomerDetailsComplianceProfileBusinessComplianceProfileRepresentativeJSON) RawJSON() string {
	return r.raw
}

type CustomerReviewV1DataCustomerDetailsDevice struct {
	// The customer's IP address at the time of profile creation. Use `0.0.0.0` to
	// represent an offline customer registration.
	IPAddress string                                        `json:"ip_address,required" format:"ipv4"`
	JSON      customerReviewV1DataCustomerDetailsDeviceJSON `json:"-"`
}

// customerReviewV1DataCustomerDetailsDeviceJSON contains the JSON metadata for the
// struct [CustomerReviewV1DataCustomerDetailsDevice]
type customerReviewV1DataCustomerDetailsDeviceJSON struct {
	IPAddress   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerReviewV1DataCustomerDetailsDevice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerReviewV1DataCustomerDetailsDeviceJSON) RawJSON() string {
	return r.raw
}

type CustomerReviewV1DataIdentityDetails struct {
	// Detailed breakdown of the customer verification results, including decisions,
	// risk scores, correlation score, and more.
	Breakdown CustomerReviewV1DataIdentityDetailsBreakdown `json:"breakdown,required"`
	// Timestamp of when the review was initiated.
	CreatedAt time.Time                                   `json:"created_at,required" format:"date-time"`
	Decision  CustomerReviewV1DataIdentityDetailsDecision `json:"decision,required"`
	// Unique identifier for the review.
	ReviewID string `json:"review_id,required" format:"uuid"`
	// Timestamp of the most recent update to the review.
	UpdatedAt time.Time                              `json:"updated_at,required" format:"date-time"`
	KYC       CustomerReviewV1DataIdentityDetailsKYC `json:"kyc"`
	// Dictionary of all messages from the customer verification process.
	Messages      map[string]string                                `json:"messages,nullable"`
	NetworkAlerts CustomerReviewV1DataIdentityDetailsNetworkAlerts `json:"network_alerts"`
	WatchList     CustomerReviewV1DataIdentityDetailsWatchList     `json:"watch_list"`
	JSON          customerReviewV1DataIdentityDetailsJSON          `json:"-"`
}

// customerReviewV1DataIdentityDetailsJSON contains the JSON metadata for the
// struct [CustomerReviewV1DataIdentityDetails]
type customerReviewV1DataIdentityDetailsJSON struct {
	Breakdown     apijson.Field
	CreatedAt     apijson.Field
	Decision      apijson.Field
	ReviewID      apijson.Field
	UpdatedAt     apijson.Field
	KYC           apijson.Field
	Messages      apijson.Field
	NetworkAlerts apijson.Field
	WatchList     apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *CustomerReviewV1DataIdentityDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerReviewV1DataIdentityDetailsJSON) RawJSON() string {
	return r.raw
}

// Detailed breakdown of the customer verification results, including decisions,
// risk scores, correlation score, and more.
type CustomerReviewV1DataIdentityDetailsBreakdown struct {
	Address                IdentityVerificationBreakdownV1                  `json:"address"`
	BusinessEvaluation     IdentityVerificationBreakdownV1                  `json:"business_evaluation"`
	BusinessIdentification IdentityVerificationBreakdownV1                  `json:"business_identification"`
	BusinessValidation     IdentityVerificationBreakdownV1                  `json:"business_validation"`
	Email                  IdentityVerificationBreakdownV1                  `json:"email"`
	Fraud                  IdentityVerificationBreakdownV1                  `json:"fraud"`
	Phone                  IdentityVerificationBreakdownV1                  `json:"phone"`
	Synthetic              IdentityVerificationBreakdownV1                  `json:"synthetic"`
	JSON                   customerReviewV1DataIdentityDetailsBreakdownJSON `json:"-"`
}

// customerReviewV1DataIdentityDetailsBreakdownJSON contains the JSON metadata for
// the struct [CustomerReviewV1DataIdentityDetailsBreakdown]
type customerReviewV1DataIdentityDetailsBreakdownJSON struct {
	Address                apijson.Field
	BusinessEvaluation     apijson.Field
	BusinessIdentification apijson.Field
	BusinessValidation     apijson.Field
	Email                  apijson.Field
	Fraud                  apijson.Field
	Phone                  apijson.Field
	Synthetic              apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *CustomerReviewV1DataIdentityDetailsBreakdown) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerReviewV1DataIdentityDetailsBreakdownJSON) RawJSON() string {
	return r.raw
}

type CustomerReviewV1DataIdentityDetailsDecision string

const (
	CustomerReviewV1DataIdentityDetailsDecisionAccept CustomerReviewV1DataIdentityDetailsDecision = "accept"
	CustomerReviewV1DataIdentityDetailsDecisionReject CustomerReviewV1DataIdentityDetailsDecision = "reject"
	CustomerReviewV1DataIdentityDetailsDecisionReview CustomerReviewV1DataIdentityDetailsDecision = "review"
)

func (r CustomerReviewV1DataIdentityDetailsDecision) IsKnown() bool {
	switch r {
	case CustomerReviewV1DataIdentityDetailsDecisionAccept, CustomerReviewV1DataIdentityDetailsDecisionReject, CustomerReviewV1DataIdentityDetailsDecisionReview:
		return true
	}
	return false
}

type CustomerReviewV1DataIdentityDetailsKYC struct {
	// Boolean values indicating the result of each validation in the KYC process.
	Validations CustomerReviewV1DataIdentityDetailsKYCValidations `json:"validations,required"`
	// List of specific result codes from the KYC screening process.
	Codes    []string                                       `json:"codes,nullable"`
	Decision CustomerReviewV1DataIdentityDetailsKYCDecision `json:"decision"`
	JSON     customerReviewV1DataIdentityDetailsKYCJSON     `json:"-"`
}

// customerReviewV1DataIdentityDetailsKYCJSON contains the JSON metadata for the
// struct [CustomerReviewV1DataIdentityDetailsKYC]
type customerReviewV1DataIdentityDetailsKYCJSON struct {
	Validations apijson.Field
	Codes       apijson.Field
	Decision    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerReviewV1DataIdentityDetailsKYC) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerReviewV1DataIdentityDetailsKYCJSON) RawJSON() string {
	return r.raw
}

// Boolean values indicating the result of each validation in the KYC process.
type CustomerReviewV1DataIdentityDetailsKYCValidations struct {
	Address   bool                                                  `json:"address"`
	City      bool                                                  `json:"city"`
	Dob       bool                                                  `json:"dob"`
	Email     bool                                                  `json:"email"`
	FirstName bool                                                  `json:"first_name"`
	LastName  bool                                                  `json:"last_name"`
	Phone     bool                                                  `json:"phone"`
	Ssn       bool                                                  `json:"ssn"`
	State     bool                                                  `json:"state"`
	Zip       bool                                                  `json:"zip"`
	JSON      customerReviewV1DataIdentityDetailsKYCValidationsJSON `json:"-"`
}

// customerReviewV1DataIdentityDetailsKYCValidationsJSON contains the JSON metadata
// for the struct [CustomerReviewV1DataIdentityDetailsKYCValidations]
type customerReviewV1DataIdentityDetailsKYCValidationsJSON struct {
	Address     apijson.Field
	City        apijson.Field
	Dob         apijson.Field
	Email       apijson.Field
	FirstName   apijson.Field
	LastName    apijson.Field
	Phone       apijson.Field
	Ssn         apijson.Field
	State       apijson.Field
	Zip         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerReviewV1DataIdentityDetailsKYCValidations) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerReviewV1DataIdentityDetailsKYCValidationsJSON) RawJSON() string {
	return r.raw
}

type CustomerReviewV1DataIdentityDetailsKYCDecision string

const (
	CustomerReviewV1DataIdentityDetailsKYCDecisionAccept CustomerReviewV1DataIdentityDetailsKYCDecision = "accept"
	CustomerReviewV1DataIdentityDetailsKYCDecisionReject CustomerReviewV1DataIdentityDetailsKYCDecision = "reject"
	CustomerReviewV1DataIdentityDetailsKYCDecisionReview CustomerReviewV1DataIdentityDetailsKYCDecision = "review"
)

func (r CustomerReviewV1DataIdentityDetailsKYCDecision) IsKnown() bool {
	switch r {
	case CustomerReviewV1DataIdentityDetailsKYCDecisionAccept, CustomerReviewV1DataIdentityDetailsKYCDecisionReject, CustomerReviewV1DataIdentityDetailsKYCDecisionReview:
		return true
	}
	return false
}

type CustomerReviewV1DataIdentityDetailsNetworkAlerts struct {
	// Any alerts or flags raised during the consortium alert screening.
	Alerts []string `json:"alerts,nullable"`
	// List of specific result codes from the consortium alert screening.
	Codes    []string                                                 `json:"codes,nullable"`
	Decision CustomerReviewV1DataIdentityDetailsNetworkAlertsDecision `json:"decision"`
	JSON     customerReviewV1DataIdentityDetailsNetworkAlertsJSON     `json:"-"`
}

// customerReviewV1DataIdentityDetailsNetworkAlertsJSON contains the JSON metadata
// for the struct [CustomerReviewV1DataIdentityDetailsNetworkAlerts]
type customerReviewV1DataIdentityDetailsNetworkAlertsJSON struct {
	Alerts      apijson.Field
	Codes       apijson.Field
	Decision    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerReviewV1DataIdentityDetailsNetworkAlerts) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerReviewV1DataIdentityDetailsNetworkAlertsJSON) RawJSON() string {
	return r.raw
}

type CustomerReviewV1DataIdentityDetailsNetworkAlertsDecision string

const (
	CustomerReviewV1DataIdentityDetailsNetworkAlertsDecisionAccept CustomerReviewV1DataIdentityDetailsNetworkAlertsDecision = "accept"
	CustomerReviewV1DataIdentityDetailsNetworkAlertsDecisionReject CustomerReviewV1DataIdentityDetailsNetworkAlertsDecision = "reject"
	CustomerReviewV1DataIdentityDetailsNetworkAlertsDecisionReview CustomerReviewV1DataIdentityDetailsNetworkAlertsDecision = "review"
)

func (r CustomerReviewV1DataIdentityDetailsNetworkAlertsDecision) IsKnown() bool {
	switch r {
	case CustomerReviewV1DataIdentityDetailsNetworkAlertsDecisionAccept, CustomerReviewV1DataIdentityDetailsNetworkAlertsDecisionReject, CustomerReviewV1DataIdentityDetailsNetworkAlertsDecisionReview:
		return true
	}
	return false
}

type CustomerReviewV1DataIdentityDetailsWatchList struct {
	// Specific codes related to the Straddle watchlist screening results.
	Codes    []string                                             `json:"codes,nullable"`
	Decision CustomerReviewV1DataIdentityDetailsWatchListDecision `json:"decision"`
	// Information about any matches found during screening.
	Matched []string `json:"matched,nullable"`
	// Information about any matches found during screening.
	Matches []CustomerReviewV1DataIdentityDetailsWatchListMatch `json:"matches,nullable"`
	JSON    customerReviewV1DataIdentityDetailsWatchListJSON    `json:"-"`
}

// customerReviewV1DataIdentityDetailsWatchListJSON contains the JSON metadata for
// the struct [CustomerReviewV1DataIdentityDetailsWatchList]
type customerReviewV1DataIdentityDetailsWatchListJSON struct {
	Codes       apijson.Field
	Decision    apijson.Field
	Matched     apijson.Field
	Matches     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerReviewV1DataIdentityDetailsWatchList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerReviewV1DataIdentityDetailsWatchListJSON) RawJSON() string {
	return r.raw
}

type CustomerReviewV1DataIdentityDetailsWatchListDecision string

const (
	CustomerReviewV1DataIdentityDetailsWatchListDecisionAccept CustomerReviewV1DataIdentityDetailsWatchListDecision = "accept"
	CustomerReviewV1DataIdentityDetailsWatchListDecisionReject CustomerReviewV1DataIdentityDetailsWatchListDecision = "reject"
	CustomerReviewV1DataIdentityDetailsWatchListDecisionReview CustomerReviewV1DataIdentityDetailsWatchListDecision = "review"
)

func (r CustomerReviewV1DataIdentityDetailsWatchListDecision) IsKnown() bool {
	switch r {
	case CustomerReviewV1DataIdentityDetailsWatchListDecisionAccept, CustomerReviewV1DataIdentityDetailsWatchListDecisionReject, CustomerReviewV1DataIdentityDetailsWatchListDecisionReview:
		return true
	}
	return false
}

type CustomerReviewV1DataIdentityDetailsWatchListMatch struct {
	Correlation CustomerReviewV1DataIdentityDetailsWatchListMatchesCorrelation `json:"correlation,required"`
	// The name of the list the match was found.
	ListName string `json:"list_name,required"`
	// Data fields that matched.
	MatchFields []string `json:"match_fields,required"`
	// Relevent Urls to review.
	URLs []string                                              `json:"urls,required"`
	JSON customerReviewV1DataIdentityDetailsWatchListMatchJSON `json:"-"`
}

// customerReviewV1DataIdentityDetailsWatchListMatchJSON contains the JSON metadata
// for the struct [CustomerReviewV1DataIdentityDetailsWatchListMatch]
type customerReviewV1DataIdentityDetailsWatchListMatchJSON struct {
	Correlation apijson.Field
	ListName    apijson.Field
	MatchFields apijson.Field
	URLs        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerReviewV1DataIdentityDetailsWatchListMatch) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerReviewV1DataIdentityDetailsWatchListMatchJSON) RawJSON() string {
	return r.raw
}

type CustomerReviewV1DataIdentityDetailsWatchListMatchesCorrelation string

const (
	CustomerReviewV1DataIdentityDetailsWatchListMatchesCorrelationLowConfidence  CustomerReviewV1DataIdentityDetailsWatchListMatchesCorrelation = "low_confidence"
	CustomerReviewV1DataIdentityDetailsWatchListMatchesCorrelationPotentialMatch CustomerReviewV1DataIdentityDetailsWatchListMatchesCorrelation = "potential_match"
	CustomerReviewV1DataIdentityDetailsWatchListMatchesCorrelationLikelyMatch    CustomerReviewV1DataIdentityDetailsWatchListMatchesCorrelation = "likely_match"
	CustomerReviewV1DataIdentityDetailsWatchListMatchesCorrelationHighConfidence CustomerReviewV1DataIdentityDetailsWatchListMatchesCorrelation = "high_confidence"
	CustomerReviewV1DataIdentityDetailsWatchListMatchesCorrelationUnknown        CustomerReviewV1DataIdentityDetailsWatchListMatchesCorrelation = "unknown"
)

func (r CustomerReviewV1DataIdentityDetailsWatchListMatchesCorrelation) IsKnown() bool {
	switch r {
	case CustomerReviewV1DataIdentityDetailsWatchListMatchesCorrelationLowConfidence, CustomerReviewV1DataIdentityDetailsWatchListMatchesCorrelationPotentialMatch, CustomerReviewV1DataIdentityDetailsWatchListMatchesCorrelationLikelyMatch, CustomerReviewV1DataIdentityDetailsWatchListMatchesCorrelationHighConfidence, CustomerReviewV1DataIdentityDetailsWatchListMatchesCorrelationUnknown:
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
type CustomerReviewV1ResponseType string

const (
	CustomerReviewV1ResponseTypeObject CustomerReviewV1ResponseType = "object"
	CustomerReviewV1ResponseTypeArray  CustomerReviewV1ResponseType = "array"
	CustomerReviewV1ResponseTypeError  CustomerReviewV1ResponseType = "error"
	CustomerReviewV1ResponseTypeNone   CustomerReviewV1ResponseType = "none"
)

func (r CustomerReviewV1ResponseType) IsKnown() bool {
	switch r {
	case CustomerReviewV1ResponseTypeObject, CustomerReviewV1ResponseTypeArray, CustomerReviewV1ResponseTypeError, CustomerReviewV1ResponseTypeNone:
		return true
	}
	return false
}

type IdentityVerificationBreakdownV1 struct {
	// List of specific result codes from the fraud and risk screening.
	Codes []string `json:"codes,nullable"`
	// Represents the strength of the correlation between provided and known
	// information. A higher score indicates a stronger correlation.
	CorrelationScore float64                                 `json:"correlation_score,nullable"`
	Decision         IdentityVerificationBreakdownV1Decision `json:"decision"`
	// Predicts the inherent risk associated with the customer for a given module. A
	// higher score indicates a greater likelihood of fraud.
	RiskScore float64                             `json:"risk_score,nullable"`
	JSON      identityVerificationBreakdownV1JSON `json:"-"`
}

// identityVerificationBreakdownV1JSON contains the JSON metadata for the struct
// [IdentityVerificationBreakdownV1]
type identityVerificationBreakdownV1JSON struct {
	Codes            apijson.Field
	CorrelationScore apijson.Field
	Decision         apijson.Field
	RiskScore        apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *IdentityVerificationBreakdownV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityVerificationBreakdownV1JSON) RawJSON() string {
	return r.raw
}

type IdentityVerificationBreakdownV1Decision string

const (
	IdentityVerificationBreakdownV1DecisionAccept IdentityVerificationBreakdownV1Decision = "accept"
	IdentityVerificationBreakdownV1DecisionReject IdentityVerificationBreakdownV1Decision = "reject"
	IdentityVerificationBreakdownV1DecisionReview IdentityVerificationBreakdownV1Decision = "review"
)

func (r IdentityVerificationBreakdownV1Decision) IsKnown() bool {
	switch r {
	case IdentityVerificationBreakdownV1DecisionAccept, IdentityVerificationBreakdownV1DecisionReject, IdentityVerificationBreakdownV1DecisionReview:
		return true
	}
	return false
}

type CustomerReviewDecisionParams struct {
	// The final status of the customer review.
	Status            param.Field[CustomerReviewDecisionParamsStatus] `json:"status,required"`
	CorrelationID     param.Field[string]                             `header:"Correlation-Id"`
	RequestID         param.Field[string]                             `header:"Request-Id"`
	StraddleAccountID param.Field[string]                             `header:"Straddle-Account-Id" format:"uuid"`
}

func (r CustomerReviewDecisionParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The final status of the customer review.
type CustomerReviewDecisionParamsStatus string

const (
	CustomerReviewDecisionParamsStatusVerified CustomerReviewDecisionParamsStatus = "verified"
	CustomerReviewDecisionParamsStatusRejected CustomerReviewDecisionParamsStatus = "rejected"
)

func (r CustomerReviewDecisionParamsStatus) IsKnown() bool {
	switch r {
	case CustomerReviewDecisionParamsStatusVerified, CustomerReviewDecisionParamsStatusRejected:
		return true
	}
	return false
}

type CustomerReviewGetParams struct {
	CorrelationID     param.Field[string] `header:"Correlation-Id"`
	RequestID         param.Field[string] `header:"Request-Id"`
	StraddleAccountID param.Field[string] `header:"Straddle-Account-Id" format:"uuid"`
}
