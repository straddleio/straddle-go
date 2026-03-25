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

func TestPaymentListWithOptionalParams(t *testing.T) {
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
	_, err := client.Payments.List(context.TODO(), straddle.PaymentListParams{
		CustomerID:        straddle.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		DefaultPageSize:   straddle.Int(0),
		DefaultSort:       straddle.PaymentListParamsDefaultSortCreatedAt,
		DefaultSortOrder:  straddle.PaymentListParamsDefaultSortOrderAsc,
		ExternalID:        straddle.String("external_id"),
		FundingID:         straddle.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		IncludeMetadata:   straddle.Bool(true),
		MaxAmount:         straddle.Int(0),
		MaxCreatedAt:      straddle.Time(time.Now()),
		MaxEffectiveAt:    straddle.Time(time.Now()),
		MaxPaymentDate:    straddle.Time(time.Now()),
		MinAmount:         straddle.Int(0),
		MinCreatedAt:      straddle.Time(time.Now()),
		MinEffectiveAt:    straddle.Time(time.Now()),
		MinPaymentDate:    straddle.Time(time.Now()),
		PageNumber:        straddle.Int(0),
		PageSize:          straddle.Int(0),
		Paykey:            straddle.String("paykey"),
		PaykeyID:          straddle.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		PaymentID:         straddle.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		PaymentStatus:     []string{"created"},
		PaymentType:       []string{"charge"},
		SearchText:        straddle.String("search_text"),
		SortBy:            straddle.PaymentListParamsSortByCreatedAt,
		SortOrder:         straddle.PaymentListParamsSortOrderAsc,
		StatusReason:      []string{"insufficient_funds"},
		StatusSource:      []string{"watchtower"},
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
