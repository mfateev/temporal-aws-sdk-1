// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/detective"
	"github.com/aws/aws-sdk-go/service/detective/detectiveiface"
	"go.temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken

type _ request.Option

type DetectiveActivities struct {
	client detectiveiface.DetectiveAPI

	sessionFactory SessionFactory
}

func NewDetectiveActivities(sess *session.Session, config ...*aws.Config) *DetectiveActivities {
	client := detective.New(sess, config...)
	return &DetectiveActivities{client: client}
}

func NewDetectiveActivitiesWithSessionFactory(sessionFactory SessionFactory) *DetectiveActivities {
	return &DetectiveActivities{sessionFactory: sessionFactory}
}

func (a *DetectiveActivities) getClient(ctx context.Context) (detectiveiface.DetectiveAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, err
	}

	return detective.New(sess), nil
}

func (a *DetectiveActivities) AcceptInvitation(ctx context.Context, input *detective.AcceptInvitationInput) (*detective.AcceptInvitationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.AcceptInvitationWithContext(ctx, input)
}

func (a *DetectiveActivities) CreateGraph(ctx context.Context, input *detective.CreateGraphInput) (*detective.CreateGraphOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateGraphWithContext(ctx, input)
}

func (a *DetectiveActivities) CreateMembers(ctx context.Context, input *detective.CreateMembersInput) (*detective.CreateMembersOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateMembersWithContext(ctx, input)
}

func (a *DetectiveActivities) DeleteGraph(ctx context.Context, input *detective.DeleteGraphInput) (*detective.DeleteGraphOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteGraphWithContext(ctx, input)
}

func (a *DetectiveActivities) DeleteMembers(ctx context.Context, input *detective.DeleteMembersInput) (*detective.DeleteMembersOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteMembersWithContext(ctx, input)
}

func (a *DetectiveActivities) DisassociateMembership(ctx context.Context, input *detective.DisassociateMembershipInput) (*detective.DisassociateMembershipOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DisassociateMembershipWithContext(ctx, input)
}

func (a *DetectiveActivities) GetMembers(ctx context.Context, input *detective.GetMembersInput) (*detective.GetMembersOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetMembersWithContext(ctx, input)
}

func (a *DetectiveActivities) ListGraphs(ctx context.Context, input *detective.ListGraphsInput) (*detective.ListGraphsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListGraphsWithContext(ctx, input)
}

func (a *DetectiveActivities) ListInvitations(ctx context.Context, input *detective.ListInvitationsInput) (*detective.ListInvitationsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListInvitationsWithContext(ctx, input)
}

func (a *DetectiveActivities) ListMembers(ctx context.Context, input *detective.ListMembersInput) (*detective.ListMembersOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListMembersWithContext(ctx, input)
}

func (a *DetectiveActivities) RejectInvitation(ctx context.Context, input *detective.RejectInvitationInput) (*detective.RejectInvitationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.RejectInvitationWithContext(ctx, input)
}

func (a *DetectiveActivities) StartMonitoringMember(ctx context.Context, input *detective.StartMonitoringMemberInput) (*detective.StartMonitoringMemberOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.StartMonitoringMemberWithContext(ctx, input)
}
