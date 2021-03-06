// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package greengrassstub

import (
	"github.com/aws/aws-sdk-go/service/greengrass"
	"go.temporal.io/sdk/workflow"

	"go.temporal.io/aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	AssociateRoleToGroup(ctx workflow.Context, input *greengrass.AssociateRoleToGroupInput) (*greengrass.AssociateRoleToGroupOutput, error)
	AssociateRoleToGroupAsync(ctx workflow.Context, input *greengrass.AssociateRoleToGroupInput) *GreengrassAssociateRoleToGroupFuture

	AssociateServiceRoleToAccount(ctx workflow.Context, input *greengrass.AssociateServiceRoleToAccountInput) (*greengrass.AssociateServiceRoleToAccountOutput, error)
	AssociateServiceRoleToAccountAsync(ctx workflow.Context, input *greengrass.AssociateServiceRoleToAccountInput) *GreengrassAssociateServiceRoleToAccountFuture

	CreateConnectorDefinition(ctx workflow.Context, input *greengrass.CreateConnectorDefinitionInput) (*greengrass.CreateConnectorDefinitionOutput, error)
	CreateConnectorDefinitionAsync(ctx workflow.Context, input *greengrass.CreateConnectorDefinitionInput) *GreengrassCreateConnectorDefinitionFuture

	CreateConnectorDefinitionVersion(ctx workflow.Context, input *greengrass.CreateConnectorDefinitionVersionInput) (*greengrass.CreateConnectorDefinitionVersionOutput, error)
	CreateConnectorDefinitionVersionAsync(ctx workflow.Context, input *greengrass.CreateConnectorDefinitionVersionInput) *GreengrassCreateConnectorDefinitionVersionFuture

	CreateCoreDefinition(ctx workflow.Context, input *greengrass.CreateCoreDefinitionInput) (*greengrass.CreateCoreDefinitionOutput, error)
	CreateCoreDefinitionAsync(ctx workflow.Context, input *greengrass.CreateCoreDefinitionInput) *GreengrassCreateCoreDefinitionFuture

	CreateCoreDefinitionVersion(ctx workflow.Context, input *greengrass.CreateCoreDefinitionVersionInput) (*greengrass.CreateCoreDefinitionVersionOutput, error)
	CreateCoreDefinitionVersionAsync(ctx workflow.Context, input *greengrass.CreateCoreDefinitionVersionInput) *GreengrassCreateCoreDefinitionVersionFuture

	CreateDeployment(ctx workflow.Context, input *greengrass.CreateDeploymentInput) (*greengrass.CreateDeploymentOutput, error)
	CreateDeploymentAsync(ctx workflow.Context, input *greengrass.CreateDeploymentInput) *GreengrassCreateDeploymentFuture

	CreateDeviceDefinition(ctx workflow.Context, input *greengrass.CreateDeviceDefinitionInput) (*greengrass.CreateDeviceDefinitionOutput, error)
	CreateDeviceDefinitionAsync(ctx workflow.Context, input *greengrass.CreateDeviceDefinitionInput) *GreengrassCreateDeviceDefinitionFuture

	CreateDeviceDefinitionVersion(ctx workflow.Context, input *greengrass.CreateDeviceDefinitionVersionInput) (*greengrass.CreateDeviceDefinitionVersionOutput, error)
	CreateDeviceDefinitionVersionAsync(ctx workflow.Context, input *greengrass.CreateDeviceDefinitionVersionInput) *GreengrassCreateDeviceDefinitionVersionFuture

	CreateFunctionDefinition(ctx workflow.Context, input *greengrass.CreateFunctionDefinitionInput) (*greengrass.CreateFunctionDefinitionOutput, error)
	CreateFunctionDefinitionAsync(ctx workflow.Context, input *greengrass.CreateFunctionDefinitionInput) *GreengrassCreateFunctionDefinitionFuture

	CreateFunctionDefinitionVersion(ctx workflow.Context, input *greengrass.CreateFunctionDefinitionVersionInput) (*greengrass.CreateFunctionDefinitionVersionOutput, error)
	CreateFunctionDefinitionVersionAsync(ctx workflow.Context, input *greengrass.CreateFunctionDefinitionVersionInput) *GreengrassCreateFunctionDefinitionVersionFuture

	CreateGroup(ctx workflow.Context, input *greengrass.CreateGroupInput) (*greengrass.CreateGroupOutput, error)
	CreateGroupAsync(ctx workflow.Context, input *greengrass.CreateGroupInput) *GreengrassCreateGroupFuture

	CreateGroupCertificateAuthority(ctx workflow.Context, input *greengrass.CreateGroupCertificateAuthorityInput) (*greengrass.CreateGroupCertificateAuthorityOutput, error)
	CreateGroupCertificateAuthorityAsync(ctx workflow.Context, input *greengrass.CreateGroupCertificateAuthorityInput) *GreengrassCreateGroupCertificateAuthorityFuture

	CreateGroupVersion(ctx workflow.Context, input *greengrass.CreateGroupVersionInput) (*greengrass.CreateGroupVersionOutput, error)
	CreateGroupVersionAsync(ctx workflow.Context, input *greengrass.CreateGroupVersionInput) *GreengrassCreateGroupVersionFuture

	CreateLoggerDefinition(ctx workflow.Context, input *greengrass.CreateLoggerDefinitionInput) (*greengrass.CreateLoggerDefinitionOutput, error)
	CreateLoggerDefinitionAsync(ctx workflow.Context, input *greengrass.CreateLoggerDefinitionInput) *GreengrassCreateLoggerDefinitionFuture

	CreateLoggerDefinitionVersion(ctx workflow.Context, input *greengrass.CreateLoggerDefinitionVersionInput) (*greengrass.CreateLoggerDefinitionVersionOutput, error)
	CreateLoggerDefinitionVersionAsync(ctx workflow.Context, input *greengrass.CreateLoggerDefinitionVersionInput) *GreengrassCreateLoggerDefinitionVersionFuture

	CreateResourceDefinition(ctx workflow.Context, input *greengrass.CreateResourceDefinitionInput) (*greengrass.CreateResourceDefinitionOutput, error)
	CreateResourceDefinitionAsync(ctx workflow.Context, input *greengrass.CreateResourceDefinitionInput) *GreengrassCreateResourceDefinitionFuture

	CreateResourceDefinitionVersion(ctx workflow.Context, input *greengrass.CreateResourceDefinitionVersionInput) (*greengrass.CreateResourceDefinitionVersionOutput, error)
	CreateResourceDefinitionVersionAsync(ctx workflow.Context, input *greengrass.CreateResourceDefinitionVersionInput) *GreengrassCreateResourceDefinitionVersionFuture

	CreateSoftwareUpdateJob(ctx workflow.Context, input *greengrass.CreateSoftwareUpdateJobInput) (*greengrass.CreateSoftwareUpdateJobOutput, error)
	CreateSoftwareUpdateJobAsync(ctx workflow.Context, input *greengrass.CreateSoftwareUpdateJobInput) *GreengrassCreateSoftwareUpdateJobFuture

	CreateSubscriptionDefinition(ctx workflow.Context, input *greengrass.CreateSubscriptionDefinitionInput) (*greengrass.CreateSubscriptionDefinitionOutput, error)
	CreateSubscriptionDefinitionAsync(ctx workflow.Context, input *greengrass.CreateSubscriptionDefinitionInput) *GreengrassCreateSubscriptionDefinitionFuture

	CreateSubscriptionDefinitionVersion(ctx workflow.Context, input *greengrass.CreateSubscriptionDefinitionVersionInput) (*greengrass.CreateSubscriptionDefinitionVersionOutput, error)
	CreateSubscriptionDefinitionVersionAsync(ctx workflow.Context, input *greengrass.CreateSubscriptionDefinitionVersionInput) *GreengrassCreateSubscriptionDefinitionVersionFuture

	DeleteConnectorDefinition(ctx workflow.Context, input *greengrass.DeleteConnectorDefinitionInput) (*greengrass.DeleteConnectorDefinitionOutput, error)
	DeleteConnectorDefinitionAsync(ctx workflow.Context, input *greengrass.DeleteConnectorDefinitionInput) *GreengrassDeleteConnectorDefinitionFuture

	DeleteCoreDefinition(ctx workflow.Context, input *greengrass.DeleteCoreDefinitionInput) (*greengrass.DeleteCoreDefinitionOutput, error)
	DeleteCoreDefinitionAsync(ctx workflow.Context, input *greengrass.DeleteCoreDefinitionInput) *GreengrassDeleteCoreDefinitionFuture

	DeleteDeviceDefinition(ctx workflow.Context, input *greengrass.DeleteDeviceDefinitionInput) (*greengrass.DeleteDeviceDefinitionOutput, error)
	DeleteDeviceDefinitionAsync(ctx workflow.Context, input *greengrass.DeleteDeviceDefinitionInput) *GreengrassDeleteDeviceDefinitionFuture

	DeleteFunctionDefinition(ctx workflow.Context, input *greengrass.DeleteFunctionDefinitionInput) (*greengrass.DeleteFunctionDefinitionOutput, error)
	DeleteFunctionDefinitionAsync(ctx workflow.Context, input *greengrass.DeleteFunctionDefinitionInput) *GreengrassDeleteFunctionDefinitionFuture

	DeleteGroup(ctx workflow.Context, input *greengrass.DeleteGroupInput) (*greengrass.DeleteGroupOutput, error)
	DeleteGroupAsync(ctx workflow.Context, input *greengrass.DeleteGroupInput) *GreengrassDeleteGroupFuture

	DeleteLoggerDefinition(ctx workflow.Context, input *greengrass.DeleteLoggerDefinitionInput) (*greengrass.DeleteLoggerDefinitionOutput, error)
	DeleteLoggerDefinitionAsync(ctx workflow.Context, input *greengrass.DeleteLoggerDefinitionInput) *GreengrassDeleteLoggerDefinitionFuture

	DeleteResourceDefinition(ctx workflow.Context, input *greengrass.DeleteResourceDefinitionInput) (*greengrass.DeleteResourceDefinitionOutput, error)
	DeleteResourceDefinitionAsync(ctx workflow.Context, input *greengrass.DeleteResourceDefinitionInput) *GreengrassDeleteResourceDefinitionFuture

	DeleteSubscriptionDefinition(ctx workflow.Context, input *greengrass.DeleteSubscriptionDefinitionInput) (*greengrass.DeleteSubscriptionDefinitionOutput, error)
	DeleteSubscriptionDefinitionAsync(ctx workflow.Context, input *greengrass.DeleteSubscriptionDefinitionInput) *GreengrassDeleteSubscriptionDefinitionFuture

	DisassociateRoleFromGroup(ctx workflow.Context, input *greengrass.DisassociateRoleFromGroupInput) (*greengrass.DisassociateRoleFromGroupOutput, error)
	DisassociateRoleFromGroupAsync(ctx workflow.Context, input *greengrass.DisassociateRoleFromGroupInput) *GreengrassDisassociateRoleFromGroupFuture

	DisassociateServiceRoleFromAccount(ctx workflow.Context, input *greengrass.DisassociateServiceRoleFromAccountInput) (*greengrass.DisassociateServiceRoleFromAccountOutput, error)
	DisassociateServiceRoleFromAccountAsync(ctx workflow.Context, input *greengrass.DisassociateServiceRoleFromAccountInput) *GreengrassDisassociateServiceRoleFromAccountFuture

	GetAssociatedRole(ctx workflow.Context, input *greengrass.GetAssociatedRoleInput) (*greengrass.GetAssociatedRoleOutput, error)
	GetAssociatedRoleAsync(ctx workflow.Context, input *greengrass.GetAssociatedRoleInput) *GreengrassGetAssociatedRoleFuture

	GetBulkDeploymentStatus(ctx workflow.Context, input *greengrass.GetBulkDeploymentStatusInput) (*greengrass.GetBulkDeploymentStatusOutput, error)
	GetBulkDeploymentStatusAsync(ctx workflow.Context, input *greengrass.GetBulkDeploymentStatusInput) *GreengrassGetBulkDeploymentStatusFuture

	GetConnectivityInfo(ctx workflow.Context, input *greengrass.GetConnectivityInfoInput) (*greengrass.GetConnectivityInfoOutput, error)
	GetConnectivityInfoAsync(ctx workflow.Context, input *greengrass.GetConnectivityInfoInput) *GreengrassGetConnectivityInfoFuture

	GetConnectorDefinition(ctx workflow.Context, input *greengrass.GetConnectorDefinitionInput) (*greengrass.GetConnectorDefinitionOutput, error)
	GetConnectorDefinitionAsync(ctx workflow.Context, input *greengrass.GetConnectorDefinitionInput) *GreengrassGetConnectorDefinitionFuture

	GetConnectorDefinitionVersion(ctx workflow.Context, input *greengrass.GetConnectorDefinitionVersionInput) (*greengrass.GetConnectorDefinitionVersionOutput, error)
	GetConnectorDefinitionVersionAsync(ctx workflow.Context, input *greengrass.GetConnectorDefinitionVersionInput) *GreengrassGetConnectorDefinitionVersionFuture

	GetCoreDefinition(ctx workflow.Context, input *greengrass.GetCoreDefinitionInput) (*greengrass.GetCoreDefinitionOutput, error)
	GetCoreDefinitionAsync(ctx workflow.Context, input *greengrass.GetCoreDefinitionInput) *GreengrassGetCoreDefinitionFuture

	GetCoreDefinitionVersion(ctx workflow.Context, input *greengrass.GetCoreDefinitionVersionInput) (*greengrass.GetCoreDefinitionVersionOutput, error)
	GetCoreDefinitionVersionAsync(ctx workflow.Context, input *greengrass.GetCoreDefinitionVersionInput) *GreengrassGetCoreDefinitionVersionFuture

	GetDeploymentStatus(ctx workflow.Context, input *greengrass.GetDeploymentStatusInput) (*greengrass.GetDeploymentStatusOutput, error)
	GetDeploymentStatusAsync(ctx workflow.Context, input *greengrass.GetDeploymentStatusInput) *GreengrassGetDeploymentStatusFuture

	GetDeviceDefinition(ctx workflow.Context, input *greengrass.GetDeviceDefinitionInput) (*greengrass.GetDeviceDefinitionOutput, error)
	GetDeviceDefinitionAsync(ctx workflow.Context, input *greengrass.GetDeviceDefinitionInput) *GreengrassGetDeviceDefinitionFuture

	GetDeviceDefinitionVersion(ctx workflow.Context, input *greengrass.GetDeviceDefinitionVersionInput) (*greengrass.GetDeviceDefinitionVersionOutput, error)
	GetDeviceDefinitionVersionAsync(ctx workflow.Context, input *greengrass.GetDeviceDefinitionVersionInput) *GreengrassGetDeviceDefinitionVersionFuture

	GetFunctionDefinition(ctx workflow.Context, input *greengrass.GetFunctionDefinitionInput) (*greengrass.GetFunctionDefinitionOutput, error)
	GetFunctionDefinitionAsync(ctx workflow.Context, input *greengrass.GetFunctionDefinitionInput) *GreengrassGetFunctionDefinitionFuture

	GetFunctionDefinitionVersion(ctx workflow.Context, input *greengrass.GetFunctionDefinitionVersionInput) (*greengrass.GetFunctionDefinitionVersionOutput, error)
	GetFunctionDefinitionVersionAsync(ctx workflow.Context, input *greengrass.GetFunctionDefinitionVersionInput) *GreengrassGetFunctionDefinitionVersionFuture

	GetGroup(ctx workflow.Context, input *greengrass.GetGroupInput) (*greengrass.GetGroupOutput, error)
	GetGroupAsync(ctx workflow.Context, input *greengrass.GetGroupInput) *GreengrassGetGroupFuture

	GetGroupCertificateAuthority(ctx workflow.Context, input *greengrass.GetGroupCertificateAuthorityInput) (*greengrass.GetGroupCertificateAuthorityOutput, error)
	GetGroupCertificateAuthorityAsync(ctx workflow.Context, input *greengrass.GetGroupCertificateAuthorityInput) *GreengrassGetGroupCertificateAuthorityFuture

	GetGroupCertificateConfiguration(ctx workflow.Context, input *greengrass.GetGroupCertificateConfigurationInput) (*greengrass.GetGroupCertificateConfigurationOutput, error)
	GetGroupCertificateConfigurationAsync(ctx workflow.Context, input *greengrass.GetGroupCertificateConfigurationInput) *GreengrassGetGroupCertificateConfigurationFuture

	GetGroupVersion(ctx workflow.Context, input *greengrass.GetGroupVersionInput) (*greengrass.GetGroupVersionOutput, error)
	GetGroupVersionAsync(ctx workflow.Context, input *greengrass.GetGroupVersionInput) *GreengrassGetGroupVersionFuture

	GetLoggerDefinition(ctx workflow.Context, input *greengrass.GetLoggerDefinitionInput) (*greengrass.GetLoggerDefinitionOutput, error)
	GetLoggerDefinitionAsync(ctx workflow.Context, input *greengrass.GetLoggerDefinitionInput) *GreengrassGetLoggerDefinitionFuture

	GetLoggerDefinitionVersion(ctx workflow.Context, input *greengrass.GetLoggerDefinitionVersionInput) (*greengrass.GetLoggerDefinitionVersionOutput, error)
	GetLoggerDefinitionVersionAsync(ctx workflow.Context, input *greengrass.GetLoggerDefinitionVersionInput) *GreengrassGetLoggerDefinitionVersionFuture

	GetResourceDefinition(ctx workflow.Context, input *greengrass.GetResourceDefinitionInput) (*greengrass.GetResourceDefinitionOutput, error)
	GetResourceDefinitionAsync(ctx workflow.Context, input *greengrass.GetResourceDefinitionInput) *GreengrassGetResourceDefinitionFuture

	GetResourceDefinitionVersion(ctx workflow.Context, input *greengrass.GetResourceDefinitionVersionInput) (*greengrass.GetResourceDefinitionVersionOutput, error)
	GetResourceDefinitionVersionAsync(ctx workflow.Context, input *greengrass.GetResourceDefinitionVersionInput) *GreengrassGetResourceDefinitionVersionFuture

	GetServiceRoleForAccount(ctx workflow.Context, input *greengrass.GetServiceRoleForAccountInput) (*greengrass.GetServiceRoleForAccountOutput, error)
	GetServiceRoleForAccountAsync(ctx workflow.Context, input *greengrass.GetServiceRoleForAccountInput) *GreengrassGetServiceRoleForAccountFuture

	GetSubscriptionDefinition(ctx workflow.Context, input *greengrass.GetSubscriptionDefinitionInput) (*greengrass.GetSubscriptionDefinitionOutput, error)
	GetSubscriptionDefinitionAsync(ctx workflow.Context, input *greengrass.GetSubscriptionDefinitionInput) *GreengrassGetSubscriptionDefinitionFuture

	GetSubscriptionDefinitionVersion(ctx workflow.Context, input *greengrass.GetSubscriptionDefinitionVersionInput) (*greengrass.GetSubscriptionDefinitionVersionOutput, error)
	GetSubscriptionDefinitionVersionAsync(ctx workflow.Context, input *greengrass.GetSubscriptionDefinitionVersionInput) *GreengrassGetSubscriptionDefinitionVersionFuture

	GetThingRuntimeConfiguration(ctx workflow.Context, input *greengrass.GetThingRuntimeConfigurationInput) (*greengrass.GetThingRuntimeConfigurationOutput, error)
	GetThingRuntimeConfigurationAsync(ctx workflow.Context, input *greengrass.GetThingRuntimeConfigurationInput) *GreengrassGetThingRuntimeConfigurationFuture

	ListBulkDeploymentDetailedReports(ctx workflow.Context, input *greengrass.ListBulkDeploymentDetailedReportsInput) (*greengrass.ListBulkDeploymentDetailedReportsOutput, error)
	ListBulkDeploymentDetailedReportsAsync(ctx workflow.Context, input *greengrass.ListBulkDeploymentDetailedReportsInput) *GreengrassListBulkDeploymentDetailedReportsFuture

	ListBulkDeployments(ctx workflow.Context, input *greengrass.ListBulkDeploymentsInput) (*greengrass.ListBulkDeploymentsOutput, error)
	ListBulkDeploymentsAsync(ctx workflow.Context, input *greengrass.ListBulkDeploymentsInput) *GreengrassListBulkDeploymentsFuture

	ListConnectorDefinitionVersions(ctx workflow.Context, input *greengrass.ListConnectorDefinitionVersionsInput) (*greengrass.ListConnectorDefinitionVersionsOutput, error)
	ListConnectorDefinitionVersionsAsync(ctx workflow.Context, input *greengrass.ListConnectorDefinitionVersionsInput) *GreengrassListConnectorDefinitionVersionsFuture

	ListConnectorDefinitions(ctx workflow.Context, input *greengrass.ListConnectorDefinitionsInput) (*greengrass.ListConnectorDefinitionsOutput, error)
	ListConnectorDefinitionsAsync(ctx workflow.Context, input *greengrass.ListConnectorDefinitionsInput) *GreengrassListConnectorDefinitionsFuture

	ListCoreDefinitionVersions(ctx workflow.Context, input *greengrass.ListCoreDefinitionVersionsInput) (*greengrass.ListCoreDefinitionVersionsOutput, error)
	ListCoreDefinitionVersionsAsync(ctx workflow.Context, input *greengrass.ListCoreDefinitionVersionsInput) *GreengrassListCoreDefinitionVersionsFuture

	ListCoreDefinitions(ctx workflow.Context, input *greengrass.ListCoreDefinitionsInput) (*greengrass.ListCoreDefinitionsOutput, error)
	ListCoreDefinitionsAsync(ctx workflow.Context, input *greengrass.ListCoreDefinitionsInput) *GreengrassListCoreDefinitionsFuture

	ListDeployments(ctx workflow.Context, input *greengrass.ListDeploymentsInput) (*greengrass.ListDeploymentsOutput, error)
	ListDeploymentsAsync(ctx workflow.Context, input *greengrass.ListDeploymentsInput) *GreengrassListDeploymentsFuture

	ListDeviceDefinitionVersions(ctx workflow.Context, input *greengrass.ListDeviceDefinitionVersionsInput) (*greengrass.ListDeviceDefinitionVersionsOutput, error)
	ListDeviceDefinitionVersionsAsync(ctx workflow.Context, input *greengrass.ListDeviceDefinitionVersionsInput) *GreengrassListDeviceDefinitionVersionsFuture

	ListDeviceDefinitions(ctx workflow.Context, input *greengrass.ListDeviceDefinitionsInput) (*greengrass.ListDeviceDefinitionsOutput, error)
	ListDeviceDefinitionsAsync(ctx workflow.Context, input *greengrass.ListDeviceDefinitionsInput) *GreengrassListDeviceDefinitionsFuture

	ListFunctionDefinitionVersions(ctx workflow.Context, input *greengrass.ListFunctionDefinitionVersionsInput) (*greengrass.ListFunctionDefinitionVersionsOutput, error)
	ListFunctionDefinitionVersionsAsync(ctx workflow.Context, input *greengrass.ListFunctionDefinitionVersionsInput) *GreengrassListFunctionDefinitionVersionsFuture

	ListFunctionDefinitions(ctx workflow.Context, input *greengrass.ListFunctionDefinitionsInput) (*greengrass.ListFunctionDefinitionsOutput, error)
	ListFunctionDefinitionsAsync(ctx workflow.Context, input *greengrass.ListFunctionDefinitionsInput) *GreengrassListFunctionDefinitionsFuture

	ListGroupCertificateAuthorities(ctx workflow.Context, input *greengrass.ListGroupCertificateAuthoritiesInput) (*greengrass.ListGroupCertificateAuthoritiesOutput, error)
	ListGroupCertificateAuthoritiesAsync(ctx workflow.Context, input *greengrass.ListGroupCertificateAuthoritiesInput) *GreengrassListGroupCertificateAuthoritiesFuture

	ListGroupVersions(ctx workflow.Context, input *greengrass.ListGroupVersionsInput) (*greengrass.ListGroupVersionsOutput, error)
	ListGroupVersionsAsync(ctx workflow.Context, input *greengrass.ListGroupVersionsInput) *GreengrassListGroupVersionsFuture

	ListGroups(ctx workflow.Context, input *greengrass.ListGroupsInput) (*greengrass.ListGroupsOutput, error)
	ListGroupsAsync(ctx workflow.Context, input *greengrass.ListGroupsInput) *GreengrassListGroupsFuture

	ListLoggerDefinitionVersions(ctx workflow.Context, input *greengrass.ListLoggerDefinitionVersionsInput) (*greengrass.ListLoggerDefinitionVersionsOutput, error)
	ListLoggerDefinitionVersionsAsync(ctx workflow.Context, input *greengrass.ListLoggerDefinitionVersionsInput) *GreengrassListLoggerDefinitionVersionsFuture

	ListLoggerDefinitions(ctx workflow.Context, input *greengrass.ListLoggerDefinitionsInput) (*greengrass.ListLoggerDefinitionsOutput, error)
	ListLoggerDefinitionsAsync(ctx workflow.Context, input *greengrass.ListLoggerDefinitionsInput) *GreengrassListLoggerDefinitionsFuture

	ListResourceDefinitionVersions(ctx workflow.Context, input *greengrass.ListResourceDefinitionVersionsInput) (*greengrass.ListResourceDefinitionVersionsOutput, error)
	ListResourceDefinitionVersionsAsync(ctx workflow.Context, input *greengrass.ListResourceDefinitionVersionsInput) *GreengrassListResourceDefinitionVersionsFuture

	ListResourceDefinitions(ctx workflow.Context, input *greengrass.ListResourceDefinitionsInput) (*greengrass.ListResourceDefinitionsOutput, error)
	ListResourceDefinitionsAsync(ctx workflow.Context, input *greengrass.ListResourceDefinitionsInput) *GreengrassListResourceDefinitionsFuture

	ListSubscriptionDefinitionVersions(ctx workflow.Context, input *greengrass.ListSubscriptionDefinitionVersionsInput) (*greengrass.ListSubscriptionDefinitionVersionsOutput, error)
	ListSubscriptionDefinitionVersionsAsync(ctx workflow.Context, input *greengrass.ListSubscriptionDefinitionVersionsInput) *GreengrassListSubscriptionDefinitionVersionsFuture

	ListSubscriptionDefinitions(ctx workflow.Context, input *greengrass.ListSubscriptionDefinitionsInput) (*greengrass.ListSubscriptionDefinitionsOutput, error)
	ListSubscriptionDefinitionsAsync(ctx workflow.Context, input *greengrass.ListSubscriptionDefinitionsInput) *GreengrassListSubscriptionDefinitionsFuture

	ListTagsForResource(ctx workflow.Context, input *greengrass.ListTagsForResourceInput) (*greengrass.ListTagsForResourceOutput, error)
	ListTagsForResourceAsync(ctx workflow.Context, input *greengrass.ListTagsForResourceInput) *GreengrassListTagsForResourceFuture

	ResetDeployments(ctx workflow.Context, input *greengrass.ResetDeploymentsInput) (*greengrass.ResetDeploymentsOutput, error)
	ResetDeploymentsAsync(ctx workflow.Context, input *greengrass.ResetDeploymentsInput) *GreengrassResetDeploymentsFuture

	StartBulkDeployment(ctx workflow.Context, input *greengrass.StartBulkDeploymentInput) (*greengrass.StartBulkDeploymentOutput, error)
	StartBulkDeploymentAsync(ctx workflow.Context, input *greengrass.StartBulkDeploymentInput) *GreengrassStartBulkDeploymentFuture

	StopBulkDeployment(ctx workflow.Context, input *greengrass.StopBulkDeploymentInput) (*greengrass.StopBulkDeploymentOutput, error)
	StopBulkDeploymentAsync(ctx workflow.Context, input *greengrass.StopBulkDeploymentInput) *GreengrassStopBulkDeploymentFuture

	TagResource(ctx workflow.Context, input *greengrass.TagResourceInput) (*greengrass.TagResourceOutput, error)
	TagResourceAsync(ctx workflow.Context, input *greengrass.TagResourceInput) *GreengrassTagResourceFuture

	UntagResource(ctx workflow.Context, input *greengrass.UntagResourceInput) (*greengrass.UntagResourceOutput, error)
	UntagResourceAsync(ctx workflow.Context, input *greengrass.UntagResourceInput) *GreengrassUntagResourceFuture

	UpdateConnectivityInfo(ctx workflow.Context, input *greengrass.UpdateConnectivityInfoInput) (*greengrass.UpdateConnectivityInfoOutput, error)
	UpdateConnectivityInfoAsync(ctx workflow.Context, input *greengrass.UpdateConnectivityInfoInput) *GreengrassUpdateConnectivityInfoFuture

	UpdateConnectorDefinition(ctx workflow.Context, input *greengrass.UpdateConnectorDefinitionInput) (*greengrass.UpdateConnectorDefinitionOutput, error)
	UpdateConnectorDefinitionAsync(ctx workflow.Context, input *greengrass.UpdateConnectorDefinitionInput) *GreengrassUpdateConnectorDefinitionFuture

	UpdateCoreDefinition(ctx workflow.Context, input *greengrass.UpdateCoreDefinitionInput) (*greengrass.UpdateCoreDefinitionOutput, error)
	UpdateCoreDefinitionAsync(ctx workflow.Context, input *greengrass.UpdateCoreDefinitionInput) *GreengrassUpdateCoreDefinitionFuture

	UpdateDeviceDefinition(ctx workflow.Context, input *greengrass.UpdateDeviceDefinitionInput) (*greengrass.UpdateDeviceDefinitionOutput, error)
	UpdateDeviceDefinitionAsync(ctx workflow.Context, input *greengrass.UpdateDeviceDefinitionInput) *GreengrassUpdateDeviceDefinitionFuture

	UpdateFunctionDefinition(ctx workflow.Context, input *greengrass.UpdateFunctionDefinitionInput) (*greengrass.UpdateFunctionDefinitionOutput, error)
	UpdateFunctionDefinitionAsync(ctx workflow.Context, input *greengrass.UpdateFunctionDefinitionInput) *GreengrassUpdateFunctionDefinitionFuture

	UpdateGroup(ctx workflow.Context, input *greengrass.UpdateGroupInput) (*greengrass.UpdateGroupOutput, error)
	UpdateGroupAsync(ctx workflow.Context, input *greengrass.UpdateGroupInput) *GreengrassUpdateGroupFuture

	UpdateGroupCertificateConfiguration(ctx workflow.Context, input *greengrass.UpdateGroupCertificateConfigurationInput) (*greengrass.UpdateGroupCertificateConfigurationOutput, error)
	UpdateGroupCertificateConfigurationAsync(ctx workflow.Context, input *greengrass.UpdateGroupCertificateConfigurationInput) *GreengrassUpdateGroupCertificateConfigurationFuture

	UpdateLoggerDefinition(ctx workflow.Context, input *greengrass.UpdateLoggerDefinitionInput) (*greengrass.UpdateLoggerDefinitionOutput, error)
	UpdateLoggerDefinitionAsync(ctx workflow.Context, input *greengrass.UpdateLoggerDefinitionInput) *GreengrassUpdateLoggerDefinitionFuture

	UpdateResourceDefinition(ctx workflow.Context, input *greengrass.UpdateResourceDefinitionInput) (*greengrass.UpdateResourceDefinitionOutput, error)
	UpdateResourceDefinitionAsync(ctx workflow.Context, input *greengrass.UpdateResourceDefinitionInput) *GreengrassUpdateResourceDefinitionFuture

	UpdateSubscriptionDefinition(ctx workflow.Context, input *greengrass.UpdateSubscriptionDefinitionInput) (*greengrass.UpdateSubscriptionDefinitionOutput, error)
	UpdateSubscriptionDefinitionAsync(ctx workflow.Context, input *greengrass.UpdateSubscriptionDefinitionInput) *GreengrassUpdateSubscriptionDefinitionFuture

	UpdateThingRuntimeConfiguration(ctx workflow.Context, input *greengrass.UpdateThingRuntimeConfigurationInput) (*greengrass.UpdateThingRuntimeConfigurationOutput, error)
	UpdateThingRuntimeConfigurationAsync(ctx workflow.Context, input *greengrass.UpdateThingRuntimeConfigurationInput) *GreengrassUpdateThingRuntimeConfigurationFuture
}

func NewClient() Client {
	return &stub{}
}
