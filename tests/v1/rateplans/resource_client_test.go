package rateplans_test

import (
	sdk "apigee_api/client"
	nullable "apigee_api/nullable"
	rateplans "apigee_api/resources/v1/rateplans"
	types "apigee_api/types"
	fmt "fmt"
	testing "testing"
)

func TestList200GeneratedSuccess(t *testing.T) {
	// Success case test
	client := sdk.NewClient(sdk.WithOauth2("API_TOKEN"), sdk.WithOauth2c("API_TOKEN"), sdk.WithBaseURL("https://api.sideko.dev/v1/mock/public/apigee/0.1.0"))
	res, err := client.V1.Rateplans.List(rateplans.ListRequest{Parent: "string"})

	if err != nil {
		t.Fatalf("TestList200GeneratedSuccess - failed making request with error: %#v", err)
	}

	fmt.Printf("response - %#v\n", res)
}

func TestCreate200SuccessDefault(t *testing.T) {
	// Success test for Default body
	client := sdk.NewClient(sdk.WithOauth2("API_TOKEN"), sdk.WithOauth2c("API_TOKEN"), sdk.WithBaseURL("https://api.sideko.dev/v1/mock/public/apigee/0.1.0"))
	res, err := client.V1.Rateplans.Create(rateplans.CreateRequest{Parent: "string", Data: types.GoogleCloudApigeeV1RatePlan{Apiproduct: nullable.NewValue("string"), BillingPeriod: nullable.NewValue(types.GoogleCloudApigeeV1RatePlanBillingPeriodEnumBillingPeriodUnspecified), ConsumptionPricingRates: nullable.NewValue([]types.GoogleCloudApigeeV1RateRange{types.GoogleCloudApigeeV1RateRange{End: nullable.NewValue("string"), Fee: nullable.NewValue(types.GoogleTypeMoney{CurrencyCode: nullable.NewValue("string"), Nanos: nullable.NewValue(123), Units: nullable.NewValue("string")}), Start: nullable.NewValue("string")}}), ConsumptionPricingType: nullable.NewValue(types.GoogleCloudApigeeV1RatePlanConsumptionPricingTypeEnumBanded), CreatedAt: nullable.NewValue("string"), CurrencyCode: nullable.NewValue("string"), Description: nullable.NewValue("string"), DisplayName: nullable.NewValue("string"), EndTime: nullable.NewValue("string"), FixedFeeFrequency: nullable.NewValue(123), FixedRecurringFee: nullable.NewValue(types.GoogleTypeMoney{CurrencyCode: nullable.NewValue("string"), Nanos: nullable.NewValue(123), Units: nullable.NewValue("string")}), LastModifiedAt: nullable.NewValue("string"), Name: nullable.NewValue("string"), PaymentFundingModel: nullable.NewValue(types.GoogleCloudApigeeV1RatePlanPaymentFundingModelEnumPaymentFundingModelUnspecified), RevenueShareRates: nullable.NewValue([]types.GoogleCloudApigeeV1RevenueShareRange{types.GoogleCloudApigeeV1RevenueShareRange{End: nullable.NewValue("string"), SharePercentage: nullable.NewValue(123.45), Start: nullable.NewValue("string")}}), RevenueShareType: nullable.NewValue(types.GoogleCloudApigeeV1RatePlanRevenueShareTypeEnumFixed), SetupFee: nullable.NewValue(types.GoogleTypeMoney{CurrencyCode: nullable.NewValue("string"), Nanos: nullable.NewValue(123), Units: nullable.NewValue("string")}), StartTime: nullable.NewValue("string"), State: nullable.NewValue(types.GoogleCloudApigeeV1RatePlanStateEnumDraft)}})

	if err != nil {
		t.Fatalf("TestCreate200SuccessDefault - failed making request with error: %#v", err)
	}

	fmt.Printf("response - %#v\n", res)
}
