package awsclients

import (
	"github.com/aws/aws-sdk-go/service/apigateway"
	"go.temporal.io/sdk/workflow"
	"temporal.io/aws-sdk/awsactivities"
)

type APIGatewayClient interface {
	CreateApiKey(ctx workflow.Context, input *apigateway.CreateApiKeyInput) (*apigateway.ApiKey, error)
	CreateApiKeyAsync(ctx workflow.Context, input *apigateway.CreateApiKeyInput) *ApigatewayCreateApiKeyResult

	CreateAuthorizer(ctx workflow.Context, input *apigateway.CreateAuthorizerInput) (*apigateway.Authorizer, error)
	CreateAuthorizerAsync(ctx workflow.Context, input *apigateway.CreateAuthorizerInput) *ApigatewayCreateAuthorizerResult

	CreateBasePathMapping(ctx workflow.Context, input *apigateway.CreateBasePathMappingInput) (*apigateway.BasePathMapping, error)
	CreateBasePathMappingAsync(ctx workflow.Context, input *apigateway.CreateBasePathMappingInput) *ApigatewayCreateBasePathMappingResult

	CreateDeployment(ctx workflow.Context, input *apigateway.CreateDeploymentInput) (*apigateway.Deployment, error)
	CreateDeploymentAsync(ctx workflow.Context, input *apigateway.CreateDeploymentInput) *ApigatewayCreateDeploymentResult

	CreateDocumentationPart(ctx workflow.Context, input *apigateway.CreateDocumentationPartInput) (*apigateway.DocumentationPart, error)
	CreateDocumentationPartAsync(ctx workflow.Context, input *apigateway.CreateDocumentationPartInput) *ApigatewayCreateDocumentationPartResult

	CreateDocumentationVersion(ctx workflow.Context, input *apigateway.CreateDocumentationVersionInput) (*apigateway.DocumentationVersion, error)
	CreateDocumentationVersionAsync(ctx workflow.Context, input *apigateway.CreateDocumentationVersionInput) *ApigatewayCreateDocumentationVersionResult

	CreateDomainName(ctx workflow.Context, input *apigateway.CreateDomainNameInput) (*apigateway.DomainName, error)
	CreateDomainNameAsync(ctx workflow.Context, input *apigateway.CreateDomainNameInput) *ApigatewayCreateDomainNameResult

	CreateModel(ctx workflow.Context, input *apigateway.CreateModelInput) (*apigateway.Model, error)
	CreateModelAsync(ctx workflow.Context, input *apigateway.CreateModelInput) *ApigatewayCreateModelResult

	CreateRequestValidator(ctx workflow.Context, input *apigateway.CreateRequestValidatorInput) (*apigateway.UpdateRequestValidatorOutput, error)
	CreateRequestValidatorAsync(ctx workflow.Context, input *apigateway.CreateRequestValidatorInput) *ApigatewayCreateRequestValidatorResult

	CreateResource(ctx workflow.Context, input *apigateway.CreateResourceInput) (*apigateway.Resource, error)
	CreateResourceAsync(ctx workflow.Context, input *apigateway.CreateResourceInput) *ApigatewayCreateResourceResult

	CreateRestApi(ctx workflow.Context, input *apigateway.CreateRestApiInput) (*apigateway.RestApi, error)
	CreateRestApiAsync(ctx workflow.Context, input *apigateway.CreateRestApiInput) *ApigatewayCreateRestApiResult

	CreateStage(ctx workflow.Context, input *apigateway.CreateStageInput) (*apigateway.Stage, error)
	CreateStageAsync(ctx workflow.Context, input *apigateway.CreateStageInput) *ApigatewayCreateStageResult

	CreateUsagePlan(ctx workflow.Context, input *apigateway.CreateUsagePlanInput) (*apigateway.UsagePlan, error)
	CreateUsagePlanAsync(ctx workflow.Context, input *apigateway.CreateUsagePlanInput) *ApigatewayCreateUsagePlanResult

	CreateUsagePlanKey(ctx workflow.Context, input *apigateway.CreateUsagePlanKeyInput) (*apigateway.UsagePlanKey, error)
	CreateUsagePlanKeyAsync(ctx workflow.Context, input *apigateway.CreateUsagePlanKeyInput) *ApigatewayCreateUsagePlanKeyResult

	CreateVpcLink(ctx workflow.Context, input *apigateway.CreateVpcLinkInput) (*apigateway.UpdateVpcLinkOutput, error)
	CreateVpcLinkAsync(ctx workflow.Context, input *apigateway.CreateVpcLinkInput) *ApigatewayCreateVpcLinkResult

	DeleteApiKey(ctx workflow.Context, input *apigateway.DeleteApiKeyInput) (*apigateway.DeleteApiKeyOutput, error)
	DeleteApiKeyAsync(ctx workflow.Context, input *apigateway.DeleteApiKeyInput) *ApigatewayDeleteApiKeyResult

	DeleteAuthorizer(ctx workflow.Context, input *apigateway.DeleteAuthorizerInput) (*apigateway.DeleteAuthorizerOutput, error)
	DeleteAuthorizerAsync(ctx workflow.Context, input *apigateway.DeleteAuthorizerInput) *ApigatewayDeleteAuthorizerResult

	DeleteBasePathMapping(ctx workflow.Context, input *apigateway.DeleteBasePathMappingInput) (*apigateway.DeleteBasePathMappingOutput, error)
	DeleteBasePathMappingAsync(ctx workflow.Context, input *apigateway.DeleteBasePathMappingInput) *ApigatewayDeleteBasePathMappingResult

	DeleteClientCertificate(ctx workflow.Context, input *apigateway.DeleteClientCertificateInput) (*apigateway.DeleteClientCertificateOutput, error)
	DeleteClientCertificateAsync(ctx workflow.Context, input *apigateway.DeleteClientCertificateInput) *ApigatewayDeleteClientCertificateResult

	DeleteDeployment(ctx workflow.Context, input *apigateway.DeleteDeploymentInput) (*apigateway.DeleteDeploymentOutput, error)
	DeleteDeploymentAsync(ctx workflow.Context, input *apigateway.DeleteDeploymentInput) *ApigatewayDeleteDeploymentResult

	DeleteDocumentationPart(ctx workflow.Context, input *apigateway.DeleteDocumentationPartInput) (*apigateway.DeleteDocumentationPartOutput, error)
	DeleteDocumentationPartAsync(ctx workflow.Context, input *apigateway.DeleteDocumentationPartInput) *ApigatewayDeleteDocumentationPartResult

	DeleteDocumentationVersion(ctx workflow.Context, input *apigateway.DeleteDocumentationVersionInput) (*apigateway.DeleteDocumentationVersionOutput, error)
	DeleteDocumentationVersionAsync(ctx workflow.Context, input *apigateway.DeleteDocumentationVersionInput) *ApigatewayDeleteDocumentationVersionResult

	DeleteDomainName(ctx workflow.Context, input *apigateway.DeleteDomainNameInput) (*apigateway.DeleteDomainNameOutput, error)
	DeleteDomainNameAsync(ctx workflow.Context, input *apigateway.DeleteDomainNameInput) *ApigatewayDeleteDomainNameResult

	DeleteGatewayResponse(ctx workflow.Context, input *apigateway.DeleteGatewayResponseInput) (*apigateway.DeleteGatewayResponseOutput, error)
	DeleteGatewayResponseAsync(ctx workflow.Context, input *apigateway.DeleteGatewayResponseInput) *ApigatewayDeleteGatewayResponseResult

	DeleteIntegration(ctx workflow.Context, input *apigateway.DeleteIntegrationInput) (*apigateway.DeleteIntegrationOutput, error)
	DeleteIntegrationAsync(ctx workflow.Context, input *apigateway.DeleteIntegrationInput) *ApigatewayDeleteIntegrationResult

	DeleteIntegrationResponse(ctx workflow.Context, input *apigateway.DeleteIntegrationResponseInput) (*apigateway.DeleteIntegrationResponseOutput, error)
	DeleteIntegrationResponseAsync(ctx workflow.Context, input *apigateway.DeleteIntegrationResponseInput) *ApigatewayDeleteIntegrationResponseResult

	DeleteMethod(ctx workflow.Context, input *apigateway.DeleteMethodInput) (*apigateway.DeleteMethodOutput, error)
	DeleteMethodAsync(ctx workflow.Context, input *apigateway.DeleteMethodInput) *ApigatewayDeleteMethodResult

	DeleteMethodResponse(ctx workflow.Context, input *apigateway.DeleteMethodResponseInput) (*apigateway.DeleteMethodResponseOutput, error)
	DeleteMethodResponseAsync(ctx workflow.Context, input *apigateway.DeleteMethodResponseInput) *ApigatewayDeleteMethodResponseResult

	DeleteModel(ctx workflow.Context, input *apigateway.DeleteModelInput) (*apigateway.DeleteModelOutput, error)
	DeleteModelAsync(ctx workflow.Context, input *apigateway.DeleteModelInput) *ApigatewayDeleteModelResult

	DeleteRequestValidator(ctx workflow.Context, input *apigateway.DeleteRequestValidatorInput) (*apigateway.DeleteRequestValidatorOutput, error)
	DeleteRequestValidatorAsync(ctx workflow.Context, input *apigateway.DeleteRequestValidatorInput) *ApigatewayDeleteRequestValidatorResult

	DeleteResource(ctx workflow.Context, input *apigateway.DeleteResourceInput) (*apigateway.DeleteResourceOutput, error)
	DeleteResourceAsync(ctx workflow.Context, input *apigateway.DeleteResourceInput) *ApigatewayDeleteResourceResult

	DeleteRestApi(ctx workflow.Context, input *apigateway.DeleteRestApiInput) (*apigateway.DeleteRestApiOutput, error)
	DeleteRestApiAsync(ctx workflow.Context, input *apigateway.DeleteRestApiInput) *ApigatewayDeleteRestApiResult

	DeleteStage(ctx workflow.Context, input *apigateway.DeleteStageInput) (*apigateway.DeleteStageOutput, error)
	DeleteStageAsync(ctx workflow.Context, input *apigateway.DeleteStageInput) *ApigatewayDeleteStageResult

	DeleteUsagePlan(ctx workflow.Context, input *apigateway.DeleteUsagePlanInput) (*apigateway.DeleteUsagePlanOutput, error)
	DeleteUsagePlanAsync(ctx workflow.Context, input *apigateway.DeleteUsagePlanInput) *ApigatewayDeleteUsagePlanResult

	DeleteUsagePlanKey(ctx workflow.Context, input *apigateway.DeleteUsagePlanKeyInput) (*apigateway.DeleteUsagePlanKeyOutput, error)
	DeleteUsagePlanKeyAsync(ctx workflow.Context, input *apigateway.DeleteUsagePlanKeyInput) *ApigatewayDeleteUsagePlanKeyResult

	DeleteVpcLink(ctx workflow.Context, input *apigateway.DeleteVpcLinkInput) (*apigateway.DeleteVpcLinkOutput, error)
	DeleteVpcLinkAsync(ctx workflow.Context, input *apigateway.DeleteVpcLinkInput) *ApigatewayDeleteVpcLinkResult

	FlushStageAuthorizersCache(ctx workflow.Context, input *apigateway.FlushStageAuthorizersCacheInput) (*apigateway.FlushStageAuthorizersCacheOutput, error)
	FlushStageAuthorizersCacheAsync(ctx workflow.Context, input *apigateway.FlushStageAuthorizersCacheInput) *ApigatewayFlushStageAuthorizersCacheResult

	FlushStageCache(ctx workflow.Context, input *apigateway.FlushStageCacheInput) (*apigateway.FlushStageCacheOutput, error)
	FlushStageCacheAsync(ctx workflow.Context, input *apigateway.FlushStageCacheInput) *ApigatewayFlushStageCacheResult

	GenerateClientCertificate(ctx workflow.Context, input *apigateway.GenerateClientCertificateInput) (*apigateway.ClientCertificate, error)
	GenerateClientCertificateAsync(ctx workflow.Context, input *apigateway.GenerateClientCertificateInput) *ApigatewayGenerateClientCertificateResult

	GetAccount(ctx workflow.Context, input *apigateway.GetAccountInput) (*apigateway.Account, error)
	GetAccountAsync(ctx workflow.Context, input *apigateway.GetAccountInput) *ApigatewayGetAccountResult

	GetApiKey(ctx workflow.Context, input *apigateway.GetApiKeyInput) (*apigateway.ApiKey, error)
	GetApiKeyAsync(ctx workflow.Context, input *apigateway.GetApiKeyInput) *ApigatewayGetApiKeyResult

	GetApiKeys(ctx workflow.Context, input *apigateway.GetApiKeysInput) (*apigateway.GetApiKeysOutput, error)
	GetApiKeysAsync(ctx workflow.Context, input *apigateway.GetApiKeysInput) *ApigatewayGetApiKeysResult

	GetAuthorizer(ctx workflow.Context, input *apigateway.GetAuthorizerInput) (*apigateway.Authorizer, error)
	GetAuthorizerAsync(ctx workflow.Context, input *apigateway.GetAuthorizerInput) *ApigatewayGetAuthorizerResult

	GetAuthorizers(ctx workflow.Context, input *apigateway.GetAuthorizersInput) (*apigateway.GetAuthorizersOutput, error)
	GetAuthorizersAsync(ctx workflow.Context, input *apigateway.GetAuthorizersInput) *ApigatewayGetAuthorizersResult

	GetBasePathMapping(ctx workflow.Context, input *apigateway.GetBasePathMappingInput) (*apigateway.BasePathMapping, error)
	GetBasePathMappingAsync(ctx workflow.Context, input *apigateway.GetBasePathMappingInput) *ApigatewayGetBasePathMappingResult

	GetBasePathMappings(ctx workflow.Context, input *apigateway.GetBasePathMappingsInput) (*apigateway.GetBasePathMappingsOutput, error)
	GetBasePathMappingsAsync(ctx workflow.Context, input *apigateway.GetBasePathMappingsInput) *ApigatewayGetBasePathMappingsResult

	GetClientCertificate(ctx workflow.Context, input *apigateway.GetClientCertificateInput) (*apigateway.ClientCertificate, error)
	GetClientCertificateAsync(ctx workflow.Context, input *apigateway.GetClientCertificateInput) *ApigatewayGetClientCertificateResult

	GetClientCertificates(ctx workflow.Context, input *apigateway.GetClientCertificatesInput) (*apigateway.GetClientCertificatesOutput, error)
	GetClientCertificatesAsync(ctx workflow.Context, input *apigateway.GetClientCertificatesInput) *ApigatewayGetClientCertificatesResult

	GetDeployment(ctx workflow.Context, input *apigateway.GetDeploymentInput) (*apigateway.Deployment, error)
	GetDeploymentAsync(ctx workflow.Context, input *apigateway.GetDeploymentInput) *ApigatewayGetDeploymentResult

	GetDeployments(ctx workflow.Context, input *apigateway.GetDeploymentsInput) (*apigateway.GetDeploymentsOutput, error)
	GetDeploymentsAsync(ctx workflow.Context, input *apigateway.GetDeploymentsInput) *ApigatewayGetDeploymentsResult

	GetDocumentationPart(ctx workflow.Context, input *apigateway.GetDocumentationPartInput) (*apigateway.DocumentationPart, error)
	GetDocumentationPartAsync(ctx workflow.Context, input *apigateway.GetDocumentationPartInput) *ApigatewayGetDocumentationPartResult

	GetDocumentationParts(ctx workflow.Context, input *apigateway.GetDocumentationPartsInput) (*apigateway.GetDocumentationPartsOutput, error)
	GetDocumentationPartsAsync(ctx workflow.Context, input *apigateway.GetDocumentationPartsInput) *ApigatewayGetDocumentationPartsResult

	GetDocumentationVersion(ctx workflow.Context, input *apigateway.GetDocumentationVersionInput) (*apigateway.DocumentationVersion, error)
	GetDocumentationVersionAsync(ctx workflow.Context, input *apigateway.GetDocumentationVersionInput) *ApigatewayGetDocumentationVersionResult

	GetDocumentationVersions(ctx workflow.Context, input *apigateway.GetDocumentationVersionsInput) (*apigateway.GetDocumentationVersionsOutput, error)
	GetDocumentationVersionsAsync(ctx workflow.Context, input *apigateway.GetDocumentationVersionsInput) *ApigatewayGetDocumentationVersionsResult

	GetDomainName(ctx workflow.Context, input *apigateway.GetDomainNameInput) (*apigateway.DomainName, error)
	GetDomainNameAsync(ctx workflow.Context, input *apigateway.GetDomainNameInput) *ApigatewayGetDomainNameResult

	GetDomainNames(ctx workflow.Context, input *apigateway.GetDomainNamesInput) (*apigateway.GetDomainNamesOutput, error)
	GetDomainNamesAsync(ctx workflow.Context, input *apigateway.GetDomainNamesInput) *ApigatewayGetDomainNamesResult

	GetExport(ctx workflow.Context, input *apigateway.GetExportInput) (*apigateway.GetExportOutput, error)
	GetExportAsync(ctx workflow.Context, input *apigateway.GetExportInput) *ApigatewayGetExportResult

	GetGatewayResponse(ctx workflow.Context, input *apigateway.GetGatewayResponseInput) (*apigateway.UpdateGatewayResponseOutput, error)
	GetGatewayResponseAsync(ctx workflow.Context, input *apigateway.GetGatewayResponseInput) *ApigatewayGetGatewayResponseResult

	GetGatewayResponses(ctx workflow.Context, input *apigateway.GetGatewayResponsesInput) (*apigateway.GetGatewayResponsesOutput, error)
	GetGatewayResponsesAsync(ctx workflow.Context, input *apigateway.GetGatewayResponsesInput) *ApigatewayGetGatewayResponsesResult

	GetIntegration(ctx workflow.Context, input *apigateway.GetIntegrationInput) (*apigateway.Integration, error)
	GetIntegrationAsync(ctx workflow.Context, input *apigateway.GetIntegrationInput) *ApigatewayGetIntegrationResult

	GetIntegrationResponse(ctx workflow.Context, input *apigateway.GetIntegrationResponseInput) (*apigateway.IntegrationResponse, error)
	GetIntegrationResponseAsync(ctx workflow.Context, input *apigateway.GetIntegrationResponseInput) *ApigatewayGetIntegrationResponseResult

	GetMethod(ctx workflow.Context, input *apigateway.GetMethodInput) (*apigateway.Method, error)
	GetMethodAsync(ctx workflow.Context, input *apigateway.GetMethodInput) *ApigatewayGetMethodResult

	GetMethodResponse(ctx workflow.Context, input *apigateway.GetMethodResponseInput) (*apigateway.MethodResponse, error)
	GetMethodResponseAsync(ctx workflow.Context, input *apigateway.GetMethodResponseInput) *ApigatewayGetMethodResponseResult

	GetModel(ctx workflow.Context, input *apigateway.GetModelInput) (*apigateway.Model, error)
	GetModelAsync(ctx workflow.Context, input *apigateway.GetModelInput) *ApigatewayGetModelResult

	GetModelTemplate(ctx workflow.Context, input *apigateway.GetModelTemplateInput) (*apigateway.GetModelTemplateOutput, error)
	GetModelTemplateAsync(ctx workflow.Context, input *apigateway.GetModelTemplateInput) *ApigatewayGetModelTemplateResult

	GetModels(ctx workflow.Context, input *apigateway.GetModelsInput) (*apigateway.GetModelsOutput, error)
	GetModelsAsync(ctx workflow.Context, input *apigateway.GetModelsInput) *ApigatewayGetModelsResult

	GetRequestValidator(ctx workflow.Context, input *apigateway.GetRequestValidatorInput) (*apigateway.UpdateRequestValidatorOutput, error)
	GetRequestValidatorAsync(ctx workflow.Context, input *apigateway.GetRequestValidatorInput) *ApigatewayGetRequestValidatorResult

	GetRequestValidators(ctx workflow.Context, input *apigateway.GetRequestValidatorsInput) (*apigateway.GetRequestValidatorsOutput, error)
	GetRequestValidatorsAsync(ctx workflow.Context, input *apigateway.GetRequestValidatorsInput) *ApigatewayGetRequestValidatorsResult

	GetResource(ctx workflow.Context, input *apigateway.GetResourceInput) (*apigateway.Resource, error)
	GetResourceAsync(ctx workflow.Context, input *apigateway.GetResourceInput) *ApigatewayGetResourceResult

	GetResources(ctx workflow.Context, input *apigateway.GetResourcesInput) (*apigateway.GetResourcesOutput, error)
	GetResourcesAsync(ctx workflow.Context, input *apigateway.GetResourcesInput) *ApigatewayGetResourcesResult

	GetRestApi(ctx workflow.Context, input *apigateway.GetRestApiInput) (*apigateway.RestApi, error)
	GetRestApiAsync(ctx workflow.Context, input *apigateway.GetRestApiInput) *ApigatewayGetRestApiResult

	GetRestApis(ctx workflow.Context, input *apigateway.GetRestApisInput) (*apigateway.GetRestApisOutput, error)
	GetRestApisAsync(ctx workflow.Context, input *apigateway.GetRestApisInput) *ApigatewayGetRestApisResult

	GetSdk(ctx workflow.Context, input *apigateway.GetSdkInput) (*apigateway.GetSdkOutput, error)
	GetSdkAsync(ctx workflow.Context, input *apigateway.GetSdkInput) *ApigatewayGetSdkResult

	GetSdkType(ctx workflow.Context, input *apigateway.GetSdkTypeInput) (*apigateway.SdkType, error)
	GetSdkTypeAsync(ctx workflow.Context, input *apigateway.GetSdkTypeInput) *ApigatewayGetSdkTypeResult

	GetSdkTypes(ctx workflow.Context, input *apigateway.GetSdkTypesInput) (*apigateway.GetSdkTypesOutput, error)
	GetSdkTypesAsync(ctx workflow.Context, input *apigateway.GetSdkTypesInput) *ApigatewayGetSdkTypesResult

	GetStage(ctx workflow.Context, input *apigateway.GetStageInput) (*apigateway.Stage, error)
	GetStageAsync(ctx workflow.Context, input *apigateway.GetStageInput) *ApigatewayGetStageResult

	GetStages(ctx workflow.Context, input *apigateway.GetStagesInput) (*apigateway.GetStagesOutput, error)
	GetStagesAsync(ctx workflow.Context, input *apigateway.GetStagesInput) *ApigatewayGetStagesResult

	GetTags(ctx workflow.Context, input *apigateway.GetTagsInput) (*apigateway.GetTagsOutput, error)
	GetTagsAsync(ctx workflow.Context, input *apigateway.GetTagsInput) *ApigatewayGetTagsResult

	GetUsage(ctx workflow.Context, input *apigateway.GetUsageInput) (*apigateway.Usage, error)
	GetUsageAsync(ctx workflow.Context, input *apigateway.GetUsageInput) *ApigatewayGetUsageResult

	GetUsagePlan(ctx workflow.Context, input *apigateway.GetUsagePlanInput) (*apigateway.UsagePlan, error)
	GetUsagePlanAsync(ctx workflow.Context, input *apigateway.GetUsagePlanInput) *ApigatewayGetUsagePlanResult

	GetUsagePlanKey(ctx workflow.Context, input *apigateway.GetUsagePlanKeyInput) (*apigateway.UsagePlanKey, error)
	GetUsagePlanKeyAsync(ctx workflow.Context, input *apigateway.GetUsagePlanKeyInput) *ApigatewayGetUsagePlanKeyResult

	GetUsagePlanKeys(ctx workflow.Context, input *apigateway.GetUsagePlanKeysInput) (*apigateway.GetUsagePlanKeysOutput, error)
	GetUsagePlanKeysAsync(ctx workflow.Context, input *apigateway.GetUsagePlanKeysInput) *ApigatewayGetUsagePlanKeysResult

	GetUsagePlans(ctx workflow.Context, input *apigateway.GetUsagePlansInput) (*apigateway.GetUsagePlansOutput, error)
	GetUsagePlansAsync(ctx workflow.Context, input *apigateway.GetUsagePlansInput) *ApigatewayGetUsagePlansResult

	GetVpcLink(ctx workflow.Context, input *apigateway.GetVpcLinkInput) (*apigateway.UpdateVpcLinkOutput, error)
	GetVpcLinkAsync(ctx workflow.Context, input *apigateway.GetVpcLinkInput) *ApigatewayGetVpcLinkResult

	GetVpcLinks(ctx workflow.Context, input *apigateway.GetVpcLinksInput) (*apigateway.GetVpcLinksOutput, error)
	GetVpcLinksAsync(ctx workflow.Context, input *apigateway.GetVpcLinksInput) *ApigatewayGetVpcLinksResult

	ImportApiKeys(ctx workflow.Context, input *apigateway.ImportApiKeysInput) (*apigateway.ImportApiKeysOutput, error)
	ImportApiKeysAsync(ctx workflow.Context, input *apigateway.ImportApiKeysInput) *ApigatewayImportApiKeysResult

	ImportDocumentationParts(ctx workflow.Context, input *apigateway.ImportDocumentationPartsInput) (*apigateway.ImportDocumentationPartsOutput, error)
	ImportDocumentationPartsAsync(ctx workflow.Context, input *apigateway.ImportDocumentationPartsInput) *ApigatewayImportDocumentationPartsResult

	ImportRestApi(ctx workflow.Context, input *apigateway.ImportRestApiInput) (*apigateway.RestApi, error)
	ImportRestApiAsync(ctx workflow.Context, input *apigateway.ImportRestApiInput) *ApigatewayImportRestApiResult

	PutGatewayResponse(ctx workflow.Context, input *apigateway.PutGatewayResponseInput) (*apigateway.UpdateGatewayResponseOutput, error)
	PutGatewayResponseAsync(ctx workflow.Context, input *apigateway.PutGatewayResponseInput) *ApigatewayPutGatewayResponseResult

	PutIntegration(ctx workflow.Context, input *apigateway.PutIntegrationInput) (*apigateway.Integration, error)
	PutIntegrationAsync(ctx workflow.Context, input *apigateway.PutIntegrationInput) *ApigatewayPutIntegrationResult

	PutIntegrationResponse(ctx workflow.Context, input *apigateway.PutIntegrationResponseInput) (*apigateway.IntegrationResponse, error)
	PutIntegrationResponseAsync(ctx workflow.Context, input *apigateway.PutIntegrationResponseInput) *ApigatewayPutIntegrationResponseResult

	PutMethod(ctx workflow.Context, input *apigateway.PutMethodInput) (*apigateway.Method, error)
	PutMethodAsync(ctx workflow.Context, input *apigateway.PutMethodInput) *ApigatewayPutMethodResult

	PutMethodResponse(ctx workflow.Context, input *apigateway.PutMethodResponseInput) (*apigateway.MethodResponse, error)
	PutMethodResponseAsync(ctx workflow.Context, input *apigateway.PutMethodResponseInput) *ApigatewayPutMethodResponseResult

	PutRestApi(ctx workflow.Context, input *apigateway.PutRestApiInput) (*apigateway.RestApi, error)
	PutRestApiAsync(ctx workflow.Context, input *apigateway.PutRestApiInput) *ApigatewayPutRestApiResult

	TagResource(ctx workflow.Context, input *apigateway.TagResourceInput) (*apigateway.TagResourceOutput, error)
	TagResourceAsync(ctx workflow.Context, input *apigateway.TagResourceInput) *ApigatewayTagResourceResult

	TestInvokeAuthorizer(ctx workflow.Context, input *apigateway.TestInvokeAuthorizerInput) (*apigateway.TestInvokeAuthorizerOutput, error)
	TestInvokeAuthorizerAsync(ctx workflow.Context, input *apigateway.TestInvokeAuthorizerInput) *ApigatewayTestInvokeAuthorizerResult

	TestInvokeMethod(ctx workflow.Context, input *apigateway.TestInvokeMethodInput) (*apigateway.TestInvokeMethodOutput, error)
	TestInvokeMethodAsync(ctx workflow.Context, input *apigateway.TestInvokeMethodInput) *ApigatewayTestInvokeMethodResult

	UntagResource(ctx workflow.Context, input *apigateway.UntagResourceInput) (*apigateway.UntagResourceOutput, error)
	UntagResourceAsync(ctx workflow.Context, input *apigateway.UntagResourceInput) *ApigatewayUntagResourceResult

	UpdateAccount(ctx workflow.Context, input *apigateway.UpdateAccountInput) (*apigateway.Account, error)
	UpdateAccountAsync(ctx workflow.Context, input *apigateway.UpdateAccountInput) *ApigatewayUpdateAccountResult

	UpdateApiKey(ctx workflow.Context, input *apigateway.UpdateApiKeyInput) (*apigateway.ApiKey, error)
	UpdateApiKeyAsync(ctx workflow.Context, input *apigateway.UpdateApiKeyInput) *ApigatewayUpdateApiKeyResult

	UpdateAuthorizer(ctx workflow.Context, input *apigateway.UpdateAuthorizerInput) (*apigateway.Authorizer, error)
	UpdateAuthorizerAsync(ctx workflow.Context, input *apigateway.UpdateAuthorizerInput) *ApigatewayUpdateAuthorizerResult

	UpdateBasePathMapping(ctx workflow.Context, input *apigateway.UpdateBasePathMappingInput) (*apigateway.BasePathMapping, error)
	UpdateBasePathMappingAsync(ctx workflow.Context, input *apigateway.UpdateBasePathMappingInput) *ApigatewayUpdateBasePathMappingResult

	UpdateClientCertificate(ctx workflow.Context, input *apigateway.UpdateClientCertificateInput) (*apigateway.ClientCertificate, error)
	UpdateClientCertificateAsync(ctx workflow.Context, input *apigateway.UpdateClientCertificateInput) *ApigatewayUpdateClientCertificateResult

	UpdateDeployment(ctx workflow.Context, input *apigateway.UpdateDeploymentInput) (*apigateway.Deployment, error)
	UpdateDeploymentAsync(ctx workflow.Context, input *apigateway.UpdateDeploymentInput) *ApigatewayUpdateDeploymentResult

	UpdateDocumentationPart(ctx workflow.Context, input *apigateway.UpdateDocumentationPartInput) (*apigateway.DocumentationPart, error)
	UpdateDocumentationPartAsync(ctx workflow.Context, input *apigateway.UpdateDocumentationPartInput) *ApigatewayUpdateDocumentationPartResult

	UpdateDocumentationVersion(ctx workflow.Context, input *apigateway.UpdateDocumentationVersionInput) (*apigateway.DocumentationVersion, error)
	UpdateDocumentationVersionAsync(ctx workflow.Context, input *apigateway.UpdateDocumentationVersionInput) *ApigatewayUpdateDocumentationVersionResult

	UpdateDomainName(ctx workflow.Context, input *apigateway.UpdateDomainNameInput) (*apigateway.DomainName, error)
	UpdateDomainNameAsync(ctx workflow.Context, input *apigateway.UpdateDomainNameInput) *ApigatewayUpdateDomainNameResult

	UpdateGatewayResponse(ctx workflow.Context, input *apigateway.UpdateGatewayResponseInput) (*apigateway.UpdateGatewayResponseOutput, error)
	UpdateGatewayResponseAsync(ctx workflow.Context, input *apigateway.UpdateGatewayResponseInput) *ApigatewayUpdateGatewayResponseResult

	UpdateIntegration(ctx workflow.Context, input *apigateway.UpdateIntegrationInput) (*apigateway.Integration, error)
	UpdateIntegrationAsync(ctx workflow.Context, input *apigateway.UpdateIntegrationInput) *ApigatewayUpdateIntegrationResult

	UpdateIntegrationResponse(ctx workflow.Context, input *apigateway.UpdateIntegrationResponseInput) (*apigateway.IntegrationResponse, error)
	UpdateIntegrationResponseAsync(ctx workflow.Context, input *apigateway.UpdateIntegrationResponseInput) *ApigatewayUpdateIntegrationResponseResult

	UpdateMethod(ctx workflow.Context, input *apigateway.UpdateMethodInput) (*apigateway.Method, error)
	UpdateMethodAsync(ctx workflow.Context, input *apigateway.UpdateMethodInput) *ApigatewayUpdateMethodResult

	UpdateMethodResponse(ctx workflow.Context, input *apigateway.UpdateMethodResponseInput) (*apigateway.MethodResponse, error)
	UpdateMethodResponseAsync(ctx workflow.Context, input *apigateway.UpdateMethodResponseInput) *ApigatewayUpdateMethodResponseResult

	UpdateModel(ctx workflow.Context, input *apigateway.UpdateModelInput) (*apigateway.Model, error)
	UpdateModelAsync(ctx workflow.Context, input *apigateway.UpdateModelInput) *ApigatewayUpdateModelResult

	UpdateRequestValidator(ctx workflow.Context, input *apigateway.UpdateRequestValidatorInput) (*apigateway.UpdateRequestValidatorOutput, error)
	UpdateRequestValidatorAsync(ctx workflow.Context, input *apigateway.UpdateRequestValidatorInput) *ApigatewayUpdateRequestValidatorResult

	UpdateResource(ctx workflow.Context, input *apigateway.UpdateResourceInput) (*apigateway.Resource, error)
	UpdateResourceAsync(ctx workflow.Context, input *apigateway.UpdateResourceInput) *ApigatewayUpdateResourceResult

	UpdateRestApi(ctx workflow.Context, input *apigateway.UpdateRestApiInput) (*apigateway.RestApi, error)
	UpdateRestApiAsync(ctx workflow.Context, input *apigateway.UpdateRestApiInput) *ApigatewayUpdateRestApiResult

	UpdateStage(ctx workflow.Context, input *apigateway.UpdateStageInput) (*apigateway.Stage, error)
	UpdateStageAsync(ctx workflow.Context, input *apigateway.UpdateStageInput) *ApigatewayUpdateStageResult

	UpdateUsage(ctx workflow.Context, input *apigateway.UpdateUsageInput) (*apigateway.Usage, error)
	UpdateUsageAsync(ctx workflow.Context, input *apigateway.UpdateUsageInput) *ApigatewayUpdateUsageResult

	UpdateUsagePlan(ctx workflow.Context, input *apigateway.UpdateUsagePlanInput) (*apigateway.UsagePlan, error)
	UpdateUsagePlanAsync(ctx workflow.Context, input *apigateway.UpdateUsagePlanInput) *ApigatewayUpdateUsagePlanResult

	UpdateVpcLink(ctx workflow.Context, input *apigateway.UpdateVpcLinkInput) (*apigateway.UpdateVpcLinkOutput, error)
	UpdateVpcLinkAsync(ctx workflow.Context, input *apigateway.UpdateVpcLinkInput) *ApigatewayUpdateVpcLinkResult
}

type ApigatewayCreateApiKeyResult struct {
	Result workflow.Future
}

func (r *ApigatewayCreateApiKeyResult) Get(ctx workflow.Context) (*apigateway.ApiKey, error) {
	var output apigateway.ApiKey
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ApigatewayCreateAuthorizerResult struct {
	Result workflow.Future
}

func (r *ApigatewayCreateAuthorizerResult) Get(ctx workflow.Context) (*apigateway.Authorizer, error) {
	var output apigateway.Authorizer
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ApigatewayCreateBasePathMappingResult struct {
	Result workflow.Future
}

func (r *ApigatewayCreateBasePathMappingResult) Get(ctx workflow.Context) (*apigateway.BasePathMapping, error) {
	var output apigateway.BasePathMapping
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ApigatewayCreateDeploymentResult struct {
	Result workflow.Future
}

func (r *ApigatewayCreateDeploymentResult) Get(ctx workflow.Context) (*apigateway.Deployment, error) {
	var output apigateway.Deployment
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ApigatewayCreateDocumentationPartResult struct {
	Result workflow.Future
}

func (r *ApigatewayCreateDocumentationPartResult) Get(ctx workflow.Context) (*apigateway.DocumentationPart, error) {
	var output apigateway.DocumentationPart
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ApigatewayCreateDocumentationVersionResult struct {
	Result workflow.Future
}

func (r *ApigatewayCreateDocumentationVersionResult) Get(ctx workflow.Context) (*apigateway.DocumentationVersion, error) {
	var output apigateway.DocumentationVersion
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ApigatewayCreateDomainNameResult struct {
	Result workflow.Future
}

func (r *ApigatewayCreateDomainNameResult) Get(ctx workflow.Context) (*apigateway.DomainName, error) {
	var output apigateway.DomainName
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ApigatewayCreateModelResult struct {
	Result workflow.Future
}

func (r *ApigatewayCreateModelResult) Get(ctx workflow.Context) (*apigateway.Model, error) {
	var output apigateway.Model
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ApigatewayCreateRequestValidatorResult struct {
	Result workflow.Future
}

func (r *ApigatewayCreateRequestValidatorResult) Get(ctx workflow.Context) (*apigateway.UpdateRequestValidatorOutput, error) {
	var output apigateway.UpdateRequestValidatorOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ApigatewayCreateResourceResult struct {
	Result workflow.Future
}

func (r *ApigatewayCreateResourceResult) Get(ctx workflow.Context) (*apigateway.Resource, error) {
	var output apigateway.Resource
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ApigatewayCreateRestApiResult struct {
	Result workflow.Future
}

func (r *ApigatewayCreateRestApiResult) Get(ctx workflow.Context) (*apigateway.RestApi, error) {
	var output apigateway.RestApi
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ApigatewayCreateStageResult struct {
	Result workflow.Future
}

func (r *ApigatewayCreateStageResult) Get(ctx workflow.Context) (*apigateway.Stage, error) {
	var output apigateway.Stage
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ApigatewayCreateUsagePlanResult struct {
	Result workflow.Future
}

func (r *ApigatewayCreateUsagePlanResult) Get(ctx workflow.Context) (*apigateway.UsagePlan, error) {
	var output apigateway.UsagePlan
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ApigatewayCreateUsagePlanKeyResult struct {
	Result workflow.Future
}

func (r *ApigatewayCreateUsagePlanKeyResult) Get(ctx workflow.Context) (*apigateway.UsagePlanKey, error) {
	var output apigateway.UsagePlanKey
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ApigatewayCreateVpcLinkResult struct {
	Result workflow.Future
}

func (r *ApigatewayCreateVpcLinkResult) Get(ctx workflow.Context) (*apigateway.UpdateVpcLinkOutput, error) {
	var output apigateway.UpdateVpcLinkOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ApigatewayDeleteApiKeyResult struct {
	Result workflow.Future
}

func (r *ApigatewayDeleteApiKeyResult) Get(ctx workflow.Context) (*apigateway.DeleteApiKeyOutput, error) {
	var output apigateway.DeleteApiKeyOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ApigatewayDeleteAuthorizerResult struct {
	Result workflow.Future
}

func (r *ApigatewayDeleteAuthorizerResult) Get(ctx workflow.Context) (*apigateway.DeleteAuthorizerOutput, error) {
	var output apigateway.DeleteAuthorizerOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ApigatewayDeleteBasePathMappingResult struct {
	Result workflow.Future
}

func (r *ApigatewayDeleteBasePathMappingResult) Get(ctx workflow.Context) (*apigateway.DeleteBasePathMappingOutput, error) {
	var output apigateway.DeleteBasePathMappingOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ApigatewayDeleteClientCertificateResult struct {
	Result workflow.Future
}

func (r *ApigatewayDeleteClientCertificateResult) Get(ctx workflow.Context) (*apigateway.DeleteClientCertificateOutput, error) {
	var output apigateway.DeleteClientCertificateOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ApigatewayDeleteDeploymentResult struct {
	Result workflow.Future
}

func (r *ApigatewayDeleteDeploymentResult) Get(ctx workflow.Context) (*apigateway.DeleteDeploymentOutput, error) {
	var output apigateway.DeleteDeploymentOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ApigatewayDeleteDocumentationPartResult struct {
	Result workflow.Future
}

func (r *ApigatewayDeleteDocumentationPartResult) Get(ctx workflow.Context) (*apigateway.DeleteDocumentationPartOutput, error) {
	var output apigateway.DeleteDocumentationPartOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ApigatewayDeleteDocumentationVersionResult struct {
	Result workflow.Future
}

func (r *ApigatewayDeleteDocumentationVersionResult) Get(ctx workflow.Context) (*apigateway.DeleteDocumentationVersionOutput, error) {
	var output apigateway.DeleteDocumentationVersionOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ApigatewayDeleteDomainNameResult struct {
	Result workflow.Future
}

func (r *ApigatewayDeleteDomainNameResult) Get(ctx workflow.Context) (*apigateway.DeleteDomainNameOutput, error) {
	var output apigateway.DeleteDomainNameOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ApigatewayDeleteGatewayResponseResult struct {
	Result workflow.Future
}

func (r *ApigatewayDeleteGatewayResponseResult) Get(ctx workflow.Context) (*apigateway.DeleteGatewayResponseOutput, error) {
	var output apigateway.DeleteGatewayResponseOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ApigatewayDeleteIntegrationResult struct {
	Result workflow.Future
}

func (r *ApigatewayDeleteIntegrationResult) Get(ctx workflow.Context) (*apigateway.DeleteIntegrationOutput, error) {
	var output apigateway.DeleteIntegrationOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ApigatewayDeleteIntegrationResponseResult struct {
	Result workflow.Future
}

func (r *ApigatewayDeleteIntegrationResponseResult) Get(ctx workflow.Context) (*apigateway.DeleteIntegrationResponseOutput, error) {
	var output apigateway.DeleteIntegrationResponseOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ApigatewayDeleteMethodResult struct {
	Result workflow.Future
}

func (r *ApigatewayDeleteMethodResult) Get(ctx workflow.Context) (*apigateway.DeleteMethodOutput, error) {
	var output apigateway.DeleteMethodOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ApigatewayDeleteMethodResponseResult struct {
	Result workflow.Future
}

func (r *ApigatewayDeleteMethodResponseResult) Get(ctx workflow.Context) (*apigateway.DeleteMethodResponseOutput, error) {
	var output apigateway.DeleteMethodResponseOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ApigatewayDeleteModelResult struct {
	Result workflow.Future
}

func (r *ApigatewayDeleteModelResult) Get(ctx workflow.Context) (*apigateway.DeleteModelOutput, error) {
	var output apigateway.DeleteModelOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ApigatewayDeleteRequestValidatorResult struct {
	Result workflow.Future
}

func (r *ApigatewayDeleteRequestValidatorResult) Get(ctx workflow.Context) (*apigateway.DeleteRequestValidatorOutput, error) {
	var output apigateway.DeleteRequestValidatorOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ApigatewayDeleteResourceResult struct {
	Result workflow.Future
}

func (r *ApigatewayDeleteResourceResult) Get(ctx workflow.Context) (*apigateway.DeleteResourceOutput, error) {
	var output apigateway.DeleteResourceOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ApigatewayDeleteRestApiResult struct {
	Result workflow.Future
}

func (r *ApigatewayDeleteRestApiResult) Get(ctx workflow.Context) (*apigateway.DeleteRestApiOutput, error) {
	var output apigateway.DeleteRestApiOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ApigatewayDeleteStageResult struct {
	Result workflow.Future
}

func (r *ApigatewayDeleteStageResult) Get(ctx workflow.Context) (*apigateway.DeleteStageOutput, error) {
	var output apigateway.DeleteStageOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ApigatewayDeleteUsagePlanResult struct {
	Result workflow.Future
}

func (r *ApigatewayDeleteUsagePlanResult) Get(ctx workflow.Context) (*apigateway.DeleteUsagePlanOutput, error) {
	var output apigateway.DeleteUsagePlanOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ApigatewayDeleteUsagePlanKeyResult struct {
	Result workflow.Future
}

func (r *ApigatewayDeleteUsagePlanKeyResult) Get(ctx workflow.Context) (*apigateway.DeleteUsagePlanKeyOutput, error) {
	var output apigateway.DeleteUsagePlanKeyOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ApigatewayDeleteVpcLinkResult struct {
	Result workflow.Future
}

func (r *ApigatewayDeleteVpcLinkResult) Get(ctx workflow.Context) (*apigateway.DeleteVpcLinkOutput, error) {
	var output apigateway.DeleteVpcLinkOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ApigatewayFlushStageAuthorizersCacheResult struct {
	Result workflow.Future
}

func (r *ApigatewayFlushStageAuthorizersCacheResult) Get(ctx workflow.Context) (*apigateway.FlushStageAuthorizersCacheOutput, error) {
	var output apigateway.FlushStageAuthorizersCacheOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ApigatewayFlushStageCacheResult struct {
	Result workflow.Future
}

func (r *ApigatewayFlushStageCacheResult) Get(ctx workflow.Context) (*apigateway.FlushStageCacheOutput, error) {
	var output apigateway.FlushStageCacheOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ApigatewayGenerateClientCertificateResult struct {
	Result workflow.Future
}

func (r *ApigatewayGenerateClientCertificateResult) Get(ctx workflow.Context) (*apigateway.ClientCertificate, error) {
	var output apigateway.ClientCertificate
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ApigatewayGetAccountResult struct {
	Result workflow.Future
}

func (r *ApigatewayGetAccountResult) Get(ctx workflow.Context) (*apigateway.Account, error) {
	var output apigateway.Account
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ApigatewayGetApiKeyResult struct {
	Result workflow.Future
}

func (r *ApigatewayGetApiKeyResult) Get(ctx workflow.Context) (*apigateway.ApiKey, error) {
	var output apigateway.ApiKey
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ApigatewayGetApiKeysResult struct {
	Result workflow.Future
}

func (r *ApigatewayGetApiKeysResult) Get(ctx workflow.Context) (*apigateway.GetApiKeysOutput, error) {
	var output apigateway.GetApiKeysOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ApigatewayGetAuthorizerResult struct {
	Result workflow.Future
}

func (r *ApigatewayGetAuthorizerResult) Get(ctx workflow.Context) (*apigateway.Authorizer, error) {
	var output apigateway.Authorizer
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ApigatewayGetAuthorizersResult struct {
	Result workflow.Future
}

func (r *ApigatewayGetAuthorizersResult) Get(ctx workflow.Context) (*apigateway.GetAuthorizersOutput, error) {
	var output apigateway.GetAuthorizersOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ApigatewayGetBasePathMappingResult struct {
	Result workflow.Future
}

func (r *ApigatewayGetBasePathMappingResult) Get(ctx workflow.Context) (*apigateway.BasePathMapping, error) {
	var output apigateway.BasePathMapping
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ApigatewayGetBasePathMappingsResult struct {
	Result workflow.Future
}

func (r *ApigatewayGetBasePathMappingsResult) Get(ctx workflow.Context) (*apigateway.GetBasePathMappingsOutput, error) {
	var output apigateway.GetBasePathMappingsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ApigatewayGetClientCertificateResult struct {
	Result workflow.Future
}

func (r *ApigatewayGetClientCertificateResult) Get(ctx workflow.Context) (*apigateway.ClientCertificate, error) {
	var output apigateway.ClientCertificate
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ApigatewayGetClientCertificatesResult struct {
	Result workflow.Future
}

func (r *ApigatewayGetClientCertificatesResult) Get(ctx workflow.Context) (*apigateway.GetClientCertificatesOutput, error) {
	var output apigateway.GetClientCertificatesOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ApigatewayGetDeploymentResult struct {
	Result workflow.Future
}

func (r *ApigatewayGetDeploymentResult) Get(ctx workflow.Context) (*apigateway.Deployment, error) {
	var output apigateway.Deployment
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ApigatewayGetDeploymentsResult struct {
	Result workflow.Future
}

func (r *ApigatewayGetDeploymentsResult) Get(ctx workflow.Context) (*apigateway.GetDeploymentsOutput, error) {
	var output apigateway.GetDeploymentsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ApigatewayGetDocumentationPartResult struct {
	Result workflow.Future
}

func (r *ApigatewayGetDocumentationPartResult) Get(ctx workflow.Context) (*apigateway.DocumentationPart, error) {
	var output apigateway.DocumentationPart
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ApigatewayGetDocumentationPartsResult struct {
	Result workflow.Future
}

func (r *ApigatewayGetDocumentationPartsResult) Get(ctx workflow.Context) (*apigateway.GetDocumentationPartsOutput, error) {
	var output apigateway.GetDocumentationPartsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ApigatewayGetDocumentationVersionResult struct {
	Result workflow.Future
}

func (r *ApigatewayGetDocumentationVersionResult) Get(ctx workflow.Context) (*apigateway.DocumentationVersion, error) {
	var output apigateway.DocumentationVersion
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ApigatewayGetDocumentationVersionsResult struct {
	Result workflow.Future
}

func (r *ApigatewayGetDocumentationVersionsResult) Get(ctx workflow.Context) (*apigateway.GetDocumentationVersionsOutput, error) {
	var output apigateway.GetDocumentationVersionsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ApigatewayGetDomainNameResult struct {
	Result workflow.Future
}

func (r *ApigatewayGetDomainNameResult) Get(ctx workflow.Context) (*apigateway.DomainName, error) {
	var output apigateway.DomainName
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ApigatewayGetDomainNamesResult struct {
	Result workflow.Future
}

func (r *ApigatewayGetDomainNamesResult) Get(ctx workflow.Context) (*apigateway.GetDomainNamesOutput, error) {
	var output apigateway.GetDomainNamesOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ApigatewayGetExportResult struct {
	Result workflow.Future
}

func (r *ApigatewayGetExportResult) Get(ctx workflow.Context) (*apigateway.GetExportOutput, error) {
	var output apigateway.GetExportOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ApigatewayGetGatewayResponseResult struct {
	Result workflow.Future
}

func (r *ApigatewayGetGatewayResponseResult) Get(ctx workflow.Context) (*apigateway.UpdateGatewayResponseOutput, error) {
	var output apigateway.UpdateGatewayResponseOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ApigatewayGetGatewayResponsesResult struct {
	Result workflow.Future
}

func (r *ApigatewayGetGatewayResponsesResult) Get(ctx workflow.Context) (*apigateway.GetGatewayResponsesOutput, error) {
	var output apigateway.GetGatewayResponsesOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ApigatewayGetIntegrationResult struct {
	Result workflow.Future
}

func (r *ApigatewayGetIntegrationResult) Get(ctx workflow.Context) (*apigateway.Integration, error) {
	var output apigateway.Integration
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ApigatewayGetIntegrationResponseResult struct {
	Result workflow.Future
}

func (r *ApigatewayGetIntegrationResponseResult) Get(ctx workflow.Context) (*apigateway.IntegrationResponse, error) {
	var output apigateway.IntegrationResponse
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ApigatewayGetMethodResult struct {
	Result workflow.Future
}

func (r *ApigatewayGetMethodResult) Get(ctx workflow.Context) (*apigateway.Method, error) {
	var output apigateway.Method
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ApigatewayGetMethodResponseResult struct {
	Result workflow.Future
}

func (r *ApigatewayGetMethodResponseResult) Get(ctx workflow.Context) (*apigateway.MethodResponse, error) {
	var output apigateway.MethodResponse
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ApigatewayGetModelResult struct {
	Result workflow.Future
}

func (r *ApigatewayGetModelResult) Get(ctx workflow.Context) (*apigateway.Model, error) {
	var output apigateway.Model
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ApigatewayGetModelTemplateResult struct {
	Result workflow.Future
}

func (r *ApigatewayGetModelTemplateResult) Get(ctx workflow.Context) (*apigateway.GetModelTemplateOutput, error) {
	var output apigateway.GetModelTemplateOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ApigatewayGetModelsResult struct {
	Result workflow.Future
}

func (r *ApigatewayGetModelsResult) Get(ctx workflow.Context) (*apigateway.GetModelsOutput, error) {
	var output apigateway.GetModelsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ApigatewayGetRequestValidatorResult struct {
	Result workflow.Future
}

func (r *ApigatewayGetRequestValidatorResult) Get(ctx workflow.Context) (*apigateway.UpdateRequestValidatorOutput, error) {
	var output apigateway.UpdateRequestValidatorOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ApigatewayGetRequestValidatorsResult struct {
	Result workflow.Future
}

func (r *ApigatewayGetRequestValidatorsResult) Get(ctx workflow.Context) (*apigateway.GetRequestValidatorsOutput, error) {
	var output apigateway.GetRequestValidatorsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ApigatewayGetResourceResult struct {
	Result workflow.Future
}

func (r *ApigatewayGetResourceResult) Get(ctx workflow.Context) (*apigateway.Resource, error) {
	var output apigateway.Resource
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ApigatewayGetResourcesResult struct {
	Result workflow.Future
}

func (r *ApigatewayGetResourcesResult) Get(ctx workflow.Context) (*apigateway.GetResourcesOutput, error) {
	var output apigateway.GetResourcesOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ApigatewayGetRestApiResult struct {
	Result workflow.Future
}

func (r *ApigatewayGetRestApiResult) Get(ctx workflow.Context) (*apigateway.RestApi, error) {
	var output apigateway.RestApi
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ApigatewayGetRestApisResult struct {
	Result workflow.Future
}

func (r *ApigatewayGetRestApisResult) Get(ctx workflow.Context) (*apigateway.GetRestApisOutput, error) {
	var output apigateway.GetRestApisOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ApigatewayGetSdkResult struct {
	Result workflow.Future
}

func (r *ApigatewayGetSdkResult) Get(ctx workflow.Context) (*apigateway.GetSdkOutput, error) {
	var output apigateway.GetSdkOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ApigatewayGetSdkTypeResult struct {
	Result workflow.Future
}

func (r *ApigatewayGetSdkTypeResult) Get(ctx workflow.Context) (*apigateway.SdkType, error) {
	var output apigateway.SdkType
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ApigatewayGetSdkTypesResult struct {
	Result workflow.Future
}

func (r *ApigatewayGetSdkTypesResult) Get(ctx workflow.Context) (*apigateway.GetSdkTypesOutput, error) {
	var output apigateway.GetSdkTypesOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ApigatewayGetStageResult struct {
	Result workflow.Future
}

func (r *ApigatewayGetStageResult) Get(ctx workflow.Context) (*apigateway.Stage, error) {
	var output apigateway.Stage
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ApigatewayGetStagesResult struct {
	Result workflow.Future
}

func (r *ApigatewayGetStagesResult) Get(ctx workflow.Context) (*apigateway.GetStagesOutput, error) {
	var output apigateway.GetStagesOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ApigatewayGetTagsResult struct {
	Result workflow.Future
}

func (r *ApigatewayGetTagsResult) Get(ctx workflow.Context) (*apigateway.GetTagsOutput, error) {
	var output apigateway.GetTagsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ApigatewayGetUsageResult struct {
	Result workflow.Future
}

func (r *ApigatewayGetUsageResult) Get(ctx workflow.Context) (*apigateway.Usage, error) {
	var output apigateway.Usage
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ApigatewayGetUsagePlanResult struct {
	Result workflow.Future
}

func (r *ApigatewayGetUsagePlanResult) Get(ctx workflow.Context) (*apigateway.UsagePlan, error) {
	var output apigateway.UsagePlan
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ApigatewayGetUsagePlanKeyResult struct {
	Result workflow.Future
}

func (r *ApigatewayGetUsagePlanKeyResult) Get(ctx workflow.Context) (*apigateway.UsagePlanKey, error) {
	var output apigateway.UsagePlanKey
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ApigatewayGetUsagePlanKeysResult struct {
	Result workflow.Future
}

func (r *ApigatewayGetUsagePlanKeysResult) Get(ctx workflow.Context) (*apigateway.GetUsagePlanKeysOutput, error) {
	var output apigateway.GetUsagePlanKeysOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ApigatewayGetUsagePlansResult struct {
	Result workflow.Future
}

func (r *ApigatewayGetUsagePlansResult) Get(ctx workflow.Context) (*apigateway.GetUsagePlansOutput, error) {
	var output apigateway.GetUsagePlansOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ApigatewayGetVpcLinkResult struct {
	Result workflow.Future
}

func (r *ApigatewayGetVpcLinkResult) Get(ctx workflow.Context) (*apigateway.UpdateVpcLinkOutput, error) {
	var output apigateway.UpdateVpcLinkOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ApigatewayGetVpcLinksResult struct {
	Result workflow.Future
}

func (r *ApigatewayGetVpcLinksResult) Get(ctx workflow.Context) (*apigateway.GetVpcLinksOutput, error) {
	var output apigateway.GetVpcLinksOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ApigatewayImportApiKeysResult struct {
	Result workflow.Future
}

func (r *ApigatewayImportApiKeysResult) Get(ctx workflow.Context) (*apigateway.ImportApiKeysOutput, error) {
	var output apigateway.ImportApiKeysOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ApigatewayImportDocumentationPartsResult struct {
	Result workflow.Future
}

func (r *ApigatewayImportDocumentationPartsResult) Get(ctx workflow.Context) (*apigateway.ImportDocumentationPartsOutput, error) {
	var output apigateway.ImportDocumentationPartsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ApigatewayImportRestApiResult struct {
	Result workflow.Future
}

func (r *ApigatewayImportRestApiResult) Get(ctx workflow.Context) (*apigateway.RestApi, error) {
	var output apigateway.RestApi
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ApigatewayPutGatewayResponseResult struct {
	Result workflow.Future
}

func (r *ApigatewayPutGatewayResponseResult) Get(ctx workflow.Context) (*apigateway.UpdateGatewayResponseOutput, error) {
	var output apigateway.UpdateGatewayResponseOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ApigatewayPutIntegrationResult struct {
	Result workflow.Future
}

func (r *ApigatewayPutIntegrationResult) Get(ctx workflow.Context) (*apigateway.Integration, error) {
	var output apigateway.Integration
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ApigatewayPutIntegrationResponseResult struct {
	Result workflow.Future
}

func (r *ApigatewayPutIntegrationResponseResult) Get(ctx workflow.Context) (*apigateway.IntegrationResponse, error) {
	var output apigateway.IntegrationResponse
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ApigatewayPutMethodResult struct {
	Result workflow.Future
}

func (r *ApigatewayPutMethodResult) Get(ctx workflow.Context) (*apigateway.Method, error) {
	var output apigateway.Method
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ApigatewayPutMethodResponseResult struct {
	Result workflow.Future
}

func (r *ApigatewayPutMethodResponseResult) Get(ctx workflow.Context) (*apigateway.MethodResponse, error) {
	var output apigateway.MethodResponse
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ApigatewayPutRestApiResult struct {
	Result workflow.Future
}

func (r *ApigatewayPutRestApiResult) Get(ctx workflow.Context) (*apigateway.RestApi, error) {
	var output apigateway.RestApi
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ApigatewayTagResourceResult struct {
	Result workflow.Future
}

func (r *ApigatewayTagResourceResult) Get(ctx workflow.Context) (*apigateway.TagResourceOutput, error) {
	var output apigateway.TagResourceOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ApigatewayTestInvokeAuthorizerResult struct {
	Result workflow.Future
}

func (r *ApigatewayTestInvokeAuthorizerResult) Get(ctx workflow.Context) (*apigateway.TestInvokeAuthorizerOutput, error) {
	var output apigateway.TestInvokeAuthorizerOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ApigatewayTestInvokeMethodResult struct {
	Result workflow.Future
}

func (r *ApigatewayTestInvokeMethodResult) Get(ctx workflow.Context) (*apigateway.TestInvokeMethodOutput, error) {
	var output apigateway.TestInvokeMethodOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ApigatewayUntagResourceResult struct {
	Result workflow.Future
}

func (r *ApigatewayUntagResourceResult) Get(ctx workflow.Context) (*apigateway.UntagResourceOutput, error) {
	var output apigateway.UntagResourceOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ApigatewayUpdateAccountResult struct {
	Result workflow.Future
}

func (r *ApigatewayUpdateAccountResult) Get(ctx workflow.Context) (*apigateway.Account, error) {
	var output apigateway.Account
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ApigatewayUpdateApiKeyResult struct {
	Result workflow.Future
}

func (r *ApigatewayUpdateApiKeyResult) Get(ctx workflow.Context) (*apigateway.ApiKey, error) {
	var output apigateway.ApiKey
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ApigatewayUpdateAuthorizerResult struct {
	Result workflow.Future
}

func (r *ApigatewayUpdateAuthorizerResult) Get(ctx workflow.Context) (*apigateway.Authorizer, error) {
	var output apigateway.Authorizer
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ApigatewayUpdateBasePathMappingResult struct {
	Result workflow.Future
}

func (r *ApigatewayUpdateBasePathMappingResult) Get(ctx workflow.Context) (*apigateway.BasePathMapping, error) {
	var output apigateway.BasePathMapping
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ApigatewayUpdateClientCertificateResult struct {
	Result workflow.Future
}

func (r *ApigatewayUpdateClientCertificateResult) Get(ctx workflow.Context) (*apigateway.ClientCertificate, error) {
	var output apigateway.ClientCertificate
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ApigatewayUpdateDeploymentResult struct {
	Result workflow.Future
}

func (r *ApigatewayUpdateDeploymentResult) Get(ctx workflow.Context) (*apigateway.Deployment, error) {
	var output apigateway.Deployment
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ApigatewayUpdateDocumentationPartResult struct {
	Result workflow.Future
}

func (r *ApigatewayUpdateDocumentationPartResult) Get(ctx workflow.Context) (*apigateway.DocumentationPart, error) {
	var output apigateway.DocumentationPart
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ApigatewayUpdateDocumentationVersionResult struct {
	Result workflow.Future
}

func (r *ApigatewayUpdateDocumentationVersionResult) Get(ctx workflow.Context) (*apigateway.DocumentationVersion, error) {
	var output apigateway.DocumentationVersion
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ApigatewayUpdateDomainNameResult struct {
	Result workflow.Future
}

func (r *ApigatewayUpdateDomainNameResult) Get(ctx workflow.Context) (*apigateway.DomainName, error) {
	var output apigateway.DomainName
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ApigatewayUpdateGatewayResponseResult struct {
	Result workflow.Future
}

func (r *ApigatewayUpdateGatewayResponseResult) Get(ctx workflow.Context) (*apigateway.UpdateGatewayResponseOutput, error) {
	var output apigateway.UpdateGatewayResponseOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ApigatewayUpdateIntegrationResult struct {
	Result workflow.Future
}

func (r *ApigatewayUpdateIntegrationResult) Get(ctx workflow.Context) (*apigateway.Integration, error) {
	var output apigateway.Integration
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ApigatewayUpdateIntegrationResponseResult struct {
	Result workflow.Future
}

func (r *ApigatewayUpdateIntegrationResponseResult) Get(ctx workflow.Context) (*apigateway.IntegrationResponse, error) {
	var output apigateway.IntegrationResponse
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ApigatewayUpdateMethodResult struct {
	Result workflow.Future
}

func (r *ApigatewayUpdateMethodResult) Get(ctx workflow.Context) (*apigateway.Method, error) {
	var output apigateway.Method
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ApigatewayUpdateMethodResponseResult struct {
	Result workflow.Future
}

func (r *ApigatewayUpdateMethodResponseResult) Get(ctx workflow.Context) (*apigateway.MethodResponse, error) {
	var output apigateway.MethodResponse
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ApigatewayUpdateModelResult struct {
	Result workflow.Future
}

func (r *ApigatewayUpdateModelResult) Get(ctx workflow.Context) (*apigateway.Model, error) {
	var output apigateway.Model
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ApigatewayUpdateRequestValidatorResult struct {
	Result workflow.Future
}

func (r *ApigatewayUpdateRequestValidatorResult) Get(ctx workflow.Context) (*apigateway.UpdateRequestValidatorOutput, error) {
	var output apigateway.UpdateRequestValidatorOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ApigatewayUpdateResourceResult struct {
	Result workflow.Future
}

func (r *ApigatewayUpdateResourceResult) Get(ctx workflow.Context) (*apigateway.Resource, error) {
	var output apigateway.Resource
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ApigatewayUpdateRestApiResult struct {
	Result workflow.Future
}

func (r *ApigatewayUpdateRestApiResult) Get(ctx workflow.Context) (*apigateway.RestApi, error) {
	var output apigateway.RestApi
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ApigatewayUpdateStageResult struct {
	Result workflow.Future
}

func (r *ApigatewayUpdateStageResult) Get(ctx workflow.Context) (*apigateway.Stage, error) {
	var output apigateway.Stage
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ApigatewayUpdateUsageResult struct {
	Result workflow.Future
}

func (r *ApigatewayUpdateUsageResult) Get(ctx workflow.Context) (*apigateway.Usage, error) {
	var output apigateway.Usage
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ApigatewayUpdateUsagePlanResult struct {
	Result workflow.Future
}

func (r *ApigatewayUpdateUsagePlanResult) Get(ctx workflow.Context) (*apigateway.UsagePlan, error) {
	var output apigateway.UsagePlan
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ApigatewayUpdateVpcLinkResult struct {
	Result workflow.Future
}

func (r *ApigatewayUpdateVpcLinkResult) Get(ctx workflow.Context) (*apigateway.UpdateVpcLinkOutput, error) {
	var output apigateway.UpdateVpcLinkOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type APIGatewayStub struct {
	activities awsactivities.APIGatewayActivities
}

func NewAPIGatewayStub() APIGatewayClient {
	return &APIGatewayStub{}
}

func (a *APIGatewayStub) CreateApiKey(ctx workflow.Context, input *apigateway.CreateApiKeyInput) (*apigateway.ApiKey, error) {
	var output apigateway.ApiKey
	err := workflow.ExecuteActivity(ctx, a.activities.CreateApiKey, input).Get(ctx, &output)
	return &output, err
}

func (a *APIGatewayStub) CreateApiKeyAsync(ctx workflow.Context, input *apigateway.CreateApiKeyInput) *ApigatewayCreateApiKeyResult {
	future := workflow.ExecuteActivity(ctx, a.activities.CreateApiKey, input)
	return &ApigatewayCreateApiKeyResult{Result: future}
}

func (a *APIGatewayStub) CreateAuthorizer(ctx workflow.Context, input *apigateway.CreateAuthorizerInput) (*apigateway.Authorizer, error) {
	var output apigateway.Authorizer
	err := workflow.ExecuteActivity(ctx, a.activities.CreateAuthorizer, input).Get(ctx, &output)
	return &output, err
}

func (a *APIGatewayStub) CreateAuthorizerAsync(ctx workflow.Context, input *apigateway.CreateAuthorizerInput) *ApigatewayCreateAuthorizerResult {
	future := workflow.ExecuteActivity(ctx, a.activities.CreateAuthorizer, input)
	return &ApigatewayCreateAuthorizerResult{Result: future}
}

func (a *APIGatewayStub) CreateBasePathMapping(ctx workflow.Context, input *apigateway.CreateBasePathMappingInput) (*apigateway.BasePathMapping, error) {
	var output apigateway.BasePathMapping
	err := workflow.ExecuteActivity(ctx, a.activities.CreateBasePathMapping, input).Get(ctx, &output)
	return &output, err
}

func (a *APIGatewayStub) CreateBasePathMappingAsync(ctx workflow.Context, input *apigateway.CreateBasePathMappingInput) *ApigatewayCreateBasePathMappingResult {
	future := workflow.ExecuteActivity(ctx, a.activities.CreateBasePathMapping, input)
	return &ApigatewayCreateBasePathMappingResult{Result: future}
}

func (a *APIGatewayStub) CreateDeployment(ctx workflow.Context, input *apigateway.CreateDeploymentInput) (*apigateway.Deployment, error) {
	var output apigateway.Deployment
	err := workflow.ExecuteActivity(ctx, a.activities.CreateDeployment, input).Get(ctx, &output)
	return &output, err
}

func (a *APIGatewayStub) CreateDeploymentAsync(ctx workflow.Context, input *apigateway.CreateDeploymentInput) *ApigatewayCreateDeploymentResult {
	future := workflow.ExecuteActivity(ctx, a.activities.CreateDeployment, input)
	return &ApigatewayCreateDeploymentResult{Result: future}
}

func (a *APIGatewayStub) CreateDocumentationPart(ctx workflow.Context, input *apigateway.CreateDocumentationPartInput) (*apigateway.DocumentationPart, error) {
	var output apigateway.DocumentationPart
	err := workflow.ExecuteActivity(ctx, a.activities.CreateDocumentationPart, input).Get(ctx, &output)
	return &output, err
}

func (a *APIGatewayStub) CreateDocumentationPartAsync(ctx workflow.Context, input *apigateway.CreateDocumentationPartInput) *ApigatewayCreateDocumentationPartResult {
	future := workflow.ExecuteActivity(ctx, a.activities.CreateDocumentationPart, input)
	return &ApigatewayCreateDocumentationPartResult{Result: future}
}

func (a *APIGatewayStub) CreateDocumentationVersion(ctx workflow.Context, input *apigateway.CreateDocumentationVersionInput) (*apigateway.DocumentationVersion, error) {
	var output apigateway.DocumentationVersion
	err := workflow.ExecuteActivity(ctx, a.activities.CreateDocumentationVersion, input).Get(ctx, &output)
	return &output, err
}

func (a *APIGatewayStub) CreateDocumentationVersionAsync(ctx workflow.Context, input *apigateway.CreateDocumentationVersionInput) *ApigatewayCreateDocumentationVersionResult {
	future := workflow.ExecuteActivity(ctx, a.activities.CreateDocumentationVersion, input)
	return &ApigatewayCreateDocumentationVersionResult{Result: future}
}

func (a *APIGatewayStub) CreateDomainName(ctx workflow.Context, input *apigateway.CreateDomainNameInput) (*apigateway.DomainName, error) {
	var output apigateway.DomainName
	err := workflow.ExecuteActivity(ctx, a.activities.CreateDomainName, input).Get(ctx, &output)
	return &output, err
}

func (a *APIGatewayStub) CreateDomainNameAsync(ctx workflow.Context, input *apigateway.CreateDomainNameInput) *ApigatewayCreateDomainNameResult {
	future := workflow.ExecuteActivity(ctx, a.activities.CreateDomainName, input)
	return &ApigatewayCreateDomainNameResult{Result: future}
}

func (a *APIGatewayStub) CreateModel(ctx workflow.Context, input *apigateway.CreateModelInput) (*apigateway.Model, error) {
	var output apigateway.Model
	err := workflow.ExecuteActivity(ctx, a.activities.CreateModel, input).Get(ctx, &output)
	return &output, err
}

func (a *APIGatewayStub) CreateModelAsync(ctx workflow.Context, input *apigateway.CreateModelInput) *ApigatewayCreateModelResult {
	future := workflow.ExecuteActivity(ctx, a.activities.CreateModel, input)
	return &ApigatewayCreateModelResult{Result: future}
}

func (a *APIGatewayStub) CreateRequestValidator(ctx workflow.Context, input *apigateway.CreateRequestValidatorInput) (*apigateway.UpdateRequestValidatorOutput, error) {
	var output apigateway.UpdateRequestValidatorOutput
	err := workflow.ExecuteActivity(ctx, a.activities.CreateRequestValidator, input).Get(ctx, &output)
	return &output, err
}

func (a *APIGatewayStub) CreateRequestValidatorAsync(ctx workflow.Context, input *apigateway.CreateRequestValidatorInput) *ApigatewayCreateRequestValidatorResult {
	future := workflow.ExecuteActivity(ctx, a.activities.CreateRequestValidator, input)
	return &ApigatewayCreateRequestValidatorResult{Result: future}
}

func (a *APIGatewayStub) CreateResource(ctx workflow.Context, input *apigateway.CreateResourceInput) (*apigateway.Resource, error) {
	var output apigateway.Resource
	err := workflow.ExecuteActivity(ctx, a.activities.CreateResource, input).Get(ctx, &output)
	return &output, err
}

func (a *APIGatewayStub) CreateResourceAsync(ctx workflow.Context, input *apigateway.CreateResourceInput) *ApigatewayCreateResourceResult {
	future := workflow.ExecuteActivity(ctx, a.activities.CreateResource, input)
	return &ApigatewayCreateResourceResult{Result: future}
}

func (a *APIGatewayStub) CreateRestApi(ctx workflow.Context, input *apigateway.CreateRestApiInput) (*apigateway.RestApi, error) {
	var output apigateway.RestApi
	err := workflow.ExecuteActivity(ctx, a.activities.CreateRestApi, input).Get(ctx, &output)
	return &output, err
}

func (a *APIGatewayStub) CreateRestApiAsync(ctx workflow.Context, input *apigateway.CreateRestApiInput) *ApigatewayCreateRestApiResult {
	future := workflow.ExecuteActivity(ctx, a.activities.CreateRestApi, input)
	return &ApigatewayCreateRestApiResult{Result: future}
}

func (a *APIGatewayStub) CreateStage(ctx workflow.Context, input *apigateway.CreateStageInput) (*apigateway.Stage, error) {
	var output apigateway.Stage
	err := workflow.ExecuteActivity(ctx, a.activities.CreateStage, input).Get(ctx, &output)
	return &output, err
}

func (a *APIGatewayStub) CreateStageAsync(ctx workflow.Context, input *apigateway.CreateStageInput) *ApigatewayCreateStageResult {
	future := workflow.ExecuteActivity(ctx, a.activities.CreateStage, input)
	return &ApigatewayCreateStageResult{Result: future}
}

func (a *APIGatewayStub) CreateUsagePlan(ctx workflow.Context, input *apigateway.CreateUsagePlanInput) (*apigateway.UsagePlan, error) {
	var output apigateway.UsagePlan
	err := workflow.ExecuteActivity(ctx, a.activities.CreateUsagePlan, input).Get(ctx, &output)
	return &output, err
}

func (a *APIGatewayStub) CreateUsagePlanAsync(ctx workflow.Context, input *apigateway.CreateUsagePlanInput) *ApigatewayCreateUsagePlanResult {
	future := workflow.ExecuteActivity(ctx, a.activities.CreateUsagePlan, input)
	return &ApigatewayCreateUsagePlanResult{Result: future}
}

func (a *APIGatewayStub) CreateUsagePlanKey(ctx workflow.Context, input *apigateway.CreateUsagePlanKeyInput) (*apigateway.UsagePlanKey, error) {
	var output apigateway.UsagePlanKey
	err := workflow.ExecuteActivity(ctx, a.activities.CreateUsagePlanKey, input).Get(ctx, &output)
	return &output, err
}

func (a *APIGatewayStub) CreateUsagePlanKeyAsync(ctx workflow.Context, input *apigateway.CreateUsagePlanKeyInput) *ApigatewayCreateUsagePlanKeyResult {
	future := workflow.ExecuteActivity(ctx, a.activities.CreateUsagePlanKey, input)
	return &ApigatewayCreateUsagePlanKeyResult{Result: future}
}

func (a *APIGatewayStub) CreateVpcLink(ctx workflow.Context, input *apigateway.CreateVpcLinkInput) (*apigateway.UpdateVpcLinkOutput, error) {
	var output apigateway.UpdateVpcLinkOutput
	err := workflow.ExecuteActivity(ctx, a.activities.CreateVpcLink, input).Get(ctx, &output)
	return &output, err
}

func (a *APIGatewayStub) CreateVpcLinkAsync(ctx workflow.Context, input *apigateway.CreateVpcLinkInput) *ApigatewayCreateVpcLinkResult {
	future := workflow.ExecuteActivity(ctx, a.activities.CreateVpcLink, input)
	return &ApigatewayCreateVpcLinkResult{Result: future}
}

func (a *APIGatewayStub) DeleteApiKey(ctx workflow.Context, input *apigateway.DeleteApiKeyInput) (*apigateway.DeleteApiKeyOutput, error) {
	var output apigateway.DeleteApiKeyOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeleteApiKey, input).Get(ctx, &output)
	return &output, err
}

func (a *APIGatewayStub) DeleteApiKeyAsync(ctx workflow.Context, input *apigateway.DeleteApiKeyInput) *ApigatewayDeleteApiKeyResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeleteApiKey, input)
	return &ApigatewayDeleteApiKeyResult{Result: future}
}

func (a *APIGatewayStub) DeleteAuthorizer(ctx workflow.Context, input *apigateway.DeleteAuthorizerInput) (*apigateway.DeleteAuthorizerOutput, error) {
	var output apigateway.DeleteAuthorizerOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeleteAuthorizer, input).Get(ctx, &output)
	return &output, err
}

func (a *APIGatewayStub) DeleteAuthorizerAsync(ctx workflow.Context, input *apigateway.DeleteAuthorizerInput) *ApigatewayDeleteAuthorizerResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeleteAuthorizer, input)
	return &ApigatewayDeleteAuthorizerResult{Result: future}
}

func (a *APIGatewayStub) DeleteBasePathMapping(ctx workflow.Context, input *apigateway.DeleteBasePathMappingInput) (*apigateway.DeleteBasePathMappingOutput, error) {
	var output apigateway.DeleteBasePathMappingOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeleteBasePathMapping, input).Get(ctx, &output)
	return &output, err
}

func (a *APIGatewayStub) DeleteBasePathMappingAsync(ctx workflow.Context, input *apigateway.DeleteBasePathMappingInput) *ApigatewayDeleteBasePathMappingResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeleteBasePathMapping, input)
	return &ApigatewayDeleteBasePathMappingResult{Result: future}
}

func (a *APIGatewayStub) DeleteClientCertificate(ctx workflow.Context, input *apigateway.DeleteClientCertificateInput) (*apigateway.DeleteClientCertificateOutput, error) {
	var output apigateway.DeleteClientCertificateOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeleteClientCertificate, input).Get(ctx, &output)
	return &output, err
}

func (a *APIGatewayStub) DeleteClientCertificateAsync(ctx workflow.Context, input *apigateway.DeleteClientCertificateInput) *ApigatewayDeleteClientCertificateResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeleteClientCertificate, input)
	return &ApigatewayDeleteClientCertificateResult{Result: future}
}

func (a *APIGatewayStub) DeleteDeployment(ctx workflow.Context, input *apigateway.DeleteDeploymentInput) (*apigateway.DeleteDeploymentOutput, error) {
	var output apigateway.DeleteDeploymentOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeleteDeployment, input).Get(ctx, &output)
	return &output, err
}

func (a *APIGatewayStub) DeleteDeploymentAsync(ctx workflow.Context, input *apigateway.DeleteDeploymentInput) *ApigatewayDeleteDeploymentResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeleteDeployment, input)
	return &ApigatewayDeleteDeploymentResult{Result: future}
}

func (a *APIGatewayStub) DeleteDocumentationPart(ctx workflow.Context, input *apigateway.DeleteDocumentationPartInput) (*apigateway.DeleteDocumentationPartOutput, error) {
	var output apigateway.DeleteDocumentationPartOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeleteDocumentationPart, input).Get(ctx, &output)
	return &output, err
}

func (a *APIGatewayStub) DeleteDocumentationPartAsync(ctx workflow.Context, input *apigateway.DeleteDocumentationPartInput) *ApigatewayDeleteDocumentationPartResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeleteDocumentationPart, input)
	return &ApigatewayDeleteDocumentationPartResult{Result: future}
}

func (a *APIGatewayStub) DeleteDocumentationVersion(ctx workflow.Context, input *apigateway.DeleteDocumentationVersionInput) (*apigateway.DeleteDocumentationVersionOutput, error) {
	var output apigateway.DeleteDocumentationVersionOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeleteDocumentationVersion, input).Get(ctx, &output)
	return &output, err
}

func (a *APIGatewayStub) DeleteDocumentationVersionAsync(ctx workflow.Context, input *apigateway.DeleteDocumentationVersionInput) *ApigatewayDeleteDocumentationVersionResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeleteDocumentationVersion, input)
	return &ApigatewayDeleteDocumentationVersionResult{Result: future}
}

func (a *APIGatewayStub) DeleteDomainName(ctx workflow.Context, input *apigateway.DeleteDomainNameInput) (*apigateway.DeleteDomainNameOutput, error) {
	var output apigateway.DeleteDomainNameOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeleteDomainName, input).Get(ctx, &output)
	return &output, err
}

func (a *APIGatewayStub) DeleteDomainNameAsync(ctx workflow.Context, input *apigateway.DeleteDomainNameInput) *ApigatewayDeleteDomainNameResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeleteDomainName, input)
	return &ApigatewayDeleteDomainNameResult{Result: future}
}

func (a *APIGatewayStub) DeleteGatewayResponse(ctx workflow.Context, input *apigateway.DeleteGatewayResponseInput) (*apigateway.DeleteGatewayResponseOutput, error) {
	var output apigateway.DeleteGatewayResponseOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeleteGatewayResponse, input).Get(ctx, &output)
	return &output, err
}

func (a *APIGatewayStub) DeleteGatewayResponseAsync(ctx workflow.Context, input *apigateway.DeleteGatewayResponseInput) *ApigatewayDeleteGatewayResponseResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeleteGatewayResponse, input)
	return &ApigatewayDeleteGatewayResponseResult{Result: future}
}

func (a *APIGatewayStub) DeleteIntegration(ctx workflow.Context, input *apigateway.DeleteIntegrationInput) (*apigateway.DeleteIntegrationOutput, error) {
	var output apigateway.DeleteIntegrationOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeleteIntegration, input).Get(ctx, &output)
	return &output, err
}

func (a *APIGatewayStub) DeleteIntegrationAsync(ctx workflow.Context, input *apigateway.DeleteIntegrationInput) *ApigatewayDeleteIntegrationResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeleteIntegration, input)
	return &ApigatewayDeleteIntegrationResult{Result: future}
}

func (a *APIGatewayStub) DeleteIntegrationResponse(ctx workflow.Context, input *apigateway.DeleteIntegrationResponseInput) (*apigateway.DeleteIntegrationResponseOutput, error) {
	var output apigateway.DeleteIntegrationResponseOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeleteIntegrationResponse, input).Get(ctx, &output)
	return &output, err
}

func (a *APIGatewayStub) DeleteIntegrationResponseAsync(ctx workflow.Context, input *apigateway.DeleteIntegrationResponseInput) *ApigatewayDeleteIntegrationResponseResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeleteIntegrationResponse, input)
	return &ApigatewayDeleteIntegrationResponseResult{Result: future}
}

func (a *APIGatewayStub) DeleteMethod(ctx workflow.Context, input *apigateway.DeleteMethodInput) (*apigateway.DeleteMethodOutput, error) {
	var output apigateway.DeleteMethodOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeleteMethod, input).Get(ctx, &output)
	return &output, err
}

func (a *APIGatewayStub) DeleteMethodAsync(ctx workflow.Context, input *apigateway.DeleteMethodInput) *ApigatewayDeleteMethodResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeleteMethod, input)
	return &ApigatewayDeleteMethodResult{Result: future}
}

func (a *APIGatewayStub) DeleteMethodResponse(ctx workflow.Context, input *apigateway.DeleteMethodResponseInput) (*apigateway.DeleteMethodResponseOutput, error) {
	var output apigateway.DeleteMethodResponseOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeleteMethodResponse, input).Get(ctx, &output)
	return &output, err
}

func (a *APIGatewayStub) DeleteMethodResponseAsync(ctx workflow.Context, input *apigateway.DeleteMethodResponseInput) *ApigatewayDeleteMethodResponseResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeleteMethodResponse, input)
	return &ApigatewayDeleteMethodResponseResult{Result: future}
}

func (a *APIGatewayStub) DeleteModel(ctx workflow.Context, input *apigateway.DeleteModelInput) (*apigateway.DeleteModelOutput, error) {
	var output apigateway.DeleteModelOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeleteModel, input).Get(ctx, &output)
	return &output, err
}

func (a *APIGatewayStub) DeleteModelAsync(ctx workflow.Context, input *apigateway.DeleteModelInput) *ApigatewayDeleteModelResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeleteModel, input)
	return &ApigatewayDeleteModelResult{Result: future}
}

func (a *APIGatewayStub) DeleteRequestValidator(ctx workflow.Context, input *apigateway.DeleteRequestValidatorInput) (*apigateway.DeleteRequestValidatorOutput, error) {
	var output apigateway.DeleteRequestValidatorOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeleteRequestValidator, input).Get(ctx, &output)
	return &output, err
}

func (a *APIGatewayStub) DeleteRequestValidatorAsync(ctx workflow.Context, input *apigateway.DeleteRequestValidatorInput) *ApigatewayDeleteRequestValidatorResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeleteRequestValidator, input)
	return &ApigatewayDeleteRequestValidatorResult{Result: future}
}

func (a *APIGatewayStub) DeleteResource(ctx workflow.Context, input *apigateway.DeleteResourceInput) (*apigateway.DeleteResourceOutput, error) {
	var output apigateway.DeleteResourceOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeleteResource, input).Get(ctx, &output)
	return &output, err
}

func (a *APIGatewayStub) DeleteResourceAsync(ctx workflow.Context, input *apigateway.DeleteResourceInput) *ApigatewayDeleteResourceResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeleteResource, input)
	return &ApigatewayDeleteResourceResult{Result: future}
}

func (a *APIGatewayStub) DeleteRestApi(ctx workflow.Context, input *apigateway.DeleteRestApiInput) (*apigateway.DeleteRestApiOutput, error) {
	var output apigateway.DeleteRestApiOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeleteRestApi, input).Get(ctx, &output)
	return &output, err
}

func (a *APIGatewayStub) DeleteRestApiAsync(ctx workflow.Context, input *apigateway.DeleteRestApiInput) *ApigatewayDeleteRestApiResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeleteRestApi, input)
	return &ApigatewayDeleteRestApiResult{Result: future}
}

func (a *APIGatewayStub) DeleteStage(ctx workflow.Context, input *apigateway.DeleteStageInput) (*apigateway.DeleteStageOutput, error) {
	var output apigateway.DeleteStageOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeleteStage, input).Get(ctx, &output)
	return &output, err
}

func (a *APIGatewayStub) DeleteStageAsync(ctx workflow.Context, input *apigateway.DeleteStageInput) *ApigatewayDeleteStageResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeleteStage, input)
	return &ApigatewayDeleteStageResult{Result: future}
}

func (a *APIGatewayStub) DeleteUsagePlan(ctx workflow.Context, input *apigateway.DeleteUsagePlanInput) (*apigateway.DeleteUsagePlanOutput, error) {
	var output apigateway.DeleteUsagePlanOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeleteUsagePlan, input).Get(ctx, &output)
	return &output, err
}

func (a *APIGatewayStub) DeleteUsagePlanAsync(ctx workflow.Context, input *apigateway.DeleteUsagePlanInput) *ApigatewayDeleteUsagePlanResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeleteUsagePlan, input)
	return &ApigatewayDeleteUsagePlanResult{Result: future}
}

func (a *APIGatewayStub) DeleteUsagePlanKey(ctx workflow.Context, input *apigateway.DeleteUsagePlanKeyInput) (*apigateway.DeleteUsagePlanKeyOutput, error) {
	var output apigateway.DeleteUsagePlanKeyOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeleteUsagePlanKey, input).Get(ctx, &output)
	return &output, err
}

func (a *APIGatewayStub) DeleteUsagePlanKeyAsync(ctx workflow.Context, input *apigateway.DeleteUsagePlanKeyInput) *ApigatewayDeleteUsagePlanKeyResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeleteUsagePlanKey, input)
	return &ApigatewayDeleteUsagePlanKeyResult{Result: future}
}

func (a *APIGatewayStub) DeleteVpcLink(ctx workflow.Context, input *apigateway.DeleteVpcLinkInput) (*apigateway.DeleteVpcLinkOutput, error) {
	var output apigateway.DeleteVpcLinkOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeleteVpcLink, input).Get(ctx, &output)
	return &output, err
}

func (a *APIGatewayStub) DeleteVpcLinkAsync(ctx workflow.Context, input *apigateway.DeleteVpcLinkInput) *ApigatewayDeleteVpcLinkResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeleteVpcLink, input)
	return &ApigatewayDeleteVpcLinkResult{Result: future}
}

func (a *APIGatewayStub) FlushStageAuthorizersCache(ctx workflow.Context, input *apigateway.FlushStageAuthorizersCacheInput) (*apigateway.FlushStageAuthorizersCacheOutput, error) {
	var output apigateway.FlushStageAuthorizersCacheOutput
	err := workflow.ExecuteActivity(ctx, a.activities.FlushStageAuthorizersCache, input).Get(ctx, &output)
	return &output, err
}

func (a *APIGatewayStub) FlushStageAuthorizersCacheAsync(ctx workflow.Context, input *apigateway.FlushStageAuthorizersCacheInput) *ApigatewayFlushStageAuthorizersCacheResult {
	future := workflow.ExecuteActivity(ctx, a.activities.FlushStageAuthorizersCache, input)
	return &ApigatewayFlushStageAuthorizersCacheResult{Result: future}
}

func (a *APIGatewayStub) FlushStageCache(ctx workflow.Context, input *apigateway.FlushStageCacheInput) (*apigateway.FlushStageCacheOutput, error) {
	var output apigateway.FlushStageCacheOutput
	err := workflow.ExecuteActivity(ctx, a.activities.FlushStageCache, input).Get(ctx, &output)
	return &output, err
}

func (a *APIGatewayStub) FlushStageCacheAsync(ctx workflow.Context, input *apigateway.FlushStageCacheInput) *ApigatewayFlushStageCacheResult {
	future := workflow.ExecuteActivity(ctx, a.activities.FlushStageCache, input)
	return &ApigatewayFlushStageCacheResult{Result: future}
}

func (a *APIGatewayStub) GenerateClientCertificate(ctx workflow.Context, input *apigateway.GenerateClientCertificateInput) (*apigateway.ClientCertificate, error) {
	var output apigateway.ClientCertificate
	err := workflow.ExecuteActivity(ctx, a.activities.GenerateClientCertificate, input).Get(ctx, &output)
	return &output, err
}

func (a *APIGatewayStub) GenerateClientCertificateAsync(ctx workflow.Context, input *apigateway.GenerateClientCertificateInput) *ApigatewayGenerateClientCertificateResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GenerateClientCertificate, input)
	return &ApigatewayGenerateClientCertificateResult{Result: future}
}

func (a *APIGatewayStub) GetAccount(ctx workflow.Context, input *apigateway.GetAccountInput) (*apigateway.Account, error) {
	var output apigateway.Account
	err := workflow.ExecuteActivity(ctx, a.activities.GetAccount, input).Get(ctx, &output)
	return &output, err
}

func (a *APIGatewayStub) GetAccountAsync(ctx workflow.Context, input *apigateway.GetAccountInput) *ApigatewayGetAccountResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetAccount, input)
	return &ApigatewayGetAccountResult{Result: future}
}

func (a *APIGatewayStub) GetApiKey(ctx workflow.Context, input *apigateway.GetApiKeyInput) (*apigateway.ApiKey, error) {
	var output apigateway.ApiKey
	err := workflow.ExecuteActivity(ctx, a.activities.GetApiKey, input).Get(ctx, &output)
	return &output, err
}

func (a *APIGatewayStub) GetApiKeyAsync(ctx workflow.Context, input *apigateway.GetApiKeyInput) *ApigatewayGetApiKeyResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetApiKey, input)
	return &ApigatewayGetApiKeyResult{Result: future}
}

func (a *APIGatewayStub) GetApiKeys(ctx workflow.Context, input *apigateway.GetApiKeysInput) (*apigateway.GetApiKeysOutput, error) {
	var output apigateway.GetApiKeysOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetApiKeys, input).Get(ctx, &output)
	return &output, err
}

func (a *APIGatewayStub) GetApiKeysAsync(ctx workflow.Context, input *apigateway.GetApiKeysInput) *ApigatewayGetApiKeysResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetApiKeys, input)
	return &ApigatewayGetApiKeysResult{Result: future}
}

func (a *APIGatewayStub) GetAuthorizer(ctx workflow.Context, input *apigateway.GetAuthorizerInput) (*apigateway.Authorizer, error) {
	var output apigateway.Authorizer
	err := workflow.ExecuteActivity(ctx, a.activities.GetAuthorizer, input).Get(ctx, &output)
	return &output, err
}

func (a *APIGatewayStub) GetAuthorizerAsync(ctx workflow.Context, input *apigateway.GetAuthorizerInput) *ApigatewayGetAuthorizerResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetAuthorizer, input)
	return &ApigatewayGetAuthorizerResult{Result: future}
}

func (a *APIGatewayStub) GetAuthorizers(ctx workflow.Context, input *apigateway.GetAuthorizersInput) (*apigateway.GetAuthorizersOutput, error) {
	var output apigateway.GetAuthorizersOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetAuthorizers, input).Get(ctx, &output)
	return &output, err
}

func (a *APIGatewayStub) GetAuthorizersAsync(ctx workflow.Context, input *apigateway.GetAuthorizersInput) *ApigatewayGetAuthorizersResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetAuthorizers, input)
	return &ApigatewayGetAuthorizersResult{Result: future}
}

func (a *APIGatewayStub) GetBasePathMapping(ctx workflow.Context, input *apigateway.GetBasePathMappingInput) (*apigateway.BasePathMapping, error) {
	var output apigateway.BasePathMapping
	err := workflow.ExecuteActivity(ctx, a.activities.GetBasePathMapping, input).Get(ctx, &output)
	return &output, err
}

func (a *APIGatewayStub) GetBasePathMappingAsync(ctx workflow.Context, input *apigateway.GetBasePathMappingInput) *ApigatewayGetBasePathMappingResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetBasePathMapping, input)
	return &ApigatewayGetBasePathMappingResult{Result: future}
}

func (a *APIGatewayStub) GetBasePathMappings(ctx workflow.Context, input *apigateway.GetBasePathMappingsInput) (*apigateway.GetBasePathMappingsOutput, error) {
	var output apigateway.GetBasePathMappingsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetBasePathMappings, input).Get(ctx, &output)
	return &output, err
}

func (a *APIGatewayStub) GetBasePathMappingsAsync(ctx workflow.Context, input *apigateway.GetBasePathMappingsInput) *ApigatewayGetBasePathMappingsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetBasePathMappings, input)
	return &ApigatewayGetBasePathMappingsResult{Result: future}
}

func (a *APIGatewayStub) GetClientCertificate(ctx workflow.Context, input *apigateway.GetClientCertificateInput) (*apigateway.ClientCertificate, error) {
	var output apigateway.ClientCertificate
	err := workflow.ExecuteActivity(ctx, a.activities.GetClientCertificate, input).Get(ctx, &output)
	return &output, err
}

func (a *APIGatewayStub) GetClientCertificateAsync(ctx workflow.Context, input *apigateway.GetClientCertificateInput) *ApigatewayGetClientCertificateResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetClientCertificate, input)
	return &ApigatewayGetClientCertificateResult{Result: future}
}

func (a *APIGatewayStub) GetClientCertificates(ctx workflow.Context, input *apigateway.GetClientCertificatesInput) (*apigateway.GetClientCertificatesOutput, error) {
	var output apigateway.GetClientCertificatesOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetClientCertificates, input).Get(ctx, &output)
	return &output, err
}

func (a *APIGatewayStub) GetClientCertificatesAsync(ctx workflow.Context, input *apigateway.GetClientCertificatesInput) *ApigatewayGetClientCertificatesResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetClientCertificates, input)
	return &ApigatewayGetClientCertificatesResult{Result: future}
}

func (a *APIGatewayStub) GetDeployment(ctx workflow.Context, input *apigateway.GetDeploymentInput) (*apigateway.Deployment, error) {
	var output apigateway.Deployment
	err := workflow.ExecuteActivity(ctx, a.activities.GetDeployment, input).Get(ctx, &output)
	return &output, err
}

func (a *APIGatewayStub) GetDeploymentAsync(ctx workflow.Context, input *apigateway.GetDeploymentInput) *ApigatewayGetDeploymentResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetDeployment, input)
	return &ApigatewayGetDeploymentResult{Result: future}
}

func (a *APIGatewayStub) GetDeployments(ctx workflow.Context, input *apigateway.GetDeploymentsInput) (*apigateway.GetDeploymentsOutput, error) {
	var output apigateway.GetDeploymentsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetDeployments, input).Get(ctx, &output)
	return &output, err
}

func (a *APIGatewayStub) GetDeploymentsAsync(ctx workflow.Context, input *apigateway.GetDeploymentsInput) *ApigatewayGetDeploymentsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetDeployments, input)
	return &ApigatewayGetDeploymentsResult{Result: future}
}

func (a *APIGatewayStub) GetDocumentationPart(ctx workflow.Context, input *apigateway.GetDocumentationPartInput) (*apigateway.DocumentationPart, error) {
	var output apigateway.DocumentationPart
	err := workflow.ExecuteActivity(ctx, a.activities.GetDocumentationPart, input).Get(ctx, &output)
	return &output, err
}

func (a *APIGatewayStub) GetDocumentationPartAsync(ctx workflow.Context, input *apigateway.GetDocumentationPartInput) *ApigatewayGetDocumentationPartResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetDocumentationPart, input)
	return &ApigatewayGetDocumentationPartResult{Result: future}
}

func (a *APIGatewayStub) GetDocumentationParts(ctx workflow.Context, input *apigateway.GetDocumentationPartsInput) (*apigateway.GetDocumentationPartsOutput, error) {
	var output apigateway.GetDocumentationPartsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetDocumentationParts, input).Get(ctx, &output)
	return &output, err
}

func (a *APIGatewayStub) GetDocumentationPartsAsync(ctx workflow.Context, input *apigateway.GetDocumentationPartsInput) *ApigatewayGetDocumentationPartsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetDocumentationParts, input)
	return &ApigatewayGetDocumentationPartsResult{Result: future}
}

func (a *APIGatewayStub) GetDocumentationVersion(ctx workflow.Context, input *apigateway.GetDocumentationVersionInput) (*apigateway.DocumentationVersion, error) {
	var output apigateway.DocumentationVersion
	err := workflow.ExecuteActivity(ctx, a.activities.GetDocumentationVersion, input).Get(ctx, &output)
	return &output, err
}

func (a *APIGatewayStub) GetDocumentationVersionAsync(ctx workflow.Context, input *apigateway.GetDocumentationVersionInput) *ApigatewayGetDocumentationVersionResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetDocumentationVersion, input)
	return &ApigatewayGetDocumentationVersionResult{Result: future}
}

func (a *APIGatewayStub) GetDocumentationVersions(ctx workflow.Context, input *apigateway.GetDocumentationVersionsInput) (*apigateway.GetDocumentationVersionsOutput, error) {
	var output apigateway.GetDocumentationVersionsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetDocumentationVersions, input).Get(ctx, &output)
	return &output, err
}

func (a *APIGatewayStub) GetDocumentationVersionsAsync(ctx workflow.Context, input *apigateway.GetDocumentationVersionsInput) *ApigatewayGetDocumentationVersionsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetDocumentationVersions, input)
	return &ApigatewayGetDocumentationVersionsResult{Result: future}
}

func (a *APIGatewayStub) GetDomainName(ctx workflow.Context, input *apigateway.GetDomainNameInput) (*apigateway.DomainName, error) {
	var output apigateway.DomainName
	err := workflow.ExecuteActivity(ctx, a.activities.GetDomainName, input).Get(ctx, &output)
	return &output, err
}

func (a *APIGatewayStub) GetDomainNameAsync(ctx workflow.Context, input *apigateway.GetDomainNameInput) *ApigatewayGetDomainNameResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetDomainName, input)
	return &ApigatewayGetDomainNameResult{Result: future}
}

func (a *APIGatewayStub) GetDomainNames(ctx workflow.Context, input *apigateway.GetDomainNamesInput) (*apigateway.GetDomainNamesOutput, error) {
	var output apigateway.GetDomainNamesOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetDomainNames, input).Get(ctx, &output)
	return &output, err
}

func (a *APIGatewayStub) GetDomainNamesAsync(ctx workflow.Context, input *apigateway.GetDomainNamesInput) *ApigatewayGetDomainNamesResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetDomainNames, input)
	return &ApigatewayGetDomainNamesResult{Result: future}
}

func (a *APIGatewayStub) GetExport(ctx workflow.Context, input *apigateway.GetExportInput) (*apigateway.GetExportOutput, error) {
	var output apigateway.GetExportOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetExport, input).Get(ctx, &output)
	return &output, err
}

func (a *APIGatewayStub) GetExportAsync(ctx workflow.Context, input *apigateway.GetExportInput) *ApigatewayGetExportResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetExport, input)
	return &ApigatewayGetExportResult{Result: future}
}

func (a *APIGatewayStub) GetGatewayResponse(ctx workflow.Context, input *apigateway.GetGatewayResponseInput) (*apigateway.UpdateGatewayResponseOutput, error) {
	var output apigateway.UpdateGatewayResponseOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetGatewayResponse, input).Get(ctx, &output)
	return &output, err
}

func (a *APIGatewayStub) GetGatewayResponseAsync(ctx workflow.Context, input *apigateway.GetGatewayResponseInput) *ApigatewayGetGatewayResponseResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetGatewayResponse, input)
	return &ApigatewayGetGatewayResponseResult{Result: future}
}

func (a *APIGatewayStub) GetGatewayResponses(ctx workflow.Context, input *apigateway.GetGatewayResponsesInput) (*apigateway.GetGatewayResponsesOutput, error) {
	var output apigateway.GetGatewayResponsesOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetGatewayResponses, input).Get(ctx, &output)
	return &output, err
}

func (a *APIGatewayStub) GetGatewayResponsesAsync(ctx workflow.Context, input *apigateway.GetGatewayResponsesInput) *ApigatewayGetGatewayResponsesResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetGatewayResponses, input)
	return &ApigatewayGetGatewayResponsesResult{Result: future}
}

func (a *APIGatewayStub) GetIntegration(ctx workflow.Context, input *apigateway.GetIntegrationInput) (*apigateway.Integration, error) {
	var output apigateway.Integration
	err := workflow.ExecuteActivity(ctx, a.activities.GetIntegration, input).Get(ctx, &output)
	return &output, err
}

func (a *APIGatewayStub) GetIntegrationAsync(ctx workflow.Context, input *apigateway.GetIntegrationInput) *ApigatewayGetIntegrationResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetIntegration, input)
	return &ApigatewayGetIntegrationResult{Result: future}
}

func (a *APIGatewayStub) GetIntegrationResponse(ctx workflow.Context, input *apigateway.GetIntegrationResponseInput) (*apigateway.IntegrationResponse, error) {
	var output apigateway.IntegrationResponse
	err := workflow.ExecuteActivity(ctx, a.activities.GetIntegrationResponse, input).Get(ctx, &output)
	return &output, err
}

func (a *APIGatewayStub) GetIntegrationResponseAsync(ctx workflow.Context, input *apigateway.GetIntegrationResponseInput) *ApigatewayGetIntegrationResponseResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetIntegrationResponse, input)
	return &ApigatewayGetIntegrationResponseResult{Result: future}
}

func (a *APIGatewayStub) GetMethod(ctx workflow.Context, input *apigateway.GetMethodInput) (*apigateway.Method, error) {
	var output apigateway.Method
	err := workflow.ExecuteActivity(ctx, a.activities.GetMethod, input).Get(ctx, &output)
	return &output, err
}

func (a *APIGatewayStub) GetMethodAsync(ctx workflow.Context, input *apigateway.GetMethodInput) *ApigatewayGetMethodResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetMethod, input)
	return &ApigatewayGetMethodResult{Result: future}
}

func (a *APIGatewayStub) GetMethodResponse(ctx workflow.Context, input *apigateway.GetMethodResponseInput) (*apigateway.MethodResponse, error) {
	var output apigateway.MethodResponse
	err := workflow.ExecuteActivity(ctx, a.activities.GetMethodResponse, input).Get(ctx, &output)
	return &output, err
}

func (a *APIGatewayStub) GetMethodResponseAsync(ctx workflow.Context, input *apigateway.GetMethodResponseInput) *ApigatewayGetMethodResponseResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetMethodResponse, input)
	return &ApigatewayGetMethodResponseResult{Result: future}
}

func (a *APIGatewayStub) GetModel(ctx workflow.Context, input *apigateway.GetModelInput) (*apigateway.Model, error) {
	var output apigateway.Model
	err := workflow.ExecuteActivity(ctx, a.activities.GetModel, input).Get(ctx, &output)
	return &output, err
}

func (a *APIGatewayStub) GetModelAsync(ctx workflow.Context, input *apigateway.GetModelInput) *ApigatewayGetModelResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetModel, input)
	return &ApigatewayGetModelResult{Result: future}
}

func (a *APIGatewayStub) GetModelTemplate(ctx workflow.Context, input *apigateway.GetModelTemplateInput) (*apigateway.GetModelTemplateOutput, error) {
	var output apigateway.GetModelTemplateOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetModelTemplate, input).Get(ctx, &output)
	return &output, err
}

func (a *APIGatewayStub) GetModelTemplateAsync(ctx workflow.Context, input *apigateway.GetModelTemplateInput) *ApigatewayGetModelTemplateResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetModelTemplate, input)
	return &ApigatewayGetModelTemplateResult{Result: future}
}

func (a *APIGatewayStub) GetModels(ctx workflow.Context, input *apigateway.GetModelsInput) (*apigateway.GetModelsOutput, error) {
	var output apigateway.GetModelsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetModels, input).Get(ctx, &output)
	return &output, err
}

func (a *APIGatewayStub) GetModelsAsync(ctx workflow.Context, input *apigateway.GetModelsInput) *ApigatewayGetModelsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetModels, input)
	return &ApigatewayGetModelsResult{Result: future}
}

func (a *APIGatewayStub) GetRequestValidator(ctx workflow.Context, input *apigateway.GetRequestValidatorInput) (*apigateway.UpdateRequestValidatorOutput, error) {
	var output apigateway.UpdateRequestValidatorOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetRequestValidator, input).Get(ctx, &output)
	return &output, err
}

func (a *APIGatewayStub) GetRequestValidatorAsync(ctx workflow.Context, input *apigateway.GetRequestValidatorInput) *ApigatewayGetRequestValidatorResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetRequestValidator, input)
	return &ApigatewayGetRequestValidatorResult{Result: future}
}

func (a *APIGatewayStub) GetRequestValidators(ctx workflow.Context, input *apigateway.GetRequestValidatorsInput) (*apigateway.GetRequestValidatorsOutput, error) {
	var output apigateway.GetRequestValidatorsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetRequestValidators, input).Get(ctx, &output)
	return &output, err
}

func (a *APIGatewayStub) GetRequestValidatorsAsync(ctx workflow.Context, input *apigateway.GetRequestValidatorsInput) *ApigatewayGetRequestValidatorsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetRequestValidators, input)
	return &ApigatewayGetRequestValidatorsResult{Result: future}
}

func (a *APIGatewayStub) GetResource(ctx workflow.Context, input *apigateway.GetResourceInput) (*apigateway.Resource, error) {
	var output apigateway.Resource
	err := workflow.ExecuteActivity(ctx, a.activities.GetResource, input).Get(ctx, &output)
	return &output, err
}

func (a *APIGatewayStub) GetResourceAsync(ctx workflow.Context, input *apigateway.GetResourceInput) *ApigatewayGetResourceResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetResource, input)
	return &ApigatewayGetResourceResult{Result: future}
}

func (a *APIGatewayStub) GetResources(ctx workflow.Context, input *apigateway.GetResourcesInput) (*apigateway.GetResourcesOutput, error) {
	var output apigateway.GetResourcesOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetResources, input).Get(ctx, &output)
	return &output, err
}

func (a *APIGatewayStub) GetResourcesAsync(ctx workflow.Context, input *apigateway.GetResourcesInput) *ApigatewayGetResourcesResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetResources, input)
	return &ApigatewayGetResourcesResult{Result: future}
}

func (a *APIGatewayStub) GetRestApi(ctx workflow.Context, input *apigateway.GetRestApiInput) (*apigateway.RestApi, error) {
	var output apigateway.RestApi
	err := workflow.ExecuteActivity(ctx, a.activities.GetRestApi, input).Get(ctx, &output)
	return &output, err
}

func (a *APIGatewayStub) GetRestApiAsync(ctx workflow.Context, input *apigateway.GetRestApiInput) *ApigatewayGetRestApiResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetRestApi, input)
	return &ApigatewayGetRestApiResult{Result: future}
}

func (a *APIGatewayStub) GetRestApis(ctx workflow.Context, input *apigateway.GetRestApisInput) (*apigateway.GetRestApisOutput, error) {
	var output apigateway.GetRestApisOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetRestApis, input).Get(ctx, &output)
	return &output, err
}

func (a *APIGatewayStub) GetRestApisAsync(ctx workflow.Context, input *apigateway.GetRestApisInput) *ApigatewayGetRestApisResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetRestApis, input)
	return &ApigatewayGetRestApisResult{Result: future}
}

func (a *APIGatewayStub) GetSdk(ctx workflow.Context, input *apigateway.GetSdkInput) (*apigateway.GetSdkOutput, error) {
	var output apigateway.GetSdkOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetSdk, input).Get(ctx, &output)
	return &output, err
}

func (a *APIGatewayStub) GetSdkAsync(ctx workflow.Context, input *apigateway.GetSdkInput) *ApigatewayGetSdkResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetSdk, input)
	return &ApigatewayGetSdkResult{Result: future}
}

func (a *APIGatewayStub) GetSdkType(ctx workflow.Context, input *apigateway.GetSdkTypeInput) (*apigateway.SdkType, error) {
	var output apigateway.SdkType
	err := workflow.ExecuteActivity(ctx, a.activities.GetSdkType, input).Get(ctx, &output)
	return &output, err
}

func (a *APIGatewayStub) GetSdkTypeAsync(ctx workflow.Context, input *apigateway.GetSdkTypeInput) *ApigatewayGetSdkTypeResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetSdkType, input)
	return &ApigatewayGetSdkTypeResult{Result: future}
}

func (a *APIGatewayStub) GetSdkTypes(ctx workflow.Context, input *apigateway.GetSdkTypesInput) (*apigateway.GetSdkTypesOutput, error) {
	var output apigateway.GetSdkTypesOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetSdkTypes, input).Get(ctx, &output)
	return &output, err
}

func (a *APIGatewayStub) GetSdkTypesAsync(ctx workflow.Context, input *apigateway.GetSdkTypesInput) *ApigatewayGetSdkTypesResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetSdkTypes, input)
	return &ApigatewayGetSdkTypesResult{Result: future}
}

func (a *APIGatewayStub) GetStage(ctx workflow.Context, input *apigateway.GetStageInput) (*apigateway.Stage, error) {
	var output apigateway.Stage
	err := workflow.ExecuteActivity(ctx, a.activities.GetStage, input).Get(ctx, &output)
	return &output, err
}

func (a *APIGatewayStub) GetStageAsync(ctx workflow.Context, input *apigateway.GetStageInput) *ApigatewayGetStageResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetStage, input)
	return &ApigatewayGetStageResult{Result: future}
}

func (a *APIGatewayStub) GetStages(ctx workflow.Context, input *apigateway.GetStagesInput) (*apigateway.GetStagesOutput, error) {
	var output apigateway.GetStagesOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetStages, input).Get(ctx, &output)
	return &output, err
}

func (a *APIGatewayStub) GetStagesAsync(ctx workflow.Context, input *apigateway.GetStagesInput) *ApigatewayGetStagesResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetStages, input)
	return &ApigatewayGetStagesResult{Result: future}
}

func (a *APIGatewayStub) GetTags(ctx workflow.Context, input *apigateway.GetTagsInput) (*apigateway.GetTagsOutput, error) {
	var output apigateway.GetTagsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetTags, input).Get(ctx, &output)
	return &output, err
}

func (a *APIGatewayStub) GetTagsAsync(ctx workflow.Context, input *apigateway.GetTagsInput) *ApigatewayGetTagsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetTags, input)
	return &ApigatewayGetTagsResult{Result: future}
}

func (a *APIGatewayStub) GetUsage(ctx workflow.Context, input *apigateway.GetUsageInput) (*apigateway.Usage, error) {
	var output apigateway.Usage
	err := workflow.ExecuteActivity(ctx, a.activities.GetUsage, input).Get(ctx, &output)
	return &output, err
}

func (a *APIGatewayStub) GetUsageAsync(ctx workflow.Context, input *apigateway.GetUsageInput) *ApigatewayGetUsageResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetUsage, input)
	return &ApigatewayGetUsageResult{Result: future}
}

func (a *APIGatewayStub) GetUsagePlan(ctx workflow.Context, input *apigateway.GetUsagePlanInput) (*apigateway.UsagePlan, error) {
	var output apigateway.UsagePlan
	err := workflow.ExecuteActivity(ctx, a.activities.GetUsagePlan, input).Get(ctx, &output)
	return &output, err
}

func (a *APIGatewayStub) GetUsagePlanAsync(ctx workflow.Context, input *apigateway.GetUsagePlanInput) *ApigatewayGetUsagePlanResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetUsagePlan, input)
	return &ApigatewayGetUsagePlanResult{Result: future}
}

func (a *APIGatewayStub) GetUsagePlanKey(ctx workflow.Context, input *apigateway.GetUsagePlanKeyInput) (*apigateway.UsagePlanKey, error) {
	var output apigateway.UsagePlanKey
	err := workflow.ExecuteActivity(ctx, a.activities.GetUsagePlanKey, input).Get(ctx, &output)
	return &output, err
}

func (a *APIGatewayStub) GetUsagePlanKeyAsync(ctx workflow.Context, input *apigateway.GetUsagePlanKeyInput) *ApigatewayGetUsagePlanKeyResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetUsagePlanKey, input)
	return &ApigatewayGetUsagePlanKeyResult{Result: future}
}

func (a *APIGatewayStub) GetUsagePlanKeys(ctx workflow.Context, input *apigateway.GetUsagePlanKeysInput) (*apigateway.GetUsagePlanKeysOutput, error) {
	var output apigateway.GetUsagePlanKeysOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetUsagePlanKeys, input).Get(ctx, &output)
	return &output, err
}

func (a *APIGatewayStub) GetUsagePlanKeysAsync(ctx workflow.Context, input *apigateway.GetUsagePlanKeysInput) *ApigatewayGetUsagePlanKeysResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetUsagePlanKeys, input)
	return &ApigatewayGetUsagePlanKeysResult{Result: future}
}

func (a *APIGatewayStub) GetUsagePlans(ctx workflow.Context, input *apigateway.GetUsagePlansInput) (*apigateway.GetUsagePlansOutput, error) {
	var output apigateway.GetUsagePlansOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetUsagePlans, input).Get(ctx, &output)
	return &output, err
}

func (a *APIGatewayStub) GetUsagePlansAsync(ctx workflow.Context, input *apigateway.GetUsagePlansInput) *ApigatewayGetUsagePlansResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetUsagePlans, input)
	return &ApigatewayGetUsagePlansResult{Result: future}
}

func (a *APIGatewayStub) GetVpcLink(ctx workflow.Context, input *apigateway.GetVpcLinkInput) (*apigateway.UpdateVpcLinkOutput, error) {
	var output apigateway.UpdateVpcLinkOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetVpcLink, input).Get(ctx, &output)
	return &output, err
}

func (a *APIGatewayStub) GetVpcLinkAsync(ctx workflow.Context, input *apigateway.GetVpcLinkInput) *ApigatewayGetVpcLinkResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetVpcLink, input)
	return &ApigatewayGetVpcLinkResult{Result: future}
}

func (a *APIGatewayStub) GetVpcLinks(ctx workflow.Context, input *apigateway.GetVpcLinksInput) (*apigateway.GetVpcLinksOutput, error) {
	var output apigateway.GetVpcLinksOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetVpcLinks, input).Get(ctx, &output)
	return &output, err
}

func (a *APIGatewayStub) GetVpcLinksAsync(ctx workflow.Context, input *apigateway.GetVpcLinksInput) *ApigatewayGetVpcLinksResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetVpcLinks, input)
	return &ApigatewayGetVpcLinksResult{Result: future}
}

func (a *APIGatewayStub) ImportApiKeys(ctx workflow.Context, input *apigateway.ImportApiKeysInput) (*apigateway.ImportApiKeysOutput, error) {
	var output apigateway.ImportApiKeysOutput
	err := workflow.ExecuteActivity(ctx, a.activities.ImportApiKeys, input).Get(ctx, &output)
	return &output, err
}

func (a *APIGatewayStub) ImportApiKeysAsync(ctx workflow.Context, input *apigateway.ImportApiKeysInput) *ApigatewayImportApiKeysResult {
	future := workflow.ExecuteActivity(ctx, a.activities.ImportApiKeys, input)
	return &ApigatewayImportApiKeysResult{Result: future}
}

func (a *APIGatewayStub) ImportDocumentationParts(ctx workflow.Context, input *apigateway.ImportDocumentationPartsInput) (*apigateway.ImportDocumentationPartsOutput, error) {
	var output apigateway.ImportDocumentationPartsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.ImportDocumentationParts, input).Get(ctx, &output)
	return &output, err
}

func (a *APIGatewayStub) ImportDocumentationPartsAsync(ctx workflow.Context, input *apigateway.ImportDocumentationPartsInput) *ApigatewayImportDocumentationPartsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.ImportDocumentationParts, input)
	return &ApigatewayImportDocumentationPartsResult{Result: future}
}

func (a *APIGatewayStub) ImportRestApi(ctx workflow.Context, input *apigateway.ImportRestApiInput) (*apigateway.RestApi, error) {
	var output apigateway.RestApi
	err := workflow.ExecuteActivity(ctx, a.activities.ImportRestApi, input).Get(ctx, &output)
	return &output, err
}

func (a *APIGatewayStub) ImportRestApiAsync(ctx workflow.Context, input *apigateway.ImportRestApiInput) *ApigatewayImportRestApiResult {
	future := workflow.ExecuteActivity(ctx, a.activities.ImportRestApi, input)
	return &ApigatewayImportRestApiResult{Result: future}
}

func (a *APIGatewayStub) PutGatewayResponse(ctx workflow.Context, input *apigateway.PutGatewayResponseInput) (*apigateway.UpdateGatewayResponseOutput, error) {
	var output apigateway.UpdateGatewayResponseOutput
	err := workflow.ExecuteActivity(ctx, a.activities.PutGatewayResponse, input).Get(ctx, &output)
	return &output, err
}

func (a *APIGatewayStub) PutGatewayResponseAsync(ctx workflow.Context, input *apigateway.PutGatewayResponseInput) *ApigatewayPutGatewayResponseResult {
	future := workflow.ExecuteActivity(ctx, a.activities.PutGatewayResponse, input)
	return &ApigatewayPutGatewayResponseResult{Result: future}
}

func (a *APIGatewayStub) PutIntegration(ctx workflow.Context, input *apigateway.PutIntegrationInput) (*apigateway.Integration, error) {
	var output apigateway.Integration
	err := workflow.ExecuteActivity(ctx, a.activities.PutIntegration, input).Get(ctx, &output)
	return &output, err
}

func (a *APIGatewayStub) PutIntegrationAsync(ctx workflow.Context, input *apigateway.PutIntegrationInput) *ApigatewayPutIntegrationResult {
	future := workflow.ExecuteActivity(ctx, a.activities.PutIntegration, input)
	return &ApigatewayPutIntegrationResult{Result: future}
}

func (a *APIGatewayStub) PutIntegrationResponse(ctx workflow.Context, input *apigateway.PutIntegrationResponseInput) (*apigateway.IntegrationResponse, error) {
	var output apigateway.IntegrationResponse
	err := workflow.ExecuteActivity(ctx, a.activities.PutIntegrationResponse, input).Get(ctx, &output)
	return &output, err
}

func (a *APIGatewayStub) PutIntegrationResponseAsync(ctx workflow.Context, input *apigateway.PutIntegrationResponseInput) *ApigatewayPutIntegrationResponseResult {
	future := workflow.ExecuteActivity(ctx, a.activities.PutIntegrationResponse, input)
	return &ApigatewayPutIntegrationResponseResult{Result: future}
}

func (a *APIGatewayStub) PutMethod(ctx workflow.Context, input *apigateway.PutMethodInput) (*apigateway.Method, error) {
	var output apigateway.Method
	err := workflow.ExecuteActivity(ctx, a.activities.PutMethod, input).Get(ctx, &output)
	return &output, err
}

func (a *APIGatewayStub) PutMethodAsync(ctx workflow.Context, input *apigateway.PutMethodInput) *ApigatewayPutMethodResult {
	future := workflow.ExecuteActivity(ctx, a.activities.PutMethod, input)
	return &ApigatewayPutMethodResult{Result: future}
}

func (a *APIGatewayStub) PutMethodResponse(ctx workflow.Context, input *apigateway.PutMethodResponseInput) (*apigateway.MethodResponse, error) {
	var output apigateway.MethodResponse
	err := workflow.ExecuteActivity(ctx, a.activities.PutMethodResponse, input).Get(ctx, &output)
	return &output, err
}

func (a *APIGatewayStub) PutMethodResponseAsync(ctx workflow.Context, input *apigateway.PutMethodResponseInput) *ApigatewayPutMethodResponseResult {
	future := workflow.ExecuteActivity(ctx, a.activities.PutMethodResponse, input)
	return &ApigatewayPutMethodResponseResult{Result: future}
}

func (a *APIGatewayStub) PutRestApi(ctx workflow.Context, input *apigateway.PutRestApiInput) (*apigateway.RestApi, error) {
	var output apigateway.RestApi
	err := workflow.ExecuteActivity(ctx, a.activities.PutRestApi, input).Get(ctx, &output)
	return &output, err
}

func (a *APIGatewayStub) PutRestApiAsync(ctx workflow.Context, input *apigateway.PutRestApiInput) *ApigatewayPutRestApiResult {
	future := workflow.ExecuteActivity(ctx, a.activities.PutRestApi, input)
	return &ApigatewayPutRestApiResult{Result: future}
}

func (a *APIGatewayStub) TagResource(ctx workflow.Context, input *apigateway.TagResourceInput) (*apigateway.TagResourceOutput, error) {
	var output apigateway.TagResourceOutput
	err := workflow.ExecuteActivity(ctx, a.activities.TagResource, input).Get(ctx, &output)
	return &output, err
}

func (a *APIGatewayStub) TagResourceAsync(ctx workflow.Context, input *apigateway.TagResourceInput) *ApigatewayTagResourceResult {
	future := workflow.ExecuteActivity(ctx, a.activities.TagResource, input)
	return &ApigatewayTagResourceResult{Result: future}
}

func (a *APIGatewayStub) TestInvokeAuthorizer(ctx workflow.Context, input *apigateway.TestInvokeAuthorizerInput) (*apigateway.TestInvokeAuthorizerOutput, error) {
	var output apigateway.TestInvokeAuthorizerOutput
	err := workflow.ExecuteActivity(ctx, a.activities.TestInvokeAuthorizer, input).Get(ctx, &output)
	return &output, err
}

func (a *APIGatewayStub) TestInvokeAuthorizerAsync(ctx workflow.Context, input *apigateway.TestInvokeAuthorizerInput) *ApigatewayTestInvokeAuthorizerResult {
	future := workflow.ExecuteActivity(ctx, a.activities.TestInvokeAuthorizer, input)
	return &ApigatewayTestInvokeAuthorizerResult{Result: future}
}

func (a *APIGatewayStub) TestInvokeMethod(ctx workflow.Context, input *apigateway.TestInvokeMethodInput) (*apigateway.TestInvokeMethodOutput, error) {
	var output apigateway.TestInvokeMethodOutput
	err := workflow.ExecuteActivity(ctx, a.activities.TestInvokeMethod, input).Get(ctx, &output)
	return &output, err
}

func (a *APIGatewayStub) TestInvokeMethodAsync(ctx workflow.Context, input *apigateway.TestInvokeMethodInput) *ApigatewayTestInvokeMethodResult {
	future := workflow.ExecuteActivity(ctx, a.activities.TestInvokeMethod, input)
	return &ApigatewayTestInvokeMethodResult{Result: future}
}

func (a *APIGatewayStub) UntagResource(ctx workflow.Context, input *apigateway.UntagResourceInput) (*apigateway.UntagResourceOutput, error) {
	var output apigateway.UntagResourceOutput
	err := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input).Get(ctx, &output)
	return &output, err
}

func (a *APIGatewayStub) UntagResourceAsync(ctx workflow.Context, input *apigateway.UntagResourceInput) *ApigatewayUntagResourceResult {
	future := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input)
	return &ApigatewayUntagResourceResult{Result: future}
}

func (a *APIGatewayStub) UpdateAccount(ctx workflow.Context, input *apigateway.UpdateAccountInput) (*apigateway.Account, error) {
	var output apigateway.Account
	err := workflow.ExecuteActivity(ctx, a.activities.UpdateAccount, input).Get(ctx, &output)
	return &output, err
}

func (a *APIGatewayStub) UpdateAccountAsync(ctx workflow.Context, input *apigateway.UpdateAccountInput) *ApigatewayUpdateAccountResult {
	future := workflow.ExecuteActivity(ctx, a.activities.UpdateAccount, input)
	return &ApigatewayUpdateAccountResult{Result: future}
}

func (a *APIGatewayStub) UpdateApiKey(ctx workflow.Context, input *apigateway.UpdateApiKeyInput) (*apigateway.ApiKey, error) {
	var output apigateway.ApiKey
	err := workflow.ExecuteActivity(ctx, a.activities.UpdateApiKey, input).Get(ctx, &output)
	return &output, err
}

func (a *APIGatewayStub) UpdateApiKeyAsync(ctx workflow.Context, input *apigateway.UpdateApiKeyInput) *ApigatewayUpdateApiKeyResult {
	future := workflow.ExecuteActivity(ctx, a.activities.UpdateApiKey, input)
	return &ApigatewayUpdateApiKeyResult{Result: future}
}

func (a *APIGatewayStub) UpdateAuthorizer(ctx workflow.Context, input *apigateway.UpdateAuthorizerInput) (*apigateway.Authorizer, error) {
	var output apigateway.Authorizer
	err := workflow.ExecuteActivity(ctx, a.activities.UpdateAuthorizer, input).Get(ctx, &output)
	return &output, err
}

func (a *APIGatewayStub) UpdateAuthorizerAsync(ctx workflow.Context, input *apigateway.UpdateAuthorizerInput) *ApigatewayUpdateAuthorizerResult {
	future := workflow.ExecuteActivity(ctx, a.activities.UpdateAuthorizer, input)
	return &ApigatewayUpdateAuthorizerResult{Result: future}
}

func (a *APIGatewayStub) UpdateBasePathMapping(ctx workflow.Context, input *apigateway.UpdateBasePathMappingInput) (*apigateway.BasePathMapping, error) {
	var output apigateway.BasePathMapping
	err := workflow.ExecuteActivity(ctx, a.activities.UpdateBasePathMapping, input).Get(ctx, &output)
	return &output, err
}

func (a *APIGatewayStub) UpdateBasePathMappingAsync(ctx workflow.Context, input *apigateway.UpdateBasePathMappingInput) *ApigatewayUpdateBasePathMappingResult {
	future := workflow.ExecuteActivity(ctx, a.activities.UpdateBasePathMapping, input)
	return &ApigatewayUpdateBasePathMappingResult{Result: future}
}

func (a *APIGatewayStub) UpdateClientCertificate(ctx workflow.Context, input *apigateway.UpdateClientCertificateInput) (*apigateway.ClientCertificate, error) {
	var output apigateway.ClientCertificate
	err := workflow.ExecuteActivity(ctx, a.activities.UpdateClientCertificate, input).Get(ctx, &output)
	return &output, err
}

func (a *APIGatewayStub) UpdateClientCertificateAsync(ctx workflow.Context, input *apigateway.UpdateClientCertificateInput) *ApigatewayUpdateClientCertificateResult {
	future := workflow.ExecuteActivity(ctx, a.activities.UpdateClientCertificate, input)
	return &ApigatewayUpdateClientCertificateResult{Result: future}
}

func (a *APIGatewayStub) UpdateDeployment(ctx workflow.Context, input *apigateway.UpdateDeploymentInput) (*apigateway.Deployment, error) {
	var output apigateway.Deployment
	err := workflow.ExecuteActivity(ctx, a.activities.UpdateDeployment, input).Get(ctx, &output)
	return &output, err
}

func (a *APIGatewayStub) UpdateDeploymentAsync(ctx workflow.Context, input *apigateway.UpdateDeploymentInput) *ApigatewayUpdateDeploymentResult {
	future := workflow.ExecuteActivity(ctx, a.activities.UpdateDeployment, input)
	return &ApigatewayUpdateDeploymentResult{Result: future}
}

func (a *APIGatewayStub) UpdateDocumentationPart(ctx workflow.Context, input *apigateway.UpdateDocumentationPartInput) (*apigateway.DocumentationPart, error) {
	var output apigateway.DocumentationPart
	err := workflow.ExecuteActivity(ctx, a.activities.UpdateDocumentationPart, input).Get(ctx, &output)
	return &output, err
}

func (a *APIGatewayStub) UpdateDocumentationPartAsync(ctx workflow.Context, input *apigateway.UpdateDocumentationPartInput) *ApigatewayUpdateDocumentationPartResult {
	future := workflow.ExecuteActivity(ctx, a.activities.UpdateDocumentationPart, input)
	return &ApigatewayUpdateDocumentationPartResult{Result: future}
}

func (a *APIGatewayStub) UpdateDocumentationVersion(ctx workflow.Context, input *apigateway.UpdateDocumentationVersionInput) (*apigateway.DocumentationVersion, error) {
	var output apigateway.DocumentationVersion
	err := workflow.ExecuteActivity(ctx, a.activities.UpdateDocumentationVersion, input).Get(ctx, &output)
	return &output, err
}

func (a *APIGatewayStub) UpdateDocumentationVersionAsync(ctx workflow.Context, input *apigateway.UpdateDocumentationVersionInput) *ApigatewayUpdateDocumentationVersionResult {
	future := workflow.ExecuteActivity(ctx, a.activities.UpdateDocumentationVersion, input)
	return &ApigatewayUpdateDocumentationVersionResult{Result: future}
}

func (a *APIGatewayStub) UpdateDomainName(ctx workflow.Context, input *apigateway.UpdateDomainNameInput) (*apigateway.DomainName, error) {
	var output apigateway.DomainName
	err := workflow.ExecuteActivity(ctx, a.activities.UpdateDomainName, input).Get(ctx, &output)
	return &output, err
}

func (a *APIGatewayStub) UpdateDomainNameAsync(ctx workflow.Context, input *apigateway.UpdateDomainNameInput) *ApigatewayUpdateDomainNameResult {
	future := workflow.ExecuteActivity(ctx, a.activities.UpdateDomainName, input)
	return &ApigatewayUpdateDomainNameResult{Result: future}
}

func (a *APIGatewayStub) UpdateGatewayResponse(ctx workflow.Context, input *apigateway.UpdateGatewayResponseInput) (*apigateway.UpdateGatewayResponseOutput, error) {
	var output apigateway.UpdateGatewayResponseOutput
	err := workflow.ExecuteActivity(ctx, a.activities.UpdateGatewayResponse, input).Get(ctx, &output)
	return &output, err
}

func (a *APIGatewayStub) UpdateGatewayResponseAsync(ctx workflow.Context, input *apigateway.UpdateGatewayResponseInput) *ApigatewayUpdateGatewayResponseResult {
	future := workflow.ExecuteActivity(ctx, a.activities.UpdateGatewayResponse, input)
	return &ApigatewayUpdateGatewayResponseResult{Result: future}
}

func (a *APIGatewayStub) UpdateIntegration(ctx workflow.Context, input *apigateway.UpdateIntegrationInput) (*apigateway.Integration, error) {
	var output apigateway.Integration
	err := workflow.ExecuteActivity(ctx, a.activities.UpdateIntegration, input).Get(ctx, &output)
	return &output, err
}

func (a *APIGatewayStub) UpdateIntegrationAsync(ctx workflow.Context, input *apigateway.UpdateIntegrationInput) *ApigatewayUpdateIntegrationResult {
	future := workflow.ExecuteActivity(ctx, a.activities.UpdateIntegration, input)
	return &ApigatewayUpdateIntegrationResult{Result: future}
}

func (a *APIGatewayStub) UpdateIntegrationResponse(ctx workflow.Context, input *apigateway.UpdateIntegrationResponseInput) (*apigateway.IntegrationResponse, error) {
	var output apigateway.IntegrationResponse
	err := workflow.ExecuteActivity(ctx, a.activities.UpdateIntegrationResponse, input).Get(ctx, &output)
	return &output, err
}

func (a *APIGatewayStub) UpdateIntegrationResponseAsync(ctx workflow.Context, input *apigateway.UpdateIntegrationResponseInput) *ApigatewayUpdateIntegrationResponseResult {
	future := workflow.ExecuteActivity(ctx, a.activities.UpdateIntegrationResponse, input)
	return &ApigatewayUpdateIntegrationResponseResult{Result: future}
}

func (a *APIGatewayStub) UpdateMethod(ctx workflow.Context, input *apigateway.UpdateMethodInput) (*apigateway.Method, error) {
	var output apigateway.Method
	err := workflow.ExecuteActivity(ctx, a.activities.UpdateMethod, input).Get(ctx, &output)
	return &output, err
}

func (a *APIGatewayStub) UpdateMethodAsync(ctx workflow.Context, input *apigateway.UpdateMethodInput) *ApigatewayUpdateMethodResult {
	future := workflow.ExecuteActivity(ctx, a.activities.UpdateMethod, input)
	return &ApigatewayUpdateMethodResult{Result: future}
}

func (a *APIGatewayStub) UpdateMethodResponse(ctx workflow.Context, input *apigateway.UpdateMethodResponseInput) (*apigateway.MethodResponse, error) {
	var output apigateway.MethodResponse
	err := workflow.ExecuteActivity(ctx, a.activities.UpdateMethodResponse, input).Get(ctx, &output)
	return &output, err
}

func (a *APIGatewayStub) UpdateMethodResponseAsync(ctx workflow.Context, input *apigateway.UpdateMethodResponseInput) *ApigatewayUpdateMethodResponseResult {
	future := workflow.ExecuteActivity(ctx, a.activities.UpdateMethodResponse, input)
	return &ApigatewayUpdateMethodResponseResult{Result: future}
}

func (a *APIGatewayStub) UpdateModel(ctx workflow.Context, input *apigateway.UpdateModelInput) (*apigateway.Model, error) {
	var output apigateway.Model
	err := workflow.ExecuteActivity(ctx, a.activities.UpdateModel, input).Get(ctx, &output)
	return &output, err
}

func (a *APIGatewayStub) UpdateModelAsync(ctx workflow.Context, input *apigateway.UpdateModelInput) *ApigatewayUpdateModelResult {
	future := workflow.ExecuteActivity(ctx, a.activities.UpdateModel, input)
	return &ApigatewayUpdateModelResult{Result: future}
}

func (a *APIGatewayStub) UpdateRequestValidator(ctx workflow.Context, input *apigateway.UpdateRequestValidatorInput) (*apigateway.UpdateRequestValidatorOutput, error) {
	var output apigateway.UpdateRequestValidatorOutput
	err := workflow.ExecuteActivity(ctx, a.activities.UpdateRequestValidator, input).Get(ctx, &output)
	return &output, err
}

func (a *APIGatewayStub) UpdateRequestValidatorAsync(ctx workflow.Context, input *apigateway.UpdateRequestValidatorInput) *ApigatewayUpdateRequestValidatorResult {
	future := workflow.ExecuteActivity(ctx, a.activities.UpdateRequestValidator, input)
	return &ApigatewayUpdateRequestValidatorResult{Result: future}
}

func (a *APIGatewayStub) UpdateResource(ctx workflow.Context, input *apigateway.UpdateResourceInput) (*apigateway.Resource, error) {
	var output apigateway.Resource
	err := workflow.ExecuteActivity(ctx, a.activities.UpdateResource, input).Get(ctx, &output)
	return &output, err
}

func (a *APIGatewayStub) UpdateResourceAsync(ctx workflow.Context, input *apigateway.UpdateResourceInput) *ApigatewayUpdateResourceResult {
	future := workflow.ExecuteActivity(ctx, a.activities.UpdateResource, input)
	return &ApigatewayUpdateResourceResult{Result: future}
}

func (a *APIGatewayStub) UpdateRestApi(ctx workflow.Context, input *apigateway.UpdateRestApiInput) (*apigateway.RestApi, error) {
	var output apigateway.RestApi
	err := workflow.ExecuteActivity(ctx, a.activities.UpdateRestApi, input).Get(ctx, &output)
	return &output, err
}

func (a *APIGatewayStub) UpdateRestApiAsync(ctx workflow.Context, input *apigateway.UpdateRestApiInput) *ApigatewayUpdateRestApiResult {
	future := workflow.ExecuteActivity(ctx, a.activities.UpdateRestApi, input)
	return &ApigatewayUpdateRestApiResult{Result: future}
}

func (a *APIGatewayStub) UpdateStage(ctx workflow.Context, input *apigateway.UpdateStageInput) (*apigateway.Stage, error) {
	var output apigateway.Stage
	err := workflow.ExecuteActivity(ctx, a.activities.UpdateStage, input).Get(ctx, &output)
	return &output, err
}

func (a *APIGatewayStub) UpdateStageAsync(ctx workflow.Context, input *apigateway.UpdateStageInput) *ApigatewayUpdateStageResult {
	future := workflow.ExecuteActivity(ctx, a.activities.UpdateStage, input)
	return &ApigatewayUpdateStageResult{Result: future}
}

func (a *APIGatewayStub) UpdateUsage(ctx workflow.Context, input *apigateway.UpdateUsageInput) (*apigateway.Usage, error) {
	var output apigateway.Usage
	err := workflow.ExecuteActivity(ctx, a.activities.UpdateUsage, input).Get(ctx, &output)
	return &output, err
}

func (a *APIGatewayStub) UpdateUsageAsync(ctx workflow.Context, input *apigateway.UpdateUsageInput) *ApigatewayUpdateUsageResult {
	future := workflow.ExecuteActivity(ctx, a.activities.UpdateUsage, input)
	return &ApigatewayUpdateUsageResult{Result: future}
}

func (a *APIGatewayStub) UpdateUsagePlan(ctx workflow.Context, input *apigateway.UpdateUsagePlanInput) (*apigateway.UsagePlan, error) {
	var output apigateway.UsagePlan
	err := workflow.ExecuteActivity(ctx, a.activities.UpdateUsagePlan, input).Get(ctx, &output)
	return &output, err
}

func (a *APIGatewayStub) UpdateUsagePlanAsync(ctx workflow.Context, input *apigateway.UpdateUsagePlanInput) *ApigatewayUpdateUsagePlanResult {
	future := workflow.ExecuteActivity(ctx, a.activities.UpdateUsagePlan, input)
	return &ApigatewayUpdateUsagePlanResult{Result: future}
}

func (a *APIGatewayStub) UpdateVpcLink(ctx workflow.Context, input *apigateway.UpdateVpcLinkInput) (*apigateway.UpdateVpcLinkOutput, error) {
	var output apigateway.UpdateVpcLinkOutput
	err := workflow.ExecuteActivity(ctx, a.activities.UpdateVpcLink, input).Get(ctx, &output)
	return &output, err
}

func (a *APIGatewayStub) UpdateVpcLinkAsync(ctx workflow.Context, input *apigateway.UpdateVpcLinkInput) *ApigatewayUpdateVpcLinkResult {
	future := workflow.ExecuteActivity(ctx, a.activities.UpdateVpcLink, input)
	return &ApigatewayUpdateVpcLinkResult{Result: future}
}
