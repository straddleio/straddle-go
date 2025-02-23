// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package straddle_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/straddle-go"
	"github.com/stainless-sdks/straddle-go/internal/testutil"
	"github.com/stainless-sdks/straddle-go/option"
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
			Businesses: straddle.F(straddle.EmbedAccountCapabilityRequestNewParamsBusinesses{
				Enable: straddle.F(true),
			}),
			Charges: straddle.F(straddle.EmbedAccountCapabilityRequestNewParamsCharges{
				DailyAmount:   straddle.F(0.000000),
				Enable:        straddle.F(true),
				MaxAmount:     straddle.F(0.000000),
				MonthlyAmount: straddle.F(0.000000),
				MonthlyCount:  straddle.F(int64(0)),
			}),
			Individuals: straddle.F(straddle.EmbedAccountCapabilityRequestNewParamsIndividuals{
				Enable: straddle.F(true),
			}),
			Internet: straddle.F(straddle.EmbedAccountCapabilityRequestNewParamsInternet{
				Enable: straddle.F(true),
			}),
			Payouts: straddle.F(straddle.EmbedAccountCapabilityRequestNewParamsPayouts{
				DailyAmount:   straddle.F(0.000000),
				Enable:        straddle.F(true),
				MaxAmount:     straddle.F(0.000000),
				MonthlyAmount: straddle.F(0.000000),
				MonthlyCount:  straddle.F(int64(0)),
			}),
			SignedAgreement: straddle.F(straddle.EmbedAccountCapabilityRequestNewParamsSignedAgreement{
				Enable: straddle.F(true),
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
			Category:      straddle.F(straddle.EmbedAccountCapabilityRequestListParamsCategoryPaymentType),
			PageNumber:    straddle.F(int64(0)),
			PageSize:      straddle.F(int64(0)),
			SortBy:        straddle.F("sort_by"),
			SortOrder:     straddle.F(straddle.EmbedAccountCapabilityRequestListParamsSortOrderAsc),
			Status:        straddle.F(straddle.EmbedAccountCapabilityRequestListParamsStatusActive),
			Type:          straddle.F(straddle.EmbedAccountCapabilityRequestListParamsTypeCharges),
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
