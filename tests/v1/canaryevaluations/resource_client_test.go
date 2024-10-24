package canaryevaluations_test

import (
	sdk "apigee_api/client"
	nullable "apigee_api/nullable"
	canaryevaluations "apigee_api/resources/v1/canaryevaluations"
	types "apigee_api/types"
	fmt "fmt"
	testing "testing"
)

func TestCreate200SuccessDefault(t *testing.T) {
	// Success test for Default body
	client := sdk.NewClient(sdk.WithOauth2("API_TOKEN"), sdk.WithOauth2c("API_TOKEN"), sdk.WithBaseURL("https://api.sideko.dev/v1/mock/public/apigee/0.1.0"))
	res, err := client.V1.Canaryevaluations.Create(canaryevaluations.CreateRequest{Parent: "string", Data: types.GoogleCloudApigeeV1CanaryEvaluation{Control: nullable.NewValue("string"), CreateTime: nullable.NewValue("string"), EndTime: nullable.NewValue("string"), MetricLabels: nullable.NewValue(types.GoogleCloudApigeeV1CanaryEvaluationMetricLabels{Env: nullable.NewValue("string"), InstanceId: nullable.NewValue("string"), Location: nullable.NewValue("string")}), Name: nullable.NewValue("string"), StartTime: nullable.NewValue("string"), State: nullable.NewValue(types.GoogleCloudApigeeV1CanaryEvaluationStateEnumRunning), Treatment: nullable.NewValue("string"), Verdict: nullable.NewValue(types.GoogleCloudApigeeV1CanaryEvaluationVerdictEnumFail)}})

	if err != nil {
		t.Fatalf("TestCreate200SuccessDefault - failed making request with error: %#v", err)
	}

	fmt.Printf("response - %#v\n", res)
}
