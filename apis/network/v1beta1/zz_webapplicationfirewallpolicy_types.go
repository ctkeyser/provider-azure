/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type CustomRulesObservation struct {
}

type CustomRulesParameters struct {

	// Type of action. Possible values are Allow, Block and Log.
	// +kubebuilder:validation:Required
	Action *string `json:"action" tf:"action,omitempty"`

	// One or more match_conditions blocks as defined below.
	// +kubebuilder:validation:Required
	MatchConditions []MatchConditionsParameters `json:"matchConditions" tf:"match_conditions,omitempty"`

	// Gets name of the resource that is unique within a policy. This name can be used to access the resource.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Describes priority of the rule. Rules with a lower value will be evaluated before rules with a higher value.
	// +kubebuilder:validation:Required
	Priority *float64 `json:"priority" tf:"priority,omitempty"`

	// Describes the type of rule. Possible values are MatchRule and Invalid.
	// +kubebuilder:validation:Required
	RuleType *string `json:"ruleType" tf:"rule_type,omitempty"`
}

type ExcludedRuleSetObservation struct {
}

type ExcludedRuleSetParameters struct {

	// One or more rule_group block defined below.
	// +kubebuilder:validation:Optional
	RuleGroup []RuleGroupParameters `json:"ruleGroup,omitempty" tf:"rule_group,omitempty"`

	// The rule set type. Possible values: Microsoft_BotManagerRuleSet and OWASP.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// The rule set version. Possible values: 0.1, 1.0, 2.2.9, 3.0, 3.1 and 3.2.
	// +kubebuilder:validation:Optional
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type ManagedRuleSetObservation struct {
}

type ManagedRuleSetParameters struct {

	// One or more rule_group_override block defined below.
	// +kubebuilder:validation:Optional
	RuleGroupOverride []RuleGroupOverrideParameters `json:"ruleGroupOverride,omitempty" tf:"rule_group_override,omitempty"`

	// The rule set type. Possible values: Microsoft_BotManagerRuleSet and OWASP.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// The rule set version. Possible values: 0.1, 1.0, 2.2.9, 3.0, 3.1 and 3.2.
	// +kubebuilder:validation:Required
	Version *string `json:"version" tf:"version,omitempty"`
}

type ManagedRulesExclusionObservation struct {
}

type ManagedRulesExclusionParameters struct {

	// One or more excluded_rule_set block defined below.
	// +kubebuilder:validation:Optional
	ExcludedRuleSet []ExcludedRuleSetParameters `json:"excludedRuleSet,omitempty" tf:"excluded_rule_set,omitempty"`

	// The name of the Match Variable. Possible values: RequestArgKeys, RequestArgNames, RequestArgValues, RequestCookieKeys, RequestCookieNames, RequestCookieValues, RequestHeaderKeys, RequestHeaderNames, RequestHeaderValues.
	// +kubebuilder:validation:Required
	MatchVariable *string `json:"matchVariable" tf:"match_variable,omitempty"`

	// Describes field of the matchVariable collection
	// +kubebuilder:validation:Required
	Selector *string `json:"selector" tf:"selector,omitempty"`

	// Describes operator to be matched. Possible values: Contains, EndsWith, Equals, EqualsAny, StartsWith.
	// +kubebuilder:validation:Required
	SelectorMatchOperator *string `json:"selectorMatchOperator" tf:"selector_match_operator,omitempty"`
}

type ManagedRulesObservation struct {
}

type ManagedRulesParameters struct {

	// One or more exclusion block defined below.
	// +kubebuilder:validation:Optional
	Exclusion []ManagedRulesExclusionParameters `json:"exclusion,omitempty" tf:"exclusion,omitempty"`

	// One or more managed_rule_set block defined below.
	// +kubebuilder:validation:Required
	ManagedRuleSet []ManagedRuleSetParameters `json:"managedRuleSet" tf:"managed_rule_set,omitempty"`
}

type MatchConditionsObservation struct {
}

type MatchConditionsParameters struct {

	// A list of match values.
	// +kubebuilder:validation:Required
	MatchValues []*string `json:"matchValues" tf:"match_values,omitempty"`

	// One or more match_variables blocks as defined below.
	// +kubebuilder:validation:Required
	MatchVariables []MatchVariablesParameters `json:"matchVariables" tf:"match_variables,omitempty"`

	// Describes if this is negate condition or not
	// +kubebuilder:validation:Optional
	NegationCondition *bool `json:"negationCondition,omitempty" tf:"negation_condition,omitempty"`

	// Describes operator to be matched. Possible values are IPMatch, GeoMatch, Equal, Contains, LessThan, GreaterThan, LessThanOrEqual, GreaterThanOrEqual, BeginsWith, EndsWith and Regex.
	// +kubebuilder:validation:Required
	Operator *string `json:"operator" tf:"operator,omitempty"`

	// A list of transformations to do before the match is attempted. Possible values are HtmlEntityDecode, Lowercase, RemoveNulls, Trim, UrlDecode and UrlEncode.
	// +kubebuilder:validation:Optional
	Transforms []*string `json:"transforms,omitempty" tf:"transforms,omitempty"`
}

type MatchVariablesObservation struct {
}

type MatchVariablesParameters struct {

	// Describes field of the matchVariable collection
	// +kubebuilder:validation:Optional
	Selector *string `json:"selector,omitempty" tf:"selector,omitempty"`

	// The name of the Match Variable. Possible values are RemoteAddr, RequestMethod, QueryString, PostArgs, RequestUri, RequestHeaders, RequestBody and RequestCookies.
	// +kubebuilder:validation:Required
	VariableName *string `json:"variableName" tf:"variable_name,omitempty"`
}

type PolicySettingsObservation struct {
}

type PolicySettingsParameters struct {

	// Describes if the policy is in enabled state or disabled state. Defaults to true.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// The File Upload Limit in MB. Accepted values are in the range 1 to 4000. Defaults to 100.
	// +kubebuilder:validation:Optional
	FileUploadLimitInMb *float64 `json:"fileUploadLimitInMb,omitempty" tf:"file_upload_limit_in_mb,omitempty"`

	// The Maximum Request Body Size in KB. Accepted values are in the range 8 to 2000. Defaults to 128.
	// +kubebuilder:validation:Optional
	MaxRequestBodySizeInKb *float64 `json:"maxRequestBodySizeInKb,omitempty" tf:"max_request_body_size_in_kb,omitempty"`

	// Describes if it is in detection mode or prevention mode at the policy level. Valid values are Detection and Prevention. Defaults to Prevention.
	// +kubebuilder:validation:Optional
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`

	// Is Request Body Inspection enabled? Defaults to true.
	// +kubebuilder:validation:Optional
	RequestBodyCheck *bool `json:"requestBodyCheck,omitempty" tf:"request_body_check,omitempty"`
}

type RuleGroupObservation struct {
}

type RuleGroupOverrideObservation struct {
}

type RuleGroupOverrideParameters struct {

	// +kubebuilder:validation:Optional
	DisabledRules []*string `json:"disabledRules,omitempty" tf:"disabled_rules,omitempty"`

	// One or more rule block defined below.
	// +kubebuilder:validation:Optional
	Rule []RuleGroupOverrideRuleParameters `json:"rule,omitempty" tf:"rule,omitempty"`

	// The name of the Rule Group. Possible values are BadBots, crs_20_protocol_violations, crs_21_protocol_anomalies, crs_23_request_limits, crs_30_http_policy, crs_35_bad_robots, crs_40_generic_attacks, crs_41_sql_injection_attacks, crs_41_xss_attacks, crs_42_tight_security, crs_45_trojans, General, GoodBots, Known-CVEs, REQUEST-911-METHOD-ENFORCEMENT, REQUEST-913-SCANNER-DETECTION, REQUEST-920-PROTOCOL-ENFORCEMENT, REQUEST-921-PROTOCOL-ATTACK, REQUEST-930-APPLICATION-ATTACK-LFI, REQUEST-931-APPLICATION-ATTACK-RFI, REQUEST-932-APPLICATION-ATTACK-RCE, REQUEST-933-APPLICATION-ATTACK-PHP, REQUEST-941-APPLICATION-ATTACK-XSS, REQUEST-942-APPLICATION-ATTACK-SQLI, REQUEST-943-APPLICATION-ATTACK-SESSION-FIXATION, REQUEST-944-APPLICATION-ATTACK-JAVA and UnknownBots.
	// +kubebuilder:validation:Required
	RuleGroupName *string `json:"ruleGroupName" tf:"rule_group_name,omitempty"`
}

type RuleGroupOverrideRuleObservation struct {
}

type RuleGroupOverrideRuleParameters struct {

	// Describes the override action to be applied when rule matches. Possible values are Allow, AnomalyScoring, Block and Log.
	// +kubebuilder:validation:Optional
	Action *string `json:"action,omitempty" tf:"action,omitempty"`

	// Describes if the managed rule is in enabled state or disabled state.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Identifier for the managed rule.
	// +kubebuilder:validation:Required
	ID *string `json:"id" tf:"id,omitempty"`
}

type RuleGroupParameters struct {

	// One or more Rule IDs for exclusion.
	// +kubebuilder:validation:Optional
	ExcludedRules []*string `json:"excludedRules,omitempty" tf:"excluded_rules,omitempty"`

	// The name of the Rule Group. Possible values are BadBots, crs_20_protocol_violations, crs_21_protocol_anomalies, crs_23_request_limits, crs_30_http_policy, crs_35_bad_robots, crs_40_generic_attacks, crs_41_sql_injection_attacks, crs_41_xss_attacks, crs_42_tight_security, crs_45_trojans, General, GoodBots, Known-CVEs, REQUEST-911-METHOD-ENFORCEMENT, REQUEST-913-SCANNER-DETECTION, REQUEST-920-PROTOCOL-ENFORCEMENT, REQUEST-921-PROTOCOL-ATTACK, REQUEST-930-APPLICATION-ATTACK-LFI, REQUEST-931-APPLICATION-ATTACK-RFI, REQUEST-932-APPLICATION-ATTACK-RCE, REQUEST-933-APPLICATION-ATTACK-PHP, REQUEST-941-APPLICATION-ATTACK-XSS, REQUEST-942-APPLICATION-ATTACK-SQLI, REQUEST-943-APPLICATION-ATTACK-SESSION-FIXATION, REQUEST-944-APPLICATION-ATTACK-JAVA and UnknownBots.
	// +kubebuilder:validation:Required
	RuleGroupName *string `json:"ruleGroupName" tf:"rule_group_name,omitempty"`
}

type WebApplicationFirewallPolicyObservation struct {

	// A list of HTTP Listener IDs from an azurerm_application_gateway.
	HTTPListenerIds []*string `json:"httpListenerIds,omitempty" tf:"http_listener_ids,omitempty"`

	// The ID of the Web Application Firewall Policy.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A list of URL Path Map Path Rule IDs from an azurerm_application_gateway.
	PathBasedRuleIds []*string `json:"pathBasedRuleIds,omitempty" tf:"path_based_rule_ids,omitempty"`
}

type WebApplicationFirewallPolicyParameters struct {

	// One or more custom_rules blocks as defined below.
	// +kubebuilder:validation:Optional
	CustomRules []CustomRulesParameters `json:"customRules,omitempty" tf:"custom_rules,omitempty"`

	// Resource location. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// A managed_rules blocks as defined below.
	// +kubebuilder:validation:Required
	ManagedRules []ManagedRulesParameters `json:"managedRules" tf:"managed_rules,omitempty"`

	// A policy_settings block as defined below.
	// +kubebuilder:validation:Optional
	PolicySettings []PolicySettingsParameters `json:"policySettings,omitempty" tf:"policy_settings,omitempty"`

	// The name of the resource group. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// A mapping of tags to assign to the Web Application Firewall Policy.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// WebApplicationFirewallPolicySpec defines the desired state of WebApplicationFirewallPolicy
type WebApplicationFirewallPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     WebApplicationFirewallPolicyParameters `json:"forProvider"`
}

// WebApplicationFirewallPolicyStatus defines the observed state of WebApplicationFirewallPolicy.
type WebApplicationFirewallPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        WebApplicationFirewallPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// WebApplicationFirewallPolicy is the Schema for the WebApplicationFirewallPolicys API. Manages a Azure Web Application Firewall Policy instance.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type WebApplicationFirewallPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              WebApplicationFirewallPolicySpec   `json:"spec"`
	Status            WebApplicationFirewallPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// WebApplicationFirewallPolicyList contains a list of WebApplicationFirewallPolicys
type WebApplicationFirewallPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []WebApplicationFirewallPolicy `json:"items"`
}

// Repository type metadata.
var (
	WebApplicationFirewallPolicy_Kind             = "WebApplicationFirewallPolicy"
	WebApplicationFirewallPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: WebApplicationFirewallPolicy_Kind}.String()
	WebApplicationFirewallPolicy_KindAPIVersion   = WebApplicationFirewallPolicy_Kind + "." + CRDGroupVersion.String()
	WebApplicationFirewallPolicy_GroupVersionKind = CRDGroupVersion.WithKind(WebApplicationFirewallPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&WebApplicationFirewallPolicy{}, &WebApplicationFirewallPolicyList{})
}
