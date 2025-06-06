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

type ClusterAnalysisTemplate struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringOutput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaOutput           `pulumi:"metadata"`
	Spec     ClusterAnalysisTemplateSpecOutput `pulumi:"spec"`
}

// NewClusterAnalysisTemplate registers a new resource with the given unique name, arguments, and options.
func NewClusterAnalysisTemplate(ctx *pulumi.Context,
	name string, args *ClusterAnalysisTemplateArgs, opts ...pulumi.ResourceOption) (*ClusterAnalysisTemplate, error) {
	if args == nil {
		args = &ClusterAnalysisTemplateArgs{}
	}

	args.ApiVersion = pulumi.StringPtr("argoproj.io/v1alpha1")
	args.Kind = pulumi.StringPtr("ClusterAnalysisTemplate")
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ClusterAnalysisTemplate
	err := ctx.RegisterResource("kubernetes:argoproj.io/v1alpha1:ClusterAnalysisTemplate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetClusterAnalysisTemplate gets an existing ClusterAnalysisTemplate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetClusterAnalysisTemplate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClusterAnalysisTemplateState, opts ...pulumi.ResourceOption) (*ClusterAnalysisTemplate, error) {
	var resource ClusterAnalysisTemplate
	err := ctx.ReadResource("kubernetes:argoproj.io/v1alpha1:ClusterAnalysisTemplate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ClusterAnalysisTemplate resources.
type clusterAnalysisTemplateState struct {
}

type ClusterAnalysisTemplateState struct {
}

func (ClusterAnalysisTemplateState) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterAnalysisTemplateState)(nil)).Elem()
}

type clusterAnalysisTemplateArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ObjectMeta           `pulumi:"metadata"`
	Spec     *ClusterAnalysisTemplateSpec `pulumi:"spec"`
}

// The set of arguments for constructing a ClusterAnalysisTemplate resource.
type ClusterAnalysisTemplateArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPtrInput
	Spec     ClusterAnalysisTemplateSpecPtrInput
}

func (ClusterAnalysisTemplateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterAnalysisTemplateArgs)(nil)).Elem()
}

type ClusterAnalysisTemplateInput interface {
	pulumi.Input

	ToClusterAnalysisTemplateOutput() ClusterAnalysisTemplateOutput
	ToClusterAnalysisTemplateOutputWithContext(ctx context.Context) ClusterAnalysisTemplateOutput
}

func (*ClusterAnalysisTemplate) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterAnalysisTemplate)(nil)).Elem()
}

func (i *ClusterAnalysisTemplate) ToClusterAnalysisTemplateOutput() ClusterAnalysisTemplateOutput {
	return i.ToClusterAnalysisTemplateOutputWithContext(context.Background())
}

func (i *ClusterAnalysisTemplate) ToClusterAnalysisTemplateOutputWithContext(ctx context.Context) ClusterAnalysisTemplateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterAnalysisTemplateOutput)
}

// ClusterAnalysisTemplateArrayInput is an input type that accepts ClusterAnalysisTemplateArray and ClusterAnalysisTemplateArrayOutput values.
// You can construct a concrete instance of `ClusterAnalysisTemplateArrayInput` via:
//
//	ClusterAnalysisTemplateArray{ ClusterAnalysisTemplateArgs{...} }
type ClusterAnalysisTemplateArrayInput interface {
	pulumi.Input

	ToClusterAnalysisTemplateArrayOutput() ClusterAnalysisTemplateArrayOutput
	ToClusterAnalysisTemplateArrayOutputWithContext(context.Context) ClusterAnalysisTemplateArrayOutput
}

type ClusterAnalysisTemplateArray []ClusterAnalysisTemplateInput

func (ClusterAnalysisTemplateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ClusterAnalysisTemplate)(nil)).Elem()
}

func (i ClusterAnalysisTemplateArray) ToClusterAnalysisTemplateArrayOutput() ClusterAnalysisTemplateArrayOutput {
	return i.ToClusterAnalysisTemplateArrayOutputWithContext(context.Background())
}

func (i ClusterAnalysisTemplateArray) ToClusterAnalysisTemplateArrayOutputWithContext(ctx context.Context) ClusterAnalysisTemplateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterAnalysisTemplateArrayOutput)
}

// ClusterAnalysisTemplateMapInput is an input type that accepts ClusterAnalysisTemplateMap and ClusterAnalysisTemplateMapOutput values.
// You can construct a concrete instance of `ClusterAnalysisTemplateMapInput` via:
//
//	ClusterAnalysisTemplateMap{ "key": ClusterAnalysisTemplateArgs{...} }
type ClusterAnalysisTemplateMapInput interface {
	pulumi.Input

	ToClusterAnalysisTemplateMapOutput() ClusterAnalysisTemplateMapOutput
	ToClusterAnalysisTemplateMapOutputWithContext(context.Context) ClusterAnalysisTemplateMapOutput
}

type ClusterAnalysisTemplateMap map[string]ClusterAnalysisTemplateInput

func (ClusterAnalysisTemplateMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ClusterAnalysisTemplate)(nil)).Elem()
}

func (i ClusterAnalysisTemplateMap) ToClusterAnalysisTemplateMapOutput() ClusterAnalysisTemplateMapOutput {
	return i.ToClusterAnalysisTemplateMapOutputWithContext(context.Background())
}

func (i ClusterAnalysisTemplateMap) ToClusterAnalysisTemplateMapOutputWithContext(ctx context.Context) ClusterAnalysisTemplateMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterAnalysisTemplateMapOutput)
}

type ClusterAnalysisTemplateOutput struct{ *pulumi.OutputState }

func (ClusterAnalysisTemplateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterAnalysisTemplate)(nil)).Elem()
}

func (o ClusterAnalysisTemplateOutput) ToClusterAnalysisTemplateOutput() ClusterAnalysisTemplateOutput {
	return o
}

func (o ClusterAnalysisTemplateOutput) ToClusterAnalysisTemplateOutputWithContext(ctx context.Context) ClusterAnalysisTemplateOutput {
	return o
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o ClusterAnalysisTemplateOutput) ApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterAnalysisTemplate) pulumi.StringOutput { return v.ApiVersion }).(pulumi.StringOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o ClusterAnalysisTemplateOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterAnalysisTemplate) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
func (o ClusterAnalysisTemplateOutput) Metadata() metav1.ObjectMetaOutput {
	return o.ApplyT(func(v *ClusterAnalysisTemplate) metav1.ObjectMetaOutput { return v.Metadata }).(metav1.ObjectMetaOutput)
}

func (o ClusterAnalysisTemplateOutput) Spec() ClusterAnalysisTemplateSpecOutput {
	return o.ApplyT(func(v *ClusterAnalysisTemplate) ClusterAnalysisTemplateSpecOutput { return v.Spec }).(ClusterAnalysisTemplateSpecOutput)
}

type ClusterAnalysisTemplateArrayOutput struct{ *pulumi.OutputState }

func (ClusterAnalysisTemplateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ClusterAnalysisTemplate)(nil)).Elem()
}

func (o ClusterAnalysisTemplateArrayOutput) ToClusterAnalysisTemplateArrayOutput() ClusterAnalysisTemplateArrayOutput {
	return o
}

func (o ClusterAnalysisTemplateArrayOutput) ToClusterAnalysisTemplateArrayOutputWithContext(ctx context.Context) ClusterAnalysisTemplateArrayOutput {
	return o
}

func (o ClusterAnalysisTemplateArrayOutput) Index(i pulumi.IntInput) ClusterAnalysisTemplateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ClusterAnalysisTemplate {
		return vs[0].([]*ClusterAnalysisTemplate)[vs[1].(int)]
	}).(ClusterAnalysisTemplateOutput)
}

type ClusterAnalysisTemplateMapOutput struct{ *pulumi.OutputState }

func (ClusterAnalysisTemplateMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ClusterAnalysisTemplate)(nil)).Elem()
}

func (o ClusterAnalysisTemplateMapOutput) ToClusterAnalysisTemplateMapOutput() ClusterAnalysisTemplateMapOutput {
	return o
}

func (o ClusterAnalysisTemplateMapOutput) ToClusterAnalysisTemplateMapOutputWithContext(ctx context.Context) ClusterAnalysisTemplateMapOutput {
	return o
}

func (o ClusterAnalysisTemplateMapOutput) MapIndex(k pulumi.StringInput) ClusterAnalysisTemplateOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ClusterAnalysisTemplate {
		return vs[0].(map[string]*ClusterAnalysisTemplate)[vs[1].(string)]
	}).(ClusterAnalysisTemplateOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterAnalysisTemplateInput)(nil)).Elem(), &ClusterAnalysisTemplate{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterAnalysisTemplateArrayInput)(nil)).Elem(), ClusterAnalysisTemplateArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterAnalysisTemplateMapInput)(nil)).Elem(), ClusterAnalysisTemplateMap{})
	pulumi.RegisterOutputType(ClusterAnalysisTemplateOutput{})
	pulumi.RegisterOutputType(ClusterAnalysisTemplateArrayOutput{})
	pulumi.RegisterOutputType(ClusterAnalysisTemplateMapOutput{})
}
