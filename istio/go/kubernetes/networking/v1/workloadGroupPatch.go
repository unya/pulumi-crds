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
type WorkloadGroupPatch struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPatchPtrOutput `pulumi:"metadata"`
	Spec     WorkloadGroupSpecPatchPtrOutput `pulumi:"spec"`
	Status   pulumi.MapOutput                `pulumi:"status"`
}

// NewWorkloadGroupPatch registers a new resource with the given unique name, arguments, and options.
func NewWorkloadGroupPatch(ctx *pulumi.Context,
	name string, args *WorkloadGroupPatchArgs, opts ...pulumi.ResourceOption) (*WorkloadGroupPatch, error) {
	if args == nil {
		args = &WorkloadGroupPatchArgs{}
	}

	args.ApiVersion = pulumi.StringPtr("networking.istio.io/v1")
	args.Kind = pulumi.StringPtr("WorkloadGroup")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("kubernetes:networking.istio.io/v1alpha3:WorkloadGroupPatch"),
		},
		{
			Type: pulumi.String("kubernetes:networking.istio.io/v1beta1:WorkloadGroupPatch"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource WorkloadGroupPatch
	err := ctx.RegisterResource("kubernetes:networking.istio.io/v1:WorkloadGroupPatch", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWorkloadGroupPatch gets an existing WorkloadGroupPatch resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWorkloadGroupPatch(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkloadGroupPatchState, opts ...pulumi.ResourceOption) (*WorkloadGroupPatch, error) {
	var resource WorkloadGroupPatch
	err := ctx.ReadResource("kubernetes:networking.istio.io/v1:WorkloadGroupPatch", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WorkloadGroupPatch resources.
type workloadGroupPatchState struct {
}

type WorkloadGroupPatchState struct {
}

func (WorkloadGroupPatchState) ElementType() reflect.Type {
	return reflect.TypeOf((*workloadGroupPatchState)(nil)).Elem()
}

type workloadGroupPatchArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ObjectMetaPatch `pulumi:"metadata"`
	Spec     *WorkloadGroupSpecPatch `pulumi:"spec"`
}

// The set of arguments for constructing a WorkloadGroupPatch resource.
type WorkloadGroupPatchArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPatchPtrInput
	Spec     WorkloadGroupSpecPatchPtrInput
}

func (WorkloadGroupPatchArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*workloadGroupPatchArgs)(nil)).Elem()
}

type WorkloadGroupPatchInput interface {
	pulumi.Input

	ToWorkloadGroupPatchOutput() WorkloadGroupPatchOutput
	ToWorkloadGroupPatchOutputWithContext(ctx context.Context) WorkloadGroupPatchOutput
}

func (*WorkloadGroupPatch) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkloadGroupPatch)(nil)).Elem()
}

func (i *WorkloadGroupPatch) ToWorkloadGroupPatchOutput() WorkloadGroupPatchOutput {
	return i.ToWorkloadGroupPatchOutputWithContext(context.Background())
}

func (i *WorkloadGroupPatch) ToWorkloadGroupPatchOutputWithContext(ctx context.Context) WorkloadGroupPatchOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkloadGroupPatchOutput)
}

// WorkloadGroupPatchArrayInput is an input type that accepts WorkloadGroupPatchArray and WorkloadGroupPatchArrayOutput values.
// You can construct a concrete instance of `WorkloadGroupPatchArrayInput` via:
//
//	WorkloadGroupPatchArray{ WorkloadGroupPatchArgs{...} }
type WorkloadGroupPatchArrayInput interface {
	pulumi.Input

	ToWorkloadGroupPatchArrayOutput() WorkloadGroupPatchArrayOutput
	ToWorkloadGroupPatchArrayOutputWithContext(context.Context) WorkloadGroupPatchArrayOutput
}

type WorkloadGroupPatchArray []WorkloadGroupPatchInput

func (WorkloadGroupPatchArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WorkloadGroupPatch)(nil)).Elem()
}

func (i WorkloadGroupPatchArray) ToWorkloadGroupPatchArrayOutput() WorkloadGroupPatchArrayOutput {
	return i.ToWorkloadGroupPatchArrayOutputWithContext(context.Background())
}

func (i WorkloadGroupPatchArray) ToWorkloadGroupPatchArrayOutputWithContext(ctx context.Context) WorkloadGroupPatchArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkloadGroupPatchArrayOutput)
}

// WorkloadGroupPatchMapInput is an input type that accepts WorkloadGroupPatchMap and WorkloadGroupPatchMapOutput values.
// You can construct a concrete instance of `WorkloadGroupPatchMapInput` via:
//
//	WorkloadGroupPatchMap{ "key": WorkloadGroupPatchArgs{...} }
type WorkloadGroupPatchMapInput interface {
	pulumi.Input

	ToWorkloadGroupPatchMapOutput() WorkloadGroupPatchMapOutput
	ToWorkloadGroupPatchMapOutputWithContext(context.Context) WorkloadGroupPatchMapOutput
}

type WorkloadGroupPatchMap map[string]WorkloadGroupPatchInput

func (WorkloadGroupPatchMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WorkloadGroupPatch)(nil)).Elem()
}

func (i WorkloadGroupPatchMap) ToWorkloadGroupPatchMapOutput() WorkloadGroupPatchMapOutput {
	return i.ToWorkloadGroupPatchMapOutputWithContext(context.Background())
}

func (i WorkloadGroupPatchMap) ToWorkloadGroupPatchMapOutputWithContext(ctx context.Context) WorkloadGroupPatchMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkloadGroupPatchMapOutput)
}

type WorkloadGroupPatchOutput struct{ *pulumi.OutputState }

func (WorkloadGroupPatchOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkloadGroupPatch)(nil)).Elem()
}

func (o WorkloadGroupPatchOutput) ToWorkloadGroupPatchOutput() WorkloadGroupPatchOutput {
	return o
}

func (o WorkloadGroupPatchOutput) ToWorkloadGroupPatchOutputWithContext(ctx context.Context) WorkloadGroupPatchOutput {
	return o
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o WorkloadGroupPatchOutput) ApiVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkloadGroupPatch) pulumi.StringPtrOutput { return v.ApiVersion }).(pulumi.StringPtrOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o WorkloadGroupPatchOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkloadGroupPatch) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
func (o WorkloadGroupPatchOutput) Metadata() metav1.ObjectMetaPatchPtrOutput {
	return o.ApplyT(func(v *WorkloadGroupPatch) metav1.ObjectMetaPatchPtrOutput { return v.Metadata }).(metav1.ObjectMetaPatchPtrOutput)
}

func (o WorkloadGroupPatchOutput) Spec() WorkloadGroupSpecPatchPtrOutput {
	return o.ApplyT(func(v *WorkloadGroupPatch) WorkloadGroupSpecPatchPtrOutput { return v.Spec }).(WorkloadGroupSpecPatchPtrOutput)
}

func (o WorkloadGroupPatchOutput) Status() pulumi.MapOutput {
	return o.ApplyT(func(v *WorkloadGroupPatch) pulumi.MapOutput { return v.Status }).(pulumi.MapOutput)
}

type WorkloadGroupPatchArrayOutput struct{ *pulumi.OutputState }

func (WorkloadGroupPatchArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WorkloadGroupPatch)(nil)).Elem()
}

func (o WorkloadGroupPatchArrayOutput) ToWorkloadGroupPatchArrayOutput() WorkloadGroupPatchArrayOutput {
	return o
}

func (o WorkloadGroupPatchArrayOutput) ToWorkloadGroupPatchArrayOutputWithContext(ctx context.Context) WorkloadGroupPatchArrayOutput {
	return o
}

func (o WorkloadGroupPatchArrayOutput) Index(i pulumi.IntInput) WorkloadGroupPatchOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *WorkloadGroupPatch {
		return vs[0].([]*WorkloadGroupPatch)[vs[1].(int)]
	}).(WorkloadGroupPatchOutput)
}

type WorkloadGroupPatchMapOutput struct{ *pulumi.OutputState }

func (WorkloadGroupPatchMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WorkloadGroupPatch)(nil)).Elem()
}

func (o WorkloadGroupPatchMapOutput) ToWorkloadGroupPatchMapOutput() WorkloadGroupPatchMapOutput {
	return o
}

func (o WorkloadGroupPatchMapOutput) ToWorkloadGroupPatchMapOutputWithContext(ctx context.Context) WorkloadGroupPatchMapOutput {
	return o
}

func (o WorkloadGroupPatchMapOutput) MapIndex(k pulumi.StringInput) WorkloadGroupPatchOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *WorkloadGroupPatch {
		return vs[0].(map[string]*WorkloadGroupPatch)[vs[1].(string)]
	}).(WorkloadGroupPatchOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WorkloadGroupPatchInput)(nil)).Elem(), &WorkloadGroupPatch{})
	pulumi.RegisterInputType(reflect.TypeOf((*WorkloadGroupPatchArrayInput)(nil)).Elem(), WorkloadGroupPatchArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*WorkloadGroupPatchMapInput)(nil)).Elem(), WorkloadGroupPatchMap{})
	pulumi.RegisterOutputType(WorkloadGroupPatchOutput{})
	pulumi.RegisterOutputType(WorkloadGroupPatchArrayOutput{})
	pulumi.RegisterOutputType(WorkloadGroupPatchMapOutput{})
}
