// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package costandusagereportservicestub

import (
	"github.com/aws/aws-sdk-go/service/costandusagereportservice"
	"go.temporal.io/sdk/workflow"

	"go.temporal.io/aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	DeleteReportDefinition(ctx workflow.Context, input *costandusagereportservice.DeleteReportDefinitionInput) (*costandusagereportservice.DeleteReportDefinitionOutput, error)
	DeleteReportDefinitionAsync(ctx workflow.Context, input *costandusagereportservice.DeleteReportDefinitionInput) *CostandUsageReportServiceDeleteReportDefinitionFuture

	DescribeReportDefinitions(ctx workflow.Context, input *costandusagereportservice.DescribeReportDefinitionsInput) (*costandusagereportservice.DescribeReportDefinitionsOutput, error)
	DescribeReportDefinitionsAsync(ctx workflow.Context, input *costandusagereportservice.DescribeReportDefinitionsInput) *CostandUsageReportServiceDescribeReportDefinitionsFuture

	ModifyReportDefinition(ctx workflow.Context, input *costandusagereportservice.ModifyReportDefinitionInput) (*costandusagereportservice.ModifyReportDefinitionOutput, error)
	ModifyReportDefinitionAsync(ctx workflow.Context, input *costandusagereportservice.ModifyReportDefinitionInput) *CostandUsageReportServiceModifyReportDefinitionFuture

	PutReportDefinition(ctx workflow.Context, input *costandusagereportservice.PutReportDefinitionInput) (*costandusagereportservice.PutReportDefinitionOutput, error)
	PutReportDefinitionAsync(ctx workflow.Context, input *costandusagereportservice.PutReportDefinitionInput) *CostandUsageReportServicePutReportDefinitionFuture
}

func NewClient() Client {
	return &stub{}
}