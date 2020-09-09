package awsclients

import (
	"github.com/aws/aws-sdk-go/service/comprehend"
	"go.temporal.io/sdk/workflow"
	"temporal.io/aws-sdk/awsactivities"
)

type ComprehendClient interface {
       BatchDetectDominantLanguage(ctx workflow.Context, input *comprehend.BatchDetectDominantLanguageInput) (*comprehend.BatchDetectDominantLanguageOutput, error)
       BatchDetectDominantLanguageAsync(ctx workflow.Context, input *comprehend.BatchDetectDominantLanguageInput) *ComprehendBatchDetectDominantLanguageResult

       BatchDetectEntities(ctx workflow.Context, input *comprehend.BatchDetectEntitiesInput) (*comprehend.BatchDetectEntitiesOutput, error)
       BatchDetectEntitiesAsync(ctx workflow.Context, input *comprehend.BatchDetectEntitiesInput) *ComprehendBatchDetectEntitiesResult

       BatchDetectKeyPhrases(ctx workflow.Context, input *comprehend.BatchDetectKeyPhrasesInput) (*comprehend.BatchDetectKeyPhrasesOutput, error)
       BatchDetectKeyPhrasesAsync(ctx workflow.Context, input *comprehend.BatchDetectKeyPhrasesInput) *ComprehendBatchDetectKeyPhrasesResult

       BatchDetectSentiment(ctx workflow.Context, input *comprehend.BatchDetectSentimentInput) (*comprehend.BatchDetectSentimentOutput, error)
       BatchDetectSentimentAsync(ctx workflow.Context, input *comprehend.BatchDetectSentimentInput) *ComprehendBatchDetectSentimentResult

       BatchDetectSyntax(ctx workflow.Context, input *comprehend.BatchDetectSyntaxInput) (*comprehend.BatchDetectSyntaxOutput, error)
       BatchDetectSyntaxAsync(ctx workflow.Context, input *comprehend.BatchDetectSyntaxInput) *ComprehendBatchDetectSyntaxResult

       ClassifyDocument(ctx workflow.Context, input *comprehend.ClassifyDocumentInput) (*comprehend.ClassifyDocumentOutput, error)
       ClassifyDocumentAsync(ctx workflow.Context, input *comprehend.ClassifyDocumentInput) *ComprehendClassifyDocumentResult

       CreateDocumentClassifier(ctx workflow.Context, input *comprehend.CreateDocumentClassifierInput) (*comprehend.CreateDocumentClassifierOutput, error)
       CreateDocumentClassifierAsync(ctx workflow.Context, input *comprehend.CreateDocumentClassifierInput) *ComprehendCreateDocumentClassifierResult

       CreateEndpoint(ctx workflow.Context, input *comprehend.CreateEndpointInput) (*comprehend.CreateEndpointOutput, error)
       CreateEndpointAsync(ctx workflow.Context, input *comprehend.CreateEndpointInput) *ComprehendCreateEndpointResult

       CreateEntityRecognizer(ctx workflow.Context, input *comprehend.CreateEntityRecognizerInput) (*comprehend.CreateEntityRecognizerOutput, error)
       CreateEntityRecognizerAsync(ctx workflow.Context, input *comprehend.CreateEntityRecognizerInput) *ComprehendCreateEntityRecognizerResult

       DeleteDocumentClassifier(ctx workflow.Context, input *comprehend.DeleteDocumentClassifierInput) (*comprehend.DeleteDocumentClassifierOutput, error)
       DeleteDocumentClassifierAsync(ctx workflow.Context, input *comprehend.DeleteDocumentClassifierInput) *ComprehendDeleteDocumentClassifierResult

       DeleteEndpoint(ctx workflow.Context, input *comprehend.DeleteEndpointInput) (*comprehend.DeleteEndpointOutput, error)
       DeleteEndpointAsync(ctx workflow.Context, input *comprehend.DeleteEndpointInput) *ComprehendDeleteEndpointResult

       DeleteEntityRecognizer(ctx workflow.Context, input *comprehend.DeleteEntityRecognizerInput) (*comprehend.DeleteEntityRecognizerOutput, error)
       DeleteEntityRecognizerAsync(ctx workflow.Context, input *comprehend.DeleteEntityRecognizerInput) *ComprehendDeleteEntityRecognizerResult

       DescribeDocumentClassificationJob(ctx workflow.Context, input *comprehend.DescribeDocumentClassificationJobInput) (*comprehend.DescribeDocumentClassificationJobOutput, error)
       DescribeDocumentClassificationJobAsync(ctx workflow.Context, input *comprehend.DescribeDocumentClassificationJobInput) *ComprehendDescribeDocumentClassificationJobResult

       DescribeDocumentClassifier(ctx workflow.Context, input *comprehend.DescribeDocumentClassifierInput) (*comprehend.DescribeDocumentClassifierOutput, error)
       DescribeDocumentClassifierAsync(ctx workflow.Context, input *comprehend.DescribeDocumentClassifierInput) *ComprehendDescribeDocumentClassifierResult

       DescribeDominantLanguageDetectionJob(ctx workflow.Context, input *comprehend.DescribeDominantLanguageDetectionJobInput) (*comprehend.DescribeDominantLanguageDetectionJobOutput, error)
       DescribeDominantLanguageDetectionJobAsync(ctx workflow.Context, input *comprehend.DescribeDominantLanguageDetectionJobInput) *ComprehendDescribeDominantLanguageDetectionJobResult

       DescribeEndpoint(ctx workflow.Context, input *comprehend.DescribeEndpointInput) (*comprehend.DescribeEndpointOutput, error)
       DescribeEndpointAsync(ctx workflow.Context, input *comprehend.DescribeEndpointInput) *ComprehendDescribeEndpointResult

       DescribeEntitiesDetectionJob(ctx workflow.Context, input *comprehend.DescribeEntitiesDetectionJobInput) (*comprehend.DescribeEntitiesDetectionJobOutput, error)
       DescribeEntitiesDetectionJobAsync(ctx workflow.Context, input *comprehend.DescribeEntitiesDetectionJobInput) *ComprehendDescribeEntitiesDetectionJobResult

       DescribeEntityRecognizer(ctx workflow.Context, input *comprehend.DescribeEntityRecognizerInput) (*comprehend.DescribeEntityRecognizerOutput, error)
       DescribeEntityRecognizerAsync(ctx workflow.Context, input *comprehend.DescribeEntityRecognizerInput) *ComprehendDescribeEntityRecognizerResult

       DescribeKeyPhrasesDetectionJob(ctx workflow.Context, input *comprehend.DescribeKeyPhrasesDetectionJobInput) (*comprehend.DescribeKeyPhrasesDetectionJobOutput, error)
       DescribeKeyPhrasesDetectionJobAsync(ctx workflow.Context, input *comprehend.DescribeKeyPhrasesDetectionJobInput) *ComprehendDescribeKeyPhrasesDetectionJobResult

       DescribeSentimentDetectionJob(ctx workflow.Context, input *comprehend.DescribeSentimentDetectionJobInput) (*comprehend.DescribeSentimentDetectionJobOutput, error)
       DescribeSentimentDetectionJobAsync(ctx workflow.Context, input *comprehend.DescribeSentimentDetectionJobInput) *ComprehendDescribeSentimentDetectionJobResult

       DescribeTopicsDetectionJob(ctx workflow.Context, input *comprehend.DescribeTopicsDetectionJobInput) (*comprehend.DescribeTopicsDetectionJobOutput, error)
       DescribeTopicsDetectionJobAsync(ctx workflow.Context, input *comprehend.DescribeTopicsDetectionJobInput) *ComprehendDescribeTopicsDetectionJobResult

       DetectDominantLanguage(ctx workflow.Context, input *comprehend.DetectDominantLanguageInput) (*comprehend.DetectDominantLanguageOutput, error)
       DetectDominantLanguageAsync(ctx workflow.Context, input *comprehend.DetectDominantLanguageInput) *ComprehendDetectDominantLanguageResult

       DetectEntities(ctx workflow.Context, input *comprehend.DetectEntitiesInput) (*comprehend.DetectEntitiesOutput, error)
       DetectEntitiesAsync(ctx workflow.Context, input *comprehend.DetectEntitiesInput) *ComprehendDetectEntitiesResult

       DetectKeyPhrases(ctx workflow.Context, input *comprehend.DetectKeyPhrasesInput) (*comprehend.DetectKeyPhrasesOutput, error)
       DetectKeyPhrasesAsync(ctx workflow.Context, input *comprehend.DetectKeyPhrasesInput) *ComprehendDetectKeyPhrasesResult

       DetectSentiment(ctx workflow.Context, input *comprehend.DetectSentimentInput) (*comprehend.DetectSentimentOutput, error)
       DetectSentimentAsync(ctx workflow.Context, input *comprehend.DetectSentimentInput) *ComprehendDetectSentimentResult

       DetectSyntax(ctx workflow.Context, input *comprehend.DetectSyntaxInput) (*comprehend.DetectSyntaxOutput, error)
       DetectSyntaxAsync(ctx workflow.Context, input *comprehend.DetectSyntaxInput) *ComprehendDetectSyntaxResult

       ListDocumentClassificationJobs(ctx workflow.Context, input *comprehend.ListDocumentClassificationJobsInput) (*comprehend.ListDocumentClassificationJobsOutput, error)
       ListDocumentClassificationJobsAsync(ctx workflow.Context, input *comprehend.ListDocumentClassificationJobsInput) *ComprehendListDocumentClassificationJobsResult

       ListDocumentClassifiers(ctx workflow.Context, input *comprehend.ListDocumentClassifiersInput) (*comprehend.ListDocumentClassifiersOutput, error)
       ListDocumentClassifiersAsync(ctx workflow.Context, input *comprehend.ListDocumentClassifiersInput) *ComprehendListDocumentClassifiersResult

       ListDominantLanguageDetectionJobs(ctx workflow.Context, input *comprehend.ListDominantLanguageDetectionJobsInput) (*comprehend.ListDominantLanguageDetectionJobsOutput, error)
       ListDominantLanguageDetectionJobsAsync(ctx workflow.Context, input *comprehend.ListDominantLanguageDetectionJobsInput) *ComprehendListDominantLanguageDetectionJobsResult

       ListEndpoints(ctx workflow.Context, input *comprehend.ListEndpointsInput) (*comprehend.ListEndpointsOutput, error)
       ListEndpointsAsync(ctx workflow.Context, input *comprehend.ListEndpointsInput) *ComprehendListEndpointsResult

       ListEntitiesDetectionJobs(ctx workflow.Context, input *comprehend.ListEntitiesDetectionJobsInput) (*comprehend.ListEntitiesDetectionJobsOutput, error)
       ListEntitiesDetectionJobsAsync(ctx workflow.Context, input *comprehend.ListEntitiesDetectionJobsInput) *ComprehendListEntitiesDetectionJobsResult

       ListEntityRecognizers(ctx workflow.Context, input *comprehend.ListEntityRecognizersInput) (*comprehend.ListEntityRecognizersOutput, error)
       ListEntityRecognizersAsync(ctx workflow.Context, input *comprehend.ListEntityRecognizersInput) *ComprehendListEntityRecognizersResult

       ListKeyPhrasesDetectionJobs(ctx workflow.Context, input *comprehend.ListKeyPhrasesDetectionJobsInput) (*comprehend.ListKeyPhrasesDetectionJobsOutput, error)
       ListKeyPhrasesDetectionJobsAsync(ctx workflow.Context, input *comprehend.ListKeyPhrasesDetectionJobsInput) *ComprehendListKeyPhrasesDetectionJobsResult

       ListSentimentDetectionJobs(ctx workflow.Context, input *comprehend.ListSentimentDetectionJobsInput) (*comprehend.ListSentimentDetectionJobsOutput, error)
       ListSentimentDetectionJobsAsync(ctx workflow.Context, input *comprehend.ListSentimentDetectionJobsInput) *ComprehendListSentimentDetectionJobsResult

       ListTagsForResource(ctx workflow.Context, input *comprehend.ListTagsForResourceInput) (*comprehend.ListTagsForResourceOutput, error)
       ListTagsForResourceAsync(ctx workflow.Context, input *comprehend.ListTagsForResourceInput) *ComprehendListTagsForResourceResult

       ListTopicsDetectionJobs(ctx workflow.Context, input *comprehend.ListTopicsDetectionJobsInput) (*comprehend.ListTopicsDetectionJobsOutput, error)
       ListTopicsDetectionJobsAsync(ctx workflow.Context, input *comprehend.ListTopicsDetectionJobsInput) *ComprehendListTopicsDetectionJobsResult

       StartDocumentClassificationJob(ctx workflow.Context, input *comprehend.StartDocumentClassificationJobInput) (*comprehend.StartDocumentClassificationJobOutput, error)
       StartDocumentClassificationJobAsync(ctx workflow.Context, input *comprehend.StartDocumentClassificationJobInput) *ComprehendStartDocumentClassificationJobResult

       StartDominantLanguageDetectionJob(ctx workflow.Context, input *comprehend.StartDominantLanguageDetectionJobInput) (*comprehend.StartDominantLanguageDetectionJobOutput, error)
       StartDominantLanguageDetectionJobAsync(ctx workflow.Context, input *comprehend.StartDominantLanguageDetectionJobInput) *ComprehendStartDominantLanguageDetectionJobResult

       StartEntitiesDetectionJob(ctx workflow.Context, input *comprehend.StartEntitiesDetectionJobInput) (*comprehend.StartEntitiesDetectionJobOutput, error)
       StartEntitiesDetectionJobAsync(ctx workflow.Context, input *comprehend.StartEntitiesDetectionJobInput) *ComprehendStartEntitiesDetectionJobResult

       StartKeyPhrasesDetectionJob(ctx workflow.Context, input *comprehend.StartKeyPhrasesDetectionJobInput) (*comprehend.StartKeyPhrasesDetectionJobOutput, error)
       StartKeyPhrasesDetectionJobAsync(ctx workflow.Context, input *comprehend.StartKeyPhrasesDetectionJobInput) *ComprehendStartKeyPhrasesDetectionJobResult

       StartSentimentDetectionJob(ctx workflow.Context, input *comprehend.StartSentimentDetectionJobInput) (*comprehend.StartSentimentDetectionJobOutput, error)
       StartSentimentDetectionJobAsync(ctx workflow.Context, input *comprehend.StartSentimentDetectionJobInput) *ComprehendStartSentimentDetectionJobResult

       StartTopicsDetectionJob(ctx workflow.Context, input *comprehend.StartTopicsDetectionJobInput) (*comprehend.StartTopicsDetectionJobOutput, error)
       StartTopicsDetectionJobAsync(ctx workflow.Context, input *comprehend.StartTopicsDetectionJobInput) *ComprehendStartTopicsDetectionJobResult

       StopDominantLanguageDetectionJob(ctx workflow.Context, input *comprehend.StopDominantLanguageDetectionJobInput) (*comprehend.StopDominantLanguageDetectionJobOutput, error)
       StopDominantLanguageDetectionJobAsync(ctx workflow.Context, input *comprehend.StopDominantLanguageDetectionJobInput) *ComprehendStopDominantLanguageDetectionJobResult

       StopEntitiesDetectionJob(ctx workflow.Context, input *comprehend.StopEntitiesDetectionJobInput) (*comprehend.StopEntitiesDetectionJobOutput, error)
       StopEntitiesDetectionJobAsync(ctx workflow.Context, input *comprehend.StopEntitiesDetectionJobInput) *ComprehendStopEntitiesDetectionJobResult

       StopKeyPhrasesDetectionJob(ctx workflow.Context, input *comprehend.StopKeyPhrasesDetectionJobInput) (*comprehend.StopKeyPhrasesDetectionJobOutput, error)
       StopKeyPhrasesDetectionJobAsync(ctx workflow.Context, input *comprehend.StopKeyPhrasesDetectionJobInput) *ComprehendStopKeyPhrasesDetectionJobResult

       StopSentimentDetectionJob(ctx workflow.Context, input *comprehend.StopSentimentDetectionJobInput) (*comprehend.StopSentimentDetectionJobOutput, error)
       StopSentimentDetectionJobAsync(ctx workflow.Context, input *comprehend.StopSentimentDetectionJobInput) *ComprehendStopSentimentDetectionJobResult

       StopTrainingDocumentClassifier(ctx workflow.Context, input *comprehend.StopTrainingDocumentClassifierInput) (*comprehend.StopTrainingDocumentClassifierOutput, error)
       StopTrainingDocumentClassifierAsync(ctx workflow.Context, input *comprehend.StopTrainingDocumentClassifierInput) *ComprehendStopTrainingDocumentClassifierResult

       StopTrainingEntityRecognizer(ctx workflow.Context, input *comprehend.StopTrainingEntityRecognizerInput) (*comprehend.StopTrainingEntityRecognizerOutput, error)
       StopTrainingEntityRecognizerAsync(ctx workflow.Context, input *comprehend.StopTrainingEntityRecognizerInput) *ComprehendStopTrainingEntityRecognizerResult

       TagResource(ctx workflow.Context, input *comprehend.TagResourceInput) (*comprehend.TagResourceOutput, error)
       TagResourceAsync(ctx workflow.Context, input *comprehend.TagResourceInput) *ComprehendTagResourceResult

       UntagResource(ctx workflow.Context, input *comprehend.UntagResourceInput) (*comprehend.UntagResourceOutput, error)
       UntagResourceAsync(ctx workflow.Context, input *comprehend.UntagResourceInput) *ComprehendUntagResourceResult

       UpdateEndpoint(ctx workflow.Context, input *comprehend.UpdateEndpointInput) (*comprehend.UpdateEndpointOutput, error)
       UpdateEndpointAsync(ctx workflow.Context, input *comprehend.UpdateEndpointInput) *ComprehendUpdateEndpointResult
}

type ComprehendBatchDetectDominantLanguageResult struct {
	Result workflow.Future
}

func (r *ComprehendBatchDetectDominantLanguageResult) Get(ctx workflow.Context) (*comprehend.BatchDetectDominantLanguageOutput, error) {
    var output comprehend.BatchDetectDominantLanguageOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ComprehendBatchDetectEntitiesResult struct {
	Result workflow.Future
}

func (r *ComprehendBatchDetectEntitiesResult) Get(ctx workflow.Context) (*comprehend.BatchDetectEntitiesOutput, error) {
    var output comprehend.BatchDetectEntitiesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ComprehendBatchDetectKeyPhrasesResult struct {
	Result workflow.Future
}

func (r *ComprehendBatchDetectKeyPhrasesResult) Get(ctx workflow.Context) (*comprehend.BatchDetectKeyPhrasesOutput, error) {
    var output comprehend.BatchDetectKeyPhrasesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ComprehendBatchDetectSentimentResult struct {
	Result workflow.Future
}

func (r *ComprehendBatchDetectSentimentResult) Get(ctx workflow.Context) (*comprehend.BatchDetectSentimentOutput, error) {
    var output comprehend.BatchDetectSentimentOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ComprehendBatchDetectSyntaxResult struct {
	Result workflow.Future
}

func (r *ComprehendBatchDetectSyntaxResult) Get(ctx workflow.Context) (*comprehend.BatchDetectSyntaxOutput, error) {
    var output comprehend.BatchDetectSyntaxOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ComprehendClassifyDocumentResult struct {
	Result workflow.Future
}

func (r *ComprehendClassifyDocumentResult) Get(ctx workflow.Context) (*comprehend.ClassifyDocumentOutput, error) {
    var output comprehend.ClassifyDocumentOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ComprehendCreateDocumentClassifierResult struct {
	Result workflow.Future
}

func (r *ComprehendCreateDocumentClassifierResult) Get(ctx workflow.Context) (*comprehend.CreateDocumentClassifierOutput, error) {
    var output comprehend.CreateDocumentClassifierOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ComprehendCreateEndpointResult struct {
	Result workflow.Future
}

func (r *ComprehendCreateEndpointResult) Get(ctx workflow.Context) (*comprehend.CreateEndpointOutput, error) {
    var output comprehend.CreateEndpointOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ComprehendCreateEntityRecognizerResult struct {
	Result workflow.Future
}

func (r *ComprehendCreateEntityRecognizerResult) Get(ctx workflow.Context) (*comprehend.CreateEntityRecognizerOutput, error) {
    var output comprehend.CreateEntityRecognizerOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ComprehendDeleteDocumentClassifierResult struct {
	Result workflow.Future
}

func (r *ComprehendDeleteDocumentClassifierResult) Get(ctx workflow.Context) (*comprehend.DeleteDocumentClassifierOutput, error) {
    var output comprehend.DeleteDocumentClassifierOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ComprehendDeleteEndpointResult struct {
	Result workflow.Future
}

func (r *ComprehendDeleteEndpointResult) Get(ctx workflow.Context) (*comprehend.DeleteEndpointOutput, error) {
    var output comprehend.DeleteEndpointOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ComprehendDeleteEntityRecognizerResult struct {
	Result workflow.Future
}

func (r *ComprehendDeleteEntityRecognizerResult) Get(ctx workflow.Context) (*comprehend.DeleteEntityRecognizerOutput, error) {
    var output comprehend.DeleteEntityRecognizerOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ComprehendDescribeDocumentClassificationJobResult struct {
	Result workflow.Future
}

func (r *ComprehendDescribeDocumentClassificationJobResult) Get(ctx workflow.Context) (*comprehend.DescribeDocumentClassificationJobOutput, error) {
    var output comprehend.DescribeDocumentClassificationJobOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ComprehendDescribeDocumentClassifierResult struct {
	Result workflow.Future
}

func (r *ComprehendDescribeDocumentClassifierResult) Get(ctx workflow.Context) (*comprehend.DescribeDocumentClassifierOutput, error) {
    var output comprehend.DescribeDocumentClassifierOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ComprehendDescribeDominantLanguageDetectionJobResult struct {
	Result workflow.Future
}

func (r *ComprehendDescribeDominantLanguageDetectionJobResult) Get(ctx workflow.Context) (*comprehend.DescribeDominantLanguageDetectionJobOutput, error) {
    var output comprehend.DescribeDominantLanguageDetectionJobOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ComprehendDescribeEndpointResult struct {
	Result workflow.Future
}

func (r *ComprehendDescribeEndpointResult) Get(ctx workflow.Context) (*comprehend.DescribeEndpointOutput, error) {
    var output comprehend.DescribeEndpointOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ComprehendDescribeEntitiesDetectionJobResult struct {
	Result workflow.Future
}

func (r *ComprehendDescribeEntitiesDetectionJobResult) Get(ctx workflow.Context) (*comprehend.DescribeEntitiesDetectionJobOutput, error) {
    var output comprehend.DescribeEntitiesDetectionJobOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ComprehendDescribeEntityRecognizerResult struct {
	Result workflow.Future
}

func (r *ComprehendDescribeEntityRecognizerResult) Get(ctx workflow.Context) (*comprehend.DescribeEntityRecognizerOutput, error) {
    var output comprehend.DescribeEntityRecognizerOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ComprehendDescribeKeyPhrasesDetectionJobResult struct {
	Result workflow.Future
}

func (r *ComprehendDescribeKeyPhrasesDetectionJobResult) Get(ctx workflow.Context) (*comprehend.DescribeKeyPhrasesDetectionJobOutput, error) {
    var output comprehend.DescribeKeyPhrasesDetectionJobOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ComprehendDescribeSentimentDetectionJobResult struct {
	Result workflow.Future
}

func (r *ComprehendDescribeSentimentDetectionJobResult) Get(ctx workflow.Context) (*comprehend.DescribeSentimentDetectionJobOutput, error) {
    var output comprehend.DescribeSentimentDetectionJobOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ComprehendDescribeTopicsDetectionJobResult struct {
	Result workflow.Future
}

func (r *ComprehendDescribeTopicsDetectionJobResult) Get(ctx workflow.Context) (*comprehend.DescribeTopicsDetectionJobOutput, error) {
    var output comprehend.DescribeTopicsDetectionJobOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ComprehendDetectDominantLanguageResult struct {
	Result workflow.Future
}

func (r *ComprehendDetectDominantLanguageResult) Get(ctx workflow.Context) (*comprehend.DetectDominantLanguageOutput, error) {
    var output comprehend.DetectDominantLanguageOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ComprehendDetectEntitiesResult struct {
	Result workflow.Future
}

func (r *ComprehendDetectEntitiesResult) Get(ctx workflow.Context) (*comprehend.DetectEntitiesOutput, error) {
    var output comprehend.DetectEntitiesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ComprehendDetectKeyPhrasesResult struct {
	Result workflow.Future
}

func (r *ComprehendDetectKeyPhrasesResult) Get(ctx workflow.Context) (*comprehend.DetectKeyPhrasesOutput, error) {
    var output comprehend.DetectKeyPhrasesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ComprehendDetectSentimentResult struct {
	Result workflow.Future
}

func (r *ComprehendDetectSentimentResult) Get(ctx workflow.Context) (*comprehend.DetectSentimentOutput, error) {
    var output comprehend.DetectSentimentOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ComprehendDetectSyntaxResult struct {
	Result workflow.Future
}

func (r *ComprehendDetectSyntaxResult) Get(ctx workflow.Context) (*comprehend.DetectSyntaxOutput, error) {
    var output comprehend.DetectSyntaxOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ComprehendListDocumentClassificationJobsResult struct {
	Result workflow.Future
}

func (r *ComprehendListDocumentClassificationJobsResult) Get(ctx workflow.Context) (*comprehend.ListDocumentClassificationJobsOutput, error) {
    var output comprehend.ListDocumentClassificationJobsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ComprehendListDocumentClassifiersResult struct {
	Result workflow.Future
}

func (r *ComprehendListDocumentClassifiersResult) Get(ctx workflow.Context) (*comprehend.ListDocumentClassifiersOutput, error) {
    var output comprehend.ListDocumentClassifiersOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ComprehendListDominantLanguageDetectionJobsResult struct {
	Result workflow.Future
}

func (r *ComprehendListDominantLanguageDetectionJobsResult) Get(ctx workflow.Context) (*comprehend.ListDominantLanguageDetectionJobsOutput, error) {
    var output comprehend.ListDominantLanguageDetectionJobsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ComprehendListEndpointsResult struct {
	Result workflow.Future
}

func (r *ComprehendListEndpointsResult) Get(ctx workflow.Context) (*comprehend.ListEndpointsOutput, error) {
    var output comprehend.ListEndpointsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ComprehendListEntitiesDetectionJobsResult struct {
	Result workflow.Future
}

func (r *ComprehendListEntitiesDetectionJobsResult) Get(ctx workflow.Context) (*comprehend.ListEntitiesDetectionJobsOutput, error) {
    var output comprehend.ListEntitiesDetectionJobsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ComprehendListEntityRecognizersResult struct {
	Result workflow.Future
}

func (r *ComprehendListEntityRecognizersResult) Get(ctx workflow.Context) (*comprehend.ListEntityRecognizersOutput, error) {
    var output comprehend.ListEntityRecognizersOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ComprehendListKeyPhrasesDetectionJobsResult struct {
	Result workflow.Future
}

func (r *ComprehendListKeyPhrasesDetectionJobsResult) Get(ctx workflow.Context) (*comprehend.ListKeyPhrasesDetectionJobsOutput, error) {
    var output comprehend.ListKeyPhrasesDetectionJobsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ComprehendListSentimentDetectionJobsResult struct {
	Result workflow.Future
}

func (r *ComprehendListSentimentDetectionJobsResult) Get(ctx workflow.Context) (*comprehend.ListSentimentDetectionJobsOutput, error) {
    var output comprehend.ListSentimentDetectionJobsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ComprehendListTagsForResourceResult struct {
	Result workflow.Future
}

func (r *ComprehendListTagsForResourceResult) Get(ctx workflow.Context) (*comprehend.ListTagsForResourceOutput, error) {
    var output comprehend.ListTagsForResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ComprehendListTopicsDetectionJobsResult struct {
	Result workflow.Future
}

func (r *ComprehendListTopicsDetectionJobsResult) Get(ctx workflow.Context) (*comprehend.ListTopicsDetectionJobsOutput, error) {
    var output comprehend.ListTopicsDetectionJobsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ComprehendStartDocumentClassificationJobResult struct {
	Result workflow.Future
}

func (r *ComprehendStartDocumentClassificationJobResult) Get(ctx workflow.Context) (*comprehend.StartDocumentClassificationJobOutput, error) {
    var output comprehend.StartDocumentClassificationJobOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ComprehendStartDominantLanguageDetectionJobResult struct {
	Result workflow.Future
}

func (r *ComprehendStartDominantLanguageDetectionJobResult) Get(ctx workflow.Context) (*comprehend.StartDominantLanguageDetectionJobOutput, error) {
    var output comprehend.StartDominantLanguageDetectionJobOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ComprehendStartEntitiesDetectionJobResult struct {
	Result workflow.Future
}

func (r *ComprehendStartEntitiesDetectionJobResult) Get(ctx workflow.Context) (*comprehend.StartEntitiesDetectionJobOutput, error) {
    var output comprehend.StartEntitiesDetectionJobOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ComprehendStartKeyPhrasesDetectionJobResult struct {
	Result workflow.Future
}

func (r *ComprehendStartKeyPhrasesDetectionJobResult) Get(ctx workflow.Context) (*comprehend.StartKeyPhrasesDetectionJobOutput, error) {
    var output comprehend.StartKeyPhrasesDetectionJobOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ComprehendStartSentimentDetectionJobResult struct {
	Result workflow.Future
}

func (r *ComprehendStartSentimentDetectionJobResult) Get(ctx workflow.Context) (*comprehend.StartSentimentDetectionJobOutput, error) {
    var output comprehend.StartSentimentDetectionJobOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ComprehendStartTopicsDetectionJobResult struct {
	Result workflow.Future
}

func (r *ComprehendStartTopicsDetectionJobResult) Get(ctx workflow.Context) (*comprehend.StartTopicsDetectionJobOutput, error) {
    var output comprehend.StartTopicsDetectionJobOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ComprehendStopDominantLanguageDetectionJobResult struct {
	Result workflow.Future
}

func (r *ComprehendStopDominantLanguageDetectionJobResult) Get(ctx workflow.Context) (*comprehend.StopDominantLanguageDetectionJobOutput, error) {
    var output comprehend.StopDominantLanguageDetectionJobOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ComprehendStopEntitiesDetectionJobResult struct {
	Result workflow.Future
}

func (r *ComprehendStopEntitiesDetectionJobResult) Get(ctx workflow.Context) (*comprehend.StopEntitiesDetectionJobOutput, error) {
    var output comprehend.StopEntitiesDetectionJobOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ComprehendStopKeyPhrasesDetectionJobResult struct {
	Result workflow.Future
}

func (r *ComprehendStopKeyPhrasesDetectionJobResult) Get(ctx workflow.Context) (*comprehend.StopKeyPhrasesDetectionJobOutput, error) {
    var output comprehend.StopKeyPhrasesDetectionJobOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ComprehendStopSentimentDetectionJobResult struct {
	Result workflow.Future
}

func (r *ComprehendStopSentimentDetectionJobResult) Get(ctx workflow.Context) (*comprehend.StopSentimentDetectionJobOutput, error) {
    var output comprehend.StopSentimentDetectionJobOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ComprehendStopTrainingDocumentClassifierResult struct {
	Result workflow.Future
}

func (r *ComprehendStopTrainingDocumentClassifierResult) Get(ctx workflow.Context) (*comprehend.StopTrainingDocumentClassifierOutput, error) {
    var output comprehend.StopTrainingDocumentClassifierOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ComprehendStopTrainingEntityRecognizerResult struct {
	Result workflow.Future
}

func (r *ComprehendStopTrainingEntityRecognizerResult) Get(ctx workflow.Context) (*comprehend.StopTrainingEntityRecognizerOutput, error) {
    var output comprehend.StopTrainingEntityRecognizerOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ComprehendTagResourceResult struct {
	Result workflow.Future
}

func (r *ComprehendTagResourceResult) Get(ctx workflow.Context) (*comprehend.TagResourceOutput, error) {
    var output comprehend.TagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ComprehendUntagResourceResult struct {
	Result workflow.Future
}

func (r *ComprehendUntagResourceResult) Get(ctx workflow.Context) (*comprehend.UntagResourceOutput, error) {
    var output comprehend.UntagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ComprehendUpdateEndpointResult struct {
	Result workflow.Future
}

func (r *ComprehendUpdateEndpointResult) Get(ctx workflow.Context) (*comprehend.UpdateEndpointOutput, error) {
    var output comprehend.UpdateEndpointOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ComprehendStub struct {
    activities awsactivities.ComprehendActivities
}

func NewComprehendStub() ComprehendClient {
    return &ComprehendStub{}
}

func (a *ComprehendStub) BatchDetectDominantLanguage(ctx workflow.Context, input *comprehend.BatchDetectDominantLanguageInput) (*comprehend.BatchDetectDominantLanguageOutput, error) {
    var output comprehend.BatchDetectDominantLanguageOutput
    err := workflow.ExecuteActivity(ctx, a.activities.BatchDetectDominantLanguage, input).Get(ctx, &output)
    return &output, err
}

func (a *ComprehendStub) BatchDetectDominantLanguageAsync(ctx workflow.Context, input *comprehend.BatchDetectDominantLanguageInput) *ComprehendBatchDetectDominantLanguageResult {
    future := workflow.ExecuteActivity(ctx, a.activities.BatchDetectDominantLanguage, input)
    return &ComprehendBatchDetectDominantLanguageResult{Result: future}
}

func (a *ComprehendStub) BatchDetectEntities(ctx workflow.Context, input *comprehend.BatchDetectEntitiesInput) (*comprehend.BatchDetectEntitiesOutput, error) {
    var output comprehend.BatchDetectEntitiesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.BatchDetectEntities, input).Get(ctx, &output)
    return &output, err
}

func (a *ComprehendStub) BatchDetectEntitiesAsync(ctx workflow.Context, input *comprehend.BatchDetectEntitiesInput) *ComprehendBatchDetectEntitiesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.BatchDetectEntities, input)
    return &ComprehendBatchDetectEntitiesResult{Result: future}
}

func (a *ComprehendStub) BatchDetectKeyPhrases(ctx workflow.Context, input *comprehend.BatchDetectKeyPhrasesInput) (*comprehend.BatchDetectKeyPhrasesOutput, error) {
    var output comprehend.BatchDetectKeyPhrasesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.BatchDetectKeyPhrases, input).Get(ctx, &output)
    return &output, err
}

func (a *ComprehendStub) BatchDetectKeyPhrasesAsync(ctx workflow.Context, input *comprehend.BatchDetectKeyPhrasesInput) *ComprehendBatchDetectKeyPhrasesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.BatchDetectKeyPhrases, input)
    return &ComprehendBatchDetectKeyPhrasesResult{Result: future}
}

func (a *ComprehendStub) BatchDetectSentiment(ctx workflow.Context, input *comprehend.BatchDetectSentimentInput) (*comprehend.BatchDetectSentimentOutput, error) {
    var output comprehend.BatchDetectSentimentOutput
    err := workflow.ExecuteActivity(ctx, a.activities.BatchDetectSentiment, input).Get(ctx, &output)
    return &output, err
}

func (a *ComprehendStub) BatchDetectSentimentAsync(ctx workflow.Context, input *comprehend.BatchDetectSentimentInput) *ComprehendBatchDetectSentimentResult {
    future := workflow.ExecuteActivity(ctx, a.activities.BatchDetectSentiment, input)
    return &ComprehendBatchDetectSentimentResult{Result: future}
}

func (a *ComprehendStub) BatchDetectSyntax(ctx workflow.Context, input *comprehend.BatchDetectSyntaxInput) (*comprehend.BatchDetectSyntaxOutput, error) {
    var output comprehend.BatchDetectSyntaxOutput
    err := workflow.ExecuteActivity(ctx, a.activities.BatchDetectSyntax, input).Get(ctx, &output)
    return &output, err
}

func (a *ComprehendStub) BatchDetectSyntaxAsync(ctx workflow.Context, input *comprehend.BatchDetectSyntaxInput) *ComprehendBatchDetectSyntaxResult {
    future := workflow.ExecuteActivity(ctx, a.activities.BatchDetectSyntax, input)
    return &ComprehendBatchDetectSyntaxResult{Result: future}
}

func (a *ComprehendStub) ClassifyDocument(ctx workflow.Context, input *comprehend.ClassifyDocumentInput) (*comprehend.ClassifyDocumentOutput, error) {
    var output comprehend.ClassifyDocumentOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ClassifyDocument, input).Get(ctx, &output)
    return &output, err
}

func (a *ComprehendStub) ClassifyDocumentAsync(ctx workflow.Context, input *comprehend.ClassifyDocumentInput) *ComprehendClassifyDocumentResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ClassifyDocument, input)
    return &ComprehendClassifyDocumentResult{Result: future}
}

func (a *ComprehendStub) CreateDocumentClassifier(ctx workflow.Context, input *comprehend.CreateDocumentClassifierInput) (*comprehend.CreateDocumentClassifierOutput, error) {
    var output comprehend.CreateDocumentClassifierOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateDocumentClassifier, input).Get(ctx, &output)
    return &output, err
}

func (a *ComprehendStub) CreateDocumentClassifierAsync(ctx workflow.Context, input *comprehend.CreateDocumentClassifierInput) *ComprehendCreateDocumentClassifierResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateDocumentClassifier, input)
    return &ComprehendCreateDocumentClassifierResult{Result: future}
}

func (a *ComprehendStub) CreateEndpoint(ctx workflow.Context, input *comprehend.CreateEndpointInput) (*comprehend.CreateEndpointOutput, error) {
    var output comprehend.CreateEndpointOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateEndpoint, input).Get(ctx, &output)
    return &output, err
}

func (a *ComprehendStub) CreateEndpointAsync(ctx workflow.Context, input *comprehend.CreateEndpointInput) *ComprehendCreateEndpointResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateEndpoint, input)
    return &ComprehendCreateEndpointResult{Result: future}
}

func (a *ComprehendStub) CreateEntityRecognizer(ctx workflow.Context, input *comprehend.CreateEntityRecognizerInput) (*comprehend.CreateEntityRecognizerOutput, error) {
    var output comprehend.CreateEntityRecognizerOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateEntityRecognizer, input).Get(ctx, &output)
    return &output, err
}

func (a *ComprehendStub) CreateEntityRecognizerAsync(ctx workflow.Context, input *comprehend.CreateEntityRecognizerInput) *ComprehendCreateEntityRecognizerResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateEntityRecognizer, input)
    return &ComprehendCreateEntityRecognizerResult{Result: future}
}

func (a *ComprehendStub) DeleteDocumentClassifier(ctx workflow.Context, input *comprehend.DeleteDocumentClassifierInput) (*comprehend.DeleteDocumentClassifierOutput, error) {
    var output comprehend.DeleteDocumentClassifierOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteDocumentClassifier, input).Get(ctx, &output)
    return &output, err
}

func (a *ComprehendStub) DeleteDocumentClassifierAsync(ctx workflow.Context, input *comprehend.DeleteDocumentClassifierInput) *ComprehendDeleteDocumentClassifierResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteDocumentClassifier, input)
    return &ComprehendDeleteDocumentClassifierResult{Result: future}
}

func (a *ComprehendStub) DeleteEndpoint(ctx workflow.Context, input *comprehend.DeleteEndpointInput) (*comprehend.DeleteEndpointOutput, error) {
    var output comprehend.DeleteEndpointOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteEndpoint, input).Get(ctx, &output)
    return &output, err
}

func (a *ComprehendStub) DeleteEndpointAsync(ctx workflow.Context, input *comprehend.DeleteEndpointInput) *ComprehendDeleteEndpointResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteEndpoint, input)
    return &ComprehendDeleteEndpointResult{Result: future}
}

func (a *ComprehendStub) DeleteEntityRecognizer(ctx workflow.Context, input *comprehend.DeleteEntityRecognizerInput) (*comprehend.DeleteEntityRecognizerOutput, error) {
    var output comprehend.DeleteEntityRecognizerOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteEntityRecognizer, input).Get(ctx, &output)
    return &output, err
}

func (a *ComprehendStub) DeleteEntityRecognizerAsync(ctx workflow.Context, input *comprehend.DeleteEntityRecognizerInput) *ComprehendDeleteEntityRecognizerResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteEntityRecognizer, input)
    return &ComprehendDeleteEntityRecognizerResult{Result: future}
}

func (a *ComprehendStub) DescribeDocumentClassificationJob(ctx workflow.Context, input *comprehend.DescribeDocumentClassificationJobInput) (*comprehend.DescribeDocumentClassificationJobOutput, error) {
    var output comprehend.DescribeDocumentClassificationJobOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeDocumentClassificationJob, input).Get(ctx, &output)
    return &output, err
}

func (a *ComprehendStub) DescribeDocumentClassificationJobAsync(ctx workflow.Context, input *comprehend.DescribeDocumentClassificationJobInput) *ComprehendDescribeDocumentClassificationJobResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeDocumentClassificationJob, input)
    return &ComprehendDescribeDocumentClassificationJobResult{Result: future}
}

func (a *ComprehendStub) DescribeDocumentClassifier(ctx workflow.Context, input *comprehend.DescribeDocumentClassifierInput) (*comprehend.DescribeDocumentClassifierOutput, error) {
    var output comprehend.DescribeDocumentClassifierOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeDocumentClassifier, input).Get(ctx, &output)
    return &output, err
}

func (a *ComprehendStub) DescribeDocumentClassifierAsync(ctx workflow.Context, input *comprehend.DescribeDocumentClassifierInput) *ComprehendDescribeDocumentClassifierResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeDocumentClassifier, input)
    return &ComprehendDescribeDocumentClassifierResult{Result: future}
}

func (a *ComprehendStub) DescribeDominantLanguageDetectionJob(ctx workflow.Context, input *comprehend.DescribeDominantLanguageDetectionJobInput) (*comprehend.DescribeDominantLanguageDetectionJobOutput, error) {
    var output comprehend.DescribeDominantLanguageDetectionJobOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeDominantLanguageDetectionJob, input).Get(ctx, &output)
    return &output, err
}

func (a *ComprehendStub) DescribeDominantLanguageDetectionJobAsync(ctx workflow.Context, input *comprehend.DescribeDominantLanguageDetectionJobInput) *ComprehendDescribeDominantLanguageDetectionJobResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeDominantLanguageDetectionJob, input)
    return &ComprehendDescribeDominantLanguageDetectionJobResult{Result: future}
}

func (a *ComprehendStub) DescribeEndpoint(ctx workflow.Context, input *comprehend.DescribeEndpointInput) (*comprehend.DescribeEndpointOutput, error) {
    var output comprehend.DescribeEndpointOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeEndpoint, input).Get(ctx, &output)
    return &output, err
}

func (a *ComprehendStub) DescribeEndpointAsync(ctx workflow.Context, input *comprehend.DescribeEndpointInput) *ComprehendDescribeEndpointResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeEndpoint, input)
    return &ComprehendDescribeEndpointResult{Result: future}
}

func (a *ComprehendStub) DescribeEntitiesDetectionJob(ctx workflow.Context, input *comprehend.DescribeEntitiesDetectionJobInput) (*comprehend.DescribeEntitiesDetectionJobOutput, error) {
    var output comprehend.DescribeEntitiesDetectionJobOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeEntitiesDetectionJob, input).Get(ctx, &output)
    return &output, err
}

func (a *ComprehendStub) DescribeEntitiesDetectionJobAsync(ctx workflow.Context, input *comprehend.DescribeEntitiesDetectionJobInput) *ComprehendDescribeEntitiesDetectionJobResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeEntitiesDetectionJob, input)
    return &ComprehendDescribeEntitiesDetectionJobResult{Result: future}
}

func (a *ComprehendStub) DescribeEntityRecognizer(ctx workflow.Context, input *comprehend.DescribeEntityRecognizerInput) (*comprehend.DescribeEntityRecognizerOutput, error) {
    var output comprehend.DescribeEntityRecognizerOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeEntityRecognizer, input).Get(ctx, &output)
    return &output, err
}

func (a *ComprehendStub) DescribeEntityRecognizerAsync(ctx workflow.Context, input *comprehend.DescribeEntityRecognizerInput) *ComprehendDescribeEntityRecognizerResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeEntityRecognizer, input)
    return &ComprehendDescribeEntityRecognizerResult{Result: future}
}

func (a *ComprehendStub) DescribeKeyPhrasesDetectionJob(ctx workflow.Context, input *comprehend.DescribeKeyPhrasesDetectionJobInput) (*comprehend.DescribeKeyPhrasesDetectionJobOutput, error) {
    var output comprehend.DescribeKeyPhrasesDetectionJobOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeKeyPhrasesDetectionJob, input).Get(ctx, &output)
    return &output, err
}

func (a *ComprehendStub) DescribeKeyPhrasesDetectionJobAsync(ctx workflow.Context, input *comprehend.DescribeKeyPhrasesDetectionJobInput) *ComprehendDescribeKeyPhrasesDetectionJobResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeKeyPhrasesDetectionJob, input)
    return &ComprehendDescribeKeyPhrasesDetectionJobResult{Result: future}
}

func (a *ComprehendStub) DescribeSentimentDetectionJob(ctx workflow.Context, input *comprehend.DescribeSentimentDetectionJobInput) (*comprehend.DescribeSentimentDetectionJobOutput, error) {
    var output comprehend.DescribeSentimentDetectionJobOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeSentimentDetectionJob, input).Get(ctx, &output)
    return &output, err
}

func (a *ComprehendStub) DescribeSentimentDetectionJobAsync(ctx workflow.Context, input *comprehend.DescribeSentimentDetectionJobInput) *ComprehendDescribeSentimentDetectionJobResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeSentimentDetectionJob, input)
    return &ComprehendDescribeSentimentDetectionJobResult{Result: future}
}

func (a *ComprehendStub) DescribeTopicsDetectionJob(ctx workflow.Context, input *comprehend.DescribeTopicsDetectionJobInput) (*comprehend.DescribeTopicsDetectionJobOutput, error) {
    var output comprehend.DescribeTopicsDetectionJobOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeTopicsDetectionJob, input).Get(ctx, &output)
    return &output, err
}

func (a *ComprehendStub) DescribeTopicsDetectionJobAsync(ctx workflow.Context, input *comprehend.DescribeTopicsDetectionJobInput) *ComprehendDescribeTopicsDetectionJobResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeTopicsDetectionJob, input)
    return &ComprehendDescribeTopicsDetectionJobResult{Result: future}
}

func (a *ComprehendStub) DetectDominantLanguage(ctx workflow.Context, input *comprehend.DetectDominantLanguageInput) (*comprehend.DetectDominantLanguageOutput, error) {
    var output comprehend.DetectDominantLanguageOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DetectDominantLanguage, input).Get(ctx, &output)
    return &output, err
}

func (a *ComprehendStub) DetectDominantLanguageAsync(ctx workflow.Context, input *comprehend.DetectDominantLanguageInput) *ComprehendDetectDominantLanguageResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DetectDominantLanguage, input)
    return &ComprehendDetectDominantLanguageResult{Result: future}
}

func (a *ComprehendStub) DetectEntities(ctx workflow.Context, input *comprehend.DetectEntitiesInput) (*comprehend.DetectEntitiesOutput, error) {
    var output comprehend.DetectEntitiesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DetectEntities, input).Get(ctx, &output)
    return &output, err
}

func (a *ComprehendStub) DetectEntitiesAsync(ctx workflow.Context, input *comprehend.DetectEntitiesInput) *ComprehendDetectEntitiesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DetectEntities, input)
    return &ComprehendDetectEntitiesResult{Result: future}
}

func (a *ComprehendStub) DetectKeyPhrases(ctx workflow.Context, input *comprehend.DetectKeyPhrasesInput) (*comprehend.DetectKeyPhrasesOutput, error) {
    var output comprehend.DetectKeyPhrasesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DetectKeyPhrases, input).Get(ctx, &output)
    return &output, err
}

func (a *ComprehendStub) DetectKeyPhrasesAsync(ctx workflow.Context, input *comprehend.DetectKeyPhrasesInput) *ComprehendDetectKeyPhrasesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DetectKeyPhrases, input)
    return &ComprehendDetectKeyPhrasesResult{Result: future}
}

func (a *ComprehendStub) DetectSentiment(ctx workflow.Context, input *comprehend.DetectSentimentInput) (*comprehend.DetectSentimentOutput, error) {
    var output comprehend.DetectSentimentOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DetectSentiment, input).Get(ctx, &output)
    return &output, err
}

func (a *ComprehendStub) DetectSentimentAsync(ctx workflow.Context, input *comprehend.DetectSentimentInput) *ComprehendDetectSentimentResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DetectSentiment, input)
    return &ComprehendDetectSentimentResult{Result: future}
}

func (a *ComprehendStub) DetectSyntax(ctx workflow.Context, input *comprehend.DetectSyntaxInput) (*comprehend.DetectSyntaxOutput, error) {
    var output comprehend.DetectSyntaxOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DetectSyntax, input).Get(ctx, &output)
    return &output, err
}

func (a *ComprehendStub) DetectSyntaxAsync(ctx workflow.Context, input *comprehend.DetectSyntaxInput) *ComprehendDetectSyntaxResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DetectSyntax, input)
    return &ComprehendDetectSyntaxResult{Result: future}
}

func (a *ComprehendStub) ListDocumentClassificationJobs(ctx workflow.Context, input *comprehend.ListDocumentClassificationJobsInput) (*comprehend.ListDocumentClassificationJobsOutput, error) {
    var output comprehend.ListDocumentClassificationJobsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListDocumentClassificationJobs, input).Get(ctx, &output)
    return &output, err
}

func (a *ComprehendStub) ListDocumentClassificationJobsAsync(ctx workflow.Context, input *comprehend.ListDocumentClassificationJobsInput) *ComprehendListDocumentClassificationJobsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListDocumentClassificationJobs, input)
    return &ComprehendListDocumentClassificationJobsResult{Result: future}
}

func (a *ComprehendStub) ListDocumentClassifiers(ctx workflow.Context, input *comprehend.ListDocumentClassifiersInput) (*comprehend.ListDocumentClassifiersOutput, error) {
    var output comprehend.ListDocumentClassifiersOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListDocumentClassifiers, input).Get(ctx, &output)
    return &output, err
}

func (a *ComprehendStub) ListDocumentClassifiersAsync(ctx workflow.Context, input *comprehend.ListDocumentClassifiersInput) *ComprehendListDocumentClassifiersResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListDocumentClassifiers, input)
    return &ComprehendListDocumentClassifiersResult{Result: future}
}

func (a *ComprehendStub) ListDominantLanguageDetectionJobs(ctx workflow.Context, input *comprehend.ListDominantLanguageDetectionJobsInput) (*comprehend.ListDominantLanguageDetectionJobsOutput, error) {
    var output comprehend.ListDominantLanguageDetectionJobsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListDominantLanguageDetectionJobs, input).Get(ctx, &output)
    return &output, err
}

func (a *ComprehendStub) ListDominantLanguageDetectionJobsAsync(ctx workflow.Context, input *comprehend.ListDominantLanguageDetectionJobsInput) *ComprehendListDominantLanguageDetectionJobsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListDominantLanguageDetectionJobs, input)
    return &ComprehendListDominantLanguageDetectionJobsResult{Result: future}
}

func (a *ComprehendStub) ListEndpoints(ctx workflow.Context, input *comprehend.ListEndpointsInput) (*comprehend.ListEndpointsOutput, error) {
    var output comprehend.ListEndpointsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListEndpoints, input).Get(ctx, &output)
    return &output, err
}

func (a *ComprehendStub) ListEndpointsAsync(ctx workflow.Context, input *comprehend.ListEndpointsInput) *ComprehendListEndpointsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListEndpoints, input)
    return &ComprehendListEndpointsResult{Result: future}
}

func (a *ComprehendStub) ListEntitiesDetectionJobs(ctx workflow.Context, input *comprehend.ListEntitiesDetectionJobsInput) (*comprehend.ListEntitiesDetectionJobsOutput, error) {
    var output comprehend.ListEntitiesDetectionJobsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListEntitiesDetectionJobs, input).Get(ctx, &output)
    return &output, err
}

func (a *ComprehendStub) ListEntitiesDetectionJobsAsync(ctx workflow.Context, input *comprehend.ListEntitiesDetectionJobsInput) *ComprehendListEntitiesDetectionJobsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListEntitiesDetectionJobs, input)
    return &ComprehendListEntitiesDetectionJobsResult{Result: future}
}

func (a *ComprehendStub) ListEntityRecognizers(ctx workflow.Context, input *comprehend.ListEntityRecognizersInput) (*comprehend.ListEntityRecognizersOutput, error) {
    var output comprehend.ListEntityRecognizersOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListEntityRecognizers, input).Get(ctx, &output)
    return &output, err
}

func (a *ComprehendStub) ListEntityRecognizersAsync(ctx workflow.Context, input *comprehend.ListEntityRecognizersInput) *ComprehendListEntityRecognizersResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListEntityRecognizers, input)
    return &ComprehendListEntityRecognizersResult{Result: future}
}

func (a *ComprehendStub) ListKeyPhrasesDetectionJobs(ctx workflow.Context, input *comprehend.ListKeyPhrasesDetectionJobsInput) (*comprehend.ListKeyPhrasesDetectionJobsOutput, error) {
    var output comprehend.ListKeyPhrasesDetectionJobsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListKeyPhrasesDetectionJobs, input).Get(ctx, &output)
    return &output, err
}

func (a *ComprehendStub) ListKeyPhrasesDetectionJobsAsync(ctx workflow.Context, input *comprehend.ListKeyPhrasesDetectionJobsInput) *ComprehendListKeyPhrasesDetectionJobsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListKeyPhrasesDetectionJobs, input)
    return &ComprehendListKeyPhrasesDetectionJobsResult{Result: future}
}

func (a *ComprehendStub) ListSentimentDetectionJobs(ctx workflow.Context, input *comprehend.ListSentimentDetectionJobsInput) (*comprehend.ListSentimentDetectionJobsOutput, error) {
    var output comprehend.ListSentimentDetectionJobsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListSentimentDetectionJobs, input).Get(ctx, &output)
    return &output, err
}

func (a *ComprehendStub) ListSentimentDetectionJobsAsync(ctx workflow.Context, input *comprehend.ListSentimentDetectionJobsInput) *ComprehendListSentimentDetectionJobsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListSentimentDetectionJobs, input)
    return &ComprehendListSentimentDetectionJobsResult{Result: future}
}

func (a *ComprehendStub) ListTagsForResource(ctx workflow.Context, input *comprehend.ListTagsForResourceInput) (*comprehend.ListTagsForResourceOutput, error) {
    var output comprehend.ListTagsForResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListTagsForResource, input).Get(ctx, &output)
    return &output, err
}

func (a *ComprehendStub) ListTagsForResourceAsync(ctx workflow.Context, input *comprehend.ListTagsForResourceInput) *ComprehendListTagsForResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListTagsForResource, input)
    return &ComprehendListTagsForResourceResult{Result: future}
}

func (a *ComprehendStub) ListTopicsDetectionJobs(ctx workflow.Context, input *comprehend.ListTopicsDetectionJobsInput) (*comprehend.ListTopicsDetectionJobsOutput, error) {
    var output comprehend.ListTopicsDetectionJobsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListTopicsDetectionJobs, input).Get(ctx, &output)
    return &output, err
}

func (a *ComprehendStub) ListTopicsDetectionJobsAsync(ctx workflow.Context, input *comprehend.ListTopicsDetectionJobsInput) *ComprehendListTopicsDetectionJobsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListTopicsDetectionJobs, input)
    return &ComprehendListTopicsDetectionJobsResult{Result: future}
}

func (a *ComprehendStub) StartDocumentClassificationJob(ctx workflow.Context, input *comprehend.StartDocumentClassificationJobInput) (*comprehend.StartDocumentClassificationJobOutput, error) {
    var output comprehend.StartDocumentClassificationJobOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StartDocumentClassificationJob, input).Get(ctx, &output)
    return &output, err
}

func (a *ComprehendStub) StartDocumentClassificationJobAsync(ctx workflow.Context, input *comprehend.StartDocumentClassificationJobInput) *ComprehendStartDocumentClassificationJobResult {
    future := workflow.ExecuteActivity(ctx, a.activities.StartDocumentClassificationJob, input)
    return &ComprehendStartDocumentClassificationJobResult{Result: future}
}

func (a *ComprehendStub) StartDominantLanguageDetectionJob(ctx workflow.Context, input *comprehend.StartDominantLanguageDetectionJobInput) (*comprehend.StartDominantLanguageDetectionJobOutput, error) {
    var output comprehend.StartDominantLanguageDetectionJobOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StartDominantLanguageDetectionJob, input).Get(ctx, &output)
    return &output, err
}

func (a *ComprehendStub) StartDominantLanguageDetectionJobAsync(ctx workflow.Context, input *comprehend.StartDominantLanguageDetectionJobInput) *ComprehendStartDominantLanguageDetectionJobResult {
    future := workflow.ExecuteActivity(ctx, a.activities.StartDominantLanguageDetectionJob, input)
    return &ComprehendStartDominantLanguageDetectionJobResult{Result: future}
}

func (a *ComprehendStub) StartEntitiesDetectionJob(ctx workflow.Context, input *comprehend.StartEntitiesDetectionJobInput) (*comprehend.StartEntitiesDetectionJobOutput, error) {
    var output comprehend.StartEntitiesDetectionJobOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StartEntitiesDetectionJob, input).Get(ctx, &output)
    return &output, err
}

func (a *ComprehendStub) StartEntitiesDetectionJobAsync(ctx workflow.Context, input *comprehend.StartEntitiesDetectionJobInput) *ComprehendStartEntitiesDetectionJobResult {
    future := workflow.ExecuteActivity(ctx, a.activities.StartEntitiesDetectionJob, input)
    return &ComprehendStartEntitiesDetectionJobResult{Result: future}
}

func (a *ComprehendStub) StartKeyPhrasesDetectionJob(ctx workflow.Context, input *comprehend.StartKeyPhrasesDetectionJobInput) (*comprehend.StartKeyPhrasesDetectionJobOutput, error) {
    var output comprehend.StartKeyPhrasesDetectionJobOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StartKeyPhrasesDetectionJob, input).Get(ctx, &output)
    return &output, err
}

func (a *ComprehendStub) StartKeyPhrasesDetectionJobAsync(ctx workflow.Context, input *comprehend.StartKeyPhrasesDetectionJobInput) *ComprehendStartKeyPhrasesDetectionJobResult {
    future := workflow.ExecuteActivity(ctx, a.activities.StartKeyPhrasesDetectionJob, input)
    return &ComprehendStartKeyPhrasesDetectionJobResult{Result: future}
}

func (a *ComprehendStub) StartSentimentDetectionJob(ctx workflow.Context, input *comprehend.StartSentimentDetectionJobInput) (*comprehend.StartSentimentDetectionJobOutput, error) {
    var output comprehend.StartSentimentDetectionJobOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StartSentimentDetectionJob, input).Get(ctx, &output)
    return &output, err
}

func (a *ComprehendStub) StartSentimentDetectionJobAsync(ctx workflow.Context, input *comprehend.StartSentimentDetectionJobInput) *ComprehendStartSentimentDetectionJobResult {
    future := workflow.ExecuteActivity(ctx, a.activities.StartSentimentDetectionJob, input)
    return &ComprehendStartSentimentDetectionJobResult{Result: future}
}

func (a *ComprehendStub) StartTopicsDetectionJob(ctx workflow.Context, input *comprehend.StartTopicsDetectionJobInput) (*comprehend.StartTopicsDetectionJobOutput, error) {
    var output comprehend.StartTopicsDetectionJobOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StartTopicsDetectionJob, input).Get(ctx, &output)
    return &output, err
}

func (a *ComprehendStub) StartTopicsDetectionJobAsync(ctx workflow.Context, input *comprehend.StartTopicsDetectionJobInput) *ComprehendStartTopicsDetectionJobResult {
    future := workflow.ExecuteActivity(ctx, a.activities.StartTopicsDetectionJob, input)
    return &ComprehendStartTopicsDetectionJobResult{Result: future}
}

func (a *ComprehendStub) StopDominantLanguageDetectionJob(ctx workflow.Context, input *comprehend.StopDominantLanguageDetectionJobInput) (*comprehend.StopDominantLanguageDetectionJobOutput, error) {
    var output comprehend.StopDominantLanguageDetectionJobOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StopDominantLanguageDetectionJob, input).Get(ctx, &output)
    return &output, err
}

func (a *ComprehendStub) StopDominantLanguageDetectionJobAsync(ctx workflow.Context, input *comprehend.StopDominantLanguageDetectionJobInput) *ComprehendStopDominantLanguageDetectionJobResult {
    future := workflow.ExecuteActivity(ctx, a.activities.StopDominantLanguageDetectionJob, input)
    return &ComprehendStopDominantLanguageDetectionJobResult{Result: future}
}

func (a *ComprehendStub) StopEntitiesDetectionJob(ctx workflow.Context, input *comprehend.StopEntitiesDetectionJobInput) (*comprehend.StopEntitiesDetectionJobOutput, error) {
    var output comprehend.StopEntitiesDetectionJobOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StopEntitiesDetectionJob, input).Get(ctx, &output)
    return &output, err
}

func (a *ComprehendStub) StopEntitiesDetectionJobAsync(ctx workflow.Context, input *comprehend.StopEntitiesDetectionJobInput) *ComprehendStopEntitiesDetectionJobResult {
    future := workflow.ExecuteActivity(ctx, a.activities.StopEntitiesDetectionJob, input)
    return &ComprehendStopEntitiesDetectionJobResult{Result: future}
}

func (a *ComprehendStub) StopKeyPhrasesDetectionJob(ctx workflow.Context, input *comprehend.StopKeyPhrasesDetectionJobInput) (*comprehend.StopKeyPhrasesDetectionJobOutput, error) {
    var output comprehend.StopKeyPhrasesDetectionJobOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StopKeyPhrasesDetectionJob, input).Get(ctx, &output)
    return &output, err
}

func (a *ComprehendStub) StopKeyPhrasesDetectionJobAsync(ctx workflow.Context, input *comprehend.StopKeyPhrasesDetectionJobInput) *ComprehendStopKeyPhrasesDetectionJobResult {
    future := workflow.ExecuteActivity(ctx, a.activities.StopKeyPhrasesDetectionJob, input)
    return &ComprehendStopKeyPhrasesDetectionJobResult{Result: future}
}

func (a *ComprehendStub) StopSentimentDetectionJob(ctx workflow.Context, input *comprehend.StopSentimentDetectionJobInput) (*comprehend.StopSentimentDetectionJobOutput, error) {
    var output comprehend.StopSentimentDetectionJobOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StopSentimentDetectionJob, input).Get(ctx, &output)
    return &output, err
}

func (a *ComprehendStub) StopSentimentDetectionJobAsync(ctx workflow.Context, input *comprehend.StopSentimentDetectionJobInput) *ComprehendStopSentimentDetectionJobResult {
    future := workflow.ExecuteActivity(ctx, a.activities.StopSentimentDetectionJob, input)
    return &ComprehendStopSentimentDetectionJobResult{Result: future}
}

func (a *ComprehendStub) StopTrainingDocumentClassifier(ctx workflow.Context, input *comprehend.StopTrainingDocumentClassifierInput) (*comprehend.StopTrainingDocumentClassifierOutput, error) {
    var output comprehend.StopTrainingDocumentClassifierOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StopTrainingDocumentClassifier, input).Get(ctx, &output)
    return &output, err
}

func (a *ComprehendStub) StopTrainingDocumentClassifierAsync(ctx workflow.Context, input *comprehend.StopTrainingDocumentClassifierInput) *ComprehendStopTrainingDocumentClassifierResult {
    future := workflow.ExecuteActivity(ctx, a.activities.StopTrainingDocumentClassifier, input)
    return &ComprehendStopTrainingDocumentClassifierResult{Result: future}
}

func (a *ComprehendStub) StopTrainingEntityRecognizer(ctx workflow.Context, input *comprehend.StopTrainingEntityRecognizerInput) (*comprehend.StopTrainingEntityRecognizerOutput, error) {
    var output comprehend.StopTrainingEntityRecognizerOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StopTrainingEntityRecognizer, input).Get(ctx, &output)
    return &output, err
}

func (a *ComprehendStub) StopTrainingEntityRecognizerAsync(ctx workflow.Context, input *comprehend.StopTrainingEntityRecognizerInput) *ComprehendStopTrainingEntityRecognizerResult {
    future := workflow.ExecuteActivity(ctx, a.activities.StopTrainingEntityRecognizer, input)
    return &ComprehendStopTrainingEntityRecognizerResult{Result: future}
}

func (a *ComprehendStub) TagResource(ctx workflow.Context, input *comprehend.TagResourceInput) (*comprehend.TagResourceOutput, error) {
    var output comprehend.TagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.TagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *ComprehendStub) TagResourceAsync(ctx workflow.Context, input *comprehend.TagResourceInput) *ComprehendTagResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.TagResource, input)
    return &ComprehendTagResourceResult{Result: future}
}

func (a *ComprehendStub) UntagResource(ctx workflow.Context, input *comprehend.UntagResourceInput) (*comprehend.UntagResourceOutput, error) {
    var output comprehend.UntagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *ComprehendStub) UntagResourceAsync(ctx workflow.Context, input *comprehend.UntagResourceInput) *ComprehendUntagResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input)
    return &ComprehendUntagResourceResult{Result: future}
}

func (a *ComprehendStub) UpdateEndpoint(ctx workflow.Context, input *comprehend.UpdateEndpointInput) (*comprehend.UpdateEndpointOutput, error) {
    var output comprehend.UpdateEndpointOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateEndpoint, input).Get(ctx, &output)
    return &output, err
}

func (a *ComprehendStub) UpdateEndpointAsync(ctx workflow.Context, input *comprehend.UpdateEndpointInput) *ComprehendUpdateEndpointResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateEndpoint, input)
    return &ComprehendUpdateEndpointResult{Result: future}
}