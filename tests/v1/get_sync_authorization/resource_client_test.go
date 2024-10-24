package get_sync_authorization_test

import (
	sdk "apigee_api/client"
	get_sync_authorization "apigee_api/resources/v1/get_sync_authorization"
	fmt "fmt"
	testing "testing"
)

func TestCreate200SuccessDefault(t *testing.T) {
	// Success test for Default body
	client := sdk.NewClient(sdk.WithOauth2("API_TOKEN"), sdk.WithOauth2c("API_TOKEN"), sdk.WithBaseURL("https://api.sideko.dev/v1/mock/public/apigee/0.1.0"))
	res, err := client.V1.GetSyncAuthorization.Create(get_sync_authorization.CreateRequest{Name: "string", Data: map[string]interface{}{}})

	if err != nil {
		t.Fatalf("TestCreate200SuccessDefault - failed making request with error: %#v", err)
	}

	fmt.Printf("response - %#v\n", res)
}
