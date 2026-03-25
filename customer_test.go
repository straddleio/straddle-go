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
)

func TestCustomerNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Customers.New(context.TODO(), straddle.CustomerNewParams{
		Device: straddle.DeviceUnmaskedV1Param{
			IPAddress: "192.168.1.1",
		},
		Email: "ron.swanson@pawnee.com",
		Name:  "Ron Swanson",
		Phone: "+12128675309",
		Type:  straddle.CustomerNewParamsTypeIndividual,
		Address: straddle.CustomerAddressV1Param{
			Address1: "123 Main St",
			City:     "Anytown",
			State:    "CA",
			Zip:      "12345",
			Address2: straddle.String("Apt 1"),
		},
		ComplianceProfile: straddle.CustomerNewParamsComplianceProfileUnion{
			OfIndividualComplianceProfile: &straddle.CustomerNewParamsComplianceProfileIndividualComplianceProfile{
				Dob: straddle.Time(time.Now()),
				Ssn: straddle.String("123-45-6789"),
			},
		},
		Config: straddle.CustomerNewParamsConfig{
			ProcessingMethod: "inline",
			SandboxOutcome:   "standard",
		},
		ExternalID: straddle.String("customer_123"),
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

func TestCustomerUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Customers.Update(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		straddle.CustomerUpdateParams{
			Device: straddle.DeviceUnmaskedV1Param{
				IPAddress: "192.168.1.1",
			},
			Email:  "dev@stainless.com",
			Name:   "name",
			Phone:  "+46991022",
			Status: straddle.CustomerUpdateParamsStatusPending,
			Address: straddle.CustomerAddressV1Param{
				Address1: "123 Main St",
				City:     "Anytown",
				State:    "CA",
				Zip:      "12345",
				Address2: straddle.String("Apt 1"),
			},
			ComplianceProfile: straddle.CustomerUpdateParamsComplianceProfileUnion{
				OfIndividualComplianceProfile: &straddle.CustomerUpdateParamsComplianceProfileIndividualComplianceProfile{
					Dob: straddle.Time(time.Now()),
					Ssn: straddle.String("123-45-6789"),
				},
			},
			ExternalID: straddle.String("external_id"),
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

func TestCustomerListWithOptionalParams(t *testing.T) {
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
	_, err := client.Customers.List(context.TODO(), straddle.CustomerListParams{
		CreatedFrom:       straddle.Time(time.Now()),
		CreatedTo:         straddle.Time(time.Now()),
		Email:             straddle.String("email"),
		ExternalID:        straddle.String("external_id"),
		Name:              straddle.String("name"),
		PageNumber:        straddle.Int(0),
		PageSize:          straddle.Int(0),
		SearchText:        straddle.String("search_text"),
		SortBy:            straddle.CustomerListParamsSortByName,
		SortOrder:         straddle.CustomerListParamsSortOrderAsc,
		Status:            []string{"pending"},
		Types:             []string{"individual"},
		CorrelationID:     straddle.String("Correlation-Id"),
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

func TestCustomerDeleteWithOptionalParams(t *testing.T) {
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
	_, err := client.Customers.Delete(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		straddle.CustomerDeleteParams{
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

func TestCustomerGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Customers.Get(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		straddle.CustomerGetParams{
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

func TestCustomerUnmaskedWithOptionalParams(t *testing.T) {
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
	_, err := client.Customers.Unmasked(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		straddle.CustomerUnmaskedParams{
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
