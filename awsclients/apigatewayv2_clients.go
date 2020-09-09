package awsclients

import (
	"github.com/aws/aws-sdk-go/service/apigatewayv2"
	"go.temporal.io/sdk/workflow"
	"temporal.io/aws-sdk/awsactivities"
)

type ApiGatewayV2Client interface {
	CreateApi(ctx workflow.Context, input *apigatewayv2.CreateApiInput) (*apigatewayv2.CreateApiOutput, error)
	CreateApiAsync(ctx workflow.Context, input *apigatewayv2.CreateApiInput) *Apigatewayv2CreateApiResult

	CreateApiMapping(ctx workflow.Context, input *apigatewayv2.CreateApiMappingInput) (*apigatewayv2.CreateApiMappingOutput, error)
	CreateApiMappingAsync(ctx workflow.Context, input *apigatewayv2.CreateApiMappingInput) *Apigatewayv2CreateApiMappingResult

	CreateAuthorizer(ctx workflow.Context, input *apigatewayv2.CreateAuthorizerInput) (*apigatewayv2.CreateAuthorizerOutput, error)
	CreateAuthorizerAsync(ctx workflow.Context, input *apigatewayv2.CreateAuthorizerInput) *Apigatewayv2CreateAuthorizerResult

	CreateDeployment(ctx workflow.Context, input *apigatewayv2.CreateDeploymentInput) (*apigatewayv2.CreateDeploymentOutput, error)
	CreateDeploymentAsync(ctx workflow.Context, input *apigatewayv2.CreateDeploymentInput) *Apigatewayv2CreateDeploymentResult

	CreateDomainName(ctx workflow.Context, input *apigatewayv2.CreateDomainNameInput) (*apigatewayv2.CreateDomainNameOutput, error)
	CreateDomainNameAsync(ctx workflow.Context, input *apigatewayv2.CreateDomainNameInput) *Apigatewayv2CreateDomainNameResult

	CreateIntegration(ctx workflow.Context, input *apigatewayv2.CreateIntegrationInput) (*apigatewayv2.CreateIntegrationOutput, error)
	CreateIntegrationAsync(ctx workflow.Context, input *apigatewayv2.CreateIntegrationInput) *Apigatewayv2CreateIntegrationResult

	CreateIntegrationResponse(ctx workflow.Context, input *apigatewayv2.CreateIntegrationResponseInput) (*apigatewayv2.CreateIntegrationResponseOutput, error)
	CreateIntegrationResponseAsync(ctx workflow.Context, input *apigatewayv2.CreateIntegrationResponseInput) *Apigatewayv2CreateIntegrationResponseResult

	CreateModel(ctx workflow.Context, input *apigatewayv2.CreateModelInput) (*apigatewayv2.CreateModelOutput, error)
	CreateModelAsync(ctx workflow.Context, input *apigatewayv2.CreateModelInput) *Apigatewayv2CreateModelResult

	CreateRoute(ctx workflow.Context, input *apigatewayv2.CreateRouteInput) (*apigatewayv2.CreateRouteOutput, error)
	CreateRouteAsync(ctx workflow.Context, input *apigatewayv2.CreateRouteInput) *Apigatewayv2CreateRouteResult

	CreateRouteResponse(ctx workflow.Context, input *apigatewayv2.CreateRouteResponseInput) (*apigatewayv2.CreateRouteResponseOutput, error)
	CreateRouteResponseAsync(ctx workflow.Context, input *apigatewayv2.CreateRouteResponseInput) *Apigatewayv2CreateRouteResponseResult

	CreateStage(ctx workflow.Context, input *apigatewayv2.CreateStageInput) (*apigatewayv2.CreateStageOutput, error)
	CreateStageAsync(ctx workflow.Context, input *apigatewayv2.CreateStageInput) *Apigatewayv2CreateStageResult

	CreateVpcLink(ctx workflow.Context, input *apigatewayv2.CreateVpcLinkInput) (*apigatewayv2.CreateVpcLinkOutput, error)
	CreateVpcLinkAsync(ctx workflow.Context, input *apigatewayv2.CreateVpcLinkInput) *Apigatewayv2CreateVpcLinkResult

	DeleteAccessLogSettings(ctx workflow.Context, input *apigatewayv2.DeleteAccessLogSettingsInput) (*apigatewayv2.DeleteAccessLogSettingsOutput, error)
	DeleteAccessLogSettingsAsync(ctx workflow.Context, input *apigatewayv2.DeleteAccessLogSettingsInput) *Apigatewayv2DeleteAccessLogSettingsResult

	DeleteApi(ctx workflow.Context, input *apigatewayv2.DeleteApiInput) (*apigatewayv2.DeleteApiOutput, error)
	DeleteApiAsync(ctx workflow.Context, input *apigatewayv2.DeleteApiInput) *Apigatewayv2DeleteApiResult

	DeleteApiMapping(ctx workflow.Context, input *apigatewayv2.DeleteApiMappingInput) (*apigatewayv2.DeleteApiMappingOutput, error)
	DeleteApiMappingAsync(ctx workflow.Context, input *apigatewayv2.DeleteApiMappingInput) *Apigatewayv2DeleteApiMappingResult

	DeleteAuthorizer(ctx workflow.Context, input *apigatewayv2.DeleteAuthorizerInput) (*apigatewayv2.DeleteAuthorizerOutput, error)
	DeleteAuthorizerAsync(ctx workflow.Context, input *apigatewayv2.DeleteAuthorizerInput) *Apigatewayv2DeleteAuthorizerResult

	DeleteCorsConfiguration(ctx workflow.Context, input *apigatewayv2.DeleteCorsConfigurationInput) (*apigatewayv2.DeleteCorsConfigurationOutput, error)
	DeleteCorsConfigurationAsync(ctx workflow.Context, input *apigatewayv2.DeleteCorsConfigurationInput) *Apigatewayv2DeleteCorsConfigurationResult

	DeleteDeployment(ctx workflow.Context, input *apigatewayv2.DeleteDeploymentInput) (*apigatewayv2.DeleteDeploymentOutput, error)
	DeleteDeploymentAsync(ctx workflow.Context, input *apigatewayv2.DeleteDeploymentInput) *Apigatewayv2DeleteDeploymentResult

	DeleteDomainName(ctx workflow.Context, input *apigatewayv2.DeleteDomainNameInput) (*apigatewayv2.DeleteDomainNameOutput, error)
	DeleteDomainNameAsync(ctx workflow.Context, input *apigatewayv2.DeleteDomainNameInput) *Apigatewayv2DeleteDomainNameResult

	DeleteIntegration(ctx workflow.Context, input *apigatewayv2.DeleteIntegrationInput) (*apigatewayv2.DeleteIntegrationOutput, error)
	DeleteIntegrationAsync(ctx workflow.Context, input *apigatewayv2.DeleteIntegrationInput) *Apigatewayv2DeleteIntegrationResult

	DeleteIntegrationResponse(ctx workflow.Context, input *apigatewayv2.DeleteIntegrationResponseInput) (*apigatewayv2.DeleteIntegrationResponseOutput, error)
	DeleteIntegrationResponseAsync(ctx workflow.Context, input *apigatewayv2.DeleteIntegrationResponseInput) *Apigatewayv2DeleteIntegrationResponseResult

	DeleteModel(ctx workflow.Context, input *apigatewayv2.DeleteModelInput) (*apigatewayv2.DeleteModelOutput, error)
	DeleteModelAsync(ctx workflow.Context, input *apigatewayv2.DeleteModelInput) *Apigatewayv2DeleteModelResult

	DeleteRoute(ctx workflow.Context, input *apigatewayv2.DeleteRouteInput) (*apigatewayv2.DeleteRouteOutput, error)
	DeleteRouteAsync(ctx workflow.Context, input *apigatewayv2.DeleteRouteInput) *Apigatewayv2DeleteRouteResult

	DeleteRouteRequestParameter(ctx workflow.Context, input *apigatewayv2.DeleteRouteRequestParameterInput) (*apigatewayv2.DeleteRouteRequestParameterOutput, error)
	DeleteRouteRequestParameterAsync(ctx workflow.Context, input *apigatewayv2.DeleteRouteRequestParameterInput) *Apigatewayv2DeleteRouteRequestParameterResult

	DeleteRouteResponse(ctx workflow.Context, input *apigatewayv2.DeleteRouteResponseInput) (*apigatewayv2.DeleteRouteResponseOutput, error)
	DeleteRouteResponseAsync(ctx workflow.Context, input *apigatewayv2.DeleteRouteResponseInput) *Apigatewayv2DeleteRouteResponseResult

	DeleteRouteSettings(ctx workflow.Context, input *apigatewayv2.DeleteRouteSettingsInput) (*apigatewayv2.DeleteRouteSettingsOutput, error)
	DeleteRouteSettingsAsync(ctx workflow.Context, input *apigatewayv2.DeleteRouteSettingsInput) *Apigatewayv2DeleteRouteSettingsResult

	DeleteStage(ctx workflow.Context, input *apigatewayv2.DeleteStageInput) (*apigatewayv2.DeleteStageOutput, error)
	DeleteStageAsync(ctx workflow.Context, input *apigatewayv2.DeleteStageInput) *Apigatewayv2DeleteStageResult

	DeleteVpcLink(ctx workflow.Context, input *apigatewayv2.DeleteVpcLinkInput) (*apigatewayv2.DeleteVpcLinkOutput, error)
	DeleteVpcLinkAsync(ctx workflow.Context, input *apigatewayv2.DeleteVpcLinkInput) *Apigatewayv2DeleteVpcLinkResult

	ExportApi(ctx workflow.Context, input *apigatewayv2.ExportApiInput) (*apigatewayv2.ExportApiOutput, error)
	ExportApiAsync(ctx workflow.Context, input *apigatewayv2.ExportApiInput) *Apigatewayv2ExportApiResult

	GetApi(ctx workflow.Context, input *apigatewayv2.GetApiInput) (*apigatewayv2.GetApiOutput, error)
	GetApiAsync(ctx workflow.Context, input *apigatewayv2.GetApiInput) *Apigatewayv2GetApiResult

	GetApiMapping(ctx workflow.Context, input *apigatewayv2.GetApiMappingInput) (*apigatewayv2.GetApiMappingOutput, error)
	GetApiMappingAsync(ctx workflow.Context, input *apigatewayv2.GetApiMappingInput) *Apigatewayv2GetApiMappingResult

	GetApiMappings(ctx workflow.Context, input *apigatewayv2.GetApiMappingsInput) (*apigatewayv2.GetApiMappingsOutput, error)
	GetApiMappingsAsync(ctx workflow.Context, input *apigatewayv2.GetApiMappingsInput) *Apigatewayv2GetApiMappingsResult

	GetApis(ctx workflow.Context, input *apigatewayv2.GetApisInput) (*apigatewayv2.GetApisOutput, error)
	GetApisAsync(ctx workflow.Context, input *apigatewayv2.GetApisInput) *Apigatewayv2GetApisResult

	GetAuthorizer(ctx workflow.Context, input *apigatewayv2.GetAuthorizerInput) (*apigatewayv2.GetAuthorizerOutput, error)
	GetAuthorizerAsync(ctx workflow.Context, input *apigatewayv2.GetAuthorizerInput) *Apigatewayv2GetAuthorizerResult

	GetAuthorizers(ctx workflow.Context, input *apigatewayv2.GetAuthorizersInput) (*apigatewayv2.GetAuthorizersOutput, error)
	GetAuthorizersAsync(ctx workflow.Context, input *apigatewayv2.GetAuthorizersInput) *Apigatewayv2GetAuthorizersResult

	GetDeployment(ctx workflow.Context, input *apigatewayv2.GetDeploymentInput) (*apigatewayv2.GetDeploymentOutput, error)
	GetDeploymentAsync(ctx workflow.Context, input *apigatewayv2.GetDeploymentInput) *Apigatewayv2GetDeploymentResult

	GetDeployments(ctx workflow.Context, input *apigatewayv2.GetDeploymentsInput) (*apigatewayv2.GetDeploymentsOutput, error)
	GetDeploymentsAsync(ctx workflow.Context, input *apigatewayv2.GetDeploymentsInput) *Apigatewayv2GetDeploymentsResult

	GetDomainName(ctx workflow.Context, input *apigatewayv2.GetDomainNameInput) (*apigatewayv2.GetDomainNameOutput, error)
	GetDomainNameAsync(ctx workflow.Context, input *apigatewayv2.GetDomainNameInput) *Apigatewayv2GetDomainNameResult

	GetDomainNames(ctx workflow.Context, input *apigatewayv2.GetDomainNamesInput) (*apigatewayv2.GetDomainNamesOutput, error)
	GetDomainNamesAsync(ctx workflow.Context, input *apigatewayv2.GetDomainNamesInput) *Apigatewayv2GetDomainNamesResult

	GetIntegration(ctx workflow.Context, input *apigatewayv2.GetIntegrationInput) (*apigatewayv2.GetIntegrationOutput, error)
	GetIntegrationAsync(ctx workflow.Context, input *apigatewayv2.GetIntegrationInput) *Apigatewayv2GetIntegrationResult

	GetIntegrationResponse(ctx workflow.Context, input *apigatewayv2.GetIntegrationResponseInput) (*apigatewayv2.GetIntegrationResponseOutput, error)
	GetIntegrationResponseAsync(ctx workflow.Context, input *apigatewayv2.GetIntegrationResponseInput) *Apigatewayv2GetIntegrationResponseResult

	GetIntegrationResponses(ctx workflow.Context, input *apigatewayv2.GetIntegrationResponsesInput) (*apigatewayv2.GetIntegrationResponsesOutput, error)
	GetIntegrationResponsesAsync(ctx workflow.Context, input *apigatewayv2.GetIntegrationResponsesInput) *Apigatewayv2GetIntegrationResponsesResult

	GetIntegrations(ctx workflow.Context, input *apigatewayv2.GetIntegrationsInput) (*apigatewayv2.GetIntegrationsOutput, error)
	GetIntegrationsAsync(ctx workflow.Context, input *apigatewayv2.GetIntegrationsInput) *Apigatewayv2GetIntegrationsResult

	GetModel(ctx workflow.Context, input *apigatewayv2.GetModelInput) (*apigatewayv2.GetModelOutput, error)
	GetModelAsync(ctx workflow.Context, input *apigatewayv2.GetModelInput) *Apigatewayv2GetModelResult

	GetModelTemplate(ctx workflow.Context, input *apigatewayv2.GetModelTemplateInput) (*apigatewayv2.GetModelTemplateOutput, error)
	GetModelTemplateAsync(ctx workflow.Context, input *apigatewayv2.GetModelTemplateInput) *Apigatewayv2GetModelTemplateResult

	GetModels(ctx workflow.Context, input *apigatewayv2.GetModelsInput) (*apigatewayv2.GetModelsOutput, error)
	GetModelsAsync(ctx workflow.Context, input *apigatewayv2.GetModelsInput) *Apigatewayv2GetModelsResult

	GetRoute(ctx workflow.Context, input *apigatewayv2.GetRouteInput) (*apigatewayv2.GetRouteOutput, error)
	GetRouteAsync(ctx workflow.Context, input *apigatewayv2.GetRouteInput) *Apigatewayv2GetRouteResult

	GetRouteResponse(ctx workflow.Context, input *apigatewayv2.GetRouteResponseInput) (*apigatewayv2.GetRouteResponseOutput, error)
	GetRouteResponseAsync(ctx workflow.Context, input *apigatewayv2.GetRouteResponseInput) *Apigatewayv2GetRouteResponseResult

	GetRouteResponses(ctx workflow.Context, input *apigatewayv2.GetRouteResponsesInput) (*apigatewayv2.GetRouteResponsesOutput, error)
	GetRouteResponsesAsync(ctx workflow.Context, input *apigatewayv2.GetRouteResponsesInput) *Apigatewayv2GetRouteResponsesResult

	GetRoutes(ctx workflow.Context, input *apigatewayv2.GetRoutesInput) (*apigatewayv2.GetRoutesOutput, error)
	GetRoutesAsync(ctx workflow.Context, input *apigatewayv2.GetRoutesInput) *Apigatewayv2GetRoutesResult

	GetStage(ctx workflow.Context, input *apigatewayv2.GetStageInput) (*apigatewayv2.GetStageOutput, error)
	GetStageAsync(ctx workflow.Context, input *apigatewayv2.GetStageInput) *Apigatewayv2GetStageResult

	GetStages(ctx workflow.Context, input *apigatewayv2.GetStagesInput) (*apigatewayv2.GetStagesOutput, error)
	GetStagesAsync(ctx workflow.Context, input *apigatewayv2.GetStagesInput) *Apigatewayv2GetStagesResult

	GetTags(ctx workflow.Context, input *apigatewayv2.GetTagsInput) (*apigatewayv2.GetTagsOutput, error)
	GetTagsAsync(ctx workflow.Context, input *apigatewayv2.GetTagsInput) *Apigatewayv2GetTagsResult

	GetVpcLink(ctx workflow.Context, input *apigatewayv2.GetVpcLinkInput) (*apigatewayv2.GetVpcLinkOutput, error)
	GetVpcLinkAsync(ctx workflow.Context, input *apigatewayv2.GetVpcLinkInput) *Apigatewayv2GetVpcLinkResult

	GetVpcLinks(ctx workflow.Context, input *apigatewayv2.GetVpcLinksInput) (*apigatewayv2.GetVpcLinksOutput, error)
	GetVpcLinksAsync(ctx workflow.Context, input *apigatewayv2.GetVpcLinksInput) *Apigatewayv2GetVpcLinksResult

	ImportApi(ctx workflow.Context, input *apigatewayv2.ImportApiInput) (*apigatewayv2.ImportApiOutput, error)
	ImportApiAsync(ctx workflow.Context, input *apigatewayv2.ImportApiInput) *Apigatewayv2ImportApiResult

	ReimportApi(ctx workflow.Context, input *apigatewayv2.ReimportApiInput) (*apigatewayv2.ReimportApiOutput, error)
	ReimportApiAsync(ctx workflow.Context, input *apigatewayv2.ReimportApiInput) *Apigatewayv2ReimportApiResult

	TagResource(ctx workflow.Context, input *apigatewayv2.TagResourceInput) (*apigatewayv2.TagResourceOutput, error)
	TagResourceAsync(ctx workflow.Context, input *apigatewayv2.TagResourceInput) *Apigatewayv2TagResourceResult

	UntagResource(ctx workflow.Context, input *apigatewayv2.UntagResourceInput) (*apigatewayv2.UntagResourceOutput, error)
	UntagResourceAsync(ctx workflow.Context, input *apigatewayv2.UntagResourceInput) *Apigatewayv2UntagResourceResult

	UpdateApi(ctx workflow.Context, input *apigatewayv2.UpdateApiInput) (*apigatewayv2.UpdateApiOutput, error)
	UpdateApiAsync(ctx workflow.Context, input *apigatewayv2.UpdateApiInput) *Apigatewayv2UpdateApiResult

	UpdateApiMapping(ctx workflow.Context, input *apigatewayv2.UpdateApiMappingInput) (*apigatewayv2.UpdateApiMappingOutput, error)
	UpdateApiMappingAsync(ctx workflow.Context, input *apigatewayv2.UpdateApiMappingInput) *Apigatewayv2UpdateApiMappingResult

	UpdateAuthorizer(ctx workflow.Context, input *apigatewayv2.UpdateAuthorizerInput) (*apigatewayv2.UpdateAuthorizerOutput, error)
	UpdateAuthorizerAsync(ctx workflow.Context, input *apigatewayv2.UpdateAuthorizerInput) *Apigatewayv2UpdateAuthorizerResult

	UpdateDeployment(ctx workflow.Context, input *apigatewayv2.UpdateDeploymentInput) (*apigatewayv2.UpdateDeploymentOutput, error)
	UpdateDeploymentAsync(ctx workflow.Context, input *apigatewayv2.UpdateDeploymentInput) *Apigatewayv2UpdateDeploymentResult

	UpdateDomainName(ctx workflow.Context, input *apigatewayv2.UpdateDomainNameInput) (*apigatewayv2.UpdateDomainNameOutput, error)
	UpdateDomainNameAsync(ctx workflow.Context, input *apigatewayv2.UpdateDomainNameInput) *Apigatewayv2UpdateDomainNameResult

	UpdateIntegration(ctx workflow.Context, input *apigatewayv2.UpdateIntegrationInput) (*apigatewayv2.UpdateIntegrationOutput, error)
	UpdateIntegrationAsync(ctx workflow.Context, input *apigatewayv2.UpdateIntegrationInput) *Apigatewayv2UpdateIntegrationResult

	UpdateIntegrationResponse(ctx workflow.Context, input *apigatewayv2.UpdateIntegrationResponseInput) (*apigatewayv2.UpdateIntegrationResponseOutput, error)
	UpdateIntegrationResponseAsync(ctx workflow.Context, input *apigatewayv2.UpdateIntegrationResponseInput) *Apigatewayv2UpdateIntegrationResponseResult

	UpdateModel(ctx workflow.Context, input *apigatewayv2.UpdateModelInput) (*apigatewayv2.UpdateModelOutput, error)
	UpdateModelAsync(ctx workflow.Context, input *apigatewayv2.UpdateModelInput) *Apigatewayv2UpdateModelResult

	UpdateRoute(ctx workflow.Context, input *apigatewayv2.UpdateRouteInput) (*apigatewayv2.UpdateRouteOutput, error)
	UpdateRouteAsync(ctx workflow.Context, input *apigatewayv2.UpdateRouteInput) *Apigatewayv2UpdateRouteResult

	UpdateRouteResponse(ctx workflow.Context, input *apigatewayv2.UpdateRouteResponseInput) (*apigatewayv2.UpdateRouteResponseOutput, error)
	UpdateRouteResponseAsync(ctx workflow.Context, input *apigatewayv2.UpdateRouteResponseInput) *Apigatewayv2UpdateRouteResponseResult

	UpdateStage(ctx workflow.Context, input *apigatewayv2.UpdateStageInput) (*apigatewayv2.UpdateStageOutput, error)
	UpdateStageAsync(ctx workflow.Context, input *apigatewayv2.UpdateStageInput) *Apigatewayv2UpdateStageResult

	UpdateVpcLink(ctx workflow.Context, input *apigatewayv2.UpdateVpcLinkInput) (*apigatewayv2.UpdateVpcLinkOutput, error)
	UpdateVpcLinkAsync(ctx workflow.Context, input *apigatewayv2.UpdateVpcLinkInput) *Apigatewayv2UpdateVpcLinkResult
}

type Apigatewayv2CreateApiResult struct {
	Result workflow.Future
}

func (r *Apigatewayv2CreateApiResult) Get(ctx workflow.Context) (*apigatewayv2.CreateApiOutput, error) {
	var output apigatewayv2.CreateApiOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type Apigatewayv2CreateApiMappingResult struct {
	Result workflow.Future
}

func (r *Apigatewayv2CreateApiMappingResult) Get(ctx workflow.Context) (*apigatewayv2.CreateApiMappingOutput, error) {
	var output apigatewayv2.CreateApiMappingOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type Apigatewayv2CreateAuthorizerResult struct {
	Result workflow.Future
}

func (r *Apigatewayv2CreateAuthorizerResult) Get(ctx workflow.Context) (*apigatewayv2.CreateAuthorizerOutput, error) {
	var output apigatewayv2.CreateAuthorizerOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type Apigatewayv2CreateDeploymentResult struct {
	Result workflow.Future
}

func (r *Apigatewayv2CreateDeploymentResult) Get(ctx workflow.Context) (*apigatewayv2.CreateDeploymentOutput, error) {
	var output apigatewayv2.CreateDeploymentOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type Apigatewayv2CreateDomainNameResult struct {
	Result workflow.Future
}

func (r *Apigatewayv2CreateDomainNameResult) Get(ctx workflow.Context) (*apigatewayv2.CreateDomainNameOutput, error) {
	var output apigatewayv2.CreateDomainNameOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type Apigatewayv2CreateIntegrationResult struct {
	Result workflow.Future
}

func (r *Apigatewayv2CreateIntegrationResult) Get(ctx workflow.Context) (*apigatewayv2.CreateIntegrationOutput, error) {
	var output apigatewayv2.CreateIntegrationOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type Apigatewayv2CreateIntegrationResponseResult struct {
	Result workflow.Future
}

func (r *Apigatewayv2CreateIntegrationResponseResult) Get(ctx workflow.Context) (*apigatewayv2.CreateIntegrationResponseOutput, error) {
	var output apigatewayv2.CreateIntegrationResponseOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type Apigatewayv2CreateModelResult struct {
	Result workflow.Future
}

func (r *Apigatewayv2CreateModelResult) Get(ctx workflow.Context) (*apigatewayv2.CreateModelOutput, error) {
	var output apigatewayv2.CreateModelOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type Apigatewayv2CreateRouteResult struct {
	Result workflow.Future
}

func (r *Apigatewayv2CreateRouteResult) Get(ctx workflow.Context) (*apigatewayv2.CreateRouteOutput, error) {
	var output apigatewayv2.CreateRouteOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type Apigatewayv2CreateRouteResponseResult struct {
	Result workflow.Future
}

func (r *Apigatewayv2CreateRouteResponseResult) Get(ctx workflow.Context) (*apigatewayv2.CreateRouteResponseOutput, error) {
	var output apigatewayv2.CreateRouteResponseOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type Apigatewayv2CreateStageResult struct {
	Result workflow.Future
}

func (r *Apigatewayv2CreateStageResult) Get(ctx workflow.Context) (*apigatewayv2.CreateStageOutput, error) {
	var output apigatewayv2.CreateStageOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type Apigatewayv2CreateVpcLinkResult struct {
	Result workflow.Future
}

func (r *Apigatewayv2CreateVpcLinkResult) Get(ctx workflow.Context) (*apigatewayv2.CreateVpcLinkOutput, error) {
	var output apigatewayv2.CreateVpcLinkOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type Apigatewayv2DeleteAccessLogSettingsResult struct {
	Result workflow.Future
}

func (r *Apigatewayv2DeleteAccessLogSettingsResult) Get(ctx workflow.Context) (*apigatewayv2.DeleteAccessLogSettingsOutput, error) {
	var output apigatewayv2.DeleteAccessLogSettingsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type Apigatewayv2DeleteApiResult struct {
	Result workflow.Future
}

func (r *Apigatewayv2DeleteApiResult) Get(ctx workflow.Context) (*apigatewayv2.DeleteApiOutput, error) {
	var output apigatewayv2.DeleteApiOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type Apigatewayv2DeleteApiMappingResult struct {
	Result workflow.Future
}

func (r *Apigatewayv2DeleteApiMappingResult) Get(ctx workflow.Context) (*apigatewayv2.DeleteApiMappingOutput, error) {
	var output apigatewayv2.DeleteApiMappingOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type Apigatewayv2DeleteAuthorizerResult struct {
	Result workflow.Future
}

func (r *Apigatewayv2DeleteAuthorizerResult) Get(ctx workflow.Context) (*apigatewayv2.DeleteAuthorizerOutput, error) {
	var output apigatewayv2.DeleteAuthorizerOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type Apigatewayv2DeleteCorsConfigurationResult struct {
	Result workflow.Future
}

func (r *Apigatewayv2DeleteCorsConfigurationResult) Get(ctx workflow.Context) (*apigatewayv2.DeleteCorsConfigurationOutput, error) {
	var output apigatewayv2.DeleteCorsConfigurationOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type Apigatewayv2DeleteDeploymentResult struct {
	Result workflow.Future
}

func (r *Apigatewayv2DeleteDeploymentResult) Get(ctx workflow.Context) (*apigatewayv2.DeleteDeploymentOutput, error) {
	var output apigatewayv2.DeleteDeploymentOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type Apigatewayv2DeleteDomainNameResult struct {
	Result workflow.Future
}

func (r *Apigatewayv2DeleteDomainNameResult) Get(ctx workflow.Context) (*apigatewayv2.DeleteDomainNameOutput, error) {
	var output apigatewayv2.DeleteDomainNameOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type Apigatewayv2DeleteIntegrationResult struct {
	Result workflow.Future
}

func (r *Apigatewayv2DeleteIntegrationResult) Get(ctx workflow.Context) (*apigatewayv2.DeleteIntegrationOutput, error) {
	var output apigatewayv2.DeleteIntegrationOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type Apigatewayv2DeleteIntegrationResponseResult struct {
	Result workflow.Future
}

func (r *Apigatewayv2DeleteIntegrationResponseResult) Get(ctx workflow.Context) (*apigatewayv2.DeleteIntegrationResponseOutput, error) {
	var output apigatewayv2.DeleteIntegrationResponseOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type Apigatewayv2DeleteModelResult struct {
	Result workflow.Future
}

func (r *Apigatewayv2DeleteModelResult) Get(ctx workflow.Context) (*apigatewayv2.DeleteModelOutput, error) {
	var output apigatewayv2.DeleteModelOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type Apigatewayv2DeleteRouteResult struct {
	Result workflow.Future
}

func (r *Apigatewayv2DeleteRouteResult) Get(ctx workflow.Context) (*apigatewayv2.DeleteRouteOutput, error) {
	var output apigatewayv2.DeleteRouteOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type Apigatewayv2DeleteRouteRequestParameterResult struct {
	Result workflow.Future
}

func (r *Apigatewayv2DeleteRouteRequestParameterResult) Get(ctx workflow.Context) (*apigatewayv2.DeleteRouteRequestParameterOutput, error) {
	var output apigatewayv2.DeleteRouteRequestParameterOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type Apigatewayv2DeleteRouteResponseResult struct {
	Result workflow.Future
}

func (r *Apigatewayv2DeleteRouteResponseResult) Get(ctx workflow.Context) (*apigatewayv2.DeleteRouteResponseOutput, error) {
	var output apigatewayv2.DeleteRouteResponseOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type Apigatewayv2DeleteRouteSettingsResult struct {
	Result workflow.Future
}

func (r *Apigatewayv2DeleteRouteSettingsResult) Get(ctx workflow.Context) (*apigatewayv2.DeleteRouteSettingsOutput, error) {
	var output apigatewayv2.DeleteRouteSettingsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type Apigatewayv2DeleteStageResult struct {
	Result workflow.Future
}

func (r *Apigatewayv2DeleteStageResult) Get(ctx workflow.Context) (*apigatewayv2.DeleteStageOutput, error) {
	var output apigatewayv2.DeleteStageOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type Apigatewayv2DeleteVpcLinkResult struct {
	Result workflow.Future
}

func (r *Apigatewayv2DeleteVpcLinkResult) Get(ctx workflow.Context) (*apigatewayv2.DeleteVpcLinkOutput, error) {
	var output apigatewayv2.DeleteVpcLinkOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type Apigatewayv2ExportApiResult struct {
	Result workflow.Future
}

func (r *Apigatewayv2ExportApiResult) Get(ctx workflow.Context) (*apigatewayv2.ExportApiOutput, error) {
	var output apigatewayv2.ExportApiOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type Apigatewayv2GetApiResult struct {
	Result workflow.Future
}

func (r *Apigatewayv2GetApiResult) Get(ctx workflow.Context) (*apigatewayv2.GetApiOutput, error) {
	var output apigatewayv2.GetApiOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type Apigatewayv2GetApiMappingResult struct {
	Result workflow.Future
}

func (r *Apigatewayv2GetApiMappingResult) Get(ctx workflow.Context) (*apigatewayv2.GetApiMappingOutput, error) {
	var output apigatewayv2.GetApiMappingOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type Apigatewayv2GetApiMappingsResult struct {
	Result workflow.Future
}

func (r *Apigatewayv2GetApiMappingsResult) Get(ctx workflow.Context) (*apigatewayv2.GetApiMappingsOutput, error) {
	var output apigatewayv2.GetApiMappingsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type Apigatewayv2GetApisResult struct {
	Result workflow.Future
}

func (r *Apigatewayv2GetApisResult) Get(ctx workflow.Context) (*apigatewayv2.GetApisOutput, error) {
	var output apigatewayv2.GetApisOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type Apigatewayv2GetAuthorizerResult struct {
	Result workflow.Future
}

func (r *Apigatewayv2GetAuthorizerResult) Get(ctx workflow.Context) (*apigatewayv2.GetAuthorizerOutput, error) {
	var output apigatewayv2.GetAuthorizerOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type Apigatewayv2GetAuthorizersResult struct {
	Result workflow.Future
}

func (r *Apigatewayv2GetAuthorizersResult) Get(ctx workflow.Context) (*apigatewayv2.GetAuthorizersOutput, error) {
	var output apigatewayv2.GetAuthorizersOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type Apigatewayv2GetDeploymentResult struct {
	Result workflow.Future
}

func (r *Apigatewayv2GetDeploymentResult) Get(ctx workflow.Context) (*apigatewayv2.GetDeploymentOutput, error) {
	var output apigatewayv2.GetDeploymentOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type Apigatewayv2GetDeploymentsResult struct {
	Result workflow.Future
}

func (r *Apigatewayv2GetDeploymentsResult) Get(ctx workflow.Context) (*apigatewayv2.GetDeploymentsOutput, error) {
	var output apigatewayv2.GetDeploymentsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type Apigatewayv2GetDomainNameResult struct {
	Result workflow.Future
}

func (r *Apigatewayv2GetDomainNameResult) Get(ctx workflow.Context) (*apigatewayv2.GetDomainNameOutput, error) {
	var output apigatewayv2.GetDomainNameOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type Apigatewayv2GetDomainNamesResult struct {
	Result workflow.Future
}

func (r *Apigatewayv2GetDomainNamesResult) Get(ctx workflow.Context) (*apigatewayv2.GetDomainNamesOutput, error) {
	var output apigatewayv2.GetDomainNamesOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type Apigatewayv2GetIntegrationResult struct {
	Result workflow.Future
}

func (r *Apigatewayv2GetIntegrationResult) Get(ctx workflow.Context) (*apigatewayv2.GetIntegrationOutput, error) {
	var output apigatewayv2.GetIntegrationOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type Apigatewayv2GetIntegrationResponseResult struct {
	Result workflow.Future
}

func (r *Apigatewayv2GetIntegrationResponseResult) Get(ctx workflow.Context) (*apigatewayv2.GetIntegrationResponseOutput, error) {
	var output apigatewayv2.GetIntegrationResponseOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type Apigatewayv2GetIntegrationResponsesResult struct {
	Result workflow.Future
}

func (r *Apigatewayv2GetIntegrationResponsesResult) Get(ctx workflow.Context) (*apigatewayv2.GetIntegrationResponsesOutput, error) {
	var output apigatewayv2.GetIntegrationResponsesOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type Apigatewayv2GetIntegrationsResult struct {
	Result workflow.Future
}

func (r *Apigatewayv2GetIntegrationsResult) Get(ctx workflow.Context) (*apigatewayv2.GetIntegrationsOutput, error) {
	var output apigatewayv2.GetIntegrationsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type Apigatewayv2GetModelResult struct {
	Result workflow.Future
}

func (r *Apigatewayv2GetModelResult) Get(ctx workflow.Context) (*apigatewayv2.GetModelOutput, error) {
	var output apigatewayv2.GetModelOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type Apigatewayv2GetModelTemplateResult struct {
	Result workflow.Future
}

func (r *Apigatewayv2GetModelTemplateResult) Get(ctx workflow.Context) (*apigatewayv2.GetModelTemplateOutput, error) {
	var output apigatewayv2.GetModelTemplateOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type Apigatewayv2GetModelsResult struct {
	Result workflow.Future
}

func (r *Apigatewayv2GetModelsResult) Get(ctx workflow.Context) (*apigatewayv2.GetModelsOutput, error) {
	var output apigatewayv2.GetModelsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type Apigatewayv2GetRouteResult struct {
	Result workflow.Future
}

func (r *Apigatewayv2GetRouteResult) Get(ctx workflow.Context) (*apigatewayv2.GetRouteOutput, error) {
	var output apigatewayv2.GetRouteOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type Apigatewayv2GetRouteResponseResult struct {
	Result workflow.Future
}

func (r *Apigatewayv2GetRouteResponseResult) Get(ctx workflow.Context) (*apigatewayv2.GetRouteResponseOutput, error) {
	var output apigatewayv2.GetRouteResponseOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type Apigatewayv2GetRouteResponsesResult struct {
	Result workflow.Future
}

func (r *Apigatewayv2GetRouteResponsesResult) Get(ctx workflow.Context) (*apigatewayv2.GetRouteResponsesOutput, error) {
	var output apigatewayv2.GetRouteResponsesOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type Apigatewayv2GetRoutesResult struct {
	Result workflow.Future
}

func (r *Apigatewayv2GetRoutesResult) Get(ctx workflow.Context) (*apigatewayv2.GetRoutesOutput, error) {
	var output apigatewayv2.GetRoutesOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type Apigatewayv2GetStageResult struct {
	Result workflow.Future
}

func (r *Apigatewayv2GetStageResult) Get(ctx workflow.Context) (*apigatewayv2.GetStageOutput, error) {
	var output apigatewayv2.GetStageOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type Apigatewayv2GetStagesResult struct {
	Result workflow.Future
}

func (r *Apigatewayv2GetStagesResult) Get(ctx workflow.Context) (*apigatewayv2.GetStagesOutput, error) {
	var output apigatewayv2.GetStagesOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type Apigatewayv2GetTagsResult struct {
	Result workflow.Future
}

func (r *Apigatewayv2GetTagsResult) Get(ctx workflow.Context) (*apigatewayv2.GetTagsOutput, error) {
	var output apigatewayv2.GetTagsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type Apigatewayv2GetVpcLinkResult struct {
	Result workflow.Future
}

func (r *Apigatewayv2GetVpcLinkResult) Get(ctx workflow.Context) (*apigatewayv2.GetVpcLinkOutput, error) {
	var output apigatewayv2.GetVpcLinkOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type Apigatewayv2GetVpcLinksResult struct {
	Result workflow.Future
}

func (r *Apigatewayv2GetVpcLinksResult) Get(ctx workflow.Context) (*apigatewayv2.GetVpcLinksOutput, error) {
	var output apigatewayv2.GetVpcLinksOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type Apigatewayv2ImportApiResult struct {
	Result workflow.Future
}

func (r *Apigatewayv2ImportApiResult) Get(ctx workflow.Context) (*apigatewayv2.ImportApiOutput, error) {
	var output apigatewayv2.ImportApiOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type Apigatewayv2ReimportApiResult struct {
	Result workflow.Future
}

func (r *Apigatewayv2ReimportApiResult) Get(ctx workflow.Context) (*apigatewayv2.ReimportApiOutput, error) {
	var output apigatewayv2.ReimportApiOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type Apigatewayv2TagResourceResult struct {
	Result workflow.Future
}

func (r *Apigatewayv2TagResourceResult) Get(ctx workflow.Context) (*apigatewayv2.TagResourceOutput, error) {
	var output apigatewayv2.TagResourceOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type Apigatewayv2UntagResourceResult struct {
	Result workflow.Future
}

func (r *Apigatewayv2UntagResourceResult) Get(ctx workflow.Context) (*apigatewayv2.UntagResourceOutput, error) {
	var output apigatewayv2.UntagResourceOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type Apigatewayv2UpdateApiResult struct {
	Result workflow.Future
}

func (r *Apigatewayv2UpdateApiResult) Get(ctx workflow.Context) (*apigatewayv2.UpdateApiOutput, error) {
	var output apigatewayv2.UpdateApiOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type Apigatewayv2UpdateApiMappingResult struct {
	Result workflow.Future
}

func (r *Apigatewayv2UpdateApiMappingResult) Get(ctx workflow.Context) (*apigatewayv2.UpdateApiMappingOutput, error) {
	var output apigatewayv2.UpdateApiMappingOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type Apigatewayv2UpdateAuthorizerResult struct {
	Result workflow.Future
}

func (r *Apigatewayv2UpdateAuthorizerResult) Get(ctx workflow.Context) (*apigatewayv2.UpdateAuthorizerOutput, error) {
	var output apigatewayv2.UpdateAuthorizerOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type Apigatewayv2UpdateDeploymentResult struct {
	Result workflow.Future
}

func (r *Apigatewayv2UpdateDeploymentResult) Get(ctx workflow.Context) (*apigatewayv2.UpdateDeploymentOutput, error) {
	var output apigatewayv2.UpdateDeploymentOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type Apigatewayv2UpdateDomainNameResult struct {
	Result workflow.Future
}

func (r *Apigatewayv2UpdateDomainNameResult) Get(ctx workflow.Context) (*apigatewayv2.UpdateDomainNameOutput, error) {
	var output apigatewayv2.UpdateDomainNameOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type Apigatewayv2UpdateIntegrationResult struct {
	Result workflow.Future
}

func (r *Apigatewayv2UpdateIntegrationResult) Get(ctx workflow.Context) (*apigatewayv2.UpdateIntegrationOutput, error) {
	var output apigatewayv2.UpdateIntegrationOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type Apigatewayv2UpdateIntegrationResponseResult struct {
	Result workflow.Future
}

func (r *Apigatewayv2UpdateIntegrationResponseResult) Get(ctx workflow.Context) (*apigatewayv2.UpdateIntegrationResponseOutput, error) {
	var output apigatewayv2.UpdateIntegrationResponseOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type Apigatewayv2UpdateModelResult struct {
	Result workflow.Future
}

func (r *Apigatewayv2UpdateModelResult) Get(ctx workflow.Context) (*apigatewayv2.UpdateModelOutput, error) {
	var output apigatewayv2.UpdateModelOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type Apigatewayv2UpdateRouteResult struct {
	Result workflow.Future
}

func (r *Apigatewayv2UpdateRouteResult) Get(ctx workflow.Context) (*apigatewayv2.UpdateRouteOutput, error) {
	var output apigatewayv2.UpdateRouteOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type Apigatewayv2UpdateRouteResponseResult struct {
	Result workflow.Future
}

func (r *Apigatewayv2UpdateRouteResponseResult) Get(ctx workflow.Context) (*apigatewayv2.UpdateRouteResponseOutput, error) {
	var output apigatewayv2.UpdateRouteResponseOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type Apigatewayv2UpdateStageResult struct {
	Result workflow.Future
}

func (r *Apigatewayv2UpdateStageResult) Get(ctx workflow.Context) (*apigatewayv2.UpdateStageOutput, error) {
	var output apigatewayv2.UpdateStageOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type Apigatewayv2UpdateVpcLinkResult struct {
	Result workflow.Future
}

func (r *Apigatewayv2UpdateVpcLinkResult) Get(ctx workflow.Context) (*apigatewayv2.UpdateVpcLinkOutput, error) {
	var output apigatewayv2.UpdateVpcLinkOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ApiGatewayV2Stub struct {
	activities awsactivities.ApiGatewayV2Activities
}

func NewApiGatewayV2Stub() ApiGatewayV2Client {
	return &ApiGatewayV2Stub{}
}

func (a *ApiGatewayV2Stub) CreateApi(ctx workflow.Context, input *apigatewayv2.CreateApiInput) (*apigatewayv2.CreateApiOutput, error) {
	var output apigatewayv2.CreateApiOutput
	err := workflow.ExecuteActivity(ctx, a.activities.CreateApi, input).Get(ctx, &output)
	return &output, err
}

func (a *ApiGatewayV2Stub) CreateApiAsync(ctx workflow.Context, input *apigatewayv2.CreateApiInput) *Apigatewayv2CreateApiResult {
	future := workflow.ExecuteActivity(ctx, a.activities.CreateApi, input)
	return &Apigatewayv2CreateApiResult{Result: future}
}

func (a *ApiGatewayV2Stub) CreateApiMapping(ctx workflow.Context, input *apigatewayv2.CreateApiMappingInput) (*apigatewayv2.CreateApiMappingOutput, error) {
	var output apigatewayv2.CreateApiMappingOutput
	err := workflow.ExecuteActivity(ctx, a.activities.CreateApiMapping, input).Get(ctx, &output)
	return &output, err
}

func (a *ApiGatewayV2Stub) CreateApiMappingAsync(ctx workflow.Context, input *apigatewayv2.CreateApiMappingInput) *Apigatewayv2CreateApiMappingResult {
	future := workflow.ExecuteActivity(ctx, a.activities.CreateApiMapping, input)
	return &Apigatewayv2CreateApiMappingResult{Result: future}
}

func (a *ApiGatewayV2Stub) CreateAuthorizer(ctx workflow.Context, input *apigatewayv2.CreateAuthorizerInput) (*apigatewayv2.CreateAuthorizerOutput, error) {
	var output apigatewayv2.CreateAuthorizerOutput
	err := workflow.ExecuteActivity(ctx, a.activities.CreateAuthorizer, input).Get(ctx, &output)
	return &output, err
}

func (a *ApiGatewayV2Stub) CreateAuthorizerAsync(ctx workflow.Context, input *apigatewayv2.CreateAuthorizerInput) *Apigatewayv2CreateAuthorizerResult {
	future := workflow.ExecuteActivity(ctx, a.activities.CreateAuthorizer, input)
	return &Apigatewayv2CreateAuthorizerResult{Result: future}
}

func (a *ApiGatewayV2Stub) CreateDeployment(ctx workflow.Context, input *apigatewayv2.CreateDeploymentInput) (*apigatewayv2.CreateDeploymentOutput, error) {
	var output apigatewayv2.CreateDeploymentOutput
	err := workflow.ExecuteActivity(ctx, a.activities.CreateDeployment, input).Get(ctx, &output)
	return &output, err
}

func (a *ApiGatewayV2Stub) CreateDeploymentAsync(ctx workflow.Context, input *apigatewayv2.CreateDeploymentInput) *Apigatewayv2CreateDeploymentResult {
	future := workflow.ExecuteActivity(ctx, a.activities.CreateDeployment, input)
	return &Apigatewayv2CreateDeploymentResult{Result: future}
}

func (a *ApiGatewayV2Stub) CreateDomainName(ctx workflow.Context, input *apigatewayv2.CreateDomainNameInput) (*apigatewayv2.CreateDomainNameOutput, error) {
	var output apigatewayv2.CreateDomainNameOutput
	err := workflow.ExecuteActivity(ctx, a.activities.CreateDomainName, input).Get(ctx, &output)
	return &output, err
}

func (a *ApiGatewayV2Stub) CreateDomainNameAsync(ctx workflow.Context, input *apigatewayv2.CreateDomainNameInput) *Apigatewayv2CreateDomainNameResult {
	future := workflow.ExecuteActivity(ctx, a.activities.CreateDomainName, input)
	return &Apigatewayv2CreateDomainNameResult{Result: future}
}

func (a *ApiGatewayV2Stub) CreateIntegration(ctx workflow.Context, input *apigatewayv2.CreateIntegrationInput) (*apigatewayv2.CreateIntegrationOutput, error) {
	var output apigatewayv2.CreateIntegrationOutput
	err := workflow.ExecuteActivity(ctx, a.activities.CreateIntegration, input).Get(ctx, &output)
	return &output, err
}

func (a *ApiGatewayV2Stub) CreateIntegrationAsync(ctx workflow.Context, input *apigatewayv2.CreateIntegrationInput) *Apigatewayv2CreateIntegrationResult {
	future := workflow.ExecuteActivity(ctx, a.activities.CreateIntegration, input)
	return &Apigatewayv2CreateIntegrationResult{Result: future}
}

func (a *ApiGatewayV2Stub) CreateIntegrationResponse(ctx workflow.Context, input *apigatewayv2.CreateIntegrationResponseInput) (*apigatewayv2.CreateIntegrationResponseOutput, error) {
	var output apigatewayv2.CreateIntegrationResponseOutput
	err := workflow.ExecuteActivity(ctx, a.activities.CreateIntegrationResponse, input).Get(ctx, &output)
	return &output, err
}

func (a *ApiGatewayV2Stub) CreateIntegrationResponseAsync(ctx workflow.Context, input *apigatewayv2.CreateIntegrationResponseInput) *Apigatewayv2CreateIntegrationResponseResult {
	future := workflow.ExecuteActivity(ctx, a.activities.CreateIntegrationResponse, input)
	return &Apigatewayv2CreateIntegrationResponseResult{Result: future}
}

func (a *ApiGatewayV2Stub) CreateModel(ctx workflow.Context, input *apigatewayv2.CreateModelInput) (*apigatewayv2.CreateModelOutput, error) {
	var output apigatewayv2.CreateModelOutput
	err := workflow.ExecuteActivity(ctx, a.activities.CreateModel, input).Get(ctx, &output)
	return &output, err
}

func (a *ApiGatewayV2Stub) CreateModelAsync(ctx workflow.Context, input *apigatewayv2.CreateModelInput) *Apigatewayv2CreateModelResult {
	future := workflow.ExecuteActivity(ctx, a.activities.CreateModel, input)
	return &Apigatewayv2CreateModelResult{Result: future}
}

func (a *ApiGatewayV2Stub) CreateRoute(ctx workflow.Context, input *apigatewayv2.CreateRouteInput) (*apigatewayv2.CreateRouteOutput, error) {
	var output apigatewayv2.CreateRouteOutput
	err := workflow.ExecuteActivity(ctx, a.activities.CreateRoute, input).Get(ctx, &output)
	return &output, err
}

func (a *ApiGatewayV2Stub) CreateRouteAsync(ctx workflow.Context, input *apigatewayv2.CreateRouteInput) *Apigatewayv2CreateRouteResult {
	future := workflow.ExecuteActivity(ctx, a.activities.CreateRoute, input)
	return &Apigatewayv2CreateRouteResult{Result: future}
}

func (a *ApiGatewayV2Stub) CreateRouteResponse(ctx workflow.Context, input *apigatewayv2.CreateRouteResponseInput) (*apigatewayv2.CreateRouteResponseOutput, error) {
	var output apigatewayv2.CreateRouteResponseOutput
	err := workflow.ExecuteActivity(ctx, a.activities.CreateRouteResponse, input).Get(ctx, &output)
	return &output, err
}

func (a *ApiGatewayV2Stub) CreateRouteResponseAsync(ctx workflow.Context, input *apigatewayv2.CreateRouteResponseInput) *Apigatewayv2CreateRouteResponseResult {
	future := workflow.ExecuteActivity(ctx, a.activities.CreateRouteResponse, input)
	return &Apigatewayv2CreateRouteResponseResult{Result: future}
}

func (a *ApiGatewayV2Stub) CreateStage(ctx workflow.Context, input *apigatewayv2.CreateStageInput) (*apigatewayv2.CreateStageOutput, error) {
	var output apigatewayv2.CreateStageOutput
	err := workflow.ExecuteActivity(ctx, a.activities.CreateStage, input).Get(ctx, &output)
	return &output, err
}

func (a *ApiGatewayV2Stub) CreateStageAsync(ctx workflow.Context, input *apigatewayv2.CreateStageInput) *Apigatewayv2CreateStageResult {
	future := workflow.ExecuteActivity(ctx, a.activities.CreateStage, input)
	return &Apigatewayv2CreateStageResult{Result: future}
}

func (a *ApiGatewayV2Stub) CreateVpcLink(ctx workflow.Context, input *apigatewayv2.CreateVpcLinkInput) (*apigatewayv2.CreateVpcLinkOutput, error) {
	var output apigatewayv2.CreateVpcLinkOutput
	err := workflow.ExecuteActivity(ctx, a.activities.CreateVpcLink, input).Get(ctx, &output)
	return &output, err
}

func (a *ApiGatewayV2Stub) CreateVpcLinkAsync(ctx workflow.Context, input *apigatewayv2.CreateVpcLinkInput) *Apigatewayv2CreateVpcLinkResult {
	future := workflow.ExecuteActivity(ctx, a.activities.CreateVpcLink, input)
	return &Apigatewayv2CreateVpcLinkResult{Result: future}
}

func (a *ApiGatewayV2Stub) DeleteAccessLogSettings(ctx workflow.Context, input *apigatewayv2.DeleteAccessLogSettingsInput) (*apigatewayv2.DeleteAccessLogSettingsOutput, error) {
	var output apigatewayv2.DeleteAccessLogSettingsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeleteAccessLogSettings, input).Get(ctx, &output)
	return &output, err
}

func (a *ApiGatewayV2Stub) DeleteAccessLogSettingsAsync(ctx workflow.Context, input *apigatewayv2.DeleteAccessLogSettingsInput) *Apigatewayv2DeleteAccessLogSettingsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeleteAccessLogSettings, input)
	return &Apigatewayv2DeleteAccessLogSettingsResult{Result: future}
}

func (a *ApiGatewayV2Stub) DeleteApi(ctx workflow.Context, input *apigatewayv2.DeleteApiInput) (*apigatewayv2.DeleteApiOutput, error) {
	var output apigatewayv2.DeleteApiOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeleteApi, input).Get(ctx, &output)
	return &output, err
}

func (a *ApiGatewayV2Stub) DeleteApiAsync(ctx workflow.Context, input *apigatewayv2.DeleteApiInput) *Apigatewayv2DeleteApiResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeleteApi, input)
	return &Apigatewayv2DeleteApiResult{Result: future}
}

func (a *ApiGatewayV2Stub) DeleteApiMapping(ctx workflow.Context, input *apigatewayv2.DeleteApiMappingInput) (*apigatewayv2.DeleteApiMappingOutput, error) {
	var output apigatewayv2.DeleteApiMappingOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeleteApiMapping, input).Get(ctx, &output)
	return &output, err
}

func (a *ApiGatewayV2Stub) DeleteApiMappingAsync(ctx workflow.Context, input *apigatewayv2.DeleteApiMappingInput) *Apigatewayv2DeleteApiMappingResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeleteApiMapping, input)
	return &Apigatewayv2DeleteApiMappingResult{Result: future}
}

func (a *ApiGatewayV2Stub) DeleteAuthorizer(ctx workflow.Context, input *apigatewayv2.DeleteAuthorizerInput) (*apigatewayv2.DeleteAuthorizerOutput, error) {
	var output apigatewayv2.DeleteAuthorizerOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeleteAuthorizer, input).Get(ctx, &output)
	return &output, err
}

func (a *ApiGatewayV2Stub) DeleteAuthorizerAsync(ctx workflow.Context, input *apigatewayv2.DeleteAuthorizerInput) *Apigatewayv2DeleteAuthorizerResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeleteAuthorizer, input)
	return &Apigatewayv2DeleteAuthorizerResult{Result: future}
}

func (a *ApiGatewayV2Stub) DeleteCorsConfiguration(ctx workflow.Context, input *apigatewayv2.DeleteCorsConfigurationInput) (*apigatewayv2.DeleteCorsConfigurationOutput, error) {
	var output apigatewayv2.DeleteCorsConfigurationOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeleteCorsConfiguration, input).Get(ctx, &output)
	return &output, err
}

func (a *ApiGatewayV2Stub) DeleteCorsConfigurationAsync(ctx workflow.Context, input *apigatewayv2.DeleteCorsConfigurationInput) *Apigatewayv2DeleteCorsConfigurationResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeleteCorsConfiguration, input)
	return &Apigatewayv2DeleteCorsConfigurationResult{Result: future}
}

func (a *ApiGatewayV2Stub) DeleteDeployment(ctx workflow.Context, input *apigatewayv2.DeleteDeploymentInput) (*apigatewayv2.DeleteDeploymentOutput, error) {
	var output apigatewayv2.DeleteDeploymentOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeleteDeployment, input).Get(ctx, &output)
	return &output, err
}

func (a *ApiGatewayV2Stub) DeleteDeploymentAsync(ctx workflow.Context, input *apigatewayv2.DeleteDeploymentInput) *Apigatewayv2DeleteDeploymentResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeleteDeployment, input)
	return &Apigatewayv2DeleteDeploymentResult{Result: future}
}

func (a *ApiGatewayV2Stub) DeleteDomainName(ctx workflow.Context, input *apigatewayv2.DeleteDomainNameInput) (*apigatewayv2.DeleteDomainNameOutput, error) {
	var output apigatewayv2.DeleteDomainNameOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeleteDomainName, input).Get(ctx, &output)
	return &output, err
}

func (a *ApiGatewayV2Stub) DeleteDomainNameAsync(ctx workflow.Context, input *apigatewayv2.DeleteDomainNameInput) *Apigatewayv2DeleteDomainNameResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeleteDomainName, input)
	return &Apigatewayv2DeleteDomainNameResult{Result: future}
}

func (a *ApiGatewayV2Stub) DeleteIntegration(ctx workflow.Context, input *apigatewayv2.DeleteIntegrationInput) (*apigatewayv2.DeleteIntegrationOutput, error) {
	var output apigatewayv2.DeleteIntegrationOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeleteIntegration, input).Get(ctx, &output)
	return &output, err
}

func (a *ApiGatewayV2Stub) DeleteIntegrationAsync(ctx workflow.Context, input *apigatewayv2.DeleteIntegrationInput) *Apigatewayv2DeleteIntegrationResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeleteIntegration, input)
	return &Apigatewayv2DeleteIntegrationResult{Result: future}
}

func (a *ApiGatewayV2Stub) DeleteIntegrationResponse(ctx workflow.Context, input *apigatewayv2.DeleteIntegrationResponseInput) (*apigatewayv2.DeleteIntegrationResponseOutput, error) {
	var output apigatewayv2.DeleteIntegrationResponseOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeleteIntegrationResponse, input).Get(ctx, &output)
	return &output, err
}

func (a *ApiGatewayV2Stub) DeleteIntegrationResponseAsync(ctx workflow.Context, input *apigatewayv2.DeleteIntegrationResponseInput) *Apigatewayv2DeleteIntegrationResponseResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeleteIntegrationResponse, input)
	return &Apigatewayv2DeleteIntegrationResponseResult{Result: future}
}

func (a *ApiGatewayV2Stub) DeleteModel(ctx workflow.Context, input *apigatewayv2.DeleteModelInput) (*apigatewayv2.DeleteModelOutput, error) {
	var output apigatewayv2.DeleteModelOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeleteModel, input).Get(ctx, &output)
	return &output, err
}

func (a *ApiGatewayV2Stub) DeleteModelAsync(ctx workflow.Context, input *apigatewayv2.DeleteModelInput) *Apigatewayv2DeleteModelResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeleteModel, input)
	return &Apigatewayv2DeleteModelResult{Result: future}
}

func (a *ApiGatewayV2Stub) DeleteRoute(ctx workflow.Context, input *apigatewayv2.DeleteRouteInput) (*apigatewayv2.DeleteRouteOutput, error) {
	var output apigatewayv2.DeleteRouteOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeleteRoute, input).Get(ctx, &output)
	return &output, err
}

func (a *ApiGatewayV2Stub) DeleteRouteAsync(ctx workflow.Context, input *apigatewayv2.DeleteRouteInput) *Apigatewayv2DeleteRouteResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeleteRoute, input)
	return &Apigatewayv2DeleteRouteResult{Result: future}
}

func (a *ApiGatewayV2Stub) DeleteRouteRequestParameter(ctx workflow.Context, input *apigatewayv2.DeleteRouteRequestParameterInput) (*apigatewayv2.DeleteRouteRequestParameterOutput, error) {
	var output apigatewayv2.DeleteRouteRequestParameterOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeleteRouteRequestParameter, input).Get(ctx, &output)
	return &output, err
}

func (a *ApiGatewayV2Stub) DeleteRouteRequestParameterAsync(ctx workflow.Context, input *apigatewayv2.DeleteRouteRequestParameterInput) *Apigatewayv2DeleteRouteRequestParameterResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeleteRouteRequestParameter, input)
	return &Apigatewayv2DeleteRouteRequestParameterResult{Result: future}
}

func (a *ApiGatewayV2Stub) DeleteRouteResponse(ctx workflow.Context, input *apigatewayv2.DeleteRouteResponseInput) (*apigatewayv2.DeleteRouteResponseOutput, error) {
	var output apigatewayv2.DeleteRouteResponseOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeleteRouteResponse, input).Get(ctx, &output)
	return &output, err
}

func (a *ApiGatewayV2Stub) DeleteRouteResponseAsync(ctx workflow.Context, input *apigatewayv2.DeleteRouteResponseInput) *Apigatewayv2DeleteRouteResponseResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeleteRouteResponse, input)
	return &Apigatewayv2DeleteRouteResponseResult{Result: future}
}

func (a *ApiGatewayV2Stub) DeleteRouteSettings(ctx workflow.Context, input *apigatewayv2.DeleteRouteSettingsInput) (*apigatewayv2.DeleteRouteSettingsOutput, error) {
	var output apigatewayv2.DeleteRouteSettingsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeleteRouteSettings, input).Get(ctx, &output)
	return &output, err
}

func (a *ApiGatewayV2Stub) DeleteRouteSettingsAsync(ctx workflow.Context, input *apigatewayv2.DeleteRouteSettingsInput) *Apigatewayv2DeleteRouteSettingsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeleteRouteSettings, input)
	return &Apigatewayv2DeleteRouteSettingsResult{Result: future}
}

func (a *ApiGatewayV2Stub) DeleteStage(ctx workflow.Context, input *apigatewayv2.DeleteStageInput) (*apigatewayv2.DeleteStageOutput, error) {
	var output apigatewayv2.DeleteStageOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeleteStage, input).Get(ctx, &output)
	return &output, err
}

func (a *ApiGatewayV2Stub) DeleteStageAsync(ctx workflow.Context, input *apigatewayv2.DeleteStageInput) *Apigatewayv2DeleteStageResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeleteStage, input)
	return &Apigatewayv2DeleteStageResult{Result: future}
}

func (a *ApiGatewayV2Stub) DeleteVpcLink(ctx workflow.Context, input *apigatewayv2.DeleteVpcLinkInput) (*apigatewayv2.DeleteVpcLinkOutput, error) {
	var output apigatewayv2.DeleteVpcLinkOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeleteVpcLink, input).Get(ctx, &output)
	return &output, err
}

func (a *ApiGatewayV2Stub) DeleteVpcLinkAsync(ctx workflow.Context, input *apigatewayv2.DeleteVpcLinkInput) *Apigatewayv2DeleteVpcLinkResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeleteVpcLink, input)
	return &Apigatewayv2DeleteVpcLinkResult{Result: future}
}

func (a *ApiGatewayV2Stub) ExportApi(ctx workflow.Context, input *apigatewayv2.ExportApiInput) (*apigatewayv2.ExportApiOutput, error) {
	var output apigatewayv2.ExportApiOutput
	err := workflow.ExecuteActivity(ctx, a.activities.ExportApi, input).Get(ctx, &output)
	return &output, err
}

func (a *ApiGatewayV2Stub) ExportApiAsync(ctx workflow.Context, input *apigatewayv2.ExportApiInput) *Apigatewayv2ExportApiResult {
	future := workflow.ExecuteActivity(ctx, a.activities.ExportApi, input)
	return &Apigatewayv2ExportApiResult{Result: future}
}

func (a *ApiGatewayV2Stub) GetApi(ctx workflow.Context, input *apigatewayv2.GetApiInput) (*apigatewayv2.GetApiOutput, error) {
	var output apigatewayv2.GetApiOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetApi, input).Get(ctx, &output)
	return &output, err
}

func (a *ApiGatewayV2Stub) GetApiAsync(ctx workflow.Context, input *apigatewayv2.GetApiInput) *Apigatewayv2GetApiResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetApi, input)
	return &Apigatewayv2GetApiResult{Result: future}
}

func (a *ApiGatewayV2Stub) GetApiMapping(ctx workflow.Context, input *apigatewayv2.GetApiMappingInput) (*apigatewayv2.GetApiMappingOutput, error) {
	var output apigatewayv2.GetApiMappingOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetApiMapping, input).Get(ctx, &output)
	return &output, err
}

func (a *ApiGatewayV2Stub) GetApiMappingAsync(ctx workflow.Context, input *apigatewayv2.GetApiMappingInput) *Apigatewayv2GetApiMappingResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetApiMapping, input)
	return &Apigatewayv2GetApiMappingResult{Result: future}
}

func (a *ApiGatewayV2Stub) GetApiMappings(ctx workflow.Context, input *apigatewayv2.GetApiMappingsInput) (*apigatewayv2.GetApiMappingsOutput, error) {
	var output apigatewayv2.GetApiMappingsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetApiMappings, input).Get(ctx, &output)
	return &output, err
}

func (a *ApiGatewayV2Stub) GetApiMappingsAsync(ctx workflow.Context, input *apigatewayv2.GetApiMappingsInput) *Apigatewayv2GetApiMappingsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetApiMappings, input)
	return &Apigatewayv2GetApiMappingsResult{Result: future}
}

func (a *ApiGatewayV2Stub) GetApis(ctx workflow.Context, input *apigatewayv2.GetApisInput) (*apigatewayv2.GetApisOutput, error) {
	var output apigatewayv2.GetApisOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetApis, input).Get(ctx, &output)
	return &output, err
}

func (a *ApiGatewayV2Stub) GetApisAsync(ctx workflow.Context, input *apigatewayv2.GetApisInput) *Apigatewayv2GetApisResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetApis, input)
	return &Apigatewayv2GetApisResult{Result: future}
}

func (a *ApiGatewayV2Stub) GetAuthorizer(ctx workflow.Context, input *apigatewayv2.GetAuthorizerInput) (*apigatewayv2.GetAuthorizerOutput, error) {
	var output apigatewayv2.GetAuthorizerOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetAuthorizer, input).Get(ctx, &output)
	return &output, err
}

func (a *ApiGatewayV2Stub) GetAuthorizerAsync(ctx workflow.Context, input *apigatewayv2.GetAuthorizerInput) *Apigatewayv2GetAuthorizerResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetAuthorizer, input)
	return &Apigatewayv2GetAuthorizerResult{Result: future}
}

func (a *ApiGatewayV2Stub) GetAuthorizers(ctx workflow.Context, input *apigatewayv2.GetAuthorizersInput) (*apigatewayv2.GetAuthorizersOutput, error) {
	var output apigatewayv2.GetAuthorizersOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetAuthorizers, input).Get(ctx, &output)
	return &output, err
}

func (a *ApiGatewayV2Stub) GetAuthorizersAsync(ctx workflow.Context, input *apigatewayv2.GetAuthorizersInput) *Apigatewayv2GetAuthorizersResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetAuthorizers, input)
	return &Apigatewayv2GetAuthorizersResult{Result: future}
}

func (a *ApiGatewayV2Stub) GetDeployment(ctx workflow.Context, input *apigatewayv2.GetDeploymentInput) (*apigatewayv2.GetDeploymentOutput, error) {
	var output apigatewayv2.GetDeploymentOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetDeployment, input).Get(ctx, &output)
	return &output, err
}

func (a *ApiGatewayV2Stub) GetDeploymentAsync(ctx workflow.Context, input *apigatewayv2.GetDeploymentInput) *Apigatewayv2GetDeploymentResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetDeployment, input)
	return &Apigatewayv2GetDeploymentResult{Result: future}
}

func (a *ApiGatewayV2Stub) GetDeployments(ctx workflow.Context, input *apigatewayv2.GetDeploymentsInput) (*apigatewayv2.GetDeploymentsOutput, error) {
	var output apigatewayv2.GetDeploymentsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetDeployments, input).Get(ctx, &output)
	return &output, err
}

func (a *ApiGatewayV2Stub) GetDeploymentsAsync(ctx workflow.Context, input *apigatewayv2.GetDeploymentsInput) *Apigatewayv2GetDeploymentsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetDeployments, input)
	return &Apigatewayv2GetDeploymentsResult{Result: future}
}

func (a *ApiGatewayV2Stub) GetDomainName(ctx workflow.Context, input *apigatewayv2.GetDomainNameInput) (*apigatewayv2.GetDomainNameOutput, error) {
	var output apigatewayv2.GetDomainNameOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetDomainName, input).Get(ctx, &output)
	return &output, err
}

func (a *ApiGatewayV2Stub) GetDomainNameAsync(ctx workflow.Context, input *apigatewayv2.GetDomainNameInput) *Apigatewayv2GetDomainNameResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetDomainName, input)
	return &Apigatewayv2GetDomainNameResult{Result: future}
}

func (a *ApiGatewayV2Stub) GetDomainNames(ctx workflow.Context, input *apigatewayv2.GetDomainNamesInput) (*apigatewayv2.GetDomainNamesOutput, error) {
	var output apigatewayv2.GetDomainNamesOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetDomainNames, input).Get(ctx, &output)
	return &output, err
}

func (a *ApiGatewayV2Stub) GetDomainNamesAsync(ctx workflow.Context, input *apigatewayv2.GetDomainNamesInput) *Apigatewayv2GetDomainNamesResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetDomainNames, input)
	return &Apigatewayv2GetDomainNamesResult{Result: future}
}

func (a *ApiGatewayV2Stub) GetIntegration(ctx workflow.Context, input *apigatewayv2.GetIntegrationInput) (*apigatewayv2.GetIntegrationOutput, error) {
	var output apigatewayv2.GetIntegrationOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetIntegration, input).Get(ctx, &output)
	return &output, err
}

func (a *ApiGatewayV2Stub) GetIntegrationAsync(ctx workflow.Context, input *apigatewayv2.GetIntegrationInput) *Apigatewayv2GetIntegrationResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetIntegration, input)
	return &Apigatewayv2GetIntegrationResult{Result: future}
}

func (a *ApiGatewayV2Stub) GetIntegrationResponse(ctx workflow.Context, input *apigatewayv2.GetIntegrationResponseInput) (*apigatewayv2.GetIntegrationResponseOutput, error) {
	var output apigatewayv2.GetIntegrationResponseOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetIntegrationResponse, input).Get(ctx, &output)
	return &output, err
}

func (a *ApiGatewayV2Stub) GetIntegrationResponseAsync(ctx workflow.Context, input *apigatewayv2.GetIntegrationResponseInput) *Apigatewayv2GetIntegrationResponseResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetIntegrationResponse, input)
	return &Apigatewayv2GetIntegrationResponseResult{Result: future}
}

func (a *ApiGatewayV2Stub) GetIntegrationResponses(ctx workflow.Context, input *apigatewayv2.GetIntegrationResponsesInput) (*apigatewayv2.GetIntegrationResponsesOutput, error) {
	var output apigatewayv2.GetIntegrationResponsesOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetIntegrationResponses, input).Get(ctx, &output)
	return &output, err
}

func (a *ApiGatewayV2Stub) GetIntegrationResponsesAsync(ctx workflow.Context, input *apigatewayv2.GetIntegrationResponsesInput) *Apigatewayv2GetIntegrationResponsesResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetIntegrationResponses, input)
	return &Apigatewayv2GetIntegrationResponsesResult{Result: future}
}

func (a *ApiGatewayV2Stub) GetIntegrations(ctx workflow.Context, input *apigatewayv2.GetIntegrationsInput) (*apigatewayv2.GetIntegrationsOutput, error) {
	var output apigatewayv2.GetIntegrationsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetIntegrations, input).Get(ctx, &output)
	return &output, err
}

func (a *ApiGatewayV2Stub) GetIntegrationsAsync(ctx workflow.Context, input *apigatewayv2.GetIntegrationsInput) *Apigatewayv2GetIntegrationsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetIntegrations, input)
	return &Apigatewayv2GetIntegrationsResult{Result: future}
}

func (a *ApiGatewayV2Stub) GetModel(ctx workflow.Context, input *apigatewayv2.GetModelInput) (*apigatewayv2.GetModelOutput, error) {
	var output apigatewayv2.GetModelOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetModel, input).Get(ctx, &output)
	return &output, err
}

func (a *ApiGatewayV2Stub) GetModelAsync(ctx workflow.Context, input *apigatewayv2.GetModelInput) *Apigatewayv2GetModelResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetModel, input)
	return &Apigatewayv2GetModelResult{Result: future}
}

func (a *ApiGatewayV2Stub) GetModelTemplate(ctx workflow.Context, input *apigatewayv2.GetModelTemplateInput) (*apigatewayv2.GetModelTemplateOutput, error) {
	var output apigatewayv2.GetModelTemplateOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetModelTemplate, input).Get(ctx, &output)
	return &output, err
}

func (a *ApiGatewayV2Stub) GetModelTemplateAsync(ctx workflow.Context, input *apigatewayv2.GetModelTemplateInput) *Apigatewayv2GetModelTemplateResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetModelTemplate, input)
	return &Apigatewayv2GetModelTemplateResult{Result: future}
}

func (a *ApiGatewayV2Stub) GetModels(ctx workflow.Context, input *apigatewayv2.GetModelsInput) (*apigatewayv2.GetModelsOutput, error) {
	var output apigatewayv2.GetModelsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetModels, input).Get(ctx, &output)
	return &output, err
}

func (a *ApiGatewayV2Stub) GetModelsAsync(ctx workflow.Context, input *apigatewayv2.GetModelsInput) *Apigatewayv2GetModelsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetModels, input)
	return &Apigatewayv2GetModelsResult{Result: future}
}

func (a *ApiGatewayV2Stub) GetRoute(ctx workflow.Context, input *apigatewayv2.GetRouteInput) (*apigatewayv2.GetRouteOutput, error) {
	var output apigatewayv2.GetRouteOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetRoute, input).Get(ctx, &output)
	return &output, err
}

func (a *ApiGatewayV2Stub) GetRouteAsync(ctx workflow.Context, input *apigatewayv2.GetRouteInput) *Apigatewayv2GetRouteResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetRoute, input)
	return &Apigatewayv2GetRouteResult{Result: future}
}

func (a *ApiGatewayV2Stub) GetRouteResponse(ctx workflow.Context, input *apigatewayv2.GetRouteResponseInput) (*apigatewayv2.GetRouteResponseOutput, error) {
	var output apigatewayv2.GetRouteResponseOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetRouteResponse, input).Get(ctx, &output)
	return &output, err
}

func (a *ApiGatewayV2Stub) GetRouteResponseAsync(ctx workflow.Context, input *apigatewayv2.GetRouteResponseInput) *Apigatewayv2GetRouteResponseResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetRouteResponse, input)
	return &Apigatewayv2GetRouteResponseResult{Result: future}
}

func (a *ApiGatewayV2Stub) GetRouteResponses(ctx workflow.Context, input *apigatewayv2.GetRouteResponsesInput) (*apigatewayv2.GetRouteResponsesOutput, error) {
	var output apigatewayv2.GetRouteResponsesOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetRouteResponses, input).Get(ctx, &output)
	return &output, err
}

func (a *ApiGatewayV2Stub) GetRouteResponsesAsync(ctx workflow.Context, input *apigatewayv2.GetRouteResponsesInput) *Apigatewayv2GetRouteResponsesResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetRouteResponses, input)
	return &Apigatewayv2GetRouteResponsesResult{Result: future}
}

func (a *ApiGatewayV2Stub) GetRoutes(ctx workflow.Context, input *apigatewayv2.GetRoutesInput) (*apigatewayv2.GetRoutesOutput, error) {
	var output apigatewayv2.GetRoutesOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetRoutes, input).Get(ctx, &output)
	return &output, err
}

func (a *ApiGatewayV2Stub) GetRoutesAsync(ctx workflow.Context, input *apigatewayv2.GetRoutesInput) *Apigatewayv2GetRoutesResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetRoutes, input)
	return &Apigatewayv2GetRoutesResult{Result: future}
}

func (a *ApiGatewayV2Stub) GetStage(ctx workflow.Context, input *apigatewayv2.GetStageInput) (*apigatewayv2.GetStageOutput, error) {
	var output apigatewayv2.GetStageOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetStage, input).Get(ctx, &output)
	return &output, err
}

func (a *ApiGatewayV2Stub) GetStageAsync(ctx workflow.Context, input *apigatewayv2.GetStageInput) *Apigatewayv2GetStageResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetStage, input)
	return &Apigatewayv2GetStageResult{Result: future}
}

func (a *ApiGatewayV2Stub) GetStages(ctx workflow.Context, input *apigatewayv2.GetStagesInput) (*apigatewayv2.GetStagesOutput, error) {
	var output apigatewayv2.GetStagesOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetStages, input).Get(ctx, &output)
	return &output, err
}

func (a *ApiGatewayV2Stub) GetStagesAsync(ctx workflow.Context, input *apigatewayv2.GetStagesInput) *Apigatewayv2GetStagesResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetStages, input)
	return &Apigatewayv2GetStagesResult{Result: future}
}

func (a *ApiGatewayV2Stub) GetTags(ctx workflow.Context, input *apigatewayv2.GetTagsInput) (*apigatewayv2.GetTagsOutput, error) {
	var output apigatewayv2.GetTagsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetTags, input).Get(ctx, &output)
	return &output, err
}

func (a *ApiGatewayV2Stub) GetTagsAsync(ctx workflow.Context, input *apigatewayv2.GetTagsInput) *Apigatewayv2GetTagsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetTags, input)
	return &Apigatewayv2GetTagsResult{Result: future}
}

func (a *ApiGatewayV2Stub) GetVpcLink(ctx workflow.Context, input *apigatewayv2.GetVpcLinkInput) (*apigatewayv2.GetVpcLinkOutput, error) {
	var output apigatewayv2.GetVpcLinkOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetVpcLink, input).Get(ctx, &output)
	return &output, err
}

func (a *ApiGatewayV2Stub) GetVpcLinkAsync(ctx workflow.Context, input *apigatewayv2.GetVpcLinkInput) *Apigatewayv2GetVpcLinkResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetVpcLink, input)
	return &Apigatewayv2GetVpcLinkResult{Result: future}
}

func (a *ApiGatewayV2Stub) GetVpcLinks(ctx workflow.Context, input *apigatewayv2.GetVpcLinksInput) (*apigatewayv2.GetVpcLinksOutput, error) {
	var output apigatewayv2.GetVpcLinksOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetVpcLinks, input).Get(ctx, &output)
	return &output, err
}

func (a *ApiGatewayV2Stub) GetVpcLinksAsync(ctx workflow.Context, input *apigatewayv2.GetVpcLinksInput) *Apigatewayv2GetVpcLinksResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetVpcLinks, input)
	return &Apigatewayv2GetVpcLinksResult{Result: future}
}

func (a *ApiGatewayV2Stub) ImportApi(ctx workflow.Context, input *apigatewayv2.ImportApiInput) (*apigatewayv2.ImportApiOutput, error) {
	var output apigatewayv2.ImportApiOutput
	err := workflow.ExecuteActivity(ctx, a.activities.ImportApi, input).Get(ctx, &output)
	return &output, err
}

func (a *ApiGatewayV2Stub) ImportApiAsync(ctx workflow.Context, input *apigatewayv2.ImportApiInput) *Apigatewayv2ImportApiResult {
	future := workflow.ExecuteActivity(ctx, a.activities.ImportApi, input)
	return &Apigatewayv2ImportApiResult{Result: future}
}

func (a *ApiGatewayV2Stub) ReimportApi(ctx workflow.Context, input *apigatewayv2.ReimportApiInput) (*apigatewayv2.ReimportApiOutput, error) {
	var output apigatewayv2.ReimportApiOutput
	err := workflow.ExecuteActivity(ctx, a.activities.ReimportApi, input).Get(ctx, &output)
	return &output, err
}

func (a *ApiGatewayV2Stub) ReimportApiAsync(ctx workflow.Context, input *apigatewayv2.ReimportApiInput) *Apigatewayv2ReimportApiResult {
	future := workflow.ExecuteActivity(ctx, a.activities.ReimportApi, input)
	return &Apigatewayv2ReimportApiResult{Result: future}
}

func (a *ApiGatewayV2Stub) TagResource(ctx workflow.Context, input *apigatewayv2.TagResourceInput) (*apigatewayv2.TagResourceOutput, error) {
	var output apigatewayv2.TagResourceOutput
	err := workflow.ExecuteActivity(ctx, a.activities.TagResource, input).Get(ctx, &output)
	return &output, err
}

func (a *ApiGatewayV2Stub) TagResourceAsync(ctx workflow.Context, input *apigatewayv2.TagResourceInput) *Apigatewayv2TagResourceResult {
	future := workflow.ExecuteActivity(ctx, a.activities.TagResource, input)
	return &Apigatewayv2TagResourceResult{Result: future}
}

func (a *ApiGatewayV2Stub) UntagResource(ctx workflow.Context, input *apigatewayv2.UntagResourceInput) (*apigatewayv2.UntagResourceOutput, error) {
	var output apigatewayv2.UntagResourceOutput
	err := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input).Get(ctx, &output)
	return &output, err
}

func (a *ApiGatewayV2Stub) UntagResourceAsync(ctx workflow.Context, input *apigatewayv2.UntagResourceInput) *Apigatewayv2UntagResourceResult {
	future := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input)
	return &Apigatewayv2UntagResourceResult{Result: future}
}

func (a *ApiGatewayV2Stub) UpdateApi(ctx workflow.Context, input *apigatewayv2.UpdateApiInput) (*apigatewayv2.UpdateApiOutput, error) {
	var output apigatewayv2.UpdateApiOutput
	err := workflow.ExecuteActivity(ctx, a.activities.UpdateApi, input).Get(ctx, &output)
	return &output, err
}

func (a *ApiGatewayV2Stub) UpdateApiAsync(ctx workflow.Context, input *apigatewayv2.UpdateApiInput) *Apigatewayv2UpdateApiResult {
	future := workflow.ExecuteActivity(ctx, a.activities.UpdateApi, input)
	return &Apigatewayv2UpdateApiResult{Result: future}
}

func (a *ApiGatewayV2Stub) UpdateApiMapping(ctx workflow.Context, input *apigatewayv2.UpdateApiMappingInput) (*apigatewayv2.UpdateApiMappingOutput, error) {
	var output apigatewayv2.UpdateApiMappingOutput
	err := workflow.ExecuteActivity(ctx, a.activities.UpdateApiMapping, input).Get(ctx, &output)
	return &output, err
}

func (a *ApiGatewayV2Stub) UpdateApiMappingAsync(ctx workflow.Context, input *apigatewayv2.UpdateApiMappingInput) *Apigatewayv2UpdateApiMappingResult {
	future := workflow.ExecuteActivity(ctx, a.activities.UpdateApiMapping, input)
	return &Apigatewayv2UpdateApiMappingResult{Result: future}
}

func (a *ApiGatewayV2Stub) UpdateAuthorizer(ctx workflow.Context, input *apigatewayv2.UpdateAuthorizerInput) (*apigatewayv2.UpdateAuthorizerOutput, error) {
	var output apigatewayv2.UpdateAuthorizerOutput
	err := workflow.ExecuteActivity(ctx, a.activities.UpdateAuthorizer, input).Get(ctx, &output)
	return &output, err
}

func (a *ApiGatewayV2Stub) UpdateAuthorizerAsync(ctx workflow.Context, input *apigatewayv2.UpdateAuthorizerInput) *Apigatewayv2UpdateAuthorizerResult {
	future := workflow.ExecuteActivity(ctx, a.activities.UpdateAuthorizer, input)
	return &Apigatewayv2UpdateAuthorizerResult{Result: future}
}

func (a *ApiGatewayV2Stub) UpdateDeployment(ctx workflow.Context, input *apigatewayv2.UpdateDeploymentInput) (*apigatewayv2.UpdateDeploymentOutput, error) {
	var output apigatewayv2.UpdateDeploymentOutput
	err := workflow.ExecuteActivity(ctx, a.activities.UpdateDeployment, input).Get(ctx, &output)
	return &output, err
}

func (a *ApiGatewayV2Stub) UpdateDeploymentAsync(ctx workflow.Context, input *apigatewayv2.UpdateDeploymentInput) *Apigatewayv2UpdateDeploymentResult {
	future := workflow.ExecuteActivity(ctx, a.activities.UpdateDeployment, input)
	return &Apigatewayv2UpdateDeploymentResult{Result: future}
}

func (a *ApiGatewayV2Stub) UpdateDomainName(ctx workflow.Context, input *apigatewayv2.UpdateDomainNameInput) (*apigatewayv2.UpdateDomainNameOutput, error) {
	var output apigatewayv2.UpdateDomainNameOutput
	err := workflow.ExecuteActivity(ctx, a.activities.UpdateDomainName, input).Get(ctx, &output)
	return &output, err
}

func (a *ApiGatewayV2Stub) UpdateDomainNameAsync(ctx workflow.Context, input *apigatewayv2.UpdateDomainNameInput) *Apigatewayv2UpdateDomainNameResult {
	future := workflow.ExecuteActivity(ctx, a.activities.UpdateDomainName, input)
	return &Apigatewayv2UpdateDomainNameResult{Result: future}
}

func (a *ApiGatewayV2Stub) UpdateIntegration(ctx workflow.Context, input *apigatewayv2.UpdateIntegrationInput) (*apigatewayv2.UpdateIntegrationOutput, error) {
	var output apigatewayv2.UpdateIntegrationOutput
	err := workflow.ExecuteActivity(ctx, a.activities.UpdateIntegration, input).Get(ctx, &output)
	return &output, err
}

func (a *ApiGatewayV2Stub) UpdateIntegrationAsync(ctx workflow.Context, input *apigatewayv2.UpdateIntegrationInput) *Apigatewayv2UpdateIntegrationResult {
	future := workflow.ExecuteActivity(ctx, a.activities.UpdateIntegration, input)
	return &Apigatewayv2UpdateIntegrationResult{Result: future}
}

func (a *ApiGatewayV2Stub) UpdateIntegrationResponse(ctx workflow.Context, input *apigatewayv2.UpdateIntegrationResponseInput) (*apigatewayv2.UpdateIntegrationResponseOutput, error) {
	var output apigatewayv2.UpdateIntegrationResponseOutput
	err := workflow.ExecuteActivity(ctx, a.activities.UpdateIntegrationResponse, input).Get(ctx, &output)
	return &output, err
}

func (a *ApiGatewayV2Stub) UpdateIntegrationResponseAsync(ctx workflow.Context, input *apigatewayv2.UpdateIntegrationResponseInput) *Apigatewayv2UpdateIntegrationResponseResult {
	future := workflow.ExecuteActivity(ctx, a.activities.UpdateIntegrationResponse, input)
	return &Apigatewayv2UpdateIntegrationResponseResult{Result: future}
}

func (a *ApiGatewayV2Stub) UpdateModel(ctx workflow.Context, input *apigatewayv2.UpdateModelInput) (*apigatewayv2.UpdateModelOutput, error) {
	var output apigatewayv2.UpdateModelOutput
	err := workflow.ExecuteActivity(ctx, a.activities.UpdateModel, input).Get(ctx, &output)
	return &output, err
}

func (a *ApiGatewayV2Stub) UpdateModelAsync(ctx workflow.Context, input *apigatewayv2.UpdateModelInput) *Apigatewayv2UpdateModelResult {
	future := workflow.ExecuteActivity(ctx, a.activities.UpdateModel, input)
	return &Apigatewayv2UpdateModelResult{Result: future}
}

func (a *ApiGatewayV2Stub) UpdateRoute(ctx workflow.Context, input *apigatewayv2.UpdateRouteInput) (*apigatewayv2.UpdateRouteOutput, error) {
	var output apigatewayv2.UpdateRouteOutput
	err := workflow.ExecuteActivity(ctx, a.activities.UpdateRoute, input).Get(ctx, &output)
	return &output, err
}

func (a *ApiGatewayV2Stub) UpdateRouteAsync(ctx workflow.Context, input *apigatewayv2.UpdateRouteInput) *Apigatewayv2UpdateRouteResult {
	future := workflow.ExecuteActivity(ctx, a.activities.UpdateRoute, input)
	return &Apigatewayv2UpdateRouteResult{Result: future}
}

func (a *ApiGatewayV2Stub) UpdateRouteResponse(ctx workflow.Context, input *apigatewayv2.UpdateRouteResponseInput) (*apigatewayv2.UpdateRouteResponseOutput, error) {
	var output apigatewayv2.UpdateRouteResponseOutput
	err := workflow.ExecuteActivity(ctx, a.activities.UpdateRouteResponse, input).Get(ctx, &output)
	return &output, err
}

func (a *ApiGatewayV2Stub) UpdateRouteResponseAsync(ctx workflow.Context, input *apigatewayv2.UpdateRouteResponseInput) *Apigatewayv2UpdateRouteResponseResult {
	future := workflow.ExecuteActivity(ctx, a.activities.UpdateRouteResponse, input)
	return &Apigatewayv2UpdateRouteResponseResult{Result: future}
}

func (a *ApiGatewayV2Stub) UpdateStage(ctx workflow.Context, input *apigatewayv2.UpdateStageInput) (*apigatewayv2.UpdateStageOutput, error) {
	var output apigatewayv2.UpdateStageOutput
	err := workflow.ExecuteActivity(ctx, a.activities.UpdateStage, input).Get(ctx, &output)
	return &output, err
}

func (a *ApiGatewayV2Stub) UpdateStageAsync(ctx workflow.Context, input *apigatewayv2.UpdateStageInput) *Apigatewayv2UpdateStageResult {
	future := workflow.ExecuteActivity(ctx, a.activities.UpdateStage, input)
	return &Apigatewayv2UpdateStageResult{Result: future}
}

func (a *ApiGatewayV2Stub) UpdateVpcLink(ctx workflow.Context, input *apigatewayv2.UpdateVpcLinkInput) (*apigatewayv2.UpdateVpcLinkOutput, error) {
	var output apigatewayv2.UpdateVpcLinkOutput
	err := workflow.ExecuteActivity(ctx, a.activities.UpdateVpcLink, input).Get(ctx, &output)
	return &output, err
}

func (a *ApiGatewayV2Stub) UpdateVpcLinkAsync(ctx workflow.Context, input *apigatewayv2.UpdateVpcLinkInput) *Apigatewayv2UpdateVpcLinkResult {
	future := workflow.ExecuteActivity(ctx, a.activities.UpdateVpcLink, input)
	return &Apigatewayv2UpdateVpcLinkResult{Result: future}
}
