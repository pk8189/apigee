package subscribe_test

import (
	sdk "apigee_api/client"
	subscribe "apigee_api/resources/v1/subscribe"
	fmt "fmt"
	testing "testing"
)

func TestCreate200GeneratedSuccess(t *testing.T) {
	// Success case test
	client := sdk.NewClient(sdk.WithOauth2("API_TOKEN"), sdk.WithOauth2c("API_TOKEN"), sdk.WithBaseURL("https://api.sideko.dev/v1/mock/public/apigee/0.1.0"))
	res, err := client.V1.Subscribe.Create(subscribe.CreateRequest{Parent: "string"})

	if err != nil {
		t.Fatalf("TestCreate200GeneratedSuccess - failed making request with error: %#v", err)
	}

	fmt.Printf("response - %#v\n", res)
}
