package references_test

import (
	sdk "apigee_api/client"
	nullable "apigee_api/nullable"
	references "apigee_api/resources/v1/references"
	types "apigee_api/types"
	fmt "fmt"
	testing "testing"
)

func TestCreate200SuccessDefault(t *testing.T) {
	// Success test for Default body
	client := sdk.NewClient(sdk.WithOauth2("API_TOKEN"), sdk.WithOauth2c("API_TOKEN"), sdk.WithBaseURL("https://api.sideko.dev/v1/mock/public/apigee/0.1.0"))
	res, err := client.V1.References.Create(references.CreateRequest{Parent: "string", Data: types.GoogleCloudApigeeV1Reference{Description: nullable.NewValue("string"), Name: nullable.NewValue("string"), Refers: nullable.NewValue("string"), ResourceType: nullable.NewValue("string")}})

	if err != nil {
		t.Fatalf("TestCreate200SuccessDefault - failed making request with error: %#v", err)
	}

	fmt.Printf("response - %#v\n", res)
}
