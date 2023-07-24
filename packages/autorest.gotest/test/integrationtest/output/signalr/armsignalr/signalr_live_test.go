//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsignalr_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/internal/recording"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/internal/testutil"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/signalr/armsignalr"
	"github.com/stretchr/testify/suite"
)

type SignalRTestSuite struct {
	suite.Suite

	ctx               context.Context
	cred              azcore.TokenCredential
	options           *arm.ClientOptions
	armEndpoint       string
	globalLocation    string
	publicPort        int
	location          string
	resourceGroupName string
	subscriptionId    string
}

func (testsuite *SignalRTestSuite) SetupSuite() {
	testutil.StartRecording(testsuite.T(), "sdk/resourcemanager/signalr/armsignalr/testdata")

	testsuite.ctx = context.Background()
	testsuite.cred, testsuite.options = testutil.GetCredAndClientOptions(testsuite.T())
	testsuite.armEndpoint = "https://management.azure.com"
	testsuite.globalLocation = "Global"
	testsuite.publicPort = 6910
	testsuite.location = recording.GetEnvVariable("LOCATION", "westus")
	testsuite.resourceGroupName = recording.GetEnvVariable("RESOURCE_GROUP_NAME", "scenarioTestTempGroup")
	testsuite.subscriptionId = recording.GetEnvVariable("AZURE_SUBSCRIPTION_ID", "00000000-0000-0000-0000-000000000000")
	resourceGroup, _, err := testutil.CreateResourceGroup(testsuite.ctx, testsuite.subscriptionId, testsuite.cred, testsuite.options, testsuite.location)
	testsuite.Require().NoError(err)
	testsuite.resourceGroupName = *resourceGroup.Name
}

func (testsuite *SignalRTestSuite) TearDownSuite() {
	_, err := testutil.DeleteResourceGroup(testsuite.ctx, testsuite.subscriptionId, testsuite.cred, testsuite.options, testsuite.resourceGroupName)
	testsuite.Require().NoError(err)
	testutil.StopRecording(testsuite.T())
}

func TestSignalRTestSuite(t *testing.T) {
	suite.Run(t, new(SignalRTestSuite))
}

// Microsoft.SignalRService/Basic_CRUD
func (testsuite *SignalRTestSuite) TestSignalR() {
	var resourceName string
	var err error
	// From step Generate_Unique_Name
	template := map[string]any{
		"$schema":        "https://schema.management.azure.com/schemas/2019-04-01/deploymentTemplate.json#",
		"contentVersion": "1.0.0.0",
		"outputs": map[string]any{
			"resourceName": map[string]any{
				"type":  "string",
				"value": "[variables('name').value]",
			},
		},
		"resources": []any{},
		"variables": map[string]any{
			"name": map[string]any{
				"type": "string",
				"metadata": map[string]any{
					"description": "Name of the SignalR service.",
				},
				"value": "[concat('sw',uniqueString(resourceGroup().id))]",
			},
		},
	}
	deployment := armresources.Deployment{
		Properties: &armresources.DeploymentProperties{
			Template: template,
			Mode:     to.Ptr(armresources.DeploymentModeIncremental),
		},
	}
	deploymentExtend, err := testutil.CreateDeployment(testsuite.ctx, testsuite.subscriptionId, testsuite.cred, testsuite.options, testsuite.resourceGroupName, "Generate_Unique_Name", &deployment)
	testsuite.Require().NoError(err)
	resourceName = deploymentExtend.Properties.Outputs.(map[string]interface{})["resourceName"].(map[string]interface{})["value"].(string)

	// From step SignalR_CheckNameAvailability
	fmt.Println("Call operation: SignalR_CheckNameAvailability")
	client, err := armsignalr.NewClient(testsuite.subscriptionId, testsuite.cred, testsuite.options)
	testsuite.Require().NoError(err)
	_, err = client.CheckNameAvailability(testsuite.ctx, testsuite.location, armsignalr.NameAvailabilityParameters{
		Name: to.Ptr(resourceName),
		Type: to.Ptr("Microsoft.SignalRService/SignalR"),
	}, nil)
	testsuite.Require().NoError(err)

	// From step SignalR_CreateOrUpdate
	fmt.Println("Call operation: SignalR_CreateOrUpdate")
	clientCreateOrUpdateResponsePoller, err := client.BeginCreateOrUpdate(testsuite.ctx, testsuite.resourceGroupName, resourceName, armsignalr.ResourceInfo{
		Location: to.Ptr(testsuite.globalLocation + "-test2"),
		Tags: map[string]*string{
			"key1": to.Ptr("value1"),
		},
		Identity: &armsignalr.ManagedIdentity{
			Type: to.Ptr(armsignalr.ManagedIdentityTypeSystemAssigned),
		},
		Kind: to.Ptr(armsignalr.ServiceKindSignalR),
		Properties: &armsignalr.Properties{
			Cors: &armsignalr.CorsSettings{
				AllowedOrigins: []*string{
					to.Ptr("https://foo.com"),
					to.Ptr("https://bar.com")},
			},
			DisableAADAuth:   to.Ptr(false),
			DisableLocalAuth: to.Ptr(false),
			Features: []*armsignalr.Feature{
				{
					Flag:       to.Ptr(armsignalr.FeatureFlagsServiceMode),
					Properties: map[string]*string{},
					Value:      to.Ptr("Serverless"),
				},
				{
					Flag:       to.Ptr(armsignalr.FeatureFlagsEnableConnectivityLogs),
					Properties: map[string]*string{},
					Value:      to.Ptr("True"),
				},
				{
					Flag:       to.Ptr(armsignalr.FeatureFlagsEnableMessagingLogs),
					Properties: map[string]*string{},
					Value:      to.Ptr("False"),
				},
				{
					Flag:       to.Ptr(armsignalr.FeatureFlagsEnableLiveTrace),
					Properties: map[string]*string{},
					Value:      to.Ptr("False"),
				}},
			NetworkACLs: &armsignalr.NetworkACLs{
				DefaultAction: to.Ptr(armsignalr.ACLActionDeny),
				PrivateEndpoints: []*armsignalr.PrivateEndpointACL{
					{
						Allow: []*armsignalr.SignalRRequestType{
							to.Ptr(armsignalr.SignalRRequestTypeServerConnection)},
						Name: to.Ptr(resourceName + ".1fa229cd-bf3f-47f0-8c49-afb36723997e"),
					}},
				PublicNetwork: &armsignalr.NetworkACL{
					Allow: []*armsignalr.SignalRRequestType{
						to.Ptr(armsignalr.SignalRRequestTypeClientConnection)},
				},
			},
			PublicNetworkAccess: to.Ptr("Enabled"),
			PublicPort:          to.Ptr[int32](testsuite.publicPort),
			TLS: &armsignalr.TLSSettings{
				ClientCertEnabled: to.Ptr(false),
			},
			Upstream: &armsignalr.ServerlessUpstreamSettings{
				Templates: []*armsignalr.UpstreamTemplate{
					{
						Auth: &armsignalr.UpstreamAuthSettings{
							Type: to.Ptr(armsignalr.UpstreamAuthTypeManagedIdentity),
							ManagedIdentity: &armsignalr.ManagedIdentitySettings{
								Resource: to.Ptr("api://example"),
							},
						},
						CategoryPattern: to.Ptr("*"),
						EventPattern:    to.Ptr("connect,disconnect"),
						HubPattern:      to.Ptr("*"),
						URLTemplate:     to.Ptr("https://example.com/chat/api/connect"),
					}},
			},
		},
		SKU: &armsignalr.ResourceSKU{
			Name:     to.Ptr("Standard_S1"),
			Capacity: to.Ptr[int32](1),
			Tier:     to.Ptr(armsignalr.SignalRSKUTierStandard),
		},
	}, nil)
	testsuite.Require().NoError(err)
	var clientCreateOrUpdateResponse *armsignalr.ClientCreateOrUpdateResponse
	clientCreateOrUpdateResponse, err = testutil.PollForTest(testsuite.ctx, clientCreateOrUpdateResponsePoller)
	testsuite.Require().NoError(err)
	signalRId = *clientCreateOrUpdateResponse.ID

	// From step SignalR_Get
	fmt.Println("Call operation: SignalR_Get")
	_, err = client.Get(testsuite.ctx, testsuite.resourceGroupName, resourceName, nil)
	testsuite.Require().NoError(err)

	// From step SignalR_Update
	fmt.Println("Call operation: SignalR_Update")
	clientUpdateResponsePoller, err := client.BeginUpdate(testsuite.ctx, testsuite.resourceGroupName, resourceName, armsignalr.ResourceInfo{
		Location: to.Ptr(testsuite.location),
		Tags: map[string]*string{
			testsuite.subscriptionId: to.Ptr(testsuite.subscriptionId),
			"key1":                   to.Ptr("value1"),
		},
		Identity: &armsignalr.ManagedIdentity{
			Type: to.Ptr(armsignalr.ManagedIdentityTypeSystemAssigned),
		},
		Kind: to.Ptr(armsignalr.ServiceKindSignalR),
		Properties: &armsignalr.Properties{
			Cors: &armsignalr.CorsSettings{
				AllowedOrigins: []*string{
					to.Ptr("https://foo.com"),
					to.Ptr("https://bar.com")},
			},
			DisableAADAuth:   to.Ptr(false),
			DisableLocalAuth: to.Ptr(false),
			Features: []*armsignalr.Feature{
				{
					Flag:       to.Ptr(armsignalr.FeatureFlagsServiceMode),
					Properties: map[string]*string{},
					Value:      to.Ptr("Serverless"),
				},
				{
					Flag:       to.Ptr(armsignalr.FeatureFlagsEnableConnectivityLogs),
					Properties: map[string]*string{},
					Value:      to.Ptr("True"),
				},
				{
					Flag:       to.Ptr(armsignalr.FeatureFlagsEnableMessagingLogs),
					Properties: map[string]*string{},
					Value:      to.Ptr("False"),
				},
				{
					Flag:       to.Ptr(armsignalr.FeatureFlagsEnableLiveTrace),
					Properties: map[string]*string{},
					Value:      to.Ptr("False"),
				}},
			NetworkACLs: &armsignalr.NetworkACLs{
				DefaultAction: to.Ptr(armsignalr.ACLActionDeny),
				PrivateEndpoints: []*armsignalr.PrivateEndpointACL{
					{
						Allow: []*armsignalr.SignalRRequestType{
							to.Ptr(armsignalr.SignalRRequestTypeServerConnection)},
						Name: to.Ptr(resourceName + ".1fa229cd-bf3f-47f0-8c49-afb36723997e"),
					}},
				PublicNetwork: &armsignalr.NetworkACL{
					Allow: []*armsignalr.SignalRRequestType{
						to.Ptr(armsignalr.SignalRRequestTypeClientConnection)},
				},
			},
			PublicNetworkAccess: to.Ptr("Enabled"),
			TLS: &armsignalr.TLSSettings{
				ClientCertEnabled: to.Ptr(false),
			},
			Upstream: &armsignalr.ServerlessUpstreamSettings{
				Templates: []*armsignalr.UpstreamTemplate{
					{
						Auth: &armsignalr.UpstreamAuthSettings{
							Type: to.Ptr(armsignalr.UpstreamAuthTypeManagedIdentity),
							ManagedIdentity: &armsignalr.ManagedIdentitySettings{
								Resource: to.Ptr("api://example"),
							},
						},
						CategoryPattern: to.Ptr("*"),
						EventPattern:    to.Ptr("connect,disconnect"),
						HubPattern:      to.Ptr("*"),
						URLTemplate:     to.Ptr("https://example.com/chat/api/connect"),
					}},
			},
		},
		SKU: &armsignalr.ResourceSKU{
			Name:     to.Ptr("Standard_S1"),
			Capacity: to.Ptr[int32](1),
			Tier:     to.Ptr(armsignalr.SignalRSKUTierStandard),
		},
	}, nil)
	testsuite.Require().NoError(err)
	_, err = testutil.PollForTest(testsuite.ctx, clientUpdateResponsePoller)
	testsuite.Require().NoError(err)

	// From step SignalR_ListKeys
	fmt.Println("Call operation: SignalR_ListKeys")
	_, err = client.ListKeys(testsuite.ctx, testsuite.resourceGroupName, resourceName, nil)
	testsuite.Require().NoError(err)

	// From step SignalR_RegenerateKey
	fmt.Println("Call operation: SignalR_RegenerateKey")
	clientRegenerateKeyResponsePoller, err := client.BeginRegenerateKey(testsuite.ctx, testsuite.resourceGroupName, resourceName, armsignalr.RegenerateKeyParameters{
		KeyType: to.Ptr(armsignalr.KeyTypePrimary),
	}, nil)
	testsuite.Require().NoError(err)
	_, err = testutil.PollForTest(testsuite.ctx, clientRegenerateKeyResponsePoller)
	testsuite.Require().NoError(err)

	// From step SignalR_Restart
	fmt.Println("Call operation: SignalR_Restart")
	clientRestartResponsePoller, err := client.BeginRestart(testsuite.ctx, testsuite.resourceGroupName, resourceName, nil)
	testsuite.Require().NoError(err)
	_, err = testutil.PollForTest(testsuite.ctx, clientRestartResponsePoller)
	testsuite.Require().NoError(err)

	// From step Usages_List
	fmt.Println("Call operation: Usages_List")
	usagesClient, err := armsignalr.NewUsagesClient(testsuite.subscriptionId, testsuite.cred, testsuite.options)
	testsuite.Require().NoError(err)
	usagesClientNewListPager := usagesClient.NewListPager(testsuite.location, nil)
	for usagesClientNewListPager.More() {
		_, err := usagesClientNewListPager.NextPage(testsuite.ctx)
		testsuite.Require().NoError(err)
		break
	}

	// From step SignalR_ListByResourceGroup
	fmt.Println("Call operation: SignalR_ListByResourceGroup")
	clientNewListByResourceGroupPager := client.NewListByResourceGroupPager(testsuite.resourceGroupName, nil)
	for clientNewListByResourceGroupPager.More() {
		_, err := clientNewListByResourceGroupPager.NextPage(testsuite.ctx)
		testsuite.Require().NoError(err)
		break
	}

	// From step SignalR_ListBySubscription
	fmt.Println("Call operation: SignalR_ListBySubscription")
	clientNewListBySubscriptionPager := client.NewListBySubscriptionPager(nil)
	for clientNewListBySubscriptionPager.More() {
		_, err := clientNewListBySubscriptionPager.NextPage(testsuite.ctx)
		testsuite.Require().NoError(err)
		break
	}

	// From step Operations_List
	fmt.Println("Call operation: Operations_List")
	operationsClient, err := armsignalr.NewOperationsClient(testsuite.cred, testsuite.options)
	testsuite.Require().NoError(err)
	operationsClientNewListPager := operationsClient.NewListPager(nil)
	for operationsClientNewListPager.More() {
		_, err := operationsClientNewListPager.NextPage(testsuite.ctx)
		testsuite.Require().NoError(err)
		break
	}

	// From step SignalR_Delete
	fmt.Println("Call operation: SignalR_Delete")
	clientDeleteResponsePoller, err := client.BeginDelete(testsuite.ctx, testsuite.resourceGroupName, resourceName, nil)
	testsuite.Require().NoError(err)
	_, err = testutil.PollForTest(testsuite.ctx, clientDeleteResponsePoller)
	testsuite.Require().NoError(err)
}
