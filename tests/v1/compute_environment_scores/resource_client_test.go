package compute_environment_scores_test

import (
	sdk "apigee_api/client"
	nullable "apigee_api/nullable"
	compute_environment_scores "apigee_api/resources/v1/compute_environment_scores"
	types "apigee_api/types"
	fmt "fmt"
	testing "testing"
)

func TestCreate200SuccessDefault(t *testing.T) {
	// Success test for Default body
	client := sdk.NewClient(sdk.WithOauth2("API_TOKEN"), sdk.WithOauth2c("API_TOKEN"), sdk.WithBaseURL("https://api.sideko.dev/v1/mock/public/apigee/0.1.0"))
	res, err := client.V1.ComputeEnvironmentScores.Create(compute_environment_scores.CreateRequest{ProfileEnvironment: "string", Data: types.GoogleCloudApigeeV1ComputeEnvironmentScoresRequest{Filters: nullable.NewValue([]types.GoogleCloudApigeeV1ComputeEnvironmentScoresRequestFilter{types.GoogleCloudApigeeV1ComputeEnvironmentScoresRequestFilter{ScorePath: nullable.NewValue("string")}}), PageSize: nullable.NewValue(123), PageToken: nullable.NewValue("string"), TimeRange: nullable.NewValue(types.GoogleTypeInterval{EndTime: nullable.NewValue("string"), StartTime: nullable.NewValue("string")})}})

	if err != nil {
		t.Fatalf("TestCreate200SuccessDefault - failed making request with error: %#v", err)
	}

	fmt.Printf("response - %#v\n", res)
}
