package security_stats_query_time_series_stats_test

import (
	sdk "apigee_api/client"
	nullable "apigee_api/nullable"
	security_stats_query_time_series_stats "apigee_api/resources/v1/security_stats_query_time_series_stats"
	types "apigee_api/types"
	fmt "fmt"
	testing "testing"
)

func TestCreate200SuccessDefault(t *testing.T) {
	// Success test for Default body
	client := sdk.NewClient(sdk.WithOauth2("API_TOKEN"), sdk.WithOauth2c("API_TOKEN"), sdk.WithBaseURL("https://api.sideko.dev/v1/mock/public/apigee/0.1.0"))
	res, err := client.V1.SecurityStatsQueryTimeSeriesStats.Create(security_stats_query_time_series_stats.CreateRequest{Orgenv: "string", Data: types.GoogleCloudApigeeV1QueryTimeSeriesStatsRequest{Dimensions: nullable.NewValue([]string{"string"}), Filter: nullable.NewValue("string"), Metrics: nullable.NewValue([]types.GoogleCloudApigeeV1MetricAggregation{types.GoogleCloudApigeeV1MetricAggregation{Aggregation: nullable.NewValue(types.GoogleCloudApigeeV1MetricAggregationAggregationEnumAggregationFunctionUnspecified), Name: nullable.NewValue("string"), Order: nullable.NewValue(types.GoogleCloudApigeeV1MetricAggregationOrderEnumAscending)}}), PageSize: nullable.NewValue(123), PageToken: nullable.NewValue("string"), TimeRange: nullable.NewValue(types.GoogleTypeInterval{EndTime: nullable.NewValue("string"), StartTime: nullable.NewValue("string")}), TimestampOrder: nullable.NewValue(types.GoogleCloudApigeeV1QueryTimeSeriesStatsRequestTimestampOrderEnumAscending), WindowSize: nullable.NewValue(types.GoogleCloudApigeeV1QueryTimeSeriesStatsRequestWindowSizeEnumDay)}})

	if err != nil {
		t.Fatalf("TestCreate200SuccessDefault - failed making request with error: %#v", err)
	}

	fmt.Printf("response - %#v\n", res)
}
