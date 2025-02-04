// Code generated by Sideko (sideko.dev)
package set_iam_policy

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

// Sets the IAM policy on an environment, if the policy already exists it will be replaced. For more information, see [Manage users, roles, and permissions using the API](https://cloud.google.com/apigee/docs/api-platform/system-administration/manage-users-roles). You must have the `apigee.environments.setIamPolicy` permission to call this API.
func (c *Client) Create(request CreateRequest, reqModifiers ...RequestModifier) (types.GoogleIamV1Policy, error) {
	// build & send request (keep comment for code generation)

	// URL formatting
	joined, err := url.JoinPath(c.coreClient.BaseURL, "/v1/"+sdkcore.FmtStringParam(request.Resource)+":setIamPolicy")
	if err != nil {
		return types.GoogleIamV1Policy{}, err
	}
	url, err := url.Parse(joined)
	if err != nil {
		return types.GoogleIamV1Policy{}, err
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
		return types.GoogleIamV1Policy{}, err
	}
	reqBodyBuf := bytes.NewBuffer([]byte(reqBody))

	// Init request
	req, err := http.NewRequest("POST", url.String(), reqBodyBuf)
	if err != nil {
		return types.GoogleIamV1Policy{}, err
	}

	// Add headers
	req.Header.Add("x-sideko-sdk-language", "Go")
	req.Header.Add("Content-Type", "application/json")

	// Add auth
	c.coreClient.AddAuth([]string{"Oauth2", "Oauth2c"}, req)

	// Add base client & request level modifiers
	if err := c.coreClient.ApplyModifiers(req, reqModifiers); err != nil {
		return types.GoogleIamV1Policy{}, err
	}

	// Dispatch request
	resp, err := c.coreClient.HttpClient.Do(req)
	if err != nil {
		return types.GoogleIamV1Policy{}, err
	}
	defer resp.Body.Close()

	// Handle response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return types.GoogleIamV1Policy{}, err
	}

	// Check status
	if resp.StatusCode >= 300 {
		return types.GoogleIamV1Policy{}, sdkcore.NewApiError(*req, *resp, body)
	}
	var bodyData types.GoogleIamV1Policy
	err = json.Unmarshal(body, &bodyData)
	if err != nil {
		return types.GoogleIamV1Policy{}, err
	}
	return bodyData, nil
}
