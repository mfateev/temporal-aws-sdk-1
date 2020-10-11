// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package awsclients

import (
	"github.com/aws/aws-sdk-go/service/cloudwatch"
	"go.temporal.io/sdk/workflow"
)

type CloudWatchClient interface {
	DeleteAlarms(ctx workflow.Context, input *cloudwatch.DeleteAlarmsInput) (*cloudwatch.DeleteAlarmsOutput, error)
	DeleteAlarmsAsync(ctx workflow.Context, input *cloudwatch.DeleteAlarmsInput) *CloudwatchDeleteAlarmsResult

	DeleteAnomalyDetector(ctx workflow.Context, input *cloudwatch.DeleteAnomalyDetectorInput) (*cloudwatch.DeleteAnomalyDetectorOutput, error)
	DeleteAnomalyDetectorAsync(ctx workflow.Context, input *cloudwatch.DeleteAnomalyDetectorInput) *CloudwatchDeleteAnomalyDetectorResult

	DeleteDashboards(ctx workflow.Context, input *cloudwatch.DeleteDashboardsInput) (*cloudwatch.DeleteDashboardsOutput, error)
	DeleteDashboardsAsync(ctx workflow.Context, input *cloudwatch.DeleteDashboardsInput) *CloudwatchDeleteDashboardsResult

	DeleteInsightRules(ctx workflow.Context, input *cloudwatch.DeleteInsightRulesInput) (*cloudwatch.DeleteInsightRulesOutput, error)
	DeleteInsightRulesAsync(ctx workflow.Context, input *cloudwatch.DeleteInsightRulesInput) *CloudwatchDeleteInsightRulesResult

	DescribeAlarmHistory(ctx workflow.Context, input *cloudwatch.DescribeAlarmHistoryInput) (*cloudwatch.DescribeAlarmHistoryOutput, error)
	DescribeAlarmHistoryAsync(ctx workflow.Context, input *cloudwatch.DescribeAlarmHistoryInput) *CloudwatchDescribeAlarmHistoryResult

	DescribeAlarms(ctx workflow.Context, input *cloudwatch.DescribeAlarmsInput) (*cloudwatch.DescribeAlarmsOutput, error)
	DescribeAlarmsAsync(ctx workflow.Context, input *cloudwatch.DescribeAlarmsInput) *CloudwatchDescribeAlarmsResult

	DescribeAlarmsForMetric(ctx workflow.Context, input *cloudwatch.DescribeAlarmsForMetricInput) (*cloudwatch.DescribeAlarmsForMetricOutput, error)
	DescribeAlarmsForMetricAsync(ctx workflow.Context, input *cloudwatch.DescribeAlarmsForMetricInput) *CloudwatchDescribeAlarmsForMetricResult

	DescribeAnomalyDetectors(ctx workflow.Context, input *cloudwatch.DescribeAnomalyDetectorsInput) (*cloudwatch.DescribeAnomalyDetectorsOutput, error)
	DescribeAnomalyDetectorsAsync(ctx workflow.Context, input *cloudwatch.DescribeAnomalyDetectorsInput) *CloudwatchDescribeAnomalyDetectorsResult

	DescribeInsightRules(ctx workflow.Context, input *cloudwatch.DescribeInsightRulesInput) (*cloudwatch.DescribeInsightRulesOutput, error)
	DescribeInsightRulesAsync(ctx workflow.Context, input *cloudwatch.DescribeInsightRulesInput) *CloudwatchDescribeInsightRulesResult

	DisableAlarmActions(ctx workflow.Context, input *cloudwatch.DisableAlarmActionsInput) (*cloudwatch.DisableAlarmActionsOutput, error)
	DisableAlarmActionsAsync(ctx workflow.Context, input *cloudwatch.DisableAlarmActionsInput) *CloudwatchDisableAlarmActionsResult

	DisableInsightRules(ctx workflow.Context, input *cloudwatch.DisableInsightRulesInput) (*cloudwatch.DisableInsightRulesOutput, error)
	DisableInsightRulesAsync(ctx workflow.Context, input *cloudwatch.DisableInsightRulesInput) *CloudwatchDisableInsightRulesResult

	EnableAlarmActions(ctx workflow.Context, input *cloudwatch.EnableAlarmActionsInput) (*cloudwatch.EnableAlarmActionsOutput, error)
	EnableAlarmActionsAsync(ctx workflow.Context, input *cloudwatch.EnableAlarmActionsInput) *CloudwatchEnableAlarmActionsResult

	EnableInsightRules(ctx workflow.Context, input *cloudwatch.EnableInsightRulesInput) (*cloudwatch.EnableInsightRulesOutput, error)
	EnableInsightRulesAsync(ctx workflow.Context, input *cloudwatch.EnableInsightRulesInput) *CloudwatchEnableInsightRulesResult

	GetDashboard(ctx workflow.Context, input *cloudwatch.GetDashboardInput) (*cloudwatch.GetDashboardOutput, error)
	GetDashboardAsync(ctx workflow.Context, input *cloudwatch.GetDashboardInput) *CloudwatchGetDashboardResult

	GetInsightRuleReport(ctx workflow.Context, input *cloudwatch.GetInsightRuleReportInput) (*cloudwatch.GetInsightRuleReportOutput, error)
	GetInsightRuleReportAsync(ctx workflow.Context, input *cloudwatch.GetInsightRuleReportInput) *CloudwatchGetInsightRuleReportResult

	GetMetricData(ctx workflow.Context, input *cloudwatch.GetMetricDataInput) (*cloudwatch.GetMetricDataOutput, error)
	GetMetricDataAsync(ctx workflow.Context, input *cloudwatch.GetMetricDataInput) *CloudwatchGetMetricDataResult

	GetMetricStatistics(ctx workflow.Context, input *cloudwatch.GetMetricStatisticsInput) (*cloudwatch.GetMetricStatisticsOutput, error)
	GetMetricStatisticsAsync(ctx workflow.Context, input *cloudwatch.GetMetricStatisticsInput) *CloudwatchGetMetricStatisticsResult

	GetMetricWidgetImage(ctx workflow.Context, input *cloudwatch.GetMetricWidgetImageInput) (*cloudwatch.GetMetricWidgetImageOutput, error)
	GetMetricWidgetImageAsync(ctx workflow.Context, input *cloudwatch.GetMetricWidgetImageInput) *CloudwatchGetMetricWidgetImageResult

	ListDashboards(ctx workflow.Context, input *cloudwatch.ListDashboardsInput) (*cloudwatch.ListDashboardsOutput, error)
	ListDashboardsAsync(ctx workflow.Context, input *cloudwatch.ListDashboardsInput) *CloudwatchListDashboardsResult

	ListMetrics(ctx workflow.Context, input *cloudwatch.ListMetricsInput) (*cloudwatch.ListMetricsOutput, error)
	ListMetricsAsync(ctx workflow.Context, input *cloudwatch.ListMetricsInput) *CloudwatchListMetricsResult

	ListTagsForResource(ctx workflow.Context, input *cloudwatch.ListTagsForResourceInput) (*cloudwatch.ListTagsForResourceOutput, error)
	ListTagsForResourceAsync(ctx workflow.Context, input *cloudwatch.ListTagsForResourceInput) *CloudwatchListTagsForResourceResult

	PutAnomalyDetector(ctx workflow.Context, input *cloudwatch.PutAnomalyDetectorInput) (*cloudwatch.PutAnomalyDetectorOutput, error)
	PutAnomalyDetectorAsync(ctx workflow.Context, input *cloudwatch.PutAnomalyDetectorInput) *CloudwatchPutAnomalyDetectorResult

	PutCompositeAlarm(ctx workflow.Context, input *cloudwatch.PutCompositeAlarmInput) (*cloudwatch.PutCompositeAlarmOutput, error)
	PutCompositeAlarmAsync(ctx workflow.Context, input *cloudwatch.PutCompositeAlarmInput) *CloudwatchPutCompositeAlarmResult

	PutDashboard(ctx workflow.Context, input *cloudwatch.PutDashboardInput) (*cloudwatch.PutDashboardOutput, error)
	PutDashboardAsync(ctx workflow.Context, input *cloudwatch.PutDashboardInput) *CloudwatchPutDashboardResult

	PutInsightRule(ctx workflow.Context, input *cloudwatch.PutInsightRuleInput) (*cloudwatch.PutInsightRuleOutput, error)
	PutInsightRuleAsync(ctx workflow.Context, input *cloudwatch.PutInsightRuleInput) *CloudwatchPutInsightRuleResult

	PutMetricAlarm(ctx workflow.Context, input *cloudwatch.PutMetricAlarmInput) (*cloudwatch.PutMetricAlarmOutput, error)
	PutMetricAlarmAsync(ctx workflow.Context, input *cloudwatch.PutMetricAlarmInput) *CloudwatchPutMetricAlarmResult

	PutMetricData(ctx workflow.Context, input *cloudwatch.PutMetricDataInput) (*cloudwatch.PutMetricDataOutput, error)
	PutMetricDataAsync(ctx workflow.Context, input *cloudwatch.PutMetricDataInput) *CloudwatchPutMetricDataResult

	SetAlarmState(ctx workflow.Context, input *cloudwatch.SetAlarmStateInput) (*cloudwatch.SetAlarmStateOutput, error)
	SetAlarmStateAsync(ctx workflow.Context, input *cloudwatch.SetAlarmStateInput) *CloudwatchSetAlarmStateResult

	TagResource(ctx workflow.Context, input *cloudwatch.TagResourceInput) (*cloudwatch.TagResourceOutput, error)
	TagResourceAsync(ctx workflow.Context, input *cloudwatch.TagResourceInput) *CloudwatchTagResourceResult

	UntagResource(ctx workflow.Context, input *cloudwatch.UntagResourceInput) (*cloudwatch.UntagResourceOutput, error)
	UntagResourceAsync(ctx workflow.Context, input *cloudwatch.UntagResourceInput) *CloudwatchUntagResourceResult

	WaitUntilAlarmExists(ctx workflow.Context, input *cloudwatch.DescribeAlarmsInput) error

	WaitUntilCompositeAlarmExists(ctx workflow.Context, input *cloudwatch.DescribeAlarmsInput) error
}

type CloudWatchStub struct{}

func NewCloudWatchStub() CloudWatchClient {
	return &CloudWatchStub{}
}

type CloudwatchDeleteAlarmsResult struct {
	Result workflow.Future
}

func (r *CloudwatchDeleteAlarmsResult) Get(ctx workflow.Context) (*cloudwatch.DeleteAlarmsOutput, error) {
	var output cloudwatch.DeleteAlarmsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type CloudwatchDeleteAnomalyDetectorResult struct {
	Result workflow.Future
}

func (r *CloudwatchDeleteAnomalyDetectorResult) Get(ctx workflow.Context) (*cloudwatch.DeleteAnomalyDetectorOutput, error) {
	var output cloudwatch.DeleteAnomalyDetectorOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type CloudwatchDeleteDashboardsResult struct {
	Result workflow.Future
}

func (r *CloudwatchDeleteDashboardsResult) Get(ctx workflow.Context) (*cloudwatch.DeleteDashboardsOutput, error) {
	var output cloudwatch.DeleteDashboardsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type CloudwatchDeleteInsightRulesResult struct {
	Result workflow.Future
}

func (r *CloudwatchDeleteInsightRulesResult) Get(ctx workflow.Context) (*cloudwatch.DeleteInsightRulesOutput, error) {
	var output cloudwatch.DeleteInsightRulesOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type CloudwatchDescribeAlarmHistoryResult struct {
	Result workflow.Future
}

func (r *CloudwatchDescribeAlarmHistoryResult) Get(ctx workflow.Context) (*cloudwatch.DescribeAlarmHistoryOutput, error) {
	var output cloudwatch.DescribeAlarmHistoryOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type CloudwatchDescribeAlarmsResult struct {
	Result workflow.Future
}

func (r *CloudwatchDescribeAlarmsResult) Get(ctx workflow.Context) (*cloudwatch.DescribeAlarmsOutput, error) {
	var output cloudwatch.DescribeAlarmsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type CloudwatchDescribeAlarmsForMetricResult struct {
	Result workflow.Future
}

func (r *CloudwatchDescribeAlarmsForMetricResult) Get(ctx workflow.Context) (*cloudwatch.DescribeAlarmsForMetricOutput, error) {
	var output cloudwatch.DescribeAlarmsForMetricOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type CloudwatchDescribeAnomalyDetectorsResult struct {
	Result workflow.Future
}

func (r *CloudwatchDescribeAnomalyDetectorsResult) Get(ctx workflow.Context) (*cloudwatch.DescribeAnomalyDetectorsOutput, error) {
	var output cloudwatch.DescribeAnomalyDetectorsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type CloudwatchDescribeInsightRulesResult struct {
	Result workflow.Future
}

func (r *CloudwatchDescribeInsightRulesResult) Get(ctx workflow.Context) (*cloudwatch.DescribeInsightRulesOutput, error) {
	var output cloudwatch.DescribeInsightRulesOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type CloudwatchDisableAlarmActionsResult struct {
	Result workflow.Future
}

func (r *CloudwatchDisableAlarmActionsResult) Get(ctx workflow.Context) (*cloudwatch.DisableAlarmActionsOutput, error) {
	var output cloudwatch.DisableAlarmActionsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type CloudwatchDisableInsightRulesResult struct {
	Result workflow.Future
}

func (r *CloudwatchDisableInsightRulesResult) Get(ctx workflow.Context) (*cloudwatch.DisableInsightRulesOutput, error) {
	var output cloudwatch.DisableInsightRulesOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type CloudwatchEnableAlarmActionsResult struct {
	Result workflow.Future
}

func (r *CloudwatchEnableAlarmActionsResult) Get(ctx workflow.Context) (*cloudwatch.EnableAlarmActionsOutput, error) {
	var output cloudwatch.EnableAlarmActionsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type CloudwatchEnableInsightRulesResult struct {
	Result workflow.Future
}

func (r *CloudwatchEnableInsightRulesResult) Get(ctx workflow.Context) (*cloudwatch.EnableInsightRulesOutput, error) {
	var output cloudwatch.EnableInsightRulesOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type CloudwatchGetDashboardResult struct {
	Result workflow.Future
}

func (r *CloudwatchGetDashboardResult) Get(ctx workflow.Context) (*cloudwatch.GetDashboardOutput, error) {
	var output cloudwatch.GetDashboardOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type CloudwatchGetInsightRuleReportResult struct {
	Result workflow.Future
}

func (r *CloudwatchGetInsightRuleReportResult) Get(ctx workflow.Context) (*cloudwatch.GetInsightRuleReportOutput, error) {
	var output cloudwatch.GetInsightRuleReportOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type CloudwatchGetMetricDataResult struct {
	Result workflow.Future
}

func (r *CloudwatchGetMetricDataResult) Get(ctx workflow.Context) (*cloudwatch.GetMetricDataOutput, error) {
	var output cloudwatch.GetMetricDataOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type CloudwatchGetMetricStatisticsResult struct {
	Result workflow.Future
}

func (r *CloudwatchGetMetricStatisticsResult) Get(ctx workflow.Context) (*cloudwatch.GetMetricStatisticsOutput, error) {
	var output cloudwatch.GetMetricStatisticsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type CloudwatchGetMetricWidgetImageResult struct {
	Result workflow.Future
}

func (r *CloudwatchGetMetricWidgetImageResult) Get(ctx workflow.Context) (*cloudwatch.GetMetricWidgetImageOutput, error) {
	var output cloudwatch.GetMetricWidgetImageOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type CloudwatchListDashboardsResult struct {
	Result workflow.Future
}

func (r *CloudwatchListDashboardsResult) Get(ctx workflow.Context) (*cloudwatch.ListDashboardsOutput, error) {
	var output cloudwatch.ListDashboardsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type CloudwatchListMetricsResult struct {
	Result workflow.Future
}

func (r *CloudwatchListMetricsResult) Get(ctx workflow.Context) (*cloudwatch.ListMetricsOutput, error) {
	var output cloudwatch.ListMetricsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type CloudwatchListTagsForResourceResult struct {
	Result workflow.Future
}

func (r *CloudwatchListTagsForResourceResult) Get(ctx workflow.Context) (*cloudwatch.ListTagsForResourceOutput, error) {
	var output cloudwatch.ListTagsForResourceOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type CloudwatchPutAnomalyDetectorResult struct {
	Result workflow.Future
}

func (r *CloudwatchPutAnomalyDetectorResult) Get(ctx workflow.Context) (*cloudwatch.PutAnomalyDetectorOutput, error) {
	var output cloudwatch.PutAnomalyDetectorOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type CloudwatchPutCompositeAlarmResult struct {
	Result workflow.Future
}

func (r *CloudwatchPutCompositeAlarmResult) Get(ctx workflow.Context) (*cloudwatch.PutCompositeAlarmOutput, error) {
	var output cloudwatch.PutCompositeAlarmOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type CloudwatchPutDashboardResult struct {
	Result workflow.Future
}

func (r *CloudwatchPutDashboardResult) Get(ctx workflow.Context) (*cloudwatch.PutDashboardOutput, error) {
	var output cloudwatch.PutDashboardOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type CloudwatchPutInsightRuleResult struct {
	Result workflow.Future
}

func (r *CloudwatchPutInsightRuleResult) Get(ctx workflow.Context) (*cloudwatch.PutInsightRuleOutput, error) {
	var output cloudwatch.PutInsightRuleOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type CloudwatchPutMetricAlarmResult struct {
	Result workflow.Future
}

func (r *CloudwatchPutMetricAlarmResult) Get(ctx workflow.Context) (*cloudwatch.PutMetricAlarmOutput, error) {
	var output cloudwatch.PutMetricAlarmOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type CloudwatchPutMetricDataResult struct {
	Result workflow.Future
}

func (r *CloudwatchPutMetricDataResult) Get(ctx workflow.Context) (*cloudwatch.PutMetricDataOutput, error) {
	var output cloudwatch.PutMetricDataOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type CloudwatchSetAlarmStateResult struct {
	Result workflow.Future
}

func (r *CloudwatchSetAlarmStateResult) Get(ctx workflow.Context) (*cloudwatch.SetAlarmStateOutput, error) {
	var output cloudwatch.SetAlarmStateOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type CloudwatchTagResourceResult struct {
	Result workflow.Future
}

func (r *CloudwatchTagResourceResult) Get(ctx workflow.Context) (*cloudwatch.TagResourceOutput, error) {
	var output cloudwatch.TagResourceOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type CloudwatchUntagResourceResult struct {
	Result workflow.Future
}

func (r *CloudwatchUntagResourceResult) Get(ctx workflow.Context) (*cloudwatch.UntagResourceOutput, error) {
	var output cloudwatch.UntagResourceOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

func (a *CloudWatchStub) DeleteAlarms(ctx workflow.Context, input *cloudwatch.DeleteAlarmsInput) (*cloudwatch.DeleteAlarmsOutput, error) {
	var output cloudwatch.DeleteAlarmsOutput
	err := workflow.ExecuteActivity(ctx, "aws.cloudwatch.DeleteAlarms", input).Get(ctx, &output)
	return &output, err
}

func (a *CloudWatchStub) DeleteAlarmsAsync(ctx workflow.Context, input *cloudwatch.DeleteAlarmsInput) *CloudwatchDeleteAlarmsResult {
	future := workflow.ExecuteActivity(ctx, "aws.cloudwatch.DeleteAlarms", input)
	return &CloudwatchDeleteAlarmsResult{Result: future}
}

func (a *CloudWatchStub) DeleteAnomalyDetector(ctx workflow.Context, input *cloudwatch.DeleteAnomalyDetectorInput) (*cloudwatch.DeleteAnomalyDetectorOutput, error) {
	var output cloudwatch.DeleteAnomalyDetectorOutput
	err := workflow.ExecuteActivity(ctx, "aws.cloudwatch.DeleteAnomalyDetector", input).Get(ctx, &output)
	return &output, err
}

func (a *CloudWatchStub) DeleteAnomalyDetectorAsync(ctx workflow.Context, input *cloudwatch.DeleteAnomalyDetectorInput) *CloudwatchDeleteAnomalyDetectorResult {
	future := workflow.ExecuteActivity(ctx, "aws.cloudwatch.DeleteAnomalyDetector", input)
	return &CloudwatchDeleteAnomalyDetectorResult{Result: future}
}

func (a *CloudWatchStub) DeleteDashboards(ctx workflow.Context, input *cloudwatch.DeleteDashboardsInput) (*cloudwatch.DeleteDashboardsOutput, error) {
	var output cloudwatch.DeleteDashboardsOutput
	err := workflow.ExecuteActivity(ctx, "aws.cloudwatch.DeleteDashboards", input).Get(ctx, &output)
	return &output, err
}

func (a *CloudWatchStub) DeleteDashboardsAsync(ctx workflow.Context, input *cloudwatch.DeleteDashboardsInput) *CloudwatchDeleteDashboardsResult {
	future := workflow.ExecuteActivity(ctx, "aws.cloudwatch.DeleteDashboards", input)
	return &CloudwatchDeleteDashboardsResult{Result: future}
}

func (a *CloudWatchStub) DeleteInsightRules(ctx workflow.Context, input *cloudwatch.DeleteInsightRulesInput) (*cloudwatch.DeleteInsightRulesOutput, error) {
	var output cloudwatch.DeleteInsightRulesOutput
	err := workflow.ExecuteActivity(ctx, "aws.cloudwatch.DeleteInsightRules", input).Get(ctx, &output)
	return &output, err
}

func (a *CloudWatchStub) DeleteInsightRulesAsync(ctx workflow.Context, input *cloudwatch.DeleteInsightRulesInput) *CloudwatchDeleteInsightRulesResult {
	future := workflow.ExecuteActivity(ctx, "aws.cloudwatch.DeleteInsightRules", input)
	return &CloudwatchDeleteInsightRulesResult{Result: future}
}

func (a *CloudWatchStub) DescribeAlarmHistory(ctx workflow.Context, input *cloudwatch.DescribeAlarmHistoryInput) (*cloudwatch.DescribeAlarmHistoryOutput, error) {
	var output cloudwatch.DescribeAlarmHistoryOutput
	err := workflow.ExecuteActivity(ctx, "aws.cloudwatch.DescribeAlarmHistory", input).Get(ctx, &output)
	return &output, err
}

func (a *CloudWatchStub) DescribeAlarmHistoryAsync(ctx workflow.Context, input *cloudwatch.DescribeAlarmHistoryInput) *CloudwatchDescribeAlarmHistoryResult {
	future := workflow.ExecuteActivity(ctx, "aws.cloudwatch.DescribeAlarmHistory", input)
	return &CloudwatchDescribeAlarmHistoryResult{Result: future}
}

func (a *CloudWatchStub) DescribeAlarms(ctx workflow.Context, input *cloudwatch.DescribeAlarmsInput) (*cloudwatch.DescribeAlarmsOutput, error) {
	var output cloudwatch.DescribeAlarmsOutput
	err := workflow.ExecuteActivity(ctx, "aws.cloudwatch.DescribeAlarms", input).Get(ctx, &output)
	return &output, err
}

func (a *CloudWatchStub) DescribeAlarmsAsync(ctx workflow.Context, input *cloudwatch.DescribeAlarmsInput) *CloudwatchDescribeAlarmsResult {
	future := workflow.ExecuteActivity(ctx, "aws.cloudwatch.DescribeAlarms", input)
	return &CloudwatchDescribeAlarmsResult{Result: future}
}

func (a *CloudWatchStub) DescribeAlarmsForMetric(ctx workflow.Context, input *cloudwatch.DescribeAlarmsForMetricInput) (*cloudwatch.DescribeAlarmsForMetricOutput, error) {
	var output cloudwatch.DescribeAlarmsForMetricOutput
	err := workflow.ExecuteActivity(ctx, "aws.cloudwatch.DescribeAlarmsForMetric", input).Get(ctx, &output)
	return &output, err
}

func (a *CloudWatchStub) DescribeAlarmsForMetricAsync(ctx workflow.Context, input *cloudwatch.DescribeAlarmsForMetricInput) *CloudwatchDescribeAlarmsForMetricResult {
	future := workflow.ExecuteActivity(ctx, "aws.cloudwatch.DescribeAlarmsForMetric", input)
	return &CloudwatchDescribeAlarmsForMetricResult{Result: future}
}

func (a *CloudWatchStub) DescribeAnomalyDetectors(ctx workflow.Context, input *cloudwatch.DescribeAnomalyDetectorsInput) (*cloudwatch.DescribeAnomalyDetectorsOutput, error) {
	var output cloudwatch.DescribeAnomalyDetectorsOutput
	err := workflow.ExecuteActivity(ctx, "aws.cloudwatch.DescribeAnomalyDetectors", input).Get(ctx, &output)
	return &output, err
}

func (a *CloudWatchStub) DescribeAnomalyDetectorsAsync(ctx workflow.Context, input *cloudwatch.DescribeAnomalyDetectorsInput) *CloudwatchDescribeAnomalyDetectorsResult {
	future := workflow.ExecuteActivity(ctx, "aws.cloudwatch.DescribeAnomalyDetectors", input)
	return &CloudwatchDescribeAnomalyDetectorsResult{Result: future}
}

func (a *CloudWatchStub) DescribeInsightRules(ctx workflow.Context, input *cloudwatch.DescribeInsightRulesInput) (*cloudwatch.DescribeInsightRulesOutput, error) {
	var output cloudwatch.DescribeInsightRulesOutput
	err := workflow.ExecuteActivity(ctx, "aws.cloudwatch.DescribeInsightRules", input).Get(ctx, &output)
	return &output, err
}

func (a *CloudWatchStub) DescribeInsightRulesAsync(ctx workflow.Context, input *cloudwatch.DescribeInsightRulesInput) *CloudwatchDescribeInsightRulesResult {
	future := workflow.ExecuteActivity(ctx, "aws.cloudwatch.DescribeInsightRules", input)
	return &CloudwatchDescribeInsightRulesResult{Result: future}
}

func (a *CloudWatchStub) DisableAlarmActions(ctx workflow.Context, input *cloudwatch.DisableAlarmActionsInput) (*cloudwatch.DisableAlarmActionsOutput, error) {
	var output cloudwatch.DisableAlarmActionsOutput
	err := workflow.ExecuteActivity(ctx, "aws.cloudwatch.DisableAlarmActions", input).Get(ctx, &output)
	return &output, err
}

func (a *CloudWatchStub) DisableAlarmActionsAsync(ctx workflow.Context, input *cloudwatch.DisableAlarmActionsInput) *CloudwatchDisableAlarmActionsResult {
	future := workflow.ExecuteActivity(ctx, "aws.cloudwatch.DisableAlarmActions", input)
	return &CloudwatchDisableAlarmActionsResult{Result: future}
}

func (a *CloudWatchStub) DisableInsightRules(ctx workflow.Context, input *cloudwatch.DisableInsightRulesInput) (*cloudwatch.DisableInsightRulesOutput, error) {
	var output cloudwatch.DisableInsightRulesOutput
	err := workflow.ExecuteActivity(ctx, "aws.cloudwatch.DisableInsightRules", input).Get(ctx, &output)
	return &output, err
}

func (a *CloudWatchStub) DisableInsightRulesAsync(ctx workflow.Context, input *cloudwatch.DisableInsightRulesInput) *CloudwatchDisableInsightRulesResult {
	future := workflow.ExecuteActivity(ctx, "aws.cloudwatch.DisableInsightRules", input)
	return &CloudwatchDisableInsightRulesResult{Result: future}
}

func (a *CloudWatchStub) EnableAlarmActions(ctx workflow.Context, input *cloudwatch.EnableAlarmActionsInput) (*cloudwatch.EnableAlarmActionsOutput, error) {
	var output cloudwatch.EnableAlarmActionsOutput
	err := workflow.ExecuteActivity(ctx, "aws.cloudwatch.EnableAlarmActions", input).Get(ctx, &output)
	return &output, err
}

func (a *CloudWatchStub) EnableAlarmActionsAsync(ctx workflow.Context, input *cloudwatch.EnableAlarmActionsInput) *CloudwatchEnableAlarmActionsResult {
	future := workflow.ExecuteActivity(ctx, "aws.cloudwatch.EnableAlarmActions", input)
	return &CloudwatchEnableAlarmActionsResult{Result: future}
}

func (a *CloudWatchStub) EnableInsightRules(ctx workflow.Context, input *cloudwatch.EnableInsightRulesInput) (*cloudwatch.EnableInsightRulesOutput, error) {
	var output cloudwatch.EnableInsightRulesOutput
	err := workflow.ExecuteActivity(ctx, "aws.cloudwatch.EnableInsightRules", input).Get(ctx, &output)
	return &output, err
}

func (a *CloudWatchStub) EnableInsightRulesAsync(ctx workflow.Context, input *cloudwatch.EnableInsightRulesInput) *CloudwatchEnableInsightRulesResult {
	future := workflow.ExecuteActivity(ctx, "aws.cloudwatch.EnableInsightRules", input)
	return &CloudwatchEnableInsightRulesResult{Result: future}
}

func (a *CloudWatchStub) GetDashboard(ctx workflow.Context, input *cloudwatch.GetDashboardInput) (*cloudwatch.GetDashboardOutput, error) {
	var output cloudwatch.GetDashboardOutput
	err := workflow.ExecuteActivity(ctx, "aws.cloudwatch.GetDashboard", input).Get(ctx, &output)
	return &output, err
}

func (a *CloudWatchStub) GetDashboardAsync(ctx workflow.Context, input *cloudwatch.GetDashboardInput) *CloudwatchGetDashboardResult {
	future := workflow.ExecuteActivity(ctx, "aws.cloudwatch.GetDashboard", input)
	return &CloudwatchGetDashboardResult{Result: future}
}

func (a *CloudWatchStub) GetInsightRuleReport(ctx workflow.Context, input *cloudwatch.GetInsightRuleReportInput) (*cloudwatch.GetInsightRuleReportOutput, error) {
	var output cloudwatch.GetInsightRuleReportOutput
	err := workflow.ExecuteActivity(ctx, "aws.cloudwatch.GetInsightRuleReport", input).Get(ctx, &output)
	return &output, err
}

func (a *CloudWatchStub) GetInsightRuleReportAsync(ctx workflow.Context, input *cloudwatch.GetInsightRuleReportInput) *CloudwatchGetInsightRuleReportResult {
	future := workflow.ExecuteActivity(ctx, "aws.cloudwatch.GetInsightRuleReport", input)
	return &CloudwatchGetInsightRuleReportResult{Result: future}
}

func (a *CloudWatchStub) GetMetricData(ctx workflow.Context, input *cloudwatch.GetMetricDataInput) (*cloudwatch.GetMetricDataOutput, error) {
	var output cloudwatch.GetMetricDataOutput
	err := workflow.ExecuteActivity(ctx, "aws.cloudwatch.GetMetricData", input).Get(ctx, &output)
	return &output, err
}

func (a *CloudWatchStub) GetMetricDataAsync(ctx workflow.Context, input *cloudwatch.GetMetricDataInput) *CloudwatchGetMetricDataResult {
	future := workflow.ExecuteActivity(ctx, "aws.cloudwatch.GetMetricData", input)
	return &CloudwatchGetMetricDataResult{Result: future}
}

func (a *CloudWatchStub) GetMetricStatistics(ctx workflow.Context, input *cloudwatch.GetMetricStatisticsInput) (*cloudwatch.GetMetricStatisticsOutput, error) {
	var output cloudwatch.GetMetricStatisticsOutput
	err := workflow.ExecuteActivity(ctx, "aws.cloudwatch.GetMetricStatistics", input).Get(ctx, &output)
	return &output, err
}

func (a *CloudWatchStub) GetMetricStatisticsAsync(ctx workflow.Context, input *cloudwatch.GetMetricStatisticsInput) *CloudwatchGetMetricStatisticsResult {
	future := workflow.ExecuteActivity(ctx, "aws.cloudwatch.GetMetricStatistics", input)
	return &CloudwatchGetMetricStatisticsResult{Result: future}
}

func (a *CloudWatchStub) GetMetricWidgetImage(ctx workflow.Context, input *cloudwatch.GetMetricWidgetImageInput) (*cloudwatch.GetMetricWidgetImageOutput, error) {
	var output cloudwatch.GetMetricWidgetImageOutput
	err := workflow.ExecuteActivity(ctx, "aws.cloudwatch.GetMetricWidgetImage", input).Get(ctx, &output)
	return &output, err
}

func (a *CloudWatchStub) GetMetricWidgetImageAsync(ctx workflow.Context, input *cloudwatch.GetMetricWidgetImageInput) *CloudwatchGetMetricWidgetImageResult {
	future := workflow.ExecuteActivity(ctx, "aws.cloudwatch.GetMetricWidgetImage", input)
	return &CloudwatchGetMetricWidgetImageResult{Result: future}
}

func (a *CloudWatchStub) ListDashboards(ctx workflow.Context, input *cloudwatch.ListDashboardsInput) (*cloudwatch.ListDashboardsOutput, error) {
	var output cloudwatch.ListDashboardsOutput
	err := workflow.ExecuteActivity(ctx, "aws.cloudwatch.ListDashboards", input).Get(ctx, &output)
	return &output, err
}

func (a *CloudWatchStub) ListDashboardsAsync(ctx workflow.Context, input *cloudwatch.ListDashboardsInput) *CloudwatchListDashboardsResult {
	future := workflow.ExecuteActivity(ctx, "aws.cloudwatch.ListDashboards", input)
	return &CloudwatchListDashboardsResult{Result: future}
}

func (a *CloudWatchStub) ListMetrics(ctx workflow.Context, input *cloudwatch.ListMetricsInput) (*cloudwatch.ListMetricsOutput, error) {
	var output cloudwatch.ListMetricsOutput
	err := workflow.ExecuteActivity(ctx, "aws.cloudwatch.ListMetrics", input).Get(ctx, &output)
	return &output, err
}

func (a *CloudWatchStub) ListMetricsAsync(ctx workflow.Context, input *cloudwatch.ListMetricsInput) *CloudwatchListMetricsResult {
	future := workflow.ExecuteActivity(ctx, "aws.cloudwatch.ListMetrics", input)
	return &CloudwatchListMetricsResult{Result: future}
}

func (a *CloudWatchStub) ListTagsForResource(ctx workflow.Context, input *cloudwatch.ListTagsForResourceInput) (*cloudwatch.ListTagsForResourceOutput, error) {
	var output cloudwatch.ListTagsForResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.cloudwatch.ListTagsForResource", input).Get(ctx, &output)
	return &output, err
}

func (a *CloudWatchStub) ListTagsForResourceAsync(ctx workflow.Context, input *cloudwatch.ListTagsForResourceInput) *CloudwatchListTagsForResourceResult {
	future := workflow.ExecuteActivity(ctx, "aws.cloudwatch.ListTagsForResource", input)
	return &CloudwatchListTagsForResourceResult{Result: future}
}

func (a *CloudWatchStub) PutAnomalyDetector(ctx workflow.Context, input *cloudwatch.PutAnomalyDetectorInput) (*cloudwatch.PutAnomalyDetectorOutput, error) {
	var output cloudwatch.PutAnomalyDetectorOutput
	err := workflow.ExecuteActivity(ctx, "aws.cloudwatch.PutAnomalyDetector", input).Get(ctx, &output)
	return &output, err
}

func (a *CloudWatchStub) PutAnomalyDetectorAsync(ctx workflow.Context, input *cloudwatch.PutAnomalyDetectorInput) *CloudwatchPutAnomalyDetectorResult {
	future := workflow.ExecuteActivity(ctx, "aws.cloudwatch.PutAnomalyDetector", input)
	return &CloudwatchPutAnomalyDetectorResult{Result: future}
}

func (a *CloudWatchStub) PutCompositeAlarm(ctx workflow.Context, input *cloudwatch.PutCompositeAlarmInput) (*cloudwatch.PutCompositeAlarmOutput, error) {
	var output cloudwatch.PutCompositeAlarmOutput
	err := workflow.ExecuteActivity(ctx, "aws.cloudwatch.PutCompositeAlarm", input).Get(ctx, &output)
	return &output, err
}

func (a *CloudWatchStub) PutCompositeAlarmAsync(ctx workflow.Context, input *cloudwatch.PutCompositeAlarmInput) *CloudwatchPutCompositeAlarmResult {
	future := workflow.ExecuteActivity(ctx, "aws.cloudwatch.PutCompositeAlarm", input)
	return &CloudwatchPutCompositeAlarmResult{Result: future}
}

func (a *CloudWatchStub) PutDashboard(ctx workflow.Context, input *cloudwatch.PutDashboardInput) (*cloudwatch.PutDashboardOutput, error) {
	var output cloudwatch.PutDashboardOutput
	err := workflow.ExecuteActivity(ctx, "aws.cloudwatch.PutDashboard", input).Get(ctx, &output)
	return &output, err
}

func (a *CloudWatchStub) PutDashboardAsync(ctx workflow.Context, input *cloudwatch.PutDashboardInput) *CloudwatchPutDashboardResult {
	future := workflow.ExecuteActivity(ctx, "aws.cloudwatch.PutDashboard", input)
	return &CloudwatchPutDashboardResult{Result: future}
}

func (a *CloudWatchStub) PutInsightRule(ctx workflow.Context, input *cloudwatch.PutInsightRuleInput) (*cloudwatch.PutInsightRuleOutput, error) {
	var output cloudwatch.PutInsightRuleOutput
	err := workflow.ExecuteActivity(ctx, "aws.cloudwatch.PutInsightRule", input).Get(ctx, &output)
	return &output, err
}

func (a *CloudWatchStub) PutInsightRuleAsync(ctx workflow.Context, input *cloudwatch.PutInsightRuleInput) *CloudwatchPutInsightRuleResult {
	future := workflow.ExecuteActivity(ctx, "aws.cloudwatch.PutInsightRule", input)
	return &CloudwatchPutInsightRuleResult{Result: future}
}

func (a *CloudWatchStub) PutMetricAlarm(ctx workflow.Context, input *cloudwatch.PutMetricAlarmInput) (*cloudwatch.PutMetricAlarmOutput, error) {
	var output cloudwatch.PutMetricAlarmOutput
	err := workflow.ExecuteActivity(ctx, "aws.cloudwatch.PutMetricAlarm", input).Get(ctx, &output)
	return &output, err
}

func (a *CloudWatchStub) PutMetricAlarmAsync(ctx workflow.Context, input *cloudwatch.PutMetricAlarmInput) *CloudwatchPutMetricAlarmResult {
	future := workflow.ExecuteActivity(ctx, "aws.cloudwatch.PutMetricAlarm", input)
	return &CloudwatchPutMetricAlarmResult{Result: future}
}

func (a *CloudWatchStub) PutMetricData(ctx workflow.Context, input *cloudwatch.PutMetricDataInput) (*cloudwatch.PutMetricDataOutput, error) {
	var output cloudwatch.PutMetricDataOutput
	err := workflow.ExecuteActivity(ctx, "aws.cloudwatch.PutMetricData", input).Get(ctx, &output)
	return &output, err
}

func (a *CloudWatchStub) PutMetricDataAsync(ctx workflow.Context, input *cloudwatch.PutMetricDataInput) *CloudwatchPutMetricDataResult {
	future := workflow.ExecuteActivity(ctx, "aws.cloudwatch.PutMetricData", input)
	return &CloudwatchPutMetricDataResult{Result: future}
}

func (a *CloudWatchStub) SetAlarmState(ctx workflow.Context, input *cloudwatch.SetAlarmStateInput) (*cloudwatch.SetAlarmStateOutput, error) {
	var output cloudwatch.SetAlarmStateOutput
	err := workflow.ExecuteActivity(ctx, "aws.cloudwatch.SetAlarmState", input).Get(ctx, &output)
	return &output, err
}

func (a *CloudWatchStub) SetAlarmStateAsync(ctx workflow.Context, input *cloudwatch.SetAlarmStateInput) *CloudwatchSetAlarmStateResult {
	future := workflow.ExecuteActivity(ctx, "aws.cloudwatch.SetAlarmState", input)
	return &CloudwatchSetAlarmStateResult{Result: future}
}

func (a *CloudWatchStub) TagResource(ctx workflow.Context, input *cloudwatch.TagResourceInput) (*cloudwatch.TagResourceOutput, error) {
	var output cloudwatch.TagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.cloudwatch.TagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *CloudWatchStub) TagResourceAsync(ctx workflow.Context, input *cloudwatch.TagResourceInput) *CloudwatchTagResourceResult {
	future := workflow.ExecuteActivity(ctx, "aws.cloudwatch.TagResource", input)
	return &CloudwatchTagResourceResult{Result: future}
}

func (a *CloudWatchStub) UntagResource(ctx workflow.Context, input *cloudwatch.UntagResourceInput) (*cloudwatch.UntagResourceOutput, error) {
	var output cloudwatch.UntagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.cloudwatch.UntagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *CloudWatchStub) UntagResourceAsync(ctx workflow.Context, input *cloudwatch.UntagResourceInput) *CloudwatchUntagResourceResult {
	future := workflow.ExecuteActivity(ctx, "aws.cloudwatch.UntagResource", input)
	return &CloudwatchUntagResourceResult{Result: future}
}

func (a *CloudWatchStub) WaitUntilAlarmExists(ctx workflow.Context, input *cloudwatch.DescribeAlarmsInput) error {
	return workflow.ExecuteActivity(ctx, "aws.cloudwatch.WaitUntilAlarmExists", input).Get(ctx, nil)
}

func (a *CloudWatchStub) WaitUntilAlarmExistsAsync(ctx workflow.Context, input *cloudwatch.DescribeAlarmsInput) workflow.Future {
	return workflow.ExecuteActivity(ctx, "aws.cloudwatch.WaitUntilAlarmExists", input)
}

func (a *CloudWatchStub) WaitUntilCompositeAlarmExists(ctx workflow.Context, input *cloudwatch.DescribeAlarmsInput) error {
	return workflow.ExecuteActivity(ctx, "aws.cloudwatch.WaitUntilCompositeAlarmExists", input).Get(ctx, nil)
}

func (a *CloudWatchStub) WaitUntilCompositeAlarmExistsAsync(ctx workflow.Context, input *cloudwatch.DescribeAlarmsInput) workflow.Future {
	return workflow.ExecuteActivity(ctx, "aws.cloudwatch.WaitUntilCompositeAlarmExists", input)
}
