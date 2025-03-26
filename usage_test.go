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
		Amount: straddle.F(int64(0)),
		Config: straddle.F(straddle.ChargeNewParamsConfig{
			BalanceCheck: straddle.F(straddle.ChargeNewParamsConfigBalanceCheckRequired),
		}),
		ConsentType: straddle.F(straddle.ChargeNewParamsConsentTypeInternet),
		Currency:    straddle.F("currency"),
		Description: straddle.F("Monthly subscription fee"),
		Device: straddle.F(shared.DeviceInfoV1Param{
			IPAddress: straddle.F("192.168.1.1"),
		}),
		ExternalID:  straddle.F("external_id"),
		Paykey:      straddle.F("paykey"),
		PaymentDate: straddle.F(time.Now()),
	})
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%+v\n", chargeV1.Data)
}
