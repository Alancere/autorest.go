// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azartifacts

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"time"
)

type dataFlowDebugSessionClient struct {
	con *connection
}

// AddDataFlow - Add a data flow into debug session.
func (client *dataFlowDebugSessionClient) AddDataFlow(ctx context.Context, request DataFlowDebugPackage, options *DataFlowDebugSessionAddDataFlowOptions) (AddDataFlowToDebugSessionResponseResponse, error) {
	req, err := client.addDataFlowCreateRequest(ctx, request, options)
	if err != nil {
		return AddDataFlowToDebugSessionResponseResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return AddDataFlowToDebugSessionResponseResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return AddDataFlowToDebugSessionResponseResponse{}, client.addDataFlowHandleError(resp)
	}
	return client.addDataFlowHandleResponse(resp)
}

// addDataFlowCreateRequest creates the AddDataFlow request.
func (client *dataFlowDebugSessionClient) addDataFlowCreateRequest(ctx context.Context, request DataFlowDebugPackage, options *DataFlowDebugSessionAddDataFlowOptions) (*azcore.Request, error) {
	urlPath := "/addDataFlowToDebugSession"
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(request)
}

// addDataFlowHandleResponse handles the AddDataFlow response.
func (client *dataFlowDebugSessionClient) addDataFlowHandleResponse(resp *azcore.Response) (AddDataFlowToDebugSessionResponseResponse, error) {
	var val *AddDataFlowToDebugSessionResponse
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return AddDataFlowToDebugSessionResponseResponse{}, err
	}
	return AddDataFlowToDebugSessionResponseResponse{RawResponse: resp.Response, AddDataFlowToDebugSessionResponse: val}, nil
}

// addDataFlowHandleError handles the AddDataFlow error response.
func (client *dataFlowDebugSessionClient) addDataFlowHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err.InnerError); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// BeginCreateDataFlowDebugSession - Creates a data flow debug session.
func (client *dataFlowDebugSessionClient) BeginCreateDataFlowDebugSession(ctx context.Context, request CreateDataFlowDebugSessionRequest, options *DataFlowDebugSessionBeginCreateDataFlowDebugSessionOptions) (CreateDataFlowDebugSessionResponsePollerResponse, error) {
	resp, err := client.createDataFlowDebugSession(ctx, request, options)
	if err != nil {
		return CreateDataFlowDebugSessionResponsePollerResponse{}, err
	}
	result := CreateDataFlowDebugSessionResponsePollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := azcore.NewLROPoller("dataFlowDebugSessionClient.CreateDataFlowDebugSession", resp, client.con.Pipeline(), client.createDataFlowDebugSessionHandleError)
	if err != nil {
		return CreateDataFlowDebugSessionResponsePollerResponse{}, err
	}
	poller := &createDataFlowDebugSessionResponsePoller{
		pt: pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (CreateDataFlowDebugSessionResponseResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeCreateDataFlowDebugSession creates a new CreateDataFlowDebugSessionResponsePoller from the specified resume token.
// token - The value must come from a previous call to CreateDataFlowDebugSessionResponsePoller.ResumeToken().
func (client *dataFlowDebugSessionClient) ResumeCreateDataFlowDebugSession(ctx context.Context, token string) (CreateDataFlowDebugSessionResponsePollerResponse, error) {
	pt, err := azcore.NewLROPollerFromResumeToken("dataFlowDebugSessionClient.CreateDataFlowDebugSession", token, client.con.Pipeline(), client.createDataFlowDebugSessionHandleError)
	if err != nil {
		return CreateDataFlowDebugSessionResponsePollerResponse{}, err
	}
	poller := &createDataFlowDebugSessionResponsePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return CreateDataFlowDebugSessionResponsePollerResponse{}, err
	}
	result := CreateDataFlowDebugSessionResponsePollerResponse{
		RawResponse: resp,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (CreateDataFlowDebugSessionResponseResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// CreateDataFlowDebugSession - Creates a data flow debug session.
func (client *dataFlowDebugSessionClient) createDataFlowDebugSession(ctx context.Context, request CreateDataFlowDebugSessionRequest, options *DataFlowDebugSessionBeginCreateDataFlowDebugSessionOptions) (*azcore.Response, error) {
	req, err := client.createDataFlowDebugSessionCreateRequest(ctx, request, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.createDataFlowDebugSessionHandleError(resp)
	}
	return resp, nil
}

// createDataFlowDebugSessionCreateRequest creates the CreateDataFlowDebugSession request.
func (client *dataFlowDebugSessionClient) createDataFlowDebugSessionCreateRequest(ctx context.Context, request CreateDataFlowDebugSessionRequest, options *DataFlowDebugSessionBeginCreateDataFlowDebugSessionOptions) (*azcore.Request, error) {
	urlPath := "/createDataFlowDebugSession"
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(request)
}

// createDataFlowDebugSessionHandleResponse handles the CreateDataFlowDebugSession response.
func (client *dataFlowDebugSessionClient) createDataFlowDebugSessionHandleResponse(resp *azcore.Response) (CreateDataFlowDebugSessionResponseResponse, error) {
	var val *CreateDataFlowDebugSessionResponse
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return CreateDataFlowDebugSessionResponseResponse{}, err
	}
	return CreateDataFlowDebugSessionResponseResponse{RawResponse: resp.Response, CreateDataFlowDebugSessionResponse: val}, nil
}

// createDataFlowDebugSessionHandleError handles the CreateDataFlowDebugSession error response.
func (client *dataFlowDebugSessionClient) createDataFlowDebugSessionHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err.InnerError); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// DeleteDataFlowDebugSession - Deletes a data flow debug session.
func (client *dataFlowDebugSessionClient) DeleteDataFlowDebugSession(ctx context.Context, request DeleteDataFlowDebugSessionRequest, options *DataFlowDebugSessionDeleteDataFlowDebugSessionOptions) (*http.Response, error) {
	req, err := client.deleteDataFlowDebugSessionCreateRequest(ctx, request, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.deleteDataFlowDebugSessionHandleError(resp)
	}
	return resp.Response, nil
}

// deleteDataFlowDebugSessionCreateRequest creates the DeleteDataFlowDebugSession request.
func (client *dataFlowDebugSessionClient) deleteDataFlowDebugSessionCreateRequest(ctx context.Context, request DeleteDataFlowDebugSessionRequest, options *DataFlowDebugSessionDeleteDataFlowDebugSessionOptions) (*azcore.Request, error) {
	urlPath := "/deleteDataFlowDebugSession"
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(request)
}

// deleteDataFlowDebugSessionHandleError handles the DeleteDataFlowDebugSession error response.
func (client *dataFlowDebugSessionClient) deleteDataFlowDebugSessionHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err.InnerError); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// BeginExecuteCommand - Execute a data flow debug command.
func (client *dataFlowDebugSessionClient) BeginExecuteCommand(ctx context.Context, request DataFlowDebugCommandRequest, options *DataFlowDebugSessionBeginExecuteCommandOptions) (DataFlowDebugCommandResponsePollerResponse, error) {
	resp, err := client.executeCommand(ctx, request, options)
	if err != nil {
		return DataFlowDebugCommandResponsePollerResponse{}, err
	}
	result := DataFlowDebugCommandResponsePollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := azcore.NewLROPoller("dataFlowDebugSessionClient.ExecuteCommand", resp, client.con.Pipeline(), client.executeCommandHandleError)
	if err != nil {
		return DataFlowDebugCommandResponsePollerResponse{}, err
	}
	poller := &dataFlowDebugCommandResponsePoller{
		pt: pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (DataFlowDebugCommandResponseResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeExecuteCommand creates a new DataFlowDebugCommandResponsePoller from the specified resume token.
// token - The value must come from a previous call to DataFlowDebugCommandResponsePoller.ResumeToken().
func (client *dataFlowDebugSessionClient) ResumeExecuteCommand(ctx context.Context, token string) (DataFlowDebugCommandResponsePollerResponse, error) {
	pt, err := azcore.NewLROPollerFromResumeToken("dataFlowDebugSessionClient.ExecuteCommand", token, client.con.Pipeline(), client.executeCommandHandleError)
	if err != nil {
		return DataFlowDebugCommandResponsePollerResponse{}, err
	}
	poller := &dataFlowDebugCommandResponsePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return DataFlowDebugCommandResponsePollerResponse{}, err
	}
	result := DataFlowDebugCommandResponsePollerResponse{
		RawResponse: resp,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (DataFlowDebugCommandResponseResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ExecuteCommand - Execute a data flow debug command.
func (client *dataFlowDebugSessionClient) executeCommand(ctx context.Context, request DataFlowDebugCommandRequest, options *DataFlowDebugSessionBeginExecuteCommandOptions) (*azcore.Response, error) {
	req, err := client.executeCommandCreateRequest(ctx, request, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.executeCommandHandleError(resp)
	}
	return resp, nil
}

// executeCommandCreateRequest creates the ExecuteCommand request.
func (client *dataFlowDebugSessionClient) executeCommandCreateRequest(ctx context.Context, request DataFlowDebugCommandRequest, options *DataFlowDebugSessionBeginExecuteCommandOptions) (*azcore.Request, error) {
	urlPath := "/executeDataFlowDebugCommand"
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(request)
}

// executeCommandHandleResponse handles the ExecuteCommand response.
func (client *dataFlowDebugSessionClient) executeCommandHandleResponse(resp *azcore.Response) (DataFlowDebugCommandResponseResponse, error) {
	var val *DataFlowDebugCommandResponse
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return DataFlowDebugCommandResponseResponse{}, err
	}
	return DataFlowDebugCommandResponseResponse{RawResponse: resp.Response, DataFlowDebugCommandResponse: val}, nil
}

// executeCommandHandleError handles the ExecuteCommand error response.
func (client *dataFlowDebugSessionClient) executeCommandHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err.InnerError); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// QueryDataFlowDebugSessionsByWorkspace - Query all active data flow debug sessions.
func (client *dataFlowDebugSessionClient) QueryDataFlowDebugSessionsByWorkspace(options *DataFlowDebugSessionQueryDataFlowDebugSessionsByWorkspaceOptions) QueryDataFlowDebugSessionsResponsePager {
	return &queryDataFlowDebugSessionsResponsePager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.queryDataFlowDebugSessionsByWorkspaceCreateRequest(ctx, options)
		},
		responder: client.queryDataFlowDebugSessionsByWorkspaceHandleResponse,
		errorer:   client.queryDataFlowDebugSessionsByWorkspaceHandleError,
		advancer: func(ctx context.Context, resp QueryDataFlowDebugSessionsResponseResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.QueryDataFlowDebugSessionsResponse.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// queryDataFlowDebugSessionsByWorkspaceCreateRequest creates the QueryDataFlowDebugSessionsByWorkspace request.
func (client *dataFlowDebugSessionClient) queryDataFlowDebugSessionsByWorkspaceCreateRequest(ctx context.Context, options *DataFlowDebugSessionQueryDataFlowDebugSessionsByWorkspaceOptions) (*azcore.Request, error) {
	urlPath := "/queryDataFlowDebugSessions"
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// queryDataFlowDebugSessionsByWorkspaceHandleResponse handles the QueryDataFlowDebugSessionsByWorkspace response.
func (client *dataFlowDebugSessionClient) queryDataFlowDebugSessionsByWorkspaceHandleResponse(resp *azcore.Response) (QueryDataFlowDebugSessionsResponseResponse, error) {
	var val *QueryDataFlowDebugSessionsResponse
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return QueryDataFlowDebugSessionsResponseResponse{}, err
	}
	return QueryDataFlowDebugSessionsResponseResponse{RawResponse: resp.Response, QueryDataFlowDebugSessionsResponse: val}, nil
}

// queryDataFlowDebugSessionsByWorkspaceHandleError handles the QueryDataFlowDebugSessionsByWorkspace error response.
func (client *dataFlowDebugSessionClient) queryDataFlowDebugSessionsByWorkspaceHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err.InnerError); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}