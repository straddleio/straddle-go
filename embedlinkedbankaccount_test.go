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

func TestEmbedLinkedBankAccountNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Embed.LinkedBankAccounts.New(context.TODO(), straddle.EmbedLinkedBankAccountNewParams{
		AccountID: straddle.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		BankAccount: straddle.F(straddle.EmbedLinkedBankAccountNewParamsBankAccount{
			AccountHolder: straddle.F("account_holder"),
			AccountNumber: straddle.F("account_number"),
			RoutingNumber: straddle.F("xxxxxxxxx"),
		}),
		Metadata: straddle.F(map[string]string{
			"foo": "string",
		}),
		CorrelationID: straddle.F("correlation-id"),
		RequestID:     straddle.F("request-id"),
	})
	if err != nil {
		var apierr *straddle.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEmbedLinkedBankAccountUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Embed.LinkedBankAccounts.Update(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		straddle.EmbedLinkedBankAccountUpdateParams{
			BankAccount: straddle.F(straddle.EmbedLinkedBankAccountUpdateParamsBankAccount{
				AccountHolder: straddle.F("account_holder"),
				AccountNumber: straddle.F("account_number"),
				RoutingNumber: straddle.F("xxxxxxxxx"),
			}),
			Metadata: straddle.F(map[string]string{
				"foo": "string",
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

func TestEmbedLinkedBankAccountListWithOptionalParams(t *testing.T) {
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
	_, err := client.Embed.LinkedBankAccounts.List(context.TODO(), straddle.EmbedLinkedBankAccountListParams{
		AccountID:     straddle.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		PageNumber:    straddle.F(int64(0)),
		PageSize:      straddle.F(int64(0)),
		SortBy:        straddle.F("sort_by"),
		SortOrder:     straddle.F(straddle.EmbedLinkedBankAccountListParamsSortOrderAsc),
		CorrelationID: straddle.F("correlation-id"),
		RequestID:     straddle.F("request-id"),
	})
	if err != nil {
		var apierr *straddle.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEmbedLinkedBankAccountGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Embed.LinkedBankAccounts.Get(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		straddle.EmbedLinkedBankAccountGetParams{
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

func TestEmbedLinkedBankAccountUnmaskWithOptionalParams(t *testing.T) {
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
	_, err := client.Embed.LinkedBankAccounts.Unmask(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		straddle.EmbedLinkedBankAccountUnmaskParams{
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
