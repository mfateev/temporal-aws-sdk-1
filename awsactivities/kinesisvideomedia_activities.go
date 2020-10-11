// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kinesisvideomedia"
	"github.com/aws/aws-sdk-go/service/kinesisvideomedia/kinesisvideomediaiface"
	"go.temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken

type _ request.Option

type KinesisVideoMediaActivities struct {
	client kinesisvideomediaiface.KinesisVideoMediaAPI

	sessionFactory SessionFactory
}

func NewKinesisVideoMediaActivities(sess *session.Session, config ...*aws.Config) *KinesisVideoMediaActivities {
	client := kinesisvideomedia.New(sess, config...)
	return &KinesisVideoMediaActivities{client: client}
}

func NewKinesisVideoMediaActivitiesWithSessionFactory(sessionFactory SessionFactory) *KinesisVideoMediaActivities {
	return &KinesisVideoMediaActivities{sessionFactory: sessionFactory}
}

func (a *KinesisVideoMediaActivities) getClient(ctx context.Context) (kinesisvideomediaiface.KinesisVideoMediaAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, err
	}

	return kinesisvideomedia.New(sess), nil
}

func (a *KinesisVideoMediaActivities) GetMedia(ctx context.Context, input *kinesisvideomedia.GetMediaInput) (*kinesisvideomedia.GetMediaOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetMediaWithContext(ctx, input)
}
