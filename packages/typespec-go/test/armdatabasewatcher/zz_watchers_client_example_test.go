// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package armdatabasewatcher_test

import (
	"armdatabasewatcher"
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"log"
)

// Generated from example definition: 2023-09-01-preview/Watchers_CreateOrUpdate_200.json
func ExampleWatchersClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatabasewatcher.NewClientFactory("6f53185c-ea09-4fc3-9075-318dec805303", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewWatchersClient().BeginCreateOrUpdate(ctx, "apiTest-ddat4p", "databasemo3ej9ih", armdatabasewatcher.Watcher{
		Location: to.Ptr("westus"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armdatabasewatcher.WatchersClientCreateOrUpdateResponse{
	// 	Watcher: &armdatabasewatcher.Watcher{
	// 		ID: to.Ptr("/subscriptions/6f53185c-ea09-4fc3-9075-318dec805303/resourceGroups/apiTest-ddat4p/providers/Microsoft.DatabaseWatcher/watchers/databasemo3ej9ih"),
	// 		Name: to.Ptr("databasemo3ej9ih"),
	// 		Type: to.Ptr("microsoft.databasewatcher/watchers"),
	// 		Location: to.Ptr("westus"),
	// 		Properties: &armdatabasewatcher.WatcherProperties{
	// 			ProvisioningState: to.Ptr(armdatabasewatcher.DatabaseWatcherProvisioningStateSucceeded),
	// 		},
	// 	},
	// }
}

// Generated from example definition: 2023-09-01-preview/Watchers_Delete_200.json
func ExampleWatchersClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatabasewatcher.NewClientFactory("6f53185c-ea09-4fc3-9075-318dec805303", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewWatchersClient().BeginDelete(ctx, "apiTest-ddat4p", "databasemo3ej9ih", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: 2023-09-01-preview/Watchers_Get_200.json
func ExampleWatchersClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatabasewatcher.NewClientFactory("6f53185c-ea09-4fc3-9075-318dec805303", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewWatchersClient().Get(ctx, "apiTest-ddat4p", "databasemo3ej9ih", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armdatabasewatcher.WatchersClientGetResponse{
	// 	Watcher: &armdatabasewatcher.Watcher{
	// 		ID: to.Ptr("/subscriptions/6f53185c-ea09-4fc3-9075-318dec805303/resourceGroups/apiTest-ddat4p/providers/Microsoft.DatabaseWatcher/watchers/databasemo3ej9ih"),
	// 		Name: to.Ptr("databasemo3ej9ih"),
	// 		Type: to.Ptr("microsoft.databasewatcher/watchers"),
	// 		Location: to.Ptr("westus"),
	// 		Properties: &armdatabasewatcher.WatcherProperties{
	// 			ProvisioningState: to.Ptr(armdatabasewatcher.DatabaseWatcherProvisioningStateSucceeded),
	// 		},
	// 	},
	// }
}

// Generated from example definition: 2023-09-01-preview/Watchers_ListByResourceGroup_200.json
func ExampleWatchersClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatabasewatcher.NewClientFactory("6f53185c-ea09-4fc3-9075-318dec805303", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewWatchersClient().NewListByResourceGroupPager("apiTest-ddat4p", nil)
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
		// page = armdatabasewatcher.WatchersClientListByResourceGroupResponse{
		// 	WatcherListResult: armdatabasewatcher.WatcherListResult{
		// 		Value: []*armdatabasewatcher.Watcher{
		// 			{
		// 				ID: to.Ptr("/subscriptions/6f53185c-ea09-4fc3-9075-318dec805303/resourceGroups/apiTest-ddat4p/providers/Microsoft.DatabaseWatcher/watchers/databasemo3ej9ih"),
		// 				Name: to.Ptr("databasemo3ej9ih"),
		// 				Type: to.Ptr("microsoft.databasewatcher/watchers"),
		// 				Location: to.Ptr("westus"),
		// 				Properties: &armdatabasewatcher.WatcherProperties{
		// 					ProvisioningState: to.Ptr(armdatabasewatcher.DatabaseWatcherProvisioningStateSucceeded),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}

// Generated from example definition: 2023-09-01-preview/Watchers_ListBySubscription_200.json
func ExampleWatchersClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatabasewatcher.NewClientFactory("6f53185c-ea09-4fc3-9075-318dec805303", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewWatchersClient().NewListBySubscriptionPager(nil)
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
		// page = armdatabasewatcher.WatchersClientListBySubscriptionResponse{
		// 	WatcherListResult: armdatabasewatcher.WatcherListResult{
		// 		Value: []*armdatabasewatcher.Watcher{
		// 			{
		// 				ID: to.Ptr("/subscriptions/6f53185c-ea09-4fc3-9075-318dec805303/resourceGroups/dummyrg/providers/Microsoft.DatabaseWatcher/watchers/myWatcher"),
		// 				Name: to.Ptr("myWatcher"),
		// 				Type: to.Ptr("microsoft.databasewatcher/watchers"),
		// 				Location: to.Ptr("West US"),
		// 				Tags: map[string]*string{
		// 				},
		// 				Identity: &armdatabasewatcher.ManagedServiceIdentity{
		// 					Type: to.Ptr(armdatabasewatcher.ManagedServiceIdentityTypeSystemAssigned),
		// 				},
		// 				Properties: &armdatabasewatcher.WatcherProperties{
		// 					ProvisioningState: to.Ptr(armdatabasewatcher.DatabaseWatcherProvisioningStateSucceeded),
		// 				},
		// 			},
		// 			{
		// 				ID: to.Ptr("/subscriptions/6f53185c-ea09-4fc3-9075-318dec805303/resourceGroups/apiTest-4sz1jg/providers/Microsoft.DatabaseWatcher/watchers/databasemo4o4zdf"),
		// 				Name: to.Ptr("databasemo4o4zdf"),
		// 				Type: to.Ptr("microsoft.databasewatcher/watchers"),
		// 				Location: to.Ptr("westus"),
		// 				Properties: &armdatabasewatcher.WatcherProperties{
		// 					ProvisioningState: to.Ptr(armdatabasewatcher.DatabaseWatcherProvisioningStateSucceeded),
		// 				},
		// 			},
		// 			{
		// 				ID: to.Ptr("/subscriptions/6f53185c-ea09-4fc3-9075-318dec805303/resourceGroups/apiTest-d77ftn/providers/Microsoft.DatabaseWatcher/watchers/databasemosn3h6l"),
		// 				Name: to.Ptr("databasemosn3h6l"),
		// 				Type: to.Ptr("microsoft.databasewatcher/watchers"),
		// 				Location: to.Ptr("westus"),
		// 				Properties: &armdatabasewatcher.WatcherProperties{
		// 					ProvisioningState: to.Ptr(armdatabasewatcher.DatabaseWatcherProvisioningStateSucceeded),
		// 				},
		// 			},
		// 			{
		// 				ID: to.Ptr("/subscriptions/6f53185c-ea09-4fc3-9075-318dec805303/resourceGroups/apiTest-nyb4gm/providers/Microsoft.DatabaseWatcher/watchers/databasemoyb6iar"),
		// 				Name: to.Ptr("databasemoyb6iar"),
		// 				Type: to.Ptr("microsoft.databasewatcher/watchers"),
		// 				Location: to.Ptr("westus"),
		// 				Properties: &armdatabasewatcher.WatcherProperties{
		// 					ProvisioningState: to.Ptr(armdatabasewatcher.DatabaseWatcherProvisioningStateSucceeded),
		// 				},
		// 			},
		// 			{
		// 				ID: to.Ptr("/subscriptions/6f53185c-ea09-4fc3-9075-318dec805303/resourceGroups/apiTest-llis4j/providers/Microsoft.DatabaseWatcher/watchers/databasemoi04xst"),
		// 				Name: to.Ptr("databasemoi04xst"),
		// 				Type: to.Ptr("microsoft.databasewatcher/watchers"),
		// 				Location: to.Ptr("westus"),
		// 				Properties: &armdatabasewatcher.WatcherProperties{
		// 					ProvisioningState: to.Ptr(armdatabasewatcher.DatabaseWatcherProvisioningStateSucceeded),
		// 				},
		// 			},
		// 			{
		// 				ID: to.Ptr("/subscriptions/6f53185c-ea09-4fc3-9075-318dec805303/resourceGroups/apiTest-thy6zd/providers/Microsoft.DatabaseWatcher/watchers/databasemonpyl24"),
		// 				Name: to.Ptr("databasemonpyl24"),
		// 				Type: to.Ptr("microsoft.databasewatcher/watchers"),
		// 				Location: to.Ptr("westus"),
		// 				SystemData: &armdatabasewatcher.SystemData{
		// 				},
		// 				Properties: &armdatabasewatcher.WatcherProperties{
		// 					ProvisioningState: to.Ptr(armdatabasewatcher.DatabaseWatcherProvisioningStateSucceeded),
		// 				},
		// 			},
		// 			{
		// 				ID: to.Ptr("/subscriptions/6f53185c-ea09-4fc3-9075-318dec805303/resourceGroups/apiTest-ddat4p/providers/Microsoft.DatabaseWatcher/watchers/databasemo3ej9ih"),
		// 				Name: to.Ptr("databasemo3ej9ih"),
		// 				Type: to.Ptr("microsoft.databasewatcher/watchers"),
		// 				Location: to.Ptr("westus"),
		// 				Properties: &armdatabasewatcher.WatcherProperties{
		// 					ProvisioningState: to.Ptr(armdatabasewatcher.DatabaseWatcherProvisioningStateSucceeded),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}

// Generated from example definition: 2023-09-01-preview/Watchers_Start_200.json
func ExampleWatchersClient_BeginStart() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatabasewatcher.NewClientFactory("6f53185c-ea09-4fc3-9075-318dec805303", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewWatchersClient().BeginStart(ctx, "apiTest-ddat4p", "databasemo3ej9ih", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armdatabasewatcher.WatchersClientStartResponse{
	// 	Watcher: &armdatabasewatcher.Watcher{
	// 		ID: to.Ptr("/subscriptions/6f53185c-ea09-4fc3-9075-318dec805303/resourceGroups/apiTest-ddat4p/providers/Microsoft.DatabaseWatcher/watchers/databasemo3ej9ih"),
	// 		Name: to.Ptr("databasemo3ej9ih"),
	// 		Type: to.Ptr("microsoft.databasewatcher/watchers"),
	// 		Location: to.Ptr("westus"),
	// 		Properties: &armdatabasewatcher.WatcherProperties{
	// 			ProvisioningState: to.Ptr(armdatabasewatcher.DatabaseWatcherProvisioningStateSucceeded),
	// 		},
	// 	},
	// }
}

// Generated from example definition: 2023-09-01-preview/Watchers_Stop_200.json
func ExampleWatchersClient_BeginStop() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatabasewatcher.NewClientFactory("6f53185c-ea09-4fc3-9075-318dec805303", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewWatchersClient().BeginStop(ctx, "apiTest-ddat4p", "databasemo3ej9ih", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armdatabasewatcher.WatchersClientStopResponse{
	// 	Watcher: &armdatabasewatcher.Watcher{
	// 		ID: to.Ptr("/subscriptions/6f53185c-ea09-4fc3-9075-318dec805303/resourceGroups/apiTest-ddat4p/providers/Microsoft.DatabaseWatcher/watchers/databasemo3ej9ih"),
	// 		Name: to.Ptr("databasemo3ej9ih"),
	// 		Type: to.Ptr("microsoft.databasewatcher/watchers"),
	// 		Location: to.Ptr("westus"),
	// 		Properties: &armdatabasewatcher.WatcherProperties{
	// 			ProvisioningState: to.Ptr(armdatabasewatcher.DatabaseWatcherProvisioningStateSucceeded),
	// 		},
	// 	},
	// }
}

// Generated from example definition: 2023-09-01-preview/Watchers_Update_200.json
func ExampleWatchersClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatabasewatcher.NewClientFactory("6f53185c-ea09-4fc3-9075-318dec805303", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewWatchersClient().BeginUpdate(ctx, "apiTest-lzect6", "databasemohkp6fn", armdatabasewatcher.Watcher{
		Identity: &armdatabasewatcher.ManagedServiceIdentity{
			Type: to.Ptr(armdatabasewatcher.ManagedServiceIdentityTypeSystemAssigned),
		},
		Properties: &armdatabasewatcher.WatcherProperties{
			Datastore: &armdatabasewatcher.Datastore{
				AdxClusterResourceID:    to.Ptr("/subscriptions/6f53185c-ea09-4fc3-9075-318dec805303/resourceGroups/apiTest/providers/Microsoft.Kusto/clusters/apiTestKusto"),
				KustoClusterURI:         to.Ptr("https://kustouri-adx.eastus.kusto.windows.net"),
				KustoClusterDisplayName: to.Ptr("kustoUri-adx"),
				KustoDataIngestionURI:   to.Ptr("https://ingest-kustouri-adx.eastus.kusto.windows.net"),
				KustoDatabaseName:       to.Ptr("kustoDatabaseName1"),
				KustoManagementURL:      to.Ptr("https://portal.azure.com/"),
				KustoOfferingType:       to.Ptr(armdatabasewatcher.KustoOfferingTypeAdx),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armdatabasewatcher.WatchersClientUpdateResponse{
	// 	Watcher: &armdatabasewatcher.Watcher{
	// 		ID: to.Ptr("/subscriptions/6f53185c-ea09-4fc3-9075-318dec805303/resourceGroups/apiTest-lzect6/providers/Microsoft.DatabaseWatcher/watchers/databasemohkp6fn"),
	// 		Name: to.Ptr("databasemohkp6fn"),
	// 		Type: to.Ptr("microsoft.databasewatcher/watchers"),
	// 		Location: to.Ptr("westus"),
	// 		SystemData: &armdatabasewatcher.SystemData{
	// 		},
	// 		Identity: &armdatabasewatcher.ManagedServiceIdentity{
	// 			Type: to.Ptr(armdatabasewatcher.ManagedServiceIdentityTypeSystemAssigned),
	// 		},
	// 		Properties: &armdatabasewatcher.WatcherProperties{
	// 			ProvisioningState: to.Ptr(armdatabasewatcher.DatabaseWatcherProvisioningStateSucceeded),
	// 			Datastore: &armdatabasewatcher.Datastore{
	// 				AdxClusterResourceID: to.Ptr("/subscriptions/6f53185c-ea09-4fc3-9075-318dec805303/resourceGroups/apiTest/providers/Microsoft.Kusto/clusters/apiTestKusto"),
	// 				KustoClusterURI: to.Ptr("https://kustouri-adx.eastus.kusto.windows.net"),
	// 				KustoClusterDisplayName: to.Ptr("kustoUri-adx"),
	// 				KustoDataIngestionURI: to.Ptr("https://ingest-kustouri-adx.eastus.kusto.windows.net"),
	// 				KustoDatabaseName: to.Ptr("kustoDatabaseName1"),
	// 				KustoManagementURL: to.Ptr("https://portal.azure.com/"),
	// 				KustoOfferingType: to.Ptr(armdatabasewatcher.KustoOfferingTypeAdx),
	// 			},
	// 		},
	// 	},
	// }
}
