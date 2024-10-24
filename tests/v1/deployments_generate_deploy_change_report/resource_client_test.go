package deployments_generate_deploy_change_report_test

import (
	sdk "apigee_api/client"
	deployments_generate_deploy_change_report "apigee_api/resources/v1/deployments_generate_deploy_change_report"
	fmt "fmt"
	testing "testing"
)

func TestCreate200GeneratedSuccess(t *testing.T) {
	// Success case test
	client := sdk.NewClient(sdk.WithOauth2("API_TOKEN"), sdk.WithOauth2c("API_TOKEN"), sdk.WithBaseURL("https://api.sideko.dev/v1/mock/public/apigee/0.1.0"))
	res, err := client.V1.DeploymentsGenerateDeployChangeReport.Create(deployments_generate_deploy_change_report.CreateRequest{Name: "string"})

	if err != nil {
		t.Fatalf("TestCreate200GeneratedSuccess - failed making request with error: %#v", err)
	}

	fmt.Printf("response - %#v\n", res)
}
