package security_profiles_test

import (
	sdk "apigee_api/client"
	security_profiles "apigee_api/resources/v1/security_profiles"
	fmt "fmt"
	testing "testing"
)

func TestList200GeneratedSuccess(t *testing.T) {
	// Success case test
	client := sdk.NewClient(sdk.WithOauth2("API_TOKEN"), sdk.WithOauth2c("API_TOKEN"), sdk.WithBaseURL("https://api.sideko.dev/v1/mock/public/apigee/0.1.0"))
	res, err := client.V1.SecurityProfiles.List(security_profiles.ListRequest{Parent: "string"})

	if err != nil {
		t.Fatalf("TestList200GeneratedSuccess - failed making request with error: %#v", err)
	}

	fmt.Printf("response - %#v\n", res)
}
