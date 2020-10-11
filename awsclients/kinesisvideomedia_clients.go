// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package awsclients

import (
	"github.com/aws/aws-sdk-go/service/kinesisvideomedia"
	"go.temporal.io/sdk/workflow"
)

type KinesisVideoMediaClient interface {
	GetMedia(ctx workflow.Context, input *kinesisvideomedia.GetMediaInput) (*kinesisvideomedia.GetMediaOutput, error)
	GetMediaAsync(ctx workflow.Context, input *kinesisvideomedia.GetMediaInput) *KinesisVideoMediaGetMediaFuture
}

type KinesisVideoMediaStub struct{}

func NewKinesisVideoMediaStub() KinesisVideoMediaClient {
	return &KinesisVideoMediaStub{}
}

type KinesisVideoMediaGetMediaFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *KinesisVideoMediaGetMediaFuture) Get(ctx workflow.Context) (*kinesisvideomedia.GetMediaOutput, error) {
	var output kinesisvideomedia.GetMediaOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *KinesisVideoMediaStub) GetMedia(ctx workflow.Context, input *kinesisvideomedia.GetMediaInput) (*kinesisvideomedia.GetMediaOutput, error) {
	var output kinesisvideomedia.GetMediaOutput
	err := workflow.ExecuteActivity(ctx, "aws.kinesisvideomedia.GetMedia", input).Get(ctx, &output)
	return &output, err
}

func (a *KinesisVideoMediaStub) GetMediaAsync(ctx workflow.Context, input *kinesisvideomedia.GetMediaInput) *KinesisVideoMediaGetMediaFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kinesisvideomedia.GetMedia", input)
	return &KinesisVideoMediaGetMediaFuture{Future: future}
}
