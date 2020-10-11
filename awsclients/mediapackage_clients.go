// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package awsclients

import (
	"github.com/aws/aws-sdk-go/service/mediapackage"
	"go.temporal.io/sdk/workflow"
)

type MediaPackageClient interface {
	ConfigureLogs(ctx workflow.Context, input *mediapackage.ConfigureLogsInput) (*mediapackage.ConfigureLogsOutput, error)
	ConfigureLogsAsync(ctx workflow.Context, input *mediapackage.ConfigureLogsInput) *MediapackageConfigureLogsResult

	CreateChannel(ctx workflow.Context, input *mediapackage.CreateChannelInput) (*mediapackage.CreateChannelOutput, error)
	CreateChannelAsync(ctx workflow.Context, input *mediapackage.CreateChannelInput) *MediapackageCreateChannelResult

	CreateHarvestJob(ctx workflow.Context, input *mediapackage.CreateHarvestJobInput) (*mediapackage.CreateHarvestJobOutput, error)
	CreateHarvestJobAsync(ctx workflow.Context, input *mediapackage.CreateHarvestJobInput) *MediapackageCreateHarvestJobResult

	CreateOriginEndpoint(ctx workflow.Context, input *mediapackage.CreateOriginEndpointInput) (*mediapackage.CreateOriginEndpointOutput, error)
	CreateOriginEndpointAsync(ctx workflow.Context, input *mediapackage.CreateOriginEndpointInput) *MediapackageCreateOriginEndpointResult

	DeleteChannel(ctx workflow.Context, input *mediapackage.DeleteChannelInput) (*mediapackage.DeleteChannelOutput, error)
	DeleteChannelAsync(ctx workflow.Context, input *mediapackage.DeleteChannelInput) *MediapackageDeleteChannelResult

	DeleteOriginEndpoint(ctx workflow.Context, input *mediapackage.DeleteOriginEndpointInput) (*mediapackage.DeleteOriginEndpointOutput, error)
	DeleteOriginEndpointAsync(ctx workflow.Context, input *mediapackage.DeleteOriginEndpointInput) *MediapackageDeleteOriginEndpointResult

	DescribeChannel(ctx workflow.Context, input *mediapackage.DescribeChannelInput) (*mediapackage.DescribeChannelOutput, error)
	DescribeChannelAsync(ctx workflow.Context, input *mediapackage.DescribeChannelInput) *MediapackageDescribeChannelResult

	DescribeHarvestJob(ctx workflow.Context, input *mediapackage.DescribeHarvestJobInput) (*mediapackage.DescribeHarvestJobOutput, error)
	DescribeHarvestJobAsync(ctx workflow.Context, input *mediapackage.DescribeHarvestJobInput) *MediapackageDescribeHarvestJobResult

	DescribeOriginEndpoint(ctx workflow.Context, input *mediapackage.DescribeOriginEndpointInput) (*mediapackage.DescribeOriginEndpointOutput, error)
	DescribeOriginEndpointAsync(ctx workflow.Context, input *mediapackage.DescribeOriginEndpointInput) *MediapackageDescribeOriginEndpointResult

	ListChannels(ctx workflow.Context, input *mediapackage.ListChannelsInput) (*mediapackage.ListChannelsOutput, error)
	ListChannelsAsync(ctx workflow.Context, input *mediapackage.ListChannelsInput) *MediapackageListChannelsResult

	ListHarvestJobs(ctx workflow.Context, input *mediapackage.ListHarvestJobsInput) (*mediapackage.ListHarvestJobsOutput, error)
	ListHarvestJobsAsync(ctx workflow.Context, input *mediapackage.ListHarvestJobsInput) *MediapackageListHarvestJobsResult

	ListOriginEndpoints(ctx workflow.Context, input *mediapackage.ListOriginEndpointsInput) (*mediapackage.ListOriginEndpointsOutput, error)
	ListOriginEndpointsAsync(ctx workflow.Context, input *mediapackage.ListOriginEndpointsInput) *MediapackageListOriginEndpointsResult

	ListTagsForResource(ctx workflow.Context, input *mediapackage.ListTagsForResourceInput) (*mediapackage.ListTagsForResourceOutput, error)
	ListTagsForResourceAsync(ctx workflow.Context, input *mediapackage.ListTagsForResourceInput) *MediapackageListTagsForResourceResult

	RotateChannelCredentials(ctx workflow.Context, input *mediapackage.RotateChannelCredentialsInput) (*mediapackage.RotateChannelCredentialsOutput, error)
	RotateChannelCredentialsAsync(ctx workflow.Context, input *mediapackage.RotateChannelCredentialsInput) *MediapackageRotateChannelCredentialsResult

	RotateIngestEndpointCredentials(ctx workflow.Context, input *mediapackage.RotateIngestEndpointCredentialsInput) (*mediapackage.RotateIngestEndpointCredentialsOutput, error)
	RotateIngestEndpointCredentialsAsync(ctx workflow.Context, input *mediapackage.RotateIngestEndpointCredentialsInput) *MediapackageRotateIngestEndpointCredentialsResult

	TagResource(ctx workflow.Context, input *mediapackage.TagResourceInput) (*mediapackage.TagResourceOutput, error)
	TagResourceAsync(ctx workflow.Context, input *mediapackage.TagResourceInput) *MediapackageTagResourceResult

	UntagResource(ctx workflow.Context, input *mediapackage.UntagResourceInput) (*mediapackage.UntagResourceOutput, error)
	UntagResourceAsync(ctx workflow.Context, input *mediapackage.UntagResourceInput) *MediapackageUntagResourceResult

	UpdateChannel(ctx workflow.Context, input *mediapackage.UpdateChannelInput) (*mediapackage.UpdateChannelOutput, error)
	UpdateChannelAsync(ctx workflow.Context, input *mediapackage.UpdateChannelInput) *MediapackageUpdateChannelResult

	UpdateOriginEndpoint(ctx workflow.Context, input *mediapackage.UpdateOriginEndpointInput) (*mediapackage.UpdateOriginEndpointOutput, error)
	UpdateOriginEndpointAsync(ctx workflow.Context, input *mediapackage.UpdateOriginEndpointInput) *MediapackageUpdateOriginEndpointResult
}

type MediaPackageStub struct{}

func NewMediaPackageStub() MediaPackageClient {
	return &MediaPackageStub{}
}

type MediapackageConfigureLogsResult struct {
	Result workflow.Future
}

func (r *MediapackageConfigureLogsResult) Get(ctx workflow.Context) (*mediapackage.ConfigureLogsOutput, error) {
	var output mediapackage.ConfigureLogsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type MediapackageCreateChannelResult struct {
	Result workflow.Future
}

func (r *MediapackageCreateChannelResult) Get(ctx workflow.Context) (*mediapackage.CreateChannelOutput, error) {
	var output mediapackage.CreateChannelOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type MediapackageCreateHarvestJobResult struct {
	Result workflow.Future
}

func (r *MediapackageCreateHarvestJobResult) Get(ctx workflow.Context) (*mediapackage.CreateHarvestJobOutput, error) {
	var output mediapackage.CreateHarvestJobOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type MediapackageCreateOriginEndpointResult struct {
	Result workflow.Future
}

func (r *MediapackageCreateOriginEndpointResult) Get(ctx workflow.Context) (*mediapackage.CreateOriginEndpointOutput, error) {
	var output mediapackage.CreateOriginEndpointOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type MediapackageDeleteChannelResult struct {
	Result workflow.Future
}

func (r *MediapackageDeleteChannelResult) Get(ctx workflow.Context) (*mediapackage.DeleteChannelOutput, error) {
	var output mediapackage.DeleteChannelOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type MediapackageDeleteOriginEndpointResult struct {
	Result workflow.Future
}

func (r *MediapackageDeleteOriginEndpointResult) Get(ctx workflow.Context) (*mediapackage.DeleteOriginEndpointOutput, error) {
	var output mediapackage.DeleteOriginEndpointOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type MediapackageDescribeChannelResult struct {
	Result workflow.Future
}

func (r *MediapackageDescribeChannelResult) Get(ctx workflow.Context) (*mediapackage.DescribeChannelOutput, error) {
	var output mediapackage.DescribeChannelOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type MediapackageDescribeHarvestJobResult struct {
	Result workflow.Future
}

func (r *MediapackageDescribeHarvestJobResult) Get(ctx workflow.Context) (*mediapackage.DescribeHarvestJobOutput, error) {
	var output mediapackage.DescribeHarvestJobOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type MediapackageDescribeOriginEndpointResult struct {
	Result workflow.Future
}

func (r *MediapackageDescribeOriginEndpointResult) Get(ctx workflow.Context) (*mediapackage.DescribeOriginEndpointOutput, error) {
	var output mediapackage.DescribeOriginEndpointOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type MediapackageListChannelsResult struct {
	Result workflow.Future
}

func (r *MediapackageListChannelsResult) Get(ctx workflow.Context) (*mediapackage.ListChannelsOutput, error) {
	var output mediapackage.ListChannelsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type MediapackageListHarvestJobsResult struct {
	Result workflow.Future
}

func (r *MediapackageListHarvestJobsResult) Get(ctx workflow.Context) (*mediapackage.ListHarvestJobsOutput, error) {
	var output mediapackage.ListHarvestJobsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type MediapackageListOriginEndpointsResult struct {
	Result workflow.Future
}

func (r *MediapackageListOriginEndpointsResult) Get(ctx workflow.Context) (*mediapackage.ListOriginEndpointsOutput, error) {
	var output mediapackage.ListOriginEndpointsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type MediapackageListTagsForResourceResult struct {
	Result workflow.Future
}

func (r *MediapackageListTagsForResourceResult) Get(ctx workflow.Context) (*mediapackage.ListTagsForResourceOutput, error) {
	var output mediapackage.ListTagsForResourceOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type MediapackageRotateChannelCredentialsResult struct {
	Result workflow.Future
}

func (r *MediapackageRotateChannelCredentialsResult) Get(ctx workflow.Context) (*mediapackage.RotateChannelCredentialsOutput, error) {
	var output mediapackage.RotateChannelCredentialsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type MediapackageRotateIngestEndpointCredentialsResult struct {
	Result workflow.Future
}

func (r *MediapackageRotateIngestEndpointCredentialsResult) Get(ctx workflow.Context) (*mediapackage.RotateIngestEndpointCredentialsOutput, error) {
	var output mediapackage.RotateIngestEndpointCredentialsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type MediapackageTagResourceResult struct {
	Result workflow.Future
}

func (r *MediapackageTagResourceResult) Get(ctx workflow.Context) (*mediapackage.TagResourceOutput, error) {
	var output mediapackage.TagResourceOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type MediapackageUntagResourceResult struct {
	Result workflow.Future
}

func (r *MediapackageUntagResourceResult) Get(ctx workflow.Context) (*mediapackage.UntagResourceOutput, error) {
	var output mediapackage.UntagResourceOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type MediapackageUpdateChannelResult struct {
	Result workflow.Future
}

func (r *MediapackageUpdateChannelResult) Get(ctx workflow.Context) (*mediapackage.UpdateChannelOutput, error) {
	var output mediapackage.UpdateChannelOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type MediapackageUpdateOriginEndpointResult struct {
	Result workflow.Future
}

func (r *MediapackageUpdateOriginEndpointResult) Get(ctx workflow.Context) (*mediapackage.UpdateOriginEndpointOutput, error) {
	var output mediapackage.UpdateOriginEndpointOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

func (a *MediaPackageStub) ConfigureLogs(ctx workflow.Context, input *mediapackage.ConfigureLogsInput) (*mediapackage.ConfigureLogsOutput, error) {
	var output mediapackage.ConfigureLogsOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediapackage.ConfigureLogs", input).Get(ctx, &output)
	return &output, err
}

func (a *MediaPackageStub) ConfigureLogsAsync(ctx workflow.Context, input *mediapackage.ConfigureLogsInput) *MediapackageConfigureLogsResult {
	future := workflow.ExecuteActivity(ctx, "aws.mediapackage.ConfigureLogs", input)
	return &MediapackageConfigureLogsResult{Result: future}
}

func (a *MediaPackageStub) CreateChannel(ctx workflow.Context, input *mediapackage.CreateChannelInput) (*mediapackage.CreateChannelOutput, error) {
	var output mediapackage.CreateChannelOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediapackage.CreateChannel", input).Get(ctx, &output)
	return &output, err
}

func (a *MediaPackageStub) CreateChannelAsync(ctx workflow.Context, input *mediapackage.CreateChannelInput) *MediapackageCreateChannelResult {
	future := workflow.ExecuteActivity(ctx, "aws.mediapackage.CreateChannel", input)
	return &MediapackageCreateChannelResult{Result: future}
}

func (a *MediaPackageStub) CreateHarvestJob(ctx workflow.Context, input *mediapackage.CreateHarvestJobInput) (*mediapackage.CreateHarvestJobOutput, error) {
	var output mediapackage.CreateHarvestJobOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediapackage.CreateHarvestJob", input).Get(ctx, &output)
	return &output, err
}

func (a *MediaPackageStub) CreateHarvestJobAsync(ctx workflow.Context, input *mediapackage.CreateHarvestJobInput) *MediapackageCreateHarvestJobResult {
	future := workflow.ExecuteActivity(ctx, "aws.mediapackage.CreateHarvestJob", input)
	return &MediapackageCreateHarvestJobResult{Result: future}
}

func (a *MediaPackageStub) CreateOriginEndpoint(ctx workflow.Context, input *mediapackage.CreateOriginEndpointInput) (*mediapackage.CreateOriginEndpointOutput, error) {
	var output mediapackage.CreateOriginEndpointOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediapackage.CreateOriginEndpoint", input).Get(ctx, &output)
	return &output, err
}

func (a *MediaPackageStub) CreateOriginEndpointAsync(ctx workflow.Context, input *mediapackage.CreateOriginEndpointInput) *MediapackageCreateOriginEndpointResult {
	future := workflow.ExecuteActivity(ctx, "aws.mediapackage.CreateOriginEndpoint", input)
	return &MediapackageCreateOriginEndpointResult{Result: future}
}

func (a *MediaPackageStub) DeleteChannel(ctx workflow.Context, input *mediapackage.DeleteChannelInput) (*mediapackage.DeleteChannelOutput, error) {
	var output mediapackage.DeleteChannelOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediapackage.DeleteChannel", input).Get(ctx, &output)
	return &output, err
}

func (a *MediaPackageStub) DeleteChannelAsync(ctx workflow.Context, input *mediapackage.DeleteChannelInput) *MediapackageDeleteChannelResult {
	future := workflow.ExecuteActivity(ctx, "aws.mediapackage.DeleteChannel", input)
	return &MediapackageDeleteChannelResult{Result: future}
}

func (a *MediaPackageStub) DeleteOriginEndpoint(ctx workflow.Context, input *mediapackage.DeleteOriginEndpointInput) (*mediapackage.DeleteOriginEndpointOutput, error) {
	var output mediapackage.DeleteOriginEndpointOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediapackage.DeleteOriginEndpoint", input).Get(ctx, &output)
	return &output, err
}

func (a *MediaPackageStub) DeleteOriginEndpointAsync(ctx workflow.Context, input *mediapackage.DeleteOriginEndpointInput) *MediapackageDeleteOriginEndpointResult {
	future := workflow.ExecuteActivity(ctx, "aws.mediapackage.DeleteOriginEndpoint", input)
	return &MediapackageDeleteOriginEndpointResult{Result: future}
}

func (a *MediaPackageStub) DescribeChannel(ctx workflow.Context, input *mediapackage.DescribeChannelInput) (*mediapackage.DescribeChannelOutput, error) {
	var output mediapackage.DescribeChannelOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediapackage.DescribeChannel", input).Get(ctx, &output)
	return &output, err
}

func (a *MediaPackageStub) DescribeChannelAsync(ctx workflow.Context, input *mediapackage.DescribeChannelInput) *MediapackageDescribeChannelResult {
	future := workflow.ExecuteActivity(ctx, "aws.mediapackage.DescribeChannel", input)
	return &MediapackageDescribeChannelResult{Result: future}
}

func (a *MediaPackageStub) DescribeHarvestJob(ctx workflow.Context, input *mediapackage.DescribeHarvestJobInput) (*mediapackage.DescribeHarvestJobOutput, error) {
	var output mediapackage.DescribeHarvestJobOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediapackage.DescribeHarvestJob", input).Get(ctx, &output)
	return &output, err
}

func (a *MediaPackageStub) DescribeHarvestJobAsync(ctx workflow.Context, input *mediapackage.DescribeHarvestJobInput) *MediapackageDescribeHarvestJobResult {
	future := workflow.ExecuteActivity(ctx, "aws.mediapackage.DescribeHarvestJob", input)
	return &MediapackageDescribeHarvestJobResult{Result: future}
}

func (a *MediaPackageStub) DescribeOriginEndpoint(ctx workflow.Context, input *mediapackage.DescribeOriginEndpointInput) (*mediapackage.DescribeOriginEndpointOutput, error) {
	var output mediapackage.DescribeOriginEndpointOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediapackage.DescribeOriginEndpoint", input).Get(ctx, &output)
	return &output, err
}

func (a *MediaPackageStub) DescribeOriginEndpointAsync(ctx workflow.Context, input *mediapackage.DescribeOriginEndpointInput) *MediapackageDescribeOriginEndpointResult {
	future := workflow.ExecuteActivity(ctx, "aws.mediapackage.DescribeOriginEndpoint", input)
	return &MediapackageDescribeOriginEndpointResult{Result: future}
}

func (a *MediaPackageStub) ListChannels(ctx workflow.Context, input *mediapackage.ListChannelsInput) (*mediapackage.ListChannelsOutput, error) {
	var output mediapackage.ListChannelsOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediapackage.ListChannels", input).Get(ctx, &output)
	return &output, err
}

func (a *MediaPackageStub) ListChannelsAsync(ctx workflow.Context, input *mediapackage.ListChannelsInput) *MediapackageListChannelsResult {
	future := workflow.ExecuteActivity(ctx, "aws.mediapackage.ListChannels", input)
	return &MediapackageListChannelsResult{Result: future}
}

func (a *MediaPackageStub) ListHarvestJobs(ctx workflow.Context, input *mediapackage.ListHarvestJobsInput) (*mediapackage.ListHarvestJobsOutput, error) {
	var output mediapackage.ListHarvestJobsOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediapackage.ListHarvestJobs", input).Get(ctx, &output)
	return &output, err
}

func (a *MediaPackageStub) ListHarvestJobsAsync(ctx workflow.Context, input *mediapackage.ListHarvestJobsInput) *MediapackageListHarvestJobsResult {
	future := workflow.ExecuteActivity(ctx, "aws.mediapackage.ListHarvestJobs", input)
	return &MediapackageListHarvestJobsResult{Result: future}
}

func (a *MediaPackageStub) ListOriginEndpoints(ctx workflow.Context, input *mediapackage.ListOriginEndpointsInput) (*mediapackage.ListOriginEndpointsOutput, error) {
	var output mediapackage.ListOriginEndpointsOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediapackage.ListOriginEndpoints", input).Get(ctx, &output)
	return &output, err
}

func (a *MediaPackageStub) ListOriginEndpointsAsync(ctx workflow.Context, input *mediapackage.ListOriginEndpointsInput) *MediapackageListOriginEndpointsResult {
	future := workflow.ExecuteActivity(ctx, "aws.mediapackage.ListOriginEndpoints", input)
	return &MediapackageListOriginEndpointsResult{Result: future}
}

func (a *MediaPackageStub) ListTagsForResource(ctx workflow.Context, input *mediapackage.ListTagsForResourceInput) (*mediapackage.ListTagsForResourceOutput, error) {
	var output mediapackage.ListTagsForResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediapackage.ListTagsForResource", input).Get(ctx, &output)
	return &output, err
}

func (a *MediaPackageStub) ListTagsForResourceAsync(ctx workflow.Context, input *mediapackage.ListTagsForResourceInput) *MediapackageListTagsForResourceResult {
	future := workflow.ExecuteActivity(ctx, "aws.mediapackage.ListTagsForResource", input)
	return &MediapackageListTagsForResourceResult{Result: future}
}

func (a *MediaPackageStub) RotateChannelCredentials(ctx workflow.Context, input *mediapackage.RotateChannelCredentialsInput) (*mediapackage.RotateChannelCredentialsOutput, error) {
	var output mediapackage.RotateChannelCredentialsOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediapackage.RotateChannelCredentials", input).Get(ctx, &output)
	return &output, err
}

func (a *MediaPackageStub) RotateChannelCredentialsAsync(ctx workflow.Context, input *mediapackage.RotateChannelCredentialsInput) *MediapackageRotateChannelCredentialsResult {
	future := workflow.ExecuteActivity(ctx, "aws.mediapackage.RotateChannelCredentials", input)
	return &MediapackageRotateChannelCredentialsResult{Result: future}
}

func (a *MediaPackageStub) RotateIngestEndpointCredentials(ctx workflow.Context, input *mediapackage.RotateIngestEndpointCredentialsInput) (*mediapackage.RotateIngestEndpointCredentialsOutput, error) {
	var output mediapackage.RotateIngestEndpointCredentialsOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediapackage.RotateIngestEndpointCredentials", input).Get(ctx, &output)
	return &output, err
}

func (a *MediaPackageStub) RotateIngestEndpointCredentialsAsync(ctx workflow.Context, input *mediapackage.RotateIngestEndpointCredentialsInput) *MediapackageRotateIngestEndpointCredentialsResult {
	future := workflow.ExecuteActivity(ctx, "aws.mediapackage.RotateIngestEndpointCredentials", input)
	return &MediapackageRotateIngestEndpointCredentialsResult{Result: future}
}

func (a *MediaPackageStub) TagResource(ctx workflow.Context, input *mediapackage.TagResourceInput) (*mediapackage.TagResourceOutput, error) {
	var output mediapackage.TagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediapackage.TagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *MediaPackageStub) TagResourceAsync(ctx workflow.Context, input *mediapackage.TagResourceInput) *MediapackageTagResourceResult {
	future := workflow.ExecuteActivity(ctx, "aws.mediapackage.TagResource", input)
	return &MediapackageTagResourceResult{Result: future}
}

func (a *MediaPackageStub) UntagResource(ctx workflow.Context, input *mediapackage.UntagResourceInput) (*mediapackage.UntagResourceOutput, error) {
	var output mediapackage.UntagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediapackage.UntagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *MediaPackageStub) UntagResourceAsync(ctx workflow.Context, input *mediapackage.UntagResourceInput) *MediapackageUntagResourceResult {
	future := workflow.ExecuteActivity(ctx, "aws.mediapackage.UntagResource", input)
	return &MediapackageUntagResourceResult{Result: future}
}

func (a *MediaPackageStub) UpdateChannel(ctx workflow.Context, input *mediapackage.UpdateChannelInput) (*mediapackage.UpdateChannelOutput, error) {
	var output mediapackage.UpdateChannelOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediapackage.UpdateChannel", input).Get(ctx, &output)
	return &output, err
}

func (a *MediaPackageStub) UpdateChannelAsync(ctx workflow.Context, input *mediapackage.UpdateChannelInput) *MediapackageUpdateChannelResult {
	future := workflow.ExecuteActivity(ctx, "aws.mediapackage.UpdateChannel", input)
	return &MediapackageUpdateChannelResult{Result: future}
}

func (a *MediaPackageStub) UpdateOriginEndpoint(ctx workflow.Context, input *mediapackage.UpdateOriginEndpointInput) (*mediapackage.UpdateOriginEndpointOutput, error) {
	var output mediapackage.UpdateOriginEndpointOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediapackage.UpdateOriginEndpoint", input).Get(ctx, &output)
	return &output, err
}

func (a *MediaPackageStub) UpdateOriginEndpointAsync(ctx workflow.Context, input *mediapackage.UpdateOriginEndpointInput) *MediapackageUpdateOriginEndpointResult {
	future := workflow.ExecuteActivity(ctx, "aws.mediapackage.UpdateOriginEndpoint", input)
	return &MediapackageUpdateOriginEndpointResult{Result: future}
}
