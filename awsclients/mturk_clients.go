package awsclients

import (
	"github.com/aws/aws-sdk-go/service/mturk"
	"go.temporal.io/sdk/workflow"
	"temporal.io/aws-sdk/awsactivities"
)

type MTurkClient interface {
       ApproveAssignment(ctx workflow.Context, input *mturk.ApproveAssignmentInput) (*mturk.ApproveAssignmentOutput, error)
       ApproveAssignmentAsync(ctx workflow.Context, input *mturk.ApproveAssignmentInput) *MturkApproveAssignmentResult

       AssociateQualificationWithWorker(ctx workflow.Context, input *mturk.AssociateQualificationWithWorkerInput) (*mturk.AssociateQualificationWithWorkerOutput, error)
       AssociateQualificationWithWorkerAsync(ctx workflow.Context, input *mturk.AssociateQualificationWithWorkerInput) *MturkAssociateQualificationWithWorkerResult

       CreateAdditionalAssignmentsForHIT(ctx workflow.Context, input *mturk.CreateAdditionalAssignmentsForHITInput) (*mturk.CreateAdditionalAssignmentsForHITOutput, error)
       CreateAdditionalAssignmentsForHITAsync(ctx workflow.Context, input *mturk.CreateAdditionalAssignmentsForHITInput) *MturkCreateAdditionalAssignmentsForHITResult

       CreateHIT(ctx workflow.Context, input *mturk.CreateHITInput) (*mturk.CreateHITOutput, error)
       CreateHITAsync(ctx workflow.Context, input *mturk.CreateHITInput) *MturkCreateHITResult

       CreateHITType(ctx workflow.Context, input *mturk.CreateHITTypeInput) (*mturk.CreateHITTypeOutput, error)
       CreateHITTypeAsync(ctx workflow.Context, input *mturk.CreateHITTypeInput) *MturkCreateHITTypeResult

       CreateHITWithHITType(ctx workflow.Context, input *mturk.CreateHITWithHITTypeInput) (*mturk.CreateHITWithHITTypeOutput, error)
       CreateHITWithHITTypeAsync(ctx workflow.Context, input *mturk.CreateHITWithHITTypeInput) *MturkCreateHITWithHITTypeResult

       CreateQualificationType(ctx workflow.Context, input *mturk.CreateQualificationTypeInput) (*mturk.CreateQualificationTypeOutput, error)
       CreateQualificationTypeAsync(ctx workflow.Context, input *mturk.CreateQualificationTypeInput) *MturkCreateQualificationTypeResult

       CreateWorkerBlock(ctx workflow.Context, input *mturk.CreateWorkerBlockInput) (*mturk.CreateWorkerBlockOutput, error)
       CreateWorkerBlockAsync(ctx workflow.Context, input *mturk.CreateWorkerBlockInput) *MturkCreateWorkerBlockResult

       DeleteHIT(ctx workflow.Context, input *mturk.DeleteHITInput) (*mturk.DeleteHITOutput, error)
       DeleteHITAsync(ctx workflow.Context, input *mturk.DeleteHITInput) *MturkDeleteHITResult

       DeleteQualificationType(ctx workflow.Context, input *mturk.DeleteQualificationTypeInput) (*mturk.DeleteQualificationTypeOutput, error)
       DeleteQualificationTypeAsync(ctx workflow.Context, input *mturk.DeleteQualificationTypeInput) *MturkDeleteQualificationTypeResult

       DeleteWorkerBlock(ctx workflow.Context, input *mturk.DeleteWorkerBlockInput) (*mturk.DeleteWorkerBlockOutput, error)
       DeleteWorkerBlockAsync(ctx workflow.Context, input *mturk.DeleteWorkerBlockInput) *MturkDeleteWorkerBlockResult

       DisassociateQualificationFromWorker(ctx workflow.Context, input *mturk.DisassociateQualificationFromWorkerInput) (*mturk.DisassociateQualificationFromWorkerOutput, error)
       DisassociateQualificationFromWorkerAsync(ctx workflow.Context, input *mturk.DisassociateQualificationFromWorkerInput) *MturkDisassociateQualificationFromWorkerResult

       GetAccountBalance(ctx workflow.Context, input *mturk.GetAccountBalanceInput) (*mturk.GetAccountBalanceOutput, error)
       GetAccountBalanceAsync(ctx workflow.Context, input *mturk.GetAccountBalanceInput) *MturkGetAccountBalanceResult

       GetAssignment(ctx workflow.Context, input *mturk.GetAssignmentInput) (*mturk.GetAssignmentOutput, error)
       GetAssignmentAsync(ctx workflow.Context, input *mturk.GetAssignmentInput) *MturkGetAssignmentResult

       GetFileUploadURL(ctx workflow.Context, input *mturk.GetFileUploadURLInput) (*mturk.GetFileUploadURLOutput, error)
       GetFileUploadURLAsync(ctx workflow.Context, input *mturk.GetFileUploadURLInput) *MturkGetFileUploadURLResult

       GetHIT(ctx workflow.Context, input *mturk.GetHITInput) (*mturk.GetHITOutput, error)
       GetHITAsync(ctx workflow.Context, input *mturk.GetHITInput) *MturkGetHITResult

       GetQualificationScore(ctx workflow.Context, input *mturk.GetQualificationScoreInput) (*mturk.GetQualificationScoreOutput, error)
       GetQualificationScoreAsync(ctx workflow.Context, input *mturk.GetQualificationScoreInput) *MturkGetQualificationScoreResult

       GetQualificationType(ctx workflow.Context, input *mturk.GetQualificationTypeInput) (*mturk.GetQualificationTypeOutput, error)
       GetQualificationTypeAsync(ctx workflow.Context, input *mturk.GetQualificationTypeInput) *MturkGetQualificationTypeResult

       ListAssignmentsForHIT(ctx workflow.Context, input *mturk.ListAssignmentsForHITInput) (*mturk.ListAssignmentsForHITOutput, error)
       ListAssignmentsForHITAsync(ctx workflow.Context, input *mturk.ListAssignmentsForHITInput) *MturkListAssignmentsForHITResult

       ListBonusPayments(ctx workflow.Context, input *mturk.ListBonusPaymentsInput) (*mturk.ListBonusPaymentsOutput, error)
       ListBonusPaymentsAsync(ctx workflow.Context, input *mturk.ListBonusPaymentsInput) *MturkListBonusPaymentsResult

       ListHITs(ctx workflow.Context, input *mturk.ListHITsInput) (*mturk.ListHITsOutput, error)
       ListHITsAsync(ctx workflow.Context, input *mturk.ListHITsInput) *MturkListHITsResult

       ListHITsForQualificationType(ctx workflow.Context, input *mturk.ListHITsForQualificationTypeInput) (*mturk.ListHITsForQualificationTypeOutput, error)
       ListHITsForQualificationTypeAsync(ctx workflow.Context, input *mturk.ListHITsForQualificationTypeInput) *MturkListHITsForQualificationTypeResult

       ListQualificationRequests(ctx workflow.Context, input *mturk.ListQualificationRequestsInput) (*mturk.ListQualificationRequestsOutput, error)
       ListQualificationRequestsAsync(ctx workflow.Context, input *mturk.ListQualificationRequestsInput) *MturkListQualificationRequestsResult

       ListQualificationTypes(ctx workflow.Context, input *mturk.ListQualificationTypesInput) (*mturk.ListQualificationTypesOutput, error)
       ListQualificationTypesAsync(ctx workflow.Context, input *mturk.ListQualificationTypesInput) *MturkListQualificationTypesResult

       ListReviewPolicyResultsForHIT(ctx workflow.Context, input *mturk.ListReviewPolicyResultsForHITInput) (*mturk.ListReviewPolicyResultsForHITOutput, error)
       ListReviewPolicyResultsForHITAsync(ctx workflow.Context, input *mturk.ListReviewPolicyResultsForHITInput) *MturkListReviewPolicyResultsForHITResult

       ListReviewableHITs(ctx workflow.Context, input *mturk.ListReviewableHITsInput) (*mturk.ListReviewableHITsOutput, error)
       ListReviewableHITsAsync(ctx workflow.Context, input *mturk.ListReviewableHITsInput) *MturkListReviewableHITsResult

       ListWorkerBlocks(ctx workflow.Context, input *mturk.ListWorkerBlocksInput) (*mturk.ListWorkerBlocksOutput, error)
       ListWorkerBlocksAsync(ctx workflow.Context, input *mturk.ListWorkerBlocksInput) *MturkListWorkerBlocksResult

       ListWorkersWithQualificationType(ctx workflow.Context, input *mturk.ListWorkersWithQualificationTypeInput) (*mturk.ListWorkersWithQualificationTypeOutput, error)
       ListWorkersWithQualificationTypeAsync(ctx workflow.Context, input *mturk.ListWorkersWithQualificationTypeInput) *MturkListWorkersWithQualificationTypeResult

       NotifyWorkers(ctx workflow.Context, input *mturk.NotifyWorkersInput) (*mturk.NotifyWorkersOutput, error)
       NotifyWorkersAsync(ctx workflow.Context, input *mturk.NotifyWorkersInput) *MturkNotifyWorkersResult

       RejectAssignment(ctx workflow.Context, input *mturk.RejectAssignmentInput) (*mturk.RejectAssignmentOutput, error)
       RejectAssignmentAsync(ctx workflow.Context, input *mturk.RejectAssignmentInput) *MturkRejectAssignmentResult

       SendBonus(ctx workflow.Context, input *mturk.SendBonusInput) (*mturk.SendBonusOutput, error)
       SendBonusAsync(ctx workflow.Context, input *mturk.SendBonusInput) *MturkSendBonusResult

       SendTestEventNotification(ctx workflow.Context, input *mturk.SendTestEventNotificationInput) (*mturk.SendTestEventNotificationOutput, error)
       SendTestEventNotificationAsync(ctx workflow.Context, input *mturk.SendTestEventNotificationInput) *MturkSendTestEventNotificationResult

       UpdateExpirationForHIT(ctx workflow.Context, input *mturk.UpdateExpirationForHITInput) (*mturk.UpdateExpirationForHITOutput, error)
       UpdateExpirationForHITAsync(ctx workflow.Context, input *mturk.UpdateExpirationForHITInput) *MturkUpdateExpirationForHITResult

       UpdateHITReviewStatus(ctx workflow.Context, input *mturk.UpdateHITReviewStatusInput) (*mturk.UpdateHITReviewStatusOutput, error)
       UpdateHITReviewStatusAsync(ctx workflow.Context, input *mturk.UpdateHITReviewStatusInput) *MturkUpdateHITReviewStatusResult

       UpdateHITTypeOfHIT(ctx workflow.Context, input *mturk.UpdateHITTypeOfHITInput) (*mturk.UpdateHITTypeOfHITOutput, error)
       UpdateHITTypeOfHITAsync(ctx workflow.Context, input *mturk.UpdateHITTypeOfHITInput) *MturkUpdateHITTypeOfHITResult

       UpdateNotificationSettings(ctx workflow.Context, input *mturk.UpdateNotificationSettingsInput) (*mturk.UpdateNotificationSettingsOutput, error)
       UpdateNotificationSettingsAsync(ctx workflow.Context, input *mturk.UpdateNotificationSettingsInput) *MturkUpdateNotificationSettingsResult

       UpdateQualificationType(ctx workflow.Context, input *mturk.UpdateQualificationTypeInput) (*mturk.UpdateQualificationTypeOutput, error)
       UpdateQualificationTypeAsync(ctx workflow.Context, input *mturk.UpdateQualificationTypeInput) *MturkUpdateQualificationTypeResult
}

type MturkApproveAssignmentResult struct {
	Result workflow.Future
}

func (r *MturkApproveAssignmentResult) Get(ctx workflow.Context) (*mturk.ApproveAssignmentOutput, error) {
    var output mturk.ApproveAssignmentOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MturkAssociateQualificationWithWorkerResult struct {
	Result workflow.Future
}

func (r *MturkAssociateQualificationWithWorkerResult) Get(ctx workflow.Context) (*mturk.AssociateQualificationWithWorkerOutput, error) {
    var output mturk.AssociateQualificationWithWorkerOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MturkCreateAdditionalAssignmentsForHITResult struct {
	Result workflow.Future
}

func (r *MturkCreateAdditionalAssignmentsForHITResult) Get(ctx workflow.Context) (*mturk.CreateAdditionalAssignmentsForHITOutput, error) {
    var output mturk.CreateAdditionalAssignmentsForHITOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MturkCreateHITResult struct {
	Result workflow.Future
}

func (r *MturkCreateHITResult) Get(ctx workflow.Context) (*mturk.CreateHITOutput, error) {
    var output mturk.CreateHITOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MturkCreateHITTypeResult struct {
	Result workflow.Future
}

func (r *MturkCreateHITTypeResult) Get(ctx workflow.Context) (*mturk.CreateHITTypeOutput, error) {
    var output mturk.CreateHITTypeOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MturkCreateHITWithHITTypeResult struct {
	Result workflow.Future
}

func (r *MturkCreateHITWithHITTypeResult) Get(ctx workflow.Context) (*mturk.CreateHITWithHITTypeOutput, error) {
    var output mturk.CreateHITWithHITTypeOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MturkCreateQualificationTypeResult struct {
	Result workflow.Future
}

func (r *MturkCreateQualificationTypeResult) Get(ctx workflow.Context) (*mturk.CreateQualificationTypeOutput, error) {
    var output mturk.CreateQualificationTypeOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MturkCreateWorkerBlockResult struct {
	Result workflow.Future
}

func (r *MturkCreateWorkerBlockResult) Get(ctx workflow.Context) (*mturk.CreateWorkerBlockOutput, error) {
    var output mturk.CreateWorkerBlockOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MturkDeleteHITResult struct {
	Result workflow.Future
}

func (r *MturkDeleteHITResult) Get(ctx workflow.Context) (*mturk.DeleteHITOutput, error) {
    var output mturk.DeleteHITOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MturkDeleteQualificationTypeResult struct {
	Result workflow.Future
}

func (r *MturkDeleteQualificationTypeResult) Get(ctx workflow.Context) (*mturk.DeleteQualificationTypeOutput, error) {
    var output mturk.DeleteQualificationTypeOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MturkDeleteWorkerBlockResult struct {
	Result workflow.Future
}

func (r *MturkDeleteWorkerBlockResult) Get(ctx workflow.Context) (*mturk.DeleteWorkerBlockOutput, error) {
    var output mturk.DeleteWorkerBlockOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MturkDisassociateQualificationFromWorkerResult struct {
	Result workflow.Future
}

func (r *MturkDisassociateQualificationFromWorkerResult) Get(ctx workflow.Context) (*mturk.DisassociateQualificationFromWorkerOutput, error) {
    var output mturk.DisassociateQualificationFromWorkerOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MturkGetAccountBalanceResult struct {
	Result workflow.Future
}

func (r *MturkGetAccountBalanceResult) Get(ctx workflow.Context) (*mturk.GetAccountBalanceOutput, error) {
    var output mturk.GetAccountBalanceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MturkGetAssignmentResult struct {
	Result workflow.Future
}

func (r *MturkGetAssignmentResult) Get(ctx workflow.Context) (*mturk.GetAssignmentOutput, error) {
    var output mturk.GetAssignmentOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MturkGetFileUploadURLResult struct {
	Result workflow.Future
}

func (r *MturkGetFileUploadURLResult) Get(ctx workflow.Context) (*mturk.GetFileUploadURLOutput, error) {
    var output mturk.GetFileUploadURLOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MturkGetHITResult struct {
	Result workflow.Future
}

func (r *MturkGetHITResult) Get(ctx workflow.Context) (*mturk.GetHITOutput, error) {
    var output mturk.GetHITOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MturkGetQualificationScoreResult struct {
	Result workflow.Future
}

func (r *MturkGetQualificationScoreResult) Get(ctx workflow.Context) (*mturk.GetQualificationScoreOutput, error) {
    var output mturk.GetQualificationScoreOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MturkGetQualificationTypeResult struct {
	Result workflow.Future
}

func (r *MturkGetQualificationTypeResult) Get(ctx workflow.Context) (*mturk.GetQualificationTypeOutput, error) {
    var output mturk.GetQualificationTypeOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MturkListAssignmentsForHITResult struct {
	Result workflow.Future
}

func (r *MturkListAssignmentsForHITResult) Get(ctx workflow.Context) (*mturk.ListAssignmentsForHITOutput, error) {
    var output mturk.ListAssignmentsForHITOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MturkListBonusPaymentsResult struct {
	Result workflow.Future
}

func (r *MturkListBonusPaymentsResult) Get(ctx workflow.Context) (*mturk.ListBonusPaymentsOutput, error) {
    var output mturk.ListBonusPaymentsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MturkListHITsResult struct {
	Result workflow.Future
}

func (r *MturkListHITsResult) Get(ctx workflow.Context) (*mturk.ListHITsOutput, error) {
    var output mturk.ListHITsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MturkListHITsForQualificationTypeResult struct {
	Result workflow.Future
}

func (r *MturkListHITsForQualificationTypeResult) Get(ctx workflow.Context) (*mturk.ListHITsForQualificationTypeOutput, error) {
    var output mturk.ListHITsForQualificationTypeOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MturkListQualificationRequestsResult struct {
	Result workflow.Future
}

func (r *MturkListQualificationRequestsResult) Get(ctx workflow.Context) (*mturk.ListQualificationRequestsOutput, error) {
    var output mturk.ListQualificationRequestsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MturkListQualificationTypesResult struct {
	Result workflow.Future
}

func (r *MturkListQualificationTypesResult) Get(ctx workflow.Context) (*mturk.ListQualificationTypesOutput, error) {
    var output mturk.ListQualificationTypesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MturkListReviewPolicyResultsForHITResult struct {
	Result workflow.Future
}

func (r *MturkListReviewPolicyResultsForHITResult) Get(ctx workflow.Context) (*mturk.ListReviewPolicyResultsForHITOutput, error) {
    var output mturk.ListReviewPolicyResultsForHITOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MturkListReviewableHITsResult struct {
	Result workflow.Future
}

func (r *MturkListReviewableHITsResult) Get(ctx workflow.Context) (*mturk.ListReviewableHITsOutput, error) {
    var output mturk.ListReviewableHITsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MturkListWorkerBlocksResult struct {
	Result workflow.Future
}

func (r *MturkListWorkerBlocksResult) Get(ctx workflow.Context) (*mturk.ListWorkerBlocksOutput, error) {
    var output mturk.ListWorkerBlocksOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MturkListWorkersWithQualificationTypeResult struct {
	Result workflow.Future
}

func (r *MturkListWorkersWithQualificationTypeResult) Get(ctx workflow.Context) (*mturk.ListWorkersWithQualificationTypeOutput, error) {
    var output mturk.ListWorkersWithQualificationTypeOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MturkNotifyWorkersResult struct {
	Result workflow.Future
}

func (r *MturkNotifyWorkersResult) Get(ctx workflow.Context) (*mturk.NotifyWorkersOutput, error) {
    var output mturk.NotifyWorkersOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MturkRejectAssignmentResult struct {
	Result workflow.Future
}

func (r *MturkRejectAssignmentResult) Get(ctx workflow.Context) (*mturk.RejectAssignmentOutput, error) {
    var output mturk.RejectAssignmentOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MturkSendBonusResult struct {
	Result workflow.Future
}

func (r *MturkSendBonusResult) Get(ctx workflow.Context) (*mturk.SendBonusOutput, error) {
    var output mturk.SendBonusOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MturkSendTestEventNotificationResult struct {
	Result workflow.Future
}

func (r *MturkSendTestEventNotificationResult) Get(ctx workflow.Context) (*mturk.SendTestEventNotificationOutput, error) {
    var output mturk.SendTestEventNotificationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MturkUpdateExpirationForHITResult struct {
	Result workflow.Future
}

func (r *MturkUpdateExpirationForHITResult) Get(ctx workflow.Context) (*mturk.UpdateExpirationForHITOutput, error) {
    var output mturk.UpdateExpirationForHITOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MturkUpdateHITReviewStatusResult struct {
	Result workflow.Future
}

func (r *MturkUpdateHITReviewStatusResult) Get(ctx workflow.Context) (*mturk.UpdateHITReviewStatusOutput, error) {
    var output mturk.UpdateHITReviewStatusOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MturkUpdateHITTypeOfHITResult struct {
	Result workflow.Future
}

func (r *MturkUpdateHITTypeOfHITResult) Get(ctx workflow.Context) (*mturk.UpdateHITTypeOfHITOutput, error) {
    var output mturk.UpdateHITTypeOfHITOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MturkUpdateNotificationSettingsResult struct {
	Result workflow.Future
}

func (r *MturkUpdateNotificationSettingsResult) Get(ctx workflow.Context) (*mturk.UpdateNotificationSettingsOutput, error) {
    var output mturk.UpdateNotificationSettingsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MturkUpdateQualificationTypeResult struct {
	Result workflow.Future
}

func (r *MturkUpdateQualificationTypeResult) Get(ctx workflow.Context) (*mturk.UpdateQualificationTypeOutput, error) {
    var output mturk.UpdateQualificationTypeOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MTurkStub struct {
    activities awsactivities.MTurkActivities
}

func NewMTurkStub() MTurkClient {
    return &MTurkStub{}
}

func (a *MTurkStub) ApproveAssignment(ctx workflow.Context, input *mturk.ApproveAssignmentInput) (*mturk.ApproveAssignmentOutput, error) {
    var output mturk.ApproveAssignmentOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ApproveAssignment, input).Get(ctx, &output)
    return &output, err
}

func (a *MTurkStub) ApproveAssignmentAsync(ctx workflow.Context, input *mturk.ApproveAssignmentInput) *MturkApproveAssignmentResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ApproveAssignment, input)
    return &MturkApproveAssignmentResult{Result: future}
}

func (a *MTurkStub) AssociateQualificationWithWorker(ctx workflow.Context, input *mturk.AssociateQualificationWithWorkerInput) (*mturk.AssociateQualificationWithWorkerOutput, error) {
    var output mturk.AssociateQualificationWithWorkerOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AssociateQualificationWithWorker, input).Get(ctx, &output)
    return &output, err
}

func (a *MTurkStub) AssociateQualificationWithWorkerAsync(ctx workflow.Context, input *mturk.AssociateQualificationWithWorkerInput) *MturkAssociateQualificationWithWorkerResult {
    future := workflow.ExecuteActivity(ctx, a.activities.AssociateQualificationWithWorker, input)
    return &MturkAssociateQualificationWithWorkerResult{Result: future}
}

func (a *MTurkStub) CreateAdditionalAssignmentsForHIT(ctx workflow.Context, input *mturk.CreateAdditionalAssignmentsForHITInput) (*mturk.CreateAdditionalAssignmentsForHITOutput, error) {
    var output mturk.CreateAdditionalAssignmentsForHITOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateAdditionalAssignmentsForHIT, input).Get(ctx, &output)
    return &output, err
}

func (a *MTurkStub) CreateAdditionalAssignmentsForHITAsync(ctx workflow.Context, input *mturk.CreateAdditionalAssignmentsForHITInput) *MturkCreateAdditionalAssignmentsForHITResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateAdditionalAssignmentsForHIT, input)
    return &MturkCreateAdditionalAssignmentsForHITResult{Result: future}
}

func (a *MTurkStub) CreateHIT(ctx workflow.Context, input *mturk.CreateHITInput) (*mturk.CreateHITOutput, error) {
    var output mturk.CreateHITOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateHIT, input).Get(ctx, &output)
    return &output, err
}

func (a *MTurkStub) CreateHITAsync(ctx workflow.Context, input *mturk.CreateHITInput) *MturkCreateHITResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateHIT, input)
    return &MturkCreateHITResult{Result: future}
}

func (a *MTurkStub) CreateHITType(ctx workflow.Context, input *mturk.CreateHITTypeInput) (*mturk.CreateHITTypeOutput, error) {
    var output mturk.CreateHITTypeOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateHITType, input).Get(ctx, &output)
    return &output, err
}

func (a *MTurkStub) CreateHITTypeAsync(ctx workflow.Context, input *mturk.CreateHITTypeInput) *MturkCreateHITTypeResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateHITType, input)
    return &MturkCreateHITTypeResult{Result: future}
}

func (a *MTurkStub) CreateHITWithHITType(ctx workflow.Context, input *mturk.CreateHITWithHITTypeInput) (*mturk.CreateHITWithHITTypeOutput, error) {
    var output mturk.CreateHITWithHITTypeOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateHITWithHITType, input).Get(ctx, &output)
    return &output, err
}

func (a *MTurkStub) CreateHITWithHITTypeAsync(ctx workflow.Context, input *mturk.CreateHITWithHITTypeInput) *MturkCreateHITWithHITTypeResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateHITWithHITType, input)
    return &MturkCreateHITWithHITTypeResult{Result: future}
}

func (a *MTurkStub) CreateQualificationType(ctx workflow.Context, input *mturk.CreateQualificationTypeInput) (*mturk.CreateQualificationTypeOutput, error) {
    var output mturk.CreateQualificationTypeOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateQualificationType, input).Get(ctx, &output)
    return &output, err
}

func (a *MTurkStub) CreateQualificationTypeAsync(ctx workflow.Context, input *mturk.CreateQualificationTypeInput) *MturkCreateQualificationTypeResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateQualificationType, input)
    return &MturkCreateQualificationTypeResult{Result: future}
}

func (a *MTurkStub) CreateWorkerBlock(ctx workflow.Context, input *mturk.CreateWorkerBlockInput) (*mturk.CreateWorkerBlockOutput, error) {
    var output mturk.CreateWorkerBlockOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateWorkerBlock, input).Get(ctx, &output)
    return &output, err
}

func (a *MTurkStub) CreateWorkerBlockAsync(ctx workflow.Context, input *mturk.CreateWorkerBlockInput) *MturkCreateWorkerBlockResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateWorkerBlock, input)
    return &MturkCreateWorkerBlockResult{Result: future}
}

func (a *MTurkStub) DeleteHIT(ctx workflow.Context, input *mturk.DeleteHITInput) (*mturk.DeleteHITOutput, error) {
    var output mturk.DeleteHITOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteHIT, input).Get(ctx, &output)
    return &output, err
}

func (a *MTurkStub) DeleteHITAsync(ctx workflow.Context, input *mturk.DeleteHITInput) *MturkDeleteHITResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteHIT, input)
    return &MturkDeleteHITResult{Result: future}
}

func (a *MTurkStub) DeleteQualificationType(ctx workflow.Context, input *mturk.DeleteQualificationTypeInput) (*mturk.DeleteQualificationTypeOutput, error) {
    var output mturk.DeleteQualificationTypeOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteQualificationType, input).Get(ctx, &output)
    return &output, err
}

func (a *MTurkStub) DeleteQualificationTypeAsync(ctx workflow.Context, input *mturk.DeleteQualificationTypeInput) *MturkDeleteQualificationTypeResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteQualificationType, input)
    return &MturkDeleteQualificationTypeResult{Result: future}
}

func (a *MTurkStub) DeleteWorkerBlock(ctx workflow.Context, input *mturk.DeleteWorkerBlockInput) (*mturk.DeleteWorkerBlockOutput, error) {
    var output mturk.DeleteWorkerBlockOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteWorkerBlock, input).Get(ctx, &output)
    return &output, err
}

func (a *MTurkStub) DeleteWorkerBlockAsync(ctx workflow.Context, input *mturk.DeleteWorkerBlockInput) *MturkDeleteWorkerBlockResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteWorkerBlock, input)
    return &MturkDeleteWorkerBlockResult{Result: future}
}

func (a *MTurkStub) DisassociateQualificationFromWorker(ctx workflow.Context, input *mturk.DisassociateQualificationFromWorkerInput) (*mturk.DisassociateQualificationFromWorkerOutput, error) {
    var output mturk.DisassociateQualificationFromWorkerOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DisassociateQualificationFromWorker, input).Get(ctx, &output)
    return &output, err
}

func (a *MTurkStub) DisassociateQualificationFromWorkerAsync(ctx workflow.Context, input *mturk.DisassociateQualificationFromWorkerInput) *MturkDisassociateQualificationFromWorkerResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DisassociateQualificationFromWorker, input)
    return &MturkDisassociateQualificationFromWorkerResult{Result: future}
}

func (a *MTurkStub) GetAccountBalance(ctx workflow.Context, input *mturk.GetAccountBalanceInput) (*mturk.GetAccountBalanceOutput, error) {
    var output mturk.GetAccountBalanceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetAccountBalance, input).Get(ctx, &output)
    return &output, err
}

func (a *MTurkStub) GetAccountBalanceAsync(ctx workflow.Context, input *mturk.GetAccountBalanceInput) *MturkGetAccountBalanceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetAccountBalance, input)
    return &MturkGetAccountBalanceResult{Result: future}
}

func (a *MTurkStub) GetAssignment(ctx workflow.Context, input *mturk.GetAssignmentInput) (*mturk.GetAssignmentOutput, error) {
    var output mturk.GetAssignmentOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetAssignment, input).Get(ctx, &output)
    return &output, err
}

func (a *MTurkStub) GetAssignmentAsync(ctx workflow.Context, input *mturk.GetAssignmentInput) *MturkGetAssignmentResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetAssignment, input)
    return &MturkGetAssignmentResult{Result: future}
}

func (a *MTurkStub) GetFileUploadURL(ctx workflow.Context, input *mturk.GetFileUploadURLInput) (*mturk.GetFileUploadURLOutput, error) {
    var output mturk.GetFileUploadURLOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetFileUploadURL, input).Get(ctx, &output)
    return &output, err
}

func (a *MTurkStub) GetFileUploadURLAsync(ctx workflow.Context, input *mturk.GetFileUploadURLInput) *MturkGetFileUploadURLResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetFileUploadURL, input)
    return &MturkGetFileUploadURLResult{Result: future}
}

func (a *MTurkStub) GetHIT(ctx workflow.Context, input *mturk.GetHITInput) (*mturk.GetHITOutput, error) {
    var output mturk.GetHITOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetHIT, input).Get(ctx, &output)
    return &output, err
}

func (a *MTurkStub) GetHITAsync(ctx workflow.Context, input *mturk.GetHITInput) *MturkGetHITResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetHIT, input)
    return &MturkGetHITResult{Result: future}
}

func (a *MTurkStub) GetQualificationScore(ctx workflow.Context, input *mturk.GetQualificationScoreInput) (*mturk.GetQualificationScoreOutput, error) {
    var output mturk.GetQualificationScoreOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetQualificationScore, input).Get(ctx, &output)
    return &output, err
}

func (a *MTurkStub) GetQualificationScoreAsync(ctx workflow.Context, input *mturk.GetQualificationScoreInput) *MturkGetQualificationScoreResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetQualificationScore, input)
    return &MturkGetQualificationScoreResult{Result: future}
}

func (a *MTurkStub) GetQualificationType(ctx workflow.Context, input *mturk.GetQualificationTypeInput) (*mturk.GetQualificationTypeOutput, error) {
    var output mturk.GetQualificationTypeOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetQualificationType, input).Get(ctx, &output)
    return &output, err
}

func (a *MTurkStub) GetQualificationTypeAsync(ctx workflow.Context, input *mturk.GetQualificationTypeInput) *MturkGetQualificationTypeResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetQualificationType, input)
    return &MturkGetQualificationTypeResult{Result: future}
}

func (a *MTurkStub) ListAssignmentsForHIT(ctx workflow.Context, input *mturk.ListAssignmentsForHITInput) (*mturk.ListAssignmentsForHITOutput, error) {
    var output mturk.ListAssignmentsForHITOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListAssignmentsForHIT, input).Get(ctx, &output)
    return &output, err
}

func (a *MTurkStub) ListAssignmentsForHITAsync(ctx workflow.Context, input *mturk.ListAssignmentsForHITInput) *MturkListAssignmentsForHITResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListAssignmentsForHIT, input)
    return &MturkListAssignmentsForHITResult{Result: future}
}

func (a *MTurkStub) ListBonusPayments(ctx workflow.Context, input *mturk.ListBonusPaymentsInput) (*mturk.ListBonusPaymentsOutput, error) {
    var output mturk.ListBonusPaymentsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListBonusPayments, input).Get(ctx, &output)
    return &output, err
}

func (a *MTurkStub) ListBonusPaymentsAsync(ctx workflow.Context, input *mturk.ListBonusPaymentsInput) *MturkListBonusPaymentsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListBonusPayments, input)
    return &MturkListBonusPaymentsResult{Result: future}
}

func (a *MTurkStub) ListHITs(ctx workflow.Context, input *mturk.ListHITsInput) (*mturk.ListHITsOutput, error) {
    var output mturk.ListHITsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListHITs, input).Get(ctx, &output)
    return &output, err
}

func (a *MTurkStub) ListHITsAsync(ctx workflow.Context, input *mturk.ListHITsInput) *MturkListHITsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListHITs, input)
    return &MturkListHITsResult{Result: future}
}

func (a *MTurkStub) ListHITsForQualificationType(ctx workflow.Context, input *mturk.ListHITsForQualificationTypeInput) (*mturk.ListHITsForQualificationTypeOutput, error) {
    var output mturk.ListHITsForQualificationTypeOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListHITsForQualificationType, input).Get(ctx, &output)
    return &output, err
}

func (a *MTurkStub) ListHITsForQualificationTypeAsync(ctx workflow.Context, input *mturk.ListHITsForQualificationTypeInput) *MturkListHITsForQualificationTypeResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListHITsForQualificationType, input)
    return &MturkListHITsForQualificationTypeResult{Result: future}
}

func (a *MTurkStub) ListQualificationRequests(ctx workflow.Context, input *mturk.ListQualificationRequestsInput) (*mturk.ListQualificationRequestsOutput, error) {
    var output mturk.ListQualificationRequestsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListQualificationRequests, input).Get(ctx, &output)
    return &output, err
}

func (a *MTurkStub) ListQualificationRequestsAsync(ctx workflow.Context, input *mturk.ListQualificationRequestsInput) *MturkListQualificationRequestsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListQualificationRequests, input)
    return &MturkListQualificationRequestsResult{Result: future}
}

func (a *MTurkStub) ListQualificationTypes(ctx workflow.Context, input *mturk.ListQualificationTypesInput) (*mturk.ListQualificationTypesOutput, error) {
    var output mturk.ListQualificationTypesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListQualificationTypes, input).Get(ctx, &output)
    return &output, err
}

func (a *MTurkStub) ListQualificationTypesAsync(ctx workflow.Context, input *mturk.ListQualificationTypesInput) *MturkListQualificationTypesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListQualificationTypes, input)
    return &MturkListQualificationTypesResult{Result: future}
}

func (a *MTurkStub) ListReviewPolicyResultsForHIT(ctx workflow.Context, input *mturk.ListReviewPolicyResultsForHITInput) (*mturk.ListReviewPolicyResultsForHITOutput, error) {
    var output mturk.ListReviewPolicyResultsForHITOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListReviewPolicyResultsForHIT, input).Get(ctx, &output)
    return &output, err
}

func (a *MTurkStub) ListReviewPolicyResultsForHITAsync(ctx workflow.Context, input *mturk.ListReviewPolicyResultsForHITInput) *MturkListReviewPolicyResultsForHITResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListReviewPolicyResultsForHIT, input)
    return &MturkListReviewPolicyResultsForHITResult{Result: future}
}

func (a *MTurkStub) ListReviewableHITs(ctx workflow.Context, input *mturk.ListReviewableHITsInput) (*mturk.ListReviewableHITsOutput, error) {
    var output mturk.ListReviewableHITsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListReviewableHITs, input).Get(ctx, &output)
    return &output, err
}

func (a *MTurkStub) ListReviewableHITsAsync(ctx workflow.Context, input *mturk.ListReviewableHITsInput) *MturkListReviewableHITsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListReviewableHITs, input)
    return &MturkListReviewableHITsResult{Result: future}
}

func (a *MTurkStub) ListWorkerBlocks(ctx workflow.Context, input *mturk.ListWorkerBlocksInput) (*mturk.ListWorkerBlocksOutput, error) {
    var output mturk.ListWorkerBlocksOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListWorkerBlocks, input).Get(ctx, &output)
    return &output, err
}

func (a *MTurkStub) ListWorkerBlocksAsync(ctx workflow.Context, input *mturk.ListWorkerBlocksInput) *MturkListWorkerBlocksResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListWorkerBlocks, input)
    return &MturkListWorkerBlocksResult{Result: future}
}

func (a *MTurkStub) ListWorkersWithQualificationType(ctx workflow.Context, input *mturk.ListWorkersWithQualificationTypeInput) (*mturk.ListWorkersWithQualificationTypeOutput, error) {
    var output mturk.ListWorkersWithQualificationTypeOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListWorkersWithQualificationType, input).Get(ctx, &output)
    return &output, err
}

func (a *MTurkStub) ListWorkersWithQualificationTypeAsync(ctx workflow.Context, input *mturk.ListWorkersWithQualificationTypeInput) *MturkListWorkersWithQualificationTypeResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListWorkersWithQualificationType, input)
    return &MturkListWorkersWithQualificationTypeResult{Result: future}
}

func (a *MTurkStub) NotifyWorkers(ctx workflow.Context, input *mturk.NotifyWorkersInput) (*mturk.NotifyWorkersOutput, error) {
    var output mturk.NotifyWorkersOutput
    err := workflow.ExecuteActivity(ctx, a.activities.NotifyWorkers, input).Get(ctx, &output)
    return &output, err
}

func (a *MTurkStub) NotifyWorkersAsync(ctx workflow.Context, input *mturk.NotifyWorkersInput) *MturkNotifyWorkersResult {
    future := workflow.ExecuteActivity(ctx, a.activities.NotifyWorkers, input)
    return &MturkNotifyWorkersResult{Result: future}
}

func (a *MTurkStub) RejectAssignment(ctx workflow.Context, input *mturk.RejectAssignmentInput) (*mturk.RejectAssignmentOutput, error) {
    var output mturk.RejectAssignmentOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RejectAssignment, input).Get(ctx, &output)
    return &output, err
}

func (a *MTurkStub) RejectAssignmentAsync(ctx workflow.Context, input *mturk.RejectAssignmentInput) *MturkRejectAssignmentResult {
    future := workflow.ExecuteActivity(ctx, a.activities.RejectAssignment, input)
    return &MturkRejectAssignmentResult{Result: future}
}

func (a *MTurkStub) SendBonus(ctx workflow.Context, input *mturk.SendBonusInput) (*mturk.SendBonusOutput, error) {
    var output mturk.SendBonusOutput
    err := workflow.ExecuteActivity(ctx, a.activities.SendBonus, input).Get(ctx, &output)
    return &output, err
}

func (a *MTurkStub) SendBonusAsync(ctx workflow.Context, input *mturk.SendBonusInput) *MturkSendBonusResult {
    future := workflow.ExecuteActivity(ctx, a.activities.SendBonus, input)
    return &MturkSendBonusResult{Result: future}
}

func (a *MTurkStub) SendTestEventNotification(ctx workflow.Context, input *mturk.SendTestEventNotificationInput) (*mturk.SendTestEventNotificationOutput, error) {
    var output mturk.SendTestEventNotificationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.SendTestEventNotification, input).Get(ctx, &output)
    return &output, err
}

func (a *MTurkStub) SendTestEventNotificationAsync(ctx workflow.Context, input *mturk.SendTestEventNotificationInput) *MturkSendTestEventNotificationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.SendTestEventNotification, input)
    return &MturkSendTestEventNotificationResult{Result: future}
}

func (a *MTurkStub) UpdateExpirationForHIT(ctx workflow.Context, input *mturk.UpdateExpirationForHITInput) (*mturk.UpdateExpirationForHITOutput, error) {
    var output mturk.UpdateExpirationForHITOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateExpirationForHIT, input).Get(ctx, &output)
    return &output, err
}

func (a *MTurkStub) UpdateExpirationForHITAsync(ctx workflow.Context, input *mturk.UpdateExpirationForHITInput) *MturkUpdateExpirationForHITResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateExpirationForHIT, input)
    return &MturkUpdateExpirationForHITResult{Result: future}
}

func (a *MTurkStub) UpdateHITReviewStatus(ctx workflow.Context, input *mturk.UpdateHITReviewStatusInput) (*mturk.UpdateHITReviewStatusOutput, error) {
    var output mturk.UpdateHITReviewStatusOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateHITReviewStatus, input).Get(ctx, &output)
    return &output, err
}

func (a *MTurkStub) UpdateHITReviewStatusAsync(ctx workflow.Context, input *mturk.UpdateHITReviewStatusInput) *MturkUpdateHITReviewStatusResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateHITReviewStatus, input)
    return &MturkUpdateHITReviewStatusResult{Result: future}
}

func (a *MTurkStub) UpdateHITTypeOfHIT(ctx workflow.Context, input *mturk.UpdateHITTypeOfHITInput) (*mturk.UpdateHITTypeOfHITOutput, error) {
    var output mturk.UpdateHITTypeOfHITOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateHITTypeOfHIT, input).Get(ctx, &output)
    return &output, err
}

func (a *MTurkStub) UpdateHITTypeOfHITAsync(ctx workflow.Context, input *mturk.UpdateHITTypeOfHITInput) *MturkUpdateHITTypeOfHITResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateHITTypeOfHIT, input)
    return &MturkUpdateHITTypeOfHITResult{Result: future}
}

func (a *MTurkStub) UpdateNotificationSettings(ctx workflow.Context, input *mturk.UpdateNotificationSettingsInput) (*mturk.UpdateNotificationSettingsOutput, error) {
    var output mturk.UpdateNotificationSettingsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateNotificationSettings, input).Get(ctx, &output)
    return &output, err
}

func (a *MTurkStub) UpdateNotificationSettingsAsync(ctx workflow.Context, input *mturk.UpdateNotificationSettingsInput) *MturkUpdateNotificationSettingsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateNotificationSettings, input)
    return &MturkUpdateNotificationSettingsResult{Result: future}
}

func (a *MTurkStub) UpdateQualificationType(ctx workflow.Context, input *mturk.UpdateQualificationTypeInput) (*mturk.UpdateQualificationTypeOutput, error) {
    var output mturk.UpdateQualificationTypeOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateQualificationType, input).Get(ctx, &output)
    return &output, err
}

func (a *MTurkStub) UpdateQualificationTypeAsync(ctx workflow.Context, input *mturk.UpdateQualificationTypeInput) *MturkUpdateQualificationTypeResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateQualificationType, input)
    return &MturkUpdateQualificationTypeResult{Result: future}
}