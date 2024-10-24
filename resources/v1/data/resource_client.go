// Code generated by Sideko (sideko.dev)
package data

import (
	sdkcore "apigee_api/core"
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

// Deletes the data from a debug session. This does not cancel the debug session or prevent further data from being collected if the session is still active in runtime pods.
func (c *Client) Delete(request DeleteRequest, reqModifiers ...RequestModifier) (map[string]interface{}, error) {
	// build & send request (keep comment for code generation)

	// URL formatting
	joined, err := url.JoinPath(c.coreClient.BaseURL, "/v1/"+sdkcore.FmtStringParam(request.Name)+"/data")
	if err != nil {
		return map[string]interface{}{}, err
	}
	url, err := url.Parse(joined)
	if err != nil {
		return map[string]interface{}{}, err
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
	req, err := http.NewRequest("DELETE", url.String(), nil)
	if err != nil {
		return map[string]interface{}{}, err
	}

	// Add headers
	req.Header.Add("x-sideko-sdk-language", "Go")

	// Add auth
	c.coreClient.AddAuth([]string{"Oauth2", "Oauth2c"}, req)

	// Add base client & request level modifiers
	if err := c.coreClient.ApplyModifiers(req, reqModifiers); err != nil {
		return map[string]interface{}{}, err
	}

	// Dispatch request
	resp, err := c.coreClient.HttpClient.Do(req)
	if err != nil {
		return map[string]interface{}{}, err
	}
	defer resp.Body.Close()

	// Handle response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return map[string]interface{}{}, err
	}

	// Check status
	if resp.StatusCode >= 300 {
		return map[string]interface{}{}, sdkcore.NewApiError(*req, *resp, body)
	}
	var bodyData map[string]interface{}
	err = json.Unmarshal(body, &bodyData)
	if err != nil {
		return map[string]interface{}{}, err
	}
	return bodyData, nil
}
