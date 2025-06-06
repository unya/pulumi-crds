// Code generated by crd2pulumi DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1alpha1

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
	case "kubernetes:argoproj.io/v1alpha1:EventBus":
		r = &EventBus{}
	case "kubernetes:argoproj.io/v1alpha1:EventBusList":
		r = &EventBusList{}
	case "kubernetes:argoproj.io/v1alpha1:EventBusPatch":
		r = &EventBusPatch{}
	case "kubernetes:argoproj.io/v1alpha1:EventSource":
		r = &EventSource{}
	case "kubernetes:argoproj.io/v1alpha1:EventSourceList":
		r = &EventSourceList{}
	case "kubernetes:argoproj.io/v1alpha1:EventSourcePatch":
		r = &EventSourcePatch{}
	case "kubernetes:argoproj.io/v1alpha1:Sensor":
		r = &Sensor{}
	case "kubernetes:argoproj.io/v1alpha1:SensorList":
		r = &SensorList{}
	case "kubernetes:argoproj.io/v1alpha1:SensorPatch":
		r = &SensorPatch{}
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
		"argo-events",
		"argoproj.io/v1alpha1",
		&module{version},
	)
}
