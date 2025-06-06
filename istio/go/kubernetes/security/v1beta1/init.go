// Code generated by crd2pulumi DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "kubernetes:security.istio.io/v1beta1:AuthorizationPolicy":
		r = &AuthorizationPolicy{}
	case "kubernetes:security.istio.io/v1beta1:AuthorizationPolicyList":
		r = &AuthorizationPolicyList{}
	case "kubernetes:security.istio.io/v1beta1:AuthorizationPolicyPatch":
		r = &AuthorizationPolicyPatch{}
	case "kubernetes:security.istio.io/v1beta1:PeerAuthentication":
		r = &PeerAuthentication{}
	case "kubernetes:security.istio.io/v1beta1:PeerAuthenticationList":
		r = &PeerAuthenticationList{}
	case "kubernetes:security.istio.io/v1beta1:PeerAuthenticationPatch":
		r = &PeerAuthenticationPatch{}
	case "kubernetes:security.istio.io/v1beta1:RequestAuthentication":
		r = &RequestAuthentication{}
	case "kubernetes:security.istio.io/v1beta1:RequestAuthenticationList":
		r = &RequestAuthenticationList{}
	case "kubernetes:security.istio.io/v1beta1:RequestAuthenticationPatch":
		r = &RequestAuthenticationPatch{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

func init() {
	version, err := utilities.PkgVersion()
	if err != nil {
		version = semver.Version{Major: 1}
	}
	pulumi.RegisterResourceModule(
		"istio",
		"security.istio.io/v1beta1",
		&module{version},
	)
}
