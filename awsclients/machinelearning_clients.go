package awsclients

import (
	"github.com/aws/aws-sdk-go/service/machinelearning"
	"go.temporal.io/sdk/workflow"
	"temporal.io/aws-sdk/awsactivities"
)

type MachineLearningClient interface {
	AddTags(ctx workflow.Context, input *machinelearning.AddTagsInput) (*machinelearning.AddTagsOutput, error)
	AddTagsAsync(ctx workflow.Context, input *machinelearning.AddTagsInput) *MachinelearningAddTagsResult

	CreateBatchPrediction(ctx workflow.Context, input *machinelearning.CreateBatchPredictionInput) (*machinelearning.CreateBatchPredictionOutput, error)
	CreateBatchPredictionAsync(ctx workflow.Context, input *machinelearning.CreateBatchPredictionInput) *MachinelearningCreateBatchPredictionResult

	CreateDataSourceFromRDS(ctx workflow.Context, input *machinelearning.CreateDataSourceFromRDSInput) (*machinelearning.CreateDataSourceFromRDSOutput, error)
	CreateDataSourceFromRDSAsync(ctx workflow.Context, input *machinelearning.CreateDataSourceFromRDSInput) *MachinelearningCreateDataSourceFromRDSResult

	CreateDataSourceFromRedshift(ctx workflow.Context, input *machinelearning.CreateDataSourceFromRedshiftInput) (*machinelearning.CreateDataSourceFromRedshiftOutput, error)
	CreateDataSourceFromRedshiftAsync(ctx workflow.Context, input *machinelearning.CreateDataSourceFromRedshiftInput) *MachinelearningCreateDataSourceFromRedshiftResult

	CreateDataSourceFromS3(ctx workflow.Context, input *machinelearning.CreateDataSourceFromS3Input) (*machinelearning.CreateDataSourceFromS3Output, error)
	CreateDataSourceFromS3Async(ctx workflow.Context, input *machinelearning.CreateDataSourceFromS3Input) *MachinelearningCreateDataSourceFromS3Result

	CreateEvaluation(ctx workflow.Context, input *machinelearning.CreateEvaluationInput) (*machinelearning.CreateEvaluationOutput, error)
	CreateEvaluationAsync(ctx workflow.Context, input *machinelearning.CreateEvaluationInput) *MachinelearningCreateEvaluationResult

	CreateMLModel(ctx workflow.Context, input *machinelearning.CreateMLModelInput) (*machinelearning.CreateMLModelOutput, error)
	CreateMLModelAsync(ctx workflow.Context, input *machinelearning.CreateMLModelInput) *MachinelearningCreateMLModelResult

	CreateRealtimeEndpoint(ctx workflow.Context, input *machinelearning.CreateRealtimeEndpointInput) (*machinelearning.CreateRealtimeEndpointOutput, error)
	CreateRealtimeEndpointAsync(ctx workflow.Context, input *machinelearning.CreateRealtimeEndpointInput) *MachinelearningCreateRealtimeEndpointResult

	DeleteBatchPrediction(ctx workflow.Context, input *machinelearning.DeleteBatchPredictionInput) (*machinelearning.DeleteBatchPredictionOutput, error)
	DeleteBatchPredictionAsync(ctx workflow.Context, input *machinelearning.DeleteBatchPredictionInput) *MachinelearningDeleteBatchPredictionResult

	DeleteDataSource(ctx workflow.Context, input *machinelearning.DeleteDataSourceInput) (*machinelearning.DeleteDataSourceOutput, error)
	DeleteDataSourceAsync(ctx workflow.Context, input *machinelearning.DeleteDataSourceInput) *MachinelearningDeleteDataSourceResult

	DeleteEvaluation(ctx workflow.Context, input *machinelearning.DeleteEvaluationInput) (*machinelearning.DeleteEvaluationOutput, error)
	DeleteEvaluationAsync(ctx workflow.Context, input *machinelearning.DeleteEvaluationInput) *MachinelearningDeleteEvaluationResult

	DeleteMLModel(ctx workflow.Context, input *machinelearning.DeleteMLModelInput) (*machinelearning.DeleteMLModelOutput, error)
	DeleteMLModelAsync(ctx workflow.Context, input *machinelearning.DeleteMLModelInput) *MachinelearningDeleteMLModelResult

	DeleteRealtimeEndpoint(ctx workflow.Context, input *machinelearning.DeleteRealtimeEndpointInput) (*machinelearning.DeleteRealtimeEndpointOutput, error)
	DeleteRealtimeEndpointAsync(ctx workflow.Context, input *machinelearning.DeleteRealtimeEndpointInput) *MachinelearningDeleteRealtimeEndpointResult

	DeleteTags(ctx workflow.Context, input *machinelearning.DeleteTagsInput) (*machinelearning.DeleteTagsOutput, error)
	DeleteTagsAsync(ctx workflow.Context, input *machinelearning.DeleteTagsInput) *MachinelearningDeleteTagsResult

	DescribeBatchPredictions(ctx workflow.Context, input *machinelearning.DescribeBatchPredictionsInput) (*machinelearning.DescribeBatchPredictionsOutput, error)
	DescribeBatchPredictionsAsync(ctx workflow.Context, input *machinelearning.DescribeBatchPredictionsInput) *MachinelearningDescribeBatchPredictionsResult

	DescribeDataSources(ctx workflow.Context, input *machinelearning.DescribeDataSourcesInput) (*machinelearning.DescribeDataSourcesOutput, error)
	DescribeDataSourcesAsync(ctx workflow.Context, input *machinelearning.DescribeDataSourcesInput) *MachinelearningDescribeDataSourcesResult

	DescribeEvaluations(ctx workflow.Context, input *machinelearning.DescribeEvaluationsInput) (*machinelearning.DescribeEvaluationsOutput, error)
	DescribeEvaluationsAsync(ctx workflow.Context, input *machinelearning.DescribeEvaluationsInput) *MachinelearningDescribeEvaluationsResult

	DescribeMLModels(ctx workflow.Context, input *machinelearning.DescribeMLModelsInput) (*machinelearning.DescribeMLModelsOutput, error)
	DescribeMLModelsAsync(ctx workflow.Context, input *machinelearning.DescribeMLModelsInput) *MachinelearningDescribeMLModelsResult

	DescribeTags(ctx workflow.Context, input *machinelearning.DescribeTagsInput) (*machinelearning.DescribeTagsOutput, error)
	DescribeTagsAsync(ctx workflow.Context, input *machinelearning.DescribeTagsInput) *MachinelearningDescribeTagsResult

	GetBatchPrediction(ctx workflow.Context, input *machinelearning.GetBatchPredictionInput) (*machinelearning.GetBatchPredictionOutput, error)
	GetBatchPredictionAsync(ctx workflow.Context, input *machinelearning.GetBatchPredictionInput) *MachinelearningGetBatchPredictionResult

	GetDataSource(ctx workflow.Context, input *machinelearning.GetDataSourceInput) (*machinelearning.GetDataSourceOutput, error)
	GetDataSourceAsync(ctx workflow.Context, input *machinelearning.GetDataSourceInput) *MachinelearningGetDataSourceResult

	GetEvaluation(ctx workflow.Context, input *machinelearning.GetEvaluationInput) (*machinelearning.GetEvaluationOutput, error)
	GetEvaluationAsync(ctx workflow.Context, input *machinelearning.GetEvaluationInput) *MachinelearningGetEvaluationResult

	GetMLModel(ctx workflow.Context, input *machinelearning.GetMLModelInput) (*machinelearning.GetMLModelOutput, error)
	GetMLModelAsync(ctx workflow.Context, input *machinelearning.GetMLModelInput) *MachinelearningGetMLModelResult

	Predict(ctx workflow.Context, input *machinelearning.PredictInput) (*machinelearning.PredictOutput, error)
	PredictAsync(ctx workflow.Context, input *machinelearning.PredictInput) *MachinelearningPredictResult

	UpdateBatchPrediction(ctx workflow.Context, input *machinelearning.UpdateBatchPredictionInput) (*machinelearning.UpdateBatchPredictionOutput, error)
	UpdateBatchPredictionAsync(ctx workflow.Context, input *machinelearning.UpdateBatchPredictionInput) *MachinelearningUpdateBatchPredictionResult

	UpdateDataSource(ctx workflow.Context, input *machinelearning.UpdateDataSourceInput) (*machinelearning.UpdateDataSourceOutput, error)
	UpdateDataSourceAsync(ctx workflow.Context, input *machinelearning.UpdateDataSourceInput) *MachinelearningUpdateDataSourceResult

	UpdateEvaluation(ctx workflow.Context, input *machinelearning.UpdateEvaluationInput) (*machinelearning.UpdateEvaluationOutput, error)
	UpdateEvaluationAsync(ctx workflow.Context, input *machinelearning.UpdateEvaluationInput) *MachinelearningUpdateEvaluationResult

	UpdateMLModel(ctx workflow.Context, input *machinelearning.UpdateMLModelInput) (*machinelearning.UpdateMLModelOutput, error)
	UpdateMLModelAsync(ctx workflow.Context, input *machinelearning.UpdateMLModelInput) *MachinelearningUpdateMLModelResult

	WaitUntilBatchPredictionAvailable(ctx workflow.Context, input *machinelearning.DescribeBatchPredictionsInput) error
	WaitUntilDataSourceAvailable(ctx workflow.Context, input *machinelearning.DescribeDataSourcesInput) error
	WaitUntilEvaluationAvailable(ctx workflow.Context, input *machinelearning.DescribeEvaluationsInput) error
	WaitUntilMLModelAvailable(ctx workflow.Context, input *machinelearning.DescribeMLModelsInput) error
}

type MachinelearningAddTagsResult struct {
	Result workflow.Future
}

func (r *MachinelearningAddTagsResult) Get(ctx workflow.Context) (*machinelearning.AddTagsOutput, error) {
	var output machinelearning.AddTagsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type MachinelearningCreateBatchPredictionResult struct {
	Result workflow.Future
}

func (r *MachinelearningCreateBatchPredictionResult) Get(ctx workflow.Context) (*machinelearning.CreateBatchPredictionOutput, error) {
	var output machinelearning.CreateBatchPredictionOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type MachinelearningCreateDataSourceFromRDSResult struct {
	Result workflow.Future
}

func (r *MachinelearningCreateDataSourceFromRDSResult) Get(ctx workflow.Context) (*machinelearning.CreateDataSourceFromRDSOutput, error) {
	var output machinelearning.CreateDataSourceFromRDSOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type MachinelearningCreateDataSourceFromRedshiftResult struct {
	Result workflow.Future
}

func (r *MachinelearningCreateDataSourceFromRedshiftResult) Get(ctx workflow.Context) (*machinelearning.CreateDataSourceFromRedshiftOutput, error) {
	var output machinelearning.CreateDataSourceFromRedshiftOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type MachinelearningCreateDataSourceFromS3Result struct {
	Result workflow.Future
}

func (r *MachinelearningCreateDataSourceFromS3Result) Get(ctx workflow.Context) (*machinelearning.CreateDataSourceFromS3Output, error) {
	var output machinelearning.CreateDataSourceFromS3Output
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type MachinelearningCreateEvaluationResult struct {
	Result workflow.Future
}

func (r *MachinelearningCreateEvaluationResult) Get(ctx workflow.Context) (*machinelearning.CreateEvaluationOutput, error) {
	var output machinelearning.CreateEvaluationOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type MachinelearningCreateMLModelResult struct {
	Result workflow.Future
}

func (r *MachinelearningCreateMLModelResult) Get(ctx workflow.Context) (*machinelearning.CreateMLModelOutput, error) {
	var output machinelearning.CreateMLModelOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type MachinelearningCreateRealtimeEndpointResult struct {
	Result workflow.Future
}

func (r *MachinelearningCreateRealtimeEndpointResult) Get(ctx workflow.Context) (*machinelearning.CreateRealtimeEndpointOutput, error) {
	var output machinelearning.CreateRealtimeEndpointOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type MachinelearningDeleteBatchPredictionResult struct {
	Result workflow.Future
}

func (r *MachinelearningDeleteBatchPredictionResult) Get(ctx workflow.Context) (*machinelearning.DeleteBatchPredictionOutput, error) {
	var output machinelearning.DeleteBatchPredictionOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type MachinelearningDeleteDataSourceResult struct {
	Result workflow.Future
}

func (r *MachinelearningDeleteDataSourceResult) Get(ctx workflow.Context) (*machinelearning.DeleteDataSourceOutput, error) {
	var output machinelearning.DeleteDataSourceOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type MachinelearningDeleteEvaluationResult struct {
	Result workflow.Future
}

func (r *MachinelearningDeleteEvaluationResult) Get(ctx workflow.Context) (*machinelearning.DeleteEvaluationOutput, error) {
	var output machinelearning.DeleteEvaluationOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type MachinelearningDeleteMLModelResult struct {
	Result workflow.Future
}

func (r *MachinelearningDeleteMLModelResult) Get(ctx workflow.Context) (*machinelearning.DeleteMLModelOutput, error) {
	var output machinelearning.DeleteMLModelOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type MachinelearningDeleteRealtimeEndpointResult struct {
	Result workflow.Future
}

func (r *MachinelearningDeleteRealtimeEndpointResult) Get(ctx workflow.Context) (*machinelearning.DeleteRealtimeEndpointOutput, error) {
	var output machinelearning.DeleteRealtimeEndpointOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type MachinelearningDeleteTagsResult struct {
	Result workflow.Future
}

func (r *MachinelearningDeleteTagsResult) Get(ctx workflow.Context) (*machinelearning.DeleteTagsOutput, error) {
	var output machinelearning.DeleteTagsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type MachinelearningDescribeBatchPredictionsResult struct {
	Result workflow.Future
}

func (r *MachinelearningDescribeBatchPredictionsResult) Get(ctx workflow.Context) (*machinelearning.DescribeBatchPredictionsOutput, error) {
	var output machinelearning.DescribeBatchPredictionsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type MachinelearningDescribeDataSourcesResult struct {
	Result workflow.Future
}

func (r *MachinelearningDescribeDataSourcesResult) Get(ctx workflow.Context) (*machinelearning.DescribeDataSourcesOutput, error) {
	var output machinelearning.DescribeDataSourcesOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type MachinelearningDescribeEvaluationsResult struct {
	Result workflow.Future
}

func (r *MachinelearningDescribeEvaluationsResult) Get(ctx workflow.Context) (*machinelearning.DescribeEvaluationsOutput, error) {
	var output machinelearning.DescribeEvaluationsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type MachinelearningDescribeMLModelsResult struct {
	Result workflow.Future
}

func (r *MachinelearningDescribeMLModelsResult) Get(ctx workflow.Context) (*machinelearning.DescribeMLModelsOutput, error) {
	var output machinelearning.DescribeMLModelsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type MachinelearningDescribeTagsResult struct {
	Result workflow.Future
}

func (r *MachinelearningDescribeTagsResult) Get(ctx workflow.Context) (*machinelearning.DescribeTagsOutput, error) {
	var output machinelearning.DescribeTagsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type MachinelearningGetBatchPredictionResult struct {
	Result workflow.Future
}

func (r *MachinelearningGetBatchPredictionResult) Get(ctx workflow.Context) (*machinelearning.GetBatchPredictionOutput, error) {
	var output machinelearning.GetBatchPredictionOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type MachinelearningGetDataSourceResult struct {
	Result workflow.Future
}

func (r *MachinelearningGetDataSourceResult) Get(ctx workflow.Context) (*machinelearning.GetDataSourceOutput, error) {
	var output machinelearning.GetDataSourceOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type MachinelearningGetEvaluationResult struct {
	Result workflow.Future
}

func (r *MachinelearningGetEvaluationResult) Get(ctx workflow.Context) (*machinelearning.GetEvaluationOutput, error) {
	var output machinelearning.GetEvaluationOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type MachinelearningGetMLModelResult struct {
	Result workflow.Future
}

func (r *MachinelearningGetMLModelResult) Get(ctx workflow.Context) (*machinelearning.GetMLModelOutput, error) {
	var output machinelearning.GetMLModelOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type MachinelearningPredictResult struct {
	Result workflow.Future
}

func (r *MachinelearningPredictResult) Get(ctx workflow.Context) (*machinelearning.PredictOutput, error) {
	var output machinelearning.PredictOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type MachinelearningUpdateBatchPredictionResult struct {
	Result workflow.Future
}

func (r *MachinelearningUpdateBatchPredictionResult) Get(ctx workflow.Context) (*machinelearning.UpdateBatchPredictionOutput, error) {
	var output machinelearning.UpdateBatchPredictionOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type MachinelearningUpdateDataSourceResult struct {
	Result workflow.Future
}

func (r *MachinelearningUpdateDataSourceResult) Get(ctx workflow.Context) (*machinelearning.UpdateDataSourceOutput, error) {
	var output machinelearning.UpdateDataSourceOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type MachinelearningUpdateEvaluationResult struct {
	Result workflow.Future
}

func (r *MachinelearningUpdateEvaluationResult) Get(ctx workflow.Context) (*machinelearning.UpdateEvaluationOutput, error) {
	var output machinelearning.UpdateEvaluationOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type MachinelearningUpdateMLModelResult struct {
	Result workflow.Future
}

func (r *MachinelearningUpdateMLModelResult) Get(ctx workflow.Context) (*machinelearning.UpdateMLModelOutput, error) {
	var output machinelearning.UpdateMLModelOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type MachineLearningStub struct {
	activities awsactivities.MachineLearningActivities
}

func NewMachineLearningStub() MachineLearningClient {
	return &MachineLearningStub{}
}

func (a *MachineLearningStub) AddTags(ctx workflow.Context, input *machinelearning.AddTagsInput) (*machinelearning.AddTagsOutput, error) {
	var output machinelearning.AddTagsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.AddTags, input).Get(ctx, &output)
	return &output, err
}

func (a *MachineLearningStub) AddTagsAsync(ctx workflow.Context, input *machinelearning.AddTagsInput) *MachinelearningAddTagsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.AddTags, input)
	return &MachinelearningAddTagsResult{Result: future}
}

func (a *MachineLearningStub) CreateBatchPrediction(ctx workflow.Context, input *machinelearning.CreateBatchPredictionInput) (*machinelearning.CreateBatchPredictionOutput, error) {
	var output machinelearning.CreateBatchPredictionOutput
	err := workflow.ExecuteActivity(ctx, a.activities.CreateBatchPrediction, input).Get(ctx, &output)
	return &output, err
}

func (a *MachineLearningStub) CreateBatchPredictionAsync(ctx workflow.Context, input *machinelearning.CreateBatchPredictionInput) *MachinelearningCreateBatchPredictionResult {
	future := workflow.ExecuteActivity(ctx, a.activities.CreateBatchPrediction, input)
	return &MachinelearningCreateBatchPredictionResult{Result: future}
}

func (a *MachineLearningStub) CreateDataSourceFromRDS(ctx workflow.Context, input *machinelearning.CreateDataSourceFromRDSInput) (*machinelearning.CreateDataSourceFromRDSOutput, error) {
	var output machinelearning.CreateDataSourceFromRDSOutput
	err := workflow.ExecuteActivity(ctx, a.activities.CreateDataSourceFromRDS, input).Get(ctx, &output)
	return &output, err
}

func (a *MachineLearningStub) CreateDataSourceFromRDSAsync(ctx workflow.Context, input *machinelearning.CreateDataSourceFromRDSInput) *MachinelearningCreateDataSourceFromRDSResult {
	future := workflow.ExecuteActivity(ctx, a.activities.CreateDataSourceFromRDS, input)
	return &MachinelearningCreateDataSourceFromRDSResult{Result: future}
}

func (a *MachineLearningStub) CreateDataSourceFromRedshift(ctx workflow.Context, input *machinelearning.CreateDataSourceFromRedshiftInput) (*machinelearning.CreateDataSourceFromRedshiftOutput, error) {
	var output machinelearning.CreateDataSourceFromRedshiftOutput
	err := workflow.ExecuteActivity(ctx, a.activities.CreateDataSourceFromRedshift, input).Get(ctx, &output)
	return &output, err
}

func (a *MachineLearningStub) CreateDataSourceFromRedshiftAsync(ctx workflow.Context, input *machinelearning.CreateDataSourceFromRedshiftInput) *MachinelearningCreateDataSourceFromRedshiftResult {
	future := workflow.ExecuteActivity(ctx, a.activities.CreateDataSourceFromRedshift, input)
	return &MachinelearningCreateDataSourceFromRedshiftResult{Result: future}
}

func (a *MachineLearningStub) CreateDataSourceFromS3(ctx workflow.Context, input *machinelearning.CreateDataSourceFromS3Input) (*machinelearning.CreateDataSourceFromS3Output, error) {
	var output machinelearning.CreateDataSourceFromS3Output
	err := workflow.ExecuteActivity(ctx, a.activities.CreateDataSourceFromS3, input).Get(ctx, &output)
	return &output, err
}

func (a *MachineLearningStub) CreateDataSourceFromS3Async(ctx workflow.Context, input *machinelearning.CreateDataSourceFromS3Input) *MachinelearningCreateDataSourceFromS3Result {
	future := workflow.ExecuteActivity(ctx, a.activities.CreateDataSourceFromS3, input)
	return &MachinelearningCreateDataSourceFromS3Result{Result: future}
}

func (a *MachineLearningStub) CreateEvaluation(ctx workflow.Context, input *machinelearning.CreateEvaluationInput) (*machinelearning.CreateEvaluationOutput, error) {
	var output machinelearning.CreateEvaluationOutput
	err := workflow.ExecuteActivity(ctx, a.activities.CreateEvaluation, input).Get(ctx, &output)
	return &output, err
}

func (a *MachineLearningStub) CreateEvaluationAsync(ctx workflow.Context, input *machinelearning.CreateEvaluationInput) *MachinelearningCreateEvaluationResult {
	future := workflow.ExecuteActivity(ctx, a.activities.CreateEvaluation, input)
	return &MachinelearningCreateEvaluationResult{Result: future}
}

func (a *MachineLearningStub) CreateMLModel(ctx workflow.Context, input *machinelearning.CreateMLModelInput) (*machinelearning.CreateMLModelOutput, error) {
	var output machinelearning.CreateMLModelOutput
	err := workflow.ExecuteActivity(ctx, a.activities.CreateMLModel, input).Get(ctx, &output)
	return &output, err
}

func (a *MachineLearningStub) CreateMLModelAsync(ctx workflow.Context, input *machinelearning.CreateMLModelInput) *MachinelearningCreateMLModelResult {
	future := workflow.ExecuteActivity(ctx, a.activities.CreateMLModel, input)
	return &MachinelearningCreateMLModelResult{Result: future}
}

func (a *MachineLearningStub) CreateRealtimeEndpoint(ctx workflow.Context, input *machinelearning.CreateRealtimeEndpointInput) (*machinelearning.CreateRealtimeEndpointOutput, error) {
	var output machinelearning.CreateRealtimeEndpointOutput
	err := workflow.ExecuteActivity(ctx, a.activities.CreateRealtimeEndpoint, input).Get(ctx, &output)
	return &output, err
}

func (a *MachineLearningStub) CreateRealtimeEndpointAsync(ctx workflow.Context, input *machinelearning.CreateRealtimeEndpointInput) *MachinelearningCreateRealtimeEndpointResult {
	future := workflow.ExecuteActivity(ctx, a.activities.CreateRealtimeEndpoint, input)
	return &MachinelearningCreateRealtimeEndpointResult{Result: future}
}

func (a *MachineLearningStub) DeleteBatchPrediction(ctx workflow.Context, input *machinelearning.DeleteBatchPredictionInput) (*machinelearning.DeleteBatchPredictionOutput, error) {
	var output machinelearning.DeleteBatchPredictionOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeleteBatchPrediction, input).Get(ctx, &output)
	return &output, err
}

func (a *MachineLearningStub) DeleteBatchPredictionAsync(ctx workflow.Context, input *machinelearning.DeleteBatchPredictionInput) *MachinelearningDeleteBatchPredictionResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeleteBatchPrediction, input)
	return &MachinelearningDeleteBatchPredictionResult{Result: future}
}

func (a *MachineLearningStub) DeleteDataSource(ctx workflow.Context, input *machinelearning.DeleteDataSourceInput) (*machinelearning.DeleteDataSourceOutput, error) {
	var output machinelearning.DeleteDataSourceOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeleteDataSource, input).Get(ctx, &output)
	return &output, err
}

func (a *MachineLearningStub) DeleteDataSourceAsync(ctx workflow.Context, input *machinelearning.DeleteDataSourceInput) *MachinelearningDeleteDataSourceResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeleteDataSource, input)
	return &MachinelearningDeleteDataSourceResult{Result: future}
}

func (a *MachineLearningStub) DeleteEvaluation(ctx workflow.Context, input *machinelearning.DeleteEvaluationInput) (*machinelearning.DeleteEvaluationOutput, error) {
	var output machinelearning.DeleteEvaluationOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeleteEvaluation, input).Get(ctx, &output)
	return &output, err
}

func (a *MachineLearningStub) DeleteEvaluationAsync(ctx workflow.Context, input *machinelearning.DeleteEvaluationInput) *MachinelearningDeleteEvaluationResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeleteEvaluation, input)
	return &MachinelearningDeleteEvaluationResult{Result: future}
}

func (a *MachineLearningStub) DeleteMLModel(ctx workflow.Context, input *machinelearning.DeleteMLModelInput) (*machinelearning.DeleteMLModelOutput, error) {
	var output machinelearning.DeleteMLModelOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeleteMLModel, input).Get(ctx, &output)
	return &output, err
}

func (a *MachineLearningStub) DeleteMLModelAsync(ctx workflow.Context, input *machinelearning.DeleteMLModelInput) *MachinelearningDeleteMLModelResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeleteMLModel, input)
	return &MachinelearningDeleteMLModelResult{Result: future}
}

func (a *MachineLearningStub) DeleteRealtimeEndpoint(ctx workflow.Context, input *machinelearning.DeleteRealtimeEndpointInput) (*machinelearning.DeleteRealtimeEndpointOutput, error) {
	var output machinelearning.DeleteRealtimeEndpointOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeleteRealtimeEndpoint, input).Get(ctx, &output)
	return &output, err
}

func (a *MachineLearningStub) DeleteRealtimeEndpointAsync(ctx workflow.Context, input *machinelearning.DeleteRealtimeEndpointInput) *MachinelearningDeleteRealtimeEndpointResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeleteRealtimeEndpoint, input)
	return &MachinelearningDeleteRealtimeEndpointResult{Result: future}
}

func (a *MachineLearningStub) DeleteTags(ctx workflow.Context, input *machinelearning.DeleteTagsInput) (*machinelearning.DeleteTagsOutput, error) {
	var output machinelearning.DeleteTagsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeleteTags, input).Get(ctx, &output)
	return &output, err
}

func (a *MachineLearningStub) DeleteTagsAsync(ctx workflow.Context, input *machinelearning.DeleteTagsInput) *MachinelearningDeleteTagsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeleteTags, input)
	return &MachinelearningDeleteTagsResult{Result: future}
}

func (a *MachineLearningStub) DescribeBatchPredictions(ctx workflow.Context, input *machinelearning.DescribeBatchPredictionsInput) (*machinelearning.DescribeBatchPredictionsOutput, error) {
	var output machinelearning.DescribeBatchPredictionsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DescribeBatchPredictions, input).Get(ctx, &output)
	return &output, err
}

func (a *MachineLearningStub) DescribeBatchPredictionsAsync(ctx workflow.Context, input *machinelearning.DescribeBatchPredictionsInput) *MachinelearningDescribeBatchPredictionsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DescribeBatchPredictions, input)
	return &MachinelearningDescribeBatchPredictionsResult{Result: future}
}

func (a *MachineLearningStub) DescribeDataSources(ctx workflow.Context, input *machinelearning.DescribeDataSourcesInput) (*machinelearning.DescribeDataSourcesOutput, error) {
	var output machinelearning.DescribeDataSourcesOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DescribeDataSources, input).Get(ctx, &output)
	return &output, err
}

func (a *MachineLearningStub) DescribeDataSourcesAsync(ctx workflow.Context, input *machinelearning.DescribeDataSourcesInput) *MachinelearningDescribeDataSourcesResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DescribeDataSources, input)
	return &MachinelearningDescribeDataSourcesResult{Result: future}
}

func (a *MachineLearningStub) DescribeEvaluations(ctx workflow.Context, input *machinelearning.DescribeEvaluationsInput) (*machinelearning.DescribeEvaluationsOutput, error) {
	var output machinelearning.DescribeEvaluationsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DescribeEvaluations, input).Get(ctx, &output)
	return &output, err
}

func (a *MachineLearningStub) DescribeEvaluationsAsync(ctx workflow.Context, input *machinelearning.DescribeEvaluationsInput) *MachinelearningDescribeEvaluationsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DescribeEvaluations, input)
	return &MachinelearningDescribeEvaluationsResult{Result: future}
}

func (a *MachineLearningStub) DescribeMLModels(ctx workflow.Context, input *machinelearning.DescribeMLModelsInput) (*machinelearning.DescribeMLModelsOutput, error) {
	var output machinelearning.DescribeMLModelsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DescribeMLModels, input).Get(ctx, &output)
	return &output, err
}

func (a *MachineLearningStub) DescribeMLModelsAsync(ctx workflow.Context, input *machinelearning.DescribeMLModelsInput) *MachinelearningDescribeMLModelsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DescribeMLModels, input)
	return &MachinelearningDescribeMLModelsResult{Result: future}
}

func (a *MachineLearningStub) DescribeTags(ctx workflow.Context, input *machinelearning.DescribeTagsInput) (*machinelearning.DescribeTagsOutput, error) {
	var output machinelearning.DescribeTagsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DescribeTags, input).Get(ctx, &output)
	return &output, err
}

func (a *MachineLearningStub) DescribeTagsAsync(ctx workflow.Context, input *machinelearning.DescribeTagsInput) *MachinelearningDescribeTagsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DescribeTags, input)
	return &MachinelearningDescribeTagsResult{Result: future}
}

func (a *MachineLearningStub) GetBatchPrediction(ctx workflow.Context, input *machinelearning.GetBatchPredictionInput) (*machinelearning.GetBatchPredictionOutput, error) {
	var output machinelearning.GetBatchPredictionOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetBatchPrediction, input).Get(ctx, &output)
	return &output, err
}

func (a *MachineLearningStub) GetBatchPredictionAsync(ctx workflow.Context, input *machinelearning.GetBatchPredictionInput) *MachinelearningGetBatchPredictionResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetBatchPrediction, input)
	return &MachinelearningGetBatchPredictionResult{Result: future}
}

func (a *MachineLearningStub) GetDataSource(ctx workflow.Context, input *machinelearning.GetDataSourceInput) (*machinelearning.GetDataSourceOutput, error) {
	var output machinelearning.GetDataSourceOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetDataSource, input).Get(ctx, &output)
	return &output, err
}

func (a *MachineLearningStub) GetDataSourceAsync(ctx workflow.Context, input *machinelearning.GetDataSourceInput) *MachinelearningGetDataSourceResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetDataSource, input)
	return &MachinelearningGetDataSourceResult{Result: future}
}

func (a *MachineLearningStub) GetEvaluation(ctx workflow.Context, input *machinelearning.GetEvaluationInput) (*machinelearning.GetEvaluationOutput, error) {
	var output machinelearning.GetEvaluationOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetEvaluation, input).Get(ctx, &output)
	return &output, err
}

func (a *MachineLearningStub) GetEvaluationAsync(ctx workflow.Context, input *machinelearning.GetEvaluationInput) *MachinelearningGetEvaluationResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetEvaluation, input)
	return &MachinelearningGetEvaluationResult{Result: future}
}

func (a *MachineLearningStub) GetMLModel(ctx workflow.Context, input *machinelearning.GetMLModelInput) (*machinelearning.GetMLModelOutput, error) {
	var output machinelearning.GetMLModelOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetMLModel, input).Get(ctx, &output)
	return &output, err
}

func (a *MachineLearningStub) GetMLModelAsync(ctx workflow.Context, input *machinelearning.GetMLModelInput) *MachinelearningGetMLModelResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetMLModel, input)
	return &MachinelearningGetMLModelResult{Result: future}
}

func (a *MachineLearningStub) Predict(ctx workflow.Context, input *machinelearning.PredictInput) (*machinelearning.PredictOutput, error) {
	var output machinelearning.PredictOutput
	err := workflow.ExecuteActivity(ctx, a.activities.Predict, input).Get(ctx, &output)
	return &output, err
}

func (a *MachineLearningStub) PredictAsync(ctx workflow.Context, input *machinelearning.PredictInput) *MachinelearningPredictResult {
	future := workflow.ExecuteActivity(ctx, a.activities.Predict, input)
	return &MachinelearningPredictResult{Result: future}
}

func (a *MachineLearningStub) UpdateBatchPrediction(ctx workflow.Context, input *machinelearning.UpdateBatchPredictionInput) (*machinelearning.UpdateBatchPredictionOutput, error) {
	var output machinelearning.UpdateBatchPredictionOutput
	err := workflow.ExecuteActivity(ctx, a.activities.UpdateBatchPrediction, input).Get(ctx, &output)
	return &output, err
}

func (a *MachineLearningStub) UpdateBatchPredictionAsync(ctx workflow.Context, input *machinelearning.UpdateBatchPredictionInput) *MachinelearningUpdateBatchPredictionResult {
	future := workflow.ExecuteActivity(ctx, a.activities.UpdateBatchPrediction, input)
	return &MachinelearningUpdateBatchPredictionResult{Result: future}
}

func (a *MachineLearningStub) UpdateDataSource(ctx workflow.Context, input *machinelearning.UpdateDataSourceInput) (*machinelearning.UpdateDataSourceOutput, error) {
	var output machinelearning.UpdateDataSourceOutput
	err := workflow.ExecuteActivity(ctx, a.activities.UpdateDataSource, input).Get(ctx, &output)
	return &output, err
}

func (a *MachineLearningStub) UpdateDataSourceAsync(ctx workflow.Context, input *machinelearning.UpdateDataSourceInput) *MachinelearningUpdateDataSourceResult {
	future := workflow.ExecuteActivity(ctx, a.activities.UpdateDataSource, input)
	return &MachinelearningUpdateDataSourceResult{Result: future}
}

func (a *MachineLearningStub) UpdateEvaluation(ctx workflow.Context, input *machinelearning.UpdateEvaluationInput) (*machinelearning.UpdateEvaluationOutput, error) {
	var output machinelearning.UpdateEvaluationOutput
	err := workflow.ExecuteActivity(ctx, a.activities.UpdateEvaluation, input).Get(ctx, &output)
	return &output, err
}

func (a *MachineLearningStub) UpdateEvaluationAsync(ctx workflow.Context, input *machinelearning.UpdateEvaluationInput) *MachinelearningUpdateEvaluationResult {
	future := workflow.ExecuteActivity(ctx, a.activities.UpdateEvaluation, input)
	return &MachinelearningUpdateEvaluationResult{Result: future}
}

func (a *MachineLearningStub) UpdateMLModel(ctx workflow.Context, input *machinelearning.UpdateMLModelInput) (*machinelearning.UpdateMLModelOutput, error) {
	var output machinelearning.UpdateMLModelOutput
	err := workflow.ExecuteActivity(ctx, a.activities.UpdateMLModel, input).Get(ctx, &output)
	return &output, err
}

func (a *MachineLearningStub) UpdateMLModelAsync(ctx workflow.Context, input *machinelearning.UpdateMLModelInput) *MachinelearningUpdateMLModelResult {
	future := workflow.ExecuteActivity(ctx, a.activities.UpdateMLModel, input)
	return &MachinelearningUpdateMLModelResult{Result: future}
}

func (a *MachineLearningStub) WaitUntilBatchPredictionAvailable(ctx workflow.Context, input *machinelearning.DescribeBatchPredictionsInput) error {
	return workflow.ExecuteActivity(ctx, a.activities.WaitUntilBatchPredictionAvailable, input).Get(ctx, nil)
}

func (a *MachineLearningStub) WaitUntilBatchPredictionAvailableAsync(ctx workflow.Context, input *machinelearning.DescribeBatchPredictionsInput) workflow.Future {
	return workflow.ExecuteActivity(ctx, a.activities.WaitUntilBatchPredictionAvailable, input)
}

func (a *MachineLearningStub) WaitUntilDataSourceAvailable(ctx workflow.Context, input *machinelearning.DescribeDataSourcesInput) error {
	return workflow.ExecuteActivity(ctx, a.activities.WaitUntilDataSourceAvailable, input).Get(ctx, nil)
}

func (a *MachineLearningStub) WaitUntilDataSourceAvailableAsync(ctx workflow.Context, input *machinelearning.DescribeDataSourcesInput) workflow.Future {
	return workflow.ExecuteActivity(ctx, a.activities.WaitUntilDataSourceAvailable, input)
}

func (a *MachineLearningStub) WaitUntilEvaluationAvailable(ctx workflow.Context, input *machinelearning.DescribeEvaluationsInput) error {
	return workflow.ExecuteActivity(ctx, a.activities.WaitUntilEvaluationAvailable, input).Get(ctx, nil)
}

func (a *MachineLearningStub) WaitUntilEvaluationAvailableAsync(ctx workflow.Context, input *machinelearning.DescribeEvaluationsInput) workflow.Future {
	return workflow.ExecuteActivity(ctx, a.activities.WaitUntilEvaluationAvailable, input)
}

func (a *MachineLearningStub) WaitUntilMLModelAvailable(ctx workflow.Context, input *machinelearning.DescribeMLModelsInput) error {
	return workflow.ExecuteActivity(ctx, a.activities.WaitUntilMLModelAvailable, input).Get(ctx, nil)
}

func (a *MachineLearningStub) WaitUntilMLModelAvailableAsync(ctx workflow.Context, input *machinelearning.DescribeMLModelsInput) workflow.Future {
	return workflow.ExecuteActivity(ctx, a.activities.WaitUntilMLModelAvailable, input)
}
