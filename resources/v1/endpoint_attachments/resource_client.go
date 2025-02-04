// Code generated by Sideko (sideko.dev)
package endpoint_attachments

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

// Creates an endpoint attachment. **Note:** Not supported for Apigee hybrid.
func (c *Client) Create(request CreateRequest, reqModifiers ...RequestModifier) (types.GoogleLongrunningOperation, error) {
	// build & send request (keep comment for code generation)

	// URL formatting
	joined, err := url.JoinPath(c.coreClient.BaseURL, "/v1/"+sdkcore.FmtStringParam(request.Parent)+"/endpointAttachments")
	if err != nil {
		return types.GoogleLongrunningOperation{}, err
	}
	url, err := url.Parse(joined)
	if err != nil {
		return types.GoogleLongrunningOperation{}, err
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
	endpointAttachmentId, err := request.EndpointAttachmentId.Value()
	if err != nil {
		sdkcore.AddQueryParam(params, "endpointAttachmentId", endpointAttachmentId, false)
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
		return types.GoogleLongrunningOperation{}, err
	}
	reqBodyBuf := bytes.NewBuffer([]byte(reqBody))

	// Init request
	req, err := http.NewRequest("POST", url.String(), reqBodyBuf)
	if err != nil {
		return types.GoogleLongrunningOperation{}, err
	}

	// Add headers
	req.Header.Add("x-sideko-sdk-language", "Go")
	req.Header.Add("Content-Type", "application/json")

	// Add auth
	c.coreClient.AddAuth([]string{"Oauth2", "Oauth2c"}, req)

	// Add base client & request level modifiers
	if err := c.coreClient.ApplyModifiers(req, reqModifiers); err != nil {
		return types.GoogleLongrunningOperation{}, err
	}

	// Dispatch request
	resp, err := c.coreClient.HttpClient.Do(req)
	if err != nil {
		return types.GoogleLongrunningOperation{}, err
	}
	defer resp.Body.Close()

	// Handle response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return types.GoogleLongrunningOperation{}, err
	}

	// Check status
	if resp.StatusCode >= 300 {
		return types.GoogleLongrunningOperation{}, sdkcore.NewApiError(*req, *resp, body)
	}
	var bodyData types.GoogleLongrunningOperation
	err = json.Unmarshal(body, &bodyData)
	if err != nil {
		return types.GoogleLongrunningOperation{}, err
	}
	return bodyData, nil
}

// Lists the endpoint attachments in an organization.
func (c *Client) List(request ListRequest, reqModifiers ...RequestModifier) (types.GoogleCloudApigeeV1ListEndpointAttachmentsResponse, error) {
	// build & send request (keep comment for code generation)

	// URL formatting
	joined, err := url.JoinPath(c.coreClient.BaseURL, "/v1/"+sdkcore.FmtStringParam(request.Parent)+"/endpointAttachments")
	if err != nil {
		return types.GoogleCloudApigeeV1ListEndpointAttachmentsResponse{}, err
	}
	url, err := url.Parse(joined)
	if err != nil {
		return types.GoogleCloudApigeeV1ListEndpointAttachmentsResponse{}, err
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
		return types.GoogleCloudApigeeV1ListEndpointAttachmentsResponse{}, err
	}

	// Add headers
	req.Header.Add("x-sideko-sdk-language", "Go")

	// Add auth
	c.coreClient.AddAuth([]string{"Oauth2", "Oauth2c"}, req)

	// Add base client & request level modifiers
	if err := c.coreClient.ApplyModifiers(req, reqModifiers); err != nil {
		return types.GoogleCloudApigeeV1ListEndpointAttachmentsResponse{}, err
	}

	// Dispatch request
	resp, err := c.coreClient.HttpClient.Do(req)
	if err != nil {
		return types.GoogleCloudApigeeV1ListEndpointAttachmentsResponse{}, err
	}
	defer resp.Body.Close()

	// Handle response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return types.GoogleCloudApigeeV1ListEndpointAttachmentsResponse{}, err
	}

	// Check status
	if resp.StatusCode >= 300 {
		return types.GoogleCloudApigeeV1ListEndpointAttachmentsResponse{}, sdkcore.NewApiError(*req, *resp, body)
	}
	var bodyData types.GoogleCloudApigeeV1ListEndpointAttachmentsResponse
	err = json.Unmarshal(body, &bodyData)
	if err != nil {
		return types.GoogleCloudApigeeV1ListEndpointAttachmentsResponse{}, err
	}
	return bodyData, nil
}
