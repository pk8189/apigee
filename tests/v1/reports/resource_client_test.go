package reports_test

import (
	sdk "apigee_api/client"
	nullable "apigee_api/nullable"
	reports "apigee_api/resources/v1/reports"
	types "apigee_api/types"
	fmt "fmt"
	testing "testing"
)

func TestList200GeneratedSuccess(t *testing.T) {
	// Success case test
	client := sdk.NewClient(sdk.WithOauth2("API_TOKEN"), sdk.WithOauth2c("API_TOKEN"), sdk.WithBaseURL("https://api.sideko.dev/v1/mock/public/apigee/0.1.0"))
	res, err := client.V1.Reports.List(reports.ListRequest{Parent: "string"})

	if err != nil {
		t.Fatalf("TestList200GeneratedSuccess - failed making request with error: %#v", err)
	}

	fmt.Printf("response - %#v\n", res)
}

func TestCreate200SuccessDefault(t *testing.T) {
	// Success test for Default body
	client := sdk.NewClient(sdk.WithOauth2("API_TOKEN"), sdk.WithOauth2c("API_TOKEN"), sdk.WithBaseURL("https://api.sideko.dev/v1/mock/public/apigee/0.1.0"))
	res, err := client.V1.Reports.Create(reports.CreateRequest{Parent: "string", Data: types.GoogleCloudApigeeV1CustomReport{ChartType: nullable.NewValue("string"), Comments: nullable.NewValue([]string{"string"}), CreatedAt: nullable.NewValue("string"), Dimensions: nullable.NewValue([]string{"string"}), DisplayName: nullable.NewValue("string"), Environment: nullable.NewValue("string"), Filter: nullable.NewValue("string"), FromTime: nullable.NewValue("string"), LastModifiedAt: nullable.NewValue("string"), LastViewedAt: nullable.NewValue("string"), Limit: nullable.NewValue("string"), Metrics: nullable.NewValue([]types.GoogleCloudApigeeV1CustomReportMetric{types.GoogleCloudApigeeV1CustomReportMetric{Function: nullable.NewValue("string"), Name: nullable.NewValue("string")}}), Name: nullable.NewValue("string"), Offset: nullable.NewValue("string"), Organization: nullable.NewValue("string"), Properties: nullable.NewValue([]types.GoogleCloudApigeeV1ReportProperty{types.GoogleCloudApigeeV1ReportProperty{Property: nullable.NewValue("string"), Value: nullable.NewValue([]types.GoogleCloudApigeeV1Attribute{types.GoogleCloudApigeeV1Attribute{Name: nullable.NewValue("string"), Value: nullable.NewValue("string")}})}}), SortByCols: nullable.NewValue([]string{"string"}), SortOrder: nullable.NewValue("string"), Tags: nullable.NewValue([]string{"string"}), TimeUnit: nullable.NewValue("string"), ToTime: nullable.NewValue("string"), Topk: nullable.NewValue("string")}})

	if err != nil {
		t.Fatalf("TestCreate200SuccessDefault - failed making request with error: %#v", err)
	}

	fmt.Printf("response - %#v\n", res)
}
