// Code generated by crd2pulumi DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1alpha1

import (
	"context"
	"reflect"

	"errors"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// DatabaseList is a list of Database
type DatabaseList struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringOutput `pulumi:"apiVersion"`
	// List of databases. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md
	Items DatabaseTypeArrayOutput `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata metav1.ListMetaOutput `pulumi:"metadata"`
}

// NewDatabaseList registers a new resource with the given unique name, arguments, and options.
func NewDatabaseList(ctx *pulumi.Context,
	name string, args *DatabaseListArgs, opts ...pulumi.ResourceOption) (*DatabaseList, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Items == nil {
		return nil, errors.New("invalid value for required argument 'Items'")
	}
	args.ApiVersion = pulumi.StringPtr("k8s.mariadb.com/v1alpha1")
	args.Kind = pulumi.StringPtr("DatabaseList")
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource DatabaseList
	err := ctx.RegisterResource("kubernetes:k8s.mariadb.com/v1alpha1:DatabaseList", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDatabaseList gets an existing DatabaseList resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDatabaseList(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatabaseListState, opts ...pulumi.ResourceOption) (*DatabaseList, error) {
	var resource DatabaseList
	err := ctx.ReadResource("kubernetes:k8s.mariadb.com/v1alpha1:DatabaseList", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DatabaseList resources.
type databaseListState struct {
}

type DatabaseListState struct {
}

func (DatabaseListState) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseListState)(nil)).Elem()
}

type databaseListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// List of databases. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md
	Items []DatabaseType `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata *metav1.ListMeta `pulumi:"metadata"`
}

// The set of arguments for constructing a DatabaseList resource.
type DatabaseListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// List of databases. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md
	Items DatabaseTypeArrayInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata metav1.ListMetaPtrInput
}

func (DatabaseListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseListArgs)(nil)).Elem()
}

type DatabaseListInput interface {
	pulumi.Input

	ToDatabaseListOutput() DatabaseListOutput
	ToDatabaseListOutputWithContext(ctx context.Context) DatabaseListOutput
}

func (*DatabaseList) ElementType() reflect.Type {
	return reflect.TypeOf((**DatabaseList)(nil)).Elem()
}

func (i *DatabaseList) ToDatabaseListOutput() DatabaseListOutput {
	return i.ToDatabaseListOutputWithContext(context.Background())
}

func (i *DatabaseList) ToDatabaseListOutputWithContext(ctx context.Context) DatabaseListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseListOutput)
}

// DatabaseListArrayInput is an input type that accepts DatabaseListArray and DatabaseListArrayOutput values.
// You can construct a concrete instance of `DatabaseListArrayInput` via:
//
//	DatabaseListArray{ DatabaseListArgs{...} }
type DatabaseListArrayInput interface {
	pulumi.Input

	ToDatabaseListArrayOutput() DatabaseListArrayOutput
	ToDatabaseListArrayOutputWithContext(context.Context) DatabaseListArrayOutput
}

type DatabaseListArray []DatabaseListInput

func (DatabaseListArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DatabaseList)(nil)).Elem()
}

func (i DatabaseListArray) ToDatabaseListArrayOutput() DatabaseListArrayOutput {
	return i.ToDatabaseListArrayOutputWithContext(context.Background())
}

func (i DatabaseListArray) ToDatabaseListArrayOutputWithContext(ctx context.Context) DatabaseListArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseListArrayOutput)
}

// DatabaseListMapInput is an input type that accepts DatabaseListMap and DatabaseListMapOutput values.
// You can construct a concrete instance of `DatabaseListMapInput` via:
//
//	DatabaseListMap{ "key": DatabaseListArgs{...} }
type DatabaseListMapInput interface {
	pulumi.Input

	ToDatabaseListMapOutput() DatabaseListMapOutput
	ToDatabaseListMapOutputWithContext(context.Context) DatabaseListMapOutput
}

type DatabaseListMap map[string]DatabaseListInput

func (DatabaseListMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DatabaseList)(nil)).Elem()
}

func (i DatabaseListMap) ToDatabaseListMapOutput() DatabaseListMapOutput {
	return i.ToDatabaseListMapOutputWithContext(context.Background())
}

func (i DatabaseListMap) ToDatabaseListMapOutputWithContext(ctx context.Context) DatabaseListMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseListMapOutput)
}

type DatabaseListOutput struct{ *pulumi.OutputState }

func (DatabaseListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DatabaseList)(nil)).Elem()
}

func (o DatabaseListOutput) ToDatabaseListOutput() DatabaseListOutput {
	return o
}

func (o DatabaseListOutput) ToDatabaseListOutputWithContext(ctx context.Context) DatabaseListOutput {
	return o
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o DatabaseListOutput) ApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseList) pulumi.StringOutput { return v.ApiVersion }).(pulumi.StringOutput)
}

// List of databases. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md
func (o DatabaseListOutput) Items() DatabaseTypeArrayOutput {
	return o.ApplyT(func(v *DatabaseList) DatabaseTypeArrayOutput { return v.Items }).(DatabaseTypeArrayOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o DatabaseListOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseList) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o DatabaseListOutput) Metadata() metav1.ListMetaOutput {
	return o.ApplyT(func(v *DatabaseList) metav1.ListMetaOutput { return v.Metadata }).(metav1.ListMetaOutput)
}

type DatabaseListArrayOutput struct{ *pulumi.OutputState }

func (DatabaseListArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DatabaseList)(nil)).Elem()
}

func (o DatabaseListArrayOutput) ToDatabaseListArrayOutput() DatabaseListArrayOutput {
	return o
}

func (o DatabaseListArrayOutput) ToDatabaseListArrayOutputWithContext(ctx context.Context) DatabaseListArrayOutput {
	return o
}

func (o DatabaseListArrayOutput) Index(i pulumi.IntInput) DatabaseListOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DatabaseList {
		return vs[0].([]*DatabaseList)[vs[1].(int)]
	}).(DatabaseListOutput)
}

type DatabaseListMapOutput struct{ *pulumi.OutputState }

func (DatabaseListMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DatabaseList)(nil)).Elem()
}

func (o DatabaseListMapOutput) ToDatabaseListMapOutput() DatabaseListMapOutput {
	return o
}

func (o DatabaseListMapOutput) ToDatabaseListMapOutputWithContext(ctx context.Context) DatabaseListMapOutput {
	return o
}

func (o DatabaseListMapOutput) MapIndex(k pulumi.StringInput) DatabaseListOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DatabaseList {
		return vs[0].(map[string]*DatabaseList)[vs[1].(string)]
	}).(DatabaseListOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DatabaseListInput)(nil)).Elem(), &DatabaseList{})
	pulumi.RegisterInputType(reflect.TypeOf((*DatabaseListArrayInput)(nil)).Elem(), DatabaseListArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DatabaseListMapInput)(nil)).Elem(), DatabaseListMap{})
	pulumi.RegisterOutputType(DatabaseListOutput{})
	pulumi.RegisterOutputType(DatabaseListArrayOutput{})
	pulumi.RegisterOutputType(DatabaseListMapOutput{})
}
