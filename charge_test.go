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
	"github.com/straddleio/straddle-go/shared"
)

func TestChargeNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Charges.New(context.TODO(), straddle.ChargeNewParams{
		Amount: 10000,
		Config: straddle.ChargeNewParamsConfig{
			BalanceCheck:    "required",
			AutoHold:        straddle.Bool(true),
			AutoHoldMessage: straddle.String("auto_hold_message"),
			SandboxOutcome:  "standard",
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
		Metadata: map[string]string{
			"foo": "string",
		},
		CorrelationID:     straddle.String("Correlation-Id"),
		IdempotencyKey:    straddle.String("xxxxxxxxxx"),
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

func TestChargeUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Charges.Update(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		straddle.ChargeUpdateParams{
			Amount:      10000,
			Description: straddle.String("Monthly subscription fee"),
			PaymentDate: time.Now(),
			Metadata: map[string]string{
				"foo": "string",
			},
			CorrelationID:     straddle.String("Correlation-Id"),
			IdempotencyKey:    straddle.String("xxxxxxxxxx"),
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

func TestChargeCancelWithOptionalParams(t *testing.T) {
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
	_, err := client.Charges.Cancel(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		straddle.ChargeCancelParams{
			Reason:            straddle.String("reason"),
			CorrelationID:     straddle.String("Correlation-Id"),
			IdempotencyKey:    straddle.String("xxxxxxxxxx"),
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

func TestChargeGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Charges.Get(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		straddle.ChargeGetParams{
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

func TestChargeHoldWithOptionalParams(t *testing.T) {
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
	_, err := client.Charges.Hold(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		straddle.ChargeHoldParams{
			Reason:            straddle.String("reason"),
			CorrelationID:     straddle.String("Correlation-Id"),
			IdempotencyKey:    straddle.String("xxxxxxxxxx"),
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

func TestChargeReleaseWithOptionalParams(t *testing.T) {
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
	_, err := client.Charges.Release(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		straddle.ChargeReleaseParams{
			Reason:            straddle.String("reason"),
			CorrelationID:     straddle.String("Correlation-Id"),
			IdempotencyKey:    straddle.String("xxxxxxxxxx"),
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

func TestChargeUnmaskWithOptionalParams(t *testing.T) {
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
	_, err := client.Charges.Unmask(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		straddle.ChargeUnmaskParams{
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
