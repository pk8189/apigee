package archive_deployments_test

import (
	sdk "apigee_api/client"
	nullable "apigee_api/nullable"
	archive_deployments "apigee_api/resources/v1/archive_deployments"
	types "apigee_api/types"
	fmt "fmt"
	testing "testing"
)

func TestList200GeneratedSuccess(t *testing.T) {
	// Success case test
	client := sdk.NewClient(sdk.WithOauth2("API_TOKEN"), sdk.WithOauth2c("API_TOKEN"), sdk.WithBaseURL("https://api.sideko.dev/v1/mock/public/apigee/0.1.0"))
	res, err := client.V1.ArchiveDeployments.List(archive_deployments.ListRequest{Parent: "string"})

	if err != nil {
		t.Fatalf("TestList200GeneratedSuccess - failed making request with error: %#v", err)
	}

	fmt.Printf("response - %#v\n", res)
}

func TestCreate200SuccessDefault(t *testing.T) {
	// Success test for Default body
	client := sdk.NewClient(sdk.WithOauth2("API_TOKEN"), sdk.WithOauth2c("API_TOKEN"), sdk.WithBaseURL("https://api.sideko.dev/v1/mock/public/apigee/0.1.0"))
	res, err := client.V1.ArchiveDeployments.Create(archive_deployments.CreateRequest{Parent: "string", Data: types.GoogleCloudApigeeV1ArchiveDeployment{CreatedAt: nullable.NewValue("string"), GcsUri: nullable.NewValue("string"), Labels: nullable.NewValue(types.GoogleCloudApigeeV1ArchiveDeploymentLabels{}), Name: nullable.NewValue("string"), Operation: nullable.NewValue("string"), UpdatedAt: nullable.NewValue("string")}})

	if err != nil {
		t.Fatalf("TestCreate200SuccessDefault - failed making request with error: %#v", err)
	}

	fmt.Printf("response - %#v\n", res)
}
