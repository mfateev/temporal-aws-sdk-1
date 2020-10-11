// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iotthingsgraph"
	"github.com/aws/aws-sdk-go/service/iotthingsgraph/iotthingsgraphiface"
	"go.temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken

type _ request.Option

type IoTThingsGraphActivities struct {
	client iotthingsgraphiface.IoTThingsGraphAPI

	sessionFactory SessionFactory
}

func NewIoTThingsGraphActivities(sess *session.Session, config ...*aws.Config) *IoTThingsGraphActivities {
	client := iotthingsgraph.New(sess, config...)
	return &IoTThingsGraphActivities{client: client}
}

func NewIoTThingsGraphActivitiesWithSessionFactory(sessionFactory SessionFactory) *IoTThingsGraphActivities {
	return &IoTThingsGraphActivities{sessionFactory: sessionFactory}
}

func (a *IoTThingsGraphActivities) getClient(ctx context.Context) (iotthingsgraphiface.IoTThingsGraphAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, err
	}

	return iotthingsgraph.New(sess), nil
}

func (a *IoTThingsGraphActivities) AssociateEntityToThing(ctx context.Context, input *iotthingsgraph.AssociateEntityToThingInput) (*iotthingsgraph.AssociateEntityToThingOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.AssociateEntityToThingWithContext(ctx, input)
}

func (a *IoTThingsGraphActivities) CreateFlowTemplate(ctx context.Context, input *iotthingsgraph.CreateFlowTemplateInput) (*iotthingsgraph.CreateFlowTemplateOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateFlowTemplateWithContext(ctx, input)
}

func (a *IoTThingsGraphActivities) CreateSystemInstance(ctx context.Context, input *iotthingsgraph.CreateSystemInstanceInput) (*iotthingsgraph.CreateSystemInstanceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateSystemInstanceWithContext(ctx, input)
}

func (a *IoTThingsGraphActivities) CreateSystemTemplate(ctx context.Context, input *iotthingsgraph.CreateSystemTemplateInput) (*iotthingsgraph.CreateSystemTemplateOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateSystemTemplateWithContext(ctx, input)
}

func (a *IoTThingsGraphActivities) DeleteFlowTemplate(ctx context.Context, input *iotthingsgraph.DeleteFlowTemplateInput) (*iotthingsgraph.DeleteFlowTemplateOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteFlowTemplateWithContext(ctx, input)
}

func (a *IoTThingsGraphActivities) DeleteNamespace(ctx context.Context, input *iotthingsgraph.DeleteNamespaceInput) (*iotthingsgraph.DeleteNamespaceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteNamespaceWithContext(ctx, input)
}

func (a *IoTThingsGraphActivities) DeleteSystemInstance(ctx context.Context, input *iotthingsgraph.DeleteSystemInstanceInput) (*iotthingsgraph.DeleteSystemInstanceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteSystemInstanceWithContext(ctx, input)
}

func (a *IoTThingsGraphActivities) DeleteSystemTemplate(ctx context.Context, input *iotthingsgraph.DeleteSystemTemplateInput) (*iotthingsgraph.DeleteSystemTemplateOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteSystemTemplateWithContext(ctx, input)
}

func (a *IoTThingsGraphActivities) DeploySystemInstance(ctx context.Context, input *iotthingsgraph.DeploySystemInstanceInput) (*iotthingsgraph.DeploySystemInstanceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeploySystemInstanceWithContext(ctx, input)
}

func (a *IoTThingsGraphActivities) DeprecateFlowTemplate(ctx context.Context, input *iotthingsgraph.DeprecateFlowTemplateInput) (*iotthingsgraph.DeprecateFlowTemplateOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeprecateFlowTemplateWithContext(ctx, input)
}

func (a *IoTThingsGraphActivities) DeprecateSystemTemplate(ctx context.Context, input *iotthingsgraph.DeprecateSystemTemplateInput) (*iotthingsgraph.DeprecateSystemTemplateOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeprecateSystemTemplateWithContext(ctx, input)
}

func (a *IoTThingsGraphActivities) DescribeNamespace(ctx context.Context, input *iotthingsgraph.DescribeNamespaceInput) (*iotthingsgraph.DescribeNamespaceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeNamespaceWithContext(ctx, input)
}

func (a *IoTThingsGraphActivities) DissociateEntityFromThing(ctx context.Context, input *iotthingsgraph.DissociateEntityFromThingInput) (*iotthingsgraph.DissociateEntityFromThingOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DissociateEntityFromThingWithContext(ctx, input)
}

func (a *IoTThingsGraphActivities) GetEntities(ctx context.Context, input *iotthingsgraph.GetEntitiesInput) (*iotthingsgraph.GetEntitiesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetEntitiesWithContext(ctx, input)
}

func (a *IoTThingsGraphActivities) GetFlowTemplate(ctx context.Context, input *iotthingsgraph.GetFlowTemplateInput) (*iotthingsgraph.GetFlowTemplateOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetFlowTemplateWithContext(ctx, input)
}

func (a *IoTThingsGraphActivities) GetFlowTemplateRevisions(ctx context.Context, input *iotthingsgraph.GetFlowTemplateRevisionsInput) (*iotthingsgraph.GetFlowTemplateRevisionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetFlowTemplateRevisionsWithContext(ctx, input)
}

func (a *IoTThingsGraphActivities) GetNamespaceDeletionStatus(ctx context.Context, input *iotthingsgraph.GetNamespaceDeletionStatusInput) (*iotthingsgraph.GetNamespaceDeletionStatusOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetNamespaceDeletionStatusWithContext(ctx, input)
}

func (a *IoTThingsGraphActivities) GetSystemInstance(ctx context.Context, input *iotthingsgraph.GetSystemInstanceInput) (*iotthingsgraph.GetSystemInstanceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetSystemInstanceWithContext(ctx, input)
}

func (a *IoTThingsGraphActivities) GetSystemTemplate(ctx context.Context, input *iotthingsgraph.GetSystemTemplateInput) (*iotthingsgraph.GetSystemTemplateOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetSystemTemplateWithContext(ctx, input)
}

func (a *IoTThingsGraphActivities) GetSystemTemplateRevisions(ctx context.Context, input *iotthingsgraph.GetSystemTemplateRevisionsInput) (*iotthingsgraph.GetSystemTemplateRevisionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetSystemTemplateRevisionsWithContext(ctx, input)
}

func (a *IoTThingsGraphActivities) GetUploadStatus(ctx context.Context, input *iotthingsgraph.GetUploadStatusInput) (*iotthingsgraph.GetUploadStatusOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetUploadStatusWithContext(ctx, input)
}

func (a *IoTThingsGraphActivities) ListFlowExecutionMessages(ctx context.Context, input *iotthingsgraph.ListFlowExecutionMessagesInput) (*iotthingsgraph.ListFlowExecutionMessagesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListFlowExecutionMessagesWithContext(ctx, input)
}

func (a *IoTThingsGraphActivities) ListTagsForResource(ctx context.Context, input *iotthingsgraph.ListTagsForResourceInput) (*iotthingsgraph.ListTagsForResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListTagsForResourceWithContext(ctx, input)
}

func (a *IoTThingsGraphActivities) SearchEntities(ctx context.Context, input *iotthingsgraph.SearchEntitiesInput) (*iotthingsgraph.SearchEntitiesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.SearchEntitiesWithContext(ctx, input)
}

func (a *IoTThingsGraphActivities) SearchFlowExecutions(ctx context.Context, input *iotthingsgraph.SearchFlowExecutionsInput) (*iotthingsgraph.SearchFlowExecutionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.SearchFlowExecutionsWithContext(ctx, input)
}

func (a *IoTThingsGraphActivities) SearchFlowTemplates(ctx context.Context, input *iotthingsgraph.SearchFlowTemplatesInput) (*iotthingsgraph.SearchFlowTemplatesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.SearchFlowTemplatesWithContext(ctx, input)
}

func (a *IoTThingsGraphActivities) SearchSystemInstances(ctx context.Context, input *iotthingsgraph.SearchSystemInstancesInput) (*iotthingsgraph.SearchSystemInstancesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.SearchSystemInstancesWithContext(ctx, input)
}

func (a *IoTThingsGraphActivities) SearchSystemTemplates(ctx context.Context, input *iotthingsgraph.SearchSystemTemplatesInput) (*iotthingsgraph.SearchSystemTemplatesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.SearchSystemTemplatesWithContext(ctx, input)
}

func (a *IoTThingsGraphActivities) SearchThings(ctx context.Context, input *iotthingsgraph.SearchThingsInput) (*iotthingsgraph.SearchThingsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.SearchThingsWithContext(ctx, input)
}

func (a *IoTThingsGraphActivities) TagResource(ctx context.Context, input *iotthingsgraph.TagResourceInput) (*iotthingsgraph.TagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.TagResourceWithContext(ctx, input)
}

func (a *IoTThingsGraphActivities) UndeploySystemInstance(ctx context.Context, input *iotthingsgraph.UndeploySystemInstanceInput) (*iotthingsgraph.UndeploySystemInstanceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UndeploySystemInstanceWithContext(ctx, input)
}

func (a *IoTThingsGraphActivities) UntagResource(ctx context.Context, input *iotthingsgraph.UntagResourceInput) (*iotthingsgraph.UntagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UntagResourceWithContext(ctx, input)
}

func (a *IoTThingsGraphActivities) UpdateFlowTemplate(ctx context.Context, input *iotthingsgraph.UpdateFlowTemplateInput) (*iotthingsgraph.UpdateFlowTemplateOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateFlowTemplateWithContext(ctx, input)
}

func (a *IoTThingsGraphActivities) UpdateSystemTemplate(ctx context.Context, input *iotthingsgraph.UpdateSystemTemplateInput) (*iotthingsgraph.UpdateSystemTemplateOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateSystemTemplateWithContext(ctx, input)
}

func (a *IoTThingsGraphActivities) UploadEntityDefinitions(ctx context.Context, input *iotthingsgraph.UploadEntityDefinitionsInput) (*iotthingsgraph.UploadEntityDefinitionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UploadEntityDefinitionsWithContext(ctx, input)
}
