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
	options []option.RequestOption
	// Accounts represent businesses using Straddle through your platform. Each account
	// must complete automated verification before processing payments. Use accounts to
	// manage your users' payment capabilities, track verification status, and control
	// access to features. Accounts can be instantly created in sandbox and require
	// additional verification for production access.
	Accounts EmbedAccountService
	// Linked bank accounts connect your platform users' external bank accounts to
	// Straddle for settlements and payment funding. Each linked account undergoes
	// automated verification and continuous monitoring. Use linked accounts to manage
	// where clients receive deposits, fund payouts, and track settlement preferences.
	LinkedBankAccounts EmbedLinkedBankAccountService
	// Organizations are a powerful feature in Straddle that allow you to manage
	// multiple accounts under a single umbrella. This hierarchical structure is
	// particularly useful for businesses with complex operations, multiple
	// departments, or legally related entities.
	Organizations EmbedOrganizationService
	// Representatives are individuals who have legal authority or significant
	// responsibility within a business entity associated with a Straddle account. Each
	// representative undergoes automated verification as part of KYC/KYB compliance.
	// Use representatives to collect and verify beneficial owners, control persons,
	// and authorized signers required for account onboarding. Representatives also
	// determine who can legally operate the account and make important changes.
	Representatives EmbedRepresentativeService
}

// NewEmbedService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewEmbedService(opts ...option.RequestOption) (r EmbedService) {
	r = EmbedService{}
	r.options = opts
	r.Accounts = NewEmbedAccountService(opts...)
	r.LinkedBankAccounts = NewEmbedLinkedBankAccountService(opts...)
	r.Organizations = NewEmbedOrganizationService(opts...)
	r.Representatives = NewEmbedRepresentativeService(opts...)
	return
}
