package developers_test

import (
	sdk "apigee_api/client"
	nullable "apigee_api/nullable"
	developers "apigee_api/resources/v1/developers"
	types "apigee_api/types"
	fmt "fmt"
	testing "testing"
)

func TestList200GeneratedSuccess(t *testing.T) {
	// Success case test
	client := sdk.NewClient(sdk.WithOauth2("API_TOKEN"), sdk.WithOauth2c("API_TOKEN"), sdk.WithBaseURL("https://api.sideko.dev/v1/mock/public/apigee/0.1.0"))
	res, err := client.V1.Developers.List(developers.ListRequest{Parent: "string"})

	if err != nil {
		t.Fatalf("TestList200GeneratedSuccess - failed making request with error: %#v", err)
	}

	fmt.Printf("response - %#v\n", res)
}

func TestCreate200SuccessDefault(t *testing.T) {
	// Success test for Default body
	client := sdk.NewClient(sdk.WithOauth2("API_TOKEN"), sdk.WithOauth2c("API_TOKEN"), sdk.WithBaseURL("https://api.sideko.dev/v1/mock/public/apigee/0.1.0"))
	res, err := client.V1.Developers.Create(developers.CreateRequest{Parent: "string", Data: types.GoogleCloudApigeeV1Developer{AccessType: nullable.NewValue("string"), AppFamily: nullable.NewValue("string"), Apps: nullable.NewValue([]string{"string"}), Attributes: nullable.NewValue([]types.GoogleCloudApigeeV1Attribute{types.GoogleCloudApigeeV1Attribute{Name: nullable.NewValue("string"), Value: nullable.NewValue("string")}}), Companies: nullable.NewValue([]string{"string"}), CreatedAt: nullable.NewValue("string"), DeveloperId: nullable.NewValue("string"), Email: nullable.NewValue("string"), FirstName: nullable.NewValue("string"), LastModifiedAt: nullable.NewValue("string"), LastName: nullable.NewValue("string"), OrganizationName: nullable.NewValue("string"), Status: nullable.NewValue("string"), UserName: nullable.NewValue("string")}})

	if err != nil {
		t.Fatalf("TestCreate200SuccessDefault - failed making request with error: %#v", err)
	}

	fmt.Printf("response - %#v\n", res)
}
