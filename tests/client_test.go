package client_test

import (
	sdk "apigee_api/client"
	fmt "fmt"
	testing "testing"
)

func TestClientInit(t *testing.T) {
	client := sdk.NewClient(sdk.WithOauth2("API_TOKEN"), sdk.WithOauth2c("API_TOKEN"), sdk.WithBaseURL("https://api.sideko.dev/v1/mock/public/apigee/0.1.0"))
	fmt.Printf("%#v\n", client)
}
