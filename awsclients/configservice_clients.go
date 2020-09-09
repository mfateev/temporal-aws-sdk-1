package awsclients

import (
	"github.com/aws/aws-sdk-go/service/configservice"
	"go.temporal.io/sdk/workflow"
	"temporal.io/aws-sdk/awsactivities"
)

type ConfigServiceClient interface {
       BatchGetAggregateResourceConfig(ctx workflow.Context, input *configservice.BatchGetAggregateResourceConfigInput) (*configservice.BatchGetAggregateResourceConfigOutput, error)
       BatchGetAggregateResourceConfigAsync(ctx workflow.Context, input *configservice.BatchGetAggregateResourceConfigInput) *ConfigserviceBatchGetAggregateResourceConfigResult

       BatchGetResourceConfig(ctx workflow.Context, input *configservice.BatchGetResourceConfigInput) (*configservice.BatchGetResourceConfigOutput, error)
       BatchGetResourceConfigAsync(ctx workflow.Context, input *configservice.BatchGetResourceConfigInput) *ConfigserviceBatchGetResourceConfigResult

       DeleteAggregationAuthorization(ctx workflow.Context, input *configservice.DeleteAggregationAuthorizationInput) (*configservice.DeleteAggregationAuthorizationOutput, error)
       DeleteAggregationAuthorizationAsync(ctx workflow.Context, input *configservice.DeleteAggregationAuthorizationInput) *ConfigserviceDeleteAggregationAuthorizationResult

       DeleteConfigRule(ctx workflow.Context, input *configservice.DeleteConfigRuleInput) (*configservice.DeleteConfigRuleOutput, error)
       DeleteConfigRuleAsync(ctx workflow.Context, input *configservice.DeleteConfigRuleInput) *ConfigserviceDeleteConfigRuleResult

       DeleteConfigurationAggregator(ctx workflow.Context, input *configservice.DeleteConfigurationAggregatorInput) (*configservice.DeleteConfigurationAggregatorOutput, error)
       DeleteConfigurationAggregatorAsync(ctx workflow.Context, input *configservice.DeleteConfigurationAggregatorInput) *ConfigserviceDeleteConfigurationAggregatorResult

       DeleteConfigurationRecorder(ctx workflow.Context, input *configservice.DeleteConfigurationRecorderInput) (*configservice.DeleteConfigurationRecorderOutput, error)
       DeleteConfigurationRecorderAsync(ctx workflow.Context, input *configservice.DeleteConfigurationRecorderInput) *ConfigserviceDeleteConfigurationRecorderResult

       DeleteConformancePack(ctx workflow.Context, input *configservice.DeleteConformancePackInput) (*configservice.DeleteConformancePackOutput, error)
       DeleteConformancePackAsync(ctx workflow.Context, input *configservice.DeleteConformancePackInput) *ConfigserviceDeleteConformancePackResult

       DeleteDeliveryChannel(ctx workflow.Context, input *configservice.DeleteDeliveryChannelInput) (*configservice.DeleteDeliveryChannelOutput, error)
       DeleteDeliveryChannelAsync(ctx workflow.Context, input *configservice.DeleteDeliveryChannelInput) *ConfigserviceDeleteDeliveryChannelResult

       DeleteEvaluationResults(ctx workflow.Context, input *configservice.DeleteEvaluationResultsInput) (*configservice.DeleteEvaluationResultsOutput, error)
       DeleteEvaluationResultsAsync(ctx workflow.Context, input *configservice.DeleteEvaluationResultsInput) *ConfigserviceDeleteEvaluationResultsResult

       DeleteOrganizationConfigRule(ctx workflow.Context, input *configservice.DeleteOrganizationConfigRuleInput) (*configservice.DeleteOrganizationConfigRuleOutput, error)
       DeleteOrganizationConfigRuleAsync(ctx workflow.Context, input *configservice.DeleteOrganizationConfigRuleInput) *ConfigserviceDeleteOrganizationConfigRuleResult

       DeleteOrganizationConformancePack(ctx workflow.Context, input *configservice.DeleteOrganizationConformancePackInput) (*configservice.DeleteOrganizationConformancePackOutput, error)
       DeleteOrganizationConformancePackAsync(ctx workflow.Context, input *configservice.DeleteOrganizationConformancePackInput) *ConfigserviceDeleteOrganizationConformancePackResult

       DeleteRemediationConfiguration(ctx workflow.Context, input *configservice.DeleteRemediationConfigurationInput) (*configservice.DeleteRemediationConfigurationOutput, error)
       DeleteRemediationConfigurationAsync(ctx workflow.Context, input *configservice.DeleteRemediationConfigurationInput) *ConfigserviceDeleteRemediationConfigurationResult

       DeleteRemediationExceptions(ctx workflow.Context, input *configservice.DeleteRemediationExceptionsInput) (*configservice.DeleteRemediationExceptionsOutput, error)
       DeleteRemediationExceptionsAsync(ctx workflow.Context, input *configservice.DeleteRemediationExceptionsInput) *ConfigserviceDeleteRemediationExceptionsResult

       DeleteResourceConfig(ctx workflow.Context, input *configservice.DeleteResourceConfigInput) (*configservice.DeleteResourceConfigOutput, error)
       DeleteResourceConfigAsync(ctx workflow.Context, input *configservice.DeleteResourceConfigInput) *ConfigserviceDeleteResourceConfigResult

       DeleteRetentionConfiguration(ctx workflow.Context, input *configservice.DeleteRetentionConfigurationInput) (*configservice.DeleteRetentionConfigurationOutput, error)
       DeleteRetentionConfigurationAsync(ctx workflow.Context, input *configservice.DeleteRetentionConfigurationInput) *ConfigserviceDeleteRetentionConfigurationResult

       DeliverConfigSnapshot(ctx workflow.Context, input *configservice.DeliverConfigSnapshotInput) (*configservice.DeliverConfigSnapshotOutput, error)
       DeliverConfigSnapshotAsync(ctx workflow.Context, input *configservice.DeliverConfigSnapshotInput) *ConfigserviceDeliverConfigSnapshotResult

       DescribeAggregateComplianceByConfigRules(ctx workflow.Context, input *configservice.DescribeAggregateComplianceByConfigRulesInput) (*configservice.DescribeAggregateComplianceByConfigRulesOutput, error)
       DescribeAggregateComplianceByConfigRulesAsync(ctx workflow.Context, input *configservice.DescribeAggregateComplianceByConfigRulesInput) *ConfigserviceDescribeAggregateComplianceByConfigRulesResult

       DescribeAggregationAuthorizations(ctx workflow.Context, input *configservice.DescribeAggregationAuthorizationsInput) (*configservice.DescribeAggregationAuthorizationsOutput, error)
       DescribeAggregationAuthorizationsAsync(ctx workflow.Context, input *configservice.DescribeAggregationAuthorizationsInput) *ConfigserviceDescribeAggregationAuthorizationsResult

       DescribeComplianceByConfigRule(ctx workflow.Context, input *configservice.DescribeComplianceByConfigRuleInput) (*configservice.DescribeComplianceByConfigRuleOutput, error)
       DescribeComplianceByConfigRuleAsync(ctx workflow.Context, input *configservice.DescribeComplianceByConfigRuleInput) *ConfigserviceDescribeComplianceByConfigRuleResult

       DescribeComplianceByResource(ctx workflow.Context, input *configservice.DescribeComplianceByResourceInput) (*configservice.DescribeComplianceByResourceOutput, error)
       DescribeComplianceByResourceAsync(ctx workflow.Context, input *configservice.DescribeComplianceByResourceInput) *ConfigserviceDescribeComplianceByResourceResult

       DescribeConfigRuleEvaluationStatus(ctx workflow.Context, input *configservice.DescribeConfigRuleEvaluationStatusInput) (*configservice.DescribeConfigRuleEvaluationStatusOutput, error)
       DescribeConfigRuleEvaluationStatusAsync(ctx workflow.Context, input *configservice.DescribeConfigRuleEvaluationStatusInput) *ConfigserviceDescribeConfigRuleEvaluationStatusResult

       DescribeConfigRules(ctx workflow.Context, input *configservice.DescribeConfigRulesInput) (*configservice.DescribeConfigRulesOutput, error)
       DescribeConfigRulesAsync(ctx workflow.Context, input *configservice.DescribeConfigRulesInput) *ConfigserviceDescribeConfigRulesResult

       DescribeConfigurationAggregatorSourcesStatus(ctx workflow.Context, input *configservice.DescribeConfigurationAggregatorSourcesStatusInput) (*configservice.DescribeConfigurationAggregatorSourcesStatusOutput, error)
       DescribeConfigurationAggregatorSourcesStatusAsync(ctx workflow.Context, input *configservice.DescribeConfigurationAggregatorSourcesStatusInput) *ConfigserviceDescribeConfigurationAggregatorSourcesStatusResult

       DescribeConfigurationAggregators(ctx workflow.Context, input *configservice.DescribeConfigurationAggregatorsInput) (*configservice.DescribeConfigurationAggregatorsOutput, error)
       DescribeConfigurationAggregatorsAsync(ctx workflow.Context, input *configservice.DescribeConfigurationAggregatorsInput) *ConfigserviceDescribeConfigurationAggregatorsResult

       DescribeConfigurationRecorderStatus(ctx workflow.Context, input *configservice.DescribeConfigurationRecorderStatusInput) (*configservice.DescribeConfigurationRecorderStatusOutput, error)
       DescribeConfigurationRecorderStatusAsync(ctx workflow.Context, input *configservice.DescribeConfigurationRecorderStatusInput) *ConfigserviceDescribeConfigurationRecorderStatusResult

       DescribeConfigurationRecorders(ctx workflow.Context, input *configservice.DescribeConfigurationRecordersInput) (*configservice.DescribeConfigurationRecordersOutput, error)
       DescribeConfigurationRecordersAsync(ctx workflow.Context, input *configservice.DescribeConfigurationRecordersInput) *ConfigserviceDescribeConfigurationRecordersResult

       DescribeConformancePackCompliance(ctx workflow.Context, input *configservice.DescribeConformancePackComplianceInput) (*configservice.DescribeConformancePackComplianceOutput, error)
       DescribeConformancePackComplianceAsync(ctx workflow.Context, input *configservice.DescribeConformancePackComplianceInput) *ConfigserviceDescribeConformancePackComplianceResult

       DescribeConformancePackStatus(ctx workflow.Context, input *configservice.DescribeConformancePackStatusInput) (*configservice.DescribeConformancePackStatusOutput, error)
       DescribeConformancePackStatusAsync(ctx workflow.Context, input *configservice.DescribeConformancePackStatusInput) *ConfigserviceDescribeConformancePackStatusResult

       DescribeConformancePacks(ctx workflow.Context, input *configservice.DescribeConformancePacksInput) (*configservice.DescribeConformancePacksOutput, error)
       DescribeConformancePacksAsync(ctx workflow.Context, input *configservice.DescribeConformancePacksInput) *ConfigserviceDescribeConformancePacksResult

       DescribeDeliveryChannelStatus(ctx workflow.Context, input *configservice.DescribeDeliveryChannelStatusInput) (*configservice.DescribeDeliveryChannelStatusOutput, error)
       DescribeDeliveryChannelStatusAsync(ctx workflow.Context, input *configservice.DescribeDeliveryChannelStatusInput) *ConfigserviceDescribeDeliveryChannelStatusResult

       DescribeDeliveryChannels(ctx workflow.Context, input *configservice.DescribeDeliveryChannelsInput) (*configservice.DescribeDeliveryChannelsOutput, error)
       DescribeDeliveryChannelsAsync(ctx workflow.Context, input *configservice.DescribeDeliveryChannelsInput) *ConfigserviceDescribeDeliveryChannelsResult

       DescribeOrganizationConfigRuleStatuses(ctx workflow.Context, input *configservice.DescribeOrganizationConfigRuleStatusesInput) (*configservice.DescribeOrganizationConfigRuleStatusesOutput, error)
       DescribeOrganizationConfigRuleStatusesAsync(ctx workflow.Context, input *configservice.DescribeOrganizationConfigRuleStatusesInput) *ConfigserviceDescribeOrganizationConfigRuleStatusesResult

       DescribeOrganizationConfigRules(ctx workflow.Context, input *configservice.DescribeOrganizationConfigRulesInput) (*configservice.DescribeOrganizationConfigRulesOutput, error)
       DescribeOrganizationConfigRulesAsync(ctx workflow.Context, input *configservice.DescribeOrganizationConfigRulesInput) *ConfigserviceDescribeOrganizationConfigRulesResult

       DescribeOrganizationConformancePackStatuses(ctx workflow.Context, input *configservice.DescribeOrganizationConformancePackStatusesInput) (*configservice.DescribeOrganizationConformancePackStatusesOutput, error)
       DescribeOrganizationConformancePackStatusesAsync(ctx workflow.Context, input *configservice.DescribeOrganizationConformancePackStatusesInput) *ConfigserviceDescribeOrganizationConformancePackStatusesResult

       DescribeOrganizationConformancePacks(ctx workflow.Context, input *configservice.DescribeOrganizationConformancePacksInput) (*configservice.DescribeOrganizationConformancePacksOutput, error)
       DescribeOrganizationConformancePacksAsync(ctx workflow.Context, input *configservice.DescribeOrganizationConformancePacksInput) *ConfigserviceDescribeOrganizationConformancePacksResult

       DescribePendingAggregationRequests(ctx workflow.Context, input *configservice.DescribePendingAggregationRequestsInput) (*configservice.DescribePendingAggregationRequestsOutput, error)
       DescribePendingAggregationRequestsAsync(ctx workflow.Context, input *configservice.DescribePendingAggregationRequestsInput) *ConfigserviceDescribePendingAggregationRequestsResult

       DescribeRemediationConfigurations(ctx workflow.Context, input *configservice.DescribeRemediationConfigurationsInput) (*configservice.DescribeRemediationConfigurationsOutput, error)
       DescribeRemediationConfigurationsAsync(ctx workflow.Context, input *configservice.DescribeRemediationConfigurationsInput) *ConfigserviceDescribeRemediationConfigurationsResult

       DescribeRemediationExceptions(ctx workflow.Context, input *configservice.DescribeRemediationExceptionsInput) (*configservice.DescribeRemediationExceptionsOutput, error)
       DescribeRemediationExceptionsAsync(ctx workflow.Context, input *configservice.DescribeRemediationExceptionsInput) *ConfigserviceDescribeRemediationExceptionsResult

       DescribeRemediationExecutionStatus(ctx workflow.Context, input *configservice.DescribeRemediationExecutionStatusInput) (*configservice.DescribeRemediationExecutionStatusOutput, error)
       DescribeRemediationExecutionStatusAsync(ctx workflow.Context, input *configservice.DescribeRemediationExecutionStatusInput) *ConfigserviceDescribeRemediationExecutionStatusResult

       DescribeRetentionConfigurations(ctx workflow.Context, input *configservice.DescribeRetentionConfigurationsInput) (*configservice.DescribeRetentionConfigurationsOutput, error)
       DescribeRetentionConfigurationsAsync(ctx workflow.Context, input *configservice.DescribeRetentionConfigurationsInput) *ConfigserviceDescribeRetentionConfigurationsResult

       GetAggregateComplianceDetailsByConfigRule(ctx workflow.Context, input *configservice.GetAggregateComplianceDetailsByConfigRuleInput) (*configservice.GetAggregateComplianceDetailsByConfigRuleOutput, error)
       GetAggregateComplianceDetailsByConfigRuleAsync(ctx workflow.Context, input *configservice.GetAggregateComplianceDetailsByConfigRuleInput) *ConfigserviceGetAggregateComplianceDetailsByConfigRuleResult

       GetAggregateConfigRuleComplianceSummary(ctx workflow.Context, input *configservice.GetAggregateConfigRuleComplianceSummaryInput) (*configservice.GetAggregateConfigRuleComplianceSummaryOutput, error)
       GetAggregateConfigRuleComplianceSummaryAsync(ctx workflow.Context, input *configservice.GetAggregateConfigRuleComplianceSummaryInput) *ConfigserviceGetAggregateConfigRuleComplianceSummaryResult

       GetAggregateDiscoveredResourceCounts(ctx workflow.Context, input *configservice.GetAggregateDiscoveredResourceCountsInput) (*configservice.GetAggregateDiscoveredResourceCountsOutput, error)
       GetAggregateDiscoveredResourceCountsAsync(ctx workflow.Context, input *configservice.GetAggregateDiscoveredResourceCountsInput) *ConfigserviceGetAggregateDiscoveredResourceCountsResult

       GetAggregateResourceConfig(ctx workflow.Context, input *configservice.GetAggregateResourceConfigInput) (*configservice.GetAggregateResourceConfigOutput, error)
       GetAggregateResourceConfigAsync(ctx workflow.Context, input *configservice.GetAggregateResourceConfigInput) *ConfigserviceGetAggregateResourceConfigResult

       GetComplianceDetailsByConfigRule(ctx workflow.Context, input *configservice.GetComplianceDetailsByConfigRuleInput) (*configservice.GetComplianceDetailsByConfigRuleOutput, error)
       GetComplianceDetailsByConfigRuleAsync(ctx workflow.Context, input *configservice.GetComplianceDetailsByConfigRuleInput) *ConfigserviceGetComplianceDetailsByConfigRuleResult

       GetComplianceDetailsByResource(ctx workflow.Context, input *configservice.GetComplianceDetailsByResourceInput) (*configservice.GetComplianceDetailsByResourceOutput, error)
       GetComplianceDetailsByResourceAsync(ctx workflow.Context, input *configservice.GetComplianceDetailsByResourceInput) *ConfigserviceGetComplianceDetailsByResourceResult

       GetComplianceSummaryByConfigRule(ctx workflow.Context, input *configservice.GetComplianceSummaryByConfigRuleInput) (*configservice.GetComplianceSummaryByConfigRuleOutput, error)
       GetComplianceSummaryByConfigRuleAsync(ctx workflow.Context, input *configservice.GetComplianceSummaryByConfigRuleInput) *ConfigserviceGetComplianceSummaryByConfigRuleResult

       GetComplianceSummaryByResourceType(ctx workflow.Context, input *configservice.GetComplianceSummaryByResourceTypeInput) (*configservice.GetComplianceSummaryByResourceTypeOutput, error)
       GetComplianceSummaryByResourceTypeAsync(ctx workflow.Context, input *configservice.GetComplianceSummaryByResourceTypeInput) *ConfigserviceGetComplianceSummaryByResourceTypeResult

       GetConformancePackComplianceDetails(ctx workflow.Context, input *configservice.GetConformancePackComplianceDetailsInput) (*configservice.GetConformancePackComplianceDetailsOutput, error)
       GetConformancePackComplianceDetailsAsync(ctx workflow.Context, input *configservice.GetConformancePackComplianceDetailsInput) *ConfigserviceGetConformancePackComplianceDetailsResult

       GetConformancePackComplianceSummary(ctx workflow.Context, input *configservice.GetConformancePackComplianceSummaryInput) (*configservice.GetConformancePackComplianceSummaryOutput, error)
       GetConformancePackComplianceSummaryAsync(ctx workflow.Context, input *configservice.GetConformancePackComplianceSummaryInput) *ConfigserviceGetConformancePackComplianceSummaryResult

       GetDiscoveredResourceCounts(ctx workflow.Context, input *configservice.GetDiscoveredResourceCountsInput) (*configservice.GetDiscoveredResourceCountsOutput, error)
       GetDiscoveredResourceCountsAsync(ctx workflow.Context, input *configservice.GetDiscoveredResourceCountsInput) *ConfigserviceGetDiscoveredResourceCountsResult

       GetOrganizationConfigRuleDetailedStatus(ctx workflow.Context, input *configservice.GetOrganizationConfigRuleDetailedStatusInput) (*configservice.GetOrganizationConfigRuleDetailedStatusOutput, error)
       GetOrganizationConfigRuleDetailedStatusAsync(ctx workflow.Context, input *configservice.GetOrganizationConfigRuleDetailedStatusInput) *ConfigserviceGetOrganizationConfigRuleDetailedStatusResult

       GetOrganizationConformancePackDetailedStatus(ctx workflow.Context, input *configservice.GetOrganizationConformancePackDetailedStatusInput) (*configservice.GetOrganizationConformancePackDetailedStatusOutput, error)
       GetOrganizationConformancePackDetailedStatusAsync(ctx workflow.Context, input *configservice.GetOrganizationConformancePackDetailedStatusInput) *ConfigserviceGetOrganizationConformancePackDetailedStatusResult

       GetResourceConfigHistory(ctx workflow.Context, input *configservice.GetResourceConfigHistoryInput) (*configservice.GetResourceConfigHistoryOutput, error)
       GetResourceConfigHistoryAsync(ctx workflow.Context, input *configservice.GetResourceConfigHistoryInput) *ConfigserviceGetResourceConfigHistoryResult

       ListAggregateDiscoveredResources(ctx workflow.Context, input *configservice.ListAggregateDiscoveredResourcesInput) (*configservice.ListAggregateDiscoveredResourcesOutput, error)
       ListAggregateDiscoveredResourcesAsync(ctx workflow.Context, input *configservice.ListAggregateDiscoveredResourcesInput) *ConfigserviceListAggregateDiscoveredResourcesResult

       ListDiscoveredResources(ctx workflow.Context, input *configservice.ListDiscoveredResourcesInput) (*configservice.ListDiscoveredResourcesOutput, error)
       ListDiscoveredResourcesAsync(ctx workflow.Context, input *configservice.ListDiscoveredResourcesInput) *ConfigserviceListDiscoveredResourcesResult

       ListTagsForResource(ctx workflow.Context, input *configservice.ListTagsForResourceInput) (*configservice.ListTagsForResourceOutput, error)
       ListTagsForResourceAsync(ctx workflow.Context, input *configservice.ListTagsForResourceInput) *ConfigserviceListTagsForResourceResult

       PutAggregationAuthorization(ctx workflow.Context, input *configservice.PutAggregationAuthorizationInput) (*configservice.PutAggregationAuthorizationOutput, error)
       PutAggregationAuthorizationAsync(ctx workflow.Context, input *configservice.PutAggregationAuthorizationInput) *ConfigservicePutAggregationAuthorizationResult

       PutConfigRule(ctx workflow.Context, input *configservice.PutConfigRuleInput) (*configservice.PutConfigRuleOutput, error)
       PutConfigRuleAsync(ctx workflow.Context, input *configservice.PutConfigRuleInput) *ConfigservicePutConfigRuleResult

       PutConfigurationAggregator(ctx workflow.Context, input *configservice.PutConfigurationAggregatorInput) (*configservice.PutConfigurationAggregatorOutput, error)
       PutConfigurationAggregatorAsync(ctx workflow.Context, input *configservice.PutConfigurationAggregatorInput) *ConfigservicePutConfigurationAggregatorResult

       PutConfigurationRecorder(ctx workflow.Context, input *configservice.PutConfigurationRecorderInput) (*configservice.PutConfigurationRecorderOutput, error)
       PutConfigurationRecorderAsync(ctx workflow.Context, input *configservice.PutConfigurationRecorderInput) *ConfigservicePutConfigurationRecorderResult

       PutConformancePack(ctx workflow.Context, input *configservice.PutConformancePackInput) (*configservice.PutConformancePackOutput, error)
       PutConformancePackAsync(ctx workflow.Context, input *configservice.PutConformancePackInput) *ConfigservicePutConformancePackResult

       PutDeliveryChannel(ctx workflow.Context, input *configservice.PutDeliveryChannelInput) (*configservice.PutDeliveryChannelOutput, error)
       PutDeliveryChannelAsync(ctx workflow.Context, input *configservice.PutDeliveryChannelInput) *ConfigservicePutDeliveryChannelResult

       PutEvaluations(ctx workflow.Context, input *configservice.PutEvaluationsInput) (*configservice.PutEvaluationsOutput, error)
       PutEvaluationsAsync(ctx workflow.Context, input *configservice.PutEvaluationsInput) *ConfigservicePutEvaluationsResult

       PutOrganizationConfigRule(ctx workflow.Context, input *configservice.PutOrganizationConfigRuleInput) (*configservice.PutOrganizationConfigRuleOutput, error)
       PutOrganizationConfigRuleAsync(ctx workflow.Context, input *configservice.PutOrganizationConfigRuleInput) *ConfigservicePutOrganizationConfigRuleResult

       PutOrganizationConformancePack(ctx workflow.Context, input *configservice.PutOrganizationConformancePackInput) (*configservice.PutOrganizationConformancePackOutput, error)
       PutOrganizationConformancePackAsync(ctx workflow.Context, input *configservice.PutOrganizationConformancePackInput) *ConfigservicePutOrganizationConformancePackResult

       PutRemediationConfigurations(ctx workflow.Context, input *configservice.PutRemediationConfigurationsInput) (*configservice.PutRemediationConfigurationsOutput, error)
       PutRemediationConfigurationsAsync(ctx workflow.Context, input *configservice.PutRemediationConfigurationsInput) *ConfigservicePutRemediationConfigurationsResult

       PutRemediationExceptions(ctx workflow.Context, input *configservice.PutRemediationExceptionsInput) (*configservice.PutRemediationExceptionsOutput, error)
       PutRemediationExceptionsAsync(ctx workflow.Context, input *configservice.PutRemediationExceptionsInput) *ConfigservicePutRemediationExceptionsResult

       PutResourceConfig(ctx workflow.Context, input *configservice.PutResourceConfigInput) (*configservice.PutResourceConfigOutput, error)
       PutResourceConfigAsync(ctx workflow.Context, input *configservice.PutResourceConfigInput) *ConfigservicePutResourceConfigResult

       PutRetentionConfiguration(ctx workflow.Context, input *configservice.PutRetentionConfigurationInput) (*configservice.PutRetentionConfigurationOutput, error)
       PutRetentionConfigurationAsync(ctx workflow.Context, input *configservice.PutRetentionConfigurationInput) *ConfigservicePutRetentionConfigurationResult

       SelectAggregateResourceConfig(ctx workflow.Context, input *configservice.SelectAggregateResourceConfigInput) (*configservice.SelectAggregateResourceConfigOutput, error)
       SelectAggregateResourceConfigAsync(ctx workflow.Context, input *configservice.SelectAggregateResourceConfigInput) *ConfigserviceSelectAggregateResourceConfigResult

       SelectResourceConfig(ctx workflow.Context, input *configservice.SelectResourceConfigInput) (*configservice.SelectResourceConfigOutput, error)
       SelectResourceConfigAsync(ctx workflow.Context, input *configservice.SelectResourceConfigInput) *ConfigserviceSelectResourceConfigResult

       StartConfigRulesEvaluation(ctx workflow.Context, input *configservice.StartConfigRulesEvaluationInput) (*configservice.StartConfigRulesEvaluationOutput, error)
       StartConfigRulesEvaluationAsync(ctx workflow.Context, input *configservice.StartConfigRulesEvaluationInput) *ConfigserviceStartConfigRulesEvaluationResult

       StartConfigurationRecorder(ctx workflow.Context, input *configservice.StartConfigurationRecorderInput) (*configservice.StartConfigurationRecorderOutput, error)
       StartConfigurationRecorderAsync(ctx workflow.Context, input *configservice.StartConfigurationRecorderInput) *ConfigserviceStartConfigurationRecorderResult

       StartRemediationExecution(ctx workflow.Context, input *configservice.StartRemediationExecutionInput) (*configservice.StartRemediationExecutionOutput, error)
       StartRemediationExecutionAsync(ctx workflow.Context, input *configservice.StartRemediationExecutionInput) *ConfigserviceStartRemediationExecutionResult

       StopConfigurationRecorder(ctx workflow.Context, input *configservice.StopConfigurationRecorderInput) (*configservice.StopConfigurationRecorderOutput, error)
       StopConfigurationRecorderAsync(ctx workflow.Context, input *configservice.StopConfigurationRecorderInput) *ConfigserviceStopConfigurationRecorderResult

       TagResource(ctx workflow.Context, input *configservice.TagResourceInput) (*configservice.TagResourceOutput, error)
       TagResourceAsync(ctx workflow.Context, input *configservice.TagResourceInput) *ConfigserviceTagResourceResult

       UntagResource(ctx workflow.Context, input *configservice.UntagResourceInput) (*configservice.UntagResourceOutput, error)
       UntagResourceAsync(ctx workflow.Context, input *configservice.UntagResourceInput) *ConfigserviceUntagResourceResult
}

type ConfigserviceBatchGetAggregateResourceConfigResult struct {
	Result workflow.Future
}

func (r *ConfigserviceBatchGetAggregateResourceConfigResult) Get(ctx workflow.Context) (*configservice.BatchGetAggregateResourceConfigOutput, error) {
    var output configservice.BatchGetAggregateResourceConfigOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ConfigserviceBatchGetResourceConfigResult struct {
	Result workflow.Future
}

func (r *ConfigserviceBatchGetResourceConfigResult) Get(ctx workflow.Context) (*configservice.BatchGetResourceConfigOutput, error) {
    var output configservice.BatchGetResourceConfigOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ConfigserviceDeleteAggregationAuthorizationResult struct {
	Result workflow.Future
}

func (r *ConfigserviceDeleteAggregationAuthorizationResult) Get(ctx workflow.Context) (*configservice.DeleteAggregationAuthorizationOutput, error) {
    var output configservice.DeleteAggregationAuthorizationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ConfigserviceDeleteConfigRuleResult struct {
	Result workflow.Future
}

func (r *ConfigserviceDeleteConfigRuleResult) Get(ctx workflow.Context) (*configservice.DeleteConfigRuleOutput, error) {
    var output configservice.DeleteConfigRuleOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ConfigserviceDeleteConfigurationAggregatorResult struct {
	Result workflow.Future
}

func (r *ConfigserviceDeleteConfigurationAggregatorResult) Get(ctx workflow.Context) (*configservice.DeleteConfigurationAggregatorOutput, error) {
    var output configservice.DeleteConfigurationAggregatorOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ConfigserviceDeleteConfigurationRecorderResult struct {
	Result workflow.Future
}

func (r *ConfigserviceDeleteConfigurationRecorderResult) Get(ctx workflow.Context) (*configservice.DeleteConfigurationRecorderOutput, error) {
    var output configservice.DeleteConfigurationRecorderOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ConfigserviceDeleteConformancePackResult struct {
	Result workflow.Future
}

func (r *ConfigserviceDeleteConformancePackResult) Get(ctx workflow.Context) (*configservice.DeleteConformancePackOutput, error) {
    var output configservice.DeleteConformancePackOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ConfigserviceDeleteDeliveryChannelResult struct {
	Result workflow.Future
}

func (r *ConfigserviceDeleteDeliveryChannelResult) Get(ctx workflow.Context) (*configservice.DeleteDeliveryChannelOutput, error) {
    var output configservice.DeleteDeliveryChannelOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ConfigserviceDeleteEvaluationResultsResult struct {
	Result workflow.Future
}

func (r *ConfigserviceDeleteEvaluationResultsResult) Get(ctx workflow.Context) (*configservice.DeleteEvaluationResultsOutput, error) {
    var output configservice.DeleteEvaluationResultsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ConfigserviceDeleteOrganizationConfigRuleResult struct {
	Result workflow.Future
}

func (r *ConfigserviceDeleteOrganizationConfigRuleResult) Get(ctx workflow.Context) (*configservice.DeleteOrganizationConfigRuleOutput, error) {
    var output configservice.DeleteOrganizationConfigRuleOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ConfigserviceDeleteOrganizationConformancePackResult struct {
	Result workflow.Future
}

func (r *ConfigserviceDeleteOrganizationConformancePackResult) Get(ctx workflow.Context) (*configservice.DeleteOrganizationConformancePackOutput, error) {
    var output configservice.DeleteOrganizationConformancePackOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ConfigserviceDeleteRemediationConfigurationResult struct {
	Result workflow.Future
}

func (r *ConfigserviceDeleteRemediationConfigurationResult) Get(ctx workflow.Context) (*configservice.DeleteRemediationConfigurationOutput, error) {
    var output configservice.DeleteRemediationConfigurationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ConfigserviceDeleteRemediationExceptionsResult struct {
	Result workflow.Future
}

func (r *ConfigserviceDeleteRemediationExceptionsResult) Get(ctx workflow.Context) (*configservice.DeleteRemediationExceptionsOutput, error) {
    var output configservice.DeleteRemediationExceptionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ConfigserviceDeleteResourceConfigResult struct {
	Result workflow.Future
}

func (r *ConfigserviceDeleteResourceConfigResult) Get(ctx workflow.Context) (*configservice.DeleteResourceConfigOutput, error) {
    var output configservice.DeleteResourceConfigOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ConfigserviceDeleteRetentionConfigurationResult struct {
	Result workflow.Future
}

func (r *ConfigserviceDeleteRetentionConfigurationResult) Get(ctx workflow.Context) (*configservice.DeleteRetentionConfigurationOutput, error) {
    var output configservice.DeleteRetentionConfigurationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ConfigserviceDeliverConfigSnapshotResult struct {
	Result workflow.Future
}

func (r *ConfigserviceDeliverConfigSnapshotResult) Get(ctx workflow.Context) (*configservice.DeliverConfigSnapshotOutput, error) {
    var output configservice.DeliverConfigSnapshotOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ConfigserviceDescribeAggregateComplianceByConfigRulesResult struct {
	Result workflow.Future
}

func (r *ConfigserviceDescribeAggregateComplianceByConfigRulesResult) Get(ctx workflow.Context) (*configservice.DescribeAggregateComplianceByConfigRulesOutput, error) {
    var output configservice.DescribeAggregateComplianceByConfigRulesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ConfigserviceDescribeAggregationAuthorizationsResult struct {
	Result workflow.Future
}

func (r *ConfigserviceDescribeAggregationAuthorizationsResult) Get(ctx workflow.Context) (*configservice.DescribeAggregationAuthorizationsOutput, error) {
    var output configservice.DescribeAggregationAuthorizationsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ConfigserviceDescribeComplianceByConfigRuleResult struct {
	Result workflow.Future
}

func (r *ConfigserviceDescribeComplianceByConfigRuleResult) Get(ctx workflow.Context) (*configservice.DescribeComplianceByConfigRuleOutput, error) {
    var output configservice.DescribeComplianceByConfigRuleOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ConfigserviceDescribeComplianceByResourceResult struct {
	Result workflow.Future
}

func (r *ConfigserviceDescribeComplianceByResourceResult) Get(ctx workflow.Context) (*configservice.DescribeComplianceByResourceOutput, error) {
    var output configservice.DescribeComplianceByResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ConfigserviceDescribeConfigRuleEvaluationStatusResult struct {
	Result workflow.Future
}

func (r *ConfigserviceDescribeConfigRuleEvaluationStatusResult) Get(ctx workflow.Context) (*configservice.DescribeConfigRuleEvaluationStatusOutput, error) {
    var output configservice.DescribeConfigRuleEvaluationStatusOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ConfigserviceDescribeConfigRulesResult struct {
	Result workflow.Future
}

func (r *ConfigserviceDescribeConfigRulesResult) Get(ctx workflow.Context) (*configservice.DescribeConfigRulesOutput, error) {
    var output configservice.DescribeConfigRulesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ConfigserviceDescribeConfigurationAggregatorSourcesStatusResult struct {
	Result workflow.Future
}

func (r *ConfigserviceDescribeConfigurationAggregatorSourcesStatusResult) Get(ctx workflow.Context) (*configservice.DescribeConfigurationAggregatorSourcesStatusOutput, error) {
    var output configservice.DescribeConfigurationAggregatorSourcesStatusOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ConfigserviceDescribeConfigurationAggregatorsResult struct {
	Result workflow.Future
}

func (r *ConfigserviceDescribeConfigurationAggregatorsResult) Get(ctx workflow.Context) (*configservice.DescribeConfigurationAggregatorsOutput, error) {
    var output configservice.DescribeConfigurationAggregatorsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ConfigserviceDescribeConfigurationRecorderStatusResult struct {
	Result workflow.Future
}

func (r *ConfigserviceDescribeConfigurationRecorderStatusResult) Get(ctx workflow.Context) (*configservice.DescribeConfigurationRecorderStatusOutput, error) {
    var output configservice.DescribeConfigurationRecorderStatusOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ConfigserviceDescribeConfigurationRecordersResult struct {
	Result workflow.Future
}

func (r *ConfigserviceDescribeConfigurationRecordersResult) Get(ctx workflow.Context) (*configservice.DescribeConfigurationRecordersOutput, error) {
    var output configservice.DescribeConfigurationRecordersOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ConfigserviceDescribeConformancePackComplianceResult struct {
	Result workflow.Future
}

func (r *ConfigserviceDescribeConformancePackComplianceResult) Get(ctx workflow.Context) (*configservice.DescribeConformancePackComplianceOutput, error) {
    var output configservice.DescribeConformancePackComplianceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ConfigserviceDescribeConformancePackStatusResult struct {
	Result workflow.Future
}

func (r *ConfigserviceDescribeConformancePackStatusResult) Get(ctx workflow.Context) (*configservice.DescribeConformancePackStatusOutput, error) {
    var output configservice.DescribeConformancePackStatusOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ConfigserviceDescribeConformancePacksResult struct {
	Result workflow.Future
}

func (r *ConfigserviceDescribeConformancePacksResult) Get(ctx workflow.Context) (*configservice.DescribeConformancePacksOutput, error) {
    var output configservice.DescribeConformancePacksOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ConfigserviceDescribeDeliveryChannelStatusResult struct {
	Result workflow.Future
}

func (r *ConfigserviceDescribeDeliveryChannelStatusResult) Get(ctx workflow.Context) (*configservice.DescribeDeliveryChannelStatusOutput, error) {
    var output configservice.DescribeDeliveryChannelStatusOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ConfigserviceDescribeDeliveryChannelsResult struct {
	Result workflow.Future
}

func (r *ConfigserviceDescribeDeliveryChannelsResult) Get(ctx workflow.Context) (*configservice.DescribeDeliveryChannelsOutput, error) {
    var output configservice.DescribeDeliveryChannelsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ConfigserviceDescribeOrganizationConfigRuleStatusesResult struct {
	Result workflow.Future
}

func (r *ConfigserviceDescribeOrganizationConfigRuleStatusesResult) Get(ctx workflow.Context) (*configservice.DescribeOrganizationConfigRuleStatusesOutput, error) {
    var output configservice.DescribeOrganizationConfigRuleStatusesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ConfigserviceDescribeOrganizationConfigRulesResult struct {
	Result workflow.Future
}

func (r *ConfigserviceDescribeOrganizationConfigRulesResult) Get(ctx workflow.Context) (*configservice.DescribeOrganizationConfigRulesOutput, error) {
    var output configservice.DescribeOrganizationConfigRulesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ConfigserviceDescribeOrganizationConformancePackStatusesResult struct {
	Result workflow.Future
}

func (r *ConfigserviceDescribeOrganizationConformancePackStatusesResult) Get(ctx workflow.Context) (*configservice.DescribeOrganizationConformancePackStatusesOutput, error) {
    var output configservice.DescribeOrganizationConformancePackStatusesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ConfigserviceDescribeOrganizationConformancePacksResult struct {
	Result workflow.Future
}

func (r *ConfigserviceDescribeOrganizationConformancePacksResult) Get(ctx workflow.Context) (*configservice.DescribeOrganizationConformancePacksOutput, error) {
    var output configservice.DescribeOrganizationConformancePacksOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ConfigserviceDescribePendingAggregationRequestsResult struct {
	Result workflow.Future
}

func (r *ConfigserviceDescribePendingAggregationRequestsResult) Get(ctx workflow.Context) (*configservice.DescribePendingAggregationRequestsOutput, error) {
    var output configservice.DescribePendingAggregationRequestsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ConfigserviceDescribeRemediationConfigurationsResult struct {
	Result workflow.Future
}

func (r *ConfigserviceDescribeRemediationConfigurationsResult) Get(ctx workflow.Context) (*configservice.DescribeRemediationConfigurationsOutput, error) {
    var output configservice.DescribeRemediationConfigurationsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ConfigserviceDescribeRemediationExceptionsResult struct {
	Result workflow.Future
}

func (r *ConfigserviceDescribeRemediationExceptionsResult) Get(ctx workflow.Context) (*configservice.DescribeRemediationExceptionsOutput, error) {
    var output configservice.DescribeRemediationExceptionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ConfigserviceDescribeRemediationExecutionStatusResult struct {
	Result workflow.Future
}

func (r *ConfigserviceDescribeRemediationExecutionStatusResult) Get(ctx workflow.Context) (*configservice.DescribeRemediationExecutionStatusOutput, error) {
    var output configservice.DescribeRemediationExecutionStatusOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ConfigserviceDescribeRetentionConfigurationsResult struct {
	Result workflow.Future
}

func (r *ConfigserviceDescribeRetentionConfigurationsResult) Get(ctx workflow.Context) (*configservice.DescribeRetentionConfigurationsOutput, error) {
    var output configservice.DescribeRetentionConfigurationsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ConfigserviceGetAggregateComplianceDetailsByConfigRuleResult struct {
	Result workflow.Future
}

func (r *ConfigserviceGetAggregateComplianceDetailsByConfigRuleResult) Get(ctx workflow.Context) (*configservice.GetAggregateComplianceDetailsByConfigRuleOutput, error) {
    var output configservice.GetAggregateComplianceDetailsByConfigRuleOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ConfigserviceGetAggregateConfigRuleComplianceSummaryResult struct {
	Result workflow.Future
}

func (r *ConfigserviceGetAggregateConfigRuleComplianceSummaryResult) Get(ctx workflow.Context) (*configservice.GetAggregateConfigRuleComplianceSummaryOutput, error) {
    var output configservice.GetAggregateConfigRuleComplianceSummaryOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ConfigserviceGetAggregateDiscoveredResourceCountsResult struct {
	Result workflow.Future
}

func (r *ConfigserviceGetAggregateDiscoveredResourceCountsResult) Get(ctx workflow.Context) (*configservice.GetAggregateDiscoveredResourceCountsOutput, error) {
    var output configservice.GetAggregateDiscoveredResourceCountsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ConfigserviceGetAggregateResourceConfigResult struct {
	Result workflow.Future
}

func (r *ConfigserviceGetAggregateResourceConfigResult) Get(ctx workflow.Context) (*configservice.GetAggregateResourceConfigOutput, error) {
    var output configservice.GetAggregateResourceConfigOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ConfigserviceGetComplianceDetailsByConfigRuleResult struct {
	Result workflow.Future
}

func (r *ConfigserviceGetComplianceDetailsByConfigRuleResult) Get(ctx workflow.Context) (*configservice.GetComplianceDetailsByConfigRuleOutput, error) {
    var output configservice.GetComplianceDetailsByConfigRuleOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ConfigserviceGetComplianceDetailsByResourceResult struct {
	Result workflow.Future
}

func (r *ConfigserviceGetComplianceDetailsByResourceResult) Get(ctx workflow.Context) (*configservice.GetComplianceDetailsByResourceOutput, error) {
    var output configservice.GetComplianceDetailsByResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ConfigserviceGetComplianceSummaryByConfigRuleResult struct {
	Result workflow.Future
}

func (r *ConfigserviceGetComplianceSummaryByConfigRuleResult) Get(ctx workflow.Context) (*configservice.GetComplianceSummaryByConfigRuleOutput, error) {
    var output configservice.GetComplianceSummaryByConfigRuleOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ConfigserviceGetComplianceSummaryByResourceTypeResult struct {
	Result workflow.Future
}

func (r *ConfigserviceGetComplianceSummaryByResourceTypeResult) Get(ctx workflow.Context) (*configservice.GetComplianceSummaryByResourceTypeOutput, error) {
    var output configservice.GetComplianceSummaryByResourceTypeOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ConfigserviceGetConformancePackComplianceDetailsResult struct {
	Result workflow.Future
}

func (r *ConfigserviceGetConformancePackComplianceDetailsResult) Get(ctx workflow.Context) (*configservice.GetConformancePackComplianceDetailsOutput, error) {
    var output configservice.GetConformancePackComplianceDetailsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ConfigserviceGetConformancePackComplianceSummaryResult struct {
	Result workflow.Future
}

func (r *ConfigserviceGetConformancePackComplianceSummaryResult) Get(ctx workflow.Context) (*configservice.GetConformancePackComplianceSummaryOutput, error) {
    var output configservice.GetConformancePackComplianceSummaryOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ConfigserviceGetDiscoveredResourceCountsResult struct {
	Result workflow.Future
}

func (r *ConfigserviceGetDiscoveredResourceCountsResult) Get(ctx workflow.Context) (*configservice.GetDiscoveredResourceCountsOutput, error) {
    var output configservice.GetDiscoveredResourceCountsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ConfigserviceGetOrganizationConfigRuleDetailedStatusResult struct {
	Result workflow.Future
}

func (r *ConfigserviceGetOrganizationConfigRuleDetailedStatusResult) Get(ctx workflow.Context) (*configservice.GetOrganizationConfigRuleDetailedStatusOutput, error) {
    var output configservice.GetOrganizationConfigRuleDetailedStatusOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ConfigserviceGetOrganizationConformancePackDetailedStatusResult struct {
	Result workflow.Future
}

func (r *ConfigserviceGetOrganizationConformancePackDetailedStatusResult) Get(ctx workflow.Context) (*configservice.GetOrganizationConformancePackDetailedStatusOutput, error) {
    var output configservice.GetOrganizationConformancePackDetailedStatusOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ConfigserviceGetResourceConfigHistoryResult struct {
	Result workflow.Future
}

func (r *ConfigserviceGetResourceConfigHistoryResult) Get(ctx workflow.Context) (*configservice.GetResourceConfigHistoryOutput, error) {
    var output configservice.GetResourceConfigHistoryOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ConfigserviceListAggregateDiscoveredResourcesResult struct {
	Result workflow.Future
}

func (r *ConfigserviceListAggregateDiscoveredResourcesResult) Get(ctx workflow.Context) (*configservice.ListAggregateDiscoveredResourcesOutput, error) {
    var output configservice.ListAggregateDiscoveredResourcesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ConfigserviceListDiscoveredResourcesResult struct {
	Result workflow.Future
}

func (r *ConfigserviceListDiscoveredResourcesResult) Get(ctx workflow.Context) (*configservice.ListDiscoveredResourcesOutput, error) {
    var output configservice.ListDiscoveredResourcesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ConfigserviceListTagsForResourceResult struct {
	Result workflow.Future
}

func (r *ConfigserviceListTagsForResourceResult) Get(ctx workflow.Context) (*configservice.ListTagsForResourceOutput, error) {
    var output configservice.ListTagsForResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ConfigservicePutAggregationAuthorizationResult struct {
	Result workflow.Future
}

func (r *ConfigservicePutAggregationAuthorizationResult) Get(ctx workflow.Context) (*configservice.PutAggregationAuthorizationOutput, error) {
    var output configservice.PutAggregationAuthorizationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ConfigservicePutConfigRuleResult struct {
	Result workflow.Future
}

func (r *ConfigservicePutConfigRuleResult) Get(ctx workflow.Context) (*configservice.PutConfigRuleOutput, error) {
    var output configservice.PutConfigRuleOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ConfigservicePutConfigurationAggregatorResult struct {
	Result workflow.Future
}

func (r *ConfigservicePutConfigurationAggregatorResult) Get(ctx workflow.Context) (*configservice.PutConfigurationAggregatorOutput, error) {
    var output configservice.PutConfigurationAggregatorOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ConfigservicePutConfigurationRecorderResult struct {
	Result workflow.Future
}

func (r *ConfigservicePutConfigurationRecorderResult) Get(ctx workflow.Context) (*configservice.PutConfigurationRecorderOutput, error) {
    var output configservice.PutConfigurationRecorderOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ConfigservicePutConformancePackResult struct {
	Result workflow.Future
}

func (r *ConfigservicePutConformancePackResult) Get(ctx workflow.Context) (*configservice.PutConformancePackOutput, error) {
    var output configservice.PutConformancePackOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ConfigservicePutDeliveryChannelResult struct {
	Result workflow.Future
}

func (r *ConfigservicePutDeliveryChannelResult) Get(ctx workflow.Context) (*configservice.PutDeliveryChannelOutput, error) {
    var output configservice.PutDeliveryChannelOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ConfigservicePutEvaluationsResult struct {
	Result workflow.Future
}

func (r *ConfigservicePutEvaluationsResult) Get(ctx workflow.Context) (*configservice.PutEvaluationsOutput, error) {
    var output configservice.PutEvaluationsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ConfigservicePutOrganizationConfigRuleResult struct {
	Result workflow.Future
}

func (r *ConfigservicePutOrganizationConfigRuleResult) Get(ctx workflow.Context) (*configservice.PutOrganizationConfigRuleOutput, error) {
    var output configservice.PutOrganizationConfigRuleOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ConfigservicePutOrganizationConformancePackResult struct {
	Result workflow.Future
}

func (r *ConfigservicePutOrganizationConformancePackResult) Get(ctx workflow.Context) (*configservice.PutOrganizationConformancePackOutput, error) {
    var output configservice.PutOrganizationConformancePackOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ConfigservicePutRemediationConfigurationsResult struct {
	Result workflow.Future
}

func (r *ConfigservicePutRemediationConfigurationsResult) Get(ctx workflow.Context) (*configservice.PutRemediationConfigurationsOutput, error) {
    var output configservice.PutRemediationConfigurationsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ConfigservicePutRemediationExceptionsResult struct {
	Result workflow.Future
}

func (r *ConfigservicePutRemediationExceptionsResult) Get(ctx workflow.Context) (*configservice.PutRemediationExceptionsOutput, error) {
    var output configservice.PutRemediationExceptionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ConfigservicePutResourceConfigResult struct {
	Result workflow.Future
}

func (r *ConfigservicePutResourceConfigResult) Get(ctx workflow.Context) (*configservice.PutResourceConfigOutput, error) {
    var output configservice.PutResourceConfigOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ConfigservicePutRetentionConfigurationResult struct {
	Result workflow.Future
}

func (r *ConfigservicePutRetentionConfigurationResult) Get(ctx workflow.Context) (*configservice.PutRetentionConfigurationOutput, error) {
    var output configservice.PutRetentionConfigurationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ConfigserviceSelectAggregateResourceConfigResult struct {
	Result workflow.Future
}

func (r *ConfigserviceSelectAggregateResourceConfigResult) Get(ctx workflow.Context) (*configservice.SelectAggregateResourceConfigOutput, error) {
    var output configservice.SelectAggregateResourceConfigOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ConfigserviceSelectResourceConfigResult struct {
	Result workflow.Future
}

func (r *ConfigserviceSelectResourceConfigResult) Get(ctx workflow.Context) (*configservice.SelectResourceConfigOutput, error) {
    var output configservice.SelectResourceConfigOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ConfigserviceStartConfigRulesEvaluationResult struct {
	Result workflow.Future
}

func (r *ConfigserviceStartConfigRulesEvaluationResult) Get(ctx workflow.Context) (*configservice.StartConfigRulesEvaluationOutput, error) {
    var output configservice.StartConfigRulesEvaluationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ConfigserviceStartConfigurationRecorderResult struct {
	Result workflow.Future
}

func (r *ConfigserviceStartConfigurationRecorderResult) Get(ctx workflow.Context) (*configservice.StartConfigurationRecorderOutput, error) {
    var output configservice.StartConfigurationRecorderOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ConfigserviceStartRemediationExecutionResult struct {
	Result workflow.Future
}

func (r *ConfigserviceStartRemediationExecutionResult) Get(ctx workflow.Context) (*configservice.StartRemediationExecutionOutput, error) {
    var output configservice.StartRemediationExecutionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ConfigserviceStopConfigurationRecorderResult struct {
	Result workflow.Future
}

func (r *ConfigserviceStopConfigurationRecorderResult) Get(ctx workflow.Context) (*configservice.StopConfigurationRecorderOutput, error) {
    var output configservice.StopConfigurationRecorderOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ConfigserviceTagResourceResult struct {
	Result workflow.Future
}

func (r *ConfigserviceTagResourceResult) Get(ctx workflow.Context) (*configservice.TagResourceOutput, error) {
    var output configservice.TagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ConfigserviceUntagResourceResult struct {
	Result workflow.Future
}

func (r *ConfigserviceUntagResourceResult) Get(ctx workflow.Context) (*configservice.UntagResourceOutput, error) {
    var output configservice.UntagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ConfigServiceStub struct {
    activities awsactivities.ConfigServiceActivities
}

func NewConfigServiceStub() ConfigServiceClient {
    return &ConfigServiceStub{}
}

func (a *ConfigServiceStub) BatchGetAggregateResourceConfig(ctx workflow.Context, input *configservice.BatchGetAggregateResourceConfigInput) (*configservice.BatchGetAggregateResourceConfigOutput, error) {
    var output configservice.BatchGetAggregateResourceConfigOutput
    err := workflow.ExecuteActivity(ctx, a.activities.BatchGetAggregateResourceConfig, input).Get(ctx, &output)
    return &output, err
}

func (a *ConfigServiceStub) BatchGetAggregateResourceConfigAsync(ctx workflow.Context, input *configservice.BatchGetAggregateResourceConfigInput) *ConfigserviceBatchGetAggregateResourceConfigResult {
    future := workflow.ExecuteActivity(ctx, a.activities.BatchGetAggregateResourceConfig, input)
    return &ConfigserviceBatchGetAggregateResourceConfigResult{Result: future}
}

func (a *ConfigServiceStub) BatchGetResourceConfig(ctx workflow.Context, input *configservice.BatchGetResourceConfigInput) (*configservice.BatchGetResourceConfigOutput, error) {
    var output configservice.BatchGetResourceConfigOutput
    err := workflow.ExecuteActivity(ctx, a.activities.BatchGetResourceConfig, input).Get(ctx, &output)
    return &output, err
}

func (a *ConfigServiceStub) BatchGetResourceConfigAsync(ctx workflow.Context, input *configservice.BatchGetResourceConfigInput) *ConfigserviceBatchGetResourceConfigResult {
    future := workflow.ExecuteActivity(ctx, a.activities.BatchGetResourceConfig, input)
    return &ConfigserviceBatchGetResourceConfigResult{Result: future}
}

func (a *ConfigServiceStub) DeleteAggregationAuthorization(ctx workflow.Context, input *configservice.DeleteAggregationAuthorizationInput) (*configservice.DeleteAggregationAuthorizationOutput, error) {
    var output configservice.DeleteAggregationAuthorizationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteAggregationAuthorization, input).Get(ctx, &output)
    return &output, err
}

func (a *ConfigServiceStub) DeleteAggregationAuthorizationAsync(ctx workflow.Context, input *configservice.DeleteAggregationAuthorizationInput) *ConfigserviceDeleteAggregationAuthorizationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteAggregationAuthorization, input)
    return &ConfigserviceDeleteAggregationAuthorizationResult{Result: future}
}

func (a *ConfigServiceStub) DeleteConfigRule(ctx workflow.Context, input *configservice.DeleteConfigRuleInput) (*configservice.DeleteConfigRuleOutput, error) {
    var output configservice.DeleteConfigRuleOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteConfigRule, input).Get(ctx, &output)
    return &output, err
}

func (a *ConfigServiceStub) DeleteConfigRuleAsync(ctx workflow.Context, input *configservice.DeleteConfigRuleInput) *ConfigserviceDeleteConfigRuleResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteConfigRule, input)
    return &ConfigserviceDeleteConfigRuleResult{Result: future}
}

func (a *ConfigServiceStub) DeleteConfigurationAggregator(ctx workflow.Context, input *configservice.DeleteConfigurationAggregatorInput) (*configservice.DeleteConfigurationAggregatorOutput, error) {
    var output configservice.DeleteConfigurationAggregatorOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteConfigurationAggregator, input).Get(ctx, &output)
    return &output, err
}

func (a *ConfigServiceStub) DeleteConfigurationAggregatorAsync(ctx workflow.Context, input *configservice.DeleteConfigurationAggregatorInput) *ConfigserviceDeleteConfigurationAggregatorResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteConfigurationAggregator, input)
    return &ConfigserviceDeleteConfigurationAggregatorResult{Result: future}
}

func (a *ConfigServiceStub) DeleteConfigurationRecorder(ctx workflow.Context, input *configservice.DeleteConfigurationRecorderInput) (*configservice.DeleteConfigurationRecorderOutput, error) {
    var output configservice.DeleteConfigurationRecorderOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteConfigurationRecorder, input).Get(ctx, &output)
    return &output, err
}

func (a *ConfigServiceStub) DeleteConfigurationRecorderAsync(ctx workflow.Context, input *configservice.DeleteConfigurationRecorderInput) *ConfigserviceDeleteConfigurationRecorderResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteConfigurationRecorder, input)
    return &ConfigserviceDeleteConfigurationRecorderResult{Result: future}
}

func (a *ConfigServiceStub) DeleteConformancePack(ctx workflow.Context, input *configservice.DeleteConformancePackInput) (*configservice.DeleteConformancePackOutput, error) {
    var output configservice.DeleteConformancePackOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteConformancePack, input).Get(ctx, &output)
    return &output, err
}

func (a *ConfigServiceStub) DeleteConformancePackAsync(ctx workflow.Context, input *configservice.DeleteConformancePackInput) *ConfigserviceDeleteConformancePackResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteConformancePack, input)
    return &ConfigserviceDeleteConformancePackResult{Result: future}
}

func (a *ConfigServiceStub) DeleteDeliveryChannel(ctx workflow.Context, input *configservice.DeleteDeliveryChannelInput) (*configservice.DeleteDeliveryChannelOutput, error) {
    var output configservice.DeleteDeliveryChannelOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteDeliveryChannel, input).Get(ctx, &output)
    return &output, err
}

func (a *ConfigServiceStub) DeleteDeliveryChannelAsync(ctx workflow.Context, input *configservice.DeleteDeliveryChannelInput) *ConfigserviceDeleteDeliveryChannelResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteDeliveryChannel, input)
    return &ConfigserviceDeleteDeliveryChannelResult{Result: future}
}

func (a *ConfigServiceStub) DeleteEvaluationResults(ctx workflow.Context, input *configservice.DeleteEvaluationResultsInput) (*configservice.DeleteEvaluationResultsOutput, error) {
    var output configservice.DeleteEvaluationResultsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteEvaluationResults, input).Get(ctx, &output)
    return &output, err
}

func (a *ConfigServiceStub) DeleteEvaluationResultsAsync(ctx workflow.Context, input *configservice.DeleteEvaluationResultsInput) *ConfigserviceDeleteEvaluationResultsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteEvaluationResults, input)
    return &ConfigserviceDeleteEvaluationResultsResult{Result: future}
}

func (a *ConfigServiceStub) DeleteOrganizationConfigRule(ctx workflow.Context, input *configservice.DeleteOrganizationConfigRuleInput) (*configservice.DeleteOrganizationConfigRuleOutput, error) {
    var output configservice.DeleteOrganizationConfigRuleOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteOrganizationConfigRule, input).Get(ctx, &output)
    return &output, err
}

func (a *ConfigServiceStub) DeleteOrganizationConfigRuleAsync(ctx workflow.Context, input *configservice.DeleteOrganizationConfigRuleInput) *ConfigserviceDeleteOrganizationConfigRuleResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteOrganizationConfigRule, input)
    return &ConfigserviceDeleteOrganizationConfigRuleResult{Result: future}
}

func (a *ConfigServiceStub) DeleteOrganizationConformancePack(ctx workflow.Context, input *configservice.DeleteOrganizationConformancePackInput) (*configservice.DeleteOrganizationConformancePackOutput, error) {
    var output configservice.DeleteOrganizationConformancePackOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteOrganizationConformancePack, input).Get(ctx, &output)
    return &output, err
}

func (a *ConfigServiceStub) DeleteOrganizationConformancePackAsync(ctx workflow.Context, input *configservice.DeleteOrganizationConformancePackInput) *ConfigserviceDeleteOrganizationConformancePackResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteOrganizationConformancePack, input)
    return &ConfigserviceDeleteOrganizationConformancePackResult{Result: future}
}

func (a *ConfigServiceStub) DeleteRemediationConfiguration(ctx workflow.Context, input *configservice.DeleteRemediationConfigurationInput) (*configservice.DeleteRemediationConfigurationOutput, error) {
    var output configservice.DeleteRemediationConfigurationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteRemediationConfiguration, input).Get(ctx, &output)
    return &output, err
}

func (a *ConfigServiceStub) DeleteRemediationConfigurationAsync(ctx workflow.Context, input *configservice.DeleteRemediationConfigurationInput) *ConfigserviceDeleteRemediationConfigurationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteRemediationConfiguration, input)
    return &ConfigserviceDeleteRemediationConfigurationResult{Result: future}
}

func (a *ConfigServiceStub) DeleteRemediationExceptions(ctx workflow.Context, input *configservice.DeleteRemediationExceptionsInput) (*configservice.DeleteRemediationExceptionsOutput, error) {
    var output configservice.DeleteRemediationExceptionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteRemediationExceptions, input).Get(ctx, &output)
    return &output, err
}

func (a *ConfigServiceStub) DeleteRemediationExceptionsAsync(ctx workflow.Context, input *configservice.DeleteRemediationExceptionsInput) *ConfigserviceDeleteRemediationExceptionsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteRemediationExceptions, input)
    return &ConfigserviceDeleteRemediationExceptionsResult{Result: future}
}

func (a *ConfigServiceStub) DeleteResourceConfig(ctx workflow.Context, input *configservice.DeleteResourceConfigInput) (*configservice.DeleteResourceConfigOutput, error) {
    var output configservice.DeleteResourceConfigOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteResourceConfig, input).Get(ctx, &output)
    return &output, err
}

func (a *ConfigServiceStub) DeleteResourceConfigAsync(ctx workflow.Context, input *configservice.DeleteResourceConfigInput) *ConfigserviceDeleteResourceConfigResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteResourceConfig, input)
    return &ConfigserviceDeleteResourceConfigResult{Result: future}
}

func (a *ConfigServiceStub) DeleteRetentionConfiguration(ctx workflow.Context, input *configservice.DeleteRetentionConfigurationInput) (*configservice.DeleteRetentionConfigurationOutput, error) {
    var output configservice.DeleteRetentionConfigurationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteRetentionConfiguration, input).Get(ctx, &output)
    return &output, err
}

func (a *ConfigServiceStub) DeleteRetentionConfigurationAsync(ctx workflow.Context, input *configservice.DeleteRetentionConfigurationInput) *ConfigserviceDeleteRetentionConfigurationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteRetentionConfiguration, input)
    return &ConfigserviceDeleteRetentionConfigurationResult{Result: future}
}

func (a *ConfigServiceStub) DeliverConfigSnapshot(ctx workflow.Context, input *configservice.DeliverConfigSnapshotInput) (*configservice.DeliverConfigSnapshotOutput, error) {
    var output configservice.DeliverConfigSnapshotOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeliverConfigSnapshot, input).Get(ctx, &output)
    return &output, err
}

func (a *ConfigServiceStub) DeliverConfigSnapshotAsync(ctx workflow.Context, input *configservice.DeliverConfigSnapshotInput) *ConfigserviceDeliverConfigSnapshotResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeliverConfigSnapshot, input)
    return &ConfigserviceDeliverConfigSnapshotResult{Result: future}
}

func (a *ConfigServiceStub) DescribeAggregateComplianceByConfigRules(ctx workflow.Context, input *configservice.DescribeAggregateComplianceByConfigRulesInput) (*configservice.DescribeAggregateComplianceByConfigRulesOutput, error) {
    var output configservice.DescribeAggregateComplianceByConfigRulesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeAggregateComplianceByConfigRules, input).Get(ctx, &output)
    return &output, err
}

func (a *ConfigServiceStub) DescribeAggregateComplianceByConfigRulesAsync(ctx workflow.Context, input *configservice.DescribeAggregateComplianceByConfigRulesInput) *ConfigserviceDescribeAggregateComplianceByConfigRulesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeAggregateComplianceByConfigRules, input)
    return &ConfigserviceDescribeAggregateComplianceByConfigRulesResult{Result: future}
}

func (a *ConfigServiceStub) DescribeAggregationAuthorizations(ctx workflow.Context, input *configservice.DescribeAggregationAuthorizationsInput) (*configservice.DescribeAggregationAuthorizationsOutput, error) {
    var output configservice.DescribeAggregationAuthorizationsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeAggregationAuthorizations, input).Get(ctx, &output)
    return &output, err
}

func (a *ConfigServiceStub) DescribeAggregationAuthorizationsAsync(ctx workflow.Context, input *configservice.DescribeAggregationAuthorizationsInput) *ConfigserviceDescribeAggregationAuthorizationsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeAggregationAuthorizations, input)
    return &ConfigserviceDescribeAggregationAuthorizationsResult{Result: future}
}

func (a *ConfigServiceStub) DescribeComplianceByConfigRule(ctx workflow.Context, input *configservice.DescribeComplianceByConfigRuleInput) (*configservice.DescribeComplianceByConfigRuleOutput, error) {
    var output configservice.DescribeComplianceByConfigRuleOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeComplianceByConfigRule, input).Get(ctx, &output)
    return &output, err
}

func (a *ConfigServiceStub) DescribeComplianceByConfigRuleAsync(ctx workflow.Context, input *configservice.DescribeComplianceByConfigRuleInput) *ConfigserviceDescribeComplianceByConfigRuleResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeComplianceByConfigRule, input)
    return &ConfigserviceDescribeComplianceByConfigRuleResult{Result: future}
}

func (a *ConfigServiceStub) DescribeComplianceByResource(ctx workflow.Context, input *configservice.DescribeComplianceByResourceInput) (*configservice.DescribeComplianceByResourceOutput, error) {
    var output configservice.DescribeComplianceByResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeComplianceByResource, input).Get(ctx, &output)
    return &output, err
}

func (a *ConfigServiceStub) DescribeComplianceByResourceAsync(ctx workflow.Context, input *configservice.DescribeComplianceByResourceInput) *ConfigserviceDescribeComplianceByResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeComplianceByResource, input)
    return &ConfigserviceDescribeComplianceByResourceResult{Result: future}
}

func (a *ConfigServiceStub) DescribeConfigRuleEvaluationStatus(ctx workflow.Context, input *configservice.DescribeConfigRuleEvaluationStatusInput) (*configservice.DescribeConfigRuleEvaluationStatusOutput, error) {
    var output configservice.DescribeConfigRuleEvaluationStatusOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeConfigRuleEvaluationStatus, input).Get(ctx, &output)
    return &output, err
}

func (a *ConfigServiceStub) DescribeConfigRuleEvaluationStatusAsync(ctx workflow.Context, input *configservice.DescribeConfigRuleEvaluationStatusInput) *ConfigserviceDescribeConfigRuleEvaluationStatusResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeConfigRuleEvaluationStatus, input)
    return &ConfigserviceDescribeConfigRuleEvaluationStatusResult{Result: future}
}

func (a *ConfigServiceStub) DescribeConfigRules(ctx workflow.Context, input *configservice.DescribeConfigRulesInput) (*configservice.DescribeConfigRulesOutput, error) {
    var output configservice.DescribeConfigRulesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeConfigRules, input).Get(ctx, &output)
    return &output, err
}

func (a *ConfigServiceStub) DescribeConfigRulesAsync(ctx workflow.Context, input *configservice.DescribeConfigRulesInput) *ConfigserviceDescribeConfigRulesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeConfigRules, input)
    return &ConfigserviceDescribeConfigRulesResult{Result: future}
}

func (a *ConfigServiceStub) DescribeConfigurationAggregatorSourcesStatus(ctx workflow.Context, input *configservice.DescribeConfigurationAggregatorSourcesStatusInput) (*configservice.DescribeConfigurationAggregatorSourcesStatusOutput, error) {
    var output configservice.DescribeConfigurationAggregatorSourcesStatusOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeConfigurationAggregatorSourcesStatus, input).Get(ctx, &output)
    return &output, err
}

func (a *ConfigServiceStub) DescribeConfigurationAggregatorSourcesStatusAsync(ctx workflow.Context, input *configservice.DescribeConfigurationAggregatorSourcesStatusInput) *ConfigserviceDescribeConfigurationAggregatorSourcesStatusResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeConfigurationAggregatorSourcesStatus, input)
    return &ConfigserviceDescribeConfigurationAggregatorSourcesStatusResult{Result: future}
}

func (a *ConfigServiceStub) DescribeConfigurationAggregators(ctx workflow.Context, input *configservice.DescribeConfigurationAggregatorsInput) (*configservice.DescribeConfigurationAggregatorsOutput, error) {
    var output configservice.DescribeConfigurationAggregatorsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeConfigurationAggregators, input).Get(ctx, &output)
    return &output, err
}

func (a *ConfigServiceStub) DescribeConfigurationAggregatorsAsync(ctx workflow.Context, input *configservice.DescribeConfigurationAggregatorsInput) *ConfigserviceDescribeConfigurationAggregatorsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeConfigurationAggregators, input)
    return &ConfigserviceDescribeConfigurationAggregatorsResult{Result: future}
}

func (a *ConfigServiceStub) DescribeConfigurationRecorderStatus(ctx workflow.Context, input *configservice.DescribeConfigurationRecorderStatusInput) (*configservice.DescribeConfigurationRecorderStatusOutput, error) {
    var output configservice.DescribeConfigurationRecorderStatusOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeConfigurationRecorderStatus, input).Get(ctx, &output)
    return &output, err
}

func (a *ConfigServiceStub) DescribeConfigurationRecorderStatusAsync(ctx workflow.Context, input *configservice.DescribeConfigurationRecorderStatusInput) *ConfigserviceDescribeConfigurationRecorderStatusResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeConfigurationRecorderStatus, input)
    return &ConfigserviceDescribeConfigurationRecorderStatusResult{Result: future}
}

func (a *ConfigServiceStub) DescribeConfigurationRecorders(ctx workflow.Context, input *configservice.DescribeConfigurationRecordersInput) (*configservice.DescribeConfigurationRecordersOutput, error) {
    var output configservice.DescribeConfigurationRecordersOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeConfigurationRecorders, input).Get(ctx, &output)
    return &output, err
}

func (a *ConfigServiceStub) DescribeConfigurationRecordersAsync(ctx workflow.Context, input *configservice.DescribeConfigurationRecordersInput) *ConfigserviceDescribeConfigurationRecordersResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeConfigurationRecorders, input)
    return &ConfigserviceDescribeConfigurationRecordersResult{Result: future}
}

func (a *ConfigServiceStub) DescribeConformancePackCompliance(ctx workflow.Context, input *configservice.DescribeConformancePackComplianceInput) (*configservice.DescribeConformancePackComplianceOutput, error) {
    var output configservice.DescribeConformancePackComplianceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeConformancePackCompliance, input).Get(ctx, &output)
    return &output, err
}

func (a *ConfigServiceStub) DescribeConformancePackComplianceAsync(ctx workflow.Context, input *configservice.DescribeConformancePackComplianceInput) *ConfigserviceDescribeConformancePackComplianceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeConformancePackCompliance, input)
    return &ConfigserviceDescribeConformancePackComplianceResult{Result: future}
}

func (a *ConfigServiceStub) DescribeConformancePackStatus(ctx workflow.Context, input *configservice.DescribeConformancePackStatusInput) (*configservice.DescribeConformancePackStatusOutput, error) {
    var output configservice.DescribeConformancePackStatusOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeConformancePackStatus, input).Get(ctx, &output)
    return &output, err
}

func (a *ConfigServiceStub) DescribeConformancePackStatusAsync(ctx workflow.Context, input *configservice.DescribeConformancePackStatusInput) *ConfigserviceDescribeConformancePackStatusResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeConformancePackStatus, input)
    return &ConfigserviceDescribeConformancePackStatusResult{Result: future}
}

func (a *ConfigServiceStub) DescribeConformancePacks(ctx workflow.Context, input *configservice.DescribeConformancePacksInput) (*configservice.DescribeConformancePacksOutput, error) {
    var output configservice.DescribeConformancePacksOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeConformancePacks, input).Get(ctx, &output)
    return &output, err
}

func (a *ConfigServiceStub) DescribeConformancePacksAsync(ctx workflow.Context, input *configservice.DescribeConformancePacksInput) *ConfigserviceDescribeConformancePacksResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeConformancePacks, input)
    return &ConfigserviceDescribeConformancePacksResult{Result: future}
}

func (a *ConfigServiceStub) DescribeDeliveryChannelStatus(ctx workflow.Context, input *configservice.DescribeDeliveryChannelStatusInput) (*configservice.DescribeDeliveryChannelStatusOutput, error) {
    var output configservice.DescribeDeliveryChannelStatusOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeDeliveryChannelStatus, input).Get(ctx, &output)
    return &output, err
}

func (a *ConfigServiceStub) DescribeDeliveryChannelStatusAsync(ctx workflow.Context, input *configservice.DescribeDeliveryChannelStatusInput) *ConfigserviceDescribeDeliveryChannelStatusResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeDeliveryChannelStatus, input)
    return &ConfigserviceDescribeDeliveryChannelStatusResult{Result: future}
}

func (a *ConfigServiceStub) DescribeDeliveryChannels(ctx workflow.Context, input *configservice.DescribeDeliveryChannelsInput) (*configservice.DescribeDeliveryChannelsOutput, error) {
    var output configservice.DescribeDeliveryChannelsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeDeliveryChannels, input).Get(ctx, &output)
    return &output, err
}

func (a *ConfigServiceStub) DescribeDeliveryChannelsAsync(ctx workflow.Context, input *configservice.DescribeDeliveryChannelsInput) *ConfigserviceDescribeDeliveryChannelsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeDeliveryChannels, input)
    return &ConfigserviceDescribeDeliveryChannelsResult{Result: future}
}

func (a *ConfigServiceStub) DescribeOrganizationConfigRuleStatuses(ctx workflow.Context, input *configservice.DescribeOrganizationConfigRuleStatusesInput) (*configservice.DescribeOrganizationConfigRuleStatusesOutput, error) {
    var output configservice.DescribeOrganizationConfigRuleStatusesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeOrganizationConfigRuleStatuses, input).Get(ctx, &output)
    return &output, err
}

func (a *ConfigServiceStub) DescribeOrganizationConfigRuleStatusesAsync(ctx workflow.Context, input *configservice.DescribeOrganizationConfigRuleStatusesInput) *ConfigserviceDescribeOrganizationConfigRuleStatusesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeOrganizationConfigRuleStatuses, input)
    return &ConfigserviceDescribeOrganizationConfigRuleStatusesResult{Result: future}
}

func (a *ConfigServiceStub) DescribeOrganizationConfigRules(ctx workflow.Context, input *configservice.DescribeOrganizationConfigRulesInput) (*configservice.DescribeOrganizationConfigRulesOutput, error) {
    var output configservice.DescribeOrganizationConfigRulesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeOrganizationConfigRules, input).Get(ctx, &output)
    return &output, err
}

func (a *ConfigServiceStub) DescribeOrganizationConfigRulesAsync(ctx workflow.Context, input *configservice.DescribeOrganizationConfigRulesInput) *ConfigserviceDescribeOrganizationConfigRulesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeOrganizationConfigRules, input)
    return &ConfigserviceDescribeOrganizationConfigRulesResult{Result: future}
}

func (a *ConfigServiceStub) DescribeOrganizationConformancePackStatuses(ctx workflow.Context, input *configservice.DescribeOrganizationConformancePackStatusesInput) (*configservice.DescribeOrganizationConformancePackStatusesOutput, error) {
    var output configservice.DescribeOrganizationConformancePackStatusesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeOrganizationConformancePackStatuses, input).Get(ctx, &output)
    return &output, err
}

func (a *ConfigServiceStub) DescribeOrganizationConformancePackStatusesAsync(ctx workflow.Context, input *configservice.DescribeOrganizationConformancePackStatusesInput) *ConfigserviceDescribeOrganizationConformancePackStatusesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeOrganizationConformancePackStatuses, input)
    return &ConfigserviceDescribeOrganizationConformancePackStatusesResult{Result: future}
}

func (a *ConfigServiceStub) DescribeOrganizationConformancePacks(ctx workflow.Context, input *configservice.DescribeOrganizationConformancePacksInput) (*configservice.DescribeOrganizationConformancePacksOutput, error) {
    var output configservice.DescribeOrganizationConformancePacksOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeOrganizationConformancePacks, input).Get(ctx, &output)
    return &output, err
}

func (a *ConfigServiceStub) DescribeOrganizationConformancePacksAsync(ctx workflow.Context, input *configservice.DescribeOrganizationConformancePacksInput) *ConfigserviceDescribeOrganizationConformancePacksResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeOrganizationConformancePacks, input)
    return &ConfigserviceDescribeOrganizationConformancePacksResult{Result: future}
}

func (a *ConfigServiceStub) DescribePendingAggregationRequests(ctx workflow.Context, input *configservice.DescribePendingAggregationRequestsInput) (*configservice.DescribePendingAggregationRequestsOutput, error) {
    var output configservice.DescribePendingAggregationRequestsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribePendingAggregationRequests, input).Get(ctx, &output)
    return &output, err
}

func (a *ConfigServiceStub) DescribePendingAggregationRequestsAsync(ctx workflow.Context, input *configservice.DescribePendingAggregationRequestsInput) *ConfigserviceDescribePendingAggregationRequestsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribePendingAggregationRequests, input)
    return &ConfigserviceDescribePendingAggregationRequestsResult{Result: future}
}

func (a *ConfigServiceStub) DescribeRemediationConfigurations(ctx workflow.Context, input *configservice.DescribeRemediationConfigurationsInput) (*configservice.DescribeRemediationConfigurationsOutput, error) {
    var output configservice.DescribeRemediationConfigurationsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeRemediationConfigurations, input).Get(ctx, &output)
    return &output, err
}

func (a *ConfigServiceStub) DescribeRemediationConfigurationsAsync(ctx workflow.Context, input *configservice.DescribeRemediationConfigurationsInput) *ConfigserviceDescribeRemediationConfigurationsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeRemediationConfigurations, input)
    return &ConfigserviceDescribeRemediationConfigurationsResult{Result: future}
}

func (a *ConfigServiceStub) DescribeRemediationExceptions(ctx workflow.Context, input *configservice.DescribeRemediationExceptionsInput) (*configservice.DescribeRemediationExceptionsOutput, error) {
    var output configservice.DescribeRemediationExceptionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeRemediationExceptions, input).Get(ctx, &output)
    return &output, err
}

func (a *ConfigServiceStub) DescribeRemediationExceptionsAsync(ctx workflow.Context, input *configservice.DescribeRemediationExceptionsInput) *ConfigserviceDescribeRemediationExceptionsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeRemediationExceptions, input)
    return &ConfigserviceDescribeRemediationExceptionsResult{Result: future}
}

func (a *ConfigServiceStub) DescribeRemediationExecutionStatus(ctx workflow.Context, input *configservice.DescribeRemediationExecutionStatusInput) (*configservice.DescribeRemediationExecutionStatusOutput, error) {
    var output configservice.DescribeRemediationExecutionStatusOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeRemediationExecutionStatus, input).Get(ctx, &output)
    return &output, err
}

func (a *ConfigServiceStub) DescribeRemediationExecutionStatusAsync(ctx workflow.Context, input *configservice.DescribeRemediationExecutionStatusInput) *ConfigserviceDescribeRemediationExecutionStatusResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeRemediationExecutionStatus, input)
    return &ConfigserviceDescribeRemediationExecutionStatusResult{Result: future}
}

func (a *ConfigServiceStub) DescribeRetentionConfigurations(ctx workflow.Context, input *configservice.DescribeRetentionConfigurationsInput) (*configservice.DescribeRetentionConfigurationsOutput, error) {
    var output configservice.DescribeRetentionConfigurationsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeRetentionConfigurations, input).Get(ctx, &output)
    return &output, err
}

func (a *ConfigServiceStub) DescribeRetentionConfigurationsAsync(ctx workflow.Context, input *configservice.DescribeRetentionConfigurationsInput) *ConfigserviceDescribeRetentionConfigurationsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeRetentionConfigurations, input)
    return &ConfigserviceDescribeRetentionConfigurationsResult{Result: future}
}

func (a *ConfigServiceStub) GetAggregateComplianceDetailsByConfigRule(ctx workflow.Context, input *configservice.GetAggregateComplianceDetailsByConfigRuleInput) (*configservice.GetAggregateComplianceDetailsByConfigRuleOutput, error) {
    var output configservice.GetAggregateComplianceDetailsByConfigRuleOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetAggregateComplianceDetailsByConfigRule, input).Get(ctx, &output)
    return &output, err
}

func (a *ConfigServiceStub) GetAggregateComplianceDetailsByConfigRuleAsync(ctx workflow.Context, input *configservice.GetAggregateComplianceDetailsByConfigRuleInput) *ConfigserviceGetAggregateComplianceDetailsByConfigRuleResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetAggregateComplianceDetailsByConfigRule, input)
    return &ConfigserviceGetAggregateComplianceDetailsByConfigRuleResult{Result: future}
}

func (a *ConfigServiceStub) GetAggregateConfigRuleComplianceSummary(ctx workflow.Context, input *configservice.GetAggregateConfigRuleComplianceSummaryInput) (*configservice.GetAggregateConfigRuleComplianceSummaryOutput, error) {
    var output configservice.GetAggregateConfigRuleComplianceSummaryOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetAggregateConfigRuleComplianceSummary, input).Get(ctx, &output)
    return &output, err
}

func (a *ConfigServiceStub) GetAggregateConfigRuleComplianceSummaryAsync(ctx workflow.Context, input *configservice.GetAggregateConfigRuleComplianceSummaryInput) *ConfigserviceGetAggregateConfigRuleComplianceSummaryResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetAggregateConfigRuleComplianceSummary, input)
    return &ConfigserviceGetAggregateConfigRuleComplianceSummaryResult{Result: future}
}

func (a *ConfigServiceStub) GetAggregateDiscoveredResourceCounts(ctx workflow.Context, input *configservice.GetAggregateDiscoveredResourceCountsInput) (*configservice.GetAggregateDiscoveredResourceCountsOutput, error) {
    var output configservice.GetAggregateDiscoveredResourceCountsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetAggregateDiscoveredResourceCounts, input).Get(ctx, &output)
    return &output, err
}

func (a *ConfigServiceStub) GetAggregateDiscoveredResourceCountsAsync(ctx workflow.Context, input *configservice.GetAggregateDiscoveredResourceCountsInput) *ConfigserviceGetAggregateDiscoveredResourceCountsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetAggregateDiscoveredResourceCounts, input)
    return &ConfigserviceGetAggregateDiscoveredResourceCountsResult{Result: future}
}

func (a *ConfigServiceStub) GetAggregateResourceConfig(ctx workflow.Context, input *configservice.GetAggregateResourceConfigInput) (*configservice.GetAggregateResourceConfigOutput, error) {
    var output configservice.GetAggregateResourceConfigOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetAggregateResourceConfig, input).Get(ctx, &output)
    return &output, err
}

func (a *ConfigServiceStub) GetAggregateResourceConfigAsync(ctx workflow.Context, input *configservice.GetAggregateResourceConfigInput) *ConfigserviceGetAggregateResourceConfigResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetAggregateResourceConfig, input)
    return &ConfigserviceGetAggregateResourceConfigResult{Result: future}
}

func (a *ConfigServiceStub) GetComplianceDetailsByConfigRule(ctx workflow.Context, input *configservice.GetComplianceDetailsByConfigRuleInput) (*configservice.GetComplianceDetailsByConfigRuleOutput, error) {
    var output configservice.GetComplianceDetailsByConfigRuleOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetComplianceDetailsByConfigRule, input).Get(ctx, &output)
    return &output, err
}

func (a *ConfigServiceStub) GetComplianceDetailsByConfigRuleAsync(ctx workflow.Context, input *configservice.GetComplianceDetailsByConfigRuleInput) *ConfigserviceGetComplianceDetailsByConfigRuleResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetComplianceDetailsByConfigRule, input)
    return &ConfigserviceGetComplianceDetailsByConfigRuleResult{Result: future}
}

func (a *ConfigServiceStub) GetComplianceDetailsByResource(ctx workflow.Context, input *configservice.GetComplianceDetailsByResourceInput) (*configservice.GetComplianceDetailsByResourceOutput, error) {
    var output configservice.GetComplianceDetailsByResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetComplianceDetailsByResource, input).Get(ctx, &output)
    return &output, err
}

func (a *ConfigServiceStub) GetComplianceDetailsByResourceAsync(ctx workflow.Context, input *configservice.GetComplianceDetailsByResourceInput) *ConfigserviceGetComplianceDetailsByResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetComplianceDetailsByResource, input)
    return &ConfigserviceGetComplianceDetailsByResourceResult{Result: future}
}

func (a *ConfigServiceStub) GetComplianceSummaryByConfigRule(ctx workflow.Context, input *configservice.GetComplianceSummaryByConfigRuleInput) (*configservice.GetComplianceSummaryByConfigRuleOutput, error) {
    var output configservice.GetComplianceSummaryByConfigRuleOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetComplianceSummaryByConfigRule, input).Get(ctx, &output)
    return &output, err
}

func (a *ConfigServiceStub) GetComplianceSummaryByConfigRuleAsync(ctx workflow.Context, input *configservice.GetComplianceSummaryByConfigRuleInput) *ConfigserviceGetComplianceSummaryByConfigRuleResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetComplianceSummaryByConfigRule, input)
    return &ConfigserviceGetComplianceSummaryByConfigRuleResult{Result: future}
}

func (a *ConfigServiceStub) GetComplianceSummaryByResourceType(ctx workflow.Context, input *configservice.GetComplianceSummaryByResourceTypeInput) (*configservice.GetComplianceSummaryByResourceTypeOutput, error) {
    var output configservice.GetComplianceSummaryByResourceTypeOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetComplianceSummaryByResourceType, input).Get(ctx, &output)
    return &output, err
}

func (a *ConfigServiceStub) GetComplianceSummaryByResourceTypeAsync(ctx workflow.Context, input *configservice.GetComplianceSummaryByResourceTypeInput) *ConfigserviceGetComplianceSummaryByResourceTypeResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetComplianceSummaryByResourceType, input)
    return &ConfigserviceGetComplianceSummaryByResourceTypeResult{Result: future}
}

func (a *ConfigServiceStub) GetConformancePackComplianceDetails(ctx workflow.Context, input *configservice.GetConformancePackComplianceDetailsInput) (*configservice.GetConformancePackComplianceDetailsOutput, error) {
    var output configservice.GetConformancePackComplianceDetailsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetConformancePackComplianceDetails, input).Get(ctx, &output)
    return &output, err
}

func (a *ConfigServiceStub) GetConformancePackComplianceDetailsAsync(ctx workflow.Context, input *configservice.GetConformancePackComplianceDetailsInput) *ConfigserviceGetConformancePackComplianceDetailsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetConformancePackComplianceDetails, input)
    return &ConfigserviceGetConformancePackComplianceDetailsResult{Result: future}
}

func (a *ConfigServiceStub) GetConformancePackComplianceSummary(ctx workflow.Context, input *configservice.GetConformancePackComplianceSummaryInput) (*configservice.GetConformancePackComplianceSummaryOutput, error) {
    var output configservice.GetConformancePackComplianceSummaryOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetConformancePackComplianceSummary, input).Get(ctx, &output)
    return &output, err
}

func (a *ConfigServiceStub) GetConformancePackComplianceSummaryAsync(ctx workflow.Context, input *configservice.GetConformancePackComplianceSummaryInput) *ConfigserviceGetConformancePackComplianceSummaryResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetConformancePackComplianceSummary, input)
    return &ConfigserviceGetConformancePackComplianceSummaryResult{Result: future}
}

func (a *ConfigServiceStub) GetDiscoveredResourceCounts(ctx workflow.Context, input *configservice.GetDiscoveredResourceCountsInput) (*configservice.GetDiscoveredResourceCountsOutput, error) {
    var output configservice.GetDiscoveredResourceCountsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetDiscoveredResourceCounts, input).Get(ctx, &output)
    return &output, err
}

func (a *ConfigServiceStub) GetDiscoveredResourceCountsAsync(ctx workflow.Context, input *configservice.GetDiscoveredResourceCountsInput) *ConfigserviceGetDiscoveredResourceCountsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetDiscoveredResourceCounts, input)
    return &ConfigserviceGetDiscoveredResourceCountsResult{Result: future}
}

func (a *ConfigServiceStub) GetOrganizationConfigRuleDetailedStatus(ctx workflow.Context, input *configservice.GetOrganizationConfigRuleDetailedStatusInput) (*configservice.GetOrganizationConfigRuleDetailedStatusOutput, error) {
    var output configservice.GetOrganizationConfigRuleDetailedStatusOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetOrganizationConfigRuleDetailedStatus, input).Get(ctx, &output)
    return &output, err
}

func (a *ConfigServiceStub) GetOrganizationConfigRuleDetailedStatusAsync(ctx workflow.Context, input *configservice.GetOrganizationConfigRuleDetailedStatusInput) *ConfigserviceGetOrganizationConfigRuleDetailedStatusResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetOrganizationConfigRuleDetailedStatus, input)
    return &ConfigserviceGetOrganizationConfigRuleDetailedStatusResult{Result: future}
}

func (a *ConfigServiceStub) GetOrganizationConformancePackDetailedStatus(ctx workflow.Context, input *configservice.GetOrganizationConformancePackDetailedStatusInput) (*configservice.GetOrganizationConformancePackDetailedStatusOutput, error) {
    var output configservice.GetOrganizationConformancePackDetailedStatusOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetOrganizationConformancePackDetailedStatus, input).Get(ctx, &output)
    return &output, err
}

func (a *ConfigServiceStub) GetOrganizationConformancePackDetailedStatusAsync(ctx workflow.Context, input *configservice.GetOrganizationConformancePackDetailedStatusInput) *ConfigserviceGetOrganizationConformancePackDetailedStatusResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetOrganizationConformancePackDetailedStatus, input)
    return &ConfigserviceGetOrganizationConformancePackDetailedStatusResult{Result: future}
}

func (a *ConfigServiceStub) GetResourceConfigHistory(ctx workflow.Context, input *configservice.GetResourceConfigHistoryInput) (*configservice.GetResourceConfigHistoryOutput, error) {
    var output configservice.GetResourceConfigHistoryOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetResourceConfigHistory, input).Get(ctx, &output)
    return &output, err
}

func (a *ConfigServiceStub) GetResourceConfigHistoryAsync(ctx workflow.Context, input *configservice.GetResourceConfigHistoryInput) *ConfigserviceGetResourceConfigHistoryResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetResourceConfigHistory, input)
    return &ConfigserviceGetResourceConfigHistoryResult{Result: future}
}

func (a *ConfigServiceStub) ListAggregateDiscoveredResources(ctx workflow.Context, input *configservice.ListAggregateDiscoveredResourcesInput) (*configservice.ListAggregateDiscoveredResourcesOutput, error) {
    var output configservice.ListAggregateDiscoveredResourcesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListAggregateDiscoveredResources, input).Get(ctx, &output)
    return &output, err
}

func (a *ConfigServiceStub) ListAggregateDiscoveredResourcesAsync(ctx workflow.Context, input *configservice.ListAggregateDiscoveredResourcesInput) *ConfigserviceListAggregateDiscoveredResourcesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListAggregateDiscoveredResources, input)
    return &ConfigserviceListAggregateDiscoveredResourcesResult{Result: future}
}

func (a *ConfigServiceStub) ListDiscoveredResources(ctx workflow.Context, input *configservice.ListDiscoveredResourcesInput) (*configservice.ListDiscoveredResourcesOutput, error) {
    var output configservice.ListDiscoveredResourcesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListDiscoveredResources, input).Get(ctx, &output)
    return &output, err
}

func (a *ConfigServiceStub) ListDiscoveredResourcesAsync(ctx workflow.Context, input *configservice.ListDiscoveredResourcesInput) *ConfigserviceListDiscoveredResourcesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListDiscoveredResources, input)
    return &ConfigserviceListDiscoveredResourcesResult{Result: future}
}

func (a *ConfigServiceStub) ListTagsForResource(ctx workflow.Context, input *configservice.ListTagsForResourceInput) (*configservice.ListTagsForResourceOutput, error) {
    var output configservice.ListTagsForResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListTagsForResource, input).Get(ctx, &output)
    return &output, err
}

func (a *ConfigServiceStub) ListTagsForResourceAsync(ctx workflow.Context, input *configservice.ListTagsForResourceInput) *ConfigserviceListTagsForResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListTagsForResource, input)
    return &ConfigserviceListTagsForResourceResult{Result: future}
}

func (a *ConfigServiceStub) PutAggregationAuthorization(ctx workflow.Context, input *configservice.PutAggregationAuthorizationInput) (*configservice.PutAggregationAuthorizationOutput, error) {
    var output configservice.PutAggregationAuthorizationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutAggregationAuthorization, input).Get(ctx, &output)
    return &output, err
}

func (a *ConfigServiceStub) PutAggregationAuthorizationAsync(ctx workflow.Context, input *configservice.PutAggregationAuthorizationInput) *ConfigservicePutAggregationAuthorizationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.PutAggregationAuthorization, input)
    return &ConfigservicePutAggregationAuthorizationResult{Result: future}
}

func (a *ConfigServiceStub) PutConfigRule(ctx workflow.Context, input *configservice.PutConfigRuleInput) (*configservice.PutConfigRuleOutput, error) {
    var output configservice.PutConfigRuleOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutConfigRule, input).Get(ctx, &output)
    return &output, err
}

func (a *ConfigServiceStub) PutConfigRuleAsync(ctx workflow.Context, input *configservice.PutConfigRuleInput) *ConfigservicePutConfigRuleResult {
    future := workflow.ExecuteActivity(ctx, a.activities.PutConfigRule, input)
    return &ConfigservicePutConfigRuleResult{Result: future}
}

func (a *ConfigServiceStub) PutConfigurationAggregator(ctx workflow.Context, input *configservice.PutConfigurationAggregatorInput) (*configservice.PutConfigurationAggregatorOutput, error) {
    var output configservice.PutConfigurationAggregatorOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutConfigurationAggregator, input).Get(ctx, &output)
    return &output, err
}

func (a *ConfigServiceStub) PutConfigurationAggregatorAsync(ctx workflow.Context, input *configservice.PutConfigurationAggregatorInput) *ConfigservicePutConfigurationAggregatorResult {
    future := workflow.ExecuteActivity(ctx, a.activities.PutConfigurationAggregator, input)
    return &ConfigservicePutConfigurationAggregatorResult{Result: future}
}

func (a *ConfigServiceStub) PutConfigurationRecorder(ctx workflow.Context, input *configservice.PutConfigurationRecorderInput) (*configservice.PutConfigurationRecorderOutput, error) {
    var output configservice.PutConfigurationRecorderOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutConfigurationRecorder, input).Get(ctx, &output)
    return &output, err
}

func (a *ConfigServiceStub) PutConfigurationRecorderAsync(ctx workflow.Context, input *configservice.PutConfigurationRecorderInput) *ConfigservicePutConfigurationRecorderResult {
    future := workflow.ExecuteActivity(ctx, a.activities.PutConfigurationRecorder, input)
    return &ConfigservicePutConfigurationRecorderResult{Result: future}
}

func (a *ConfigServiceStub) PutConformancePack(ctx workflow.Context, input *configservice.PutConformancePackInput) (*configservice.PutConformancePackOutput, error) {
    var output configservice.PutConformancePackOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutConformancePack, input).Get(ctx, &output)
    return &output, err
}

func (a *ConfigServiceStub) PutConformancePackAsync(ctx workflow.Context, input *configservice.PutConformancePackInput) *ConfigservicePutConformancePackResult {
    future := workflow.ExecuteActivity(ctx, a.activities.PutConformancePack, input)
    return &ConfigservicePutConformancePackResult{Result: future}
}

func (a *ConfigServiceStub) PutDeliveryChannel(ctx workflow.Context, input *configservice.PutDeliveryChannelInput) (*configservice.PutDeliveryChannelOutput, error) {
    var output configservice.PutDeliveryChannelOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutDeliveryChannel, input).Get(ctx, &output)
    return &output, err
}

func (a *ConfigServiceStub) PutDeliveryChannelAsync(ctx workflow.Context, input *configservice.PutDeliveryChannelInput) *ConfigservicePutDeliveryChannelResult {
    future := workflow.ExecuteActivity(ctx, a.activities.PutDeliveryChannel, input)
    return &ConfigservicePutDeliveryChannelResult{Result: future}
}

func (a *ConfigServiceStub) PutEvaluations(ctx workflow.Context, input *configservice.PutEvaluationsInput) (*configservice.PutEvaluationsOutput, error) {
    var output configservice.PutEvaluationsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutEvaluations, input).Get(ctx, &output)
    return &output, err
}

func (a *ConfigServiceStub) PutEvaluationsAsync(ctx workflow.Context, input *configservice.PutEvaluationsInput) *ConfigservicePutEvaluationsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.PutEvaluations, input)
    return &ConfigservicePutEvaluationsResult{Result: future}
}

func (a *ConfigServiceStub) PutOrganizationConfigRule(ctx workflow.Context, input *configservice.PutOrganizationConfigRuleInput) (*configservice.PutOrganizationConfigRuleOutput, error) {
    var output configservice.PutOrganizationConfigRuleOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutOrganizationConfigRule, input).Get(ctx, &output)
    return &output, err
}

func (a *ConfigServiceStub) PutOrganizationConfigRuleAsync(ctx workflow.Context, input *configservice.PutOrganizationConfigRuleInput) *ConfigservicePutOrganizationConfigRuleResult {
    future := workflow.ExecuteActivity(ctx, a.activities.PutOrganizationConfigRule, input)
    return &ConfigservicePutOrganizationConfigRuleResult{Result: future}
}

func (a *ConfigServiceStub) PutOrganizationConformancePack(ctx workflow.Context, input *configservice.PutOrganizationConformancePackInput) (*configservice.PutOrganizationConformancePackOutput, error) {
    var output configservice.PutOrganizationConformancePackOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutOrganizationConformancePack, input).Get(ctx, &output)
    return &output, err
}

func (a *ConfigServiceStub) PutOrganizationConformancePackAsync(ctx workflow.Context, input *configservice.PutOrganizationConformancePackInput) *ConfigservicePutOrganizationConformancePackResult {
    future := workflow.ExecuteActivity(ctx, a.activities.PutOrganizationConformancePack, input)
    return &ConfigservicePutOrganizationConformancePackResult{Result: future}
}

func (a *ConfigServiceStub) PutRemediationConfigurations(ctx workflow.Context, input *configservice.PutRemediationConfigurationsInput) (*configservice.PutRemediationConfigurationsOutput, error) {
    var output configservice.PutRemediationConfigurationsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutRemediationConfigurations, input).Get(ctx, &output)
    return &output, err
}

func (a *ConfigServiceStub) PutRemediationConfigurationsAsync(ctx workflow.Context, input *configservice.PutRemediationConfigurationsInput) *ConfigservicePutRemediationConfigurationsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.PutRemediationConfigurations, input)
    return &ConfigservicePutRemediationConfigurationsResult{Result: future}
}

func (a *ConfigServiceStub) PutRemediationExceptions(ctx workflow.Context, input *configservice.PutRemediationExceptionsInput) (*configservice.PutRemediationExceptionsOutput, error) {
    var output configservice.PutRemediationExceptionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutRemediationExceptions, input).Get(ctx, &output)
    return &output, err
}

func (a *ConfigServiceStub) PutRemediationExceptionsAsync(ctx workflow.Context, input *configservice.PutRemediationExceptionsInput) *ConfigservicePutRemediationExceptionsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.PutRemediationExceptions, input)
    return &ConfigservicePutRemediationExceptionsResult{Result: future}
}

func (a *ConfigServiceStub) PutResourceConfig(ctx workflow.Context, input *configservice.PutResourceConfigInput) (*configservice.PutResourceConfigOutput, error) {
    var output configservice.PutResourceConfigOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutResourceConfig, input).Get(ctx, &output)
    return &output, err
}

func (a *ConfigServiceStub) PutResourceConfigAsync(ctx workflow.Context, input *configservice.PutResourceConfigInput) *ConfigservicePutResourceConfigResult {
    future := workflow.ExecuteActivity(ctx, a.activities.PutResourceConfig, input)
    return &ConfigservicePutResourceConfigResult{Result: future}
}

func (a *ConfigServiceStub) PutRetentionConfiguration(ctx workflow.Context, input *configservice.PutRetentionConfigurationInput) (*configservice.PutRetentionConfigurationOutput, error) {
    var output configservice.PutRetentionConfigurationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutRetentionConfiguration, input).Get(ctx, &output)
    return &output, err
}

func (a *ConfigServiceStub) PutRetentionConfigurationAsync(ctx workflow.Context, input *configservice.PutRetentionConfigurationInput) *ConfigservicePutRetentionConfigurationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.PutRetentionConfiguration, input)
    return &ConfigservicePutRetentionConfigurationResult{Result: future}
}

func (a *ConfigServiceStub) SelectAggregateResourceConfig(ctx workflow.Context, input *configservice.SelectAggregateResourceConfigInput) (*configservice.SelectAggregateResourceConfigOutput, error) {
    var output configservice.SelectAggregateResourceConfigOutput
    err := workflow.ExecuteActivity(ctx, a.activities.SelectAggregateResourceConfig, input).Get(ctx, &output)
    return &output, err
}

func (a *ConfigServiceStub) SelectAggregateResourceConfigAsync(ctx workflow.Context, input *configservice.SelectAggregateResourceConfigInput) *ConfigserviceSelectAggregateResourceConfigResult {
    future := workflow.ExecuteActivity(ctx, a.activities.SelectAggregateResourceConfig, input)
    return &ConfigserviceSelectAggregateResourceConfigResult{Result: future}
}

func (a *ConfigServiceStub) SelectResourceConfig(ctx workflow.Context, input *configservice.SelectResourceConfigInput) (*configservice.SelectResourceConfigOutput, error) {
    var output configservice.SelectResourceConfigOutput
    err := workflow.ExecuteActivity(ctx, a.activities.SelectResourceConfig, input).Get(ctx, &output)
    return &output, err
}

func (a *ConfigServiceStub) SelectResourceConfigAsync(ctx workflow.Context, input *configservice.SelectResourceConfigInput) *ConfigserviceSelectResourceConfigResult {
    future := workflow.ExecuteActivity(ctx, a.activities.SelectResourceConfig, input)
    return &ConfigserviceSelectResourceConfigResult{Result: future}
}

func (a *ConfigServiceStub) StartConfigRulesEvaluation(ctx workflow.Context, input *configservice.StartConfigRulesEvaluationInput) (*configservice.StartConfigRulesEvaluationOutput, error) {
    var output configservice.StartConfigRulesEvaluationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StartConfigRulesEvaluation, input).Get(ctx, &output)
    return &output, err
}

func (a *ConfigServiceStub) StartConfigRulesEvaluationAsync(ctx workflow.Context, input *configservice.StartConfigRulesEvaluationInput) *ConfigserviceStartConfigRulesEvaluationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.StartConfigRulesEvaluation, input)
    return &ConfigserviceStartConfigRulesEvaluationResult{Result: future}
}

func (a *ConfigServiceStub) StartConfigurationRecorder(ctx workflow.Context, input *configservice.StartConfigurationRecorderInput) (*configservice.StartConfigurationRecorderOutput, error) {
    var output configservice.StartConfigurationRecorderOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StartConfigurationRecorder, input).Get(ctx, &output)
    return &output, err
}

func (a *ConfigServiceStub) StartConfigurationRecorderAsync(ctx workflow.Context, input *configservice.StartConfigurationRecorderInput) *ConfigserviceStartConfigurationRecorderResult {
    future := workflow.ExecuteActivity(ctx, a.activities.StartConfigurationRecorder, input)
    return &ConfigserviceStartConfigurationRecorderResult{Result: future}
}

func (a *ConfigServiceStub) StartRemediationExecution(ctx workflow.Context, input *configservice.StartRemediationExecutionInput) (*configservice.StartRemediationExecutionOutput, error) {
    var output configservice.StartRemediationExecutionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StartRemediationExecution, input).Get(ctx, &output)
    return &output, err
}

func (a *ConfigServiceStub) StartRemediationExecutionAsync(ctx workflow.Context, input *configservice.StartRemediationExecutionInput) *ConfigserviceStartRemediationExecutionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.StartRemediationExecution, input)
    return &ConfigserviceStartRemediationExecutionResult{Result: future}
}

func (a *ConfigServiceStub) StopConfigurationRecorder(ctx workflow.Context, input *configservice.StopConfigurationRecorderInput) (*configservice.StopConfigurationRecorderOutput, error) {
    var output configservice.StopConfigurationRecorderOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StopConfigurationRecorder, input).Get(ctx, &output)
    return &output, err
}

func (a *ConfigServiceStub) StopConfigurationRecorderAsync(ctx workflow.Context, input *configservice.StopConfigurationRecorderInput) *ConfigserviceStopConfigurationRecorderResult {
    future := workflow.ExecuteActivity(ctx, a.activities.StopConfigurationRecorder, input)
    return &ConfigserviceStopConfigurationRecorderResult{Result: future}
}

func (a *ConfigServiceStub) TagResource(ctx workflow.Context, input *configservice.TagResourceInput) (*configservice.TagResourceOutput, error) {
    var output configservice.TagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.TagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *ConfigServiceStub) TagResourceAsync(ctx workflow.Context, input *configservice.TagResourceInput) *ConfigserviceTagResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.TagResource, input)
    return &ConfigserviceTagResourceResult{Result: future}
}

func (a *ConfigServiceStub) UntagResource(ctx workflow.Context, input *configservice.UntagResourceInput) (*configservice.UntagResourceOutput, error) {
    var output configservice.UntagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *ConfigServiceStub) UntagResourceAsync(ctx workflow.Context, input *configservice.UntagResourceInput) *ConfigserviceUntagResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input)
    return &ConfigserviceUntagResourceResult{Result: future}
}