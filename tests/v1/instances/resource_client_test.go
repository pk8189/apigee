package instances_test

import (
	sdk "apigee_api/client"
	nullable "apigee_api/nullable"
	instances "apigee_api/resources/v1/instances"
	types "apigee_api/types"
	fmt "fmt"
	testing "testing"
)

func TestList200GeneratedSuccess(t *testing.T) {
	// Success case test
	client := sdk.NewClient(sdk.WithOauth2("API_TOKEN"), sdk.WithOauth2c("API_TOKEN"), sdk.WithBaseURL("https://api.sideko.dev/v1/mock/public/apigee/0.1.0"))
	res, err := client.V1.Instances.List(instances.ListRequest{Parent: "string"})

	if err != nil {
		t.Fatalf("TestList200GeneratedSuccess - failed making request with error: %#v", err)
	}

	fmt.Printf("response - %#v\n", res)
}

func TestCreate200SuccessDefault(t *testing.T) {
	// Success test for Default body
	client := sdk.NewClient(sdk.WithOauth2("API_TOKEN"), sdk.WithOauth2c("API_TOKEN"), sdk.WithBaseURL("https://api.sideko.dev/v1/mock/public/apigee/0.1.0"))
	res, err := client.V1.Instances.Create(instances.CreateRequest{Parent: "string", Data: types.GoogleCloudApigeeV1Instance{ConsumerAcceptList: nullable.NewValue([]string{"string"}), CreatedAt: nullable.NewValue("string"), Description: nullable.NewValue("string"), DiskEncryptionKeyName: nullable.NewValue("string"), DisplayName: nullable.NewValue("string"), Host: nullable.NewValue("string"), IpRange: nullable.NewValue("string"), LastModifiedAt: nullable.NewValue("string"), Location: nullable.NewValue("string"), Name: nullable.NewValue("string"), PeeringCidrRange: nullable.NewValue(types.GoogleCloudApigeeV1InstancePeeringCidrRangeEnumCidrRangeUnspecified), Port: nullable.NewValue("string"), RuntimeVersion: nullable.NewValue("string"), ServiceAttachment: nullable.NewValue("string"), State: nullable.NewValue(types.GoogleCloudApigeeV1InstanceStateEnumActive)}})

	if err != nil {
		t.Fatalf("TestCreate200SuccessDefault - failed making request with error: %#v", err)
	}

	fmt.Printf("response - %#v\n", res)
}
