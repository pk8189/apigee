package attributes_test

import (
	sdk "apigee_api/client"
	nullable "apigee_api/nullable"
	attributes "apigee_api/resources/v1/attributes"
	types "apigee_api/types"
	fmt "fmt"
	testing "testing"
)

func TestList200GeneratedSuccess(t *testing.T) {
	// Success case test
	client := sdk.NewClient(sdk.WithOauth2("API_TOKEN"), sdk.WithOauth2c("API_TOKEN"), sdk.WithBaseURL("https://api.sideko.dev/v1/mock/public/apigee/0.1.0"))
	res, err := client.V1.Attributes.List(attributes.ListRequest{Parent: "string"})

	if err != nil {
		t.Fatalf("TestList200GeneratedSuccess - failed making request with error: %#v", err)
	}

	fmt.Printf("response - %#v\n", res)
}

func TestCreate200SuccessDefault(t *testing.T) {
	// Success test for Default body
	client := sdk.NewClient(sdk.WithOauth2("API_TOKEN"), sdk.WithOauth2c("API_TOKEN"), sdk.WithBaseURL("https://api.sideko.dev/v1/mock/public/apigee/0.1.0"))
	res, err := client.V1.Attributes.Create(attributes.CreateRequest{Name: "string", Data: types.GoogleCloudApigeeV1Attributes{Attribute: nullable.NewValue([]types.GoogleCloudApigeeV1Attribute{types.GoogleCloudApigeeV1Attribute{Name: nullable.NewValue("string"), Value: nullable.NewValue("string")}})}})

	if err != nil {
		t.Fatalf("TestCreate200SuccessDefault - failed making request with error: %#v", err)
	}

	fmt.Printf("response - %#v\n", res)
}

func TestCreate0200SuccessDefault(t *testing.T) {
	// Success test for Default body
	client := sdk.NewClient(sdk.WithOauth2("API_TOKEN"), sdk.WithOauth2c("API_TOKEN"), sdk.WithBaseURL("https://api.sideko.dev/v1/mock/public/apigee/0.1.0"))
	res, err := client.V1.Attributes.Create0(attributes.Create0Request{Parent: "string", Data: types.GoogleCloudApigeeV1Attributes{Attribute: nullable.NewValue([]types.GoogleCloudApigeeV1Attribute{types.GoogleCloudApigeeV1Attribute{Name: nullable.NewValue("string"), Value: nullable.NewValue("string")}})}})

	if err != nil {
		t.Fatalf("TestCreate0200SuccessDefault - failed making request with error: %#v", err)
	}

	fmt.Printf("response - %#v\n", res)
}
