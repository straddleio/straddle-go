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
		AccountID:    "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		Dob:          time.Now(),
		Email:        "ron.swanson@pawnee.com",
		FirstName:    "first_name",
		LastName:     "last_name",
		MobileNumber: "+12128675309",
		Relationship: straddle.EmbedRepresentativeNewParamsRelationship{
			Control:          true,
			Owner:            true,
			Primary:          true,
			PercentOwnership: straddle.Float(0),
			Title:            straddle.String("title"),
		},
		SsnLast4:   "1234",
		ExternalID: straddle.String("external_id"),
		Metadata: map[string]string{
			"foo": "string",
		},
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
			Dob:          time.Now(),
			Email:        "ron.swanson@pawnee.com",
			FirstName:    "Ron",
			LastName:     "Swanson",
			MobileNumber: "+12128675309",
			Relationship: straddle.EmbedRepresentativeUpdateParamsRelationship{
				Control:          true,
				Owner:            true,
				Primary:          true,
				PercentOwnership: straddle.Float(0),
				Title:            straddle.String("title"),
			},
			SsnLast4:   "1234",
			ExternalID: straddle.String("external_id"),
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
		AccountID:      straddle.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Level:          straddle.EmbedRepresentativeListParamsLevelAccount,
		OrganizationID: straddle.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		PageNumber:     straddle.Int(0),
		PageSize:       straddle.Int(0),
		PlatformID:     straddle.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		SortBy:         straddle.String("sort_by"),
		SortOrder:      straddle.EmbedRepresentativeListParamsSortOrderAsc,
		CorrelationID:  straddle.String("correlation-id"),
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
