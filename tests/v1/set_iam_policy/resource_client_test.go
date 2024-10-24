package set_iam_policy_test

import (
	sdk "apigee_api/client"
	nullable "apigee_api/nullable"
	set_iam_policy "apigee_api/resources/v1/set_iam_policy"
	types "apigee_api/types"
	fmt "fmt"
	testing "testing"
)

func TestCreate200SuccessDefault(t *testing.T) {
	// Success test for Default body
	client := sdk.NewClient(sdk.WithOauth2("API_TOKEN"), sdk.WithOauth2c("API_TOKEN"), sdk.WithBaseURL("https://api.sideko.dev/v1/mock/public/apigee/0.1.0"))
	res, err := client.V1.SetIamPolicy.Create(set_iam_policy.CreateRequest{Resource: "string", Data: types.GoogleIamV1SetIamPolicyRequest{Policy: nullable.NewValue(types.GoogleIamV1Policy{AuditConfigs: nullable.NewValue([]types.GoogleIamV1AuditConfig{types.GoogleIamV1AuditConfig{AuditLogConfigs: nullable.NewValue([]types.GoogleIamV1AuditLogConfig{types.GoogleIamV1AuditLogConfig{ExemptedMembers: nullable.NewValue([]string{"string"}), LogType: nullable.NewValue(types.GoogleIamV1AuditLogConfigLogTypeEnumAdminRead)}}), Service: nullable.NewValue("string")}}), Bindings: nullable.NewValue([]types.GoogleIamV1Binding{types.GoogleIamV1Binding{Condition: nullable.NewValue(types.GoogleTypeExpr{Description: nullable.NewValue("string"), Expression: nullable.NewValue("string"), Location: nullable.NewValue("string"), Title: nullable.NewValue("string")}), Members: nullable.NewValue([]string{"string"}), Role: nullable.NewValue("string")}}), Etag: nullable.NewValue("string"), Version: nullable.NewValue(123)}), UpdateMask: nullable.NewValue("string")}})

	if err != nil {
		t.Fatalf("TestCreate200SuccessDefault - failed making request with error: %#v", err)
	}

	fmt.Printf("response - %#v\n", res)
}
