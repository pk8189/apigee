package keyvaluemaps_test

import (
	sdk "apigee_api/client"
	nullable "apigee_api/nullable"
	keyvaluemaps "apigee_api/resources/v1/keyvaluemaps"
	types "apigee_api/types"
	fmt "fmt"
	testing "testing"
)

func TestCreate200SuccessDefault(t *testing.T) {
	// Success test for Default body
	client := sdk.NewClient(sdk.WithOauth2("API_TOKEN"), sdk.WithOauth2c("API_TOKEN"), sdk.WithBaseURL("https://api.sideko.dev/v1/mock/public/apigee/0.1.0"))
	res, err := client.V1.Keyvaluemaps.Create(keyvaluemaps.CreateRequest{Parent: "string", Data: types.GoogleCloudApigeeV1KeyValueMap{Encrypted: nullable.NewValue(true), Name: nullable.NewValue("string")}})

	if err != nil {
		t.Fatalf("TestCreate200SuccessDefault - failed making request with error: %#v", err)
	}

	fmt.Printf("response - %#v\n", res)
}
