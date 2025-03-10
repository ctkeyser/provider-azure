/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"

// GetCondition of this LogAnalyticsSolution.
func (mg *LogAnalyticsSolution) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this LogAnalyticsSolution.
func (mg *LogAnalyticsSolution) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this LogAnalyticsSolution.
func (mg *LogAnalyticsSolution) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this LogAnalyticsSolution.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *LogAnalyticsSolution) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetPublishConnectionDetailsTo of this LogAnalyticsSolution.
func (mg *LogAnalyticsSolution) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this LogAnalyticsSolution.
func (mg *LogAnalyticsSolution) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this LogAnalyticsSolution.
func (mg *LogAnalyticsSolution) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this LogAnalyticsSolution.
func (mg *LogAnalyticsSolution) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this LogAnalyticsSolution.
func (mg *LogAnalyticsSolution) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this LogAnalyticsSolution.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *LogAnalyticsSolution) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetPublishConnectionDetailsTo of this LogAnalyticsSolution.
func (mg *LogAnalyticsSolution) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this LogAnalyticsSolution.
func (mg *LogAnalyticsSolution) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}
