// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package opsworkscmstub

import (
	"github.com/aws/aws-sdk-go/service/opsworkscm"
	"go.temporal.io/sdk/workflow"

	"go.temporal.io/aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	AssociateNode(ctx workflow.Context, input *opsworkscm.AssociateNodeInput) (*opsworkscm.AssociateNodeOutput, error)
	AssociateNodeAsync(ctx workflow.Context, input *opsworkscm.AssociateNodeInput) *OpsWorksCMAssociateNodeFuture

	CreateBackup(ctx workflow.Context, input *opsworkscm.CreateBackupInput) (*opsworkscm.CreateBackupOutput, error)
	CreateBackupAsync(ctx workflow.Context, input *opsworkscm.CreateBackupInput) *OpsWorksCMCreateBackupFuture

	CreateServer(ctx workflow.Context, input *opsworkscm.CreateServerInput) (*opsworkscm.CreateServerOutput, error)
	CreateServerAsync(ctx workflow.Context, input *opsworkscm.CreateServerInput) *OpsWorksCMCreateServerFuture

	DeleteBackup(ctx workflow.Context, input *opsworkscm.DeleteBackupInput) (*opsworkscm.DeleteBackupOutput, error)
	DeleteBackupAsync(ctx workflow.Context, input *opsworkscm.DeleteBackupInput) *OpsWorksCMDeleteBackupFuture

	DeleteServer(ctx workflow.Context, input *opsworkscm.DeleteServerInput) (*opsworkscm.DeleteServerOutput, error)
	DeleteServerAsync(ctx workflow.Context, input *opsworkscm.DeleteServerInput) *OpsWorksCMDeleteServerFuture

	DescribeAccountAttributes(ctx workflow.Context, input *opsworkscm.DescribeAccountAttributesInput) (*opsworkscm.DescribeAccountAttributesOutput, error)
	DescribeAccountAttributesAsync(ctx workflow.Context, input *opsworkscm.DescribeAccountAttributesInput) *OpsWorksCMDescribeAccountAttributesFuture

	DescribeBackups(ctx workflow.Context, input *opsworkscm.DescribeBackupsInput) (*opsworkscm.DescribeBackupsOutput, error)
	DescribeBackupsAsync(ctx workflow.Context, input *opsworkscm.DescribeBackupsInput) *OpsWorksCMDescribeBackupsFuture

	DescribeEvents(ctx workflow.Context, input *opsworkscm.DescribeEventsInput) (*opsworkscm.DescribeEventsOutput, error)
	DescribeEventsAsync(ctx workflow.Context, input *opsworkscm.DescribeEventsInput) *OpsWorksCMDescribeEventsFuture

	DescribeNodeAssociationStatus(ctx workflow.Context, input *opsworkscm.DescribeNodeAssociationStatusInput) (*opsworkscm.DescribeNodeAssociationStatusOutput, error)
	DescribeNodeAssociationStatusAsync(ctx workflow.Context, input *opsworkscm.DescribeNodeAssociationStatusInput) *OpsWorksCMDescribeNodeAssociationStatusFuture

	DescribeServers(ctx workflow.Context, input *opsworkscm.DescribeServersInput) (*opsworkscm.DescribeServersOutput, error)
	DescribeServersAsync(ctx workflow.Context, input *opsworkscm.DescribeServersInput) *OpsWorksCMDescribeServersFuture

	DisassociateNode(ctx workflow.Context, input *opsworkscm.DisassociateNodeInput) (*opsworkscm.DisassociateNodeOutput, error)
	DisassociateNodeAsync(ctx workflow.Context, input *opsworkscm.DisassociateNodeInput) *OpsWorksCMDisassociateNodeFuture

	ExportServerEngineAttribute(ctx workflow.Context, input *opsworkscm.ExportServerEngineAttributeInput) (*opsworkscm.ExportServerEngineAttributeOutput, error)
	ExportServerEngineAttributeAsync(ctx workflow.Context, input *opsworkscm.ExportServerEngineAttributeInput) *OpsWorksCMExportServerEngineAttributeFuture

	ListTagsForResource(ctx workflow.Context, input *opsworkscm.ListTagsForResourceInput) (*opsworkscm.ListTagsForResourceOutput, error)
	ListTagsForResourceAsync(ctx workflow.Context, input *opsworkscm.ListTagsForResourceInput) *OpsWorksCMListTagsForResourceFuture

	RestoreServer(ctx workflow.Context, input *opsworkscm.RestoreServerInput) (*opsworkscm.RestoreServerOutput, error)
	RestoreServerAsync(ctx workflow.Context, input *opsworkscm.RestoreServerInput) *OpsWorksCMRestoreServerFuture

	StartMaintenance(ctx workflow.Context, input *opsworkscm.StartMaintenanceInput) (*opsworkscm.StartMaintenanceOutput, error)
	StartMaintenanceAsync(ctx workflow.Context, input *opsworkscm.StartMaintenanceInput) *OpsWorksCMStartMaintenanceFuture

	TagResource(ctx workflow.Context, input *opsworkscm.TagResourceInput) (*opsworkscm.TagResourceOutput, error)
	TagResourceAsync(ctx workflow.Context, input *opsworkscm.TagResourceInput) *OpsWorksCMTagResourceFuture

	UntagResource(ctx workflow.Context, input *opsworkscm.UntagResourceInput) (*opsworkscm.UntagResourceOutput, error)
	UntagResourceAsync(ctx workflow.Context, input *opsworkscm.UntagResourceInput) *OpsWorksCMUntagResourceFuture

	UpdateServer(ctx workflow.Context, input *opsworkscm.UpdateServerInput) (*opsworkscm.UpdateServerOutput, error)
	UpdateServerAsync(ctx workflow.Context, input *opsworkscm.UpdateServerInput) *OpsWorksCMUpdateServerFuture

	UpdateServerEngineAttributes(ctx workflow.Context, input *opsworkscm.UpdateServerEngineAttributesInput) (*opsworkscm.UpdateServerEngineAttributesOutput, error)
	UpdateServerEngineAttributesAsync(ctx workflow.Context, input *opsworkscm.UpdateServerEngineAttributesInput) *OpsWorksCMUpdateServerEngineAttributesFuture

	WaitUntilNodeAssociated(ctx workflow.Context, input *opsworkscm.DescribeNodeAssociationStatusInput) error
	WaitUntilNodeAssociatedAsync(ctx workflow.Context, input *opsworkscm.DescribeNodeAssociationStatusInput) *clients.VoidFuture
}

func NewClient() Client {
	return &stub{}
}