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
// ProjectConfig is a resource type that describes the configuration of a
// Project.
type ProjectConfigPatch struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPatchPtrOutput   `pulumi:"metadata"`
	Spec     ProjectConfigSpecPatchPtrOutput   `pulumi:"spec"`
	Status   ProjectConfigStatusPatchPtrOutput `pulumi:"status"`
}

// NewProjectConfigPatch registers a new resource with the given unique name, arguments, and options.
func NewProjectConfigPatch(ctx *pulumi.Context,
	name string, args *ProjectConfigPatchArgs, opts ...pulumi.ResourceOption) (*ProjectConfigPatch, error) {
	if args == nil {
		args = &ProjectConfigPatchArgs{}
	}

	args.ApiVersion = pulumi.StringPtr("kargo.akuity.io/v1alpha1")
	args.Kind = pulumi.StringPtr("ProjectConfig")
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ProjectConfigPatch
	err := ctx.RegisterResource("kubernetes:kargo.akuity.io/v1alpha1:ProjectConfigPatch", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProjectConfigPatch gets an existing ProjectConfigPatch resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProjectConfigPatch(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProjectConfigPatchState, opts ...pulumi.ResourceOption) (*ProjectConfigPatch, error) {
	var resource ProjectConfigPatch
	err := ctx.ReadResource("kubernetes:kargo.akuity.io/v1alpha1:ProjectConfigPatch", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ProjectConfigPatch resources.
type projectConfigPatchState struct {
}

type ProjectConfigPatchState struct {
}

func (ProjectConfigPatchState) ElementType() reflect.Type {
	return reflect.TypeOf((*projectConfigPatchState)(nil)).Elem()
}

type projectConfigPatchArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ObjectMetaPatch `pulumi:"metadata"`
	Spec     *ProjectConfigSpecPatch `pulumi:"spec"`
}

// The set of arguments for constructing a ProjectConfigPatch resource.
type ProjectConfigPatchArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPatchPtrInput
	Spec     ProjectConfigSpecPatchPtrInput
}

func (ProjectConfigPatchArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*projectConfigPatchArgs)(nil)).Elem()
}

type ProjectConfigPatchInput interface {
	pulumi.Input

	ToProjectConfigPatchOutput() ProjectConfigPatchOutput
	ToProjectConfigPatchOutputWithContext(ctx context.Context) ProjectConfigPatchOutput
}

func (*ProjectConfigPatch) ElementType() reflect.Type {
	return reflect.TypeOf((**ProjectConfigPatch)(nil)).Elem()
}

func (i *ProjectConfigPatch) ToProjectConfigPatchOutput() ProjectConfigPatchOutput {
	return i.ToProjectConfigPatchOutputWithContext(context.Background())
}

func (i *ProjectConfigPatch) ToProjectConfigPatchOutputWithContext(ctx context.Context) ProjectConfigPatchOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectConfigPatchOutput)
}

// ProjectConfigPatchArrayInput is an input type that accepts ProjectConfigPatchArray and ProjectConfigPatchArrayOutput values.
// You can construct a concrete instance of `ProjectConfigPatchArrayInput` via:
//
//	ProjectConfigPatchArray{ ProjectConfigPatchArgs{...} }
type ProjectConfigPatchArrayInput interface {
	pulumi.Input

	ToProjectConfigPatchArrayOutput() ProjectConfigPatchArrayOutput
	ToProjectConfigPatchArrayOutputWithContext(context.Context) ProjectConfigPatchArrayOutput
}

type ProjectConfigPatchArray []ProjectConfigPatchInput

func (ProjectConfigPatchArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ProjectConfigPatch)(nil)).Elem()
}

func (i ProjectConfigPatchArray) ToProjectConfigPatchArrayOutput() ProjectConfigPatchArrayOutput {
	return i.ToProjectConfigPatchArrayOutputWithContext(context.Background())
}

func (i ProjectConfigPatchArray) ToProjectConfigPatchArrayOutputWithContext(ctx context.Context) ProjectConfigPatchArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectConfigPatchArrayOutput)
}

// ProjectConfigPatchMapInput is an input type that accepts ProjectConfigPatchMap and ProjectConfigPatchMapOutput values.
// You can construct a concrete instance of `ProjectConfigPatchMapInput` via:
//
//	ProjectConfigPatchMap{ "key": ProjectConfigPatchArgs{...} }
type ProjectConfigPatchMapInput interface {
	pulumi.Input

	ToProjectConfigPatchMapOutput() ProjectConfigPatchMapOutput
	ToProjectConfigPatchMapOutputWithContext(context.Context) ProjectConfigPatchMapOutput
}

type ProjectConfigPatchMap map[string]ProjectConfigPatchInput

func (ProjectConfigPatchMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ProjectConfigPatch)(nil)).Elem()
}

func (i ProjectConfigPatchMap) ToProjectConfigPatchMapOutput() ProjectConfigPatchMapOutput {
	return i.ToProjectConfigPatchMapOutputWithContext(context.Background())
}

func (i ProjectConfigPatchMap) ToProjectConfigPatchMapOutputWithContext(ctx context.Context) ProjectConfigPatchMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectConfigPatchMapOutput)
}

type ProjectConfigPatchOutput struct{ *pulumi.OutputState }

func (ProjectConfigPatchOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProjectConfigPatch)(nil)).Elem()
}

func (o ProjectConfigPatchOutput) ToProjectConfigPatchOutput() ProjectConfigPatchOutput {
	return o
}

func (o ProjectConfigPatchOutput) ToProjectConfigPatchOutputWithContext(ctx context.Context) ProjectConfigPatchOutput {
	return o
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o ProjectConfigPatchOutput) ApiVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProjectConfigPatch) pulumi.StringPtrOutput { return v.ApiVersion }).(pulumi.StringPtrOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o ProjectConfigPatchOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProjectConfigPatch) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
func (o ProjectConfigPatchOutput) Metadata() metav1.ObjectMetaPatchPtrOutput {
	return o.ApplyT(func(v *ProjectConfigPatch) metav1.ObjectMetaPatchPtrOutput { return v.Metadata }).(metav1.ObjectMetaPatchPtrOutput)
}

func (o ProjectConfigPatchOutput) Spec() ProjectConfigSpecPatchPtrOutput {
	return o.ApplyT(func(v *ProjectConfigPatch) ProjectConfigSpecPatchPtrOutput { return v.Spec }).(ProjectConfigSpecPatchPtrOutput)
}

func (o ProjectConfigPatchOutput) Status() ProjectConfigStatusPatchPtrOutput {
	return o.ApplyT(func(v *ProjectConfigPatch) ProjectConfigStatusPatchPtrOutput { return v.Status }).(ProjectConfigStatusPatchPtrOutput)
}

type ProjectConfigPatchArrayOutput struct{ *pulumi.OutputState }

func (ProjectConfigPatchArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ProjectConfigPatch)(nil)).Elem()
}

func (o ProjectConfigPatchArrayOutput) ToProjectConfigPatchArrayOutput() ProjectConfigPatchArrayOutput {
	return o
}

func (o ProjectConfigPatchArrayOutput) ToProjectConfigPatchArrayOutputWithContext(ctx context.Context) ProjectConfigPatchArrayOutput {
	return o
}

func (o ProjectConfigPatchArrayOutput) Index(i pulumi.IntInput) ProjectConfigPatchOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ProjectConfigPatch {
		return vs[0].([]*ProjectConfigPatch)[vs[1].(int)]
	}).(ProjectConfigPatchOutput)
}

type ProjectConfigPatchMapOutput struct{ *pulumi.OutputState }

func (ProjectConfigPatchMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ProjectConfigPatch)(nil)).Elem()
}

func (o ProjectConfigPatchMapOutput) ToProjectConfigPatchMapOutput() ProjectConfigPatchMapOutput {
	return o
}

func (o ProjectConfigPatchMapOutput) ToProjectConfigPatchMapOutputWithContext(ctx context.Context) ProjectConfigPatchMapOutput {
	return o
}

func (o ProjectConfigPatchMapOutput) MapIndex(k pulumi.StringInput) ProjectConfigPatchOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ProjectConfigPatch {
		return vs[0].(map[string]*ProjectConfigPatch)[vs[1].(string)]
	}).(ProjectConfigPatchOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectConfigPatchInput)(nil)).Elem(), &ProjectConfigPatch{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectConfigPatchArrayInput)(nil)).Elem(), ProjectConfigPatchArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectConfigPatchMapInput)(nil)).Elem(), ProjectConfigPatchMap{})
	pulumi.RegisterOutputType(ProjectConfigPatchOutput{})
	pulumi.RegisterOutputType(ProjectConfigPatchArrayOutput{})
	pulumi.RegisterOutputType(ProjectConfigPatchMapOutput{})
}
