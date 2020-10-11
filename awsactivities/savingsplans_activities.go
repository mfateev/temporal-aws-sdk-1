// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/savingsplans"
	"github.com/aws/aws-sdk-go/service/savingsplans/savingsplansiface"
	"go.temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken

type _ request.Option

type SavingsPlansActivities struct {
	client savingsplansiface.SavingsPlansAPI

	sessionFactory SessionFactory
}

func NewSavingsPlansActivities(sess *session.Session, config ...*aws.Config) *SavingsPlansActivities {
	client := savingsplans.New(sess, config...)
	return &SavingsPlansActivities{client: client}
}

func NewSavingsPlansActivitiesWithSessionFactory(sessionFactory SessionFactory) *SavingsPlansActivities {
	return &SavingsPlansActivities{sessionFactory: sessionFactory}
}

func (a *SavingsPlansActivities) getClient(ctx context.Context) (savingsplansiface.SavingsPlansAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, err
	}

	return savingsplans.New(sess), nil
}

func (a *SavingsPlansActivities) CreateSavingsPlan(ctx context.Context, input *savingsplans.CreateSavingsPlanInput) (*savingsplans.CreateSavingsPlanOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	internal.SetClientToken(ctx, &input.ClientToken)
	return client.CreateSavingsPlanWithContext(ctx, input)
}

func (a *SavingsPlansActivities) DeleteQueuedSavingsPlan(ctx context.Context, input *savingsplans.DeleteQueuedSavingsPlanInput) (*savingsplans.DeleteQueuedSavingsPlanOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteQueuedSavingsPlanWithContext(ctx, input)
}

func (a *SavingsPlansActivities) DescribeSavingsPlanRates(ctx context.Context, input *savingsplans.DescribeSavingsPlanRatesInput) (*savingsplans.DescribeSavingsPlanRatesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeSavingsPlanRatesWithContext(ctx, input)
}

func (a *SavingsPlansActivities) DescribeSavingsPlans(ctx context.Context, input *savingsplans.DescribeSavingsPlansInput) (*savingsplans.DescribeSavingsPlansOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeSavingsPlansWithContext(ctx, input)
}

func (a *SavingsPlansActivities) DescribeSavingsPlansOfferingRates(ctx context.Context, input *savingsplans.DescribeSavingsPlansOfferingRatesInput) (*savingsplans.DescribeSavingsPlansOfferingRatesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeSavingsPlansOfferingRatesWithContext(ctx, input)
}

func (a *SavingsPlansActivities) DescribeSavingsPlansOfferings(ctx context.Context, input *savingsplans.DescribeSavingsPlansOfferingsInput) (*savingsplans.DescribeSavingsPlansOfferingsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeSavingsPlansOfferingsWithContext(ctx, input)
}

func (a *SavingsPlansActivities) ListTagsForResource(ctx context.Context, input *savingsplans.ListTagsForResourceInput) (*savingsplans.ListTagsForResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListTagsForResourceWithContext(ctx, input)
}

func (a *SavingsPlansActivities) TagResource(ctx context.Context, input *savingsplans.TagResourceInput) (*savingsplans.TagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.TagResourceWithContext(ctx, input)
}

func (a *SavingsPlansActivities) UntagResource(ctx context.Context, input *savingsplans.UntagResourceInput) (*savingsplans.UntagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UntagResourceWithContext(ctx, input)
}
