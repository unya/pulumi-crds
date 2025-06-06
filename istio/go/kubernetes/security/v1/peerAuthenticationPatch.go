// Code generated by crd2pulumi DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Patch resources are used to modify existing Kubernetes resources by using
// Server-Side Apply updates. The name of the resource must be specified, but all other properties are optional. More than
// one patch may be applied to the same resource, and a random FieldManager name will be used for each Patch resource.
// Conflicts will result in an error by default, but can be forced using the "pulumi.com/patchForce" annotation. See the
// [Server-Side Apply Docs](https://www.pulumi.com/registry/packages/kubernetes/how-to-guides/managing-resources-with-server-side-apply/) for
// additional information about using Server-Side Apply to manage Kubernetes resources with Pulumi.
type PeerAuthenticationPatch struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPatchPtrOutput      `pulumi:"metadata"`
	Spec     PeerAuthenticationSpecPatchPtrOutput `pulumi:"spec"`
	Status   pulumi.MapOutput                     `pulumi:"status"`
}

// NewPeerAuthenticationPatch registers a new resource with the given unique name, arguments, and options.
func NewPeerAuthenticationPatch(ctx *pulumi.Context,
	name string, args *PeerAuthenticationPatchArgs, opts ...pulumi.ResourceOption) (*PeerAuthenticationPatch, error) {
	if args == nil {
		args = &PeerAuthenticationPatchArgs{}
	}

	args.ApiVersion = pulumi.StringPtr("security.istio.io/v1")
	args.Kind = pulumi.StringPtr("PeerAuthentication")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("kubernetes:security.istio.io/v1beta1:PeerAuthenticationPatch"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource PeerAuthenticationPatch
	err := ctx.RegisterResource("kubernetes:security.istio.io/v1:PeerAuthenticationPatch", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPeerAuthenticationPatch gets an existing PeerAuthenticationPatch resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPeerAuthenticationPatch(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PeerAuthenticationPatchState, opts ...pulumi.ResourceOption) (*PeerAuthenticationPatch, error) {
	var resource PeerAuthenticationPatch
	err := ctx.ReadResource("kubernetes:security.istio.io/v1:PeerAuthenticationPatch", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PeerAuthenticationPatch resources.
type peerAuthenticationPatchState struct {
}

type PeerAuthenticationPatchState struct {
}

func (PeerAuthenticationPatchState) ElementType() reflect.Type {
	return reflect.TypeOf((*peerAuthenticationPatchState)(nil)).Elem()
}

type peerAuthenticationPatchArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ObjectMetaPatch      `pulumi:"metadata"`
	Spec     *PeerAuthenticationSpecPatch `pulumi:"spec"`
}

// The set of arguments for constructing a PeerAuthenticationPatch resource.
type PeerAuthenticationPatchArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPatchPtrInput
	Spec     PeerAuthenticationSpecPatchPtrInput
}

func (PeerAuthenticationPatchArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*peerAuthenticationPatchArgs)(nil)).Elem()
}

type PeerAuthenticationPatchInput interface {
	pulumi.Input

	ToPeerAuthenticationPatchOutput() PeerAuthenticationPatchOutput
	ToPeerAuthenticationPatchOutputWithContext(ctx context.Context) PeerAuthenticationPatchOutput
}

func (*PeerAuthenticationPatch) ElementType() reflect.Type {
	return reflect.TypeOf((**PeerAuthenticationPatch)(nil)).Elem()
}

func (i *PeerAuthenticationPatch) ToPeerAuthenticationPatchOutput() PeerAuthenticationPatchOutput {
	return i.ToPeerAuthenticationPatchOutputWithContext(context.Background())
}

func (i *PeerAuthenticationPatch) ToPeerAuthenticationPatchOutputWithContext(ctx context.Context) PeerAuthenticationPatchOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PeerAuthenticationPatchOutput)
}

// PeerAuthenticationPatchArrayInput is an input type that accepts PeerAuthenticationPatchArray and PeerAuthenticationPatchArrayOutput values.
// You can construct a concrete instance of `PeerAuthenticationPatchArrayInput` via:
//
//	PeerAuthenticationPatchArray{ PeerAuthenticationPatchArgs{...} }
type PeerAuthenticationPatchArrayInput interface {
	pulumi.Input

	ToPeerAuthenticationPatchArrayOutput() PeerAuthenticationPatchArrayOutput
	ToPeerAuthenticationPatchArrayOutputWithContext(context.Context) PeerAuthenticationPatchArrayOutput
}

type PeerAuthenticationPatchArray []PeerAuthenticationPatchInput

func (PeerAuthenticationPatchArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PeerAuthenticationPatch)(nil)).Elem()
}

func (i PeerAuthenticationPatchArray) ToPeerAuthenticationPatchArrayOutput() PeerAuthenticationPatchArrayOutput {
	return i.ToPeerAuthenticationPatchArrayOutputWithContext(context.Background())
}

func (i PeerAuthenticationPatchArray) ToPeerAuthenticationPatchArrayOutputWithContext(ctx context.Context) PeerAuthenticationPatchArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PeerAuthenticationPatchArrayOutput)
}

// PeerAuthenticationPatchMapInput is an input type that accepts PeerAuthenticationPatchMap and PeerAuthenticationPatchMapOutput values.
// You can construct a concrete instance of `PeerAuthenticationPatchMapInput` via:
//
//	PeerAuthenticationPatchMap{ "key": PeerAuthenticationPatchArgs{...} }
type PeerAuthenticationPatchMapInput interface {
	pulumi.Input

	ToPeerAuthenticationPatchMapOutput() PeerAuthenticationPatchMapOutput
	ToPeerAuthenticationPatchMapOutputWithContext(context.Context) PeerAuthenticationPatchMapOutput
}

type PeerAuthenticationPatchMap map[string]PeerAuthenticationPatchInput

func (PeerAuthenticationPatchMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PeerAuthenticationPatch)(nil)).Elem()
}

func (i PeerAuthenticationPatchMap) ToPeerAuthenticationPatchMapOutput() PeerAuthenticationPatchMapOutput {
	return i.ToPeerAuthenticationPatchMapOutputWithContext(context.Background())
}

func (i PeerAuthenticationPatchMap) ToPeerAuthenticationPatchMapOutputWithContext(ctx context.Context) PeerAuthenticationPatchMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PeerAuthenticationPatchMapOutput)
}

type PeerAuthenticationPatchOutput struct{ *pulumi.OutputState }

func (PeerAuthenticationPatchOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PeerAuthenticationPatch)(nil)).Elem()
}

func (o PeerAuthenticationPatchOutput) ToPeerAuthenticationPatchOutput() PeerAuthenticationPatchOutput {
	return o
}

func (o PeerAuthenticationPatchOutput) ToPeerAuthenticationPatchOutputWithContext(ctx context.Context) PeerAuthenticationPatchOutput {
	return o
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o PeerAuthenticationPatchOutput) ApiVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PeerAuthenticationPatch) pulumi.StringPtrOutput { return v.ApiVersion }).(pulumi.StringPtrOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o PeerAuthenticationPatchOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PeerAuthenticationPatch) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
func (o PeerAuthenticationPatchOutput) Metadata() metav1.ObjectMetaPatchPtrOutput {
	return o.ApplyT(func(v *PeerAuthenticationPatch) metav1.ObjectMetaPatchPtrOutput { return v.Metadata }).(metav1.ObjectMetaPatchPtrOutput)
}

func (o PeerAuthenticationPatchOutput) Spec() PeerAuthenticationSpecPatchPtrOutput {
	return o.ApplyT(func(v *PeerAuthenticationPatch) PeerAuthenticationSpecPatchPtrOutput { return v.Spec }).(PeerAuthenticationSpecPatchPtrOutput)
}

func (o PeerAuthenticationPatchOutput) Status() pulumi.MapOutput {
	return o.ApplyT(func(v *PeerAuthenticationPatch) pulumi.MapOutput { return v.Status }).(pulumi.MapOutput)
}

type PeerAuthenticationPatchArrayOutput struct{ *pulumi.OutputState }

func (PeerAuthenticationPatchArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PeerAuthenticationPatch)(nil)).Elem()
}

func (o PeerAuthenticationPatchArrayOutput) ToPeerAuthenticationPatchArrayOutput() PeerAuthenticationPatchArrayOutput {
	return o
}

func (o PeerAuthenticationPatchArrayOutput) ToPeerAuthenticationPatchArrayOutputWithContext(ctx context.Context) PeerAuthenticationPatchArrayOutput {
	return o
}

func (o PeerAuthenticationPatchArrayOutput) Index(i pulumi.IntInput) PeerAuthenticationPatchOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *PeerAuthenticationPatch {
		return vs[0].([]*PeerAuthenticationPatch)[vs[1].(int)]
	}).(PeerAuthenticationPatchOutput)
}

type PeerAuthenticationPatchMapOutput struct{ *pulumi.OutputState }

func (PeerAuthenticationPatchMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PeerAuthenticationPatch)(nil)).Elem()
}

func (o PeerAuthenticationPatchMapOutput) ToPeerAuthenticationPatchMapOutput() PeerAuthenticationPatchMapOutput {
	return o
}

func (o PeerAuthenticationPatchMapOutput) ToPeerAuthenticationPatchMapOutputWithContext(ctx context.Context) PeerAuthenticationPatchMapOutput {
	return o
}

func (o PeerAuthenticationPatchMapOutput) MapIndex(k pulumi.StringInput) PeerAuthenticationPatchOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *PeerAuthenticationPatch {
		return vs[0].(map[string]*PeerAuthenticationPatch)[vs[1].(string)]
	}).(PeerAuthenticationPatchOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PeerAuthenticationPatchInput)(nil)).Elem(), &PeerAuthenticationPatch{})
	pulumi.RegisterInputType(reflect.TypeOf((*PeerAuthenticationPatchArrayInput)(nil)).Elem(), PeerAuthenticationPatchArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PeerAuthenticationPatchMapInput)(nil)).Elem(), PeerAuthenticationPatchMap{})
	pulumi.RegisterOutputType(PeerAuthenticationPatchOutput{})
	pulumi.RegisterOutputType(PeerAuthenticationPatchArrayOutput{})
	pulumi.RegisterOutputType(PeerAuthenticationPatchMapOutput{})
}
