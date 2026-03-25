// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package straddle_test

import (
	"context"
	"os"
	"testing"

	"github.com/straddleio/straddle-go"
	"github.com/straddleio/straddle-go/internal/testutil"
	"github.com/straddleio/straddle-go/option"
)

func TestAutoPagination(t *testing.T) {
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
	iter := client.Payments.ListAutoPaging(context.TODO(), straddle.PaymentListParams{})
	// The mock server isn't going to give us real pagination
	for i := 0; i < 3 && iter.Next(); i++ {
		payment := iter.Current()
		t.Logf("%+v\n", payment.ID)
	}
	if err := iter.Err(); err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
