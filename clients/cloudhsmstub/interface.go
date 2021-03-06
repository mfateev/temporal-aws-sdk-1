// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package cloudhsmstub

import (
	"github.com/aws/aws-sdk-go/service/cloudhsm"
	"go.temporal.io/sdk/workflow"

	"go.temporal.io/aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	AddTagsToResource(ctx workflow.Context, input *cloudhsm.AddTagsToResourceInput) (*cloudhsm.AddTagsToResourceOutput, error)
	AddTagsToResourceAsync(ctx workflow.Context, input *cloudhsm.AddTagsToResourceInput) *CloudHSMAddTagsToResourceFuture

	CreateHapg(ctx workflow.Context, input *cloudhsm.CreateHapgInput) (*cloudhsm.CreateHapgOutput, error)
	CreateHapgAsync(ctx workflow.Context, input *cloudhsm.CreateHapgInput) *CloudHSMCreateHapgFuture

	CreateHsm(ctx workflow.Context, input *cloudhsm.CreateHsmInput) (*cloudhsm.CreateHsmOutput, error)
	CreateHsmAsync(ctx workflow.Context, input *cloudhsm.CreateHsmInput) *CloudHSMCreateHsmFuture

	CreateLunaClient(ctx workflow.Context, input *cloudhsm.CreateLunaClientInput) (*cloudhsm.CreateLunaClientOutput, error)
	CreateLunaClientAsync(ctx workflow.Context, input *cloudhsm.CreateLunaClientInput) *CloudHSMCreateLunaClientFuture

	DeleteHapg(ctx workflow.Context, input *cloudhsm.DeleteHapgInput) (*cloudhsm.DeleteHapgOutput, error)
	DeleteHapgAsync(ctx workflow.Context, input *cloudhsm.DeleteHapgInput) *CloudHSMDeleteHapgFuture

	DeleteHsm(ctx workflow.Context, input *cloudhsm.DeleteHsmInput) (*cloudhsm.DeleteHsmOutput, error)
	DeleteHsmAsync(ctx workflow.Context, input *cloudhsm.DeleteHsmInput) *CloudHSMDeleteHsmFuture

	DeleteLunaClient(ctx workflow.Context, input *cloudhsm.DeleteLunaClientInput) (*cloudhsm.DeleteLunaClientOutput, error)
	DeleteLunaClientAsync(ctx workflow.Context, input *cloudhsm.DeleteLunaClientInput) *CloudHSMDeleteLunaClientFuture

	DescribeHapg(ctx workflow.Context, input *cloudhsm.DescribeHapgInput) (*cloudhsm.DescribeHapgOutput, error)
	DescribeHapgAsync(ctx workflow.Context, input *cloudhsm.DescribeHapgInput) *CloudHSMDescribeHapgFuture

	DescribeHsm(ctx workflow.Context, input *cloudhsm.DescribeHsmInput) (*cloudhsm.DescribeHsmOutput, error)
	DescribeHsmAsync(ctx workflow.Context, input *cloudhsm.DescribeHsmInput) *CloudHSMDescribeHsmFuture

	DescribeLunaClient(ctx workflow.Context, input *cloudhsm.DescribeLunaClientInput) (*cloudhsm.DescribeLunaClientOutput, error)
	DescribeLunaClientAsync(ctx workflow.Context, input *cloudhsm.DescribeLunaClientInput) *CloudHSMDescribeLunaClientFuture

	GetConfig(ctx workflow.Context, input *cloudhsm.GetConfigInput) (*cloudhsm.GetConfigOutput, error)
	GetConfigAsync(ctx workflow.Context, input *cloudhsm.GetConfigInput) *CloudHSMGetConfigFuture

	ListAvailableZones(ctx workflow.Context, input *cloudhsm.ListAvailableZonesInput) (*cloudhsm.ListAvailableZonesOutput, error)
	ListAvailableZonesAsync(ctx workflow.Context, input *cloudhsm.ListAvailableZonesInput) *CloudHSMListAvailableZonesFuture

	ListHapgs(ctx workflow.Context, input *cloudhsm.ListHapgsInput) (*cloudhsm.ListHapgsOutput, error)
	ListHapgsAsync(ctx workflow.Context, input *cloudhsm.ListHapgsInput) *CloudHSMListHapgsFuture

	ListHsms(ctx workflow.Context, input *cloudhsm.ListHsmsInput) (*cloudhsm.ListHsmsOutput, error)
	ListHsmsAsync(ctx workflow.Context, input *cloudhsm.ListHsmsInput) *CloudHSMListHsmsFuture

	ListLunaClients(ctx workflow.Context, input *cloudhsm.ListLunaClientsInput) (*cloudhsm.ListLunaClientsOutput, error)
	ListLunaClientsAsync(ctx workflow.Context, input *cloudhsm.ListLunaClientsInput) *CloudHSMListLunaClientsFuture

	ListTagsForResource(ctx workflow.Context, input *cloudhsm.ListTagsForResourceInput) (*cloudhsm.ListTagsForResourceOutput, error)
	ListTagsForResourceAsync(ctx workflow.Context, input *cloudhsm.ListTagsForResourceInput) *CloudHSMListTagsForResourceFuture

	ModifyHapg(ctx workflow.Context, input *cloudhsm.ModifyHapgInput) (*cloudhsm.ModifyHapgOutput, error)
	ModifyHapgAsync(ctx workflow.Context, input *cloudhsm.ModifyHapgInput) *CloudHSMModifyHapgFuture

	ModifyHsm(ctx workflow.Context, input *cloudhsm.ModifyHsmInput) (*cloudhsm.ModifyHsmOutput, error)
	ModifyHsmAsync(ctx workflow.Context, input *cloudhsm.ModifyHsmInput) *CloudHSMModifyHsmFuture

	ModifyLunaClient(ctx workflow.Context, input *cloudhsm.ModifyLunaClientInput) (*cloudhsm.ModifyLunaClientOutput, error)
	ModifyLunaClientAsync(ctx workflow.Context, input *cloudhsm.ModifyLunaClientInput) *CloudHSMModifyLunaClientFuture

	RemoveTagsFromResource(ctx workflow.Context, input *cloudhsm.RemoveTagsFromResourceInput) (*cloudhsm.RemoveTagsFromResourceOutput, error)
	RemoveTagsFromResourceAsync(ctx workflow.Context, input *cloudhsm.RemoveTagsFromResourceInput) *CloudHSMRemoveTagsFromResourceFuture
}

func NewClient() Client {
	return &stub{}
}
