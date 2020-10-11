// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package awsclients

import (
	"github.com/aws/aws-sdk-go/service/rdsdataservice"
	"go.temporal.io/sdk/workflow"
)

type RDSDataServiceClient interface {
	BatchExecuteStatement(ctx workflow.Context, input *rdsdataservice.BatchExecuteStatementInput) (*rdsdataservice.BatchExecuteStatementOutput, error)
	BatchExecuteStatementAsync(ctx workflow.Context, input *rdsdataservice.BatchExecuteStatementInput) *RdsdataserviceBatchExecuteStatementResult

	BeginTransaction(ctx workflow.Context, input *rdsdataservice.BeginTransactionInput) (*rdsdataservice.BeginTransactionOutput, error)
	BeginTransactionAsync(ctx workflow.Context, input *rdsdataservice.BeginTransactionInput) *RdsdataserviceBeginTransactionResult

	CommitTransaction(ctx workflow.Context, input *rdsdataservice.CommitTransactionInput) (*rdsdataservice.CommitTransactionOutput, error)
	CommitTransactionAsync(ctx workflow.Context, input *rdsdataservice.CommitTransactionInput) *RdsdataserviceCommitTransactionResult

	ExecuteSql(ctx workflow.Context, input *rdsdataservice.ExecuteSqlInput) (*rdsdataservice.ExecuteSqlOutput, error)
	ExecuteSqlAsync(ctx workflow.Context, input *rdsdataservice.ExecuteSqlInput) *RdsdataserviceExecuteSqlResult

	ExecuteStatement(ctx workflow.Context, input *rdsdataservice.ExecuteStatementInput) (*rdsdataservice.ExecuteStatementOutput, error)
	ExecuteStatementAsync(ctx workflow.Context, input *rdsdataservice.ExecuteStatementInput) *RdsdataserviceExecuteStatementResult

	RollbackTransaction(ctx workflow.Context, input *rdsdataservice.RollbackTransactionInput) (*rdsdataservice.RollbackTransactionOutput, error)
	RollbackTransactionAsync(ctx workflow.Context, input *rdsdataservice.RollbackTransactionInput) *RdsdataserviceRollbackTransactionResult
}

type RDSDataServiceStub struct{}

func NewRDSDataServiceStub() RDSDataServiceClient {
	return &RDSDataServiceStub{}
}

type RdsdataserviceBatchExecuteStatementResult struct {
	Result workflow.Future
}

func (r *RdsdataserviceBatchExecuteStatementResult) Get(ctx workflow.Context) (*rdsdataservice.BatchExecuteStatementOutput, error) {
	var output rdsdataservice.BatchExecuteStatementOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type RdsdataserviceBeginTransactionResult struct {
	Result workflow.Future
}

func (r *RdsdataserviceBeginTransactionResult) Get(ctx workflow.Context) (*rdsdataservice.BeginTransactionOutput, error) {
	var output rdsdataservice.BeginTransactionOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type RdsdataserviceCommitTransactionResult struct {
	Result workflow.Future
}

func (r *RdsdataserviceCommitTransactionResult) Get(ctx workflow.Context) (*rdsdataservice.CommitTransactionOutput, error) {
	var output rdsdataservice.CommitTransactionOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type RdsdataserviceExecuteSqlResult struct {
	Result workflow.Future
}

func (r *RdsdataserviceExecuteSqlResult) Get(ctx workflow.Context) (*rdsdataservice.ExecuteSqlOutput, error) {
	var output rdsdataservice.ExecuteSqlOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type RdsdataserviceExecuteStatementResult struct {
	Result workflow.Future
}

func (r *RdsdataserviceExecuteStatementResult) Get(ctx workflow.Context) (*rdsdataservice.ExecuteStatementOutput, error) {
	var output rdsdataservice.ExecuteStatementOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type RdsdataserviceRollbackTransactionResult struct {
	Result workflow.Future
}

func (r *RdsdataserviceRollbackTransactionResult) Get(ctx workflow.Context) (*rdsdataservice.RollbackTransactionOutput, error) {
	var output rdsdataservice.RollbackTransactionOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

func (a *RDSDataServiceStub) BatchExecuteStatement(ctx workflow.Context, input *rdsdataservice.BatchExecuteStatementInput) (*rdsdataservice.BatchExecuteStatementOutput, error) {
	var output rdsdataservice.BatchExecuteStatementOutput
	err := workflow.ExecuteActivity(ctx, "aws.rdsdataservice.BatchExecuteStatement", input).Get(ctx, &output)
	return &output, err
}

func (a *RDSDataServiceStub) BatchExecuteStatementAsync(ctx workflow.Context, input *rdsdataservice.BatchExecuteStatementInput) *RdsdataserviceBatchExecuteStatementResult {
	future := workflow.ExecuteActivity(ctx, "aws.rdsdataservice.BatchExecuteStatement", input)
	return &RdsdataserviceBatchExecuteStatementResult{Result: future}
}

func (a *RDSDataServiceStub) BeginTransaction(ctx workflow.Context, input *rdsdataservice.BeginTransactionInput) (*rdsdataservice.BeginTransactionOutput, error) {
	var output rdsdataservice.BeginTransactionOutput
	err := workflow.ExecuteActivity(ctx, "aws.rdsdataservice.BeginTransaction", input).Get(ctx, &output)
	return &output, err
}

func (a *RDSDataServiceStub) BeginTransactionAsync(ctx workflow.Context, input *rdsdataservice.BeginTransactionInput) *RdsdataserviceBeginTransactionResult {
	future := workflow.ExecuteActivity(ctx, "aws.rdsdataservice.BeginTransaction", input)
	return &RdsdataserviceBeginTransactionResult{Result: future}
}

func (a *RDSDataServiceStub) CommitTransaction(ctx workflow.Context, input *rdsdataservice.CommitTransactionInput) (*rdsdataservice.CommitTransactionOutput, error) {
	var output rdsdataservice.CommitTransactionOutput
	err := workflow.ExecuteActivity(ctx, "aws.rdsdataservice.CommitTransaction", input).Get(ctx, &output)
	return &output, err
}

func (a *RDSDataServiceStub) CommitTransactionAsync(ctx workflow.Context, input *rdsdataservice.CommitTransactionInput) *RdsdataserviceCommitTransactionResult {
	future := workflow.ExecuteActivity(ctx, "aws.rdsdataservice.CommitTransaction", input)
	return &RdsdataserviceCommitTransactionResult{Result: future}
}

func (a *RDSDataServiceStub) ExecuteSql(ctx workflow.Context, input *rdsdataservice.ExecuteSqlInput) (*rdsdataservice.ExecuteSqlOutput, error) {
	var output rdsdataservice.ExecuteSqlOutput
	err := workflow.ExecuteActivity(ctx, "aws.rdsdataservice.ExecuteSql", input).Get(ctx, &output)
	return &output, err
}

func (a *RDSDataServiceStub) ExecuteSqlAsync(ctx workflow.Context, input *rdsdataservice.ExecuteSqlInput) *RdsdataserviceExecuteSqlResult {
	future := workflow.ExecuteActivity(ctx, "aws.rdsdataservice.ExecuteSql", input)
	return &RdsdataserviceExecuteSqlResult{Result: future}
}

func (a *RDSDataServiceStub) ExecuteStatement(ctx workflow.Context, input *rdsdataservice.ExecuteStatementInput) (*rdsdataservice.ExecuteStatementOutput, error) {
	var output rdsdataservice.ExecuteStatementOutput
	err := workflow.ExecuteActivity(ctx, "aws.rdsdataservice.ExecuteStatement", input).Get(ctx, &output)
	return &output, err
}

func (a *RDSDataServiceStub) ExecuteStatementAsync(ctx workflow.Context, input *rdsdataservice.ExecuteStatementInput) *RdsdataserviceExecuteStatementResult {
	future := workflow.ExecuteActivity(ctx, "aws.rdsdataservice.ExecuteStatement", input)
	return &RdsdataserviceExecuteStatementResult{Result: future}
}

func (a *RDSDataServiceStub) RollbackTransaction(ctx workflow.Context, input *rdsdataservice.RollbackTransactionInput) (*rdsdataservice.RollbackTransactionOutput, error) {
	var output rdsdataservice.RollbackTransactionOutput
	err := workflow.ExecuteActivity(ctx, "aws.rdsdataservice.RollbackTransaction", input).Get(ctx, &output)
	return &output, err
}

func (a *RDSDataServiceStub) RollbackTransactionAsync(ctx workflow.Context, input *rdsdataservice.RollbackTransactionInput) *RdsdataserviceRollbackTransactionResult {
	future := workflow.ExecuteActivity(ctx, "aws.rdsdataservice.RollbackTransaction", input)
	return &RdsdataserviceRollbackTransactionResult{Result: future}
}
