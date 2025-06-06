// Code generated by crd2pulumi DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

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
// Vhost is the Schema for the vhosts API
type VhostPatch struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPatchPtrOutput `pulumi:"metadata"`
	Spec     VhostSpecPatchPtrOutput         `pulumi:"spec"`
	Status   VhostStatusPatchPtrOutput       `pulumi:"status"`
}

// NewVhostPatch registers a new resource with the given unique name, arguments, and options.
func NewVhostPatch(ctx *pulumi.Context,
	name string, args *VhostPatchArgs, opts ...pulumi.ResourceOption) (*VhostPatch, error) {
	if args == nil {
		args = &VhostPatchArgs{}
	}

	args.ApiVersion = pulumi.StringPtr("rabbitmq.com/v1beta1")
	args.Kind = pulumi.StringPtr("Vhost")
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource VhostPatch
	err := ctx.RegisterResource("kubernetes:rabbitmq.com/v1beta1:VhostPatch", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVhostPatch gets an existing VhostPatch resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVhostPatch(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VhostPatchState, opts ...pulumi.ResourceOption) (*VhostPatch, error) {
	var resource VhostPatch
	err := ctx.ReadResource("kubernetes:rabbitmq.com/v1beta1:VhostPatch", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VhostPatch resources.
type vhostPatchState struct {
}

type VhostPatchState struct {
}

func (VhostPatchState) ElementType() reflect.Type {
	return reflect.TypeOf((*vhostPatchState)(nil)).Elem()
}

type vhostPatchArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ObjectMetaPatch `pulumi:"metadata"`
	Spec     *VhostSpecPatch         `pulumi:"spec"`
}

// The set of arguments for constructing a VhostPatch resource.
type VhostPatchArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPatchPtrInput
	Spec     VhostSpecPatchPtrInput
}

func (VhostPatchArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vhostPatchArgs)(nil)).Elem()
}

type VhostPatchInput interface {
	pulumi.Input

	ToVhostPatchOutput() VhostPatchOutput
	ToVhostPatchOutputWithContext(ctx context.Context) VhostPatchOutput
}

func (*VhostPatch) ElementType() reflect.Type {
	return reflect.TypeOf((**VhostPatch)(nil)).Elem()
}

func (i *VhostPatch) ToVhostPatchOutput() VhostPatchOutput {
	return i.ToVhostPatchOutputWithContext(context.Background())
}

func (i *VhostPatch) ToVhostPatchOutputWithContext(ctx context.Context) VhostPatchOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VhostPatchOutput)
}

// VhostPatchArrayInput is an input type that accepts VhostPatchArray and VhostPatchArrayOutput values.
// You can construct a concrete instance of `VhostPatchArrayInput` via:
//
//	VhostPatchArray{ VhostPatchArgs{...} }
type VhostPatchArrayInput interface {
	pulumi.Input

	ToVhostPatchArrayOutput() VhostPatchArrayOutput
	ToVhostPatchArrayOutputWithContext(context.Context) VhostPatchArrayOutput
}

type VhostPatchArray []VhostPatchInput

func (VhostPatchArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VhostPatch)(nil)).Elem()
}

func (i VhostPatchArray) ToVhostPatchArrayOutput() VhostPatchArrayOutput {
	return i.ToVhostPatchArrayOutputWithContext(context.Background())
}

func (i VhostPatchArray) ToVhostPatchArrayOutputWithContext(ctx context.Context) VhostPatchArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VhostPatchArrayOutput)
}

// VhostPatchMapInput is an input type that accepts VhostPatchMap and VhostPatchMapOutput values.
// You can construct a concrete instance of `VhostPatchMapInput` via:
//
//	VhostPatchMap{ "key": VhostPatchArgs{...} }
type VhostPatchMapInput interface {
	pulumi.Input

	ToVhostPatchMapOutput() VhostPatchMapOutput
	ToVhostPatchMapOutputWithContext(context.Context) VhostPatchMapOutput
}

type VhostPatchMap map[string]VhostPatchInput

func (VhostPatchMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VhostPatch)(nil)).Elem()
}

func (i VhostPatchMap) ToVhostPatchMapOutput() VhostPatchMapOutput {
	return i.ToVhostPatchMapOutputWithContext(context.Background())
}

func (i VhostPatchMap) ToVhostPatchMapOutputWithContext(ctx context.Context) VhostPatchMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VhostPatchMapOutput)
}

type VhostPatchOutput struct{ *pulumi.OutputState }

func (VhostPatchOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VhostPatch)(nil)).Elem()
}

func (o VhostPatchOutput) ToVhostPatchOutput() VhostPatchOutput {
	return o
}

func (o VhostPatchOutput) ToVhostPatchOutputWithContext(ctx context.Context) VhostPatchOutput {
	return o
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o VhostPatchOutput) ApiVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VhostPatch) pulumi.StringPtrOutput { return v.ApiVersion }).(pulumi.StringPtrOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o VhostPatchOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VhostPatch) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
func (o VhostPatchOutput) Metadata() metav1.ObjectMetaPatchPtrOutput {
	return o.ApplyT(func(v *VhostPatch) metav1.ObjectMetaPatchPtrOutput { return v.Metadata }).(metav1.ObjectMetaPatchPtrOutput)
}

func (o VhostPatchOutput) Spec() VhostSpecPatchPtrOutput {
	return o.ApplyT(func(v *VhostPatch) VhostSpecPatchPtrOutput { return v.Spec }).(VhostSpecPatchPtrOutput)
}

func (o VhostPatchOutput) Status() VhostStatusPatchPtrOutput {
	return o.ApplyT(func(v *VhostPatch) VhostStatusPatchPtrOutput { return v.Status }).(VhostStatusPatchPtrOutput)
}

type VhostPatchArrayOutput struct{ *pulumi.OutputState }

func (VhostPatchArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VhostPatch)(nil)).Elem()
}

func (o VhostPatchArrayOutput) ToVhostPatchArrayOutput() VhostPatchArrayOutput {
	return o
}

func (o VhostPatchArrayOutput) ToVhostPatchArrayOutputWithContext(ctx context.Context) VhostPatchArrayOutput {
	return o
}

func (o VhostPatchArrayOutput) Index(i pulumi.IntInput) VhostPatchOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *VhostPatch {
		return vs[0].([]*VhostPatch)[vs[1].(int)]
	}).(VhostPatchOutput)
}

type VhostPatchMapOutput struct{ *pulumi.OutputState }

func (VhostPatchMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VhostPatch)(nil)).Elem()
}

func (o VhostPatchMapOutput) ToVhostPatchMapOutput() VhostPatchMapOutput {
	return o
}

func (o VhostPatchMapOutput) ToVhostPatchMapOutputWithContext(ctx context.Context) VhostPatchMapOutput {
	return o
}

func (o VhostPatchMapOutput) MapIndex(k pulumi.StringInput) VhostPatchOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *VhostPatch {
		return vs[0].(map[string]*VhostPatch)[vs[1].(string)]
	}).(VhostPatchOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VhostPatchInput)(nil)).Elem(), &VhostPatch{})
	pulumi.RegisterInputType(reflect.TypeOf((*VhostPatchArrayInput)(nil)).Elem(), VhostPatchArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VhostPatchMapInput)(nil)).Elem(), VhostPatchMap{})
	pulumi.RegisterOutputType(VhostPatchOutput{})
	pulumi.RegisterOutputType(VhostPatchArrayOutput{})
	pulumi.RegisterOutputType(VhostPatchMapOutput{})
}
