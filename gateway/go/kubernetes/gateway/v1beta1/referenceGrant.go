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

// ReferenceGrant identifies kinds of resources in other namespaces that are
// trusted to reference the specified kinds of resources in the same namespace
// as the policy.
//
// Each ReferenceGrant can be used to represent a unique trust relationship.
// Additional Reference Grants can be used to add to the set of trusted
// sources of inbound references for the namespace they are defined within.
//
// All cross-namespace references in Gateway API (with the exception of cross-namespace
// Gateway-route attachment) require a ReferenceGrant.
//
// ReferenceGrant is a form of runtime verification allowing users to assert
// which cross-namespace object references are permitted. Implementations that
// support ReferenceGrant MUST NOT permit cross-namespace references which have
// no grant, and MUST respond to the removal of a grant by revoking the access
// that the grant allowed.
type ReferenceGrant struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringOutput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaOutput  `pulumi:"metadata"`
	Spec     ReferenceGrantSpecOutput `pulumi:"spec"`
}

// NewReferenceGrant registers a new resource with the given unique name, arguments, and options.
func NewReferenceGrant(ctx *pulumi.Context,
	name string, args *ReferenceGrantArgs, opts ...pulumi.ResourceOption) (*ReferenceGrant, error) {
	if args == nil {
		args = &ReferenceGrantArgs{}
	}

	args.ApiVersion = pulumi.StringPtr("gateway.networking.k8s.io/v1beta1")
	args.Kind = pulumi.StringPtr("ReferenceGrant")
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ReferenceGrant
	err := ctx.RegisterResource("kubernetes:gateway.networking.k8s.io/v1beta1:ReferenceGrant", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetReferenceGrant gets an existing ReferenceGrant resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetReferenceGrant(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ReferenceGrantState, opts ...pulumi.ResourceOption) (*ReferenceGrant, error) {
	var resource ReferenceGrant
	err := ctx.ReadResource("kubernetes:gateway.networking.k8s.io/v1beta1:ReferenceGrant", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ReferenceGrant resources.
type referenceGrantState struct {
}

type ReferenceGrantState struct {
}

func (ReferenceGrantState) ElementType() reflect.Type {
	return reflect.TypeOf((*referenceGrantState)(nil)).Elem()
}

type referenceGrantArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ObjectMeta  `pulumi:"metadata"`
	Spec     *ReferenceGrantSpec `pulumi:"spec"`
}

// The set of arguments for constructing a ReferenceGrant resource.
type ReferenceGrantArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPtrInput
	Spec     ReferenceGrantSpecPtrInput
}

func (ReferenceGrantArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*referenceGrantArgs)(nil)).Elem()
}

type ReferenceGrantInput interface {
	pulumi.Input

	ToReferenceGrantOutput() ReferenceGrantOutput
	ToReferenceGrantOutputWithContext(ctx context.Context) ReferenceGrantOutput
}

func (*ReferenceGrant) ElementType() reflect.Type {
	return reflect.TypeOf((**ReferenceGrant)(nil)).Elem()
}

func (i *ReferenceGrant) ToReferenceGrantOutput() ReferenceGrantOutput {
	return i.ToReferenceGrantOutputWithContext(context.Background())
}

func (i *ReferenceGrant) ToReferenceGrantOutputWithContext(ctx context.Context) ReferenceGrantOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReferenceGrantOutput)
}

// ReferenceGrantArrayInput is an input type that accepts ReferenceGrantArray and ReferenceGrantArrayOutput values.
// You can construct a concrete instance of `ReferenceGrantArrayInput` via:
//
//	ReferenceGrantArray{ ReferenceGrantArgs{...} }
type ReferenceGrantArrayInput interface {
	pulumi.Input

	ToReferenceGrantArrayOutput() ReferenceGrantArrayOutput
	ToReferenceGrantArrayOutputWithContext(context.Context) ReferenceGrantArrayOutput
}

type ReferenceGrantArray []ReferenceGrantInput

func (ReferenceGrantArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ReferenceGrant)(nil)).Elem()
}

func (i ReferenceGrantArray) ToReferenceGrantArrayOutput() ReferenceGrantArrayOutput {
	return i.ToReferenceGrantArrayOutputWithContext(context.Background())
}

func (i ReferenceGrantArray) ToReferenceGrantArrayOutputWithContext(ctx context.Context) ReferenceGrantArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReferenceGrantArrayOutput)
}

// ReferenceGrantMapInput is an input type that accepts ReferenceGrantMap and ReferenceGrantMapOutput values.
// You can construct a concrete instance of `ReferenceGrantMapInput` via:
//
//	ReferenceGrantMap{ "key": ReferenceGrantArgs{...} }
type ReferenceGrantMapInput interface {
	pulumi.Input

	ToReferenceGrantMapOutput() ReferenceGrantMapOutput
	ToReferenceGrantMapOutputWithContext(context.Context) ReferenceGrantMapOutput
}

type ReferenceGrantMap map[string]ReferenceGrantInput

func (ReferenceGrantMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ReferenceGrant)(nil)).Elem()
}

func (i ReferenceGrantMap) ToReferenceGrantMapOutput() ReferenceGrantMapOutput {
	return i.ToReferenceGrantMapOutputWithContext(context.Background())
}

func (i ReferenceGrantMap) ToReferenceGrantMapOutputWithContext(ctx context.Context) ReferenceGrantMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReferenceGrantMapOutput)
}

type ReferenceGrantOutput struct{ *pulumi.OutputState }

func (ReferenceGrantOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReferenceGrant)(nil)).Elem()
}

func (o ReferenceGrantOutput) ToReferenceGrantOutput() ReferenceGrantOutput {
	return o
}

func (o ReferenceGrantOutput) ToReferenceGrantOutputWithContext(ctx context.Context) ReferenceGrantOutput {
	return o
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o ReferenceGrantOutput) ApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *ReferenceGrant) pulumi.StringOutput { return v.ApiVersion }).(pulumi.StringOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o ReferenceGrantOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *ReferenceGrant) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
func (o ReferenceGrantOutput) Metadata() metav1.ObjectMetaOutput {
	return o.ApplyT(func(v *ReferenceGrant) metav1.ObjectMetaOutput { return v.Metadata }).(metav1.ObjectMetaOutput)
}

func (o ReferenceGrantOutput) Spec() ReferenceGrantSpecOutput {
	return o.ApplyT(func(v *ReferenceGrant) ReferenceGrantSpecOutput { return v.Spec }).(ReferenceGrantSpecOutput)
}

type ReferenceGrantArrayOutput struct{ *pulumi.OutputState }

func (ReferenceGrantArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ReferenceGrant)(nil)).Elem()
}

func (o ReferenceGrantArrayOutput) ToReferenceGrantArrayOutput() ReferenceGrantArrayOutput {
	return o
}

func (o ReferenceGrantArrayOutput) ToReferenceGrantArrayOutputWithContext(ctx context.Context) ReferenceGrantArrayOutput {
	return o
}

func (o ReferenceGrantArrayOutput) Index(i pulumi.IntInput) ReferenceGrantOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ReferenceGrant {
		return vs[0].([]*ReferenceGrant)[vs[1].(int)]
	}).(ReferenceGrantOutput)
}

type ReferenceGrantMapOutput struct{ *pulumi.OutputState }

func (ReferenceGrantMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ReferenceGrant)(nil)).Elem()
}

func (o ReferenceGrantMapOutput) ToReferenceGrantMapOutput() ReferenceGrantMapOutput {
	return o
}

func (o ReferenceGrantMapOutput) ToReferenceGrantMapOutputWithContext(ctx context.Context) ReferenceGrantMapOutput {
	return o
}

func (o ReferenceGrantMapOutput) MapIndex(k pulumi.StringInput) ReferenceGrantOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ReferenceGrant {
		return vs[0].(map[string]*ReferenceGrant)[vs[1].(string)]
	}).(ReferenceGrantOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ReferenceGrantInput)(nil)).Elem(), &ReferenceGrant{})
	pulumi.RegisterInputType(reflect.TypeOf((*ReferenceGrantArrayInput)(nil)).Elem(), ReferenceGrantArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ReferenceGrantMapInput)(nil)).Elem(), ReferenceGrantMap{})
	pulumi.RegisterOutputType(ReferenceGrantOutput{})
	pulumi.RegisterOutputType(ReferenceGrantArrayOutput{})
	pulumi.RegisterOutputType(ReferenceGrantMapOutput{})
}
