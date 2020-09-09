package awsclients

import (
	"github.com/aws/aws-sdk-go/service/waf"
	"github.com/aws/aws-sdk-go/service/wafregional"
	"go.temporal.io/sdk/workflow"
	"temporal.io/aws-sdk/awsactivities"
)

type WAFRegionalClient interface {
	AssociateWebACL(ctx workflow.Context, input *wafregional.AssociateWebACLInput) (*wafregional.AssociateWebACLOutput, error)
	AssociateWebACLAsync(ctx workflow.Context, input *wafregional.AssociateWebACLInput) *WafregionalAssociateWebACLResult

	CreateByteMatchSet(ctx workflow.Context, input *waf.CreateByteMatchSetInput) (*waf.CreateByteMatchSetOutput, error)
	CreateByteMatchSetAsync(ctx workflow.Context, input *waf.CreateByteMatchSetInput) *WafCreateByteMatchSetResult

	CreateGeoMatchSet(ctx workflow.Context, input *waf.CreateGeoMatchSetInput) (*waf.CreateGeoMatchSetOutput, error)
	CreateGeoMatchSetAsync(ctx workflow.Context, input *waf.CreateGeoMatchSetInput) *WafCreateGeoMatchSetResult

	CreateIPSet(ctx workflow.Context, input *waf.CreateIPSetInput) (*waf.CreateIPSetOutput, error)
	CreateIPSetAsync(ctx workflow.Context, input *waf.CreateIPSetInput) *WafCreateIPSetResult

	CreateRateBasedRule(ctx workflow.Context, input *waf.CreateRateBasedRuleInput) (*waf.CreateRateBasedRuleOutput, error)
	CreateRateBasedRuleAsync(ctx workflow.Context, input *waf.CreateRateBasedRuleInput) *WafCreateRateBasedRuleResult

	CreateRegexMatchSet(ctx workflow.Context, input *waf.CreateRegexMatchSetInput) (*waf.CreateRegexMatchSetOutput, error)
	CreateRegexMatchSetAsync(ctx workflow.Context, input *waf.CreateRegexMatchSetInput) *WafCreateRegexMatchSetResult

	CreateRegexPatternSet(ctx workflow.Context, input *waf.CreateRegexPatternSetInput) (*waf.CreateRegexPatternSetOutput, error)
	CreateRegexPatternSetAsync(ctx workflow.Context, input *waf.CreateRegexPatternSetInput) *WafCreateRegexPatternSetResult

	CreateRule(ctx workflow.Context, input *waf.CreateRuleInput) (*waf.CreateRuleOutput, error)
	CreateRuleAsync(ctx workflow.Context, input *waf.CreateRuleInput) *WafCreateRuleResult

	CreateRuleGroup(ctx workflow.Context, input *waf.CreateRuleGroupInput) (*waf.CreateRuleGroupOutput, error)
	CreateRuleGroupAsync(ctx workflow.Context, input *waf.CreateRuleGroupInput) *WafCreateRuleGroupResult

	CreateSizeConstraintSet(ctx workflow.Context, input *waf.CreateSizeConstraintSetInput) (*waf.CreateSizeConstraintSetOutput, error)
	CreateSizeConstraintSetAsync(ctx workflow.Context, input *waf.CreateSizeConstraintSetInput) *WafCreateSizeConstraintSetResult

	CreateSqlInjectionMatchSet(ctx workflow.Context, input *waf.CreateSqlInjectionMatchSetInput) (*waf.CreateSqlInjectionMatchSetOutput, error)
	CreateSqlInjectionMatchSetAsync(ctx workflow.Context, input *waf.CreateSqlInjectionMatchSetInput) *WafCreateSqlInjectionMatchSetResult

	CreateWebACL(ctx workflow.Context, input *waf.CreateWebACLInput) (*waf.CreateWebACLOutput, error)
	CreateWebACLAsync(ctx workflow.Context, input *waf.CreateWebACLInput) *WafCreateWebACLResult

	CreateWebACLMigrationStack(ctx workflow.Context, input *waf.CreateWebACLMigrationStackInput) (*waf.CreateWebACLMigrationStackOutput, error)
	CreateWebACLMigrationStackAsync(ctx workflow.Context, input *waf.CreateWebACLMigrationStackInput) *WafCreateWebACLMigrationStackResult

	CreateXssMatchSet(ctx workflow.Context, input *waf.CreateXssMatchSetInput) (*waf.CreateXssMatchSetOutput, error)
	CreateXssMatchSetAsync(ctx workflow.Context, input *waf.CreateXssMatchSetInput) *WafCreateXssMatchSetResult

	DeleteByteMatchSet(ctx workflow.Context, input *waf.DeleteByteMatchSetInput) (*waf.DeleteByteMatchSetOutput, error)
	DeleteByteMatchSetAsync(ctx workflow.Context, input *waf.DeleteByteMatchSetInput) *WafDeleteByteMatchSetResult

	DeleteGeoMatchSet(ctx workflow.Context, input *waf.DeleteGeoMatchSetInput) (*waf.DeleteGeoMatchSetOutput, error)
	DeleteGeoMatchSetAsync(ctx workflow.Context, input *waf.DeleteGeoMatchSetInput) *WafDeleteGeoMatchSetResult

	DeleteIPSet(ctx workflow.Context, input *waf.DeleteIPSetInput) (*waf.DeleteIPSetOutput, error)
	DeleteIPSetAsync(ctx workflow.Context, input *waf.DeleteIPSetInput) *WafDeleteIPSetResult

	DeleteLoggingConfiguration(ctx workflow.Context, input *waf.DeleteLoggingConfigurationInput) (*waf.DeleteLoggingConfigurationOutput, error)
	DeleteLoggingConfigurationAsync(ctx workflow.Context, input *waf.DeleteLoggingConfigurationInput) *WafDeleteLoggingConfigurationResult

	DeletePermissionPolicy(ctx workflow.Context, input *waf.DeletePermissionPolicyInput) (*waf.DeletePermissionPolicyOutput, error)
	DeletePermissionPolicyAsync(ctx workflow.Context, input *waf.DeletePermissionPolicyInput) *WafDeletePermissionPolicyResult

	DeleteRateBasedRule(ctx workflow.Context, input *waf.DeleteRateBasedRuleInput) (*waf.DeleteRateBasedRuleOutput, error)
	DeleteRateBasedRuleAsync(ctx workflow.Context, input *waf.DeleteRateBasedRuleInput) *WafDeleteRateBasedRuleResult

	DeleteRegexMatchSet(ctx workflow.Context, input *waf.DeleteRegexMatchSetInput) (*waf.DeleteRegexMatchSetOutput, error)
	DeleteRegexMatchSetAsync(ctx workflow.Context, input *waf.DeleteRegexMatchSetInput) *WafDeleteRegexMatchSetResult

	DeleteRegexPatternSet(ctx workflow.Context, input *waf.DeleteRegexPatternSetInput) (*waf.DeleteRegexPatternSetOutput, error)
	DeleteRegexPatternSetAsync(ctx workflow.Context, input *waf.DeleteRegexPatternSetInput) *WafDeleteRegexPatternSetResult

	DeleteRule(ctx workflow.Context, input *waf.DeleteRuleInput) (*waf.DeleteRuleOutput, error)
	DeleteRuleAsync(ctx workflow.Context, input *waf.DeleteRuleInput) *WafDeleteRuleResult

	DeleteRuleGroup(ctx workflow.Context, input *waf.DeleteRuleGroupInput) (*waf.DeleteRuleGroupOutput, error)
	DeleteRuleGroupAsync(ctx workflow.Context, input *waf.DeleteRuleGroupInput) *WafDeleteRuleGroupResult

	DeleteSizeConstraintSet(ctx workflow.Context, input *waf.DeleteSizeConstraintSetInput) (*waf.DeleteSizeConstraintSetOutput, error)
	DeleteSizeConstraintSetAsync(ctx workflow.Context, input *waf.DeleteSizeConstraintSetInput) *WafDeleteSizeConstraintSetResult

	DeleteSqlInjectionMatchSet(ctx workflow.Context, input *waf.DeleteSqlInjectionMatchSetInput) (*waf.DeleteSqlInjectionMatchSetOutput, error)
	DeleteSqlInjectionMatchSetAsync(ctx workflow.Context, input *waf.DeleteSqlInjectionMatchSetInput) *WafDeleteSqlInjectionMatchSetResult

	DeleteWebACL(ctx workflow.Context, input *waf.DeleteWebACLInput) (*waf.DeleteWebACLOutput, error)
	DeleteWebACLAsync(ctx workflow.Context, input *waf.DeleteWebACLInput) *WafDeleteWebACLResult

	DeleteXssMatchSet(ctx workflow.Context, input *waf.DeleteXssMatchSetInput) (*waf.DeleteXssMatchSetOutput, error)
	DeleteXssMatchSetAsync(ctx workflow.Context, input *waf.DeleteXssMatchSetInput) *WafDeleteXssMatchSetResult

	DisassociateWebACL(ctx workflow.Context, input *wafregional.DisassociateWebACLInput) (*wafregional.DisassociateWebACLOutput, error)
	DisassociateWebACLAsync(ctx workflow.Context, input *wafregional.DisassociateWebACLInput) *WafregionalDisassociateWebACLResult

	GetByteMatchSet(ctx workflow.Context, input *waf.GetByteMatchSetInput) (*waf.GetByteMatchSetOutput, error)
	GetByteMatchSetAsync(ctx workflow.Context, input *waf.GetByteMatchSetInput) *WafGetByteMatchSetResult

	GetChangeToken(ctx workflow.Context, input *waf.GetChangeTokenInput) (*waf.GetChangeTokenOutput, error)
	GetChangeTokenAsync(ctx workflow.Context, input *waf.GetChangeTokenInput) *WafGetChangeTokenResult

	GetChangeTokenStatus(ctx workflow.Context, input *waf.GetChangeTokenStatusInput) (*waf.GetChangeTokenStatusOutput, error)
	GetChangeTokenStatusAsync(ctx workflow.Context, input *waf.GetChangeTokenStatusInput) *WafGetChangeTokenStatusResult

	GetGeoMatchSet(ctx workflow.Context, input *waf.GetGeoMatchSetInput) (*waf.GetGeoMatchSetOutput, error)
	GetGeoMatchSetAsync(ctx workflow.Context, input *waf.GetGeoMatchSetInput) *WafGetGeoMatchSetResult

	GetIPSet(ctx workflow.Context, input *waf.GetIPSetInput) (*waf.GetIPSetOutput, error)
	GetIPSetAsync(ctx workflow.Context, input *waf.GetIPSetInput) *WafGetIPSetResult

	GetLoggingConfiguration(ctx workflow.Context, input *waf.GetLoggingConfigurationInput) (*waf.GetLoggingConfigurationOutput, error)
	GetLoggingConfigurationAsync(ctx workflow.Context, input *waf.GetLoggingConfigurationInput) *WafGetLoggingConfigurationResult

	GetPermissionPolicy(ctx workflow.Context, input *waf.GetPermissionPolicyInput) (*waf.GetPermissionPolicyOutput, error)
	GetPermissionPolicyAsync(ctx workflow.Context, input *waf.GetPermissionPolicyInput) *WafGetPermissionPolicyResult

	GetRateBasedRule(ctx workflow.Context, input *waf.GetRateBasedRuleInput) (*waf.GetRateBasedRuleOutput, error)
	GetRateBasedRuleAsync(ctx workflow.Context, input *waf.GetRateBasedRuleInput) *WafGetRateBasedRuleResult

	GetRateBasedRuleManagedKeys(ctx workflow.Context, input *waf.GetRateBasedRuleManagedKeysInput) (*waf.GetRateBasedRuleManagedKeysOutput, error)
	GetRateBasedRuleManagedKeysAsync(ctx workflow.Context, input *waf.GetRateBasedRuleManagedKeysInput) *WafGetRateBasedRuleManagedKeysResult

	GetRegexMatchSet(ctx workflow.Context, input *waf.GetRegexMatchSetInput) (*waf.GetRegexMatchSetOutput, error)
	GetRegexMatchSetAsync(ctx workflow.Context, input *waf.GetRegexMatchSetInput) *WafGetRegexMatchSetResult

	GetRegexPatternSet(ctx workflow.Context, input *waf.GetRegexPatternSetInput) (*waf.GetRegexPatternSetOutput, error)
	GetRegexPatternSetAsync(ctx workflow.Context, input *waf.GetRegexPatternSetInput) *WafGetRegexPatternSetResult

	GetRule(ctx workflow.Context, input *waf.GetRuleInput) (*waf.GetRuleOutput, error)
	GetRuleAsync(ctx workflow.Context, input *waf.GetRuleInput) *WafGetRuleResult

	GetRuleGroup(ctx workflow.Context, input *waf.GetRuleGroupInput) (*waf.GetRuleGroupOutput, error)
	GetRuleGroupAsync(ctx workflow.Context, input *waf.GetRuleGroupInput) *WafGetRuleGroupResult

	GetSampledRequests(ctx workflow.Context, input *waf.GetSampledRequestsInput) (*waf.GetSampledRequestsOutput, error)
	GetSampledRequestsAsync(ctx workflow.Context, input *waf.GetSampledRequestsInput) *WafGetSampledRequestsResult

	GetSizeConstraintSet(ctx workflow.Context, input *waf.GetSizeConstraintSetInput) (*waf.GetSizeConstraintSetOutput, error)
	GetSizeConstraintSetAsync(ctx workflow.Context, input *waf.GetSizeConstraintSetInput) *WafGetSizeConstraintSetResult

	GetSqlInjectionMatchSet(ctx workflow.Context, input *waf.GetSqlInjectionMatchSetInput) (*waf.GetSqlInjectionMatchSetOutput, error)
	GetSqlInjectionMatchSetAsync(ctx workflow.Context, input *waf.GetSqlInjectionMatchSetInput) *WafGetSqlInjectionMatchSetResult

	GetWebACL(ctx workflow.Context, input *waf.GetWebACLInput) (*waf.GetWebACLOutput, error)
	GetWebACLAsync(ctx workflow.Context, input *waf.GetWebACLInput) *WafGetWebACLResult

	GetWebACLForResource(ctx workflow.Context, input *wafregional.GetWebACLForResourceInput) (*wafregional.GetWebACLForResourceOutput, error)
	GetWebACLForResourceAsync(ctx workflow.Context, input *wafregional.GetWebACLForResourceInput) *WafregionalGetWebACLForResourceResult

	GetXssMatchSet(ctx workflow.Context, input *waf.GetXssMatchSetInput) (*waf.GetXssMatchSetOutput, error)
	GetXssMatchSetAsync(ctx workflow.Context, input *waf.GetXssMatchSetInput) *WafGetXssMatchSetResult

	ListActivatedRulesInRuleGroup(ctx workflow.Context, input *waf.ListActivatedRulesInRuleGroupInput) (*waf.ListActivatedRulesInRuleGroupOutput, error)
	ListActivatedRulesInRuleGroupAsync(ctx workflow.Context, input *waf.ListActivatedRulesInRuleGroupInput) *WafListActivatedRulesInRuleGroupResult

	ListByteMatchSets(ctx workflow.Context, input *waf.ListByteMatchSetsInput) (*waf.ListByteMatchSetsOutput, error)
	ListByteMatchSetsAsync(ctx workflow.Context, input *waf.ListByteMatchSetsInput) *WafListByteMatchSetsResult

	ListGeoMatchSets(ctx workflow.Context, input *waf.ListGeoMatchSetsInput) (*waf.ListGeoMatchSetsOutput, error)
	ListGeoMatchSetsAsync(ctx workflow.Context, input *waf.ListGeoMatchSetsInput) *WafListGeoMatchSetsResult

	ListIPSets(ctx workflow.Context, input *waf.ListIPSetsInput) (*waf.ListIPSetsOutput, error)
	ListIPSetsAsync(ctx workflow.Context, input *waf.ListIPSetsInput) *WafListIPSetsResult

	ListLoggingConfigurations(ctx workflow.Context, input *waf.ListLoggingConfigurationsInput) (*waf.ListLoggingConfigurationsOutput, error)
	ListLoggingConfigurationsAsync(ctx workflow.Context, input *waf.ListLoggingConfigurationsInput) *WafListLoggingConfigurationsResult

	ListRateBasedRules(ctx workflow.Context, input *waf.ListRateBasedRulesInput) (*waf.ListRateBasedRulesOutput, error)
	ListRateBasedRulesAsync(ctx workflow.Context, input *waf.ListRateBasedRulesInput) *WafListRateBasedRulesResult

	ListRegexMatchSets(ctx workflow.Context, input *waf.ListRegexMatchSetsInput) (*waf.ListRegexMatchSetsOutput, error)
	ListRegexMatchSetsAsync(ctx workflow.Context, input *waf.ListRegexMatchSetsInput) *WafListRegexMatchSetsResult

	ListRegexPatternSets(ctx workflow.Context, input *waf.ListRegexPatternSetsInput) (*waf.ListRegexPatternSetsOutput, error)
	ListRegexPatternSetsAsync(ctx workflow.Context, input *waf.ListRegexPatternSetsInput) *WafListRegexPatternSetsResult

	ListResourcesForWebACL(ctx workflow.Context, input *wafregional.ListResourcesForWebACLInput) (*wafregional.ListResourcesForWebACLOutput, error)
	ListResourcesForWebACLAsync(ctx workflow.Context, input *wafregional.ListResourcesForWebACLInput) *WafregionalListResourcesForWebACLResult

	ListRuleGroups(ctx workflow.Context, input *waf.ListRuleGroupsInput) (*waf.ListRuleGroupsOutput, error)
	ListRuleGroupsAsync(ctx workflow.Context, input *waf.ListRuleGroupsInput) *WafListRuleGroupsResult

	ListRules(ctx workflow.Context, input *waf.ListRulesInput) (*waf.ListRulesOutput, error)
	ListRulesAsync(ctx workflow.Context, input *waf.ListRulesInput) *WafListRulesResult

	ListSizeConstraintSets(ctx workflow.Context, input *waf.ListSizeConstraintSetsInput) (*waf.ListSizeConstraintSetsOutput, error)
	ListSizeConstraintSetsAsync(ctx workflow.Context, input *waf.ListSizeConstraintSetsInput) *WafListSizeConstraintSetsResult

	ListSqlInjectionMatchSets(ctx workflow.Context, input *waf.ListSqlInjectionMatchSetsInput) (*waf.ListSqlInjectionMatchSetsOutput, error)
	ListSqlInjectionMatchSetsAsync(ctx workflow.Context, input *waf.ListSqlInjectionMatchSetsInput) *WafListSqlInjectionMatchSetsResult

	ListSubscribedRuleGroups(ctx workflow.Context, input *waf.ListSubscribedRuleGroupsInput) (*waf.ListSubscribedRuleGroupsOutput, error)
	ListSubscribedRuleGroupsAsync(ctx workflow.Context, input *waf.ListSubscribedRuleGroupsInput) *WafListSubscribedRuleGroupsResult

	ListTagsForResource(ctx workflow.Context, input *waf.ListTagsForResourceInput) (*waf.ListTagsForResourceOutput, error)
	ListTagsForResourceAsync(ctx workflow.Context, input *waf.ListTagsForResourceInput) *WafListTagsForResourceResult

	ListWebACLs(ctx workflow.Context, input *waf.ListWebACLsInput) (*waf.ListWebACLsOutput, error)
	ListWebACLsAsync(ctx workflow.Context, input *waf.ListWebACLsInput) *WafListWebACLsResult

	ListXssMatchSets(ctx workflow.Context, input *waf.ListXssMatchSetsInput) (*waf.ListXssMatchSetsOutput, error)
	ListXssMatchSetsAsync(ctx workflow.Context, input *waf.ListXssMatchSetsInput) *WafListXssMatchSetsResult

	PutLoggingConfiguration(ctx workflow.Context, input *waf.PutLoggingConfigurationInput) (*waf.PutLoggingConfigurationOutput, error)
	PutLoggingConfigurationAsync(ctx workflow.Context, input *waf.PutLoggingConfigurationInput) *WafPutLoggingConfigurationResult

	PutPermissionPolicy(ctx workflow.Context, input *waf.PutPermissionPolicyInput) (*waf.PutPermissionPolicyOutput, error)
	PutPermissionPolicyAsync(ctx workflow.Context, input *waf.PutPermissionPolicyInput) *WafPutPermissionPolicyResult

	TagResource(ctx workflow.Context, input *waf.TagResourceInput) (*waf.TagResourceOutput, error)
	TagResourceAsync(ctx workflow.Context, input *waf.TagResourceInput) *WafTagResourceResult

	UntagResource(ctx workflow.Context, input *waf.UntagResourceInput) (*waf.UntagResourceOutput, error)
	UntagResourceAsync(ctx workflow.Context, input *waf.UntagResourceInput) *WafUntagResourceResult

	UpdateByteMatchSet(ctx workflow.Context, input *waf.UpdateByteMatchSetInput) (*waf.UpdateByteMatchSetOutput, error)
	UpdateByteMatchSetAsync(ctx workflow.Context, input *waf.UpdateByteMatchSetInput) *WafUpdateByteMatchSetResult

	UpdateGeoMatchSet(ctx workflow.Context, input *waf.UpdateGeoMatchSetInput) (*waf.UpdateGeoMatchSetOutput, error)
	UpdateGeoMatchSetAsync(ctx workflow.Context, input *waf.UpdateGeoMatchSetInput) *WafUpdateGeoMatchSetResult

	UpdateIPSet(ctx workflow.Context, input *waf.UpdateIPSetInput) (*waf.UpdateIPSetOutput, error)
	UpdateIPSetAsync(ctx workflow.Context, input *waf.UpdateIPSetInput) *WafUpdateIPSetResult

	UpdateRateBasedRule(ctx workflow.Context, input *waf.UpdateRateBasedRuleInput) (*waf.UpdateRateBasedRuleOutput, error)
	UpdateRateBasedRuleAsync(ctx workflow.Context, input *waf.UpdateRateBasedRuleInput) *WafUpdateRateBasedRuleResult

	UpdateRegexMatchSet(ctx workflow.Context, input *waf.UpdateRegexMatchSetInput) (*waf.UpdateRegexMatchSetOutput, error)
	UpdateRegexMatchSetAsync(ctx workflow.Context, input *waf.UpdateRegexMatchSetInput) *WafUpdateRegexMatchSetResult

	UpdateRegexPatternSet(ctx workflow.Context, input *waf.UpdateRegexPatternSetInput) (*waf.UpdateRegexPatternSetOutput, error)
	UpdateRegexPatternSetAsync(ctx workflow.Context, input *waf.UpdateRegexPatternSetInput) *WafUpdateRegexPatternSetResult

	UpdateRule(ctx workflow.Context, input *waf.UpdateRuleInput) (*waf.UpdateRuleOutput, error)
	UpdateRuleAsync(ctx workflow.Context, input *waf.UpdateRuleInput) *WafUpdateRuleResult

	UpdateRuleGroup(ctx workflow.Context, input *waf.UpdateRuleGroupInput) (*waf.UpdateRuleGroupOutput, error)
	UpdateRuleGroupAsync(ctx workflow.Context, input *waf.UpdateRuleGroupInput) *WafUpdateRuleGroupResult

	UpdateSizeConstraintSet(ctx workflow.Context, input *waf.UpdateSizeConstraintSetInput) (*waf.UpdateSizeConstraintSetOutput, error)
	UpdateSizeConstraintSetAsync(ctx workflow.Context, input *waf.UpdateSizeConstraintSetInput) *WafUpdateSizeConstraintSetResult

	UpdateSqlInjectionMatchSet(ctx workflow.Context, input *waf.UpdateSqlInjectionMatchSetInput) (*waf.UpdateSqlInjectionMatchSetOutput, error)
	UpdateSqlInjectionMatchSetAsync(ctx workflow.Context, input *waf.UpdateSqlInjectionMatchSetInput) *WafUpdateSqlInjectionMatchSetResult

	UpdateWebACL(ctx workflow.Context, input *waf.UpdateWebACLInput) (*waf.UpdateWebACLOutput, error)
	UpdateWebACLAsync(ctx workflow.Context, input *waf.UpdateWebACLInput) *WafUpdateWebACLResult

	UpdateXssMatchSet(ctx workflow.Context, input *waf.UpdateXssMatchSetInput) (*waf.UpdateXssMatchSetOutput, error)
	UpdateXssMatchSetAsync(ctx workflow.Context, input *waf.UpdateXssMatchSetInput) *WafUpdateXssMatchSetResult
}

type WafregionalAssociateWebACLResult struct {
	Result workflow.Future
}

func (r *WafregionalAssociateWebACLResult) Get(ctx workflow.Context) (*wafregional.AssociateWebACLOutput, error) {
	var output wafregional.AssociateWebACLOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type WafregionalDisassociateWebACLResult struct {
	Result workflow.Future
}

func (r *WafregionalDisassociateWebACLResult) Get(ctx workflow.Context) (*wafregional.DisassociateWebACLOutput, error) {
	var output wafregional.DisassociateWebACLOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type WafregionalGetWebACLForResourceResult struct {
	Result workflow.Future
}

func (r *WafregionalGetWebACLForResourceResult) Get(ctx workflow.Context) (*wafregional.GetWebACLForResourceOutput, error) {
	var output wafregional.GetWebACLForResourceOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type WafregionalListResourcesForWebACLResult struct {
	Result workflow.Future
}

func (r *WafregionalListResourcesForWebACLResult) Get(ctx workflow.Context) (*wafregional.ListResourcesForWebACLOutput, error) {
	var output wafregional.ListResourcesForWebACLOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type WAFRegionalStub struct {
	activities awsactivities.WAFRegionalActivities
}

func NewWAFRegionalStub() WAFRegionalClient {
	return &WAFRegionalStub{}
}

func (a *WAFRegionalStub) AssociateWebACL(ctx workflow.Context, input *wafregional.AssociateWebACLInput) (*wafregional.AssociateWebACLOutput, error) {
	var output wafregional.AssociateWebACLOutput
	err := workflow.ExecuteActivity(ctx, a.activities.AssociateWebACL, input).Get(ctx, &output)
	return &output, err
}

func (a *WAFRegionalStub) AssociateWebACLAsync(ctx workflow.Context, input *wafregional.AssociateWebACLInput) *WafregionalAssociateWebACLResult {
	future := workflow.ExecuteActivity(ctx, a.activities.AssociateWebACL, input)
	return &WafregionalAssociateWebACLResult{Result: future}
}

func (a *WAFRegionalStub) CreateByteMatchSet(ctx workflow.Context, input *waf.CreateByteMatchSetInput) (*waf.CreateByteMatchSetOutput, error) {
	var output waf.CreateByteMatchSetOutput
	err := workflow.ExecuteActivity(ctx, a.activities.CreateByteMatchSet, input).Get(ctx, &output)
	return &output, err
}

func (a *WAFRegionalStub) CreateByteMatchSetAsync(ctx workflow.Context, input *waf.CreateByteMatchSetInput) *WafCreateByteMatchSetResult {
	future := workflow.ExecuteActivity(ctx, a.activities.CreateByteMatchSet, input)
	return &WafCreateByteMatchSetResult{Result: future}
}

func (a *WAFRegionalStub) CreateGeoMatchSet(ctx workflow.Context, input *waf.CreateGeoMatchSetInput) (*waf.CreateGeoMatchSetOutput, error) {
	var output waf.CreateGeoMatchSetOutput
	err := workflow.ExecuteActivity(ctx, a.activities.CreateGeoMatchSet, input).Get(ctx, &output)
	return &output, err
}

func (a *WAFRegionalStub) CreateGeoMatchSetAsync(ctx workflow.Context, input *waf.CreateGeoMatchSetInput) *WafCreateGeoMatchSetResult {
	future := workflow.ExecuteActivity(ctx, a.activities.CreateGeoMatchSet, input)
	return &WafCreateGeoMatchSetResult{Result: future}
}

func (a *WAFRegionalStub) CreateIPSet(ctx workflow.Context, input *waf.CreateIPSetInput) (*waf.CreateIPSetOutput, error) {
	var output waf.CreateIPSetOutput
	err := workflow.ExecuteActivity(ctx, a.activities.CreateIPSet, input).Get(ctx, &output)
	return &output, err
}

func (a *WAFRegionalStub) CreateIPSetAsync(ctx workflow.Context, input *waf.CreateIPSetInput) *WafCreateIPSetResult {
	future := workflow.ExecuteActivity(ctx, a.activities.CreateIPSet, input)
	return &WafCreateIPSetResult{Result: future}
}

func (a *WAFRegionalStub) CreateRateBasedRule(ctx workflow.Context, input *waf.CreateRateBasedRuleInput) (*waf.CreateRateBasedRuleOutput, error) {
	var output waf.CreateRateBasedRuleOutput
	err := workflow.ExecuteActivity(ctx, a.activities.CreateRateBasedRule, input).Get(ctx, &output)
	return &output, err
}

func (a *WAFRegionalStub) CreateRateBasedRuleAsync(ctx workflow.Context, input *waf.CreateRateBasedRuleInput) *WafCreateRateBasedRuleResult {
	future := workflow.ExecuteActivity(ctx, a.activities.CreateRateBasedRule, input)
	return &WafCreateRateBasedRuleResult{Result: future}
}

func (a *WAFRegionalStub) CreateRegexMatchSet(ctx workflow.Context, input *waf.CreateRegexMatchSetInput) (*waf.CreateRegexMatchSetOutput, error) {
	var output waf.CreateRegexMatchSetOutput
	err := workflow.ExecuteActivity(ctx, a.activities.CreateRegexMatchSet, input).Get(ctx, &output)
	return &output, err
}

func (a *WAFRegionalStub) CreateRegexMatchSetAsync(ctx workflow.Context, input *waf.CreateRegexMatchSetInput) *WafCreateRegexMatchSetResult {
	future := workflow.ExecuteActivity(ctx, a.activities.CreateRegexMatchSet, input)
	return &WafCreateRegexMatchSetResult{Result: future}
}

func (a *WAFRegionalStub) CreateRegexPatternSet(ctx workflow.Context, input *waf.CreateRegexPatternSetInput) (*waf.CreateRegexPatternSetOutput, error) {
	var output waf.CreateRegexPatternSetOutput
	err := workflow.ExecuteActivity(ctx, a.activities.CreateRegexPatternSet, input).Get(ctx, &output)
	return &output, err
}

func (a *WAFRegionalStub) CreateRegexPatternSetAsync(ctx workflow.Context, input *waf.CreateRegexPatternSetInput) *WafCreateRegexPatternSetResult {
	future := workflow.ExecuteActivity(ctx, a.activities.CreateRegexPatternSet, input)
	return &WafCreateRegexPatternSetResult{Result: future}
}

func (a *WAFRegionalStub) CreateRule(ctx workflow.Context, input *waf.CreateRuleInput) (*waf.CreateRuleOutput, error) {
	var output waf.CreateRuleOutput
	err := workflow.ExecuteActivity(ctx, a.activities.CreateRule, input).Get(ctx, &output)
	return &output, err
}

func (a *WAFRegionalStub) CreateRuleAsync(ctx workflow.Context, input *waf.CreateRuleInput) *WafCreateRuleResult {
	future := workflow.ExecuteActivity(ctx, a.activities.CreateRule, input)
	return &WafCreateRuleResult{Result: future}
}

func (a *WAFRegionalStub) CreateRuleGroup(ctx workflow.Context, input *waf.CreateRuleGroupInput) (*waf.CreateRuleGroupOutput, error) {
	var output waf.CreateRuleGroupOutput
	err := workflow.ExecuteActivity(ctx, a.activities.CreateRuleGroup, input).Get(ctx, &output)
	return &output, err
}

func (a *WAFRegionalStub) CreateRuleGroupAsync(ctx workflow.Context, input *waf.CreateRuleGroupInput) *WafCreateRuleGroupResult {
	future := workflow.ExecuteActivity(ctx, a.activities.CreateRuleGroup, input)
	return &WafCreateRuleGroupResult{Result: future}
}

func (a *WAFRegionalStub) CreateSizeConstraintSet(ctx workflow.Context, input *waf.CreateSizeConstraintSetInput) (*waf.CreateSizeConstraintSetOutput, error) {
	var output waf.CreateSizeConstraintSetOutput
	err := workflow.ExecuteActivity(ctx, a.activities.CreateSizeConstraintSet, input).Get(ctx, &output)
	return &output, err
}

func (a *WAFRegionalStub) CreateSizeConstraintSetAsync(ctx workflow.Context, input *waf.CreateSizeConstraintSetInput) *WafCreateSizeConstraintSetResult {
	future := workflow.ExecuteActivity(ctx, a.activities.CreateSizeConstraintSet, input)
	return &WafCreateSizeConstraintSetResult{Result: future}
}

func (a *WAFRegionalStub) CreateSqlInjectionMatchSet(ctx workflow.Context, input *waf.CreateSqlInjectionMatchSetInput) (*waf.CreateSqlInjectionMatchSetOutput, error) {
	var output waf.CreateSqlInjectionMatchSetOutput
	err := workflow.ExecuteActivity(ctx, a.activities.CreateSqlInjectionMatchSet, input).Get(ctx, &output)
	return &output, err
}

func (a *WAFRegionalStub) CreateSqlInjectionMatchSetAsync(ctx workflow.Context, input *waf.CreateSqlInjectionMatchSetInput) *WafCreateSqlInjectionMatchSetResult {
	future := workflow.ExecuteActivity(ctx, a.activities.CreateSqlInjectionMatchSet, input)
	return &WafCreateSqlInjectionMatchSetResult{Result: future}
}

func (a *WAFRegionalStub) CreateWebACL(ctx workflow.Context, input *waf.CreateWebACLInput) (*waf.CreateWebACLOutput, error) {
	var output waf.CreateWebACLOutput
	err := workflow.ExecuteActivity(ctx, a.activities.CreateWebACL, input).Get(ctx, &output)
	return &output, err
}

func (a *WAFRegionalStub) CreateWebACLAsync(ctx workflow.Context, input *waf.CreateWebACLInput) *WafCreateWebACLResult {
	future := workflow.ExecuteActivity(ctx, a.activities.CreateWebACL, input)
	return &WafCreateWebACLResult{Result: future}
}

func (a *WAFRegionalStub) CreateWebACLMigrationStack(ctx workflow.Context, input *waf.CreateWebACLMigrationStackInput) (*waf.CreateWebACLMigrationStackOutput, error) {
	var output waf.CreateWebACLMigrationStackOutput
	err := workflow.ExecuteActivity(ctx, a.activities.CreateWebACLMigrationStack, input).Get(ctx, &output)
	return &output, err
}

func (a *WAFRegionalStub) CreateWebACLMigrationStackAsync(ctx workflow.Context, input *waf.CreateWebACLMigrationStackInput) *WafCreateWebACLMigrationStackResult {
	future := workflow.ExecuteActivity(ctx, a.activities.CreateWebACLMigrationStack, input)
	return &WafCreateWebACLMigrationStackResult{Result: future}
}

func (a *WAFRegionalStub) CreateXssMatchSet(ctx workflow.Context, input *waf.CreateXssMatchSetInput) (*waf.CreateXssMatchSetOutput, error) {
	var output waf.CreateXssMatchSetOutput
	err := workflow.ExecuteActivity(ctx, a.activities.CreateXssMatchSet, input).Get(ctx, &output)
	return &output, err
}

func (a *WAFRegionalStub) CreateXssMatchSetAsync(ctx workflow.Context, input *waf.CreateXssMatchSetInput) *WafCreateXssMatchSetResult {
	future := workflow.ExecuteActivity(ctx, a.activities.CreateXssMatchSet, input)
	return &WafCreateXssMatchSetResult{Result: future}
}

func (a *WAFRegionalStub) DeleteByteMatchSet(ctx workflow.Context, input *waf.DeleteByteMatchSetInput) (*waf.DeleteByteMatchSetOutput, error) {
	var output waf.DeleteByteMatchSetOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeleteByteMatchSet, input).Get(ctx, &output)
	return &output, err
}

func (a *WAFRegionalStub) DeleteByteMatchSetAsync(ctx workflow.Context, input *waf.DeleteByteMatchSetInput) *WafDeleteByteMatchSetResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeleteByteMatchSet, input)
	return &WafDeleteByteMatchSetResult{Result: future}
}

func (a *WAFRegionalStub) DeleteGeoMatchSet(ctx workflow.Context, input *waf.DeleteGeoMatchSetInput) (*waf.DeleteGeoMatchSetOutput, error) {
	var output waf.DeleteGeoMatchSetOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeleteGeoMatchSet, input).Get(ctx, &output)
	return &output, err
}

func (a *WAFRegionalStub) DeleteGeoMatchSetAsync(ctx workflow.Context, input *waf.DeleteGeoMatchSetInput) *WafDeleteGeoMatchSetResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeleteGeoMatchSet, input)
	return &WafDeleteGeoMatchSetResult{Result: future}
}

func (a *WAFRegionalStub) DeleteIPSet(ctx workflow.Context, input *waf.DeleteIPSetInput) (*waf.DeleteIPSetOutput, error) {
	var output waf.DeleteIPSetOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeleteIPSet, input).Get(ctx, &output)
	return &output, err
}

func (a *WAFRegionalStub) DeleteIPSetAsync(ctx workflow.Context, input *waf.DeleteIPSetInput) *WafDeleteIPSetResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeleteIPSet, input)
	return &WafDeleteIPSetResult{Result: future}
}

func (a *WAFRegionalStub) DeleteLoggingConfiguration(ctx workflow.Context, input *waf.DeleteLoggingConfigurationInput) (*waf.DeleteLoggingConfigurationOutput, error) {
	var output waf.DeleteLoggingConfigurationOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeleteLoggingConfiguration, input).Get(ctx, &output)
	return &output, err
}

func (a *WAFRegionalStub) DeleteLoggingConfigurationAsync(ctx workflow.Context, input *waf.DeleteLoggingConfigurationInput) *WafDeleteLoggingConfigurationResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeleteLoggingConfiguration, input)
	return &WafDeleteLoggingConfigurationResult{Result: future}
}

func (a *WAFRegionalStub) DeletePermissionPolicy(ctx workflow.Context, input *waf.DeletePermissionPolicyInput) (*waf.DeletePermissionPolicyOutput, error) {
	var output waf.DeletePermissionPolicyOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeletePermissionPolicy, input).Get(ctx, &output)
	return &output, err
}

func (a *WAFRegionalStub) DeletePermissionPolicyAsync(ctx workflow.Context, input *waf.DeletePermissionPolicyInput) *WafDeletePermissionPolicyResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeletePermissionPolicy, input)
	return &WafDeletePermissionPolicyResult{Result: future}
}

func (a *WAFRegionalStub) DeleteRateBasedRule(ctx workflow.Context, input *waf.DeleteRateBasedRuleInput) (*waf.DeleteRateBasedRuleOutput, error) {
	var output waf.DeleteRateBasedRuleOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeleteRateBasedRule, input).Get(ctx, &output)
	return &output, err
}

func (a *WAFRegionalStub) DeleteRateBasedRuleAsync(ctx workflow.Context, input *waf.DeleteRateBasedRuleInput) *WafDeleteRateBasedRuleResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeleteRateBasedRule, input)
	return &WafDeleteRateBasedRuleResult{Result: future}
}

func (a *WAFRegionalStub) DeleteRegexMatchSet(ctx workflow.Context, input *waf.DeleteRegexMatchSetInput) (*waf.DeleteRegexMatchSetOutput, error) {
	var output waf.DeleteRegexMatchSetOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeleteRegexMatchSet, input).Get(ctx, &output)
	return &output, err
}

func (a *WAFRegionalStub) DeleteRegexMatchSetAsync(ctx workflow.Context, input *waf.DeleteRegexMatchSetInput) *WafDeleteRegexMatchSetResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeleteRegexMatchSet, input)
	return &WafDeleteRegexMatchSetResult{Result: future}
}

func (a *WAFRegionalStub) DeleteRegexPatternSet(ctx workflow.Context, input *waf.DeleteRegexPatternSetInput) (*waf.DeleteRegexPatternSetOutput, error) {
	var output waf.DeleteRegexPatternSetOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeleteRegexPatternSet, input).Get(ctx, &output)
	return &output, err
}

func (a *WAFRegionalStub) DeleteRegexPatternSetAsync(ctx workflow.Context, input *waf.DeleteRegexPatternSetInput) *WafDeleteRegexPatternSetResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeleteRegexPatternSet, input)
	return &WafDeleteRegexPatternSetResult{Result: future}
}

func (a *WAFRegionalStub) DeleteRule(ctx workflow.Context, input *waf.DeleteRuleInput) (*waf.DeleteRuleOutput, error) {
	var output waf.DeleteRuleOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeleteRule, input).Get(ctx, &output)
	return &output, err
}

func (a *WAFRegionalStub) DeleteRuleAsync(ctx workflow.Context, input *waf.DeleteRuleInput) *WafDeleteRuleResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeleteRule, input)
	return &WafDeleteRuleResult{Result: future}
}

func (a *WAFRegionalStub) DeleteRuleGroup(ctx workflow.Context, input *waf.DeleteRuleGroupInput) (*waf.DeleteRuleGroupOutput, error) {
	var output waf.DeleteRuleGroupOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeleteRuleGroup, input).Get(ctx, &output)
	return &output, err
}

func (a *WAFRegionalStub) DeleteRuleGroupAsync(ctx workflow.Context, input *waf.DeleteRuleGroupInput) *WafDeleteRuleGroupResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeleteRuleGroup, input)
	return &WafDeleteRuleGroupResult{Result: future}
}

func (a *WAFRegionalStub) DeleteSizeConstraintSet(ctx workflow.Context, input *waf.DeleteSizeConstraintSetInput) (*waf.DeleteSizeConstraintSetOutput, error) {
	var output waf.DeleteSizeConstraintSetOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeleteSizeConstraintSet, input).Get(ctx, &output)
	return &output, err
}

func (a *WAFRegionalStub) DeleteSizeConstraintSetAsync(ctx workflow.Context, input *waf.DeleteSizeConstraintSetInput) *WafDeleteSizeConstraintSetResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeleteSizeConstraintSet, input)
	return &WafDeleteSizeConstraintSetResult{Result: future}
}

func (a *WAFRegionalStub) DeleteSqlInjectionMatchSet(ctx workflow.Context, input *waf.DeleteSqlInjectionMatchSetInput) (*waf.DeleteSqlInjectionMatchSetOutput, error) {
	var output waf.DeleteSqlInjectionMatchSetOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeleteSqlInjectionMatchSet, input).Get(ctx, &output)
	return &output, err
}

func (a *WAFRegionalStub) DeleteSqlInjectionMatchSetAsync(ctx workflow.Context, input *waf.DeleteSqlInjectionMatchSetInput) *WafDeleteSqlInjectionMatchSetResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeleteSqlInjectionMatchSet, input)
	return &WafDeleteSqlInjectionMatchSetResult{Result: future}
}

func (a *WAFRegionalStub) DeleteWebACL(ctx workflow.Context, input *waf.DeleteWebACLInput) (*waf.DeleteWebACLOutput, error) {
	var output waf.DeleteWebACLOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeleteWebACL, input).Get(ctx, &output)
	return &output, err
}

func (a *WAFRegionalStub) DeleteWebACLAsync(ctx workflow.Context, input *waf.DeleteWebACLInput) *WafDeleteWebACLResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeleteWebACL, input)
	return &WafDeleteWebACLResult{Result: future}
}

func (a *WAFRegionalStub) DeleteXssMatchSet(ctx workflow.Context, input *waf.DeleteXssMatchSetInput) (*waf.DeleteXssMatchSetOutput, error) {
	var output waf.DeleteXssMatchSetOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeleteXssMatchSet, input).Get(ctx, &output)
	return &output, err
}

func (a *WAFRegionalStub) DeleteXssMatchSetAsync(ctx workflow.Context, input *waf.DeleteXssMatchSetInput) *WafDeleteXssMatchSetResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeleteXssMatchSet, input)
	return &WafDeleteXssMatchSetResult{Result: future}
}

func (a *WAFRegionalStub) DisassociateWebACL(ctx workflow.Context, input *wafregional.DisassociateWebACLInput) (*wafregional.DisassociateWebACLOutput, error) {
	var output wafregional.DisassociateWebACLOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DisassociateWebACL, input).Get(ctx, &output)
	return &output, err
}

func (a *WAFRegionalStub) DisassociateWebACLAsync(ctx workflow.Context, input *wafregional.DisassociateWebACLInput) *WafregionalDisassociateWebACLResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DisassociateWebACL, input)
	return &WafregionalDisassociateWebACLResult{Result: future}
}

func (a *WAFRegionalStub) GetByteMatchSet(ctx workflow.Context, input *waf.GetByteMatchSetInput) (*waf.GetByteMatchSetOutput, error) {
	var output waf.GetByteMatchSetOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetByteMatchSet, input).Get(ctx, &output)
	return &output, err
}

func (a *WAFRegionalStub) GetByteMatchSetAsync(ctx workflow.Context, input *waf.GetByteMatchSetInput) *WafGetByteMatchSetResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetByteMatchSet, input)
	return &WafGetByteMatchSetResult{Result: future}
}

func (a *WAFRegionalStub) GetChangeToken(ctx workflow.Context, input *waf.GetChangeTokenInput) (*waf.GetChangeTokenOutput, error) {
	var output waf.GetChangeTokenOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetChangeToken, input).Get(ctx, &output)
	return &output, err
}

func (a *WAFRegionalStub) GetChangeTokenAsync(ctx workflow.Context, input *waf.GetChangeTokenInput) *WafGetChangeTokenResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetChangeToken, input)
	return &WafGetChangeTokenResult{Result: future}
}

func (a *WAFRegionalStub) GetChangeTokenStatus(ctx workflow.Context, input *waf.GetChangeTokenStatusInput) (*waf.GetChangeTokenStatusOutput, error) {
	var output waf.GetChangeTokenStatusOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetChangeTokenStatus, input).Get(ctx, &output)
	return &output, err
}

func (a *WAFRegionalStub) GetChangeTokenStatusAsync(ctx workflow.Context, input *waf.GetChangeTokenStatusInput) *WafGetChangeTokenStatusResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetChangeTokenStatus, input)
	return &WafGetChangeTokenStatusResult{Result: future}
}

func (a *WAFRegionalStub) GetGeoMatchSet(ctx workflow.Context, input *waf.GetGeoMatchSetInput) (*waf.GetGeoMatchSetOutput, error) {
	var output waf.GetGeoMatchSetOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetGeoMatchSet, input).Get(ctx, &output)
	return &output, err
}

func (a *WAFRegionalStub) GetGeoMatchSetAsync(ctx workflow.Context, input *waf.GetGeoMatchSetInput) *WafGetGeoMatchSetResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetGeoMatchSet, input)
	return &WafGetGeoMatchSetResult{Result: future}
}

func (a *WAFRegionalStub) GetIPSet(ctx workflow.Context, input *waf.GetIPSetInput) (*waf.GetIPSetOutput, error) {
	var output waf.GetIPSetOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetIPSet, input).Get(ctx, &output)
	return &output, err
}

func (a *WAFRegionalStub) GetIPSetAsync(ctx workflow.Context, input *waf.GetIPSetInput) *WafGetIPSetResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetIPSet, input)
	return &WafGetIPSetResult{Result: future}
}

func (a *WAFRegionalStub) GetLoggingConfiguration(ctx workflow.Context, input *waf.GetLoggingConfigurationInput) (*waf.GetLoggingConfigurationOutput, error) {
	var output waf.GetLoggingConfigurationOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetLoggingConfiguration, input).Get(ctx, &output)
	return &output, err
}

func (a *WAFRegionalStub) GetLoggingConfigurationAsync(ctx workflow.Context, input *waf.GetLoggingConfigurationInput) *WafGetLoggingConfigurationResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetLoggingConfiguration, input)
	return &WafGetLoggingConfigurationResult{Result: future}
}

func (a *WAFRegionalStub) GetPermissionPolicy(ctx workflow.Context, input *waf.GetPermissionPolicyInput) (*waf.GetPermissionPolicyOutput, error) {
	var output waf.GetPermissionPolicyOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetPermissionPolicy, input).Get(ctx, &output)
	return &output, err
}

func (a *WAFRegionalStub) GetPermissionPolicyAsync(ctx workflow.Context, input *waf.GetPermissionPolicyInput) *WafGetPermissionPolicyResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetPermissionPolicy, input)
	return &WafGetPermissionPolicyResult{Result: future}
}

func (a *WAFRegionalStub) GetRateBasedRule(ctx workflow.Context, input *waf.GetRateBasedRuleInput) (*waf.GetRateBasedRuleOutput, error) {
	var output waf.GetRateBasedRuleOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetRateBasedRule, input).Get(ctx, &output)
	return &output, err
}

func (a *WAFRegionalStub) GetRateBasedRuleAsync(ctx workflow.Context, input *waf.GetRateBasedRuleInput) *WafGetRateBasedRuleResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetRateBasedRule, input)
	return &WafGetRateBasedRuleResult{Result: future}
}

func (a *WAFRegionalStub) GetRateBasedRuleManagedKeys(ctx workflow.Context, input *waf.GetRateBasedRuleManagedKeysInput) (*waf.GetRateBasedRuleManagedKeysOutput, error) {
	var output waf.GetRateBasedRuleManagedKeysOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetRateBasedRuleManagedKeys, input).Get(ctx, &output)
	return &output, err
}

func (a *WAFRegionalStub) GetRateBasedRuleManagedKeysAsync(ctx workflow.Context, input *waf.GetRateBasedRuleManagedKeysInput) *WafGetRateBasedRuleManagedKeysResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetRateBasedRuleManagedKeys, input)
	return &WafGetRateBasedRuleManagedKeysResult{Result: future}
}

func (a *WAFRegionalStub) GetRegexMatchSet(ctx workflow.Context, input *waf.GetRegexMatchSetInput) (*waf.GetRegexMatchSetOutput, error) {
	var output waf.GetRegexMatchSetOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetRegexMatchSet, input).Get(ctx, &output)
	return &output, err
}

func (a *WAFRegionalStub) GetRegexMatchSetAsync(ctx workflow.Context, input *waf.GetRegexMatchSetInput) *WafGetRegexMatchSetResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetRegexMatchSet, input)
	return &WafGetRegexMatchSetResult{Result: future}
}

func (a *WAFRegionalStub) GetRegexPatternSet(ctx workflow.Context, input *waf.GetRegexPatternSetInput) (*waf.GetRegexPatternSetOutput, error) {
	var output waf.GetRegexPatternSetOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetRegexPatternSet, input).Get(ctx, &output)
	return &output, err
}

func (a *WAFRegionalStub) GetRegexPatternSetAsync(ctx workflow.Context, input *waf.GetRegexPatternSetInput) *WafGetRegexPatternSetResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetRegexPatternSet, input)
	return &WafGetRegexPatternSetResult{Result: future}
}

func (a *WAFRegionalStub) GetRule(ctx workflow.Context, input *waf.GetRuleInput) (*waf.GetRuleOutput, error) {
	var output waf.GetRuleOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetRule, input).Get(ctx, &output)
	return &output, err
}

func (a *WAFRegionalStub) GetRuleAsync(ctx workflow.Context, input *waf.GetRuleInput) *WafGetRuleResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetRule, input)
	return &WafGetRuleResult{Result: future}
}

func (a *WAFRegionalStub) GetRuleGroup(ctx workflow.Context, input *waf.GetRuleGroupInput) (*waf.GetRuleGroupOutput, error) {
	var output waf.GetRuleGroupOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetRuleGroup, input).Get(ctx, &output)
	return &output, err
}

func (a *WAFRegionalStub) GetRuleGroupAsync(ctx workflow.Context, input *waf.GetRuleGroupInput) *WafGetRuleGroupResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetRuleGroup, input)
	return &WafGetRuleGroupResult{Result: future}
}

func (a *WAFRegionalStub) GetSampledRequests(ctx workflow.Context, input *waf.GetSampledRequestsInput) (*waf.GetSampledRequestsOutput, error) {
	var output waf.GetSampledRequestsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetSampledRequests, input).Get(ctx, &output)
	return &output, err
}

func (a *WAFRegionalStub) GetSampledRequestsAsync(ctx workflow.Context, input *waf.GetSampledRequestsInput) *WafGetSampledRequestsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetSampledRequests, input)
	return &WafGetSampledRequestsResult{Result: future}
}

func (a *WAFRegionalStub) GetSizeConstraintSet(ctx workflow.Context, input *waf.GetSizeConstraintSetInput) (*waf.GetSizeConstraintSetOutput, error) {
	var output waf.GetSizeConstraintSetOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetSizeConstraintSet, input).Get(ctx, &output)
	return &output, err
}

func (a *WAFRegionalStub) GetSizeConstraintSetAsync(ctx workflow.Context, input *waf.GetSizeConstraintSetInput) *WafGetSizeConstraintSetResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetSizeConstraintSet, input)
	return &WafGetSizeConstraintSetResult{Result: future}
}

func (a *WAFRegionalStub) GetSqlInjectionMatchSet(ctx workflow.Context, input *waf.GetSqlInjectionMatchSetInput) (*waf.GetSqlInjectionMatchSetOutput, error) {
	var output waf.GetSqlInjectionMatchSetOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetSqlInjectionMatchSet, input).Get(ctx, &output)
	return &output, err
}

func (a *WAFRegionalStub) GetSqlInjectionMatchSetAsync(ctx workflow.Context, input *waf.GetSqlInjectionMatchSetInput) *WafGetSqlInjectionMatchSetResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetSqlInjectionMatchSet, input)
	return &WafGetSqlInjectionMatchSetResult{Result: future}
}

func (a *WAFRegionalStub) GetWebACL(ctx workflow.Context, input *waf.GetWebACLInput) (*waf.GetWebACLOutput, error) {
	var output waf.GetWebACLOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetWebACL, input).Get(ctx, &output)
	return &output, err
}

func (a *WAFRegionalStub) GetWebACLAsync(ctx workflow.Context, input *waf.GetWebACLInput) *WafGetWebACLResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetWebACL, input)
	return &WafGetWebACLResult{Result: future}
}

func (a *WAFRegionalStub) GetWebACLForResource(ctx workflow.Context, input *wafregional.GetWebACLForResourceInput) (*wafregional.GetWebACLForResourceOutput, error) {
	var output wafregional.GetWebACLForResourceOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetWebACLForResource, input).Get(ctx, &output)
	return &output, err
}

func (a *WAFRegionalStub) GetWebACLForResourceAsync(ctx workflow.Context, input *wafregional.GetWebACLForResourceInput) *WafregionalGetWebACLForResourceResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetWebACLForResource, input)
	return &WafregionalGetWebACLForResourceResult{Result: future}
}

func (a *WAFRegionalStub) GetXssMatchSet(ctx workflow.Context, input *waf.GetXssMatchSetInput) (*waf.GetXssMatchSetOutput, error) {
	var output waf.GetXssMatchSetOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetXssMatchSet, input).Get(ctx, &output)
	return &output, err
}

func (a *WAFRegionalStub) GetXssMatchSetAsync(ctx workflow.Context, input *waf.GetXssMatchSetInput) *WafGetXssMatchSetResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetXssMatchSet, input)
	return &WafGetXssMatchSetResult{Result: future}
}

func (a *WAFRegionalStub) ListActivatedRulesInRuleGroup(ctx workflow.Context, input *waf.ListActivatedRulesInRuleGroupInput) (*waf.ListActivatedRulesInRuleGroupOutput, error) {
	var output waf.ListActivatedRulesInRuleGroupOutput
	err := workflow.ExecuteActivity(ctx, a.activities.ListActivatedRulesInRuleGroup, input).Get(ctx, &output)
	return &output, err
}

func (a *WAFRegionalStub) ListActivatedRulesInRuleGroupAsync(ctx workflow.Context, input *waf.ListActivatedRulesInRuleGroupInput) *WafListActivatedRulesInRuleGroupResult {
	future := workflow.ExecuteActivity(ctx, a.activities.ListActivatedRulesInRuleGroup, input)
	return &WafListActivatedRulesInRuleGroupResult{Result: future}
}

func (a *WAFRegionalStub) ListByteMatchSets(ctx workflow.Context, input *waf.ListByteMatchSetsInput) (*waf.ListByteMatchSetsOutput, error) {
	var output waf.ListByteMatchSetsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.ListByteMatchSets, input).Get(ctx, &output)
	return &output, err
}

func (a *WAFRegionalStub) ListByteMatchSetsAsync(ctx workflow.Context, input *waf.ListByteMatchSetsInput) *WafListByteMatchSetsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.ListByteMatchSets, input)
	return &WafListByteMatchSetsResult{Result: future}
}

func (a *WAFRegionalStub) ListGeoMatchSets(ctx workflow.Context, input *waf.ListGeoMatchSetsInput) (*waf.ListGeoMatchSetsOutput, error) {
	var output waf.ListGeoMatchSetsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.ListGeoMatchSets, input).Get(ctx, &output)
	return &output, err
}

func (a *WAFRegionalStub) ListGeoMatchSetsAsync(ctx workflow.Context, input *waf.ListGeoMatchSetsInput) *WafListGeoMatchSetsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.ListGeoMatchSets, input)
	return &WafListGeoMatchSetsResult{Result: future}
}

func (a *WAFRegionalStub) ListIPSets(ctx workflow.Context, input *waf.ListIPSetsInput) (*waf.ListIPSetsOutput, error) {
	var output waf.ListIPSetsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.ListIPSets, input).Get(ctx, &output)
	return &output, err
}

func (a *WAFRegionalStub) ListIPSetsAsync(ctx workflow.Context, input *waf.ListIPSetsInput) *WafListIPSetsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.ListIPSets, input)
	return &WafListIPSetsResult{Result: future}
}

func (a *WAFRegionalStub) ListLoggingConfigurations(ctx workflow.Context, input *waf.ListLoggingConfigurationsInput) (*waf.ListLoggingConfigurationsOutput, error) {
	var output waf.ListLoggingConfigurationsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.ListLoggingConfigurations, input).Get(ctx, &output)
	return &output, err
}

func (a *WAFRegionalStub) ListLoggingConfigurationsAsync(ctx workflow.Context, input *waf.ListLoggingConfigurationsInput) *WafListLoggingConfigurationsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.ListLoggingConfigurations, input)
	return &WafListLoggingConfigurationsResult{Result: future}
}

func (a *WAFRegionalStub) ListRateBasedRules(ctx workflow.Context, input *waf.ListRateBasedRulesInput) (*waf.ListRateBasedRulesOutput, error) {
	var output waf.ListRateBasedRulesOutput
	err := workflow.ExecuteActivity(ctx, a.activities.ListRateBasedRules, input).Get(ctx, &output)
	return &output, err
}

func (a *WAFRegionalStub) ListRateBasedRulesAsync(ctx workflow.Context, input *waf.ListRateBasedRulesInput) *WafListRateBasedRulesResult {
	future := workflow.ExecuteActivity(ctx, a.activities.ListRateBasedRules, input)
	return &WafListRateBasedRulesResult{Result: future}
}

func (a *WAFRegionalStub) ListRegexMatchSets(ctx workflow.Context, input *waf.ListRegexMatchSetsInput) (*waf.ListRegexMatchSetsOutput, error) {
	var output waf.ListRegexMatchSetsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.ListRegexMatchSets, input).Get(ctx, &output)
	return &output, err
}

func (a *WAFRegionalStub) ListRegexMatchSetsAsync(ctx workflow.Context, input *waf.ListRegexMatchSetsInput) *WafListRegexMatchSetsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.ListRegexMatchSets, input)
	return &WafListRegexMatchSetsResult{Result: future}
}

func (a *WAFRegionalStub) ListRegexPatternSets(ctx workflow.Context, input *waf.ListRegexPatternSetsInput) (*waf.ListRegexPatternSetsOutput, error) {
	var output waf.ListRegexPatternSetsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.ListRegexPatternSets, input).Get(ctx, &output)
	return &output, err
}

func (a *WAFRegionalStub) ListRegexPatternSetsAsync(ctx workflow.Context, input *waf.ListRegexPatternSetsInput) *WafListRegexPatternSetsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.ListRegexPatternSets, input)
	return &WafListRegexPatternSetsResult{Result: future}
}

func (a *WAFRegionalStub) ListResourcesForWebACL(ctx workflow.Context, input *wafregional.ListResourcesForWebACLInput) (*wafregional.ListResourcesForWebACLOutput, error) {
	var output wafregional.ListResourcesForWebACLOutput
	err := workflow.ExecuteActivity(ctx, a.activities.ListResourcesForWebACL, input).Get(ctx, &output)
	return &output, err
}

func (a *WAFRegionalStub) ListResourcesForWebACLAsync(ctx workflow.Context, input *wafregional.ListResourcesForWebACLInput) *WafregionalListResourcesForWebACLResult {
	future := workflow.ExecuteActivity(ctx, a.activities.ListResourcesForWebACL, input)
	return &WafregionalListResourcesForWebACLResult{Result: future}
}

func (a *WAFRegionalStub) ListRuleGroups(ctx workflow.Context, input *waf.ListRuleGroupsInput) (*waf.ListRuleGroupsOutput, error) {
	var output waf.ListRuleGroupsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.ListRuleGroups, input).Get(ctx, &output)
	return &output, err
}

func (a *WAFRegionalStub) ListRuleGroupsAsync(ctx workflow.Context, input *waf.ListRuleGroupsInput) *WafListRuleGroupsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.ListRuleGroups, input)
	return &WafListRuleGroupsResult{Result: future}
}

func (a *WAFRegionalStub) ListRules(ctx workflow.Context, input *waf.ListRulesInput) (*waf.ListRulesOutput, error) {
	var output waf.ListRulesOutput
	err := workflow.ExecuteActivity(ctx, a.activities.ListRules, input).Get(ctx, &output)
	return &output, err
}

func (a *WAFRegionalStub) ListRulesAsync(ctx workflow.Context, input *waf.ListRulesInput) *WafListRulesResult {
	future := workflow.ExecuteActivity(ctx, a.activities.ListRules, input)
	return &WafListRulesResult{Result: future}
}

func (a *WAFRegionalStub) ListSizeConstraintSets(ctx workflow.Context, input *waf.ListSizeConstraintSetsInput) (*waf.ListSizeConstraintSetsOutput, error) {
	var output waf.ListSizeConstraintSetsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.ListSizeConstraintSets, input).Get(ctx, &output)
	return &output, err
}

func (a *WAFRegionalStub) ListSizeConstraintSetsAsync(ctx workflow.Context, input *waf.ListSizeConstraintSetsInput) *WafListSizeConstraintSetsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.ListSizeConstraintSets, input)
	return &WafListSizeConstraintSetsResult{Result: future}
}

func (a *WAFRegionalStub) ListSqlInjectionMatchSets(ctx workflow.Context, input *waf.ListSqlInjectionMatchSetsInput) (*waf.ListSqlInjectionMatchSetsOutput, error) {
	var output waf.ListSqlInjectionMatchSetsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.ListSqlInjectionMatchSets, input).Get(ctx, &output)
	return &output, err
}

func (a *WAFRegionalStub) ListSqlInjectionMatchSetsAsync(ctx workflow.Context, input *waf.ListSqlInjectionMatchSetsInput) *WafListSqlInjectionMatchSetsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.ListSqlInjectionMatchSets, input)
	return &WafListSqlInjectionMatchSetsResult{Result: future}
}

func (a *WAFRegionalStub) ListSubscribedRuleGroups(ctx workflow.Context, input *waf.ListSubscribedRuleGroupsInput) (*waf.ListSubscribedRuleGroupsOutput, error) {
	var output waf.ListSubscribedRuleGroupsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.ListSubscribedRuleGroups, input).Get(ctx, &output)
	return &output, err
}

func (a *WAFRegionalStub) ListSubscribedRuleGroupsAsync(ctx workflow.Context, input *waf.ListSubscribedRuleGroupsInput) *WafListSubscribedRuleGroupsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.ListSubscribedRuleGroups, input)
	return &WafListSubscribedRuleGroupsResult{Result: future}
}

func (a *WAFRegionalStub) ListTagsForResource(ctx workflow.Context, input *waf.ListTagsForResourceInput) (*waf.ListTagsForResourceOutput, error) {
	var output waf.ListTagsForResourceOutput
	err := workflow.ExecuteActivity(ctx, a.activities.ListTagsForResource, input).Get(ctx, &output)
	return &output, err
}

func (a *WAFRegionalStub) ListTagsForResourceAsync(ctx workflow.Context, input *waf.ListTagsForResourceInput) *WafListTagsForResourceResult {
	future := workflow.ExecuteActivity(ctx, a.activities.ListTagsForResource, input)
	return &WafListTagsForResourceResult{Result: future}
}

func (a *WAFRegionalStub) ListWebACLs(ctx workflow.Context, input *waf.ListWebACLsInput) (*waf.ListWebACLsOutput, error) {
	var output waf.ListWebACLsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.ListWebACLs, input).Get(ctx, &output)
	return &output, err
}

func (a *WAFRegionalStub) ListWebACLsAsync(ctx workflow.Context, input *waf.ListWebACLsInput) *WafListWebACLsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.ListWebACLs, input)
	return &WafListWebACLsResult{Result: future}
}

func (a *WAFRegionalStub) ListXssMatchSets(ctx workflow.Context, input *waf.ListXssMatchSetsInput) (*waf.ListXssMatchSetsOutput, error) {
	var output waf.ListXssMatchSetsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.ListXssMatchSets, input).Get(ctx, &output)
	return &output, err
}

func (a *WAFRegionalStub) ListXssMatchSetsAsync(ctx workflow.Context, input *waf.ListXssMatchSetsInput) *WafListXssMatchSetsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.ListXssMatchSets, input)
	return &WafListXssMatchSetsResult{Result: future}
}

func (a *WAFRegionalStub) PutLoggingConfiguration(ctx workflow.Context, input *waf.PutLoggingConfigurationInput) (*waf.PutLoggingConfigurationOutput, error) {
	var output waf.PutLoggingConfigurationOutput
	err := workflow.ExecuteActivity(ctx, a.activities.PutLoggingConfiguration, input).Get(ctx, &output)
	return &output, err
}

func (a *WAFRegionalStub) PutLoggingConfigurationAsync(ctx workflow.Context, input *waf.PutLoggingConfigurationInput) *WafPutLoggingConfigurationResult {
	future := workflow.ExecuteActivity(ctx, a.activities.PutLoggingConfiguration, input)
	return &WafPutLoggingConfigurationResult{Result: future}
}

func (a *WAFRegionalStub) PutPermissionPolicy(ctx workflow.Context, input *waf.PutPermissionPolicyInput) (*waf.PutPermissionPolicyOutput, error) {
	var output waf.PutPermissionPolicyOutput
	err := workflow.ExecuteActivity(ctx, a.activities.PutPermissionPolicy, input).Get(ctx, &output)
	return &output, err
}

func (a *WAFRegionalStub) PutPermissionPolicyAsync(ctx workflow.Context, input *waf.PutPermissionPolicyInput) *WafPutPermissionPolicyResult {
	future := workflow.ExecuteActivity(ctx, a.activities.PutPermissionPolicy, input)
	return &WafPutPermissionPolicyResult{Result: future}
}

func (a *WAFRegionalStub) TagResource(ctx workflow.Context, input *waf.TagResourceInput) (*waf.TagResourceOutput, error) {
	var output waf.TagResourceOutput
	err := workflow.ExecuteActivity(ctx, a.activities.TagResource, input).Get(ctx, &output)
	return &output, err
}

func (a *WAFRegionalStub) TagResourceAsync(ctx workflow.Context, input *waf.TagResourceInput) *WafTagResourceResult {
	future := workflow.ExecuteActivity(ctx, a.activities.TagResource, input)
	return &WafTagResourceResult{Result: future}
}

func (a *WAFRegionalStub) UntagResource(ctx workflow.Context, input *waf.UntagResourceInput) (*waf.UntagResourceOutput, error) {
	var output waf.UntagResourceOutput
	err := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input).Get(ctx, &output)
	return &output, err
}

func (a *WAFRegionalStub) UntagResourceAsync(ctx workflow.Context, input *waf.UntagResourceInput) *WafUntagResourceResult {
	future := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input)
	return &WafUntagResourceResult{Result: future}
}

func (a *WAFRegionalStub) UpdateByteMatchSet(ctx workflow.Context, input *waf.UpdateByteMatchSetInput) (*waf.UpdateByteMatchSetOutput, error) {
	var output waf.UpdateByteMatchSetOutput
	err := workflow.ExecuteActivity(ctx, a.activities.UpdateByteMatchSet, input).Get(ctx, &output)
	return &output, err
}

func (a *WAFRegionalStub) UpdateByteMatchSetAsync(ctx workflow.Context, input *waf.UpdateByteMatchSetInput) *WafUpdateByteMatchSetResult {
	future := workflow.ExecuteActivity(ctx, a.activities.UpdateByteMatchSet, input)
	return &WafUpdateByteMatchSetResult{Result: future}
}

func (a *WAFRegionalStub) UpdateGeoMatchSet(ctx workflow.Context, input *waf.UpdateGeoMatchSetInput) (*waf.UpdateGeoMatchSetOutput, error) {
	var output waf.UpdateGeoMatchSetOutput
	err := workflow.ExecuteActivity(ctx, a.activities.UpdateGeoMatchSet, input).Get(ctx, &output)
	return &output, err
}

func (a *WAFRegionalStub) UpdateGeoMatchSetAsync(ctx workflow.Context, input *waf.UpdateGeoMatchSetInput) *WafUpdateGeoMatchSetResult {
	future := workflow.ExecuteActivity(ctx, a.activities.UpdateGeoMatchSet, input)
	return &WafUpdateGeoMatchSetResult{Result: future}
}

func (a *WAFRegionalStub) UpdateIPSet(ctx workflow.Context, input *waf.UpdateIPSetInput) (*waf.UpdateIPSetOutput, error) {
	var output waf.UpdateIPSetOutput
	err := workflow.ExecuteActivity(ctx, a.activities.UpdateIPSet, input).Get(ctx, &output)
	return &output, err
}

func (a *WAFRegionalStub) UpdateIPSetAsync(ctx workflow.Context, input *waf.UpdateIPSetInput) *WafUpdateIPSetResult {
	future := workflow.ExecuteActivity(ctx, a.activities.UpdateIPSet, input)
	return &WafUpdateIPSetResult{Result: future}
}

func (a *WAFRegionalStub) UpdateRateBasedRule(ctx workflow.Context, input *waf.UpdateRateBasedRuleInput) (*waf.UpdateRateBasedRuleOutput, error) {
	var output waf.UpdateRateBasedRuleOutput
	err := workflow.ExecuteActivity(ctx, a.activities.UpdateRateBasedRule, input).Get(ctx, &output)
	return &output, err
}

func (a *WAFRegionalStub) UpdateRateBasedRuleAsync(ctx workflow.Context, input *waf.UpdateRateBasedRuleInput) *WafUpdateRateBasedRuleResult {
	future := workflow.ExecuteActivity(ctx, a.activities.UpdateRateBasedRule, input)
	return &WafUpdateRateBasedRuleResult{Result: future}
}

func (a *WAFRegionalStub) UpdateRegexMatchSet(ctx workflow.Context, input *waf.UpdateRegexMatchSetInput) (*waf.UpdateRegexMatchSetOutput, error) {
	var output waf.UpdateRegexMatchSetOutput
	err := workflow.ExecuteActivity(ctx, a.activities.UpdateRegexMatchSet, input).Get(ctx, &output)
	return &output, err
}

func (a *WAFRegionalStub) UpdateRegexMatchSetAsync(ctx workflow.Context, input *waf.UpdateRegexMatchSetInput) *WafUpdateRegexMatchSetResult {
	future := workflow.ExecuteActivity(ctx, a.activities.UpdateRegexMatchSet, input)
	return &WafUpdateRegexMatchSetResult{Result: future}
}

func (a *WAFRegionalStub) UpdateRegexPatternSet(ctx workflow.Context, input *waf.UpdateRegexPatternSetInput) (*waf.UpdateRegexPatternSetOutput, error) {
	var output waf.UpdateRegexPatternSetOutput
	err := workflow.ExecuteActivity(ctx, a.activities.UpdateRegexPatternSet, input).Get(ctx, &output)
	return &output, err
}

func (a *WAFRegionalStub) UpdateRegexPatternSetAsync(ctx workflow.Context, input *waf.UpdateRegexPatternSetInput) *WafUpdateRegexPatternSetResult {
	future := workflow.ExecuteActivity(ctx, a.activities.UpdateRegexPatternSet, input)
	return &WafUpdateRegexPatternSetResult{Result: future}
}

func (a *WAFRegionalStub) UpdateRule(ctx workflow.Context, input *waf.UpdateRuleInput) (*waf.UpdateRuleOutput, error) {
	var output waf.UpdateRuleOutput
	err := workflow.ExecuteActivity(ctx, a.activities.UpdateRule, input).Get(ctx, &output)
	return &output, err
}

func (a *WAFRegionalStub) UpdateRuleAsync(ctx workflow.Context, input *waf.UpdateRuleInput) *WafUpdateRuleResult {
	future := workflow.ExecuteActivity(ctx, a.activities.UpdateRule, input)
	return &WafUpdateRuleResult{Result: future}
}

func (a *WAFRegionalStub) UpdateRuleGroup(ctx workflow.Context, input *waf.UpdateRuleGroupInput) (*waf.UpdateRuleGroupOutput, error) {
	var output waf.UpdateRuleGroupOutput
	err := workflow.ExecuteActivity(ctx, a.activities.UpdateRuleGroup, input).Get(ctx, &output)
	return &output, err
}

func (a *WAFRegionalStub) UpdateRuleGroupAsync(ctx workflow.Context, input *waf.UpdateRuleGroupInput) *WafUpdateRuleGroupResult {
	future := workflow.ExecuteActivity(ctx, a.activities.UpdateRuleGroup, input)
	return &WafUpdateRuleGroupResult{Result: future}
}

func (a *WAFRegionalStub) UpdateSizeConstraintSet(ctx workflow.Context, input *waf.UpdateSizeConstraintSetInput) (*waf.UpdateSizeConstraintSetOutput, error) {
	var output waf.UpdateSizeConstraintSetOutput
	err := workflow.ExecuteActivity(ctx, a.activities.UpdateSizeConstraintSet, input).Get(ctx, &output)
	return &output, err
}

func (a *WAFRegionalStub) UpdateSizeConstraintSetAsync(ctx workflow.Context, input *waf.UpdateSizeConstraintSetInput) *WafUpdateSizeConstraintSetResult {
	future := workflow.ExecuteActivity(ctx, a.activities.UpdateSizeConstraintSet, input)
	return &WafUpdateSizeConstraintSetResult{Result: future}
}

func (a *WAFRegionalStub) UpdateSqlInjectionMatchSet(ctx workflow.Context, input *waf.UpdateSqlInjectionMatchSetInput) (*waf.UpdateSqlInjectionMatchSetOutput, error) {
	var output waf.UpdateSqlInjectionMatchSetOutput
	err := workflow.ExecuteActivity(ctx, a.activities.UpdateSqlInjectionMatchSet, input).Get(ctx, &output)
	return &output, err
}

func (a *WAFRegionalStub) UpdateSqlInjectionMatchSetAsync(ctx workflow.Context, input *waf.UpdateSqlInjectionMatchSetInput) *WafUpdateSqlInjectionMatchSetResult {
	future := workflow.ExecuteActivity(ctx, a.activities.UpdateSqlInjectionMatchSet, input)
	return &WafUpdateSqlInjectionMatchSetResult{Result: future}
}

func (a *WAFRegionalStub) UpdateWebACL(ctx workflow.Context, input *waf.UpdateWebACLInput) (*waf.UpdateWebACLOutput, error) {
	var output waf.UpdateWebACLOutput
	err := workflow.ExecuteActivity(ctx, a.activities.UpdateWebACL, input).Get(ctx, &output)
	return &output, err
}

func (a *WAFRegionalStub) UpdateWebACLAsync(ctx workflow.Context, input *waf.UpdateWebACLInput) *WafUpdateWebACLResult {
	future := workflow.ExecuteActivity(ctx, a.activities.UpdateWebACL, input)
	return &WafUpdateWebACLResult{Result: future}
}

func (a *WAFRegionalStub) UpdateXssMatchSet(ctx workflow.Context, input *waf.UpdateXssMatchSetInput) (*waf.UpdateXssMatchSetOutput, error) {
	var output waf.UpdateXssMatchSetOutput
	err := workflow.ExecuteActivity(ctx, a.activities.UpdateXssMatchSet, input).Get(ctx, &output)
	return &output, err
}

func (a *WAFRegionalStub) UpdateXssMatchSetAsync(ctx workflow.Context, input *waf.UpdateXssMatchSetInput) *WafUpdateXssMatchSetResult {
	future := workflow.ExecuteActivity(ctx, a.activities.UpdateXssMatchSet, input)
	return &WafUpdateXssMatchSetResult{Result: future}
}
