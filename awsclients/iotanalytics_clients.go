package awsclients

import (
	"github.com/aws/aws-sdk-go/service/iotanalytics"
	"go.temporal.io/sdk/workflow"
	"temporal.io/aws-sdk/awsactivities"
)

type IoTAnalyticsClient interface {
       BatchPutMessage(ctx workflow.Context, input *iotanalytics.BatchPutMessageInput) (*iotanalytics.BatchPutMessageOutput, error)
       BatchPutMessageAsync(ctx workflow.Context, input *iotanalytics.BatchPutMessageInput) *IotanalyticsBatchPutMessageResult

       CancelPipelineReprocessing(ctx workflow.Context, input *iotanalytics.CancelPipelineReprocessingInput) (*iotanalytics.CancelPipelineReprocessingOutput, error)
       CancelPipelineReprocessingAsync(ctx workflow.Context, input *iotanalytics.CancelPipelineReprocessingInput) *IotanalyticsCancelPipelineReprocessingResult

       CreateChannel(ctx workflow.Context, input *iotanalytics.CreateChannelInput) (*iotanalytics.CreateChannelOutput, error)
       CreateChannelAsync(ctx workflow.Context, input *iotanalytics.CreateChannelInput) *IotanalyticsCreateChannelResult

       CreateDataset(ctx workflow.Context, input *iotanalytics.CreateDatasetInput) (*iotanalytics.CreateDatasetOutput, error)
       CreateDatasetAsync(ctx workflow.Context, input *iotanalytics.CreateDatasetInput) *IotanalyticsCreateDatasetResult

       CreateDatasetContent(ctx workflow.Context, input *iotanalytics.CreateDatasetContentInput) (*iotanalytics.CreateDatasetContentOutput, error)
       CreateDatasetContentAsync(ctx workflow.Context, input *iotanalytics.CreateDatasetContentInput) *IotanalyticsCreateDatasetContentResult

       CreateDatastore(ctx workflow.Context, input *iotanalytics.CreateDatastoreInput) (*iotanalytics.CreateDatastoreOutput, error)
       CreateDatastoreAsync(ctx workflow.Context, input *iotanalytics.CreateDatastoreInput) *IotanalyticsCreateDatastoreResult

       CreatePipeline(ctx workflow.Context, input *iotanalytics.CreatePipelineInput) (*iotanalytics.CreatePipelineOutput, error)
       CreatePipelineAsync(ctx workflow.Context, input *iotanalytics.CreatePipelineInput) *IotanalyticsCreatePipelineResult

       DeleteChannel(ctx workflow.Context, input *iotanalytics.DeleteChannelInput) (*iotanalytics.DeleteChannelOutput, error)
       DeleteChannelAsync(ctx workflow.Context, input *iotanalytics.DeleteChannelInput) *IotanalyticsDeleteChannelResult

       DeleteDataset(ctx workflow.Context, input *iotanalytics.DeleteDatasetInput) (*iotanalytics.DeleteDatasetOutput, error)
       DeleteDatasetAsync(ctx workflow.Context, input *iotanalytics.DeleteDatasetInput) *IotanalyticsDeleteDatasetResult

       DeleteDatasetContent(ctx workflow.Context, input *iotanalytics.DeleteDatasetContentInput) (*iotanalytics.DeleteDatasetContentOutput, error)
       DeleteDatasetContentAsync(ctx workflow.Context, input *iotanalytics.DeleteDatasetContentInput) *IotanalyticsDeleteDatasetContentResult

       DeleteDatastore(ctx workflow.Context, input *iotanalytics.DeleteDatastoreInput) (*iotanalytics.DeleteDatastoreOutput, error)
       DeleteDatastoreAsync(ctx workflow.Context, input *iotanalytics.DeleteDatastoreInput) *IotanalyticsDeleteDatastoreResult

       DeletePipeline(ctx workflow.Context, input *iotanalytics.DeletePipelineInput) (*iotanalytics.DeletePipelineOutput, error)
       DeletePipelineAsync(ctx workflow.Context, input *iotanalytics.DeletePipelineInput) *IotanalyticsDeletePipelineResult

       DescribeChannel(ctx workflow.Context, input *iotanalytics.DescribeChannelInput) (*iotanalytics.DescribeChannelOutput, error)
       DescribeChannelAsync(ctx workflow.Context, input *iotanalytics.DescribeChannelInput) *IotanalyticsDescribeChannelResult

       DescribeDataset(ctx workflow.Context, input *iotanalytics.DescribeDatasetInput) (*iotanalytics.DescribeDatasetOutput, error)
       DescribeDatasetAsync(ctx workflow.Context, input *iotanalytics.DescribeDatasetInput) *IotanalyticsDescribeDatasetResult

       DescribeDatastore(ctx workflow.Context, input *iotanalytics.DescribeDatastoreInput) (*iotanalytics.DescribeDatastoreOutput, error)
       DescribeDatastoreAsync(ctx workflow.Context, input *iotanalytics.DescribeDatastoreInput) *IotanalyticsDescribeDatastoreResult

       DescribeLoggingOptions(ctx workflow.Context, input *iotanalytics.DescribeLoggingOptionsInput) (*iotanalytics.DescribeLoggingOptionsOutput, error)
       DescribeLoggingOptionsAsync(ctx workflow.Context, input *iotanalytics.DescribeLoggingOptionsInput) *IotanalyticsDescribeLoggingOptionsResult

       DescribePipeline(ctx workflow.Context, input *iotanalytics.DescribePipelineInput) (*iotanalytics.DescribePipelineOutput, error)
       DescribePipelineAsync(ctx workflow.Context, input *iotanalytics.DescribePipelineInput) *IotanalyticsDescribePipelineResult

       GetDatasetContent(ctx workflow.Context, input *iotanalytics.GetDatasetContentInput) (*iotanalytics.GetDatasetContentOutput, error)
       GetDatasetContentAsync(ctx workflow.Context, input *iotanalytics.GetDatasetContentInput) *IotanalyticsGetDatasetContentResult

       ListChannels(ctx workflow.Context, input *iotanalytics.ListChannelsInput) (*iotanalytics.ListChannelsOutput, error)
       ListChannelsAsync(ctx workflow.Context, input *iotanalytics.ListChannelsInput) *IotanalyticsListChannelsResult

       ListDatasetContents(ctx workflow.Context, input *iotanalytics.ListDatasetContentsInput) (*iotanalytics.ListDatasetContentsOutput, error)
       ListDatasetContentsAsync(ctx workflow.Context, input *iotanalytics.ListDatasetContentsInput) *IotanalyticsListDatasetContentsResult

       ListDatasets(ctx workflow.Context, input *iotanalytics.ListDatasetsInput) (*iotanalytics.ListDatasetsOutput, error)
       ListDatasetsAsync(ctx workflow.Context, input *iotanalytics.ListDatasetsInput) *IotanalyticsListDatasetsResult

       ListDatastores(ctx workflow.Context, input *iotanalytics.ListDatastoresInput) (*iotanalytics.ListDatastoresOutput, error)
       ListDatastoresAsync(ctx workflow.Context, input *iotanalytics.ListDatastoresInput) *IotanalyticsListDatastoresResult

       ListPipelines(ctx workflow.Context, input *iotanalytics.ListPipelinesInput) (*iotanalytics.ListPipelinesOutput, error)
       ListPipelinesAsync(ctx workflow.Context, input *iotanalytics.ListPipelinesInput) *IotanalyticsListPipelinesResult

       ListTagsForResource(ctx workflow.Context, input *iotanalytics.ListTagsForResourceInput) (*iotanalytics.ListTagsForResourceOutput, error)
       ListTagsForResourceAsync(ctx workflow.Context, input *iotanalytics.ListTagsForResourceInput) *IotanalyticsListTagsForResourceResult

       PutLoggingOptions(ctx workflow.Context, input *iotanalytics.PutLoggingOptionsInput) (*iotanalytics.PutLoggingOptionsOutput, error)
       PutLoggingOptionsAsync(ctx workflow.Context, input *iotanalytics.PutLoggingOptionsInput) *IotanalyticsPutLoggingOptionsResult

       RunPipelineActivity(ctx workflow.Context, input *iotanalytics.RunPipelineActivityInput) (*iotanalytics.RunPipelineActivityOutput, error)
       RunPipelineActivityAsync(ctx workflow.Context, input *iotanalytics.RunPipelineActivityInput) *IotanalyticsRunPipelineActivityResult

       SampleChannelData(ctx workflow.Context, input *iotanalytics.SampleChannelDataInput) (*iotanalytics.SampleChannelDataOutput, error)
       SampleChannelDataAsync(ctx workflow.Context, input *iotanalytics.SampleChannelDataInput) *IotanalyticsSampleChannelDataResult

       StartPipelineReprocessing(ctx workflow.Context, input *iotanalytics.StartPipelineReprocessingInput) (*iotanalytics.StartPipelineReprocessingOutput, error)
       StartPipelineReprocessingAsync(ctx workflow.Context, input *iotanalytics.StartPipelineReprocessingInput) *IotanalyticsStartPipelineReprocessingResult

       TagResource(ctx workflow.Context, input *iotanalytics.TagResourceInput) (*iotanalytics.TagResourceOutput, error)
       TagResourceAsync(ctx workflow.Context, input *iotanalytics.TagResourceInput) *IotanalyticsTagResourceResult

       UntagResource(ctx workflow.Context, input *iotanalytics.UntagResourceInput) (*iotanalytics.UntagResourceOutput, error)
       UntagResourceAsync(ctx workflow.Context, input *iotanalytics.UntagResourceInput) *IotanalyticsUntagResourceResult

       UpdateChannel(ctx workflow.Context, input *iotanalytics.UpdateChannelInput) (*iotanalytics.UpdateChannelOutput, error)
       UpdateChannelAsync(ctx workflow.Context, input *iotanalytics.UpdateChannelInput) *IotanalyticsUpdateChannelResult

       UpdateDataset(ctx workflow.Context, input *iotanalytics.UpdateDatasetInput) (*iotanalytics.UpdateDatasetOutput, error)
       UpdateDatasetAsync(ctx workflow.Context, input *iotanalytics.UpdateDatasetInput) *IotanalyticsUpdateDatasetResult

       UpdateDatastore(ctx workflow.Context, input *iotanalytics.UpdateDatastoreInput) (*iotanalytics.UpdateDatastoreOutput, error)
       UpdateDatastoreAsync(ctx workflow.Context, input *iotanalytics.UpdateDatastoreInput) *IotanalyticsUpdateDatastoreResult

       UpdatePipeline(ctx workflow.Context, input *iotanalytics.UpdatePipelineInput) (*iotanalytics.UpdatePipelineOutput, error)
       UpdatePipelineAsync(ctx workflow.Context, input *iotanalytics.UpdatePipelineInput) *IotanalyticsUpdatePipelineResult
}

type IotanalyticsBatchPutMessageResult struct {
	Result workflow.Future
}

func (r *IotanalyticsBatchPutMessageResult) Get(ctx workflow.Context) (*iotanalytics.BatchPutMessageOutput, error) {
    var output iotanalytics.BatchPutMessageOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotanalyticsCancelPipelineReprocessingResult struct {
	Result workflow.Future
}

func (r *IotanalyticsCancelPipelineReprocessingResult) Get(ctx workflow.Context) (*iotanalytics.CancelPipelineReprocessingOutput, error) {
    var output iotanalytics.CancelPipelineReprocessingOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotanalyticsCreateChannelResult struct {
	Result workflow.Future
}

func (r *IotanalyticsCreateChannelResult) Get(ctx workflow.Context) (*iotanalytics.CreateChannelOutput, error) {
    var output iotanalytics.CreateChannelOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotanalyticsCreateDatasetResult struct {
	Result workflow.Future
}

func (r *IotanalyticsCreateDatasetResult) Get(ctx workflow.Context) (*iotanalytics.CreateDatasetOutput, error) {
    var output iotanalytics.CreateDatasetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotanalyticsCreateDatasetContentResult struct {
	Result workflow.Future
}

func (r *IotanalyticsCreateDatasetContentResult) Get(ctx workflow.Context) (*iotanalytics.CreateDatasetContentOutput, error) {
    var output iotanalytics.CreateDatasetContentOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotanalyticsCreateDatastoreResult struct {
	Result workflow.Future
}

func (r *IotanalyticsCreateDatastoreResult) Get(ctx workflow.Context) (*iotanalytics.CreateDatastoreOutput, error) {
    var output iotanalytics.CreateDatastoreOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotanalyticsCreatePipelineResult struct {
	Result workflow.Future
}

func (r *IotanalyticsCreatePipelineResult) Get(ctx workflow.Context) (*iotanalytics.CreatePipelineOutput, error) {
    var output iotanalytics.CreatePipelineOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotanalyticsDeleteChannelResult struct {
	Result workflow.Future
}

func (r *IotanalyticsDeleteChannelResult) Get(ctx workflow.Context) (*iotanalytics.DeleteChannelOutput, error) {
    var output iotanalytics.DeleteChannelOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotanalyticsDeleteDatasetResult struct {
	Result workflow.Future
}

func (r *IotanalyticsDeleteDatasetResult) Get(ctx workflow.Context) (*iotanalytics.DeleteDatasetOutput, error) {
    var output iotanalytics.DeleteDatasetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotanalyticsDeleteDatasetContentResult struct {
	Result workflow.Future
}

func (r *IotanalyticsDeleteDatasetContentResult) Get(ctx workflow.Context) (*iotanalytics.DeleteDatasetContentOutput, error) {
    var output iotanalytics.DeleteDatasetContentOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotanalyticsDeleteDatastoreResult struct {
	Result workflow.Future
}

func (r *IotanalyticsDeleteDatastoreResult) Get(ctx workflow.Context) (*iotanalytics.DeleteDatastoreOutput, error) {
    var output iotanalytics.DeleteDatastoreOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotanalyticsDeletePipelineResult struct {
	Result workflow.Future
}

func (r *IotanalyticsDeletePipelineResult) Get(ctx workflow.Context) (*iotanalytics.DeletePipelineOutput, error) {
    var output iotanalytics.DeletePipelineOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotanalyticsDescribeChannelResult struct {
	Result workflow.Future
}

func (r *IotanalyticsDescribeChannelResult) Get(ctx workflow.Context) (*iotanalytics.DescribeChannelOutput, error) {
    var output iotanalytics.DescribeChannelOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotanalyticsDescribeDatasetResult struct {
	Result workflow.Future
}

func (r *IotanalyticsDescribeDatasetResult) Get(ctx workflow.Context) (*iotanalytics.DescribeDatasetOutput, error) {
    var output iotanalytics.DescribeDatasetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotanalyticsDescribeDatastoreResult struct {
	Result workflow.Future
}

func (r *IotanalyticsDescribeDatastoreResult) Get(ctx workflow.Context) (*iotanalytics.DescribeDatastoreOutput, error) {
    var output iotanalytics.DescribeDatastoreOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotanalyticsDescribeLoggingOptionsResult struct {
	Result workflow.Future
}

func (r *IotanalyticsDescribeLoggingOptionsResult) Get(ctx workflow.Context) (*iotanalytics.DescribeLoggingOptionsOutput, error) {
    var output iotanalytics.DescribeLoggingOptionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotanalyticsDescribePipelineResult struct {
	Result workflow.Future
}

func (r *IotanalyticsDescribePipelineResult) Get(ctx workflow.Context) (*iotanalytics.DescribePipelineOutput, error) {
    var output iotanalytics.DescribePipelineOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotanalyticsGetDatasetContentResult struct {
	Result workflow.Future
}

func (r *IotanalyticsGetDatasetContentResult) Get(ctx workflow.Context) (*iotanalytics.GetDatasetContentOutput, error) {
    var output iotanalytics.GetDatasetContentOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotanalyticsListChannelsResult struct {
	Result workflow.Future
}

func (r *IotanalyticsListChannelsResult) Get(ctx workflow.Context) (*iotanalytics.ListChannelsOutput, error) {
    var output iotanalytics.ListChannelsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotanalyticsListDatasetContentsResult struct {
	Result workflow.Future
}

func (r *IotanalyticsListDatasetContentsResult) Get(ctx workflow.Context) (*iotanalytics.ListDatasetContentsOutput, error) {
    var output iotanalytics.ListDatasetContentsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotanalyticsListDatasetsResult struct {
	Result workflow.Future
}

func (r *IotanalyticsListDatasetsResult) Get(ctx workflow.Context) (*iotanalytics.ListDatasetsOutput, error) {
    var output iotanalytics.ListDatasetsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotanalyticsListDatastoresResult struct {
	Result workflow.Future
}

func (r *IotanalyticsListDatastoresResult) Get(ctx workflow.Context) (*iotanalytics.ListDatastoresOutput, error) {
    var output iotanalytics.ListDatastoresOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotanalyticsListPipelinesResult struct {
	Result workflow.Future
}

func (r *IotanalyticsListPipelinesResult) Get(ctx workflow.Context) (*iotanalytics.ListPipelinesOutput, error) {
    var output iotanalytics.ListPipelinesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotanalyticsListTagsForResourceResult struct {
	Result workflow.Future
}

func (r *IotanalyticsListTagsForResourceResult) Get(ctx workflow.Context) (*iotanalytics.ListTagsForResourceOutput, error) {
    var output iotanalytics.ListTagsForResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotanalyticsPutLoggingOptionsResult struct {
	Result workflow.Future
}

func (r *IotanalyticsPutLoggingOptionsResult) Get(ctx workflow.Context) (*iotanalytics.PutLoggingOptionsOutput, error) {
    var output iotanalytics.PutLoggingOptionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotanalyticsRunPipelineActivityResult struct {
	Result workflow.Future
}

func (r *IotanalyticsRunPipelineActivityResult) Get(ctx workflow.Context) (*iotanalytics.RunPipelineActivityOutput, error) {
    var output iotanalytics.RunPipelineActivityOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotanalyticsSampleChannelDataResult struct {
	Result workflow.Future
}

func (r *IotanalyticsSampleChannelDataResult) Get(ctx workflow.Context) (*iotanalytics.SampleChannelDataOutput, error) {
    var output iotanalytics.SampleChannelDataOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotanalyticsStartPipelineReprocessingResult struct {
	Result workflow.Future
}

func (r *IotanalyticsStartPipelineReprocessingResult) Get(ctx workflow.Context) (*iotanalytics.StartPipelineReprocessingOutput, error) {
    var output iotanalytics.StartPipelineReprocessingOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotanalyticsTagResourceResult struct {
	Result workflow.Future
}

func (r *IotanalyticsTagResourceResult) Get(ctx workflow.Context) (*iotanalytics.TagResourceOutput, error) {
    var output iotanalytics.TagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotanalyticsUntagResourceResult struct {
	Result workflow.Future
}

func (r *IotanalyticsUntagResourceResult) Get(ctx workflow.Context) (*iotanalytics.UntagResourceOutput, error) {
    var output iotanalytics.UntagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotanalyticsUpdateChannelResult struct {
	Result workflow.Future
}

func (r *IotanalyticsUpdateChannelResult) Get(ctx workflow.Context) (*iotanalytics.UpdateChannelOutput, error) {
    var output iotanalytics.UpdateChannelOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotanalyticsUpdateDatasetResult struct {
	Result workflow.Future
}

func (r *IotanalyticsUpdateDatasetResult) Get(ctx workflow.Context) (*iotanalytics.UpdateDatasetOutput, error) {
    var output iotanalytics.UpdateDatasetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotanalyticsUpdateDatastoreResult struct {
	Result workflow.Future
}

func (r *IotanalyticsUpdateDatastoreResult) Get(ctx workflow.Context) (*iotanalytics.UpdateDatastoreOutput, error) {
    var output iotanalytics.UpdateDatastoreOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotanalyticsUpdatePipelineResult struct {
	Result workflow.Future
}

func (r *IotanalyticsUpdatePipelineResult) Get(ctx workflow.Context) (*iotanalytics.UpdatePipelineOutput, error) {
    var output iotanalytics.UpdatePipelineOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IoTAnalyticsStub struct {
    activities awsactivities.IoTAnalyticsActivities
}

func NewIoTAnalyticsStub() IoTAnalyticsClient {
    return &IoTAnalyticsStub{}
}

func (a *IoTAnalyticsStub) BatchPutMessage(ctx workflow.Context, input *iotanalytics.BatchPutMessageInput) (*iotanalytics.BatchPutMessageOutput, error) {
    var output iotanalytics.BatchPutMessageOutput
    err := workflow.ExecuteActivity(ctx, a.activities.BatchPutMessage, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTAnalyticsStub) BatchPutMessageAsync(ctx workflow.Context, input *iotanalytics.BatchPutMessageInput) *IotanalyticsBatchPutMessageResult {
    future := workflow.ExecuteActivity(ctx, a.activities.BatchPutMessage, input)
    return &IotanalyticsBatchPutMessageResult{Result: future}
}

func (a *IoTAnalyticsStub) CancelPipelineReprocessing(ctx workflow.Context, input *iotanalytics.CancelPipelineReprocessingInput) (*iotanalytics.CancelPipelineReprocessingOutput, error) {
    var output iotanalytics.CancelPipelineReprocessingOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CancelPipelineReprocessing, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTAnalyticsStub) CancelPipelineReprocessingAsync(ctx workflow.Context, input *iotanalytics.CancelPipelineReprocessingInput) *IotanalyticsCancelPipelineReprocessingResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CancelPipelineReprocessing, input)
    return &IotanalyticsCancelPipelineReprocessingResult{Result: future}
}

func (a *IoTAnalyticsStub) CreateChannel(ctx workflow.Context, input *iotanalytics.CreateChannelInput) (*iotanalytics.CreateChannelOutput, error) {
    var output iotanalytics.CreateChannelOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateChannel, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTAnalyticsStub) CreateChannelAsync(ctx workflow.Context, input *iotanalytics.CreateChannelInput) *IotanalyticsCreateChannelResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateChannel, input)
    return &IotanalyticsCreateChannelResult{Result: future}
}

func (a *IoTAnalyticsStub) CreateDataset(ctx workflow.Context, input *iotanalytics.CreateDatasetInput) (*iotanalytics.CreateDatasetOutput, error) {
    var output iotanalytics.CreateDatasetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateDataset, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTAnalyticsStub) CreateDatasetAsync(ctx workflow.Context, input *iotanalytics.CreateDatasetInput) *IotanalyticsCreateDatasetResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateDataset, input)
    return &IotanalyticsCreateDatasetResult{Result: future}
}

func (a *IoTAnalyticsStub) CreateDatasetContent(ctx workflow.Context, input *iotanalytics.CreateDatasetContentInput) (*iotanalytics.CreateDatasetContentOutput, error) {
    var output iotanalytics.CreateDatasetContentOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateDatasetContent, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTAnalyticsStub) CreateDatasetContentAsync(ctx workflow.Context, input *iotanalytics.CreateDatasetContentInput) *IotanalyticsCreateDatasetContentResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateDatasetContent, input)
    return &IotanalyticsCreateDatasetContentResult{Result: future}
}

func (a *IoTAnalyticsStub) CreateDatastore(ctx workflow.Context, input *iotanalytics.CreateDatastoreInput) (*iotanalytics.CreateDatastoreOutput, error) {
    var output iotanalytics.CreateDatastoreOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateDatastore, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTAnalyticsStub) CreateDatastoreAsync(ctx workflow.Context, input *iotanalytics.CreateDatastoreInput) *IotanalyticsCreateDatastoreResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateDatastore, input)
    return &IotanalyticsCreateDatastoreResult{Result: future}
}

func (a *IoTAnalyticsStub) CreatePipeline(ctx workflow.Context, input *iotanalytics.CreatePipelineInput) (*iotanalytics.CreatePipelineOutput, error) {
    var output iotanalytics.CreatePipelineOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreatePipeline, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTAnalyticsStub) CreatePipelineAsync(ctx workflow.Context, input *iotanalytics.CreatePipelineInput) *IotanalyticsCreatePipelineResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreatePipeline, input)
    return &IotanalyticsCreatePipelineResult{Result: future}
}

func (a *IoTAnalyticsStub) DeleteChannel(ctx workflow.Context, input *iotanalytics.DeleteChannelInput) (*iotanalytics.DeleteChannelOutput, error) {
    var output iotanalytics.DeleteChannelOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteChannel, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTAnalyticsStub) DeleteChannelAsync(ctx workflow.Context, input *iotanalytics.DeleteChannelInput) *IotanalyticsDeleteChannelResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteChannel, input)
    return &IotanalyticsDeleteChannelResult{Result: future}
}

func (a *IoTAnalyticsStub) DeleteDataset(ctx workflow.Context, input *iotanalytics.DeleteDatasetInput) (*iotanalytics.DeleteDatasetOutput, error) {
    var output iotanalytics.DeleteDatasetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteDataset, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTAnalyticsStub) DeleteDatasetAsync(ctx workflow.Context, input *iotanalytics.DeleteDatasetInput) *IotanalyticsDeleteDatasetResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteDataset, input)
    return &IotanalyticsDeleteDatasetResult{Result: future}
}

func (a *IoTAnalyticsStub) DeleteDatasetContent(ctx workflow.Context, input *iotanalytics.DeleteDatasetContentInput) (*iotanalytics.DeleteDatasetContentOutput, error) {
    var output iotanalytics.DeleteDatasetContentOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteDatasetContent, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTAnalyticsStub) DeleteDatasetContentAsync(ctx workflow.Context, input *iotanalytics.DeleteDatasetContentInput) *IotanalyticsDeleteDatasetContentResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteDatasetContent, input)
    return &IotanalyticsDeleteDatasetContentResult{Result: future}
}

func (a *IoTAnalyticsStub) DeleteDatastore(ctx workflow.Context, input *iotanalytics.DeleteDatastoreInput) (*iotanalytics.DeleteDatastoreOutput, error) {
    var output iotanalytics.DeleteDatastoreOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteDatastore, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTAnalyticsStub) DeleteDatastoreAsync(ctx workflow.Context, input *iotanalytics.DeleteDatastoreInput) *IotanalyticsDeleteDatastoreResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteDatastore, input)
    return &IotanalyticsDeleteDatastoreResult{Result: future}
}

func (a *IoTAnalyticsStub) DeletePipeline(ctx workflow.Context, input *iotanalytics.DeletePipelineInput) (*iotanalytics.DeletePipelineOutput, error) {
    var output iotanalytics.DeletePipelineOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeletePipeline, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTAnalyticsStub) DeletePipelineAsync(ctx workflow.Context, input *iotanalytics.DeletePipelineInput) *IotanalyticsDeletePipelineResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeletePipeline, input)
    return &IotanalyticsDeletePipelineResult{Result: future}
}

func (a *IoTAnalyticsStub) DescribeChannel(ctx workflow.Context, input *iotanalytics.DescribeChannelInput) (*iotanalytics.DescribeChannelOutput, error) {
    var output iotanalytics.DescribeChannelOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeChannel, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTAnalyticsStub) DescribeChannelAsync(ctx workflow.Context, input *iotanalytics.DescribeChannelInput) *IotanalyticsDescribeChannelResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeChannel, input)
    return &IotanalyticsDescribeChannelResult{Result: future}
}

func (a *IoTAnalyticsStub) DescribeDataset(ctx workflow.Context, input *iotanalytics.DescribeDatasetInput) (*iotanalytics.DescribeDatasetOutput, error) {
    var output iotanalytics.DescribeDatasetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeDataset, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTAnalyticsStub) DescribeDatasetAsync(ctx workflow.Context, input *iotanalytics.DescribeDatasetInput) *IotanalyticsDescribeDatasetResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeDataset, input)
    return &IotanalyticsDescribeDatasetResult{Result: future}
}

func (a *IoTAnalyticsStub) DescribeDatastore(ctx workflow.Context, input *iotanalytics.DescribeDatastoreInput) (*iotanalytics.DescribeDatastoreOutput, error) {
    var output iotanalytics.DescribeDatastoreOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeDatastore, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTAnalyticsStub) DescribeDatastoreAsync(ctx workflow.Context, input *iotanalytics.DescribeDatastoreInput) *IotanalyticsDescribeDatastoreResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeDatastore, input)
    return &IotanalyticsDescribeDatastoreResult{Result: future}
}

func (a *IoTAnalyticsStub) DescribeLoggingOptions(ctx workflow.Context, input *iotanalytics.DescribeLoggingOptionsInput) (*iotanalytics.DescribeLoggingOptionsOutput, error) {
    var output iotanalytics.DescribeLoggingOptionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeLoggingOptions, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTAnalyticsStub) DescribeLoggingOptionsAsync(ctx workflow.Context, input *iotanalytics.DescribeLoggingOptionsInput) *IotanalyticsDescribeLoggingOptionsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeLoggingOptions, input)
    return &IotanalyticsDescribeLoggingOptionsResult{Result: future}
}

func (a *IoTAnalyticsStub) DescribePipeline(ctx workflow.Context, input *iotanalytics.DescribePipelineInput) (*iotanalytics.DescribePipelineOutput, error) {
    var output iotanalytics.DescribePipelineOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribePipeline, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTAnalyticsStub) DescribePipelineAsync(ctx workflow.Context, input *iotanalytics.DescribePipelineInput) *IotanalyticsDescribePipelineResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribePipeline, input)
    return &IotanalyticsDescribePipelineResult{Result: future}
}

func (a *IoTAnalyticsStub) GetDatasetContent(ctx workflow.Context, input *iotanalytics.GetDatasetContentInput) (*iotanalytics.GetDatasetContentOutput, error) {
    var output iotanalytics.GetDatasetContentOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetDatasetContent, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTAnalyticsStub) GetDatasetContentAsync(ctx workflow.Context, input *iotanalytics.GetDatasetContentInput) *IotanalyticsGetDatasetContentResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetDatasetContent, input)
    return &IotanalyticsGetDatasetContentResult{Result: future}
}

func (a *IoTAnalyticsStub) ListChannels(ctx workflow.Context, input *iotanalytics.ListChannelsInput) (*iotanalytics.ListChannelsOutput, error) {
    var output iotanalytics.ListChannelsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListChannels, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTAnalyticsStub) ListChannelsAsync(ctx workflow.Context, input *iotanalytics.ListChannelsInput) *IotanalyticsListChannelsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListChannels, input)
    return &IotanalyticsListChannelsResult{Result: future}
}

func (a *IoTAnalyticsStub) ListDatasetContents(ctx workflow.Context, input *iotanalytics.ListDatasetContentsInput) (*iotanalytics.ListDatasetContentsOutput, error) {
    var output iotanalytics.ListDatasetContentsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListDatasetContents, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTAnalyticsStub) ListDatasetContentsAsync(ctx workflow.Context, input *iotanalytics.ListDatasetContentsInput) *IotanalyticsListDatasetContentsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListDatasetContents, input)
    return &IotanalyticsListDatasetContentsResult{Result: future}
}

func (a *IoTAnalyticsStub) ListDatasets(ctx workflow.Context, input *iotanalytics.ListDatasetsInput) (*iotanalytics.ListDatasetsOutput, error) {
    var output iotanalytics.ListDatasetsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListDatasets, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTAnalyticsStub) ListDatasetsAsync(ctx workflow.Context, input *iotanalytics.ListDatasetsInput) *IotanalyticsListDatasetsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListDatasets, input)
    return &IotanalyticsListDatasetsResult{Result: future}
}

func (a *IoTAnalyticsStub) ListDatastores(ctx workflow.Context, input *iotanalytics.ListDatastoresInput) (*iotanalytics.ListDatastoresOutput, error) {
    var output iotanalytics.ListDatastoresOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListDatastores, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTAnalyticsStub) ListDatastoresAsync(ctx workflow.Context, input *iotanalytics.ListDatastoresInput) *IotanalyticsListDatastoresResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListDatastores, input)
    return &IotanalyticsListDatastoresResult{Result: future}
}

func (a *IoTAnalyticsStub) ListPipelines(ctx workflow.Context, input *iotanalytics.ListPipelinesInput) (*iotanalytics.ListPipelinesOutput, error) {
    var output iotanalytics.ListPipelinesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListPipelines, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTAnalyticsStub) ListPipelinesAsync(ctx workflow.Context, input *iotanalytics.ListPipelinesInput) *IotanalyticsListPipelinesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListPipelines, input)
    return &IotanalyticsListPipelinesResult{Result: future}
}

func (a *IoTAnalyticsStub) ListTagsForResource(ctx workflow.Context, input *iotanalytics.ListTagsForResourceInput) (*iotanalytics.ListTagsForResourceOutput, error) {
    var output iotanalytics.ListTagsForResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListTagsForResource, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTAnalyticsStub) ListTagsForResourceAsync(ctx workflow.Context, input *iotanalytics.ListTagsForResourceInput) *IotanalyticsListTagsForResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListTagsForResource, input)
    return &IotanalyticsListTagsForResourceResult{Result: future}
}

func (a *IoTAnalyticsStub) PutLoggingOptions(ctx workflow.Context, input *iotanalytics.PutLoggingOptionsInput) (*iotanalytics.PutLoggingOptionsOutput, error) {
    var output iotanalytics.PutLoggingOptionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutLoggingOptions, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTAnalyticsStub) PutLoggingOptionsAsync(ctx workflow.Context, input *iotanalytics.PutLoggingOptionsInput) *IotanalyticsPutLoggingOptionsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.PutLoggingOptions, input)
    return &IotanalyticsPutLoggingOptionsResult{Result: future}
}

func (a *IoTAnalyticsStub) RunPipelineActivity(ctx workflow.Context, input *iotanalytics.RunPipelineActivityInput) (*iotanalytics.RunPipelineActivityOutput, error) {
    var output iotanalytics.RunPipelineActivityOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RunPipelineActivity, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTAnalyticsStub) RunPipelineActivityAsync(ctx workflow.Context, input *iotanalytics.RunPipelineActivityInput) *IotanalyticsRunPipelineActivityResult {
    future := workflow.ExecuteActivity(ctx, a.activities.RunPipelineActivity, input)
    return &IotanalyticsRunPipelineActivityResult{Result: future}
}

func (a *IoTAnalyticsStub) SampleChannelData(ctx workflow.Context, input *iotanalytics.SampleChannelDataInput) (*iotanalytics.SampleChannelDataOutput, error) {
    var output iotanalytics.SampleChannelDataOutput
    err := workflow.ExecuteActivity(ctx, a.activities.SampleChannelData, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTAnalyticsStub) SampleChannelDataAsync(ctx workflow.Context, input *iotanalytics.SampleChannelDataInput) *IotanalyticsSampleChannelDataResult {
    future := workflow.ExecuteActivity(ctx, a.activities.SampleChannelData, input)
    return &IotanalyticsSampleChannelDataResult{Result: future}
}

func (a *IoTAnalyticsStub) StartPipelineReprocessing(ctx workflow.Context, input *iotanalytics.StartPipelineReprocessingInput) (*iotanalytics.StartPipelineReprocessingOutput, error) {
    var output iotanalytics.StartPipelineReprocessingOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StartPipelineReprocessing, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTAnalyticsStub) StartPipelineReprocessingAsync(ctx workflow.Context, input *iotanalytics.StartPipelineReprocessingInput) *IotanalyticsStartPipelineReprocessingResult {
    future := workflow.ExecuteActivity(ctx, a.activities.StartPipelineReprocessing, input)
    return &IotanalyticsStartPipelineReprocessingResult{Result: future}
}

func (a *IoTAnalyticsStub) TagResource(ctx workflow.Context, input *iotanalytics.TagResourceInput) (*iotanalytics.TagResourceOutput, error) {
    var output iotanalytics.TagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.TagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTAnalyticsStub) TagResourceAsync(ctx workflow.Context, input *iotanalytics.TagResourceInput) *IotanalyticsTagResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.TagResource, input)
    return &IotanalyticsTagResourceResult{Result: future}
}

func (a *IoTAnalyticsStub) UntagResource(ctx workflow.Context, input *iotanalytics.UntagResourceInput) (*iotanalytics.UntagResourceOutput, error) {
    var output iotanalytics.UntagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTAnalyticsStub) UntagResourceAsync(ctx workflow.Context, input *iotanalytics.UntagResourceInput) *IotanalyticsUntagResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input)
    return &IotanalyticsUntagResourceResult{Result: future}
}

func (a *IoTAnalyticsStub) UpdateChannel(ctx workflow.Context, input *iotanalytics.UpdateChannelInput) (*iotanalytics.UpdateChannelOutput, error) {
    var output iotanalytics.UpdateChannelOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateChannel, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTAnalyticsStub) UpdateChannelAsync(ctx workflow.Context, input *iotanalytics.UpdateChannelInput) *IotanalyticsUpdateChannelResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateChannel, input)
    return &IotanalyticsUpdateChannelResult{Result: future}
}

func (a *IoTAnalyticsStub) UpdateDataset(ctx workflow.Context, input *iotanalytics.UpdateDatasetInput) (*iotanalytics.UpdateDatasetOutput, error) {
    var output iotanalytics.UpdateDatasetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateDataset, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTAnalyticsStub) UpdateDatasetAsync(ctx workflow.Context, input *iotanalytics.UpdateDatasetInput) *IotanalyticsUpdateDatasetResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateDataset, input)
    return &IotanalyticsUpdateDatasetResult{Result: future}
}

func (a *IoTAnalyticsStub) UpdateDatastore(ctx workflow.Context, input *iotanalytics.UpdateDatastoreInput) (*iotanalytics.UpdateDatastoreOutput, error) {
    var output iotanalytics.UpdateDatastoreOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateDatastore, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTAnalyticsStub) UpdateDatastoreAsync(ctx workflow.Context, input *iotanalytics.UpdateDatastoreInput) *IotanalyticsUpdateDatastoreResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateDatastore, input)
    return &IotanalyticsUpdateDatastoreResult{Result: future}
}

func (a *IoTAnalyticsStub) UpdatePipeline(ctx workflow.Context, input *iotanalytics.UpdatePipelineInput) (*iotanalytics.UpdatePipelineOutput, error) {
    var output iotanalytics.UpdatePipelineOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdatePipeline, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTAnalyticsStub) UpdatePipelineAsync(ctx workflow.Context, input *iotanalytics.UpdatePipelineInput) *IotanalyticsUpdatePipelineResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdatePipeline, input)
    return &IotanalyticsUpdatePipelineResult{Result: future}
}