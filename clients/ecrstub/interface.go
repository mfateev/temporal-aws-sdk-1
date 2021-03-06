// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package ecrstub

import (
	"github.com/aws/aws-sdk-go/service/ecr"
	"go.temporal.io/sdk/workflow"

	"go.temporal.io/aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	BatchCheckLayerAvailability(ctx workflow.Context, input *ecr.BatchCheckLayerAvailabilityInput) (*ecr.BatchCheckLayerAvailabilityOutput, error)
	BatchCheckLayerAvailabilityAsync(ctx workflow.Context, input *ecr.BatchCheckLayerAvailabilityInput) *ECRBatchCheckLayerAvailabilityFuture

	BatchDeleteImage(ctx workflow.Context, input *ecr.BatchDeleteImageInput) (*ecr.BatchDeleteImageOutput, error)
	BatchDeleteImageAsync(ctx workflow.Context, input *ecr.BatchDeleteImageInput) *ECRBatchDeleteImageFuture

	BatchGetImage(ctx workflow.Context, input *ecr.BatchGetImageInput) (*ecr.BatchGetImageOutput, error)
	BatchGetImageAsync(ctx workflow.Context, input *ecr.BatchGetImageInput) *ECRBatchGetImageFuture

	CompleteLayerUpload(ctx workflow.Context, input *ecr.CompleteLayerUploadInput) (*ecr.CompleteLayerUploadOutput, error)
	CompleteLayerUploadAsync(ctx workflow.Context, input *ecr.CompleteLayerUploadInput) *ECRCompleteLayerUploadFuture

	CreateRepository(ctx workflow.Context, input *ecr.CreateRepositoryInput) (*ecr.CreateRepositoryOutput, error)
	CreateRepositoryAsync(ctx workflow.Context, input *ecr.CreateRepositoryInput) *ECRCreateRepositoryFuture

	DeleteLifecyclePolicy(ctx workflow.Context, input *ecr.DeleteLifecyclePolicyInput) (*ecr.DeleteLifecyclePolicyOutput, error)
	DeleteLifecyclePolicyAsync(ctx workflow.Context, input *ecr.DeleteLifecyclePolicyInput) *ECRDeleteLifecyclePolicyFuture

	DeleteRepository(ctx workflow.Context, input *ecr.DeleteRepositoryInput) (*ecr.DeleteRepositoryOutput, error)
	DeleteRepositoryAsync(ctx workflow.Context, input *ecr.DeleteRepositoryInput) *ECRDeleteRepositoryFuture

	DeleteRepositoryPolicy(ctx workflow.Context, input *ecr.DeleteRepositoryPolicyInput) (*ecr.DeleteRepositoryPolicyOutput, error)
	DeleteRepositoryPolicyAsync(ctx workflow.Context, input *ecr.DeleteRepositoryPolicyInput) *ECRDeleteRepositoryPolicyFuture

	DescribeImageScanFindings(ctx workflow.Context, input *ecr.DescribeImageScanFindingsInput) (*ecr.DescribeImageScanFindingsOutput, error)
	DescribeImageScanFindingsAsync(ctx workflow.Context, input *ecr.DescribeImageScanFindingsInput) *ECRDescribeImageScanFindingsFuture

	DescribeImages(ctx workflow.Context, input *ecr.DescribeImagesInput) (*ecr.DescribeImagesOutput, error)
	DescribeImagesAsync(ctx workflow.Context, input *ecr.DescribeImagesInput) *ECRDescribeImagesFuture

	DescribeRepositories(ctx workflow.Context, input *ecr.DescribeRepositoriesInput) (*ecr.DescribeRepositoriesOutput, error)
	DescribeRepositoriesAsync(ctx workflow.Context, input *ecr.DescribeRepositoriesInput) *ECRDescribeRepositoriesFuture

	GetAuthorizationToken(ctx workflow.Context, input *ecr.GetAuthorizationTokenInput) (*ecr.GetAuthorizationTokenOutput, error)
	GetAuthorizationTokenAsync(ctx workflow.Context, input *ecr.GetAuthorizationTokenInput) *ECRGetAuthorizationTokenFuture

	GetDownloadUrlForLayer(ctx workflow.Context, input *ecr.GetDownloadUrlForLayerInput) (*ecr.GetDownloadUrlForLayerOutput, error)
	GetDownloadUrlForLayerAsync(ctx workflow.Context, input *ecr.GetDownloadUrlForLayerInput) *ECRGetDownloadUrlForLayerFuture

	GetLifecyclePolicy(ctx workflow.Context, input *ecr.GetLifecyclePolicyInput) (*ecr.GetLifecyclePolicyOutput, error)
	GetLifecyclePolicyAsync(ctx workflow.Context, input *ecr.GetLifecyclePolicyInput) *ECRGetLifecyclePolicyFuture

	GetLifecyclePolicyPreview(ctx workflow.Context, input *ecr.GetLifecyclePolicyPreviewInput) (*ecr.GetLifecyclePolicyPreviewOutput, error)
	GetLifecyclePolicyPreviewAsync(ctx workflow.Context, input *ecr.GetLifecyclePolicyPreviewInput) *ECRGetLifecyclePolicyPreviewFuture

	GetRepositoryPolicy(ctx workflow.Context, input *ecr.GetRepositoryPolicyInput) (*ecr.GetRepositoryPolicyOutput, error)
	GetRepositoryPolicyAsync(ctx workflow.Context, input *ecr.GetRepositoryPolicyInput) *ECRGetRepositoryPolicyFuture

	InitiateLayerUpload(ctx workflow.Context, input *ecr.InitiateLayerUploadInput) (*ecr.InitiateLayerUploadOutput, error)
	InitiateLayerUploadAsync(ctx workflow.Context, input *ecr.InitiateLayerUploadInput) *ECRInitiateLayerUploadFuture

	ListImages(ctx workflow.Context, input *ecr.ListImagesInput) (*ecr.ListImagesOutput, error)
	ListImagesAsync(ctx workflow.Context, input *ecr.ListImagesInput) *ECRListImagesFuture

	ListTagsForResource(ctx workflow.Context, input *ecr.ListTagsForResourceInput) (*ecr.ListTagsForResourceOutput, error)
	ListTagsForResourceAsync(ctx workflow.Context, input *ecr.ListTagsForResourceInput) *ECRListTagsForResourceFuture

	PutImage(ctx workflow.Context, input *ecr.PutImageInput) (*ecr.PutImageOutput, error)
	PutImageAsync(ctx workflow.Context, input *ecr.PutImageInput) *ECRPutImageFuture

	PutImageScanningConfiguration(ctx workflow.Context, input *ecr.PutImageScanningConfigurationInput) (*ecr.PutImageScanningConfigurationOutput, error)
	PutImageScanningConfigurationAsync(ctx workflow.Context, input *ecr.PutImageScanningConfigurationInput) *ECRPutImageScanningConfigurationFuture

	PutImageTagMutability(ctx workflow.Context, input *ecr.PutImageTagMutabilityInput) (*ecr.PutImageTagMutabilityOutput, error)
	PutImageTagMutabilityAsync(ctx workflow.Context, input *ecr.PutImageTagMutabilityInput) *ECRPutImageTagMutabilityFuture

	PutLifecyclePolicy(ctx workflow.Context, input *ecr.PutLifecyclePolicyInput) (*ecr.PutLifecyclePolicyOutput, error)
	PutLifecyclePolicyAsync(ctx workflow.Context, input *ecr.PutLifecyclePolicyInput) *ECRPutLifecyclePolicyFuture

	SetRepositoryPolicy(ctx workflow.Context, input *ecr.SetRepositoryPolicyInput) (*ecr.SetRepositoryPolicyOutput, error)
	SetRepositoryPolicyAsync(ctx workflow.Context, input *ecr.SetRepositoryPolicyInput) *ECRSetRepositoryPolicyFuture

	StartImageScan(ctx workflow.Context, input *ecr.StartImageScanInput) (*ecr.StartImageScanOutput, error)
	StartImageScanAsync(ctx workflow.Context, input *ecr.StartImageScanInput) *ECRStartImageScanFuture

	StartLifecyclePolicyPreview(ctx workflow.Context, input *ecr.StartLifecyclePolicyPreviewInput) (*ecr.StartLifecyclePolicyPreviewOutput, error)
	StartLifecyclePolicyPreviewAsync(ctx workflow.Context, input *ecr.StartLifecyclePolicyPreviewInput) *ECRStartLifecyclePolicyPreviewFuture

	TagResource(ctx workflow.Context, input *ecr.TagResourceInput) (*ecr.TagResourceOutput, error)
	TagResourceAsync(ctx workflow.Context, input *ecr.TagResourceInput) *ECRTagResourceFuture

	UntagResource(ctx workflow.Context, input *ecr.UntagResourceInput) (*ecr.UntagResourceOutput, error)
	UntagResourceAsync(ctx workflow.Context, input *ecr.UntagResourceInput) *ECRUntagResourceFuture

	UploadLayerPart(ctx workflow.Context, input *ecr.UploadLayerPartInput) (*ecr.UploadLayerPartOutput, error)
	UploadLayerPartAsync(ctx workflow.Context, input *ecr.UploadLayerPartInput) *ECRUploadLayerPartFuture

	WaitUntilImageScanComplete(ctx workflow.Context, input *ecr.DescribeImageScanFindingsInput) error
	WaitUntilImageScanCompleteAsync(ctx workflow.Context, input *ecr.DescribeImageScanFindingsInput) *clients.VoidFuture

	WaitUntilLifecyclePolicyPreviewComplete(ctx workflow.Context, input *ecr.GetLifecyclePolicyPreviewInput) error
	WaitUntilLifecyclePolicyPreviewCompleteAsync(ctx workflow.Context, input *ecr.GetLifecyclePolicyPreviewInput) *clients.VoidFuture
}

func NewClient() Client {
	return &stub{}
}
