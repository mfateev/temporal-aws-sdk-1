package awsclients

import (
	"github.com/aws/aws-sdk-go/service/kinesisanalytics"
	"go.temporal.io/sdk/workflow"
	"temporal.io/aws-sdk/awsactivities"
)

type KinesisAnalyticsClient interface {
	AddApplicationCloudWatchLoggingOption(ctx workflow.Context, input *kinesisanalytics.AddApplicationCloudWatchLoggingOptionInput) (*kinesisanalytics.AddApplicationCloudWatchLoggingOptionOutput, error)
	AddApplicationCloudWatchLoggingOptionAsync(ctx workflow.Context, input *kinesisanalytics.AddApplicationCloudWatchLoggingOptionInput) *KinesisanalyticsAddApplicationCloudWatchLoggingOptionResult

	AddApplicationInput(ctx workflow.Context, input *kinesisanalytics.AddApplicationInputInput) (*kinesisanalytics.AddApplicationInputOutput, error)
	AddApplicationInputAsync(ctx workflow.Context, input *kinesisanalytics.AddApplicationInputInput) *KinesisanalyticsAddApplicationInputResult

	AddApplicationInputProcessingConfiguration(ctx workflow.Context, input *kinesisanalytics.AddApplicationInputProcessingConfigurationInput) (*kinesisanalytics.AddApplicationInputProcessingConfigurationOutput, error)
	AddApplicationInputProcessingConfigurationAsync(ctx workflow.Context, input *kinesisanalytics.AddApplicationInputProcessingConfigurationInput) *KinesisanalyticsAddApplicationInputProcessingConfigurationResult

	AddApplicationOutput(ctx workflow.Context, input *kinesisanalytics.AddApplicationOutputInput) (*kinesisanalytics.AddApplicationOutputOutput, error)
	AddApplicationOutputAsync(ctx workflow.Context, input *kinesisanalytics.AddApplicationOutputInput) *KinesisanalyticsAddApplicationOutputResult

	AddApplicationReferenceDataSource(ctx workflow.Context, input *kinesisanalytics.AddApplicationReferenceDataSourceInput) (*kinesisanalytics.AddApplicationReferenceDataSourceOutput, error)
	AddApplicationReferenceDataSourceAsync(ctx workflow.Context, input *kinesisanalytics.AddApplicationReferenceDataSourceInput) *KinesisanalyticsAddApplicationReferenceDataSourceResult

	CreateApplication(ctx workflow.Context, input *kinesisanalytics.CreateApplicationInput) (*kinesisanalytics.CreateApplicationOutput, error)
	CreateApplicationAsync(ctx workflow.Context, input *kinesisanalytics.CreateApplicationInput) *KinesisanalyticsCreateApplicationResult

	DeleteApplication(ctx workflow.Context, input *kinesisanalytics.DeleteApplicationInput) (*kinesisanalytics.DeleteApplicationOutput, error)
	DeleteApplicationAsync(ctx workflow.Context, input *kinesisanalytics.DeleteApplicationInput) *KinesisanalyticsDeleteApplicationResult

	DeleteApplicationCloudWatchLoggingOption(ctx workflow.Context, input *kinesisanalytics.DeleteApplicationCloudWatchLoggingOptionInput) (*kinesisanalytics.DeleteApplicationCloudWatchLoggingOptionOutput, error)
	DeleteApplicationCloudWatchLoggingOptionAsync(ctx workflow.Context, input *kinesisanalytics.DeleteApplicationCloudWatchLoggingOptionInput) *KinesisanalyticsDeleteApplicationCloudWatchLoggingOptionResult

	DeleteApplicationInputProcessingConfiguration(ctx workflow.Context, input *kinesisanalytics.DeleteApplicationInputProcessingConfigurationInput) (*kinesisanalytics.DeleteApplicationInputProcessingConfigurationOutput, error)
	DeleteApplicationInputProcessingConfigurationAsync(ctx workflow.Context, input *kinesisanalytics.DeleteApplicationInputProcessingConfigurationInput) *KinesisanalyticsDeleteApplicationInputProcessingConfigurationResult

	DeleteApplicationOutput(ctx workflow.Context, input *kinesisanalytics.DeleteApplicationOutputInput) (*kinesisanalytics.DeleteApplicationOutputOutput, error)
	DeleteApplicationOutputAsync(ctx workflow.Context, input *kinesisanalytics.DeleteApplicationOutputInput) *KinesisanalyticsDeleteApplicationOutputResult

	DeleteApplicationReferenceDataSource(ctx workflow.Context, input *kinesisanalytics.DeleteApplicationReferenceDataSourceInput) (*kinesisanalytics.DeleteApplicationReferenceDataSourceOutput, error)
	DeleteApplicationReferenceDataSourceAsync(ctx workflow.Context, input *kinesisanalytics.DeleteApplicationReferenceDataSourceInput) *KinesisanalyticsDeleteApplicationReferenceDataSourceResult

	DescribeApplication(ctx workflow.Context, input *kinesisanalytics.DescribeApplicationInput) (*kinesisanalytics.DescribeApplicationOutput, error)
	DescribeApplicationAsync(ctx workflow.Context, input *kinesisanalytics.DescribeApplicationInput) *KinesisanalyticsDescribeApplicationResult

	DiscoverInputSchema(ctx workflow.Context, input *kinesisanalytics.DiscoverInputSchemaInput) (*kinesisanalytics.DiscoverInputSchemaOutput, error)
	DiscoverInputSchemaAsync(ctx workflow.Context, input *kinesisanalytics.DiscoverInputSchemaInput) *KinesisanalyticsDiscoverInputSchemaResult

	ListApplications(ctx workflow.Context, input *kinesisanalytics.ListApplicationsInput) (*kinesisanalytics.ListApplicationsOutput, error)
	ListApplicationsAsync(ctx workflow.Context, input *kinesisanalytics.ListApplicationsInput) *KinesisanalyticsListApplicationsResult

	ListTagsForResource(ctx workflow.Context, input *kinesisanalytics.ListTagsForResourceInput) (*kinesisanalytics.ListTagsForResourceOutput, error)
	ListTagsForResourceAsync(ctx workflow.Context, input *kinesisanalytics.ListTagsForResourceInput) *KinesisanalyticsListTagsForResourceResult

	StartApplication(ctx workflow.Context, input *kinesisanalytics.StartApplicationInput) (*kinesisanalytics.StartApplicationOutput, error)
	StartApplicationAsync(ctx workflow.Context, input *kinesisanalytics.StartApplicationInput) *KinesisanalyticsStartApplicationResult

	StopApplication(ctx workflow.Context, input *kinesisanalytics.StopApplicationInput) (*kinesisanalytics.StopApplicationOutput, error)
	StopApplicationAsync(ctx workflow.Context, input *kinesisanalytics.StopApplicationInput) *KinesisanalyticsStopApplicationResult

	TagResource(ctx workflow.Context, input *kinesisanalytics.TagResourceInput) (*kinesisanalytics.TagResourceOutput, error)
	TagResourceAsync(ctx workflow.Context, input *kinesisanalytics.TagResourceInput) *KinesisanalyticsTagResourceResult

	UntagResource(ctx workflow.Context, input *kinesisanalytics.UntagResourceInput) (*kinesisanalytics.UntagResourceOutput, error)
	UntagResourceAsync(ctx workflow.Context, input *kinesisanalytics.UntagResourceInput) *KinesisanalyticsUntagResourceResult

	UpdateApplication(ctx workflow.Context, input *kinesisanalytics.UpdateApplicationInput) (*kinesisanalytics.UpdateApplicationOutput, error)
	UpdateApplicationAsync(ctx workflow.Context, input *kinesisanalytics.UpdateApplicationInput) *KinesisanalyticsUpdateApplicationResult
}

type KinesisanalyticsAddApplicationCloudWatchLoggingOptionResult struct {
	Result workflow.Future
}

func (r *KinesisanalyticsAddApplicationCloudWatchLoggingOptionResult) Get(ctx workflow.Context) (*kinesisanalytics.AddApplicationCloudWatchLoggingOptionOutput, error) {
	var output kinesisanalytics.AddApplicationCloudWatchLoggingOptionOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type KinesisanalyticsAddApplicationInputResult struct {
	Result workflow.Future
}

func (r *KinesisanalyticsAddApplicationInputResult) Get(ctx workflow.Context) (*kinesisanalytics.AddApplicationInputOutput, error) {
	var output kinesisanalytics.AddApplicationInputOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type KinesisanalyticsAddApplicationInputProcessingConfigurationResult struct {
	Result workflow.Future
}

func (r *KinesisanalyticsAddApplicationInputProcessingConfigurationResult) Get(ctx workflow.Context) (*kinesisanalytics.AddApplicationInputProcessingConfigurationOutput, error) {
	var output kinesisanalytics.AddApplicationInputProcessingConfigurationOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type KinesisanalyticsAddApplicationOutputResult struct {
	Result workflow.Future
}

func (r *KinesisanalyticsAddApplicationOutputResult) Get(ctx workflow.Context) (*kinesisanalytics.AddApplicationOutputOutput, error) {
	var output kinesisanalytics.AddApplicationOutputOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type KinesisanalyticsAddApplicationReferenceDataSourceResult struct {
	Result workflow.Future
}

func (r *KinesisanalyticsAddApplicationReferenceDataSourceResult) Get(ctx workflow.Context) (*kinesisanalytics.AddApplicationReferenceDataSourceOutput, error) {
	var output kinesisanalytics.AddApplicationReferenceDataSourceOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type KinesisanalyticsCreateApplicationResult struct {
	Result workflow.Future
}

func (r *KinesisanalyticsCreateApplicationResult) Get(ctx workflow.Context) (*kinesisanalytics.CreateApplicationOutput, error) {
	var output kinesisanalytics.CreateApplicationOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type KinesisanalyticsDeleteApplicationResult struct {
	Result workflow.Future
}

func (r *KinesisanalyticsDeleteApplicationResult) Get(ctx workflow.Context) (*kinesisanalytics.DeleteApplicationOutput, error) {
	var output kinesisanalytics.DeleteApplicationOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type KinesisanalyticsDeleteApplicationCloudWatchLoggingOptionResult struct {
	Result workflow.Future
}

func (r *KinesisanalyticsDeleteApplicationCloudWatchLoggingOptionResult) Get(ctx workflow.Context) (*kinesisanalytics.DeleteApplicationCloudWatchLoggingOptionOutput, error) {
	var output kinesisanalytics.DeleteApplicationCloudWatchLoggingOptionOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type KinesisanalyticsDeleteApplicationInputProcessingConfigurationResult struct {
	Result workflow.Future
}

func (r *KinesisanalyticsDeleteApplicationInputProcessingConfigurationResult) Get(ctx workflow.Context) (*kinesisanalytics.DeleteApplicationInputProcessingConfigurationOutput, error) {
	var output kinesisanalytics.DeleteApplicationInputProcessingConfigurationOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type KinesisanalyticsDeleteApplicationOutputResult struct {
	Result workflow.Future
}

func (r *KinesisanalyticsDeleteApplicationOutputResult) Get(ctx workflow.Context) (*kinesisanalytics.DeleteApplicationOutputOutput, error) {
	var output kinesisanalytics.DeleteApplicationOutputOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type KinesisanalyticsDeleteApplicationReferenceDataSourceResult struct {
	Result workflow.Future
}

func (r *KinesisanalyticsDeleteApplicationReferenceDataSourceResult) Get(ctx workflow.Context) (*kinesisanalytics.DeleteApplicationReferenceDataSourceOutput, error) {
	var output kinesisanalytics.DeleteApplicationReferenceDataSourceOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type KinesisanalyticsDescribeApplicationResult struct {
	Result workflow.Future
}

func (r *KinesisanalyticsDescribeApplicationResult) Get(ctx workflow.Context) (*kinesisanalytics.DescribeApplicationOutput, error) {
	var output kinesisanalytics.DescribeApplicationOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type KinesisanalyticsDiscoverInputSchemaResult struct {
	Result workflow.Future
}

func (r *KinesisanalyticsDiscoverInputSchemaResult) Get(ctx workflow.Context) (*kinesisanalytics.DiscoverInputSchemaOutput, error) {
	var output kinesisanalytics.DiscoverInputSchemaOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type KinesisanalyticsListApplicationsResult struct {
	Result workflow.Future
}

func (r *KinesisanalyticsListApplicationsResult) Get(ctx workflow.Context) (*kinesisanalytics.ListApplicationsOutput, error) {
	var output kinesisanalytics.ListApplicationsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type KinesisanalyticsListTagsForResourceResult struct {
	Result workflow.Future
}

func (r *KinesisanalyticsListTagsForResourceResult) Get(ctx workflow.Context) (*kinesisanalytics.ListTagsForResourceOutput, error) {
	var output kinesisanalytics.ListTagsForResourceOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type KinesisanalyticsStartApplicationResult struct {
	Result workflow.Future
}

func (r *KinesisanalyticsStartApplicationResult) Get(ctx workflow.Context) (*kinesisanalytics.StartApplicationOutput, error) {
	var output kinesisanalytics.StartApplicationOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type KinesisanalyticsStopApplicationResult struct {
	Result workflow.Future
}

func (r *KinesisanalyticsStopApplicationResult) Get(ctx workflow.Context) (*kinesisanalytics.StopApplicationOutput, error) {
	var output kinesisanalytics.StopApplicationOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type KinesisanalyticsTagResourceResult struct {
	Result workflow.Future
}

func (r *KinesisanalyticsTagResourceResult) Get(ctx workflow.Context) (*kinesisanalytics.TagResourceOutput, error) {
	var output kinesisanalytics.TagResourceOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type KinesisanalyticsUntagResourceResult struct {
	Result workflow.Future
}

func (r *KinesisanalyticsUntagResourceResult) Get(ctx workflow.Context) (*kinesisanalytics.UntagResourceOutput, error) {
	var output kinesisanalytics.UntagResourceOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type KinesisanalyticsUpdateApplicationResult struct {
	Result workflow.Future
}

func (r *KinesisanalyticsUpdateApplicationResult) Get(ctx workflow.Context) (*kinesisanalytics.UpdateApplicationOutput, error) {
	var output kinesisanalytics.UpdateApplicationOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type KinesisAnalyticsStub struct {
	activities awsactivities.KinesisAnalyticsActivities
}

func NewKinesisAnalyticsStub() KinesisAnalyticsClient {
	return &KinesisAnalyticsStub{}
}

func (a *KinesisAnalyticsStub) AddApplicationCloudWatchLoggingOption(ctx workflow.Context, input *kinesisanalytics.AddApplicationCloudWatchLoggingOptionInput) (*kinesisanalytics.AddApplicationCloudWatchLoggingOptionOutput, error) {
	var output kinesisanalytics.AddApplicationCloudWatchLoggingOptionOutput
	err := workflow.ExecuteActivity(ctx, a.activities.AddApplicationCloudWatchLoggingOption, input).Get(ctx, &output)
	return &output, err
}

func (a *KinesisAnalyticsStub) AddApplicationCloudWatchLoggingOptionAsync(ctx workflow.Context, input *kinesisanalytics.AddApplicationCloudWatchLoggingOptionInput) *KinesisanalyticsAddApplicationCloudWatchLoggingOptionResult {
	future := workflow.ExecuteActivity(ctx, a.activities.AddApplicationCloudWatchLoggingOption, input)
	return &KinesisanalyticsAddApplicationCloudWatchLoggingOptionResult{Result: future}
}

func (a *KinesisAnalyticsStub) AddApplicationInput(ctx workflow.Context, input *kinesisanalytics.AddApplicationInputInput) (*kinesisanalytics.AddApplicationInputOutput, error) {
	var output kinesisanalytics.AddApplicationInputOutput
	err := workflow.ExecuteActivity(ctx, a.activities.AddApplicationInput, input).Get(ctx, &output)
	return &output, err
}

func (a *KinesisAnalyticsStub) AddApplicationInputAsync(ctx workflow.Context, input *kinesisanalytics.AddApplicationInputInput) *KinesisanalyticsAddApplicationInputResult {
	future := workflow.ExecuteActivity(ctx, a.activities.AddApplicationInput, input)
	return &KinesisanalyticsAddApplicationInputResult{Result: future}
}

func (a *KinesisAnalyticsStub) AddApplicationInputProcessingConfiguration(ctx workflow.Context, input *kinesisanalytics.AddApplicationInputProcessingConfigurationInput) (*kinesisanalytics.AddApplicationInputProcessingConfigurationOutput, error) {
	var output kinesisanalytics.AddApplicationInputProcessingConfigurationOutput
	err := workflow.ExecuteActivity(ctx, a.activities.AddApplicationInputProcessingConfiguration, input).Get(ctx, &output)
	return &output, err
}

func (a *KinesisAnalyticsStub) AddApplicationInputProcessingConfigurationAsync(ctx workflow.Context, input *kinesisanalytics.AddApplicationInputProcessingConfigurationInput) *KinesisanalyticsAddApplicationInputProcessingConfigurationResult {
	future := workflow.ExecuteActivity(ctx, a.activities.AddApplicationInputProcessingConfiguration, input)
	return &KinesisanalyticsAddApplicationInputProcessingConfigurationResult{Result: future}
}

func (a *KinesisAnalyticsStub) AddApplicationOutput(ctx workflow.Context, input *kinesisanalytics.AddApplicationOutputInput) (*kinesisanalytics.AddApplicationOutputOutput, error) {
	var output kinesisanalytics.AddApplicationOutputOutput
	err := workflow.ExecuteActivity(ctx, a.activities.AddApplicationOutput, input).Get(ctx, &output)
	return &output, err
}

func (a *KinesisAnalyticsStub) AddApplicationOutputAsync(ctx workflow.Context, input *kinesisanalytics.AddApplicationOutputInput) *KinesisanalyticsAddApplicationOutputResult {
	future := workflow.ExecuteActivity(ctx, a.activities.AddApplicationOutput, input)
	return &KinesisanalyticsAddApplicationOutputResult{Result: future}
}

func (a *KinesisAnalyticsStub) AddApplicationReferenceDataSource(ctx workflow.Context, input *kinesisanalytics.AddApplicationReferenceDataSourceInput) (*kinesisanalytics.AddApplicationReferenceDataSourceOutput, error) {
	var output kinesisanalytics.AddApplicationReferenceDataSourceOutput
	err := workflow.ExecuteActivity(ctx, a.activities.AddApplicationReferenceDataSource, input).Get(ctx, &output)
	return &output, err
}

func (a *KinesisAnalyticsStub) AddApplicationReferenceDataSourceAsync(ctx workflow.Context, input *kinesisanalytics.AddApplicationReferenceDataSourceInput) *KinesisanalyticsAddApplicationReferenceDataSourceResult {
	future := workflow.ExecuteActivity(ctx, a.activities.AddApplicationReferenceDataSource, input)
	return &KinesisanalyticsAddApplicationReferenceDataSourceResult{Result: future}
}

func (a *KinesisAnalyticsStub) CreateApplication(ctx workflow.Context, input *kinesisanalytics.CreateApplicationInput) (*kinesisanalytics.CreateApplicationOutput, error) {
	var output kinesisanalytics.CreateApplicationOutput
	err := workflow.ExecuteActivity(ctx, a.activities.CreateApplication, input).Get(ctx, &output)
	return &output, err
}

func (a *KinesisAnalyticsStub) CreateApplicationAsync(ctx workflow.Context, input *kinesisanalytics.CreateApplicationInput) *KinesisanalyticsCreateApplicationResult {
	future := workflow.ExecuteActivity(ctx, a.activities.CreateApplication, input)
	return &KinesisanalyticsCreateApplicationResult{Result: future}
}

func (a *KinesisAnalyticsStub) DeleteApplication(ctx workflow.Context, input *kinesisanalytics.DeleteApplicationInput) (*kinesisanalytics.DeleteApplicationOutput, error) {
	var output kinesisanalytics.DeleteApplicationOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeleteApplication, input).Get(ctx, &output)
	return &output, err
}

func (a *KinesisAnalyticsStub) DeleteApplicationAsync(ctx workflow.Context, input *kinesisanalytics.DeleteApplicationInput) *KinesisanalyticsDeleteApplicationResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeleteApplication, input)
	return &KinesisanalyticsDeleteApplicationResult{Result: future}
}

func (a *KinesisAnalyticsStub) DeleteApplicationCloudWatchLoggingOption(ctx workflow.Context, input *kinesisanalytics.DeleteApplicationCloudWatchLoggingOptionInput) (*kinesisanalytics.DeleteApplicationCloudWatchLoggingOptionOutput, error) {
	var output kinesisanalytics.DeleteApplicationCloudWatchLoggingOptionOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeleteApplicationCloudWatchLoggingOption, input).Get(ctx, &output)
	return &output, err
}

func (a *KinesisAnalyticsStub) DeleteApplicationCloudWatchLoggingOptionAsync(ctx workflow.Context, input *kinesisanalytics.DeleteApplicationCloudWatchLoggingOptionInput) *KinesisanalyticsDeleteApplicationCloudWatchLoggingOptionResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeleteApplicationCloudWatchLoggingOption, input)
	return &KinesisanalyticsDeleteApplicationCloudWatchLoggingOptionResult{Result: future}
}

func (a *KinesisAnalyticsStub) DeleteApplicationInputProcessingConfiguration(ctx workflow.Context, input *kinesisanalytics.DeleteApplicationInputProcessingConfigurationInput) (*kinesisanalytics.DeleteApplicationInputProcessingConfigurationOutput, error) {
	var output kinesisanalytics.DeleteApplicationInputProcessingConfigurationOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeleteApplicationInputProcessingConfiguration, input).Get(ctx, &output)
	return &output, err
}

func (a *KinesisAnalyticsStub) DeleteApplicationInputProcessingConfigurationAsync(ctx workflow.Context, input *kinesisanalytics.DeleteApplicationInputProcessingConfigurationInput) *KinesisanalyticsDeleteApplicationInputProcessingConfigurationResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeleteApplicationInputProcessingConfiguration, input)
	return &KinesisanalyticsDeleteApplicationInputProcessingConfigurationResult{Result: future}
}

func (a *KinesisAnalyticsStub) DeleteApplicationOutput(ctx workflow.Context, input *kinesisanalytics.DeleteApplicationOutputInput) (*kinesisanalytics.DeleteApplicationOutputOutput, error) {
	var output kinesisanalytics.DeleteApplicationOutputOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeleteApplicationOutput, input).Get(ctx, &output)
	return &output, err
}

func (a *KinesisAnalyticsStub) DeleteApplicationOutputAsync(ctx workflow.Context, input *kinesisanalytics.DeleteApplicationOutputInput) *KinesisanalyticsDeleteApplicationOutputResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeleteApplicationOutput, input)
	return &KinesisanalyticsDeleteApplicationOutputResult{Result: future}
}

func (a *KinesisAnalyticsStub) DeleteApplicationReferenceDataSource(ctx workflow.Context, input *kinesisanalytics.DeleteApplicationReferenceDataSourceInput) (*kinesisanalytics.DeleteApplicationReferenceDataSourceOutput, error) {
	var output kinesisanalytics.DeleteApplicationReferenceDataSourceOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeleteApplicationReferenceDataSource, input).Get(ctx, &output)
	return &output, err
}

func (a *KinesisAnalyticsStub) DeleteApplicationReferenceDataSourceAsync(ctx workflow.Context, input *kinesisanalytics.DeleteApplicationReferenceDataSourceInput) *KinesisanalyticsDeleteApplicationReferenceDataSourceResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeleteApplicationReferenceDataSource, input)
	return &KinesisanalyticsDeleteApplicationReferenceDataSourceResult{Result: future}
}

func (a *KinesisAnalyticsStub) DescribeApplication(ctx workflow.Context, input *kinesisanalytics.DescribeApplicationInput) (*kinesisanalytics.DescribeApplicationOutput, error) {
	var output kinesisanalytics.DescribeApplicationOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DescribeApplication, input).Get(ctx, &output)
	return &output, err
}

func (a *KinesisAnalyticsStub) DescribeApplicationAsync(ctx workflow.Context, input *kinesisanalytics.DescribeApplicationInput) *KinesisanalyticsDescribeApplicationResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DescribeApplication, input)
	return &KinesisanalyticsDescribeApplicationResult{Result: future}
}

func (a *KinesisAnalyticsStub) DiscoverInputSchema(ctx workflow.Context, input *kinesisanalytics.DiscoverInputSchemaInput) (*kinesisanalytics.DiscoverInputSchemaOutput, error) {
	var output kinesisanalytics.DiscoverInputSchemaOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DiscoverInputSchema, input).Get(ctx, &output)
	return &output, err
}

func (a *KinesisAnalyticsStub) DiscoverInputSchemaAsync(ctx workflow.Context, input *kinesisanalytics.DiscoverInputSchemaInput) *KinesisanalyticsDiscoverInputSchemaResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DiscoverInputSchema, input)
	return &KinesisanalyticsDiscoverInputSchemaResult{Result: future}
}

func (a *KinesisAnalyticsStub) ListApplications(ctx workflow.Context, input *kinesisanalytics.ListApplicationsInput) (*kinesisanalytics.ListApplicationsOutput, error) {
	var output kinesisanalytics.ListApplicationsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.ListApplications, input).Get(ctx, &output)
	return &output, err
}

func (a *KinesisAnalyticsStub) ListApplicationsAsync(ctx workflow.Context, input *kinesisanalytics.ListApplicationsInput) *KinesisanalyticsListApplicationsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.ListApplications, input)
	return &KinesisanalyticsListApplicationsResult{Result: future}
}

func (a *KinesisAnalyticsStub) ListTagsForResource(ctx workflow.Context, input *kinesisanalytics.ListTagsForResourceInput) (*kinesisanalytics.ListTagsForResourceOutput, error) {
	var output kinesisanalytics.ListTagsForResourceOutput
	err := workflow.ExecuteActivity(ctx, a.activities.ListTagsForResource, input).Get(ctx, &output)
	return &output, err
}

func (a *KinesisAnalyticsStub) ListTagsForResourceAsync(ctx workflow.Context, input *kinesisanalytics.ListTagsForResourceInput) *KinesisanalyticsListTagsForResourceResult {
	future := workflow.ExecuteActivity(ctx, a.activities.ListTagsForResource, input)
	return &KinesisanalyticsListTagsForResourceResult{Result: future}
}

func (a *KinesisAnalyticsStub) StartApplication(ctx workflow.Context, input *kinesisanalytics.StartApplicationInput) (*kinesisanalytics.StartApplicationOutput, error) {
	var output kinesisanalytics.StartApplicationOutput
	err := workflow.ExecuteActivity(ctx, a.activities.StartApplication, input).Get(ctx, &output)
	return &output, err
}

func (a *KinesisAnalyticsStub) StartApplicationAsync(ctx workflow.Context, input *kinesisanalytics.StartApplicationInput) *KinesisanalyticsStartApplicationResult {
	future := workflow.ExecuteActivity(ctx, a.activities.StartApplication, input)
	return &KinesisanalyticsStartApplicationResult{Result: future}
}

func (a *KinesisAnalyticsStub) StopApplication(ctx workflow.Context, input *kinesisanalytics.StopApplicationInput) (*kinesisanalytics.StopApplicationOutput, error) {
	var output kinesisanalytics.StopApplicationOutput
	err := workflow.ExecuteActivity(ctx, a.activities.StopApplication, input).Get(ctx, &output)
	return &output, err
}

func (a *KinesisAnalyticsStub) StopApplicationAsync(ctx workflow.Context, input *kinesisanalytics.StopApplicationInput) *KinesisanalyticsStopApplicationResult {
	future := workflow.ExecuteActivity(ctx, a.activities.StopApplication, input)
	return &KinesisanalyticsStopApplicationResult{Result: future}
}

func (a *KinesisAnalyticsStub) TagResource(ctx workflow.Context, input *kinesisanalytics.TagResourceInput) (*kinesisanalytics.TagResourceOutput, error) {
	var output kinesisanalytics.TagResourceOutput
	err := workflow.ExecuteActivity(ctx, a.activities.TagResource, input).Get(ctx, &output)
	return &output, err
}

func (a *KinesisAnalyticsStub) TagResourceAsync(ctx workflow.Context, input *kinesisanalytics.TagResourceInput) *KinesisanalyticsTagResourceResult {
	future := workflow.ExecuteActivity(ctx, a.activities.TagResource, input)
	return &KinesisanalyticsTagResourceResult{Result: future}
}

func (a *KinesisAnalyticsStub) UntagResource(ctx workflow.Context, input *kinesisanalytics.UntagResourceInput) (*kinesisanalytics.UntagResourceOutput, error) {
	var output kinesisanalytics.UntagResourceOutput
	err := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input).Get(ctx, &output)
	return &output, err
}

func (a *KinesisAnalyticsStub) UntagResourceAsync(ctx workflow.Context, input *kinesisanalytics.UntagResourceInput) *KinesisanalyticsUntagResourceResult {
	future := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input)
	return &KinesisanalyticsUntagResourceResult{Result: future}
}

func (a *KinesisAnalyticsStub) UpdateApplication(ctx workflow.Context, input *kinesisanalytics.UpdateApplicationInput) (*kinesisanalytics.UpdateApplicationOutput, error) {
	var output kinesisanalytics.UpdateApplicationOutput
	err := workflow.ExecuteActivity(ctx, a.activities.UpdateApplication, input).Get(ctx, &output)
	return &output, err
}

func (a *KinesisAnalyticsStub) UpdateApplicationAsync(ctx workflow.Context, input *kinesisanalytics.UpdateApplicationInput) *KinesisanalyticsUpdateApplicationResult {
	future := workflow.ExecuteActivity(ctx, a.activities.UpdateApplication, input)
	return &KinesisanalyticsUpdateApplicationResult{Result: future}
}
