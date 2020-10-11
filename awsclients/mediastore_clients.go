// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package awsclients

import (
	"github.com/aws/aws-sdk-go/service/mediastore"
	"go.temporal.io/sdk/workflow"
)

type MediaStoreClient interface {
	CreateContainer(ctx workflow.Context, input *mediastore.CreateContainerInput) (*mediastore.CreateContainerOutput, error)
	CreateContainerAsync(ctx workflow.Context, input *mediastore.CreateContainerInput) *MediaStoreCreateContainerFuture

	DeleteContainer(ctx workflow.Context, input *mediastore.DeleteContainerInput) (*mediastore.DeleteContainerOutput, error)
	DeleteContainerAsync(ctx workflow.Context, input *mediastore.DeleteContainerInput) *MediaStoreDeleteContainerFuture

	DeleteContainerPolicy(ctx workflow.Context, input *mediastore.DeleteContainerPolicyInput) (*mediastore.DeleteContainerPolicyOutput, error)
	DeleteContainerPolicyAsync(ctx workflow.Context, input *mediastore.DeleteContainerPolicyInput) *MediaStoreDeleteContainerPolicyFuture

	DeleteCorsPolicy(ctx workflow.Context, input *mediastore.DeleteCorsPolicyInput) (*mediastore.DeleteCorsPolicyOutput, error)
	DeleteCorsPolicyAsync(ctx workflow.Context, input *mediastore.DeleteCorsPolicyInput) *MediaStoreDeleteCorsPolicyFuture

	DeleteLifecyclePolicy(ctx workflow.Context, input *mediastore.DeleteLifecyclePolicyInput) (*mediastore.DeleteLifecyclePolicyOutput, error)
	DeleteLifecyclePolicyAsync(ctx workflow.Context, input *mediastore.DeleteLifecyclePolicyInput) *MediaStoreDeleteLifecyclePolicyFuture

	DeleteMetricPolicy(ctx workflow.Context, input *mediastore.DeleteMetricPolicyInput) (*mediastore.DeleteMetricPolicyOutput, error)
	DeleteMetricPolicyAsync(ctx workflow.Context, input *mediastore.DeleteMetricPolicyInput) *MediaStoreDeleteMetricPolicyFuture

	DescribeContainer(ctx workflow.Context, input *mediastore.DescribeContainerInput) (*mediastore.DescribeContainerOutput, error)
	DescribeContainerAsync(ctx workflow.Context, input *mediastore.DescribeContainerInput) *MediaStoreDescribeContainerFuture

	GetContainerPolicy(ctx workflow.Context, input *mediastore.GetContainerPolicyInput) (*mediastore.GetContainerPolicyOutput, error)
	GetContainerPolicyAsync(ctx workflow.Context, input *mediastore.GetContainerPolicyInput) *MediaStoreGetContainerPolicyFuture

	GetCorsPolicy(ctx workflow.Context, input *mediastore.GetCorsPolicyInput) (*mediastore.GetCorsPolicyOutput, error)
	GetCorsPolicyAsync(ctx workflow.Context, input *mediastore.GetCorsPolicyInput) *MediaStoreGetCorsPolicyFuture

	GetLifecyclePolicy(ctx workflow.Context, input *mediastore.GetLifecyclePolicyInput) (*mediastore.GetLifecyclePolicyOutput, error)
	GetLifecyclePolicyAsync(ctx workflow.Context, input *mediastore.GetLifecyclePolicyInput) *MediaStoreGetLifecyclePolicyFuture

	GetMetricPolicy(ctx workflow.Context, input *mediastore.GetMetricPolicyInput) (*mediastore.GetMetricPolicyOutput, error)
	GetMetricPolicyAsync(ctx workflow.Context, input *mediastore.GetMetricPolicyInput) *MediaStoreGetMetricPolicyFuture

	ListContainers(ctx workflow.Context, input *mediastore.ListContainersInput) (*mediastore.ListContainersOutput, error)
	ListContainersAsync(ctx workflow.Context, input *mediastore.ListContainersInput) *MediaStoreListContainersFuture

	ListTagsForResource(ctx workflow.Context, input *mediastore.ListTagsForResourceInput) (*mediastore.ListTagsForResourceOutput, error)
	ListTagsForResourceAsync(ctx workflow.Context, input *mediastore.ListTagsForResourceInput) *MediaStoreListTagsForResourceFuture

	PutContainerPolicy(ctx workflow.Context, input *mediastore.PutContainerPolicyInput) (*mediastore.PutContainerPolicyOutput, error)
	PutContainerPolicyAsync(ctx workflow.Context, input *mediastore.PutContainerPolicyInput) *MediaStorePutContainerPolicyFuture

	PutCorsPolicy(ctx workflow.Context, input *mediastore.PutCorsPolicyInput) (*mediastore.PutCorsPolicyOutput, error)
	PutCorsPolicyAsync(ctx workflow.Context, input *mediastore.PutCorsPolicyInput) *MediaStorePutCorsPolicyFuture

	PutLifecyclePolicy(ctx workflow.Context, input *mediastore.PutLifecyclePolicyInput) (*mediastore.PutLifecyclePolicyOutput, error)
	PutLifecyclePolicyAsync(ctx workflow.Context, input *mediastore.PutLifecyclePolicyInput) *MediaStorePutLifecyclePolicyFuture

	PutMetricPolicy(ctx workflow.Context, input *mediastore.PutMetricPolicyInput) (*mediastore.PutMetricPolicyOutput, error)
	PutMetricPolicyAsync(ctx workflow.Context, input *mediastore.PutMetricPolicyInput) *MediaStorePutMetricPolicyFuture

	StartAccessLogging(ctx workflow.Context, input *mediastore.StartAccessLoggingInput) (*mediastore.StartAccessLoggingOutput, error)
	StartAccessLoggingAsync(ctx workflow.Context, input *mediastore.StartAccessLoggingInput) *MediaStoreStartAccessLoggingFuture

	StopAccessLogging(ctx workflow.Context, input *mediastore.StopAccessLoggingInput) (*mediastore.StopAccessLoggingOutput, error)
	StopAccessLoggingAsync(ctx workflow.Context, input *mediastore.StopAccessLoggingInput) *MediaStoreStopAccessLoggingFuture

	TagResource(ctx workflow.Context, input *mediastore.TagResourceInput) (*mediastore.TagResourceOutput, error)
	TagResourceAsync(ctx workflow.Context, input *mediastore.TagResourceInput) *MediaStoreTagResourceFuture

	UntagResource(ctx workflow.Context, input *mediastore.UntagResourceInput) (*mediastore.UntagResourceOutput, error)
	UntagResourceAsync(ctx workflow.Context, input *mediastore.UntagResourceInput) *MediaStoreUntagResourceFuture
}

type MediaStoreStub struct{}

func NewMediaStoreStub() MediaStoreClient {
	return &MediaStoreStub{}
}

type MediaStoreCreateContainerFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MediaStoreCreateContainerFuture) Get(ctx workflow.Context) (*mediastore.CreateContainerOutput, error) {
	var output mediastore.CreateContainerOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MediaStoreDeleteContainerFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MediaStoreDeleteContainerFuture) Get(ctx workflow.Context) (*mediastore.DeleteContainerOutput, error) {
	var output mediastore.DeleteContainerOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MediaStoreDeleteContainerPolicyFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MediaStoreDeleteContainerPolicyFuture) Get(ctx workflow.Context) (*mediastore.DeleteContainerPolicyOutput, error) {
	var output mediastore.DeleteContainerPolicyOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MediaStoreDeleteCorsPolicyFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MediaStoreDeleteCorsPolicyFuture) Get(ctx workflow.Context) (*mediastore.DeleteCorsPolicyOutput, error) {
	var output mediastore.DeleteCorsPolicyOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MediaStoreDeleteLifecyclePolicyFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MediaStoreDeleteLifecyclePolicyFuture) Get(ctx workflow.Context) (*mediastore.DeleteLifecyclePolicyOutput, error) {
	var output mediastore.DeleteLifecyclePolicyOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MediaStoreDeleteMetricPolicyFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MediaStoreDeleteMetricPolicyFuture) Get(ctx workflow.Context) (*mediastore.DeleteMetricPolicyOutput, error) {
	var output mediastore.DeleteMetricPolicyOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MediaStoreDescribeContainerFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MediaStoreDescribeContainerFuture) Get(ctx workflow.Context) (*mediastore.DescribeContainerOutput, error) {
	var output mediastore.DescribeContainerOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MediaStoreGetContainerPolicyFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MediaStoreGetContainerPolicyFuture) Get(ctx workflow.Context) (*mediastore.GetContainerPolicyOutput, error) {
	var output mediastore.GetContainerPolicyOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MediaStoreGetCorsPolicyFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MediaStoreGetCorsPolicyFuture) Get(ctx workflow.Context) (*mediastore.GetCorsPolicyOutput, error) {
	var output mediastore.GetCorsPolicyOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MediaStoreGetLifecyclePolicyFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MediaStoreGetLifecyclePolicyFuture) Get(ctx workflow.Context) (*mediastore.GetLifecyclePolicyOutput, error) {
	var output mediastore.GetLifecyclePolicyOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MediaStoreGetMetricPolicyFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MediaStoreGetMetricPolicyFuture) Get(ctx workflow.Context) (*mediastore.GetMetricPolicyOutput, error) {
	var output mediastore.GetMetricPolicyOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MediaStoreListContainersFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MediaStoreListContainersFuture) Get(ctx workflow.Context) (*mediastore.ListContainersOutput, error) {
	var output mediastore.ListContainersOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MediaStoreListTagsForResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MediaStoreListTagsForResourceFuture) Get(ctx workflow.Context) (*mediastore.ListTagsForResourceOutput, error) {
	var output mediastore.ListTagsForResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MediaStorePutContainerPolicyFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MediaStorePutContainerPolicyFuture) Get(ctx workflow.Context) (*mediastore.PutContainerPolicyOutput, error) {
	var output mediastore.PutContainerPolicyOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MediaStorePutCorsPolicyFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MediaStorePutCorsPolicyFuture) Get(ctx workflow.Context) (*mediastore.PutCorsPolicyOutput, error) {
	var output mediastore.PutCorsPolicyOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MediaStorePutLifecyclePolicyFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MediaStorePutLifecyclePolicyFuture) Get(ctx workflow.Context) (*mediastore.PutLifecyclePolicyOutput, error) {
	var output mediastore.PutLifecyclePolicyOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MediaStorePutMetricPolicyFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MediaStorePutMetricPolicyFuture) Get(ctx workflow.Context) (*mediastore.PutMetricPolicyOutput, error) {
	var output mediastore.PutMetricPolicyOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MediaStoreStartAccessLoggingFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MediaStoreStartAccessLoggingFuture) Get(ctx workflow.Context) (*mediastore.StartAccessLoggingOutput, error) {
	var output mediastore.StartAccessLoggingOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MediaStoreStopAccessLoggingFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MediaStoreStopAccessLoggingFuture) Get(ctx workflow.Context) (*mediastore.StopAccessLoggingOutput, error) {
	var output mediastore.StopAccessLoggingOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MediaStoreTagResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MediaStoreTagResourceFuture) Get(ctx workflow.Context) (*mediastore.TagResourceOutput, error) {
	var output mediastore.TagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MediaStoreUntagResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MediaStoreUntagResourceFuture) Get(ctx workflow.Context) (*mediastore.UntagResourceOutput, error) {
	var output mediastore.UntagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *MediaStoreStub) CreateContainer(ctx workflow.Context, input *mediastore.CreateContainerInput) (*mediastore.CreateContainerOutput, error) {
	var output mediastore.CreateContainerOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediastore.CreateContainer", input).Get(ctx, &output)
	return &output, err
}

func (a *MediaStoreStub) CreateContainerAsync(ctx workflow.Context, input *mediastore.CreateContainerInput) *MediaStoreCreateContainerFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mediastore.CreateContainer", input)
	return &MediaStoreCreateContainerFuture{Future: future}
}

func (a *MediaStoreStub) DeleteContainer(ctx workflow.Context, input *mediastore.DeleteContainerInput) (*mediastore.DeleteContainerOutput, error) {
	var output mediastore.DeleteContainerOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediastore.DeleteContainer", input).Get(ctx, &output)
	return &output, err
}

func (a *MediaStoreStub) DeleteContainerAsync(ctx workflow.Context, input *mediastore.DeleteContainerInput) *MediaStoreDeleteContainerFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mediastore.DeleteContainer", input)
	return &MediaStoreDeleteContainerFuture{Future: future}
}

func (a *MediaStoreStub) DeleteContainerPolicy(ctx workflow.Context, input *mediastore.DeleteContainerPolicyInput) (*mediastore.DeleteContainerPolicyOutput, error) {
	var output mediastore.DeleteContainerPolicyOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediastore.DeleteContainerPolicy", input).Get(ctx, &output)
	return &output, err
}

func (a *MediaStoreStub) DeleteContainerPolicyAsync(ctx workflow.Context, input *mediastore.DeleteContainerPolicyInput) *MediaStoreDeleteContainerPolicyFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mediastore.DeleteContainerPolicy", input)
	return &MediaStoreDeleteContainerPolicyFuture{Future: future}
}

func (a *MediaStoreStub) DeleteCorsPolicy(ctx workflow.Context, input *mediastore.DeleteCorsPolicyInput) (*mediastore.DeleteCorsPolicyOutput, error) {
	var output mediastore.DeleteCorsPolicyOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediastore.DeleteCorsPolicy", input).Get(ctx, &output)
	return &output, err
}

func (a *MediaStoreStub) DeleteCorsPolicyAsync(ctx workflow.Context, input *mediastore.DeleteCorsPolicyInput) *MediaStoreDeleteCorsPolicyFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mediastore.DeleteCorsPolicy", input)
	return &MediaStoreDeleteCorsPolicyFuture{Future: future}
}

func (a *MediaStoreStub) DeleteLifecyclePolicy(ctx workflow.Context, input *mediastore.DeleteLifecyclePolicyInput) (*mediastore.DeleteLifecyclePolicyOutput, error) {
	var output mediastore.DeleteLifecyclePolicyOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediastore.DeleteLifecyclePolicy", input).Get(ctx, &output)
	return &output, err
}

func (a *MediaStoreStub) DeleteLifecyclePolicyAsync(ctx workflow.Context, input *mediastore.DeleteLifecyclePolicyInput) *MediaStoreDeleteLifecyclePolicyFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mediastore.DeleteLifecyclePolicy", input)
	return &MediaStoreDeleteLifecyclePolicyFuture{Future: future}
}

func (a *MediaStoreStub) DeleteMetricPolicy(ctx workflow.Context, input *mediastore.DeleteMetricPolicyInput) (*mediastore.DeleteMetricPolicyOutput, error) {
	var output mediastore.DeleteMetricPolicyOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediastore.DeleteMetricPolicy", input).Get(ctx, &output)
	return &output, err
}

func (a *MediaStoreStub) DeleteMetricPolicyAsync(ctx workflow.Context, input *mediastore.DeleteMetricPolicyInput) *MediaStoreDeleteMetricPolicyFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mediastore.DeleteMetricPolicy", input)
	return &MediaStoreDeleteMetricPolicyFuture{Future: future}
}

func (a *MediaStoreStub) DescribeContainer(ctx workflow.Context, input *mediastore.DescribeContainerInput) (*mediastore.DescribeContainerOutput, error) {
	var output mediastore.DescribeContainerOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediastore.DescribeContainer", input).Get(ctx, &output)
	return &output, err
}

func (a *MediaStoreStub) DescribeContainerAsync(ctx workflow.Context, input *mediastore.DescribeContainerInput) *MediaStoreDescribeContainerFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mediastore.DescribeContainer", input)
	return &MediaStoreDescribeContainerFuture{Future: future}
}

func (a *MediaStoreStub) GetContainerPolicy(ctx workflow.Context, input *mediastore.GetContainerPolicyInput) (*mediastore.GetContainerPolicyOutput, error) {
	var output mediastore.GetContainerPolicyOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediastore.GetContainerPolicy", input).Get(ctx, &output)
	return &output, err
}

func (a *MediaStoreStub) GetContainerPolicyAsync(ctx workflow.Context, input *mediastore.GetContainerPolicyInput) *MediaStoreGetContainerPolicyFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mediastore.GetContainerPolicy", input)
	return &MediaStoreGetContainerPolicyFuture{Future: future}
}

func (a *MediaStoreStub) GetCorsPolicy(ctx workflow.Context, input *mediastore.GetCorsPolicyInput) (*mediastore.GetCorsPolicyOutput, error) {
	var output mediastore.GetCorsPolicyOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediastore.GetCorsPolicy", input).Get(ctx, &output)
	return &output, err
}

func (a *MediaStoreStub) GetCorsPolicyAsync(ctx workflow.Context, input *mediastore.GetCorsPolicyInput) *MediaStoreGetCorsPolicyFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mediastore.GetCorsPolicy", input)
	return &MediaStoreGetCorsPolicyFuture{Future: future}
}

func (a *MediaStoreStub) GetLifecyclePolicy(ctx workflow.Context, input *mediastore.GetLifecyclePolicyInput) (*mediastore.GetLifecyclePolicyOutput, error) {
	var output mediastore.GetLifecyclePolicyOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediastore.GetLifecyclePolicy", input).Get(ctx, &output)
	return &output, err
}

func (a *MediaStoreStub) GetLifecyclePolicyAsync(ctx workflow.Context, input *mediastore.GetLifecyclePolicyInput) *MediaStoreGetLifecyclePolicyFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mediastore.GetLifecyclePolicy", input)
	return &MediaStoreGetLifecyclePolicyFuture{Future: future}
}

func (a *MediaStoreStub) GetMetricPolicy(ctx workflow.Context, input *mediastore.GetMetricPolicyInput) (*mediastore.GetMetricPolicyOutput, error) {
	var output mediastore.GetMetricPolicyOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediastore.GetMetricPolicy", input).Get(ctx, &output)
	return &output, err
}

func (a *MediaStoreStub) GetMetricPolicyAsync(ctx workflow.Context, input *mediastore.GetMetricPolicyInput) *MediaStoreGetMetricPolicyFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mediastore.GetMetricPolicy", input)
	return &MediaStoreGetMetricPolicyFuture{Future: future}
}

func (a *MediaStoreStub) ListContainers(ctx workflow.Context, input *mediastore.ListContainersInput) (*mediastore.ListContainersOutput, error) {
	var output mediastore.ListContainersOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediastore.ListContainers", input).Get(ctx, &output)
	return &output, err
}

func (a *MediaStoreStub) ListContainersAsync(ctx workflow.Context, input *mediastore.ListContainersInput) *MediaStoreListContainersFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mediastore.ListContainers", input)
	return &MediaStoreListContainersFuture{Future: future}
}

func (a *MediaStoreStub) ListTagsForResource(ctx workflow.Context, input *mediastore.ListTagsForResourceInput) (*mediastore.ListTagsForResourceOutput, error) {
	var output mediastore.ListTagsForResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediastore.ListTagsForResource", input).Get(ctx, &output)
	return &output, err
}

func (a *MediaStoreStub) ListTagsForResourceAsync(ctx workflow.Context, input *mediastore.ListTagsForResourceInput) *MediaStoreListTagsForResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mediastore.ListTagsForResource", input)
	return &MediaStoreListTagsForResourceFuture{Future: future}
}

func (a *MediaStoreStub) PutContainerPolicy(ctx workflow.Context, input *mediastore.PutContainerPolicyInput) (*mediastore.PutContainerPolicyOutput, error) {
	var output mediastore.PutContainerPolicyOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediastore.PutContainerPolicy", input).Get(ctx, &output)
	return &output, err
}

func (a *MediaStoreStub) PutContainerPolicyAsync(ctx workflow.Context, input *mediastore.PutContainerPolicyInput) *MediaStorePutContainerPolicyFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mediastore.PutContainerPolicy", input)
	return &MediaStorePutContainerPolicyFuture{Future: future}
}

func (a *MediaStoreStub) PutCorsPolicy(ctx workflow.Context, input *mediastore.PutCorsPolicyInput) (*mediastore.PutCorsPolicyOutput, error) {
	var output mediastore.PutCorsPolicyOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediastore.PutCorsPolicy", input).Get(ctx, &output)
	return &output, err
}

func (a *MediaStoreStub) PutCorsPolicyAsync(ctx workflow.Context, input *mediastore.PutCorsPolicyInput) *MediaStorePutCorsPolicyFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mediastore.PutCorsPolicy", input)
	return &MediaStorePutCorsPolicyFuture{Future: future}
}

func (a *MediaStoreStub) PutLifecyclePolicy(ctx workflow.Context, input *mediastore.PutLifecyclePolicyInput) (*mediastore.PutLifecyclePolicyOutput, error) {
	var output mediastore.PutLifecyclePolicyOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediastore.PutLifecyclePolicy", input).Get(ctx, &output)
	return &output, err
}

func (a *MediaStoreStub) PutLifecyclePolicyAsync(ctx workflow.Context, input *mediastore.PutLifecyclePolicyInput) *MediaStorePutLifecyclePolicyFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mediastore.PutLifecyclePolicy", input)
	return &MediaStorePutLifecyclePolicyFuture{Future: future}
}

func (a *MediaStoreStub) PutMetricPolicy(ctx workflow.Context, input *mediastore.PutMetricPolicyInput) (*mediastore.PutMetricPolicyOutput, error) {
	var output mediastore.PutMetricPolicyOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediastore.PutMetricPolicy", input).Get(ctx, &output)
	return &output, err
}

func (a *MediaStoreStub) PutMetricPolicyAsync(ctx workflow.Context, input *mediastore.PutMetricPolicyInput) *MediaStorePutMetricPolicyFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mediastore.PutMetricPolicy", input)
	return &MediaStorePutMetricPolicyFuture{Future: future}
}

func (a *MediaStoreStub) StartAccessLogging(ctx workflow.Context, input *mediastore.StartAccessLoggingInput) (*mediastore.StartAccessLoggingOutput, error) {
	var output mediastore.StartAccessLoggingOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediastore.StartAccessLogging", input).Get(ctx, &output)
	return &output, err
}

func (a *MediaStoreStub) StartAccessLoggingAsync(ctx workflow.Context, input *mediastore.StartAccessLoggingInput) *MediaStoreStartAccessLoggingFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mediastore.StartAccessLogging", input)
	return &MediaStoreStartAccessLoggingFuture{Future: future}
}

func (a *MediaStoreStub) StopAccessLogging(ctx workflow.Context, input *mediastore.StopAccessLoggingInput) (*mediastore.StopAccessLoggingOutput, error) {
	var output mediastore.StopAccessLoggingOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediastore.StopAccessLogging", input).Get(ctx, &output)
	return &output, err
}

func (a *MediaStoreStub) StopAccessLoggingAsync(ctx workflow.Context, input *mediastore.StopAccessLoggingInput) *MediaStoreStopAccessLoggingFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mediastore.StopAccessLogging", input)
	return &MediaStoreStopAccessLoggingFuture{Future: future}
}

func (a *MediaStoreStub) TagResource(ctx workflow.Context, input *mediastore.TagResourceInput) (*mediastore.TagResourceOutput, error) {
	var output mediastore.TagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediastore.TagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *MediaStoreStub) TagResourceAsync(ctx workflow.Context, input *mediastore.TagResourceInput) *MediaStoreTagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mediastore.TagResource", input)
	return &MediaStoreTagResourceFuture{Future: future}
}

func (a *MediaStoreStub) UntagResource(ctx workflow.Context, input *mediastore.UntagResourceInput) (*mediastore.UntagResourceOutput, error) {
	var output mediastore.UntagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediastore.UntagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *MediaStoreStub) UntagResourceAsync(ctx workflow.Context, input *mediastore.UntagResourceInput) *MediaStoreUntagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mediastore.UntagResource", input)
	return &MediaStoreUntagResourceFuture{Future: future}
}
