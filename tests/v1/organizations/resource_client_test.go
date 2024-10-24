package organizations_test

import (
	sdk "apigee_api/client"
	nullable "apigee_api/nullable"
	organizations "apigee_api/resources/v1/organizations"
	types "apigee_api/types"
	fmt "fmt"
	testing "testing"
)

func TestCreate200SuccessDefault(t *testing.T) {
	// Success test for Default body
	client := sdk.NewClient(sdk.WithOauth2("API_TOKEN"), sdk.WithOauth2c("API_TOKEN"), sdk.WithBaseURL("https://api.sideko.dev/v1/mock/public/apigee/0.1.0"))
	res, err := client.V1.Organizations.Create(organizations.CreateRequest{Data: types.GoogleCloudApigeeV1Organization{AddonsConfig: nullable.NewValue(types.GoogleCloudApigeeV1AddonsConfig{AdvancedApiOpsConfig: nullable.NewValue(types.GoogleCloudApigeeV1AdvancedApiOpsConfig{Enabled: nullable.NewValue(true)}), ApiSecurityConfig: nullable.NewValue(types.GoogleCloudApigeeV1ApiSecurityConfig{Enabled: nullable.NewValue(true), ExpiresAt: nullable.NewValue("string")}), ConnectorsPlatformConfig: nullable.NewValue(types.GoogleCloudApigeeV1ConnectorsPlatformConfig{Enabled: nullable.NewValue(true), ExpiresAt: nullable.NewValue("string")}), IntegrationConfig: nullable.NewValue(types.GoogleCloudApigeeV1IntegrationConfig{Enabled: nullable.NewValue(true)}), MonetizationConfig: nullable.NewValue(types.GoogleCloudApigeeV1MonetizationConfig{Enabled: nullable.NewValue(true)})}), AnalyticsRegion: nullable.NewValue("string"), ApiConsumerDataEncryptionKeyName: nullable.NewValue("string"), ApiConsumerDataLocation: nullable.NewValue("string"), ApigeeProjectId: nullable.NewValue("string"), Attributes: nullable.NewValue([]string{"string"}), AuthorizedNetwork: nullable.NewValue("string"), BillingType: nullable.NewValue(types.GoogleCloudApigeeV1OrganizationBillingTypeEnumBillingTypeUnspecified), CaCertificate: nullable.NewValue("string"), ControlPlaneEncryptionKeyName: nullable.NewValue("string"), CreatedAt: nullable.NewValue("string"), CustomerName: nullable.NewValue("string"), Description: nullable.NewValue("string"), DisplayName: nullable.NewValue("string"), Environments: nullable.NewValue([]string{"string"}), ExpiresAt: nullable.NewValue("string"), LastModifiedAt: nullable.NewValue("string"), Name: nullable.NewValue("string"), PortalDisabled: nullable.NewValue(true), ProjectId: nullable.NewValue("string"), Properties: nullable.NewValue(types.GoogleCloudApigeeV1Properties{Property: nullable.NewValue([]types.GoogleCloudApigeeV1Property{types.GoogleCloudApigeeV1Property{Name: nullable.NewValue("string"), Value: nullable.NewValue("string")}})}), RuntimeDatabaseEncryptionKeyName: nullable.NewValue("string"), RuntimeType: nullable.NewValue(types.GoogleCloudApigeeV1OrganizationRuntimeTypeEnumCloud), State: nullable.NewValue(types.GoogleCloudApigeeV1OrganizationStateEnumActive), SubscriptionType: nullable.NewValue(types.GoogleCloudApigeeV1OrganizationSubscriptionTypeEnumPaid), Type: nullable.NewValue(types.GoogleCloudApigeeV1OrganizationTypeEnumTypeInternal)}})

	if err != nil {
		t.Fatalf("TestCreate200SuccessDefault - failed making request with error: %#v", err)
	}

	fmt.Printf("response - %#v\n", res)
}
