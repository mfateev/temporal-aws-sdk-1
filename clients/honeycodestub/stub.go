// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package honeycodestub

import (
	"github.com/aws/aws-sdk-go/service/honeycode"
	"go.temporal.io/sdk/workflow"

	"go.temporal.io/aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type stub struct{}

type HoneycodeGetScreenDataFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *HoneycodeGetScreenDataFuture) Get(ctx workflow.Context) (*honeycode.GetScreenDataOutput, error) {
	var output honeycode.GetScreenDataOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type HoneycodeInvokeScreenAutomationFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *HoneycodeInvokeScreenAutomationFuture) Get(ctx workflow.Context) (*honeycode.InvokeScreenAutomationOutput, error) {
	var output honeycode.InvokeScreenAutomationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *stub) GetScreenData(ctx workflow.Context, input *honeycode.GetScreenDataInput) (*honeycode.GetScreenDataOutput, error) {
	var output honeycode.GetScreenDataOutput
	err := workflow.ExecuteActivity(ctx, "aws.honeycode.GetScreenData", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetScreenDataAsync(ctx workflow.Context, input *honeycode.GetScreenDataInput) *HoneycodeGetScreenDataFuture {
	future := workflow.ExecuteActivity(ctx, "aws.honeycode.GetScreenData", input)
	return &HoneycodeGetScreenDataFuture{Future: future}
}

func (a *stub) InvokeScreenAutomation(ctx workflow.Context, input *honeycode.InvokeScreenAutomationInput) (*honeycode.InvokeScreenAutomationOutput, error) {
	var output honeycode.InvokeScreenAutomationOutput
	err := workflow.ExecuteActivity(ctx, "aws.honeycode.InvokeScreenAutomation", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) InvokeScreenAutomationAsync(ctx workflow.Context, input *honeycode.InvokeScreenAutomationInput) *HoneycodeInvokeScreenAutomationFuture {
	future := workflow.ExecuteActivity(ctx, "aws.honeycode.InvokeScreenAutomation", input)
	return &HoneycodeInvokeScreenAutomationFuture{Future: future}
}
