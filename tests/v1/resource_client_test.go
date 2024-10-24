package v1_test

import (
	sdk "apigee_api/client"
	nullable "apigee_api/nullable"
	v1 "apigee_api/resources/v1"
	types "apigee_api/types"
	fmt "fmt"
	testing "testing"
)

func TestDelete200GeneratedSuccess(t *testing.T) {
	// Success case test
	client := sdk.NewClient(sdk.WithOauth2("API_TOKEN"), sdk.WithOauth2c("API_TOKEN"), sdk.WithBaseURL("https://api.sideko.dev/v1/mock/public/apigee/0.1.0"))
	res, err := client.V1.Delete(v1.DeleteRequest{Name: "string"})

	if err != nil {
		t.Fatalf("TestDelete200GeneratedSuccess - failed making request with error: %#v", err)
	}

	fmt.Printf("response - %#v\n", res)
}

func TestGetByName200GeneratedSuccess(t *testing.T) {
	// Success case test
	client := sdk.NewClient(sdk.WithOauth2("API_TOKEN"), sdk.WithOauth2c("API_TOKEN"), sdk.WithBaseURL("https://api.sideko.dev/v1/mock/public/apigee/0.1.0"))
	res, err := client.V1.GetByName(v1.GetByNameRequest{Name: "string"})

	if err != nil {
		t.Fatalf("TestGetByName200GeneratedSuccess - failed making request with error: %#v", err)
	}

	fmt.Printf("response - %#v\n", res)
}

func TestGetByParent200GeneratedSuccess(t *testing.T) {
	// Success case test
	client := sdk.NewClient(sdk.WithOauth2("API_TOKEN"), sdk.WithOauth2c("API_TOKEN"), sdk.WithBaseURL("https://api.sideko.dev/v1/mock/public/apigee/0.1.0"))
	res, err := client.V1.GetByParent(v1.GetByParentRequest{Parent: "string"})

	if err != nil {
		t.Fatalf("TestGetByParent200GeneratedSuccess - failed making request with error: %#v", err)
	}

	fmt.Printf("response - %#v\n", res)
}

func TestPatch200SuccessDefault(t *testing.T) {
	// Success test for Default body
	client := sdk.NewClient(sdk.WithOauth2("API_TOKEN"), sdk.WithOauth2c("API_TOKEN"), sdk.WithBaseURL("https://api.sideko.dev/v1/mock/public/apigee/0.1.0"))
	res, err := client.V1.Patch(v1.PatchRequest{Name: "string", Data: types.GoogleCloudApigeeV1ApiCategoryData{Id: nullable.NewValue("string"), Name: nullable.NewValue("string"), SiteId: nullable.NewValue("string"), UpdateTime: nullable.NewValue("string")}})

	if err != nil {
		t.Fatalf("TestPatch200SuccessDefault - failed making request with error: %#v", err)
	}

	fmt.Printf("response - %#v\n", res)
}

func TestCreate200SuccessDefault(t *testing.T) {
	// Success test for Default body
	client := sdk.NewClient(sdk.WithOauth2("API_TOKEN"), sdk.WithOauth2c("API_TOKEN"), sdk.WithBaseURL("https://api.sideko.dev/v1/mock/public/apigee/0.1.0"))
	res, err := client.V1.Create(v1.CreateRequest{Name: "string", Data: types.GoogleApiHttpBody{ContentType: nullable.NewValue("string"), Data: nullable.NewValue("string"), Extensions: nullable.NewValue([]map[string]interface{}{map[string]interface{}{}})}})

	if err != nil {
		t.Fatalf("TestCreate200SuccessDefault - failed making request with error: %#v", err)
	}

	fmt.Printf("response - %#v\n", res)
}

func TestPut200SuccessDefault(t *testing.T) {
	// Success test for Default body
	client := sdk.NewClient(sdk.WithOauth2("API_TOKEN"), sdk.WithOauth2c("API_TOKEN"), sdk.WithBaseURL("https://api.sideko.dev/v1/mock/public/apigee/0.1.0"))
	res, err := client.V1.Put(v1.PutRequest{Name: "string", Data: types.GoogleCloudApigeeV1CustomReport{ChartType: nullable.NewValue("string"), Comments: nullable.NewValue([]string{"string"}), CreatedAt: nullable.NewValue("string"), Dimensions: nullable.NewValue([]string{"string"}), DisplayName: nullable.NewValue("string"), Environment: nullable.NewValue("string"), Filter: nullable.NewValue("string"), FromTime: nullable.NewValue("string"), LastModifiedAt: nullable.NewValue("string"), LastViewedAt: nullable.NewValue("string"), Limit: nullable.NewValue("string"), Metrics: nullable.NewValue([]types.GoogleCloudApigeeV1CustomReportMetric{types.GoogleCloudApigeeV1CustomReportMetric{Function: nullable.NewValue("string"), Name: nullable.NewValue("string")}}), Name: nullable.NewValue("string"), Offset: nullable.NewValue("string"), Organization: nullable.NewValue("string"), Properties: nullable.NewValue([]types.GoogleCloudApigeeV1ReportProperty{types.GoogleCloudApigeeV1ReportProperty{Property: nullable.NewValue("string"), Value: nullable.NewValue([]types.GoogleCloudApigeeV1Attribute{types.GoogleCloudApigeeV1Attribute{Name: nullable.NewValue("string"), Value: nullable.NewValue("string")}})}}), SortByCols: nullable.NewValue([]string{"string"}), SortOrder: nullable.NewValue("string"), Tags: nullable.NewValue([]string{"string"}), TimeUnit: nullable.NewValue("string"), ToTime: nullable.NewValue("string"), Topk: nullable.NewValue("string")}})

	if err != nil {
		t.Fatalf("TestPut200SuccessDefault - failed making request with error: %#v", err)
	}

	fmt.Printf("response - %#v\n", res)
}
