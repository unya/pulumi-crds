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

// OCIRepository is the Schema for the ocirepositories API
type OCIRepository struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringOutput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaOutput      `pulumi:"metadata"`
	Spec     OCIRepositorySpecOutput      `pulumi:"spec"`
	Status   OCIRepositoryStatusPtrOutput `pulumi:"status"`
}

// NewOCIRepository registers a new resource with the given unique name, arguments, and options.
func NewOCIRepository(ctx *pulumi.Context,
	name string, args *OCIRepositoryArgs, opts ...pulumi.ResourceOption) (*OCIRepository, error) {
	if args == nil {
		args = &OCIRepositoryArgs{}
	}

	args.ApiVersion = pulumi.StringPtr("source.toolkit.fluxcd.io/v1")
	args.Kind = pulumi.StringPtr("OCIRepository")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("kubernetes:source.toolkit.fluxcd.io/v1beta2:OCIRepository"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource OCIRepository
	err := ctx.RegisterResource("kubernetes:source.toolkit.fluxcd.io/v1:OCIRepository", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOCIRepository gets an existing OCIRepository resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOCIRepository(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OCIRepositoryState, opts ...pulumi.ResourceOption) (*OCIRepository, error) {
	var resource OCIRepository
	err := ctx.ReadResource("kubernetes:source.toolkit.fluxcd.io/v1:OCIRepository", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering OCIRepository resources.
type ocirepositoryState struct {
}

type OCIRepositoryState struct {
}

func (OCIRepositoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*ocirepositoryState)(nil)).Elem()
}

type ocirepositoryArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ObjectMeta `pulumi:"metadata"`
	Spec     *OCIRepositorySpec `pulumi:"spec"`
}

// The set of arguments for constructing a OCIRepository resource.
type OCIRepositoryArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPtrInput
	Spec     OCIRepositorySpecPtrInput
}

func (OCIRepositoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ocirepositoryArgs)(nil)).Elem()
}

type OCIRepositoryInput interface {
	pulumi.Input

	ToOCIRepositoryOutput() OCIRepositoryOutput
	ToOCIRepositoryOutputWithContext(ctx context.Context) OCIRepositoryOutput
}

func (*OCIRepository) ElementType() reflect.Type {
	return reflect.TypeOf((**OCIRepository)(nil)).Elem()
}

func (i *OCIRepository) ToOCIRepositoryOutput() OCIRepositoryOutput {
	return i.ToOCIRepositoryOutputWithContext(context.Background())
}

func (i *OCIRepository) ToOCIRepositoryOutputWithContext(ctx context.Context) OCIRepositoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OCIRepositoryOutput)
}

// OCIRepositoryArrayInput is an input type that accepts OCIRepositoryArray and OCIRepositoryArrayOutput values.
// You can construct a concrete instance of `OCIRepositoryArrayInput` via:
//
//	OCIRepositoryArray{ OCIRepositoryArgs{...} }
type OCIRepositoryArrayInput interface {
	pulumi.Input

	ToOCIRepositoryArrayOutput() OCIRepositoryArrayOutput
	ToOCIRepositoryArrayOutputWithContext(context.Context) OCIRepositoryArrayOutput
}

type OCIRepositoryArray []OCIRepositoryInput

func (OCIRepositoryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*OCIRepository)(nil)).Elem()
}

func (i OCIRepositoryArray) ToOCIRepositoryArrayOutput() OCIRepositoryArrayOutput {
	return i.ToOCIRepositoryArrayOutputWithContext(context.Background())
}

func (i OCIRepositoryArray) ToOCIRepositoryArrayOutputWithContext(ctx context.Context) OCIRepositoryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OCIRepositoryArrayOutput)
}

// OCIRepositoryMapInput is an input type that accepts OCIRepositoryMap and OCIRepositoryMapOutput values.
// You can construct a concrete instance of `OCIRepositoryMapInput` via:
//
//	OCIRepositoryMap{ "key": OCIRepositoryArgs{...} }
type OCIRepositoryMapInput interface {
	pulumi.Input

	ToOCIRepositoryMapOutput() OCIRepositoryMapOutput
	ToOCIRepositoryMapOutputWithContext(context.Context) OCIRepositoryMapOutput
}

type OCIRepositoryMap map[string]OCIRepositoryInput

func (OCIRepositoryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*OCIRepository)(nil)).Elem()
}

func (i OCIRepositoryMap) ToOCIRepositoryMapOutput() OCIRepositoryMapOutput {
	return i.ToOCIRepositoryMapOutputWithContext(context.Background())
}

func (i OCIRepositoryMap) ToOCIRepositoryMapOutputWithContext(ctx context.Context) OCIRepositoryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OCIRepositoryMapOutput)
}

type OCIRepositoryOutput struct{ *pulumi.OutputState }

func (OCIRepositoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OCIRepository)(nil)).Elem()
}

func (o OCIRepositoryOutput) ToOCIRepositoryOutput() OCIRepositoryOutput {
	return o
}

func (o OCIRepositoryOutput) ToOCIRepositoryOutputWithContext(ctx context.Context) OCIRepositoryOutput {
	return o
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o OCIRepositoryOutput) ApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *OCIRepository) pulumi.StringOutput { return v.ApiVersion }).(pulumi.StringOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o OCIRepositoryOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *OCIRepository) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
func (o OCIRepositoryOutput) Metadata() metav1.ObjectMetaOutput {
	return o.ApplyT(func(v *OCIRepository) metav1.ObjectMetaOutput { return v.Metadata }).(metav1.ObjectMetaOutput)
}

func (o OCIRepositoryOutput) Spec() OCIRepositorySpecOutput {
	return o.ApplyT(func(v *OCIRepository) OCIRepositorySpecOutput { return v.Spec }).(OCIRepositorySpecOutput)
}

func (o OCIRepositoryOutput) Status() OCIRepositoryStatusPtrOutput {
	return o.ApplyT(func(v *OCIRepository) OCIRepositoryStatusPtrOutput { return v.Status }).(OCIRepositoryStatusPtrOutput)
}

type OCIRepositoryArrayOutput struct{ *pulumi.OutputState }

func (OCIRepositoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*OCIRepository)(nil)).Elem()
}

func (o OCIRepositoryArrayOutput) ToOCIRepositoryArrayOutput() OCIRepositoryArrayOutput {
	return o
}

func (o OCIRepositoryArrayOutput) ToOCIRepositoryArrayOutputWithContext(ctx context.Context) OCIRepositoryArrayOutput {
	return o
}

func (o OCIRepositoryArrayOutput) Index(i pulumi.IntInput) OCIRepositoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *OCIRepository {
		return vs[0].([]*OCIRepository)[vs[1].(int)]
	}).(OCIRepositoryOutput)
}

type OCIRepositoryMapOutput struct{ *pulumi.OutputState }

func (OCIRepositoryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*OCIRepository)(nil)).Elem()
}

func (o OCIRepositoryMapOutput) ToOCIRepositoryMapOutput() OCIRepositoryMapOutput {
	return o
}

func (o OCIRepositoryMapOutput) ToOCIRepositoryMapOutputWithContext(ctx context.Context) OCIRepositoryMapOutput {
	return o
}

func (o OCIRepositoryMapOutput) MapIndex(k pulumi.StringInput) OCIRepositoryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *OCIRepository {
		return vs[0].(map[string]*OCIRepository)[vs[1].(string)]
	}).(OCIRepositoryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*OCIRepositoryInput)(nil)).Elem(), &OCIRepository{})
	pulumi.RegisterInputType(reflect.TypeOf((*OCIRepositoryArrayInput)(nil)).Elem(), OCIRepositoryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*OCIRepositoryMapInput)(nil)).Elem(), OCIRepositoryMap{})
	pulumi.RegisterOutputType(OCIRepositoryOutput{})
	pulumi.RegisterOutputType(OCIRepositoryArrayOutput{})
	pulumi.RegisterOutputType(OCIRepositoryMapOutput{})
}
