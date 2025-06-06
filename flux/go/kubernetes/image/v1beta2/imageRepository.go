// Code generated by crd2pulumi DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta2

import (
	"context"
	"reflect"

	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ImageRepository is the Schema for the imagerepositories API
type ImageRepository struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringOutput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaOutput        `pulumi:"metadata"`
	Spec     ImageRepositorySpecOutput      `pulumi:"spec"`
	Status   ImageRepositoryStatusPtrOutput `pulumi:"status"`
}

// NewImageRepository registers a new resource with the given unique name, arguments, and options.
func NewImageRepository(ctx *pulumi.Context,
	name string, args *ImageRepositoryArgs, opts ...pulumi.ResourceOption) (*ImageRepository, error) {
	if args == nil {
		args = &ImageRepositoryArgs{}
	}

	args.ApiVersion = pulumi.StringPtr("image.toolkit.fluxcd.io/v1beta2")
	args.Kind = pulumi.StringPtr("ImageRepository")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("kubernetes:image.toolkit.fluxcd.io/v1beta1:ImageRepository"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ImageRepository
	err := ctx.RegisterResource("kubernetes:image.toolkit.fluxcd.io/v1beta2:ImageRepository", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetImageRepository gets an existing ImageRepository resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetImageRepository(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ImageRepositoryState, opts ...pulumi.ResourceOption) (*ImageRepository, error) {
	var resource ImageRepository
	err := ctx.ReadResource("kubernetes:image.toolkit.fluxcd.io/v1beta2:ImageRepository", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ImageRepository resources.
type imageRepositoryState struct {
}

type ImageRepositoryState struct {
}

func (ImageRepositoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*imageRepositoryState)(nil)).Elem()
}

type imageRepositoryArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ObjectMeta   `pulumi:"metadata"`
	Spec     *ImageRepositorySpec `pulumi:"spec"`
}

// The set of arguments for constructing a ImageRepository resource.
type ImageRepositoryArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPtrInput
	Spec     ImageRepositorySpecPtrInput
}

func (ImageRepositoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*imageRepositoryArgs)(nil)).Elem()
}

type ImageRepositoryInput interface {
	pulumi.Input

	ToImageRepositoryOutput() ImageRepositoryOutput
	ToImageRepositoryOutputWithContext(ctx context.Context) ImageRepositoryOutput
}

func (*ImageRepository) ElementType() reflect.Type {
	return reflect.TypeOf((**ImageRepository)(nil)).Elem()
}

func (i *ImageRepository) ToImageRepositoryOutput() ImageRepositoryOutput {
	return i.ToImageRepositoryOutputWithContext(context.Background())
}

func (i *ImageRepository) ToImageRepositoryOutputWithContext(ctx context.Context) ImageRepositoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageRepositoryOutput)
}

// ImageRepositoryArrayInput is an input type that accepts ImageRepositoryArray and ImageRepositoryArrayOutput values.
// You can construct a concrete instance of `ImageRepositoryArrayInput` via:
//
//	ImageRepositoryArray{ ImageRepositoryArgs{...} }
type ImageRepositoryArrayInput interface {
	pulumi.Input

	ToImageRepositoryArrayOutput() ImageRepositoryArrayOutput
	ToImageRepositoryArrayOutputWithContext(context.Context) ImageRepositoryArrayOutput
}

type ImageRepositoryArray []ImageRepositoryInput

func (ImageRepositoryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ImageRepository)(nil)).Elem()
}

func (i ImageRepositoryArray) ToImageRepositoryArrayOutput() ImageRepositoryArrayOutput {
	return i.ToImageRepositoryArrayOutputWithContext(context.Background())
}

func (i ImageRepositoryArray) ToImageRepositoryArrayOutputWithContext(ctx context.Context) ImageRepositoryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageRepositoryArrayOutput)
}

// ImageRepositoryMapInput is an input type that accepts ImageRepositoryMap and ImageRepositoryMapOutput values.
// You can construct a concrete instance of `ImageRepositoryMapInput` via:
//
//	ImageRepositoryMap{ "key": ImageRepositoryArgs{...} }
type ImageRepositoryMapInput interface {
	pulumi.Input

	ToImageRepositoryMapOutput() ImageRepositoryMapOutput
	ToImageRepositoryMapOutputWithContext(context.Context) ImageRepositoryMapOutput
}

type ImageRepositoryMap map[string]ImageRepositoryInput

func (ImageRepositoryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ImageRepository)(nil)).Elem()
}

func (i ImageRepositoryMap) ToImageRepositoryMapOutput() ImageRepositoryMapOutput {
	return i.ToImageRepositoryMapOutputWithContext(context.Background())
}

func (i ImageRepositoryMap) ToImageRepositoryMapOutputWithContext(ctx context.Context) ImageRepositoryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageRepositoryMapOutput)
}

type ImageRepositoryOutput struct{ *pulumi.OutputState }

func (ImageRepositoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ImageRepository)(nil)).Elem()
}

func (o ImageRepositoryOutput) ToImageRepositoryOutput() ImageRepositoryOutput {
	return o
}

func (o ImageRepositoryOutput) ToImageRepositoryOutputWithContext(ctx context.Context) ImageRepositoryOutput {
	return o
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o ImageRepositoryOutput) ApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *ImageRepository) pulumi.StringOutput { return v.ApiVersion }).(pulumi.StringOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o ImageRepositoryOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *ImageRepository) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
func (o ImageRepositoryOutput) Metadata() metav1.ObjectMetaOutput {
	return o.ApplyT(func(v *ImageRepository) metav1.ObjectMetaOutput { return v.Metadata }).(metav1.ObjectMetaOutput)
}

func (o ImageRepositoryOutput) Spec() ImageRepositorySpecOutput {
	return o.ApplyT(func(v *ImageRepository) ImageRepositorySpecOutput { return v.Spec }).(ImageRepositorySpecOutput)
}

func (o ImageRepositoryOutput) Status() ImageRepositoryStatusPtrOutput {
	return o.ApplyT(func(v *ImageRepository) ImageRepositoryStatusPtrOutput { return v.Status }).(ImageRepositoryStatusPtrOutput)
}

type ImageRepositoryArrayOutput struct{ *pulumi.OutputState }

func (ImageRepositoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ImageRepository)(nil)).Elem()
}

func (o ImageRepositoryArrayOutput) ToImageRepositoryArrayOutput() ImageRepositoryArrayOutput {
	return o
}

func (o ImageRepositoryArrayOutput) ToImageRepositoryArrayOutputWithContext(ctx context.Context) ImageRepositoryArrayOutput {
	return o
}

func (o ImageRepositoryArrayOutput) Index(i pulumi.IntInput) ImageRepositoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ImageRepository {
		return vs[0].([]*ImageRepository)[vs[1].(int)]
	}).(ImageRepositoryOutput)
}

type ImageRepositoryMapOutput struct{ *pulumi.OutputState }

func (ImageRepositoryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ImageRepository)(nil)).Elem()
}

func (o ImageRepositoryMapOutput) ToImageRepositoryMapOutput() ImageRepositoryMapOutput {
	return o
}

func (o ImageRepositoryMapOutput) ToImageRepositoryMapOutputWithContext(ctx context.Context) ImageRepositoryMapOutput {
	return o
}

func (o ImageRepositoryMapOutput) MapIndex(k pulumi.StringInput) ImageRepositoryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ImageRepository {
		return vs[0].(map[string]*ImageRepository)[vs[1].(string)]
	}).(ImageRepositoryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ImageRepositoryInput)(nil)).Elem(), &ImageRepository{})
	pulumi.RegisterInputType(reflect.TypeOf((*ImageRepositoryArrayInput)(nil)).Elem(), ImageRepositoryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ImageRepositoryMapInput)(nil)).Elem(), ImageRepositoryMap{})
	pulumi.RegisterOutputType(ImageRepositoryOutput{})
	pulumi.RegisterOutputType(ImageRepositoryArrayOutput{})
	pulumi.RegisterOutputType(ImageRepositoryMapOutput{})
}
