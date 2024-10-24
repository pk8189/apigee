package exports_test

import (
	sdk "apigee_api/client"
	nullable "apigee_api/nullable"
	exports "apigee_api/resources/v1/analytics/exports"
	types "apigee_api/types"
	fmt "fmt"
	testing "testing"
)

func TestList200GeneratedSuccess(t *testing.T) {
	// Success case test
	client := sdk.NewClient(sdk.WithOauth2("API_TOKEN"), sdk.WithOauth2c("API_TOKEN"), sdk.WithBaseURL("https://api.sideko.dev/v1/mock/public/apigee/0.1.0"))
	res, err := client.V1.Analytics.Exports.List(exports.ListRequest{Parent: "string"})

	if err != nil {
		t.Fatalf("TestList200GeneratedSuccess - failed making request with error: %#v", err)
	}

	fmt.Printf("response - %#v\n", res)
}

func TestCreate200SuccessDefault(t *testing.T) {
	// Success test for Default body
	client := sdk.NewClient(sdk.WithOauth2("API_TOKEN"), sdk.WithOauth2c("API_TOKEN"), sdk.WithBaseURL("https://api.sideko.dev/v1/mock/public/apigee/0.1.0"))
	res, err := client.V1.Analytics.Exports.Create(exports.CreateRequest{Parent: "string", Data: types.GoogleCloudApigeeV1ExportRequest{CsvDelimiter: nullable.NewValue("string"), DatastoreName: nullable.NewValue("string"), DateRange: nullable.NewValue(types.GoogleCloudApigeeV1DateRange{End: nullable.NewValue("string"), Start: nullable.NewValue("string")}), Description: nullable.NewValue("string"), Name: nullable.NewValue("string"), OutputFormat: nullable.NewValue("string")}})

	if err != nil {
		t.Fatalf("TestCreate200SuccessDefault - failed making request with error: %#v", err)
	}

	fmt.Printf("response - %#v\n", res)
}
