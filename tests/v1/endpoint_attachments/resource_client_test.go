package endpoint_attachments_test

import (
	sdk "apigee_api/client"
	nullable "apigee_api/nullable"
	endpoint_attachments "apigee_api/resources/v1/endpoint_attachments"
	types "apigee_api/types"
	fmt "fmt"
	testing "testing"
)

func TestList200GeneratedSuccess(t *testing.T) {
	// Success case test
	client := sdk.NewClient(sdk.WithOauth2("API_TOKEN"), sdk.WithOauth2c("API_TOKEN"), sdk.WithBaseURL("https://api.sideko.dev/v1/mock/public/apigee/0.1.0"))
	res, err := client.V1.EndpointAttachments.List(endpoint_attachments.ListRequest{Parent: "string"})

	if err != nil {
		t.Fatalf("TestList200GeneratedSuccess - failed making request with error: %#v", err)
	}

	fmt.Printf("response - %#v\n", res)
}

func TestCreate200SuccessDefault(t *testing.T) {
	// Success test for Default body
	client := sdk.NewClient(sdk.WithOauth2("API_TOKEN"), sdk.WithOauth2c("API_TOKEN"), sdk.WithBaseURL("https://api.sideko.dev/v1/mock/public/apigee/0.1.0"))
	res, err := client.V1.EndpointAttachments.Create(endpoint_attachments.CreateRequest{Parent: "string", Data: types.GoogleCloudApigeeV1EndpointAttachment{ConnectionState: nullable.NewValue(types.GoogleCloudApigeeV1EndpointAttachmentConnectionStateEnumAccepted), Host: nullable.NewValue("string"), Location: nullable.NewValue("string"), Name: nullable.NewValue("string"), ServiceAttachment: nullable.NewValue("string"), State: nullable.NewValue(types.GoogleCloudApigeeV1EndpointAttachmentStateEnumActive)}})

	if err != nil {
		t.Fatalf("TestCreate200SuccessDefault - failed making request with error: %#v", err)
	}

	fmt.Printf("response - %#v\n", res)
}
