// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package ecsstub

import (
	"github.com/aws/aws-sdk-go/service/ecs"
	"go.temporal.io/sdk/workflow"

	"go.temporal.io/aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	CreateCapacityProvider(ctx workflow.Context, input *ecs.CreateCapacityProviderInput) (*ecs.CreateCapacityProviderOutput, error)
	CreateCapacityProviderAsync(ctx workflow.Context, input *ecs.CreateCapacityProviderInput) *ECSCreateCapacityProviderFuture

	CreateCluster(ctx workflow.Context, input *ecs.CreateClusterInput) (*ecs.CreateClusterOutput, error)
	CreateClusterAsync(ctx workflow.Context, input *ecs.CreateClusterInput) *ECSCreateClusterFuture

	CreateService(ctx workflow.Context, input *ecs.CreateServiceInput) (*ecs.CreateServiceOutput, error)
	CreateServiceAsync(ctx workflow.Context, input *ecs.CreateServiceInput) *ECSCreateServiceFuture

	CreateTaskSet(ctx workflow.Context, input *ecs.CreateTaskSetInput) (*ecs.CreateTaskSetOutput, error)
	CreateTaskSetAsync(ctx workflow.Context, input *ecs.CreateTaskSetInput) *ECSCreateTaskSetFuture

	DeleteAccountSetting(ctx workflow.Context, input *ecs.DeleteAccountSettingInput) (*ecs.DeleteAccountSettingOutput, error)
	DeleteAccountSettingAsync(ctx workflow.Context, input *ecs.DeleteAccountSettingInput) *ECSDeleteAccountSettingFuture

	DeleteAttributes(ctx workflow.Context, input *ecs.DeleteAttributesInput) (*ecs.DeleteAttributesOutput, error)
	DeleteAttributesAsync(ctx workflow.Context, input *ecs.DeleteAttributesInput) *ECSDeleteAttributesFuture

	DeleteCapacityProvider(ctx workflow.Context, input *ecs.DeleteCapacityProviderInput) (*ecs.DeleteCapacityProviderOutput, error)
	DeleteCapacityProviderAsync(ctx workflow.Context, input *ecs.DeleteCapacityProviderInput) *ECSDeleteCapacityProviderFuture

	DeleteCluster(ctx workflow.Context, input *ecs.DeleteClusterInput) (*ecs.DeleteClusterOutput, error)
	DeleteClusterAsync(ctx workflow.Context, input *ecs.DeleteClusterInput) *ECSDeleteClusterFuture

	DeleteService(ctx workflow.Context, input *ecs.DeleteServiceInput) (*ecs.DeleteServiceOutput, error)
	DeleteServiceAsync(ctx workflow.Context, input *ecs.DeleteServiceInput) *ECSDeleteServiceFuture

	DeleteTaskSet(ctx workflow.Context, input *ecs.DeleteTaskSetInput) (*ecs.DeleteTaskSetOutput, error)
	DeleteTaskSetAsync(ctx workflow.Context, input *ecs.DeleteTaskSetInput) *ECSDeleteTaskSetFuture

	DeregisterContainerInstance(ctx workflow.Context, input *ecs.DeregisterContainerInstanceInput) (*ecs.DeregisterContainerInstanceOutput, error)
	DeregisterContainerInstanceAsync(ctx workflow.Context, input *ecs.DeregisterContainerInstanceInput) *ECSDeregisterContainerInstanceFuture

	DeregisterTaskDefinition(ctx workflow.Context, input *ecs.DeregisterTaskDefinitionInput) (*ecs.DeregisterTaskDefinitionOutput, error)
	DeregisterTaskDefinitionAsync(ctx workflow.Context, input *ecs.DeregisterTaskDefinitionInput) *ECSDeregisterTaskDefinitionFuture

	DescribeCapacityProviders(ctx workflow.Context, input *ecs.DescribeCapacityProvidersInput) (*ecs.DescribeCapacityProvidersOutput, error)
	DescribeCapacityProvidersAsync(ctx workflow.Context, input *ecs.DescribeCapacityProvidersInput) *ECSDescribeCapacityProvidersFuture

	DescribeClusters(ctx workflow.Context, input *ecs.DescribeClustersInput) (*ecs.DescribeClustersOutput, error)
	DescribeClustersAsync(ctx workflow.Context, input *ecs.DescribeClustersInput) *ECSDescribeClustersFuture

	DescribeContainerInstances(ctx workflow.Context, input *ecs.DescribeContainerInstancesInput) (*ecs.DescribeContainerInstancesOutput, error)
	DescribeContainerInstancesAsync(ctx workflow.Context, input *ecs.DescribeContainerInstancesInput) *ECSDescribeContainerInstancesFuture

	DescribeServices(ctx workflow.Context, input *ecs.DescribeServicesInput) (*ecs.DescribeServicesOutput, error)
	DescribeServicesAsync(ctx workflow.Context, input *ecs.DescribeServicesInput) *ECSDescribeServicesFuture

	DescribeTaskDefinition(ctx workflow.Context, input *ecs.DescribeTaskDefinitionInput) (*ecs.DescribeTaskDefinitionOutput, error)
	DescribeTaskDefinitionAsync(ctx workflow.Context, input *ecs.DescribeTaskDefinitionInput) *ECSDescribeTaskDefinitionFuture

	DescribeTaskSets(ctx workflow.Context, input *ecs.DescribeTaskSetsInput) (*ecs.DescribeTaskSetsOutput, error)
	DescribeTaskSetsAsync(ctx workflow.Context, input *ecs.DescribeTaskSetsInput) *ECSDescribeTaskSetsFuture

	DescribeTasks(ctx workflow.Context, input *ecs.DescribeTasksInput) (*ecs.DescribeTasksOutput, error)
	DescribeTasksAsync(ctx workflow.Context, input *ecs.DescribeTasksInput) *ECSDescribeTasksFuture

	DiscoverPollEndpoint(ctx workflow.Context, input *ecs.DiscoverPollEndpointInput) (*ecs.DiscoverPollEndpointOutput, error)
	DiscoverPollEndpointAsync(ctx workflow.Context, input *ecs.DiscoverPollEndpointInput) *ECSDiscoverPollEndpointFuture

	ListAccountSettings(ctx workflow.Context, input *ecs.ListAccountSettingsInput) (*ecs.ListAccountSettingsOutput, error)
	ListAccountSettingsAsync(ctx workflow.Context, input *ecs.ListAccountSettingsInput) *ECSListAccountSettingsFuture

	ListAttributes(ctx workflow.Context, input *ecs.ListAttributesInput) (*ecs.ListAttributesOutput, error)
	ListAttributesAsync(ctx workflow.Context, input *ecs.ListAttributesInput) *ECSListAttributesFuture

	ListClusters(ctx workflow.Context, input *ecs.ListClustersInput) (*ecs.ListClustersOutput, error)
	ListClustersAsync(ctx workflow.Context, input *ecs.ListClustersInput) *ECSListClustersFuture

	ListContainerInstances(ctx workflow.Context, input *ecs.ListContainerInstancesInput) (*ecs.ListContainerInstancesOutput, error)
	ListContainerInstancesAsync(ctx workflow.Context, input *ecs.ListContainerInstancesInput) *ECSListContainerInstancesFuture

	ListServices(ctx workflow.Context, input *ecs.ListServicesInput) (*ecs.ListServicesOutput, error)
	ListServicesAsync(ctx workflow.Context, input *ecs.ListServicesInput) *ECSListServicesFuture

	ListTagsForResource(ctx workflow.Context, input *ecs.ListTagsForResourceInput) (*ecs.ListTagsForResourceOutput, error)
	ListTagsForResourceAsync(ctx workflow.Context, input *ecs.ListTagsForResourceInput) *ECSListTagsForResourceFuture

	ListTaskDefinitionFamilies(ctx workflow.Context, input *ecs.ListTaskDefinitionFamiliesInput) (*ecs.ListTaskDefinitionFamiliesOutput, error)
	ListTaskDefinitionFamiliesAsync(ctx workflow.Context, input *ecs.ListTaskDefinitionFamiliesInput) *ECSListTaskDefinitionFamiliesFuture

	ListTaskDefinitions(ctx workflow.Context, input *ecs.ListTaskDefinitionsInput) (*ecs.ListTaskDefinitionsOutput, error)
	ListTaskDefinitionsAsync(ctx workflow.Context, input *ecs.ListTaskDefinitionsInput) *ECSListTaskDefinitionsFuture

	ListTasks(ctx workflow.Context, input *ecs.ListTasksInput) (*ecs.ListTasksOutput, error)
	ListTasksAsync(ctx workflow.Context, input *ecs.ListTasksInput) *ECSListTasksFuture

	PutAccountSetting(ctx workflow.Context, input *ecs.PutAccountSettingInput) (*ecs.PutAccountSettingOutput, error)
	PutAccountSettingAsync(ctx workflow.Context, input *ecs.PutAccountSettingInput) *ECSPutAccountSettingFuture

	PutAccountSettingDefault(ctx workflow.Context, input *ecs.PutAccountSettingDefaultInput) (*ecs.PutAccountSettingDefaultOutput, error)
	PutAccountSettingDefaultAsync(ctx workflow.Context, input *ecs.PutAccountSettingDefaultInput) *ECSPutAccountSettingDefaultFuture

	PutAttributes(ctx workflow.Context, input *ecs.PutAttributesInput) (*ecs.PutAttributesOutput, error)
	PutAttributesAsync(ctx workflow.Context, input *ecs.PutAttributesInput) *ECSPutAttributesFuture

	PutClusterCapacityProviders(ctx workflow.Context, input *ecs.PutClusterCapacityProvidersInput) (*ecs.PutClusterCapacityProvidersOutput, error)
	PutClusterCapacityProvidersAsync(ctx workflow.Context, input *ecs.PutClusterCapacityProvidersInput) *ECSPutClusterCapacityProvidersFuture

	RegisterContainerInstance(ctx workflow.Context, input *ecs.RegisterContainerInstanceInput) (*ecs.RegisterContainerInstanceOutput, error)
	RegisterContainerInstanceAsync(ctx workflow.Context, input *ecs.RegisterContainerInstanceInput) *ECSRegisterContainerInstanceFuture

	RegisterTaskDefinition(ctx workflow.Context, input *ecs.RegisterTaskDefinitionInput) (*ecs.RegisterTaskDefinitionOutput, error)
	RegisterTaskDefinitionAsync(ctx workflow.Context, input *ecs.RegisterTaskDefinitionInput) *ECSRegisterTaskDefinitionFuture

	RunTask(ctx workflow.Context, input *ecs.RunTaskInput) (*ecs.RunTaskOutput, error)
	RunTaskAsync(ctx workflow.Context, input *ecs.RunTaskInput) *ECSRunTaskFuture

	StartTask(ctx workflow.Context, input *ecs.StartTaskInput) (*ecs.StartTaskOutput, error)
	StartTaskAsync(ctx workflow.Context, input *ecs.StartTaskInput) *ECSStartTaskFuture

	StopTask(ctx workflow.Context, input *ecs.StopTaskInput) (*ecs.StopTaskOutput, error)
	StopTaskAsync(ctx workflow.Context, input *ecs.StopTaskInput) *ECSStopTaskFuture

	SubmitAttachmentStateChanges(ctx workflow.Context, input *ecs.SubmitAttachmentStateChangesInput) (*ecs.SubmitAttachmentStateChangesOutput, error)
	SubmitAttachmentStateChangesAsync(ctx workflow.Context, input *ecs.SubmitAttachmentStateChangesInput) *ECSSubmitAttachmentStateChangesFuture

	SubmitContainerStateChange(ctx workflow.Context, input *ecs.SubmitContainerStateChangeInput) (*ecs.SubmitContainerStateChangeOutput, error)
	SubmitContainerStateChangeAsync(ctx workflow.Context, input *ecs.SubmitContainerStateChangeInput) *ECSSubmitContainerStateChangeFuture

	SubmitTaskStateChange(ctx workflow.Context, input *ecs.SubmitTaskStateChangeInput) (*ecs.SubmitTaskStateChangeOutput, error)
	SubmitTaskStateChangeAsync(ctx workflow.Context, input *ecs.SubmitTaskStateChangeInput) *ECSSubmitTaskStateChangeFuture

	TagResource(ctx workflow.Context, input *ecs.TagResourceInput) (*ecs.TagResourceOutput, error)
	TagResourceAsync(ctx workflow.Context, input *ecs.TagResourceInput) *ECSTagResourceFuture

	UntagResource(ctx workflow.Context, input *ecs.UntagResourceInput) (*ecs.UntagResourceOutput, error)
	UntagResourceAsync(ctx workflow.Context, input *ecs.UntagResourceInput) *ECSUntagResourceFuture

	UpdateClusterSettings(ctx workflow.Context, input *ecs.UpdateClusterSettingsInput) (*ecs.UpdateClusterSettingsOutput, error)
	UpdateClusterSettingsAsync(ctx workflow.Context, input *ecs.UpdateClusterSettingsInput) *ECSUpdateClusterSettingsFuture

	UpdateContainerAgent(ctx workflow.Context, input *ecs.UpdateContainerAgentInput) (*ecs.UpdateContainerAgentOutput, error)
	UpdateContainerAgentAsync(ctx workflow.Context, input *ecs.UpdateContainerAgentInput) *ECSUpdateContainerAgentFuture

	UpdateContainerInstancesState(ctx workflow.Context, input *ecs.UpdateContainerInstancesStateInput) (*ecs.UpdateContainerInstancesStateOutput, error)
	UpdateContainerInstancesStateAsync(ctx workflow.Context, input *ecs.UpdateContainerInstancesStateInput) *ECSUpdateContainerInstancesStateFuture

	UpdateService(ctx workflow.Context, input *ecs.UpdateServiceInput) (*ecs.UpdateServiceOutput, error)
	UpdateServiceAsync(ctx workflow.Context, input *ecs.UpdateServiceInput) *ECSUpdateServiceFuture

	UpdateServicePrimaryTaskSet(ctx workflow.Context, input *ecs.UpdateServicePrimaryTaskSetInput) (*ecs.UpdateServicePrimaryTaskSetOutput, error)
	UpdateServicePrimaryTaskSetAsync(ctx workflow.Context, input *ecs.UpdateServicePrimaryTaskSetInput) *ECSUpdateServicePrimaryTaskSetFuture

	UpdateTaskSet(ctx workflow.Context, input *ecs.UpdateTaskSetInput) (*ecs.UpdateTaskSetOutput, error)
	UpdateTaskSetAsync(ctx workflow.Context, input *ecs.UpdateTaskSetInput) *ECSUpdateTaskSetFuture

	WaitUntilServicesInactive(ctx workflow.Context, input *ecs.DescribeServicesInput) error
	WaitUntilServicesInactiveAsync(ctx workflow.Context, input *ecs.DescribeServicesInput) *clients.VoidFuture

	WaitUntilServicesStable(ctx workflow.Context, input *ecs.DescribeServicesInput) error
	WaitUntilServicesStableAsync(ctx workflow.Context, input *ecs.DescribeServicesInput) *clients.VoidFuture

	WaitUntilTasksRunning(ctx workflow.Context, input *ecs.DescribeTasksInput) error
	WaitUntilTasksRunningAsync(ctx workflow.Context, input *ecs.DescribeTasksInput) *clients.VoidFuture

	WaitUntilTasksStopped(ctx workflow.Context, input *ecs.DescribeTasksInput) error
	WaitUntilTasksStoppedAsync(ctx workflow.Context, input *ecs.DescribeTasksInput) *clients.VoidFuture
}

func NewClient() Client {
	return &stub{}
}