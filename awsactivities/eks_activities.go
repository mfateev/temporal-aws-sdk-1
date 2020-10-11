// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/eks"
	"github.com/aws/aws-sdk-go/service/eks/eksiface"
	"go.temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken

type _ request.Option

type EKSActivities struct {
	client eksiface.EKSAPI

	sessionFactory SessionFactory
}

func NewEKSActivities(sess *session.Session, config ...*aws.Config) *EKSActivities {
	client := eks.New(sess, config...)
	return &EKSActivities{client: client}
}

func NewEKSActivitiesWithSessionFactory(sessionFactory SessionFactory) *EKSActivities {
	return &EKSActivities{sessionFactory: sessionFactory}
}

func (a *EKSActivities) getClient(ctx context.Context) (eksiface.EKSAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, err
	}

	return eks.New(sess), nil
}

func (a *EKSActivities) CreateCluster(ctx context.Context, input *eks.CreateClusterInput) (*eks.CreateClusterOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateClusterWithContext(ctx, input)
}

func (a *EKSActivities) CreateFargateProfile(ctx context.Context, input *eks.CreateFargateProfileInput) (*eks.CreateFargateProfileOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateFargateProfileWithContext(ctx, input)
}

func (a *EKSActivities) CreateNodegroup(ctx context.Context, input *eks.CreateNodegroupInput) (*eks.CreateNodegroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateNodegroupWithContext(ctx, input)
}

func (a *EKSActivities) DeleteCluster(ctx context.Context, input *eks.DeleteClusterInput) (*eks.DeleteClusterOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteClusterWithContext(ctx, input)
}

func (a *EKSActivities) DeleteFargateProfile(ctx context.Context, input *eks.DeleteFargateProfileInput) (*eks.DeleteFargateProfileOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteFargateProfileWithContext(ctx, input)
}

func (a *EKSActivities) DeleteNodegroup(ctx context.Context, input *eks.DeleteNodegroupInput) (*eks.DeleteNodegroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteNodegroupWithContext(ctx, input)
}

func (a *EKSActivities) DescribeCluster(ctx context.Context, input *eks.DescribeClusterInput) (*eks.DescribeClusterOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeClusterWithContext(ctx, input)
}

func (a *EKSActivities) DescribeFargateProfile(ctx context.Context, input *eks.DescribeFargateProfileInput) (*eks.DescribeFargateProfileOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeFargateProfileWithContext(ctx, input)
}

func (a *EKSActivities) DescribeNodegroup(ctx context.Context, input *eks.DescribeNodegroupInput) (*eks.DescribeNodegroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeNodegroupWithContext(ctx, input)
}

func (a *EKSActivities) DescribeUpdate(ctx context.Context, input *eks.DescribeUpdateInput) (*eks.DescribeUpdateOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeUpdateWithContext(ctx, input)
}

func (a *EKSActivities) ListClusters(ctx context.Context, input *eks.ListClustersInput) (*eks.ListClustersOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListClustersWithContext(ctx, input)
}

func (a *EKSActivities) ListFargateProfiles(ctx context.Context, input *eks.ListFargateProfilesInput) (*eks.ListFargateProfilesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListFargateProfilesWithContext(ctx, input)
}

func (a *EKSActivities) ListNodegroups(ctx context.Context, input *eks.ListNodegroupsInput) (*eks.ListNodegroupsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListNodegroupsWithContext(ctx, input)
}

func (a *EKSActivities) ListTagsForResource(ctx context.Context, input *eks.ListTagsForResourceInput) (*eks.ListTagsForResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListTagsForResourceWithContext(ctx, input)
}

func (a *EKSActivities) ListUpdates(ctx context.Context, input *eks.ListUpdatesInput) (*eks.ListUpdatesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListUpdatesWithContext(ctx, input)
}

func (a *EKSActivities) TagResource(ctx context.Context, input *eks.TagResourceInput) (*eks.TagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.TagResourceWithContext(ctx, input)
}

func (a *EKSActivities) UntagResource(ctx context.Context, input *eks.UntagResourceInput) (*eks.UntagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UntagResourceWithContext(ctx, input)
}

func (a *EKSActivities) UpdateClusterConfig(ctx context.Context, input *eks.UpdateClusterConfigInput) (*eks.UpdateClusterConfigOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateClusterConfigWithContext(ctx, input)
}

func (a *EKSActivities) UpdateClusterVersion(ctx context.Context, input *eks.UpdateClusterVersionInput) (*eks.UpdateClusterVersionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateClusterVersionWithContext(ctx, input)
}

func (a *EKSActivities) UpdateNodegroupConfig(ctx context.Context, input *eks.UpdateNodegroupConfigInput) (*eks.UpdateNodegroupConfigOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateNodegroupConfigWithContext(ctx, input)
}

func (a *EKSActivities) UpdateNodegroupVersion(ctx context.Context, input *eks.UpdateNodegroupVersionInput) (*eks.UpdateNodegroupVersionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateNodegroupVersionWithContext(ctx, input)
}

func (a *EKSActivities) WaitUntilClusterActive(ctx context.Context, input *eks.DescribeClusterInput) error {
	client, err := a.getClient(ctx)
	if err != nil {
		return err
	}
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return client.WaitUntilClusterActiveWithContext(ctx, input, options...)
	})
}

func (a *EKSActivities) WaitUntilClusterDeleted(ctx context.Context, input *eks.DescribeClusterInput) error {
	client, err := a.getClient(ctx)
	if err != nil {
		return err
	}
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return client.WaitUntilClusterDeletedWithContext(ctx, input, options...)
	})
}

func (a *EKSActivities) WaitUntilNodegroupActive(ctx context.Context, input *eks.DescribeNodegroupInput) error {
	client, err := a.getClient(ctx)
	if err != nil {
		return err
	}
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return client.WaitUntilNodegroupActiveWithContext(ctx, input, options...)
	})
}

func (a *EKSActivities) WaitUntilNodegroupDeleted(ctx context.Context, input *eks.DescribeNodegroupInput) error {
	client, err := a.getClient(ctx)
	if err != nil {
		return err
	}
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return client.WaitUntilNodegroupDeletedWithContext(ctx, input, options...)
	})
}
