// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/directoryservice"
	"github.com/aws/aws-sdk-go/service/directoryservice/directoryserviceiface"
	"go.temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken

type _ request.Option

type DirectoryServiceActivities struct {
	client directoryserviceiface.DirectoryServiceAPI

	sessionFactory SessionFactory
}

func NewDirectoryServiceActivities(sess *session.Session, config ...*aws.Config) *DirectoryServiceActivities {
	client := directoryservice.New(sess, config...)
	return &DirectoryServiceActivities{client: client}
}

func NewDirectoryServiceActivitiesWithSessionFactory(sessionFactory SessionFactory) *DirectoryServiceActivities {
	return &DirectoryServiceActivities{sessionFactory: sessionFactory}
}

func (a *DirectoryServiceActivities) getClient(ctx context.Context) (directoryserviceiface.DirectoryServiceAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, err
	}

	return directoryservice.New(sess), nil
}

func (a *DirectoryServiceActivities) AcceptSharedDirectory(ctx context.Context, input *directoryservice.AcceptSharedDirectoryInput) (*directoryservice.AcceptSharedDirectoryOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.AcceptSharedDirectoryWithContext(ctx, input)
}

func (a *DirectoryServiceActivities) AddIpRoutes(ctx context.Context, input *directoryservice.AddIpRoutesInput) (*directoryservice.AddIpRoutesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.AddIpRoutesWithContext(ctx, input)
}

func (a *DirectoryServiceActivities) AddTagsToResource(ctx context.Context, input *directoryservice.AddTagsToResourceInput) (*directoryservice.AddTagsToResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.AddTagsToResourceWithContext(ctx, input)
}

func (a *DirectoryServiceActivities) CancelSchemaExtension(ctx context.Context, input *directoryservice.CancelSchemaExtensionInput) (*directoryservice.CancelSchemaExtensionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CancelSchemaExtensionWithContext(ctx, input)
}

func (a *DirectoryServiceActivities) ConnectDirectory(ctx context.Context, input *directoryservice.ConnectDirectoryInput) (*directoryservice.ConnectDirectoryOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ConnectDirectoryWithContext(ctx, input)
}

func (a *DirectoryServiceActivities) CreateAlias(ctx context.Context, input *directoryservice.CreateAliasInput) (*directoryservice.CreateAliasOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateAliasWithContext(ctx, input)
}

func (a *DirectoryServiceActivities) CreateComputer(ctx context.Context, input *directoryservice.CreateComputerInput) (*directoryservice.CreateComputerOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateComputerWithContext(ctx, input)
}

func (a *DirectoryServiceActivities) CreateConditionalForwarder(ctx context.Context, input *directoryservice.CreateConditionalForwarderInput) (*directoryservice.CreateConditionalForwarderOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateConditionalForwarderWithContext(ctx, input)
}

func (a *DirectoryServiceActivities) CreateDirectory(ctx context.Context, input *directoryservice.CreateDirectoryInput) (*directoryservice.CreateDirectoryOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateDirectoryWithContext(ctx, input)
}

func (a *DirectoryServiceActivities) CreateLogSubscription(ctx context.Context, input *directoryservice.CreateLogSubscriptionInput) (*directoryservice.CreateLogSubscriptionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateLogSubscriptionWithContext(ctx, input)
}

func (a *DirectoryServiceActivities) CreateMicrosoftAD(ctx context.Context, input *directoryservice.CreateMicrosoftADInput) (*directoryservice.CreateMicrosoftADOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateMicrosoftADWithContext(ctx, input)
}

func (a *DirectoryServiceActivities) CreateSnapshot(ctx context.Context, input *directoryservice.CreateSnapshotInput) (*directoryservice.CreateSnapshotOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateSnapshotWithContext(ctx, input)
}

func (a *DirectoryServiceActivities) CreateTrust(ctx context.Context, input *directoryservice.CreateTrustInput) (*directoryservice.CreateTrustOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateTrustWithContext(ctx, input)
}

func (a *DirectoryServiceActivities) DeleteConditionalForwarder(ctx context.Context, input *directoryservice.DeleteConditionalForwarderInput) (*directoryservice.DeleteConditionalForwarderOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteConditionalForwarderWithContext(ctx, input)
}

func (a *DirectoryServiceActivities) DeleteDirectory(ctx context.Context, input *directoryservice.DeleteDirectoryInput) (*directoryservice.DeleteDirectoryOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteDirectoryWithContext(ctx, input)
}

func (a *DirectoryServiceActivities) DeleteLogSubscription(ctx context.Context, input *directoryservice.DeleteLogSubscriptionInput) (*directoryservice.DeleteLogSubscriptionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteLogSubscriptionWithContext(ctx, input)
}

func (a *DirectoryServiceActivities) DeleteSnapshot(ctx context.Context, input *directoryservice.DeleteSnapshotInput) (*directoryservice.DeleteSnapshotOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteSnapshotWithContext(ctx, input)
}

func (a *DirectoryServiceActivities) DeleteTrust(ctx context.Context, input *directoryservice.DeleteTrustInput) (*directoryservice.DeleteTrustOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteTrustWithContext(ctx, input)
}

func (a *DirectoryServiceActivities) DeregisterCertificate(ctx context.Context, input *directoryservice.DeregisterCertificateInput) (*directoryservice.DeregisterCertificateOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeregisterCertificateWithContext(ctx, input)
}

func (a *DirectoryServiceActivities) DeregisterEventTopic(ctx context.Context, input *directoryservice.DeregisterEventTopicInput) (*directoryservice.DeregisterEventTopicOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeregisterEventTopicWithContext(ctx, input)
}

func (a *DirectoryServiceActivities) DescribeCertificate(ctx context.Context, input *directoryservice.DescribeCertificateInput) (*directoryservice.DescribeCertificateOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeCertificateWithContext(ctx, input)
}

func (a *DirectoryServiceActivities) DescribeConditionalForwarders(ctx context.Context, input *directoryservice.DescribeConditionalForwardersInput) (*directoryservice.DescribeConditionalForwardersOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeConditionalForwardersWithContext(ctx, input)
}

func (a *DirectoryServiceActivities) DescribeDirectories(ctx context.Context, input *directoryservice.DescribeDirectoriesInput) (*directoryservice.DescribeDirectoriesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeDirectoriesWithContext(ctx, input)
}

func (a *DirectoryServiceActivities) DescribeDomainControllers(ctx context.Context, input *directoryservice.DescribeDomainControllersInput) (*directoryservice.DescribeDomainControllersOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeDomainControllersWithContext(ctx, input)
}

func (a *DirectoryServiceActivities) DescribeEventTopics(ctx context.Context, input *directoryservice.DescribeEventTopicsInput) (*directoryservice.DescribeEventTopicsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeEventTopicsWithContext(ctx, input)
}

func (a *DirectoryServiceActivities) DescribeLDAPSSettings(ctx context.Context, input *directoryservice.DescribeLDAPSSettingsInput) (*directoryservice.DescribeLDAPSSettingsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeLDAPSSettingsWithContext(ctx, input)
}

func (a *DirectoryServiceActivities) DescribeSharedDirectories(ctx context.Context, input *directoryservice.DescribeSharedDirectoriesInput) (*directoryservice.DescribeSharedDirectoriesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeSharedDirectoriesWithContext(ctx, input)
}

func (a *DirectoryServiceActivities) DescribeSnapshots(ctx context.Context, input *directoryservice.DescribeSnapshotsInput) (*directoryservice.DescribeSnapshotsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeSnapshotsWithContext(ctx, input)
}

func (a *DirectoryServiceActivities) DescribeTrusts(ctx context.Context, input *directoryservice.DescribeTrustsInput) (*directoryservice.DescribeTrustsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeTrustsWithContext(ctx, input)
}

func (a *DirectoryServiceActivities) DisableLDAPS(ctx context.Context, input *directoryservice.DisableLDAPSInput) (*directoryservice.DisableLDAPSOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DisableLDAPSWithContext(ctx, input)
}

func (a *DirectoryServiceActivities) DisableRadius(ctx context.Context, input *directoryservice.DisableRadiusInput) (*directoryservice.DisableRadiusOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DisableRadiusWithContext(ctx, input)
}

func (a *DirectoryServiceActivities) DisableSso(ctx context.Context, input *directoryservice.DisableSsoInput) (*directoryservice.DisableSsoOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DisableSsoWithContext(ctx, input)
}

func (a *DirectoryServiceActivities) EnableLDAPS(ctx context.Context, input *directoryservice.EnableLDAPSInput) (*directoryservice.EnableLDAPSOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.EnableLDAPSWithContext(ctx, input)
}

func (a *DirectoryServiceActivities) EnableRadius(ctx context.Context, input *directoryservice.EnableRadiusInput) (*directoryservice.EnableRadiusOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.EnableRadiusWithContext(ctx, input)
}

func (a *DirectoryServiceActivities) EnableSso(ctx context.Context, input *directoryservice.EnableSsoInput) (*directoryservice.EnableSsoOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.EnableSsoWithContext(ctx, input)
}

func (a *DirectoryServiceActivities) GetDirectoryLimits(ctx context.Context, input *directoryservice.GetDirectoryLimitsInput) (*directoryservice.GetDirectoryLimitsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetDirectoryLimitsWithContext(ctx, input)
}

func (a *DirectoryServiceActivities) GetSnapshotLimits(ctx context.Context, input *directoryservice.GetSnapshotLimitsInput) (*directoryservice.GetSnapshotLimitsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetSnapshotLimitsWithContext(ctx, input)
}

func (a *DirectoryServiceActivities) ListCertificates(ctx context.Context, input *directoryservice.ListCertificatesInput) (*directoryservice.ListCertificatesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListCertificatesWithContext(ctx, input)
}

func (a *DirectoryServiceActivities) ListIpRoutes(ctx context.Context, input *directoryservice.ListIpRoutesInput) (*directoryservice.ListIpRoutesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListIpRoutesWithContext(ctx, input)
}

func (a *DirectoryServiceActivities) ListLogSubscriptions(ctx context.Context, input *directoryservice.ListLogSubscriptionsInput) (*directoryservice.ListLogSubscriptionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListLogSubscriptionsWithContext(ctx, input)
}

func (a *DirectoryServiceActivities) ListSchemaExtensions(ctx context.Context, input *directoryservice.ListSchemaExtensionsInput) (*directoryservice.ListSchemaExtensionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListSchemaExtensionsWithContext(ctx, input)
}

func (a *DirectoryServiceActivities) ListTagsForResource(ctx context.Context, input *directoryservice.ListTagsForResourceInput) (*directoryservice.ListTagsForResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListTagsForResourceWithContext(ctx, input)
}

func (a *DirectoryServiceActivities) RegisterCertificate(ctx context.Context, input *directoryservice.RegisterCertificateInput) (*directoryservice.RegisterCertificateOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.RegisterCertificateWithContext(ctx, input)
}

func (a *DirectoryServiceActivities) RegisterEventTopic(ctx context.Context, input *directoryservice.RegisterEventTopicInput) (*directoryservice.RegisterEventTopicOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.RegisterEventTopicWithContext(ctx, input)
}

func (a *DirectoryServiceActivities) RejectSharedDirectory(ctx context.Context, input *directoryservice.RejectSharedDirectoryInput) (*directoryservice.RejectSharedDirectoryOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.RejectSharedDirectoryWithContext(ctx, input)
}

func (a *DirectoryServiceActivities) RemoveIpRoutes(ctx context.Context, input *directoryservice.RemoveIpRoutesInput) (*directoryservice.RemoveIpRoutesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.RemoveIpRoutesWithContext(ctx, input)
}

func (a *DirectoryServiceActivities) RemoveTagsFromResource(ctx context.Context, input *directoryservice.RemoveTagsFromResourceInput) (*directoryservice.RemoveTagsFromResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.RemoveTagsFromResourceWithContext(ctx, input)
}

func (a *DirectoryServiceActivities) ResetUserPassword(ctx context.Context, input *directoryservice.ResetUserPasswordInput) (*directoryservice.ResetUserPasswordOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ResetUserPasswordWithContext(ctx, input)
}

func (a *DirectoryServiceActivities) RestoreFromSnapshot(ctx context.Context, input *directoryservice.RestoreFromSnapshotInput) (*directoryservice.RestoreFromSnapshotOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.RestoreFromSnapshotWithContext(ctx, input)
}

func (a *DirectoryServiceActivities) ShareDirectory(ctx context.Context, input *directoryservice.ShareDirectoryInput) (*directoryservice.ShareDirectoryOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ShareDirectoryWithContext(ctx, input)
}

func (a *DirectoryServiceActivities) StartSchemaExtension(ctx context.Context, input *directoryservice.StartSchemaExtensionInput) (*directoryservice.StartSchemaExtensionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.StartSchemaExtensionWithContext(ctx, input)
}

func (a *DirectoryServiceActivities) UnshareDirectory(ctx context.Context, input *directoryservice.UnshareDirectoryInput) (*directoryservice.UnshareDirectoryOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UnshareDirectoryWithContext(ctx, input)
}

func (a *DirectoryServiceActivities) UpdateConditionalForwarder(ctx context.Context, input *directoryservice.UpdateConditionalForwarderInput) (*directoryservice.UpdateConditionalForwarderOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateConditionalForwarderWithContext(ctx, input)
}

func (a *DirectoryServiceActivities) UpdateNumberOfDomainControllers(ctx context.Context, input *directoryservice.UpdateNumberOfDomainControllersInput) (*directoryservice.UpdateNumberOfDomainControllersOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateNumberOfDomainControllersWithContext(ctx, input)
}

func (a *DirectoryServiceActivities) UpdateRadius(ctx context.Context, input *directoryservice.UpdateRadiusInput) (*directoryservice.UpdateRadiusOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateRadiusWithContext(ctx, input)
}

func (a *DirectoryServiceActivities) UpdateTrust(ctx context.Context, input *directoryservice.UpdateTrustInput) (*directoryservice.UpdateTrustOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateTrustWithContext(ctx, input)
}

func (a *DirectoryServiceActivities) VerifyTrust(ctx context.Context, input *directoryservice.VerifyTrustInput) (*directoryservice.VerifyTrustOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.VerifyTrustWithContext(ctx, input)
}
