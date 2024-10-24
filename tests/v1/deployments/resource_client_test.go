package deployments_test

import (
	sdk "apigee_api/client"
	deployments "apigee_api/resources/v1/deployments"
	fmt "fmt"
	testing "testing"
)

func TestDelete200GeneratedSuccess(t *testing.T) {
	// Success case test
	client := sdk.NewClient(sdk.WithOauth2("API_TOKEN"), sdk.WithOauth2c("API_TOKEN"), sdk.WithBaseURL("https://api.sideko.dev/v1/mock/public/apigee/0.1.0"))
	res, err := client.V1.Deployments.Delete(deployments.DeleteRequest{Name: "string"})

	if err != nil {
		t.Fatalf("TestDelete200GeneratedSuccess - failed making request with error: %#v", err)
	}

	fmt.Printf("response - %#v\n", res)
}

func TestCreate200GeneratedSuccess(t *testing.T) {
	// Success case test
	client := sdk.NewClient(sdk.WithOauth2("API_TOKEN"), sdk.WithOauth2c("API_TOKEN"), sdk.WithBaseURL("https://api.sideko.dev/v1/mock/public/apigee/0.1.0"))
	res, err := client.V1.Deployments.Create(deployments.CreateRequest{Name: "string"})

	if err != nil {
		t.Fatalf("TestCreate200GeneratedSuccess - failed making request with error: %#v", err)
	}

	fmt.Printf("response - %#v\n", res)
}
