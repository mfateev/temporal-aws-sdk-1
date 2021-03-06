// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package applicationdiscoveryservicestub

import (
	"github.com/aws/aws-sdk-go/service/applicationdiscoveryservice"
	"go.temporal.io/sdk/workflow"

	"go.temporal.io/aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type stub struct{}

type ApplicationDiscoveryServiceAssociateConfigurationItemsToApplicationFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ApplicationDiscoveryServiceAssociateConfigurationItemsToApplicationFuture) Get(ctx workflow.Context) (*applicationdiscoveryservice.AssociateConfigurationItemsToApplicationOutput, error) {
	var output applicationdiscoveryservice.AssociateConfigurationItemsToApplicationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ApplicationDiscoveryServiceBatchDeleteImportDataFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ApplicationDiscoveryServiceBatchDeleteImportDataFuture) Get(ctx workflow.Context) (*applicationdiscoveryservice.BatchDeleteImportDataOutput, error) {
	var output applicationdiscoveryservice.BatchDeleteImportDataOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ApplicationDiscoveryServiceCreateApplicationFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ApplicationDiscoveryServiceCreateApplicationFuture) Get(ctx workflow.Context) (*applicationdiscoveryservice.CreateApplicationOutput, error) {
	var output applicationdiscoveryservice.CreateApplicationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ApplicationDiscoveryServiceCreateTagsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ApplicationDiscoveryServiceCreateTagsFuture) Get(ctx workflow.Context) (*applicationdiscoveryservice.CreateTagsOutput, error) {
	var output applicationdiscoveryservice.CreateTagsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ApplicationDiscoveryServiceDeleteApplicationsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ApplicationDiscoveryServiceDeleteApplicationsFuture) Get(ctx workflow.Context) (*applicationdiscoveryservice.DeleteApplicationsOutput, error) {
	var output applicationdiscoveryservice.DeleteApplicationsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ApplicationDiscoveryServiceDeleteTagsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ApplicationDiscoveryServiceDeleteTagsFuture) Get(ctx workflow.Context) (*applicationdiscoveryservice.DeleteTagsOutput, error) {
	var output applicationdiscoveryservice.DeleteTagsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ApplicationDiscoveryServiceDescribeAgentsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ApplicationDiscoveryServiceDescribeAgentsFuture) Get(ctx workflow.Context) (*applicationdiscoveryservice.DescribeAgentsOutput, error) {
	var output applicationdiscoveryservice.DescribeAgentsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ApplicationDiscoveryServiceDescribeConfigurationsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ApplicationDiscoveryServiceDescribeConfigurationsFuture) Get(ctx workflow.Context) (*applicationdiscoveryservice.DescribeConfigurationsOutput, error) {
	var output applicationdiscoveryservice.DescribeConfigurationsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ApplicationDiscoveryServiceDescribeContinuousExportsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ApplicationDiscoveryServiceDescribeContinuousExportsFuture) Get(ctx workflow.Context) (*applicationdiscoveryservice.DescribeContinuousExportsOutput, error) {
	var output applicationdiscoveryservice.DescribeContinuousExportsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ApplicationDiscoveryServiceDescribeExportConfigurationsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ApplicationDiscoveryServiceDescribeExportConfigurationsFuture) Get(ctx workflow.Context) (*applicationdiscoveryservice.DescribeExportConfigurationsOutput, error) {
	var output applicationdiscoveryservice.DescribeExportConfigurationsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ApplicationDiscoveryServiceDescribeExportTasksFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ApplicationDiscoveryServiceDescribeExportTasksFuture) Get(ctx workflow.Context) (*applicationdiscoveryservice.DescribeExportTasksOutput, error) {
	var output applicationdiscoveryservice.DescribeExportTasksOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ApplicationDiscoveryServiceDescribeImportTasksFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ApplicationDiscoveryServiceDescribeImportTasksFuture) Get(ctx workflow.Context) (*applicationdiscoveryservice.DescribeImportTasksOutput, error) {
	var output applicationdiscoveryservice.DescribeImportTasksOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ApplicationDiscoveryServiceDescribeTagsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ApplicationDiscoveryServiceDescribeTagsFuture) Get(ctx workflow.Context) (*applicationdiscoveryservice.DescribeTagsOutput, error) {
	var output applicationdiscoveryservice.DescribeTagsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ApplicationDiscoveryServiceDisassociateConfigurationItemsFromApplicationFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ApplicationDiscoveryServiceDisassociateConfigurationItemsFromApplicationFuture) Get(ctx workflow.Context) (*applicationdiscoveryservice.DisassociateConfigurationItemsFromApplicationOutput, error) {
	var output applicationdiscoveryservice.DisassociateConfigurationItemsFromApplicationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ApplicationDiscoveryServiceExportConfigurationsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ApplicationDiscoveryServiceExportConfigurationsFuture) Get(ctx workflow.Context) (*applicationdiscoveryservice.ExportConfigurationsOutput, error) {
	var output applicationdiscoveryservice.ExportConfigurationsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ApplicationDiscoveryServiceGetDiscoverySummaryFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ApplicationDiscoveryServiceGetDiscoverySummaryFuture) Get(ctx workflow.Context) (*applicationdiscoveryservice.GetDiscoverySummaryOutput, error) {
	var output applicationdiscoveryservice.GetDiscoverySummaryOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ApplicationDiscoveryServiceListConfigurationsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ApplicationDiscoveryServiceListConfigurationsFuture) Get(ctx workflow.Context) (*applicationdiscoveryservice.ListConfigurationsOutput, error) {
	var output applicationdiscoveryservice.ListConfigurationsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ApplicationDiscoveryServiceListServerNeighborsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ApplicationDiscoveryServiceListServerNeighborsFuture) Get(ctx workflow.Context) (*applicationdiscoveryservice.ListServerNeighborsOutput, error) {
	var output applicationdiscoveryservice.ListServerNeighborsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ApplicationDiscoveryServiceStartContinuousExportFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ApplicationDiscoveryServiceStartContinuousExportFuture) Get(ctx workflow.Context) (*applicationdiscoveryservice.StartContinuousExportOutput, error) {
	var output applicationdiscoveryservice.StartContinuousExportOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ApplicationDiscoveryServiceStartDataCollectionByAgentIdsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ApplicationDiscoveryServiceStartDataCollectionByAgentIdsFuture) Get(ctx workflow.Context) (*applicationdiscoveryservice.StartDataCollectionByAgentIdsOutput, error) {
	var output applicationdiscoveryservice.StartDataCollectionByAgentIdsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ApplicationDiscoveryServiceStartExportTaskFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ApplicationDiscoveryServiceStartExportTaskFuture) Get(ctx workflow.Context) (*applicationdiscoveryservice.StartExportTaskOutput, error) {
	var output applicationdiscoveryservice.StartExportTaskOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ApplicationDiscoveryServiceStartImportTaskFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ApplicationDiscoveryServiceStartImportTaskFuture) Get(ctx workflow.Context) (*applicationdiscoveryservice.StartImportTaskOutput, error) {
	var output applicationdiscoveryservice.StartImportTaskOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ApplicationDiscoveryServiceStopContinuousExportFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ApplicationDiscoveryServiceStopContinuousExportFuture) Get(ctx workflow.Context) (*applicationdiscoveryservice.StopContinuousExportOutput, error) {
	var output applicationdiscoveryservice.StopContinuousExportOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ApplicationDiscoveryServiceStopDataCollectionByAgentIdsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ApplicationDiscoveryServiceStopDataCollectionByAgentIdsFuture) Get(ctx workflow.Context) (*applicationdiscoveryservice.StopDataCollectionByAgentIdsOutput, error) {
	var output applicationdiscoveryservice.StopDataCollectionByAgentIdsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ApplicationDiscoveryServiceUpdateApplicationFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ApplicationDiscoveryServiceUpdateApplicationFuture) Get(ctx workflow.Context) (*applicationdiscoveryservice.UpdateApplicationOutput, error) {
	var output applicationdiscoveryservice.UpdateApplicationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *stub) AssociateConfigurationItemsToApplication(ctx workflow.Context, input *applicationdiscoveryservice.AssociateConfigurationItemsToApplicationInput) (*applicationdiscoveryservice.AssociateConfigurationItemsToApplicationOutput, error) {
	var output applicationdiscoveryservice.AssociateConfigurationItemsToApplicationOutput
	err := workflow.ExecuteActivity(ctx, "aws.applicationdiscoveryservice.AssociateConfigurationItemsToApplication", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) AssociateConfigurationItemsToApplicationAsync(ctx workflow.Context, input *applicationdiscoveryservice.AssociateConfigurationItemsToApplicationInput) *ApplicationDiscoveryServiceAssociateConfigurationItemsToApplicationFuture {
	future := workflow.ExecuteActivity(ctx, "aws.applicationdiscoveryservice.AssociateConfigurationItemsToApplication", input)
	return &ApplicationDiscoveryServiceAssociateConfigurationItemsToApplicationFuture{Future: future}
}

func (a *stub) BatchDeleteImportData(ctx workflow.Context, input *applicationdiscoveryservice.BatchDeleteImportDataInput) (*applicationdiscoveryservice.BatchDeleteImportDataOutput, error) {
	var output applicationdiscoveryservice.BatchDeleteImportDataOutput
	err := workflow.ExecuteActivity(ctx, "aws.applicationdiscoveryservice.BatchDeleteImportData", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) BatchDeleteImportDataAsync(ctx workflow.Context, input *applicationdiscoveryservice.BatchDeleteImportDataInput) *ApplicationDiscoveryServiceBatchDeleteImportDataFuture {
	future := workflow.ExecuteActivity(ctx, "aws.applicationdiscoveryservice.BatchDeleteImportData", input)
	return &ApplicationDiscoveryServiceBatchDeleteImportDataFuture{Future: future}
}

func (a *stub) CreateApplication(ctx workflow.Context, input *applicationdiscoveryservice.CreateApplicationInput) (*applicationdiscoveryservice.CreateApplicationOutput, error) {
	var output applicationdiscoveryservice.CreateApplicationOutput
	err := workflow.ExecuteActivity(ctx, "aws.applicationdiscoveryservice.CreateApplication", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateApplicationAsync(ctx workflow.Context, input *applicationdiscoveryservice.CreateApplicationInput) *ApplicationDiscoveryServiceCreateApplicationFuture {
	future := workflow.ExecuteActivity(ctx, "aws.applicationdiscoveryservice.CreateApplication", input)
	return &ApplicationDiscoveryServiceCreateApplicationFuture{Future: future}
}

func (a *stub) CreateTags(ctx workflow.Context, input *applicationdiscoveryservice.CreateTagsInput) (*applicationdiscoveryservice.CreateTagsOutput, error) {
	var output applicationdiscoveryservice.CreateTagsOutput
	err := workflow.ExecuteActivity(ctx, "aws.applicationdiscoveryservice.CreateTags", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateTagsAsync(ctx workflow.Context, input *applicationdiscoveryservice.CreateTagsInput) *ApplicationDiscoveryServiceCreateTagsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.applicationdiscoveryservice.CreateTags", input)
	return &ApplicationDiscoveryServiceCreateTagsFuture{Future: future}
}

func (a *stub) DeleteApplications(ctx workflow.Context, input *applicationdiscoveryservice.DeleteApplicationsInput) (*applicationdiscoveryservice.DeleteApplicationsOutput, error) {
	var output applicationdiscoveryservice.DeleteApplicationsOutput
	err := workflow.ExecuteActivity(ctx, "aws.applicationdiscoveryservice.DeleteApplications", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteApplicationsAsync(ctx workflow.Context, input *applicationdiscoveryservice.DeleteApplicationsInput) *ApplicationDiscoveryServiceDeleteApplicationsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.applicationdiscoveryservice.DeleteApplications", input)
	return &ApplicationDiscoveryServiceDeleteApplicationsFuture{Future: future}
}

func (a *stub) DeleteTags(ctx workflow.Context, input *applicationdiscoveryservice.DeleteTagsInput) (*applicationdiscoveryservice.DeleteTagsOutput, error) {
	var output applicationdiscoveryservice.DeleteTagsOutput
	err := workflow.ExecuteActivity(ctx, "aws.applicationdiscoveryservice.DeleteTags", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteTagsAsync(ctx workflow.Context, input *applicationdiscoveryservice.DeleteTagsInput) *ApplicationDiscoveryServiceDeleteTagsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.applicationdiscoveryservice.DeleteTags", input)
	return &ApplicationDiscoveryServiceDeleteTagsFuture{Future: future}
}

func (a *stub) DescribeAgents(ctx workflow.Context, input *applicationdiscoveryservice.DescribeAgentsInput) (*applicationdiscoveryservice.DescribeAgentsOutput, error) {
	var output applicationdiscoveryservice.DescribeAgentsOutput
	err := workflow.ExecuteActivity(ctx, "aws.applicationdiscoveryservice.DescribeAgents", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeAgentsAsync(ctx workflow.Context, input *applicationdiscoveryservice.DescribeAgentsInput) *ApplicationDiscoveryServiceDescribeAgentsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.applicationdiscoveryservice.DescribeAgents", input)
	return &ApplicationDiscoveryServiceDescribeAgentsFuture{Future: future}
}

func (a *stub) DescribeConfigurations(ctx workflow.Context, input *applicationdiscoveryservice.DescribeConfigurationsInput) (*applicationdiscoveryservice.DescribeConfigurationsOutput, error) {
	var output applicationdiscoveryservice.DescribeConfigurationsOutput
	err := workflow.ExecuteActivity(ctx, "aws.applicationdiscoveryservice.DescribeConfigurations", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeConfigurationsAsync(ctx workflow.Context, input *applicationdiscoveryservice.DescribeConfigurationsInput) *ApplicationDiscoveryServiceDescribeConfigurationsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.applicationdiscoveryservice.DescribeConfigurations", input)
	return &ApplicationDiscoveryServiceDescribeConfigurationsFuture{Future: future}
}

func (a *stub) DescribeContinuousExports(ctx workflow.Context, input *applicationdiscoveryservice.DescribeContinuousExportsInput) (*applicationdiscoveryservice.DescribeContinuousExportsOutput, error) {
	var output applicationdiscoveryservice.DescribeContinuousExportsOutput
	err := workflow.ExecuteActivity(ctx, "aws.applicationdiscoveryservice.DescribeContinuousExports", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeContinuousExportsAsync(ctx workflow.Context, input *applicationdiscoveryservice.DescribeContinuousExportsInput) *ApplicationDiscoveryServiceDescribeContinuousExportsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.applicationdiscoveryservice.DescribeContinuousExports", input)
	return &ApplicationDiscoveryServiceDescribeContinuousExportsFuture{Future: future}
}

func (a *stub) DescribeExportConfigurations(ctx workflow.Context, input *applicationdiscoveryservice.DescribeExportConfigurationsInput) (*applicationdiscoveryservice.DescribeExportConfigurationsOutput, error) {
	var output applicationdiscoveryservice.DescribeExportConfigurationsOutput
	err := workflow.ExecuteActivity(ctx, "aws.applicationdiscoveryservice.DescribeExportConfigurations", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeExportConfigurationsAsync(ctx workflow.Context, input *applicationdiscoveryservice.DescribeExportConfigurationsInput) *ApplicationDiscoveryServiceDescribeExportConfigurationsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.applicationdiscoveryservice.DescribeExportConfigurations", input)
	return &ApplicationDiscoveryServiceDescribeExportConfigurationsFuture{Future: future}
}

func (a *stub) DescribeExportTasks(ctx workflow.Context, input *applicationdiscoveryservice.DescribeExportTasksInput) (*applicationdiscoveryservice.DescribeExportTasksOutput, error) {
	var output applicationdiscoveryservice.DescribeExportTasksOutput
	err := workflow.ExecuteActivity(ctx, "aws.applicationdiscoveryservice.DescribeExportTasks", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeExportTasksAsync(ctx workflow.Context, input *applicationdiscoveryservice.DescribeExportTasksInput) *ApplicationDiscoveryServiceDescribeExportTasksFuture {
	future := workflow.ExecuteActivity(ctx, "aws.applicationdiscoveryservice.DescribeExportTasks", input)
	return &ApplicationDiscoveryServiceDescribeExportTasksFuture{Future: future}
}

func (a *stub) DescribeImportTasks(ctx workflow.Context, input *applicationdiscoveryservice.DescribeImportTasksInput) (*applicationdiscoveryservice.DescribeImportTasksOutput, error) {
	var output applicationdiscoveryservice.DescribeImportTasksOutput
	err := workflow.ExecuteActivity(ctx, "aws.applicationdiscoveryservice.DescribeImportTasks", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeImportTasksAsync(ctx workflow.Context, input *applicationdiscoveryservice.DescribeImportTasksInput) *ApplicationDiscoveryServiceDescribeImportTasksFuture {
	future := workflow.ExecuteActivity(ctx, "aws.applicationdiscoveryservice.DescribeImportTasks", input)
	return &ApplicationDiscoveryServiceDescribeImportTasksFuture{Future: future}
}

func (a *stub) DescribeTags(ctx workflow.Context, input *applicationdiscoveryservice.DescribeTagsInput) (*applicationdiscoveryservice.DescribeTagsOutput, error) {
	var output applicationdiscoveryservice.DescribeTagsOutput
	err := workflow.ExecuteActivity(ctx, "aws.applicationdiscoveryservice.DescribeTags", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeTagsAsync(ctx workflow.Context, input *applicationdiscoveryservice.DescribeTagsInput) *ApplicationDiscoveryServiceDescribeTagsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.applicationdiscoveryservice.DescribeTags", input)
	return &ApplicationDiscoveryServiceDescribeTagsFuture{Future: future}
}

func (a *stub) DisassociateConfigurationItemsFromApplication(ctx workflow.Context, input *applicationdiscoveryservice.DisassociateConfigurationItemsFromApplicationInput) (*applicationdiscoveryservice.DisassociateConfigurationItemsFromApplicationOutput, error) {
	var output applicationdiscoveryservice.DisassociateConfigurationItemsFromApplicationOutput
	err := workflow.ExecuteActivity(ctx, "aws.applicationdiscoveryservice.DisassociateConfigurationItemsFromApplication", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DisassociateConfigurationItemsFromApplicationAsync(ctx workflow.Context, input *applicationdiscoveryservice.DisassociateConfigurationItemsFromApplicationInput) *ApplicationDiscoveryServiceDisassociateConfigurationItemsFromApplicationFuture {
	future := workflow.ExecuteActivity(ctx, "aws.applicationdiscoveryservice.DisassociateConfigurationItemsFromApplication", input)
	return &ApplicationDiscoveryServiceDisassociateConfigurationItemsFromApplicationFuture{Future: future}
}

func (a *stub) ExportConfigurations(ctx workflow.Context, input *applicationdiscoveryservice.ExportConfigurationsInput) (*applicationdiscoveryservice.ExportConfigurationsOutput, error) {
	var output applicationdiscoveryservice.ExportConfigurationsOutput
	err := workflow.ExecuteActivity(ctx, "aws.applicationdiscoveryservice.ExportConfigurations", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ExportConfigurationsAsync(ctx workflow.Context, input *applicationdiscoveryservice.ExportConfigurationsInput) *ApplicationDiscoveryServiceExportConfigurationsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.applicationdiscoveryservice.ExportConfigurations", input)
	return &ApplicationDiscoveryServiceExportConfigurationsFuture{Future: future}
}

func (a *stub) GetDiscoverySummary(ctx workflow.Context, input *applicationdiscoveryservice.GetDiscoverySummaryInput) (*applicationdiscoveryservice.GetDiscoverySummaryOutput, error) {
	var output applicationdiscoveryservice.GetDiscoverySummaryOutput
	err := workflow.ExecuteActivity(ctx, "aws.applicationdiscoveryservice.GetDiscoverySummary", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetDiscoverySummaryAsync(ctx workflow.Context, input *applicationdiscoveryservice.GetDiscoverySummaryInput) *ApplicationDiscoveryServiceGetDiscoverySummaryFuture {
	future := workflow.ExecuteActivity(ctx, "aws.applicationdiscoveryservice.GetDiscoverySummary", input)
	return &ApplicationDiscoveryServiceGetDiscoverySummaryFuture{Future: future}
}

func (a *stub) ListConfigurations(ctx workflow.Context, input *applicationdiscoveryservice.ListConfigurationsInput) (*applicationdiscoveryservice.ListConfigurationsOutput, error) {
	var output applicationdiscoveryservice.ListConfigurationsOutput
	err := workflow.ExecuteActivity(ctx, "aws.applicationdiscoveryservice.ListConfigurations", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListConfigurationsAsync(ctx workflow.Context, input *applicationdiscoveryservice.ListConfigurationsInput) *ApplicationDiscoveryServiceListConfigurationsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.applicationdiscoveryservice.ListConfigurations", input)
	return &ApplicationDiscoveryServiceListConfigurationsFuture{Future: future}
}

func (a *stub) ListServerNeighbors(ctx workflow.Context, input *applicationdiscoveryservice.ListServerNeighborsInput) (*applicationdiscoveryservice.ListServerNeighborsOutput, error) {
	var output applicationdiscoveryservice.ListServerNeighborsOutput
	err := workflow.ExecuteActivity(ctx, "aws.applicationdiscoveryservice.ListServerNeighbors", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListServerNeighborsAsync(ctx workflow.Context, input *applicationdiscoveryservice.ListServerNeighborsInput) *ApplicationDiscoveryServiceListServerNeighborsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.applicationdiscoveryservice.ListServerNeighbors", input)
	return &ApplicationDiscoveryServiceListServerNeighborsFuture{Future: future}
}

func (a *stub) StartContinuousExport(ctx workflow.Context, input *applicationdiscoveryservice.StartContinuousExportInput) (*applicationdiscoveryservice.StartContinuousExportOutput, error) {
	var output applicationdiscoveryservice.StartContinuousExportOutput
	err := workflow.ExecuteActivity(ctx, "aws.applicationdiscoveryservice.StartContinuousExport", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) StartContinuousExportAsync(ctx workflow.Context, input *applicationdiscoveryservice.StartContinuousExportInput) *ApplicationDiscoveryServiceStartContinuousExportFuture {
	future := workflow.ExecuteActivity(ctx, "aws.applicationdiscoveryservice.StartContinuousExport", input)
	return &ApplicationDiscoveryServiceStartContinuousExportFuture{Future: future}
}

func (a *stub) StartDataCollectionByAgentIds(ctx workflow.Context, input *applicationdiscoveryservice.StartDataCollectionByAgentIdsInput) (*applicationdiscoveryservice.StartDataCollectionByAgentIdsOutput, error) {
	var output applicationdiscoveryservice.StartDataCollectionByAgentIdsOutput
	err := workflow.ExecuteActivity(ctx, "aws.applicationdiscoveryservice.StartDataCollectionByAgentIds", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) StartDataCollectionByAgentIdsAsync(ctx workflow.Context, input *applicationdiscoveryservice.StartDataCollectionByAgentIdsInput) *ApplicationDiscoveryServiceStartDataCollectionByAgentIdsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.applicationdiscoveryservice.StartDataCollectionByAgentIds", input)
	return &ApplicationDiscoveryServiceStartDataCollectionByAgentIdsFuture{Future: future}
}

func (a *stub) StartExportTask(ctx workflow.Context, input *applicationdiscoveryservice.StartExportTaskInput) (*applicationdiscoveryservice.StartExportTaskOutput, error) {
	var output applicationdiscoveryservice.StartExportTaskOutput
	err := workflow.ExecuteActivity(ctx, "aws.applicationdiscoveryservice.StartExportTask", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) StartExportTaskAsync(ctx workflow.Context, input *applicationdiscoveryservice.StartExportTaskInput) *ApplicationDiscoveryServiceStartExportTaskFuture {
	future := workflow.ExecuteActivity(ctx, "aws.applicationdiscoveryservice.StartExportTask", input)
	return &ApplicationDiscoveryServiceStartExportTaskFuture{Future: future}
}

func (a *stub) StartImportTask(ctx workflow.Context, input *applicationdiscoveryservice.StartImportTaskInput) (*applicationdiscoveryservice.StartImportTaskOutput, error) {
	var output applicationdiscoveryservice.StartImportTaskOutput
	err := workflow.ExecuteActivity(ctx, "aws.applicationdiscoveryservice.StartImportTask", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) StartImportTaskAsync(ctx workflow.Context, input *applicationdiscoveryservice.StartImportTaskInput) *ApplicationDiscoveryServiceStartImportTaskFuture {
	future := workflow.ExecuteActivity(ctx, "aws.applicationdiscoveryservice.StartImportTask", input)
	return &ApplicationDiscoveryServiceStartImportTaskFuture{Future: future}
}

func (a *stub) StopContinuousExport(ctx workflow.Context, input *applicationdiscoveryservice.StopContinuousExportInput) (*applicationdiscoveryservice.StopContinuousExportOutput, error) {
	var output applicationdiscoveryservice.StopContinuousExportOutput
	err := workflow.ExecuteActivity(ctx, "aws.applicationdiscoveryservice.StopContinuousExport", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) StopContinuousExportAsync(ctx workflow.Context, input *applicationdiscoveryservice.StopContinuousExportInput) *ApplicationDiscoveryServiceStopContinuousExportFuture {
	future := workflow.ExecuteActivity(ctx, "aws.applicationdiscoveryservice.StopContinuousExport", input)
	return &ApplicationDiscoveryServiceStopContinuousExportFuture{Future: future}
}

func (a *stub) StopDataCollectionByAgentIds(ctx workflow.Context, input *applicationdiscoveryservice.StopDataCollectionByAgentIdsInput) (*applicationdiscoveryservice.StopDataCollectionByAgentIdsOutput, error) {
	var output applicationdiscoveryservice.StopDataCollectionByAgentIdsOutput
	err := workflow.ExecuteActivity(ctx, "aws.applicationdiscoveryservice.StopDataCollectionByAgentIds", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) StopDataCollectionByAgentIdsAsync(ctx workflow.Context, input *applicationdiscoveryservice.StopDataCollectionByAgentIdsInput) *ApplicationDiscoveryServiceStopDataCollectionByAgentIdsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.applicationdiscoveryservice.StopDataCollectionByAgentIds", input)
	return &ApplicationDiscoveryServiceStopDataCollectionByAgentIdsFuture{Future: future}
}

func (a *stub) UpdateApplication(ctx workflow.Context, input *applicationdiscoveryservice.UpdateApplicationInput) (*applicationdiscoveryservice.UpdateApplicationOutput, error) {
	var output applicationdiscoveryservice.UpdateApplicationOutput
	err := workflow.ExecuteActivity(ctx, "aws.applicationdiscoveryservice.UpdateApplication", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateApplicationAsync(ctx workflow.Context, input *applicationdiscoveryservice.UpdateApplicationInput) *ApplicationDiscoveryServiceUpdateApplicationFuture {
	future := workflow.ExecuteActivity(ctx, "aws.applicationdiscoveryservice.UpdateApplication", input)
	return &ApplicationDiscoveryServiceUpdateApplicationFuture{Future: future}
}
