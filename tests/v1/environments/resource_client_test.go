package environments_test

import (
	sdk "apigee_api/client"
	nullable "apigee_api/nullable"
	environments "apigee_api/resources/v1/environments"
	types "apigee_api/types"
	fmt "fmt"
	testing "testing"
)

func TestCreate200SuccessDefault(t *testing.T) {
	// Success test for Default body
	client := sdk.NewClient(sdk.WithOauth2("API_TOKEN"), sdk.WithOauth2c("API_TOKEN"), sdk.WithBaseURL("https://api.sideko.dev/v1/mock/public/apigee/0.1.0"))
	res, err := client.V1.Environments.Create(environments.CreateRequest{Parent: "string", Data: types.GoogleCloudApigeeV1SecurityProfileEnvironmentAssociation{AttachTime: nullable.NewValue("string"), Name: nullable.NewValue("string"), SecurityProfileRevisionId: nullable.NewValue("string")}})

	if err != nil {
		t.Fatalf("TestCreate200SuccessDefault - failed making request with error: %#v", err)
	}

	fmt.Printf("response - %#v\n", res)
}
