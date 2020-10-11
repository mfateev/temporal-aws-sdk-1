// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iotsitewise"
	"github.com/aws/aws-sdk-go/service/iotsitewise/iotsitewiseiface"
	"go.temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken
type _ request.Option

type IoTSiteWiseActivities struct {
	client iotsitewiseiface.IoTSiteWiseAPI

	sessionFactory SessionFactory
}

func NewIoTSiteWiseActivities(sess *session.Session, config ...*aws.Config) *IoTSiteWiseActivities {
	client := iotsitewise.New(sess, config...)
	return &IoTSiteWiseActivities{client: client}
}

func NewIoTSiteWiseActivitiesWithSessionFactory(sessionFactory SessionFactory) *IoTSiteWiseActivities {
	return &IoTSiteWiseActivities{sessionFactory: sessionFactory}
}

func (a *IoTSiteWiseActivities) getClient(ctx context.Context) (iotsitewiseiface.IoTSiteWiseAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, err
	}

	return iotsitewise.New(sess), nil
}

func (a *IoTSiteWiseActivities) AssociateAssets(ctx context.Context, input *iotsitewise.AssociateAssetsInput) (*iotsitewise.AssociateAssetsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	internal.SetClientToken(ctx, &input.ClientToken)
	return client.AssociateAssetsWithContext(ctx, input)
}

func (a *IoTSiteWiseActivities) BatchAssociateProjectAssets(ctx context.Context, input *iotsitewise.BatchAssociateProjectAssetsInput) (*iotsitewise.BatchAssociateProjectAssetsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	internal.SetClientToken(ctx, &input.ClientToken)
	return client.BatchAssociateProjectAssetsWithContext(ctx, input)
}

func (a *IoTSiteWiseActivities) BatchDisassociateProjectAssets(ctx context.Context, input *iotsitewise.BatchDisassociateProjectAssetsInput) (*iotsitewise.BatchDisassociateProjectAssetsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	internal.SetClientToken(ctx, &input.ClientToken)
	return client.BatchDisassociateProjectAssetsWithContext(ctx, input)
}

func (a *IoTSiteWiseActivities) BatchPutAssetPropertyValue(ctx context.Context, input *iotsitewise.BatchPutAssetPropertyValueInput) (*iotsitewise.BatchPutAssetPropertyValueOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.BatchPutAssetPropertyValueWithContext(ctx, input)
}

func (a *IoTSiteWiseActivities) CreateAccessPolicy(ctx context.Context, input *iotsitewise.CreateAccessPolicyInput) (*iotsitewise.CreateAccessPolicyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	internal.SetClientToken(ctx, &input.ClientToken)
	return client.CreateAccessPolicyWithContext(ctx, input)
}

func (a *IoTSiteWiseActivities) CreateAsset(ctx context.Context, input *iotsitewise.CreateAssetInput) (*iotsitewise.CreateAssetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	internal.SetClientToken(ctx, &input.ClientToken)
	return client.CreateAssetWithContext(ctx, input)
}

func (a *IoTSiteWiseActivities) CreateAssetModel(ctx context.Context, input *iotsitewise.CreateAssetModelInput) (*iotsitewise.CreateAssetModelOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	internal.SetClientToken(ctx, &input.ClientToken)
	return client.CreateAssetModelWithContext(ctx, input)
}

func (a *IoTSiteWiseActivities) CreateDashboard(ctx context.Context, input *iotsitewise.CreateDashboardInput) (*iotsitewise.CreateDashboardOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	internal.SetClientToken(ctx, &input.ClientToken)
	return client.CreateDashboardWithContext(ctx, input)
}

func (a *IoTSiteWiseActivities) CreateGateway(ctx context.Context, input *iotsitewise.CreateGatewayInput) (*iotsitewise.CreateGatewayOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateGatewayWithContext(ctx, input)
}

func (a *IoTSiteWiseActivities) CreatePortal(ctx context.Context, input *iotsitewise.CreatePortalInput) (*iotsitewise.CreatePortalOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	internal.SetClientToken(ctx, &input.ClientToken)
	return client.CreatePortalWithContext(ctx, input)
}

func (a *IoTSiteWiseActivities) CreatePresignedPortalUrl(ctx context.Context, input *iotsitewise.CreatePresignedPortalUrlInput) (*iotsitewise.CreatePresignedPortalUrlOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreatePresignedPortalUrlWithContext(ctx, input)
}

func (a *IoTSiteWiseActivities) CreateProject(ctx context.Context, input *iotsitewise.CreateProjectInput) (*iotsitewise.CreateProjectOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	internal.SetClientToken(ctx, &input.ClientToken)
	return client.CreateProjectWithContext(ctx, input)
}

func (a *IoTSiteWiseActivities) DeleteAccessPolicy(ctx context.Context, input *iotsitewise.DeleteAccessPolicyInput) (*iotsitewise.DeleteAccessPolicyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	internal.SetClientToken(ctx, &input.ClientToken)
	return client.DeleteAccessPolicyWithContext(ctx, input)
}

func (a *IoTSiteWiseActivities) DeleteAsset(ctx context.Context, input *iotsitewise.DeleteAssetInput) (*iotsitewise.DeleteAssetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	internal.SetClientToken(ctx, &input.ClientToken)
	return client.DeleteAssetWithContext(ctx, input)
}

func (a *IoTSiteWiseActivities) DeleteAssetModel(ctx context.Context, input *iotsitewise.DeleteAssetModelInput) (*iotsitewise.DeleteAssetModelOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	internal.SetClientToken(ctx, &input.ClientToken)
	return client.DeleteAssetModelWithContext(ctx, input)
}

func (a *IoTSiteWiseActivities) DeleteDashboard(ctx context.Context, input *iotsitewise.DeleteDashboardInput) (*iotsitewise.DeleteDashboardOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	internal.SetClientToken(ctx, &input.ClientToken)
	return client.DeleteDashboardWithContext(ctx, input)
}

func (a *IoTSiteWiseActivities) DeleteGateway(ctx context.Context, input *iotsitewise.DeleteGatewayInput) (*iotsitewise.DeleteGatewayOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteGatewayWithContext(ctx, input)
}

func (a *IoTSiteWiseActivities) DeletePortal(ctx context.Context, input *iotsitewise.DeletePortalInput) (*iotsitewise.DeletePortalOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	internal.SetClientToken(ctx, &input.ClientToken)
	return client.DeletePortalWithContext(ctx, input)
}

func (a *IoTSiteWiseActivities) DeleteProject(ctx context.Context, input *iotsitewise.DeleteProjectInput) (*iotsitewise.DeleteProjectOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	internal.SetClientToken(ctx, &input.ClientToken)
	return client.DeleteProjectWithContext(ctx, input)
}

func (a *IoTSiteWiseActivities) DescribeAccessPolicy(ctx context.Context, input *iotsitewise.DescribeAccessPolicyInput) (*iotsitewise.DescribeAccessPolicyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeAccessPolicyWithContext(ctx, input)
}

func (a *IoTSiteWiseActivities) DescribeAsset(ctx context.Context, input *iotsitewise.DescribeAssetInput) (*iotsitewise.DescribeAssetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeAssetWithContext(ctx, input)
}

func (a *IoTSiteWiseActivities) DescribeAssetModel(ctx context.Context, input *iotsitewise.DescribeAssetModelInput) (*iotsitewise.DescribeAssetModelOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeAssetModelWithContext(ctx, input)
}

func (a *IoTSiteWiseActivities) DescribeAssetProperty(ctx context.Context, input *iotsitewise.DescribeAssetPropertyInput) (*iotsitewise.DescribeAssetPropertyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeAssetPropertyWithContext(ctx, input)
}

func (a *IoTSiteWiseActivities) DescribeDashboard(ctx context.Context, input *iotsitewise.DescribeDashboardInput) (*iotsitewise.DescribeDashboardOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeDashboardWithContext(ctx, input)
}

func (a *IoTSiteWiseActivities) DescribeGateway(ctx context.Context, input *iotsitewise.DescribeGatewayInput) (*iotsitewise.DescribeGatewayOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeGatewayWithContext(ctx, input)
}

func (a *IoTSiteWiseActivities) DescribeGatewayCapabilityConfiguration(ctx context.Context, input *iotsitewise.DescribeGatewayCapabilityConfigurationInput) (*iotsitewise.DescribeGatewayCapabilityConfigurationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeGatewayCapabilityConfigurationWithContext(ctx, input)
}

func (a *IoTSiteWiseActivities) DescribeLoggingOptions(ctx context.Context, input *iotsitewise.DescribeLoggingOptionsInput) (*iotsitewise.DescribeLoggingOptionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeLoggingOptionsWithContext(ctx, input)
}

func (a *IoTSiteWiseActivities) DescribePortal(ctx context.Context, input *iotsitewise.DescribePortalInput) (*iotsitewise.DescribePortalOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribePortalWithContext(ctx, input)
}

func (a *IoTSiteWiseActivities) DescribeProject(ctx context.Context, input *iotsitewise.DescribeProjectInput) (*iotsitewise.DescribeProjectOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeProjectWithContext(ctx, input)
}

func (a *IoTSiteWiseActivities) DisassociateAssets(ctx context.Context, input *iotsitewise.DisassociateAssetsInput) (*iotsitewise.DisassociateAssetsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	internal.SetClientToken(ctx, &input.ClientToken)
	return client.DisassociateAssetsWithContext(ctx, input)
}

func (a *IoTSiteWiseActivities) GetAssetPropertyAggregates(ctx context.Context, input *iotsitewise.GetAssetPropertyAggregatesInput) (*iotsitewise.GetAssetPropertyAggregatesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetAssetPropertyAggregatesWithContext(ctx, input)
}

func (a *IoTSiteWiseActivities) GetAssetPropertyValue(ctx context.Context, input *iotsitewise.GetAssetPropertyValueInput) (*iotsitewise.GetAssetPropertyValueOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetAssetPropertyValueWithContext(ctx, input)
}

func (a *IoTSiteWiseActivities) GetAssetPropertyValueHistory(ctx context.Context, input *iotsitewise.GetAssetPropertyValueHistoryInput) (*iotsitewise.GetAssetPropertyValueHistoryOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetAssetPropertyValueHistoryWithContext(ctx, input)
}

func (a *IoTSiteWiseActivities) ListAccessPolicies(ctx context.Context, input *iotsitewise.ListAccessPoliciesInput) (*iotsitewise.ListAccessPoliciesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListAccessPoliciesWithContext(ctx, input)
}

func (a *IoTSiteWiseActivities) ListAssetModels(ctx context.Context, input *iotsitewise.ListAssetModelsInput) (*iotsitewise.ListAssetModelsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListAssetModelsWithContext(ctx, input)
}

func (a *IoTSiteWiseActivities) ListAssets(ctx context.Context, input *iotsitewise.ListAssetsInput) (*iotsitewise.ListAssetsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListAssetsWithContext(ctx, input)
}

func (a *IoTSiteWiseActivities) ListAssociatedAssets(ctx context.Context, input *iotsitewise.ListAssociatedAssetsInput) (*iotsitewise.ListAssociatedAssetsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListAssociatedAssetsWithContext(ctx, input)
}

func (a *IoTSiteWiseActivities) ListDashboards(ctx context.Context, input *iotsitewise.ListDashboardsInput) (*iotsitewise.ListDashboardsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListDashboardsWithContext(ctx, input)
}

func (a *IoTSiteWiseActivities) ListGateways(ctx context.Context, input *iotsitewise.ListGatewaysInput) (*iotsitewise.ListGatewaysOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListGatewaysWithContext(ctx, input)
}

func (a *IoTSiteWiseActivities) ListPortals(ctx context.Context, input *iotsitewise.ListPortalsInput) (*iotsitewise.ListPortalsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListPortalsWithContext(ctx, input)
}

func (a *IoTSiteWiseActivities) ListProjectAssets(ctx context.Context, input *iotsitewise.ListProjectAssetsInput) (*iotsitewise.ListProjectAssetsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListProjectAssetsWithContext(ctx, input)
}

func (a *IoTSiteWiseActivities) ListProjects(ctx context.Context, input *iotsitewise.ListProjectsInput) (*iotsitewise.ListProjectsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListProjectsWithContext(ctx, input)
}

func (a *IoTSiteWiseActivities) ListTagsForResource(ctx context.Context, input *iotsitewise.ListTagsForResourceInput) (*iotsitewise.ListTagsForResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListTagsForResourceWithContext(ctx, input)
}

func (a *IoTSiteWiseActivities) PutLoggingOptions(ctx context.Context, input *iotsitewise.PutLoggingOptionsInput) (*iotsitewise.PutLoggingOptionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.PutLoggingOptionsWithContext(ctx, input)
}

func (a *IoTSiteWiseActivities) TagResource(ctx context.Context, input *iotsitewise.TagResourceInput) (*iotsitewise.TagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.TagResourceWithContext(ctx, input)
}

func (a *IoTSiteWiseActivities) UntagResource(ctx context.Context, input *iotsitewise.UntagResourceInput) (*iotsitewise.UntagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UntagResourceWithContext(ctx, input)
}

func (a *IoTSiteWiseActivities) UpdateAccessPolicy(ctx context.Context, input *iotsitewise.UpdateAccessPolicyInput) (*iotsitewise.UpdateAccessPolicyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	internal.SetClientToken(ctx, &input.ClientToken)
	return client.UpdateAccessPolicyWithContext(ctx, input)
}

func (a *IoTSiteWiseActivities) UpdateAsset(ctx context.Context, input *iotsitewise.UpdateAssetInput) (*iotsitewise.UpdateAssetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	internal.SetClientToken(ctx, &input.ClientToken)
	return client.UpdateAssetWithContext(ctx, input)
}

func (a *IoTSiteWiseActivities) UpdateAssetModel(ctx context.Context, input *iotsitewise.UpdateAssetModelInput) (*iotsitewise.UpdateAssetModelOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	internal.SetClientToken(ctx, &input.ClientToken)
	return client.UpdateAssetModelWithContext(ctx, input)
}

func (a *IoTSiteWiseActivities) UpdateAssetProperty(ctx context.Context, input *iotsitewise.UpdateAssetPropertyInput) (*iotsitewise.UpdateAssetPropertyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	internal.SetClientToken(ctx, &input.ClientToken)
	return client.UpdateAssetPropertyWithContext(ctx, input)
}

func (a *IoTSiteWiseActivities) UpdateDashboard(ctx context.Context, input *iotsitewise.UpdateDashboardInput) (*iotsitewise.UpdateDashboardOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	internal.SetClientToken(ctx, &input.ClientToken)
	return client.UpdateDashboardWithContext(ctx, input)
}

func (a *IoTSiteWiseActivities) UpdateGateway(ctx context.Context, input *iotsitewise.UpdateGatewayInput) (*iotsitewise.UpdateGatewayOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateGatewayWithContext(ctx, input)
}

func (a *IoTSiteWiseActivities) UpdateGatewayCapabilityConfiguration(ctx context.Context, input *iotsitewise.UpdateGatewayCapabilityConfigurationInput) (*iotsitewise.UpdateGatewayCapabilityConfigurationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateGatewayCapabilityConfigurationWithContext(ctx, input)
}

func (a *IoTSiteWiseActivities) UpdatePortal(ctx context.Context, input *iotsitewise.UpdatePortalInput) (*iotsitewise.UpdatePortalOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	internal.SetClientToken(ctx, &input.ClientToken)
	return client.UpdatePortalWithContext(ctx, input)
}

func (a *IoTSiteWiseActivities) UpdateProject(ctx context.Context, input *iotsitewise.UpdateProjectInput) (*iotsitewise.UpdateProjectOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	internal.SetClientToken(ctx, &input.ClientToken)
	return client.UpdateProjectWithContext(ctx, input)
}

func (a *IoTSiteWiseActivities) WaitUntilAssetActive(ctx context.Context, input *iotsitewise.DescribeAssetInput) error {
	client, err := a.getClient(ctx)
	if err != nil {
		return err
	}
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return client.WaitUntilAssetActiveWithContext(ctx, input, options...)
	})
}

func (a *IoTSiteWiseActivities) WaitUntilAssetModelActive(ctx context.Context, input *iotsitewise.DescribeAssetModelInput) error {
	client, err := a.getClient(ctx)
	if err != nil {
		return err
	}
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return client.WaitUntilAssetModelActiveWithContext(ctx, input, options...)
	})
}

func (a *IoTSiteWiseActivities) WaitUntilAssetModelNotExists(ctx context.Context, input *iotsitewise.DescribeAssetModelInput) error {
	client, err := a.getClient(ctx)
	if err != nil {
		return err
	}
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return client.WaitUntilAssetModelNotExistsWithContext(ctx, input, options...)
	})
}

func (a *IoTSiteWiseActivities) WaitUntilAssetNotExists(ctx context.Context, input *iotsitewise.DescribeAssetInput) error {
	client, err := a.getClient(ctx)
	if err != nil {
		return err
	}
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return client.WaitUntilAssetNotExistsWithContext(ctx, input, options...)
	})
}

func (a *IoTSiteWiseActivities) WaitUntilPortalActive(ctx context.Context, input *iotsitewise.DescribePortalInput) error {
	client, err := a.getClient(ctx)
	if err != nil {
		return err
	}
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return client.WaitUntilPortalActiveWithContext(ctx, input, options...)
	})
}

func (a *IoTSiteWiseActivities) WaitUntilPortalNotExists(ctx context.Context, input *iotsitewise.DescribePortalInput) error {
	client, err := a.getClient(ctx)
	if err != nil {
		return err
	}
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return client.WaitUntilPortalNotExistsWithContext(ctx, input, options...)
	})
}
