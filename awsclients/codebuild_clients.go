package awsclients

import (
	"github.com/aws/aws-sdk-go/service/codebuild"
	"go.temporal.io/sdk/workflow"
	"temporal.io/aws-sdk/awsactivities"
)

type CodeBuildClient interface {
	BatchDeleteBuilds(ctx workflow.Context, input *codebuild.BatchDeleteBuildsInput) (*codebuild.BatchDeleteBuildsOutput, error)
	BatchDeleteBuildsAsync(ctx workflow.Context, input *codebuild.BatchDeleteBuildsInput) *CodebuildBatchDeleteBuildsResult

	BatchGetBuildBatches(ctx workflow.Context, input *codebuild.BatchGetBuildBatchesInput) (*codebuild.BatchGetBuildBatchesOutput, error)
	BatchGetBuildBatchesAsync(ctx workflow.Context, input *codebuild.BatchGetBuildBatchesInput) *CodebuildBatchGetBuildBatchesResult

	BatchGetBuilds(ctx workflow.Context, input *codebuild.BatchGetBuildsInput) (*codebuild.BatchGetBuildsOutput, error)
	BatchGetBuildsAsync(ctx workflow.Context, input *codebuild.BatchGetBuildsInput) *CodebuildBatchGetBuildsResult

	BatchGetProjects(ctx workflow.Context, input *codebuild.BatchGetProjectsInput) (*codebuild.BatchGetProjectsOutput, error)
	BatchGetProjectsAsync(ctx workflow.Context, input *codebuild.BatchGetProjectsInput) *CodebuildBatchGetProjectsResult

	BatchGetReportGroups(ctx workflow.Context, input *codebuild.BatchGetReportGroupsInput) (*codebuild.BatchGetReportGroupsOutput, error)
	BatchGetReportGroupsAsync(ctx workflow.Context, input *codebuild.BatchGetReportGroupsInput) *CodebuildBatchGetReportGroupsResult

	BatchGetReports(ctx workflow.Context, input *codebuild.BatchGetReportsInput) (*codebuild.BatchGetReportsOutput, error)
	BatchGetReportsAsync(ctx workflow.Context, input *codebuild.BatchGetReportsInput) *CodebuildBatchGetReportsResult

	CreateProject(ctx workflow.Context, input *codebuild.CreateProjectInput) (*codebuild.CreateProjectOutput, error)
	CreateProjectAsync(ctx workflow.Context, input *codebuild.CreateProjectInput) *CodebuildCreateProjectResult

	CreateReportGroup(ctx workflow.Context, input *codebuild.CreateReportGroupInput) (*codebuild.CreateReportGroupOutput, error)
	CreateReportGroupAsync(ctx workflow.Context, input *codebuild.CreateReportGroupInput) *CodebuildCreateReportGroupResult

	CreateWebhook(ctx workflow.Context, input *codebuild.CreateWebhookInput) (*codebuild.CreateWebhookOutput, error)
	CreateWebhookAsync(ctx workflow.Context, input *codebuild.CreateWebhookInput) *CodebuildCreateWebhookResult

	DeleteBuildBatch(ctx workflow.Context, input *codebuild.DeleteBuildBatchInput) (*codebuild.DeleteBuildBatchOutput, error)
	DeleteBuildBatchAsync(ctx workflow.Context, input *codebuild.DeleteBuildBatchInput) *CodebuildDeleteBuildBatchResult

	DeleteProject(ctx workflow.Context, input *codebuild.DeleteProjectInput) (*codebuild.DeleteProjectOutput, error)
	DeleteProjectAsync(ctx workflow.Context, input *codebuild.DeleteProjectInput) *CodebuildDeleteProjectResult

	DeleteReport(ctx workflow.Context, input *codebuild.DeleteReportInput) (*codebuild.DeleteReportOutput, error)
	DeleteReportAsync(ctx workflow.Context, input *codebuild.DeleteReportInput) *CodebuildDeleteReportResult

	DeleteReportGroup(ctx workflow.Context, input *codebuild.DeleteReportGroupInput) (*codebuild.DeleteReportGroupOutput, error)
	DeleteReportGroupAsync(ctx workflow.Context, input *codebuild.DeleteReportGroupInput) *CodebuildDeleteReportGroupResult

	DeleteResourcePolicy(ctx workflow.Context, input *codebuild.DeleteResourcePolicyInput) (*codebuild.DeleteResourcePolicyOutput, error)
	DeleteResourcePolicyAsync(ctx workflow.Context, input *codebuild.DeleteResourcePolicyInput) *CodebuildDeleteResourcePolicyResult

	DeleteSourceCredentials(ctx workflow.Context, input *codebuild.DeleteSourceCredentialsInput) (*codebuild.DeleteSourceCredentialsOutput, error)
	DeleteSourceCredentialsAsync(ctx workflow.Context, input *codebuild.DeleteSourceCredentialsInput) *CodebuildDeleteSourceCredentialsResult

	DeleteWebhook(ctx workflow.Context, input *codebuild.DeleteWebhookInput) (*codebuild.DeleteWebhookOutput, error)
	DeleteWebhookAsync(ctx workflow.Context, input *codebuild.DeleteWebhookInput) *CodebuildDeleteWebhookResult

	DescribeCodeCoverages(ctx workflow.Context, input *codebuild.DescribeCodeCoveragesInput) (*codebuild.DescribeCodeCoveragesOutput, error)
	DescribeCodeCoveragesAsync(ctx workflow.Context, input *codebuild.DescribeCodeCoveragesInput) *CodebuildDescribeCodeCoveragesResult

	DescribeTestCases(ctx workflow.Context, input *codebuild.DescribeTestCasesInput) (*codebuild.DescribeTestCasesOutput, error)
	DescribeTestCasesAsync(ctx workflow.Context, input *codebuild.DescribeTestCasesInput) *CodebuildDescribeTestCasesResult

	GetResourcePolicy(ctx workflow.Context, input *codebuild.GetResourcePolicyInput) (*codebuild.GetResourcePolicyOutput, error)
	GetResourcePolicyAsync(ctx workflow.Context, input *codebuild.GetResourcePolicyInput) *CodebuildGetResourcePolicyResult

	ImportSourceCredentials(ctx workflow.Context, input *codebuild.ImportSourceCredentialsInput) (*codebuild.ImportSourceCredentialsOutput, error)
	ImportSourceCredentialsAsync(ctx workflow.Context, input *codebuild.ImportSourceCredentialsInput) *CodebuildImportSourceCredentialsResult

	InvalidateProjectCache(ctx workflow.Context, input *codebuild.InvalidateProjectCacheInput) (*codebuild.InvalidateProjectCacheOutput, error)
	InvalidateProjectCacheAsync(ctx workflow.Context, input *codebuild.InvalidateProjectCacheInput) *CodebuildInvalidateProjectCacheResult

	ListBuildBatches(ctx workflow.Context, input *codebuild.ListBuildBatchesInput) (*codebuild.ListBuildBatchesOutput, error)
	ListBuildBatchesAsync(ctx workflow.Context, input *codebuild.ListBuildBatchesInput) *CodebuildListBuildBatchesResult

	ListBuildBatchesForProject(ctx workflow.Context, input *codebuild.ListBuildBatchesForProjectInput) (*codebuild.ListBuildBatchesForProjectOutput, error)
	ListBuildBatchesForProjectAsync(ctx workflow.Context, input *codebuild.ListBuildBatchesForProjectInput) *CodebuildListBuildBatchesForProjectResult

	ListBuilds(ctx workflow.Context, input *codebuild.ListBuildsInput) (*codebuild.ListBuildsOutput, error)
	ListBuildsAsync(ctx workflow.Context, input *codebuild.ListBuildsInput) *CodebuildListBuildsResult

	ListBuildsForProject(ctx workflow.Context, input *codebuild.ListBuildsForProjectInput) (*codebuild.ListBuildsForProjectOutput, error)
	ListBuildsForProjectAsync(ctx workflow.Context, input *codebuild.ListBuildsForProjectInput) *CodebuildListBuildsForProjectResult

	ListCuratedEnvironmentImages(ctx workflow.Context, input *codebuild.ListCuratedEnvironmentImagesInput) (*codebuild.ListCuratedEnvironmentImagesOutput, error)
	ListCuratedEnvironmentImagesAsync(ctx workflow.Context, input *codebuild.ListCuratedEnvironmentImagesInput) *CodebuildListCuratedEnvironmentImagesResult

	ListProjects(ctx workflow.Context, input *codebuild.ListProjectsInput) (*codebuild.ListProjectsOutput, error)
	ListProjectsAsync(ctx workflow.Context, input *codebuild.ListProjectsInput) *CodebuildListProjectsResult

	ListReportGroups(ctx workflow.Context, input *codebuild.ListReportGroupsInput) (*codebuild.ListReportGroupsOutput, error)
	ListReportGroupsAsync(ctx workflow.Context, input *codebuild.ListReportGroupsInput) *CodebuildListReportGroupsResult

	ListReports(ctx workflow.Context, input *codebuild.ListReportsInput) (*codebuild.ListReportsOutput, error)
	ListReportsAsync(ctx workflow.Context, input *codebuild.ListReportsInput) *CodebuildListReportsResult

	ListReportsForReportGroup(ctx workflow.Context, input *codebuild.ListReportsForReportGroupInput) (*codebuild.ListReportsForReportGroupOutput, error)
	ListReportsForReportGroupAsync(ctx workflow.Context, input *codebuild.ListReportsForReportGroupInput) *CodebuildListReportsForReportGroupResult

	ListSharedProjects(ctx workflow.Context, input *codebuild.ListSharedProjectsInput) (*codebuild.ListSharedProjectsOutput, error)
	ListSharedProjectsAsync(ctx workflow.Context, input *codebuild.ListSharedProjectsInput) *CodebuildListSharedProjectsResult

	ListSharedReportGroups(ctx workflow.Context, input *codebuild.ListSharedReportGroupsInput) (*codebuild.ListSharedReportGroupsOutput, error)
	ListSharedReportGroupsAsync(ctx workflow.Context, input *codebuild.ListSharedReportGroupsInput) *CodebuildListSharedReportGroupsResult

	ListSourceCredentials(ctx workflow.Context, input *codebuild.ListSourceCredentialsInput) (*codebuild.ListSourceCredentialsOutput, error)
	ListSourceCredentialsAsync(ctx workflow.Context, input *codebuild.ListSourceCredentialsInput) *CodebuildListSourceCredentialsResult

	PutResourcePolicy(ctx workflow.Context, input *codebuild.PutResourcePolicyInput) (*codebuild.PutResourcePolicyOutput, error)
	PutResourcePolicyAsync(ctx workflow.Context, input *codebuild.PutResourcePolicyInput) *CodebuildPutResourcePolicyResult

	RetryBuild(ctx workflow.Context, input *codebuild.RetryBuildInput) (*codebuild.RetryBuildOutput, error)
	RetryBuildAsync(ctx workflow.Context, input *codebuild.RetryBuildInput) *CodebuildRetryBuildResult

	RetryBuildBatch(ctx workflow.Context, input *codebuild.RetryBuildBatchInput) (*codebuild.RetryBuildBatchOutput, error)
	RetryBuildBatchAsync(ctx workflow.Context, input *codebuild.RetryBuildBatchInput) *CodebuildRetryBuildBatchResult

	StartBuild(ctx workflow.Context, input *codebuild.StartBuildInput) (*codebuild.StartBuildOutput, error)
	StartBuildAsync(ctx workflow.Context, input *codebuild.StartBuildInput) *CodebuildStartBuildResult

	StartBuildBatch(ctx workflow.Context, input *codebuild.StartBuildBatchInput) (*codebuild.StartBuildBatchOutput, error)
	StartBuildBatchAsync(ctx workflow.Context, input *codebuild.StartBuildBatchInput) *CodebuildStartBuildBatchResult

	StopBuild(ctx workflow.Context, input *codebuild.StopBuildInput) (*codebuild.StopBuildOutput, error)
	StopBuildAsync(ctx workflow.Context, input *codebuild.StopBuildInput) *CodebuildStopBuildResult

	StopBuildBatch(ctx workflow.Context, input *codebuild.StopBuildBatchInput) (*codebuild.StopBuildBatchOutput, error)
	StopBuildBatchAsync(ctx workflow.Context, input *codebuild.StopBuildBatchInput) *CodebuildStopBuildBatchResult

	UpdateProject(ctx workflow.Context, input *codebuild.UpdateProjectInput) (*codebuild.UpdateProjectOutput, error)
	UpdateProjectAsync(ctx workflow.Context, input *codebuild.UpdateProjectInput) *CodebuildUpdateProjectResult

	UpdateReportGroup(ctx workflow.Context, input *codebuild.UpdateReportGroupInput) (*codebuild.UpdateReportGroupOutput, error)
	UpdateReportGroupAsync(ctx workflow.Context, input *codebuild.UpdateReportGroupInput) *CodebuildUpdateReportGroupResult

	UpdateWebhook(ctx workflow.Context, input *codebuild.UpdateWebhookInput) (*codebuild.UpdateWebhookOutput, error)
	UpdateWebhookAsync(ctx workflow.Context, input *codebuild.UpdateWebhookInput) *CodebuildUpdateWebhookResult
}

type CodebuildBatchDeleteBuildsResult struct {
	Result workflow.Future
}

func (r *CodebuildBatchDeleteBuildsResult) Get(ctx workflow.Context) (*codebuild.BatchDeleteBuildsOutput, error) {
	var output codebuild.BatchDeleteBuildsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type CodebuildBatchGetBuildBatchesResult struct {
	Result workflow.Future
}

func (r *CodebuildBatchGetBuildBatchesResult) Get(ctx workflow.Context) (*codebuild.BatchGetBuildBatchesOutput, error) {
	var output codebuild.BatchGetBuildBatchesOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type CodebuildBatchGetBuildsResult struct {
	Result workflow.Future
}

func (r *CodebuildBatchGetBuildsResult) Get(ctx workflow.Context) (*codebuild.BatchGetBuildsOutput, error) {
	var output codebuild.BatchGetBuildsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type CodebuildBatchGetProjectsResult struct {
	Result workflow.Future
}

func (r *CodebuildBatchGetProjectsResult) Get(ctx workflow.Context) (*codebuild.BatchGetProjectsOutput, error) {
	var output codebuild.BatchGetProjectsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type CodebuildBatchGetReportGroupsResult struct {
	Result workflow.Future
}

func (r *CodebuildBatchGetReportGroupsResult) Get(ctx workflow.Context) (*codebuild.BatchGetReportGroupsOutput, error) {
	var output codebuild.BatchGetReportGroupsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type CodebuildBatchGetReportsResult struct {
	Result workflow.Future
}

func (r *CodebuildBatchGetReportsResult) Get(ctx workflow.Context) (*codebuild.BatchGetReportsOutput, error) {
	var output codebuild.BatchGetReportsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type CodebuildCreateProjectResult struct {
	Result workflow.Future
}

func (r *CodebuildCreateProjectResult) Get(ctx workflow.Context) (*codebuild.CreateProjectOutput, error) {
	var output codebuild.CreateProjectOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type CodebuildCreateReportGroupResult struct {
	Result workflow.Future
}

func (r *CodebuildCreateReportGroupResult) Get(ctx workflow.Context) (*codebuild.CreateReportGroupOutput, error) {
	var output codebuild.CreateReportGroupOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type CodebuildCreateWebhookResult struct {
	Result workflow.Future
}

func (r *CodebuildCreateWebhookResult) Get(ctx workflow.Context) (*codebuild.CreateWebhookOutput, error) {
	var output codebuild.CreateWebhookOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type CodebuildDeleteBuildBatchResult struct {
	Result workflow.Future
}

func (r *CodebuildDeleteBuildBatchResult) Get(ctx workflow.Context) (*codebuild.DeleteBuildBatchOutput, error) {
	var output codebuild.DeleteBuildBatchOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type CodebuildDeleteProjectResult struct {
	Result workflow.Future
}

func (r *CodebuildDeleteProjectResult) Get(ctx workflow.Context) (*codebuild.DeleteProjectOutput, error) {
	var output codebuild.DeleteProjectOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type CodebuildDeleteReportResult struct {
	Result workflow.Future
}

func (r *CodebuildDeleteReportResult) Get(ctx workflow.Context) (*codebuild.DeleteReportOutput, error) {
	var output codebuild.DeleteReportOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type CodebuildDeleteReportGroupResult struct {
	Result workflow.Future
}

func (r *CodebuildDeleteReportGroupResult) Get(ctx workflow.Context) (*codebuild.DeleteReportGroupOutput, error) {
	var output codebuild.DeleteReportGroupOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type CodebuildDeleteResourcePolicyResult struct {
	Result workflow.Future
}

func (r *CodebuildDeleteResourcePolicyResult) Get(ctx workflow.Context) (*codebuild.DeleteResourcePolicyOutput, error) {
	var output codebuild.DeleteResourcePolicyOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type CodebuildDeleteSourceCredentialsResult struct {
	Result workflow.Future
}

func (r *CodebuildDeleteSourceCredentialsResult) Get(ctx workflow.Context) (*codebuild.DeleteSourceCredentialsOutput, error) {
	var output codebuild.DeleteSourceCredentialsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type CodebuildDeleteWebhookResult struct {
	Result workflow.Future
}

func (r *CodebuildDeleteWebhookResult) Get(ctx workflow.Context) (*codebuild.DeleteWebhookOutput, error) {
	var output codebuild.DeleteWebhookOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type CodebuildDescribeCodeCoveragesResult struct {
	Result workflow.Future
}

func (r *CodebuildDescribeCodeCoveragesResult) Get(ctx workflow.Context) (*codebuild.DescribeCodeCoveragesOutput, error) {
	var output codebuild.DescribeCodeCoveragesOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type CodebuildDescribeTestCasesResult struct {
	Result workflow.Future
}

func (r *CodebuildDescribeTestCasesResult) Get(ctx workflow.Context) (*codebuild.DescribeTestCasesOutput, error) {
	var output codebuild.DescribeTestCasesOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type CodebuildGetResourcePolicyResult struct {
	Result workflow.Future
}

func (r *CodebuildGetResourcePolicyResult) Get(ctx workflow.Context) (*codebuild.GetResourcePolicyOutput, error) {
	var output codebuild.GetResourcePolicyOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type CodebuildImportSourceCredentialsResult struct {
	Result workflow.Future
}

func (r *CodebuildImportSourceCredentialsResult) Get(ctx workflow.Context) (*codebuild.ImportSourceCredentialsOutput, error) {
	var output codebuild.ImportSourceCredentialsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type CodebuildInvalidateProjectCacheResult struct {
	Result workflow.Future
}

func (r *CodebuildInvalidateProjectCacheResult) Get(ctx workflow.Context) (*codebuild.InvalidateProjectCacheOutput, error) {
	var output codebuild.InvalidateProjectCacheOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type CodebuildListBuildBatchesResult struct {
	Result workflow.Future
}

func (r *CodebuildListBuildBatchesResult) Get(ctx workflow.Context) (*codebuild.ListBuildBatchesOutput, error) {
	var output codebuild.ListBuildBatchesOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type CodebuildListBuildBatchesForProjectResult struct {
	Result workflow.Future
}

func (r *CodebuildListBuildBatchesForProjectResult) Get(ctx workflow.Context) (*codebuild.ListBuildBatchesForProjectOutput, error) {
	var output codebuild.ListBuildBatchesForProjectOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type CodebuildListBuildsResult struct {
	Result workflow.Future
}

func (r *CodebuildListBuildsResult) Get(ctx workflow.Context) (*codebuild.ListBuildsOutput, error) {
	var output codebuild.ListBuildsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type CodebuildListBuildsForProjectResult struct {
	Result workflow.Future
}

func (r *CodebuildListBuildsForProjectResult) Get(ctx workflow.Context) (*codebuild.ListBuildsForProjectOutput, error) {
	var output codebuild.ListBuildsForProjectOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type CodebuildListCuratedEnvironmentImagesResult struct {
	Result workflow.Future
}

func (r *CodebuildListCuratedEnvironmentImagesResult) Get(ctx workflow.Context) (*codebuild.ListCuratedEnvironmentImagesOutput, error) {
	var output codebuild.ListCuratedEnvironmentImagesOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type CodebuildListProjectsResult struct {
	Result workflow.Future
}

func (r *CodebuildListProjectsResult) Get(ctx workflow.Context) (*codebuild.ListProjectsOutput, error) {
	var output codebuild.ListProjectsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type CodebuildListReportGroupsResult struct {
	Result workflow.Future
}

func (r *CodebuildListReportGroupsResult) Get(ctx workflow.Context) (*codebuild.ListReportGroupsOutput, error) {
	var output codebuild.ListReportGroupsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type CodebuildListReportsResult struct {
	Result workflow.Future
}

func (r *CodebuildListReportsResult) Get(ctx workflow.Context) (*codebuild.ListReportsOutput, error) {
	var output codebuild.ListReportsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type CodebuildListReportsForReportGroupResult struct {
	Result workflow.Future
}

func (r *CodebuildListReportsForReportGroupResult) Get(ctx workflow.Context) (*codebuild.ListReportsForReportGroupOutput, error) {
	var output codebuild.ListReportsForReportGroupOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type CodebuildListSharedProjectsResult struct {
	Result workflow.Future
}

func (r *CodebuildListSharedProjectsResult) Get(ctx workflow.Context) (*codebuild.ListSharedProjectsOutput, error) {
	var output codebuild.ListSharedProjectsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type CodebuildListSharedReportGroupsResult struct {
	Result workflow.Future
}

func (r *CodebuildListSharedReportGroupsResult) Get(ctx workflow.Context) (*codebuild.ListSharedReportGroupsOutput, error) {
	var output codebuild.ListSharedReportGroupsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type CodebuildListSourceCredentialsResult struct {
	Result workflow.Future
}

func (r *CodebuildListSourceCredentialsResult) Get(ctx workflow.Context) (*codebuild.ListSourceCredentialsOutput, error) {
	var output codebuild.ListSourceCredentialsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type CodebuildPutResourcePolicyResult struct {
	Result workflow.Future
}

func (r *CodebuildPutResourcePolicyResult) Get(ctx workflow.Context) (*codebuild.PutResourcePolicyOutput, error) {
	var output codebuild.PutResourcePolicyOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type CodebuildRetryBuildResult struct {
	Result workflow.Future
}

func (r *CodebuildRetryBuildResult) Get(ctx workflow.Context) (*codebuild.RetryBuildOutput, error) {
	var output codebuild.RetryBuildOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type CodebuildRetryBuildBatchResult struct {
	Result workflow.Future
}

func (r *CodebuildRetryBuildBatchResult) Get(ctx workflow.Context) (*codebuild.RetryBuildBatchOutput, error) {
	var output codebuild.RetryBuildBatchOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type CodebuildStartBuildResult struct {
	Result workflow.Future
}

func (r *CodebuildStartBuildResult) Get(ctx workflow.Context) (*codebuild.StartBuildOutput, error) {
	var output codebuild.StartBuildOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type CodebuildStartBuildBatchResult struct {
	Result workflow.Future
}

func (r *CodebuildStartBuildBatchResult) Get(ctx workflow.Context) (*codebuild.StartBuildBatchOutput, error) {
	var output codebuild.StartBuildBatchOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type CodebuildStopBuildResult struct {
	Result workflow.Future
}

func (r *CodebuildStopBuildResult) Get(ctx workflow.Context) (*codebuild.StopBuildOutput, error) {
	var output codebuild.StopBuildOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type CodebuildStopBuildBatchResult struct {
	Result workflow.Future
}

func (r *CodebuildStopBuildBatchResult) Get(ctx workflow.Context) (*codebuild.StopBuildBatchOutput, error) {
	var output codebuild.StopBuildBatchOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type CodebuildUpdateProjectResult struct {
	Result workflow.Future
}

func (r *CodebuildUpdateProjectResult) Get(ctx workflow.Context) (*codebuild.UpdateProjectOutput, error) {
	var output codebuild.UpdateProjectOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type CodebuildUpdateReportGroupResult struct {
	Result workflow.Future
}

func (r *CodebuildUpdateReportGroupResult) Get(ctx workflow.Context) (*codebuild.UpdateReportGroupOutput, error) {
	var output codebuild.UpdateReportGroupOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type CodebuildUpdateWebhookResult struct {
	Result workflow.Future
}

func (r *CodebuildUpdateWebhookResult) Get(ctx workflow.Context) (*codebuild.UpdateWebhookOutput, error) {
	var output codebuild.UpdateWebhookOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type CodeBuildStub struct {
	activities awsactivities.CodeBuildActivities
}

func NewCodeBuildStub() CodeBuildClient {
	return &CodeBuildStub{}
}

func (a *CodeBuildStub) BatchDeleteBuilds(ctx workflow.Context, input *codebuild.BatchDeleteBuildsInput) (*codebuild.BatchDeleteBuildsOutput, error) {
	var output codebuild.BatchDeleteBuildsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.BatchDeleteBuilds, input).Get(ctx, &output)
	return &output, err
}

func (a *CodeBuildStub) BatchDeleteBuildsAsync(ctx workflow.Context, input *codebuild.BatchDeleteBuildsInput) *CodebuildBatchDeleteBuildsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.BatchDeleteBuilds, input)
	return &CodebuildBatchDeleteBuildsResult{Result: future}
}

func (a *CodeBuildStub) BatchGetBuildBatches(ctx workflow.Context, input *codebuild.BatchGetBuildBatchesInput) (*codebuild.BatchGetBuildBatchesOutput, error) {
	var output codebuild.BatchGetBuildBatchesOutput
	err := workflow.ExecuteActivity(ctx, a.activities.BatchGetBuildBatches, input).Get(ctx, &output)
	return &output, err
}

func (a *CodeBuildStub) BatchGetBuildBatchesAsync(ctx workflow.Context, input *codebuild.BatchGetBuildBatchesInput) *CodebuildBatchGetBuildBatchesResult {
	future := workflow.ExecuteActivity(ctx, a.activities.BatchGetBuildBatches, input)
	return &CodebuildBatchGetBuildBatchesResult{Result: future}
}

func (a *CodeBuildStub) BatchGetBuilds(ctx workflow.Context, input *codebuild.BatchGetBuildsInput) (*codebuild.BatchGetBuildsOutput, error) {
	var output codebuild.BatchGetBuildsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.BatchGetBuilds, input).Get(ctx, &output)
	return &output, err
}

func (a *CodeBuildStub) BatchGetBuildsAsync(ctx workflow.Context, input *codebuild.BatchGetBuildsInput) *CodebuildBatchGetBuildsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.BatchGetBuilds, input)
	return &CodebuildBatchGetBuildsResult{Result: future}
}

func (a *CodeBuildStub) BatchGetProjects(ctx workflow.Context, input *codebuild.BatchGetProjectsInput) (*codebuild.BatchGetProjectsOutput, error) {
	var output codebuild.BatchGetProjectsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.BatchGetProjects, input).Get(ctx, &output)
	return &output, err
}

func (a *CodeBuildStub) BatchGetProjectsAsync(ctx workflow.Context, input *codebuild.BatchGetProjectsInput) *CodebuildBatchGetProjectsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.BatchGetProjects, input)
	return &CodebuildBatchGetProjectsResult{Result: future}
}

func (a *CodeBuildStub) BatchGetReportGroups(ctx workflow.Context, input *codebuild.BatchGetReportGroupsInput) (*codebuild.BatchGetReportGroupsOutput, error) {
	var output codebuild.BatchGetReportGroupsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.BatchGetReportGroups, input).Get(ctx, &output)
	return &output, err
}

func (a *CodeBuildStub) BatchGetReportGroupsAsync(ctx workflow.Context, input *codebuild.BatchGetReportGroupsInput) *CodebuildBatchGetReportGroupsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.BatchGetReportGroups, input)
	return &CodebuildBatchGetReportGroupsResult{Result: future}
}

func (a *CodeBuildStub) BatchGetReports(ctx workflow.Context, input *codebuild.BatchGetReportsInput) (*codebuild.BatchGetReportsOutput, error) {
	var output codebuild.BatchGetReportsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.BatchGetReports, input).Get(ctx, &output)
	return &output, err
}

func (a *CodeBuildStub) BatchGetReportsAsync(ctx workflow.Context, input *codebuild.BatchGetReportsInput) *CodebuildBatchGetReportsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.BatchGetReports, input)
	return &CodebuildBatchGetReportsResult{Result: future}
}

func (a *CodeBuildStub) CreateProject(ctx workflow.Context, input *codebuild.CreateProjectInput) (*codebuild.CreateProjectOutput, error) {
	var output codebuild.CreateProjectOutput
	err := workflow.ExecuteActivity(ctx, a.activities.CreateProject, input).Get(ctx, &output)
	return &output, err
}

func (a *CodeBuildStub) CreateProjectAsync(ctx workflow.Context, input *codebuild.CreateProjectInput) *CodebuildCreateProjectResult {
	future := workflow.ExecuteActivity(ctx, a.activities.CreateProject, input)
	return &CodebuildCreateProjectResult{Result: future}
}

func (a *CodeBuildStub) CreateReportGroup(ctx workflow.Context, input *codebuild.CreateReportGroupInput) (*codebuild.CreateReportGroupOutput, error) {
	var output codebuild.CreateReportGroupOutput
	err := workflow.ExecuteActivity(ctx, a.activities.CreateReportGroup, input).Get(ctx, &output)
	return &output, err
}

func (a *CodeBuildStub) CreateReportGroupAsync(ctx workflow.Context, input *codebuild.CreateReportGroupInput) *CodebuildCreateReportGroupResult {
	future := workflow.ExecuteActivity(ctx, a.activities.CreateReportGroup, input)
	return &CodebuildCreateReportGroupResult{Result: future}
}

func (a *CodeBuildStub) CreateWebhook(ctx workflow.Context, input *codebuild.CreateWebhookInput) (*codebuild.CreateWebhookOutput, error) {
	var output codebuild.CreateWebhookOutput
	err := workflow.ExecuteActivity(ctx, a.activities.CreateWebhook, input).Get(ctx, &output)
	return &output, err
}

func (a *CodeBuildStub) CreateWebhookAsync(ctx workflow.Context, input *codebuild.CreateWebhookInput) *CodebuildCreateWebhookResult {
	future := workflow.ExecuteActivity(ctx, a.activities.CreateWebhook, input)
	return &CodebuildCreateWebhookResult{Result: future}
}

func (a *CodeBuildStub) DeleteBuildBatch(ctx workflow.Context, input *codebuild.DeleteBuildBatchInput) (*codebuild.DeleteBuildBatchOutput, error) {
	var output codebuild.DeleteBuildBatchOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeleteBuildBatch, input).Get(ctx, &output)
	return &output, err
}

func (a *CodeBuildStub) DeleteBuildBatchAsync(ctx workflow.Context, input *codebuild.DeleteBuildBatchInput) *CodebuildDeleteBuildBatchResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeleteBuildBatch, input)
	return &CodebuildDeleteBuildBatchResult{Result: future}
}

func (a *CodeBuildStub) DeleteProject(ctx workflow.Context, input *codebuild.DeleteProjectInput) (*codebuild.DeleteProjectOutput, error) {
	var output codebuild.DeleteProjectOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeleteProject, input).Get(ctx, &output)
	return &output, err
}

func (a *CodeBuildStub) DeleteProjectAsync(ctx workflow.Context, input *codebuild.DeleteProjectInput) *CodebuildDeleteProjectResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeleteProject, input)
	return &CodebuildDeleteProjectResult{Result: future}
}

func (a *CodeBuildStub) DeleteReport(ctx workflow.Context, input *codebuild.DeleteReportInput) (*codebuild.DeleteReportOutput, error) {
	var output codebuild.DeleteReportOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeleteReport, input).Get(ctx, &output)
	return &output, err
}

func (a *CodeBuildStub) DeleteReportAsync(ctx workflow.Context, input *codebuild.DeleteReportInput) *CodebuildDeleteReportResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeleteReport, input)
	return &CodebuildDeleteReportResult{Result: future}
}

func (a *CodeBuildStub) DeleteReportGroup(ctx workflow.Context, input *codebuild.DeleteReportGroupInput) (*codebuild.DeleteReportGroupOutput, error) {
	var output codebuild.DeleteReportGroupOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeleteReportGroup, input).Get(ctx, &output)
	return &output, err
}

func (a *CodeBuildStub) DeleteReportGroupAsync(ctx workflow.Context, input *codebuild.DeleteReportGroupInput) *CodebuildDeleteReportGroupResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeleteReportGroup, input)
	return &CodebuildDeleteReportGroupResult{Result: future}
}

func (a *CodeBuildStub) DeleteResourcePolicy(ctx workflow.Context, input *codebuild.DeleteResourcePolicyInput) (*codebuild.DeleteResourcePolicyOutput, error) {
	var output codebuild.DeleteResourcePolicyOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeleteResourcePolicy, input).Get(ctx, &output)
	return &output, err
}

func (a *CodeBuildStub) DeleteResourcePolicyAsync(ctx workflow.Context, input *codebuild.DeleteResourcePolicyInput) *CodebuildDeleteResourcePolicyResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeleteResourcePolicy, input)
	return &CodebuildDeleteResourcePolicyResult{Result: future}
}

func (a *CodeBuildStub) DeleteSourceCredentials(ctx workflow.Context, input *codebuild.DeleteSourceCredentialsInput) (*codebuild.DeleteSourceCredentialsOutput, error) {
	var output codebuild.DeleteSourceCredentialsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeleteSourceCredentials, input).Get(ctx, &output)
	return &output, err
}

func (a *CodeBuildStub) DeleteSourceCredentialsAsync(ctx workflow.Context, input *codebuild.DeleteSourceCredentialsInput) *CodebuildDeleteSourceCredentialsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeleteSourceCredentials, input)
	return &CodebuildDeleteSourceCredentialsResult{Result: future}
}

func (a *CodeBuildStub) DeleteWebhook(ctx workflow.Context, input *codebuild.DeleteWebhookInput) (*codebuild.DeleteWebhookOutput, error) {
	var output codebuild.DeleteWebhookOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeleteWebhook, input).Get(ctx, &output)
	return &output, err
}

func (a *CodeBuildStub) DeleteWebhookAsync(ctx workflow.Context, input *codebuild.DeleteWebhookInput) *CodebuildDeleteWebhookResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeleteWebhook, input)
	return &CodebuildDeleteWebhookResult{Result: future}
}

func (a *CodeBuildStub) DescribeCodeCoverages(ctx workflow.Context, input *codebuild.DescribeCodeCoveragesInput) (*codebuild.DescribeCodeCoveragesOutput, error) {
	var output codebuild.DescribeCodeCoveragesOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DescribeCodeCoverages, input).Get(ctx, &output)
	return &output, err
}

func (a *CodeBuildStub) DescribeCodeCoveragesAsync(ctx workflow.Context, input *codebuild.DescribeCodeCoveragesInput) *CodebuildDescribeCodeCoveragesResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DescribeCodeCoverages, input)
	return &CodebuildDescribeCodeCoveragesResult{Result: future}
}

func (a *CodeBuildStub) DescribeTestCases(ctx workflow.Context, input *codebuild.DescribeTestCasesInput) (*codebuild.DescribeTestCasesOutput, error) {
	var output codebuild.DescribeTestCasesOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DescribeTestCases, input).Get(ctx, &output)
	return &output, err
}

func (a *CodeBuildStub) DescribeTestCasesAsync(ctx workflow.Context, input *codebuild.DescribeTestCasesInput) *CodebuildDescribeTestCasesResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DescribeTestCases, input)
	return &CodebuildDescribeTestCasesResult{Result: future}
}

func (a *CodeBuildStub) GetResourcePolicy(ctx workflow.Context, input *codebuild.GetResourcePolicyInput) (*codebuild.GetResourcePolicyOutput, error) {
	var output codebuild.GetResourcePolicyOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetResourcePolicy, input).Get(ctx, &output)
	return &output, err
}

func (a *CodeBuildStub) GetResourcePolicyAsync(ctx workflow.Context, input *codebuild.GetResourcePolicyInput) *CodebuildGetResourcePolicyResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetResourcePolicy, input)
	return &CodebuildGetResourcePolicyResult{Result: future}
}

func (a *CodeBuildStub) ImportSourceCredentials(ctx workflow.Context, input *codebuild.ImportSourceCredentialsInput) (*codebuild.ImportSourceCredentialsOutput, error) {
	var output codebuild.ImportSourceCredentialsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.ImportSourceCredentials, input).Get(ctx, &output)
	return &output, err
}

func (a *CodeBuildStub) ImportSourceCredentialsAsync(ctx workflow.Context, input *codebuild.ImportSourceCredentialsInput) *CodebuildImportSourceCredentialsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.ImportSourceCredentials, input)
	return &CodebuildImportSourceCredentialsResult{Result: future}
}

func (a *CodeBuildStub) InvalidateProjectCache(ctx workflow.Context, input *codebuild.InvalidateProjectCacheInput) (*codebuild.InvalidateProjectCacheOutput, error) {
	var output codebuild.InvalidateProjectCacheOutput
	err := workflow.ExecuteActivity(ctx, a.activities.InvalidateProjectCache, input).Get(ctx, &output)
	return &output, err
}

func (a *CodeBuildStub) InvalidateProjectCacheAsync(ctx workflow.Context, input *codebuild.InvalidateProjectCacheInput) *CodebuildInvalidateProjectCacheResult {
	future := workflow.ExecuteActivity(ctx, a.activities.InvalidateProjectCache, input)
	return &CodebuildInvalidateProjectCacheResult{Result: future}
}

func (a *CodeBuildStub) ListBuildBatches(ctx workflow.Context, input *codebuild.ListBuildBatchesInput) (*codebuild.ListBuildBatchesOutput, error) {
	var output codebuild.ListBuildBatchesOutput
	err := workflow.ExecuteActivity(ctx, a.activities.ListBuildBatches, input).Get(ctx, &output)
	return &output, err
}

func (a *CodeBuildStub) ListBuildBatchesAsync(ctx workflow.Context, input *codebuild.ListBuildBatchesInput) *CodebuildListBuildBatchesResult {
	future := workflow.ExecuteActivity(ctx, a.activities.ListBuildBatches, input)
	return &CodebuildListBuildBatchesResult{Result: future}
}

func (a *CodeBuildStub) ListBuildBatchesForProject(ctx workflow.Context, input *codebuild.ListBuildBatchesForProjectInput) (*codebuild.ListBuildBatchesForProjectOutput, error) {
	var output codebuild.ListBuildBatchesForProjectOutput
	err := workflow.ExecuteActivity(ctx, a.activities.ListBuildBatchesForProject, input).Get(ctx, &output)
	return &output, err
}

func (a *CodeBuildStub) ListBuildBatchesForProjectAsync(ctx workflow.Context, input *codebuild.ListBuildBatchesForProjectInput) *CodebuildListBuildBatchesForProjectResult {
	future := workflow.ExecuteActivity(ctx, a.activities.ListBuildBatchesForProject, input)
	return &CodebuildListBuildBatchesForProjectResult{Result: future}
}

func (a *CodeBuildStub) ListBuilds(ctx workflow.Context, input *codebuild.ListBuildsInput) (*codebuild.ListBuildsOutput, error) {
	var output codebuild.ListBuildsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.ListBuilds, input).Get(ctx, &output)
	return &output, err
}

func (a *CodeBuildStub) ListBuildsAsync(ctx workflow.Context, input *codebuild.ListBuildsInput) *CodebuildListBuildsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.ListBuilds, input)
	return &CodebuildListBuildsResult{Result: future}
}

func (a *CodeBuildStub) ListBuildsForProject(ctx workflow.Context, input *codebuild.ListBuildsForProjectInput) (*codebuild.ListBuildsForProjectOutput, error) {
	var output codebuild.ListBuildsForProjectOutput
	err := workflow.ExecuteActivity(ctx, a.activities.ListBuildsForProject, input).Get(ctx, &output)
	return &output, err
}

func (a *CodeBuildStub) ListBuildsForProjectAsync(ctx workflow.Context, input *codebuild.ListBuildsForProjectInput) *CodebuildListBuildsForProjectResult {
	future := workflow.ExecuteActivity(ctx, a.activities.ListBuildsForProject, input)
	return &CodebuildListBuildsForProjectResult{Result: future}
}

func (a *CodeBuildStub) ListCuratedEnvironmentImages(ctx workflow.Context, input *codebuild.ListCuratedEnvironmentImagesInput) (*codebuild.ListCuratedEnvironmentImagesOutput, error) {
	var output codebuild.ListCuratedEnvironmentImagesOutput
	err := workflow.ExecuteActivity(ctx, a.activities.ListCuratedEnvironmentImages, input).Get(ctx, &output)
	return &output, err
}

func (a *CodeBuildStub) ListCuratedEnvironmentImagesAsync(ctx workflow.Context, input *codebuild.ListCuratedEnvironmentImagesInput) *CodebuildListCuratedEnvironmentImagesResult {
	future := workflow.ExecuteActivity(ctx, a.activities.ListCuratedEnvironmentImages, input)
	return &CodebuildListCuratedEnvironmentImagesResult{Result: future}
}

func (a *CodeBuildStub) ListProjects(ctx workflow.Context, input *codebuild.ListProjectsInput) (*codebuild.ListProjectsOutput, error) {
	var output codebuild.ListProjectsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.ListProjects, input).Get(ctx, &output)
	return &output, err
}

func (a *CodeBuildStub) ListProjectsAsync(ctx workflow.Context, input *codebuild.ListProjectsInput) *CodebuildListProjectsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.ListProjects, input)
	return &CodebuildListProjectsResult{Result: future}
}

func (a *CodeBuildStub) ListReportGroups(ctx workflow.Context, input *codebuild.ListReportGroupsInput) (*codebuild.ListReportGroupsOutput, error) {
	var output codebuild.ListReportGroupsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.ListReportGroups, input).Get(ctx, &output)
	return &output, err
}

func (a *CodeBuildStub) ListReportGroupsAsync(ctx workflow.Context, input *codebuild.ListReportGroupsInput) *CodebuildListReportGroupsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.ListReportGroups, input)
	return &CodebuildListReportGroupsResult{Result: future}
}

func (a *CodeBuildStub) ListReports(ctx workflow.Context, input *codebuild.ListReportsInput) (*codebuild.ListReportsOutput, error) {
	var output codebuild.ListReportsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.ListReports, input).Get(ctx, &output)
	return &output, err
}

func (a *CodeBuildStub) ListReportsAsync(ctx workflow.Context, input *codebuild.ListReportsInput) *CodebuildListReportsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.ListReports, input)
	return &CodebuildListReportsResult{Result: future}
}

func (a *CodeBuildStub) ListReportsForReportGroup(ctx workflow.Context, input *codebuild.ListReportsForReportGroupInput) (*codebuild.ListReportsForReportGroupOutput, error) {
	var output codebuild.ListReportsForReportGroupOutput
	err := workflow.ExecuteActivity(ctx, a.activities.ListReportsForReportGroup, input).Get(ctx, &output)
	return &output, err
}

func (a *CodeBuildStub) ListReportsForReportGroupAsync(ctx workflow.Context, input *codebuild.ListReportsForReportGroupInput) *CodebuildListReportsForReportGroupResult {
	future := workflow.ExecuteActivity(ctx, a.activities.ListReportsForReportGroup, input)
	return &CodebuildListReportsForReportGroupResult{Result: future}
}

func (a *CodeBuildStub) ListSharedProjects(ctx workflow.Context, input *codebuild.ListSharedProjectsInput) (*codebuild.ListSharedProjectsOutput, error) {
	var output codebuild.ListSharedProjectsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.ListSharedProjects, input).Get(ctx, &output)
	return &output, err
}

func (a *CodeBuildStub) ListSharedProjectsAsync(ctx workflow.Context, input *codebuild.ListSharedProjectsInput) *CodebuildListSharedProjectsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.ListSharedProjects, input)
	return &CodebuildListSharedProjectsResult{Result: future}
}

func (a *CodeBuildStub) ListSharedReportGroups(ctx workflow.Context, input *codebuild.ListSharedReportGroupsInput) (*codebuild.ListSharedReportGroupsOutput, error) {
	var output codebuild.ListSharedReportGroupsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.ListSharedReportGroups, input).Get(ctx, &output)
	return &output, err
}

func (a *CodeBuildStub) ListSharedReportGroupsAsync(ctx workflow.Context, input *codebuild.ListSharedReportGroupsInput) *CodebuildListSharedReportGroupsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.ListSharedReportGroups, input)
	return &CodebuildListSharedReportGroupsResult{Result: future}
}

func (a *CodeBuildStub) ListSourceCredentials(ctx workflow.Context, input *codebuild.ListSourceCredentialsInput) (*codebuild.ListSourceCredentialsOutput, error) {
	var output codebuild.ListSourceCredentialsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.ListSourceCredentials, input).Get(ctx, &output)
	return &output, err
}

func (a *CodeBuildStub) ListSourceCredentialsAsync(ctx workflow.Context, input *codebuild.ListSourceCredentialsInput) *CodebuildListSourceCredentialsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.ListSourceCredentials, input)
	return &CodebuildListSourceCredentialsResult{Result: future}
}

func (a *CodeBuildStub) PutResourcePolicy(ctx workflow.Context, input *codebuild.PutResourcePolicyInput) (*codebuild.PutResourcePolicyOutput, error) {
	var output codebuild.PutResourcePolicyOutput
	err := workflow.ExecuteActivity(ctx, a.activities.PutResourcePolicy, input).Get(ctx, &output)
	return &output, err
}

func (a *CodeBuildStub) PutResourcePolicyAsync(ctx workflow.Context, input *codebuild.PutResourcePolicyInput) *CodebuildPutResourcePolicyResult {
	future := workflow.ExecuteActivity(ctx, a.activities.PutResourcePolicy, input)
	return &CodebuildPutResourcePolicyResult{Result: future}
}

func (a *CodeBuildStub) RetryBuild(ctx workflow.Context, input *codebuild.RetryBuildInput) (*codebuild.RetryBuildOutput, error) {
	var output codebuild.RetryBuildOutput
	err := workflow.ExecuteActivity(ctx, a.activities.RetryBuild, input).Get(ctx, &output)
	return &output, err
}

func (a *CodeBuildStub) RetryBuildAsync(ctx workflow.Context, input *codebuild.RetryBuildInput) *CodebuildRetryBuildResult {
	future := workflow.ExecuteActivity(ctx, a.activities.RetryBuild, input)
	return &CodebuildRetryBuildResult{Result: future}
}

func (a *CodeBuildStub) RetryBuildBatch(ctx workflow.Context, input *codebuild.RetryBuildBatchInput) (*codebuild.RetryBuildBatchOutput, error) {
	var output codebuild.RetryBuildBatchOutput
	err := workflow.ExecuteActivity(ctx, a.activities.RetryBuildBatch, input).Get(ctx, &output)
	return &output, err
}

func (a *CodeBuildStub) RetryBuildBatchAsync(ctx workflow.Context, input *codebuild.RetryBuildBatchInput) *CodebuildRetryBuildBatchResult {
	future := workflow.ExecuteActivity(ctx, a.activities.RetryBuildBatch, input)
	return &CodebuildRetryBuildBatchResult{Result: future}
}

func (a *CodeBuildStub) StartBuild(ctx workflow.Context, input *codebuild.StartBuildInput) (*codebuild.StartBuildOutput, error) {
	var output codebuild.StartBuildOutput
	err := workflow.ExecuteActivity(ctx, a.activities.StartBuild, input).Get(ctx, &output)
	return &output, err
}

func (a *CodeBuildStub) StartBuildAsync(ctx workflow.Context, input *codebuild.StartBuildInput) *CodebuildStartBuildResult {
	future := workflow.ExecuteActivity(ctx, a.activities.StartBuild, input)
	return &CodebuildStartBuildResult{Result: future}
}

func (a *CodeBuildStub) StartBuildBatch(ctx workflow.Context, input *codebuild.StartBuildBatchInput) (*codebuild.StartBuildBatchOutput, error) {
	var output codebuild.StartBuildBatchOutput
	err := workflow.ExecuteActivity(ctx, a.activities.StartBuildBatch, input).Get(ctx, &output)
	return &output, err
}

func (a *CodeBuildStub) StartBuildBatchAsync(ctx workflow.Context, input *codebuild.StartBuildBatchInput) *CodebuildStartBuildBatchResult {
	future := workflow.ExecuteActivity(ctx, a.activities.StartBuildBatch, input)
	return &CodebuildStartBuildBatchResult{Result: future}
}

func (a *CodeBuildStub) StopBuild(ctx workflow.Context, input *codebuild.StopBuildInput) (*codebuild.StopBuildOutput, error) {
	var output codebuild.StopBuildOutput
	err := workflow.ExecuteActivity(ctx, a.activities.StopBuild, input).Get(ctx, &output)
	return &output, err
}

func (a *CodeBuildStub) StopBuildAsync(ctx workflow.Context, input *codebuild.StopBuildInput) *CodebuildStopBuildResult {
	future := workflow.ExecuteActivity(ctx, a.activities.StopBuild, input)
	return &CodebuildStopBuildResult{Result: future}
}

func (a *CodeBuildStub) StopBuildBatch(ctx workflow.Context, input *codebuild.StopBuildBatchInput) (*codebuild.StopBuildBatchOutput, error) {
	var output codebuild.StopBuildBatchOutput
	err := workflow.ExecuteActivity(ctx, a.activities.StopBuildBatch, input).Get(ctx, &output)
	return &output, err
}

func (a *CodeBuildStub) StopBuildBatchAsync(ctx workflow.Context, input *codebuild.StopBuildBatchInput) *CodebuildStopBuildBatchResult {
	future := workflow.ExecuteActivity(ctx, a.activities.StopBuildBatch, input)
	return &CodebuildStopBuildBatchResult{Result: future}
}

func (a *CodeBuildStub) UpdateProject(ctx workflow.Context, input *codebuild.UpdateProjectInput) (*codebuild.UpdateProjectOutput, error) {
	var output codebuild.UpdateProjectOutput
	err := workflow.ExecuteActivity(ctx, a.activities.UpdateProject, input).Get(ctx, &output)
	return &output, err
}

func (a *CodeBuildStub) UpdateProjectAsync(ctx workflow.Context, input *codebuild.UpdateProjectInput) *CodebuildUpdateProjectResult {
	future := workflow.ExecuteActivity(ctx, a.activities.UpdateProject, input)
	return &CodebuildUpdateProjectResult{Result: future}
}

func (a *CodeBuildStub) UpdateReportGroup(ctx workflow.Context, input *codebuild.UpdateReportGroupInput) (*codebuild.UpdateReportGroupOutput, error) {
	var output codebuild.UpdateReportGroupOutput
	err := workflow.ExecuteActivity(ctx, a.activities.UpdateReportGroup, input).Get(ctx, &output)
	return &output, err
}

func (a *CodeBuildStub) UpdateReportGroupAsync(ctx workflow.Context, input *codebuild.UpdateReportGroupInput) *CodebuildUpdateReportGroupResult {
	future := workflow.ExecuteActivity(ctx, a.activities.UpdateReportGroup, input)
	return &CodebuildUpdateReportGroupResult{Result: future}
}

func (a *CodeBuildStub) UpdateWebhook(ctx workflow.Context, input *codebuild.UpdateWebhookInput) (*codebuild.UpdateWebhookOutput, error) {
	var output codebuild.UpdateWebhookOutput
	err := workflow.ExecuteActivity(ctx, a.activities.UpdateWebhook, input).Get(ctx, &output)
	return &output, err
}

func (a *CodeBuildStub) UpdateWebhookAsync(ctx workflow.Context, input *codebuild.UpdateWebhookInput) *CodebuildUpdateWebhookResult {
	future := workflow.ExecuteActivity(ctx, a.activities.UpdateWebhook, input)
	return &CodebuildUpdateWebhookResult{Result: future}
}
