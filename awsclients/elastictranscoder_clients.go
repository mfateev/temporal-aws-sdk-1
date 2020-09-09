package awsclients

import (
	"github.com/aws/aws-sdk-go/service/elastictranscoder"
	"go.temporal.io/sdk/workflow"
	"temporal.io/aws-sdk/awsactivities"
)

type ElasticTranscoderClient interface {
	CancelJob(ctx workflow.Context, input *elastictranscoder.CancelJobInput) (*elastictranscoder.CancelJobOutput, error)
	CancelJobAsync(ctx workflow.Context, input *elastictranscoder.CancelJobInput) *ElastictranscoderCancelJobResult

	CreateJob(ctx workflow.Context, input *elastictranscoder.CreateJobInput) (*elastictranscoder.CreateJobResponse, error)
	CreateJobAsync(ctx workflow.Context, input *elastictranscoder.CreateJobInput) *ElastictranscoderCreateJobResult

	CreatePipeline(ctx workflow.Context, input *elastictranscoder.CreatePipelineInput) (*elastictranscoder.CreatePipelineOutput, error)
	CreatePipelineAsync(ctx workflow.Context, input *elastictranscoder.CreatePipelineInput) *ElastictranscoderCreatePipelineResult

	CreatePreset(ctx workflow.Context, input *elastictranscoder.CreatePresetInput) (*elastictranscoder.CreatePresetOutput, error)
	CreatePresetAsync(ctx workflow.Context, input *elastictranscoder.CreatePresetInput) *ElastictranscoderCreatePresetResult

	DeletePipeline(ctx workflow.Context, input *elastictranscoder.DeletePipelineInput) (*elastictranscoder.DeletePipelineOutput, error)
	DeletePipelineAsync(ctx workflow.Context, input *elastictranscoder.DeletePipelineInput) *ElastictranscoderDeletePipelineResult

	DeletePreset(ctx workflow.Context, input *elastictranscoder.DeletePresetInput) (*elastictranscoder.DeletePresetOutput, error)
	DeletePresetAsync(ctx workflow.Context, input *elastictranscoder.DeletePresetInput) *ElastictranscoderDeletePresetResult

	ListJobsByPipeline(ctx workflow.Context, input *elastictranscoder.ListJobsByPipelineInput) (*elastictranscoder.ListJobsByPipelineOutput, error)
	ListJobsByPipelineAsync(ctx workflow.Context, input *elastictranscoder.ListJobsByPipelineInput) *ElastictranscoderListJobsByPipelineResult

	ListJobsByStatus(ctx workflow.Context, input *elastictranscoder.ListJobsByStatusInput) (*elastictranscoder.ListJobsByStatusOutput, error)
	ListJobsByStatusAsync(ctx workflow.Context, input *elastictranscoder.ListJobsByStatusInput) *ElastictranscoderListJobsByStatusResult

	ListPipelines(ctx workflow.Context, input *elastictranscoder.ListPipelinesInput) (*elastictranscoder.ListPipelinesOutput, error)
	ListPipelinesAsync(ctx workflow.Context, input *elastictranscoder.ListPipelinesInput) *ElastictranscoderListPipelinesResult

	ListPresets(ctx workflow.Context, input *elastictranscoder.ListPresetsInput) (*elastictranscoder.ListPresetsOutput, error)
	ListPresetsAsync(ctx workflow.Context, input *elastictranscoder.ListPresetsInput) *ElastictranscoderListPresetsResult

	ReadJob(ctx workflow.Context, input *elastictranscoder.ReadJobInput) (*elastictranscoder.ReadJobOutput, error)
	ReadJobAsync(ctx workflow.Context, input *elastictranscoder.ReadJobInput) *ElastictranscoderReadJobResult

	ReadPipeline(ctx workflow.Context, input *elastictranscoder.ReadPipelineInput) (*elastictranscoder.ReadPipelineOutput, error)
	ReadPipelineAsync(ctx workflow.Context, input *elastictranscoder.ReadPipelineInput) *ElastictranscoderReadPipelineResult

	ReadPreset(ctx workflow.Context, input *elastictranscoder.ReadPresetInput) (*elastictranscoder.ReadPresetOutput, error)
	ReadPresetAsync(ctx workflow.Context, input *elastictranscoder.ReadPresetInput) *ElastictranscoderReadPresetResult

	TestRole(ctx workflow.Context, input *elastictranscoder.TestRoleInput) (*elastictranscoder.TestRoleOutput, error)
	TestRoleAsync(ctx workflow.Context, input *elastictranscoder.TestRoleInput) *ElastictranscoderTestRoleResult

	UpdatePipeline(ctx workflow.Context, input *elastictranscoder.UpdatePipelineInput) (*elastictranscoder.UpdatePipelineOutput, error)
	UpdatePipelineAsync(ctx workflow.Context, input *elastictranscoder.UpdatePipelineInput) *ElastictranscoderUpdatePipelineResult

	UpdatePipelineNotifications(ctx workflow.Context, input *elastictranscoder.UpdatePipelineNotificationsInput) (*elastictranscoder.UpdatePipelineNotificationsOutput, error)
	UpdatePipelineNotificationsAsync(ctx workflow.Context, input *elastictranscoder.UpdatePipelineNotificationsInput) *ElastictranscoderUpdatePipelineNotificationsResult

	UpdatePipelineStatus(ctx workflow.Context, input *elastictranscoder.UpdatePipelineStatusInput) (*elastictranscoder.UpdatePipelineStatusOutput, error)
	UpdatePipelineStatusAsync(ctx workflow.Context, input *elastictranscoder.UpdatePipelineStatusInput) *ElastictranscoderUpdatePipelineStatusResult

	WaitUntilJobComplete(ctx workflow.Context, input *elastictranscoder.ReadJobInput) error
}

type ElastictranscoderCancelJobResult struct {
	Result workflow.Future
}

func (r *ElastictranscoderCancelJobResult) Get(ctx workflow.Context) (*elastictranscoder.CancelJobOutput, error) {
	var output elastictranscoder.CancelJobOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ElastictranscoderCreateJobResult struct {
	Result workflow.Future
}

func (r *ElastictranscoderCreateJobResult) Get(ctx workflow.Context) (*elastictranscoder.CreateJobResponse, error) {
	var output elastictranscoder.CreateJobResponse
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ElastictranscoderCreatePipelineResult struct {
	Result workflow.Future
}

func (r *ElastictranscoderCreatePipelineResult) Get(ctx workflow.Context) (*elastictranscoder.CreatePipelineOutput, error) {
	var output elastictranscoder.CreatePipelineOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ElastictranscoderCreatePresetResult struct {
	Result workflow.Future
}

func (r *ElastictranscoderCreatePresetResult) Get(ctx workflow.Context) (*elastictranscoder.CreatePresetOutput, error) {
	var output elastictranscoder.CreatePresetOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ElastictranscoderDeletePipelineResult struct {
	Result workflow.Future
}

func (r *ElastictranscoderDeletePipelineResult) Get(ctx workflow.Context) (*elastictranscoder.DeletePipelineOutput, error) {
	var output elastictranscoder.DeletePipelineOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ElastictranscoderDeletePresetResult struct {
	Result workflow.Future
}

func (r *ElastictranscoderDeletePresetResult) Get(ctx workflow.Context) (*elastictranscoder.DeletePresetOutput, error) {
	var output elastictranscoder.DeletePresetOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ElastictranscoderListJobsByPipelineResult struct {
	Result workflow.Future
}

func (r *ElastictranscoderListJobsByPipelineResult) Get(ctx workflow.Context) (*elastictranscoder.ListJobsByPipelineOutput, error) {
	var output elastictranscoder.ListJobsByPipelineOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ElastictranscoderListJobsByStatusResult struct {
	Result workflow.Future
}

func (r *ElastictranscoderListJobsByStatusResult) Get(ctx workflow.Context) (*elastictranscoder.ListJobsByStatusOutput, error) {
	var output elastictranscoder.ListJobsByStatusOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ElastictranscoderListPipelinesResult struct {
	Result workflow.Future
}

func (r *ElastictranscoderListPipelinesResult) Get(ctx workflow.Context) (*elastictranscoder.ListPipelinesOutput, error) {
	var output elastictranscoder.ListPipelinesOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ElastictranscoderListPresetsResult struct {
	Result workflow.Future
}

func (r *ElastictranscoderListPresetsResult) Get(ctx workflow.Context) (*elastictranscoder.ListPresetsOutput, error) {
	var output elastictranscoder.ListPresetsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ElastictranscoderReadJobResult struct {
	Result workflow.Future
}

func (r *ElastictranscoderReadJobResult) Get(ctx workflow.Context) (*elastictranscoder.ReadJobOutput, error) {
	var output elastictranscoder.ReadJobOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ElastictranscoderReadPipelineResult struct {
	Result workflow.Future
}

func (r *ElastictranscoderReadPipelineResult) Get(ctx workflow.Context) (*elastictranscoder.ReadPipelineOutput, error) {
	var output elastictranscoder.ReadPipelineOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ElastictranscoderReadPresetResult struct {
	Result workflow.Future
}

func (r *ElastictranscoderReadPresetResult) Get(ctx workflow.Context) (*elastictranscoder.ReadPresetOutput, error) {
	var output elastictranscoder.ReadPresetOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ElastictranscoderTestRoleResult struct {
	Result workflow.Future
}

func (r *ElastictranscoderTestRoleResult) Get(ctx workflow.Context) (*elastictranscoder.TestRoleOutput, error) {
	var output elastictranscoder.TestRoleOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ElastictranscoderUpdatePipelineResult struct {
	Result workflow.Future
}

func (r *ElastictranscoderUpdatePipelineResult) Get(ctx workflow.Context) (*elastictranscoder.UpdatePipelineOutput, error) {
	var output elastictranscoder.UpdatePipelineOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ElastictranscoderUpdatePipelineNotificationsResult struct {
	Result workflow.Future
}

func (r *ElastictranscoderUpdatePipelineNotificationsResult) Get(ctx workflow.Context) (*elastictranscoder.UpdatePipelineNotificationsOutput, error) {
	var output elastictranscoder.UpdatePipelineNotificationsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ElastictranscoderUpdatePipelineStatusResult struct {
	Result workflow.Future
}

func (r *ElastictranscoderUpdatePipelineStatusResult) Get(ctx workflow.Context) (*elastictranscoder.UpdatePipelineStatusOutput, error) {
	var output elastictranscoder.UpdatePipelineStatusOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ElasticTranscoderStub struct {
	activities awsactivities.ElasticTranscoderActivities
}

func NewElasticTranscoderStub() ElasticTranscoderClient {
	return &ElasticTranscoderStub{}
}

func (a *ElasticTranscoderStub) CancelJob(ctx workflow.Context, input *elastictranscoder.CancelJobInput) (*elastictranscoder.CancelJobOutput, error) {
	var output elastictranscoder.CancelJobOutput
	err := workflow.ExecuteActivity(ctx, a.activities.CancelJob, input).Get(ctx, &output)
	return &output, err
}

func (a *ElasticTranscoderStub) CancelJobAsync(ctx workflow.Context, input *elastictranscoder.CancelJobInput) *ElastictranscoderCancelJobResult {
	future := workflow.ExecuteActivity(ctx, a.activities.CancelJob, input)
	return &ElastictranscoderCancelJobResult{Result: future}
}

func (a *ElasticTranscoderStub) CreateJob(ctx workflow.Context, input *elastictranscoder.CreateJobInput) (*elastictranscoder.CreateJobResponse, error) {
	var output elastictranscoder.CreateJobResponse
	err := workflow.ExecuteActivity(ctx, a.activities.CreateJob, input).Get(ctx, &output)
	return &output, err
}

func (a *ElasticTranscoderStub) CreateJobAsync(ctx workflow.Context, input *elastictranscoder.CreateJobInput) *ElastictranscoderCreateJobResult {
	future := workflow.ExecuteActivity(ctx, a.activities.CreateJob, input)
	return &ElastictranscoderCreateJobResult{Result: future}
}

func (a *ElasticTranscoderStub) CreatePipeline(ctx workflow.Context, input *elastictranscoder.CreatePipelineInput) (*elastictranscoder.CreatePipelineOutput, error) {
	var output elastictranscoder.CreatePipelineOutput
	err := workflow.ExecuteActivity(ctx, a.activities.CreatePipeline, input).Get(ctx, &output)
	return &output, err
}

func (a *ElasticTranscoderStub) CreatePipelineAsync(ctx workflow.Context, input *elastictranscoder.CreatePipelineInput) *ElastictranscoderCreatePipelineResult {
	future := workflow.ExecuteActivity(ctx, a.activities.CreatePipeline, input)
	return &ElastictranscoderCreatePipelineResult{Result: future}
}

func (a *ElasticTranscoderStub) CreatePreset(ctx workflow.Context, input *elastictranscoder.CreatePresetInput) (*elastictranscoder.CreatePresetOutput, error) {
	var output elastictranscoder.CreatePresetOutput
	err := workflow.ExecuteActivity(ctx, a.activities.CreatePreset, input).Get(ctx, &output)
	return &output, err
}

func (a *ElasticTranscoderStub) CreatePresetAsync(ctx workflow.Context, input *elastictranscoder.CreatePresetInput) *ElastictranscoderCreatePresetResult {
	future := workflow.ExecuteActivity(ctx, a.activities.CreatePreset, input)
	return &ElastictranscoderCreatePresetResult{Result: future}
}

func (a *ElasticTranscoderStub) DeletePipeline(ctx workflow.Context, input *elastictranscoder.DeletePipelineInput) (*elastictranscoder.DeletePipelineOutput, error) {
	var output elastictranscoder.DeletePipelineOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeletePipeline, input).Get(ctx, &output)
	return &output, err
}

func (a *ElasticTranscoderStub) DeletePipelineAsync(ctx workflow.Context, input *elastictranscoder.DeletePipelineInput) *ElastictranscoderDeletePipelineResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeletePipeline, input)
	return &ElastictranscoderDeletePipelineResult{Result: future}
}

func (a *ElasticTranscoderStub) DeletePreset(ctx workflow.Context, input *elastictranscoder.DeletePresetInput) (*elastictranscoder.DeletePresetOutput, error) {
	var output elastictranscoder.DeletePresetOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeletePreset, input).Get(ctx, &output)
	return &output, err
}

func (a *ElasticTranscoderStub) DeletePresetAsync(ctx workflow.Context, input *elastictranscoder.DeletePresetInput) *ElastictranscoderDeletePresetResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeletePreset, input)
	return &ElastictranscoderDeletePresetResult{Result: future}
}

func (a *ElasticTranscoderStub) ListJobsByPipeline(ctx workflow.Context, input *elastictranscoder.ListJobsByPipelineInput) (*elastictranscoder.ListJobsByPipelineOutput, error) {
	var output elastictranscoder.ListJobsByPipelineOutput
	err := workflow.ExecuteActivity(ctx, a.activities.ListJobsByPipeline, input).Get(ctx, &output)
	return &output, err
}

func (a *ElasticTranscoderStub) ListJobsByPipelineAsync(ctx workflow.Context, input *elastictranscoder.ListJobsByPipelineInput) *ElastictranscoderListJobsByPipelineResult {
	future := workflow.ExecuteActivity(ctx, a.activities.ListJobsByPipeline, input)
	return &ElastictranscoderListJobsByPipelineResult{Result: future}
}

func (a *ElasticTranscoderStub) ListJobsByStatus(ctx workflow.Context, input *elastictranscoder.ListJobsByStatusInput) (*elastictranscoder.ListJobsByStatusOutput, error) {
	var output elastictranscoder.ListJobsByStatusOutput
	err := workflow.ExecuteActivity(ctx, a.activities.ListJobsByStatus, input).Get(ctx, &output)
	return &output, err
}

func (a *ElasticTranscoderStub) ListJobsByStatusAsync(ctx workflow.Context, input *elastictranscoder.ListJobsByStatusInput) *ElastictranscoderListJobsByStatusResult {
	future := workflow.ExecuteActivity(ctx, a.activities.ListJobsByStatus, input)
	return &ElastictranscoderListJobsByStatusResult{Result: future}
}

func (a *ElasticTranscoderStub) ListPipelines(ctx workflow.Context, input *elastictranscoder.ListPipelinesInput) (*elastictranscoder.ListPipelinesOutput, error) {
	var output elastictranscoder.ListPipelinesOutput
	err := workflow.ExecuteActivity(ctx, a.activities.ListPipelines, input).Get(ctx, &output)
	return &output, err
}

func (a *ElasticTranscoderStub) ListPipelinesAsync(ctx workflow.Context, input *elastictranscoder.ListPipelinesInput) *ElastictranscoderListPipelinesResult {
	future := workflow.ExecuteActivity(ctx, a.activities.ListPipelines, input)
	return &ElastictranscoderListPipelinesResult{Result: future}
}

func (a *ElasticTranscoderStub) ListPresets(ctx workflow.Context, input *elastictranscoder.ListPresetsInput) (*elastictranscoder.ListPresetsOutput, error) {
	var output elastictranscoder.ListPresetsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.ListPresets, input).Get(ctx, &output)
	return &output, err
}

func (a *ElasticTranscoderStub) ListPresetsAsync(ctx workflow.Context, input *elastictranscoder.ListPresetsInput) *ElastictranscoderListPresetsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.ListPresets, input)
	return &ElastictranscoderListPresetsResult{Result: future}
}

func (a *ElasticTranscoderStub) ReadJob(ctx workflow.Context, input *elastictranscoder.ReadJobInput) (*elastictranscoder.ReadJobOutput, error) {
	var output elastictranscoder.ReadJobOutput
	err := workflow.ExecuteActivity(ctx, a.activities.ReadJob, input).Get(ctx, &output)
	return &output, err
}

func (a *ElasticTranscoderStub) ReadJobAsync(ctx workflow.Context, input *elastictranscoder.ReadJobInput) *ElastictranscoderReadJobResult {
	future := workflow.ExecuteActivity(ctx, a.activities.ReadJob, input)
	return &ElastictranscoderReadJobResult{Result: future}
}

func (a *ElasticTranscoderStub) ReadPipeline(ctx workflow.Context, input *elastictranscoder.ReadPipelineInput) (*elastictranscoder.ReadPipelineOutput, error) {
	var output elastictranscoder.ReadPipelineOutput
	err := workflow.ExecuteActivity(ctx, a.activities.ReadPipeline, input).Get(ctx, &output)
	return &output, err
}

func (a *ElasticTranscoderStub) ReadPipelineAsync(ctx workflow.Context, input *elastictranscoder.ReadPipelineInput) *ElastictranscoderReadPipelineResult {
	future := workflow.ExecuteActivity(ctx, a.activities.ReadPipeline, input)
	return &ElastictranscoderReadPipelineResult{Result: future}
}

func (a *ElasticTranscoderStub) ReadPreset(ctx workflow.Context, input *elastictranscoder.ReadPresetInput) (*elastictranscoder.ReadPresetOutput, error) {
	var output elastictranscoder.ReadPresetOutput
	err := workflow.ExecuteActivity(ctx, a.activities.ReadPreset, input).Get(ctx, &output)
	return &output, err
}

func (a *ElasticTranscoderStub) ReadPresetAsync(ctx workflow.Context, input *elastictranscoder.ReadPresetInput) *ElastictranscoderReadPresetResult {
	future := workflow.ExecuteActivity(ctx, a.activities.ReadPreset, input)
	return &ElastictranscoderReadPresetResult{Result: future}
}

func (a *ElasticTranscoderStub) TestRole(ctx workflow.Context, input *elastictranscoder.TestRoleInput) (*elastictranscoder.TestRoleOutput, error) {
	var output elastictranscoder.TestRoleOutput
	err := workflow.ExecuteActivity(ctx, a.activities.TestRole, input).Get(ctx, &output)
	return &output, err
}

func (a *ElasticTranscoderStub) TestRoleAsync(ctx workflow.Context, input *elastictranscoder.TestRoleInput) *ElastictranscoderTestRoleResult {
	future := workflow.ExecuteActivity(ctx, a.activities.TestRole, input)
	return &ElastictranscoderTestRoleResult{Result: future}
}

func (a *ElasticTranscoderStub) UpdatePipeline(ctx workflow.Context, input *elastictranscoder.UpdatePipelineInput) (*elastictranscoder.UpdatePipelineOutput, error) {
	var output elastictranscoder.UpdatePipelineOutput
	err := workflow.ExecuteActivity(ctx, a.activities.UpdatePipeline, input).Get(ctx, &output)
	return &output, err
}

func (a *ElasticTranscoderStub) UpdatePipelineAsync(ctx workflow.Context, input *elastictranscoder.UpdatePipelineInput) *ElastictranscoderUpdatePipelineResult {
	future := workflow.ExecuteActivity(ctx, a.activities.UpdatePipeline, input)
	return &ElastictranscoderUpdatePipelineResult{Result: future}
}

func (a *ElasticTranscoderStub) UpdatePipelineNotifications(ctx workflow.Context, input *elastictranscoder.UpdatePipelineNotificationsInput) (*elastictranscoder.UpdatePipelineNotificationsOutput, error) {
	var output elastictranscoder.UpdatePipelineNotificationsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.UpdatePipelineNotifications, input).Get(ctx, &output)
	return &output, err
}

func (a *ElasticTranscoderStub) UpdatePipelineNotificationsAsync(ctx workflow.Context, input *elastictranscoder.UpdatePipelineNotificationsInput) *ElastictranscoderUpdatePipelineNotificationsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.UpdatePipelineNotifications, input)
	return &ElastictranscoderUpdatePipelineNotificationsResult{Result: future}
}

func (a *ElasticTranscoderStub) UpdatePipelineStatus(ctx workflow.Context, input *elastictranscoder.UpdatePipelineStatusInput) (*elastictranscoder.UpdatePipelineStatusOutput, error) {
	var output elastictranscoder.UpdatePipelineStatusOutput
	err := workflow.ExecuteActivity(ctx, a.activities.UpdatePipelineStatus, input).Get(ctx, &output)
	return &output, err
}

func (a *ElasticTranscoderStub) UpdatePipelineStatusAsync(ctx workflow.Context, input *elastictranscoder.UpdatePipelineStatusInput) *ElastictranscoderUpdatePipelineStatusResult {
	future := workflow.ExecuteActivity(ctx, a.activities.UpdatePipelineStatus, input)
	return &ElastictranscoderUpdatePipelineStatusResult{Result: future}
}

func (a *ElasticTranscoderStub) WaitUntilJobComplete(ctx workflow.Context, input *elastictranscoder.ReadJobInput) error {
	return workflow.ExecuteActivity(ctx, a.activities.WaitUntilJobComplete, input).Get(ctx, nil)
}

func (a *ElasticTranscoderStub) WaitUntilJobCompleteAsync(ctx workflow.Context, input *elastictranscoder.ReadJobInput) workflow.Future {
	return workflow.ExecuteActivity(ctx, a.activities.WaitUntilJobComplete, input)
}
