//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armlogz_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/logz/armlogz"
)

// x-ms-original-file: specification/logz/resource-manager/Microsoft.Logz/stable/2020-10-01/examples/SubAccount_ListByResourceGroup.json
func ExampleSubAccountClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armlogz.NewSubAccountClient("<subscription-id>", cred, nil)
	pager := client.List("<resource-group-name>",
		"<monitor-name>",
		nil)
	for pager.NextPage(ctx) {
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("LogzMonitorResource.ID: %s\n", *v.ID)
		}
	}
}

// x-ms-original-file: specification/logz/resource-manager/Microsoft.Logz/stable/2020-10-01/examples/SubAccount_Create.json
func ExampleSubAccountClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armlogz.NewSubAccountClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreate(ctx,
		"<resource-group-name>",
		"<monitor-name>",
		"<sub-account-name>",
		&armlogz.SubAccountBeginCreateOptions{Body: nil})
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("LogzMonitorResource.ID: %s\n", *res.ID)
}

// x-ms-original-file: specification/logz/resource-manager/Microsoft.Logz/stable/2020-10-01/examples/SubAccount_Get.json
func ExampleSubAccountClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armlogz.NewSubAccountClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<monitor-name>",
		"<sub-account-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("LogzMonitorResource.ID: %s\n", *res.ID)
}

// x-ms-original-file: specification/logz/resource-manager/Microsoft.Logz/stable/2020-10-01/examples/SubAccount_Delete.json
func ExampleSubAccountClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armlogz.NewSubAccountClient("<subscription-id>", cred, nil)
	poller, err := client.BeginDelete(ctx,
		"<resource-group-name>",
		"<monitor-name>",
		"<sub-account-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/logz/resource-manager/Microsoft.Logz/stable/2020-10-01/examples/SubAccount_Update.json
func ExampleSubAccountClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armlogz.NewSubAccountClient("<subscription-id>", cred, nil)
	res, err := client.Update(ctx,
		"<resource-group-name>",
		"<monitor-name>",
		"<sub-account-name>",
		&armlogz.SubAccountUpdateOptions{Body: nil})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("LogzMonitorResource.ID: %s\n", *res.ID)
}

// x-ms-original-file: specification/logz/resource-manager/Microsoft.Logz/stable/2020-10-01/examples/SubAccount_MonitoredResources_List.json
func ExampleSubAccountClient_ListMonitoredResources() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armlogz.NewSubAccountClient("<subscription-id>", cred, nil)
	pager := client.ListMonitoredResources("<resource-group-name>",
		"<monitor-name>",
		"<sub-account-name>",
		nil)
	for pager.NextPage(ctx) {
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("MonitoredResource.ID: %s\n", *v.ID)
		}
	}
}

// x-ms-original-file: specification/logz/resource-manager/Microsoft.Logz/stable/2020-10-01/examples/SubAccount_VMHosts_Payload.json
func ExampleSubAccountClient_VMHostPayload() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armlogz.NewSubAccountClient("<subscription-id>", cred, nil)
	_, err = client.VMHostPayload(ctx,
		"<resource-group-name>",
		"<monitor-name>",
		"<sub-account-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/logz/resource-manager/Microsoft.Logz/stable/2020-10-01/examples/SubAccount_VMHosts_Update.json
func ExampleSubAccountClient_ListVMHostUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armlogz.NewSubAccountClient("<subscription-id>", cred, nil)
	pager := client.ListVMHostUpdate("<resource-group-name>",
		"<monitor-name>",
		"<sub-account-name>",
		&armlogz.SubAccountListVMHostUpdateOptions{Body: nil})
	for pager.NextPage(ctx) {
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("VMResources.ID: %s\n", *v.ID)
		}
	}
}

// x-ms-original-file: specification/logz/resource-manager/Microsoft.Logz/stable/2020-10-01/examples/SubAccount_VMHosts_List.json
func ExampleSubAccountClient_ListVMHosts() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armlogz.NewSubAccountClient("<subscription-id>", cred, nil)
	pager := client.ListVMHosts("<resource-group-name>",
		"<monitor-name>",
		"<sub-account-name>",
		nil)
	for pager.NextPage(ctx) {
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("VMResources.ID: %s\n", *v.ID)
		}
	}
}