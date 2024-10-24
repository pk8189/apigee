package data_test

import (
	sdk "apigee_api/client"
	data "apigee_api/resources/v1/data"
	fmt "fmt"
	testing "testing"
)

func TestDelete200GeneratedSuccess(t *testing.T) {
	// Success case test
	client := sdk.NewClient(sdk.WithOauth2("API_TOKEN"), sdk.WithOauth2c("API_TOKEN"), sdk.WithBaseURL("https://api.sideko.dev/v1/mock/public/apigee/0.1.0"))
	res, err := client.V1.Data.Delete(data.DeleteRequest{Name: "string"})

	if err != nil {
		t.Fatalf("TestDelete200GeneratedSuccess - failed making request with error: %#v", err)
	}

	fmt.Printf("response - %#v\n", res)
}
