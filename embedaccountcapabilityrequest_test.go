// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package straddle_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/straddleio/straddle-go"
	"github.com/straddleio/straddle-go/internal/testutil"
	"github.com/straddleio/straddle-go/option"
)

func TestEmbedAccountCapabilityRequestNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Embed.Accounts.CapabilityRequests.New(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		straddle.EmbedAccountCapabilityRequestNewParams{
			Businesses: straddle.EmbedAccountCapabilityRequestNewParamsBusinesses{
				Enable: true,
			},
			Charges: straddle.EmbedAccountCapabilityRequestNewParamsCharges{
				DailyAmount:   0,
				Enable:        true,
				MaxAmount:     0,
				MonthlyAmount: 0,
				MonthlyCount:  0,
			},
			Individuals: straddle.EmbedAccountCapabilityRequestNewParamsIndividuals{
				Enable: true,
			},
			Internet: straddle.EmbedAccountCapabilityRequestNewParamsInternet{
				Enable: true,
			},
			Payouts: straddle.EmbedAccountCapabilityRequestNewParamsPayouts{
				DailyAmount:   0,
				Enable:        true,
				MaxAmount:     0,
				MonthlyAmount: 0,
				MonthlyCount:  0,
			},
			SignedAgreement: straddle.EmbedAccountCapabilityRequestNewParamsSignedAgreement{
				Enable: true,
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

func TestEmbedAccountCapabilityRequestListWithOptionalParams(t *testing.T) {
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
	_, err := client.Embed.Accounts.CapabilityRequests.List(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		straddle.EmbedAccountCapabilityRequestListParams{
			Category:      straddle.EmbedAccountCapabilityRequestListParamsCategoryPaymentType,
			PageNumber:    straddle.Int(0),
			PageSize:      straddle.Int(0),
			SortBy:        straddle.String("sort_by"),
			SortOrder:     straddle.EmbedAccountCapabilityRequestListParamsSortOrderAsc,
			Status:        straddle.EmbedAccountCapabilityRequestListParamsStatusActive,
			Type:          straddle.EmbedAccountCapabilityRequestListParamsTypeCharges,
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
