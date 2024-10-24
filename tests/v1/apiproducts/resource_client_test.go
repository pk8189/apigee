package apiproducts_test

import (
	sdk "apigee_api/client"
	nullable "apigee_api/nullable"
	apiproducts "apigee_api/resources/v1/apiproducts"
	types "apigee_api/types"
	fmt "fmt"
	testing "testing"
)

func TestList200GeneratedSuccess(t *testing.T) {
	// Success case test
	client := sdk.NewClient(sdk.WithOauth2("API_TOKEN"), sdk.WithOauth2c("API_TOKEN"), sdk.WithBaseURL("https://api.sideko.dev/v1/mock/public/apigee/0.1.0"))
	res, err := client.V1.Apiproducts.List(apiproducts.ListRequest{Parent: "string"})

	if err != nil {
		t.Fatalf("TestList200GeneratedSuccess - failed making request with error: %#v", err)
	}

	fmt.Printf("response - %#v\n", res)
}

func TestCreate200SuccessDefault(t *testing.T) {
	// Success test for Default body
	client := sdk.NewClient(sdk.WithOauth2("API_TOKEN"), sdk.WithOauth2c("API_TOKEN"), sdk.WithBaseURL("https://api.sideko.dev/v1/mock/public/apigee/0.1.0"))
	res, err := client.V1.Apiproducts.Create(apiproducts.CreateRequest{Parent: "string", Data: types.GoogleCloudApigeeV1ApiProduct{ApiResources: nullable.NewValue([]string{"string"}), ApprovalType: nullable.NewValue("string"), Attributes: nullable.NewValue([]types.GoogleCloudApigeeV1Attribute{types.GoogleCloudApigeeV1Attribute{Name: nullable.NewValue("string"), Value: nullable.NewValue("string")}}), CreatedAt: nullable.NewValue("string"), Description: nullable.NewValue("string"), DisplayName: nullable.NewValue("string"), Environments: nullable.NewValue([]string{"string"}), GraphqlOperationGroup: nullable.NewValue(types.GoogleCloudApigeeV1GraphQlOperationGroup{OperationConfigType: nullable.NewValue("string"), OperationConfigs: nullable.NewValue([]types.GoogleCloudApigeeV1GraphQlOperationConfig{types.GoogleCloudApigeeV1GraphQlOperationConfig{ApiSource: nullable.NewValue("string"), Attributes: nullable.NewValue([]types.GoogleCloudApigeeV1Attribute{types.GoogleCloudApigeeV1Attribute{Name: nullable.NewValue("string"), Value: nullable.NewValue("string")}}), Operations: nullable.NewValue([]types.GoogleCloudApigeeV1GraphQlOperation{types.GoogleCloudApigeeV1GraphQlOperation{Operation: nullable.NewValue("string"), OperationTypes: nullable.NewValue([]string{"string"})}}), Quota: nullable.NewValue(types.GoogleCloudApigeeV1Quota{Interval: nullable.NewValue("string"), Limit: nullable.NewValue("string"), TimeUnit: nullable.NewValue("string")})}})}), LastModifiedAt: nullable.NewValue("string"), Name: nullable.NewValue("string"), OperationGroup: nullable.NewValue(types.GoogleCloudApigeeV1OperationGroup{OperationConfigType: nullable.NewValue("string"), OperationConfigs: nullable.NewValue([]types.GoogleCloudApigeeV1OperationConfig{types.GoogleCloudApigeeV1OperationConfig{ApiSource: nullable.NewValue("string"), Attributes: nullable.NewValue([]types.GoogleCloudApigeeV1Attribute{types.GoogleCloudApigeeV1Attribute{Name: nullable.NewValue("string"), Value: nullable.NewValue("string")}}), Operations: nullable.NewValue([]types.GoogleCloudApigeeV1Operation{types.GoogleCloudApigeeV1Operation{Methods: nullable.NewValue([]string{"string"}), Resource: nullable.NewValue("string")}}), Quota: nullable.NewValue(types.GoogleCloudApigeeV1Quota{Interval: nullable.NewValue("string"), Limit: nullable.NewValue("string"), TimeUnit: nullable.NewValue("string")})}})}), Proxies: nullable.NewValue([]string{"string"}), Quota: nullable.NewValue("string"), QuotaCounterScope: nullable.NewValue(types.GoogleCloudApigeeV1ApiProductQuotaCounterScopeEnumOperation), QuotaInterval: nullable.NewValue("string"), QuotaTimeUnit: nullable.NewValue("string"), Scopes: nullable.NewValue([]string{"string"})}})

	if err != nil {
		t.Fatalf("TestCreate200SuccessDefault - failed making request with error: %#v", err)
	}

	fmt.Printf("response - %#v\n", res)
}
