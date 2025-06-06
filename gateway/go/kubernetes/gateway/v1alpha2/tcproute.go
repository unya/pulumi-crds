// Code generated by crd2pulumi DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1alpha2

import (
	"context"
	"reflect"

	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// TCPRoute provides a way to route TCP requests. When combined with a Gateway
// listener, it can be used to forward connections on the port specified by the
// listener to a set of backends specified by the TCPRoute.
type TCPRoute struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringOutput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaOutput `pulumi:"metadata"`
	Spec     TCPRouteSpecOutput      `pulumi:"spec"`
	Status   TCPRouteStatusPtrOutput `pulumi:"status"`
}

// NewTCPRoute registers a new resource with the given unique name, arguments, and options.
func NewTCPRoute(ctx *pulumi.Context,
	name string, args *TCPRouteArgs, opts ...pulumi.ResourceOption) (*TCPRoute, error) {
	if args == nil {
		args = &TCPRouteArgs{}
	}

	args.ApiVersion = pulumi.StringPtr("gateway.networking.k8s.io/v1alpha2")
	args.Kind = pulumi.StringPtr("TCPRoute")
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource TCPRoute
	err := ctx.RegisterResource("kubernetes:gateway.networking.k8s.io/v1alpha2:TCPRoute", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTCPRoute gets an existing TCPRoute resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTCPRoute(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TCPRouteState, opts ...pulumi.ResourceOption) (*TCPRoute, error) {
	var resource TCPRoute
	err := ctx.ReadResource("kubernetes:gateway.networking.k8s.io/v1alpha2:TCPRoute", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TCPRoute resources.
type tcprouteState struct {
}

type TCPRouteState struct {
}

func (TCPRouteState) ElementType() reflect.Type {
	return reflect.TypeOf((*tcprouteState)(nil)).Elem()
}

type tcprouteArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ObjectMeta `pulumi:"metadata"`
	Spec     *TCPRouteSpec      `pulumi:"spec"`
}

// The set of arguments for constructing a TCPRoute resource.
type TCPRouteArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPtrInput
	Spec     TCPRouteSpecPtrInput
}

func (TCPRouteArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tcprouteArgs)(nil)).Elem()
}

type TCPRouteInput interface {
	pulumi.Input

	ToTCPRouteOutput() TCPRouteOutput
	ToTCPRouteOutputWithContext(ctx context.Context) TCPRouteOutput
}

func (*TCPRoute) ElementType() reflect.Type {
	return reflect.TypeOf((**TCPRoute)(nil)).Elem()
}

func (i *TCPRoute) ToTCPRouteOutput() TCPRouteOutput {
	return i.ToTCPRouteOutputWithContext(context.Background())
}

func (i *TCPRoute) ToTCPRouteOutputWithContext(ctx context.Context) TCPRouteOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TCPRouteOutput)
}

// TCPRouteArrayInput is an input type that accepts TCPRouteArray and TCPRouteArrayOutput values.
// You can construct a concrete instance of `TCPRouteArrayInput` via:
//
//	TCPRouteArray{ TCPRouteArgs{...} }
type TCPRouteArrayInput interface {
	pulumi.Input

	ToTCPRouteArrayOutput() TCPRouteArrayOutput
	ToTCPRouteArrayOutputWithContext(context.Context) TCPRouteArrayOutput
}

type TCPRouteArray []TCPRouteInput

func (TCPRouteArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TCPRoute)(nil)).Elem()
}

func (i TCPRouteArray) ToTCPRouteArrayOutput() TCPRouteArrayOutput {
	return i.ToTCPRouteArrayOutputWithContext(context.Background())
}

func (i TCPRouteArray) ToTCPRouteArrayOutputWithContext(ctx context.Context) TCPRouteArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TCPRouteArrayOutput)
}

// TCPRouteMapInput is an input type that accepts TCPRouteMap and TCPRouteMapOutput values.
// You can construct a concrete instance of `TCPRouteMapInput` via:
//
//	TCPRouteMap{ "key": TCPRouteArgs{...} }
type TCPRouteMapInput interface {
	pulumi.Input

	ToTCPRouteMapOutput() TCPRouteMapOutput
	ToTCPRouteMapOutputWithContext(context.Context) TCPRouteMapOutput
}

type TCPRouteMap map[string]TCPRouteInput

func (TCPRouteMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TCPRoute)(nil)).Elem()
}

func (i TCPRouteMap) ToTCPRouteMapOutput() TCPRouteMapOutput {
	return i.ToTCPRouteMapOutputWithContext(context.Background())
}

func (i TCPRouteMap) ToTCPRouteMapOutputWithContext(ctx context.Context) TCPRouteMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TCPRouteMapOutput)
}

type TCPRouteOutput struct{ *pulumi.OutputState }

func (TCPRouteOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TCPRoute)(nil)).Elem()
}

func (o TCPRouteOutput) ToTCPRouteOutput() TCPRouteOutput {
	return o
}

func (o TCPRouteOutput) ToTCPRouteOutputWithContext(ctx context.Context) TCPRouteOutput {
	return o
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o TCPRouteOutput) ApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *TCPRoute) pulumi.StringOutput { return v.ApiVersion }).(pulumi.StringOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o TCPRouteOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *TCPRoute) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
func (o TCPRouteOutput) Metadata() metav1.ObjectMetaOutput {
	return o.ApplyT(func(v *TCPRoute) metav1.ObjectMetaOutput { return v.Metadata }).(metav1.ObjectMetaOutput)
}

func (o TCPRouteOutput) Spec() TCPRouteSpecOutput {
	return o.ApplyT(func(v *TCPRoute) TCPRouteSpecOutput { return v.Spec }).(TCPRouteSpecOutput)
}

func (o TCPRouteOutput) Status() TCPRouteStatusPtrOutput {
	return o.ApplyT(func(v *TCPRoute) TCPRouteStatusPtrOutput { return v.Status }).(TCPRouteStatusPtrOutput)
}

type TCPRouteArrayOutput struct{ *pulumi.OutputState }

func (TCPRouteArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TCPRoute)(nil)).Elem()
}

func (o TCPRouteArrayOutput) ToTCPRouteArrayOutput() TCPRouteArrayOutput {
	return o
}

func (o TCPRouteArrayOutput) ToTCPRouteArrayOutputWithContext(ctx context.Context) TCPRouteArrayOutput {
	return o
}

func (o TCPRouteArrayOutput) Index(i pulumi.IntInput) TCPRouteOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *TCPRoute {
		return vs[0].([]*TCPRoute)[vs[1].(int)]
	}).(TCPRouteOutput)
}

type TCPRouteMapOutput struct{ *pulumi.OutputState }

func (TCPRouteMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TCPRoute)(nil)).Elem()
}

func (o TCPRouteMapOutput) ToTCPRouteMapOutput() TCPRouteMapOutput {
	return o
}

func (o TCPRouteMapOutput) ToTCPRouteMapOutputWithContext(ctx context.Context) TCPRouteMapOutput {
	return o
}

func (o TCPRouteMapOutput) MapIndex(k pulumi.StringInput) TCPRouteOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *TCPRoute {
		return vs[0].(map[string]*TCPRoute)[vs[1].(string)]
	}).(TCPRouteOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TCPRouteInput)(nil)).Elem(), &TCPRoute{})
	pulumi.RegisterInputType(reflect.TypeOf((*TCPRouteArrayInput)(nil)).Elem(), TCPRouteArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TCPRouteMapInput)(nil)).Elem(), TCPRouteMap{})
	pulumi.RegisterOutputType(TCPRouteOutput{})
	pulumi.RegisterOutputType(TCPRouteArrayOutput{})
	pulumi.RegisterOutputType(TCPRouteMapOutput{})
}
