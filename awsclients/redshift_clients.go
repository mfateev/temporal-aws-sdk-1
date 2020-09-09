package awsclients

import (
	"github.com/aws/aws-sdk-go/service/redshift"
	"go.temporal.io/sdk/workflow"
	"temporal.io/aws-sdk/awsactivities"
)

type RedshiftClient interface {
	AcceptReservedNodeExchange(ctx workflow.Context, input *redshift.AcceptReservedNodeExchangeInput) (*redshift.AcceptReservedNodeExchangeOutput, error)
	AcceptReservedNodeExchangeAsync(ctx workflow.Context, input *redshift.AcceptReservedNodeExchangeInput) *RedshiftAcceptReservedNodeExchangeResult

	AuthorizeClusterSecurityGroupIngress(ctx workflow.Context, input *redshift.AuthorizeClusterSecurityGroupIngressInput) (*redshift.AuthorizeClusterSecurityGroupIngressOutput, error)
	AuthorizeClusterSecurityGroupIngressAsync(ctx workflow.Context, input *redshift.AuthorizeClusterSecurityGroupIngressInput) *RedshiftAuthorizeClusterSecurityGroupIngressResult

	AuthorizeSnapshotAccess(ctx workflow.Context, input *redshift.AuthorizeSnapshotAccessInput) (*redshift.AuthorizeSnapshotAccessOutput, error)
	AuthorizeSnapshotAccessAsync(ctx workflow.Context, input *redshift.AuthorizeSnapshotAccessInput) *RedshiftAuthorizeSnapshotAccessResult

	BatchDeleteClusterSnapshots(ctx workflow.Context, input *redshift.BatchDeleteClusterSnapshotsInput) (*redshift.BatchDeleteClusterSnapshotsOutput, error)
	BatchDeleteClusterSnapshotsAsync(ctx workflow.Context, input *redshift.BatchDeleteClusterSnapshotsInput) *RedshiftBatchDeleteClusterSnapshotsResult

	BatchModifyClusterSnapshots(ctx workflow.Context, input *redshift.BatchModifyClusterSnapshotsInput) (*redshift.BatchModifyClusterSnapshotsOutput, error)
	BatchModifyClusterSnapshotsAsync(ctx workflow.Context, input *redshift.BatchModifyClusterSnapshotsInput) *RedshiftBatchModifyClusterSnapshotsResult

	CancelResize(ctx workflow.Context, input *redshift.CancelResizeInput) (*redshift.CancelResizeOutput, error)
	CancelResizeAsync(ctx workflow.Context, input *redshift.CancelResizeInput) *RedshiftCancelResizeResult

	CopyClusterSnapshot(ctx workflow.Context, input *redshift.CopyClusterSnapshotInput) (*redshift.CopyClusterSnapshotOutput, error)
	CopyClusterSnapshotAsync(ctx workflow.Context, input *redshift.CopyClusterSnapshotInput) *RedshiftCopyClusterSnapshotResult

	CreateCluster(ctx workflow.Context, input *redshift.CreateClusterInput) (*redshift.CreateClusterOutput, error)
	CreateClusterAsync(ctx workflow.Context, input *redshift.CreateClusterInput) *RedshiftCreateClusterResult

	CreateClusterParameterGroup(ctx workflow.Context, input *redshift.CreateClusterParameterGroupInput) (*redshift.CreateClusterParameterGroupOutput, error)
	CreateClusterParameterGroupAsync(ctx workflow.Context, input *redshift.CreateClusterParameterGroupInput) *RedshiftCreateClusterParameterGroupResult

	CreateClusterSecurityGroup(ctx workflow.Context, input *redshift.CreateClusterSecurityGroupInput) (*redshift.CreateClusterSecurityGroupOutput, error)
	CreateClusterSecurityGroupAsync(ctx workflow.Context, input *redshift.CreateClusterSecurityGroupInput) *RedshiftCreateClusterSecurityGroupResult

	CreateClusterSnapshot(ctx workflow.Context, input *redshift.CreateClusterSnapshotInput) (*redshift.CreateClusterSnapshotOutput, error)
	CreateClusterSnapshotAsync(ctx workflow.Context, input *redshift.CreateClusterSnapshotInput) *RedshiftCreateClusterSnapshotResult

	CreateClusterSubnetGroup(ctx workflow.Context, input *redshift.CreateClusterSubnetGroupInput) (*redshift.CreateClusterSubnetGroupOutput, error)
	CreateClusterSubnetGroupAsync(ctx workflow.Context, input *redshift.CreateClusterSubnetGroupInput) *RedshiftCreateClusterSubnetGroupResult

	CreateEventSubscription(ctx workflow.Context, input *redshift.CreateEventSubscriptionInput) (*redshift.CreateEventSubscriptionOutput, error)
	CreateEventSubscriptionAsync(ctx workflow.Context, input *redshift.CreateEventSubscriptionInput) *RedshiftCreateEventSubscriptionResult

	CreateHsmClientCertificate(ctx workflow.Context, input *redshift.CreateHsmClientCertificateInput) (*redshift.CreateHsmClientCertificateOutput, error)
	CreateHsmClientCertificateAsync(ctx workflow.Context, input *redshift.CreateHsmClientCertificateInput) *RedshiftCreateHsmClientCertificateResult

	CreateHsmConfiguration(ctx workflow.Context, input *redshift.CreateHsmConfigurationInput) (*redshift.CreateHsmConfigurationOutput, error)
	CreateHsmConfigurationAsync(ctx workflow.Context, input *redshift.CreateHsmConfigurationInput) *RedshiftCreateHsmConfigurationResult

	CreateScheduledAction(ctx workflow.Context, input *redshift.CreateScheduledActionInput) (*redshift.CreateScheduledActionOutput, error)
	CreateScheduledActionAsync(ctx workflow.Context, input *redshift.CreateScheduledActionInput) *RedshiftCreateScheduledActionResult

	CreateSnapshotCopyGrant(ctx workflow.Context, input *redshift.CreateSnapshotCopyGrantInput) (*redshift.CreateSnapshotCopyGrantOutput, error)
	CreateSnapshotCopyGrantAsync(ctx workflow.Context, input *redshift.CreateSnapshotCopyGrantInput) *RedshiftCreateSnapshotCopyGrantResult

	CreateSnapshotSchedule(ctx workflow.Context, input *redshift.CreateSnapshotScheduleInput) (*redshift.CreateSnapshotScheduleOutput, error)
	CreateSnapshotScheduleAsync(ctx workflow.Context, input *redshift.CreateSnapshotScheduleInput) *RedshiftCreateSnapshotScheduleResult

	CreateTags(ctx workflow.Context, input *redshift.CreateTagsInput) (*redshift.CreateTagsOutput, error)
	CreateTagsAsync(ctx workflow.Context, input *redshift.CreateTagsInput) *RedshiftCreateTagsResult

	CreateUsageLimit(ctx workflow.Context, input *redshift.CreateUsageLimitInput) (*redshift.CreateUsageLimitOutput, error)
	CreateUsageLimitAsync(ctx workflow.Context, input *redshift.CreateUsageLimitInput) *RedshiftCreateUsageLimitResult

	DeleteCluster(ctx workflow.Context, input *redshift.DeleteClusterInput) (*redshift.DeleteClusterOutput, error)
	DeleteClusterAsync(ctx workflow.Context, input *redshift.DeleteClusterInput) *RedshiftDeleteClusterResult

	DeleteClusterParameterGroup(ctx workflow.Context, input *redshift.DeleteClusterParameterGroupInput) (*redshift.DeleteClusterParameterGroupOutput, error)
	DeleteClusterParameterGroupAsync(ctx workflow.Context, input *redshift.DeleteClusterParameterGroupInput) *RedshiftDeleteClusterParameterGroupResult

	DeleteClusterSecurityGroup(ctx workflow.Context, input *redshift.DeleteClusterSecurityGroupInput) (*redshift.DeleteClusterSecurityGroupOutput, error)
	DeleteClusterSecurityGroupAsync(ctx workflow.Context, input *redshift.DeleteClusterSecurityGroupInput) *RedshiftDeleteClusterSecurityGroupResult

	DeleteClusterSnapshot(ctx workflow.Context, input *redshift.DeleteClusterSnapshotInput) (*redshift.DeleteClusterSnapshotOutput, error)
	DeleteClusterSnapshotAsync(ctx workflow.Context, input *redshift.DeleteClusterSnapshotInput) *RedshiftDeleteClusterSnapshotResult

	DeleteClusterSubnetGroup(ctx workflow.Context, input *redshift.DeleteClusterSubnetGroupInput) (*redshift.DeleteClusterSubnetGroupOutput, error)
	DeleteClusterSubnetGroupAsync(ctx workflow.Context, input *redshift.DeleteClusterSubnetGroupInput) *RedshiftDeleteClusterSubnetGroupResult

	DeleteEventSubscription(ctx workflow.Context, input *redshift.DeleteEventSubscriptionInput) (*redshift.DeleteEventSubscriptionOutput, error)
	DeleteEventSubscriptionAsync(ctx workflow.Context, input *redshift.DeleteEventSubscriptionInput) *RedshiftDeleteEventSubscriptionResult

	DeleteHsmClientCertificate(ctx workflow.Context, input *redshift.DeleteHsmClientCertificateInput) (*redshift.DeleteHsmClientCertificateOutput, error)
	DeleteHsmClientCertificateAsync(ctx workflow.Context, input *redshift.DeleteHsmClientCertificateInput) *RedshiftDeleteHsmClientCertificateResult

	DeleteHsmConfiguration(ctx workflow.Context, input *redshift.DeleteHsmConfigurationInput) (*redshift.DeleteHsmConfigurationOutput, error)
	DeleteHsmConfigurationAsync(ctx workflow.Context, input *redshift.DeleteHsmConfigurationInput) *RedshiftDeleteHsmConfigurationResult

	DeleteScheduledAction(ctx workflow.Context, input *redshift.DeleteScheduledActionInput) (*redshift.DeleteScheduledActionOutput, error)
	DeleteScheduledActionAsync(ctx workflow.Context, input *redshift.DeleteScheduledActionInput) *RedshiftDeleteScheduledActionResult

	DeleteSnapshotCopyGrant(ctx workflow.Context, input *redshift.DeleteSnapshotCopyGrantInput) (*redshift.DeleteSnapshotCopyGrantOutput, error)
	DeleteSnapshotCopyGrantAsync(ctx workflow.Context, input *redshift.DeleteSnapshotCopyGrantInput) *RedshiftDeleteSnapshotCopyGrantResult

	DeleteSnapshotSchedule(ctx workflow.Context, input *redshift.DeleteSnapshotScheduleInput) (*redshift.DeleteSnapshotScheduleOutput, error)
	DeleteSnapshotScheduleAsync(ctx workflow.Context, input *redshift.DeleteSnapshotScheduleInput) *RedshiftDeleteSnapshotScheduleResult

	DeleteTags(ctx workflow.Context, input *redshift.DeleteTagsInput) (*redshift.DeleteTagsOutput, error)
	DeleteTagsAsync(ctx workflow.Context, input *redshift.DeleteTagsInput) *RedshiftDeleteTagsResult

	DeleteUsageLimit(ctx workflow.Context, input *redshift.DeleteUsageLimitInput) (*redshift.DeleteUsageLimitOutput, error)
	DeleteUsageLimitAsync(ctx workflow.Context, input *redshift.DeleteUsageLimitInput) *RedshiftDeleteUsageLimitResult

	DescribeAccountAttributes(ctx workflow.Context, input *redshift.DescribeAccountAttributesInput) (*redshift.DescribeAccountAttributesOutput, error)
	DescribeAccountAttributesAsync(ctx workflow.Context, input *redshift.DescribeAccountAttributesInput) *RedshiftDescribeAccountAttributesResult

	DescribeClusterDbRevisions(ctx workflow.Context, input *redshift.DescribeClusterDbRevisionsInput) (*redshift.DescribeClusterDbRevisionsOutput, error)
	DescribeClusterDbRevisionsAsync(ctx workflow.Context, input *redshift.DescribeClusterDbRevisionsInput) *RedshiftDescribeClusterDbRevisionsResult

	DescribeClusterParameterGroups(ctx workflow.Context, input *redshift.DescribeClusterParameterGroupsInput) (*redshift.DescribeClusterParameterGroupsOutput, error)
	DescribeClusterParameterGroupsAsync(ctx workflow.Context, input *redshift.DescribeClusterParameterGroupsInput) *RedshiftDescribeClusterParameterGroupsResult

	DescribeClusterParameters(ctx workflow.Context, input *redshift.DescribeClusterParametersInput) (*redshift.DescribeClusterParametersOutput, error)
	DescribeClusterParametersAsync(ctx workflow.Context, input *redshift.DescribeClusterParametersInput) *RedshiftDescribeClusterParametersResult

	DescribeClusterSecurityGroups(ctx workflow.Context, input *redshift.DescribeClusterSecurityGroupsInput) (*redshift.DescribeClusterSecurityGroupsOutput, error)
	DescribeClusterSecurityGroupsAsync(ctx workflow.Context, input *redshift.DescribeClusterSecurityGroupsInput) *RedshiftDescribeClusterSecurityGroupsResult

	DescribeClusterSnapshots(ctx workflow.Context, input *redshift.DescribeClusterSnapshotsInput) (*redshift.DescribeClusterSnapshotsOutput, error)
	DescribeClusterSnapshotsAsync(ctx workflow.Context, input *redshift.DescribeClusterSnapshotsInput) *RedshiftDescribeClusterSnapshotsResult

	DescribeClusterSubnetGroups(ctx workflow.Context, input *redshift.DescribeClusterSubnetGroupsInput) (*redshift.DescribeClusterSubnetGroupsOutput, error)
	DescribeClusterSubnetGroupsAsync(ctx workflow.Context, input *redshift.DescribeClusterSubnetGroupsInput) *RedshiftDescribeClusterSubnetGroupsResult

	DescribeClusterTracks(ctx workflow.Context, input *redshift.DescribeClusterTracksInput) (*redshift.DescribeClusterTracksOutput, error)
	DescribeClusterTracksAsync(ctx workflow.Context, input *redshift.DescribeClusterTracksInput) *RedshiftDescribeClusterTracksResult

	DescribeClusterVersions(ctx workflow.Context, input *redshift.DescribeClusterVersionsInput) (*redshift.DescribeClusterVersionsOutput, error)
	DescribeClusterVersionsAsync(ctx workflow.Context, input *redshift.DescribeClusterVersionsInput) *RedshiftDescribeClusterVersionsResult

	DescribeClusters(ctx workflow.Context, input *redshift.DescribeClustersInput) (*redshift.DescribeClustersOutput, error)
	DescribeClustersAsync(ctx workflow.Context, input *redshift.DescribeClustersInput) *RedshiftDescribeClustersResult

	DescribeDefaultClusterParameters(ctx workflow.Context, input *redshift.DescribeDefaultClusterParametersInput) (*redshift.DescribeDefaultClusterParametersOutput, error)
	DescribeDefaultClusterParametersAsync(ctx workflow.Context, input *redshift.DescribeDefaultClusterParametersInput) *RedshiftDescribeDefaultClusterParametersResult

	DescribeEventCategories(ctx workflow.Context, input *redshift.DescribeEventCategoriesInput) (*redshift.DescribeEventCategoriesOutput, error)
	DescribeEventCategoriesAsync(ctx workflow.Context, input *redshift.DescribeEventCategoriesInput) *RedshiftDescribeEventCategoriesResult

	DescribeEventSubscriptions(ctx workflow.Context, input *redshift.DescribeEventSubscriptionsInput) (*redshift.DescribeEventSubscriptionsOutput, error)
	DescribeEventSubscriptionsAsync(ctx workflow.Context, input *redshift.DescribeEventSubscriptionsInput) *RedshiftDescribeEventSubscriptionsResult

	DescribeEvents(ctx workflow.Context, input *redshift.DescribeEventsInput) (*redshift.DescribeEventsOutput, error)
	DescribeEventsAsync(ctx workflow.Context, input *redshift.DescribeEventsInput) *RedshiftDescribeEventsResult

	DescribeHsmClientCertificates(ctx workflow.Context, input *redshift.DescribeHsmClientCertificatesInput) (*redshift.DescribeHsmClientCertificatesOutput, error)
	DescribeHsmClientCertificatesAsync(ctx workflow.Context, input *redshift.DescribeHsmClientCertificatesInput) *RedshiftDescribeHsmClientCertificatesResult

	DescribeHsmConfigurations(ctx workflow.Context, input *redshift.DescribeHsmConfigurationsInput) (*redshift.DescribeHsmConfigurationsOutput, error)
	DescribeHsmConfigurationsAsync(ctx workflow.Context, input *redshift.DescribeHsmConfigurationsInput) *RedshiftDescribeHsmConfigurationsResult

	DescribeLoggingStatus(ctx workflow.Context, input *redshift.DescribeLoggingStatusInput) (*redshift.LoggingStatus, error)
	DescribeLoggingStatusAsync(ctx workflow.Context, input *redshift.DescribeLoggingStatusInput) *RedshiftDescribeLoggingStatusResult

	DescribeNodeConfigurationOptions(ctx workflow.Context, input *redshift.DescribeNodeConfigurationOptionsInput) (*redshift.DescribeNodeConfigurationOptionsOutput, error)
	DescribeNodeConfigurationOptionsAsync(ctx workflow.Context, input *redshift.DescribeNodeConfigurationOptionsInput) *RedshiftDescribeNodeConfigurationOptionsResult

	DescribeOrderableClusterOptions(ctx workflow.Context, input *redshift.DescribeOrderableClusterOptionsInput) (*redshift.DescribeOrderableClusterOptionsOutput, error)
	DescribeOrderableClusterOptionsAsync(ctx workflow.Context, input *redshift.DescribeOrderableClusterOptionsInput) *RedshiftDescribeOrderableClusterOptionsResult

	DescribeReservedNodeOfferings(ctx workflow.Context, input *redshift.DescribeReservedNodeOfferingsInput) (*redshift.DescribeReservedNodeOfferingsOutput, error)
	DescribeReservedNodeOfferingsAsync(ctx workflow.Context, input *redshift.DescribeReservedNodeOfferingsInput) *RedshiftDescribeReservedNodeOfferingsResult

	DescribeReservedNodes(ctx workflow.Context, input *redshift.DescribeReservedNodesInput) (*redshift.DescribeReservedNodesOutput, error)
	DescribeReservedNodesAsync(ctx workflow.Context, input *redshift.DescribeReservedNodesInput) *RedshiftDescribeReservedNodesResult

	DescribeResize(ctx workflow.Context, input *redshift.DescribeResizeInput) (*redshift.DescribeResizeOutput, error)
	DescribeResizeAsync(ctx workflow.Context, input *redshift.DescribeResizeInput) *RedshiftDescribeResizeResult

	DescribeScheduledActions(ctx workflow.Context, input *redshift.DescribeScheduledActionsInput) (*redshift.DescribeScheduledActionsOutput, error)
	DescribeScheduledActionsAsync(ctx workflow.Context, input *redshift.DescribeScheduledActionsInput) *RedshiftDescribeScheduledActionsResult

	DescribeSnapshotCopyGrants(ctx workflow.Context, input *redshift.DescribeSnapshotCopyGrantsInput) (*redshift.DescribeSnapshotCopyGrantsOutput, error)
	DescribeSnapshotCopyGrantsAsync(ctx workflow.Context, input *redshift.DescribeSnapshotCopyGrantsInput) *RedshiftDescribeSnapshotCopyGrantsResult

	DescribeSnapshotSchedules(ctx workflow.Context, input *redshift.DescribeSnapshotSchedulesInput) (*redshift.DescribeSnapshotSchedulesOutput, error)
	DescribeSnapshotSchedulesAsync(ctx workflow.Context, input *redshift.DescribeSnapshotSchedulesInput) *RedshiftDescribeSnapshotSchedulesResult

	DescribeStorage(ctx workflow.Context, input *redshift.DescribeStorageInput) (*redshift.DescribeStorageOutput, error)
	DescribeStorageAsync(ctx workflow.Context, input *redshift.DescribeStorageInput) *RedshiftDescribeStorageResult

	DescribeTableRestoreStatus(ctx workflow.Context, input *redshift.DescribeTableRestoreStatusInput) (*redshift.DescribeTableRestoreStatusOutput, error)
	DescribeTableRestoreStatusAsync(ctx workflow.Context, input *redshift.DescribeTableRestoreStatusInput) *RedshiftDescribeTableRestoreStatusResult

	DescribeTags(ctx workflow.Context, input *redshift.DescribeTagsInput) (*redshift.DescribeTagsOutput, error)
	DescribeTagsAsync(ctx workflow.Context, input *redshift.DescribeTagsInput) *RedshiftDescribeTagsResult

	DescribeUsageLimits(ctx workflow.Context, input *redshift.DescribeUsageLimitsInput) (*redshift.DescribeUsageLimitsOutput, error)
	DescribeUsageLimitsAsync(ctx workflow.Context, input *redshift.DescribeUsageLimitsInput) *RedshiftDescribeUsageLimitsResult

	DisableLogging(ctx workflow.Context, input *redshift.DisableLoggingInput) (*redshift.LoggingStatus, error)
	DisableLoggingAsync(ctx workflow.Context, input *redshift.DisableLoggingInput) *RedshiftDisableLoggingResult

	DisableSnapshotCopy(ctx workflow.Context, input *redshift.DisableSnapshotCopyInput) (*redshift.DisableSnapshotCopyOutput, error)
	DisableSnapshotCopyAsync(ctx workflow.Context, input *redshift.DisableSnapshotCopyInput) *RedshiftDisableSnapshotCopyResult

	EnableLogging(ctx workflow.Context, input *redshift.EnableLoggingInput) (*redshift.LoggingStatus, error)
	EnableLoggingAsync(ctx workflow.Context, input *redshift.EnableLoggingInput) *RedshiftEnableLoggingResult

	EnableSnapshotCopy(ctx workflow.Context, input *redshift.EnableSnapshotCopyInput) (*redshift.EnableSnapshotCopyOutput, error)
	EnableSnapshotCopyAsync(ctx workflow.Context, input *redshift.EnableSnapshotCopyInput) *RedshiftEnableSnapshotCopyResult

	GetClusterCredentials(ctx workflow.Context, input *redshift.GetClusterCredentialsInput) (*redshift.GetClusterCredentialsOutput, error)
	GetClusterCredentialsAsync(ctx workflow.Context, input *redshift.GetClusterCredentialsInput) *RedshiftGetClusterCredentialsResult

	GetReservedNodeExchangeOfferings(ctx workflow.Context, input *redshift.GetReservedNodeExchangeOfferingsInput) (*redshift.GetReservedNodeExchangeOfferingsOutput, error)
	GetReservedNodeExchangeOfferingsAsync(ctx workflow.Context, input *redshift.GetReservedNodeExchangeOfferingsInput) *RedshiftGetReservedNodeExchangeOfferingsResult

	ModifyCluster(ctx workflow.Context, input *redshift.ModifyClusterInput) (*redshift.ModifyClusterOutput, error)
	ModifyClusterAsync(ctx workflow.Context, input *redshift.ModifyClusterInput) *RedshiftModifyClusterResult

	ModifyClusterDbRevision(ctx workflow.Context, input *redshift.ModifyClusterDbRevisionInput) (*redshift.ModifyClusterDbRevisionOutput, error)
	ModifyClusterDbRevisionAsync(ctx workflow.Context, input *redshift.ModifyClusterDbRevisionInput) *RedshiftModifyClusterDbRevisionResult

	ModifyClusterIamRoles(ctx workflow.Context, input *redshift.ModifyClusterIamRolesInput) (*redshift.ModifyClusterIamRolesOutput, error)
	ModifyClusterIamRolesAsync(ctx workflow.Context, input *redshift.ModifyClusterIamRolesInput) *RedshiftModifyClusterIamRolesResult

	ModifyClusterMaintenance(ctx workflow.Context, input *redshift.ModifyClusterMaintenanceInput) (*redshift.ModifyClusterMaintenanceOutput, error)
	ModifyClusterMaintenanceAsync(ctx workflow.Context, input *redshift.ModifyClusterMaintenanceInput) *RedshiftModifyClusterMaintenanceResult

	ModifyClusterParameterGroup(ctx workflow.Context, input *redshift.ModifyClusterParameterGroupInput) (*redshift.ClusterParameterGroupNameMessage, error)
	ModifyClusterParameterGroupAsync(ctx workflow.Context, input *redshift.ModifyClusterParameterGroupInput) *RedshiftModifyClusterParameterGroupResult

	ModifyClusterSnapshot(ctx workflow.Context, input *redshift.ModifyClusterSnapshotInput) (*redshift.ModifyClusterSnapshotOutput, error)
	ModifyClusterSnapshotAsync(ctx workflow.Context, input *redshift.ModifyClusterSnapshotInput) *RedshiftModifyClusterSnapshotResult

	ModifyClusterSnapshotSchedule(ctx workflow.Context, input *redshift.ModifyClusterSnapshotScheduleInput) (*redshift.ModifyClusterSnapshotScheduleOutput, error)
	ModifyClusterSnapshotScheduleAsync(ctx workflow.Context, input *redshift.ModifyClusterSnapshotScheduleInput) *RedshiftModifyClusterSnapshotScheduleResult

	ModifyClusterSubnetGroup(ctx workflow.Context, input *redshift.ModifyClusterSubnetGroupInput) (*redshift.ModifyClusterSubnetGroupOutput, error)
	ModifyClusterSubnetGroupAsync(ctx workflow.Context, input *redshift.ModifyClusterSubnetGroupInput) *RedshiftModifyClusterSubnetGroupResult

	ModifyEventSubscription(ctx workflow.Context, input *redshift.ModifyEventSubscriptionInput) (*redshift.ModifyEventSubscriptionOutput, error)
	ModifyEventSubscriptionAsync(ctx workflow.Context, input *redshift.ModifyEventSubscriptionInput) *RedshiftModifyEventSubscriptionResult

	ModifyScheduledAction(ctx workflow.Context, input *redshift.ModifyScheduledActionInput) (*redshift.ModifyScheduledActionOutput, error)
	ModifyScheduledActionAsync(ctx workflow.Context, input *redshift.ModifyScheduledActionInput) *RedshiftModifyScheduledActionResult

	ModifySnapshotCopyRetentionPeriod(ctx workflow.Context, input *redshift.ModifySnapshotCopyRetentionPeriodInput) (*redshift.ModifySnapshotCopyRetentionPeriodOutput, error)
	ModifySnapshotCopyRetentionPeriodAsync(ctx workflow.Context, input *redshift.ModifySnapshotCopyRetentionPeriodInput) *RedshiftModifySnapshotCopyRetentionPeriodResult

	ModifySnapshotSchedule(ctx workflow.Context, input *redshift.ModifySnapshotScheduleInput) (*redshift.ModifySnapshotScheduleOutput, error)
	ModifySnapshotScheduleAsync(ctx workflow.Context, input *redshift.ModifySnapshotScheduleInput) *RedshiftModifySnapshotScheduleResult

	ModifyUsageLimit(ctx workflow.Context, input *redshift.ModifyUsageLimitInput) (*redshift.ModifyUsageLimitOutput, error)
	ModifyUsageLimitAsync(ctx workflow.Context, input *redshift.ModifyUsageLimitInput) *RedshiftModifyUsageLimitResult

	PauseCluster(ctx workflow.Context, input *redshift.PauseClusterInput) (*redshift.PauseClusterOutput, error)
	PauseClusterAsync(ctx workflow.Context, input *redshift.PauseClusterInput) *RedshiftPauseClusterResult

	PurchaseReservedNodeOffering(ctx workflow.Context, input *redshift.PurchaseReservedNodeOfferingInput) (*redshift.PurchaseReservedNodeOfferingOutput, error)
	PurchaseReservedNodeOfferingAsync(ctx workflow.Context, input *redshift.PurchaseReservedNodeOfferingInput) *RedshiftPurchaseReservedNodeOfferingResult

	RebootCluster(ctx workflow.Context, input *redshift.RebootClusterInput) (*redshift.RebootClusterOutput, error)
	RebootClusterAsync(ctx workflow.Context, input *redshift.RebootClusterInput) *RedshiftRebootClusterResult

	ResetClusterParameterGroup(ctx workflow.Context, input *redshift.ResetClusterParameterGroupInput) (*redshift.ClusterParameterGroupNameMessage, error)
	ResetClusterParameterGroupAsync(ctx workflow.Context, input *redshift.ResetClusterParameterGroupInput) *RedshiftResetClusterParameterGroupResult

	ResizeCluster(ctx workflow.Context, input *redshift.ResizeClusterInput) (*redshift.ResizeClusterOutput, error)
	ResizeClusterAsync(ctx workflow.Context, input *redshift.ResizeClusterInput) *RedshiftResizeClusterResult

	RestoreFromClusterSnapshot(ctx workflow.Context, input *redshift.RestoreFromClusterSnapshotInput) (*redshift.RestoreFromClusterSnapshotOutput, error)
	RestoreFromClusterSnapshotAsync(ctx workflow.Context, input *redshift.RestoreFromClusterSnapshotInput) *RedshiftRestoreFromClusterSnapshotResult

	RestoreTableFromClusterSnapshot(ctx workflow.Context, input *redshift.RestoreTableFromClusterSnapshotInput) (*redshift.RestoreTableFromClusterSnapshotOutput, error)
	RestoreTableFromClusterSnapshotAsync(ctx workflow.Context, input *redshift.RestoreTableFromClusterSnapshotInput) *RedshiftRestoreTableFromClusterSnapshotResult

	ResumeCluster(ctx workflow.Context, input *redshift.ResumeClusterInput) (*redshift.ResumeClusterOutput, error)
	ResumeClusterAsync(ctx workflow.Context, input *redshift.ResumeClusterInput) *RedshiftResumeClusterResult

	RevokeClusterSecurityGroupIngress(ctx workflow.Context, input *redshift.RevokeClusterSecurityGroupIngressInput) (*redshift.RevokeClusterSecurityGroupIngressOutput, error)
	RevokeClusterSecurityGroupIngressAsync(ctx workflow.Context, input *redshift.RevokeClusterSecurityGroupIngressInput) *RedshiftRevokeClusterSecurityGroupIngressResult

	RevokeSnapshotAccess(ctx workflow.Context, input *redshift.RevokeSnapshotAccessInput) (*redshift.RevokeSnapshotAccessOutput, error)
	RevokeSnapshotAccessAsync(ctx workflow.Context, input *redshift.RevokeSnapshotAccessInput) *RedshiftRevokeSnapshotAccessResult

	RotateEncryptionKey(ctx workflow.Context, input *redshift.RotateEncryptionKeyInput) (*redshift.RotateEncryptionKeyOutput, error)
	RotateEncryptionKeyAsync(ctx workflow.Context, input *redshift.RotateEncryptionKeyInput) *RedshiftRotateEncryptionKeyResult

	WaitUntilClusterAvailable(ctx workflow.Context, input *redshift.DescribeClustersInput) error
	WaitUntilClusterDeleted(ctx workflow.Context, input *redshift.DescribeClustersInput) error
	WaitUntilClusterRestored(ctx workflow.Context, input *redshift.DescribeClustersInput) error
	WaitUntilSnapshotAvailable(ctx workflow.Context, input *redshift.DescribeClusterSnapshotsInput) error
}

type RedshiftAcceptReservedNodeExchangeResult struct {
	Result workflow.Future
}

func (r *RedshiftAcceptReservedNodeExchangeResult) Get(ctx workflow.Context) (*redshift.AcceptReservedNodeExchangeOutput, error) {
	var output redshift.AcceptReservedNodeExchangeOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type RedshiftAuthorizeClusterSecurityGroupIngressResult struct {
	Result workflow.Future
}

func (r *RedshiftAuthorizeClusterSecurityGroupIngressResult) Get(ctx workflow.Context) (*redshift.AuthorizeClusterSecurityGroupIngressOutput, error) {
	var output redshift.AuthorizeClusterSecurityGroupIngressOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type RedshiftAuthorizeSnapshotAccessResult struct {
	Result workflow.Future
}

func (r *RedshiftAuthorizeSnapshotAccessResult) Get(ctx workflow.Context) (*redshift.AuthorizeSnapshotAccessOutput, error) {
	var output redshift.AuthorizeSnapshotAccessOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type RedshiftBatchDeleteClusterSnapshotsResult struct {
	Result workflow.Future
}

func (r *RedshiftBatchDeleteClusterSnapshotsResult) Get(ctx workflow.Context) (*redshift.BatchDeleteClusterSnapshotsOutput, error) {
	var output redshift.BatchDeleteClusterSnapshotsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type RedshiftBatchModifyClusterSnapshotsResult struct {
	Result workflow.Future
}

func (r *RedshiftBatchModifyClusterSnapshotsResult) Get(ctx workflow.Context) (*redshift.BatchModifyClusterSnapshotsOutput, error) {
	var output redshift.BatchModifyClusterSnapshotsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type RedshiftCancelResizeResult struct {
	Result workflow.Future
}

func (r *RedshiftCancelResizeResult) Get(ctx workflow.Context) (*redshift.CancelResizeOutput, error) {
	var output redshift.CancelResizeOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type RedshiftCopyClusterSnapshotResult struct {
	Result workflow.Future
}

func (r *RedshiftCopyClusterSnapshotResult) Get(ctx workflow.Context) (*redshift.CopyClusterSnapshotOutput, error) {
	var output redshift.CopyClusterSnapshotOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type RedshiftCreateClusterResult struct {
	Result workflow.Future
}

func (r *RedshiftCreateClusterResult) Get(ctx workflow.Context) (*redshift.CreateClusterOutput, error) {
	var output redshift.CreateClusterOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type RedshiftCreateClusterParameterGroupResult struct {
	Result workflow.Future
}

func (r *RedshiftCreateClusterParameterGroupResult) Get(ctx workflow.Context) (*redshift.CreateClusterParameterGroupOutput, error) {
	var output redshift.CreateClusterParameterGroupOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type RedshiftCreateClusterSecurityGroupResult struct {
	Result workflow.Future
}

func (r *RedshiftCreateClusterSecurityGroupResult) Get(ctx workflow.Context) (*redshift.CreateClusterSecurityGroupOutput, error) {
	var output redshift.CreateClusterSecurityGroupOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type RedshiftCreateClusterSnapshotResult struct {
	Result workflow.Future
}

func (r *RedshiftCreateClusterSnapshotResult) Get(ctx workflow.Context) (*redshift.CreateClusterSnapshotOutput, error) {
	var output redshift.CreateClusterSnapshotOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type RedshiftCreateClusterSubnetGroupResult struct {
	Result workflow.Future
}

func (r *RedshiftCreateClusterSubnetGroupResult) Get(ctx workflow.Context) (*redshift.CreateClusterSubnetGroupOutput, error) {
	var output redshift.CreateClusterSubnetGroupOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type RedshiftCreateEventSubscriptionResult struct {
	Result workflow.Future
}

func (r *RedshiftCreateEventSubscriptionResult) Get(ctx workflow.Context) (*redshift.CreateEventSubscriptionOutput, error) {
	var output redshift.CreateEventSubscriptionOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type RedshiftCreateHsmClientCertificateResult struct {
	Result workflow.Future
}

func (r *RedshiftCreateHsmClientCertificateResult) Get(ctx workflow.Context) (*redshift.CreateHsmClientCertificateOutput, error) {
	var output redshift.CreateHsmClientCertificateOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type RedshiftCreateHsmConfigurationResult struct {
	Result workflow.Future
}

func (r *RedshiftCreateHsmConfigurationResult) Get(ctx workflow.Context) (*redshift.CreateHsmConfigurationOutput, error) {
	var output redshift.CreateHsmConfigurationOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type RedshiftCreateScheduledActionResult struct {
	Result workflow.Future
}

func (r *RedshiftCreateScheduledActionResult) Get(ctx workflow.Context) (*redshift.CreateScheduledActionOutput, error) {
	var output redshift.CreateScheduledActionOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type RedshiftCreateSnapshotCopyGrantResult struct {
	Result workflow.Future
}

func (r *RedshiftCreateSnapshotCopyGrantResult) Get(ctx workflow.Context) (*redshift.CreateSnapshotCopyGrantOutput, error) {
	var output redshift.CreateSnapshotCopyGrantOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type RedshiftCreateSnapshotScheduleResult struct {
	Result workflow.Future
}

func (r *RedshiftCreateSnapshotScheduleResult) Get(ctx workflow.Context) (*redshift.CreateSnapshotScheduleOutput, error) {
	var output redshift.CreateSnapshotScheduleOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type RedshiftCreateTagsResult struct {
	Result workflow.Future
}

func (r *RedshiftCreateTagsResult) Get(ctx workflow.Context) (*redshift.CreateTagsOutput, error) {
	var output redshift.CreateTagsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type RedshiftCreateUsageLimitResult struct {
	Result workflow.Future
}

func (r *RedshiftCreateUsageLimitResult) Get(ctx workflow.Context) (*redshift.CreateUsageLimitOutput, error) {
	var output redshift.CreateUsageLimitOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type RedshiftDeleteClusterResult struct {
	Result workflow.Future
}

func (r *RedshiftDeleteClusterResult) Get(ctx workflow.Context) (*redshift.DeleteClusterOutput, error) {
	var output redshift.DeleteClusterOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type RedshiftDeleteClusterParameterGroupResult struct {
	Result workflow.Future
}

func (r *RedshiftDeleteClusterParameterGroupResult) Get(ctx workflow.Context) (*redshift.DeleteClusterParameterGroupOutput, error) {
	var output redshift.DeleteClusterParameterGroupOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type RedshiftDeleteClusterSecurityGroupResult struct {
	Result workflow.Future
}

func (r *RedshiftDeleteClusterSecurityGroupResult) Get(ctx workflow.Context) (*redshift.DeleteClusterSecurityGroupOutput, error) {
	var output redshift.DeleteClusterSecurityGroupOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type RedshiftDeleteClusterSnapshotResult struct {
	Result workflow.Future
}

func (r *RedshiftDeleteClusterSnapshotResult) Get(ctx workflow.Context) (*redshift.DeleteClusterSnapshotOutput, error) {
	var output redshift.DeleteClusterSnapshotOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type RedshiftDeleteClusterSubnetGroupResult struct {
	Result workflow.Future
}

func (r *RedshiftDeleteClusterSubnetGroupResult) Get(ctx workflow.Context) (*redshift.DeleteClusterSubnetGroupOutput, error) {
	var output redshift.DeleteClusterSubnetGroupOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type RedshiftDeleteEventSubscriptionResult struct {
	Result workflow.Future
}

func (r *RedshiftDeleteEventSubscriptionResult) Get(ctx workflow.Context) (*redshift.DeleteEventSubscriptionOutput, error) {
	var output redshift.DeleteEventSubscriptionOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type RedshiftDeleteHsmClientCertificateResult struct {
	Result workflow.Future
}

func (r *RedshiftDeleteHsmClientCertificateResult) Get(ctx workflow.Context) (*redshift.DeleteHsmClientCertificateOutput, error) {
	var output redshift.DeleteHsmClientCertificateOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type RedshiftDeleteHsmConfigurationResult struct {
	Result workflow.Future
}

func (r *RedshiftDeleteHsmConfigurationResult) Get(ctx workflow.Context) (*redshift.DeleteHsmConfigurationOutput, error) {
	var output redshift.DeleteHsmConfigurationOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type RedshiftDeleteScheduledActionResult struct {
	Result workflow.Future
}

func (r *RedshiftDeleteScheduledActionResult) Get(ctx workflow.Context) (*redshift.DeleteScheduledActionOutput, error) {
	var output redshift.DeleteScheduledActionOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type RedshiftDeleteSnapshotCopyGrantResult struct {
	Result workflow.Future
}

func (r *RedshiftDeleteSnapshotCopyGrantResult) Get(ctx workflow.Context) (*redshift.DeleteSnapshotCopyGrantOutput, error) {
	var output redshift.DeleteSnapshotCopyGrantOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type RedshiftDeleteSnapshotScheduleResult struct {
	Result workflow.Future
}

func (r *RedshiftDeleteSnapshotScheduleResult) Get(ctx workflow.Context) (*redshift.DeleteSnapshotScheduleOutput, error) {
	var output redshift.DeleteSnapshotScheduleOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type RedshiftDeleteTagsResult struct {
	Result workflow.Future
}

func (r *RedshiftDeleteTagsResult) Get(ctx workflow.Context) (*redshift.DeleteTagsOutput, error) {
	var output redshift.DeleteTagsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type RedshiftDeleteUsageLimitResult struct {
	Result workflow.Future
}

func (r *RedshiftDeleteUsageLimitResult) Get(ctx workflow.Context) (*redshift.DeleteUsageLimitOutput, error) {
	var output redshift.DeleteUsageLimitOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type RedshiftDescribeAccountAttributesResult struct {
	Result workflow.Future
}

func (r *RedshiftDescribeAccountAttributesResult) Get(ctx workflow.Context) (*redshift.DescribeAccountAttributesOutput, error) {
	var output redshift.DescribeAccountAttributesOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type RedshiftDescribeClusterDbRevisionsResult struct {
	Result workflow.Future
}

func (r *RedshiftDescribeClusterDbRevisionsResult) Get(ctx workflow.Context) (*redshift.DescribeClusterDbRevisionsOutput, error) {
	var output redshift.DescribeClusterDbRevisionsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type RedshiftDescribeClusterParameterGroupsResult struct {
	Result workflow.Future
}

func (r *RedshiftDescribeClusterParameterGroupsResult) Get(ctx workflow.Context) (*redshift.DescribeClusterParameterGroupsOutput, error) {
	var output redshift.DescribeClusterParameterGroupsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type RedshiftDescribeClusterParametersResult struct {
	Result workflow.Future
}

func (r *RedshiftDescribeClusterParametersResult) Get(ctx workflow.Context) (*redshift.DescribeClusterParametersOutput, error) {
	var output redshift.DescribeClusterParametersOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type RedshiftDescribeClusterSecurityGroupsResult struct {
	Result workflow.Future
}

func (r *RedshiftDescribeClusterSecurityGroupsResult) Get(ctx workflow.Context) (*redshift.DescribeClusterSecurityGroupsOutput, error) {
	var output redshift.DescribeClusterSecurityGroupsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type RedshiftDescribeClusterSnapshotsResult struct {
	Result workflow.Future
}

func (r *RedshiftDescribeClusterSnapshotsResult) Get(ctx workflow.Context) (*redshift.DescribeClusterSnapshotsOutput, error) {
	var output redshift.DescribeClusterSnapshotsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type RedshiftDescribeClusterSubnetGroupsResult struct {
	Result workflow.Future
}

func (r *RedshiftDescribeClusterSubnetGroupsResult) Get(ctx workflow.Context) (*redshift.DescribeClusterSubnetGroupsOutput, error) {
	var output redshift.DescribeClusterSubnetGroupsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type RedshiftDescribeClusterTracksResult struct {
	Result workflow.Future
}

func (r *RedshiftDescribeClusterTracksResult) Get(ctx workflow.Context) (*redshift.DescribeClusterTracksOutput, error) {
	var output redshift.DescribeClusterTracksOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type RedshiftDescribeClusterVersionsResult struct {
	Result workflow.Future
}

func (r *RedshiftDescribeClusterVersionsResult) Get(ctx workflow.Context) (*redshift.DescribeClusterVersionsOutput, error) {
	var output redshift.DescribeClusterVersionsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type RedshiftDescribeClustersResult struct {
	Result workflow.Future
}

func (r *RedshiftDescribeClustersResult) Get(ctx workflow.Context) (*redshift.DescribeClustersOutput, error) {
	var output redshift.DescribeClustersOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type RedshiftDescribeDefaultClusterParametersResult struct {
	Result workflow.Future
}

func (r *RedshiftDescribeDefaultClusterParametersResult) Get(ctx workflow.Context) (*redshift.DescribeDefaultClusterParametersOutput, error) {
	var output redshift.DescribeDefaultClusterParametersOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type RedshiftDescribeEventCategoriesResult struct {
	Result workflow.Future
}

func (r *RedshiftDescribeEventCategoriesResult) Get(ctx workflow.Context) (*redshift.DescribeEventCategoriesOutput, error) {
	var output redshift.DescribeEventCategoriesOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type RedshiftDescribeEventSubscriptionsResult struct {
	Result workflow.Future
}

func (r *RedshiftDescribeEventSubscriptionsResult) Get(ctx workflow.Context) (*redshift.DescribeEventSubscriptionsOutput, error) {
	var output redshift.DescribeEventSubscriptionsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type RedshiftDescribeEventsResult struct {
	Result workflow.Future
}

func (r *RedshiftDescribeEventsResult) Get(ctx workflow.Context) (*redshift.DescribeEventsOutput, error) {
	var output redshift.DescribeEventsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type RedshiftDescribeHsmClientCertificatesResult struct {
	Result workflow.Future
}

func (r *RedshiftDescribeHsmClientCertificatesResult) Get(ctx workflow.Context) (*redshift.DescribeHsmClientCertificatesOutput, error) {
	var output redshift.DescribeHsmClientCertificatesOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type RedshiftDescribeHsmConfigurationsResult struct {
	Result workflow.Future
}

func (r *RedshiftDescribeHsmConfigurationsResult) Get(ctx workflow.Context) (*redshift.DescribeHsmConfigurationsOutput, error) {
	var output redshift.DescribeHsmConfigurationsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type RedshiftDescribeLoggingStatusResult struct {
	Result workflow.Future
}

func (r *RedshiftDescribeLoggingStatusResult) Get(ctx workflow.Context) (*redshift.LoggingStatus, error) {
	var output redshift.LoggingStatus
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type RedshiftDescribeNodeConfigurationOptionsResult struct {
	Result workflow.Future
}

func (r *RedshiftDescribeNodeConfigurationOptionsResult) Get(ctx workflow.Context) (*redshift.DescribeNodeConfigurationOptionsOutput, error) {
	var output redshift.DescribeNodeConfigurationOptionsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type RedshiftDescribeOrderableClusterOptionsResult struct {
	Result workflow.Future
}

func (r *RedshiftDescribeOrderableClusterOptionsResult) Get(ctx workflow.Context) (*redshift.DescribeOrderableClusterOptionsOutput, error) {
	var output redshift.DescribeOrderableClusterOptionsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type RedshiftDescribeReservedNodeOfferingsResult struct {
	Result workflow.Future
}

func (r *RedshiftDescribeReservedNodeOfferingsResult) Get(ctx workflow.Context) (*redshift.DescribeReservedNodeOfferingsOutput, error) {
	var output redshift.DescribeReservedNodeOfferingsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type RedshiftDescribeReservedNodesResult struct {
	Result workflow.Future
}

func (r *RedshiftDescribeReservedNodesResult) Get(ctx workflow.Context) (*redshift.DescribeReservedNodesOutput, error) {
	var output redshift.DescribeReservedNodesOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type RedshiftDescribeResizeResult struct {
	Result workflow.Future
}

func (r *RedshiftDescribeResizeResult) Get(ctx workflow.Context) (*redshift.DescribeResizeOutput, error) {
	var output redshift.DescribeResizeOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type RedshiftDescribeScheduledActionsResult struct {
	Result workflow.Future
}

func (r *RedshiftDescribeScheduledActionsResult) Get(ctx workflow.Context) (*redshift.DescribeScheduledActionsOutput, error) {
	var output redshift.DescribeScheduledActionsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type RedshiftDescribeSnapshotCopyGrantsResult struct {
	Result workflow.Future
}

func (r *RedshiftDescribeSnapshotCopyGrantsResult) Get(ctx workflow.Context) (*redshift.DescribeSnapshotCopyGrantsOutput, error) {
	var output redshift.DescribeSnapshotCopyGrantsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type RedshiftDescribeSnapshotSchedulesResult struct {
	Result workflow.Future
}

func (r *RedshiftDescribeSnapshotSchedulesResult) Get(ctx workflow.Context) (*redshift.DescribeSnapshotSchedulesOutput, error) {
	var output redshift.DescribeSnapshotSchedulesOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type RedshiftDescribeStorageResult struct {
	Result workflow.Future
}

func (r *RedshiftDescribeStorageResult) Get(ctx workflow.Context) (*redshift.DescribeStorageOutput, error) {
	var output redshift.DescribeStorageOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type RedshiftDescribeTableRestoreStatusResult struct {
	Result workflow.Future
}

func (r *RedshiftDescribeTableRestoreStatusResult) Get(ctx workflow.Context) (*redshift.DescribeTableRestoreStatusOutput, error) {
	var output redshift.DescribeTableRestoreStatusOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type RedshiftDescribeTagsResult struct {
	Result workflow.Future
}

func (r *RedshiftDescribeTagsResult) Get(ctx workflow.Context) (*redshift.DescribeTagsOutput, error) {
	var output redshift.DescribeTagsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type RedshiftDescribeUsageLimitsResult struct {
	Result workflow.Future
}

func (r *RedshiftDescribeUsageLimitsResult) Get(ctx workflow.Context) (*redshift.DescribeUsageLimitsOutput, error) {
	var output redshift.DescribeUsageLimitsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type RedshiftDisableLoggingResult struct {
	Result workflow.Future
}

func (r *RedshiftDisableLoggingResult) Get(ctx workflow.Context) (*redshift.LoggingStatus, error) {
	var output redshift.LoggingStatus
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type RedshiftDisableSnapshotCopyResult struct {
	Result workflow.Future
}

func (r *RedshiftDisableSnapshotCopyResult) Get(ctx workflow.Context) (*redshift.DisableSnapshotCopyOutput, error) {
	var output redshift.DisableSnapshotCopyOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type RedshiftEnableLoggingResult struct {
	Result workflow.Future
}

func (r *RedshiftEnableLoggingResult) Get(ctx workflow.Context) (*redshift.LoggingStatus, error) {
	var output redshift.LoggingStatus
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type RedshiftEnableSnapshotCopyResult struct {
	Result workflow.Future
}

func (r *RedshiftEnableSnapshotCopyResult) Get(ctx workflow.Context) (*redshift.EnableSnapshotCopyOutput, error) {
	var output redshift.EnableSnapshotCopyOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type RedshiftGetClusterCredentialsResult struct {
	Result workflow.Future
}

func (r *RedshiftGetClusterCredentialsResult) Get(ctx workflow.Context) (*redshift.GetClusterCredentialsOutput, error) {
	var output redshift.GetClusterCredentialsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type RedshiftGetReservedNodeExchangeOfferingsResult struct {
	Result workflow.Future
}

func (r *RedshiftGetReservedNodeExchangeOfferingsResult) Get(ctx workflow.Context) (*redshift.GetReservedNodeExchangeOfferingsOutput, error) {
	var output redshift.GetReservedNodeExchangeOfferingsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type RedshiftModifyClusterResult struct {
	Result workflow.Future
}

func (r *RedshiftModifyClusterResult) Get(ctx workflow.Context) (*redshift.ModifyClusterOutput, error) {
	var output redshift.ModifyClusterOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type RedshiftModifyClusterDbRevisionResult struct {
	Result workflow.Future
}

func (r *RedshiftModifyClusterDbRevisionResult) Get(ctx workflow.Context) (*redshift.ModifyClusterDbRevisionOutput, error) {
	var output redshift.ModifyClusterDbRevisionOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type RedshiftModifyClusterIamRolesResult struct {
	Result workflow.Future
}

func (r *RedshiftModifyClusterIamRolesResult) Get(ctx workflow.Context) (*redshift.ModifyClusterIamRolesOutput, error) {
	var output redshift.ModifyClusterIamRolesOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type RedshiftModifyClusterMaintenanceResult struct {
	Result workflow.Future
}

func (r *RedshiftModifyClusterMaintenanceResult) Get(ctx workflow.Context) (*redshift.ModifyClusterMaintenanceOutput, error) {
	var output redshift.ModifyClusterMaintenanceOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type RedshiftModifyClusterParameterGroupResult struct {
	Result workflow.Future
}

func (r *RedshiftModifyClusterParameterGroupResult) Get(ctx workflow.Context) (*redshift.ClusterParameterGroupNameMessage, error) {
	var output redshift.ClusterParameterGroupNameMessage
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type RedshiftModifyClusterSnapshotResult struct {
	Result workflow.Future
}

func (r *RedshiftModifyClusterSnapshotResult) Get(ctx workflow.Context) (*redshift.ModifyClusterSnapshotOutput, error) {
	var output redshift.ModifyClusterSnapshotOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type RedshiftModifyClusterSnapshotScheduleResult struct {
	Result workflow.Future
}

func (r *RedshiftModifyClusterSnapshotScheduleResult) Get(ctx workflow.Context) (*redshift.ModifyClusterSnapshotScheduleOutput, error) {
	var output redshift.ModifyClusterSnapshotScheduleOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type RedshiftModifyClusterSubnetGroupResult struct {
	Result workflow.Future
}

func (r *RedshiftModifyClusterSubnetGroupResult) Get(ctx workflow.Context) (*redshift.ModifyClusterSubnetGroupOutput, error) {
	var output redshift.ModifyClusterSubnetGroupOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type RedshiftModifyEventSubscriptionResult struct {
	Result workflow.Future
}

func (r *RedshiftModifyEventSubscriptionResult) Get(ctx workflow.Context) (*redshift.ModifyEventSubscriptionOutput, error) {
	var output redshift.ModifyEventSubscriptionOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type RedshiftModifyScheduledActionResult struct {
	Result workflow.Future
}

func (r *RedshiftModifyScheduledActionResult) Get(ctx workflow.Context) (*redshift.ModifyScheduledActionOutput, error) {
	var output redshift.ModifyScheduledActionOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type RedshiftModifySnapshotCopyRetentionPeriodResult struct {
	Result workflow.Future
}

func (r *RedshiftModifySnapshotCopyRetentionPeriodResult) Get(ctx workflow.Context) (*redshift.ModifySnapshotCopyRetentionPeriodOutput, error) {
	var output redshift.ModifySnapshotCopyRetentionPeriodOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type RedshiftModifySnapshotScheduleResult struct {
	Result workflow.Future
}

func (r *RedshiftModifySnapshotScheduleResult) Get(ctx workflow.Context) (*redshift.ModifySnapshotScheduleOutput, error) {
	var output redshift.ModifySnapshotScheduleOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type RedshiftModifyUsageLimitResult struct {
	Result workflow.Future
}

func (r *RedshiftModifyUsageLimitResult) Get(ctx workflow.Context) (*redshift.ModifyUsageLimitOutput, error) {
	var output redshift.ModifyUsageLimitOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type RedshiftPauseClusterResult struct {
	Result workflow.Future
}

func (r *RedshiftPauseClusterResult) Get(ctx workflow.Context) (*redshift.PauseClusterOutput, error) {
	var output redshift.PauseClusterOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type RedshiftPurchaseReservedNodeOfferingResult struct {
	Result workflow.Future
}

func (r *RedshiftPurchaseReservedNodeOfferingResult) Get(ctx workflow.Context) (*redshift.PurchaseReservedNodeOfferingOutput, error) {
	var output redshift.PurchaseReservedNodeOfferingOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type RedshiftRebootClusterResult struct {
	Result workflow.Future
}

func (r *RedshiftRebootClusterResult) Get(ctx workflow.Context) (*redshift.RebootClusterOutput, error) {
	var output redshift.RebootClusterOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type RedshiftResetClusterParameterGroupResult struct {
	Result workflow.Future
}

func (r *RedshiftResetClusterParameterGroupResult) Get(ctx workflow.Context) (*redshift.ClusterParameterGroupNameMessage, error) {
	var output redshift.ClusterParameterGroupNameMessage
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type RedshiftResizeClusterResult struct {
	Result workflow.Future
}

func (r *RedshiftResizeClusterResult) Get(ctx workflow.Context) (*redshift.ResizeClusterOutput, error) {
	var output redshift.ResizeClusterOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type RedshiftRestoreFromClusterSnapshotResult struct {
	Result workflow.Future
}

func (r *RedshiftRestoreFromClusterSnapshotResult) Get(ctx workflow.Context) (*redshift.RestoreFromClusterSnapshotOutput, error) {
	var output redshift.RestoreFromClusterSnapshotOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type RedshiftRestoreTableFromClusterSnapshotResult struct {
	Result workflow.Future
}

func (r *RedshiftRestoreTableFromClusterSnapshotResult) Get(ctx workflow.Context) (*redshift.RestoreTableFromClusterSnapshotOutput, error) {
	var output redshift.RestoreTableFromClusterSnapshotOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type RedshiftResumeClusterResult struct {
	Result workflow.Future
}

func (r *RedshiftResumeClusterResult) Get(ctx workflow.Context) (*redshift.ResumeClusterOutput, error) {
	var output redshift.ResumeClusterOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type RedshiftRevokeClusterSecurityGroupIngressResult struct {
	Result workflow.Future
}

func (r *RedshiftRevokeClusterSecurityGroupIngressResult) Get(ctx workflow.Context) (*redshift.RevokeClusterSecurityGroupIngressOutput, error) {
	var output redshift.RevokeClusterSecurityGroupIngressOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type RedshiftRevokeSnapshotAccessResult struct {
	Result workflow.Future
}

func (r *RedshiftRevokeSnapshotAccessResult) Get(ctx workflow.Context) (*redshift.RevokeSnapshotAccessOutput, error) {
	var output redshift.RevokeSnapshotAccessOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type RedshiftRotateEncryptionKeyResult struct {
	Result workflow.Future
}

func (r *RedshiftRotateEncryptionKeyResult) Get(ctx workflow.Context) (*redshift.RotateEncryptionKeyOutput, error) {
	var output redshift.RotateEncryptionKeyOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type RedshiftStub struct {
	activities awsactivities.RedshiftActivities
}

func NewRedshiftStub() RedshiftClient {
	return &RedshiftStub{}
}

func (a *RedshiftStub) AcceptReservedNodeExchange(ctx workflow.Context, input *redshift.AcceptReservedNodeExchangeInput) (*redshift.AcceptReservedNodeExchangeOutput, error) {
	var output redshift.AcceptReservedNodeExchangeOutput
	err := workflow.ExecuteActivity(ctx, a.activities.AcceptReservedNodeExchange, input).Get(ctx, &output)
	return &output, err
}

func (a *RedshiftStub) AcceptReservedNodeExchangeAsync(ctx workflow.Context, input *redshift.AcceptReservedNodeExchangeInput) *RedshiftAcceptReservedNodeExchangeResult {
	future := workflow.ExecuteActivity(ctx, a.activities.AcceptReservedNodeExchange, input)
	return &RedshiftAcceptReservedNodeExchangeResult{Result: future}
}

func (a *RedshiftStub) AuthorizeClusterSecurityGroupIngress(ctx workflow.Context, input *redshift.AuthorizeClusterSecurityGroupIngressInput) (*redshift.AuthorizeClusterSecurityGroupIngressOutput, error) {
	var output redshift.AuthorizeClusterSecurityGroupIngressOutput
	err := workflow.ExecuteActivity(ctx, a.activities.AuthorizeClusterSecurityGroupIngress, input).Get(ctx, &output)
	return &output, err
}

func (a *RedshiftStub) AuthorizeClusterSecurityGroupIngressAsync(ctx workflow.Context, input *redshift.AuthorizeClusterSecurityGroupIngressInput) *RedshiftAuthorizeClusterSecurityGroupIngressResult {
	future := workflow.ExecuteActivity(ctx, a.activities.AuthorizeClusterSecurityGroupIngress, input)
	return &RedshiftAuthorizeClusterSecurityGroupIngressResult{Result: future}
}

func (a *RedshiftStub) AuthorizeSnapshotAccess(ctx workflow.Context, input *redshift.AuthorizeSnapshotAccessInput) (*redshift.AuthorizeSnapshotAccessOutput, error) {
	var output redshift.AuthorizeSnapshotAccessOutput
	err := workflow.ExecuteActivity(ctx, a.activities.AuthorizeSnapshotAccess, input).Get(ctx, &output)
	return &output, err
}

func (a *RedshiftStub) AuthorizeSnapshotAccessAsync(ctx workflow.Context, input *redshift.AuthorizeSnapshotAccessInput) *RedshiftAuthorizeSnapshotAccessResult {
	future := workflow.ExecuteActivity(ctx, a.activities.AuthorizeSnapshotAccess, input)
	return &RedshiftAuthorizeSnapshotAccessResult{Result: future}
}

func (a *RedshiftStub) BatchDeleteClusterSnapshots(ctx workflow.Context, input *redshift.BatchDeleteClusterSnapshotsInput) (*redshift.BatchDeleteClusterSnapshotsOutput, error) {
	var output redshift.BatchDeleteClusterSnapshotsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.BatchDeleteClusterSnapshots, input).Get(ctx, &output)
	return &output, err
}

func (a *RedshiftStub) BatchDeleteClusterSnapshotsAsync(ctx workflow.Context, input *redshift.BatchDeleteClusterSnapshotsInput) *RedshiftBatchDeleteClusterSnapshotsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.BatchDeleteClusterSnapshots, input)
	return &RedshiftBatchDeleteClusterSnapshotsResult{Result: future}
}

func (a *RedshiftStub) BatchModifyClusterSnapshots(ctx workflow.Context, input *redshift.BatchModifyClusterSnapshotsInput) (*redshift.BatchModifyClusterSnapshotsOutput, error) {
	var output redshift.BatchModifyClusterSnapshotsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.BatchModifyClusterSnapshots, input).Get(ctx, &output)
	return &output, err
}

func (a *RedshiftStub) BatchModifyClusterSnapshotsAsync(ctx workflow.Context, input *redshift.BatchModifyClusterSnapshotsInput) *RedshiftBatchModifyClusterSnapshotsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.BatchModifyClusterSnapshots, input)
	return &RedshiftBatchModifyClusterSnapshotsResult{Result: future}
}

func (a *RedshiftStub) CancelResize(ctx workflow.Context, input *redshift.CancelResizeInput) (*redshift.CancelResizeOutput, error) {
	var output redshift.CancelResizeOutput
	err := workflow.ExecuteActivity(ctx, a.activities.CancelResize, input).Get(ctx, &output)
	return &output, err
}

func (a *RedshiftStub) CancelResizeAsync(ctx workflow.Context, input *redshift.CancelResizeInput) *RedshiftCancelResizeResult {
	future := workflow.ExecuteActivity(ctx, a.activities.CancelResize, input)
	return &RedshiftCancelResizeResult{Result: future}
}

func (a *RedshiftStub) CopyClusterSnapshot(ctx workflow.Context, input *redshift.CopyClusterSnapshotInput) (*redshift.CopyClusterSnapshotOutput, error) {
	var output redshift.CopyClusterSnapshotOutput
	err := workflow.ExecuteActivity(ctx, a.activities.CopyClusterSnapshot, input).Get(ctx, &output)
	return &output, err
}

func (a *RedshiftStub) CopyClusterSnapshotAsync(ctx workflow.Context, input *redshift.CopyClusterSnapshotInput) *RedshiftCopyClusterSnapshotResult {
	future := workflow.ExecuteActivity(ctx, a.activities.CopyClusterSnapshot, input)
	return &RedshiftCopyClusterSnapshotResult{Result: future}
}

func (a *RedshiftStub) CreateCluster(ctx workflow.Context, input *redshift.CreateClusterInput) (*redshift.CreateClusterOutput, error) {
	var output redshift.CreateClusterOutput
	err := workflow.ExecuteActivity(ctx, a.activities.CreateCluster, input).Get(ctx, &output)
	return &output, err
}

func (a *RedshiftStub) CreateClusterAsync(ctx workflow.Context, input *redshift.CreateClusterInput) *RedshiftCreateClusterResult {
	future := workflow.ExecuteActivity(ctx, a.activities.CreateCluster, input)
	return &RedshiftCreateClusterResult{Result: future}
}

func (a *RedshiftStub) CreateClusterParameterGroup(ctx workflow.Context, input *redshift.CreateClusterParameterGroupInput) (*redshift.CreateClusterParameterGroupOutput, error) {
	var output redshift.CreateClusterParameterGroupOutput
	err := workflow.ExecuteActivity(ctx, a.activities.CreateClusterParameterGroup, input).Get(ctx, &output)
	return &output, err
}

func (a *RedshiftStub) CreateClusterParameterGroupAsync(ctx workflow.Context, input *redshift.CreateClusterParameterGroupInput) *RedshiftCreateClusterParameterGroupResult {
	future := workflow.ExecuteActivity(ctx, a.activities.CreateClusterParameterGroup, input)
	return &RedshiftCreateClusterParameterGroupResult{Result: future}
}

func (a *RedshiftStub) CreateClusterSecurityGroup(ctx workflow.Context, input *redshift.CreateClusterSecurityGroupInput) (*redshift.CreateClusterSecurityGroupOutput, error) {
	var output redshift.CreateClusterSecurityGroupOutput
	err := workflow.ExecuteActivity(ctx, a.activities.CreateClusterSecurityGroup, input).Get(ctx, &output)
	return &output, err
}

func (a *RedshiftStub) CreateClusterSecurityGroupAsync(ctx workflow.Context, input *redshift.CreateClusterSecurityGroupInput) *RedshiftCreateClusterSecurityGroupResult {
	future := workflow.ExecuteActivity(ctx, a.activities.CreateClusterSecurityGroup, input)
	return &RedshiftCreateClusterSecurityGroupResult{Result: future}
}

func (a *RedshiftStub) CreateClusterSnapshot(ctx workflow.Context, input *redshift.CreateClusterSnapshotInput) (*redshift.CreateClusterSnapshotOutput, error) {
	var output redshift.CreateClusterSnapshotOutput
	err := workflow.ExecuteActivity(ctx, a.activities.CreateClusterSnapshot, input).Get(ctx, &output)
	return &output, err
}

func (a *RedshiftStub) CreateClusterSnapshotAsync(ctx workflow.Context, input *redshift.CreateClusterSnapshotInput) *RedshiftCreateClusterSnapshotResult {
	future := workflow.ExecuteActivity(ctx, a.activities.CreateClusterSnapshot, input)
	return &RedshiftCreateClusterSnapshotResult{Result: future}
}

func (a *RedshiftStub) CreateClusterSubnetGroup(ctx workflow.Context, input *redshift.CreateClusterSubnetGroupInput) (*redshift.CreateClusterSubnetGroupOutput, error) {
	var output redshift.CreateClusterSubnetGroupOutput
	err := workflow.ExecuteActivity(ctx, a.activities.CreateClusterSubnetGroup, input).Get(ctx, &output)
	return &output, err
}

func (a *RedshiftStub) CreateClusterSubnetGroupAsync(ctx workflow.Context, input *redshift.CreateClusterSubnetGroupInput) *RedshiftCreateClusterSubnetGroupResult {
	future := workflow.ExecuteActivity(ctx, a.activities.CreateClusterSubnetGroup, input)
	return &RedshiftCreateClusterSubnetGroupResult{Result: future}
}

func (a *RedshiftStub) CreateEventSubscription(ctx workflow.Context, input *redshift.CreateEventSubscriptionInput) (*redshift.CreateEventSubscriptionOutput, error) {
	var output redshift.CreateEventSubscriptionOutput
	err := workflow.ExecuteActivity(ctx, a.activities.CreateEventSubscription, input).Get(ctx, &output)
	return &output, err
}

func (a *RedshiftStub) CreateEventSubscriptionAsync(ctx workflow.Context, input *redshift.CreateEventSubscriptionInput) *RedshiftCreateEventSubscriptionResult {
	future := workflow.ExecuteActivity(ctx, a.activities.CreateEventSubscription, input)
	return &RedshiftCreateEventSubscriptionResult{Result: future}
}

func (a *RedshiftStub) CreateHsmClientCertificate(ctx workflow.Context, input *redshift.CreateHsmClientCertificateInput) (*redshift.CreateHsmClientCertificateOutput, error) {
	var output redshift.CreateHsmClientCertificateOutput
	err := workflow.ExecuteActivity(ctx, a.activities.CreateHsmClientCertificate, input).Get(ctx, &output)
	return &output, err
}

func (a *RedshiftStub) CreateHsmClientCertificateAsync(ctx workflow.Context, input *redshift.CreateHsmClientCertificateInput) *RedshiftCreateHsmClientCertificateResult {
	future := workflow.ExecuteActivity(ctx, a.activities.CreateHsmClientCertificate, input)
	return &RedshiftCreateHsmClientCertificateResult{Result: future}
}

func (a *RedshiftStub) CreateHsmConfiguration(ctx workflow.Context, input *redshift.CreateHsmConfigurationInput) (*redshift.CreateHsmConfigurationOutput, error) {
	var output redshift.CreateHsmConfigurationOutput
	err := workflow.ExecuteActivity(ctx, a.activities.CreateHsmConfiguration, input).Get(ctx, &output)
	return &output, err
}

func (a *RedshiftStub) CreateHsmConfigurationAsync(ctx workflow.Context, input *redshift.CreateHsmConfigurationInput) *RedshiftCreateHsmConfigurationResult {
	future := workflow.ExecuteActivity(ctx, a.activities.CreateHsmConfiguration, input)
	return &RedshiftCreateHsmConfigurationResult{Result: future}
}

func (a *RedshiftStub) CreateScheduledAction(ctx workflow.Context, input *redshift.CreateScheduledActionInput) (*redshift.CreateScheduledActionOutput, error) {
	var output redshift.CreateScheduledActionOutput
	err := workflow.ExecuteActivity(ctx, a.activities.CreateScheduledAction, input).Get(ctx, &output)
	return &output, err
}

func (a *RedshiftStub) CreateScheduledActionAsync(ctx workflow.Context, input *redshift.CreateScheduledActionInput) *RedshiftCreateScheduledActionResult {
	future := workflow.ExecuteActivity(ctx, a.activities.CreateScheduledAction, input)
	return &RedshiftCreateScheduledActionResult{Result: future}
}

func (a *RedshiftStub) CreateSnapshotCopyGrant(ctx workflow.Context, input *redshift.CreateSnapshotCopyGrantInput) (*redshift.CreateSnapshotCopyGrantOutput, error) {
	var output redshift.CreateSnapshotCopyGrantOutput
	err := workflow.ExecuteActivity(ctx, a.activities.CreateSnapshotCopyGrant, input).Get(ctx, &output)
	return &output, err
}

func (a *RedshiftStub) CreateSnapshotCopyGrantAsync(ctx workflow.Context, input *redshift.CreateSnapshotCopyGrantInput) *RedshiftCreateSnapshotCopyGrantResult {
	future := workflow.ExecuteActivity(ctx, a.activities.CreateSnapshotCopyGrant, input)
	return &RedshiftCreateSnapshotCopyGrantResult{Result: future}
}

func (a *RedshiftStub) CreateSnapshotSchedule(ctx workflow.Context, input *redshift.CreateSnapshotScheduleInput) (*redshift.CreateSnapshotScheduleOutput, error) {
	var output redshift.CreateSnapshotScheduleOutput
	err := workflow.ExecuteActivity(ctx, a.activities.CreateSnapshotSchedule, input).Get(ctx, &output)
	return &output, err
}

func (a *RedshiftStub) CreateSnapshotScheduleAsync(ctx workflow.Context, input *redshift.CreateSnapshotScheduleInput) *RedshiftCreateSnapshotScheduleResult {
	future := workflow.ExecuteActivity(ctx, a.activities.CreateSnapshotSchedule, input)
	return &RedshiftCreateSnapshotScheduleResult{Result: future}
}

func (a *RedshiftStub) CreateTags(ctx workflow.Context, input *redshift.CreateTagsInput) (*redshift.CreateTagsOutput, error) {
	var output redshift.CreateTagsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.CreateTags, input).Get(ctx, &output)
	return &output, err
}

func (a *RedshiftStub) CreateTagsAsync(ctx workflow.Context, input *redshift.CreateTagsInput) *RedshiftCreateTagsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.CreateTags, input)
	return &RedshiftCreateTagsResult{Result: future}
}

func (a *RedshiftStub) CreateUsageLimit(ctx workflow.Context, input *redshift.CreateUsageLimitInput) (*redshift.CreateUsageLimitOutput, error) {
	var output redshift.CreateUsageLimitOutput
	err := workflow.ExecuteActivity(ctx, a.activities.CreateUsageLimit, input).Get(ctx, &output)
	return &output, err
}

func (a *RedshiftStub) CreateUsageLimitAsync(ctx workflow.Context, input *redshift.CreateUsageLimitInput) *RedshiftCreateUsageLimitResult {
	future := workflow.ExecuteActivity(ctx, a.activities.CreateUsageLimit, input)
	return &RedshiftCreateUsageLimitResult{Result: future}
}

func (a *RedshiftStub) DeleteCluster(ctx workflow.Context, input *redshift.DeleteClusterInput) (*redshift.DeleteClusterOutput, error) {
	var output redshift.DeleteClusterOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeleteCluster, input).Get(ctx, &output)
	return &output, err
}

func (a *RedshiftStub) DeleteClusterAsync(ctx workflow.Context, input *redshift.DeleteClusterInput) *RedshiftDeleteClusterResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeleteCluster, input)
	return &RedshiftDeleteClusterResult{Result: future}
}

func (a *RedshiftStub) DeleteClusterParameterGroup(ctx workflow.Context, input *redshift.DeleteClusterParameterGroupInput) (*redshift.DeleteClusterParameterGroupOutput, error) {
	var output redshift.DeleteClusterParameterGroupOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeleteClusterParameterGroup, input).Get(ctx, &output)
	return &output, err
}

func (a *RedshiftStub) DeleteClusterParameterGroupAsync(ctx workflow.Context, input *redshift.DeleteClusterParameterGroupInput) *RedshiftDeleteClusterParameterGroupResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeleteClusterParameterGroup, input)
	return &RedshiftDeleteClusterParameterGroupResult{Result: future}
}

func (a *RedshiftStub) DeleteClusterSecurityGroup(ctx workflow.Context, input *redshift.DeleteClusterSecurityGroupInput) (*redshift.DeleteClusterSecurityGroupOutput, error) {
	var output redshift.DeleteClusterSecurityGroupOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeleteClusterSecurityGroup, input).Get(ctx, &output)
	return &output, err
}

func (a *RedshiftStub) DeleteClusterSecurityGroupAsync(ctx workflow.Context, input *redshift.DeleteClusterSecurityGroupInput) *RedshiftDeleteClusterSecurityGroupResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeleteClusterSecurityGroup, input)
	return &RedshiftDeleteClusterSecurityGroupResult{Result: future}
}

func (a *RedshiftStub) DeleteClusterSnapshot(ctx workflow.Context, input *redshift.DeleteClusterSnapshotInput) (*redshift.DeleteClusterSnapshotOutput, error) {
	var output redshift.DeleteClusterSnapshotOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeleteClusterSnapshot, input).Get(ctx, &output)
	return &output, err
}

func (a *RedshiftStub) DeleteClusterSnapshotAsync(ctx workflow.Context, input *redshift.DeleteClusterSnapshotInput) *RedshiftDeleteClusterSnapshotResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeleteClusterSnapshot, input)
	return &RedshiftDeleteClusterSnapshotResult{Result: future}
}

func (a *RedshiftStub) DeleteClusterSubnetGroup(ctx workflow.Context, input *redshift.DeleteClusterSubnetGroupInput) (*redshift.DeleteClusterSubnetGroupOutput, error) {
	var output redshift.DeleteClusterSubnetGroupOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeleteClusterSubnetGroup, input).Get(ctx, &output)
	return &output, err
}

func (a *RedshiftStub) DeleteClusterSubnetGroupAsync(ctx workflow.Context, input *redshift.DeleteClusterSubnetGroupInput) *RedshiftDeleteClusterSubnetGroupResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeleteClusterSubnetGroup, input)
	return &RedshiftDeleteClusterSubnetGroupResult{Result: future}
}

func (a *RedshiftStub) DeleteEventSubscription(ctx workflow.Context, input *redshift.DeleteEventSubscriptionInput) (*redshift.DeleteEventSubscriptionOutput, error) {
	var output redshift.DeleteEventSubscriptionOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeleteEventSubscription, input).Get(ctx, &output)
	return &output, err
}

func (a *RedshiftStub) DeleteEventSubscriptionAsync(ctx workflow.Context, input *redshift.DeleteEventSubscriptionInput) *RedshiftDeleteEventSubscriptionResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeleteEventSubscription, input)
	return &RedshiftDeleteEventSubscriptionResult{Result: future}
}

func (a *RedshiftStub) DeleteHsmClientCertificate(ctx workflow.Context, input *redshift.DeleteHsmClientCertificateInput) (*redshift.DeleteHsmClientCertificateOutput, error) {
	var output redshift.DeleteHsmClientCertificateOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeleteHsmClientCertificate, input).Get(ctx, &output)
	return &output, err
}

func (a *RedshiftStub) DeleteHsmClientCertificateAsync(ctx workflow.Context, input *redshift.DeleteHsmClientCertificateInput) *RedshiftDeleteHsmClientCertificateResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeleteHsmClientCertificate, input)
	return &RedshiftDeleteHsmClientCertificateResult{Result: future}
}

func (a *RedshiftStub) DeleteHsmConfiguration(ctx workflow.Context, input *redshift.DeleteHsmConfigurationInput) (*redshift.DeleteHsmConfigurationOutput, error) {
	var output redshift.DeleteHsmConfigurationOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeleteHsmConfiguration, input).Get(ctx, &output)
	return &output, err
}

func (a *RedshiftStub) DeleteHsmConfigurationAsync(ctx workflow.Context, input *redshift.DeleteHsmConfigurationInput) *RedshiftDeleteHsmConfigurationResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeleteHsmConfiguration, input)
	return &RedshiftDeleteHsmConfigurationResult{Result: future}
}

func (a *RedshiftStub) DeleteScheduledAction(ctx workflow.Context, input *redshift.DeleteScheduledActionInput) (*redshift.DeleteScheduledActionOutput, error) {
	var output redshift.DeleteScheduledActionOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeleteScheduledAction, input).Get(ctx, &output)
	return &output, err
}

func (a *RedshiftStub) DeleteScheduledActionAsync(ctx workflow.Context, input *redshift.DeleteScheduledActionInput) *RedshiftDeleteScheduledActionResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeleteScheduledAction, input)
	return &RedshiftDeleteScheduledActionResult{Result: future}
}

func (a *RedshiftStub) DeleteSnapshotCopyGrant(ctx workflow.Context, input *redshift.DeleteSnapshotCopyGrantInput) (*redshift.DeleteSnapshotCopyGrantOutput, error) {
	var output redshift.DeleteSnapshotCopyGrantOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeleteSnapshotCopyGrant, input).Get(ctx, &output)
	return &output, err
}

func (a *RedshiftStub) DeleteSnapshotCopyGrantAsync(ctx workflow.Context, input *redshift.DeleteSnapshotCopyGrantInput) *RedshiftDeleteSnapshotCopyGrantResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeleteSnapshotCopyGrant, input)
	return &RedshiftDeleteSnapshotCopyGrantResult{Result: future}
}

func (a *RedshiftStub) DeleteSnapshotSchedule(ctx workflow.Context, input *redshift.DeleteSnapshotScheduleInput) (*redshift.DeleteSnapshotScheduleOutput, error) {
	var output redshift.DeleteSnapshotScheduleOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeleteSnapshotSchedule, input).Get(ctx, &output)
	return &output, err
}

func (a *RedshiftStub) DeleteSnapshotScheduleAsync(ctx workflow.Context, input *redshift.DeleteSnapshotScheduleInput) *RedshiftDeleteSnapshotScheduleResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeleteSnapshotSchedule, input)
	return &RedshiftDeleteSnapshotScheduleResult{Result: future}
}

func (a *RedshiftStub) DeleteTags(ctx workflow.Context, input *redshift.DeleteTagsInput) (*redshift.DeleteTagsOutput, error) {
	var output redshift.DeleteTagsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeleteTags, input).Get(ctx, &output)
	return &output, err
}

func (a *RedshiftStub) DeleteTagsAsync(ctx workflow.Context, input *redshift.DeleteTagsInput) *RedshiftDeleteTagsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeleteTags, input)
	return &RedshiftDeleteTagsResult{Result: future}
}

func (a *RedshiftStub) DeleteUsageLimit(ctx workflow.Context, input *redshift.DeleteUsageLimitInput) (*redshift.DeleteUsageLimitOutput, error) {
	var output redshift.DeleteUsageLimitOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeleteUsageLimit, input).Get(ctx, &output)
	return &output, err
}

func (a *RedshiftStub) DeleteUsageLimitAsync(ctx workflow.Context, input *redshift.DeleteUsageLimitInput) *RedshiftDeleteUsageLimitResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeleteUsageLimit, input)
	return &RedshiftDeleteUsageLimitResult{Result: future}
}

func (a *RedshiftStub) DescribeAccountAttributes(ctx workflow.Context, input *redshift.DescribeAccountAttributesInput) (*redshift.DescribeAccountAttributesOutput, error) {
	var output redshift.DescribeAccountAttributesOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DescribeAccountAttributes, input).Get(ctx, &output)
	return &output, err
}

func (a *RedshiftStub) DescribeAccountAttributesAsync(ctx workflow.Context, input *redshift.DescribeAccountAttributesInput) *RedshiftDescribeAccountAttributesResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DescribeAccountAttributes, input)
	return &RedshiftDescribeAccountAttributesResult{Result: future}
}

func (a *RedshiftStub) DescribeClusterDbRevisions(ctx workflow.Context, input *redshift.DescribeClusterDbRevisionsInput) (*redshift.DescribeClusterDbRevisionsOutput, error) {
	var output redshift.DescribeClusterDbRevisionsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DescribeClusterDbRevisions, input).Get(ctx, &output)
	return &output, err
}

func (a *RedshiftStub) DescribeClusterDbRevisionsAsync(ctx workflow.Context, input *redshift.DescribeClusterDbRevisionsInput) *RedshiftDescribeClusterDbRevisionsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DescribeClusterDbRevisions, input)
	return &RedshiftDescribeClusterDbRevisionsResult{Result: future}
}

func (a *RedshiftStub) DescribeClusterParameterGroups(ctx workflow.Context, input *redshift.DescribeClusterParameterGroupsInput) (*redshift.DescribeClusterParameterGroupsOutput, error) {
	var output redshift.DescribeClusterParameterGroupsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DescribeClusterParameterGroups, input).Get(ctx, &output)
	return &output, err
}

func (a *RedshiftStub) DescribeClusterParameterGroupsAsync(ctx workflow.Context, input *redshift.DescribeClusterParameterGroupsInput) *RedshiftDescribeClusterParameterGroupsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DescribeClusterParameterGroups, input)
	return &RedshiftDescribeClusterParameterGroupsResult{Result: future}
}

func (a *RedshiftStub) DescribeClusterParameters(ctx workflow.Context, input *redshift.DescribeClusterParametersInput) (*redshift.DescribeClusterParametersOutput, error) {
	var output redshift.DescribeClusterParametersOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DescribeClusterParameters, input).Get(ctx, &output)
	return &output, err
}

func (a *RedshiftStub) DescribeClusterParametersAsync(ctx workflow.Context, input *redshift.DescribeClusterParametersInput) *RedshiftDescribeClusterParametersResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DescribeClusterParameters, input)
	return &RedshiftDescribeClusterParametersResult{Result: future}
}

func (a *RedshiftStub) DescribeClusterSecurityGroups(ctx workflow.Context, input *redshift.DescribeClusterSecurityGroupsInput) (*redshift.DescribeClusterSecurityGroupsOutput, error) {
	var output redshift.DescribeClusterSecurityGroupsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DescribeClusterSecurityGroups, input).Get(ctx, &output)
	return &output, err
}

func (a *RedshiftStub) DescribeClusterSecurityGroupsAsync(ctx workflow.Context, input *redshift.DescribeClusterSecurityGroupsInput) *RedshiftDescribeClusterSecurityGroupsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DescribeClusterSecurityGroups, input)
	return &RedshiftDescribeClusterSecurityGroupsResult{Result: future}
}

func (a *RedshiftStub) DescribeClusterSnapshots(ctx workflow.Context, input *redshift.DescribeClusterSnapshotsInput) (*redshift.DescribeClusterSnapshotsOutput, error) {
	var output redshift.DescribeClusterSnapshotsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DescribeClusterSnapshots, input).Get(ctx, &output)
	return &output, err
}

func (a *RedshiftStub) DescribeClusterSnapshotsAsync(ctx workflow.Context, input *redshift.DescribeClusterSnapshotsInput) *RedshiftDescribeClusterSnapshotsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DescribeClusterSnapshots, input)
	return &RedshiftDescribeClusterSnapshotsResult{Result: future}
}

func (a *RedshiftStub) DescribeClusterSubnetGroups(ctx workflow.Context, input *redshift.DescribeClusterSubnetGroupsInput) (*redshift.DescribeClusterSubnetGroupsOutput, error) {
	var output redshift.DescribeClusterSubnetGroupsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DescribeClusterSubnetGroups, input).Get(ctx, &output)
	return &output, err
}

func (a *RedshiftStub) DescribeClusterSubnetGroupsAsync(ctx workflow.Context, input *redshift.DescribeClusterSubnetGroupsInput) *RedshiftDescribeClusterSubnetGroupsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DescribeClusterSubnetGroups, input)
	return &RedshiftDescribeClusterSubnetGroupsResult{Result: future}
}

func (a *RedshiftStub) DescribeClusterTracks(ctx workflow.Context, input *redshift.DescribeClusterTracksInput) (*redshift.DescribeClusterTracksOutput, error) {
	var output redshift.DescribeClusterTracksOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DescribeClusterTracks, input).Get(ctx, &output)
	return &output, err
}

func (a *RedshiftStub) DescribeClusterTracksAsync(ctx workflow.Context, input *redshift.DescribeClusterTracksInput) *RedshiftDescribeClusterTracksResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DescribeClusterTracks, input)
	return &RedshiftDescribeClusterTracksResult{Result: future}
}

func (a *RedshiftStub) DescribeClusterVersions(ctx workflow.Context, input *redshift.DescribeClusterVersionsInput) (*redshift.DescribeClusterVersionsOutput, error) {
	var output redshift.DescribeClusterVersionsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DescribeClusterVersions, input).Get(ctx, &output)
	return &output, err
}

func (a *RedshiftStub) DescribeClusterVersionsAsync(ctx workflow.Context, input *redshift.DescribeClusterVersionsInput) *RedshiftDescribeClusterVersionsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DescribeClusterVersions, input)
	return &RedshiftDescribeClusterVersionsResult{Result: future}
}

func (a *RedshiftStub) DescribeClusters(ctx workflow.Context, input *redshift.DescribeClustersInput) (*redshift.DescribeClustersOutput, error) {
	var output redshift.DescribeClustersOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DescribeClusters, input).Get(ctx, &output)
	return &output, err
}

func (a *RedshiftStub) DescribeClustersAsync(ctx workflow.Context, input *redshift.DescribeClustersInput) *RedshiftDescribeClustersResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DescribeClusters, input)
	return &RedshiftDescribeClustersResult{Result: future}
}

func (a *RedshiftStub) DescribeDefaultClusterParameters(ctx workflow.Context, input *redshift.DescribeDefaultClusterParametersInput) (*redshift.DescribeDefaultClusterParametersOutput, error) {
	var output redshift.DescribeDefaultClusterParametersOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DescribeDefaultClusterParameters, input).Get(ctx, &output)
	return &output, err
}

func (a *RedshiftStub) DescribeDefaultClusterParametersAsync(ctx workflow.Context, input *redshift.DescribeDefaultClusterParametersInput) *RedshiftDescribeDefaultClusterParametersResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DescribeDefaultClusterParameters, input)
	return &RedshiftDescribeDefaultClusterParametersResult{Result: future}
}

func (a *RedshiftStub) DescribeEventCategories(ctx workflow.Context, input *redshift.DescribeEventCategoriesInput) (*redshift.DescribeEventCategoriesOutput, error) {
	var output redshift.DescribeEventCategoriesOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DescribeEventCategories, input).Get(ctx, &output)
	return &output, err
}

func (a *RedshiftStub) DescribeEventCategoriesAsync(ctx workflow.Context, input *redshift.DescribeEventCategoriesInput) *RedshiftDescribeEventCategoriesResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DescribeEventCategories, input)
	return &RedshiftDescribeEventCategoriesResult{Result: future}
}

func (a *RedshiftStub) DescribeEventSubscriptions(ctx workflow.Context, input *redshift.DescribeEventSubscriptionsInput) (*redshift.DescribeEventSubscriptionsOutput, error) {
	var output redshift.DescribeEventSubscriptionsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DescribeEventSubscriptions, input).Get(ctx, &output)
	return &output, err
}

func (a *RedshiftStub) DescribeEventSubscriptionsAsync(ctx workflow.Context, input *redshift.DescribeEventSubscriptionsInput) *RedshiftDescribeEventSubscriptionsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DescribeEventSubscriptions, input)
	return &RedshiftDescribeEventSubscriptionsResult{Result: future}
}

func (a *RedshiftStub) DescribeEvents(ctx workflow.Context, input *redshift.DescribeEventsInput) (*redshift.DescribeEventsOutput, error) {
	var output redshift.DescribeEventsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DescribeEvents, input).Get(ctx, &output)
	return &output, err
}

func (a *RedshiftStub) DescribeEventsAsync(ctx workflow.Context, input *redshift.DescribeEventsInput) *RedshiftDescribeEventsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DescribeEvents, input)
	return &RedshiftDescribeEventsResult{Result: future}
}

func (a *RedshiftStub) DescribeHsmClientCertificates(ctx workflow.Context, input *redshift.DescribeHsmClientCertificatesInput) (*redshift.DescribeHsmClientCertificatesOutput, error) {
	var output redshift.DescribeHsmClientCertificatesOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DescribeHsmClientCertificates, input).Get(ctx, &output)
	return &output, err
}

func (a *RedshiftStub) DescribeHsmClientCertificatesAsync(ctx workflow.Context, input *redshift.DescribeHsmClientCertificatesInput) *RedshiftDescribeHsmClientCertificatesResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DescribeHsmClientCertificates, input)
	return &RedshiftDescribeHsmClientCertificatesResult{Result: future}
}

func (a *RedshiftStub) DescribeHsmConfigurations(ctx workflow.Context, input *redshift.DescribeHsmConfigurationsInput) (*redshift.DescribeHsmConfigurationsOutput, error) {
	var output redshift.DescribeHsmConfigurationsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DescribeHsmConfigurations, input).Get(ctx, &output)
	return &output, err
}

func (a *RedshiftStub) DescribeHsmConfigurationsAsync(ctx workflow.Context, input *redshift.DescribeHsmConfigurationsInput) *RedshiftDescribeHsmConfigurationsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DescribeHsmConfigurations, input)
	return &RedshiftDescribeHsmConfigurationsResult{Result: future}
}

func (a *RedshiftStub) DescribeLoggingStatus(ctx workflow.Context, input *redshift.DescribeLoggingStatusInput) (*redshift.LoggingStatus, error) {
	var output redshift.LoggingStatus
	err := workflow.ExecuteActivity(ctx, a.activities.DescribeLoggingStatus, input).Get(ctx, &output)
	return &output, err
}

func (a *RedshiftStub) DescribeLoggingStatusAsync(ctx workflow.Context, input *redshift.DescribeLoggingStatusInput) *RedshiftDescribeLoggingStatusResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DescribeLoggingStatus, input)
	return &RedshiftDescribeLoggingStatusResult{Result: future}
}

func (a *RedshiftStub) DescribeNodeConfigurationOptions(ctx workflow.Context, input *redshift.DescribeNodeConfigurationOptionsInput) (*redshift.DescribeNodeConfigurationOptionsOutput, error) {
	var output redshift.DescribeNodeConfigurationOptionsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DescribeNodeConfigurationOptions, input).Get(ctx, &output)
	return &output, err
}

func (a *RedshiftStub) DescribeNodeConfigurationOptionsAsync(ctx workflow.Context, input *redshift.DescribeNodeConfigurationOptionsInput) *RedshiftDescribeNodeConfigurationOptionsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DescribeNodeConfigurationOptions, input)
	return &RedshiftDescribeNodeConfigurationOptionsResult{Result: future}
}

func (a *RedshiftStub) DescribeOrderableClusterOptions(ctx workflow.Context, input *redshift.DescribeOrderableClusterOptionsInput) (*redshift.DescribeOrderableClusterOptionsOutput, error) {
	var output redshift.DescribeOrderableClusterOptionsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DescribeOrderableClusterOptions, input).Get(ctx, &output)
	return &output, err
}

func (a *RedshiftStub) DescribeOrderableClusterOptionsAsync(ctx workflow.Context, input *redshift.DescribeOrderableClusterOptionsInput) *RedshiftDescribeOrderableClusterOptionsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DescribeOrderableClusterOptions, input)
	return &RedshiftDescribeOrderableClusterOptionsResult{Result: future}
}

func (a *RedshiftStub) DescribeReservedNodeOfferings(ctx workflow.Context, input *redshift.DescribeReservedNodeOfferingsInput) (*redshift.DescribeReservedNodeOfferingsOutput, error) {
	var output redshift.DescribeReservedNodeOfferingsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DescribeReservedNodeOfferings, input).Get(ctx, &output)
	return &output, err
}

func (a *RedshiftStub) DescribeReservedNodeOfferingsAsync(ctx workflow.Context, input *redshift.DescribeReservedNodeOfferingsInput) *RedshiftDescribeReservedNodeOfferingsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DescribeReservedNodeOfferings, input)
	return &RedshiftDescribeReservedNodeOfferingsResult{Result: future}
}

func (a *RedshiftStub) DescribeReservedNodes(ctx workflow.Context, input *redshift.DescribeReservedNodesInput) (*redshift.DescribeReservedNodesOutput, error) {
	var output redshift.DescribeReservedNodesOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DescribeReservedNodes, input).Get(ctx, &output)
	return &output, err
}

func (a *RedshiftStub) DescribeReservedNodesAsync(ctx workflow.Context, input *redshift.DescribeReservedNodesInput) *RedshiftDescribeReservedNodesResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DescribeReservedNodes, input)
	return &RedshiftDescribeReservedNodesResult{Result: future}
}

func (a *RedshiftStub) DescribeResize(ctx workflow.Context, input *redshift.DescribeResizeInput) (*redshift.DescribeResizeOutput, error) {
	var output redshift.DescribeResizeOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DescribeResize, input).Get(ctx, &output)
	return &output, err
}

func (a *RedshiftStub) DescribeResizeAsync(ctx workflow.Context, input *redshift.DescribeResizeInput) *RedshiftDescribeResizeResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DescribeResize, input)
	return &RedshiftDescribeResizeResult{Result: future}
}

func (a *RedshiftStub) DescribeScheduledActions(ctx workflow.Context, input *redshift.DescribeScheduledActionsInput) (*redshift.DescribeScheduledActionsOutput, error) {
	var output redshift.DescribeScheduledActionsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DescribeScheduledActions, input).Get(ctx, &output)
	return &output, err
}

func (a *RedshiftStub) DescribeScheduledActionsAsync(ctx workflow.Context, input *redshift.DescribeScheduledActionsInput) *RedshiftDescribeScheduledActionsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DescribeScheduledActions, input)
	return &RedshiftDescribeScheduledActionsResult{Result: future}
}

func (a *RedshiftStub) DescribeSnapshotCopyGrants(ctx workflow.Context, input *redshift.DescribeSnapshotCopyGrantsInput) (*redshift.DescribeSnapshotCopyGrantsOutput, error) {
	var output redshift.DescribeSnapshotCopyGrantsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DescribeSnapshotCopyGrants, input).Get(ctx, &output)
	return &output, err
}

func (a *RedshiftStub) DescribeSnapshotCopyGrantsAsync(ctx workflow.Context, input *redshift.DescribeSnapshotCopyGrantsInput) *RedshiftDescribeSnapshotCopyGrantsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DescribeSnapshotCopyGrants, input)
	return &RedshiftDescribeSnapshotCopyGrantsResult{Result: future}
}

func (a *RedshiftStub) DescribeSnapshotSchedules(ctx workflow.Context, input *redshift.DescribeSnapshotSchedulesInput) (*redshift.DescribeSnapshotSchedulesOutput, error) {
	var output redshift.DescribeSnapshotSchedulesOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DescribeSnapshotSchedules, input).Get(ctx, &output)
	return &output, err
}

func (a *RedshiftStub) DescribeSnapshotSchedulesAsync(ctx workflow.Context, input *redshift.DescribeSnapshotSchedulesInput) *RedshiftDescribeSnapshotSchedulesResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DescribeSnapshotSchedules, input)
	return &RedshiftDescribeSnapshotSchedulesResult{Result: future}
}

func (a *RedshiftStub) DescribeStorage(ctx workflow.Context, input *redshift.DescribeStorageInput) (*redshift.DescribeStorageOutput, error) {
	var output redshift.DescribeStorageOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DescribeStorage, input).Get(ctx, &output)
	return &output, err
}

func (a *RedshiftStub) DescribeStorageAsync(ctx workflow.Context, input *redshift.DescribeStorageInput) *RedshiftDescribeStorageResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DescribeStorage, input)
	return &RedshiftDescribeStorageResult{Result: future}
}

func (a *RedshiftStub) DescribeTableRestoreStatus(ctx workflow.Context, input *redshift.DescribeTableRestoreStatusInput) (*redshift.DescribeTableRestoreStatusOutput, error) {
	var output redshift.DescribeTableRestoreStatusOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DescribeTableRestoreStatus, input).Get(ctx, &output)
	return &output, err
}

func (a *RedshiftStub) DescribeTableRestoreStatusAsync(ctx workflow.Context, input *redshift.DescribeTableRestoreStatusInput) *RedshiftDescribeTableRestoreStatusResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DescribeTableRestoreStatus, input)
	return &RedshiftDescribeTableRestoreStatusResult{Result: future}
}

func (a *RedshiftStub) DescribeTags(ctx workflow.Context, input *redshift.DescribeTagsInput) (*redshift.DescribeTagsOutput, error) {
	var output redshift.DescribeTagsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DescribeTags, input).Get(ctx, &output)
	return &output, err
}

func (a *RedshiftStub) DescribeTagsAsync(ctx workflow.Context, input *redshift.DescribeTagsInput) *RedshiftDescribeTagsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DescribeTags, input)
	return &RedshiftDescribeTagsResult{Result: future}
}

func (a *RedshiftStub) DescribeUsageLimits(ctx workflow.Context, input *redshift.DescribeUsageLimitsInput) (*redshift.DescribeUsageLimitsOutput, error) {
	var output redshift.DescribeUsageLimitsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DescribeUsageLimits, input).Get(ctx, &output)
	return &output, err
}

func (a *RedshiftStub) DescribeUsageLimitsAsync(ctx workflow.Context, input *redshift.DescribeUsageLimitsInput) *RedshiftDescribeUsageLimitsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DescribeUsageLimits, input)
	return &RedshiftDescribeUsageLimitsResult{Result: future}
}

func (a *RedshiftStub) DisableLogging(ctx workflow.Context, input *redshift.DisableLoggingInput) (*redshift.LoggingStatus, error) {
	var output redshift.LoggingStatus
	err := workflow.ExecuteActivity(ctx, a.activities.DisableLogging, input).Get(ctx, &output)
	return &output, err
}

func (a *RedshiftStub) DisableLoggingAsync(ctx workflow.Context, input *redshift.DisableLoggingInput) *RedshiftDisableLoggingResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DisableLogging, input)
	return &RedshiftDisableLoggingResult{Result: future}
}

func (a *RedshiftStub) DisableSnapshotCopy(ctx workflow.Context, input *redshift.DisableSnapshotCopyInput) (*redshift.DisableSnapshotCopyOutput, error) {
	var output redshift.DisableSnapshotCopyOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DisableSnapshotCopy, input).Get(ctx, &output)
	return &output, err
}

func (a *RedshiftStub) DisableSnapshotCopyAsync(ctx workflow.Context, input *redshift.DisableSnapshotCopyInput) *RedshiftDisableSnapshotCopyResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DisableSnapshotCopy, input)
	return &RedshiftDisableSnapshotCopyResult{Result: future}
}

func (a *RedshiftStub) EnableLogging(ctx workflow.Context, input *redshift.EnableLoggingInput) (*redshift.LoggingStatus, error) {
	var output redshift.LoggingStatus
	err := workflow.ExecuteActivity(ctx, a.activities.EnableLogging, input).Get(ctx, &output)
	return &output, err
}

func (a *RedshiftStub) EnableLoggingAsync(ctx workflow.Context, input *redshift.EnableLoggingInput) *RedshiftEnableLoggingResult {
	future := workflow.ExecuteActivity(ctx, a.activities.EnableLogging, input)
	return &RedshiftEnableLoggingResult{Result: future}
}

func (a *RedshiftStub) EnableSnapshotCopy(ctx workflow.Context, input *redshift.EnableSnapshotCopyInput) (*redshift.EnableSnapshotCopyOutput, error) {
	var output redshift.EnableSnapshotCopyOutput
	err := workflow.ExecuteActivity(ctx, a.activities.EnableSnapshotCopy, input).Get(ctx, &output)
	return &output, err
}

func (a *RedshiftStub) EnableSnapshotCopyAsync(ctx workflow.Context, input *redshift.EnableSnapshotCopyInput) *RedshiftEnableSnapshotCopyResult {
	future := workflow.ExecuteActivity(ctx, a.activities.EnableSnapshotCopy, input)
	return &RedshiftEnableSnapshotCopyResult{Result: future}
}

func (a *RedshiftStub) GetClusterCredentials(ctx workflow.Context, input *redshift.GetClusterCredentialsInput) (*redshift.GetClusterCredentialsOutput, error) {
	var output redshift.GetClusterCredentialsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetClusterCredentials, input).Get(ctx, &output)
	return &output, err
}

func (a *RedshiftStub) GetClusterCredentialsAsync(ctx workflow.Context, input *redshift.GetClusterCredentialsInput) *RedshiftGetClusterCredentialsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetClusterCredentials, input)
	return &RedshiftGetClusterCredentialsResult{Result: future}
}

func (a *RedshiftStub) GetReservedNodeExchangeOfferings(ctx workflow.Context, input *redshift.GetReservedNodeExchangeOfferingsInput) (*redshift.GetReservedNodeExchangeOfferingsOutput, error) {
	var output redshift.GetReservedNodeExchangeOfferingsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetReservedNodeExchangeOfferings, input).Get(ctx, &output)
	return &output, err
}

func (a *RedshiftStub) GetReservedNodeExchangeOfferingsAsync(ctx workflow.Context, input *redshift.GetReservedNodeExchangeOfferingsInput) *RedshiftGetReservedNodeExchangeOfferingsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetReservedNodeExchangeOfferings, input)
	return &RedshiftGetReservedNodeExchangeOfferingsResult{Result: future}
}

func (a *RedshiftStub) ModifyCluster(ctx workflow.Context, input *redshift.ModifyClusterInput) (*redshift.ModifyClusterOutput, error) {
	var output redshift.ModifyClusterOutput
	err := workflow.ExecuteActivity(ctx, a.activities.ModifyCluster, input).Get(ctx, &output)
	return &output, err
}

func (a *RedshiftStub) ModifyClusterAsync(ctx workflow.Context, input *redshift.ModifyClusterInput) *RedshiftModifyClusterResult {
	future := workflow.ExecuteActivity(ctx, a.activities.ModifyCluster, input)
	return &RedshiftModifyClusterResult{Result: future}
}

func (a *RedshiftStub) ModifyClusterDbRevision(ctx workflow.Context, input *redshift.ModifyClusterDbRevisionInput) (*redshift.ModifyClusterDbRevisionOutput, error) {
	var output redshift.ModifyClusterDbRevisionOutput
	err := workflow.ExecuteActivity(ctx, a.activities.ModifyClusterDbRevision, input).Get(ctx, &output)
	return &output, err
}

func (a *RedshiftStub) ModifyClusterDbRevisionAsync(ctx workflow.Context, input *redshift.ModifyClusterDbRevisionInput) *RedshiftModifyClusterDbRevisionResult {
	future := workflow.ExecuteActivity(ctx, a.activities.ModifyClusterDbRevision, input)
	return &RedshiftModifyClusterDbRevisionResult{Result: future}
}

func (a *RedshiftStub) ModifyClusterIamRoles(ctx workflow.Context, input *redshift.ModifyClusterIamRolesInput) (*redshift.ModifyClusterIamRolesOutput, error) {
	var output redshift.ModifyClusterIamRolesOutput
	err := workflow.ExecuteActivity(ctx, a.activities.ModifyClusterIamRoles, input).Get(ctx, &output)
	return &output, err
}

func (a *RedshiftStub) ModifyClusterIamRolesAsync(ctx workflow.Context, input *redshift.ModifyClusterIamRolesInput) *RedshiftModifyClusterIamRolesResult {
	future := workflow.ExecuteActivity(ctx, a.activities.ModifyClusterIamRoles, input)
	return &RedshiftModifyClusterIamRolesResult{Result: future}
}

func (a *RedshiftStub) ModifyClusterMaintenance(ctx workflow.Context, input *redshift.ModifyClusterMaintenanceInput) (*redshift.ModifyClusterMaintenanceOutput, error) {
	var output redshift.ModifyClusterMaintenanceOutput
	err := workflow.ExecuteActivity(ctx, a.activities.ModifyClusterMaintenance, input).Get(ctx, &output)
	return &output, err
}

func (a *RedshiftStub) ModifyClusterMaintenanceAsync(ctx workflow.Context, input *redshift.ModifyClusterMaintenanceInput) *RedshiftModifyClusterMaintenanceResult {
	future := workflow.ExecuteActivity(ctx, a.activities.ModifyClusterMaintenance, input)
	return &RedshiftModifyClusterMaintenanceResult{Result: future}
}

func (a *RedshiftStub) ModifyClusterParameterGroup(ctx workflow.Context, input *redshift.ModifyClusterParameterGroupInput) (*redshift.ClusterParameterGroupNameMessage, error) {
	var output redshift.ClusterParameterGroupNameMessage
	err := workflow.ExecuteActivity(ctx, a.activities.ModifyClusterParameterGroup, input).Get(ctx, &output)
	return &output, err
}

func (a *RedshiftStub) ModifyClusterParameterGroupAsync(ctx workflow.Context, input *redshift.ModifyClusterParameterGroupInput) *RedshiftModifyClusterParameterGroupResult {
	future := workflow.ExecuteActivity(ctx, a.activities.ModifyClusterParameterGroup, input)
	return &RedshiftModifyClusterParameterGroupResult{Result: future}
}

func (a *RedshiftStub) ModifyClusterSnapshot(ctx workflow.Context, input *redshift.ModifyClusterSnapshotInput) (*redshift.ModifyClusterSnapshotOutput, error) {
	var output redshift.ModifyClusterSnapshotOutput
	err := workflow.ExecuteActivity(ctx, a.activities.ModifyClusterSnapshot, input).Get(ctx, &output)
	return &output, err
}

func (a *RedshiftStub) ModifyClusterSnapshotAsync(ctx workflow.Context, input *redshift.ModifyClusterSnapshotInput) *RedshiftModifyClusterSnapshotResult {
	future := workflow.ExecuteActivity(ctx, a.activities.ModifyClusterSnapshot, input)
	return &RedshiftModifyClusterSnapshotResult{Result: future}
}

func (a *RedshiftStub) ModifyClusterSnapshotSchedule(ctx workflow.Context, input *redshift.ModifyClusterSnapshotScheduleInput) (*redshift.ModifyClusterSnapshotScheduleOutput, error) {
	var output redshift.ModifyClusterSnapshotScheduleOutput
	err := workflow.ExecuteActivity(ctx, a.activities.ModifyClusterSnapshotSchedule, input).Get(ctx, &output)
	return &output, err
}

func (a *RedshiftStub) ModifyClusterSnapshotScheduleAsync(ctx workflow.Context, input *redshift.ModifyClusterSnapshotScheduleInput) *RedshiftModifyClusterSnapshotScheduleResult {
	future := workflow.ExecuteActivity(ctx, a.activities.ModifyClusterSnapshotSchedule, input)
	return &RedshiftModifyClusterSnapshotScheduleResult{Result: future}
}

func (a *RedshiftStub) ModifyClusterSubnetGroup(ctx workflow.Context, input *redshift.ModifyClusterSubnetGroupInput) (*redshift.ModifyClusterSubnetGroupOutput, error) {
	var output redshift.ModifyClusterSubnetGroupOutput
	err := workflow.ExecuteActivity(ctx, a.activities.ModifyClusterSubnetGroup, input).Get(ctx, &output)
	return &output, err
}

func (a *RedshiftStub) ModifyClusterSubnetGroupAsync(ctx workflow.Context, input *redshift.ModifyClusterSubnetGroupInput) *RedshiftModifyClusterSubnetGroupResult {
	future := workflow.ExecuteActivity(ctx, a.activities.ModifyClusterSubnetGroup, input)
	return &RedshiftModifyClusterSubnetGroupResult{Result: future}
}

func (a *RedshiftStub) ModifyEventSubscription(ctx workflow.Context, input *redshift.ModifyEventSubscriptionInput) (*redshift.ModifyEventSubscriptionOutput, error) {
	var output redshift.ModifyEventSubscriptionOutput
	err := workflow.ExecuteActivity(ctx, a.activities.ModifyEventSubscription, input).Get(ctx, &output)
	return &output, err
}

func (a *RedshiftStub) ModifyEventSubscriptionAsync(ctx workflow.Context, input *redshift.ModifyEventSubscriptionInput) *RedshiftModifyEventSubscriptionResult {
	future := workflow.ExecuteActivity(ctx, a.activities.ModifyEventSubscription, input)
	return &RedshiftModifyEventSubscriptionResult{Result: future}
}

func (a *RedshiftStub) ModifyScheduledAction(ctx workflow.Context, input *redshift.ModifyScheduledActionInput) (*redshift.ModifyScheduledActionOutput, error) {
	var output redshift.ModifyScheduledActionOutput
	err := workflow.ExecuteActivity(ctx, a.activities.ModifyScheduledAction, input).Get(ctx, &output)
	return &output, err
}

func (a *RedshiftStub) ModifyScheduledActionAsync(ctx workflow.Context, input *redshift.ModifyScheduledActionInput) *RedshiftModifyScheduledActionResult {
	future := workflow.ExecuteActivity(ctx, a.activities.ModifyScheduledAction, input)
	return &RedshiftModifyScheduledActionResult{Result: future}
}

func (a *RedshiftStub) ModifySnapshotCopyRetentionPeriod(ctx workflow.Context, input *redshift.ModifySnapshotCopyRetentionPeriodInput) (*redshift.ModifySnapshotCopyRetentionPeriodOutput, error) {
	var output redshift.ModifySnapshotCopyRetentionPeriodOutput
	err := workflow.ExecuteActivity(ctx, a.activities.ModifySnapshotCopyRetentionPeriod, input).Get(ctx, &output)
	return &output, err
}

func (a *RedshiftStub) ModifySnapshotCopyRetentionPeriodAsync(ctx workflow.Context, input *redshift.ModifySnapshotCopyRetentionPeriodInput) *RedshiftModifySnapshotCopyRetentionPeriodResult {
	future := workflow.ExecuteActivity(ctx, a.activities.ModifySnapshotCopyRetentionPeriod, input)
	return &RedshiftModifySnapshotCopyRetentionPeriodResult{Result: future}
}

func (a *RedshiftStub) ModifySnapshotSchedule(ctx workflow.Context, input *redshift.ModifySnapshotScheduleInput) (*redshift.ModifySnapshotScheduleOutput, error) {
	var output redshift.ModifySnapshotScheduleOutput
	err := workflow.ExecuteActivity(ctx, a.activities.ModifySnapshotSchedule, input).Get(ctx, &output)
	return &output, err
}

func (a *RedshiftStub) ModifySnapshotScheduleAsync(ctx workflow.Context, input *redshift.ModifySnapshotScheduleInput) *RedshiftModifySnapshotScheduleResult {
	future := workflow.ExecuteActivity(ctx, a.activities.ModifySnapshotSchedule, input)
	return &RedshiftModifySnapshotScheduleResult{Result: future}
}

func (a *RedshiftStub) ModifyUsageLimit(ctx workflow.Context, input *redshift.ModifyUsageLimitInput) (*redshift.ModifyUsageLimitOutput, error) {
	var output redshift.ModifyUsageLimitOutput
	err := workflow.ExecuteActivity(ctx, a.activities.ModifyUsageLimit, input).Get(ctx, &output)
	return &output, err
}

func (a *RedshiftStub) ModifyUsageLimitAsync(ctx workflow.Context, input *redshift.ModifyUsageLimitInput) *RedshiftModifyUsageLimitResult {
	future := workflow.ExecuteActivity(ctx, a.activities.ModifyUsageLimit, input)
	return &RedshiftModifyUsageLimitResult{Result: future}
}

func (a *RedshiftStub) PauseCluster(ctx workflow.Context, input *redshift.PauseClusterInput) (*redshift.PauseClusterOutput, error) {
	var output redshift.PauseClusterOutput
	err := workflow.ExecuteActivity(ctx, a.activities.PauseCluster, input).Get(ctx, &output)
	return &output, err
}

func (a *RedshiftStub) PauseClusterAsync(ctx workflow.Context, input *redshift.PauseClusterInput) *RedshiftPauseClusterResult {
	future := workflow.ExecuteActivity(ctx, a.activities.PauseCluster, input)
	return &RedshiftPauseClusterResult{Result: future}
}

func (a *RedshiftStub) PurchaseReservedNodeOffering(ctx workflow.Context, input *redshift.PurchaseReservedNodeOfferingInput) (*redshift.PurchaseReservedNodeOfferingOutput, error) {
	var output redshift.PurchaseReservedNodeOfferingOutput
	err := workflow.ExecuteActivity(ctx, a.activities.PurchaseReservedNodeOffering, input).Get(ctx, &output)
	return &output, err
}

func (a *RedshiftStub) PurchaseReservedNodeOfferingAsync(ctx workflow.Context, input *redshift.PurchaseReservedNodeOfferingInput) *RedshiftPurchaseReservedNodeOfferingResult {
	future := workflow.ExecuteActivity(ctx, a.activities.PurchaseReservedNodeOffering, input)
	return &RedshiftPurchaseReservedNodeOfferingResult{Result: future}
}

func (a *RedshiftStub) RebootCluster(ctx workflow.Context, input *redshift.RebootClusterInput) (*redshift.RebootClusterOutput, error) {
	var output redshift.RebootClusterOutput
	err := workflow.ExecuteActivity(ctx, a.activities.RebootCluster, input).Get(ctx, &output)
	return &output, err
}

func (a *RedshiftStub) RebootClusterAsync(ctx workflow.Context, input *redshift.RebootClusterInput) *RedshiftRebootClusterResult {
	future := workflow.ExecuteActivity(ctx, a.activities.RebootCluster, input)
	return &RedshiftRebootClusterResult{Result: future}
}

func (a *RedshiftStub) ResetClusterParameterGroup(ctx workflow.Context, input *redshift.ResetClusterParameterGroupInput) (*redshift.ClusterParameterGroupNameMessage, error) {
	var output redshift.ClusterParameterGroupNameMessage
	err := workflow.ExecuteActivity(ctx, a.activities.ResetClusterParameterGroup, input).Get(ctx, &output)
	return &output, err
}

func (a *RedshiftStub) ResetClusterParameterGroupAsync(ctx workflow.Context, input *redshift.ResetClusterParameterGroupInput) *RedshiftResetClusterParameterGroupResult {
	future := workflow.ExecuteActivity(ctx, a.activities.ResetClusterParameterGroup, input)
	return &RedshiftResetClusterParameterGroupResult{Result: future}
}

func (a *RedshiftStub) ResizeCluster(ctx workflow.Context, input *redshift.ResizeClusterInput) (*redshift.ResizeClusterOutput, error) {
	var output redshift.ResizeClusterOutput
	err := workflow.ExecuteActivity(ctx, a.activities.ResizeCluster, input).Get(ctx, &output)
	return &output, err
}

func (a *RedshiftStub) ResizeClusterAsync(ctx workflow.Context, input *redshift.ResizeClusterInput) *RedshiftResizeClusterResult {
	future := workflow.ExecuteActivity(ctx, a.activities.ResizeCluster, input)
	return &RedshiftResizeClusterResult{Result: future}
}

func (a *RedshiftStub) RestoreFromClusterSnapshot(ctx workflow.Context, input *redshift.RestoreFromClusterSnapshotInput) (*redshift.RestoreFromClusterSnapshotOutput, error) {
	var output redshift.RestoreFromClusterSnapshotOutput
	err := workflow.ExecuteActivity(ctx, a.activities.RestoreFromClusterSnapshot, input).Get(ctx, &output)
	return &output, err
}

func (a *RedshiftStub) RestoreFromClusterSnapshotAsync(ctx workflow.Context, input *redshift.RestoreFromClusterSnapshotInput) *RedshiftRestoreFromClusterSnapshotResult {
	future := workflow.ExecuteActivity(ctx, a.activities.RestoreFromClusterSnapshot, input)
	return &RedshiftRestoreFromClusterSnapshotResult{Result: future}
}

func (a *RedshiftStub) RestoreTableFromClusterSnapshot(ctx workflow.Context, input *redshift.RestoreTableFromClusterSnapshotInput) (*redshift.RestoreTableFromClusterSnapshotOutput, error) {
	var output redshift.RestoreTableFromClusterSnapshotOutput
	err := workflow.ExecuteActivity(ctx, a.activities.RestoreTableFromClusterSnapshot, input).Get(ctx, &output)
	return &output, err
}

func (a *RedshiftStub) RestoreTableFromClusterSnapshotAsync(ctx workflow.Context, input *redshift.RestoreTableFromClusterSnapshotInput) *RedshiftRestoreTableFromClusterSnapshotResult {
	future := workflow.ExecuteActivity(ctx, a.activities.RestoreTableFromClusterSnapshot, input)
	return &RedshiftRestoreTableFromClusterSnapshotResult{Result: future}
}

func (a *RedshiftStub) ResumeCluster(ctx workflow.Context, input *redshift.ResumeClusterInput) (*redshift.ResumeClusterOutput, error) {
	var output redshift.ResumeClusterOutput
	err := workflow.ExecuteActivity(ctx, a.activities.ResumeCluster, input).Get(ctx, &output)
	return &output, err
}

func (a *RedshiftStub) ResumeClusterAsync(ctx workflow.Context, input *redshift.ResumeClusterInput) *RedshiftResumeClusterResult {
	future := workflow.ExecuteActivity(ctx, a.activities.ResumeCluster, input)
	return &RedshiftResumeClusterResult{Result: future}
}

func (a *RedshiftStub) RevokeClusterSecurityGroupIngress(ctx workflow.Context, input *redshift.RevokeClusterSecurityGroupIngressInput) (*redshift.RevokeClusterSecurityGroupIngressOutput, error) {
	var output redshift.RevokeClusterSecurityGroupIngressOutput
	err := workflow.ExecuteActivity(ctx, a.activities.RevokeClusterSecurityGroupIngress, input).Get(ctx, &output)
	return &output, err
}

func (a *RedshiftStub) RevokeClusterSecurityGroupIngressAsync(ctx workflow.Context, input *redshift.RevokeClusterSecurityGroupIngressInput) *RedshiftRevokeClusterSecurityGroupIngressResult {
	future := workflow.ExecuteActivity(ctx, a.activities.RevokeClusterSecurityGroupIngress, input)
	return &RedshiftRevokeClusterSecurityGroupIngressResult{Result: future}
}

func (a *RedshiftStub) RevokeSnapshotAccess(ctx workflow.Context, input *redshift.RevokeSnapshotAccessInput) (*redshift.RevokeSnapshotAccessOutput, error) {
	var output redshift.RevokeSnapshotAccessOutput
	err := workflow.ExecuteActivity(ctx, a.activities.RevokeSnapshotAccess, input).Get(ctx, &output)
	return &output, err
}

func (a *RedshiftStub) RevokeSnapshotAccessAsync(ctx workflow.Context, input *redshift.RevokeSnapshotAccessInput) *RedshiftRevokeSnapshotAccessResult {
	future := workflow.ExecuteActivity(ctx, a.activities.RevokeSnapshotAccess, input)
	return &RedshiftRevokeSnapshotAccessResult{Result: future}
}

func (a *RedshiftStub) RotateEncryptionKey(ctx workflow.Context, input *redshift.RotateEncryptionKeyInput) (*redshift.RotateEncryptionKeyOutput, error) {
	var output redshift.RotateEncryptionKeyOutput
	err := workflow.ExecuteActivity(ctx, a.activities.RotateEncryptionKey, input).Get(ctx, &output)
	return &output, err
}

func (a *RedshiftStub) RotateEncryptionKeyAsync(ctx workflow.Context, input *redshift.RotateEncryptionKeyInput) *RedshiftRotateEncryptionKeyResult {
	future := workflow.ExecuteActivity(ctx, a.activities.RotateEncryptionKey, input)
	return &RedshiftRotateEncryptionKeyResult{Result: future}
}

func (a *RedshiftStub) WaitUntilClusterAvailable(ctx workflow.Context, input *redshift.DescribeClustersInput) error {
	return workflow.ExecuteActivity(ctx, a.activities.WaitUntilClusterAvailable, input).Get(ctx, nil)
}

func (a *RedshiftStub) WaitUntilClusterAvailableAsync(ctx workflow.Context, input *redshift.DescribeClustersInput) workflow.Future {
	return workflow.ExecuteActivity(ctx, a.activities.WaitUntilClusterAvailable, input)
}

func (a *RedshiftStub) WaitUntilClusterDeleted(ctx workflow.Context, input *redshift.DescribeClustersInput) error {
	return workflow.ExecuteActivity(ctx, a.activities.WaitUntilClusterDeleted, input).Get(ctx, nil)
}

func (a *RedshiftStub) WaitUntilClusterDeletedAsync(ctx workflow.Context, input *redshift.DescribeClustersInput) workflow.Future {
	return workflow.ExecuteActivity(ctx, a.activities.WaitUntilClusterDeleted, input)
}

func (a *RedshiftStub) WaitUntilClusterRestored(ctx workflow.Context, input *redshift.DescribeClustersInput) error {
	return workflow.ExecuteActivity(ctx, a.activities.WaitUntilClusterRestored, input).Get(ctx, nil)
}

func (a *RedshiftStub) WaitUntilClusterRestoredAsync(ctx workflow.Context, input *redshift.DescribeClustersInput) workflow.Future {
	return workflow.ExecuteActivity(ctx, a.activities.WaitUntilClusterRestored, input)
}

func (a *RedshiftStub) WaitUntilSnapshotAvailable(ctx workflow.Context, input *redshift.DescribeClusterSnapshotsInput) error {
	return workflow.ExecuteActivity(ctx, a.activities.WaitUntilSnapshotAvailable, input).Get(ctx, nil)
}

func (a *RedshiftStub) WaitUntilSnapshotAvailableAsync(ctx workflow.Context, input *redshift.DescribeClusterSnapshotsInput) workflow.Future {
	return workflow.ExecuteActivity(ctx, a.activities.WaitUntilSnapshotAvailable, input)
}
