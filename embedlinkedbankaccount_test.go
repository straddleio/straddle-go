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
		AccountID: straddle.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		BankAccount: straddle.EmbedLinkedBankAccountNewParamsBankAccount{
			AccountHolder: "account_holder",
			AccountNumber: "account_number",
			RoutingNumber: "xxxxxxxxx",
		},
		Description: straddle.String("description"),
		Metadata: map[string]string{
			"foo": "string",
		},
		PlatformID:     straddle.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Purposes:       []string{"charges"},
		CorrelationID:  straddle.String("correlation-id"),
		IdempotencyKey: straddle.String("xxxxxxxxxx"),
		RequestID:      straddle.String("request-id"),
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
			BankAccount: straddle.EmbedLinkedBankAccountUpdateParamsBankAccount{
				AccountHolder: "account_holder",
				AccountNumber: "account_number",
				RoutingNumber: "xxxxxxxxx",
			},
			Metadata: map[string]string{
				"foo": "string",
			},
			CorrelationID:  straddle.String("correlation-id"),
			IdempotencyKey: straddle.String("xxxxxxxxxx"),
			RequestID:      straddle.String("request-id"),
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
		AccountID:     straddle.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Level:         straddle.EmbedLinkedBankAccountListParamsLevelAccount,
		PageNumber:    straddle.Int(0),
		PageSize:      straddle.Int(0),
		Purpose:       straddle.EmbedLinkedBankAccountListParamsPurposeCharges,
		SortBy:        straddle.String("sort_by"),
		SortOrder:     straddle.EmbedLinkedBankAccountListParamsSortOrderAsc,
		Status:        straddle.EmbedLinkedBankAccountListParamsStatusCreated,
		CorrelationID: straddle.String("correlation-id"),
		RequestID:     straddle.String("request-id"),
	})
	if err != nil {
		var apierr *straddle.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEmbedLinkedBankAccountCancelWithOptionalParams(t *testing.T) {
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
	_, err := client.Embed.LinkedBankAccounts.Cancel(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		straddle.EmbedLinkedBankAccountCancelParams{
			CorrelationID:  straddle.String("correlation-id"),
			IdempotencyKey: straddle.String("xxxxxxxxxx"),
			RequestID:      straddle.String("request-id"),
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
			CorrelationID: straddle.String("correlation-id"),
			RequestID:     straddle.String("request-id"),
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
			CorrelationID: straddle.String("correlation-id"),
			RequestID:     straddle.String("request-id"),
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
