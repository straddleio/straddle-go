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
	"github.com/stainless-sdks/straddle-go/shared"
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
		CustomerID:        straddle.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		DefaultPageSize:   straddle.F(int64(0)),
		DefaultSort:       straddle.F(shared.PaymentSortByV1CreatedAt),
		DefaultSortOrder:  straddle.F(shared.SortOrderAsc),
		ExternalID:        straddle.F("external_id"),
		FundingID:         straddle.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		MaxAmount:         straddle.F(int64(0)),
		MaxCreatedAt:      straddle.F(time.Now()),
		MaxEffectiveAt:    straddle.F(time.Now()),
		MaxPaymentDate:    straddle.F(time.Now()),
		MinAmount:         straddle.F(int64(0)),
		MinCreatedAt:      straddle.F(time.Now()),
		MinEffectiveAt:    straddle.F(time.Now()),
		MinPaymentDate:    straddle.F(time.Now()),
		PageNumber:        straddle.F(int64(0)),
		PageSize:          straddle.F(int64(0)),
		Paykey:            straddle.F("paykey"),
		PaykeyID:          straddle.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		PaymentID:         straddle.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		PaymentStatus:     straddle.F([]shared.PaymentStatusV1{shared.PaymentStatusV1Created}),
		PaymentType:       straddle.F([]shared.PaymentTypeV1{shared.PaymentTypeV1Charge}),
		SearchText:        straddle.F("search_text"),
		SortBy:            straddle.F(shared.PaymentSortByV1CreatedAt),
		SortOrder:         straddle.F(shared.SortOrderAsc),
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
