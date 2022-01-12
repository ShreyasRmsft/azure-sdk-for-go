//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armapplicationinsights

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
	"strconv"
	"strings"
)

// FavoritesClient contains the methods for the Favorites group.
// Don't use this type directly, use NewFavoritesClient() instead.
type FavoritesClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewFavoritesClient creates a new instance of FavoritesClient with the specified values.
func NewFavoritesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *FavoritesClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &FavoritesClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// Add - Adds a new favorites to an Application Insights component.
// If the operation fails it returns a generic error.
func (client *FavoritesClient) Add(ctx context.Context, resourceGroupName string, resourceName string, favoriteID string, favoriteProperties ApplicationInsightsComponentFavorite, options *FavoritesAddOptions) (FavoritesAddResponse, error) {
	req, err := client.addCreateRequest(ctx, resourceGroupName, resourceName, favoriteID, favoriteProperties, options)
	if err != nil {
		return FavoritesAddResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return FavoritesAddResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return FavoritesAddResponse{}, client.addHandleError(resp)
	}
	return client.addHandleResponse(resp)
}

// addCreateRequest creates the Add request.
func (client *FavoritesClient) addCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, favoriteID string, favoriteProperties ApplicationInsightsComponentFavorite, options *FavoritesAddOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}/favorites/{favoriteId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if favoriteID == "" {
		return nil, errors.New("parameter favoriteID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{favoriteId}", url.PathEscape(favoriteID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2015-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, favoriteProperties)
}

// addHandleResponse handles the Add response.
func (client *FavoritesClient) addHandleResponse(resp *http.Response) (FavoritesAddResponse, error) {
	result := FavoritesAddResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ApplicationInsightsComponentFavorite); err != nil {
		return FavoritesAddResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// addHandleError handles the Add error response.
func (client *FavoritesClient) addHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// Delete - Remove a favorite that is associated to an Application Insights component.
// If the operation fails it returns a generic error.
func (client *FavoritesClient) Delete(ctx context.Context, resourceGroupName string, resourceName string, favoriteID string, options *FavoritesDeleteOptions) (FavoritesDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, resourceName, favoriteID, options)
	if err != nil {
		return FavoritesDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return FavoritesDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return FavoritesDeleteResponse{}, client.deleteHandleError(resp)
	}
	return FavoritesDeleteResponse{RawResponse: resp}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *FavoritesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, favoriteID string, options *FavoritesDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}/favorites/{favoriteId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if favoriteID == "" {
		return nil, errors.New("parameter favoriteID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{favoriteId}", url.PathEscape(favoriteID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2015-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *FavoritesClient) deleteHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// Get - Get a single favorite by its FavoriteId, defined within an Application Insights component.
// If the operation fails it returns a generic error.
func (client *FavoritesClient) Get(ctx context.Context, resourceGroupName string, resourceName string, favoriteID string, options *FavoritesGetOptions) (FavoritesGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, resourceName, favoriteID, options)
	if err != nil {
		return FavoritesGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return FavoritesGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return FavoritesGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *FavoritesClient) getCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, favoriteID string, options *FavoritesGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}/favorites/{favoriteId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if favoriteID == "" {
		return nil, errors.New("parameter favoriteID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{favoriteId}", url.PathEscape(favoriteID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2015-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *FavoritesClient) getHandleResponse(resp *http.Response) (FavoritesGetResponse, error) {
	result := FavoritesGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ApplicationInsightsComponentFavorite); err != nil {
		return FavoritesGetResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *FavoritesClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// List - Gets a list of favorites defined within an Application Insights component.
// If the operation fails it returns a generic error.
func (client *FavoritesClient) List(ctx context.Context, resourceGroupName string, resourceName string, options *FavoritesListOptions) (FavoritesListResponse, error) {
	req, err := client.listCreateRequest(ctx, resourceGroupName, resourceName, options)
	if err != nil {
		return FavoritesListResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return FavoritesListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return FavoritesListResponse{}, client.listHandleError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *FavoritesClient) listCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, options *FavoritesListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}/favorites"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2015-05-01")
	if options != nil && options.FavoriteType != nil {
		reqQP.Set("favoriteType", string(*options.FavoriteType))
	}
	if options != nil && options.SourceType != nil {
		reqQP.Set("sourceType", string(*options.SourceType))
	}
	if options != nil && options.CanFetchContent != nil {
		reqQP.Set("canFetchContent", strconv.FormatBool(*options.CanFetchContent))
	}
	if options != nil && options.Tags != nil {
		reqQP.Set("tags", strings.Join(options.Tags, ","))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *FavoritesClient) listHandleResponse(resp *http.Response) (FavoritesListResponse, error) {
	result := FavoritesListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ApplicationInsightsComponentFavoriteArray); err != nil {
		return FavoritesListResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *FavoritesClient) listHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// Update - Updates a favorite that has already been added to an Application Insights component.
// If the operation fails it returns a generic error.
func (client *FavoritesClient) Update(ctx context.Context, resourceGroupName string, resourceName string, favoriteID string, favoriteProperties ApplicationInsightsComponentFavorite, options *FavoritesUpdateOptions) (FavoritesUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, resourceName, favoriteID, favoriteProperties, options)
	if err != nil {
		return FavoritesUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return FavoritesUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return FavoritesUpdateResponse{}, client.updateHandleError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *FavoritesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, favoriteID string, favoriteProperties ApplicationInsightsComponentFavorite, options *FavoritesUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}/favorites/{favoriteId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if favoriteID == "" {
		return nil, errors.New("parameter favoriteID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{favoriteId}", url.PathEscape(favoriteID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2015-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, favoriteProperties)
}

// updateHandleResponse handles the Update response.
func (client *FavoritesClient) updateHandleResponse(resp *http.Response) (FavoritesUpdateResponse, error) {
	result := FavoritesUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ApplicationInsightsComponentFavorite); err != nil {
		return FavoritesUpdateResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// updateHandleError handles the Update error response.
func (client *FavoritesClient) updateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}