package report_status_test

import (
	sdk "apigee_api/client"
	nullable "apigee_api/nullable"
	report_status "apigee_api/resources/v1/report_status"
	types "apigee_api/types"
	fmt "fmt"
	testing "testing"
)

func TestCreate200SuccessDefault(t *testing.T) {
	// Success test for Default body
	client := sdk.NewClient(sdk.WithOauth2("API_TOKEN"), sdk.WithOauth2c("API_TOKEN"), sdk.WithBaseURL("https://api.sideko.dev/v1/mock/public/apigee/0.1.0"))
	res, err := client.V1.ReportStatus.Create(report_status.CreateRequest{Instance: "string", Data: types.GoogleCloudApigeeV1ReportInstanceStatusRequest{InstanceUid: nullable.NewValue("string"), ReportTime: nullable.NewValue("string"), Resources: nullable.NewValue([]types.GoogleCloudApigeeV1ResourceStatus{types.GoogleCloudApigeeV1ResourceStatus{Resource: nullable.NewValue("string"), Revisions: nullable.NewValue([]types.GoogleCloudApigeeV1RevisionStatus{types.GoogleCloudApigeeV1RevisionStatus{Errors: nullable.NewValue([]types.GoogleCloudApigeeV1UpdateError{types.GoogleCloudApigeeV1UpdateError{Code: nullable.NewValue(types.GoogleCloudApigeeV1UpdateErrorCodeEnumAborted), Message: nullable.NewValue("string"), Resource: nullable.NewValue("string"), Type: nullable.NewValue("string")}}), JsonSpec: nullable.NewValue("string"), Replicas: nullable.NewValue(123), RevisionId: nullable.NewValue("string")}}), TotalReplicas: nullable.NewValue(123), Uid: nullable.NewValue("string")}})}})

	if err != nil {
		t.Fatalf("TestCreate200SuccessDefault - failed making request with error: %#v", err)
	}

	fmt.Printf("response - %#v\n", res)
}
