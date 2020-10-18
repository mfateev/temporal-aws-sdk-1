// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package qldbstub

import (
	"github.com/aws/aws-sdk-go/service/qldb"
	"go.temporal.io/sdk/workflow"

	"go.temporal.io/aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	CancelJournalKinesisStream(ctx workflow.Context, input *qldb.CancelJournalKinesisStreamInput) (*qldb.CancelJournalKinesisStreamOutput, error)
	CancelJournalKinesisStreamAsync(ctx workflow.Context, input *qldb.CancelJournalKinesisStreamInput) *QLDBCancelJournalKinesisStreamFuture

	CreateLedger(ctx workflow.Context, input *qldb.CreateLedgerInput) (*qldb.CreateLedgerOutput, error)
	CreateLedgerAsync(ctx workflow.Context, input *qldb.CreateLedgerInput) *QLDBCreateLedgerFuture

	DeleteLedger(ctx workflow.Context, input *qldb.DeleteLedgerInput) (*qldb.DeleteLedgerOutput, error)
	DeleteLedgerAsync(ctx workflow.Context, input *qldb.DeleteLedgerInput) *QLDBDeleteLedgerFuture

	DescribeJournalKinesisStream(ctx workflow.Context, input *qldb.DescribeJournalKinesisStreamInput) (*qldb.DescribeJournalKinesisStreamOutput, error)
	DescribeJournalKinesisStreamAsync(ctx workflow.Context, input *qldb.DescribeJournalKinesisStreamInput) *QLDBDescribeJournalKinesisStreamFuture

	DescribeJournalS3Export(ctx workflow.Context, input *qldb.DescribeJournalS3ExportInput) (*qldb.DescribeJournalS3ExportOutput, error)
	DescribeJournalS3ExportAsync(ctx workflow.Context, input *qldb.DescribeJournalS3ExportInput) *QLDBDescribeJournalS3ExportFuture

	DescribeLedger(ctx workflow.Context, input *qldb.DescribeLedgerInput) (*qldb.DescribeLedgerOutput, error)
	DescribeLedgerAsync(ctx workflow.Context, input *qldb.DescribeLedgerInput) *QLDBDescribeLedgerFuture

	ExportJournalToS3(ctx workflow.Context, input *qldb.ExportJournalToS3Input) (*qldb.ExportJournalToS3Output, error)
	ExportJournalToS3Async(ctx workflow.Context, input *qldb.ExportJournalToS3Input) *QLDBExportJournalToS3Future

	GetBlock(ctx workflow.Context, input *qldb.GetBlockInput) (*qldb.GetBlockOutput, error)
	GetBlockAsync(ctx workflow.Context, input *qldb.GetBlockInput) *QLDBGetBlockFuture

	GetDigest(ctx workflow.Context, input *qldb.GetDigestInput) (*qldb.GetDigestOutput, error)
	GetDigestAsync(ctx workflow.Context, input *qldb.GetDigestInput) *QLDBGetDigestFuture

	GetRevision(ctx workflow.Context, input *qldb.GetRevisionInput) (*qldb.GetRevisionOutput, error)
	GetRevisionAsync(ctx workflow.Context, input *qldb.GetRevisionInput) *QLDBGetRevisionFuture

	ListJournalKinesisStreamsForLedger(ctx workflow.Context, input *qldb.ListJournalKinesisStreamsForLedgerInput) (*qldb.ListJournalKinesisStreamsForLedgerOutput, error)
	ListJournalKinesisStreamsForLedgerAsync(ctx workflow.Context, input *qldb.ListJournalKinesisStreamsForLedgerInput) *QLDBListJournalKinesisStreamsForLedgerFuture

	ListJournalS3Exports(ctx workflow.Context, input *qldb.ListJournalS3ExportsInput) (*qldb.ListJournalS3ExportsOutput, error)
	ListJournalS3ExportsAsync(ctx workflow.Context, input *qldb.ListJournalS3ExportsInput) *QLDBListJournalS3ExportsFuture

	ListJournalS3ExportsForLedger(ctx workflow.Context, input *qldb.ListJournalS3ExportsForLedgerInput) (*qldb.ListJournalS3ExportsForLedgerOutput, error)
	ListJournalS3ExportsForLedgerAsync(ctx workflow.Context, input *qldb.ListJournalS3ExportsForLedgerInput) *QLDBListJournalS3ExportsForLedgerFuture

	ListLedgers(ctx workflow.Context, input *qldb.ListLedgersInput) (*qldb.ListLedgersOutput, error)
	ListLedgersAsync(ctx workflow.Context, input *qldb.ListLedgersInput) *QLDBListLedgersFuture

	ListTagsForResource(ctx workflow.Context, input *qldb.ListTagsForResourceInput) (*qldb.ListTagsForResourceOutput, error)
	ListTagsForResourceAsync(ctx workflow.Context, input *qldb.ListTagsForResourceInput) *QLDBListTagsForResourceFuture

	StreamJournalToKinesis(ctx workflow.Context, input *qldb.StreamJournalToKinesisInput) (*qldb.StreamJournalToKinesisOutput, error)
	StreamJournalToKinesisAsync(ctx workflow.Context, input *qldb.StreamJournalToKinesisInput) *QLDBStreamJournalToKinesisFuture

	TagResource(ctx workflow.Context, input *qldb.TagResourceInput) (*qldb.TagResourceOutput, error)
	TagResourceAsync(ctx workflow.Context, input *qldb.TagResourceInput) *QLDBTagResourceFuture

	UntagResource(ctx workflow.Context, input *qldb.UntagResourceInput) (*qldb.UntagResourceOutput, error)
	UntagResourceAsync(ctx workflow.Context, input *qldb.UntagResourceInput) *QLDBUntagResourceFuture

	UpdateLedger(ctx workflow.Context, input *qldb.UpdateLedgerInput) (*qldb.UpdateLedgerOutput, error)
	UpdateLedgerAsync(ctx workflow.Context, input *qldb.UpdateLedgerInput) *QLDBUpdateLedgerFuture
}

func NewClient() Client {
	return &stub{}
}