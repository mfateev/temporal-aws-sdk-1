// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package appsyncstub

import (
	"github.com/aws/aws-sdk-go/service/appsync"
	"go.temporal.io/sdk/workflow"

	"go.temporal.io/aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type stub struct{}

type AppSyncCreateApiCacheFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AppSyncCreateApiCacheFuture) Get(ctx workflow.Context) (*appsync.CreateApiCacheOutput, error) {
	var output appsync.CreateApiCacheOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AppSyncCreateApiKeyFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AppSyncCreateApiKeyFuture) Get(ctx workflow.Context) (*appsync.CreateApiKeyOutput, error) {
	var output appsync.CreateApiKeyOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AppSyncCreateDataSourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AppSyncCreateDataSourceFuture) Get(ctx workflow.Context) (*appsync.CreateDataSourceOutput, error) {
	var output appsync.CreateDataSourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AppSyncCreateFunctionFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AppSyncCreateFunctionFuture) Get(ctx workflow.Context) (*appsync.CreateFunctionOutput, error) {
	var output appsync.CreateFunctionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AppSyncCreateGraphqlApiFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AppSyncCreateGraphqlApiFuture) Get(ctx workflow.Context) (*appsync.CreateGraphqlApiOutput, error) {
	var output appsync.CreateGraphqlApiOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AppSyncCreateResolverFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AppSyncCreateResolverFuture) Get(ctx workflow.Context) (*appsync.CreateResolverOutput, error) {
	var output appsync.CreateResolverOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AppSyncCreateTypeFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AppSyncCreateTypeFuture) Get(ctx workflow.Context) (*appsync.CreateTypeOutput, error) {
	var output appsync.CreateTypeOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AppSyncDeleteApiCacheFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AppSyncDeleteApiCacheFuture) Get(ctx workflow.Context) (*appsync.DeleteApiCacheOutput, error) {
	var output appsync.DeleteApiCacheOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AppSyncDeleteApiKeyFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AppSyncDeleteApiKeyFuture) Get(ctx workflow.Context) (*appsync.DeleteApiKeyOutput, error) {
	var output appsync.DeleteApiKeyOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AppSyncDeleteDataSourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AppSyncDeleteDataSourceFuture) Get(ctx workflow.Context) (*appsync.DeleteDataSourceOutput, error) {
	var output appsync.DeleteDataSourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AppSyncDeleteFunctionFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AppSyncDeleteFunctionFuture) Get(ctx workflow.Context) (*appsync.DeleteFunctionOutput, error) {
	var output appsync.DeleteFunctionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AppSyncDeleteGraphqlApiFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AppSyncDeleteGraphqlApiFuture) Get(ctx workflow.Context) (*appsync.DeleteGraphqlApiOutput, error) {
	var output appsync.DeleteGraphqlApiOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AppSyncDeleteResolverFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AppSyncDeleteResolverFuture) Get(ctx workflow.Context) (*appsync.DeleteResolverOutput, error) {
	var output appsync.DeleteResolverOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AppSyncDeleteTypeFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AppSyncDeleteTypeFuture) Get(ctx workflow.Context) (*appsync.DeleteTypeOutput, error) {
	var output appsync.DeleteTypeOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AppSyncFlushApiCacheFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AppSyncFlushApiCacheFuture) Get(ctx workflow.Context) (*appsync.FlushApiCacheOutput, error) {
	var output appsync.FlushApiCacheOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AppSyncGetApiCacheFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AppSyncGetApiCacheFuture) Get(ctx workflow.Context) (*appsync.GetApiCacheOutput, error) {
	var output appsync.GetApiCacheOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AppSyncGetDataSourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AppSyncGetDataSourceFuture) Get(ctx workflow.Context) (*appsync.GetDataSourceOutput, error) {
	var output appsync.GetDataSourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AppSyncGetFunctionFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AppSyncGetFunctionFuture) Get(ctx workflow.Context) (*appsync.GetFunctionOutput, error) {
	var output appsync.GetFunctionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AppSyncGetGraphqlApiFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AppSyncGetGraphqlApiFuture) Get(ctx workflow.Context) (*appsync.GetGraphqlApiOutput, error) {
	var output appsync.GetGraphqlApiOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AppSyncGetIntrospectionSchemaFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AppSyncGetIntrospectionSchemaFuture) Get(ctx workflow.Context) (*appsync.GetIntrospectionSchemaOutput, error) {
	var output appsync.GetIntrospectionSchemaOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AppSyncGetResolverFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AppSyncGetResolverFuture) Get(ctx workflow.Context) (*appsync.GetResolverOutput, error) {
	var output appsync.GetResolverOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AppSyncGetSchemaCreationStatusFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AppSyncGetSchemaCreationStatusFuture) Get(ctx workflow.Context) (*appsync.GetSchemaCreationStatusOutput, error) {
	var output appsync.GetSchemaCreationStatusOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AppSyncGetTypeFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AppSyncGetTypeFuture) Get(ctx workflow.Context) (*appsync.GetTypeOutput, error) {
	var output appsync.GetTypeOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AppSyncListApiKeysFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AppSyncListApiKeysFuture) Get(ctx workflow.Context) (*appsync.ListApiKeysOutput, error) {
	var output appsync.ListApiKeysOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AppSyncListDataSourcesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AppSyncListDataSourcesFuture) Get(ctx workflow.Context) (*appsync.ListDataSourcesOutput, error) {
	var output appsync.ListDataSourcesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AppSyncListFunctionsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AppSyncListFunctionsFuture) Get(ctx workflow.Context) (*appsync.ListFunctionsOutput, error) {
	var output appsync.ListFunctionsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AppSyncListGraphqlApisFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AppSyncListGraphqlApisFuture) Get(ctx workflow.Context) (*appsync.ListGraphqlApisOutput, error) {
	var output appsync.ListGraphqlApisOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AppSyncListResolversFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AppSyncListResolversFuture) Get(ctx workflow.Context) (*appsync.ListResolversOutput, error) {
	var output appsync.ListResolversOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AppSyncListResolversByFunctionFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AppSyncListResolversByFunctionFuture) Get(ctx workflow.Context) (*appsync.ListResolversByFunctionOutput, error) {
	var output appsync.ListResolversByFunctionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AppSyncListTagsForResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AppSyncListTagsForResourceFuture) Get(ctx workflow.Context) (*appsync.ListTagsForResourceOutput, error) {
	var output appsync.ListTagsForResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AppSyncListTypesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AppSyncListTypesFuture) Get(ctx workflow.Context) (*appsync.ListTypesOutput, error) {
	var output appsync.ListTypesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AppSyncStartSchemaCreationFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AppSyncStartSchemaCreationFuture) Get(ctx workflow.Context) (*appsync.StartSchemaCreationOutput, error) {
	var output appsync.StartSchemaCreationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AppSyncTagResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AppSyncTagResourceFuture) Get(ctx workflow.Context) (*appsync.TagResourceOutput, error) {
	var output appsync.TagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AppSyncUntagResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AppSyncUntagResourceFuture) Get(ctx workflow.Context) (*appsync.UntagResourceOutput, error) {
	var output appsync.UntagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AppSyncUpdateApiCacheFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AppSyncUpdateApiCacheFuture) Get(ctx workflow.Context) (*appsync.UpdateApiCacheOutput, error) {
	var output appsync.UpdateApiCacheOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AppSyncUpdateApiKeyFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AppSyncUpdateApiKeyFuture) Get(ctx workflow.Context) (*appsync.UpdateApiKeyOutput, error) {
	var output appsync.UpdateApiKeyOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AppSyncUpdateDataSourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AppSyncUpdateDataSourceFuture) Get(ctx workflow.Context) (*appsync.UpdateDataSourceOutput, error) {
	var output appsync.UpdateDataSourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AppSyncUpdateFunctionFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AppSyncUpdateFunctionFuture) Get(ctx workflow.Context) (*appsync.UpdateFunctionOutput, error) {
	var output appsync.UpdateFunctionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AppSyncUpdateGraphqlApiFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AppSyncUpdateGraphqlApiFuture) Get(ctx workflow.Context) (*appsync.UpdateGraphqlApiOutput, error) {
	var output appsync.UpdateGraphqlApiOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AppSyncUpdateResolverFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AppSyncUpdateResolverFuture) Get(ctx workflow.Context) (*appsync.UpdateResolverOutput, error) {
	var output appsync.UpdateResolverOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AppSyncUpdateTypeFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AppSyncUpdateTypeFuture) Get(ctx workflow.Context) (*appsync.UpdateTypeOutput, error) {
	var output appsync.UpdateTypeOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateApiCache(ctx workflow.Context, input *appsync.CreateApiCacheInput) (*appsync.CreateApiCacheOutput, error) {
	var output appsync.CreateApiCacheOutput
	err := workflow.ExecuteActivity(ctx, "aws.appsync.CreateApiCache", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateApiCacheAsync(ctx workflow.Context, input *appsync.CreateApiCacheInput) *AppSyncCreateApiCacheFuture {
	future := workflow.ExecuteActivity(ctx, "aws.appsync.CreateApiCache", input)
	return &AppSyncCreateApiCacheFuture{Future: future}
}

func (a *stub) CreateApiKey(ctx workflow.Context, input *appsync.CreateApiKeyInput) (*appsync.CreateApiKeyOutput, error) {
	var output appsync.CreateApiKeyOutput
	err := workflow.ExecuteActivity(ctx, "aws.appsync.CreateApiKey", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateApiKeyAsync(ctx workflow.Context, input *appsync.CreateApiKeyInput) *AppSyncCreateApiKeyFuture {
	future := workflow.ExecuteActivity(ctx, "aws.appsync.CreateApiKey", input)
	return &AppSyncCreateApiKeyFuture{Future: future}
}

func (a *stub) CreateDataSource(ctx workflow.Context, input *appsync.CreateDataSourceInput) (*appsync.CreateDataSourceOutput, error) {
	var output appsync.CreateDataSourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.appsync.CreateDataSource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateDataSourceAsync(ctx workflow.Context, input *appsync.CreateDataSourceInput) *AppSyncCreateDataSourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.appsync.CreateDataSource", input)
	return &AppSyncCreateDataSourceFuture{Future: future}
}

func (a *stub) CreateFunction(ctx workflow.Context, input *appsync.CreateFunctionInput) (*appsync.CreateFunctionOutput, error) {
	var output appsync.CreateFunctionOutput
	err := workflow.ExecuteActivity(ctx, "aws.appsync.CreateFunction", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateFunctionAsync(ctx workflow.Context, input *appsync.CreateFunctionInput) *AppSyncCreateFunctionFuture {
	future := workflow.ExecuteActivity(ctx, "aws.appsync.CreateFunction", input)
	return &AppSyncCreateFunctionFuture{Future: future}
}

func (a *stub) CreateGraphqlApi(ctx workflow.Context, input *appsync.CreateGraphqlApiInput) (*appsync.CreateGraphqlApiOutput, error) {
	var output appsync.CreateGraphqlApiOutput
	err := workflow.ExecuteActivity(ctx, "aws.appsync.CreateGraphqlApi", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateGraphqlApiAsync(ctx workflow.Context, input *appsync.CreateGraphqlApiInput) *AppSyncCreateGraphqlApiFuture {
	future := workflow.ExecuteActivity(ctx, "aws.appsync.CreateGraphqlApi", input)
	return &AppSyncCreateGraphqlApiFuture{Future: future}
}

func (a *stub) CreateResolver(ctx workflow.Context, input *appsync.CreateResolverInput) (*appsync.CreateResolverOutput, error) {
	var output appsync.CreateResolverOutput
	err := workflow.ExecuteActivity(ctx, "aws.appsync.CreateResolver", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateResolverAsync(ctx workflow.Context, input *appsync.CreateResolverInput) *AppSyncCreateResolverFuture {
	future := workflow.ExecuteActivity(ctx, "aws.appsync.CreateResolver", input)
	return &AppSyncCreateResolverFuture{Future: future}
}

func (a *stub) CreateType(ctx workflow.Context, input *appsync.CreateTypeInput) (*appsync.CreateTypeOutput, error) {
	var output appsync.CreateTypeOutput
	err := workflow.ExecuteActivity(ctx, "aws.appsync.CreateType", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateTypeAsync(ctx workflow.Context, input *appsync.CreateTypeInput) *AppSyncCreateTypeFuture {
	future := workflow.ExecuteActivity(ctx, "aws.appsync.CreateType", input)
	return &AppSyncCreateTypeFuture{Future: future}
}

func (a *stub) DeleteApiCache(ctx workflow.Context, input *appsync.DeleteApiCacheInput) (*appsync.DeleteApiCacheOutput, error) {
	var output appsync.DeleteApiCacheOutput
	err := workflow.ExecuteActivity(ctx, "aws.appsync.DeleteApiCache", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteApiCacheAsync(ctx workflow.Context, input *appsync.DeleteApiCacheInput) *AppSyncDeleteApiCacheFuture {
	future := workflow.ExecuteActivity(ctx, "aws.appsync.DeleteApiCache", input)
	return &AppSyncDeleteApiCacheFuture{Future: future}
}

func (a *stub) DeleteApiKey(ctx workflow.Context, input *appsync.DeleteApiKeyInput) (*appsync.DeleteApiKeyOutput, error) {
	var output appsync.DeleteApiKeyOutput
	err := workflow.ExecuteActivity(ctx, "aws.appsync.DeleteApiKey", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteApiKeyAsync(ctx workflow.Context, input *appsync.DeleteApiKeyInput) *AppSyncDeleteApiKeyFuture {
	future := workflow.ExecuteActivity(ctx, "aws.appsync.DeleteApiKey", input)
	return &AppSyncDeleteApiKeyFuture{Future: future}
}

func (a *stub) DeleteDataSource(ctx workflow.Context, input *appsync.DeleteDataSourceInput) (*appsync.DeleteDataSourceOutput, error) {
	var output appsync.DeleteDataSourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.appsync.DeleteDataSource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteDataSourceAsync(ctx workflow.Context, input *appsync.DeleteDataSourceInput) *AppSyncDeleteDataSourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.appsync.DeleteDataSource", input)
	return &AppSyncDeleteDataSourceFuture{Future: future}
}

func (a *stub) DeleteFunction(ctx workflow.Context, input *appsync.DeleteFunctionInput) (*appsync.DeleteFunctionOutput, error) {
	var output appsync.DeleteFunctionOutput
	err := workflow.ExecuteActivity(ctx, "aws.appsync.DeleteFunction", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteFunctionAsync(ctx workflow.Context, input *appsync.DeleteFunctionInput) *AppSyncDeleteFunctionFuture {
	future := workflow.ExecuteActivity(ctx, "aws.appsync.DeleteFunction", input)
	return &AppSyncDeleteFunctionFuture{Future: future}
}

func (a *stub) DeleteGraphqlApi(ctx workflow.Context, input *appsync.DeleteGraphqlApiInput) (*appsync.DeleteGraphqlApiOutput, error) {
	var output appsync.DeleteGraphqlApiOutput
	err := workflow.ExecuteActivity(ctx, "aws.appsync.DeleteGraphqlApi", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteGraphqlApiAsync(ctx workflow.Context, input *appsync.DeleteGraphqlApiInput) *AppSyncDeleteGraphqlApiFuture {
	future := workflow.ExecuteActivity(ctx, "aws.appsync.DeleteGraphqlApi", input)
	return &AppSyncDeleteGraphqlApiFuture{Future: future}
}

func (a *stub) DeleteResolver(ctx workflow.Context, input *appsync.DeleteResolverInput) (*appsync.DeleteResolverOutput, error) {
	var output appsync.DeleteResolverOutput
	err := workflow.ExecuteActivity(ctx, "aws.appsync.DeleteResolver", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteResolverAsync(ctx workflow.Context, input *appsync.DeleteResolverInput) *AppSyncDeleteResolverFuture {
	future := workflow.ExecuteActivity(ctx, "aws.appsync.DeleteResolver", input)
	return &AppSyncDeleteResolverFuture{Future: future}
}

func (a *stub) DeleteType(ctx workflow.Context, input *appsync.DeleteTypeInput) (*appsync.DeleteTypeOutput, error) {
	var output appsync.DeleteTypeOutput
	err := workflow.ExecuteActivity(ctx, "aws.appsync.DeleteType", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteTypeAsync(ctx workflow.Context, input *appsync.DeleteTypeInput) *AppSyncDeleteTypeFuture {
	future := workflow.ExecuteActivity(ctx, "aws.appsync.DeleteType", input)
	return &AppSyncDeleteTypeFuture{Future: future}
}

func (a *stub) FlushApiCache(ctx workflow.Context, input *appsync.FlushApiCacheInput) (*appsync.FlushApiCacheOutput, error) {
	var output appsync.FlushApiCacheOutput
	err := workflow.ExecuteActivity(ctx, "aws.appsync.FlushApiCache", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) FlushApiCacheAsync(ctx workflow.Context, input *appsync.FlushApiCacheInput) *AppSyncFlushApiCacheFuture {
	future := workflow.ExecuteActivity(ctx, "aws.appsync.FlushApiCache", input)
	return &AppSyncFlushApiCacheFuture{Future: future}
}

func (a *stub) GetApiCache(ctx workflow.Context, input *appsync.GetApiCacheInput) (*appsync.GetApiCacheOutput, error) {
	var output appsync.GetApiCacheOutput
	err := workflow.ExecuteActivity(ctx, "aws.appsync.GetApiCache", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetApiCacheAsync(ctx workflow.Context, input *appsync.GetApiCacheInput) *AppSyncGetApiCacheFuture {
	future := workflow.ExecuteActivity(ctx, "aws.appsync.GetApiCache", input)
	return &AppSyncGetApiCacheFuture{Future: future}
}

func (a *stub) GetDataSource(ctx workflow.Context, input *appsync.GetDataSourceInput) (*appsync.GetDataSourceOutput, error) {
	var output appsync.GetDataSourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.appsync.GetDataSource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetDataSourceAsync(ctx workflow.Context, input *appsync.GetDataSourceInput) *AppSyncGetDataSourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.appsync.GetDataSource", input)
	return &AppSyncGetDataSourceFuture{Future: future}
}

func (a *stub) GetFunction(ctx workflow.Context, input *appsync.GetFunctionInput) (*appsync.GetFunctionOutput, error) {
	var output appsync.GetFunctionOutput
	err := workflow.ExecuteActivity(ctx, "aws.appsync.GetFunction", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetFunctionAsync(ctx workflow.Context, input *appsync.GetFunctionInput) *AppSyncGetFunctionFuture {
	future := workflow.ExecuteActivity(ctx, "aws.appsync.GetFunction", input)
	return &AppSyncGetFunctionFuture{Future: future}
}

func (a *stub) GetGraphqlApi(ctx workflow.Context, input *appsync.GetGraphqlApiInput) (*appsync.GetGraphqlApiOutput, error) {
	var output appsync.GetGraphqlApiOutput
	err := workflow.ExecuteActivity(ctx, "aws.appsync.GetGraphqlApi", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetGraphqlApiAsync(ctx workflow.Context, input *appsync.GetGraphqlApiInput) *AppSyncGetGraphqlApiFuture {
	future := workflow.ExecuteActivity(ctx, "aws.appsync.GetGraphqlApi", input)
	return &AppSyncGetGraphqlApiFuture{Future: future}
}

func (a *stub) GetIntrospectionSchema(ctx workflow.Context, input *appsync.GetIntrospectionSchemaInput) (*appsync.GetIntrospectionSchemaOutput, error) {
	var output appsync.GetIntrospectionSchemaOutput
	err := workflow.ExecuteActivity(ctx, "aws.appsync.GetIntrospectionSchema", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetIntrospectionSchemaAsync(ctx workflow.Context, input *appsync.GetIntrospectionSchemaInput) *AppSyncGetIntrospectionSchemaFuture {
	future := workflow.ExecuteActivity(ctx, "aws.appsync.GetIntrospectionSchema", input)
	return &AppSyncGetIntrospectionSchemaFuture{Future: future}
}

func (a *stub) GetResolver(ctx workflow.Context, input *appsync.GetResolverInput) (*appsync.GetResolverOutput, error) {
	var output appsync.GetResolverOutput
	err := workflow.ExecuteActivity(ctx, "aws.appsync.GetResolver", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetResolverAsync(ctx workflow.Context, input *appsync.GetResolverInput) *AppSyncGetResolverFuture {
	future := workflow.ExecuteActivity(ctx, "aws.appsync.GetResolver", input)
	return &AppSyncGetResolverFuture{Future: future}
}

func (a *stub) GetSchemaCreationStatus(ctx workflow.Context, input *appsync.GetSchemaCreationStatusInput) (*appsync.GetSchemaCreationStatusOutput, error) {
	var output appsync.GetSchemaCreationStatusOutput
	err := workflow.ExecuteActivity(ctx, "aws.appsync.GetSchemaCreationStatus", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetSchemaCreationStatusAsync(ctx workflow.Context, input *appsync.GetSchemaCreationStatusInput) *AppSyncGetSchemaCreationStatusFuture {
	future := workflow.ExecuteActivity(ctx, "aws.appsync.GetSchemaCreationStatus", input)
	return &AppSyncGetSchemaCreationStatusFuture{Future: future}
}

func (a *stub) GetType(ctx workflow.Context, input *appsync.GetTypeInput) (*appsync.GetTypeOutput, error) {
	var output appsync.GetTypeOutput
	err := workflow.ExecuteActivity(ctx, "aws.appsync.GetType", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetTypeAsync(ctx workflow.Context, input *appsync.GetTypeInput) *AppSyncGetTypeFuture {
	future := workflow.ExecuteActivity(ctx, "aws.appsync.GetType", input)
	return &AppSyncGetTypeFuture{Future: future}
}

func (a *stub) ListApiKeys(ctx workflow.Context, input *appsync.ListApiKeysInput) (*appsync.ListApiKeysOutput, error) {
	var output appsync.ListApiKeysOutput
	err := workflow.ExecuteActivity(ctx, "aws.appsync.ListApiKeys", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListApiKeysAsync(ctx workflow.Context, input *appsync.ListApiKeysInput) *AppSyncListApiKeysFuture {
	future := workflow.ExecuteActivity(ctx, "aws.appsync.ListApiKeys", input)
	return &AppSyncListApiKeysFuture{Future: future}
}

func (a *stub) ListDataSources(ctx workflow.Context, input *appsync.ListDataSourcesInput) (*appsync.ListDataSourcesOutput, error) {
	var output appsync.ListDataSourcesOutput
	err := workflow.ExecuteActivity(ctx, "aws.appsync.ListDataSources", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListDataSourcesAsync(ctx workflow.Context, input *appsync.ListDataSourcesInput) *AppSyncListDataSourcesFuture {
	future := workflow.ExecuteActivity(ctx, "aws.appsync.ListDataSources", input)
	return &AppSyncListDataSourcesFuture{Future: future}
}

func (a *stub) ListFunctions(ctx workflow.Context, input *appsync.ListFunctionsInput) (*appsync.ListFunctionsOutput, error) {
	var output appsync.ListFunctionsOutput
	err := workflow.ExecuteActivity(ctx, "aws.appsync.ListFunctions", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListFunctionsAsync(ctx workflow.Context, input *appsync.ListFunctionsInput) *AppSyncListFunctionsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.appsync.ListFunctions", input)
	return &AppSyncListFunctionsFuture{Future: future}
}

func (a *stub) ListGraphqlApis(ctx workflow.Context, input *appsync.ListGraphqlApisInput) (*appsync.ListGraphqlApisOutput, error) {
	var output appsync.ListGraphqlApisOutput
	err := workflow.ExecuteActivity(ctx, "aws.appsync.ListGraphqlApis", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListGraphqlApisAsync(ctx workflow.Context, input *appsync.ListGraphqlApisInput) *AppSyncListGraphqlApisFuture {
	future := workflow.ExecuteActivity(ctx, "aws.appsync.ListGraphqlApis", input)
	return &AppSyncListGraphqlApisFuture{Future: future}
}

func (a *stub) ListResolvers(ctx workflow.Context, input *appsync.ListResolversInput) (*appsync.ListResolversOutput, error) {
	var output appsync.ListResolversOutput
	err := workflow.ExecuteActivity(ctx, "aws.appsync.ListResolvers", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListResolversAsync(ctx workflow.Context, input *appsync.ListResolversInput) *AppSyncListResolversFuture {
	future := workflow.ExecuteActivity(ctx, "aws.appsync.ListResolvers", input)
	return &AppSyncListResolversFuture{Future: future}
}

func (a *stub) ListResolversByFunction(ctx workflow.Context, input *appsync.ListResolversByFunctionInput) (*appsync.ListResolversByFunctionOutput, error) {
	var output appsync.ListResolversByFunctionOutput
	err := workflow.ExecuteActivity(ctx, "aws.appsync.ListResolversByFunction", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListResolversByFunctionAsync(ctx workflow.Context, input *appsync.ListResolversByFunctionInput) *AppSyncListResolversByFunctionFuture {
	future := workflow.ExecuteActivity(ctx, "aws.appsync.ListResolversByFunction", input)
	return &AppSyncListResolversByFunctionFuture{Future: future}
}

func (a *stub) ListTagsForResource(ctx workflow.Context, input *appsync.ListTagsForResourceInput) (*appsync.ListTagsForResourceOutput, error) {
	var output appsync.ListTagsForResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.appsync.ListTagsForResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListTagsForResourceAsync(ctx workflow.Context, input *appsync.ListTagsForResourceInput) *AppSyncListTagsForResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.appsync.ListTagsForResource", input)
	return &AppSyncListTagsForResourceFuture{Future: future}
}

func (a *stub) ListTypes(ctx workflow.Context, input *appsync.ListTypesInput) (*appsync.ListTypesOutput, error) {
	var output appsync.ListTypesOutput
	err := workflow.ExecuteActivity(ctx, "aws.appsync.ListTypes", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListTypesAsync(ctx workflow.Context, input *appsync.ListTypesInput) *AppSyncListTypesFuture {
	future := workflow.ExecuteActivity(ctx, "aws.appsync.ListTypes", input)
	return &AppSyncListTypesFuture{Future: future}
}

func (a *stub) StartSchemaCreation(ctx workflow.Context, input *appsync.StartSchemaCreationInput) (*appsync.StartSchemaCreationOutput, error) {
	var output appsync.StartSchemaCreationOutput
	err := workflow.ExecuteActivity(ctx, "aws.appsync.StartSchemaCreation", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) StartSchemaCreationAsync(ctx workflow.Context, input *appsync.StartSchemaCreationInput) *AppSyncStartSchemaCreationFuture {
	future := workflow.ExecuteActivity(ctx, "aws.appsync.StartSchemaCreation", input)
	return &AppSyncStartSchemaCreationFuture{Future: future}
}

func (a *stub) TagResource(ctx workflow.Context, input *appsync.TagResourceInput) (*appsync.TagResourceOutput, error) {
	var output appsync.TagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.appsync.TagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) TagResourceAsync(ctx workflow.Context, input *appsync.TagResourceInput) *AppSyncTagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.appsync.TagResource", input)
	return &AppSyncTagResourceFuture{Future: future}
}

func (a *stub) UntagResource(ctx workflow.Context, input *appsync.UntagResourceInput) (*appsync.UntagResourceOutput, error) {
	var output appsync.UntagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.appsync.UntagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UntagResourceAsync(ctx workflow.Context, input *appsync.UntagResourceInput) *AppSyncUntagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.appsync.UntagResource", input)
	return &AppSyncUntagResourceFuture{Future: future}
}

func (a *stub) UpdateApiCache(ctx workflow.Context, input *appsync.UpdateApiCacheInput) (*appsync.UpdateApiCacheOutput, error) {
	var output appsync.UpdateApiCacheOutput
	err := workflow.ExecuteActivity(ctx, "aws.appsync.UpdateApiCache", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateApiCacheAsync(ctx workflow.Context, input *appsync.UpdateApiCacheInput) *AppSyncUpdateApiCacheFuture {
	future := workflow.ExecuteActivity(ctx, "aws.appsync.UpdateApiCache", input)
	return &AppSyncUpdateApiCacheFuture{Future: future}
}

func (a *stub) UpdateApiKey(ctx workflow.Context, input *appsync.UpdateApiKeyInput) (*appsync.UpdateApiKeyOutput, error) {
	var output appsync.UpdateApiKeyOutput
	err := workflow.ExecuteActivity(ctx, "aws.appsync.UpdateApiKey", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateApiKeyAsync(ctx workflow.Context, input *appsync.UpdateApiKeyInput) *AppSyncUpdateApiKeyFuture {
	future := workflow.ExecuteActivity(ctx, "aws.appsync.UpdateApiKey", input)
	return &AppSyncUpdateApiKeyFuture{Future: future}
}

func (a *stub) UpdateDataSource(ctx workflow.Context, input *appsync.UpdateDataSourceInput) (*appsync.UpdateDataSourceOutput, error) {
	var output appsync.UpdateDataSourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.appsync.UpdateDataSource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateDataSourceAsync(ctx workflow.Context, input *appsync.UpdateDataSourceInput) *AppSyncUpdateDataSourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.appsync.UpdateDataSource", input)
	return &AppSyncUpdateDataSourceFuture{Future: future}
}

func (a *stub) UpdateFunction(ctx workflow.Context, input *appsync.UpdateFunctionInput) (*appsync.UpdateFunctionOutput, error) {
	var output appsync.UpdateFunctionOutput
	err := workflow.ExecuteActivity(ctx, "aws.appsync.UpdateFunction", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateFunctionAsync(ctx workflow.Context, input *appsync.UpdateFunctionInput) *AppSyncUpdateFunctionFuture {
	future := workflow.ExecuteActivity(ctx, "aws.appsync.UpdateFunction", input)
	return &AppSyncUpdateFunctionFuture{Future: future}
}

func (a *stub) UpdateGraphqlApi(ctx workflow.Context, input *appsync.UpdateGraphqlApiInput) (*appsync.UpdateGraphqlApiOutput, error) {
	var output appsync.UpdateGraphqlApiOutput
	err := workflow.ExecuteActivity(ctx, "aws.appsync.UpdateGraphqlApi", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateGraphqlApiAsync(ctx workflow.Context, input *appsync.UpdateGraphqlApiInput) *AppSyncUpdateGraphqlApiFuture {
	future := workflow.ExecuteActivity(ctx, "aws.appsync.UpdateGraphqlApi", input)
	return &AppSyncUpdateGraphqlApiFuture{Future: future}
}

func (a *stub) UpdateResolver(ctx workflow.Context, input *appsync.UpdateResolverInput) (*appsync.UpdateResolverOutput, error) {
	var output appsync.UpdateResolverOutput
	err := workflow.ExecuteActivity(ctx, "aws.appsync.UpdateResolver", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateResolverAsync(ctx workflow.Context, input *appsync.UpdateResolverInput) *AppSyncUpdateResolverFuture {
	future := workflow.ExecuteActivity(ctx, "aws.appsync.UpdateResolver", input)
	return &AppSyncUpdateResolverFuture{Future: future}
}

func (a *stub) UpdateType(ctx workflow.Context, input *appsync.UpdateTypeInput) (*appsync.UpdateTypeOutput, error) {
	var output appsync.UpdateTypeOutput
	err := workflow.ExecuteActivity(ctx, "aws.appsync.UpdateType", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateTypeAsync(ctx workflow.Context, input *appsync.UpdateTypeInput) *AppSyncUpdateTypeFuture {
	future := workflow.ExecuteActivity(ctx, "aws.appsync.UpdateType", input)
	return &AppSyncUpdateTypeFuture{Future: future}
}
