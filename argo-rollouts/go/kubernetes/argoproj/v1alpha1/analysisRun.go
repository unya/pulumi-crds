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

type AnalysisRun struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringOutput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaOutput    `pulumi:"metadata"`
	Spec     AnalysisRunSpecOutput      `pulumi:"spec"`
	Status   AnalysisRunStatusPtrOutput `pulumi:"status"`
}

// NewAnalysisRun registers a new resource with the given unique name, arguments, and options.
func NewAnalysisRun(ctx *pulumi.Context,
	name string, args *AnalysisRunArgs, opts ...pulumi.ResourceOption) (*AnalysisRun, error) {
	if args == nil {
		args = &AnalysisRunArgs{}
	}

	args.ApiVersion = pulumi.StringPtr("argoproj.io/v1alpha1")
	args.Kind = pulumi.StringPtr("AnalysisRun")
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource AnalysisRun
	err := ctx.RegisterResource("kubernetes:argoproj.io/v1alpha1:AnalysisRun", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAnalysisRun gets an existing AnalysisRun resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAnalysisRun(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AnalysisRunState, opts ...pulumi.ResourceOption) (*AnalysisRun, error) {
	var resource AnalysisRun
	err := ctx.ReadResource("kubernetes:argoproj.io/v1alpha1:AnalysisRun", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AnalysisRun resources.
type analysisRunState struct {
}

type AnalysisRunState struct {
}

func (AnalysisRunState) ElementType() reflect.Type {
	return reflect.TypeOf((*analysisRunState)(nil)).Elem()
}

type analysisRunArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ObjectMeta `pulumi:"metadata"`
	Spec     *AnalysisRunSpec   `pulumi:"spec"`
}

// The set of arguments for constructing a AnalysisRun resource.
type AnalysisRunArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPtrInput
	Spec     AnalysisRunSpecPtrInput
}

func (AnalysisRunArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*analysisRunArgs)(nil)).Elem()
}

type AnalysisRunInput interface {
	pulumi.Input

	ToAnalysisRunOutput() AnalysisRunOutput
	ToAnalysisRunOutputWithContext(ctx context.Context) AnalysisRunOutput
}

func (*AnalysisRun) ElementType() reflect.Type {
	return reflect.TypeOf((**AnalysisRun)(nil)).Elem()
}

func (i *AnalysisRun) ToAnalysisRunOutput() AnalysisRunOutput {
	return i.ToAnalysisRunOutputWithContext(context.Background())
}

func (i *AnalysisRun) ToAnalysisRunOutputWithContext(ctx context.Context) AnalysisRunOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AnalysisRunOutput)
}

// AnalysisRunArrayInput is an input type that accepts AnalysisRunArray and AnalysisRunArrayOutput values.
// You can construct a concrete instance of `AnalysisRunArrayInput` via:
//
//	AnalysisRunArray{ AnalysisRunArgs{...} }
type AnalysisRunArrayInput interface {
	pulumi.Input

	ToAnalysisRunArrayOutput() AnalysisRunArrayOutput
	ToAnalysisRunArrayOutputWithContext(context.Context) AnalysisRunArrayOutput
}

type AnalysisRunArray []AnalysisRunInput

func (AnalysisRunArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AnalysisRun)(nil)).Elem()
}

func (i AnalysisRunArray) ToAnalysisRunArrayOutput() AnalysisRunArrayOutput {
	return i.ToAnalysisRunArrayOutputWithContext(context.Background())
}

func (i AnalysisRunArray) ToAnalysisRunArrayOutputWithContext(ctx context.Context) AnalysisRunArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AnalysisRunArrayOutput)
}

// AnalysisRunMapInput is an input type that accepts AnalysisRunMap and AnalysisRunMapOutput values.
// You can construct a concrete instance of `AnalysisRunMapInput` via:
//
//	AnalysisRunMap{ "key": AnalysisRunArgs{...} }
type AnalysisRunMapInput interface {
	pulumi.Input

	ToAnalysisRunMapOutput() AnalysisRunMapOutput
	ToAnalysisRunMapOutputWithContext(context.Context) AnalysisRunMapOutput
}

type AnalysisRunMap map[string]AnalysisRunInput

func (AnalysisRunMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AnalysisRun)(nil)).Elem()
}

func (i AnalysisRunMap) ToAnalysisRunMapOutput() AnalysisRunMapOutput {
	return i.ToAnalysisRunMapOutputWithContext(context.Background())
}

func (i AnalysisRunMap) ToAnalysisRunMapOutputWithContext(ctx context.Context) AnalysisRunMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AnalysisRunMapOutput)
}

type AnalysisRunOutput struct{ *pulumi.OutputState }

func (AnalysisRunOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AnalysisRun)(nil)).Elem()
}

func (o AnalysisRunOutput) ToAnalysisRunOutput() AnalysisRunOutput {
	return o
}

func (o AnalysisRunOutput) ToAnalysisRunOutputWithContext(ctx context.Context) AnalysisRunOutput {
	return o
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o AnalysisRunOutput) ApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *AnalysisRun) pulumi.StringOutput { return v.ApiVersion }).(pulumi.StringOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o AnalysisRunOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *AnalysisRun) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
func (o AnalysisRunOutput) Metadata() metav1.ObjectMetaOutput {
	return o.ApplyT(func(v *AnalysisRun) metav1.ObjectMetaOutput { return v.Metadata }).(metav1.ObjectMetaOutput)
}

func (o AnalysisRunOutput) Spec() AnalysisRunSpecOutput {
	return o.ApplyT(func(v *AnalysisRun) AnalysisRunSpecOutput { return v.Spec }).(AnalysisRunSpecOutput)
}

func (o AnalysisRunOutput) Status() AnalysisRunStatusPtrOutput {
	return o.ApplyT(func(v *AnalysisRun) AnalysisRunStatusPtrOutput { return v.Status }).(AnalysisRunStatusPtrOutput)
}

type AnalysisRunArrayOutput struct{ *pulumi.OutputState }

func (AnalysisRunArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AnalysisRun)(nil)).Elem()
}

func (o AnalysisRunArrayOutput) ToAnalysisRunArrayOutput() AnalysisRunArrayOutput {
	return o
}

func (o AnalysisRunArrayOutput) ToAnalysisRunArrayOutputWithContext(ctx context.Context) AnalysisRunArrayOutput {
	return o
}

func (o AnalysisRunArrayOutput) Index(i pulumi.IntInput) AnalysisRunOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AnalysisRun {
		return vs[0].([]*AnalysisRun)[vs[1].(int)]
	}).(AnalysisRunOutput)
}

type AnalysisRunMapOutput struct{ *pulumi.OutputState }

func (AnalysisRunMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AnalysisRun)(nil)).Elem()
}

func (o AnalysisRunMapOutput) ToAnalysisRunMapOutput() AnalysisRunMapOutput {
	return o
}

func (o AnalysisRunMapOutput) ToAnalysisRunMapOutputWithContext(ctx context.Context) AnalysisRunMapOutput {
	return o
}

func (o AnalysisRunMapOutput) MapIndex(k pulumi.StringInput) AnalysisRunOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AnalysisRun {
		return vs[0].(map[string]*AnalysisRun)[vs[1].(string)]
	}).(AnalysisRunOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AnalysisRunInput)(nil)).Elem(), &AnalysisRun{})
	pulumi.RegisterInputType(reflect.TypeOf((*AnalysisRunArrayInput)(nil)).Elem(), AnalysisRunArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AnalysisRunMapInput)(nil)).Elem(), AnalysisRunMap{})
	pulumi.RegisterOutputType(AnalysisRunOutput{})
	pulumi.RegisterOutputType(AnalysisRunArrayOutput{})
	pulumi.RegisterOutputType(AnalysisRunMapOutput{})
}
