// Code generated by crd2pulumi DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"errors"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ClusterImageCatalogList is a list of ClusterImageCatalog
type ClusterImageCatalogList struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringOutput `pulumi:"apiVersion"`
	// List of clusterimagecatalogs. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md
	Items ClusterImageCatalogTypeArrayOutput `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata metav1.ListMetaOutput `pulumi:"metadata"`
}

// NewClusterImageCatalogList registers a new resource with the given unique name, arguments, and options.
func NewClusterImageCatalogList(ctx *pulumi.Context,
	name string, args *ClusterImageCatalogListArgs, opts ...pulumi.ResourceOption) (*ClusterImageCatalogList, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Items == nil {
		return nil, errors.New("invalid value for required argument 'Items'")
	}
	args.ApiVersion = pulumi.StringPtr("postgresql.cnpg.io/v1")
	args.Kind = pulumi.StringPtr("ClusterImageCatalogList")
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ClusterImageCatalogList
	err := ctx.RegisterResource("kubernetes:postgresql.cnpg.io/v1:ClusterImageCatalogList", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetClusterImageCatalogList gets an existing ClusterImageCatalogList resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetClusterImageCatalogList(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClusterImageCatalogListState, opts ...pulumi.ResourceOption) (*ClusterImageCatalogList, error) {
	var resource ClusterImageCatalogList
	err := ctx.ReadResource("kubernetes:postgresql.cnpg.io/v1:ClusterImageCatalogList", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ClusterImageCatalogList resources.
type clusterImageCatalogListState struct {
}

type ClusterImageCatalogListState struct {
}

func (ClusterImageCatalogListState) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterImageCatalogListState)(nil)).Elem()
}

type clusterImageCatalogListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// List of clusterimagecatalogs. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md
	Items []ClusterImageCatalogType `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata *metav1.ListMeta `pulumi:"metadata"`
}

// The set of arguments for constructing a ClusterImageCatalogList resource.
type ClusterImageCatalogListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// List of clusterimagecatalogs. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md
	Items ClusterImageCatalogTypeArrayInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata metav1.ListMetaPtrInput
}

func (ClusterImageCatalogListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterImageCatalogListArgs)(nil)).Elem()
}

type ClusterImageCatalogListInput interface {
	pulumi.Input

	ToClusterImageCatalogListOutput() ClusterImageCatalogListOutput
	ToClusterImageCatalogListOutputWithContext(ctx context.Context) ClusterImageCatalogListOutput
}

func (*ClusterImageCatalogList) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterImageCatalogList)(nil)).Elem()
}

func (i *ClusterImageCatalogList) ToClusterImageCatalogListOutput() ClusterImageCatalogListOutput {
	return i.ToClusterImageCatalogListOutputWithContext(context.Background())
}

func (i *ClusterImageCatalogList) ToClusterImageCatalogListOutputWithContext(ctx context.Context) ClusterImageCatalogListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterImageCatalogListOutput)
}

// ClusterImageCatalogListArrayInput is an input type that accepts ClusterImageCatalogListArray and ClusterImageCatalogListArrayOutput values.
// You can construct a concrete instance of `ClusterImageCatalogListArrayInput` via:
//
//	ClusterImageCatalogListArray{ ClusterImageCatalogListArgs{...} }
type ClusterImageCatalogListArrayInput interface {
	pulumi.Input

	ToClusterImageCatalogListArrayOutput() ClusterImageCatalogListArrayOutput
	ToClusterImageCatalogListArrayOutputWithContext(context.Context) ClusterImageCatalogListArrayOutput
}

type ClusterImageCatalogListArray []ClusterImageCatalogListInput

func (ClusterImageCatalogListArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ClusterImageCatalogList)(nil)).Elem()
}

func (i ClusterImageCatalogListArray) ToClusterImageCatalogListArrayOutput() ClusterImageCatalogListArrayOutput {
	return i.ToClusterImageCatalogListArrayOutputWithContext(context.Background())
}

func (i ClusterImageCatalogListArray) ToClusterImageCatalogListArrayOutputWithContext(ctx context.Context) ClusterImageCatalogListArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterImageCatalogListArrayOutput)
}

// ClusterImageCatalogListMapInput is an input type that accepts ClusterImageCatalogListMap and ClusterImageCatalogListMapOutput values.
// You can construct a concrete instance of `ClusterImageCatalogListMapInput` via:
//
//	ClusterImageCatalogListMap{ "key": ClusterImageCatalogListArgs{...} }
type ClusterImageCatalogListMapInput interface {
	pulumi.Input

	ToClusterImageCatalogListMapOutput() ClusterImageCatalogListMapOutput
	ToClusterImageCatalogListMapOutputWithContext(context.Context) ClusterImageCatalogListMapOutput
}

type ClusterImageCatalogListMap map[string]ClusterImageCatalogListInput

func (ClusterImageCatalogListMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ClusterImageCatalogList)(nil)).Elem()
}

func (i ClusterImageCatalogListMap) ToClusterImageCatalogListMapOutput() ClusterImageCatalogListMapOutput {
	return i.ToClusterImageCatalogListMapOutputWithContext(context.Background())
}

func (i ClusterImageCatalogListMap) ToClusterImageCatalogListMapOutputWithContext(ctx context.Context) ClusterImageCatalogListMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterImageCatalogListMapOutput)
}

type ClusterImageCatalogListOutput struct{ *pulumi.OutputState }

func (ClusterImageCatalogListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterImageCatalogList)(nil)).Elem()
}

func (o ClusterImageCatalogListOutput) ToClusterImageCatalogListOutput() ClusterImageCatalogListOutput {
	return o
}

func (o ClusterImageCatalogListOutput) ToClusterImageCatalogListOutputWithContext(ctx context.Context) ClusterImageCatalogListOutput {
	return o
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o ClusterImageCatalogListOutput) ApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterImageCatalogList) pulumi.StringOutput { return v.ApiVersion }).(pulumi.StringOutput)
}

// List of clusterimagecatalogs. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md
func (o ClusterImageCatalogListOutput) Items() ClusterImageCatalogTypeArrayOutput {
	return o.ApplyT(func(v *ClusterImageCatalogList) ClusterImageCatalogTypeArrayOutput { return v.Items }).(ClusterImageCatalogTypeArrayOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o ClusterImageCatalogListOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterImageCatalogList) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o ClusterImageCatalogListOutput) Metadata() metav1.ListMetaOutput {
	return o.ApplyT(func(v *ClusterImageCatalogList) metav1.ListMetaOutput { return v.Metadata }).(metav1.ListMetaOutput)
}

type ClusterImageCatalogListArrayOutput struct{ *pulumi.OutputState }

func (ClusterImageCatalogListArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ClusterImageCatalogList)(nil)).Elem()
}

func (o ClusterImageCatalogListArrayOutput) ToClusterImageCatalogListArrayOutput() ClusterImageCatalogListArrayOutput {
	return o
}

func (o ClusterImageCatalogListArrayOutput) ToClusterImageCatalogListArrayOutputWithContext(ctx context.Context) ClusterImageCatalogListArrayOutput {
	return o
}

func (o ClusterImageCatalogListArrayOutput) Index(i pulumi.IntInput) ClusterImageCatalogListOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ClusterImageCatalogList {
		return vs[0].([]*ClusterImageCatalogList)[vs[1].(int)]
	}).(ClusterImageCatalogListOutput)
}

type ClusterImageCatalogListMapOutput struct{ *pulumi.OutputState }

func (ClusterImageCatalogListMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ClusterImageCatalogList)(nil)).Elem()
}

func (o ClusterImageCatalogListMapOutput) ToClusterImageCatalogListMapOutput() ClusterImageCatalogListMapOutput {
	return o
}

func (o ClusterImageCatalogListMapOutput) ToClusterImageCatalogListMapOutputWithContext(ctx context.Context) ClusterImageCatalogListMapOutput {
	return o
}

func (o ClusterImageCatalogListMapOutput) MapIndex(k pulumi.StringInput) ClusterImageCatalogListOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ClusterImageCatalogList {
		return vs[0].(map[string]*ClusterImageCatalogList)[vs[1].(string)]
	}).(ClusterImageCatalogListOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterImageCatalogListInput)(nil)).Elem(), &ClusterImageCatalogList{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterImageCatalogListArrayInput)(nil)).Elem(), ClusterImageCatalogListArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterImageCatalogListMapInput)(nil)).Elem(), ClusterImageCatalogListMap{})
	pulumi.RegisterOutputType(ClusterImageCatalogListOutput{})
	pulumi.RegisterOutputType(ClusterImageCatalogListArrayOutput{})
	pulumi.RegisterOutputType(ClusterImageCatalogListMapOutput{})
}
