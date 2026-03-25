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

func TestBridgeLinkBankAccountWithOptionalParams(t *testing.T) {
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
	_, err := client.Bridge.Link.BankAccount(context.TODO(), straddle.BridgeLinkBankAccountParams{
		AccountNumber: "account_number",
		AccountType:   straddle.BridgeLinkBankAccountParamsAccountTypeChecking,
		CustomerID:    "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		RoutingNumber: "xxxxxxxxx",
		Config: straddle.BridgeLinkBankAccountParamsConfig{
			ProcessingMethod: "inline",
			SandboxOutcome:   "standard",
		},
		ExternalID: straddle.String("external_id"),
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

func TestBridgeLinkNewPaykeyWithOptionalParams(t *testing.T) {
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
	_, err := client.Bridge.Link.NewPaykey(context.TODO(), straddle.BridgeLinkNewPaykeyParams{
		CustomerID:  "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		QuilttToken: "quiltt_token",
		Config: straddle.BridgeLinkNewPaykeyParamsConfig{
			ProcessingMethod: "inline",
			SandboxOutcome:   "standard",
		},
		ExternalID: straddle.String("external_id"),
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

func TestBridgeLinkNewTanWithOptionalParams(t *testing.T) {
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
	_, err := client.Bridge.Link.NewTan(context.TODO(), straddle.BridgeLinkNewTanParams{
		AccountType:   straddle.BridgeLinkNewTanParamsAccountTypeChecking,
		CustomerID:    "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		RoutingNumber: "routing_number",
		Tan:           "tan",
		Config: straddle.BridgeLinkNewTanParamsConfig{
			ProcessingMethod: "inline",
			SandboxOutcome:   "standard",
		},
		ExternalID: straddle.String("external_id"),
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

func TestBridgeLinkPlaidWithOptionalParams(t *testing.T) {
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
	_, err := client.Bridge.Link.Plaid(context.TODO(), straddle.BridgeLinkPlaidParams{
		CustomerID: "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		PlaidToken: "plaid_token",
		Config: straddle.BridgeLinkPlaidParamsConfig{
			ProcessingMethod: "inline",
			SandboxOutcome:   "standard",
		},
		ExternalID: straddle.String("external_id"),
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
