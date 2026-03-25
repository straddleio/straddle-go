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
		CreatedFrom:       straddle.Time(time.Now()),
		CreatedTo:         straddle.Time(time.Now()),
		Direction:         straddle.FundingEventListParamsDirectionDeposit,
		EventType:         straddle.FundingEventListParamsEventTypeChargeDeposit,
		PageNumber:        straddle.Int(0),
		PageSize:          straddle.Int(0),
		SearchText:        straddle.String("search_text"),
		SortBy:            straddle.FundingEventListParamsSortByTransferDate,
		SortOrder:         straddle.FundingEventListParamsSortOrderAsc,
		Status:            []string{"created"},
		StatusReason:      []string{"insufficient_funds"},
		StatusSource:      []string{"watchtower"},
		TraceID:           straddle.String("trace_id"),
		TraceNumber:       straddle.String("trace_number"),
		CorrelationID:     straddle.String("Correlation-Id"),
		RequestID:         straddle.String("Request-Id"),
		StraddleAccountID: straddle.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
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
			CorrelationID:     straddle.String("Correlation-Id"),
			RequestID:         straddle.String("Request-Id"),
			StraddleAccountID: straddle.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
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
