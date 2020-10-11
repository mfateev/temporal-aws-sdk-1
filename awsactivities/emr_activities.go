// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/emr"
	"github.com/aws/aws-sdk-go/service/emr/emriface"
	"go.temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken

type _ request.Option

type EMRActivities struct {
	client emriface.EMRAPI

	sessionFactory SessionFactory
}

func NewEMRActivities(sess *session.Session, config ...*aws.Config) *EMRActivities {
	client := emr.New(sess, config...)
	return &EMRActivities{client: client}
}

func NewEMRActivitiesWithSessionFactory(sessionFactory SessionFactory) *EMRActivities {
	return &EMRActivities{sessionFactory: sessionFactory}
}

func (a *EMRActivities) getClient(ctx context.Context) (emriface.EMRAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, err
	}

	return emr.New(sess), nil
}

func (a *EMRActivities) AddInstanceFleet(ctx context.Context, input *emr.AddInstanceFleetInput) (*emr.AddInstanceFleetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.AddInstanceFleetWithContext(ctx, input)
}

func (a *EMRActivities) AddInstanceGroups(ctx context.Context, input *emr.AddInstanceGroupsInput) (*emr.AddInstanceGroupsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.AddInstanceGroupsWithContext(ctx, input)
}

func (a *EMRActivities) AddJobFlowSteps(ctx context.Context, input *emr.AddJobFlowStepsInput) (*emr.AddJobFlowStepsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.AddJobFlowStepsWithContext(ctx, input)
}

func (a *EMRActivities) AddTags(ctx context.Context, input *emr.AddTagsInput) (*emr.AddTagsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.AddTagsWithContext(ctx, input)
}

func (a *EMRActivities) CancelSteps(ctx context.Context, input *emr.CancelStepsInput) (*emr.CancelStepsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CancelStepsWithContext(ctx, input)
}

func (a *EMRActivities) CreateSecurityConfiguration(ctx context.Context, input *emr.CreateSecurityConfigurationInput) (*emr.CreateSecurityConfigurationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateSecurityConfigurationWithContext(ctx, input)
}

func (a *EMRActivities) DeleteSecurityConfiguration(ctx context.Context, input *emr.DeleteSecurityConfigurationInput) (*emr.DeleteSecurityConfigurationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteSecurityConfigurationWithContext(ctx, input)
}

func (a *EMRActivities) DescribeCluster(ctx context.Context, input *emr.DescribeClusterInput) (*emr.DescribeClusterOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeClusterWithContext(ctx, input)
}

func (a *EMRActivities) DescribeJobFlows(ctx context.Context, input *emr.DescribeJobFlowsInput) (*emr.DescribeJobFlowsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeJobFlowsWithContext(ctx, input)
}

func (a *EMRActivities) DescribeNotebookExecution(ctx context.Context, input *emr.DescribeNotebookExecutionInput) (*emr.DescribeNotebookExecutionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeNotebookExecutionWithContext(ctx, input)
}

func (a *EMRActivities) DescribeSecurityConfiguration(ctx context.Context, input *emr.DescribeSecurityConfigurationInput) (*emr.DescribeSecurityConfigurationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeSecurityConfigurationWithContext(ctx, input)
}

func (a *EMRActivities) DescribeStep(ctx context.Context, input *emr.DescribeStepInput) (*emr.DescribeStepOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeStepWithContext(ctx, input)
}

func (a *EMRActivities) GetBlockPublicAccessConfiguration(ctx context.Context, input *emr.GetBlockPublicAccessConfigurationInput) (*emr.GetBlockPublicAccessConfigurationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetBlockPublicAccessConfigurationWithContext(ctx, input)
}

func (a *EMRActivities) GetManagedScalingPolicy(ctx context.Context, input *emr.GetManagedScalingPolicyInput) (*emr.GetManagedScalingPolicyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetManagedScalingPolicyWithContext(ctx, input)
}

func (a *EMRActivities) ListBootstrapActions(ctx context.Context, input *emr.ListBootstrapActionsInput) (*emr.ListBootstrapActionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListBootstrapActionsWithContext(ctx, input)
}

func (a *EMRActivities) ListClusters(ctx context.Context, input *emr.ListClustersInput) (*emr.ListClustersOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListClustersWithContext(ctx, input)
}

func (a *EMRActivities) ListInstanceFleets(ctx context.Context, input *emr.ListInstanceFleetsInput) (*emr.ListInstanceFleetsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListInstanceFleetsWithContext(ctx, input)
}

func (a *EMRActivities) ListInstanceGroups(ctx context.Context, input *emr.ListInstanceGroupsInput) (*emr.ListInstanceGroupsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListInstanceGroupsWithContext(ctx, input)
}

func (a *EMRActivities) ListInstances(ctx context.Context, input *emr.ListInstancesInput) (*emr.ListInstancesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListInstancesWithContext(ctx, input)
}

func (a *EMRActivities) ListNotebookExecutions(ctx context.Context, input *emr.ListNotebookExecutionsInput) (*emr.ListNotebookExecutionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListNotebookExecutionsWithContext(ctx, input)
}

func (a *EMRActivities) ListSecurityConfigurations(ctx context.Context, input *emr.ListSecurityConfigurationsInput) (*emr.ListSecurityConfigurationsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListSecurityConfigurationsWithContext(ctx, input)
}

func (a *EMRActivities) ListSteps(ctx context.Context, input *emr.ListStepsInput) (*emr.ListStepsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListStepsWithContext(ctx, input)
}

func (a *EMRActivities) ModifyCluster(ctx context.Context, input *emr.ModifyClusterInput) (*emr.ModifyClusterOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ModifyClusterWithContext(ctx, input)
}

func (a *EMRActivities) ModifyInstanceFleet(ctx context.Context, input *emr.ModifyInstanceFleetInput) (*emr.ModifyInstanceFleetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ModifyInstanceFleetWithContext(ctx, input)
}

func (a *EMRActivities) ModifyInstanceGroups(ctx context.Context, input *emr.ModifyInstanceGroupsInput) (*emr.ModifyInstanceGroupsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ModifyInstanceGroupsWithContext(ctx, input)
}

func (a *EMRActivities) PutAutoScalingPolicy(ctx context.Context, input *emr.PutAutoScalingPolicyInput) (*emr.PutAutoScalingPolicyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.PutAutoScalingPolicyWithContext(ctx, input)
}

func (a *EMRActivities) PutBlockPublicAccessConfiguration(ctx context.Context, input *emr.PutBlockPublicAccessConfigurationInput) (*emr.PutBlockPublicAccessConfigurationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.PutBlockPublicAccessConfigurationWithContext(ctx, input)
}

func (a *EMRActivities) PutManagedScalingPolicy(ctx context.Context, input *emr.PutManagedScalingPolicyInput) (*emr.PutManagedScalingPolicyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.PutManagedScalingPolicyWithContext(ctx, input)
}

func (a *EMRActivities) RemoveAutoScalingPolicy(ctx context.Context, input *emr.RemoveAutoScalingPolicyInput) (*emr.RemoveAutoScalingPolicyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.RemoveAutoScalingPolicyWithContext(ctx, input)
}

func (a *EMRActivities) RemoveManagedScalingPolicy(ctx context.Context, input *emr.RemoveManagedScalingPolicyInput) (*emr.RemoveManagedScalingPolicyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.RemoveManagedScalingPolicyWithContext(ctx, input)
}

func (a *EMRActivities) RemoveTags(ctx context.Context, input *emr.RemoveTagsInput) (*emr.RemoveTagsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.RemoveTagsWithContext(ctx, input)
}

func (a *EMRActivities) RunJobFlow(ctx context.Context, input *emr.RunJobFlowInput) (*emr.RunJobFlowOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.RunJobFlowWithContext(ctx, input)
}

func (a *EMRActivities) SetTerminationProtection(ctx context.Context, input *emr.SetTerminationProtectionInput) (*emr.SetTerminationProtectionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.SetTerminationProtectionWithContext(ctx, input)
}

func (a *EMRActivities) SetVisibleToAllUsers(ctx context.Context, input *emr.SetVisibleToAllUsersInput) (*emr.SetVisibleToAllUsersOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.SetVisibleToAllUsersWithContext(ctx, input)
}

func (a *EMRActivities) StartNotebookExecution(ctx context.Context, input *emr.StartNotebookExecutionInput) (*emr.StartNotebookExecutionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.StartNotebookExecutionWithContext(ctx, input)
}

func (a *EMRActivities) StopNotebookExecution(ctx context.Context, input *emr.StopNotebookExecutionInput) (*emr.StopNotebookExecutionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.StopNotebookExecutionWithContext(ctx, input)
}

func (a *EMRActivities) TerminateJobFlows(ctx context.Context, input *emr.TerminateJobFlowsInput) (*emr.TerminateJobFlowsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.TerminateJobFlowsWithContext(ctx, input)
}

func (a *EMRActivities) WaitUntilClusterRunning(ctx context.Context, input *emr.DescribeClusterInput) error {
	client, err := a.getClient(ctx)
	if err != nil {
		return err
	}
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return client.WaitUntilClusterRunningWithContext(ctx, input, options...)
	})
}

func (a *EMRActivities) WaitUntilClusterTerminated(ctx context.Context, input *emr.DescribeClusterInput) error {
	client, err := a.getClient(ctx)
	if err != nil {
		return err
	}
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return client.WaitUntilClusterTerminatedWithContext(ctx, input, options...)
	})
}

func (a *EMRActivities) WaitUntilStepComplete(ctx context.Context, input *emr.DescribeStepInput) error {
	client, err := a.getClient(ctx)
	if err != nil {
		return err
	}
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return client.WaitUntilStepCompleteWithContext(ctx, input, options...)
	})
}
