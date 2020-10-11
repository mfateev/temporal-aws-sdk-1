// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package awsclients

import (
	"github.com/aws/aws-sdk-go/service/acmpca"
	"go.temporal.io/sdk/workflow"
)

type ACMPCAClient interface {
	CreateCertificateAuthority(ctx workflow.Context, input *acmpca.CreateCertificateAuthorityInput) (*acmpca.CreateCertificateAuthorityOutput, error)
	CreateCertificateAuthorityAsync(ctx workflow.Context, input *acmpca.CreateCertificateAuthorityInput) *AcmpcaCreateCertificateAuthorityFuture

	CreateCertificateAuthorityAuditReport(ctx workflow.Context, input *acmpca.CreateCertificateAuthorityAuditReportInput) (*acmpca.CreateCertificateAuthorityAuditReportOutput, error)
	CreateCertificateAuthorityAuditReportAsync(ctx workflow.Context, input *acmpca.CreateCertificateAuthorityAuditReportInput) *AcmpcaCreateCertificateAuthorityAuditReportFuture

	CreatePermission(ctx workflow.Context, input *acmpca.CreatePermissionInput) (*acmpca.CreatePermissionOutput, error)
	CreatePermissionAsync(ctx workflow.Context, input *acmpca.CreatePermissionInput) *AcmpcaCreatePermissionFuture

	DeleteCertificateAuthority(ctx workflow.Context, input *acmpca.DeleteCertificateAuthorityInput) (*acmpca.DeleteCertificateAuthorityOutput, error)
	DeleteCertificateAuthorityAsync(ctx workflow.Context, input *acmpca.DeleteCertificateAuthorityInput) *AcmpcaDeleteCertificateAuthorityFuture

	DeletePermission(ctx workflow.Context, input *acmpca.DeletePermissionInput) (*acmpca.DeletePermissionOutput, error)
	DeletePermissionAsync(ctx workflow.Context, input *acmpca.DeletePermissionInput) *AcmpcaDeletePermissionFuture

	DeletePolicy(ctx workflow.Context, input *acmpca.DeletePolicyInput) (*acmpca.DeletePolicyOutput, error)
	DeletePolicyAsync(ctx workflow.Context, input *acmpca.DeletePolicyInput) *AcmpcaDeletePolicyFuture

	DescribeCertificateAuthority(ctx workflow.Context, input *acmpca.DescribeCertificateAuthorityInput) (*acmpca.DescribeCertificateAuthorityOutput, error)
	DescribeCertificateAuthorityAsync(ctx workflow.Context, input *acmpca.DescribeCertificateAuthorityInput) *AcmpcaDescribeCertificateAuthorityFuture

	DescribeCertificateAuthorityAuditReport(ctx workflow.Context, input *acmpca.DescribeCertificateAuthorityAuditReportInput) (*acmpca.DescribeCertificateAuthorityAuditReportOutput, error)
	DescribeCertificateAuthorityAuditReportAsync(ctx workflow.Context, input *acmpca.DescribeCertificateAuthorityAuditReportInput) *AcmpcaDescribeCertificateAuthorityAuditReportFuture

	GetCertificate(ctx workflow.Context, input *acmpca.GetCertificateInput) (*acmpca.GetCertificateOutput, error)
	GetCertificateAsync(ctx workflow.Context, input *acmpca.GetCertificateInput) *AcmpcaGetCertificateFuture

	GetCertificateAuthorityCertificate(ctx workflow.Context, input *acmpca.GetCertificateAuthorityCertificateInput) (*acmpca.GetCertificateAuthorityCertificateOutput, error)
	GetCertificateAuthorityCertificateAsync(ctx workflow.Context, input *acmpca.GetCertificateAuthorityCertificateInput) *AcmpcaGetCertificateAuthorityCertificateFuture

	GetCertificateAuthorityCsr(ctx workflow.Context, input *acmpca.GetCertificateAuthorityCsrInput) (*acmpca.GetCertificateAuthorityCsrOutput, error)
	GetCertificateAuthorityCsrAsync(ctx workflow.Context, input *acmpca.GetCertificateAuthorityCsrInput) *AcmpcaGetCertificateAuthorityCsrFuture

	GetPolicy(ctx workflow.Context, input *acmpca.GetPolicyInput) (*acmpca.GetPolicyOutput, error)
	GetPolicyAsync(ctx workflow.Context, input *acmpca.GetPolicyInput) *AcmpcaGetPolicyFuture

	ImportCertificateAuthorityCertificate(ctx workflow.Context, input *acmpca.ImportCertificateAuthorityCertificateInput) (*acmpca.ImportCertificateAuthorityCertificateOutput, error)
	ImportCertificateAuthorityCertificateAsync(ctx workflow.Context, input *acmpca.ImportCertificateAuthorityCertificateInput) *AcmpcaImportCertificateAuthorityCertificateFuture

	IssueCertificate(ctx workflow.Context, input *acmpca.IssueCertificateInput) (*acmpca.IssueCertificateOutput, error)
	IssueCertificateAsync(ctx workflow.Context, input *acmpca.IssueCertificateInput) *AcmpcaIssueCertificateFuture

	ListCertificateAuthorities(ctx workflow.Context, input *acmpca.ListCertificateAuthoritiesInput) (*acmpca.ListCertificateAuthoritiesOutput, error)
	ListCertificateAuthoritiesAsync(ctx workflow.Context, input *acmpca.ListCertificateAuthoritiesInput) *AcmpcaListCertificateAuthoritiesFuture

	ListPermissions(ctx workflow.Context, input *acmpca.ListPermissionsInput) (*acmpca.ListPermissionsOutput, error)
	ListPermissionsAsync(ctx workflow.Context, input *acmpca.ListPermissionsInput) *AcmpcaListPermissionsFuture

	ListTags(ctx workflow.Context, input *acmpca.ListTagsInput) (*acmpca.ListTagsOutput, error)
	ListTagsAsync(ctx workflow.Context, input *acmpca.ListTagsInput) *AcmpcaListTagsFuture

	PutPolicy(ctx workflow.Context, input *acmpca.PutPolicyInput) (*acmpca.PutPolicyOutput, error)
	PutPolicyAsync(ctx workflow.Context, input *acmpca.PutPolicyInput) *AcmpcaPutPolicyFuture

	RestoreCertificateAuthority(ctx workflow.Context, input *acmpca.RestoreCertificateAuthorityInput) (*acmpca.RestoreCertificateAuthorityOutput, error)
	RestoreCertificateAuthorityAsync(ctx workflow.Context, input *acmpca.RestoreCertificateAuthorityInput) *AcmpcaRestoreCertificateAuthorityFuture

	RevokeCertificate(ctx workflow.Context, input *acmpca.RevokeCertificateInput) (*acmpca.RevokeCertificateOutput, error)
	RevokeCertificateAsync(ctx workflow.Context, input *acmpca.RevokeCertificateInput) *AcmpcaRevokeCertificateFuture

	TagCertificateAuthority(ctx workflow.Context, input *acmpca.TagCertificateAuthorityInput) (*acmpca.TagCertificateAuthorityOutput, error)
	TagCertificateAuthorityAsync(ctx workflow.Context, input *acmpca.TagCertificateAuthorityInput) *AcmpcaTagCertificateAuthorityFuture

	UntagCertificateAuthority(ctx workflow.Context, input *acmpca.UntagCertificateAuthorityInput) (*acmpca.UntagCertificateAuthorityOutput, error)
	UntagCertificateAuthorityAsync(ctx workflow.Context, input *acmpca.UntagCertificateAuthorityInput) *AcmpcaUntagCertificateAuthorityFuture

	UpdateCertificateAuthority(ctx workflow.Context, input *acmpca.UpdateCertificateAuthorityInput) (*acmpca.UpdateCertificateAuthorityOutput, error)
	UpdateCertificateAuthorityAsync(ctx workflow.Context, input *acmpca.UpdateCertificateAuthorityInput) *AcmpcaUpdateCertificateAuthorityFuture

	WaitUntilAuditReportCreated(ctx workflow.Context, input *acmpca.DescribeCertificateAuthorityAuditReportInput) error
	WaitUntilAuditReportCreatedAsync(ctx workflow.Context, input *acmpca.DescribeCertificateAuthorityAuditReportInput) *VoidFuture

	WaitUntilCertificateAuthorityCSRCreated(ctx workflow.Context, input *acmpca.GetCertificateAuthorityCsrInput) error
	WaitUntilCertificateAuthorityCSRCreatedAsync(ctx workflow.Context, input *acmpca.GetCertificateAuthorityCsrInput) *VoidFuture

	WaitUntilCertificateIssued(ctx workflow.Context, input *acmpca.GetCertificateInput) error
	WaitUntilCertificateIssuedAsync(ctx workflow.Context, input *acmpca.GetCertificateInput) *VoidFuture
}

type ACMPCAStub struct{}

func NewACMPCAStub() ACMPCAClient {
	return &ACMPCAStub{}
}

type AcmpcaCreateCertificateAuthorityFuture struct {
	Future workflow.Future
}

func (r *AcmpcaCreateCertificateAuthorityFuture) Get(ctx workflow.Context) (*acmpca.CreateCertificateAuthorityOutput, error) {
	var output acmpca.CreateCertificateAuthorityOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AcmpcaCreateCertificateAuthorityAuditReportFuture struct {
	Future workflow.Future
}

func (r *AcmpcaCreateCertificateAuthorityAuditReportFuture) Get(ctx workflow.Context) (*acmpca.CreateCertificateAuthorityAuditReportOutput, error) {
	var output acmpca.CreateCertificateAuthorityAuditReportOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AcmpcaCreatePermissionFuture struct {
	Future workflow.Future
}

func (r *AcmpcaCreatePermissionFuture) Get(ctx workflow.Context) (*acmpca.CreatePermissionOutput, error) {
	var output acmpca.CreatePermissionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AcmpcaDeleteCertificateAuthorityFuture struct {
	Future workflow.Future
}

func (r *AcmpcaDeleteCertificateAuthorityFuture) Get(ctx workflow.Context) (*acmpca.DeleteCertificateAuthorityOutput, error) {
	var output acmpca.DeleteCertificateAuthorityOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AcmpcaDeletePermissionFuture struct {
	Future workflow.Future
}

func (r *AcmpcaDeletePermissionFuture) Get(ctx workflow.Context) (*acmpca.DeletePermissionOutput, error) {
	var output acmpca.DeletePermissionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AcmpcaDeletePolicyFuture struct {
	Future workflow.Future
}

func (r *AcmpcaDeletePolicyFuture) Get(ctx workflow.Context) (*acmpca.DeletePolicyOutput, error) {
	var output acmpca.DeletePolicyOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AcmpcaDescribeCertificateAuthorityFuture struct {
	Future workflow.Future
}

func (r *AcmpcaDescribeCertificateAuthorityFuture) Get(ctx workflow.Context) (*acmpca.DescribeCertificateAuthorityOutput, error) {
	var output acmpca.DescribeCertificateAuthorityOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AcmpcaDescribeCertificateAuthorityAuditReportFuture struct {
	Future workflow.Future
}

func (r *AcmpcaDescribeCertificateAuthorityAuditReportFuture) Get(ctx workflow.Context) (*acmpca.DescribeCertificateAuthorityAuditReportOutput, error) {
	var output acmpca.DescribeCertificateAuthorityAuditReportOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AcmpcaGetCertificateFuture struct {
	Future workflow.Future
}

func (r *AcmpcaGetCertificateFuture) Get(ctx workflow.Context) (*acmpca.GetCertificateOutput, error) {
	var output acmpca.GetCertificateOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AcmpcaGetCertificateAuthorityCertificateFuture struct {
	Future workflow.Future
}

func (r *AcmpcaGetCertificateAuthorityCertificateFuture) Get(ctx workflow.Context) (*acmpca.GetCertificateAuthorityCertificateOutput, error) {
	var output acmpca.GetCertificateAuthorityCertificateOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AcmpcaGetCertificateAuthorityCsrFuture struct {
	Future workflow.Future
}

func (r *AcmpcaGetCertificateAuthorityCsrFuture) Get(ctx workflow.Context) (*acmpca.GetCertificateAuthorityCsrOutput, error) {
	var output acmpca.GetCertificateAuthorityCsrOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AcmpcaGetPolicyFuture struct {
	Future workflow.Future
}

func (r *AcmpcaGetPolicyFuture) Get(ctx workflow.Context) (*acmpca.GetPolicyOutput, error) {
	var output acmpca.GetPolicyOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AcmpcaImportCertificateAuthorityCertificateFuture struct {
	Future workflow.Future
}

func (r *AcmpcaImportCertificateAuthorityCertificateFuture) Get(ctx workflow.Context) (*acmpca.ImportCertificateAuthorityCertificateOutput, error) {
	var output acmpca.ImportCertificateAuthorityCertificateOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AcmpcaIssueCertificateFuture struct {
	Future workflow.Future
}

func (r *AcmpcaIssueCertificateFuture) Get(ctx workflow.Context) (*acmpca.IssueCertificateOutput, error) {
	var output acmpca.IssueCertificateOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AcmpcaListCertificateAuthoritiesFuture struct {
	Future workflow.Future
}

func (r *AcmpcaListCertificateAuthoritiesFuture) Get(ctx workflow.Context) (*acmpca.ListCertificateAuthoritiesOutput, error) {
	var output acmpca.ListCertificateAuthoritiesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AcmpcaListPermissionsFuture struct {
	Future workflow.Future
}

func (r *AcmpcaListPermissionsFuture) Get(ctx workflow.Context) (*acmpca.ListPermissionsOutput, error) {
	var output acmpca.ListPermissionsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AcmpcaListTagsFuture struct {
	Future workflow.Future
}

func (r *AcmpcaListTagsFuture) Get(ctx workflow.Context) (*acmpca.ListTagsOutput, error) {
	var output acmpca.ListTagsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AcmpcaPutPolicyFuture struct {
	Future workflow.Future
}

func (r *AcmpcaPutPolicyFuture) Get(ctx workflow.Context) (*acmpca.PutPolicyOutput, error) {
	var output acmpca.PutPolicyOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AcmpcaRestoreCertificateAuthorityFuture struct {
	Future workflow.Future
}

func (r *AcmpcaRestoreCertificateAuthorityFuture) Get(ctx workflow.Context) (*acmpca.RestoreCertificateAuthorityOutput, error) {
	var output acmpca.RestoreCertificateAuthorityOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AcmpcaRevokeCertificateFuture struct {
	Future workflow.Future
}

func (r *AcmpcaRevokeCertificateFuture) Get(ctx workflow.Context) (*acmpca.RevokeCertificateOutput, error) {
	var output acmpca.RevokeCertificateOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AcmpcaTagCertificateAuthorityFuture struct {
	Future workflow.Future
}

func (r *AcmpcaTagCertificateAuthorityFuture) Get(ctx workflow.Context) (*acmpca.TagCertificateAuthorityOutput, error) {
	var output acmpca.TagCertificateAuthorityOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AcmpcaUntagCertificateAuthorityFuture struct {
	Future workflow.Future
}

func (r *AcmpcaUntagCertificateAuthorityFuture) Get(ctx workflow.Context) (*acmpca.UntagCertificateAuthorityOutput, error) {
	var output acmpca.UntagCertificateAuthorityOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AcmpcaUpdateCertificateAuthorityFuture struct {
	Future workflow.Future
}

func (r *AcmpcaUpdateCertificateAuthorityFuture) Get(ctx workflow.Context) (*acmpca.UpdateCertificateAuthorityOutput, error) {
	var output acmpca.UpdateCertificateAuthorityOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *ACMPCAStub) CreateCertificateAuthority(ctx workflow.Context, input *acmpca.CreateCertificateAuthorityInput) (*acmpca.CreateCertificateAuthorityOutput, error) {
	var output acmpca.CreateCertificateAuthorityOutput
	err := workflow.ExecuteActivity(ctx, "aws.acmpca.CreateCertificateAuthority", input).Get(ctx, &output)
	return &output, err
}

func (a *ACMPCAStub) CreateCertificateAuthorityAsync(ctx workflow.Context, input *acmpca.CreateCertificateAuthorityInput) *AcmpcaCreateCertificateAuthorityFuture {
	future := workflow.ExecuteActivity(ctx, "aws.acmpca.CreateCertificateAuthority", input)
	return &AcmpcaCreateCertificateAuthorityFuture{Future: future}
}

func (a *ACMPCAStub) CreateCertificateAuthorityAuditReport(ctx workflow.Context, input *acmpca.CreateCertificateAuthorityAuditReportInput) (*acmpca.CreateCertificateAuthorityAuditReportOutput, error) {
	var output acmpca.CreateCertificateAuthorityAuditReportOutput
	err := workflow.ExecuteActivity(ctx, "aws.acmpca.CreateCertificateAuthorityAuditReport", input).Get(ctx, &output)
	return &output, err
}

func (a *ACMPCAStub) CreateCertificateAuthorityAuditReportAsync(ctx workflow.Context, input *acmpca.CreateCertificateAuthorityAuditReportInput) *AcmpcaCreateCertificateAuthorityAuditReportFuture {
	future := workflow.ExecuteActivity(ctx, "aws.acmpca.CreateCertificateAuthorityAuditReport", input)
	return &AcmpcaCreateCertificateAuthorityAuditReportFuture{Future: future}
}

func (a *ACMPCAStub) CreatePermission(ctx workflow.Context, input *acmpca.CreatePermissionInput) (*acmpca.CreatePermissionOutput, error) {
	var output acmpca.CreatePermissionOutput
	err := workflow.ExecuteActivity(ctx, "aws.acmpca.CreatePermission", input).Get(ctx, &output)
	return &output, err
}

func (a *ACMPCAStub) CreatePermissionAsync(ctx workflow.Context, input *acmpca.CreatePermissionInput) *AcmpcaCreatePermissionFuture {
	future := workflow.ExecuteActivity(ctx, "aws.acmpca.CreatePermission", input)
	return &AcmpcaCreatePermissionFuture{Future: future}
}

func (a *ACMPCAStub) DeleteCertificateAuthority(ctx workflow.Context, input *acmpca.DeleteCertificateAuthorityInput) (*acmpca.DeleteCertificateAuthorityOutput, error) {
	var output acmpca.DeleteCertificateAuthorityOutput
	err := workflow.ExecuteActivity(ctx, "aws.acmpca.DeleteCertificateAuthority", input).Get(ctx, &output)
	return &output, err
}

func (a *ACMPCAStub) DeleteCertificateAuthorityAsync(ctx workflow.Context, input *acmpca.DeleteCertificateAuthorityInput) *AcmpcaDeleteCertificateAuthorityFuture {
	future := workflow.ExecuteActivity(ctx, "aws.acmpca.DeleteCertificateAuthority", input)
	return &AcmpcaDeleteCertificateAuthorityFuture{Future: future}
}

func (a *ACMPCAStub) DeletePermission(ctx workflow.Context, input *acmpca.DeletePermissionInput) (*acmpca.DeletePermissionOutput, error) {
	var output acmpca.DeletePermissionOutput
	err := workflow.ExecuteActivity(ctx, "aws.acmpca.DeletePermission", input).Get(ctx, &output)
	return &output, err
}

func (a *ACMPCAStub) DeletePermissionAsync(ctx workflow.Context, input *acmpca.DeletePermissionInput) *AcmpcaDeletePermissionFuture {
	future := workflow.ExecuteActivity(ctx, "aws.acmpca.DeletePermission", input)
	return &AcmpcaDeletePermissionFuture{Future: future}
}

func (a *ACMPCAStub) DeletePolicy(ctx workflow.Context, input *acmpca.DeletePolicyInput) (*acmpca.DeletePolicyOutput, error) {
	var output acmpca.DeletePolicyOutput
	err := workflow.ExecuteActivity(ctx, "aws.acmpca.DeletePolicy", input).Get(ctx, &output)
	return &output, err
}

func (a *ACMPCAStub) DeletePolicyAsync(ctx workflow.Context, input *acmpca.DeletePolicyInput) *AcmpcaDeletePolicyFuture {
	future := workflow.ExecuteActivity(ctx, "aws.acmpca.DeletePolicy", input)
	return &AcmpcaDeletePolicyFuture{Future: future}
}

func (a *ACMPCAStub) DescribeCertificateAuthority(ctx workflow.Context, input *acmpca.DescribeCertificateAuthorityInput) (*acmpca.DescribeCertificateAuthorityOutput, error) {
	var output acmpca.DescribeCertificateAuthorityOutput
	err := workflow.ExecuteActivity(ctx, "aws.acmpca.DescribeCertificateAuthority", input).Get(ctx, &output)
	return &output, err
}

func (a *ACMPCAStub) DescribeCertificateAuthorityAsync(ctx workflow.Context, input *acmpca.DescribeCertificateAuthorityInput) *AcmpcaDescribeCertificateAuthorityFuture {
	future := workflow.ExecuteActivity(ctx, "aws.acmpca.DescribeCertificateAuthority", input)
	return &AcmpcaDescribeCertificateAuthorityFuture{Future: future}
}

func (a *ACMPCAStub) DescribeCertificateAuthorityAuditReport(ctx workflow.Context, input *acmpca.DescribeCertificateAuthorityAuditReportInput) (*acmpca.DescribeCertificateAuthorityAuditReportOutput, error) {
	var output acmpca.DescribeCertificateAuthorityAuditReportOutput
	err := workflow.ExecuteActivity(ctx, "aws.acmpca.DescribeCertificateAuthorityAuditReport", input).Get(ctx, &output)
	return &output, err
}

func (a *ACMPCAStub) DescribeCertificateAuthorityAuditReportAsync(ctx workflow.Context, input *acmpca.DescribeCertificateAuthorityAuditReportInput) *AcmpcaDescribeCertificateAuthorityAuditReportFuture {
	future := workflow.ExecuteActivity(ctx, "aws.acmpca.DescribeCertificateAuthorityAuditReport", input)
	return &AcmpcaDescribeCertificateAuthorityAuditReportFuture{Future: future}
}

func (a *ACMPCAStub) GetCertificate(ctx workflow.Context, input *acmpca.GetCertificateInput) (*acmpca.GetCertificateOutput, error) {
	var output acmpca.GetCertificateOutput
	err := workflow.ExecuteActivity(ctx, "aws.acmpca.GetCertificate", input).Get(ctx, &output)
	return &output, err
}

func (a *ACMPCAStub) GetCertificateAsync(ctx workflow.Context, input *acmpca.GetCertificateInput) *AcmpcaGetCertificateFuture {
	future := workflow.ExecuteActivity(ctx, "aws.acmpca.GetCertificate", input)
	return &AcmpcaGetCertificateFuture{Future: future}
}

func (a *ACMPCAStub) GetCertificateAuthorityCertificate(ctx workflow.Context, input *acmpca.GetCertificateAuthorityCertificateInput) (*acmpca.GetCertificateAuthorityCertificateOutput, error) {
	var output acmpca.GetCertificateAuthorityCertificateOutput
	err := workflow.ExecuteActivity(ctx, "aws.acmpca.GetCertificateAuthorityCertificate", input).Get(ctx, &output)
	return &output, err
}

func (a *ACMPCAStub) GetCertificateAuthorityCertificateAsync(ctx workflow.Context, input *acmpca.GetCertificateAuthorityCertificateInput) *AcmpcaGetCertificateAuthorityCertificateFuture {
	future := workflow.ExecuteActivity(ctx, "aws.acmpca.GetCertificateAuthorityCertificate", input)
	return &AcmpcaGetCertificateAuthorityCertificateFuture{Future: future}
}

func (a *ACMPCAStub) GetCertificateAuthorityCsr(ctx workflow.Context, input *acmpca.GetCertificateAuthorityCsrInput) (*acmpca.GetCertificateAuthorityCsrOutput, error) {
	var output acmpca.GetCertificateAuthorityCsrOutput
	err := workflow.ExecuteActivity(ctx, "aws.acmpca.GetCertificateAuthorityCsr", input).Get(ctx, &output)
	return &output, err
}

func (a *ACMPCAStub) GetCertificateAuthorityCsrAsync(ctx workflow.Context, input *acmpca.GetCertificateAuthorityCsrInput) *AcmpcaGetCertificateAuthorityCsrFuture {
	future := workflow.ExecuteActivity(ctx, "aws.acmpca.GetCertificateAuthorityCsr", input)
	return &AcmpcaGetCertificateAuthorityCsrFuture{Future: future}
}

func (a *ACMPCAStub) GetPolicy(ctx workflow.Context, input *acmpca.GetPolicyInput) (*acmpca.GetPolicyOutput, error) {
	var output acmpca.GetPolicyOutput
	err := workflow.ExecuteActivity(ctx, "aws.acmpca.GetPolicy", input).Get(ctx, &output)
	return &output, err
}

func (a *ACMPCAStub) GetPolicyAsync(ctx workflow.Context, input *acmpca.GetPolicyInput) *AcmpcaGetPolicyFuture {
	future := workflow.ExecuteActivity(ctx, "aws.acmpca.GetPolicy", input)
	return &AcmpcaGetPolicyFuture{Future: future}
}

func (a *ACMPCAStub) ImportCertificateAuthorityCertificate(ctx workflow.Context, input *acmpca.ImportCertificateAuthorityCertificateInput) (*acmpca.ImportCertificateAuthorityCertificateOutput, error) {
	var output acmpca.ImportCertificateAuthorityCertificateOutput
	err := workflow.ExecuteActivity(ctx, "aws.acmpca.ImportCertificateAuthorityCertificate", input).Get(ctx, &output)
	return &output, err
}

func (a *ACMPCAStub) ImportCertificateAuthorityCertificateAsync(ctx workflow.Context, input *acmpca.ImportCertificateAuthorityCertificateInput) *AcmpcaImportCertificateAuthorityCertificateFuture {
	future := workflow.ExecuteActivity(ctx, "aws.acmpca.ImportCertificateAuthorityCertificate", input)
	return &AcmpcaImportCertificateAuthorityCertificateFuture{Future: future}
}

func (a *ACMPCAStub) IssueCertificate(ctx workflow.Context, input *acmpca.IssueCertificateInput) (*acmpca.IssueCertificateOutput, error) {
	var output acmpca.IssueCertificateOutput
	err := workflow.ExecuteActivity(ctx, "aws.acmpca.IssueCertificate", input).Get(ctx, &output)
	return &output, err
}

func (a *ACMPCAStub) IssueCertificateAsync(ctx workflow.Context, input *acmpca.IssueCertificateInput) *AcmpcaIssueCertificateFuture {
	future := workflow.ExecuteActivity(ctx, "aws.acmpca.IssueCertificate", input)
	return &AcmpcaIssueCertificateFuture{Future: future}
}

func (a *ACMPCAStub) ListCertificateAuthorities(ctx workflow.Context, input *acmpca.ListCertificateAuthoritiesInput) (*acmpca.ListCertificateAuthoritiesOutput, error) {
	var output acmpca.ListCertificateAuthoritiesOutput
	err := workflow.ExecuteActivity(ctx, "aws.acmpca.ListCertificateAuthorities", input).Get(ctx, &output)
	return &output, err
}

func (a *ACMPCAStub) ListCertificateAuthoritiesAsync(ctx workflow.Context, input *acmpca.ListCertificateAuthoritiesInput) *AcmpcaListCertificateAuthoritiesFuture {
	future := workflow.ExecuteActivity(ctx, "aws.acmpca.ListCertificateAuthorities", input)
	return &AcmpcaListCertificateAuthoritiesFuture{Future: future}
}

func (a *ACMPCAStub) ListPermissions(ctx workflow.Context, input *acmpca.ListPermissionsInput) (*acmpca.ListPermissionsOutput, error) {
	var output acmpca.ListPermissionsOutput
	err := workflow.ExecuteActivity(ctx, "aws.acmpca.ListPermissions", input).Get(ctx, &output)
	return &output, err
}

func (a *ACMPCAStub) ListPermissionsAsync(ctx workflow.Context, input *acmpca.ListPermissionsInput) *AcmpcaListPermissionsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.acmpca.ListPermissions", input)
	return &AcmpcaListPermissionsFuture{Future: future}
}

func (a *ACMPCAStub) ListTags(ctx workflow.Context, input *acmpca.ListTagsInput) (*acmpca.ListTagsOutput, error) {
	var output acmpca.ListTagsOutput
	err := workflow.ExecuteActivity(ctx, "aws.acmpca.ListTags", input).Get(ctx, &output)
	return &output, err
}

func (a *ACMPCAStub) ListTagsAsync(ctx workflow.Context, input *acmpca.ListTagsInput) *AcmpcaListTagsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.acmpca.ListTags", input)
	return &AcmpcaListTagsFuture{Future: future}
}

func (a *ACMPCAStub) PutPolicy(ctx workflow.Context, input *acmpca.PutPolicyInput) (*acmpca.PutPolicyOutput, error) {
	var output acmpca.PutPolicyOutput
	err := workflow.ExecuteActivity(ctx, "aws.acmpca.PutPolicy", input).Get(ctx, &output)
	return &output, err
}

func (a *ACMPCAStub) PutPolicyAsync(ctx workflow.Context, input *acmpca.PutPolicyInput) *AcmpcaPutPolicyFuture {
	future := workflow.ExecuteActivity(ctx, "aws.acmpca.PutPolicy", input)
	return &AcmpcaPutPolicyFuture{Future: future}
}

func (a *ACMPCAStub) RestoreCertificateAuthority(ctx workflow.Context, input *acmpca.RestoreCertificateAuthorityInput) (*acmpca.RestoreCertificateAuthorityOutput, error) {
	var output acmpca.RestoreCertificateAuthorityOutput
	err := workflow.ExecuteActivity(ctx, "aws.acmpca.RestoreCertificateAuthority", input).Get(ctx, &output)
	return &output, err
}

func (a *ACMPCAStub) RestoreCertificateAuthorityAsync(ctx workflow.Context, input *acmpca.RestoreCertificateAuthorityInput) *AcmpcaRestoreCertificateAuthorityFuture {
	future := workflow.ExecuteActivity(ctx, "aws.acmpca.RestoreCertificateAuthority", input)
	return &AcmpcaRestoreCertificateAuthorityFuture{Future: future}
}

func (a *ACMPCAStub) RevokeCertificate(ctx workflow.Context, input *acmpca.RevokeCertificateInput) (*acmpca.RevokeCertificateOutput, error) {
	var output acmpca.RevokeCertificateOutput
	err := workflow.ExecuteActivity(ctx, "aws.acmpca.RevokeCertificate", input).Get(ctx, &output)
	return &output, err
}

func (a *ACMPCAStub) RevokeCertificateAsync(ctx workflow.Context, input *acmpca.RevokeCertificateInput) *AcmpcaRevokeCertificateFuture {
	future := workflow.ExecuteActivity(ctx, "aws.acmpca.RevokeCertificate", input)
	return &AcmpcaRevokeCertificateFuture{Future: future}
}

func (a *ACMPCAStub) TagCertificateAuthority(ctx workflow.Context, input *acmpca.TagCertificateAuthorityInput) (*acmpca.TagCertificateAuthorityOutput, error) {
	var output acmpca.TagCertificateAuthorityOutput
	err := workflow.ExecuteActivity(ctx, "aws.acmpca.TagCertificateAuthority", input).Get(ctx, &output)
	return &output, err
}

func (a *ACMPCAStub) TagCertificateAuthorityAsync(ctx workflow.Context, input *acmpca.TagCertificateAuthorityInput) *AcmpcaTagCertificateAuthorityFuture {
	future := workflow.ExecuteActivity(ctx, "aws.acmpca.TagCertificateAuthority", input)
	return &AcmpcaTagCertificateAuthorityFuture{Future: future}
}

func (a *ACMPCAStub) UntagCertificateAuthority(ctx workflow.Context, input *acmpca.UntagCertificateAuthorityInput) (*acmpca.UntagCertificateAuthorityOutput, error) {
	var output acmpca.UntagCertificateAuthorityOutput
	err := workflow.ExecuteActivity(ctx, "aws.acmpca.UntagCertificateAuthority", input).Get(ctx, &output)
	return &output, err
}

func (a *ACMPCAStub) UntagCertificateAuthorityAsync(ctx workflow.Context, input *acmpca.UntagCertificateAuthorityInput) *AcmpcaUntagCertificateAuthorityFuture {
	future := workflow.ExecuteActivity(ctx, "aws.acmpca.UntagCertificateAuthority", input)
	return &AcmpcaUntagCertificateAuthorityFuture{Future: future}
}

func (a *ACMPCAStub) UpdateCertificateAuthority(ctx workflow.Context, input *acmpca.UpdateCertificateAuthorityInput) (*acmpca.UpdateCertificateAuthorityOutput, error) {
	var output acmpca.UpdateCertificateAuthorityOutput
	err := workflow.ExecuteActivity(ctx, "aws.acmpca.UpdateCertificateAuthority", input).Get(ctx, &output)
	return &output, err
}

func (a *ACMPCAStub) UpdateCertificateAuthorityAsync(ctx workflow.Context, input *acmpca.UpdateCertificateAuthorityInput) *AcmpcaUpdateCertificateAuthorityFuture {
	future := workflow.ExecuteActivity(ctx, "aws.acmpca.UpdateCertificateAuthority", input)
	return &AcmpcaUpdateCertificateAuthorityFuture{Future: future}
}

func (a *ACMPCAStub) WaitUntilAuditReportCreated(ctx workflow.Context, input *acmpca.DescribeCertificateAuthorityAuditReportInput) error {
	return workflow.ExecuteActivity(ctx, "aws.acmpca.WaitUntilAuditReportCreated", input).Get(ctx, nil)
}

func (a *ACMPCAStub) WaitUntilAuditReportCreatedAsync(ctx workflow.Context, input *acmpca.DescribeCertificateAuthorityAuditReportInput) *VoidFuture {
	future := workflow.ExecuteActivity(ctx, "aws.acmpca.WaitUntilAuditReportCreated", input)
	return NewVoidFuture(future)
}

func (a *ACMPCAStub) WaitUntilCertificateAuthorityCSRCreated(ctx workflow.Context, input *acmpca.GetCertificateAuthorityCsrInput) error {
	return workflow.ExecuteActivity(ctx, "aws.acmpca.WaitUntilCertificateAuthorityCSRCreated", input).Get(ctx, nil)
}

func (a *ACMPCAStub) WaitUntilCertificateAuthorityCSRCreatedAsync(ctx workflow.Context, input *acmpca.GetCertificateAuthorityCsrInput) *VoidFuture {
	future := workflow.ExecuteActivity(ctx, "aws.acmpca.WaitUntilCertificateAuthorityCSRCreated", input)
	return NewVoidFuture(future)
}

func (a *ACMPCAStub) WaitUntilCertificateIssued(ctx workflow.Context, input *acmpca.GetCertificateInput) error {
	return workflow.ExecuteActivity(ctx, "aws.acmpca.WaitUntilCertificateIssued", input).Get(ctx, nil)
}

func (a *ACMPCAStub) WaitUntilCertificateIssuedAsync(ctx workflow.Context, input *acmpca.GetCertificateInput) *VoidFuture {
	future := workflow.ExecuteActivity(ctx, "aws.acmpca.WaitUntilCertificateIssued", input)
	return NewVoidFuture(future)
}
