// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package straddle

import (
	"context"
	"net/http"
	"os"
	"slices"

	"github.com/straddleio/straddle-go/internal/requestconfig"
	"github.com/straddleio/straddle-go/option"
)

// Client creates a struct with services and top level methods that help with
// interacting with the straddle API. You should not instantiate this client
// directly, and instead use the [NewClient] method instead.
type Client struct {
	options []option.RequestOption
	Embed   EmbedService
	// Bridge provides a comprehensive suite of tools for connecting customer bank
	// accounts. Use it to generate secure widget sessions for instant account
	// verification, accept tokens from major providers like Plaid and Finicity, or
	// verify accounts directly via our API. Bridge handles all sensitive banking
	// credentials and ensures secure, compliant connections with support for 90% of US
	// bank accounts.
	Bridge BridgeService
	// Customers represent the end users who send or receive payments through your
	// integration. Each customer undergoes automatic identity verification and fraud
	// screening upon creation. Use customers to track payment history, manage bank
	// account connections, and maintain a secure record of all transactions associated
	// with a user. Customers can be either individuals or businesses with appropriate
	// compliance checks for each type.
	Customers CustomerService
	// Paykeys are secure tokens that link verified customer identities to their bank
	// accounts. Each Paykey includes built-in balance checking, fraud detection
	// through LSTM machine learning models, and can be reused for subscriptions and
	// recurring payments without storing sensitive data. Paykeys eliminate fraud by
	// ensuring the person initiating payment owns the funding account.
	Paykeys PaykeyService
	// Charges represent attempts to debit money from a customer's bank account using a
	// Paykey. Each charge includes automatic balance verification, real-time fraud
	// screening, and multi-rail optimization and detailed status tracking throughout
	// the payment lifecycle. Use charges to accept bank payments with confidence
	// knowing every transaction is protected.
	Charges ChargeService
	// Funding events represent all money movement between Straddle and an Account's
	// external bank accounts. They are automatically generated when charges settle or
	// payouts are initiated. Each event provides detailed tracking of settlement
	// status, fee breakdowns, and reconciliation data across both incoming and
	// outgoing transfers. Use funding events to monitor your platform's entire money
	// movement lifecycle.
	FundingEvents FundingEventService
	// Payments provide endpoints to filter both Charges and Payouts with multiple
	// different parameters.
	Payments PaymentService
	// Payouts represent transfers from Straddle to customer bank accounts. Create
	// payouts to handle disbursements, process refunds, or manage marketplace
	// settlements. Use payouts to send money quickly and securely with the most
	// cost-effective rail automatically selected.
	Payouts PayoutService
	Reports ReportService
}

// DefaultClientOptions read from the environment (STRADDLE_API_KEY,
// STRADDLE_BASE_URL). This should be used to initialize new clients.
func DefaultClientOptions() []option.RequestOption {
	defaults := []option.RequestOption{option.WithEnvironmentSandbox()}
	if o, ok := os.LookupEnv("STRADDLE_BASE_URL"); ok {
		defaults = append(defaults, option.WithBaseURL(o))
	}
	if o, ok := os.LookupEnv("STRADDLE_API_KEY"); ok {
		defaults = append(defaults, option.WithAPIKey(o))
	}
	return defaults
}

// NewClient generates a new client with the default option read from the
// environment (STRADDLE_API_KEY, STRADDLE_BASE_URL). The option passed in as
// arguments are applied after these default arguments, and all option will be
// passed down to the services and requests that this client makes.
func NewClient(opts ...option.RequestOption) (r Client) {
	opts = append(DefaultClientOptions(), opts...)

	r = Client{options: opts}

	r.Embed = NewEmbedService(opts...)
	r.Bridge = NewBridgeService(opts...)
	r.Customers = NewCustomerService(opts...)
	r.Paykeys = NewPaykeyService(opts...)
	r.Charges = NewChargeService(opts...)
	r.FundingEvents = NewFundingEventService(opts...)
	r.Payments = NewPaymentService(opts...)
	r.Payouts = NewPayoutService(opts...)
	r.Reports = NewReportService(opts...)

	return
}

// Execute makes a request with the given context, method, URL, request params,
// response, and request options. This is useful for hitting undocumented endpoints
// while retaining the base URL, auth, retries, and other options from the client.
//
// If a byte slice or an [io.Reader] is supplied to params, it will be used as-is
// for the request body.
//
// The params is by default serialized into the body using [encoding/json]. If your
// type implements a MarshalJSON function, it will be used instead to serialize the
// request. If a URLQuery method is implemented, the returned [url.Values] will be
// used as query strings to the url.
//
// If your params struct uses [param.Field], you must provide either [MarshalJSON],
// [URLQuery], and/or [MarshalForm] functions. It is undefined behavior to use a
// struct uses [param.Field] without specifying how it is serialized.
//
// Any "…Params" object defined in this library can be used as the request
// argument. Note that 'path' arguments will not be forwarded into the url.
//
// The response body will be deserialized into the res variable, depending on its
// type:
//
//   - A pointer to a [*http.Response] is populated by the raw response.
//   - A pointer to a byte array will be populated with the contents of the request
//     body.
//   - A pointer to any other type uses this library's default JSON decoding, which
//     respects UnmarshalJSON if it is defined on the type.
//   - A nil value will not read the response body.
//
// For even greater flexibility, see [option.WithResponseInto] and
// [option.WithResponseBodyInto].
func (r *Client) Execute(ctx context.Context, method string, path string, params any, res any, opts ...option.RequestOption) error {
	opts = slices.Concat(r.options, opts)
	return requestconfig.ExecuteNewRequest(ctx, method, path, params, res, opts...)
}

// Get makes a GET request with the given URL, params, and optionally deserializes
// to a response. See [Execute] documentation on the params and response.
func (r *Client) Get(ctx context.Context, path string, params any, res any, opts ...option.RequestOption) error {
	return r.Execute(ctx, http.MethodGet, path, params, res, opts...)
}

// Post makes a POST request with the given URL, params, and optionally
// deserializes to a response. See [Execute] documentation on the params and
// response.
func (r *Client) Post(ctx context.Context, path string, params any, res any, opts ...option.RequestOption) error {
	return r.Execute(ctx, http.MethodPost, path, params, res, opts...)
}

// Put makes a PUT request with the given URL, params, and optionally deserializes
// to a response. See [Execute] documentation on the params and response.
func (r *Client) Put(ctx context.Context, path string, params any, res any, opts ...option.RequestOption) error {
	return r.Execute(ctx, http.MethodPut, path, params, res, opts...)
}

// Patch makes a PATCH request with the given URL, params, and optionally
// deserializes to a response. See [Execute] documentation on the params and
// response.
func (r *Client) Patch(ctx context.Context, path string, params any, res any, opts ...option.RequestOption) error {
	return r.Execute(ctx, http.MethodPatch, path, params, res, opts...)
}

// Delete makes a DELETE request with the given URL, params, and optionally
// deserializes to a response. See [Execute] documentation on the params and
// response.
func (r *Client) Delete(ctx context.Context, path string, params any, res any, opts ...option.RequestOption) error {
	return r.Execute(ctx, http.MethodDelete, path, params, res, opts...)
}
