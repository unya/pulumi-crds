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

// ShovelList is a list of Shovel
type ShovelList struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringOutput `pulumi:"apiVersion"`
	// List of shovels. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md
	Items ShovelTypeArrayOutput `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata metav1.ListMetaOutput `pulumi:"metadata"`
}

// NewShovelList registers a new resource with the given unique name, arguments, and options.
func NewShovelList(ctx *pulumi.Context,
	name string, args *ShovelListArgs, opts ...pulumi.ResourceOption) (*ShovelList, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Items == nil {
		return nil, errors.New("invalid value for required argument 'Items'")
	}
	args.ApiVersion = pulumi.StringPtr("rabbitmq.com/v1beta1")
	args.Kind = pulumi.StringPtr("ShovelList")
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ShovelList
	err := ctx.RegisterResource("kubernetes:rabbitmq.com/v1beta1:ShovelList", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetShovelList gets an existing ShovelList resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetShovelList(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ShovelListState, opts ...pulumi.ResourceOption) (*ShovelList, error) {
	var resource ShovelList
	err := ctx.ReadResource("kubernetes:rabbitmq.com/v1beta1:ShovelList", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ShovelList resources.
type shovelListState struct {
}

type ShovelListState struct {
}

func (ShovelListState) ElementType() reflect.Type {
	return reflect.TypeOf((*shovelListState)(nil)).Elem()
}

type shovelListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// List of shovels. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md
	Items []ShovelType `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata *metav1.ListMeta `pulumi:"metadata"`
}

// The set of arguments for constructing a ShovelList resource.
type ShovelListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// List of shovels. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md
	Items ShovelTypeArrayInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata metav1.ListMetaPtrInput
}

func (ShovelListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*shovelListArgs)(nil)).Elem()
}

type ShovelListInput interface {
	pulumi.Input

	ToShovelListOutput() ShovelListOutput
	ToShovelListOutputWithContext(ctx context.Context) ShovelListOutput
}

func (*ShovelList) ElementType() reflect.Type {
	return reflect.TypeOf((**ShovelList)(nil)).Elem()
}

func (i *ShovelList) ToShovelListOutput() ShovelListOutput {
	return i.ToShovelListOutputWithContext(context.Background())
}

func (i *ShovelList) ToShovelListOutputWithContext(ctx context.Context) ShovelListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ShovelListOutput)
}

// ShovelListArrayInput is an input type that accepts ShovelListArray and ShovelListArrayOutput values.
// You can construct a concrete instance of `ShovelListArrayInput` via:
//
//	ShovelListArray{ ShovelListArgs{...} }
type ShovelListArrayInput interface {
	pulumi.Input

	ToShovelListArrayOutput() ShovelListArrayOutput
	ToShovelListArrayOutputWithContext(context.Context) ShovelListArrayOutput
}

type ShovelListArray []ShovelListInput

func (ShovelListArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ShovelList)(nil)).Elem()
}

func (i ShovelListArray) ToShovelListArrayOutput() ShovelListArrayOutput {
	return i.ToShovelListArrayOutputWithContext(context.Background())
}

func (i ShovelListArray) ToShovelListArrayOutputWithContext(ctx context.Context) ShovelListArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ShovelListArrayOutput)
}

// ShovelListMapInput is an input type that accepts ShovelListMap and ShovelListMapOutput values.
// You can construct a concrete instance of `ShovelListMapInput` via:
//
//	ShovelListMap{ "key": ShovelListArgs{...} }
type ShovelListMapInput interface {
	pulumi.Input

	ToShovelListMapOutput() ShovelListMapOutput
	ToShovelListMapOutputWithContext(context.Context) ShovelListMapOutput
}

type ShovelListMap map[string]ShovelListInput

func (ShovelListMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ShovelList)(nil)).Elem()
}

func (i ShovelListMap) ToShovelListMapOutput() ShovelListMapOutput {
	return i.ToShovelListMapOutputWithContext(context.Background())
}

func (i ShovelListMap) ToShovelListMapOutputWithContext(ctx context.Context) ShovelListMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ShovelListMapOutput)
}

type ShovelListOutput struct{ *pulumi.OutputState }

func (ShovelListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ShovelList)(nil)).Elem()
}

func (o ShovelListOutput) ToShovelListOutput() ShovelListOutput {
	return o
}

func (o ShovelListOutput) ToShovelListOutputWithContext(ctx context.Context) ShovelListOutput {
	return o
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o ShovelListOutput) ApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *ShovelList) pulumi.StringOutput { return v.ApiVersion }).(pulumi.StringOutput)
}

// List of shovels. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md
func (o ShovelListOutput) Items() ShovelTypeArrayOutput {
	return o.ApplyT(func(v *ShovelList) ShovelTypeArrayOutput { return v.Items }).(ShovelTypeArrayOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o ShovelListOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *ShovelList) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o ShovelListOutput) Metadata() metav1.ListMetaOutput {
	return o.ApplyT(func(v *ShovelList) metav1.ListMetaOutput { return v.Metadata }).(metav1.ListMetaOutput)
}

type ShovelListArrayOutput struct{ *pulumi.OutputState }

func (ShovelListArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ShovelList)(nil)).Elem()
}

func (o ShovelListArrayOutput) ToShovelListArrayOutput() ShovelListArrayOutput {
	return o
}

func (o ShovelListArrayOutput) ToShovelListArrayOutputWithContext(ctx context.Context) ShovelListArrayOutput {
	return o
}

func (o ShovelListArrayOutput) Index(i pulumi.IntInput) ShovelListOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ShovelList {
		return vs[0].([]*ShovelList)[vs[1].(int)]
	}).(ShovelListOutput)
}

type ShovelListMapOutput struct{ *pulumi.OutputState }

func (ShovelListMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ShovelList)(nil)).Elem()
}

func (o ShovelListMapOutput) ToShovelListMapOutput() ShovelListMapOutput {
	return o
}

func (o ShovelListMapOutput) ToShovelListMapOutputWithContext(ctx context.Context) ShovelListMapOutput {
	return o
}

func (o ShovelListMapOutput) MapIndex(k pulumi.StringInput) ShovelListOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ShovelList {
		return vs[0].(map[string]*ShovelList)[vs[1].(string)]
	}).(ShovelListOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ShovelListInput)(nil)).Elem(), &ShovelList{})
	pulumi.RegisterInputType(reflect.TypeOf((*ShovelListArrayInput)(nil)).Elem(), ShovelListArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ShovelListMapInput)(nil)).Elem(), ShovelListMap{})
	pulumi.RegisterOutputType(ShovelListOutput{})
	pulumi.RegisterOutputType(ShovelListArrayOutput{})
	pulumi.RegisterOutputType(ShovelListMapOutput{})
}
