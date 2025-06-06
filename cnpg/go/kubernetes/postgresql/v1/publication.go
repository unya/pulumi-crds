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

// Publication is the Schema for the publications API
type Publication struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringOutput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaOutput    `pulumi:"metadata"`
	Spec     PublicationSpecOutput      `pulumi:"spec"`
	Status   PublicationStatusPtrOutput `pulumi:"status"`
}

// NewPublication registers a new resource with the given unique name, arguments, and options.
func NewPublication(ctx *pulumi.Context,
	name string, args *PublicationArgs, opts ...pulumi.ResourceOption) (*Publication, error) {
	if args == nil {
		args = &PublicationArgs{}
	}

	args.ApiVersion = pulumi.StringPtr("postgresql.cnpg.io/v1")
	args.Kind = pulumi.StringPtr("Publication")
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Publication
	err := ctx.RegisterResource("kubernetes:postgresql.cnpg.io/v1:Publication", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPublication gets an existing Publication resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPublication(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PublicationState, opts ...pulumi.ResourceOption) (*Publication, error) {
	var resource Publication
	err := ctx.ReadResource("kubernetes:postgresql.cnpg.io/v1:Publication", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Publication resources.
type publicationState struct {
}

type PublicationState struct {
}

func (PublicationState) ElementType() reflect.Type {
	return reflect.TypeOf((*publicationState)(nil)).Elem()
}

type publicationArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ObjectMeta `pulumi:"metadata"`
	Spec     *PublicationSpec   `pulumi:"spec"`
}

// The set of arguments for constructing a Publication resource.
type PublicationArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPtrInput
	Spec     PublicationSpecPtrInput
}

func (PublicationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*publicationArgs)(nil)).Elem()
}

type PublicationInput interface {
	pulumi.Input

	ToPublicationOutput() PublicationOutput
	ToPublicationOutputWithContext(ctx context.Context) PublicationOutput
}

func (*Publication) ElementType() reflect.Type {
	return reflect.TypeOf((**Publication)(nil)).Elem()
}

func (i *Publication) ToPublicationOutput() PublicationOutput {
	return i.ToPublicationOutputWithContext(context.Background())
}

func (i *Publication) ToPublicationOutputWithContext(ctx context.Context) PublicationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PublicationOutput)
}

// PublicationArrayInput is an input type that accepts PublicationArray and PublicationArrayOutput values.
// You can construct a concrete instance of `PublicationArrayInput` via:
//
//	PublicationArray{ PublicationArgs{...} }
type PublicationArrayInput interface {
	pulumi.Input

	ToPublicationArrayOutput() PublicationArrayOutput
	ToPublicationArrayOutputWithContext(context.Context) PublicationArrayOutput
}

type PublicationArray []PublicationInput

func (PublicationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Publication)(nil)).Elem()
}

func (i PublicationArray) ToPublicationArrayOutput() PublicationArrayOutput {
	return i.ToPublicationArrayOutputWithContext(context.Background())
}

func (i PublicationArray) ToPublicationArrayOutputWithContext(ctx context.Context) PublicationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PublicationArrayOutput)
}

// PublicationMapInput is an input type that accepts PublicationMap and PublicationMapOutput values.
// You can construct a concrete instance of `PublicationMapInput` via:
//
//	PublicationMap{ "key": PublicationArgs{...} }
type PublicationMapInput interface {
	pulumi.Input

	ToPublicationMapOutput() PublicationMapOutput
	ToPublicationMapOutputWithContext(context.Context) PublicationMapOutput
}

type PublicationMap map[string]PublicationInput

func (PublicationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Publication)(nil)).Elem()
}

func (i PublicationMap) ToPublicationMapOutput() PublicationMapOutput {
	return i.ToPublicationMapOutputWithContext(context.Background())
}

func (i PublicationMap) ToPublicationMapOutputWithContext(ctx context.Context) PublicationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PublicationMapOutput)
}

type PublicationOutput struct{ *pulumi.OutputState }

func (PublicationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Publication)(nil)).Elem()
}

func (o PublicationOutput) ToPublicationOutput() PublicationOutput {
	return o
}

func (o PublicationOutput) ToPublicationOutputWithContext(ctx context.Context) PublicationOutput {
	return o
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o PublicationOutput) ApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *Publication) pulumi.StringOutput { return v.ApiVersion }).(pulumi.StringOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o PublicationOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *Publication) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
func (o PublicationOutput) Metadata() metav1.ObjectMetaOutput {
	return o.ApplyT(func(v *Publication) metav1.ObjectMetaOutput { return v.Metadata }).(metav1.ObjectMetaOutput)
}

func (o PublicationOutput) Spec() PublicationSpecOutput {
	return o.ApplyT(func(v *Publication) PublicationSpecOutput { return v.Spec }).(PublicationSpecOutput)
}

func (o PublicationOutput) Status() PublicationStatusPtrOutput {
	return o.ApplyT(func(v *Publication) PublicationStatusPtrOutput { return v.Status }).(PublicationStatusPtrOutput)
}

type PublicationArrayOutput struct{ *pulumi.OutputState }

func (PublicationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Publication)(nil)).Elem()
}

func (o PublicationArrayOutput) ToPublicationArrayOutput() PublicationArrayOutput {
	return o
}

func (o PublicationArrayOutput) ToPublicationArrayOutputWithContext(ctx context.Context) PublicationArrayOutput {
	return o
}

func (o PublicationArrayOutput) Index(i pulumi.IntInput) PublicationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Publication {
		return vs[0].([]*Publication)[vs[1].(int)]
	}).(PublicationOutput)
}

type PublicationMapOutput struct{ *pulumi.OutputState }

func (PublicationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Publication)(nil)).Elem()
}

func (o PublicationMapOutput) ToPublicationMapOutput() PublicationMapOutput {
	return o
}

func (o PublicationMapOutput) ToPublicationMapOutputWithContext(ctx context.Context) PublicationMapOutput {
	return o
}

func (o PublicationMapOutput) MapIndex(k pulumi.StringInput) PublicationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Publication {
		return vs[0].(map[string]*Publication)[vs[1].(string)]
	}).(PublicationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PublicationInput)(nil)).Elem(), &Publication{})
	pulumi.RegisterInputType(reflect.TypeOf((*PublicationArrayInput)(nil)).Elem(), PublicationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PublicationMapInput)(nil)).Elem(), PublicationMap{})
	pulumi.RegisterOutputType(PublicationOutput{})
	pulumi.RegisterOutputType(PublicationArrayOutput{})
	pulumi.RegisterOutputType(PublicationMapOutput{})
}
