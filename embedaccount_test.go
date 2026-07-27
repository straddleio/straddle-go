// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package straddle_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/straddleio/straddle-go"
	"github.com/straddleio/straddle-go/internal/testutil"
	"github.com/straddleio/straddle-go/option"
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
		AccessLevel: straddle.EmbedAccountNewParamsAccessLevelStandard,
		AccountType: straddle.EmbedAccountNewParamsAccountTypeBusiness,
		BusinessProfile: straddle.BusinessProfileV1Param{
			Name:    "name",
			Website: "https://example.com",
			Address: straddle.AddressV1Param{
				Address1:   "address1",
				City:       straddle.String("city"),
				Line1:      straddle.String("line1"),
				PostalCode: straddle.String("21029-1360"),
				State:      straddle.String("SE"),
				Zip:        "zip",
				Address2:   straddle.String("address2"),
				Country:    straddle.String("country"),
				Line2:      straddle.String("line2"),
			},
			Description: straddle.String("description"),
			Industry: straddle.IndustryV1Param{
				Category: straddle.String("category"),
				Mcc:      straddle.String("mcc"),
				Sector:   straddle.String("sector"),
			},
			LegalName: straddle.String("legal_name"),
			Phone:     straddle.String("+46991022"),
			SupportChannels: straddle.SupportChannelsV1Param{
				Email: straddle.String("dev@stainless.com"),
				Phone: straddle.String("+46991022"),
				URL:   straddle.String("https://example.com"),
			},
			TaxID:   straddle.String("210297980"),
			UseCase: straddle.String("use_case"),
		},
		OrganizationID: "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		ExternalID:     straddle.String("external_id"),
		Metadata: map[string]string{
			"foo": "string",
		},
		CorrelationID:  straddle.String("correlation-id"),
		IdempotencyKey: straddle.String("xxxxxxxxxx"),
		RequestID:      straddle.String("request-id"),
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
			BusinessProfile: straddle.BusinessProfileV1Param{
				Name:    "name",
				Website: "https://example.com",
				Address: straddle.AddressV1Param{
					Address1:   "address1",
					City:       straddle.String("city"),
					Line1:      straddle.String("line1"),
					PostalCode: straddle.String("21029-1360"),
					State:      straddle.String("SE"),
					Zip:        "zip",
					Address2:   straddle.String("address2"),
					Country:    straddle.String("country"),
					Line2:      straddle.String("line2"),
				},
				Description: straddle.String("description"),
				Industry: straddle.IndustryV1Param{
					Category: straddle.String("category"),
					Mcc:      straddle.String("mcc"),
					Sector:   straddle.String("sector"),
				},
				LegalName: straddle.String("legal_name"),
				Phone:     straddle.String("+46991022"),
				SupportChannels: straddle.SupportChannelsV1Param{
					Email: straddle.String("dev@stainless.com"),
					Phone: straddle.String("+46991022"),
					URL:   straddle.String("https://example.com"),
				},
				TaxID:   straddle.String("210297980"),
				UseCase: straddle.String("use_case"),
			},
			ExternalID: straddle.String("external_id"),
			Metadata: map[string]string{
				"foo": "string",
			},
			CorrelationID:  straddle.String("correlation-id"),
			IdempotencyKey: straddle.String("xxxxxxxxxx"),
			RequestID:      straddle.String("request-id"),
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
		ExternalID:    straddle.String("external_id"),
		PageNumber:    straddle.Int(0),
		PageSize:      straddle.Int(0),
		SearchText:    straddle.String("search_text"),
		SortBy:        straddle.String("sort_by"),
		SortOrder:     straddle.EmbedAccountListParamsSortOrderAsc,
		Status:        straddle.EmbedAccountListParamsStatusCreated,
		Type:          straddle.EmbedAccountListParamsTypeBusiness,
		CorrelationID: straddle.String("correlation-id"),
		RequestID:     straddle.String("request-id"),
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
			CorrelationID: straddle.String("correlation-id"),
			RequestID:     straddle.String("request-id"),
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
			TermsOfService: straddle.TermsOfServiceV1Param{
				AcceptedDate:      time.Now(),
				AgreementType:     straddle.TermsOfServiceV1AgreementTypeEmbedded,
				AgreementURL:      straddle.String("agreement_url"),
				AcceptedIP:        straddle.String("accepted_ip"),
				AcceptedUserAgent: straddle.String("accepted_user_agent"),
			},
			CorrelationID:  straddle.String("correlation-id"),
			IdempotencyKey: straddle.String("xxxxxxxxxx"),
			RequestID:      straddle.String("request-id"),
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
			FinalStatus:    straddle.EmbedAccountSimulateParamsFinalStatusOnboarding,
			CorrelationID:  straddle.String("correlation-id"),
			IdempotencyKey: straddle.String("xxxxxxxxxx"),
			RequestID:      straddle.String("request-id"),
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
