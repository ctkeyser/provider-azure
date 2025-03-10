/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"

// GetCondition of this ResourcePolicyRemediation.
func (mg *ResourcePolicyRemediation) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this ResourcePolicyRemediation.
func (mg *ResourcePolicyRemediation) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this ResourcePolicyRemediation.
func (mg *ResourcePolicyRemediation) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this ResourcePolicyRemediation.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *ResourcePolicyRemediation) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetPublishConnectionDetailsTo of this ResourcePolicyRemediation.
func (mg *ResourcePolicyRemediation) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this ResourcePolicyRemediation.
func (mg *ResourcePolicyRemediation) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this ResourcePolicyRemediation.
func (mg *ResourcePolicyRemediation) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this ResourcePolicyRemediation.
func (mg *ResourcePolicyRemediation) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this ResourcePolicyRemediation.
func (mg *ResourcePolicyRemediation) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this ResourcePolicyRemediation.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *ResourcePolicyRemediation) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetPublishConnectionDetailsTo of this ResourcePolicyRemediation.
func (mg *ResourcePolicyRemediation) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this ResourcePolicyRemediation.
func (mg *ResourcePolicyRemediation) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this SubscriptionPolicyRemediation.
func (mg *SubscriptionPolicyRemediation) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this SubscriptionPolicyRemediation.
func (mg *SubscriptionPolicyRemediation) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this SubscriptionPolicyRemediation.
func (mg *SubscriptionPolicyRemediation) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this SubscriptionPolicyRemediation.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *SubscriptionPolicyRemediation) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetPublishConnectionDetailsTo of this SubscriptionPolicyRemediation.
func (mg *SubscriptionPolicyRemediation) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this SubscriptionPolicyRemediation.
func (mg *SubscriptionPolicyRemediation) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this SubscriptionPolicyRemediation.
func (mg *SubscriptionPolicyRemediation) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this SubscriptionPolicyRemediation.
func (mg *SubscriptionPolicyRemediation) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this SubscriptionPolicyRemediation.
func (mg *SubscriptionPolicyRemediation) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this SubscriptionPolicyRemediation.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *SubscriptionPolicyRemediation) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetPublishConnectionDetailsTo of this SubscriptionPolicyRemediation.
func (mg *SubscriptionPolicyRemediation) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this SubscriptionPolicyRemediation.
func (mg *SubscriptionPolicyRemediation) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}
