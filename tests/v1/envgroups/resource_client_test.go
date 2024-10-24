package envgroups_test

import (
	sdk "apigee_api/client"
	nullable "apigee_api/nullable"
	envgroups "apigee_api/resources/v1/envgroups"
	types "apigee_api/types"
	fmt "fmt"
	testing "testing"
)

func TestList200GeneratedSuccess(t *testing.T) {
	// Success case test
	client := sdk.NewClient(sdk.WithOauth2("API_TOKEN"), sdk.WithOauth2c("API_TOKEN"), sdk.WithBaseURL("https://api.sideko.dev/v1/mock/public/apigee/0.1.0"))
	res, err := client.V1.Envgroups.List(envgroups.ListRequest{Parent: "string"})

	if err != nil {
		t.Fatalf("TestList200GeneratedSuccess - failed making request with error: %#v", err)
	}

	fmt.Printf("response - %#v\n", res)
}

func TestCreate200SuccessDefault(t *testing.T) {
	// Success test for Default body
	client := sdk.NewClient(sdk.WithOauth2("API_TOKEN"), sdk.WithOauth2c("API_TOKEN"), sdk.WithBaseURL("https://api.sideko.dev/v1/mock/public/apigee/0.1.0"))
	res, err := client.V1.Envgroups.Create(envgroups.CreateRequest{Parent: "string", Data: types.GoogleCloudApigeeV1EnvironmentGroup{CreatedAt: nullable.NewValue("string"), Hostnames: nullable.NewValue([]string{"string"}), LastModifiedAt: nullable.NewValue("string"), Name: nullable.NewValue("string"), State: nullable.NewValue(types.GoogleCloudApigeeV1EnvironmentGroupStateEnumActive)}})

	if err != nil {
		t.Fatalf("TestCreate200SuccessDefault - failed making request with error: %#v", err)
	}

	fmt.Printf("response - %#v\n", res)
}
