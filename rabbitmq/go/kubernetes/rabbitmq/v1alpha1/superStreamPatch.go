// Code generated by crd2pulumi DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1alpha1

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
// SuperStream is the Schema for the queues API
type SuperStreamPatch struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPatchPtrOutput `pulumi:"metadata"`
	Spec     SuperStreamSpecPatchPtrOutput   `pulumi:"spec"`
	Status   SuperStreamStatusPatchPtrOutput `pulumi:"status"`
}

// NewSuperStreamPatch registers a new resource with the given unique name, arguments, and options.
func NewSuperStreamPatch(ctx *pulumi.Context,
	name string, args *SuperStreamPatchArgs, opts ...pulumi.ResourceOption) (*SuperStreamPatch, error) {
	if args == nil {
		args = &SuperStreamPatchArgs{}
	}

	args.ApiVersion = pulumi.StringPtr("rabbitmq.com/v1alpha1")
	args.Kind = pulumi.StringPtr("SuperStream")
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource SuperStreamPatch
	err := ctx.RegisterResource("kubernetes:rabbitmq.com/v1alpha1:SuperStreamPatch", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSuperStreamPatch gets an existing SuperStreamPatch resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSuperStreamPatch(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SuperStreamPatchState, opts ...pulumi.ResourceOption) (*SuperStreamPatch, error) {
	var resource SuperStreamPatch
	err := ctx.ReadResource("kubernetes:rabbitmq.com/v1alpha1:SuperStreamPatch", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SuperStreamPatch resources.
type superStreamPatchState struct {
}

type SuperStreamPatchState struct {
}

func (SuperStreamPatchState) ElementType() reflect.Type {
	return reflect.TypeOf((*superStreamPatchState)(nil)).Elem()
}

type superStreamPatchArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ObjectMetaPatch `pulumi:"metadata"`
	Spec     *SuperStreamSpecPatch   `pulumi:"spec"`
}

// The set of arguments for constructing a SuperStreamPatch resource.
type SuperStreamPatchArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPatchPtrInput
	Spec     SuperStreamSpecPatchPtrInput
}

func (SuperStreamPatchArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*superStreamPatchArgs)(nil)).Elem()
}

type SuperStreamPatchInput interface {
	pulumi.Input

	ToSuperStreamPatchOutput() SuperStreamPatchOutput
	ToSuperStreamPatchOutputWithContext(ctx context.Context) SuperStreamPatchOutput
}

func (*SuperStreamPatch) ElementType() reflect.Type {
	return reflect.TypeOf((**SuperStreamPatch)(nil)).Elem()
}

func (i *SuperStreamPatch) ToSuperStreamPatchOutput() SuperStreamPatchOutput {
	return i.ToSuperStreamPatchOutputWithContext(context.Background())
}

func (i *SuperStreamPatch) ToSuperStreamPatchOutputWithContext(ctx context.Context) SuperStreamPatchOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SuperStreamPatchOutput)
}

// SuperStreamPatchArrayInput is an input type that accepts SuperStreamPatchArray and SuperStreamPatchArrayOutput values.
// You can construct a concrete instance of `SuperStreamPatchArrayInput` via:
//
//	SuperStreamPatchArray{ SuperStreamPatchArgs{...} }
type SuperStreamPatchArrayInput interface {
	pulumi.Input

	ToSuperStreamPatchArrayOutput() SuperStreamPatchArrayOutput
	ToSuperStreamPatchArrayOutputWithContext(context.Context) SuperStreamPatchArrayOutput
}

type SuperStreamPatchArray []SuperStreamPatchInput

func (SuperStreamPatchArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SuperStreamPatch)(nil)).Elem()
}

func (i SuperStreamPatchArray) ToSuperStreamPatchArrayOutput() SuperStreamPatchArrayOutput {
	return i.ToSuperStreamPatchArrayOutputWithContext(context.Background())
}

func (i SuperStreamPatchArray) ToSuperStreamPatchArrayOutputWithContext(ctx context.Context) SuperStreamPatchArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SuperStreamPatchArrayOutput)
}

// SuperStreamPatchMapInput is an input type that accepts SuperStreamPatchMap and SuperStreamPatchMapOutput values.
// You can construct a concrete instance of `SuperStreamPatchMapInput` via:
//
//	SuperStreamPatchMap{ "key": SuperStreamPatchArgs{...} }
type SuperStreamPatchMapInput interface {
	pulumi.Input

	ToSuperStreamPatchMapOutput() SuperStreamPatchMapOutput
	ToSuperStreamPatchMapOutputWithContext(context.Context) SuperStreamPatchMapOutput
}

type SuperStreamPatchMap map[string]SuperStreamPatchInput

func (SuperStreamPatchMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SuperStreamPatch)(nil)).Elem()
}

func (i SuperStreamPatchMap) ToSuperStreamPatchMapOutput() SuperStreamPatchMapOutput {
	return i.ToSuperStreamPatchMapOutputWithContext(context.Background())
}

func (i SuperStreamPatchMap) ToSuperStreamPatchMapOutputWithContext(ctx context.Context) SuperStreamPatchMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SuperStreamPatchMapOutput)
}

type SuperStreamPatchOutput struct{ *pulumi.OutputState }

func (SuperStreamPatchOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SuperStreamPatch)(nil)).Elem()
}

func (o SuperStreamPatchOutput) ToSuperStreamPatchOutput() SuperStreamPatchOutput {
	return o
}

func (o SuperStreamPatchOutput) ToSuperStreamPatchOutputWithContext(ctx context.Context) SuperStreamPatchOutput {
	return o
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o SuperStreamPatchOutput) ApiVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SuperStreamPatch) pulumi.StringPtrOutput { return v.ApiVersion }).(pulumi.StringPtrOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o SuperStreamPatchOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SuperStreamPatch) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
func (o SuperStreamPatchOutput) Metadata() metav1.ObjectMetaPatchPtrOutput {
	return o.ApplyT(func(v *SuperStreamPatch) metav1.ObjectMetaPatchPtrOutput { return v.Metadata }).(metav1.ObjectMetaPatchPtrOutput)
}

func (o SuperStreamPatchOutput) Spec() SuperStreamSpecPatchPtrOutput {
	return o.ApplyT(func(v *SuperStreamPatch) SuperStreamSpecPatchPtrOutput { return v.Spec }).(SuperStreamSpecPatchPtrOutput)
}

func (o SuperStreamPatchOutput) Status() SuperStreamStatusPatchPtrOutput {
	return o.ApplyT(func(v *SuperStreamPatch) SuperStreamStatusPatchPtrOutput { return v.Status }).(SuperStreamStatusPatchPtrOutput)
}

type SuperStreamPatchArrayOutput struct{ *pulumi.OutputState }

func (SuperStreamPatchArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SuperStreamPatch)(nil)).Elem()
}

func (o SuperStreamPatchArrayOutput) ToSuperStreamPatchArrayOutput() SuperStreamPatchArrayOutput {
	return o
}

func (o SuperStreamPatchArrayOutput) ToSuperStreamPatchArrayOutputWithContext(ctx context.Context) SuperStreamPatchArrayOutput {
	return o
}

func (o SuperStreamPatchArrayOutput) Index(i pulumi.IntInput) SuperStreamPatchOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SuperStreamPatch {
		return vs[0].([]*SuperStreamPatch)[vs[1].(int)]
	}).(SuperStreamPatchOutput)
}

type SuperStreamPatchMapOutput struct{ *pulumi.OutputState }

func (SuperStreamPatchMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SuperStreamPatch)(nil)).Elem()
}

func (o SuperStreamPatchMapOutput) ToSuperStreamPatchMapOutput() SuperStreamPatchMapOutput {
	return o
}

func (o SuperStreamPatchMapOutput) ToSuperStreamPatchMapOutputWithContext(ctx context.Context) SuperStreamPatchMapOutput {
	return o
}

func (o SuperStreamPatchMapOutput) MapIndex(k pulumi.StringInput) SuperStreamPatchOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SuperStreamPatch {
		return vs[0].(map[string]*SuperStreamPatch)[vs[1].(string)]
	}).(SuperStreamPatchOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SuperStreamPatchInput)(nil)).Elem(), &SuperStreamPatch{})
	pulumi.RegisterInputType(reflect.TypeOf((*SuperStreamPatchArrayInput)(nil)).Elem(), SuperStreamPatchArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SuperStreamPatchMapInput)(nil)).Elem(), SuperStreamPatchMap{})
	pulumi.RegisterOutputType(SuperStreamPatchOutput{})
	pulumi.RegisterOutputType(SuperStreamPatchArrayOutput{})
	pulumi.RegisterOutputType(SuperStreamPatchMapOutput{})
}
