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

// Receiver is the Schema for the receivers API.
type Receiver struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringOutput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaOutput `pulumi:"metadata"`
	Spec     ReceiverSpecOutput      `pulumi:"spec"`
	Status   ReceiverStatusPtrOutput `pulumi:"status"`
}

// NewReceiver registers a new resource with the given unique name, arguments, and options.
func NewReceiver(ctx *pulumi.Context,
	name string, args *ReceiverArgs, opts ...pulumi.ResourceOption) (*Receiver, error) {
	if args == nil {
		args = &ReceiverArgs{}
	}

	args.ApiVersion = pulumi.StringPtr("notification.toolkit.fluxcd.io/v1")
	args.Kind = pulumi.StringPtr("Receiver")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("kubernetes:notification.toolkit.fluxcd.io/v1beta1:Receiver"),
		},
		{
			Type: pulumi.String("kubernetes:notification.toolkit.fluxcd.io/v1beta2:Receiver"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Receiver
	err := ctx.RegisterResource("kubernetes:notification.toolkit.fluxcd.io/v1:Receiver", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetReceiver gets an existing Receiver resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetReceiver(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ReceiverState, opts ...pulumi.ResourceOption) (*Receiver, error) {
	var resource Receiver
	err := ctx.ReadResource("kubernetes:notification.toolkit.fluxcd.io/v1:Receiver", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Receiver resources.
type receiverState struct {
}

type ReceiverState struct {
}

func (ReceiverState) ElementType() reflect.Type {
	return reflect.TypeOf((*receiverState)(nil)).Elem()
}

type receiverArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ObjectMeta `pulumi:"metadata"`
	Spec     *ReceiverSpec      `pulumi:"spec"`
}

// The set of arguments for constructing a Receiver resource.
type ReceiverArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPtrInput
	Spec     ReceiverSpecPtrInput
}

func (ReceiverArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*receiverArgs)(nil)).Elem()
}

type ReceiverInput interface {
	pulumi.Input

	ToReceiverOutput() ReceiverOutput
	ToReceiverOutputWithContext(ctx context.Context) ReceiverOutput
}

func (*Receiver) ElementType() reflect.Type {
	return reflect.TypeOf((**Receiver)(nil)).Elem()
}

func (i *Receiver) ToReceiverOutput() ReceiverOutput {
	return i.ToReceiverOutputWithContext(context.Background())
}

func (i *Receiver) ToReceiverOutputWithContext(ctx context.Context) ReceiverOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReceiverOutput)
}

// ReceiverArrayInput is an input type that accepts ReceiverArray and ReceiverArrayOutput values.
// You can construct a concrete instance of `ReceiverArrayInput` via:
//
//	ReceiverArray{ ReceiverArgs{...} }
type ReceiverArrayInput interface {
	pulumi.Input

	ToReceiverArrayOutput() ReceiverArrayOutput
	ToReceiverArrayOutputWithContext(context.Context) ReceiverArrayOutput
}

type ReceiverArray []ReceiverInput

func (ReceiverArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Receiver)(nil)).Elem()
}

func (i ReceiverArray) ToReceiverArrayOutput() ReceiverArrayOutput {
	return i.ToReceiverArrayOutputWithContext(context.Background())
}

func (i ReceiverArray) ToReceiverArrayOutputWithContext(ctx context.Context) ReceiverArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReceiverArrayOutput)
}

// ReceiverMapInput is an input type that accepts ReceiverMap and ReceiverMapOutput values.
// You can construct a concrete instance of `ReceiverMapInput` via:
//
//	ReceiverMap{ "key": ReceiverArgs{...} }
type ReceiverMapInput interface {
	pulumi.Input

	ToReceiverMapOutput() ReceiverMapOutput
	ToReceiverMapOutputWithContext(context.Context) ReceiverMapOutput
}

type ReceiverMap map[string]ReceiverInput

func (ReceiverMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Receiver)(nil)).Elem()
}

func (i ReceiverMap) ToReceiverMapOutput() ReceiverMapOutput {
	return i.ToReceiverMapOutputWithContext(context.Background())
}

func (i ReceiverMap) ToReceiverMapOutputWithContext(ctx context.Context) ReceiverMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReceiverMapOutput)
}

type ReceiverOutput struct{ *pulumi.OutputState }

func (ReceiverOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Receiver)(nil)).Elem()
}

func (o ReceiverOutput) ToReceiverOutput() ReceiverOutput {
	return o
}

func (o ReceiverOutput) ToReceiverOutputWithContext(ctx context.Context) ReceiverOutput {
	return o
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o ReceiverOutput) ApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *Receiver) pulumi.StringOutput { return v.ApiVersion }).(pulumi.StringOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o ReceiverOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *Receiver) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
func (o ReceiverOutput) Metadata() metav1.ObjectMetaOutput {
	return o.ApplyT(func(v *Receiver) metav1.ObjectMetaOutput { return v.Metadata }).(metav1.ObjectMetaOutput)
}

func (o ReceiverOutput) Spec() ReceiverSpecOutput {
	return o.ApplyT(func(v *Receiver) ReceiverSpecOutput { return v.Spec }).(ReceiverSpecOutput)
}

func (o ReceiverOutput) Status() ReceiverStatusPtrOutput {
	return o.ApplyT(func(v *Receiver) ReceiverStatusPtrOutput { return v.Status }).(ReceiverStatusPtrOutput)
}

type ReceiverArrayOutput struct{ *pulumi.OutputState }

func (ReceiverArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Receiver)(nil)).Elem()
}

func (o ReceiverArrayOutput) ToReceiverArrayOutput() ReceiverArrayOutput {
	return o
}

func (o ReceiverArrayOutput) ToReceiverArrayOutputWithContext(ctx context.Context) ReceiverArrayOutput {
	return o
}

func (o ReceiverArrayOutput) Index(i pulumi.IntInput) ReceiverOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Receiver {
		return vs[0].([]*Receiver)[vs[1].(int)]
	}).(ReceiverOutput)
}

type ReceiverMapOutput struct{ *pulumi.OutputState }

func (ReceiverMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Receiver)(nil)).Elem()
}

func (o ReceiverMapOutput) ToReceiverMapOutput() ReceiverMapOutput {
	return o
}

func (o ReceiverMapOutput) ToReceiverMapOutputWithContext(ctx context.Context) ReceiverMapOutput {
	return o
}

func (o ReceiverMapOutput) MapIndex(k pulumi.StringInput) ReceiverOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Receiver {
		return vs[0].(map[string]*Receiver)[vs[1].(string)]
	}).(ReceiverOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ReceiverInput)(nil)).Elem(), &Receiver{})
	pulumi.RegisterInputType(reflect.TypeOf((*ReceiverArrayInput)(nil)).Elem(), ReceiverArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ReceiverMapInput)(nil)).Elem(), ReceiverMap{})
	pulumi.RegisterOutputType(ReceiverOutput{})
	pulumi.RegisterOutputType(ReceiverArrayOutput{})
	pulumi.RegisterOutputType(ReceiverMapOutput{})
}
