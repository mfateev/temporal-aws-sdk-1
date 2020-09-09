package awsclients

import (
	"github.com/aws/aws-sdk-go/service/mediastore"
	"go.temporal.io/sdk/workflow"
	"temporal.io/aws-sdk/awsactivities"
)

type MediaStoreClient interface {
	CreateContainer(ctx workflow.Context, input *mediastore.CreateContainerInput) (*mediastore.CreateContainerOutput, error)
	CreateContainerAsync(ctx workflow.Context, input *mediastore.CreateContainerInput) *MediastoreCreateContainerResult

	DeleteContainer(ctx workflow.Context, input *mediastore.DeleteContainerInput) (*mediastore.DeleteContainerOutput, error)
	DeleteContainerAsync(ctx workflow.Context, input *mediastore.DeleteContainerInput) *MediastoreDeleteContainerResult

	DeleteContainerPolicy(ctx workflow.Context, input *mediastore.DeleteContainerPolicyInput) (*mediastore.DeleteContainerPolicyOutput, error)
	DeleteContainerPolicyAsync(ctx workflow.Context, input *mediastore.DeleteContainerPolicyInput) *MediastoreDeleteContainerPolicyResult

	DeleteCorsPolicy(ctx workflow.Context, input *mediastore.DeleteCorsPolicyInput) (*mediastore.DeleteCorsPolicyOutput, error)
	DeleteCorsPolicyAsync(ctx workflow.Context, input *mediastore.DeleteCorsPolicyInput) *MediastoreDeleteCorsPolicyResult

	DeleteLifecyclePolicy(ctx workflow.Context, input *mediastore.DeleteLifecyclePolicyInput) (*mediastore.DeleteLifecyclePolicyOutput, error)
	DeleteLifecyclePolicyAsync(ctx workflow.Context, input *mediastore.DeleteLifecyclePolicyInput) *MediastoreDeleteLifecyclePolicyResult

	DeleteMetricPolicy(ctx workflow.Context, input *mediastore.DeleteMetricPolicyInput) (*mediastore.DeleteMetricPolicyOutput, error)
	DeleteMetricPolicyAsync(ctx workflow.Context, input *mediastore.DeleteMetricPolicyInput) *MediastoreDeleteMetricPolicyResult

	DescribeContainer(ctx workflow.Context, input *mediastore.DescribeContainerInput) (*mediastore.DescribeContainerOutput, error)
	DescribeContainerAsync(ctx workflow.Context, input *mediastore.DescribeContainerInput) *MediastoreDescribeContainerResult

	GetContainerPolicy(ctx workflow.Context, input *mediastore.GetContainerPolicyInput) (*mediastore.GetContainerPolicyOutput, error)
	GetContainerPolicyAsync(ctx workflow.Context, input *mediastore.GetContainerPolicyInput) *MediastoreGetContainerPolicyResult

	GetCorsPolicy(ctx workflow.Context, input *mediastore.GetCorsPolicyInput) (*mediastore.GetCorsPolicyOutput, error)
	GetCorsPolicyAsync(ctx workflow.Context, input *mediastore.GetCorsPolicyInput) *MediastoreGetCorsPolicyResult

	GetLifecyclePolicy(ctx workflow.Context, input *mediastore.GetLifecyclePolicyInput) (*mediastore.GetLifecyclePolicyOutput, error)
	GetLifecyclePolicyAsync(ctx workflow.Context, input *mediastore.GetLifecyclePolicyInput) *MediastoreGetLifecyclePolicyResult

	GetMetricPolicy(ctx workflow.Context, input *mediastore.GetMetricPolicyInput) (*mediastore.GetMetricPolicyOutput, error)
	GetMetricPolicyAsync(ctx workflow.Context, input *mediastore.GetMetricPolicyInput) *MediastoreGetMetricPolicyResult

	ListContainers(ctx workflow.Context, input *mediastore.ListContainersInput) (*mediastore.ListContainersOutput, error)
	ListContainersAsync(ctx workflow.Context, input *mediastore.ListContainersInput) *MediastoreListContainersResult

	ListTagsForResource(ctx workflow.Context, input *mediastore.ListTagsForResourceInput) (*mediastore.ListTagsForResourceOutput, error)
	ListTagsForResourceAsync(ctx workflow.Context, input *mediastore.ListTagsForResourceInput) *MediastoreListTagsForResourceResult

	PutContainerPolicy(ctx workflow.Context, input *mediastore.PutContainerPolicyInput) (*mediastore.PutContainerPolicyOutput, error)
	PutContainerPolicyAsync(ctx workflow.Context, input *mediastore.PutContainerPolicyInput) *MediastorePutContainerPolicyResult

	PutCorsPolicy(ctx workflow.Context, input *mediastore.PutCorsPolicyInput) (*mediastore.PutCorsPolicyOutput, error)
	PutCorsPolicyAsync(ctx workflow.Context, input *mediastore.PutCorsPolicyInput) *MediastorePutCorsPolicyResult

	PutLifecyclePolicy(ctx workflow.Context, input *mediastore.PutLifecyclePolicyInput) (*mediastore.PutLifecyclePolicyOutput, error)
	PutLifecyclePolicyAsync(ctx workflow.Context, input *mediastore.PutLifecyclePolicyInput) *MediastorePutLifecyclePolicyResult

	PutMetricPolicy(ctx workflow.Context, input *mediastore.PutMetricPolicyInput) (*mediastore.PutMetricPolicyOutput, error)
	PutMetricPolicyAsync(ctx workflow.Context, input *mediastore.PutMetricPolicyInput) *MediastorePutMetricPolicyResult

	StartAccessLogging(ctx workflow.Context, input *mediastore.StartAccessLoggingInput) (*mediastore.StartAccessLoggingOutput, error)
	StartAccessLoggingAsync(ctx workflow.Context, input *mediastore.StartAccessLoggingInput) *MediastoreStartAccessLoggingResult

	StopAccessLogging(ctx workflow.Context, input *mediastore.StopAccessLoggingInput) (*mediastore.StopAccessLoggingOutput, error)
	StopAccessLoggingAsync(ctx workflow.Context, input *mediastore.StopAccessLoggingInput) *MediastoreStopAccessLoggingResult

	TagResource(ctx workflow.Context, input *mediastore.TagResourceInput) (*mediastore.TagResourceOutput, error)
	TagResourceAsync(ctx workflow.Context, input *mediastore.TagResourceInput) *MediastoreTagResourceResult

	UntagResource(ctx workflow.Context, input *mediastore.UntagResourceInput) (*mediastore.UntagResourceOutput, error)
	UntagResourceAsync(ctx workflow.Context, input *mediastore.UntagResourceInput) *MediastoreUntagResourceResult
}

type MediastoreCreateContainerResult struct {
	Result workflow.Future
}

func (r *MediastoreCreateContainerResult) Get(ctx workflow.Context) (*mediastore.CreateContainerOutput, error) {
	var output mediastore.CreateContainerOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type MediastoreDeleteContainerResult struct {
	Result workflow.Future
}

func (r *MediastoreDeleteContainerResult) Get(ctx workflow.Context) (*mediastore.DeleteContainerOutput, error) {
	var output mediastore.DeleteContainerOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type MediastoreDeleteContainerPolicyResult struct {
	Result workflow.Future
}

func (r *MediastoreDeleteContainerPolicyResult) Get(ctx workflow.Context) (*mediastore.DeleteContainerPolicyOutput, error) {
	var output mediastore.DeleteContainerPolicyOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type MediastoreDeleteCorsPolicyResult struct {
	Result workflow.Future
}

func (r *MediastoreDeleteCorsPolicyResult) Get(ctx workflow.Context) (*mediastore.DeleteCorsPolicyOutput, error) {
	var output mediastore.DeleteCorsPolicyOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type MediastoreDeleteLifecyclePolicyResult struct {
	Result workflow.Future
}

func (r *MediastoreDeleteLifecyclePolicyResult) Get(ctx workflow.Context) (*mediastore.DeleteLifecyclePolicyOutput, error) {
	var output mediastore.DeleteLifecyclePolicyOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type MediastoreDeleteMetricPolicyResult struct {
	Result workflow.Future
}

func (r *MediastoreDeleteMetricPolicyResult) Get(ctx workflow.Context) (*mediastore.DeleteMetricPolicyOutput, error) {
	var output mediastore.DeleteMetricPolicyOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type MediastoreDescribeContainerResult struct {
	Result workflow.Future
}

func (r *MediastoreDescribeContainerResult) Get(ctx workflow.Context) (*mediastore.DescribeContainerOutput, error) {
	var output mediastore.DescribeContainerOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type MediastoreGetContainerPolicyResult struct {
	Result workflow.Future
}

func (r *MediastoreGetContainerPolicyResult) Get(ctx workflow.Context) (*mediastore.GetContainerPolicyOutput, error) {
	var output mediastore.GetContainerPolicyOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type MediastoreGetCorsPolicyResult struct {
	Result workflow.Future
}

func (r *MediastoreGetCorsPolicyResult) Get(ctx workflow.Context) (*mediastore.GetCorsPolicyOutput, error) {
	var output mediastore.GetCorsPolicyOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type MediastoreGetLifecyclePolicyResult struct {
	Result workflow.Future
}

func (r *MediastoreGetLifecyclePolicyResult) Get(ctx workflow.Context) (*mediastore.GetLifecyclePolicyOutput, error) {
	var output mediastore.GetLifecyclePolicyOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type MediastoreGetMetricPolicyResult struct {
	Result workflow.Future
}

func (r *MediastoreGetMetricPolicyResult) Get(ctx workflow.Context) (*mediastore.GetMetricPolicyOutput, error) {
	var output mediastore.GetMetricPolicyOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type MediastoreListContainersResult struct {
	Result workflow.Future
}

func (r *MediastoreListContainersResult) Get(ctx workflow.Context) (*mediastore.ListContainersOutput, error) {
	var output mediastore.ListContainersOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type MediastoreListTagsForResourceResult struct {
	Result workflow.Future
}

func (r *MediastoreListTagsForResourceResult) Get(ctx workflow.Context) (*mediastore.ListTagsForResourceOutput, error) {
	var output mediastore.ListTagsForResourceOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type MediastorePutContainerPolicyResult struct {
	Result workflow.Future
}

func (r *MediastorePutContainerPolicyResult) Get(ctx workflow.Context) (*mediastore.PutContainerPolicyOutput, error) {
	var output mediastore.PutContainerPolicyOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type MediastorePutCorsPolicyResult struct {
	Result workflow.Future
}

func (r *MediastorePutCorsPolicyResult) Get(ctx workflow.Context) (*mediastore.PutCorsPolicyOutput, error) {
	var output mediastore.PutCorsPolicyOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type MediastorePutLifecyclePolicyResult struct {
	Result workflow.Future
}

func (r *MediastorePutLifecyclePolicyResult) Get(ctx workflow.Context) (*mediastore.PutLifecyclePolicyOutput, error) {
	var output mediastore.PutLifecyclePolicyOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type MediastorePutMetricPolicyResult struct {
	Result workflow.Future
}

func (r *MediastorePutMetricPolicyResult) Get(ctx workflow.Context) (*mediastore.PutMetricPolicyOutput, error) {
	var output mediastore.PutMetricPolicyOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type MediastoreStartAccessLoggingResult struct {
	Result workflow.Future
}

func (r *MediastoreStartAccessLoggingResult) Get(ctx workflow.Context) (*mediastore.StartAccessLoggingOutput, error) {
	var output mediastore.StartAccessLoggingOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type MediastoreStopAccessLoggingResult struct {
	Result workflow.Future
}

func (r *MediastoreStopAccessLoggingResult) Get(ctx workflow.Context) (*mediastore.StopAccessLoggingOutput, error) {
	var output mediastore.StopAccessLoggingOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type MediastoreTagResourceResult struct {
	Result workflow.Future
}

func (r *MediastoreTagResourceResult) Get(ctx workflow.Context) (*mediastore.TagResourceOutput, error) {
	var output mediastore.TagResourceOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type MediastoreUntagResourceResult struct {
	Result workflow.Future
}

func (r *MediastoreUntagResourceResult) Get(ctx workflow.Context) (*mediastore.UntagResourceOutput, error) {
	var output mediastore.UntagResourceOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type MediaStoreStub struct {
	activities awsactivities.MediaStoreActivities
}

func NewMediaStoreStub() MediaStoreClient {
	return &MediaStoreStub{}
}

func (a *MediaStoreStub) CreateContainer(ctx workflow.Context, input *mediastore.CreateContainerInput) (*mediastore.CreateContainerOutput, error) {
	var output mediastore.CreateContainerOutput
	err := workflow.ExecuteActivity(ctx, a.activities.CreateContainer, input).Get(ctx, &output)
	return &output, err
}

func (a *MediaStoreStub) CreateContainerAsync(ctx workflow.Context, input *mediastore.CreateContainerInput) *MediastoreCreateContainerResult {
	future := workflow.ExecuteActivity(ctx, a.activities.CreateContainer, input)
	return &MediastoreCreateContainerResult{Result: future}
}

func (a *MediaStoreStub) DeleteContainer(ctx workflow.Context, input *mediastore.DeleteContainerInput) (*mediastore.DeleteContainerOutput, error) {
	var output mediastore.DeleteContainerOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeleteContainer, input).Get(ctx, &output)
	return &output, err
}

func (a *MediaStoreStub) DeleteContainerAsync(ctx workflow.Context, input *mediastore.DeleteContainerInput) *MediastoreDeleteContainerResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeleteContainer, input)
	return &MediastoreDeleteContainerResult{Result: future}
}

func (a *MediaStoreStub) DeleteContainerPolicy(ctx workflow.Context, input *mediastore.DeleteContainerPolicyInput) (*mediastore.DeleteContainerPolicyOutput, error) {
	var output mediastore.DeleteContainerPolicyOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeleteContainerPolicy, input).Get(ctx, &output)
	return &output, err
}

func (a *MediaStoreStub) DeleteContainerPolicyAsync(ctx workflow.Context, input *mediastore.DeleteContainerPolicyInput) *MediastoreDeleteContainerPolicyResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeleteContainerPolicy, input)
	return &MediastoreDeleteContainerPolicyResult{Result: future}
}

func (a *MediaStoreStub) DeleteCorsPolicy(ctx workflow.Context, input *mediastore.DeleteCorsPolicyInput) (*mediastore.DeleteCorsPolicyOutput, error) {
	var output mediastore.DeleteCorsPolicyOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeleteCorsPolicy, input).Get(ctx, &output)
	return &output, err
}

func (a *MediaStoreStub) DeleteCorsPolicyAsync(ctx workflow.Context, input *mediastore.DeleteCorsPolicyInput) *MediastoreDeleteCorsPolicyResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeleteCorsPolicy, input)
	return &MediastoreDeleteCorsPolicyResult{Result: future}
}

func (a *MediaStoreStub) DeleteLifecyclePolicy(ctx workflow.Context, input *mediastore.DeleteLifecyclePolicyInput) (*mediastore.DeleteLifecyclePolicyOutput, error) {
	var output mediastore.DeleteLifecyclePolicyOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeleteLifecyclePolicy, input).Get(ctx, &output)
	return &output, err
}

func (a *MediaStoreStub) DeleteLifecyclePolicyAsync(ctx workflow.Context, input *mediastore.DeleteLifecyclePolicyInput) *MediastoreDeleteLifecyclePolicyResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeleteLifecyclePolicy, input)
	return &MediastoreDeleteLifecyclePolicyResult{Result: future}
}

func (a *MediaStoreStub) DeleteMetricPolicy(ctx workflow.Context, input *mediastore.DeleteMetricPolicyInput) (*mediastore.DeleteMetricPolicyOutput, error) {
	var output mediastore.DeleteMetricPolicyOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeleteMetricPolicy, input).Get(ctx, &output)
	return &output, err
}

func (a *MediaStoreStub) DeleteMetricPolicyAsync(ctx workflow.Context, input *mediastore.DeleteMetricPolicyInput) *MediastoreDeleteMetricPolicyResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeleteMetricPolicy, input)
	return &MediastoreDeleteMetricPolicyResult{Result: future}
}

func (a *MediaStoreStub) DescribeContainer(ctx workflow.Context, input *mediastore.DescribeContainerInput) (*mediastore.DescribeContainerOutput, error) {
	var output mediastore.DescribeContainerOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DescribeContainer, input).Get(ctx, &output)
	return &output, err
}

func (a *MediaStoreStub) DescribeContainerAsync(ctx workflow.Context, input *mediastore.DescribeContainerInput) *MediastoreDescribeContainerResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DescribeContainer, input)
	return &MediastoreDescribeContainerResult{Result: future}
}

func (a *MediaStoreStub) GetContainerPolicy(ctx workflow.Context, input *mediastore.GetContainerPolicyInput) (*mediastore.GetContainerPolicyOutput, error) {
	var output mediastore.GetContainerPolicyOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetContainerPolicy, input).Get(ctx, &output)
	return &output, err
}

func (a *MediaStoreStub) GetContainerPolicyAsync(ctx workflow.Context, input *mediastore.GetContainerPolicyInput) *MediastoreGetContainerPolicyResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetContainerPolicy, input)
	return &MediastoreGetContainerPolicyResult{Result: future}
}

func (a *MediaStoreStub) GetCorsPolicy(ctx workflow.Context, input *mediastore.GetCorsPolicyInput) (*mediastore.GetCorsPolicyOutput, error) {
	var output mediastore.GetCorsPolicyOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetCorsPolicy, input).Get(ctx, &output)
	return &output, err
}

func (a *MediaStoreStub) GetCorsPolicyAsync(ctx workflow.Context, input *mediastore.GetCorsPolicyInput) *MediastoreGetCorsPolicyResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetCorsPolicy, input)
	return &MediastoreGetCorsPolicyResult{Result: future}
}

func (a *MediaStoreStub) GetLifecyclePolicy(ctx workflow.Context, input *mediastore.GetLifecyclePolicyInput) (*mediastore.GetLifecyclePolicyOutput, error) {
	var output mediastore.GetLifecyclePolicyOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetLifecyclePolicy, input).Get(ctx, &output)
	return &output, err
}

func (a *MediaStoreStub) GetLifecyclePolicyAsync(ctx workflow.Context, input *mediastore.GetLifecyclePolicyInput) *MediastoreGetLifecyclePolicyResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetLifecyclePolicy, input)
	return &MediastoreGetLifecyclePolicyResult{Result: future}
}

func (a *MediaStoreStub) GetMetricPolicy(ctx workflow.Context, input *mediastore.GetMetricPolicyInput) (*mediastore.GetMetricPolicyOutput, error) {
	var output mediastore.GetMetricPolicyOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetMetricPolicy, input).Get(ctx, &output)
	return &output, err
}

func (a *MediaStoreStub) GetMetricPolicyAsync(ctx workflow.Context, input *mediastore.GetMetricPolicyInput) *MediastoreGetMetricPolicyResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetMetricPolicy, input)
	return &MediastoreGetMetricPolicyResult{Result: future}
}

func (a *MediaStoreStub) ListContainers(ctx workflow.Context, input *mediastore.ListContainersInput) (*mediastore.ListContainersOutput, error) {
	var output mediastore.ListContainersOutput
	err := workflow.ExecuteActivity(ctx, a.activities.ListContainers, input).Get(ctx, &output)
	return &output, err
}

func (a *MediaStoreStub) ListContainersAsync(ctx workflow.Context, input *mediastore.ListContainersInput) *MediastoreListContainersResult {
	future := workflow.ExecuteActivity(ctx, a.activities.ListContainers, input)
	return &MediastoreListContainersResult{Result: future}
}

func (a *MediaStoreStub) ListTagsForResource(ctx workflow.Context, input *mediastore.ListTagsForResourceInput) (*mediastore.ListTagsForResourceOutput, error) {
	var output mediastore.ListTagsForResourceOutput
	err := workflow.ExecuteActivity(ctx, a.activities.ListTagsForResource, input).Get(ctx, &output)
	return &output, err
}

func (a *MediaStoreStub) ListTagsForResourceAsync(ctx workflow.Context, input *mediastore.ListTagsForResourceInput) *MediastoreListTagsForResourceResult {
	future := workflow.ExecuteActivity(ctx, a.activities.ListTagsForResource, input)
	return &MediastoreListTagsForResourceResult{Result: future}
}

func (a *MediaStoreStub) PutContainerPolicy(ctx workflow.Context, input *mediastore.PutContainerPolicyInput) (*mediastore.PutContainerPolicyOutput, error) {
	var output mediastore.PutContainerPolicyOutput
	err := workflow.ExecuteActivity(ctx, a.activities.PutContainerPolicy, input).Get(ctx, &output)
	return &output, err
}

func (a *MediaStoreStub) PutContainerPolicyAsync(ctx workflow.Context, input *mediastore.PutContainerPolicyInput) *MediastorePutContainerPolicyResult {
	future := workflow.ExecuteActivity(ctx, a.activities.PutContainerPolicy, input)
	return &MediastorePutContainerPolicyResult{Result: future}
}

func (a *MediaStoreStub) PutCorsPolicy(ctx workflow.Context, input *mediastore.PutCorsPolicyInput) (*mediastore.PutCorsPolicyOutput, error) {
	var output mediastore.PutCorsPolicyOutput
	err := workflow.ExecuteActivity(ctx, a.activities.PutCorsPolicy, input).Get(ctx, &output)
	return &output, err
}

func (a *MediaStoreStub) PutCorsPolicyAsync(ctx workflow.Context, input *mediastore.PutCorsPolicyInput) *MediastorePutCorsPolicyResult {
	future := workflow.ExecuteActivity(ctx, a.activities.PutCorsPolicy, input)
	return &MediastorePutCorsPolicyResult{Result: future}
}

func (a *MediaStoreStub) PutLifecyclePolicy(ctx workflow.Context, input *mediastore.PutLifecyclePolicyInput) (*mediastore.PutLifecyclePolicyOutput, error) {
	var output mediastore.PutLifecyclePolicyOutput
	err := workflow.ExecuteActivity(ctx, a.activities.PutLifecyclePolicy, input).Get(ctx, &output)
	return &output, err
}

func (a *MediaStoreStub) PutLifecyclePolicyAsync(ctx workflow.Context, input *mediastore.PutLifecyclePolicyInput) *MediastorePutLifecyclePolicyResult {
	future := workflow.ExecuteActivity(ctx, a.activities.PutLifecyclePolicy, input)
	return &MediastorePutLifecyclePolicyResult{Result: future}
}

func (a *MediaStoreStub) PutMetricPolicy(ctx workflow.Context, input *mediastore.PutMetricPolicyInput) (*mediastore.PutMetricPolicyOutput, error) {
	var output mediastore.PutMetricPolicyOutput
	err := workflow.ExecuteActivity(ctx, a.activities.PutMetricPolicy, input).Get(ctx, &output)
	return &output, err
}

func (a *MediaStoreStub) PutMetricPolicyAsync(ctx workflow.Context, input *mediastore.PutMetricPolicyInput) *MediastorePutMetricPolicyResult {
	future := workflow.ExecuteActivity(ctx, a.activities.PutMetricPolicy, input)
	return &MediastorePutMetricPolicyResult{Result: future}
}

func (a *MediaStoreStub) StartAccessLogging(ctx workflow.Context, input *mediastore.StartAccessLoggingInput) (*mediastore.StartAccessLoggingOutput, error) {
	var output mediastore.StartAccessLoggingOutput
	err := workflow.ExecuteActivity(ctx, a.activities.StartAccessLogging, input).Get(ctx, &output)
	return &output, err
}

func (a *MediaStoreStub) StartAccessLoggingAsync(ctx workflow.Context, input *mediastore.StartAccessLoggingInput) *MediastoreStartAccessLoggingResult {
	future := workflow.ExecuteActivity(ctx, a.activities.StartAccessLogging, input)
	return &MediastoreStartAccessLoggingResult{Result: future}
}

func (a *MediaStoreStub) StopAccessLogging(ctx workflow.Context, input *mediastore.StopAccessLoggingInput) (*mediastore.StopAccessLoggingOutput, error) {
	var output mediastore.StopAccessLoggingOutput
	err := workflow.ExecuteActivity(ctx, a.activities.StopAccessLogging, input).Get(ctx, &output)
	return &output, err
}

func (a *MediaStoreStub) StopAccessLoggingAsync(ctx workflow.Context, input *mediastore.StopAccessLoggingInput) *MediastoreStopAccessLoggingResult {
	future := workflow.ExecuteActivity(ctx, a.activities.StopAccessLogging, input)
	return &MediastoreStopAccessLoggingResult{Result: future}
}

func (a *MediaStoreStub) TagResource(ctx workflow.Context, input *mediastore.TagResourceInput) (*mediastore.TagResourceOutput, error) {
	var output mediastore.TagResourceOutput
	err := workflow.ExecuteActivity(ctx, a.activities.TagResource, input).Get(ctx, &output)
	return &output, err
}

func (a *MediaStoreStub) TagResourceAsync(ctx workflow.Context, input *mediastore.TagResourceInput) *MediastoreTagResourceResult {
	future := workflow.ExecuteActivity(ctx, a.activities.TagResource, input)
	return &MediastoreTagResourceResult{Result: future}
}

func (a *MediaStoreStub) UntagResource(ctx workflow.Context, input *mediastore.UntagResourceInput) (*mediastore.UntagResourceOutput, error) {
	var output mediastore.UntagResourceOutput
	err := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input).Get(ctx, &output)
	return &output, err
}

func (a *MediaStoreStub) UntagResourceAsync(ctx workflow.Context, input *mediastore.UntagResourceInput) *MediastoreUntagResourceResult {
	future := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input)
	return &MediastoreUntagResourceResult{Result: future}
}
