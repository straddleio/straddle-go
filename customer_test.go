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
		Device: straddle.F(straddle.DeviceUnmaskedV1Param{
			IPAddress: straddle.F("192.168.1.1"),
		}),
		Email: straddle.F("ron.swanson@pawnee.com"),
		Name:  straddle.F("Ron Swanson"),
		Phone: straddle.F("+12128675309"),
		Type:  straddle.F(straddle.CustomerNewParamsTypeIndividual),
		Address: straddle.F(straddle.CustomerAddressV1Param{
			Address1: straddle.F("123 Main St"),
			City:     straddle.F("Anytown"),
			State:    straddle.F("CA"),
			Zip:      straddle.F("94105"),
			Address2: straddle.F("address2"),
		}),
		ComplianceProfile: straddle.F[straddle.CustomerNewParamsComplianceProfileUnion](straddle.CustomerNewParamsComplianceProfileIndividualCustomerComplianceProfile{
			Dob: straddle.F(time.Now()),
			Ssn: straddle.F("123-45-6789"),
		}),
		ExternalID:        straddle.F("customer_123"),
		Metadata:          straddle.F(map[string]string{}),
		CorrelationID:     straddle.F("Correlation-Id"),
		RequestID:         straddle.F("Request-Id"),
		StraddleAccountID: straddle.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
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
			Device: straddle.F(straddle.DeviceUnmaskedV1Param{
				IPAddress: straddle.F("192.168.1.1"),
			}),
			Email:  straddle.F("dev@stainless.com"),
			Name:   straddle.F("name"),
			Phone:  straddle.F("+46991022"),
			Status: straddle.F(straddle.CustomerUpdateParamsStatusPending),
			Address: straddle.F(straddle.CustomerAddressV1Param{
				Address1: straddle.F("123 Main St"),
				City:     straddle.F("Anytown"),
				State:    straddle.F("CA"),
				Zip:      straddle.F("94105"),
				Address2: straddle.Null[string](),
			}),
			ComplianceProfile: straddle.F[straddle.CustomerUpdateParamsComplianceProfileUnion](straddle.CustomerUpdateParamsComplianceProfileIndividualCustomerComplianceProfile{
				Dob: straddle.F(time.Now()),
				Ssn: straddle.F("123-45-6789"),
			}),
			ExternalID: straddle.F("external_id"),
			Metadata: straddle.F(map[string]string{
				"foo": "string",
			}),
			CorrelationID:     straddle.F("Correlation-Id"),
			RequestID:         straddle.F("Request-Id"),
			StraddleAccountID: straddle.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
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
		CreatedFrom:       straddle.F(time.Now()),
		CreatedTo:         straddle.F(time.Now()),
		Email:             straddle.F("email"),
		ExternalID:        straddle.F("external_id"),
		Name:              straddle.F("name"),
		PageNumber:        straddle.F(int64(0)),
		PageSize:          straddle.F(int64(0)),
		SearchText:        straddle.F("search_text"),
		SortBy:            straddle.F(straddle.CustomerListParamsSortByName),
		SortOrder:         straddle.F(straddle.CustomerListParamsSortOrderAsc),
		Status:            straddle.F([]straddle.CustomerListParamsStatus{straddle.CustomerListParamsStatusPending}),
		Types:             straddle.F([]straddle.CustomerListParamsType{straddle.CustomerListParamsTypeIndividual}),
		CorrelationID:     straddle.F("Correlation-Id"),
		RequestID:         straddle.F("Request-Id"),
		StraddleAccountID: straddle.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
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
			CorrelationID:     straddle.F("Correlation-Id"),
			RequestID:         straddle.F("Request-Id"),
			StraddleAccountID: straddle.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
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
			CorrelationID:     straddle.F("Correlation-Id"),
			RequestID:         straddle.F("Request-Id"),
			StraddleAccountID: straddle.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
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

func TestCustomerRefreshReviewWithOptionalParams(t *testing.T) {
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
	_, err := client.Customers.RefreshReview(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		straddle.CustomerRefreshReviewParams{
			CorrelationID:     straddle.F("Correlation-Id"),
			RequestID:         straddle.F("Request-Id"),
			StraddleAccountID: straddle.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
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
			CorrelationID:     straddle.F("Correlation-Id"),
			RequestID:         straddle.F("Request-Id"),
			StraddleAccountID: straddle.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
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
