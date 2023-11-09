//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmachinelearning

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
)

// ClientFactory is a client factory used to create any client in this module.
// Don't use this type directly, use NewClientFactory instead.
type ClientFactory struct {
	subscriptionID string
	credential     azcore.TokenCredential
	options        *arm.ClientOptions
}

// NewClientFactory creates a new instance of ClientFactory with the specified values.
// The parameter values will be propagated to any client created from this factory.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewClientFactory(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ClientFactory, error) {
	_, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	return &ClientFactory{
		subscriptionID: subscriptionID, credential: credential,
		options: options.Clone(),
	}, nil
}

// NewBatchDeploymentsClient creates a new instance of BatchDeploymentsClient.
func (c *ClientFactory) NewBatchDeploymentsClient() *BatchDeploymentsClient {
	subClient, _ := NewBatchDeploymentsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewBatchEndpointsClient creates a new instance of BatchEndpointsClient.
func (c *ClientFactory) NewBatchEndpointsClient() *BatchEndpointsClient {
	subClient, _ := NewBatchEndpointsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewCodeContainersClient creates a new instance of CodeContainersClient.
func (c *ClientFactory) NewCodeContainersClient() *CodeContainersClient {
	subClient, _ := NewCodeContainersClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewCodeVersionsClient creates a new instance of CodeVersionsClient.
func (c *ClientFactory) NewCodeVersionsClient() *CodeVersionsClient {
	subClient, _ := NewCodeVersionsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewComponentContainersClient creates a new instance of ComponentContainersClient.
func (c *ClientFactory) NewComponentContainersClient() *ComponentContainersClient {
	subClient, _ := NewComponentContainersClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewComponentVersionsClient creates a new instance of ComponentVersionsClient.
func (c *ClientFactory) NewComponentVersionsClient() *ComponentVersionsClient {
	subClient, _ := NewComponentVersionsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewComputeClient creates a new instance of ComputeClient.
func (c *ClientFactory) NewComputeClient() *ComputeClient {
	subClient, _ := NewComputeClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewDataContainersClient creates a new instance of DataContainersClient.
func (c *ClientFactory) NewDataContainersClient() *DataContainersClient {
	subClient, _ := NewDataContainersClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewDataVersionsClient creates a new instance of DataVersionsClient.
func (c *ClientFactory) NewDataVersionsClient() *DataVersionsClient {
	subClient, _ := NewDataVersionsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewDatastoresClient creates a new instance of DatastoresClient.
func (c *ClientFactory) NewDatastoresClient() *DatastoresClient {
	subClient, _ := NewDatastoresClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewEnvironmentContainersClient creates a new instance of EnvironmentContainersClient.
func (c *ClientFactory) NewEnvironmentContainersClient() *EnvironmentContainersClient {
	subClient, _ := NewEnvironmentContainersClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewEnvironmentVersionsClient creates a new instance of EnvironmentVersionsClient.
func (c *ClientFactory) NewEnvironmentVersionsClient() *EnvironmentVersionsClient {
	subClient, _ := NewEnvironmentVersionsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewJobsClient creates a new instance of JobsClient.
func (c *ClientFactory) NewJobsClient() *JobsClient {
	subClient, _ := NewJobsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewModelContainersClient creates a new instance of ModelContainersClient.
func (c *ClientFactory) NewModelContainersClient() *ModelContainersClient {
	subClient, _ := NewModelContainersClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewModelVersionsClient creates a new instance of ModelVersionsClient.
func (c *ClientFactory) NewModelVersionsClient() *ModelVersionsClient {
	subClient, _ := NewModelVersionsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewOnlineDeploymentsClient creates a new instance of OnlineDeploymentsClient.
func (c *ClientFactory) NewOnlineDeploymentsClient() *OnlineDeploymentsClient {
	subClient, _ := NewOnlineDeploymentsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewOnlineEndpointsClient creates a new instance of OnlineEndpointsClient.
func (c *ClientFactory) NewOnlineEndpointsClient() *OnlineEndpointsClient {
	subClient, _ := NewOnlineEndpointsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewOperationsClient creates a new instance of OperationsClient.
func (c *ClientFactory) NewOperationsClient() *OperationsClient {
	subClient, _ := NewOperationsClient(c.credential, c.options)
	return subClient
}

// NewPrivateEndpointConnectionsClient creates a new instance of PrivateEndpointConnectionsClient.
func (c *ClientFactory) NewPrivateEndpointConnectionsClient() *PrivateEndpointConnectionsClient {
	subClient, _ := NewPrivateEndpointConnectionsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewPrivateLinkResourcesClient creates a new instance of PrivateLinkResourcesClient.
func (c *ClientFactory) NewPrivateLinkResourcesClient() *PrivateLinkResourcesClient {
	subClient, _ := NewPrivateLinkResourcesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewQuotasClient creates a new instance of QuotasClient.
func (c *ClientFactory) NewQuotasClient() *QuotasClient {
	subClient, _ := NewQuotasClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewUsagesClient creates a new instance of UsagesClient.
func (c *ClientFactory) NewUsagesClient() *UsagesClient {
	subClient, _ := NewUsagesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewVirtualMachineSizesClient creates a new instance of VirtualMachineSizesClient.
func (c *ClientFactory) NewVirtualMachineSizesClient() *VirtualMachineSizesClient {
	subClient, _ := NewVirtualMachineSizesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewWorkspaceConnectionsClient creates a new instance of WorkspaceConnectionsClient.
func (c *ClientFactory) NewWorkspaceConnectionsClient() *WorkspaceConnectionsClient {
	subClient, _ := NewWorkspaceConnectionsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewWorkspaceFeaturesClient creates a new instance of WorkspaceFeaturesClient.
func (c *ClientFactory) NewWorkspaceFeaturesClient() *WorkspaceFeaturesClient {
	subClient, _ := NewWorkspaceFeaturesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewWorkspacesClient creates a new instance of WorkspacesClient.
func (c *ClientFactory) NewWorkspacesClient() *WorkspacesClient {
	subClient, _ := NewWorkspacesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}
