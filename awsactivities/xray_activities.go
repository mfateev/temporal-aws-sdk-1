// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/xray"
	"github.com/aws/aws-sdk-go/service/xray/xrayiface"
	"go.temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken

type _ request.Option

type XRayActivities struct {
	client xrayiface.XRayAPI

	sessionFactory SessionFactory
}

func NewXRayActivities(sess *session.Session, config ...*aws.Config) *XRayActivities {
	client := xray.New(sess, config...)
	return &XRayActivities{client: client}
}

func NewXRayActivitiesWithSessionFactory(sessionFactory SessionFactory) *XRayActivities {
	return &XRayActivities{sessionFactory: sessionFactory}
}

func (a *XRayActivities) getClient(ctx context.Context) (xrayiface.XRayAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, err
	}

	return xray.New(sess), nil
}

func (a *XRayActivities) BatchGetTraces(ctx context.Context, input *xray.BatchGetTracesInput) (*xray.BatchGetTracesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.BatchGetTracesWithContext(ctx, input)
}

func (a *XRayActivities) CreateGroup(ctx context.Context, input *xray.CreateGroupInput) (*xray.CreateGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateGroupWithContext(ctx, input)
}

func (a *XRayActivities) CreateSamplingRule(ctx context.Context, input *xray.CreateSamplingRuleInput) (*xray.CreateSamplingRuleOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateSamplingRuleWithContext(ctx, input)
}

func (a *XRayActivities) DeleteGroup(ctx context.Context, input *xray.DeleteGroupInput) (*xray.DeleteGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteGroupWithContext(ctx, input)
}

func (a *XRayActivities) DeleteSamplingRule(ctx context.Context, input *xray.DeleteSamplingRuleInput) (*xray.DeleteSamplingRuleOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteSamplingRuleWithContext(ctx, input)
}

func (a *XRayActivities) GetEncryptionConfig(ctx context.Context, input *xray.GetEncryptionConfigInput) (*xray.GetEncryptionConfigOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetEncryptionConfigWithContext(ctx, input)
}

func (a *XRayActivities) GetGroup(ctx context.Context, input *xray.GetGroupInput) (*xray.GetGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetGroupWithContext(ctx, input)
}

func (a *XRayActivities) GetGroups(ctx context.Context, input *xray.GetGroupsInput) (*xray.GetGroupsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetGroupsWithContext(ctx, input)
}

func (a *XRayActivities) GetSamplingRules(ctx context.Context, input *xray.GetSamplingRulesInput) (*xray.GetSamplingRulesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetSamplingRulesWithContext(ctx, input)
}

func (a *XRayActivities) GetSamplingStatisticSummaries(ctx context.Context, input *xray.GetSamplingStatisticSummariesInput) (*xray.GetSamplingStatisticSummariesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetSamplingStatisticSummariesWithContext(ctx, input)
}

func (a *XRayActivities) GetSamplingTargets(ctx context.Context, input *xray.GetSamplingTargetsInput) (*xray.GetSamplingTargetsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetSamplingTargetsWithContext(ctx, input)
}

func (a *XRayActivities) GetServiceGraph(ctx context.Context, input *xray.GetServiceGraphInput) (*xray.GetServiceGraphOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetServiceGraphWithContext(ctx, input)
}

func (a *XRayActivities) GetTimeSeriesServiceStatistics(ctx context.Context, input *xray.GetTimeSeriesServiceStatisticsInput) (*xray.GetTimeSeriesServiceStatisticsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetTimeSeriesServiceStatisticsWithContext(ctx, input)
}

func (a *XRayActivities) GetTraceGraph(ctx context.Context, input *xray.GetTraceGraphInput) (*xray.GetTraceGraphOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetTraceGraphWithContext(ctx, input)
}

func (a *XRayActivities) GetTraceSummaries(ctx context.Context, input *xray.GetTraceSummariesInput) (*xray.GetTraceSummariesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetTraceSummariesWithContext(ctx, input)
}

func (a *XRayActivities) ListTagsForResource(ctx context.Context, input *xray.ListTagsForResourceInput) (*xray.ListTagsForResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListTagsForResourceWithContext(ctx, input)
}

func (a *XRayActivities) PutEncryptionConfig(ctx context.Context, input *xray.PutEncryptionConfigInput) (*xray.PutEncryptionConfigOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.PutEncryptionConfigWithContext(ctx, input)
}

func (a *XRayActivities) PutTelemetryRecords(ctx context.Context, input *xray.PutTelemetryRecordsInput) (*xray.PutTelemetryRecordsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.PutTelemetryRecordsWithContext(ctx, input)
}

func (a *XRayActivities) PutTraceSegments(ctx context.Context, input *xray.PutTraceSegmentsInput) (*xray.PutTraceSegmentsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.PutTraceSegmentsWithContext(ctx, input)
}

func (a *XRayActivities) TagResource(ctx context.Context, input *xray.TagResourceInput) (*xray.TagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.TagResourceWithContext(ctx, input)
}

func (a *XRayActivities) UntagResource(ctx context.Context, input *xray.UntagResourceInput) (*xray.UntagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UntagResourceWithContext(ctx, input)
}

func (a *XRayActivities) UpdateGroup(ctx context.Context, input *xray.UpdateGroupInput) (*xray.UpdateGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateGroupWithContext(ctx, input)
}

func (a *XRayActivities) UpdateSamplingRule(ctx context.Context, input *xray.UpdateSamplingRuleInput) (*xray.UpdateSamplingRuleOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateSamplingRuleWithContext(ctx, input)
}

