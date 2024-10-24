package apps_test

import (
	sdk "apigee_api/client"
	nullable "apigee_api/nullable"
	apps "apigee_api/resources/v1/apps"
	types "apigee_api/types"
	fmt "fmt"
	testing "testing"
)

func TestList200GeneratedSuccess(t *testing.T) {
	// Success case test
	client := sdk.NewClient(sdk.WithOauth2("API_TOKEN"), sdk.WithOauth2c("API_TOKEN"), sdk.WithBaseURL("https://api.sideko.dev/v1/mock/public/apigee/0.1.0"))
	res, err := client.V1.Apps.List(apps.ListRequest{Parent: "string"})

	if err != nil {
		t.Fatalf("TestList200GeneratedSuccess - failed making request with error: %#v", err)
	}

	fmt.Printf("response - %#v\n", res)
}

func TestCreate200SuccessDefault(t *testing.T) {
	// Success test for Default body
	client := sdk.NewClient(sdk.WithOauth2("API_TOKEN"), sdk.WithOauth2c("API_TOKEN"), sdk.WithBaseURL("https://api.sideko.dev/v1/mock/public/apigee/0.1.0"))
	res, err := client.V1.Apps.Create(apps.CreateRequest{Parent: "string", Data: types.GoogleCloudApigeeV1DeveloperApp{ApiProducts: nullable.NewValue([]string{"string"}), AppFamily: nullable.NewValue("string"), AppId: nullable.NewValue("string"), Attributes: nullable.NewValue([]types.GoogleCloudApigeeV1Attribute{types.GoogleCloudApigeeV1Attribute{Name: nullable.NewValue("string"), Value: nullable.NewValue("string")}}), CallbackUrl: nullable.NewValue("string"), CreatedAt: nullable.NewValue("string"), Credentials: nullable.NewValue([]types.GoogleCloudApigeeV1Credential{types.GoogleCloudApigeeV1Credential{ApiProducts: nullable.NewValue([]types.GoogleCloudApigeeV1ApiProductRef{types.GoogleCloudApigeeV1ApiProductRef{Apiproduct: nullable.NewValue("string"), Status: nullable.NewValue("string")}}), Attributes: nullable.NewValue([]types.GoogleCloudApigeeV1Attribute{types.GoogleCloudApigeeV1Attribute{Name: nullable.NewValue("string"), Value: nullable.NewValue("string")}}), ConsumerKey: nullable.NewValue("string"), ConsumerSecret: nullable.NewValue("string"), ExpiresAt: nullable.NewValue("string"), IssuedAt: nullable.NewValue("string"), Scopes: nullable.NewValue([]string{"string"}), Status: nullable.NewValue("string")}}), DeveloperId: nullable.NewValue("string"), KeyExpiresIn: nullable.NewValue("string"), LastModifiedAt: nullable.NewValue("string"), Name: nullable.NewValue("string"), Scopes: nullable.NewValue([]string{"string"}), Status: nullable.NewValue("string")}})

	if err != nil {
		t.Fatalf("TestCreate200SuccessDefault - failed making request with error: %#v", err)
	}

	fmt.Printf("response - %#v\n", res)
}
