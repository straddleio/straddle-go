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

func TestFundingEventListWithOptionalParams(t *testing.T) {
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
	_, err := client.FundingEvents.List(context.TODO(), straddle.FundingEventListParams{
		CreatedFrom:       straddle.F(time.Now()),
		CreatedTo:         straddle.F(time.Now()),
		Direction:         straddle.F(straddle.FundingEventListParamsDirectionDeposit),
		EventType:         straddle.F(straddle.FundingEventListParamsEventTypeChargeDeposit),
		PageNumber:        straddle.F(int64(0)),
		PageSize:          straddle.F(int64(0)),
		SortBy:            straddle.F(straddle.FundingEventListParamsSortByTransferDate),
		SortOrder:         straddle.F(straddle.FundingEventListParamsSortOrderAsc),
		TraceNumber:       straddle.F("trace_number"),
		CorrelationID:     straddle.F("Correlation-Id"),
		RequestID:         straddle.F("Request-Id"),
		StraddleAccountID: straddle.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
	})
	if err != nil {
		var apierr *straddle.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestFundingEventGetWithOptionalParams(t *testing.T) {
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
	_, err := client.FundingEvents.Get(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		straddle.FundingEventGetParams{
			CorrelationID:     straddle.F("Correlation-Id"),
			RequestID:         straddle.F("Request-Id"),
			StraddleAccountID: straddle.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
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
