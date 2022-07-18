/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	v1beta1 "github.com/upbound/official-providers/provider-azure/apis/azure/v1beta1"
	v1beta11 "github.com/upbound/official-providers/provider-azure/apis/network/v1beta1"
	rconfig "github.com/upbound/official-providers/provider-azure/apis/rconfig"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this FlexibleDatabase.
func (mg *FlexibleDatabase) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceGroupName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ResourceGroupNameRef,
		Selector:     mg.Spec.ForProvider.ResourceGroupNameSelector,
		To: reference.To{
			List:    &v1beta1.ResourceGroupList{},
			Managed: &v1beta1.ResourceGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceGroupName")
	}
	mg.Spec.ForProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceGroupNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ServerName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ServerNameRef,
		Selector:     mg.Spec.ForProvider.ServerNameSelector,
		To: reference.To{
			List:    &FlexibleServerList{},
			Managed: &FlexibleServer{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ServerName")
	}
	mg.Spec.ForProvider.ServerName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ServerNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this FlexibleServer.
func (mg *FlexibleServer) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DelegatedSubnetID),
		Extract:      rconfig.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.DelegatedSubnetIDRef,
		Selector:     mg.Spec.ForProvider.DelegatedSubnetIDSelector,
		To: reference.To{
			List:    &v1beta11.SubnetList{},
			Managed: &v1beta11.Subnet{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.DelegatedSubnetID")
	}
	mg.Spec.ForProvider.DelegatedSubnetID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.DelegatedSubnetIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.PrivateDNSZoneID),
		Extract:      rconfig.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.PrivateDNSZoneIDRef,
		Selector:     mg.Spec.ForProvider.PrivateDNSZoneIDSelector,
		To: reference.To{
			List:    &v1beta11.PrivateDNSZoneList{},
			Managed: &v1beta11.PrivateDNSZone{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.PrivateDNSZoneID")
	}
	mg.Spec.ForProvider.PrivateDNSZoneID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.PrivateDNSZoneIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceGroupName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ResourceGroupNameRef,
		Selector:     mg.Spec.ForProvider.ResourceGroupNameSelector,
		To: reference.To{
			List:    &v1beta1.ResourceGroupList{},
			Managed: &v1beta1.ResourceGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceGroupName")
	}
	mg.Spec.ForProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceGroupNameRef = rsp.ResolvedReference

	return nil
}
