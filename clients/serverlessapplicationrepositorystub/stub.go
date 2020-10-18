// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package serverlessapplicationrepositorystub

import (
	"github.com/aws/aws-sdk-go/service/serverlessapplicationrepository"
	"go.temporal.io/sdk/workflow"

	"go.temporal.io/aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type stub struct{}

type ServerlessApplicationRepositoryCreateApplicationFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ServerlessApplicationRepositoryCreateApplicationFuture) Get(ctx workflow.Context) (*serverlessapplicationrepository.CreateApplicationOutput, error) {
	var output serverlessapplicationrepository.CreateApplicationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ServerlessApplicationRepositoryCreateApplicationVersionFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ServerlessApplicationRepositoryCreateApplicationVersionFuture) Get(ctx workflow.Context) (*serverlessapplicationrepository.CreateApplicationVersionOutput, error) {
	var output serverlessapplicationrepository.CreateApplicationVersionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ServerlessApplicationRepositoryCreateCloudFormationChangeSetFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ServerlessApplicationRepositoryCreateCloudFormationChangeSetFuture) Get(ctx workflow.Context) (*serverlessapplicationrepository.CreateCloudFormationChangeSetOutput, error) {
	var output serverlessapplicationrepository.CreateCloudFormationChangeSetOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ServerlessApplicationRepositoryCreateCloudFormationTemplateFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ServerlessApplicationRepositoryCreateCloudFormationTemplateFuture) Get(ctx workflow.Context) (*serverlessapplicationrepository.CreateCloudFormationTemplateOutput, error) {
	var output serverlessapplicationrepository.CreateCloudFormationTemplateOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ServerlessApplicationRepositoryDeleteApplicationFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ServerlessApplicationRepositoryDeleteApplicationFuture) Get(ctx workflow.Context) (*serverlessapplicationrepository.DeleteApplicationOutput, error) {
	var output serverlessapplicationrepository.DeleteApplicationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ServerlessApplicationRepositoryGetApplicationFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ServerlessApplicationRepositoryGetApplicationFuture) Get(ctx workflow.Context) (*serverlessapplicationrepository.GetApplicationOutput, error) {
	var output serverlessapplicationrepository.GetApplicationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ServerlessApplicationRepositoryGetApplicationPolicyFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ServerlessApplicationRepositoryGetApplicationPolicyFuture) Get(ctx workflow.Context) (*serverlessapplicationrepository.GetApplicationPolicyOutput, error) {
	var output serverlessapplicationrepository.GetApplicationPolicyOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ServerlessApplicationRepositoryGetCloudFormationTemplateFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ServerlessApplicationRepositoryGetCloudFormationTemplateFuture) Get(ctx workflow.Context) (*serverlessapplicationrepository.GetCloudFormationTemplateOutput, error) {
	var output serverlessapplicationrepository.GetCloudFormationTemplateOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ServerlessApplicationRepositoryListApplicationDependenciesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ServerlessApplicationRepositoryListApplicationDependenciesFuture) Get(ctx workflow.Context) (*serverlessapplicationrepository.ListApplicationDependenciesOutput, error) {
	var output serverlessapplicationrepository.ListApplicationDependenciesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ServerlessApplicationRepositoryListApplicationVersionsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ServerlessApplicationRepositoryListApplicationVersionsFuture) Get(ctx workflow.Context) (*serverlessapplicationrepository.ListApplicationVersionsOutput, error) {
	var output serverlessapplicationrepository.ListApplicationVersionsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ServerlessApplicationRepositoryListApplicationsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ServerlessApplicationRepositoryListApplicationsFuture) Get(ctx workflow.Context) (*serverlessapplicationrepository.ListApplicationsOutput, error) {
	var output serverlessapplicationrepository.ListApplicationsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ServerlessApplicationRepositoryPutApplicationPolicyFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ServerlessApplicationRepositoryPutApplicationPolicyFuture) Get(ctx workflow.Context) (*serverlessapplicationrepository.PutApplicationPolicyOutput, error) {
	var output serverlessapplicationrepository.PutApplicationPolicyOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ServerlessApplicationRepositoryUnshareApplicationFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ServerlessApplicationRepositoryUnshareApplicationFuture) Get(ctx workflow.Context) (*serverlessapplicationrepository.UnshareApplicationOutput, error) {
	var output serverlessapplicationrepository.UnshareApplicationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ServerlessApplicationRepositoryUpdateApplicationFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ServerlessApplicationRepositoryUpdateApplicationFuture) Get(ctx workflow.Context) (*serverlessapplicationrepository.UpdateApplicationOutput, error) {
	var output serverlessapplicationrepository.UpdateApplicationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateApplication(ctx workflow.Context, input *serverlessapplicationrepository.CreateApplicationRequest) (*serverlessapplicationrepository.CreateApplicationOutput, error) {
	var output serverlessapplicationrepository.CreateApplicationOutput
	err := workflow.ExecuteActivity(ctx, "aws.serverlessapplicationrepository.CreateApplication", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateApplicationAsync(ctx workflow.Context, input *serverlessapplicationrepository.CreateApplicationRequest) *ServerlessApplicationRepositoryCreateApplicationFuture {
	future := workflow.ExecuteActivity(ctx, "aws.serverlessapplicationrepository.CreateApplication", input)
	return &ServerlessApplicationRepositoryCreateApplicationFuture{Future: future}
}

func (a *stub) CreateApplicationVersion(ctx workflow.Context, input *serverlessapplicationrepository.CreateApplicationVersionRequest) (*serverlessapplicationrepository.CreateApplicationVersionOutput, error) {
	var output serverlessapplicationrepository.CreateApplicationVersionOutput
	err := workflow.ExecuteActivity(ctx, "aws.serverlessapplicationrepository.CreateApplicationVersion", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateApplicationVersionAsync(ctx workflow.Context, input *serverlessapplicationrepository.CreateApplicationVersionRequest) *ServerlessApplicationRepositoryCreateApplicationVersionFuture {
	future := workflow.ExecuteActivity(ctx, "aws.serverlessapplicationrepository.CreateApplicationVersion", input)
	return &ServerlessApplicationRepositoryCreateApplicationVersionFuture{Future: future}
}

func (a *stub) CreateCloudFormationChangeSet(ctx workflow.Context, input *serverlessapplicationrepository.CreateCloudFormationChangeSetRequest) (*serverlessapplicationrepository.CreateCloudFormationChangeSetOutput, error) {
	var output serverlessapplicationrepository.CreateCloudFormationChangeSetOutput
	err := workflow.ExecuteActivity(ctx, "aws.serverlessapplicationrepository.CreateCloudFormationChangeSet", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateCloudFormationChangeSetAsync(ctx workflow.Context, input *serverlessapplicationrepository.CreateCloudFormationChangeSetRequest) *ServerlessApplicationRepositoryCreateCloudFormationChangeSetFuture {
	future := workflow.ExecuteActivity(ctx, "aws.serverlessapplicationrepository.CreateCloudFormationChangeSet", input)
	return &ServerlessApplicationRepositoryCreateCloudFormationChangeSetFuture{Future: future}
}

func (a *stub) CreateCloudFormationTemplate(ctx workflow.Context, input *serverlessapplicationrepository.CreateCloudFormationTemplateInput) (*serverlessapplicationrepository.CreateCloudFormationTemplateOutput, error) {
	var output serverlessapplicationrepository.CreateCloudFormationTemplateOutput
	err := workflow.ExecuteActivity(ctx, "aws.serverlessapplicationrepository.CreateCloudFormationTemplate", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateCloudFormationTemplateAsync(ctx workflow.Context, input *serverlessapplicationrepository.CreateCloudFormationTemplateInput) *ServerlessApplicationRepositoryCreateCloudFormationTemplateFuture {
	future := workflow.ExecuteActivity(ctx, "aws.serverlessapplicationrepository.CreateCloudFormationTemplate", input)
	return &ServerlessApplicationRepositoryCreateCloudFormationTemplateFuture{Future: future}
}

func (a *stub) DeleteApplication(ctx workflow.Context, input *serverlessapplicationrepository.DeleteApplicationInput) (*serverlessapplicationrepository.DeleteApplicationOutput, error) {
	var output serverlessapplicationrepository.DeleteApplicationOutput
	err := workflow.ExecuteActivity(ctx, "aws.serverlessapplicationrepository.DeleteApplication", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteApplicationAsync(ctx workflow.Context, input *serverlessapplicationrepository.DeleteApplicationInput) *ServerlessApplicationRepositoryDeleteApplicationFuture {
	future := workflow.ExecuteActivity(ctx, "aws.serverlessapplicationrepository.DeleteApplication", input)
	return &ServerlessApplicationRepositoryDeleteApplicationFuture{Future: future}
}

func (a *stub) GetApplication(ctx workflow.Context, input *serverlessapplicationrepository.GetApplicationInput) (*serverlessapplicationrepository.GetApplicationOutput, error) {
	var output serverlessapplicationrepository.GetApplicationOutput
	err := workflow.ExecuteActivity(ctx, "aws.serverlessapplicationrepository.GetApplication", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetApplicationAsync(ctx workflow.Context, input *serverlessapplicationrepository.GetApplicationInput) *ServerlessApplicationRepositoryGetApplicationFuture {
	future := workflow.ExecuteActivity(ctx, "aws.serverlessapplicationrepository.GetApplication", input)
	return &ServerlessApplicationRepositoryGetApplicationFuture{Future: future}
}

func (a *stub) GetApplicationPolicy(ctx workflow.Context, input *serverlessapplicationrepository.GetApplicationPolicyInput) (*serverlessapplicationrepository.GetApplicationPolicyOutput, error) {
	var output serverlessapplicationrepository.GetApplicationPolicyOutput
	err := workflow.ExecuteActivity(ctx, "aws.serverlessapplicationrepository.GetApplicationPolicy", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetApplicationPolicyAsync(ctx workflow.Context, input *serverlessapplicationrepository.GetApplicationPolicyInput) *ServerlessApplicationRepositoryGetApplicationPolicyFuture {
	future := workflow.ExecuteActivity(ctx, "aws.serverlessapplicationrepository.GetApplicationPolicy", input)
	return &ServerlessApplicationRepositoryGetApplicationPolicyFuture{Future: future}
}

func (a *stub) GetCloudFormationTemplate(ctx workflow.Context, input *serverlessapplicationrepository.GetCloudFormationTemplateInput) (*serverlessapplicationrepository.GetCloudFormationTemplateOutput, error) {
	var output serverlessapplicationrepository.GetCloudFormationTemplateOutput
	err := workflow.ExecuteActivity(ctx, "aws.serverlessapplicationrepository.GetCloudFormationTemplate", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetCloudFormationTemplateAsync(ctx workflow.Context, input *serverlessapplicationrepository.GetCloudFormationTemplateInput) *ServerlessApplicationRepositoryGetCloudFormationTemplateFuture {
	future := workflow.ExecuteActivity(ctx, "aws.serverlessapplicationrepository.GetCloudFormationTemplate", input)
	return &ServerlessApplicationRepositoryGetCloudFormationTemplateFuture{Future: future}
}

func (a *stub) ListApplicationDependencies(ctx workflow.Context, input *serverlessapplicationrepository.ListApplicationDependenciesInput) (*serverlessapplicationrepository.ListApplicationDependenciesOutput, error) {
	var output serverlessapplicationrepository.ListApplicationDependenciesOutput
	err := workflow.ExecuteActivity(ctx, "aws.serverlessapplicationrepository.ListApplicationDependencies", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListApplicationDependenciesAsync(ctx workflow.Context, input *serverlessapplicationrepository.ListApplicationDependenciesInput) *ServerlessApplicationRepositoryListApplicationDependenciesFuture {
	future := workflow.ExecuteActivity(ctx, "aws.serverlessapplicationrepository.ListApplicationDependencies", input)
	return &ServerlessApplicationRepositoryListApplicationDependenciesFuture{Future: future}
}

func (a *stub) ListApplicationVersions(ctx workflow.Context, input *serverlessapplicationrepository.ListApplicationVersionsInput) (*serverlessapplicationrepository.ListApplicationVersionsOutput, error) {
	var output serverlessapplicationrepository.ListApplicationVersionsOutput
	err := workflow.ExecuteActivity(ctx, "aws.serverlessapplicationrepository.ListApplicationVersions", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListApplicationVersionsAsync(ctx workflow.Context, input *serverlessapplicationrepository.ListApplicationVersionsInput) *ServerlessApplicationRepositoryListApplicationVersionsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.serverlessapplicationrepository.ListApplicationVersions", input)
	return &ServerlessApplicationRepositoryListApplicationVersionsFuture{Future: future}
}

func (a *stub) ListApplications(ctx workflow.Context, input *serverlessapplicationrepository.ListApplicationsInput) (*serverlessapplicationrepository.ListApplicationsOutput, error) {
	var output serverlessapplicationrepository.ListApplicationsOutput
	err := workflow.ExecuteActivity(ctx, "aws.serverlessapplicationrepository.ListApplications", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListApplicationsAsync(ctx workflow.Context, input *serverlessapplicationrepository.ListApplicationsInput) *ServerlessApplicationRepositoryListApplicationsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.serverlessapplicationrepository.ListApplications", input)
	return &ServerlessApplicationRepositoryListApplicationsFuture{Future: future}
}

func (a *stub) PutApplicationPolicy(ctx workflow.Context, input *serverlessapplicationrepository.PutApplicationPolicyInput) (*serverlessapplicationrepository.PutApplicationPolicyOutput, error) {
	var output serverlessapplicationrepository.PutApplicationPolicyOutput
	err := workflow.ExecuteActivity(ctx, "aws.serverlessapplicationrepository.PutApplicationPolicy", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) PutApplicationPolicyAsync(ctx workflow.Context, input *serverlessapplicationrepository.PutApplicationPolicyInput) *ServerlessApplicationRepositoryPutApplicationPolicyFuture {
	future := workflow.ExecuteActivity(ctx, "aws.serverlessapplicationrepository.PutApplicationPolicy", input)
	return &ServerlessApplicationRepositoryPutApplicationPolicyFuture{Future: future}
}

func (a *stub) UnshareApplication(ctx workflow.Context, input *serverlessapplicationrepository.UnshareApplicationInput) (*serverlessapplicationrepository.UnshareApplicationOutput, error) {
	var output serverlessapplicationrepository.UnshareApplicationOutput
	err := workflow.ExecuteActivity(ctx, "aws.serverlessapplicationrepository.UnshareApplication", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UnshareApplicationAsync(ctx workflow.Context, input *serverlessapplicationrepository.UnshareApplicationInput) *ServerlessApplicationRepositoryUnshareApplicationFuture {
	future := workflow.ExecuteActivity(ctx, "aws.serverlessapplicationrepository.UnshareApplication", input)
	return &ServerlessApplicationRepositoryUnshareApplicationFuture{Future: future}
}

func (a *stub) UpdateApplication(ctx workflow.Context, input *serverlessapplicationrepository.UpdateApplicationRequest) (*serverlessapplicationrepository.UpdateApplicationOutput, error) {
	var output serverlessapplicationrepository.UpdateApplicationOutput
	err := workflow.ExecuteActivity(ctx, "aws.serverlessapplicationrepository.UpdateApplication", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateApplicationAsync(ctx workflow.Context, input *serverlessapplicationrepository.UpdateApplicationRequest) *ServerlessApplicationRepositoryUpdateApplicationFuture {
	future := workflow.ExecuteActivity(ctx, "aws.serverlessapplicationrepository.UpdateApplication", input)
	return &ServerlessApplicationRepositoryUpdateApplicationFuture{Future: future}
}