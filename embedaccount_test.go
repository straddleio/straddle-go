// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package straddle_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/stainless-sdks/straddle-go"
	"github.com/stainless-sdks/straddle-go/internal/testutil"
	"github.com/stainless-sdks/straddle-go/option"
)

func TestEmbedAccountNewWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := straddle.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Embed.Accounts.New(context.TODO(), straddle.EmbedAccountNewParams{
		AccessLevel: straddle.F(straddle.EmbedAccountNewParamsAccessLevelStandard),
		AccountType: straddle.F(straddle.EmbedAccountNewParamsAccountTypeBusiness),
		BusinessProfile: straddle.F(straddle.BusinessProfileV1Param{
			Name:    straddle.F("name"),
			Website: straddle.F("https://example.com"),
			Address: straddle.F(straddle.AddressV1Param{
				City:       straddle.F("city"),
				Country:    straddle.F("country"),
				Line1:      straddle.F("line1"),
				Line2:      straddle.F("line2"),
				PostalCode: straddle.F("21029-1360"),
				State:      straddle.F("SE"),
			}),
			Description: straddle.F("description"),
			Industry: straddle.F(straddle.IndustryV1Param{
				Category: straddle.F("category"),
				Mcc:      straddle.F("mcc"),
				Sector:   straddle.F("sector"),
			}),
			LegalName: straddle.F("legal_name"),
			Phone:     straddle.F("+46991022"),
			SupportChannels: straddle.F(straddle.SupportChannelsV1Param{
				Email: straddle.F("dev@stainless.com"),
				Phone: straddle.F("+46991022"),
				URL:   straddle.F("https://example.com"),
			}),
			TaxID:   straddle.F("210297980"),
			UseCase: straddle.F("use_case"),
		}),
		OrganizationID: straddle.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		ExternalID:     straddle.F("external_id"),
		Metadata: straddle.F(map[string]string{
			"foo": "string",
		}),
		CorrelationID: straddle.F("correlation-id"),
		RequestID:     straddle.F("request-id"),
	})
	if err != nil {
		var apierr *straddle.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEmbedAccountUpdateWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := straddle.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Embed.Accounts.Update(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		straddle.EmbedAccountUpdateParams{
			BusinessProfile: straddle.F(straddle.BusinessProfileV1Param{
				Name:    straddle.F("name"),
				Website: straddle.F("https://example.com"),
				Address: straddle.F(straddle.AddressV1Param{
					City:       straddle.F("city"),
					Country:    straddle.F("country"),
					Line1:      straddle.F("line1"),
					Line2:      straddle.F("line2"),
					PostalCode: straddle.F("21029-1360"),
					State:      straddle.F("SE"),
				}),
				Description: straddle.F("description"),
				Industry: straddle.F(straddle.IndustryV1Param{
					Category: straddle.F("category"),
					Mcc:      straddle.F("mcc"),
					Sector:   straddle.F("sector"),
				}),
				LegalName: straddle.F("legal_name"),
				Phone:     straddle.F("+46991022"),
				SupportChannels: straddle.F(straddle.SupportChannelsV1Param{
					Email: straddle.F("dev@stainless.com"),
					Phone: straddle.F("+46991022"),
					URL:   straddle.F("https://example.com"),
				}),
				TaxID:   straddle.F("210297980"),
				UseCase: straddle.F("use_case"),
			}),
			ExternalID: straddle.F("external_id"),
			Metadata: straddle.F(map[string]string{
				"foo": "string",
			}),
			CorrelationID: straddle.F("correlation-id"),
			RequestID:     straddle.F("request-id"),
		},
	)
	if err != nil {
		var apierr *straddle.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEmbedAccountListWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := straddle.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Embed.Accounts.List(context.TODO(), straddle.EmbedAccountListParams{
		PageNumber:    straddle.F(int64(0)),
		PageSize:      straddle.F(int64(0)),
		SearchText:    straddle.F("search_text"),
		SortBy:        straddle.F("sort_by"),
		SortOrder:     straddle.F(straddle.EmbedAccountListParamsSortOrderAsc),
		CorrelationID: straddle.F("correlation-id"),
		RequestID:     straddle.F("request-id"),
	})
	if err != nil {
		var apierr *straddle.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEmbedAccountGetWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := straddle.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Embed.Accounts.Get(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		straddle.EmbedAccountGetParams{
			CorrelationID: straddle.F("correlation-id"),
			RequestID:     straddle.F("request-id"),
		},
	)
	if err != nil {
		var apierr *straddle.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEmbedAccountOnboardWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := straddle.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Embed.Accounts.Onboard(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		straddle.EmbedAccountOnboardParams{
			TermsOfService: straddle.F(straddle.TermsOfServiceV1Param{
				AcceptedDate:      straddle.F(time.Now()),
				AgreementType:     straddle.F(straddle.TermsOfServiceV1AgreementTypeEmbedded),
				AgreementURL:      straddle.F("agreement_url"),
				AcceptedIP:        straddle.F("accepted_ip"),
				AcceptedUserAgent: straddle.F("accepted_user_agent"),
			}),
			CorrelationID: straddle.F("correlation-id"),
			RequestID:     straddle.F("request-id"),
		},
	)
	if err != nil {
		var apierr *straddle.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEmbedAccountSimulateWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := straddle.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Embed.Accounts.Simulate(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		straddle.EmbedAccountSimulateParams{
			FinalStatus:   straddle.F(straddle.EmbedAccountSimulateParamsFinalStatusOnboarding),
			CorrelationID: straddle.F("correlation-id"),
			RequestID:     straddle.F("request-id"),
		},
	)
	if err != nil {
		var apierr *straddle.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
