// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package awsclients

import (
	"github.com/aws/aws-sdk-go/service/s3outposts"
	"go.temporal.io/sdk/workflow"
)

type S3OutpostsClient interface {
	CreateEndpoint(ctx workflow.Context, input *s3outposts.CreateEndpointInput) (*s3outposts.CreateEndpointOutput, error)
	CreateEndpointAsync(ctx workflow.Context, input *s3outposts.CreateEndpointInput) *S3outpostsCreateEndpointResult

	DeleteEndpoint(ctx workflow.Context, input *s3outposts.DeleteEndpointInput) (*s3outposts.DeleteEndpointOutput, error)
	DeleteEndpointAsync(ctx workflow.Context, input *s3outposts.DeleteEndpointInput) *S3outpostsDeleteEndpointResult

	ListEndpoints(ctx workflow.Context, input *s3outposts.ListEndpointsInput) (*s3outposts.ListEndpointsOutput, error)
	ListEndpointsAsync(ctx workflow.Context, input *s3outposts.ListEndpointsInput) *S3outpostsListEndpointsResult
}

type S3OutpostsStub struct{}

func NewS3OutpostsStub() S3OutpostsClient {
	return &S3OutpostsStub{}
}

type S3outpostsCreateEndpointResult struct {
	Result workflow.Future
}

func (r *S3outpostsCreateEndpointResult) Get(ctx workflow.Context) (*s3outposts.CreateEndpointOutput, error) {
	var output s3outposts.CreateEndpointOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type S3outpostsDeleteEndpointResult struct {
	Result workflow.Future
}

func (r *S3outpostsDeleteEndpointResult) Get(ctx workflow.Context) (*s3outposts.DeleteEndpointOutput, error) {
	var output s3outposts.DeleteEndpointOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type S3outpostsListEndpointsResult struct {
	Result workflow.Future
}

func (r *S3outpostsListEndpointsResult) Get(ctx workflow.Context) (*s3outposts.ListEndpointsOutput, error) {
	var output s3outposts.ListEndpointsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

func (a *S3OutpostsStub) CreateEndpoint(ctx workflow.Context, input *s3outposts.CreateEndpointInput) (*s3outposts.CreateEndpointOutput, error) {
	var output s3outposts.CreateEndpointOutput
	err := workflow.ExecuteActivity(ctx, "aws.s3outposts.CreateEndpoint", input).Get(ctx, &output)
	return &output, err
}

func (a *S3OutpostsStub) CreateEndpointAsync(ctx workflow.Context, input *s3outposts.CreateEndpointInput) *S3outpostsCreateEndpointResult {
	future := workflow.ExecuteActivity(ctx, "aws.s3outposts.CreateEndpoint", input)
	return &S3outpostsCreateEndpointResult{Result: future}
}

func (a *S3OutpostsStub) DeleteEndpoint(ctx workflow.Context, input *s3outposts.DeleteEndpointInput) (*s3outposts.DeleteEndpointOutput, error) {
	var output s3outposts.DeleteEndpointOutput
	err := workflow.ExecuteActivity(ctx, "aws.s3outposts.DeleteEndpoint", input).Get(ctx, &output)
	return &output, err
}

func (a *S3OutpostsStub) DeleteEndpointAsync(ctx workflow.Context, input *s3outposts.DeleteEndpointInput) *S3outpostsDeleteEndpointResult {
	future := workflow.ExecuteActivity(ctx, "aws.s3outposts.DeleteEndpoint", input)
	return &S3outpostsDeleteEndpointResult{Result: future}
}

func (a *S3OutpostsStub) ListEndpoints(ctx workflow.Context, input *s3outposts.ListEndpointsInput) (*s3outposts.ListEndpointsOutput, error) {
	var output s3outposts.ListEndpointsOutput
	err := workflow.ExecuteActivity(ctx, "aws.s3outposts.ListEndpoints", input).Get(ctx, &output)
	return &output, err
}

func (a *S3OutpostsStub) ListEndpointsAsync(ctx workflow.Context, input *s3outposts.ListEndpointsInput) *S3outpostsListEndpointsResult {
	future := workflow.ExecuteActivity(ctx, "aws.s3outposts.ListEndpoints", input)
	return &S3outpostsListEndpointsResult{Result: future}
}
