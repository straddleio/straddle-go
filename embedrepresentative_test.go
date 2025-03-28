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
)

func TestEmbedRepresentativeNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Embed.Representatives.New(context.TODO(), straddle.EmbedRepresentativeNewParams{
		AccountID:    straddle.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Dob:          straddle.F(time.Now()),
		Email:        straddle.F("ron.swanson@pawnee.com"),
		FirstName:    straddle.F("first_name"),
		LastName:     straddle.F("last_name"),
		MobileNumber: straddle.F("+12128675309"),
		Relationship: straddle.F(straddle.EmbedRepresentativeNewParamsRelationship{
			Control:          straddle.F(true),
			Owner:            straddle.F(true),
			Primary:          straddle.F(true),
			PercentOwnership: straddle.F(0.000000),
			Title:            straddle.F("title"),
		}),
		SsnLast4:      straddle.F("1234"),
		ExternalID:    straddle.F("external_id"),
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

func TestEmbedRepresentativeUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Embed.Representatives.Update(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		straddle.EmbedRepresentativeUpdateParams{
			Dob:          straddle.F(time.Now()),
			Email:        straddle.F("ron.swanson@pawnee.com"),
			FirstName:    straddle.F("Ron"),
			LastName:     straddle.F("Swanson"),
			MobileNumber: straddle.F("+12128675309"),
			Relationship: straddle.F(straddle.EmbedRepresentativeUpdateParamsRelationship{
				Control:          straddle.F(true),
				Owner:            straddle.F(true),
				Primary:          straddle.F(true),
				PercentOwnership: straddle.F(0.000000),
				Title:            straddle.F("title"),
			}),
			SsnLast4:      straddle.F("1234"),
			ExternalID:    straddle.F("external_id"),
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

func TestEmbedRepresentativeListWithOptionalParams(t *testing.T) {
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
	_, err := client.Embed.Representatives.List(context.TODO(), straddle.EmbedRepresentativeListParams{
		AccountID:      straddle.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Level:          straddle.F(straddle.EmbedRepresentativeListParamsLevelAccount),
		OrganizationID: straddle.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		PageNumber:     straddle.F(int64(0)),
		PageSize:       straddle.F(int64(0)),
		PlatformID:     straddle.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		SortBy:         straddle.F("sort_by"),
		SortOrder:      straddle.F(straddle.EmbedRepresentativeListParamsSortOrderAsc),
		CorrelationID:  straddle.F("correlation-id"),
		RequestID:      straddle.F("request-id"),
	})
	if err != nil {
		var apierr *straddle.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEmbedRepresentativeGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Embed.Representatives.Get(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		straddle.EmbedRepresentativeGetParams{
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

func TestEmbedRepresentativeUnmaskWithOptionalParams(t *testing.T) {
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
	_, err := client.Embed.Representatives.Unmask(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		straddle.EmbedRepresentativeUnmaskParams{
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
