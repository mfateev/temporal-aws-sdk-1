// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/guardduty"
	"github.com/aws/aws-sdk-go/service/guardduty/guarddutyiface"
	"go.temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken

type _ request.Option

type GuardDutyActivities struct {
	client guarddutyiface.GuardDutyAPI

	sessionFactory SessionFactory
}

func NewGuardDutyActivities(sess *session.Session, config ...*aws.Config) *GuardDutyActivities {
	client := guardduty.New(sess, config...)
	return &GuardDutyActivities{client: client}
}

func NewGuardDutyActivitiesWithSessionFactory(sessionFactory SessionFactory) *GuardDutyActivities {
	return &GuardDutyActivities{sessionFactory: sessionFactory}
}

func (a *GuardDutyActivities) getClient(ctx context.Context) (guarddutyiface.GuardDutyAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, err
	}

	return guardduty.New(sess), nil
}

func (a *GuardDutyActivities) AcceptInvitation(ctx context.Context, input *guardduty.AcceptInvitationInput) (*guardduty.AcceptInvitationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.AcceptInvitationWithContext(ctx, input)
}

func (a *GuardDutyActivities) ArchiveFindings(ctx context.Context, input *guardduty.ArchiveFindingsInput) (*guardduty.ArchiveFindingsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ArchiveFindingsWithContext(ctx, input)
}

func (a *GuardDutyActivities) CreateDetector(ctx context.Context, input *guardduty.CreateDetectorInput) (*guardduty.CreateDetectorOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	internal.SetClientToken(ctx, &input.ClientToken)
	return client.CreateDetectorWithContext(ctx, input)
}

func (a *GuardDutyActivities) CreateFilter(ctx context.Context, input *guardduty.CreateFilterInput) (*guardduty.CreateFilterOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	internal.SetClientToken(ctx, &input.ClientToken)
	return client.CreateFilterWithContext(ctx, input)
}

func (a *GuardDutyActivities) CreateIPSet(ctx context.Context, input *guardduty.CreateIPSetInput) (*guardduty.CreateIPSetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	internal.SetClientToken(ctx, &input.ClientToken)
	return client.CreateIPSetWithContext(ctx, input)
}

func (a *GuardDutyActivities) CreateMembers(ctx context.Context, input *guardduty.CreateMembersInput) (*guardduty.CreateMembersOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateMembersWithContext(ctx, input)
}

func (a *GuardDutyActivities) CreatePublishingDestination(ctx context.Context, input *guardduty.CreatePublishingDestinationInput) (*guardduty.CreatePublishingDestinationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	internal.SetClientToken(ctx, &input.ClientToken)
	return client.CreatePublishingDestinationWithContext(ctx, input)
}

func (a *GuardDutyActivities) CreateSampleFindings(ctx context.Context, input *guardduty.CreateSampleFindingsInput) (*guardduty.CreateSampleFindingsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateSampleFindingsWithContext(ctx, input)
}

func (a *GuardDutyActivities) CreateThreatIntelSet(ctx context.Context, input *guardduty.CreateThreatIntelSetInput) (*guardduty.CreateThreatIntelSetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	internal.SetClientToken(ctx, &input.ClientToken)
	return client.CreateThreatIntelSetWithContext(ctx, input)
}

func (a *GuardDutyActivities) DeclineInvitations(ctx context.Context, input *guardduty.DeclineInvitationsInput) (*guardduty.DeclineInvitationsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeclineInvitationsWithContext(ctx, input)
}

func (a *GuardDutyActivities) DeleteDetector(ctx context.Context, input *guardduty.DeleteDetectorInput) (*guardduty.DeleteDetectorOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteDetectorWithContext(ctx, input)
}

func (a *GuardDutyActivities) DeleteFilter(ctx context.Context, input *guardduty.DeleteFilterInput) (*guardduty.DeleteFilterOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteFilterWithContext(ctx, input)
}

func (a *GuardDutyActivities) DeleteIPSet(ctx context.Context, input *guardduty.DeleteIPSetInput) (*guardduty.DeleteIPSetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteIPSetWithContext(ctx, input)
}

func (a *GuardDutyActivities) DeleteInvitations(ctx context.Context, input *guardduty.DeleteInvitationsInput) (*guardduty.DeleteInvitationsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteInvitationsWithContext(ctx, input)
}

func (a *GuardDutyActivities) DeleteMembers(ctx context.Context, input *guardduty.DeleteMembersInput) (*guardduty.DeleteMembersOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteMembersWithContext(ctx, input)
}

func (a *GuardDutyActivities) DeletePublishingDestination(ctx context.Context, input *guardduty.DeletePublishingDestinationInput) (*guardduty.DeletePublishingDestinationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeletePublishingDestinationWithContext(ctx, input)
}

func (a *GuardDutyActivities) DeleteThreatIntelSet(ctx context.Context, input *guardduty.DeleteThreatIntelSetInput) (*guardduty.DeleteThreatIntelSetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteThreatIntelSetWithContext(ctx, input)
}

func (a *GuardDutyActivities) DescribeOrganizationConfiguration(ctx context.Context, input *guardduty.DescribeOrganizationConfigurationInput) (*guardduty.DescribeOrganizationConfigurationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeOrganizationConfigurationWithContext(ctx, input)
}

func (a *GuardDutyActivities) DescribePublishingDestination(ctx context.Context, input *guardduty.DescribePublishingDestinationInput) (*guardduty.DescribePublishingDestinationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribePublishingDestinationWithContext(ctx, input)
}

func (a *GuardDutyActivities) DisableOrganizationAdminAccount(ctx context.Context, input *guardduty.DisableOrganizationAdminAccountInput) (*guardduty.DisableOrganizationAdminAccountOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DisableOrganizationAdminAccountWithContext(ctx, input)
}

func (a *GuardDutyActivities) DisassociateFromMasterAccount(ctx context.Context, input *guardduty.DisassociateFromMasterAccountInput) (*guardduty.DisassociateFromMasterAccountOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DisassociateFromMasterAccountWithContext(ctx, input)
}

func (a *GuardDutyActivities) DisassociateMembers(ctx context.Context, input *guardduty.DisassociateMembersInput) (*guardduty.DisassociateMembersOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DisassociateMembersWithContext(ctx, input)
}

func (a *GuardDutyActivities) EnableOrganizationAdminAccount(ctx context.Context, input *guardduty.EnableOrganizationAdminAccountInput) (*guardduty.EnableOrganizationAdminAccountOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.EnableOrganizationAdminAccountWithContext(ctx, input)
}

func (a *GuardDutyActivities) GetDetector(ctx context.Context, input *guardduty.GetDetectorInput) (*guardduty.GetDetectorOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetDetectorWithContext(ctx, input)
}

func (a *GuardDutyActivities) GetFilter(ctx context.Context, input *guardduty.GetFilterInput) (*guardduty.GetFilterOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetFilterWithContext(ctx, input)
}

func (a *GuardDutyActivities) GetFindings(ctx context.Context, input *guardduty.GetFindingsInput) (*guardduty.GetFindingsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetFindingsWithContext(ctx, input)
}

func (a *GuardDutyActivities) GetFindingsStatistics(ctx context.Context, input *guardduty.GetFindingsStatisticsInput) (*guardduty.GetFindingsStatisticsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetFindingsStatisticsWithContext(ctx, input)
}

func (a *GuardDutyActivities) GetIPSet(ctx context.Context, input *guardduty.GetIPSetInput) (*guardduty.GetIPSetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetIPSetWithContext(ctx, input)
}

func (a *GuardDutyActivities) GetInvitationsCount(ctx context.Context, input *guardduty.GetInvitationsCountInput) (*guardduty.GetInvitationsCountOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetInvitationsCountWithContext(ctx, input)
}

func (a *GuardDutyActivities) GetMasterAccount(ctx context.Context, input *guardduty.GetMasterAccountInput) (*guardduty.GetMasterAccountOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetMasterAccountWithContext(ctx, input)
}

func (a *GuardDutyActivities) GetMemberDetectors(ctx context.Context, input *guardduty.GetMemberDetectorsInput) (*guardduty.GetMemberDetectorsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetMemberDetectorsWithContext(ctx, input)
}

func (a *GuardDutyActivities) GetMembers(ctx context.Context, input *guardduty.GetMembersInput) (*guardduty.GetMembersOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetMembersWithContext(ctx, input)
}

func (a *GuardDutyActivities) GetThreatIntelSet(ctx context.Context, input *guardduty.GetThreatIntelSetInput) (*guardduty.GetThreatIntelSetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetThreatIntelSetWithContext(ctx, input)
}

func (a *GuardDutyActivities) GetUsageStatistics(ctx context.Context, input *guardduty.GetUsageStatisticsInput) (*guardduty.GetUsageStatisticsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetUsageStatisticsWithContext(ctx, input)
}

func (a *GuardDutyActivities) InviteMembers(ctx context.Context, input *guardduty.InviteMembersInput) (*guardduty.InviteMembersOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.InviteMembersWithContext(ctx, input)
}

func (a *GuardDutyActivities) ListDetectors(ctx context.Context, input *guardduty.ListDetectorsInput) (*guardduty.ListDetectorsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListDetectorsWithContext(ctx, input)
}

func (a *GuardDutyActivities) ListFilters(ctx context.Context, input *guardduty.ListFiltersInput) (*guardduty.ListFiltersOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListFiltersWithContext(ctx, input)
}

func (a *GuardDutyActivities) ListFindings(ctx context.Context, input *guardduty.ListFindingsInput) (*guardduty.ListFindingsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListFindingsWithContext(ctx, input)
}

func (a *GuardDutyActivities) ListIPSets(ctx context.Context, input *guardduty.ListIPSetsInput) (*guardduty.ListIPSetsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListIPSetsWithContext(ctx, input)
}

func (a *GuardDutyActivities) ListInvitations(ctx context.Context, input *guardduty.ListInvitationsInput) (*guardduty.ListInvitationsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListInvitationsWithContext(ctx, input)
}

func (a *GuardDutyActivities) ListMembers(ctx context.Context, input *guardduty.ListMembersInput) (*guardduty.ListMembersOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListMembersWithContext(ctx, input)
}

func (a *GuardDutyActivities) ListOrganizationAdminAccounts(ctx context.Context, input *guardduty.ListOrganizationAdminAccountsInput) (*guardduty.ListOrganizationAdminAccountsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListOrganizationAdminAccountsWithContext(ctx, input)
}

func (a *GuardDutyActivities) ListPublishingDestinations(ctx context.Context, input *guardduty.ListPublishingDestinationsInput) (*guardduty.ListPublishingDestinationsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListPublishingDestinationsWithContext(ctx, input)
}

func (a *GuardDutyActivities) ListTagsForResource(ctx context.Context, input *guardduty.ListTagsForResourceInput) (*guardduty.ListTagsForResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListTagsForResourceWithContext(ctx, input)
}

func (a *GuardDutyActivities) ListThreatIntelSets(ctx context.Context, input *guardduty.ListThreatIntelSetsInput) (*guardduty.ListThreatIntelSetsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListThreatIntelSetsWithContext(ctx, input)
}

func (a *GuardDutyActivities) StartMonitoringMembers(ctx context.Context, input *guardduty.StartMonitoringMembersInput) (*guardduty.StartMonitoringMembersOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.StartMonitoringMembersWithContext(ctx, input)
}

func (a *GuardDutyActivities) StopMonitoringMembers(ctx context.Context, input *guardduty.StopMonitoringMembersInput) (*guardduty.StopMonitoringMembersOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.StopMonitoringMembersWithContext(ctx, input)
}

func (a *GuardDutyActivities) TagResource(ctx context.Context, input *guardduty.TagResourceInput) (*guardduty.TagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.TagResourceWithContext(ctx, input)
}

func (a *GuardDutyActivities) UnarchiveFindings(ctx context.Context, input *guardduty.UnarchiveFindingsInput) (*guardduty.UnarchiveFindingsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UnarchiveFindingsWithContext(ctx, input)
}

func (a *GuardDutyActivities) UntagResource(ctx context.Context, input *guardduty.UntagResourceInput) (*guardduty.UntagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UntagResourceWithContext(ctx, input)
}

func (a *GuardDutyActivities) UpdateDetector(ctx context.Context, input *guardduty.UpdateDetectorInput) (*guardduty.UpdateDetectorOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateDetectorWithContext(ctx, input)
}

func (a *GuardDutyActivities) UpdateFilter(ctx context.Context, input *guardduty.UpdateFilterInput) (*guardduty.UpdateFilterOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateFilterWithContext(ctx, input)
}

func (a *GuardDutyActivities) UpdateFindingsFeedback(ctx context.Context, input *guardduty.UpdateFindingsFeedbackInput) (*guardduty.UpdateFindingsFeedbackOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateFindingsFeedbackWithContext(ctx, input)
}

func (a *GuardDutyActivities) UpdateIPSet(ctx context.Context, input *guardduty.UpdateIPSetInput) (*guardduty.UpdateIPSetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateIPSetWithContext(ctx, input)
}

func (a *GuardDutyActivities) UpdateMemberDetectors(ctx context.Context, input *guardduty.UpdateMemberDetectorsInput) (*guardduty.UpdateMemberDetectorsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateMemberDetectorsWithContext(ctx, input)
}

func (a *GuardDutyActivities) UpdateOrganizationConfiguration(ctx context.Context, input *guardduty.UpdateOrganizationConfigurationInput) (*guardduty.UpdateOrganizationConfigurationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateOrganizationConfigurationWithContext(ctx, input)
}

func (a *GuardDutyActivities) UpdatePublishingDestination(ctx context.Context, input *guardduty.UpdatePublishingDestinationInput) (*guardduty.UpdatePublishingDestinationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdatePublishingDestinationWithContext(ctx, input)
}

func (a *GuardDutyActivities) UpdateThreatIntelSet(ctx context.Context, input *guardduty.UpdateThreatIntelSetInput) (*guardduty.UpdateThreatIntelSetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateThreatIntelSetWithContext(ctx, input)
}
