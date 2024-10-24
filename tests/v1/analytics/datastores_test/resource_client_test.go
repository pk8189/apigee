package datastores_test_test

import (
	sdk "apigee_api/client"
	nullable "apigee_api/nullable"
	datastores_test "apigee_api/resources/v1/analytics/datastores_test"
	types "apigee_api/types"
	fmt "fmt"
	testing "testing"
)

func TestCreate200SuccessDefault(t *testing.T) {
	// Success test for Default body
	client := sdk.NewClient(sdk.WithOauth2("API_TOKEN"), sdk.WithOauth2c("API_TOKEN"), sdk.WithBaseURL("https://api.sideko.dev/v1/mock/public/apigee/0.1.0"))
	res, err := client.V1.Analytics.DatastoresTest.Create(datastores_test.CreateRequest{Parent: "string", Data: types.GoogleCloudApigeeV1Datastore{CreateTime: nullable.NewValue("string"), DatastoreConfig: nullable.NewValue(types.GoogleCloudApigeeV1DatastoreConfig{BucketName: nullable.NewValue("string"), DatasetName: nullable.NewValue("string"), Path: nullable.NewValue("string"), ProjectId: nullable.NewValue("string"), TablePrefix: nullable.NewValue("string")}), DisplayName: nullable.NewValue("string"), LastUpdateTime: nullable.NewValue("string"), Org: nullable.NewValue("string"), Self: nullable.NewValue("string"), TargetType: nullable.NewValue("string")}})

	if err != nil {
		t.Fatalf("TestCreate200SuccessDefault - failed making request with error: %#v", err)
	}

	fmt.Printf("response - %#v\n", res)
}
