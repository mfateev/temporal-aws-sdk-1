package awsclients

import (
	"github.com/aws/aws-sdk-go/service/health"
	"go.temporal.io/sdk/workflow"
	"temporal.io/aws-sdk/awsactivities"
)

type HealthClient interface {
	DescribeAffectedAccountsForOrganization(ctx workflow.Context, input *health.DescribeAffectedAccountsForOrganizationInput) (*health.DescribeAffectedAccountsForOrganizationOutput, error)
	DescribeAffectedAccountsForOrganizationAsync(ctx workflow.Context, input *health.DescribeAffectedAccountsForOrganizationInput) *HealthDescribeAffectedAccountsForOrganizationResult

	DescribeAffectedEntities(ctx workflow.Context, input *health.DescribeAffectedEntitiesInput) (*health.DescribeAffectedEntitiesOutput, error)
	DescribeAffectedEntitiesAsync(ctx workflow.Context, input *health.DescribeAffectedEntitiesInput) *HealthDescribeAffectedEntitiesResult

	DescribeAffectedEntitiesForOrganization(ctx workflow.Context, input *health.DescribeAffectedEntitiesForOrganizationInput) (*health.DescribeAffectedEntitiesForOrganizationOutput, error)
	DescribeAffectedEntitiesForOrganizationAsync(ctx workflow.Context, input *health.DescribeAffectedEntitiesForOrganizationInput) *HealthDescribeAffectedEntitiesForOrganizationResult

	DescribeEntityAggregates(ctx workflow.Context, input *health.DescribeEntityAggregatesInput) (*health.DescribeEntityAggregatesOutput, error)
	DescribeEntityAggregatesAsync(ctx workflow.Context, input *health.DescribeEntityAggregatesInput) *HealthDescribeEntityAggregatesResult

	DescribeEventAggregates(ctx workflow.Context, input *health.DescribeEventAggregatesInput) (*health.DescribeEventAggregatesOutput, error)
	DescribeEventAggregatesAsync(ctx workflow.Context, input *health.DescribeEventAggregatesInput) *HealthDescribeEventAggregatesResult

	DescribeEventDetails(ctx workflow.Context, input *health.DescribeEventDetailsInput) (*health.DescribeEventDetailsOutput, error)
	DescribeEventDetailsAsync(ctx workflow.Context, input *health.DescribeEventDetailsInput) *HealthDescribeEventDetailsResult

	DescribeEventDetailsForOrganization(ctx workflow.Context, input *health.DescribeEventDetailsForOrganizationInput) (*health.DescribeEventDetailsForOrganizationOutput, error)
	DescribeEventDetailsForOrganizationAsync(ctx workflow.Context, input *health.DescribeEventDetailsForOrganizationInput) *HealthDescribeEventDetailsForOrganizationResult

	DescribeEventTypes(ctx workflow.Context, input *health.DescribeEventTypesInput) (*health.DescribeEventTypesOutput, error)
	DescribeEventTypesAsync(ctx workflow.Context, input *health.DescribeEventTypesInput) *HealthDescribeEventTypesResult

	DescribeEvents(ctx workflow.Context, input *health.DescribeEventsInput) (*health.DescribeEventsOutput, error)
	DescribeEventsAsync(ctx workflow.Context, input *health.DescribeEventsInput) *HealthDescribeEventsResult

	DescribeEventsForOrganization(ctx workflow.Context, input *health.DescribeEventsForOrganizationInput) (*health.DescribeEventsForOrganizationOutput, error)
	DescribeEventsForOrganizationAsync(ctx workflow.Context, input *health.DescribeEventsForOrganizationInput) *HealthDescribeEventsForOrganizationResult

	DescribeHealthServiceStatusForOrganization(ctx workflow.Context, input *health.DescribeHealthServiceStatusForOrganizationInput) (*health.DescribeHealthServiceStatusForOrganizationOutput, error)
	DescribeHealthServiceStatusForOrganizationAsync(ctx workflow.Context, input *health.DescribeHealthServiceStatusForOrganizationInput) *HealthDescribeHealthServiceStatusForOrganizationResult

	DisableHealthServiceAccessForOrganization(ctx workflow.Context, input *health.DisableHealthServiceAccessForOrganizationInput) (*health.DisableHealthServiceAccessForOrganizationOutput, error)
	DisableHealthServiceAccessForOrganizationAsync(ctx workflow.Context, input *health.DisableHealthServiceAccessForOrganizationInput) *HealthDisableHealthServiceAccessForOrganizationResult

	EnableHealthServiceAccessForOrganization(ctx workflow.Context, input *health.EnableHealthServiceAccessForOrganizationInput) (*health.EnableHealthServiceAccessForOrganizationOutput, error)
	EnableHealthServiceAccessForOrganizationAsync(ctx workflow.Context, input *health.EnableHealthServiceAccessForOrganizationInput) *HealthEnableHealthServiceAccessForOrganizationResult
}

type HealthDescribeAffectedAccountsForOrganizationResult struct {
	Result workflow.Future
}

func (r *HealthDescribeAffectedAccountsForOrganizationResult) Get(ctx workflow.Context) (*health.DescribeAffectedAccountsForOrganizationOutput, error) {
	var output health.DescribeAffectedAccountsForOrganizationOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type HealthDescribeAffectedEntitiesResult struct {
	Result workflow.Future
}

func (r *HealthDescribeAffectedEntitiesResult) Get(ctx workflow.Context) (*health.DescribeAffectedEntitiesOutput, error) {
	var output health.DescribeAffectedEntitiesOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type HealthDescribeAffectedEntitiesForOrganizationResult struct {
	Result workflow.Future
}

func (r *HealthDescribeAffectedEntitiesForOrganizationResult) Get(ctx workflow.Context) (*health.DescribeAffectedEntitiesForOrganizationOutput, error) {
	var output health.DescribeAffectedEntitiesForOrganizationOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type HealthDescribeEntityAggregatesResult struct {
	Result workflow.Future
}

func (r *HealthDescribeEntityAggregatesResult) Get(ctx workflow.Context) (*health.DescribeEntityAggregatesOutput, error) {
	var output health.DescribeEntityAggregatesOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type HealthDescribeEventAggregatesResult struct {
	Result workflow.Future
}

func (r *HealthDescribeEventAggregatesResult) Get(ctx workflow.Context) (*health.DescribeEventAggregatesOutput, error) {
	var output health.DescribeEventAggregatesOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type HealthDescribeEventDetailsResult struct {
	Result workflow.Future
}

func (r *HealthDescribeEventDetailsResult) Get(ctx workflow.Context) (*health.DescribeEventDetailsOutput, error) {
	var output health.DescribeEventDetailsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type HealthDescribeEventDetailsForOrganizationResult struct {
	Result workflow.Future
}

func (r *HealthDescribeEventDetailsForOrganizationResult) Get(ctx workflow.Context) (*health.DescribeEventDetailsForOrganizationOutput, error) {
	var output health.DescribeEventDetailsForOrganizationOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type HealthDescribeEventTypesResult struct {
	Result workflow.Future
}

func (r *HealthDescribeEventTypesResult) Get(ctx workflow.Context) (*health.DescribeEventTypesOutput, error) {
	var output health.DescribeEventTypesOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type HealthDescribeEventsResult struct {
	Result workflow.Future
}

func (r *HealthDescribeEventsResult) Get(ctx workflow.Context) (*health.DescribeEventsOutput, error) {
	var output health.DescribeEventsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type HealthDescribeEventsForOrganizationResult struct {
	Result workflow.Future
}

func (r *HealthDescribeEventsForOrganizationResult) Get(ctx workflow.Context) (*health.DescribeEventsForOrganizationOutput, error) {
	var output health.DescribeEventsForOrganizationOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type HealthDescribeHealthServiceStatusForOrganizationResult struct {
	Result workflow.Future
}

func (r *HealthDescribeHealthServiceStatusForOrganizationResult) Get(ctx workflow.Context) (*health.DescribeHealthServiceStatusForOrganizationOutput, error) {
	var output health.DescribeHealthServiceStatusForOrganizationOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type HealthDisableHealthServiceAccessForOrganizationResult struct {
	Result workflow.Future
}

func (r *HealthDisableHealthServiceAccessForOrganizationResult) Get(ctx workflow.Context) (*health.DisableHealthServiceAccessForOrganizationOutput, error) {
	var output health.DisableHealthServiceAccessForOrganizationOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type HealthEnableHealthServiceAccessForOrganizationResult struct {
	Result workflow.Future
}

func (r *HealthEnableHealthServiceAccessForOrganizationResult) Get(ctx workflow.Context) (*health.EnableHealthServiceAccessForOrganizationOutput, error) {
	var output health.EnableHealthServiceAccessForOrganizationOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type HealthStub struct {
	activities awsactivities.HealthActivities
}

func NewHealthStub() HealthClient {
	return &HealthStub{}
}

func (a *HealthStub) DescribeAffectedAccountsForOrganization(ctx workflow.Context, input *health.DescribeAffectedAccountsForOrganizationInput) (*health.DescribeAffectedAccountsForOrganizationOutput, error) {
	var output health.DescribeAffectedAccountsForOrganizationOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DescribeAffectedAccountsForOrganization, input).Get(ctx, &output)
	return &output, err
}

func (a *HealthStub) DescribeAffectedAccountsForOrganizationAsync(ctx workflow.Context, input *health.DescribeAffectedAccountsForOrganizationInput) *HealthDescribeAffectedAccountsForOrganizationResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DescribeAffectedAccountsForOrganization, input)
	return &HealthDescribeAffectedAccountsForOrganizationResult{Result: future}
}

func (a *HealthStub) DescribeAffectedEntities(ctx workflow.Context, input *health.DescribeAffectedEntitiesInput) (*health.DescribeAffectedEntitiesOutput, error) {
	var output health.DescribeAffectedEntitiesOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DescribeAffectedEntities, input).Get(ctx, &output)
	return &output, err
}

func (a *HealthStub) DescribeAffectedEntitiesAsync(ctx workflow.Context, input *health.DescribeAffectedEntitiesInput) *HealthDescribeAffectedEntitiesResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DescribeAffectedEntities, input)
	return &HealthDescribeAffectedEntitiesResult{Result: future}
}

func (a *HealthStub) DescribeAffectedEntitiesForOrganization(ctx workflow.Context, input *health.DescribeAffectedEntitiesForOrganizationInput) (*health.DescribeAffectedEntitiesForOrganizationOutput, error) {
	var output health.DescribeAffectedEntitiesForOrganizationOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DescribeAffectedEntitiesForOrganization, input).Get(ctx, &output)
	return &output, err
}

func (a *HealthStub) DescribeAffectedEntitiesForOrganizationAsync(ctx workflow.Context, input *health.DescribeAffectedEntitiesForOrganizationInput) *HealthDescribeAffectedEntitiesForOrganizationResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DescribeAffectedEntitiesForOrganization, input)
	return &HealthDescribeAffectedEntitiesForOrganizationResult{Result: future}
}

func (a *HealthStub) DescribeEntityAggregates(ctx workflow.Context, input *health.DescribeEntityAggregatesInput) (*health.DescribeEntityAggregatesOutput, error) {
	var output health.DescribeEntityAggregatesOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DescribeEntityAggregates, input).Get(ctx, &output)
	return &output, err
}

func (a *HealthStub) DescribeEntityAggregatesAsync(ctx workflow.Context, input *health.DescribeEntityAggregatesInput) *HealthDescribeEntityAggregatesResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DescribeEntityAggregates, input)
	return &HealthDescribeEntityAggregatesResult{Result: future}
}

func (a *HealthStub) DescribeEventAggregates(ctx workflow.Context, input *health.DescribeEventAggregatesInput) (*health.DescribeEventAggregatesOutput, error) {
	var output health.DescribeEventAggregatesOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DescribeEventAggregates, input).Get(ctx, &output)
	return &output, err
}

func (a *HealthStub) DescribeEventAggregatesAsync(ctx workflow.Context, input *health.DescribeEventAggregatesInput) *HealthDescribeEventAggregatesResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DescribeEventAggregates, input)
	return &HealthDescribeEventAggregatesResult{Result: future}
}

func (a *HealthStub) DescribeEventDetails(ctx workflow.Context, input *health.DescribeEventDetailsInput) (*health.DescribeEventDetailsOutput, error) {
	var output health.DescribeEventDetailsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DescribeEventDetails, input).Get(ctx, &output)
	return &output, err
}

func (a *HealthStub) DescribeEventDetailsAsync(ctx workflow.Context, input *health.DescribeEventDetailsInput) *HealthDescribeEventDetailsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DescribeEventDetails, input)
	return &HealthDescribeEventDetailsResult{Result: future}
}

func (a *HealthStub) DescribeEventDetailsForOrganization(ctx workflow.Context, input *health.DescribeEventDetailsForOrganizationInput) (*health.DescribeEventDetailsForOrganizationOutput, error) {
	var output health.DescribeEventDetailsForOrganizationOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DescribeEventDetailsForOrganization, input).Get(ctx, &output)
	return &output, err
}

func (a *HealthStub) DescribeEventDetailsForOrganizationAsync(ctx workflow.Context, input *health.DescribeEventDetailsForOrganizationInput) *HealthDescribeEventDetailsForOrganizationResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DescribeEventDetailsForOrganization, input)
	return &HealthDescribeEventDetailsForOrganizationResult{Result: future}
}

func (a *HealthStub) DescribeEventTypes(ctx workflow.Context, input *health.DescribeEventTypesInput) (*health.DescribeEventTypesOutput, error) {
	var output health.DescribeEventTypesOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DescribeEventTypes, input).Get(ctx, &output)
	return &output, err
}

func (a *HealthStub) DescribeEventTypesAsync(ctx workflow.Context, input *health.DescribeEventTypesInput) *HealthDescribeEventTypesResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DescribeEventTypes, input)
	return &HealthDescribeEventTypesResult{Result: future}
}

func (a *HealthStub) DescribeEvents(ctx workflow.Context, input *health.DescribeEventsInput) (*health.DescribeEventsOutput, error) {
	var output health.DescribeEventsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DescribeEvents, input).Get(ctx, &output)
	return &output, err
}

func (a *HealthStub) DescribeEventsAsync(ctx workflow.Context, input *health.DescribeEventsInput) *HealthDescribeEventsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DescribeEvents, input)
	return &HealthDescribeEventsResult{Result: future}
}

func (a *HealthStub) DescribeEventsForOrganization(ctx workflow.Context, input *health.DescribeEventsForOrganizationInput) (*health.DescribeEventsForOrganizationOutput, error) {
	var output health.DescribeEventsForOrganizationOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DescribeEventsForOrganization, input).Get(ctx, &output)
	return &output, err
}

func (a *HealthStub) DescribeEventsForOrganizationAsync(ctx workflow.Context, input *health.DescribeEventsForOrganizationInput) *HealthDescribeEventsForOrganizationResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DescribeEventsForOrganization, input)
	return &HealthDescribeEventsForOrganizationResult{Result: future}
}

func (a *HealthStub) DescribeHealthServiceStatusForOrganization(ctx workflow.Context, input *health.DescribeHealthServiceStatusForOrganizationInput) (*health.DescribeHealthServiceStatusForOrganizationOutput, error) {
	var output health.DescribeHealthServiceStatusForOrganizationOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DescribeHealthServiceStatusForOrganization, input).Get(ctx, &output)
	return &output, err
}

func (a *HealthStub) DescribeHealthServiceStatusForOrganizationAsync(ctx workflow.Context, input *health.DescribeHealthServiceStatusForOrganizationInput) *HealthDescribeHealthServiceStatusForOrganizationResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DescribeHealthServiceStatusForOrganization, input)
	return &HealthDescribeHealthServiceStatusForOrganizationResult{Result: future}
}

func (a *HealthStub) DisableHealthServiceAccessForOrganization(ctx workflow.Context, input *health.DisableHealthServiceAccessForOrganizationInput) (*health.DisableHealthServiceAccessForOrganizationOutput, error) {
	var output health.DisableHealthServiceAccessForOrganizationOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DisableHealthServiceAccessForOrganization, input).Get(ctx, &output)
	return &output, err
}

func (a *HealthStub) DisableHealthServiceAccessForOrganizationAsync(ctx workflow.Context, input *health.DisableHealthServiceAccessForOrganizationInput) *HealthDisableHealthServiceAccessForOrganizationResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DisableHealthServiceAccessForOrganization, input)
	return &HealthDisableHealthServiceAccessForOrganizationResult{Result: future}
}

func (a *HealthStub) EnableHealthServiceAccessForOrganization(ctx workflow.Context, input *health.EnableHealthServiceAccessForOrganizationInput) (*health.EnableHealthServiceAccessForOrganizationOutput, error) {
	var output health.EnableHealthServiceAccessForOrganizationOutput
	err := workflow.ExecuteActivity(ctx, a.activities.EnableHealthServiceAccessForOrganization, input).Get(ctx, &output)
	return &output, err
}

func (a *HealthStub) EnableHealthServiceAccessForOrganizationAsync(ctx workflow.Context, input *health.EnableHealthServiceAccessForOrganizationInput) *HealthEnableHealthServiceAccessForOrganizationResult {
	future := workflow.ExecuteActivity(ctx, a.activities.EnableHealthServiceAccessForOrganization, input)
	return &HealthEnableHealthServiceAccessForOrganizationResult{Result: future}
}
