// Code generated by Sideko (sideko.dev)
package exports

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

// Submit a data export job to be processed in the background. If the request is successful, the API returns a 201 status, a URI that can be used to retrieve the status of the export job, and the `state` value of "enqueued".
func (c *Client) Create(request CreateRequest, reqModifiers ...RequestModifier) (types.GoogleCloudApigeeV1Export, error) {
	// build & send request (keep comment for code generation)

	// URL formatting
	joined, err := url.JoinPath(c.coreClient.BaseURL, "/v1/"+sdkcore.FmtStringParam(request.Parent)+"/analytics/"+"exports")
	if err != nil {
		return types.GoogleCloudApigeeV1Export{}, err
	}
	url, err := url.Parse(joined)
	if err != nil {
		return types.GoogleCloudApigeeV1Export{}, err
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
		return types.GoogleCloudApigeeV1Export{}, err
	}
	reqBodyBuf := bytes.NewBuffer([]byte(reqBody))

	// Init request
	req, err := http.NewRequest("POST", url.String(), reqBodyBuf)
	if err != nil {
		return types.GoogleCloudApigeeV1Export{}, err
	}

	// Add headers
	req.Header.Add("x-sideko-sdk-language", "Go")
	req.Header.Add("Content-Type", "application/json")

	// Add auth
	c.coreClient.AddAuth([]string{"Oauth2", "Oauth2c"}, req)

	// Add base client & request level modifiers
	if err := c.coreClient.ApplyModifiers(req, reqModifiers); err != nil {
		return types.GoogleCloudApigeeV1Export{}, err
	}

	// Dispatch request
	resp, err := c.coreClient.HttpClient.Do(req)
	if err != nil {
		return types.GoogleCloudApigeeV1Export{}, err
	}
	defer resp.Body.Close()

	// Handle response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return types.GoogleCloudApigeeV1Export{}, err
	}

	// Check status
	if resp.StatusCode >= 300 {
		return types.GoogleCloudApigeeV1Export{}, sdkcore.NewApiError(*req, *resp, body)
	}
	var bodyData types.GoogleCloudApigeeV1Export
	err = json.Unmarshal(body, &bodyData)
	if err != nil {
		return types.GoogleCloudApigeeV1Export{}, err
	}
	return bodyData, nil
}

// Lists the details and status of all analytics export jobs belonging to the parent organization and environment.
func (c *Client) List(request ListRequest, reqModifiers ...RequestModifier) (types.GoogleCloudApigeeV1ListExportsResponse, error) {
	// build & send request (keep comment for code generation)

	// URL formatting
	joined, err := url.JoinPath(c.coreClient.BaseURL, "/v1/"+sdkcore.FmtStringParam(request.Parent)+"/analytics/"+"exports")
	if err != nil {
		return types.GoogleCloudApigeeV1ListExportsResponse{}, err
	}
	url, err := url.Parse(joined)
	if err != nil {
		return types.GoogleCloudApigeeV1ListExportsResponse{}, err
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

	// Init request
	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		return types.GoogleCloudApigeeV1ListExportsResponse{}, err
	}

	// Add headers
	req.Header.Add("x-sideko-sdk-language", "Go")

	// Add auth
	c.coreClient.AddAuth([]string{"Oauth2", "Oauth2c"}, req)

	// Add base client & request level modifiers
	if err := c.coreClient.ApplyModifiers(req, reqModifiers); err != nil {
		return types.GoogleCloudApigeeV1ListExportsResponse{}, err
	}

	// Dispatch request
	resp, err := c.coreClient.HttpClient.Do(req)
	if err != nil {
		return types.GoogleCloudApigeeV1ListExportsResponse{}, err
	}
	defer resp.Body.Close()

	// Handle response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return types.GoogleCloudApigeeV1ListExportsResponse{}, err
	}

	// Check status
	if resp.StatusCode >= 300 {
		return types.GoogleCloudApigeeV1ListExportsResponse{}, sdkcore.NewApiError(*req, *resp, body)
	}
	var bodyData types.GoogleCloudApigeeV1ListExportsResponse
	err = json.Unmarshal(body, &bodyData)
	if err != nil {
		return types.GoogleCloudApigeeV1ListExportsResponse{}, err
	}
	return bodyData, nil
}
