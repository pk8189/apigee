package set_addons_test

import (
	sdk "apigee_api/client"
	nullable "apigee_api/nullable"
	set_addons "apigee_api/resources/v1/set_addons"
	types "apigee_api/types"
	fmt "fmt"
	testing "testing"
)

func TestCreate200SuccessDefault(t *testing.T) {
	// Success test for Default body
	client := sdk.NewClient(sdk.WithOauth2("API_TOKEN"), sdk.WithOauth2c("API_TOKEN"), sdk.WithBaseURL("https://api.sideko.dev/v1/mock/public/apigee/0.1.0"))
	res, err := client.V1.SetAddons.Create(set_addons.CreateRequest{Org: "string", Data: types.GoogleCloudApigeeV1SetAddonsRequest{AddonsConfig: nullable.NewValue(types.GoogleCloudApigeeV1AddonsConfig{AdvancedApiOpsConfig: nullable.NewValue(types.GoogleCloudApigeeV1AdvancedApiOpsConfig{Enabled: nullable.NewValue(true)}), ApiSecurityConfig: nullable.NewValue(types.GoogleCloudApigeeV1ApiSecurityConfig{Enabled: nullable.NewValue(true), ExpiresAt: nullable.NewValue("string")}), ConnectorsPlatformConfig: nullable.NewValue(types.GoogleCloudApigeeV1ConnectorsPlatformConfig{Enabled: nullable.NewValue(true), ExpiresAt: nullable.NewValue("string")}), IntegrationConfig: nullable.NewValue(types.GoogleCloudApigeeV1IntegrationConfig{Enabled: nullable.NewValue(true)}), MonetizationConfig: nullable.NewValue(types.GoogleCloudApigeeV1MonetizationConfig{Enabled: nullable.NewValue(true)})})}})

	if err != nil {
		t.Fatalf("TestCreate200SuccessDefault - failed making request with error: %#v", err)
	}

	fmt.Printf("response - %#v\n", res)
}
