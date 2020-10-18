// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package autoscalingplans

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"

	"go.temporal.io/aws-sdk/internal"
	"github.com/aws/aws-sdk-go/service/autoscalingplans"
	"github.com/aws/aws-sdk-go/service/autoscalingplans/autoscalingplansiface"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken

type _ request.Option

// SessionFactory returns an aws.Session based on the activity context.
type SessionFactory interface {
	Session(ctx context.Context) (*session.Session, error)
}

type Activities struct {
	client autoscalingplansiface.AutoScalingPlansAPI

	sessionFactory SessionFactory
}

func NewActivities(sess *session.Session, config ...*aws.Config) *Activities {
	client := autoscalingplans.New(sess, config...)
	return &Activities{client: client}
}

func NewActivitiesWithSessionFactory(sessionFactory SessionFactory) *Activities {
	return &Activities{sessionFactory: sessionFactory}
}

func (a *Activities) getClient(ctx context.Context) (autoscalingplansiface.AutoScalingPlansAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}

	return autoscalingplans.New(sess), nil
}

func (a *Activities) CreateScalingPlan(ctx context.Context, input *autoscalingplans.CreateScalingPlanInput) (*autoscalingplans.CreateScalingPlanOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateScalingPlanWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteScalingPlan(ctx context.Context, input *autoscalingplans.DeleteScalingPlanInput) (*autoscalingplans.DeleteScalingPlanOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteScalingPlanWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeScalingPlanResources(ctx context.Context, input *autoscalingplans.DescribeScalingPlanResourcesInput) (*autoscalingplans.DescribeScalingPlanResourcesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeScalingPlanResourcesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeScalingPlans(ctx context.Context, input *autoscalingplans.DescribeScalingPlansInput) (*autoscalingplans.DescribeScalingPlansOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeScalingPlansWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetScalingPlanResourceForecastData(ctx context.Context, input *autoscalingplans.GetScalingPlanResourceForecastDataInput) (*autoscalingplans.GetScalingPlanResourceForecastDataOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetScalingPlanResourceForecastDataWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateScalingPlan(ctx context.Context, input *autoscalingplans.UpdateScalingPlanInput) (*autoscalingplans.UpdateScalingPlanOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateScalingPlanWithContext(ctx, input)

	return output, internal.EncodeError(err)
}