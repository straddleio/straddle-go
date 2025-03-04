// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package straddle

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/stainless-sdks/straddle-go/internal/apijson"
	"github.com/stainless-sdks/straddle-go/internal/param"
	"github.com/stainless-sdks/straddle-go/internal/requestconfig"
	"github.com/stainless-sdks/straddle-go/option"
	"github.com/stainless-sdks/straddle-go/shared"
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
func (r *CustomerReviewService) Decision(ctx context.Context, id string, params CustomerReviewDecisionParams, opts ...option.RequestOption) (res *shared.CustomerV1ItemResponse, err error) {
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
	ResponseType shared.ResponseTypeEnum `json:"response_type,required"`
	JSON         customerReviewV1JSON    `json:"-"`
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
	CustomerDetails shared.CustomerV1                   `json:"customer_details,required"`
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

type CustomerReviewV1DataIdentityDetails struct {
	// Detailed breakdown of the customer verification results, including decisions,
	// risk scores, correlation score, and more.
	Breakdown CustomerReviewV1DataIdentityDetailsBreakdown `json:"breakdown,required"`
	// Timestamp of when the review was initiated.
	CreatedAt time.Time                 `json:"created_at,required" format:"date-time"`
	Decision  shared.IdentityDecisionV1 `json:"decision,required"`
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
	Address   shared.IdentityVerificationBreakdownV1           `json:"address"`
	Email     shared.IdentityVerificationBreakdownV1           `json:"email"`
	Fraud     shared.IdentityVerificationBreakdownV1           `json:"fraud"`
	Phone     shared.IdentityVerificationBreakdownV1           `json:"phone"`
	Synthetic shared.IdentityVerificationBreakdownV1           `json:"synthetic"`
	JSON      customerReviewV1DataIdentityDetailsBreakdownJSON `json:"-"`
}

// customerReviewV1DataIdentityDetailsBreakdownJSON contains the JSON metadata for
// the struct [CustomerReviewV1DataIdentityDetailsBreakdown]
type customerReviewV1DataIdentityDetailsBreakdownJSON struct {
	Address     apijson.Field
	Email       apijson.Field
	Fraud       apijson.Field
	Phone       apijson.Field
	Synthetic   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerReviewV1DataIdentityDetailsBreakdown) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerReviewV1DataIdentityDetailsBreakdownJSON) RawJSON() string {
	return r.raw
}

type CustomerReviewV1DataIdentityDetailsKYC struct {
	// Boolean values indicating the result of each validation in the KYC process.
	Validations CustomerReviewV1DataIdentityDetailsKYCValidations `json:"validations,required"`
	// List of specific result codes from the KYC screening process.
	Codes    []string                                   `json:"codes,nullable"`
	Decision shared.IdentityDecisionV1                  `json:"decision"`
	JSON     customerReviewV1DataIdentityDetailsKYCJSON `json:"-"`
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

type CustomerReviewV1DataIdentityDetailsNetworkAlerts struct {
	// Any alerts or flags raised during the consortium alert screening.
	Alerts []string `json:"alerts,nullable"`
	// List of specific result codes from the consortium alert screening.
	Codes    []string                                             `json:"codes,nullable"`
	Decision shared.IdentityDecisionV1                            `json:"decision"`
	JSON     customerReviewV1DataIdentityDetailsNetworkAlertsJSON `json:"-"`
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

type CustomerReviewV1DataIdentityDetailsWatchList struct {
	// Specific codes related to the Straddle watchlist screening results.
	Codes    []string                  `json:"codes,nullable"`
	Decision shared.IdentityDecisionV1 `json:"decision"`
	// Information about any matches found during screening.
	Matched []string                                         `json:"matched,nullable"`
	JSON    customerReviewV1DataIdentityDetailsWatchListJSON `json:"-"`
}

// customerReviewV1DataIdentityDetailsWatchListJSON contains the JSON metadata for
// the struct [CustomerReviewV1DataIdentityDetailsWatchList]
type customerReviewV1DataIdentityDetailsWatchListJSON struct {
	Codes       apijson.Field
	Decision    apijson.Field
	Matched     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerReviewV1DataIdentityDetailsWatchList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerReviewV1DataIdentityDetailsWatchListJSON) RawJSON() string {
	return r.raw
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
