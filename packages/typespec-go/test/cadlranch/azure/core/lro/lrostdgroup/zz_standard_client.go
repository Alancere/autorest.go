// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package lrostdgroup

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// StandardClient - Illustrates bodies templated with Azure Core with long-running operation
// Don't use this type directly, use a constructor function instead.
type StandardClient struct {
	internal   *azcore.Client
	apiVersion string
}

// BeginCreateOrReplace - Adds a user or replaces a user's fields.
//   - name - The name of user.
//   - resource - The resource instance.
//   - options - StandardClientCreateOrReplaceOptions contains the optional parameters for the StandardClient.CreateOrReplace
//     method.
func (client *StandardClient) BeginCreateOrReplace(ctx context.Context, name string, resource User, options *StandardClientCreateOrReplaceOptions) (*runtime.Poller[StandardClientCreateOrReplaceResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrReplace(ctx, name, resource, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller[StandardClientCreateOrReplaceResponse](resp, client.internal.Pipeline(), nil)
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[StandardClientCreateOrReplaceResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CreateOrReplace - Adds a user or replaces a user's fields.
func (client *StandardClient) createOrReplace(ctx context.Context, name string, resource User, options *StandardClientCreateOrReplaceOptions) (*http.Response, error) {
	var err error
	req, err := client.createOrReplaceCreateRequest(ctx, name, resource, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// createOrReplaceCreateRequest creates the CreateOrReplace request.
func (client *StandardClient) createOrReplaceCreateRequest(ctx context.Context, name string, resource User, options *StandardClientCreateOrReplaceOptions) (*policy.Request, error) {
	urlPath := "/azure/core/lro/standard/users/{name}"
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", client.apiVersion)
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	req.Raw().Header["Content-Type"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, resource); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Deletes a user.
//   - name - The name of user.
//   - options - StandardClientDeleteOptions contains the optional parameters for the StandardClient.Delete method.
func (client *StandardClient) BeginDelete(ctx context.Context, name string, options *StandardClientDeleteOptions) (*runtime.Poller[StandardClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, name, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller[StandardClientDeleteResponse](resp, client.internal.Pipeline(), nil)
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[StandardClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Deletes a user.
func (client *StandardClient) deleteOperation(ctx context.Context, name string, options *StandardClientDeleteOptions) (*http.Response, error) {
	var err error
	req, err := client.deleteCreateRequest(ctx, name, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *StandardClient) deleteCreateRequest(ctx context.Context, name string, options *StandardClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/azure/core/lro/standard/users/{name}"
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", client.apiVersion)
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// BeginExport - Exports a user.
//   - name - The name of user.
//   - formatParam - The format of the data.
//   - options - StandardClientExportOptions contains the optional parameters for the StandardClient.Export method.
func (client *StandardClient) BeginExport(ctx context.Context, name string, formatParam string, options *StandardClientExportOptions) (*runtime.Poller[StandardClientExportResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.export(ctx, name, formatParam, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller[StandardClientExportResponse](resp, client.internal.Pipeline(), nil)
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[StandardClientExportResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Export - Exports a user.
func (client *StandardClient) export(ctx context.Context, name string, formatParam string, options *StandardClientExportOptions) (*http.Response, error) {
	var err error
	req, err := client.exportCreateRequest(ctx, name, formatParam, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// exportCreateRequest creates the Export request.
func (client *StandardClient) exportCreateRequest(ctx context.Context, name string, formatParam string, options *StandardClientExportOptions) (*policy.Request, error) {
	urlPath := "/azure/core/lro/standard/users/{name}:export"
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", client.apiVersion)
	reqQP.Set("format", formatParam)
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}
