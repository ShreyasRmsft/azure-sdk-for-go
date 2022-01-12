//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetapp

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// SnapshotPoliciesClient contains the methods for the SnapshotPolicies group.
// Don't use this type directly, use NewSnapshotPoliciesClient() instead.
type SnapshotPoliciesClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewSnapshotPoliciesClient creates a new instance of SnapshotPoliciesClient with the specified values.
func NewSnapshotPoliciesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *SnapshotPoliciesClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &SnapshotPoliciesClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// Create - Create a snapshot policy
// If the operation fails it returns a generic error.
func (client *SnapshotPoliciesClient) Create(ctx context.Context, resourceGroupName string, accountName string, snapshotPolicyName string, body SnapshotPolicy, options *SnapshotPoliciesCreateOptions) (SnapshotPoliciesCreateResponse, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, accountName, snapshotPolicyName, body, options)
	if err != nil {
		return SnapshotPoliciesCreateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SnapshotPoliciesCreateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return SnapshotPoliciesCreateResponse{}, client.createHandleError(resp)
	}
	return client.createHandleResponse(resp)
}

// createCreateRequest creates the Create request.
func (client *SnapshotPoliciesClient) createCreateRequest(ctx context.Context, resourceGroupName string, accountName string, snapshotPolicyName string, body SnapshotPolicy, options *SnapshotPoliciesCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts/{accountName}/snapshotPolicies/{snapshotPolicyName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if snapshotPolicyName == "" {
		return nil, errors.New("parameter snapshotPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{snapshotPolicyName}", url.PathEscape(snapshotPolicyName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, body)
}

// createHandleResponse handles the Create response.
func (client *SnapshotPoliciesClient) createHandleResponse(resp *http.Response) (SnapshotPoliciesCreateResponse, error) {
	result := SnapshotPoliciesCreateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.SnapshotPolicy); err != nil {
		return SnapshotPoliciesCreateResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// createHandleError handles the Create error response.
func (client *SnapshotPoliciesClient) createHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// BeginDelete - Delete snapshot policy
// If the operation fails it returns a generic error.
func (client *SnapshotPoliciesClient) BeginDelete(ctx context.Context, resourceGroupName string, accountName string, snapshotPolicyName string, options *SnapshotPoliciesBeginDeleteOptions) (SnapshotPoliciesDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, accountName, snapshotPolicyName, options)
	if err != nil {
		return SnapshotPoliciesDeletePollerResponse{}, err
	}
	result := SnapshotPoliciesDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("SnapshotPoliciesClient.Delete", "location", resp, client.pl, client.deleteHandleError)
	if err != nil {
		return SnapshotPoliciesDeletePollerResponse{}, err
	}
	result.Poller = &SnapshotPoliciesDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Delete snapshot policy
// If the operation fails it returns a generic error.
func (client *SnapshotPoliciesClient) deleteOperation(ctx context.Context, resourceGroupName string, accountName string, snapshotPolicyName string, options *SnapshotPoliciesBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, accountName, snapshotPolicyName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *SnapshotPoliciesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, accountName string, snapshotPolicyName string, options *SnapshotPoliciesBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts/{accountName}/snapshotPolicies/{snapshotPolicyName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if snapshotPolicyName == "" {
		return nil, errors.New("parameter snapshotPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{snapshotPolicyName}", url.PathEscape(snapshotPolicyName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *SnapshotPoliciesClient) deleteHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// Get - Get a snapshot Policy
// If the operation fails it returns a generic error.
func (client *SnapshotPoliciesClient) Get(ctx context.Context, resourceGroupName string, accountName string, snapshotPolicyName string, options *SnapshotPoliciesGetOptions) (SnapshotPoliciesGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, accountName, snapshotPolicyName, options)
	if err != nil {
		return SnapshotPoliciesGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SnapshotPoliciesGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SnapshotPoliciesGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *SnapshotPoliciesClient) getCreateRequest(ctx context.Context, resourceGroupName string, accountName string, snapshotPolicyName string, options *SnapshotPoliciesGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts/{accountName}/snapshotPolicies/{snapshotPolicyName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if snapshotPolicyName == "" {
		return nil, errors.New("parameter snapshotPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{snapshotPolicyName}", url.PathEscape(snapshotPolicyName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *SnapshotPoliciesClient) getHandleResponse(resp *http.Response) (SnapshotPoliciesGetResponse, error) {
	result := SnapshotPoliciesGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.SnapshotPolicy); err != nil {
		return SnapshotPoliciesGetResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *SnapshotPoliciesClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// List - List snapshot policy
// If the operation fails it returns a generic error.
func (client *SnapshotPoliciesClient) List(ctx context.Context, resourceGroupName string, accountName string, options *SnapshotPoliciesListOptions) (SnapshotPoliciesListResponse, error) {
	req, err := client.listCreateRequest(ctx, resourceGroupName, accountName, options)
	if err != nil {
		return SnapshotPoliciesListResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SnapshotPoliciesListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SnapshotPoliciesListResponse{}, client.listHandleError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *SnapshotPoliciesClient) listCreateRequest(ctx context.Context, resourceGroupName string, accountName string, options *SnapshotPoliciesListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts/{accountName}/snapshotPolicies"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *SnapshotPoliciesClient) listHandleResponse(resp *http.Response) (SnapshotPoliciesListResponse, error) {
	result := SnapshotPoliciesListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.SnapshotPoliciesList); err != nil {
		return SnapshotPoliciesListResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *SnapshotPoliciesClient) listHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// ListVolumes - Get volumes associated with snapshot policy
// If the operation fails it returns a generic error.
func (client *SnapshotPoliciesClient) ListVolumes(ctx context.Context, resourceGroupName string, accountName string, snapshotPolicyName string, options *SnapshotPoliciesListVolumesOptions) (SnapshotPoliciesListVolumesResponse, error) {
	req, err := client.listVolumesCreateRequest(ctx, resourceGroupName, accountName, snapshotPolicyName, options)
	if err != nil {
		return SnapshotPoliciesListVolumesResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SnapshotPoliciesListVolumesResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SnapshotPoliciesListVolumesResponse{}, client.listVolumesHandleError(resp)
	}
	return client.listVolumesHandleResponse(resp)
}

// listVolumesCreateRequest creates the ListVolumes request.
func (client *SnapshotPoliciesClient) listVolumesCreateRequest(ctx context.Context, resourceGroupName string, accountName string, snapshotPolicyName string, options *SnapshotPoliciesListVolumesOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts/{accountName}/snapshotPolicies/{snapshotPolicyName}/volumes"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if snapshotPolicyName == "" {
		return nil, errors.New("parameter snapshotPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{snapshotPolicyName}", url.PathEscape(snapshotPolicyName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listVolumesHandleResponse handles the ListVolumes response.
func (client *SnapshotPoliciesClient) listVolumesHandleResponse(resp *http.Response) (SnapshotPoliciesListVolumesResponse, error) {
	result := SnapshotPoliciesListVolumesResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.SnapshotPolicyVolumeList); err != nil {
		return SnapshotPoliciesListVolumesResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listVolumesHandleError handles the ListVolumes error response.
func (client *SnapshotPoliciesClient) listVolumesHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// BeginUpdate - Patch a snapshot policy
// If the operation fails it returns a generic error.
func (client *SnapshotPoliciesClient) BeginUpdate(ctx context.Context, resourceGroupName string, accountName string, snapshotPolicyName string, body SnapshotPolicyPatch, options *SnapshotPoliciesBeginUpdateOptions) (SnapshotPoliciesUpdatePollerResponse, error) {
	resp, err := client.update(ctx, resourceGroupName, accountName, snapshotPolicyName, body, options)
	if err != nil {
		return SnapshotPoliciesUpdatePollerResponse{}, err
	}
	result := SnapshotPoliciesUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("SnapshotPoliciesClient.Update", "location", resp, client.pl, client.updateHandleError)
	if err != nil {
		return SnapshotPoliciesUpdatePollerResponse{}, err
	}
	result.Poller = &SnapshotPoliciesUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// Update - Patch a snapshot policy
// If the operation fails it returns a generic error.
func (client *SnapshotPoliciesClient) update(ctx context.Context, resourceGroupName string, accountName string, snapshotPolicyName string, body SnapshotPolicyPatch, options *SnapshotPoliciesBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, accountName, snapshotPolicyName, body, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, client.updateHandleError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *SnapshotPoliciesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, accountName string, snapshotPolicyName string, body SnapshotPolicyPatch, options *SnapshotPoliciesBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts/{accountName}/snapshotPolicies/{snapshotPolicyName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if snapshotPolicyName == "" {
		return nil, errors.New("parameter snapshotPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{snapshotPolicyName}", url.PathEscape(snapshotPolicyName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, body)
}

// updateHandleError handles the Update error response.
func (client *SnapshotPoliciesClient) updateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}