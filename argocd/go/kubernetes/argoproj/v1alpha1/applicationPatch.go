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

// Patch resources are used to modify existing Kubernetes resources by using
// Server-Side Apply updates. The name of the resource must be specified, but all other properties are optional. More than
// one patch may be applied to the same resource, and a random FieldManager name will be used for each Patch resource.
// Conflicts will result in an error by default, but can be forced using the "pulumi.com/patchForce" annotation. See the
// [Server-Side Apply Docs](https://www.pulumi.com/registry/packages/kubernetes/how-to-guides/managing-resources-with-server-side-apply/) for
// additional information about using Server-Side Apply to manage Kubernetes resources with Pulumi.
// Application is a definition of Application resource.
type ApplicationPatch struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata  metav1.ObjectMetaPatchPtrOutput    `pulumi:"metadata"`
	Operation ApplicationOperationPatchPtrOutput `pulumi:"operation"`
	Spec      ApplicationSpecPatchPtrOutput      `pulumi:"spec"`
	Status    ApplicationStatusPatchPtrOutput    `pulumi:"status"`
}

// NewApplicationPatch registers a new resource with the given unique name, arguments, and options.
func NewApplicationPatch(ctx *pulumi.Context,
	name string, args *ApplicationPatchArgs, opts ...pulumi.ResourceOption) (*ApplicationPatch, error) {
	if args == nil {
		args = &ApplicationPatchArgs{}
	}

	args.ApiVersion = pulumi.StringPtr("argoproj.io/v1alpha1")
	args.Kind = pulumi.StringPtr("Application")
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ApplicationPatch
	err := ctx.RegisterResource("kubernetes:argoproj.io/v1alpha1:ApplicationPatch", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApplicationPatch gets an existing ApplicationPatch resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApplicationPatch(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApplicationPatchState, opts ...pulumi.ResourceOption) (*ApplicationPatch, error) {
	var resource ApplicationPatch
	err := ctx.ReadResource("kubernetes:argoproj.io/v1alpha1:ApplicationPatch", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApplicationPatch resources.
type applicationPatchState struct {
}

type ApplicationPatchState struct {
}

func (ApplicationPatchState) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationPatchState)(nil)).Elem()
}

type applicationPatchArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata  *metav1.ObjectMetaPatch    `pulumi:"metadata"`
	Operation *ApplicationOperationPatch `pulumi:"operation"`
	Spec      *ApplicationSpecPatch      `pulumi:"spec"`
}

// The set of arguments for constructing a ApplicationPatch resource.
type ApplicationPatchArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata  metav1.ObjectMetaPatchPtrInput
	Operation ApplicationOperationPatchPtrInput
	Spec      ApplicationSpecPatchPtrInput
}

func (ApplicationPatchArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationPatchArgs)(nil)).Elem()
}

type ApplicationPatchInput interface {
	pulumi.Input

	ToApplicationPatchOutput() ApplicationPatchOutput
	ToApplicationPatchOutputWithContext(ctx context.Context) ApplicationPatchOutput
}

func (*ApplicationPatch) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationPatch)(nil)).Elem()
}

func (i *ApplicationPatch) ToApplicationPatchOutput() ApplicationPatchOutput {
	return i.ToApplicationPatchOutputWithContext(context.Background())
}

func (i *ApplicationPatch) ToApplicationPatchOutputWithContext(ctx context.Context) ApplicationPatchOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationPatchOutput)
}

// ApplicationPatchArrayInput is an input type that accepts ApplicationPatchArray and ApplicationPatchArrayOutput values.
// You can construct a concrete instance of `ApplicationPatchArrayInput` via:
//
//	ApplicationPatchArray{ ApplicationPatchArgs{...} }
type ApplicationPatchArrayInput interface {
	pulumi.Input

	ToApplicationPatchArrayOutput() ApplicationPatchArrayOutput
	ToApplicationPatchArrayOutputWithContext(context.Context) ApplicationPatchArrayOutput
}

type ApplicationPatchArray []ApplicationPatchInput

func (ApplicationPatchArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ApplicationPatch)(nil)).Elem()
}

func (i ApplicationPatchArray) ToApplicationPatchArrayOutput() ApplicationPatchArrayOutput {
	return i.ToApplicationPatchArrayOutputWithContext(context.Background())
}

func (i ApplicationPatchArray) ToApplicationPatchArrayOutputWithContext(ctx context.Context) ApplicationPatchArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationPatchArrayOutput)
}

// ApplicationPatchMapInput is an input type that accepts ApplicationPatchMap and ApplicationPatchMapOutput values.
// You can construct a concrete instance of `ApplicationPatchMapInput` via:
//
//	ApplicationPatchMap{ "key": ApplicationPatchArgs{...} }
type ApplicationPatchMapInput interface {
	pulumi.Input

	ToApplicationPatchMapOutput() ApplicationPatchMapOutput
	ToApplicationPatchMapOutputWithContext(context.Context) ApplicationPatchMapOutput
}

type ApplicationPatchMap map[string]ApplicationPatchInput

func (ApplicationPatchMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ApplicationPatch)(nil)).Elem()
}

func (i ApplicationPatchMap) ToApplicationPatchMapOutput() ApplicationPatchMapOutput {
	return i.ToApplicationPatchMapOutputWithContext(context.Background())
}

func (i ApplicationPatchMap) ToApplicationPatchMapOutputWithContext(ctx context.Context) ApplicationPatchMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationPatchMapOutput)
}

type ApplicationPatchOutput struct{ *pulumi.OutputState }

func (ApplicationPatchOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationPatch)(nil)).Elem()
}

func (o ApplicationPatchOutput) ToApplicationPatchOutput() ApplicationPatchOutput {
	return o
}

func (o ApplicationPatchOutput) ToApplicationPatchOutputWithContext(ctx context.Context) ApplicationPatchOutput {
	return o
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o ApplicationPatchOutput) ApiVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationPatch) pulumi.StringPtrOutput { return v.ApiVersion }).(pulumi.StringPtrOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o ApplicationPatchOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationPatch) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
func (o ApplicationPatchOutput) Metadata() metav1.ObjectMetaPatchPtrOutput {
	return o.ApplyT(func(v *ApplicationPatch) metav1.ObjectMetaPatchPtrOutput { return v.Metadata }).(metav1.ObjectMetaPatchPtrOutput)
}

func (o ApplicationPatchOutput) Operation() ApplicationOperationPatchPtrOutput {
	return o.ApplyT(func(v *ApplicationPatch) ApplicationOperationPatchPtrOutput { return v.Operation }).(ApplicationOperationPatchPtrOutput)
}

func (o ApplicationPatchOutput) Spec() ApplicationSpecPatchPtrOutput {
	return o.ApplyT(func(v *ApplicationPatch) ApplicationSpecPatchPtrOutput { return v.Spec }).(ApplicationSpecPatchPtrOutput)
}

func (o ApplicationPatchOutput) Status() ApplicationStatusPatchPtrOutput {
	return o.ApplyT(func(v *ApplicationPatch) ApplicationStatusPatchPtrOutput { return v.Status }).(ApplicationStatusPatchPtrOutput)
}

type ApplicationPatchArrayOutput struct{ *pulumi.OutputState }

func (ApplicationPatchArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ApplicationPatch)(nil)).Elem()
}

func (o ApplicationPatchArrayOutput) ToApplicationPatchArrayOutput() ApplicationPatchArrayOutput {
	return o
}

func (o ApplicationPatchArrayOutput) ToApplicationPatchArrayOutputWithContext(ctx context.Context) ApplicationPatchArrayOutput {
	return o
}

func (o ApplicationPatchArrayOutput) Index(i pulumi.IntInput) ApplicationPatchOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ApplicationPatch {
		return vs[0].([]*ApplicationPatch)[vs[1].(int)]
	}).(ApplicationPatchOutput)
}

type ApplicationPatchMapOutput struct{ *pulumi.OutputState }

func (ApplicationPatchMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ApplicationPatch)(nil)).Elem()
}

func (o ApplicationPatchMapOutput) ToApplicationPatchMapOutput() ApplicationPatchMapOutput {
	return o
}

func (o ApplicationPatchMapOutput) ToApplicationPatchMapOutputWithContext(ctx context.Context) ApplicationPatchMapOutput {
	return o
}

func (o ApplicationPatchMapOutput) MapIndex(k pulumi.StringInput) ApplicationPatchOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ApplicationPatch {
		return vs[0].(map[string]*ApplicationPatch)[vs[1].(string)]
	}).(ApplicationPatchOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationPatchInput)(nil)).Elem(), &ApplicationPatch{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationPatchArrayInput)(nil)).Elem(), ApplicationPatchArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationPatchMapInput)(nil)).Elem(), ApplicationPatchMap{})
	pulumi.RegisterOutputType(ApplicationPatchOutput{})
	pulumi.RegisterOutputType(ApplicationPatchArrayOutput{})
	pulumi.RegisterOutputType(ApplicationPatchMapOutput{})
}
