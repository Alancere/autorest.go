// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package armdatabasewatcher_test

import (
	"armdatabasewatcher"
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"log"
)

// Generated from example definition: 2024-07-19-preview/Operations_List_MaximumSet_Gen.json
func ExampleOperationsClient_NewListPager_operationsListGeneratedByMaximumSetRule() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatabasewatcher.NewClientFactory("<subscriptionID>", cred, nil)
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
		// page = armdatabasewatcher.OperationsClientListResponse{
		// 	OperationListResult: armdatabasewatcher.OperationListResult{
		// 		Value: []*armdatabasewatcher.Operation{
		// 			{
		// 				Name: to.Ptr("snzrdvltunnrz"),
		// 				IsDataAction: to.Ptr(true),
		// 				Display: &armdatabasewatcher.OperationDisplay{
		// 					Provider: to.Ptr("dtfrqzamclscchmghtxn"),
		// 					Resource: to.Ptr("lvlhnsfnquorjuuutjxex"),
		// 					Operation: to.Ptr("vbgvamoxqwthpbdghxzaw"),
		// 					Description: to.Ptr("nvbtuwwjfehylzmoatd"),
		// 				},
		// 				Origin: to.Ptr(armdatabasewatcher.OriginUser),
		// 				ActionType: to.Ptr(armdatabasewatcher.ActionTypeInternal),
		// 			},
		// 		},
		// 		NextLink: to.Ptr("https://microsoft.com/awfba"),
		// 	},
		// }
	}
}

// Generated from example definition: 2024-07-19-preview/Operations_List_MinimumSet_Gen.json
func ExampleOperationsClient_NewListPager_operationsListGeneratedByMinimumSetRule() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatabasewatcher.NewClientFactory("<subscriptionID>", cred, nil)
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
		// page = armdatabasewatcher.OperationsClientListResponse{
		// 	OperationListResult: armdatabasewatcher.OperationListResult{
		// 	},
		// }
	}
}
