package targetservers_test

import (
	sdk "apigee_api/client"
	nullable "apigee_api/nullable"
	targetservers "apigee_api/resources/v1/targetservers"
	types "apigee_api/types"
	fmt "fmt"
	testing "testing"
)

func TestCreate200SuccessDefault(t *testing.T) {
	// Success test for Default body
	client := sdk.NewClient(sdk.WithOauth2("API_TOKEN"), sdk.WithOauth2c("API_TOKEN"), sdk.WithBaseURL("https://api.sideko.dev/v1/mock/public/apigee/0.1.0"))
	res, err := client.V1.Targetservers.Create(targetservers.CreateRequest{Parent: "string", Data: types.GoogleCloudApigeeV1TargetServer{Description: nullable.NewValue("string"), Host: nullable.NewValue("string"), IsEnabled: nullable.NewValue(true), Name: nullable.NewValue("string"), Port: nullable.NewValue(123), Protocol: nullable.NewValue(types.GoogleCloudApigeeV1TargetServerProtocolEnumGrpc), SSlInfo: nullable.NewValue(types.GoogleCloudApigeeV1TlsInfo{Ciphers: nullable.NewValue([]string{"string"}), ClientAuthEnabled: nullable.NewValue(true), CommonName: nullable.NewValue(types.GoogleCloudApigeeV1TlsInfoCommonName{Value: nullable.NewValue("string"), WildcardMatch: nullable.NewValue(true)}), Enabled: nullable.NewValue(true), IgnoreValidationErrors: nullable.NewValue(true), KeyAlias: nullable.NewValue("string"), KeyStore: nullable.NewValue("string"), Protocols: nullable.NewValue([]string{"string"}), TrustStore: nullable.NewValue("string")})}})

	if err != nil {
		t.Fatalf("TestCreate200SuccessDefault - failed making request with error: %#v", err)
	}

	fmt.Printf("response - %#v\n", res)
}
