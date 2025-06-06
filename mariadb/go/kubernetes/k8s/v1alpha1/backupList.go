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

// BackupList is a list of Backup
type BackupList struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringOutput `pulumi:"apiVersion"`
	// List of backups. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md
	Items BackupTypeArrayOutput `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata metav1.ListMetaOutput `pulumi:"metadata"`
}

// NewBackupList registers a new resource with the given unique name, arguments, and options.
func NewBackupList(ctx *pulumi.Context,
	name string, args *BackupListArgs, opts ...pulumi.ResourceOption) (*BackupList, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Items == nil {
		return nil, errors.New("invalid value for required argument 'Items'")
	}
	args.ApiVersion = pulumi.StringPtr("k8s.mariadb.com/v1alpha1")
	args.Kind = pulumi.StringPtr("BackupList")
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource BackupList
	err := ctx.RegisterResource("kubernetes:k8s.mariadb.com/v1alpha1:BackupList", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBackupList gets an existing BackupList resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBackupList(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BackupListState, opts ...pulumi.ResourceOption) (*BackupList, error) {
	var resource BackupList
	err := ctx.ReadResource("kubernetes:k8s.mariadb.com/v1alpha1:BackupList", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BackupList resources.
type backupListState struct {
}

type BackupListState struct {
}

func (BackupListState) ElementType() reflect.Type {
	return reflect.TypeOf((*backupListState)(nil)).Elem()
}

type backupListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// List of backups. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md
	Items []BackupType `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata *metav1.ListMeta `pulumi:"metadata"`
}

// The set of arguments for constructing a BackupList resource.
type BackupListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// List of backups. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md
	Items BackupTypeArrayInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata metav1.ListMetaPtrInput
}

func (BackupListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*backupListArgs)(nil)).Elem()
}

type BackupListInput interface {
	pulumi.Input

	ToBackupListOutput() BackupListOutput
	ToBackupListOutputWithContext(ctx context.Context) BackupListOutput
}

func (*BackupList) ElementType() reflect.Type {
	return reflect.TypeOf((**BackupList)(nil)).Elem()
}

func (i *BackupList) ToBackupListOutput() BackupListOutput {
	return i.ToBackupListOutputWithContext(context.Background())
}

func (i *BackupList) ToBackupListOutputWithContext(ctx context.Context) BackupListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackupListOutput)
}

// BackupListArrayInput is an input type that accepts BackupListArray and BackupListArrayOutput values.
// You can construct a concrete instance of `BackupListArrayInput` via:
//
//	BackupListArray{ BackupListArgs{...} }
type BackupListArrayInput interface {
	pulumi.Input

	ToBackupListArrayOutput() BackupListArrayOutput
	ToBackupListArrayOutputWithContext(context.Context) BackupListArrayOutput
}

type BackupListArray []BackupListInput

func (BackupListArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BackupList)(nil)).Elem()
}

func (i BackupListArray) ToBackupListArrayOutput() BackupListArrayOutput {
	return i.ToBackupListArrayOutputWithContext(context.Background())
}

func (i BackupListArray) ToBackupListArrayOutputWithContext(ctx context.Context) BackupListArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackupListArrayOutput)
}

// BackupListMapInput is an input type that accepts BackupListMap and BackupListMapOutput values.
// You can construct a concrete instance of `BackupListMapInput` via:
//
//	BackupListMap{ "key": BackupListArgs{...} }
type BackupListMapInput interface {
	pulumi.Input

	ToBackupListMapOutput() BackupListMapOutput
	ToBackupListMapOutputWithContext(context.Context) BackupListMapOutput
}

type BackupListMap map[string]BackupListInput

func (BackupListMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BackupList)(nil)).Elem()
}

func (i BackupListMap) ToBackupListMapOutput() BackupListMapOutput {
	return i.ToBackupListMapOutputWithContext(context.Background())
}

func (i BackupListMap) ToBackupListMapOutputWithContext(ctx context.Context) BackupListMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackupListMapOutput)
}

type BackupListOutput struct{ *pulumi.OutputState }

func (BackupListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BackupList)(nil)).Elem()
}

func (o BackupListOutput) ToBackupListOutput() BackupListOutput {
	return o
}

func (o BackupListOutput) ToBackupListOutputWithContext(ctx context.Context) BackupListOutput {
	return o
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o BackupListOutput) ApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *BackupList) pulumi.StringOutput { return v.ApiVersion }).(pulumi.StringOutput)
}

// List of backups. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md
func (o BackupListOutput) Items() BackupTypeArrayOutput {
	return o.ApplyT(func(v *BackupList) BackupTypeArrayOutput { return v.Items }).(BackupTypeArrayOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o BackupListOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *BackupList) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o BackupListOutput) Metadata() metav1.ListMetaOutput {
	return o.ApplyT(func(v *BackupList) metav1.ListMetaOutput { return v.Metadata }).(metav1.ListMetaOutput)
}

type BackupListArrayOutput struct{ *pulumi.OutputState }

func (BackupListArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BackupList)(nil)).Elem()
}

func (o BackupListArrayOutput) ToBackupListArrayOutput() BackupListArrayOutput {
	return o
}

func (o BackupListArrayOutput) ToBackupListArrayOutputWithContext(ctx context.Context) BackupListArrayOutput {
	return o
}

func (o BackupListArrayOutput) Index(i pulumi.IntInput) BackupListOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *BackupList {
		return vs[0].([]*BackupList)[vs[1].(int)]
	}).(BackupListOutput)
}

type BackupListMapOutput struct{ *pulumi.OutputState }

func (BackupListMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BackupList)(nil)).Elem()
}

func (o BackupListMapOutput) ToBackupListMapOutput() BackupListMapOutput {
	return o
}

func (o BackupListMapOutput) ToBackupListMapOutputWithContext(ctx context.Context) BackupListMapOutput {
	return o
}

func (o BackupListMapOutput) MapIndex(k pulumi.StringInput) BackupListOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *BackupList {
		return vs[0].(map[string]*BackupList)[vs[1].(string)]
	}).(BackupListOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BackupListInput)(nil)).Elem(), &BackupList{})
	pulumi.RegisterInputType(reflect.TypeOf((*BackupListArrayInput)(nil)).Elem(), BackupListArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BackupListMapInput)(nil)).Elem(), BackupListMap{})
	pulumi.RegisterOutputType(BackupListOutput{})
	pulumi.RegisterOutputType(BackupListArrayOutput{})
	pulumi.RegisterOutputType(BackupListMapOutput{})
}
