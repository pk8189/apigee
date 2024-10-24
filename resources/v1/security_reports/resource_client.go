// Code generated by Sideko (sideko.dev)
package security_reports

import (
	sdkcore "apigee_api/core"
	types "apigee_api/types"
	bytes "bytes"
	json "encoding/json"
	io "io"
	http "net/http"
	url "net/url"
)

type Client struct {
	coreClient *sdkcore.CoreClient
}
type RequestModifier = func(req *http.Request) error

// Instantiate a new resource client
func NewClient(coreClient *sdkcore.CoreClient) *Client {
	client := Client{
		coreClient: coreClient,
	}

	return &client
}

// register api methods (keep comment for code generation)

// Submit a report request to be processed in the background. If the submission succeeds, the API returns a 200 status and an ID that refer to the report request. In addition to the HTTP status 200, the `state` of "enqueued" means that the request succeeded.
func (c *Client) Create(request CreateRequest, reqModifiers ...RequestModifier) (types.GoogleCloudApigeeV1SecurityReport, error) {
	// build & send request (keep comment for code generation)

	// URL formatting
	joined, err := url.JoinPath(c.coreClient.BaseURL, "/v1/"+sdkcore.FmtStringParam(request.Parent)+"/securityReports")
	if err != nil {
		return types.GoogleCloudApigeeV1SecurityReport{}, err
	}
	url, err := url.Parse(joined)
	if err != nil {
		return types.GoogleCloudApigeeV1SecurityReport{}, err
	}

	// Query params
	params := url.Query()
	queryXgafv, err := request.QueryXgafv.Value()
	if err != nil {
		sdkcore.AddQueryParam(params, "$.xgafv", queryXgafv, false)
	}
	accessToken, err := request.AccessToken.Value()
	if err != nil {
		sdkcore.AddQueryParam(params, "access_token", accessToken, false)
	}
	alt, err := request.Alt.Value()
	if err != nil {
		sdkcore.AddQueryParam(params, "alt", alt, false)
	}
	callback, err := request.Callback.Value()
	if err != nil {
		sdkcore.AddQueryParam(params, "callback", callback, false)
	}
	fields, err := request.Fields.Value()
	if err != nil {
		sdkcore.AddQueryParam(params, "fields", fields, false)
	}
	key, err := request.Key.Value()
	if err != nil {
		sdkcore.AddQueryParam(params, "key", key, false)
	}
	oauthToken, err := request.OauthToken.Value()
	if err != nil {
		sdkcore.AddQueryParam(params, "oauth_token", oauthToken, false)
	}
	prettyPrint, err := request.PrettyPrint.Value()
	if err != nil {
		sdkcore.AddQueryParam(params, "prettyPrint", prettyPrint, false)
	}
	quotaUser, err := request.QuotaUser.Value()
	if err != nil {
		sdkcore.AddQueryParam(params, "quotaUser", quotaUser, false)
	}
	uploadType, err := request.UploadType.Value()
	if err != nil {
		sdkcore.AddQueryParam(params, "uploadType", uploadType, false)
	}
	uploadProtocol, err := request.UploadProtocol.Value()
	if err != nil {
		sdkcore.AddQueryParam(params, "upload_protocol", uploadProtocol, false)
	}
	url.RawQuery = params.Encode()

	// Prep body
	reqBody, err := json.Marshal(request.Data)
	if err != nil {
		return types.GoogleCloudApigeeV1SecurityReport{}, err
	}
	reqBodyBuf := bytes.NewBuffer([]byte(reqBody))

	// Init request
	req, err := http.NewRequest("POST", url.String(), reqBodyBuf)
	if err != nil {
		return types.GoogleCloudApigeeV1SecurityReport{}, err
	}

	// Add headers
	req.Header.Add("x-sideko-sdk-language", "Go")
	req.Header.Add("Content-Type", "application/json")

	// Add auth
	c.coreClient.AddAuth([]string{"Oauth2", "Oauth2c"}, req)

	// Add base client & request level modifiers
	if err := c.coreClient.ApplyModifiers(req, reqModifiers); err != nil {
		return types.GoogleCloudApigeeV1SecurityReport{}, err
	}

	// Dispatch request
	resp, err := c.coreClient.HttpClient.Do(req)
	if err != nil {
		return types.GoogleCloudApigeeV1SecurityReport{}, err
	}
	defer resp.Body.Close()

	// Handle response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return types.GoogleCloudApigeeV1SecurityReport{}, err
	}

	// Check status
	if resp.StatusCode >= 300 {
		return types.GoogleCloudApigeeV1SecurityReport{}, sdkcore.NewApiError(*req, *resp, body)
	}
	var bodyData types.GoogleCloudApigeeV1SecurityReport
	err = json.Unmarshal(body, &bodyData)
	if err != nil {
		return types.GoogleCloudApigeeV1SecurityReport{}, err
	}
	return bodyData, nil
}

// Return a list of Security Reports
func (c *Client) List(request ListRequest, reqModifiers ...RequestModifier) (types.GoogleCloudApigeeV1ListSecurityReportsResponse, error) {
	// build & send request (keep comment for code generation)

	// URL formatting
	joined, err := url.JoinPath(c.coreClient.BaseURL, "/v1/"+sdkcore.FmtStringParam(request.Parent)+"/securityReports")
	if err != nil {
		return types.GoogleCloudApigeeV1ListSecurityReportsResponse{}, err
	}
	url, err := url.Parse(joined)
	if err != nil {
		return types.GoogleCloudApigeeV1ListSecurityReportsResponse{}, err
	}

	// Query params
	params := url.Query()
	queryXgafv, err := request.QueryXgafv.Value()
	if err != nil {
		sdkcore.AddQueryParam(params, "$.xgafv", queryXgafv, false)
	}
	accessToken, err := request.AccessToken.Value()
	if err != nil {
		sdkcore.AddQueryParam(params, "access_token", accessToken, false)
	}
	alt, err := request.Alt.Value()
	if err != nil {
		sdkcore.AddQueryParam(params, "alt", alt, false)
	}
	callback, err := request.Callback.Value()
	if err != nil {
		sdkcore.AddQueryParam(params, "callback", callback, false)
	}
	dataset, err := request.Dataset.Value()
	if err != nil {
		sdkcore.AddQueryParam(params, "dataset", dataset, false)
	}
	fields, err := request.Fields.Value()
	if err != nil {
		sdkcore.AddQueryParam(params, "fields", fields, false)
	}
	from, err := request.From.Value()
	if err != nil {
		sdkcore.AddQueryParam(params, "from", from, false)
	}
	key, err := request.Key.Value()
	if err != nil {
		sdkcore.AddQueryParam(params, "key", key, false)
	}
	oauthToken, err := request.OauthToken.Value()
	if err != nil {
		sdkcore.AddQueryParam(params, "oauth_token", oauthToken, false)
	}
	pageSize, err := request.PageSize.Value()
	if err != nil {
		sdkcore.AddQueryParam(params, "pageSize", pageSize, false)
	}
	pageToken, err := request.PageToken.Value()
	if err != nil {
		sdkcore.AddQueryParam(params, "pageToken", pageToken, false)
	}
	prettyPrint, err := request.PrettyPrint.Value()
	if err != nil {
		sdkcore.AddQueryParam(params, "prettyPrint", prettyPrint, false)
	}
	quotaUser, err := request.QuotaUser.Value()
	if err != nil {
		sdkcore.AddQueryParam(params, "quotaUser", quotaUser, false)
	}
	status, err := request.Status.Value()
	if err != nil {
		sdkcore.AddQueryParam(params, "status", status, false)
	}
	submittedBy, err := request.SubmittedBy.Value()
	if err != nil {
		sdkcore.AddQueryParam(params, "submittedBy", submittedBy, false)
	}
	to, err := request.To.Value()
	if err != nil {
		sdkcore.AddQueryParam(params, "to", to, false)
	}
	uploadType, err := request.UploadType.Value()
	if err != nil {
		sdkcore.AddQueryParam(params, "uploadType", uploadType, false)
	}
	uploadProtocol, err := request.UploadProtocol.Value()
	if err != nil {
		sdkcore.AddQueryParam(params, "upload_protocol", uploadProtocol, false)
	}
	url.RawQuery = params.Encode()

	// Init request
	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		return types.GoogleCloudApigeeV1ListSecurityReportsResponse{}, err
	}

	// Add headers
	req.Header.Add("x-sideko-sdk-language", "Go")

	// Add auth
	c.coreClient.AddAuth([]string{"Oauth2", "Oauth2c"}, req)

	// Add base client & request level modifiers
	if err := c.coreClient.ApplyModifiers(req, reqModifiers); err != nil {
		return types.GoogleCloudApigeeV1ListSecurityReportsResponse{}, err
	}

	// Dispatch request
	resp, err := c.coreClient.HttpClient.Do(req)
	if err != nil {
		return types.GoogleCloudApigeeV1ListSecurityReportsResponse{}, err
	}
	defer resp.Body.Close()

	// Handle response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return types.GoogleCloudApigeeV1ListSecurityReportsResponse{}, err
	}

	// Check status
	if resp.StatusCode >= 300 {
		return types.GoogleCloudApigeeV1ListSecurityReportsResponse{}, sdkcore.NewApiError(*req, *resp, body)
	}
	var bodyData types.GoogleCloudApigeeV1ListSecurityReportsResponse
	err = json.Unmarshal(body, &bodyData)
	if err != nil {
		return types.GoogleCloudApigeeV1ListSecurityReportsResponse{}, err
	}
	return bodyData, nil
}
