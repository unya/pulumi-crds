// Code generated by crd2pulumi DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1alpha3

import (
	"context"
	"reflect"

	"errors"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// BackendTLSPolicyList is a list of BackendTLSPolicy
type BackendTLSPolicyList struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringOutput `pulumi:"apiVersion"`
	// List of backendtlspolicies. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md
	Items BackendTLSPolicyTypeArrayOutput `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata metav1.ListMetaOutput `pulumi:"metadata"`
}

// NewBackendTLSPolicyList registers a new resource with the given unique name, arguments, and options.
func NewBackendTLSPolicyList(ctx *pulumi.Context,
	name string, args *BackendTLSPolicyListArgs, opts ...pulumi.ResourceOption) (*BackendTLSPolicyList, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Items == nil {
		return nil, errors.New("invalid value for required argument 'Items'")
	}
	args.ApiVersion = pulumi.StringPtr("gateway.networking.k8s.io/v1alpha3")
	args.Kind = pulumi.StringPtr("BackendTLSPolicyList")
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource BackendTLSPolicyList
	err := ctx.RegisterResource("kubernetes:gateway.networking.k8s.io/v1alpha3:BackendTLSPolicyList", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBackendTLSPolicyList gets an existing BackendTLSPolicyList resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBackendTLSPolicyList(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BackendTLSPolicyListState, opts ...pulumi.ResourceOption) (*BackendTLSPolicyList, error) {
	var resource BackendTLSPolicyList
	err := ctx.ReadResource("kubernetes:gateway.networking.k8s.io/v1alpha3:BackendTLSPolicyList", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BackendTLSPolicyList resources.
type backendTLSPolicyListState struct {
}

type BackendTLSPolicyListState struct {
}

func (BackendTLSPolicyListState) ElementType() reflect.Type {
	return reflect.TypeOf((*backendTLSPolicyListState)(nil)).Elem()
}

type backendTLSPolicyListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// List of backendtlspolicies. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md
	Items []BackendTLSPolicyType `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata *metav1.ListMeta `pulumi:"metadata"`
}

// The set of arguments for constructing a BackendTLSPolicyList resource.
type BackendTLSPolicyListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// List of backendtlspolicies. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md
	Items BackendTLSPolicyTypeArrayInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata metav1.ListMetaPtrInput
}

func (BackendTLSPolicyListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*backendTLSPolicyListArgs)(nil)).Elem()
}

type BackendTLSPolicyListInput interface {
	pulumi.Input

	ToBackendTLSPolicyListOutput() BackendTLSPolicyListOutput
	ToBackendTLSPolicyListOutputWithContext(ctx context.Context) BackendTLSPolicyListOutput
}

func (*BackendTLSPolicyList) ElementType() reflect.Type {
	return reflect.TypeOf((**BackendTLSPolicyList)(nil)).Elem()
}

func (i *BackendTLSPolicyList) ToBackendTLSPolicyListOutput() BackendTLSPolicyListOutput {
	return i.ToBackendTLSPolicyListOutputWithContext(context.Background())
}

func (i *BackendTLSPolicyList) ToBackendTLSPolicyListOutputWithContext(ctx context.Context) BackendTLSPolicyListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackendTLSPolicyListOutput)
}

// BackendTLSPolicyListArrayInput is an input type that accepts BackendTLSPolicyListArray and BackendTLSPolicyListArrayOutput values.
// You can construct a concrete instance of `BackendTLSPolicyListArrayInput` via:
//
//	BackendTLSPolicyListArray{ BackendTLSPolicyListArgs{...} }
type BackendTLSPolicyListArrayInput interface {
	pulumi.Input

	ToBackendTLSPolicyListArrayOutput() BackendTLSPolicyListArrayOutput
	ToBackendTLSPolicyListArrayOutputWithContext(context.Context) BackendTLSPolicyListArrayOutput
}

type BackendTLSPolicyListArray []BackendTLSPolicyListInput

func (BackendTLSPolicyListArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BackendTLSPolicyList)(nil)).Elem()
}

func (i BackendTLSPolicyListArray) ToBackendTLSPolicyListArrayOutput() BackendTLSPolicyListArrayOutput {
	return i.ToBackendTLSPolicyListArrayOutputWithContext(context.Background())
}

func (i BackendTLSPolicyListArray) ToBackendTLSPolicyListArrayOutputWithContext(ctx context.Context) BackendTLSPolicyListArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackendTLSPolicyListArrayOutput)
}

// BackendTLSPolicyListMapInput is an input type that accepts BackendTLSPolicyListMap and BackendTLSPolicyListMapOutput values.
// You can construct a concrete instance of `BackendTLSPolicyListMapInput` via:
//
//	BackendTLSPolicyListMap{ "key": BackendTLSPolicyListArgs{...} }
type BackendTLSPolicyListMapInput interface {
	pulumi.Input

	ToBackendTLSPolicyListMapOutput() BackendTLSPolicyListMapOutput
	ToBackendTLSPolicyListMapOutputWithContext(context.Context) BackendTLSPolicyListMapOutput
}

type BackendTLSPolicyListMap map[string]BackendTLSPolicyListInput

func (BackendTLSPolicyListMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BackendTLSPolicyList)(nil)).Elem()
}

func (i BackendTLSPolicyListMap) ToBackendTLSPolicyListMapOutput() BackendTLSPolicyListMapOutput {
	return i.ToBackendTLSPolicyListMapOutputWithContext(context.Background())
}

func (i BackendTLSPolicyListMap) ToBackendTLSPolicyListMapOutputWithContext(ctx context.Context) BackendTLSPolicyListMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackendTLSPolicyListMapOutput)
}

type BackendTLSPolicyListOutput struct{ *pulumi.OutputState }

func (BackendTLSPolicyListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BackendTLSPolicyList)(nil)).Elem()
}

func (o BackendTLSPolicyListOutput) ToBackendTLSPolicyListOutput() BackendTLSPolicyListOutput {
	return o
}

func (o BackendTLSPolicyListOutput) ToBackendTLSPolicyListOutputWithContext(ctx context.Context) BackendTLSPolicyListOutput {
	return o
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o BackendTLSPolicyListOutput) ApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *BackendTLSPolicyList) pulumi.StringOutput { return v.ApiVersion }).(pulumi.StringOutput)
}

// List of backendtlspolicies. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md
func (o BackendTLSPolicyListOutput) Items() BackendTLSPolicyTypeArrayOutput {
	return o.ApplyT(func(v *BackendTLSPolicyList) BackendTLSPolicyTypeArrayOutput { return v.Items }).(BackendTLSPolicyTypeArrayOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o BackendTLSPolicyListOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *BackendTLSPolicyList) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o BackendTLSPolicyListOutput) Metadata() metav1.ListMetaOutput {
	return o.ApplyT(func(v *BackendTLSPolicyList) metav1.ListMetaOutput { return v.Metadata }).(metav1.ListMetaOutput)
}

type BackendTLSPolicyListArrayOutput struct{ *pulumi.OutputState }

func (BackendTLSPolicyListArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BackendTLSPolicyList)(nil)).Elem()
}

func (o BackendTLSPolicyListArrayOutput) ToBackendTLSPolicyListArrayOutput() BackendTLSPolicyListArrayOutput {
	return o
}

func (o BackendTLSPolicyListArrayOutput) ToBackendTLSPolicyListArrayOutputWithContext(ctx context.Context) BackendTLSPolicyListArrayOutput {
	return o
}

func (o BackendTLSPolicyListArrayOutput) Index(i pulumi.IntInput) BackendTLSPolicyListOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *BackendTLSPolicyList {
		return vs[0].([]*BackendTLSPolicyList)[vs[1].(int)]
	}).(BackendTLSPolicyListOutput)
}

type BackendTLSPolicyListMapOutput struct{ *pulumi.OutputState }

func (BackendTLSPolicyListMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BackendTLSPolicyList)(nil)).Elem()
}

func (o BackendTLSPolicyListMapOutput) ToBackendTLSPolicyListMapOutput() BackendTLSPolicyListMapOutput {
	return o
}

func (o BackendTLSPolicyListMapOutput) ToBackendTLSPolicyListMapOutputWithContext(ctx context.Context) BackendTLSPolicyListMapOutput {
	return o
}

func (o BackendTLSPolicyListMapOutput) MapIndex(k pulumi.StringInput) BackendTLSPolicyListOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *BackendTLSPolicyList {
		return vs[0].(map[string]*BackendTLSPolicyList)[vs[1].(string)]
	}).(BackendTLSPolicyListOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BackendTLSPolicyListInput)(nil)).Elem(), &BackendTLSPolicyList{})
	pulumi.RegisterInputType(reflect.TypeOf((*BackendTLSPolicyListArrayInput)(nil)).Elem(), BackendTLSPolicyListArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BackendTLSPolicyListMapInput)(nil)).Elem(), BackendTLSPolicyListMap{})
	pulumi.RegisterOutputType(BackendTLSPolicyListOutput{})
	pulumi.RegisterOutputType(BackendTLSPolicyListArrayOutput{})
	pulumi.RegisterOutputType(BackendTLSPolicyListMapOutput{})
}
