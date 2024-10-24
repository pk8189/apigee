package debugsessions_test

import (
	sdk "apigee_api/client"
	nullable "apigee_api/nullable"
	debugsessions "apigee_api/resources/v1/debugsessions"
	types "apigee_api/types"
	fmt "fmt"
	testing "testing"
)

func TestList200GeneratedSuccess(t *testing.T) {
	// Success case test
	client := sdk.NewClient(sdk.WithOauth2("API_TOKEN"), sdk.WithOauth2c("API_TOKEN"), sdk.WithBaseURL("https://api.sideko.dev/v1/mock/public/apigee/0.1.0"))
	res, err := client.V1.Debugsessions.List(debugsessions.ListRequest{Parent: "string"})

	if err != nil {
		t.Fatalf("TestList200GeneratedSuccess - failed making request with error: %#v", err)
	}

	fmt.Printf("response - %#v\n", res)
}

func TestCreate200SuccessDefault(t *testing.T) {
	// Success test for Default body
	client := sdk.NewClient(sdk.WithOauth2("API_TOKEN"), sdk.WithOauth2c("API_TOKEN"), sdk.WithBaseURL("https://api.sideko.dev/v1/mock/public/apigee/0.1.0"))
	res, err := client.V1.Debugsessions.Create(debugsessions.CreateRequest{Parent: "string", Data: types.GoogleCloudApigeeV1DebugSession{Count: nullable.NewValue(123), CreateTime: nullable.NewValue("string"), Filter: nullable.NewValue("string"), Name: nullable.NewValue("string"), Timeout: nullable.NewValue("string"), Tracesize: nullable.NewValue(123), Validity: nullable.NewValue(123)}})

	if err != nil {
		t.Fatalf("TestCreate200SuccessDefault - failed making request with error: %#v", err)
	}

	fmt.Printf("response - %#v\n", res)
}
