package archive_deployments_generate_upload_url_test

import (
	sdk "apigee_api/client"
	archive_deployments_generate_upload_url "apigee_api/resources/v1/archive_deployments_generate_upload_url"
	fmt "fmt"
	testing "testing"
)

func TestCreate200SuccessDefault(t *testing.T) {
	// Success test for Default body
	client := sdk.NewClient(sdk.WithOauth2("API_TOKEN"), sdk.WithOauth2c("API_TOKEN"), sdk.WithBaseURL("https://api.sideko.dev/v1/mock/public/apigee/0.1.0"))
	res, err := client.V1.ArchiveDeploymentsGenerateUploadUrl.Create(archive_deployments_generate_upload_url.CreateRequest{Parent: "string", Data: map[string]interface{}{}})

	if err != nil {
		t.Fatalf("TestCreate200SuccessDefault - failed making request with error: %#v", err)
	}

	fmt.Printf("response - %#v\n", res)
}
