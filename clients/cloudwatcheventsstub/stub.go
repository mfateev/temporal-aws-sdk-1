// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package cloudwatcheventsstub

import (
	"github.com/aws/aws-sdk-go/service/cloudwatchevents"
	"go.temporal.io/sdk/workflow"

	"go.temporal.io/aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type stub struct{}

type CloudWatchEventsActivateEventSourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CloudWatchEventsActivateEventSourceFuture) Get(ctx workflow.Context) (*cloudwatchevents.ActivateEventSourceOutput, error) {
	var output cloudwatchevents.ActivateEventSourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CloudWatchEventsCreateEventBusFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CloudWatchEventsCreateEventBusFuture) Get(ctx workflow.Context) (*cloudwatchevents.CreateEventBusOutput, error) {
	var output cloudwatchevents.CreateEventBusOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CloudWatchEventsCreatePartnerEventSourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CloudWatchEventsCreatePartnerEventSourceFuture) Get(ctx workflow.Context) (*cloudwatchevents.CreatePartnerEventSourceOutput, error) {
	var output cloudwatchevents.CreatePartnerEventSourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CloudWatchEventsDeactivateEventSourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CloudWatchEventsDeactivateEventSourceFuture) Get(ctx workflow.Context) (*cloudwatchevents.DeactivateEventSourceOutput, error) {
	var output cloudwatchevents.DeactivateEventSourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CloudWatchEventsDeleteEventBusFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CloudWatchEventsDeleteEventBusFuture) Get(ctx workflow.Context) (*cloudwatchevents.DeleteEventBusOutput, error) {
	var output cloudwatchevents.DeleteEventBusOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CloudWatchEventsDeletePartnerEventSourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CloudWatchEventsDeletePartnerEventSourceFuture) Get(ctx workflow.Context) (*cloudwatchevents.DeletePartnerEventSourceOutput, error) {
	var output cloudwatchevents.DeletePartnerEventSourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CloudWatchEventsDeleteRuleFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CloudWatchEventsDeleteRuleFuture) Get(ctx workflow.Context) (*cloudwatchevents.DeleteRuleOutput, error) {
	var output cloudwatchevents.DeleteRuleOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CloudWatchEventsDescribeEventBusFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CloudWatchEventsDescribeEventBusFuture) Get(ctx workflow.Context) (*cloudwatchevents.DescribeEventBusOutput, error) {
	var output cloudwatchevents.DescribeEventBusOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CloudWatchEventsDescribeEventSourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CloudWatchEventsDescribeEventSourceFuture) Get(ctx workflow.Context) (*cloudwatchevents.DescribeEventSourceOutput, error) {
	var output cloudwatchevents.DescribeEventSourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CloudWatchEventsDescribePartnerEventSourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CloudWatchEventsDescribePartnerEventSourceFuture) Get(ctx workflow.Context) (*cloudwatchevents.DescribePartnerEventSourceOutput, error) {
	var output cloudwatchevents.DescribePartnerEventSourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CloudWatchEventsDescribeRuleFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CloudWatchEventsDescribeRuleFuture) Get(ctx workflow.Context) (*cloudwatchevents.DescribeRuleOutput, error) {
	var output cloudwatchevents.DescribeRuleOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CloudWatchEventsDisableRuleFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CloudWatchEventsDisableRuleFuture) Get(ctx workflow.Context) (*cloudwatchevents.DisableRuleOutput, error) {
	var output cloudwatchevents.DisableRuleOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CloudWatchEventsEnableRuleFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CloudWatchEventsEnableRuleFuture) Get(ctx workflow.Context) (*cloudwatchevents.EnableRuleOutput, error) {
	var output cloudwatchevents.EnableRuleOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CloudWatchEventsListEventBusesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CloudWatchEventsListEventBusesFuture) Get(ctx workflow.Context) (*cloudwatchevents.ListEventBusesOutput, error) {
	var output cloudwatchevents.ListEventBusesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CloudWatchEventsListEventSourcesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CloudWatchEventsListEventSourcesFuture) Get(ctx workflow.Context) (*cloudwatchevents.ListEventSourcesOutput, error) {
	var output cloudwatchevents.ListEventSourcesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CloudWatchEventsListPartnerEventSourceAccountsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CloudWatchEventsListPartnerEventSourceAccountsFuture) Get(ctx workflow.Context) (*cloudwatchevents.ListPartnerEventSourceAccountsOutput, error) {
	var output cloudwatchevents.ListPartnerEventSourceAccountsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CloudWatchEventsListPartnerEventSourcesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CloudWatchEventsListPartnerEventSourcesFuture) Get(ctx workflow.Context) (*cloudwatchevents.ListPartnerEventSourcesOutput, error) {
	var output cloudwatchevents.ListPartnerEventSourcesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CloudWatchEventsListRuleNamesByTargetFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CloudWatchEventsListRuleNamesByTargetFuture) Get(ctx workflow.Context) (*cloudwatchevents.ListRuleNamesByTargetOutput, error) {
	var output cloudwatchevents.ListRuleNamesByTargetOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CloudWatchEventsListRulesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CloudWatchEventsListRulesFuture) Get(ctx workflow.Context) (*cloudwatchevents.ListRulesOutput, error) {
	var output cloudwatchevents.ListRulesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CloudWatchEventsListTagsForResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CloudWatchEventsListTagsForResourceFuture) Get(ctx workflow.Context) (*cloudwatchevents.ListTagsForResourceOutput, error) {
	var output cloudwatchevents.ListTagsForResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CloudWatchEventsListTargetsByRuleFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CloudWatchEventsListTargetsByRuleFuture) Get(ctx workflow.Context) (*cloudwatchevents.ListTargetsByRuleOutput, error) {
	var output cloudwatchevents.ListTargetsByRuleOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CloudWatchEventsPutEventsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CloudWatchEventsPutEventsFuture) Get(ctx workflow.Context) (*cloudwatchevents.PutEventsOutput, error) {
	var output cloudwatchevents.PutEventsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CloudWatchEventsPutPartnerEventsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CloudWatchEventsPutPartnerEventsFuture) Get(ctx workflow.Context) (*cloudwatchevents.PutPartnerEventsOutput, error) {
	var output cloudwatchevents.PutPartnerEventsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CloudWatchEventsPutPermissionFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CloudWatchEventsPutPermissionFuture) Get(ctx workflow.Context) (*cloudwatchevents.PutPermissionOutput, error) {
	var output cloudwatchevents.PutPermissionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CloudWatchEventsPutRuleFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CloudWatchEventsPutRuleFuture) Get(ctx workflow.Context) (*cloudwatchevents.PutRuleOutput, error) {
	var output cloudwatchevents.PutRuleOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CloudWatchEventsPutTargetsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CloudWatchEventsPutTargetsFuture) Get(ctx workflow.Context) (*cloudwatchevents.PutTargetsOutput, error) {
	var output cloudwatchevents.PutTargetsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CloudWatchEventsRemovePermissionFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CloudWatchEventsRemovePermissionFuture) Get(ctx workflow.Context) (*cloudwatchevents.RemovePermissionOutput, error) {
	var output cloudwatchevents.RemovePermissionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CloudWatchEventsRemoveTargetsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CloudWatchEventsRemoveTargetsFuture) Get(ctx workflow.Context) (*cloudwatchevents.RemoveTargetsOutput, error) {
	var output cloudwatchevents.RemoveTargetsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CloudWatchEventsTagResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CloudWatchEventsTagResourceFuture) Get(ctx workflow.Context) (*cloudwatchevents.TagResourceOutput, error) {
	var output cloudwatchevents.TagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CloudWatchEventsTestEventPatternFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CloudWatchEventsTestEventPatternFuture) Get(ctx workflow.Context) (*cloudwatchevents.TestEventPatternOutput, error) {
	var output cloudwatchevents.TestEventPatternOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CloudWatchEventsUntagResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CloudWatchEventsUntagResourceFuture) Get(ctx workflow.Context) (*cloudwatchevents.UntagResourceOutput, error) {
	var output cloudwatchevents.UntagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *stub) ActivateEventSource(ctx workflow.Context, input *cloudwatchevents.ActivateEventSourceInput) (*cloudwatchevents.ActivateEventSourceOutput, error) {
	var output cloudwatchevents.ActivateEventSourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.cloudwatchevents.ActivateEventSource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ActivateEventSourceAsync(ctx workflow.Context, input *cloudwatchevents.ActivateEventSourceInput) *CloudWatchEventsActivateEventSourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cloudwatchevents.ActivateEventSource", input)
	return &CloudWatchEventsActivateEventSourceFuture{Future: future}
}

func (a *stub) CreateEventBus(ctx workflow.Context, input *cloudwatchevents.CreateEventBusInput) (*cloudwatchevents.CreateEventBusOutput, error) {
	var output cloudwatchevents.CreateEventBusOutput
	err := workflow.ExecuteActivity(ctx, "aws.cloudwatchevents.CreateEventBus", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateEventBusAsync(ctx workflow.Context, input *cloudwatchevents.CreateEventBusInput) *CloudWatchEventsCreateEventBusFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cloudwatchevents.CreateEventBus", input)
	return &CloudWatchEventsCreateEventBusFuture{Future: future}
}

func (a *stub) CreatePartnerEventSource(ctx workflow.Context, input *cloudwatchevents.CreatePartnerEventSourceInput) (*cloudwatchevents.CreatePartnerEventSourceOutput, error) {
	var output cloudwatchevents.CreatePartnerEventSourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.cloudwatchevents.CreatePartnerEventSource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreatePartnerEventSourceAsync(ctx workflow.Context, input *cloudwatchevents.CreatePartnerEventSourceInput) *CloudWatchEventsCreatePartnerEventSourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cloudwatchevents.CreatePartnerEventSource", input)
	return &CloudWatchEventsCreatePartnerEventSourceFuture{Future: future}
}

func (a *stub) DeactivateEventSource(ctx workflow.Context, input *cloudwatchevents.DeactivateEventSourceInput) (*cloudwatchevents.DeactivateEventSourceOutput, error) {
	var output cloudwatchevents.DeactivateEventSourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.cloudwatchevents.DeactivateEventSource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeactivateEventSourceAsync(ctx workflow.Context, input *cloudwatchevents.DeactivateEventSourceInput) *CloudWatchEventsDeactivateEventSourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cloudwatchevents.DeactivateEventSource", input)
	return &CloudWatchEventsDeactivateEventSourceFuture{Future: future}
}

func (a *stub) DeleteEventBus(ctx workflow.Context, input *cloudwatchevents.DeleteEventBusInput) (*cloudwatchevents.DeleteEventBusOutput, error) {
	var output cloudwatchevents.DeleteEventBusOutput
	err := workflow.ExecuteActivity(ctx, "aws.cloudwatchevents.DeleteEventBus", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteEventBusAsync(ctx workflow.Context, input *cloudwatchevents.DeleteEventBusInput) *CloudWatchEventsDeleteEventBusFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cloudwatchevents.DeleteEventBus", input)
	return &CloudWatchEventsDeleteEventBusFuture{Future: future}
}

func (a *stub) DeletePartnerEventSource(ctx workflow.Context, input *cloudwatchevents.DeletePartnerEventSourceInput) (*cloudwatchevents.DeletePartnerEventSourceOutput, error) {
	var output cloudwatchevents.DeletePartnerEventSourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.cloudwatchevents.DeletePartnerEventSource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeletePartnerEventSourceAsync(ctx workflow.Context, input *cloudwatchevents.DeletePartnerEventSourceInput) *CloudWatchEventsDeletePartnerEventSourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cloudwatchevents.DeletePartnerEventSource", input)
	return &CloudWatchEventsDeletePartnerEventSourceFuture{Future: future}
}

func (a *stub) DeleteRule(ctx workflow.Context, input *cloudwatchevents.DeleteRuleInput) (*cloudwatchevents.DeleteRuleOutput, error) {
	var output cloudwatchevents.DeleteRuleOutput
	err := workflow.ExecuteActivity(ctx, "aws.cloudwatchevents.DeleteRule", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteRuleAsync(ctx workflow.Context, input *cloudwatchevents.DeleteRuleInput) *CloudWatchEventsDeleteRuleFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cloudwatchevents.DeleteRule", input)
	return &CloudWatchEventsDeleteRuleFuture{Future: future}
}

func (a *stub) DescribeEventBus(ctx workflow.Context, input *cloudwatchevents.DescribeEventBusInput) (*cloudwatchevents.DescribeEventBusOutput, error) {
	var output cloudwatchevents.DescribeEventBusOutput
	err := workflow.ExecuteActivity(ctx, "aws.cloudwatchevents.DescribeEventBus", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeEventBusAsync(ctx workflow.Context, input *cloudwatchevents.DescribeEventBusInput) *CloudWatchEventsDescribeEventBusFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cloudwatchevents.DescribeEventBus", input)
	return &CloudWatchEventsDescribeEventBusFuture{Future: future}
}

func (a *stub) DescribeEventSource(ctx workflow.Context, input *cloudwatchevents.DescribeEventSourceInput) (*cloudwatchevents.DescribeEventSourceOutput, error) {
	var output cloudwatchevents.DescribeEventSourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.cloudwatchevents.DescribeEventSource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeEventSourceAsync(ctx workflow.Context, input *cloudwatchevents.DescribeEventSourceInput) *CloudWatchEventsDescribeEventSourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cloudwatchevents.DescribeEventSource", input)
	return &CloudWatchEventsDescribeEventSourceFuture{Future: future}
}

func (a *stub) DescribePartnerEventSource(ctx workflow.Context, input *cloudwatchevents.DescribePartnerEventSourceInput) (*cloudwatchevents.DescribePartnerEventSourceOutput, error) {
	var output cloudwatchevents.DescribePartnerEventSourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.cloudwatchevents.DescribePartnerEventSource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribePartnerEventSourceAsync(ctx workflow.Context, input *cloudwatchevents.DescribePartnerEventSourceInput) *CloudWatchEventsDescribePartnerEventSourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cloudwatchevents.DescribePartnerEventSource", input)
	return &CloudWatchEventsDescribePartnerEventSourceFuture{Future: future}
}

func (a *stub) DescribeRule(ctx workflow.Context, input *cloudwatchevents.DescribeRuleInput) (*cloudwatchevents.DescribeRuleOutput, error) {
	var output cloudwatchevents.DescribeRuleOutput
	err := workflow.ExecuteActivity(ctx, "aws.cloudwatchevents.DescribeRule", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeRuleAsync(ctx workflow.Context, input *cloudwatchevents.DescribeRuleInput) *CloudWatchEventsDescribeRuleFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cloudwatchevents.DescribeRule", input)
	return &CloudWatchEventsDescribeRuleFuture{Future: future}
}

func (a *stub) DisableRule(ctx workflow.Context, input *cloudwatchevents.DisableRuleInput) (*cloudwatchevents.DisableRuleOutput, error) {
	var output cloudwatchevents.DisableRuleOutput
	err := workflow.ExecuteActivity(ctx, "aws.cloudwatchevents.DisableRule", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DisableRuleAsync(ctx workflow.Context, input *cloudwatchevents.DisableRuleInput) *CloudWatchEventsDisableRuleFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cloudwatchevents.DisableRule", input)
	return &CloudWatchEventsDisableRuleFuture{Future: future}
}

func (a *stub) EnableRule(ctx workflow.Context, input *cloudwatchevents.EnableRuleInput) (*cloudwatchevents.EnableRuleOutput, error) {
	var output cloudwatchevents.EnableRuleOutput
	err := workflow.ExecuteActivity(ctx, "aws.cloudwatchevents.EnableRule", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) EnableRuleAsync(ctx workflow.Context, input *cloudwatchevents.EnableRuleInput) *CloudWatchEventsEnableRuleFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cloudwatchevents.EnableRule", input)
	return &CloudWatchEventsEnableRuleFuture{Future: future}
}

func (a *stub) ListEventBuses(ctx workflow.Context, input *cloudwatchevents.ListEventBusesInput) (*cloudwatchevents.ListEventBusesOutput, error) {
	var output cloudwatchevents.ListEventBusesOutput
	err := workflow.ExecuteActivity(ctx, "aws.cloudwatchevents.ListEventBuses", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListEventBusesAsync(ctx workflow.Context, input *cloudwatchevents.ListEventBusesInput) *CloudWatchEventsListEventBusesFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cloudwatchevents.ListEventBuses", input)
	return &CloudWatchEventsListEventBusesFuture{Future: future}
}

func (a *stub) ListEventSources(ctx workflow.Context, input *cloudwatchevents.ListEventSourcesInput) (*cloudwatchevents.ListEventSourcesOutput, error) {
	var output cloudwatchevents.ListEventSourcesOutput
	err := workflow.ExecuteActivity(ctx, "aws.cloudwatchevents.ListEventSources", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListEventSourcesAsync(ctx workflow.Context, input *cloudwatchevents.ListEventSourcesInput) *CloudWatchEventsListEventSourcesFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cloudwatchevents.ListEventSources", input)
	return &CloudWatchEventsListEventSourcesFuture{Future: future}
}

func (a *stub) ListPartnerEventSourceAccounts(ctx workflow.Context, input *cloudwatchevents.ListPartnerEventSourceAccountsInput) (*cloudwatchevents.ListPartnerEventSourceAccountsOutput, error) {
	var output cloudwatchevents.ListPartnerEventSourceAccountsOutput
	err := workflow.ExecuteActivity(ctx, "aws.cloudwatchevents.ListPartnerEventSourceAccounts", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListPartnerEventSourceAccountsAsync(ctx workflow.Context, input *cloudwatchevents.ListPartnerEventSourceAccountsInput) *CloudWatchEventsListPartnerEventSourceAccountsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cloudwatchevents.ListPartnerEventSourceAccounts", input)
	return &CloudWatchEventsListPartnerEventSourceAccountsFuture{Future: future}
}

func (a *stub) ListPartnerEventSources(ctx workflow.Context, input *cloudwatchevents.ListPartnerEventSourcesInput) (*cloudwatchevents.ListPartnerEventSourcesOutput, error) {
	var output cloudwatchevents.ListPartnerEventSourcesOutput
	err := workflow.ExecuteActivity(ctx, "aws.cloudwatchevents.ListPartnerEventSources", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListPartnerEventSourcesAsync(ctx workflow.Context, input *cloudwatchevents.ListPartnerEventSourcesInput) *CloudWatchEventsListPartnerEventSourcesFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cloudwatchevents.ListPartnerEventSources", input)
	return &CloudWatchEventsListPartnerEventSourcesFuture{Future: future}
}

func (a *stub) ListRuleNamesByTarget(ctx workflow.Context, input *cloudwatchevents.ListRuleNamesByTargetInput) (*cloudwatchevents.ListRuleNamesByTargetOutput, error) {
	var output cloudwatchevents.ListRuleNamesByTargetOutput
	err := workflow.ExecuteActivity(ctx, "aws.cloudwatchevents.ListRuleNamesByTarget", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListRuleNamesByTargetAsync(ctx workflow.Context, input *cloudwatchevents.ListRuleNamesByTargetInput) *CloudWatchEventsListRuleNamesByTargetFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cloudwatchevents.ListRuleNamesByTarget", input)
	return &CloudWatchEventsListRuleNamesByTargetFuture{Future: future}
}

func (a *stub) ListRules(ctx workflow.Context, input *cloudwatchevents.ListRulesInput) (*cloudwatchevents.ListRulesOutput, error) {
	var output cloudwatchevents.ListRulesOutput
	err := workflow.ExecuteActivity(ctx, "aws.cloudwatchevents.ListRules", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListRulesAsync(ctx workflow.Context, input *cloudwatchevents.ListRulesInput) *CloudWatchEventsListRulesFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cloudwatchevents.ListRules", input)
	return &CloudWatchEventsListRulesFuture{Future: future}
}

func (a *stub) ListTagsForResource(ctx workflow.Context, input *cloudwatchevents.ListTagsForResourceInput) (*cloudwatchevents.ListTagsForResourceOutput, error) {
	var output cloudwatchevents.ListTagsForResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.cloudwatchevents.ListTagsForResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListTagsForResourceAsync(ctx workflow.Context, input *cloudwatchevents.ListTagsForResourceInput) *CloudWatchEventsListTagsForResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cloudwatchevents.ListTagsForResource", input)
	return &CloudWatchEventsListTagsForResourceFuture{Future: future}
}

func (a *stub) ListTargetsByRule(ctx workflow.Context, input *cloudwatchevents.ListTargetsByRuleInput) (*cloudwatchevents.ListTargetsByRuleOutput, error) {
	var output cloudwatchevents.ListTargetsByRuleOutput
	err := workflow.ExecuteActivity(ctx, "aws.cloudwatchevents.ListTargetsByRule", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListTargetsByRuleAsync(ctx workflow.Context, input *cloudwatchevents.ListTargetsByRuleInput) *CloudWatchEventsListTargetsByRuleFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cloudwatchevents.ListTargetsByRule", input)
	return &CloudWatchEventsListTargetsByRuleFuture{Future: future}
}

func (a *stub) PutEvents(ctx workflow.Context, input *cloudwatchevents.PutEventsInput) (*cloudwatchevents.PutEventsOutput, error) {
	var output cloudwatchevents.PutEventsOutput
	err := workflow.ExecuteActivity(ctx, "aws.cloudwatchevents.PutEvents", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) PutEventsAsync(ctx workflow.Context, input *cloudwatchevents.PutEventsInput) *CloudWatchEventsPutEventsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cloudwatchevents.PutEvents", input)
	return &CloudWatchEventsPutEventsFuture{Future: future}
}

func (a *stub) PutPartnerEvents(ctx workflow.Context, input *cloudwatchevents.PutPartnerEventsInput) (*cloudwatchevents.PutPartnerEventsOutput, error) {
	var output cloudwatchevents.PutPartnerEventsOutput
	err := workflow.ExecuteActivity(ctx, "aws.cloudwatchevents.PutPartnerEvents", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) PutPartnerEventsAsync(ctx workflow.Context, input *cloudwatchevents.PutPartnerEventsInput) *CloudWatchEventsPutPartnerEventsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cloudwatchevents.PutPartnerEvents", input)
	return &CloudWatchEventsPutPartnerEventsFuture{Future: future}
}

func (a *stub) PutPermission(ctx workflow.Context, input *cloudwatchevents.PutPermissionInput) (*cloudwatchevents.PutPermissionOutput, error) {
	var output cloudwatchevents.PutPermissionOutput
	err := workflow.ExecuteActivity(ctx, "aws.cloudwatchevents.PutPermission", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) PutPermissionAsync(ctx workflow.Context, input *cloudwatchevents.PutPermissionInput) *CloudWatchEventsPutPermissionFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cloudwatchevents.PutPermission", input)
	return &CloudWatchEventsPutPermissionFuture{Future: future}
}

func (a *stub) PutRule(ctx workflow.Context, input *cloudwatchevents.PutRuleInput) (*cloudwatchevents.PutRuleOutput, error) {
	var output cloudwatchevents.PutRuleOutput
	err := workflow.ExecuteActivity(ctx, "aws.cloudwatchevents.PutRule", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) PutRuleAsync(ctx workflow.Context, input *cloudwatchevents.PutRuleInput) *CloudWatchEventsPutRuleFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cloudwatchevents.PutRule", input)
	return &CloudWatchEventsPutRuleFuture{Future: future}
}

func (a *stub) PutTargets(ctx workflow.Context, input *cloudwatchevents.PutTargetsInput) (*cloudwatchevents.PutTargetsOutput, error) {
	var output cloudwatchevents.PutTargetsOutput
	err := workflow.ExecuteActivity(ctx, "aws.cloudwatchevents.PutTargets", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) PutTargetsAsync(ctx workflow.Context, input *cloudwatchevents.PutTargetsInput) *CloudWatchEventsPutTargetsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cloudwatchevents.PutTargets", input)
	return &CloudWatchEventsPutTargetsFuture{Future: future}
}

func (a *stub) RemovePermission(ctx workflow.Context, input *cloudwatchevents.RemovePermissionInput) (*cloudwatchevents.RemovePermissionOutput, error) {
	var output cloudwatchevents.RemovePermissionOutput
	err := workflow.ExecuteActivity(ctx, "aws.cloudwatchevents.RemovePermission", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) RemovePermissionAsync(ctx workflow.Context, input *cloudwatchevents.RemovePermissionInput) *CloudWatchEventsRemovePermissionFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cloudwatchevents.RemovePermission", input)
	return &CloudWatchEventsRemovePermissionFuture{Future: future}
}

func (a *stub) RemoveTargets(ctx workflow.Context, input *cloudwatchevents.RemoveTargetsInput) (*cloudwatchevents.RemoveTargetsOutput, error) {
	var output cloudwatchevents.RemoveTargetsOutput
	err := workflow.ExecuteActivity(ctx, "aws.cloudwatchevents.RemoveTargets", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) RemoveTargetsAsync(ctx workflow.Context, input *cloudwatchevents.RemoveTargetsInput) *CloudWatchEventsRemoveTargetsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cloudwatchevents.RemoveTargets", input)
	return &CloudWatchEventsRemoveTargetsFuture{Future: future}
}

func (a *stub) TagResource(ctx workflow.Context, input *cloudwatchevents.TagResourceInput) (*cloudwatchevents.TagResourceOutput, error) {
	var output cloudwatchevents.TagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.cloudwatchevents.TagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) TagResourceAsync(ctx workflow.Context, input *cloudwatchevents.TagResourceInput) *CloudWatchEventsTagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cloudwatchevents.TagResource", input)
	return &CloudWatchEventsTagResourceFuture{Future: future}
}

func (a *stub) TestEventPattern(ctx workflow.Context, input *cloudwatchevents.TestEventPatternInput) (*cloudwatchevents.TestEventPatternOutput, error) {
	var output cloudwatchevents.TestEventPatternOutput
	err := workflow.ExecuteActivity(ctx, "aws.cloudwatchevents.TestEventPattern", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) TestEventPatternAsync(ctx workflow.Context, input *cloudwatchevents.TestEventPatternInput) *CloudWatchEventsTestEventPatternFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cloudwatchevents.TestEventPattern", input)
	return &CloudWatchEventsTestEventPatternFuture{Future: future}
}

func (a *stub) UntagResource(ctx workflow.Context, input *cloudwatchevents.UntagResourceInput) (*cloudwatchevents.UntagResourceOutput, error) {
	var output cloudwatchevents.UntagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.cloudwatchevents.UntagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UntagResourceAsync(ctx workflow.Context, input *cloudwatchevents.UntagResourceInput) *CloudWatchEventsUntagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cloudwatchevents.UntagResource", input)
	return &CloudWatchEventsUntagResourceFuture{Future: future}
}
