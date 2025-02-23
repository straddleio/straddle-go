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

func TestEmbedOrganizationNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Embed.Organizations.New(context.TODO(), straddle.EmbedOrganizationNewParams{
		Name:       straddle.F("name"),
		ExternalID: straddle.F("external_id"),
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

func TestEmbedOrganizationListWithOptionalParams(t *testing.T) {
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
	_, err := client.Embed.Organizations.List(context.TODO(), straddle.EmbedOrganizationListParams{
		ExternalID:    straddle.F("external_id"),
		Name:          straddle.F("name"),
		PageNumber:    straddle.F(int64(0)),
		PageSize:      straddle.F(int64(0)),
		SortBy:        straddle.F("sort_by"),
		SortOrder:     straddle.F(straddle.EmbedOrganizationListParamsSortOrderAsc),
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

func TestEmbedOrganizationGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Embed.Organizations.Get(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		straddle.EmbedOrganizationGetParams{
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
