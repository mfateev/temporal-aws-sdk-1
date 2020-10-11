// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package awsclients

import (
	"github.com/aws/aws-sdk-go/service/iot1clickdevicesservice"
	"go.temporal.io/sdk/workflow"
)

type IoT1ClickDevicesServiceClient interface {
	ClaimDevicesByClaimCode(ctx workflow.Context, input *iot1clickdevicesservice.ClaimDevicesByClaimCodeInput) (*iot1clickdevicesservice.ClaimDevicesByClaimCodeOutput, error)
	ClaimDevicesByClaimCodeAsync(ctx workflow.Context, input *iot1clickdevicesservice.ClaimDevicesByClaimCodeInput) *IoT1ClickDevicesServiceClaimDevicesByClaimCodeFuture

	DescribeDevice(ctx workflow.Context, input *iot1clickdevicesservice.DescribeDeviceInput) (*iot1clickdevicesservice.DescribeDeviceOutput, error)
	DescribeDeviceAsync(ctx workflow.Context, input *iot1clickdevicesservice.DescribeDeviceInput) *IoT1ClickDevicesServiceDescribeDeviceFuture

	FinalizeDeviceClaim(ctx workflow.Context, input *iot1clickdevicesservice.FinalizeDeviceClaimInput) (*iot1clickdevicesservice.FinalizeDeviceClaimOutput, error)
	FinalizeDeviceClaimAsync(ctx workflow.Context, input *iot1clickdevicesservice.FinalizeDeviceClaimInput) *IoT1ClickDevicesServiceFinalizeDeviceClaimFuture

	GetDeviceMethods(ctx workflow.Context, input *iot1clickdevicesservice.GetDeviceMethodsInput) (*iot1clickdevicesservice.GetDeviceMethodsOutput, error)
	GetDeviceMethodsAsync(ctx workflow.Context, input *iot1clickdevicesservice.GetDeviceMethodsInput) *IoT1ClickDevicesServiceGetDeviceMethodsFuture

	InitiateDeviceClaim(ctx workflow.Context, input *iot1clickdevicesservice.InitiateDeviceClaimInput) (*iot1clickdevicesservice.InitiateDeviceClaimOutput, error)
	InitiateDeviceClaimAsync(ctx workflow.Context, input *iot1clickdevicesservice.InitiateDeviceClaimInput) *IoT1ClickDevicesServiceInitiateDeviceClaimFuture

	InvokeDeviceMethod(ctx workflow.Context, input *iot1clickdevicesservice.InvokeDeviceMethodInput) (*iot1clickdevicesservice.InvokeDeviceMethodOutput, error)
	InvokeDeviceMethodAsync(ctx workflow.Context, input *iot1clickdevicesservice.InvokeDeviceMethodInput) *IoT1ClickDevicesServiceInvokeDeviceMethodFuture

	ListDeviceEvents(ctx workflow.Context, input *iot1clickdevicesservice.ListDeviceEventsInput) (*iot1clickdevicesservice.ListDeviceEventsOutput, error)
	ListDeviceEventsAsync(ctx workflow.Context, input *iot1clickdevicesservice.ListDeviceEventsInput) *IoT1ClickDevicesServiceListDeviceEventsFuture

	ListDevices(ctx workflow.Context, input *iot1clickdevicesservice.ListDevicesInput) (*iot1clickdevicesservice.ListDevicesOutput, error)
	ListDevicesAsync(ctx workflow.Context, input *iot1clickdevicesservice.ListDevicesInput) *IoT1ClickDevicesServiceListDevicesFuture

	ListTagsForResource(ctx workflow.Context, input *iot1clickdevicesservice.ListTagsForResourceInput) (*iot1clickdevicesservice.ListTagsForResourceOutput, error)
	ListTagsForResourceAsync(ctx workflow.Context, input *iot1clickdevicesservice.ListTagsForResourceInput) *IoT1ClickDevicesServiceListTagsForResourceFuture

	TagResource(ctx workflow.Context, input *iot1clickdevicesservice.TagResourceInput) (*iot1clickdevicesservice.TagResourceOutput, error)
	TagResourceAsync(ctx workflow.Context, input *iot1clickdevicesservice.TagResourceInput) *IoT1ClickDevicesServiceTagResourceFuture

	UnclaimDevice(ctx workflow.Context, input *iot1clickdevicesservice.UnclaimDeviceInput) (*iot1clickdevicesservice.UnclaimDeviceOutput, error)
	UnclaimDeviceAsync(ctx workflow.Context, input *iot1clickdevicesservice.UnclaimDeviceInput) *IoT1ClickDevicesServiceUnclaimDeviceFuture

	UntagResource(ctx workflow.Context, input *iot1clickdevicesservice.UntagResourceInput) (*iot1clickdevicesservice.UntagResourceOutput, error)
	UntagResourceAsync(ctx workflow.Context, input *iot1clickdevicesservice.UntagResourceInput) *IoT1ClickDevicesServiceUntagResourceFuture

	UpdateDeviceState(ctx workflow.Context, input *iot1clickdevicesservice.UpdateDeviceStateInput) (*iot1clickdevicesservice.UpdateDeviceStateOutput, error)
	UpdateDeviceStateAsync(ctx workflow.Context, input *iot1clickdevicesservice.UpdateDeviceStateInput) *IoT1ClickDevicesServiceUpdateDeviceStateFuture
}

type IoT1ClickDevicesServiceStub struct{}

func NewIoT1ClickDevicesServiceStub() IoT1ClickDevicesServiceClient {
	return &IoT1ClickDevicesServiceStub{}
}

type IoT1ClickDevicesServiceClaimDevicesByClaimCodeFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *IoT1ClickDevicesServiceClaimDevicesByClaimCodeFuture) Get(ctx workflow.Context) (*iot1clickdevicesservice.ClaimDevicesByClaimCodeOutput, error) {
	var output iot1clickdevicesservice.ClaimDevicesByClaimCodeOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type IoT1ClickDevicesServiceDescribeDeviceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *IoT1ClickDevicesServiceDescribeDeviceFuture) Get(ctx workflow.Context) (*iot1clickdevicesservice.DescribeDeviceOutput, error) {
	var output iot1clickdevicesservice.DescribeDeviceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type IoT1ClickDevicesServiceFinalizeDeviceClaimFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *IoT1ClickDevicesServiceFinalizeDeviceClaimFuture) Get(ctx workflow.Context) (*iot1clickdevicesservice.FinalizeDeviceClaimOutput, error) {
	var output iot1clickdevicesservice.FinalizeDeviceClaimOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type IoT1ClickDevicesServiceGetDeviceMethodsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *IoT1ClickDevicesServiceGetDeviceMethodsFuture) Get(ctx workflow.Context) (*iot1clickdevicesservice.GetDeviceMethodsOutput, error) {
	var output iot1clickdevicesservice.GetDeviceMethodsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type IoT1ClickDevicesServiceInitiateDeviceClaimFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *IoT1ClickDevicesServiceInitiateDeviceClaimFuture) Get(ctx workflow.Context) (*iot1clickdevicesservice.InitiateDeviceClaimOutput, error) {
	var output iot1clickdevicesservice.InitiateDeviceClaimOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type IoT1ClickDevicesServiceInvokeDeviceMethodFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *IoT1ClickDevicesServiceInvokeDeviceMethodFuture) Get(ctx workflow.Context) (*iot1clickdevicesservice.InvokeDeviceMethodOutput, error) {
	var output iot1clickdevicesservice.InvokeDeviceMethodOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type IoT1ClickDevicesServiceListDeviceEventsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *IoT1ClickDevicesServiceListDeviceEventsFuture) Get(ctx workflow.Context) (*iot1clickdevicesservice.ListDeviceEventsOutput, error) {
	var output iot1clickdevicesservice.ListDeviceEventsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type IoT1ClickDevicesServiceListDevicesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *IoT1ClickDevicesServiceListDevicesFuture) Get(ctx workflow.Context) (*iot1clickdevicesservice.ListDevicesOutput, error) {
	var output iot1clickdevicesservice.ListDevicesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type IoT1ClickDevicesServiceListTagsForResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *IoT1ClickDevicesServiceListTagsForResourceFuture) Get(ctx workflow.Context) (*iot1clickdevicesservice.ListTagsForResourceOutput, error) {
	var output iot1clickdevicesservice.ListTagsForResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type IoT1ClickDevicesServiceTagResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *IoT1ClickDevicesServiceTagResourceFuture) Get(ctx workflow.Context) (*iot1clickdevicesservice.TagResourceOutput, error) {
	var output iot1clickdevicesservice.TagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type IoT1ClickDevicesServiceUnclaimDeviceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *IoT1ClickDevicesServiceUnclaimDeviceFuture) Get(ctx workflow.Context) (*iot1clickdevicesservice.UnclaimDeviceOutput, error) {
	var output iot1clickdevicesservice.UnclaimDeviceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type IoT1ClickDevicesServiceUntagResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *IoT1ClickDevicesServiceUntagResourceFuture) Get(ctx workflow.Context) (*iot1clickdevicesservice.UntagResourceOutput, error) {
	var output iot1clickdevicesservice.UntagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type IoT1ClickDevicesServiceUpdateDeviceStateFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *IoT1ClickDevicesServiceUpdateDeviceStateFuture) Get(ctx workflow.Context) (*iot1clickdevicesservice.UpdateDeviceStateOutput, error) {
	var output iot1clickdevicesservice.UpdateDeviceStateOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *IoT1ClickDevicesServiceStub) ClaimDevicesByClaimCode(ctx workflow.Context, input *iot1clickdevicesservice.ClaimDevicesByClaimCodeInput) (*iot1clickdevicesservice.ClaimDevicesByClaimCodeOutput, error) {
	var output iot1clickdevicesservice.ClaimDevicesByClaimCodeOutput
	err := workflow.ExecuteActivity(ctx, "aws.iot1clickdevicesservice.ClaimDevicesByClaimCode", input).Get(ctx, &output)
	return &output, err
}

func (a *IoT1ClickDevicesServiceStub) ClaimDevicesByClaimCodeAsync(ctx workflow.Context, input *iot1clickdevicesservice.ClaimDevicesByClaimCodeInput) *IoT1ClickDevicesServiceClaimDevicesByClaimCodeFuture {
	future := workflow.ExecuteActivity(ctx, "aws.iot1clickdevicesservice.ClaimDevicesByClaimCode", input)
	return &IoT1ClickDevicesServiceClaimDevicesByClaimCodeFuture{Future: future}
}

func (a *IoT1ClickDevicesServiceStub) DescribeDevice(ctx workflow.Context, input *iot1clickdevicesservice.DescribeDeviceInput) (*iot1clickdevicesservice.DescribeDeviceOutput, error) {
	var output iot1clickdevicesservice.DescribeDeviceOutput
	err := workflow.ExecuteActivity(ctx, "aws.iot1clickdevicesservice.DescribeDevice", input).Get(ctx, &output)
	return &output, err
}

func (a *IoT1ClickDevicesServiceStub) DescribeDeviceAsync(ctx workflow.Context, input *iot1clickdevicesservice.DescribeDeviceInput) *IoT1ClickDevicesServiceDescribeDeviceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.iot1clickdevicesservice.DescribeDevice", input)
	return &IoT1ClickDevicesServiceDescribeDeviceFuture{Future: future}
}

func (a *IoT1ClickDevicesServiceStub) FinalizeDeviceClaim(ctx workflow.Context, input *iot1clickdevicesservice.FinalizeDeviceClaimInput) (*iot1clickdevicesservice.FinalizeDeviceClaimOutput, error) {
	var output iot1clickdevicesservice.FinalizeDeviceClaimOutput
	err := workflow.ExecuteActivity(ctx, "aws.iot1clickdevicesservice.FinalizeDeviceClaim", input).Get(ctx, &output)
	return &output, err
}

func (a *IoT1ClickDevicesServiceStub) FinalizeDeviceClaimAsync(ctx workflow.Context, input *iot1clickdevicesservice.FinalizeDeviceClaimInput) *IoT1ClickDevicesServiceFinalizeDeviceClaimFuture {
	future := workflow.ExecuteActivity(ctx, "aws.iot1clickdevicesservice.FinalizeDeviceClaim", input)
	return &IoT1ClickDevicesServiceFinalizeDeviceClaimFuture{Future: future}
}

func (a *IoT1ClickDevicesServiceStub) GetDeviceMethods(ctx workflow.Context, input *iot1clickdevicesservice.GetDeviceMethodsInput) (*iot1clickdevicesservice.GetDeviceMethodsOutput, error) {
	var output iot1clickdevicesservice.GetDeviceMethodsOutput
	err := workflow.ExecuteActivity(ctx, "aws.iot1clickdevicesservice.GetDeviceMethods", input).Get(ctx, &output)
	return &output, err
}

func (a *IoT1ClickDevicesServiceStub) GetDeviceMethodsAsync(ctx workflow.Context, input *iot1clickdevicesservice.GetDeviceMethodsInput) *IoT1ClickDevicesServiceGetDeviceMethodsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.iot1clickdevicesservice.GetDeviceMethods", input)
	return &IoT1ClickDevicesServiceGetDeviceMethodsFuture{Future: future}
}

func (a *IoT1ClickDevicesServiceStub) InitiateDeviceClaim(ctx workflow.Context, input *iot1clickdevicesservice.InitiateDeviceClaimInput) (*iot1clickdevicesservice.InitiateDeviceClaimOutput, error) {
	var output iot1clickdevicesservice.InitiateDeviceClaimOutput
	err := workflow.ExecuteActivity(ctx, "aws.iot1clickdevicesservice.InitiateDeviceClaim", input).Get(ctx, &output)
	return &output, err
}

func (a *IoT1ClickDevicesServiceStub) InitiateDeviceClaimAsync(ctx workflow.Context, input *iot1clickdevicesservice.InitiateDeviceClaimInput) *IoT1ClickDevicesServiceInitiateDeviceClaimFuture {
	future := workflow.ExecuteActivity(ctx, "aws.iot1clickdevicesservice.InitiateDeviceClaim", input)
	return &IoT1ClickDevicesServiceInitiateDeviceClaimFuture{Future: future}
}

func (a *IoT1ClickDevicesServiceStub) InvokeDeviceMethod(ctx workflow.Context, input *iot1clickdevicesservice.InvokeDeviceMethodInput) (*iot1clickdevicesservice.InvokeDeviceMethodOutput, error) {
	var output iot1clickdevicesservice.InvokeDeviceMethodOutput
	err := workflow.ExecuteActivity(ctx, "aws.iot1clickdevicesservice.InvokeDeviceMethod", input).Get(ctx, &output)
	return &output, err
}

func (a *IoT1ClickDevicesServiceStub) InvokeDeviceMethodAsync(ctx workflow.Context, input *iot1clickdevicesservice.InvokeDeviceMethodInput) *IoT1ClickDevicesServiceInvokeDeviceMethodFuture {
	future := workflow.ExecuteActivity(ctx, "aws.iot1clickdevicesservice.InvokeDeviceMethod", input)
	return &IoT1ClickDevicesServiceInvokeDeviceMethodFuture{Future: future}
}

func (a *IoT1ClickDevicesServiceStub) ListDeviceEvents(ctx workflow.Context, input *iot1clickdevicesservice.ListDeviceEventsInput) (*iot1clickdevicesservice.ListDeviceEventsOutput, error) {
	var output iot1clickdevicesservice.ListDeviceEventsOutput
	err := workflow.ExecuteActivity(ctx, "aws.iot1clickdevicesservice.ListDeviceEvents", input).Get(ctx, &output)
	return &output, err
}

func (a *IoT1ClickDevicesServiceStub) ListDeviceEventsAsync(ctx workflow.Context, input *iot1clickdevicesservice.ListDeviceEventsInput) *IoT1ClickDevicesServiceListDeviceEventsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.iot1clickdevicesservice.ListDeviceEvents", input)
	return &IoT1ClickDevicesServiceListDeviceEventsFuture{Future: future}
}

func (a *IoT1ClickDevicesServiceStub) ListDevices(ctx workflow.Context, input *iot1clickdevicesservice.ListDevicesInput) (*iot1clickdevicesservice.ListDevicesOutput, error) {
	var output iot1clickdevicesservice.ListDevicesOutput
	err := workflow.ExecuteActivity(ctx, "aws.iot1clickdevicesservice.ListDevices", input).Get(ctx, &output)
	return &output, err
}

func (a *IoT1ClickDevicesServiceStub) ListDevicesAsync(ctx workflow.Context, input *iot1clickdevicesservice.ListDevicesInput) *IoT1ClickDevicesServiceListDevicesFuture {
	future := workflow.ExecuteActivity(ctx, "aws.iot1clickdevicesservice.ListDevices", input)
	return &IoT1ClickDevicesServiceListDevicesFuture{Future: future}
}

func (a *IoT1ClickDevicesServiceStub) ListTagsForResource(ctx workflow.Context, input *iot1clickdevicesservice.ListTagsForResourceInput) (*iot1clickdevicesservice.ListTagsForResourceOutput, error) {
	var output iot1clickdevicesservice.ListTagsForResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.iot1clickdevicesservice.ListTagsForResource", input).Get(ctx, &output)
	return &output, err
}

func (a *IoT1ClickDevicesServiceStub) ListTagsForResourceAsync(ctx workflow.Context, input *iot1clickdevicesservice.ListTagsForResourceInput) *IoT1ClickDevicesServiceListTagsForResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.iot1clickdevicesservice.ListTagsForResource", input)
	return &IoT1ClickDevicesServiceListTagsForResourceFuture{Future: future}
}

func (a *IoT1ClickDevicesServiceStub) TagResource(ctx workflow.Context, input *iot1clickdevicesservice.TagResourceInput) (*iot1clickdevicesservice.TagResourceOutput, error) {
	var output iot1clickdevicesservice.TagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.iot1clickdevicesservice.TagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *IoT1ClickDevicesServiceStub) TagResourceAsync(ctx workflow.Context, input *iot1clickdevicesservice.TagResourceInput) *IoT1ClickDevicesServiceTagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.iot1clickdevicesservice.TagResource", input)
	return &IoT1ClickDevicesServiceTagResourceFuture{Future: future}
}

func (a *IoT1ClickDevicesServiceStub) UnclaimDevice(ctx workflow.Context, input *iot1clickdevicesservice.UnclaimDeviceInput) (*iot1clickdevicesservice.UnclaimDeviceOutput, error) {
	var output iot1clickdevicesservice.UnclaimDeviceOutput
	err := workflow.ExecuteActivity(ctx, "aws.iot1clickdevicesservice.UnclaimDevice", input).Get(ctx, &output)
	return &output, err
}

func (a *IoT1ClickDevicesServiceStub) UnclaimDeviceAsync(ctx workflow.Context, input *iot1clickdevicesservice.UnclaimDeviceInput) *IoT1ClickDevicesServiceUnclaimDeviceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.iot1clickdevicesservice.UnclaimDevice", input)
	return &IoT1ClickDevicesServiceUnclaimDeviceFuture{Future: future}
}

func (a *IoT1ClickDevicesServiceStub) UntagResource(ctx workflow.Context, input *iot1clickdevicesservice.UntagResourceInput) (*iot1clickdevicesservice.UntagResourceOutput, error) {
	var output iot1clickdevicesservice.UntagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.iot1clickdevicesservice.UntagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *IoT1ClickDevicesServiceStub) UntagResourceAsync(ctx workflow.Context, input *iot1clickdevicesservice.UntagResourceInput) *IoT1ClickDevicesServiceUntagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.iot1clickdevicesservice.UntagResource", input)
	return &IoT1ClickDevicesServiceUntagResourceFuture{Future: future}
}

func (a *IoT1ClickDevicesServiceStub) UpdateDeviceState(ctx workflow.Context, input *iot1clickdevicesservice.UpdateDeviceStateInput) (*iot1clickdevicesservice.UpdateDeviceStateOutput, error) {
	var output iot1clickdevicesservice.UpdateDeviceStateOutput
	err := workflow.ExecuteActivity(ctx, "aws.iot1clickdevicesservice.UpdateDeviceState", input).Get(ctx, &output)
	return &output, err
}

func (a *IoT1ClickDevicesServiceStub) UpdateDeviceStateAsync(ctx workflow.Context, input *iot1clickdevicesservice.UpdateDeviceStateInput) *IoT1ClickDevicesServiceUpdateDeviceStateFuture {
	future := workflow.ExecuteActivity(ctx, "aws.iot1clickdevicesservice.UpdateDeviceState", input)
	return &IoT1ClickDevicesServiceUpdateDeviceStateFuture{Future: future}
}
