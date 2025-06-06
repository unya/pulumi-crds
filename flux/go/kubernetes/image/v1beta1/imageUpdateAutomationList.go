// Code generated by crd2pulumi DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	"errors"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ImageUpdateAutomationList is a list of ImageUpdateAutomation
type ImageUpdateAutomationList struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringOutput `pulumi:"apiVersion"`
	// List of imageupdateautomations. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md
	Items ImageUpdateAutomationTypeArrayOutput `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata metav1.ListMetaOutput `pulumi:"metadata"`
}

// NewImageUpdateAutomationList registers a new resource with the given unique name, arguments, and options.
func NewImageUpdateAutomationList(ctx *pulumi.Context,
	name string, args *ImageUpdateAutomationListArgs, opts ...pulumi.ResourceOption) (*ImageUpdateAutomationList, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Items == nil {
		return nil, errors.New("invalid value for required argument 'Items'")
	}
	args.ApiVersion = pulumi.StringPtr("image.toolkit.fluxcd.io/v1beta1")
	args.Kind = pulumi.StringPtr("ImageUpdateAutomationList")
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ImageUpdateAutomationList
	err := ctx.RegisterResource("kubernetes:image.toolkit.fluxcd.io/v1beta1:ImageUpdateAutomationList", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetImageUpdateAutomationList gets an existing ImageUpdateAutomationList resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetImageUpdateAutomationList(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ImageUpdateAutomationListState, opts ...pulumi.ResourceOption) (*ImageUpdateAutomationList, error) {
	var resource ImageUpdateAutomationList
	err := ctx.ReadResource("kubernetes:image.toolkit.fluxcd.io/v1beta1:ImageUpdateAutomationList", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ImageUpdateAutomationList resources.
type imageUpdateAutomationListState struct {
}

type ImageUpdateAutomationListState struct {
}

func (ImageUpdateAutomationListState) ElementType() reflect.Type {
	return reflect.TypeOf((*imageUpdateAutomationListState)(nil)).Elem()
}

type imageUpdateAutomationListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// List of imageupdateautomations. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md
	Items []ImageUpdateAutomationType `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata *metav1.ListMeta `pulumi:"metadata"`
}

// The set of arguments for constructing a ImageUpdateAutomationList resource.
type ImageUpdateAutomationListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// List of imageupdateautomations. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md
	Items ImageUpdateAutomationTypeArrayInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata metav1.ListMetaPtrInput
}

func (ImageUpdateAutomationListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*imageUpdateAutomationListArgs)(nil)).Elem()
}

type ImageUpdateAutomationListInput interface {
	pulumi.Input

	ToImageUpdateAutomationListOutput() ImageUpdateAutomationListOutput
	ToImageUpdateAutomationListOutputWithContext(ctx context.Context) ImageUpdateAutomationListOutput
}

func (*ImageUpdateAutomationList) ElementType() reflect.Type {
	return reflect.TypeOf((**ImageUpdateAutomationList)(nil)).Elem()
}

func (i *ImageUpdateAutomationList) ToImageUpdateAutomationListOutput() ImageUpdateAutomationListOutput {
	return i.ToImageUpdateAutomationListOutputWithContext(context.Background())
}

func (i *ImageUpdateAutomationList) ToImageUpdateAutomationListOutputWithContext(ctx context.Context) ImageUpdateAutomationListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageUpdateAutomationListOutput)
}

// ImageUpdateAutomationListArrayInput is an input type that accepts ImageUpdateAutomationListArray and ImageUpdateAutomationListArrayOutput values.
// You can construct a concrete instance of `ImageUpdateAutomationListArrayInput` via:
//
//	ImageUpdateAutomationListArray{ ImageUpdateAutomationListArgs{...} }
type ImageUpdateAutomationListArrayInput interface {
	pulumi.Input

	ToImageUpdateAutomationListArrayOutput() ImageUpdateAutomationListArrayOutput
	ToImageUpdateAutomationListArrayOutputWithContext(context.Context) ImageUpdateAutomationListArrayOutput
}

type ImageUpdateAutomationListArray []ImageUpdateAutomationListInput

func (ImageUpdateAutomationListArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ImageUpdateAutomationList)(nil)).Elem()
}

func (i ImageUpdateAutomationListArray) ToImageUpdateAutomationListArrayOutput() ImageUpdateAutomationListArrayOutput {
	return i.ToImageUpdateAutomationListArrayOutputWithContext(context.Background())
}

func (i ImageUpdateAutomationListArray) ToImageUpdateAutomationListArrayOutputWithContext(ctx context.Context) ImageUpdateAutomationListArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageUpdateAutomationListArrayOutput)
}

// ImageUpdateAutomationListMapInput is an input type that accepts ImageUpdateAutomationListMap and ImageUpdateAutomationListMapOutput values.
// You can construct a concrete instance of `ImageUpdateAutomationListMapInput` via:
//
//	ImageUpdateAutomationListMap{ "key": ImageUpdateAutomationListArgs{...} }
type ImageUpdateAutomationListMapInput interface {
	pulumi.Input

	ToImageUpdateAutomationListMapOutput() ImageUpdateAutomationListMapOutput
	ToImageUpdateAutomationListMapOutputWithContext(context.Context) ImageUpdateAutomationListMapOutput
}

type ImageUpdateAutomationListMap map[string]ImageUpdateAutomationListInput

func (ImageUpdateAutomationListMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ImageUpdateAutomationList)(nil)).Elem()
}

func (i ImageUpdateAutomationListMap) ToImageUpdateAutomationListMapOutput() ImageUpdateAutomationListMapOutput {
	return i.ToImageUpdateAutomationListMapOutputWithContext(context.Background())
}

func (i ImageUpdateAutomationListMap) ToImageUpdateAutomationListMapOutputWithContext(ctx context.Context) ImageUpdateAutomationListMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageUpdateAutomationListMapOutput)
}

type ImageUpdateAutomationListOutput struct{ *pulumi.OutputState }

func (ImageUpdateAutomationListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ImageUpdateAutomationList)(nil)).Elem()
}

func (o ImageUpdateAutomationListOutput) ToImageUpdateAutomationListOutput() ImageUpdateAutomationListOutput {
	return o
}

func (o ImageUpdateAutomationListOutput) ToImageUpdateAutomationListOutputWithContext(ctx context.Context) ImageUpdateAutomationListOutput {
	return o
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o ImageUpdateAutomationListOutput) ApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *ImageUpdateAutomationList) pulumi.StringOutput { return v.ApiVersion }).(pulumi.StringOutput)
}

// List of imageupdateautomations. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md
func (o ImageUpdateAutomationListOutput) Items() ImageUpdateAutomationTypeArrayOutput {
	return o.ApplyT(func(v *ImageUpdateAutomationList) ImageUpdateAutomationTypeArrayOutput { return v.Items }).(ImageUpdateAutomationTypeArrayOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o ImageUpdateAutomationListOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *ImageUpdateAutomationList) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o ImageUpdateAutomationListOutput) Metadata() metav1.ListMetaOutput {
	return o.ApplyT(func(v *ImageUpdateAutomationList) metav1.ListMetaOutput { return v.Metadata }).(metav1.ListMetaOutput)
}

type ImageUpdateAutomationListArrayOutput struct{ *pulumi.OutputState }

func (ImageUpdateAutomationListArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ImageUpdateAutomationList)(nil)).Elem()
}

func (o ImageUpdateAutomationListArrayOutput) ToImageUpdateAutomationListArrayOutput() ImageUpdateAutomationListArrayOutput {
	return o
}

func (o ImageUpdateAutomationListArrayOutput) ToImageUpdateAutomationListArrayOutputWithContext(ctx context.Context) ImageUpdateAutomationListArrayOutput {
	return o
}

func (o ImageUpdateAutomationListArrayOutput) Index(i pulumi.IntInput) ImageUpdateAutomationListOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ImageUpdateAutomationList {
		return vs[0].([]*ImageUpdateAutomationList)[vs[1].(int)]
	}).(ImageUpdateAutomationListOutput)
}

type ImageUpdateAutomationListMapOutput struct{ *pulumi.OutputState }

func (ImageUpdateAutomationListMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ImageUpdateAutomationList)(nil)).Elem()
}

func (o ImageUpdateAutomationListMapOutput) ToImageUpdateAutomationListMapOutput() ImageUpdateAutomationListMapOutput {
	return o
}

func (o ImageUpdateAutomationListMapOutput) ToImageUpdateAutomationListMapOutputWithContext(ctx context.Context) ImageUpdateAutomationListMapOutput {
	return o
}

func (o ImageUpdateAutomationListMapOutput) MapIndex(k pulumi.StringInput) ImageUpdateAutomationListOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ImageUpdateAutomationList {
		return vs[0].(map[string]*ImageUpdateAutomationList)[vs[1].(string)]
	}).(ImageUpdateAutomationListOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ImageUpdateAutomationListInput)(nil)).Elem(), &ImageUpdateAutomationList{})
	pulumi.RegisterInputType(reflect.TypeOf((*ImageUpdateAutomationListArrayInput)(nil)).Elem(), ImageUpdateAutomationListArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ImageUpdateAutomationListMapInput)(nil)).Elem(), ImageUpdateAutomationListMap{})
	pulumi.RegisterOutputType(ImageUpdateAutomationListOutput{})
	pulumi.RegisterOutputType(ImageUpdateAutomationListArrayOutput{})
	pulumi.RegisterOutputType(ImageUpdateAutomationListMapOutput{})
}
