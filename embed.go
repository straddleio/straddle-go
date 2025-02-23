// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package straddle

import (
	"github.com/stainless-sdks/straddle-go/option"
)

// EmbedService contains methods and other services that help with interacting with
// the straddle API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEmbedService] method instead.
type EmbedService struct {
	Options            []option.RequestOption
	Accounts           *EmbedAccountService
	LinkedBankAccounts *EmbedLinkedBankAccountService
	Organizations      *EmbedOrganizationService
	Representatives    *EmbedRepresentativeService
}

// NewEmbedService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewEmbedService(opts ...option.RequestOption) (r *EmbedService) {
	r = &EmbedService{}
	r.Options = opts
	r.Accounts = NewEmbedAccountService(opts...)
	r.LinkedBankAccounts = NewEmbedLinkedBankAccountService(opts...)
	r.Organizations = NewEmbedOrganizationService(opts...)
	r.Representatives = NewEmbedRepresentativeService(opts...)
	return
}
