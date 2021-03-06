// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package amplifystub

import (
	"github.com/aws/aws-sdk-go/service/amplify"
	"go.temporal.io/sdk/workflow"

	"go.temporal.io/aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	CreateApp(ctx workflow.Context, input *amplify.CreateAppInput) (*amplify.CreateAppOutput, error)
	CreateAppAsync(ctx workflow.Context, input *amplify.CreateAppInput) *AmplifyCreateAppFuture

	CreateBackendEnvironment(ctx workflow.Context, input *amplify.CreateBackendEnvironmentInput) (*amplify.CreateBackendEnvironmentOutput, error)
	CreateBackendEnvironmentAsync(ctx workflow.Context, input *amplify.CreateBackendEnvironmentInput) *AmplifyCreateBackendEnvironmentFuture

	CreateBranch(ctx workflow.Context, input *amplify.CreateBranchInput) (*amplify.CreateBranchOutput, error)
	CreateBranchAsync(ctx workflow.Context, input *amplify.CreateBranchInput) *AmplifyCreateBranchFuture

	CreateDeployment(ctx workflow.Context, input *amplify.CreateDeploymentInput) (*amplify.CreateDeploymentOutput, error)
	CreateDeploymentAsync(ctx workflow.Context, input *amplify.CreateDeploymentInput) *AmplifyCreateDeploymentFuture

	CreateDomainAssociation(ctx workflow.Context, input *amplify.CreateDomainAssociationInput) (*amplify.CreateDomainAssociationOutput, error)
	CreateDomainAssociationAsync(ctx workflow.Context, input *amplify.CreateDomainAssociationInput) *AmplifyCreateDomainAssociationFuture

	CreateWebhook(ctx workflow.Context, input *amplify.CreateWebhookInput) (*amplify.CreateWebhookOutput, error)
	CreateWebhookAsync(ctx workflow.Context, input *amplify.CreateWebhookInput) *AmplifyCreateWebhookFuture

	DeleteApp(ctx workflow.Context, input *amplify.DeleteAppInput) (*amplify.DeleteAppOutput, error)
	DeleteAppAsync(ctx workflow.Context, input *amplify.DeleteAppInput) *AmplifyDeleteAppFuture

	DeleteBackendEnvironment(ctx workflow.Context, input *amplify.DeleteBackendEnvironmentInput) (*amplify.DeleteBackendEnvironmentOutput, error)
	DeleteBackendEnvironmentAsync(ctx workflow.Context, input *amplify.DeleteBackendEnvironmentInput) *AmplifyDeleteBackendEnvironmentFuture

	DeleteBranch(ctx workflow.Context, input *amplify.DeleteBranchInput) (*amplify.DeleteBranchOutput, error)
	DeleteBranchAsync(ctx workflow.Context, input *amplify.DeleteBranchInput) *AmplifyDeleteBranchFuture

	DeleteDomainAssociation(ctx workflow.Context, input *amplify.DeleteDomainAssociationInput) (*amplify.DeleteDomainAssociationOutput, error)
	DeleteDomainAssociationAsync(ctx workflow.Context, input *amplify.DeleteDomainAssociationInput) *AmplifyDeleteDomainAssociationFuture

	DeleteJob(ctx workflow.Context, input *amplify.DeleteJobInput) (*amplify.DeleteJobOutput, error)
	DeleteJobAsync(ctx workflow.Context, input *amplify.DeleteJobInput) *AmplifyDeleteJobFuture

	DeleteWebhook(ctx workflow.Context, input *amplify.DeleteWebhookInput) (*amplify.DeleteWebhookOutput, error)
	DeleteWebhookAsync(ctx workflow.Context, input *amplify.DeleteWebhookInput) *AmplifyDeleteWebhookFuture

	GenerateAccessLogs(ctx workflow.Context, input *amplify.GenerateAccessLogsInput) (*amplify.GenerateAccessLogsOutput, error)
	GenerateAccessLogsAsync(ctx workflow.Context, input *amplify.GenerateAccessLogsInput) *AmplifyGenerateAccessLogsFuture

	GetApp(ctx workflow.Context, input *amplify.GetAppInput) (*amplify.GetAppOutput, error)
	GetAppAsync(ctx workflow.Context, input *amplify.GetAppInput) *AmplifyGetAppFuture

	GetArtifactUrl(ctx workflow.Context, input *amplify.GetArtifactUrlInput) (*amplify.GetArtifactUrlOutput, error)
	GetArtifactUrlAsync(ctx workflow.Context, input *amplify.GetArtifactUrlInput) *AmplifyGetArtifactUrlFuture

	GetBackendEnvironment(ctx workflow.Context, input *amplify.GetBackendEnvironmentInput) (*amplify.GetBackendEnvironmentOutput, error)
	GetBackendEnvironmentAsync(ctx workflow.Context, input *amplify.GetBackendEnvironmentInput) *AmplifyGetBackendEnvironmentFuture

	GetBranch(ctx workflow.Context, input *amplify.GetBranchInput) (*amplify.GetBranchOutput, error)
	GetBranchAsync(ctx workflow.Context, input *amplify.GetBranchInput) *AmplifyGetBranchFuture

	GetDomainAssociation(ctx workflow.Context, input *amplify.GetDomainAssociationInput) (*amplify.GetDomainAssociationOutput, error)
	GetDomainAssociationAsync(ctx workflow.Context, input *amplify.GetDomainAssociationInput) *AmplifyGetDomainAssociationFuture

	GetJob(ctx workflow.Context, input *amplify.GetJobInput) (*amplify.GetJobOutput, error)
	GetJobAsync(ctx workflow.Context, input *amplify.GetJobInput) *AmplifyGetJobFuture

	GetWebhook(ctx workflow.Context, input *amplify.GetWebhookInput) (*amplify.GetWebhookOutput, error)
	GetWebhookAsync(ctx workflow.Context, input *amplify.GetWebhookInput) *AmplifyGetWebhookFuture

	ListApps(ctx workflow.Context, input *amplify.ListAppsInput) (*amplify.ListAppsOutput, error)
	ListAppsAsync(ctx workflow.Context, input *amplify.ListAppsInput) *AmplifyListAppsFuture

	ListArtifacts(ctx workflow.Context, input *amplify.ListArtifactsInput) (*amplify.ListArtifactsOutput, error)
	ListArtifactsAsync(ctx workflow.Context, input *amplify.ListArtifactsInput) *AmplifyListArtifactsFuture

	ListBackendEnvironments(ctx workflow.Context, input *amplify.ListBackendEnvironmentsInput) (*amplify.ListBackendEnvironmentsOutput, error)
	ListBackendEnvironmentsAsync(ctx workflow.Context, input *amplify.ListBackendEnvironmentsInput) *AmplifyListBackendEnvironmentsFuture

	ListBranches(ctx workflow.Context, input *amplify.ListBranchesInput) (*amplify.ListBranchesOutput, error)
	ListBranchesAsync(ctx workflow.Context, input *amplify.ListBranchesInput) *AmplifyListBranchesFuture

	ListDomainAssociations(ctx workflow.Context, input *amplify.ListDomainAssociationsInput) (*amplify.ListDomainAssociationsOutput, error)
	ListDomainAssociationsAsync(ctx workflow.Context, input *amplify.ListDomainAssociationsInput) *AmplifyListDomainAssociationsFuture

	ListJobs(ctx workflow.Context, input *amplify.ListJobsInput) (*amplify.ListJobsOutput, error)
	ListJobsAsync(ctx workflow.Context, input *amplify.ListJobsInput) *AmplifyListJobsFuture

	ListTagsForResource(ctx workflow.Context, input *amplify.ListTagsForResourceInput) (*amplify.ListTagsForResourceOutput, error)
	ListTagsForResourceAsync(ctx workflow.Context, input *amplify.ListTagsForResourceInput) *AmplifyListTagsForResourceFuture

	ListWebhooks(ctx workflow.Context, input *amplify.ListWebhooksInput) (*amplify.ListWebhooksOutput, error)
	ListWebhooksAsync(ctx workflow.Context, input *amplify.ListWebhooksInput) *AmplifyListWebhooksFuture

	StartDeployment(ctx workflow.Context, input *amplify.StartDeploymentInput) (*amplify.StartDeploymentOutput, error)
	StartDeploymentAsync(ctx workflow.Context, input *amplify.StartDeploymentInput) *AmplifyStartDeploymentFuture

	StartJob(ctx workflow.Context, input *amplify.StartJobInput) (*amplify.StartJobOutput, error)
	StartJobAsync(ctx workflow.Context, input *amplify.StartJobInput) *AmplifyStartJobFuture

	StopJob(ctx workflow.Context, input *amplify.StopJobInput) (*amplify.StopJobOutput, error)
	StopJobAsync(ctx workflow.Context, input *amplify.StopJobInput) *AmplifyStopJobFuture

	TagResource(ctx workflow.Context, input *amplify.TagResourceInput) (*amplify.TagResourceOutput, error)
	TagResourceAsync(ctx workflow.Context, input *amplify.TagResourceInput) *AmplifyTagResourceFuture

	UntagResource(ctx workflow.Context, input *amplify.UntagResourceInput) (*amplify.UntagResourceOutput, error)
	UntagResourceAsync(ctx workflow.Context, input *amplify.UntagResourceInput) *AmplifyUntagResourceFuture

	UpdateApp(ctx workflow.Context, input *amplify.UpdateAppInput) (*amplify.UpdateAppOutput, error)
	UpdateAppAsync(ctx workflow.Context, input *amplify.UpdateAppInput) *AmplifyUpdateAppFuture

	UpdateBranch(ctx workflow.Context, input *amplify.UpdateBranchInput) (*amplify.UpdateBranchOutput, error)
	UpdateBranchAsync(ctx workflow.Context, input *amplify.UpdateBranchInput) *AmplifyUpdateBranchFuture

	UpdateDomainAssociation(ctx workflow.Context, input *amplify.UpdateDomainAssociationInput) (*amplify.UpdateDomainAssociationOutput, error)
	UpdateDomainAssociationAsync(ctx workflow.Context, input *amplify.UpdateDomainAssociationInput) *AmplifyUpdateDomainAssociationFuture

	UpdateWebhook(ctx workflow.Context, input *amplify.UpdateWebhookInput) (*amplify.UpdateWebhookOutput, error)
	UpdateWebhookAsync(ctx workflow.Context, input *amplify.UpdateWebhookInput) *AmplifyUpdateWebhookFuture
}

func NewClient() Client {
	return &stub{}
}
