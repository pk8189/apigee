package x_test

import (
	sdk "apigee_api/client"
	nullable "apigee_api/nullable"
	x "apigee_api/resources/v1/resourcefiles/x"
	types "apigee_api/types"
	fmt "fmt"
	testing "testing"
)

func TestDelete200GeneratedSuccess(t *testing.T) {
	// Success case test
	client := sdk.NewClient(sdk.WithOauth2("API_TOKEN"), sdk.WithOauth2c("API_TOKEN"), sdk.WithBaseURL("https://api.sideko.dev/v1/mock/public/apigee/0.1.0"))
	res, err := client.V1.Resourcefiles.X.Delete(x.DeleteRequest{Parent: "string", Type: "string", Name: "string"})

	if err != nil {
		t.Fatalf("TestDelete200GeneratedSuccess - failed making request with error: %#v", err)
	}

	fmt.Printf("response - %#v\n", res)
}

func TestGet200GeneratedSuccess(t *testing.T) {
	// Success case test
	client := sdk.NewClient(sdk.WithOauth2("API_TOKEN"), sdk.WithOauth2c("API_TOKEN"), sdk.WithBaseURL("https://api.sideko.dev/v1/mock/public/apigee/0.1.0"))
	res, err := client.V1.Resourcefiles.X.Get(x.GetRequest{Parent: "string", Type: "string", Name: "string"})

	if err != nil {
		t.Fatalf("TestGet200GeneratedSuccess - failed making request with error: %#v", err)
	}

	fmt.Printf("response - %#v\n", res)
}

func TestPut200SuccessDefault(t *testing.T) {
	// Success test for Default body
	client := sdk.NewClient(sdk.WithOauth2("API_TOKEN"), sdk.WithOauth2c("API_TOKEN"), sdk.WithBaseURL("https://api.sideko.dev/v1/mock/public/apigee/0.1.0"))
	res, err := client.V1.Resourcefiles.X.Put(x.PutRequest{Parent: "string", Type: "string", Name: "string", Data: types.GoogleApiHttpBody{ContentType: nullable.NewValue("string"), Data: nullable.NewValue("string"), Extensions: nullable.NewValue([]map[string]interface{}{map[string]interface{}{}})}})

	if err != nil {
		t.Fatalf("TestPut200SuccessDefault - failed making request with error: %#v", err)
	}

	fmt.Printf("response - %#v\n", res)
}
