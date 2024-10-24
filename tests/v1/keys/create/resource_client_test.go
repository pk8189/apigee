package create_test

import (
	sdk "apigee_api/client"
	nullable "apigee_api/nullable"
	create "apigee_api/resources/v1/keys/create"
	types "apigee_api/types"
	fmt "fmt"
	testing "testing"
)

func TestCreate200SuccessDefault(t *testing.T) {
	// Success test for Default body
	client := sdk.NewClient(sdk.WithOauth2("API_TOKEN"), sdk.WithOauth2c("API_TOKEN"), sdk.WithBaseURL("https://api.sideko.dev/v1/mock/public/apigee/0.1.0"))
	res, err := client.V1.Keys.Create.Create(create.CreateRequest{Parent: "string", Data: types.GoogleCloudApigeeV1DeveloperAppKey{ApiProducts: nullable.NewValue([]interface{}{"could be anything"}), Attributes: nullable.NewValue([]types.GoogleCloudApigeeV1Attribute{types.GoogleCloudApigeeV1Attribute{Name: nullable.NewValue("string"), Value: nullable.NewValue("string")}}), ConsumerKey: nullable.NewValue("string"), ConsumerSecret: nullable.NewValue("string"), ExpiresAt: nullable.NewValue("string"), ExpiresInSeconds: nullable.NewValue("string"), IssuedAt: nullable.NewValue("string"), Scopes: nullable.NewValue([]string{"string"}), Status: nullable.NewValue("string")}})

	if err != nil {
		t.Fatalf("TestCreate200SuccessDefault - failed making request with error: %#v", err)
	}

	fmt.Printf("response - %#v\n", res)
}
