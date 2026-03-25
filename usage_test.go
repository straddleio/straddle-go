// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package straddle_test

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/stainless-sdks/straddle-go"
	"github.com/stainless-sdks/straddle-go/internal/testutil"
	"github.com/stainless-sdks/straddle-go/option"
	"github.com/stainless-sdks/straddle-go/shared"
)

func TestUsage(t *testing.T) {
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
	chargeV1, err := client.Charges.New(context.TODO(), straddle.ChargeNewParams{
		Amount: 10000,
		Config: straddle.ChargeNewParamsConfig{
			BalanceCheck: "required",
		},
		ConsentType: straddle.ChargeNewParamsConsentTypeInternet,
		Currency:    "currency",
		Description: straddle.String("Monthly subscription fee"),
		Device: shared.DeviceInfoV1Param{
			IPAddress: "192.168.1.1",
		},
		ExternalID:  "external_id",
		Paykey:      "paykey",
		PaymentDate: time.Now(),
	})
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	t.Logf("%+v\n", chargeV1.Data)
}
