package awsclients

import (
	"github.com/aws/aws-sdk-go/service/frauddetector"
	"go.temporal.io/sdk/workflow"
	"temporal.io/aws-sdk/awsactivities"
)

type FraudDetectorClient interface {
	BatchCreateVariable(ctx workflow.Context, input *frauddetector.BatchCreateVariableInput) (*frauddetector.BatchCreateVariableOutput, error)
	BatchCreateVariableAsync(ctx workflow.Context, input *frauddetector.BatchCreateVariableInput) *FrauddetectorBatchCreateVariableResult

	BatchGetVariable(ctx workflow.Context, input *frauddetector.BatchGetVariableInput) (*frauddetector.BatchGetVariableOutput, error)
	BatchGetVariableAsync(ctx workflow.Context, input *frauddetector.BatchGetVariableInput) *FrauddetectorBatchGetVariableResult

	CreateDetectorVersion(ctx workflow.Context, input *frauddetector.CreateDetectorVersionInput) (*frauddetector.CreateDetectorVersionOutput, error)
	CreateDetectorVersionAsync(ctx workflow.Context, input *frauddetector.CreateDetectorVersionInput) *FrauddetectorCreateDetectorVersionResult

	CreateModel(ctx workflow.Context, input *frauddetector.CreateModelInput) (*frauddetector.CreateModelOutput, error)
	CreateModelAsync(ctx workflow.Context, input *frauddetector.CreateModelInput) *FrauddetectorCreateModelResult

	CreateModelVersion(ctx workflow.Context, input *frauddetector.CreateModelVersionInput) (*frauddetector.CreateModelVersionOutput, error)
	CreateModelVersionAsync(ctx workflow.Context, input *frauddetector.CreateModelVersionInput) *FrauddetectorCreateModelVersionResult

	CreateRule(ctx workflow.Context, input *frauddetector.CreateRuleInput) (*frauddetector.CreateRuleOutput, error)
	CreateRuleAsync(ctx workflow.Context, input *frauddetector.CreateRuleInput) *FrauddetectorCreateRuleResult

	CreateVariable(ctx workflow.Context, input *frauddetector.CreateVariableInput) (*frauddetector.CreateVariableOutput, error)
	CreateVariableAsync(ctx workflow.Context, input *frauddetector.CreateVariableInput) *FrauddetectorCreateVariableResult

	DeleteDetector(ctx workflow.Context, input *frauddetector.DeleteDetectorInput) (*frauddetector.DeleteDetectorOutput, error)
	DeleteDetectorAsync(ctx workflow.Context, input *frauddetector.DeleteDetectorInput) *FrauddetectorDeleteDetectorResult

	DeleteDetectorVersion(ctx workflow.Context, input *frauddetector.DeleteDetectorVersionInput) (*frauddetector.DeleteDetectorVersionOutput, error)
	DeleteDetectorVersionAsync(ctx workflow.Context, input *frauddetector.DeleteDetectorVersionInput) *FrauddetectorDeleteDetectorVersionResult

	DeleteEvent(ctx workflow.Context, input *frauddetector.DeleteEventInput) (*frauddetector.DeleteEventOutput, error)
	DeleteEventAsync(ctx workflow.Context, input *frauddetector.DeleteEventInput) *FrauddetectorDeleteEventResult

	DeleteRule(ctx workflow.Context, input *frauddetector.DeleteRuleInput) (*frauddetector.DeleteRuleOutput, error)
	DeleteRuleAsync(ctx workflow.Context, input *frauddetector.DeleteRuleInput) *FrauddetectorDeleteRuleResult

	DescribeDetector(ctx workflow.Context, input *frauddetector.DescribeDetectorInput) (*frauddetector.DescribeDetectorOutput, error)
	DescribeDetectorAsync(ctx workflow.Context, input *frauddetector.DescribeDetectorInput) *FrauddetectorDescribeDetectorResult

	DescribeModelVersions(ctx workflow.Context, input *frauddetector.DescribeModelVersionsInput) (*frauddetector.DescribeModelVersionsOutput, error)
	DescribeModelVersionsAsync(ctx workflow.Context, input *frauddetector.DescribeModelVersionsInput) *FrauddetectorDescribeModelVersionsResult

	GetDetectorVersion(ctx workflow.Context, input *frauddetector.GetDetectorVersionInput) (*frauddetector.GetDetectorVersionOutput, error)
	GetDetectorVersionAsync(ctx workflow.Context, input *frauddetector.GetDetectorVersionInput) *FrauddetectorGetDetectorVersionResult

	GetDetectors(ctx workflow.Context, input *frauddetector.GetDetectorsInput) (*frauddetector.GetDetectorsOutput, error)
	GetDetectorsAsync(ctx workflow.Context, input *frauddetector.GetDetectorsInput) *FrauddetectorGetDetectorsResult

	GetEntityTypes(ctx workflow.Context, input *frauddetector.GetEntityTypesInput) (*frauddetector.GetEntityTypesOutput, error)
	GetEntityTypesAsync(ctx workflow.Context, input *frauddetector.GetEntityTypesInput) *FrauddetectorGetEntityTypesResult

	GetEventPrediction(ctx workflow.Context, input *frauddetector.GetEventPredictionInput) (*frauddetector.GetEventPredictionOutput, error)
	GetEventPredictionAsync(ctx workflow.Context, input *frauddetector.GetEventPredictionInput) *FrauddetectorGetEventPredictionResult

	GetEventTypes(ctx workflow.Context, input *frauddetector.GetEventTypesInput) (*frauddetector.GetEventTypesOutput, error)
	GetEventTypesAsync(ctx workflow.Context, input *frauddetector.GetEventTypesInput) *FrauddetectorGetEventTypesResult

	GetExternalModels(ctx workflow.Context, input *frauddetector.GetExternalModelsInput) (*frauddetector.GetExternalModelsOutput, error)
	GetExternalModelsAsync(ctx workflow.Context, input *frauddetector.GetExternalModelsInput) *FrauddetectorGetExternalModelsResult

	GetKMSEncryptionKey(ctx workflow.Context, input *frauddetector.GetKMSEncryptionKeyInput) (*frauddetector.GetKMSEncryptionKeyOutput, error)
	GetKMSEncryptionKeyAsync(ctx workflow.Context, input *frauddetector.GetKMSEncryptionKeyInput) *FrauddetectorGetKMSEncryptionKeyResult

	GetLabels(ctx workflow.Context, input *frauddetector.GetLabelsInput) (*frauddetector.GetLabelsOutput, error)
	GetLabelsAsync(ctx workflow.Context, input *frauddetector.GetLabelsInput) *FrauddetectorGetLabelsResult

	GetModelVersion(ctx workflow.Context, input *frauddetector.GetModelVersionInput) (*frauddetector.GetModelVersionOutput, error)
	GetModelVersionAsync(ctx workflow.Context, input *frauddetector.GetModelVersionInput) *FrauddetectorGetModelVersionResult

	GetModels(ctx workflow.Context, input *frauddetector.GetModelsInput) (*frauddetector.GetModelsOutput, error)
	GetModelsAsync(ctx workflow.Context, input *frauddetector.GetModelsInput) *FrauddetectorGetModelsResult

	GetOutcomes(ctx workflow.Context, input *frauddetector.GetOutcomesInput) (*frauddetector.GetOutcomesOutput, error)
	GetOutcomesAsync(ctx workflow.Context, input *frauddetector.GetOutcomesInput) *FrauddetectorGetOutcomesResult

	GetRules(ctx workflow.Context, input *frauddetector.GetRulesInput) (*frauddetector.GetRulesOutput, error)
	GetRulesAsync(ctx workflow.Context, input *frauddetector.GetRulesInput) *FrauddetectorGetRulesResult

	GetVariables(ctx workflow.Context, input *frauddetector.GetVariablesInput) (*frauddetector.GetVariablesOutput, error)
	GetVariablesAsync(ctx workflow.Context, input *frauddetector.GetVariablesInput) *FrauddetectorGetVariablesResult

	ListTagsForResource(ctx workflow.Context, input *frauddetector.ListTagsForResourceInput) (*frauddetector.ListTagsForResourceOutput, error)
	ListTagsForResourceAsync(ctx workflow.Context, input *frauddetector.ListTagsForResourceInput) *FrauddetectorListTagsForResourceResult

	PutDetector(ctx workflow.Context, input *frauddetector.PutDetectorInput) (*frauddetector.PutDetectorOutput, error)
	PutDetectorAsync(ctx workflow.Context, input *frauddetector.PutDetectorInput) *FrauddetectorPutDetectorResult

	PutEntityType(ctx workflow.Context, input *frauddetector.PutEntityTypeInput) (*frauddetector.PutEntityTypeOutput, error)
	PutEntityTypeAsync(ctx workflow.Context, input *frauddetector.PutEntityTypeInput) *FrauddetectorPutEntityTypeResult

	PutEventType(ctx workflow.Context, input *frauddetector.PutEventTypeInput) (*frauddetector.PutEventTypeOutput, error)
	PutEventTypeAsync(ctx workflow.Context, input *frauddetector.PutEventTypeInput) *FrauddetectorPutEventTypeResult

	PutExternalModel(ctx workflow.Context, input *frauddetector.PutExternalModelInput) (*frauddetector.PutExternalModelOutput, error)
	PutExternalModelAsync(ctx workflow.Context, input *frauddetector.PutExternalModelInput) *FrauddetectorPutExternalModelResult

	PutKMSEncryptionKey(ctx workflow.Context, input *frauddetector.PutKMSEncryptionKeyInput) (*frauddetector.PutKMSEncryptionKeyOutput, error)
	PutKMSEncryptionKeyAsync(ctx workflow.Context, input *frauddetector.PutKMSEncryptionKeyInput) *FrauddetectorPutKMSEncryptionKeyResult

	PutLabel(ctx workflow.Context, input *frauddetector.PutLabelInput) (*frauddetector.PutLabelOutput, error)
	PutLabelAsync(ctx workflow.Context, input *frauddetector.PutLabelInput) *FrauddetectorPutLabelResult

	PutOutcome(ctx workflow.Context, input *frauddetector.PutOutcomeInput) (*frauddetector.PutOutcomeOutput, error)
	PutOutcomeAsync(ctx workflow.Context, input *frauddetector.PutOutcomeInput) *FrauddetectorPutOutcomeResult

	TagResource(ctx workflow.Context, input *frauddetector.TagResourceInput) (*frauddetector.TagResourceOutput, error)
	TagResourceAsync(ctx workflow.Context, input *frauddetector.TagResourceInput) *FrauddetectorTagResourceResult

	UntagResource(ctx workflow.Context, input *frauddetector.UntagResourceInput) (*frauddetector.UntagResourceOutput, error)
	UntagResourceAsync(ctx workflow.Context, input *frauddetector.UntagResourceInput) *FrauddetectorUntagResourceResult

	UpdateDetectorVersion(ctx workflow.Context, input *frauddetector.UpdateDetectorVersionInput) (*frauddetector.UpdateDetectorVersionOutput, error)
	UpdateDetectorVersionAsync(ctx workflow.Context, input *frauddetector.UpdateDetectorVersionInput) *FrauddetectorUpdateDetectorVersionResult

	UpdateDetectorVersionMetadata(ctx workflow.Context, input *frauddetector.UpdateDetectorVersionMetadataInput) (*frauddetector.UpdateDetectorVersionMetadataOutput, error)
	UpdateDetectorVersionMetadataAsync(ctx workflow.Context, input *frauddetector.UpdateDetectorVersionMetadataInput) *FrauddetectorUpdateDetectorVersionMetadataResult

	UpdateDetectorVersionStatus(ctx workflow.Context, input *frauddetector.UpdateDetectorVersionStatusInput) (*frauddetector.UpdateDetectorVersionStatusOutput, error)
	UpdateDetectorVersionStatusAsync(ctx workflow.Context, input *frauddetector.UpdateDetectorVersionStatusInput) *FrauddetectorUpdateDetectorVersionStatusResult

	UpdateModel(ctx workflow.Context, input *frauddetector.UpdateModelInput) (*frauddetector.UpdateModelOutput, error)
	UpdateModelAsync(ctx workflow.Context, input *frauddetector.UpdateModelInput) *FrauddetectorUpdateModelResult

	UpdateModelVersion(ctx workflow.Context, input *frauddetector.UpdateModelVersionInput) (*frauddetector.UpdateModelVersionOutput, error)
	UpdateModelVersionAsync(ctx workflow.Context, input *frauddetector.UpdateModelVersionInput) *FrauddetectorUpdateModelVersionResult

	UpdateModelVersionStatus(ctx workflow.Context, input *frauddetector.UpdateModelVersionStatusInput) (*frauddetector.UpdateModelVersionStatusOutput, error)
	UpdateModelVersionStatusAsync(ctx workflow.Context, input *frauddetector.UpdateModelVersionStatusInput) *FrauddetectorUpdateModelVersionStatusResult

	UpdateRuleMetadata(ctx workflow.Context, input *frauddetector.UpdateRuleMetadataInput) (*frauddetector.UpdateRuleMetadataOutput, error)
	UpdateRuleMetadataAsync(ctx workflow.Context, input *frauddetector.UpdateRuleMetadataInput) *FrauddetectorUpdateRuleMetadataResult

	UpdateRuleVersion(ctx workflow.Context, input *frauddetector.UpdateRuleVersionInput) (*frauddetector.UpdateRuleVersionOutput, error)
	UpdateRuleVersionAsync(ctx workflow.Context, input *frauddetector.UpdateRuleVersionInput) *FrauddetectorUpdateRuleVersionResult

	UpdateVariable(ctx workflow.Context, input *frauddetector.UpdateVariableInput) (*frauddetector.UpdateVariableOutput, error)
	UpdateVariableAsync(ctx workflow.Context, input *frauddetector.UpdateVariableInput) *FrauddetectorUpdateVariableResult
}

type FrauddetectorBatchCreateVariableResult struct {
	Result workflow.Future
}

func (r *FrauddetectorBatchCreateVariableResult) Get(ctx workflow.Context) (*frauddetector.BatchCreateVariableOutput, error) {
	var output frauddetector.BatchCreateVariableOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type FrauddetectorBatchGetVariableResult struct {
	Result workflow.Future
}

func (r *FrauddetectorBatchGetVariableResult) Get(ctx workflow.Context) (*frauddetector.BatchGetVariableOutput, error) {
	var output frauddetector.BatchGetVariableOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type FrauddetectorCreateDetectorVersionResult struct {
	Result workflow.Future
}

func (r *FrauddetectorCreateDetectorVersionResult) Get(ctx workflow.Context) (*frauddetector.CreateDetectorVersionOutput, error) {
	var output frauddetector.CreateDetectorVersionOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type FrauddetectorCreateModelResult struct {
	Result workflow.Future
}

func (r *FrauddetectorCreateModelResult) Get(ctx workflow.Context) (*frauddetector.CreateModelOutput, error) {
	var output frauddetector.CreateModelOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type FrauddetectorCreateModelVersionResult struct {
	Result workflow.Future
}

func (r *FrauddetectorCreateModelVersionResult) Get(ctx workflow.Context) (*frauddetector.CreateModelVersionOutput, error) {
	var output frauddetector.CreateModelVersionOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type FrauddetectorCreateRuleResult struct {
	Result workflow.Future
}

func (r *FrauddetectorCreateRuleResult) Get(ctx workflow.Context) (*frauddetector.CreateRuleOutput, error) {
	var output frauddetector.CreateRuleOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type FrauddetectorCreateVariableResult struct {
	Result workflow.Future
}

func (r *FrauddetectorCreateVariableResult) Get(ctx workflow.Context) (*frauddetector.CreateVariableOutput, error) {
	var output frauddetector.CreateVariableOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type FrauddetectorDeleteDetectorResult struct {
	Result workflow.Future
}

func (r *FrauddetectorDeleteDetectorResult) Get(ctx workflow.Context) (*frauddetector.DeleteDetectorOutput, error) {
	var output frauddetector.DeleteDetectorOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type FrauddetectorDeleteDetectorVersionResult struct {
	Result workflow.Future
}

func (r *FrauddetectorDeleteDetectorVersionResult) Get(ctx workflow.Context) (*frauddetector.DeleteDetectorVersionOutput, error) {
	var output frauddetector.DeleteDetectorVersionOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type FrauddetectorDeleteEventResult struct {
	Result workflow.Future
}

func (r *FrauddetectorDeleteEventResult) Get(ctx workflow.Context) (*frauddetector.DeleteEventOutput, error) {
	var output frauddetector.DeleteEventOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type FrauddetectorDeleteRuleResult struct {
	Result workflow.Future
}

func (r *FrauddetectorDeleteRuleResult) Get(ctx workflow.Context) (*frauddetector.DeleteRuleOutput, error) {
	var output frauddetector.DeleteRuleOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type FrauddetectorDescribeDetectorResult struct {
	Result workflow.Future
}

func (r *FrauddetectorDescribeDetectorResult) Get(ctx workflow.Context) (*frauddetector.DescribeDetectorOutput, error) {
	var output frauddetector.DescribeDetectorOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type FrauddetectorDescribeModelVersionsResult struct {
	Result workflow.Future
}

func (r *FrauddetectorDescribeModelVersionsResult) Get(ctx workflow.Context) (*frauddetector.DescribeModelVersionsOutput, error) {
	var output frauddetector.DescribeModelVersionsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type FrauddetectorGetDetectorVersionResult struct {
	Result workflow.Future
}

func (r *FrauddetectorGetDetectorVersionResult) Get(ctx workflow.Context) (*frauddetector.GetDetectorVersionOutput, error) {
	var output frauddetector.GetDetectorVersionOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type FrauddetectorGetDetectorsResult struct {
	Result workflow.Future
}

func (r *FrauddetectorGetDetectorsResult) Get(ctx workflow.Context) (*frauddetector.GetDetectorsOutput, error) {
	var output frauddetector.GetDetectorsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type FrauddetectorGetEntityTypesResult struct {
	Result workflow.Future
}

func (r *FrauddetectorGetEntityTypesResult) Get(ctx workflow.Context) (*frauddetector.GetEntityTypesOutput, error) {
	var output frauddetector.GetEntityTypesOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type FrauddetectorGetEventPredictionResult struct {
	Result workflow.Future
}

func (r *FrauddetectorGetEventPredictionResult) Get(ctx workflow.Context) (*frauddetector.GetEventPredictionOutput, error) {
	var output frauddetector.GetEventPredictionOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type FrauddetectorGetEventTypesResult struct {
	Result workflow.Future
}

func (r *FrauddetectorGetEventTypesResult) Get(ctx workflow.Context) (*frauddetector.GetEventTypesOutput, error) {
	var output frauddetector.GetEventTypesOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type FrauddetectorGetExternalModelsResult struct {
	Result workflow.Future
}

func (r *FrauddetectorGetExternalModelsResult) Get(ctx workflow.Context) (*frauddetector.GetExternalModelsOutput, error) {
	var output frauddetector.GetExternalModelsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type FrauddetectorGetKMSEncryptionKeyResult struct {
	Result workflow.Future
}

func (r *FrauddetectorGetKMSEncryptionKeyResult) Get(ctx workflow.Context) (*frauddetector.GetKMSEncryptionKeyOutput, error) {
	var output frauddetector.GetKMSEncryptionKeyOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type FrauddetectorGetLabelsResult struct {
	Result workflow.Future
}

func (r *FrauddetectorGetLabelsResult) Get(ctx workflow.Context) (*frauddetector.GetLabelsOutput, error) {
	var output frauddetector.GetLabelsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type FrauddetectorGetModelVersionResult struct {
	Result workflow.Future
}

func (r *FrauddetectorGetModelVersionResult) Get(ctx workflow.Context) (*frauddetector.GetModelVersionOutput, error) {
	var output frauddetector.GetModelVersionOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type FrauddetectorGetModelsResult struct {
	Result workflow.Future
}

func (r *FrauddetectorGetModelsResult) Get(ctx workflow.Context) (*frauddetector.GetModelsOutput, error) {
	var output frauddetector.GetModelsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type FrauddetectorGetOutcomesResult struct {
	Result workflow.Future
}

func (r *FrauddetectorGetOutcomesResult) Get(ctx workflow.Context) (*frauddetector.GetOutcomesOutput, error) {
	var output frauddetector.GetOutcomesOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type FrauddetectorGetRulesResult struct {
	Result workflow.Future
}

func (r *FrauddetectorGetRulesResult) Get(ctx workflow.Context) (*frauddetector.GetRulesOutput, error) {
	var output frauddetector.GetRulesOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type FrauddetectorGetVariablesResult struct {
	Result workflow.Future
}

func (r *FrauddetectorGetVariablesResult) Get(ctx workflow.Context) (*frauddetector.GetVariablesOutput, error) {
	var output frauddetector.GetVariablesOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type FrauddetectorListTagsForResourceResult struct {
	Result workflow.Future
}

func (r *FrauddetectorListTagsForResourceResult) Get(ctx workflow.Context) (*frauddetector.ListTagsForResourceOutput, error) {
	var output frauddetector.ListTagsForResourceOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type FrauddetectorPutDetectorResult struct {
	Result workflow.Future
}

func (r *FrauddetectorPutDetectorResult) Get(ctx workflow.Context) (*frauddetector.PutDetectorOutput, error) {
	var output frauddetector.PutDetectorOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type FrauddetectorPutEntityTypeResult struct {
	Result workflow.Future
}

func (r *FrauddetectorPutEntityTypeResult) Get(ctx workflow.Context) (*frauddetector.PutEntityTypeOutput, error) {
	var output frauddetector.PutEntityTypeOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type FrauddetectorPutEventTypeResult struct {
	Result workflow.Future
}

func (r *FrauddetectorPutEventTypeResult) Get(ctx workflow.Context) (*frauddetector.PutEventTypeOutput, error) {
	var output frauddetector.PutEventTypeOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type FrauddetectorPutExternalModelResult struct {
	Result workflow.Future
}

func (r *FrauddetectorPutExternalModelResult) Get(ctx workflow.Context) (*frauddetector.PutExternalModelOutput, error) {
	var output frauddetector.PutExternalModelOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type FrauddetectorPutKMSEncryptionKeyResult struct {
	Result workflow.Future
}

func (r *FrauddetectorPutKMSEncryptionKeyResult) Get(ctx workflow.Context) (*frauddetector.PutKMSEncryptionKeyOutput, error) {
	var output frauddetector.PutKMSEncryptionKeyOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type FrauddetectorPutLabelResult struct {
	Result workflow.Future
}

func (r *FrauddetectorPutLabelResult) Get(ctx workflow.Context) (*frauddetector.PutLabelOutput, error) {
	var output frauddetector.PutLabelOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type FrauddetectorPutOutcomeResult struct {
	Result workflow.Future
}

func (r *FrauddetectorPutOutcomeResult) Get(ctx workflow.Context) (*frauddetector.PutOutcomeOutput, error) {
	var output frauddetector.PutOutcomeOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type FrauddetectorTagResourceResult struct {
	Result workflow.Future
}

func (r *FrauddetectorTagResourceResult) Get(ctx workflow.Context) (*frauddetector.TagResourceOutput, error) {
	var output frauddetector.TagResourceOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type FrauddetectorUntagResourceResult struct {
	Result workflow.Future
}

func (r *FrauddetectorUntagResourceResult) Get(ctx workflow.Context) (*frauddetector.UntagResourceOutput, error) {
	var output frauddetector.UntagResourceOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type FrauddetectorUpdateDetectorVersionResult struct {
	Result workflow.Future
}

func (r *FrauddetectorUpdateDetectorVersionResult) Get(ctx workflow.Context) (*frauddetector.UpdateDetectorVersionOutput, error) {
	var output frauddetector.UpdateDetectorVersionOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type FrauddetectorUpdateDetectorVersionMetadataResult struct {
	Result workflow.Future
}

func (r *FrauddetectorUpdateDetectorVersionMetadataResult) Get(ctx workflow.Context) (*frauddetector.UpdateDetectorVersionMetadataOutput, error) {
	var output frauddetector.UpdateDetectorVersionMetadataOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type FrauddetectorUpdateDetectorVersionStatusResult struct {
	Result workflow.Future
}

func (r *FrauddetectorUpdateDetectorVersionStatusResult) Get(ctx workflow.Context) (*frauddetector.UpdateDetectorVersionStatusOutput, error) {
	var output frauddetector.UpdateDetectorVersionStatusOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type FrauddetectorUpdateModelResult struct {
	Result workflow.Future
}

func (r *FrauddetectorUpdateModelResult) Get(ctx workflow.Context) (*frauddetector.UpdateModelOutput, error) {
	var output frauddetector.UpdateModelOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type FrauddetectorUpdateModelVersionResult struct {
	Result workflow.Future
}

func (r *FrauddetectorUpdateModelVersionResult) Get(ctx workflow.Context) (*frauddetector.UpdateModelVersionOutput, error) {
	var output frauddetector.UpdateModelVersionOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type FrauddetectorUpdateModelVersionStatusResult struct {
	Result workflow.Future
}

func (r *FrauddetectorUpdateModelVersionStatusResult) Get(ctx workflow.Context) (*frauddetector.UpdateModelVersionStatusOutput, error) {
	var output frauddetector.UpdateModelVersionStatusOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type FrauddetectorUpdateRuleMetadataResult struct {
	Result workflow.Future
}

func (r *FrauddetectorUpdateRuleMetadataResult) Get(ctx workflow.Context) (*frauddetector.UpdateRuleMetadataOutput, error) {
	var output frauddetector.UpdateRuleMetadataOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type FrauddetectorUpdateRuleVersionResult struct {
	Result workflow.Future
}

func (r *FrauddetectorUpdateRuleVersionResult) Get(ctx workflow.Context) (*frauddetector.UpdateRuleVersionOutput, error) {
	var output frauddetector.UpdateRuleVersionOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type FrauddetectorUpdateVariableResult struct {
	Result workflow.Future
}

func (r *FrauddetectorUpdateVariableResult) Get(ctx workflow.Context) (*frauddetector.UpdateVariableOutput, error) {
	var output frauddetector.UpdateVariableOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type FraudDetectorStub struct {
	activities awsactivities.FraudDetectorActivities
}

func NewFraudDetectorStub() FraudDetectorClient {
	return &FraudDetectorStub{}
}

func (a *FraudDetectorStub) BatchCreateVariable(ctx workflow.Context, input *frauddetector.BatchCreateVariableInput) (*frauddetector.BatchCreateVariableOutput, error) {
	var output frauddetector.BatchCreateVariableOutput
	err := workflow.ExecuteActivity(ctx, a.activities.BatchCreateVariable, input).Get(ctx, &output)
	return &output, err
}

func (a *FraudDetectorStub) BatchCreateVariableAsync(ctx workflow.Context, input *frauddetector.BatchCreateVariableInput) *FrauddetectorBatchCreateVariableResult {
	future := workflow.ExecuteActivity(ctx, a.activities.BatchCreateVariable, input)
	return &FrauddetectorBatchCreateVariableResult{Result: future}
}

func (a *FraudDetectorStub) BatchGetVariable(ctx workflow.Context, input *frauddetector.BatchGetVariableInput) (*frauddetector.BatchGetVariableOutput, error) {
	var output frauddetector.BatchGetVariableOutput
	err := workflow.ExecuteActivity(ctx, a.activities.BatchGetVariable, input).Get(ctx, &output)
	return &output, err
}

func (a *FraudDetectorStub) BatchGetVariableAsync(ctx workflow.Context, input *frauddetector.BatchGetVariableInput) *FrauddetectorBatchGetVariableResult {
	future := workflow.ExecuteActivity(ctx, a.activities.BatchGetVariable, input)
	return &FrauddetectorBatchGetVariableResult{Result: future}
}

func (a *FraudDetectorStub) CreateDetectorVersion(ctx workflow.Context, input *frauddetector.CreateDetectorVersionInput) (*frauddetector.CreateDetectorVersionOutput, error) {
	var output frauddetector.CreateDetectorVersionOutput
	err := workflow.ExecuteActivity(ctx, a.activities.CreateDetectorVersion, input).Get(ctx, &output)
	return &output, err
}

func (a *FraudDetectorStub) CreateDetectorVersionAsync(ctx workflow.Context, input *frauddetector.CreateDetectorVersionInput) *FrauddetectorCreateDetectorVersionResult {
	future := workflow.ExecuteActivity(ctx, a.activities.CreateDetectorVersion, input)
	return &FrauddetectorCreateDetectorVersionResult{Result: future}
}

func (a *FraudDetectorStub) CreateModel(ctx workflow.Context, input *frauddetector.CreateModelInput) (*frauddetector.CreateModelOutput, error) {
	var output frauddetector.CreateModelOutput
	err := workflow.ExecuteActivity(ctx, a.activities.CreateModel, input).Get(ctx, &output)
	return &output, err
}

func (a *FraudDetectorStub) CreateModelAsync(ctx workflow.Context, input *frauddetector.CreateModelInput) *FrauddetectorCreateModelResult {
	future := workflow.ExecuteActivity(ctx, a.activities.CreateModel, input)
	return &FrauddetectorCreateModelResult{Result: future}
}

func (a *FraudDetectorStub) CreateModelVersion(ctx workflow.Context, input *frauddetector.CreateModelVersionInput) (*frauddetector.CreateModelVersionOutput, error) {
	var output frauddetector.CreateModelVersionOutput
	err := workflow.ExecuteActivity(ctx, a.activities.CreateModelVersion, input).Get(ctx, &output)
	return &output, err
}

func (a *FraudDetectorStub) CreateModelVersionAsync(ctx workflow.Context, input *frauddetector.CreateModelVersionInput) *FrauddetectorCreateModelVersionResult {
	future := workflow.ExecuteActivity(ctx, a.activities.CreateModelVersion, input)
	return &FrauddetectorCreateModelVersionResult{Result: future}
}

func (a *FraudDetectorStub) CreateRule(ctx workflow.Context, input *frauddetector.CreateRuleInput) (*frauddetector.CreateRuleOutput, error) {
	var output frauddetector.CreateRuleOutput
	err := workflow.ExecuteActivity(ctx, a.activities.CreateRule, input).Get(ctx, &output)
	return &output, err
}

func (a *FraudDetectorStub) CreateRuleAsync(ctx workflow.Context, input *frauddetector.CreateRuleInput) *FrauddetectorCreateRuleResult {
	future := workflow.ExecuteActivity(ctx, a.activities.CreateRule, input)
	return &FrauddetectorCreateRuleResult{Result: future}
}

func (a *FraudDetectorStub) CreateVariable(ctx workflow.Context, input *frauddetector.CreateVariableInput) (*frauddetector.CreateVariableOutput, error) {
	var output frauddetector.CreateVariableOutput
	err := workflow.ExecuteActivity(ctx, a.activities.CreateVariable, input).Get(ctx, &output)
	return &output, err
}

func (a *FraudDetectorStub) CreateVariableAsync(ctx workflow.Context, input *frauddetector.CreateVariableInput) *FrauddetectorCreateVariableResult {
	future := workflow.ExecuteActivity(ctx, a.activities.CreateVariable, input)
	return &FrauddetectorCreateVariableResult{Result: future}
}

func (a *FraudDetectorStub) DeleteDetector(ctx workflow.Context, input *frauddetector.DeleteDetectorInput) (*frauddetector.DeleteDetectorOutput, error) {
	var output frauddetector.DeleteDetectorOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeleteDetector, input).Get(ctx, &output)
	return &output, err
}

func (a *FraudDetectorStub) DeleteDetectorAsync(ctx workflow.Context, input *frauddetector.DeleteDetectorInput) *FrauddetectorDeleteDetectorResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeleteDetector, input)
	return &FrauddetectorDeleteDetectorResult{Result: future}
}

func (a *FraudDetectorStub) DeleteDetectorVersion(ctx workflow.Context, input *frauddetector.DeleteDetectorVersionInput) (*frauddetector.DeleteDetectorVersionOutput, error) {
	var output frauddetector.DeleteDetectorVersionOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeleteDetectorVersion, input).Get(ctx, &output)
	return &output, err
}

func (a *FraudDetectorStub) DeleteDetectorVersionAsync(ctx workflow.Context, input *frauddetector.DeleteDetectorVersionInput) *FrauddetectorDeleteDetectorVersionResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeleteDetectorVersion, input)
	return &FrauddetectorDeleteDetectorVersionResult{Result: future}
}

func (a *FraudDetectorStub) DeleteEvent(ctx workflow.Context, input *frauddetector.DeleteEventInput) (*frauddetector.DeleteEventOutput, error) {
	var output frauddetector.DeleteEventOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeleteEvent, input).Get(ctx, &output)
	return &output, err
}

func (a *FraudDetectorStub) DeleteEventAsync(ctx workflow.Context, input *frauddetector.DeleteEventInput) *FrauddetectorDeleteEventResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeleteEvent, input)
	return &FrauddetectorDeleteEventResult{Result: future}
}

func (a *FraudDetectorStub) DeleteRule(ctx workflow.Context, input *frauddetector.DeleteRuleInput) (*frauddetector.DeleteRuleOutput, error) {
	var output frauddetector.DeleteRuleOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeleteRule, input).Get(ctx, &output)
	return &output, err
}

func (a *FraudDetectorStub) DeleteRuleAsync(ctx workflow.Context, input *frauddetector.DeleteRuleInput) *FrauddetectorDeleteRuleResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeleteRule, input)
	return &FrauddetectorDeleteRuleResult{Result: future}
}

func (a *FraudDetectorStub) DescribeDetector(ctx workflow.Context, input *frauddetector.DescribeDetectorInput) (*frauddetector.DescribeDetectorOutput, error) {
	var output frauddetector.DescribeDetectorOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DescribeDetector, input).Get(ctx, &output)
	return &output, err
}

func (a *FraudDetectorStub) DescribeDetectorAsync(ctx workflow.Context, input *frauddetector.DescribeDetectorInput) *FrauddetectorDescribeDetectorResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DescribeDetector, input)
	return &FrauddetectorDescribeDetectorResult{Result: future}
}

func (a *FraudDetectorStub) DescribeModelVersions(ctx workflow.Context, input *frauddetector.DescribeModelVersionsInput) (*frauddetector.DescribeModelVersionsOutput, error) {
	var output frauddetector.DescribeModelVersionsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DescribeModelVersions, input).Get(ctx, &output)
	return &output, err
}

func (a *FraudDetectorStub) DescribeModelVersionsAsync(ctx workflow.Context, input *frauddetector.DescribeModelVersionsInput) *FrauddetectorDescribeModelVersionsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DescribeModelVersions, input)
	return &FrauddetectorDescribeModelVersionsResult{Result: future}
}

func (a *FraudDetectorStub) GetDetectorVersion(ctx workflow.Context, input *frauddetector.GetDetectorVersionInput) (*frauddetector.GetDetectorVersionOutput, error) {
	var output frauddetector.GetDetectorVersionOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetDetectorVersion, input).Get(ctx, &output)
	return &output, err
}

func (a *FraudDetectorStub) GetDetectorVersionAsync(ctx workflow.Context, input *frauddetector.GetDetectorVersionInput) *FrauddetectorGetDetectorVersionResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetDetectorVersion, input)
	return &FrauddetectorGetDetectorVersionResult{Result: future}
}

func (a *FraudDetectorStub) GetDetectors(ctx workflow.Context, input *frauddetector.GetDetectorsInput) (*frauddetector.GetDetectorsOutput, error) {
	var output frauddetector.GetDetectorsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetDetectors, input).Get(ctx, &output)
	return &output, err
}

func (a *FraudDetectorStub) GetDetectorsAsync(ctx workflow.Context, input *frauddetector.GetDetectorsInput) *FrauddetectorGetDetectorsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetDetectors, input)
	return &FrauddetectorGetDetectorsResult{Result: future}
}

func (a *FraudDetectorStub) GetEntityTypes(ctx workflow.Context, input *frauddetector.GetEntityTypesInput) (*frauddetector.GetEntityTypesOutput, error) {
	var output frauddetector.GetEntityTypesOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetEntityTypes, input).Get(ctx, &output)
	return &output, err
}

func (a *FraudDetectorStub) GetEntityTypesAsync(ctx workflow.Context, input *frauddetector.GetEntityTypesInput) *FrauddetectorGetEntityTypesResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetEntityTypes, input)
	return &FrauddetectorGetEntityTypesResult{Result: future}
}

func (a *FraudDetectorStub) GetEventPrediction(ctx workflow.Context, input *frauddetector.GetEventPredictionInput) (*frauddetector.GetEventPredictionOutput, error) {
	var output frauddetector.GetEventPredictionOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetEventPrediction, input).Get(ctx, &output)
	return &output, err
}

func (a *FraudDetectorStub) GetEventPredictionAsync(ctx workflow.Context, input *frauddetector.GetEventPredictionInput) *FrauddetectorGetEventPredictionResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetEventPrediction, input)
	return &FrauddetectorGetEventPredictionResult{Result: future}
}

func (a *FraudDetectorStub) GetEventTypes(ctx workflow.Context, input *frauddetector.GetEventTypesInput) (*frauddetector.GetEventTypesOutput, error) {
	var output frauddetector.GetEventTypesOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetEventTypes, input).Get(ctx, &output)
	return &output, err
}

func (a *FraudDetectorStub) GetEventTypesAsync(ctx workflow.Context, input *frauddetector.GetEventTypesInput) *FrauddetectorGetEventTypesResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetEventTypes, input)
	return &FrauddetectorGetEventTypesResult{Result: future}
}

func (a *FraudDetectorStub) GetExternalModels(ctx workflow.Context, input *frauddetector.GetExternalModelsInput) (*frauddetector.GetExternalModelsOutput, error) {
	var output frauddetector.GetExternalModelsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetExternalModels, input).Get(ctx, &output)
	return &output, err
}

func (a *FraudDetectorStub) GetExternalModelsAsync(ctx workflow.Context, input *frauddetector.GetExternalModelsInput) *FrauddetectorGetExternalModelsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetExternalModels, input)
	return &FrauddetectorGetExternalModelsResult{Result: future}
}

func (a *FraudDetectorStub) GetKMSEncryptionKey(ctx workflow.Context, input *frauddetector.GetKMSEncryptionKeyInput) (*frauddetector.GetKMSEncryptionKeyOutput, error) {
	var output frauddetector.GetKMSEncryptionKeyOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetKMSEncryptionKey, input).Get(ctx, &output)
	return &output, err
}

func (a *FraudDetectorStub) GetKMSEncryptionKeyAsync(ctx workflow.Context, input *frauddetector.GetKMSEncryptionKeyInput) *FrauddetectorGetKMSEncryptionKeyResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetKMSEncryptionKey, input)
	return &FrauddetectorGetKMSEncryptionKeyResult{Result: future}
}

func (a *FraudDetectorStub) GetLabels(ctx workflow.Context, input *frauddetector.GetLabelsInput) (*frauddetector.GetLabelsOutput, error) {
	var output frauddetector.GetLabelsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetLabels, input).Get(ctx, &output)
	return &output, err
}

func (a *FraudDetectorStub) GetLabelsAsync(ctx workflow.Context, input *frauddetector.GetLabelsInput) *FrauddetectorGetLabelsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetLabels, input)
	return &FrauddetectorGetLabelsResult{Result: future}
}

func (a *FraudDetectorStub) GetModelVersion(ctx workflow.Context, input *frauddetector.GetModelVersionInput) (*frauddetector.GetModelVersionOutput, error) {
	var output frauddetector.GetModelVersionOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetModelVersion, input).Get(ctx, &output)
	return &output, err
}

func (a *FraudDetectorStub) GetModelVersionAsync(ctx workflow.Context, input *frauddetector.GetModelVersionInput) *FrauddetectorGetModelVersionResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetModelVersion, input)
	return &FrauddetectorGetModelVersionResult{Result: future}
}

func (a *FraudDetectorStub) GetModels(ctx workflow.Context, input *frauddetector.GetModelsInput) (*frauddetector.GetModelsOutput, error) {
	var output frauddetector.GetModelsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetModels, input).Get(ctx, &output)
	return &output, err
}

func (a *FraudDetectorStub) GetModelsAsync(ctx workflow.Context, input *frauddetector.GetModelsInput) *FrauddetectorGetModelsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetModels, input)
	return &FrauddetectorGetModelsResult{Result: future}
}

func (a *FraudDetectorStub) GetOutcomes(ctx workflow.Context, input *frauddetector.GetOutcomesInput) (*frauddetector.GetOutcomesOutput, error) {
	var output frauddetector.GetOutcomesOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetOutcomes, input).Get(ctx, &output)
	return &output, err
}

func (a *FraudDetectorStub) GetOutcomesAsync(ctx workflow.Context, input *frauddetector.GetOutcomesInput) *FrauddetectorGetOutcomesResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetOutcomes, input)
	return &FrauddetectorGetOutcomesResult{Result: future}
}

func (a *FraudDetectorStub) GetRules(ctx workflow.Context, input *frauddetector.GetRulesInput) (*frauddetector.GetRulesOutput, error) {
	var output frauddetector.GetRulesOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetRules, input).Get(ctx, &output)
	return &output, err
}

func (a *FraudDetectorStub) GetRulesAsync(ctx workflow.Context, input *frauddetector.GetRulesInput) *FrauddetectorGetRulesResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetRules, input)
	return &FrauddetectorGetRulesResult{Result: future}
}

func (a *FraudDetectorStub) GetVariables(ctx workflow.Context, input *frauddetector.GetVariablesInput) (*frauddetector.GetVariablesOutput, error) {
	var output frauddetector.GetVariablesOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetVariables, input).Get(ctx, &output)
	return &output, err
}

func (a *FraudDetectorStub) GetVariablesAsync(ctx workflow.Context, input *frauddetector.GetVariablesInput) *FrauddetectorGetVariablesResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetVariables, input)
	return &FrauddetectorGetVariablesResult{Result: future}
}

func (a *FraudDetectorStub) ListTagsForResource(ctx workflow.Context, input *frauddetector.ListTagsForResourceInput) (*frauddetector.ListTagsForResourceOutput, error) {
	var output frauddetector.ListTagsForResourceOutput
	err := workflow.ExecuteActivity(ctx, a.activities.ListTagsForResource, input).Get(ctx, &output)
	return &output, err
}

func (a *FraudDetectorStub) ListTagsForResourceAsync(ctx workflow.Context, input *frauddetector.ListTagsForResourceInput) *FrauddetectorListTagsForResourceResult {
	future := workflow.ExecuteActivity(ctx, a.activities.ListTagsForResource, input)
	return &FrauddetectorListTagsForResourceResult{Result: future}
}

func (a *FraudDetectorStub) PutDetector(ctx workflow.Context, input *frauddetector.PutDetectorInput) (*frauddetector.PutDetectorOutput, error) {
	var output frauddetector.PutDetectorOutput
	err := workflow.ExecuteActivity(ctx, a.activities.PutDetector, input).Get(ctx, &output)
	return &output, err
}

func (a *FraudDetectorStub) PutDetectorAsync(ctx workflow.Context, input *frauddetector.PutDetectorInput) *FrauddetectorPutDetectorResult {
	future := workflow.ExecuteActivity(ctx, a.activities.PutDetector, input)
	return &FrauddetectorPutDetectorResult{Result: future}
}

func (a *FraudDetectorStub) PutEntityType(ctx workflow.Context, input *frauddetector.PutEntityTypeInput) (*frauddetector.PutEntityTypeOutput, error) {
	var output frauddetector.PutEntityTypeOutput
	err := workflow.ExecuteActivity(ctx, a.activities.PutEntityType, input).Get(ctx, &output)
	return &output, err
}

func (a *FraudDetectorStub) PutEntityTypeAsync(ctx workflow.Context, input *frauddetector.PutEntityTypeInput) *FrauddetectorPutEntityTypeResult {
	future := workflow.ExecuteActivity(ctx, a.activities.PutEntityType, input)
	return &FrauddetectorPutEntityTypeResult{Result: future}
}

func (a *FraudDetectorStub) PutEventType(ctx workflow.Context, input *frauddetector.PutEventTypeInput) (*frauddetector.PutEventTypeOutput, error) {
	var output frauddetector.PutEventTypeOutput
	err := workflow.ExecuteActivity(ctx, a.activities.PutEventType, input).Get(ctx, &output)
	return &output, err
}

func (a *FraudDetectorStub) PutEventTypeAsync(ctx workflow.Context, input *frauddetector.PutEventTypeInput) *FrauddetectorPutEventTypeResult {
	future := workflow.ExecuteActivity(ctx, a.activities.PutEventType, input)
	return &FrauddetectorPutEventTypeResult{Result: future}
}

func (a *FraudDetectorStub) PutExternalModel(ctx workflow.Context, input *frauddetector.PutExternalModelInput) (*frauddetector.PutExternalModelOutput, error) {
	var output frauddetector.PutExternalModelOutput
	err := workflow.ExecuteActivity(ctx, a.activities.PutExternalModel, input).Get(ctx, &output)
	return &output, err
}

func (a *FraudDetectorStub) PutExternalModelAsync(ctx workflow.Context, input *frauddetector.PutExternalModelInput) *FrauddetectorPutExternalModelResult {
	future := workflow.ExecuteActivity(ctx, a.activities.PutExternalModel, input)
	return &FrauddetectorPutExternalModelResult{Result: future}
}

func (a *FraudDetectorStub) PutKMSEncryptionKey(ctx workflow.Context, input *frauddetector.PutKMSEncryptionKeyInput) (*frauddetector.PutKMSEncryptionKeyOutput, error) {
	var output frauddetector.PutKMSEncryptionKeyOutput
	err := workflow.ExecuteActivity(ctx, a.activities.PutKMSEncryptionKey, input).Get(ctx, &output)
	return &output, err
}

func (a *FraudDetectorStub) PutKMSEncryptionKeyAsync(ctx workflow.Context, input *frauddetector.PutKMSEncryptionKeyInput) *FrauddetectorPutKMSEncryptionKeyResult {
	future := workflow.ExecuteActivity(ctx, a.activities.PutKMSEncryptionKey, input)
	return &FrauddetectorPutKMSEncryptionKeyResult{Result: future}
}

func (a *FraudDetectorStub) PutLabel(ctx workflow.Context, input *frauddetector.PutLabelInput) (*frauddetector.PutLabelOutput, error) {
	var output frauddetector.PutLabelOutput
	err := workflow.ExecuteActivity(ctx, a.activities.PutLabel, input).Get(ctx, &output)
	return &output, err
}

func (a *FraudDetectorStub) PutLabelAsync(ctx workflow.Context, input *frauddetector.PutLabelInput) *FrauddetectorPutLabelResult {
	future := workflow.ExecuteActivity(ctx, a.activities.PutLabel, input)
	return &FrauddetectorPutLabelResult{Result: future}
}

func (a *FraudDetectorStub) PutOutcome(ctx workflow.Context, input *frauddetector.PutOutcomeInput) (*frauddetector.PutOutcomeOutput, error) {
	var output frauddetector.PutOutcomeOutput
	err := workflow.ExecuteActivity(ctx, a.activities.PutOutcome, input).Get(ctx, &output)
	return &output, err
}

func (a *FraudDetectorStub) PutOutcomeAsync(ctx workflow.Context, input *frauddetector.PutOutcomeInput) *FrauddetectorPutOutcomeResult {
	future := workflow.ExecuteActivity(ctx, a.activities.PutOutcome, input)
	return &FrauddetectorPutOutcomeResult{Result: future}
}

func (a *FraudDetectorStub) TagResource(ctx workflow.Context, input *frauddetector.TagResourceInput) (*frauddetector.TagResourceOutput, error) {
	var output frauddetector.TagResourceOutput
	err := workflow.ExecuteActivity(ctx, a.activities.TagResource, input).Get(ctx, &output)
	return &output, err
}

func (a *FraudDetectorStub) TagResourceAsync(ctx workflow.Context, input *frauddetector.TagResourceInput) *FrauddetectorTagResourceResult {
	future := workflow.ExecuteActivity(ctx, a.activities.TagResource, input)
	return &FrauddetectorTagResourceResult{Result: future}
}

func (a *FraudDetectorStub) UntagResource(ctx workflow.Context, input *frauddetector.UntagResourceInput) (*frauddetector.UntagResourceOutput, error) {
	var output frauddetector.UntagResourceOutput
	err := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input).Get(ctx, &output)
	return &output, err
}

func (a *FraudDetectorStub) UntagResourceAsync(ctx workflow.Context, input *frauddetector.UntagResourceInput) *FrauddetectorUntagResourceResult {
	future := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input)
	return &FrauddetectorUntagResourceResult{Result: future}
}

func (a *FraudDetectorStub) UpdateDetectorVersion(ctx workflow.Context, input *frauddetector.UpdateDetectorVersionInput) (*frauddetector.UpdateDetectorVersionOutput, error) {
	var output frauddetector.UpdateDetectorVersionOutput
	err := workflow.ExecuteActivity(ctx, a.activities.UpdateDetectorVersion, input).Get(ctx, &output)
	return &output, err
}

func (a *FraudDetectorStub) UpdateDetectorVersionAsync(ctx workflow.Context, input *frauddetector.UpdateDetectorVersionInput) *FrauddetectorUpdateDetectorVersionResult {
	future := workflow.ExecuteActivity(ctx, a.activities.UpdateDetectorVersion, input)
	return &FrauddetectorUpdateDetectorVersionResult{Result: future}
}

func (a *FraudDetectorStub) UpdateDetectorVersionMetadata(ctx workflow.Context, input *frauddetector.UpdateDetectorVersionMetadataInput) (*frauddetector.UpdateDetectorVersionMetadataOutput, error) {
	var output frauddetector.UpdateDetectorVersionMetadataOutput
	err := workflow.ExecuteActivity(ctx, a.activities.UpdateDetectorVersionMetadata, input).Get(ctx, &output)
	return &output, err
}

func (a *FraudDetectorStub) UpdateDetectorVersionMetadataAsync(ctx workflow.Context, input *frauddetector.UpdateDetectorVersionMetadataInput) *FrauddetectorUpdateDetectorVersionMetadataResult {
	future := workflow.ExecuteActivity(ctx, a.activities.UpdateDetectorVersionMetadata, input)
	return &FrauddetectorUpdateDetectorVersionMetadataResult{Result: future}
}

func (a *FraudDetectorStub) UpdateDetectorVersionStatus(ctx workflow.Context, input *frauddetector.UpdateDetectorVersionStatusInput) (*frauddetector.UpdateDetectorVersionStatusOutput, error) {
	var output frauddetector.UpdateDetectorVersionStatusOutput
	err := workflow.ExecuteActivity(ctx, a.activities.UpdateDetectorVersionStatus, input).Get(ctx, &output)
	return &output, err
}

func (a *FraudDetectorStub) UpdateDetectorVersionStatusAsync(ctx workflow.Context, input *frauddetector.UpdateDetectorVersionStatusInput) *FrauddetectorUpdateDetectorVersionStatusResult {
	future := workflow.ExecuteActivity(ctx, a.activities.UpdateDetectorVersionStatus, input)
	return &FrauddetectorUpdateDetectorVersionStatusResult{Result: future}
}

func (a *FraudDetectorStub) UpdateModel(ctx workflow.Context, input *frauddetector.UpdateModelInput) (*frauddetector.UpdateModelOutput, error) {
	var output frauddetector.UpdateModelOutput
	err := workflow.ExecuteActivity(ctx, a.activities.UpdateModel, input).Get(ctx, &output)
	return &output, err
}

func (a *FraudDetectorStub) UpdateModelAsync(ctx workflow.Context, input *frauddetector.UpdateModelInput) *FrauddetectorUpdateModelResult {
	future := workflow.ExecuteActivity(ctx, a.activities.UpdateModel, input)
	return &FrauddetectorUpdateModelResult{Result: future}
}

func (a *FraudDetectorStub) UpdateModelVersion(ctx workflow.Context, input *frauddetector.UpdateModelVersionInput) (*frauddetector.UpdateModelVersionOutput, error) {
	var output frauddetector.UpdateModelVersionOutput
	err := workflow.ExecuteActivity(ctx, a.activities.UpdateModelVersion, input).Get(ctx, &output)
	return &output, err
}

func (a *FraudDetectorStub) UpdateModelVersionAsync(ctx workflow.Context, input *frauddetector.UpdateModelVersionInput) *FrauddetectorUpdateModelVersionResult {
	future := workflow.ExecuteActivity(ctx, a.activities.UpdateModelVersion, input)
	return &FrauddetectorUpdateModelVersionResult{Result: future}
}

func (a *FraudDetectorStub) UpdateModelVersionStatus(ctx workflow.Context, input *frauddetector.UpdateModelVersionStatusInput) (*frauddetector.UpdateModelVersionStatusOutput, error) {
	var output frauddetector.UpdateModelVersionStatusOutput
	err := workflow.ExecuteActivity(ctx, a.activities.UpdateModelVersionStatus, input).Get(ctx, &output)
	return &output, err
}

func (a *FraudDetectorStub) UpdateModelVersionStatusAsync(ctx workflow.Context, input *frauddetector.UpdateModelVersionStatusInput) *FrauddetectorUpdateModelVersionStatusResult {
	future := workflow.ExecuteActivity(ctx, a.activities.UpdateModelVersionStatus, input)
	return &FrauddetectorUpdateModelVersionStatusResult{Result: future}
}

func (a *FraudDetectorStub) UpdateRuleMetadata(ctx workflow.Context, input *frauddetector.UpdateRuleMetadataInput) (*frauddetector.UpdateRuleMetadataOutput, error) {
	var output frauddetector.UpdateRuleMetadataOutput
	err := workflow.ExecuteActivity(ctx, a.activities.UpdateRuleMetadata, input).Get(ctx, &output)
	return &output, err
}

func (a *FraudDetectorStub) UpdateRuleMetadataAsync(ctx workflow.Context, input *frauddetector.UpdateRuleMetadataInput) *FrauddetectorUpdateRuleMetadataResult {
	future := workflow.ExecuteActivity(ctx, a.activities.UpdateRuleMetadata, input)
	return &FrauddetectorUpdateRuleMetadataResult{Result: future}
}

func (a *FraudDetectorStub) UpdateRuleVersion(ctx workflow.Context, input *frauddetector.UpdateRuleVersionInput) (*frauddetector.UpdateRuleVersionOutput, error) {
	var output frauddetector.UpdateRuleVersionOutput
	err := workflow.ExecuteActivity(ctx, a.activities.UpdateRuleVersion, input).Get(ctx, &output)
	return &output, err
}

func (a *FraudDetectorStub) UpdateRuleVersionAsync(ctx workflow.Context, input *frauddetector.UpdateRuleVersionInput) *FrauddetectorUpdateRuleVersionResult {
	future := workflow.ExecuteActivity(ctx, a.activities.UpdateRuleVersion, input)
	return &FrauddetectorUpdateRuleVersionResult{Result: future}
}

func (a *FraudDetectorStub) UpdateVariable(ctx workflow.Context, input *frauddetector.UpdateVariableInput) (*frauddetector.UpdateVariableOutput, error) {
	var output frauddetector.UpdateVariableOutput
	err := workflow.ExecuteActivity(ctx, a.activities.UpdateVariable, input).Get(ctx, &output)
	return &output, err
}

func (a *FraudDetectorStub) UpdateVariableAsync(ctx workflow.Context, input *frauddetector.UpdateVariableInput) *FrauddetectorUpdateVariableResult {
	future := workflow.ExecuteActivity(ctx, a.activities.UpdateVariable, input)
	return &FrauddetectorUpdateVariableResult{Result: future}
}
