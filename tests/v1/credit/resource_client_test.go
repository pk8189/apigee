package credit_test

import (
	sdk "apigee_api/client"
	nullable "apigee_api/nullable"
	credit "apigee_api/resources/v1/credit"
	types "apigee_api/types"
	fmt "fmt"
	testing "testing"
)

func TestCreate200SuccessDefault(t *testing.T) {
	// Success test for Default body
	client := sdk.NewClient(sdk.WithOauth2("API_TOKEN"), sdk.WithOauth2c("API_TOKEN"), sdk.WithBaseURL("https://api.sideko.dev/v1/mock/public/apigee/0.1.0"))
	res, err := client.V1.Credit.Create(credit.CreateRequest{Name: "string", Data: types.GoogleCloudApigeeV1CreditDeveloperBalanceRequest{TransactionAmount: nullable.NewValue(types.GoogleTypeMoney{CurrencyCode: nullable.NewValue("string"), Nanos: nullable.NewValue(123), Units: nullable.NewValue("string")}), TransactionId: nullable.NewValue("string")}})

	if err != nil {
		t.Fatalf("TestCreate200SuccessDefault - failed making request with error: %#v", err)
	}

	fmt.Printf("response - %#v\n", res)
}
