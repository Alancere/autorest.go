//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armappplatform_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appplatform/armappplatform/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/main/specification/appplatform/resource-manager/Microsoft.AppPlatform/preview/2020-11-01-preview/examples/Operations_List.json
func ExampleOperationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappplatform.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewOperationsClient().NewListPager(nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.AvailableOperations = armappplatform.AvailableOperations{
		// 	Value: []*armappplatform.OperationDetail{
		// 		{
		// 			Name: to.Ptr("Microsoft.AppPlatform/Spring/read"),
		// 			Display: &armappplatform.OperationDisplay{
		// 				Description: to.Ptr("Create or Update Managed Applications"),
		// 				Operation: to.Ptr("Create or Update Managed Applications"),
		// 				Provider: to.Ptr("Microsoft Azure Distributed Managed Service for Spring"),
		// 				Resource: to.Ptr("Managed Applications"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 			Origin: to.Ptr("user,system"),
		// 			Properties: &armappplatform.OperationProperties{
		// 			},
		// 	}},
		// }
	}
}
