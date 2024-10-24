package overrides_test

import (
	sdk "apigee_api/client"
	nullable "apigee_api/nullable"
	overrides "apigee_api/resources/v1/overrides"
	types "apigee_api/types"
	fmt "fmt"
	testing "testing"
)

func TestList200GeneratedSuccess(t *testing.T) {
	// Success case test
	client := sdk.NewClient(sdk.WithOauth2("API_TOKEN"), sdk.WithOauth2c("API_TOKEN"), sdk.WithBaseURL("https://api.sideko.dev/v1/mock/public/apigee/0.1.0"))
	res, err := client.V1.Overrides.List(overrides.ListRequest{Parent: "string"})

	if err != nil {
		t.Fatalf("TestList200GeneratedSuccess - failed making request with error: %#v", err)
	}

	fmt.Printf("response - %#v\n", res)
}

func TestCreate200SuccessDefault(t *testing.T) {
	// Success test for Default body
	client := sdk.NewClient(sdk.WithOauth2("API_TOKEN"), sdk.WithOauth2c("API_TOKEN"), sdk.WithBaseURL("https://api.sideko.dev/v1/mock/public/apigee/0.1.0"))
	res, err := client.V1.Overrides.Create(overrides.CreateRequest{Parent: "string", Data: types.GoogleCloudApigeeV1TraceConfigOverride{ApiProxy: nullable.NewValue("string"), Name: nullable.NewValue("string"), SamplingConfig: nullable.NewValue(types.GoogleCloudApigeeV1TraceSamplingConfig{Sampler: nullable.NewValue(types.GoogleCloudApigeeV1TraceSamplingConfigSamplerEnumOff), SamplingRate: nullable.NewValue(123.45)})}})

	if err != nil {
		t.Fatalf("TestCreate200SuccessDefault - failed making request with error: %#v", err)
	}

	fmt.Printf("response - %#v\n", res)
}
