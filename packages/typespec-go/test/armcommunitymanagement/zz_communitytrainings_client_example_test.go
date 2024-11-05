// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package armcommunitymanagement_test

import (
	"armcommunitymanagement"
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"log"
)

// Generated from example definition: 2023-11-01/CommunityTrainings_Create.json
func ExampleCommunityTrainingsClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcommunitymanagement.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewCommunityTrainingsClient().BeginCreate(ctx, "rgCommunityTaining", "ctApplication", armcommunitymanagement.CommunityTraining{
		Properties: &armcommunitymanagement.CommunityTrainingProperties{
			PortalName:                  to.Ptr("ctwebsite"),
			PortalAdminEmailAddress:     to.Ptr("ctadmin@ct.com"),
			PortalOwnerOrganizationName: to.Ptr("CT Portal Owner Organization"),
			PortalOwnerEmailAddress:     to.Ptr("ctcontact@ct.com"),
			IdentityConfiguration: &armcommunitymanagement.IdentityConfigurationProperties{
				IdentityType:            to.Ptr("ADB2C"),
				TeamsEnabled:            to.Ptr(false),
				TenantID:                to.Ptr("c1ffbb60-88cf-4b83-b54f-c47ae6220c19"),
				DomainName:              to.Ptr("cttenant"),
				ClientID:                to.Ptr("8c92390f-2f30-493d-bd13-d3c3eba3709d"),
				ClientSecret:            to.Ptr("idenityConfigurationClientSecret"),
				B2CAuthenticationPolicy: to.Ptr("B2C_1_signup_signin"),
				B2CPasswordResetPolicy:  to.Ptr("B2C_1_pwd_reset"),
				CustomLoginParameters:   to.Ptr("custom_hint"),
			},
			ZoneRedundancyEnabled:   to.Ptr(true),
			DisasterRecoveryEnabled: to.Ptr(true),
		},
		SKU: &armcommunitymanagement.SKU{
			Name: to.Ptr("Commercial"),
			Tier: to.Ptr(armcommunitymanagement.SKUTierStandard),
		},
		Location: to.Ptr("southeastasia"),
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
	// res = armcommunitymanagement.CommunityTrainingsClientCreateResponse{
	// 	CommunityTraining: &armcommunitymanagement.CommunityTraining{
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rgCommunityTraining/providers/Microsoft.Community/communityTrainings/ctApplication"),
	// 		Name: to.Ptr("ctApplication"),
	// 		Type: to.Ptr("Microsoft.Community/communityTrainings"),
	// 		Properties: &armcommunitymanagement.CommunityTrainingProperties{
	// 			PortalName: to.Ptr("ctwebsite"),
	// 			PortalAdminEmailAddress: to.Ptr("ctadmin@ct.com"),
	// 			PortalOwnerOrganizationName: to.Ptr("CT Portal Owner Organization"),
	// 			PortalOwnerEmailAddress: to.Ptr("ctcontact@ct.com"),
	// 			IdentityConfiguration: &armcommunitymanagement.IdentityConfigurationProperties{
	// 				IdentityType: to.Ptr("ADB2C"),
	// 				TeamsEnabled: to.Ptr(false),
	// 				TenantID: to.Ptr("c1ffbb60-88cf-4b83-b54f-c47ae6220c19"),
	// 				DomainName: to.Ptr("cttenant"),
	// 				ClientID: to.Ptr("8c92390f-2f30-493d-bd13-d3c3eba3709d"),
	// 				B2CAuthenticationPolicy: to.Ptr("B2C_1_signup_signin"),
	// 				B2CPasswordResetPolicy: to.Ptr("B2C_1_pwd_reset"),
	// 				CustomLoginParameters: to.Ptr("custom_hint"),
	// 			},
	// 			ZoneRedundancyEnabled: to.Ptr(true),
	// 			DisasterRecoveryEnabled: to.Ptr(true),
	// 			ProvisioningState: to.Ptr(armcommunitymanagement.ProvisioningStateSucceeded),
	// 		},
	// 		SKU: &armcommunitymanagement.SKU{
	// 			Name: to.Ptr("Commercial"),
	// 			Tier: to.Ptr(armcommunitymanagement.SKUTierStandard),
	// 		},
	// 		Tags: map[string]*string{
	// 		},
	// 		Location: to.Ptr("southeastasia"),
	// 		SystemData: &armcommunitymanagement.SystemData{
	// 			CreatedBy: to.Ptr("ctowner@ct.com"),
	// 			CreatedByType: to.Ptr(armcommunitymanagement.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-01T11:42:44.657Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("ctowner@ct.com"),
	// 			LastModifiedByType: to.Ptr(armcommunitymanagement.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-01T11:42:44.657Z"); return t}()),
	// 		},
	// 	},
	// }
}

// Generated from example definition: 2023-11-01/CommunityTrainings_Delete.json
func ExampleCommunityTrainingsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcommunitymanagement.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewCommunityTrainingsClient().BeginDelete(ctx, "rgCommunityTraining", "ctApplication", nil)
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
	// res = armcommunitymanagement.CommunityTrainingsClientDeleteResponse{
	// }
}

// Generated from example definition: 2023-11-01/CommunityTrainings_Get.json
func ExampleCommunityTrainingsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcommunitymanagement.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewCommunityTrainingsClient().Get(ctx, "rgCommunityTraining", "ctApplication", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcommunitymanagement.CommunityTrainingsClientGetResponse{
	// 	CommunityTraining: &armcommunitymanagement.CommunityTraining{
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rgCommunityTraining/providers/Microsoft.Community/communityTrainings/ctApplication"),
	// 		Name: to.Ptr("ctApplication"),
	// 		Type: to.Ptr("Microsoft.Community/communityTrainings"),
	// 		Properties: &armcommunitymanagement.CommunityTrainingProperties{
	// 			PortalName: to.Ptr("ctwebsite"),
	// 			PortalAdminEmailAddress: to.Ptr("ctadmin@ct.com"),
	// 			PortalOwnerOrganizationName: to.Ptr("CT Portal Owner Organization"),
	// 			PortalOwnerEmailAddress: to.Ptr("ctcontact@ct.com"),
	// 			IdentityConfiguration: &armcommunitymanagement.IdentityConfigurationProperties{
	// 				IdentityType: to.Ptr("ADB2C"),
	// 				TeamsEnabled: to.Ptr(false),
	// 				TenantID: to.Ptr("c1ffbb60-88cf-4b83-b54f-c47ae6220c19"),
	// 				DomainName: to.Ptr("cttenant"),
	// 				ClientID: to.Ptr("8c92390f-2f30-493d-bd13-d3c3eba3709d"),
	// 				B2CAuthenticationPolicy: to.Ptr("B2C_1_signup_signin"),
	// 				B2CPasswordResetPolicy: to.Ptr("B2C_1_pwd_reset"),
	// 				CustomLoginParameters: to.Ptr("custom_hint"),
	// 			},
	// 			ZoneRedundancyEnabled: to.Ptr(true),
	// 			DisasterRecoveryEnabled: to.Ptr(true),
	// 			ProvisioningState: to.Ptr(armcommunitymanagement.ProvisioningStateSucceeded),
	// 		},
	// 		SKU: &armcommunitymanagement.SKU{
	// 			Name: to.Ptr("Commercial"),
	// 			Tier: to.Ptr(armcommunitymanagement.SKUTierStandard),
	// 		},
	// 		Tags: map[string]*string{
	// 		},
	// 		Location: to.Ptr("southeastasia"),
	// 		SystemData: &armcommunitymanagement.SystemData{
	// 			CreatedBy: to.Ptr("ctowner@ct.com"),
	// 			CreatedByType: to.Ptr(armcommunitymanagement.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-01T11:42:44.657Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("ctowner@ct.com"),
	// 			LastModifiedByType: to.Ptr(armcommunitymanagement.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-01T11:42:44.657Z"); return t}()),
	// 		},
	// 	},
	// }
}

// Generated from example definition: 2023-11-01/CommunityTrainings_ListByResourceGroup.json
func ExampleCommunityTrainingsClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcommunitymanagement.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewCommunityTrainingsClient().NewListByResourceGroupPager("rgCommunityTraining", nil)
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
		// page = armcommunitymanagement.CommunityTrainingsClientListByResourceGroupResponse{
		// 	CommunityTrainingListResult: armcommunitymanagement.CommunityTrainingListResult{
		// 		Value: []*armcommunitymanagement.CommunityTraining{
		// 			{
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rgCommunityTraining/providers/Microsoft.Community/communityTrainings/ctApplication"),
		// 				Name: to.Ptr("ctApplication"),
		// 				Type: to.Ptr("Microsoft.Community/communityTrainings"),
		// 				Properties: &armcommunitymanagement.CommunityTrainingProperties{
		// 					PortalName: to.Ptr("ctwebsite"),
		// 					PortalAdminEmailAddress: to.Ptr("ctadmin@ct.com"),
		// 					PortalOwnerOrganizationName: to.Ptr("CT Portal Owner Organization"),
		// 					PortalOwnerEmailAddress: to.Ptr("ctcontact@ct.com"),
		// 					IdentityConfiguration: &armcommunitymanagement.IdentityConfigurationProperties{
		// 						IdentityType: to.Ptr("ADB2C"),
		// 						TeamsEnabled: to.Ptr(false),
		// 						TenantID: to.Ptr("c1ffbb60-88cf-4b83-b54f-c47ae6220c19"),
		// 						DomainName: to.Ptr("cttenant"),
		// 						ClientID: to.Ptr("8c92390f-2f30-493d-bd13-d3c3eba3709d"),
		// 						B2CAuthenticationPolicy: to.Ptr("B2C_1_signup_signin"),
		// 						B2CPasswordResetPolicy: to.Ptr("B2C_1_pwd_reset"),
		// 						CustomLoginParameters: to.Ptr("custom_hint"),
		// 					},
		// 					ZoneRedundancyEnabled: to.Ptr(true),
		// 					DisasterRecoveryEnabled: to.Ptr(true),
		// 					ProvisioningState: to.Ptr(armcommunitymanagement.ProvisioningStateSucceeded),
		// 				},
		// 				SKU: &armcommunitymanagement.SKU{
		// 					Name: to.Ptr("Commercial"),
		// 					Tier: to.Ptr(armcommunitymanagement.SKUTierStandard),
		// 				},
		// 				Tags: map[string]*string{
		// 				},
		// 				Location: to.Ptr("southeastasia"),
		// 				SystemData: &armcommunitymanagement.SystemData{
		// 					CreatedBy: to.Ptr("ctowner@ct.com"),
		// 					CreatedByType: to.Ptr(armcommunitymanagement.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-01T11:42:44.657Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("ctowner@ct.com"),
		// 					LastModifiedByType: to.Ptr(armcommunitymanagement.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-01T11:42:44.657Z"); return t}()),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}

// Generated from example definition: 2023-11-01/CommunityTrainings_ListBySubscription.json
func ExampleCommunityTrainingsClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcommunitymanagement.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewCommunityTrainingsClient().NewListBySubscriptionPager(nil)
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
		// page = armcommunitymanagement.CommunityTrainingsClientListBySubscriptionResponse{
		// 	CommunityTrainingListResult: armcommunitymanagement.CommunityTrainingListResult{
		// 		Value: []*armcommunitymanagement.CommunityTraining{
		// 			{
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rgCommunityTraining/providers/Microsoft.Community/communityTrainings/ctApplication"),
		// 				Name: to.Ptr("ctApplication"),
		// 				Type: to.Ptr("Microsoft.Community/communityTrainings"),
		// 				Properties: &armcommunitymanagement.CommunityTrainingProperties{
		// 					PortalName: to.Ptr("ctwebsite"),
		// 					PortalAdminEmailAddress: to.Ptr("ctadmin@ct.com"),
		// 					PortalOwnerOrganizationName: to.Ptr("CT Portal Owner Organization"),
		// 					PortalOwnerEmailAddress: to.Ptr("ctcontact@ct.com"),
		// 					IdentityConfiguration: &armcommunitymanagement.IdentityConfigurationProperties{
		// 						IdentityType: to.Ptr("ADB2C"),
		// 						TeamsEnabled: to.Ptr(false),
		// 						TenantID: to.Ptr("c1ffbb60-88cf-4b83-b54f-c47ae6220c19"),
		// 						DomainName: to.Ptr("cttenant"),
		// 						ClientID: to.Ptr("8c92390f-2f30-493d-bd13-d3c3eba3709d"),
		// 						B2CAuthenticationPolicy: to.Ptr("B2C_1_signup_signin"),
		// 						B2CPasswordResetPolicy: to.Ptr("B2C_1_pwd_reset"),
		// 						CustomLoginParameters: to.Ptr("custom_hint"),
		// 					},
		// 					ZoneRedundancyEnabled: to.Ptr(true),
		// 					DisasterRecoveryEnabled: to.Ptr(true),
		// 					ProvisioningState: to.Ptr(armcommunitymanagement.ProvisioningStateSucceeded),
		// 				},
		// 				SKU: &armcommunitymanagement.SKU{
		// 					Name: to.Ptr("Commercial"),
		// 					Tier: to.Ptr(armcommunitymanagement.SKUTierStandard),
		// 				},
		// 				Tags: map[string]*string{
		// 				},
		// 				Location: to.Ptr("southeastasia"),
		// 				SystemData: &armcommunitymanagement.SystemData{
		// 					CreatedBy: to.Ptr("ctowner@ct.com"),
		// 					CreatedByType: to.Ptr(armcommunitymanagement.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-01T11:42:44.657Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("ctowner@ct.com"),
		// 					LastModifiedByType: to.Ptr(armcommunitymanagement.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-01T11:42:44.657Z"); return t}()),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}

// Generated from example definition: 2023-11-01/CommunityTrainings_Update.json
func ExampleCommunityTrainingsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcommunitymanagement.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewCommunityTrainingsClient().BeginUpdate(ctx, "rgCommunityTraining", "ctApplication", armcommunitymanagement.CommunityTrainingUpdate{
		SKU: &armcommunitymanagement.SKU{
			Name: to.Ptr("Commercial"),
			Tier: to.Ptr(armcommunitymanagement.SKUTierStandard),
		},
		Tags: map[string]*string{},
		Properties: &armcommunitymanagement.CommunityTrainingUpdateProperties{
			IdentityConfiguration: &armcommunitymanagement.IdentityConfigurationProperties{
				IdentityType:            to.Ptr("ADB2C"),
				TeamsEnabled:            to.Ptr(false),
				TenantID:                to.Ptr("c1ffbb60-88cf-4b83-b54f-c47ae6220c19"),
				DomainName:              to.Ptr("cttenant"),
				ClientID:                to.Ptr("8c92390f-2f30-493d-bd13-d3c3eba3709d"),
				ClientSecret:            to.Ptr("idenityConfigurationClientSecret"),
				B2CAuthenticationPolicy: to.Ptr("B2C_1_signup_signin"),
				B2CPasswordResetPolicy:  to.Ptr("B2C_1_pwd_reset"),
				CustomLoginParameters:   to.Ptr("custom_hint"),
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
	// res = armcommunitymanagement.CommunityTrainingsClientUpdateResponse{
	// 	CommunityTraining: &armcommunitymanagement.CommunityTraining{
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rgCommunityTraining/providers/Microsoft.Community/communityTrainings/ctApplication"),
	// 		Name: to.Ptr("ctApplication"),
	// 		Type: to.Ptr("Microsoft.Community/communityTrainings"),
	// 		Properties: &armcommunitymanagement.CommunityTrainingProperties{
	// 			PortalName: to.Ptr("ctwebsite"),
	// 			PortalAdminEmailAddress: to.Ptr("ctadmin@ct.com"),
	// 			PortalOwnerOrganizationName: to.Ptr("CT Portal Owner Organization"),
	// 			PortalOwnerEmailAddress: to.Ptr("ctcontact@ct.com"),
	// 			IdentityConfiguration: &armcommunitymanagement.IdentityConfigurationProperties{
	// 				IdentityType: to.Ptr("ADB2C"),
	// 				TeamsEnabled: to.Ptr(false),
	// 				TenantID: to.Ptr("c1ffbb60-88cf-4b83-b54f-c47ae6220c19"),
	// 				DomainName: to.Ptr("cttenant"),
	// 				ClientID: to.Ptr("8c92390f-2f30-493d-bd13-d3c3eba3709d"),
	// 				B2CAuthenticationPolicy: to.Ptr("B2C_1_signup_signin"),
	// 				B2CPasswordResetPolicy: to.Ptr("B2C_1_pwd_reset"),
	// 				CustomLoginParameters: to.Ptr("custom_hint"),
	// 			},
	// 			ZoneRedundancyEnabled: to.Ptr(true),
	// 			DisasterRecoveryEnabled: to.Ptr(true),
	// 			ProvisioningState: to.Ptr(armcommunitymanagement.ProvisioningStateSucceeded),
	// 		},
	// 		SKU: &armcommunitymanagement.SKU{
	// 			Name: to.Ptr("Commercial"),
	// 			Tier: to.Ptr(armcommunitymanagement.SKUTierStandard),
	// 		},
	// 		Tags: map[string]*string{
	// 		},
	// 		Location: to.Ptr("southeastasia"),
	// 		SystemData: &armcommunitymanagement.SystemData{
	// 			CreatedBy: to.Ptr("ctowner@ct.com"),
	// 			CreatedByType: to.Ptr(armcommunitymanagement.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-01T11:42:44.657Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("ctowner@ct.com"),
	// 			LastModifiedByType: to.Ptr(armcommunitymanagement.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-01T11:42:44.657Z"); return t}()),
	// 		},
	// 	},
	// }
}
