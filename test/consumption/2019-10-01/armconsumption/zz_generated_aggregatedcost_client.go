//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armconsumption

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// AggregatedCostClient contains the methods for the AggregatedCost group.
// Don't use this type directly, use NewAggregatedCostClient() instead.
type AggregatedCostClient struct {
	host string
	pl   runtime.Pipeline
}

// NewAggregatedCostClient creates a new instance of AggregatedCostClient with the specified values.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewAggregatedCostClient(credential azcore.TokenCredential, options *arm.ClientOptions) *AggregatedCostClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	client := &AggregatedCostClient{
		host: string(cp.Host),
		pl:   armruntime.NewPipeline(module, version, credential, &cp),
	}
	return client
}

// GetByManagementGroup - Provides the aggregate cost of a management group and all child management groups by current billing
// period.
// If the operation fails it returns the *ErrorResponse error type.
// managementGroupID - Azure Management Group ID.
// options - AggregatedCostGetByManagementGroupOptions contains the optional parameters for the AggregatedCostClient.GetByManagementGroup
// method.
func (client *AggregatedCostClient) GetByManagementGroup(ctx context.Context, managementGroupID string, options *AggregatedCostGetByManagementGroupOptions) (AggregatedCostGetByManagementGroupResponse, error) {
	req, err := client.getByManagementGroupCreateRequest(ctx, managementGroupID, options)
	if err != nil {
		return AggregatedCostGetByManagementGroupResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AggregatedCostGetByManagementGroupResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AggregatedCostGetByManagementGroupResponse{}, client.getByManagementGroupHandleError(resp)
	}
	return client.getByManagementGroupHandleResponse(resp)
}

// getByManagementGroupCreateRequest creates the GetByManagementGroup request.
func (client *AggregatedCostClient) getByManagementGroupCreateRequest(ctx context.Context, managementGroupID string, options *AggregatedCostGetByManagementGroupOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Management/managementGroups/{managementGroupId}/providers/Microsoft.Consumption/aggregatedcost"
	if managementGroupID == "" {
		return nil, errors.New("parameter managementGroupID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managementGroupId}", url.PathEscape(managementGroupID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-10-01")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getByManagementGroupHandleResponse handles the GetByManagementGroup response.
func (client *AggregatedCostClient) getByManagementGroupHandleResponse(resp *http.Response) (AggregatedCostGetByManagementGroupResponse, error) {
	result := AggregatedCostGetByManagementGroupResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ManagementGroupAggregatedCostResult); err != nil {
		return AggregatedCostGetByManagementGroupResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getByManagementGroupHandleError handles the GetByManagementGroup error response.
func (client *AggregatedCostClient) getByManagementGroupHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// GetForBillingPeriodByManagementGroup - Provides the aggregate cost of a management group and all child management groups
// by specified billing period
// If the operation fails it returns the *ErrorResponse error type.
// managementGroupID - Azure Management Group ID.
// billingPeriodName - Billing Period Name.
// options - AggregatedCostGetForBillingPeriodByManagementGroupOptions contains the optional parameters for the AggregatedCostClient.GetForBillingPeriodByManagementGroup
// method.
func (client *AggregatedCostClient) GetForBillingPeriodByManagementGroup(ctx context.Context, managementGroupID string, billingPeriodName string, options *AggregatedCostGetForBillingPeriodByManagementGroupOptions) (AggregatedCostGetForBillingPeriodByManagementGroupResponse, error) {
	req, err := client.getForBillingPeriodByManagementGroupCreateRequest(ctx, managementGroupID, billingPeriodName, options)
	if err != nil {
		return AggregatedCostGetForBillingPeriodByManagementGroupResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AggregatedCostGetForBillingPeriodByManagementGroupResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AggregatedCostGetForBillingPeriodByManagementGroupResponse{}, client.getForBillingPeriodByManagementGroupHandleError(resp)
	}
	return client.getForBillingPeriodByManagementGroupHandleResponse(resp)
}

// getForBillingPeriodByManagementGroupCreateRequest creates the GetForBillingPeriodByManagementGroup request.
func (client *AggregatedCostClient) getForBillingPeriodByManagementGroupCreateRequest(ctx context.Context, managementGroupID string, billingPeriodName string, options *AggregatedCostGetForBillingPeriodByManagementGroupOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Management/managementGroups/{managementGroupId}/providers/Microsoft.Billing/billingPeriods/{billingPeriodName}/Microsoft.Consumption/aggregatedCost"
	if managementGroupID == "" {
		return nil, errors.New("parameter managementGroupID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managementGroupId}", url.PathEscape(managementGroupID))
	if billingPeriodName == "" {
		return nil, errors.New("parameter billingPeriodName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingPeriodName}", url.PathEscape(billingPeriodName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getForBillingPeriodByManagementGroupHandleResponse handles the GetForBillingPeriodByManagementGroup response.
func (client *AggregatedCostClient) getForBillingPeriodByManagementGroupHandleResponse(resp *http.Response) (AggregatedCostGetForBillingPeriodByManagementGroupResponse, error) {
	result := AggregatedCostGetForBillingPeriodByManagementGroupResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ManagementGroupAggregatedCostResult); err != nil {
		return AggregatedCostGetForBillingPeriodByManagementGroupResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getForBillingPeriodByManagementGroupHandleError handles the GetForBillingPeriodByManagementGroup error response.
func (client *AggregatedCostClient) getForBillingPeriodByManagementGroupHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
