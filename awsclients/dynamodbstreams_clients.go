package awsclients

import (
	"github.com/aws/aws-sdk-go/service/dynamodbstreams"
	"go.temporal.io/sdk/workflow"
	"temporal.io/aws-sdk/awsactivities"
)

type DynamoDBStreamsClient interface {
       DescribeStream(ctx workflow.Context, input *dynamodbstreams.DescribeStreamInput) (*dynamodbstreams.DescribeStreamOutput, error)
       DescribeStreamAsync(ctx workflow.Context, input *dynamodbstreams.DescribeStreamInput) *DynamodbstreamsDescribeStreamResult

       GetRecords(ctx workflow.Context, input *dynamodbstreams.GetRecordsInput) (*dynamodbstreams.GetRecordsOutput, error)
       GetRecordsAsync(ctx workflow.Context, input *dynamodbstreams.GetRecordsInput) *DynamodbstreamsGetRecordsResult

       GetShardIterator(ctx workflow.Context, input *dynamodbstreams.GetShardIteratorInput) (*dynamodbstreams.GetShardIteratorOutput, error)
       GetShardIteratorAsync(ctx workflow.Context, input *dynamodbstreams.GetShardIteratorInput) *DynamodbstreamsGetShardIteratorResult

       ListStreams(ctx workflow.Context, input *dynamodbstreams.ListStreamsInput) (*dynamodbstreams.ListStreamsOutput, error)
       ListStreamsAsync(ctx workflow.Context, input *dynamodbstreams.ListStreamsInput) *DynamodbstreamsListStreamsResult
}

type DynamodbstreamsDescribeStreamResult struct {
	Result workflow.Future
}

func (r *DynamodbstreamsDescribeStreamResult) Get(ctx workflow.Context) (*dynamodbstreams.DescribeStreamOutput, error) {
    var output dynamodbstreams.DescribeStreamOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DynamodbstreamsGetRecordsResult struct {
	Result workflow.Future
}

func (r *DynamodbstreamsGetRecordsResult) Get(ctx workflow.Context) (*dynamodbstreams.GetRecordsOutput, error) {
    var output dynamodbstreams.GetRecordsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DynamodbstreamsGetShardIteratorResult struct {
	Result workflow.Future
}

func (r *DynamodbstreamsGetShardIteratorResult) Get(ctx workflow.Context) (*dynamodbstreams.GetShardIteratorOutput, error) {
    var output dynamodbstreams.GetShardIteratorOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DynamodbstreamsListStreamsResult struct {
	Result workflow.Future
}

func (r *DynamodbstreamsListStreamsResult) Get(ctx workflow.Context) (*dynamodbstreams.ListStreamsOutput, error) {
    var output dynamodbstreams.ListStreamsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DynamoDBStreamsStub struct {
    activities awsactivities.DynamoDBStreamsActivities
}

func NewDynamoDBStreamsStub() DynamoDBStreamsClient {
    return &DynamoDBStreamsStub{}
}

func (a *DynamoDBStreamsStub) DescribeStream(ctx workflow.Context, input *dynamodbstreams.DescribeStreamInput) (*dynamodbstreams.DescribeStreamOutput, error) {
    var output dynamodbstreams.DescribeStreamOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeStream, input).Get(ctx, &output)
    return &output, err
}

func (a *DynamoDBStreamsStub) DescribeStreamAsync(ctx workflow.Context, input *dynamodbstreams.DescribeStreamInput) *DynamodbstreamsDescribeStreamResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeStream, input)
    return &DynamodbstreamsDescribeStreamResult{Result: future}
}

func (a *DynamoDBStreamsStub) GetRecords(ctx workflow.Context, input *dynamodbstreams.GetRecordsInput) (*dynamodbstreams.GetRecordsOutput, error) {
    var output dynamodbstreams.GetRecordsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetRecords, input).Get(ctx, &output)
    return &output, err
}

func (a *DynamoDBStreamsStub) GetRecordsAsync(ctx workflow.Context, input *dynamodbstreams.GetRecordsInput) *DynamodbstreamsGetRecordsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetRecords, input)
    return &DynamodbstreamsGetRecordsResult{Result: future}
}

func (a *DynamoDBStreamsStub) GetShardIterator(ctx workflow.Context, input *dynamodbstreams.GetShardIteratorInput) (*dynamodbstreams.GetShardIteratorOutput, error) {
    var output dynamodbstreams.GetShardIteratorOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetShardIterator, input).Get(ctx, &output)
    return &output, err
}

func (a *DynamoDBStreamsStub) GetShardIteratorAsync(ctx workflow.Context, input *dynamodbstreams.GetShardIteratorInput) *DynamodbstreamsGetShardIteratorResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetShardIterator, input)
    return &DynamodbstreamsGetShardIteratorResult{Result: future}
}

func (a *DynamoDBStreamsStub) ListStreams(ctx workflow.Context, input *dynamodbstreams.ListStreamsInput) (*dynamodbstreams.ListStreamsOutput, error) {
    var output dynamodbstreams.ListStreamsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListStreams, input).Get(ctx, &output)
    return &output, err
}

func (a *DynamoDBStreamsStub) ListStreamsAsync(ctx workflow.Context, input *dynamodbstreams.ListStreamsInput) *DynamodbstreamsListStreamsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListStreams, input)
    return &DynamodbstreamsListStreamsResult{Result: future}
}