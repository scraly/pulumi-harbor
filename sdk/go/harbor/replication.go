// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package harbor

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
	"github.com/pulumiverse/pulumi-harbor/sdk/v3/go/harbor/internal"
)

// ## Example Usage
//
// ## Import
//
// Harbor project can be imported using the `replication id` eg, `<break><break>```sh<break> $ pulumi import harbor:index/replication:Replication main /replication/policies/1 <break>```<break><break>`
type Replication struct {
	pulumi.CustomResourceState

	Action               pulumi.StringOutput          `pulumi:"action"`
	Deletion             pulumi.BoolPtrOutput         `pulumi:"deletion"`
	Description          pulumi.StringPtrOutput       `pulumi:"description"`
	DestNamespace        pulumi.StringPtrOutput       `pulumi:"destNamespace"`
	DestNamespaceReplace pulumi.IntPtrOutput          `pulumi:"destNamespaceReplace"`
	Enabled              pulumi.BoolPtrOutput         `pulumi:"enabled"`
	Filters              ReplicationFilterArrayOutput `pulumi:"filters"`
	Name                 pulumi.StringOutput          `pulumi:"name"`
	Override             pulumi.BoolPtrOutput         `pulumi:"override"`
	RegistryId           pulumi.IntOutput             `pulumi:"registryId"`
	ReplicationPolicyId  pulumi.IntOutput             `pulumi:"replicationPolicyId"`
	Schedule             pulumi.StringPtrOutput       `pulumi:"schedule"`
	Speed                pulumi.IntPtrOutput          `pulumi:"speed"`
}

// NewReplication registers a new resource with the given unique name, arguments, and options.
func NewReplication(ctx *pulumi.Context,
	name string, args *ReplicationArgs, opts ...pulumi.ResourceOption) (*Replication, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Action == nil {
		return nil, errors.New("invalid value for required argument 'Action'")
	}
	if args.RegistryId == nil {
		return nil, errors.New("invalid value for required argument 'RegistryId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Replication
	err := ctx.RegisterResource("harbor:index/replication:Replication", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetReplication gets an existing Replication resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetReplication(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ReplicationState, opts ...pulumi.ResourceOption) (*Replication, error) {
	var resource Replication
	err := ctx.ReadResource("harbor:index/replication:Replication", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Replication resources.
type replicationState struct {
	Action               *string             `pulumi:"action"`
	Deletion             *bool               `pulumi:"deletion"`
	Description          *string             `pulumi:"description"`
	DestNamespace        *string             `pulumi:"destNamespace"`
	DestNamespaceReplace *int                `pulumi:"destNamespaceReplace"`
	Enabled              *bool               `pulumi:"enabled"`
	Filters              []ReplicationFilter `pulumi:"filters"`
	Name                 *string             `pulumi:"name"`
	Override             *bool               `pulumi:"override"`
	RegistryId           *int                `pulumi:"registryId"`
	ReplicationPolicyId  *int                `pulumi:"replicationPolicyId"`
	Schedule             *string             `pulumi:"schedule"`
	Speed                *int                `pulumi:"speed"`
}

type ReplicationState struct {
	Action               pulumi.StringPtrInput
	Deletion             pulumi.BoolPtrInput
	Description          pulumi.StringPtrInput
	DestNamespace        pulumi.StringPtrInput
	DestNamespaceReplace pulumi.IntPtrInput
	Enabled              pulumi.BoolPtrInput
	Filters              ReplicationFilterArrayInput
	Name                 pulumi.StringPtrInput
	Override             pulumi.BoolPtrInput
	RegistryId           pulumi.IntPtrInput
	ReplicationPolicyId  pulumi.IntPtrInput
	Schedule             pulumi.StringPtrInput
	Speed                pulumi.IntPtrInput
}

func (ReplicationState) ElementType() reflect.Type {
	return reflect.TypeOf((*replicationState)(nil)).Elem()
}

type replicationArgs struct {
	Action               string              `pulumi:"action"`
	Deletion             *bool               `pulumi:"deletion"`
	Description          *string             `pulumi:"description"`
	DestNamespace        *string             `pulumi:"destNamespace"`
	DestNamespaceReplace *int                `pulumi:"destNamespaceReplace"`
	Enabled              *bool               `pulumi:"enabled"`
	Filters              []ReplicationFilter `pulumi:"filters"`
	Name                 *string             `pulumi:"name"`
	Override             *bool               `pulumi:"override"`
	RegistryId           int                 `pulumi:"registryId"`
	Schedule             *string             `pulumi:"schedule"`
	Speed                *int                `pulumi:"speed"`
}

// The set of arguments for constructing a Replication resource.
type ReplicationArgs struct {
	Action               pulumi.StringInput
	Deletion             pulumi.BoolPtrInput
	Description          pulumi.StringPtrInput
	DestNamespace        pulumi.StringPtrInput
	DestNamespaceReplace pulumi.IntPtrInput
	Enabled              pulumi.BoolPtrInput
	Filters              ReplicationFilterArrayInput
	Name                 pulumi.StringPtrInput
	Override             pulumi.BoolPtrInput
	RegistryId           pulumi.IntInput
	Schedule             pulumi.StringPtrInput
	Speed                pulumi.IntPtrInput
}

func (ReplicationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*replicationArgs)(nil)).Elem()
}

type ReplicationInput interface {
	pulumi.Input

	ToReplicationOutput() ReplicationOutput
	ToReplicationOutputWithContext(ctx context.Context) ReplicationOutput
}

func (*Replication) ElementType() reflect.Type {
	return reflect.TypeOf((**Replication)(nil)).Elem()
}

func (i *Replication) ToReplicationOutput() ReplicationOutput {
	return i.ToReplicationOutputWithContext(context.Background())
}

func (i *Replication) ToReplicationOutputWithContext(ctx context.Context) ReplicationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReplicationOutput)
}

func (i *Replication) ToOutput(ctx context.Context) pulumix.Output[*Replication] {
	return pulumix.Output[*Replication]{
		OutputState: i.ToReplicationOutputWithContext(ctx).OutputState,
	}
}

// ReplicationArrayInput is an input type that accepts ReplicationArray and ReplicationArrayOutput values.
// You can construct a concrete instance of `ReplicationArrayInput` via:
//
//	ReplicationArray{ ReplicationArgs{...} }
type ReplicationArrayInput interface {
	pulumi.Input

	ToReplicationArrayOutput() ReplicationArrayOutput
	ToReplicationArrayOutputWithContext(context.Context) ReplicationArrayOutput
}

type ReplicationArray []ReplicationInput

func (ReplicationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Replication)(nil)).Elem()
}

func (i ReplicationArray) ToReplicationArrayOutput() ReplicationArrayOutput {
	return i.ToReplicationArrayOutputWithContext(context.Background())
}

func (i ReplicationArray) ToReplicationArrayOutputWithContext(ctx context.Context) ReplicationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReplicationArrayOutput)
}

func (i ReplicationArray) ToOutput(ctx context.Context) pulumix.Output[[]*Replication] {
	return pulumix.Output[[]*Replication]{
		OutputState: i.ToReplicationArrayOutputWithContext(ctx).OutputState,
	}
}

// ReplicationMapInput is an input type that accepts ReplicationMap and ReplicationMapOutput values.
// You can construct a concrete instance of `ReplicationMapInput` via:
//
//	ReplicationMap{ "key": ReplicationArgs{...} }
type ReplicationMapInput interface {
	pulumi.Input

	ToReplicationMapOutput() ReplicationMapOutput
	ToReplicationMapOutputWithContext(context.Context) ReplicationMapOutput
}

type ReplicationMap map[string]ReplicationInput

func (ReplicationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Replication)(nil)).Elem()
}

func (i ReplicationMap) ToReplicationMapOutput() ReplicationMapOutput {
	return i.ToReplicationMapOutputWithContext(context.Background())
}

func (i ReplicationMap) ToReplicationMapOutputWithContext(ctx context.Context) ReplicationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReplicationMapOutput)
}

func (i ReplicationMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*Replication] {
	return pulumix.Output[map[string]*Replication]{
		OutputState: i.ToReplicationMapOutputWithContext(ctx).OutputState,
	}
}

type ReplicationOutput struct{ *pulumi.OutputState }

func (ReplicationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Replication)(nil)).Elem()
}

func (o ReplicationOutput) ToReplicationOutput() ReplicationOutput {
	return o
}

func (o ReplicationOutput) ToReplicationOutputWithContext(ctx context.Context) ReplicationOutput {
	return o
}

func (o ReplicationOutput) ToOutput(ctx context.Context) pulumix.Output[*Replication] {
	return pulumix.Output[*Replication]{
		OutputState: o.OutputState,
	}
}

func (o ReplicationOutput) Action() pulumi.StringOutput {
	return o.ApplyT(func(v *Replication) pulumi.StringOutput { return v.Action }).(pulumi.StringOutput)
}

func (o ReplicationOutput) Deletion() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Replication) pulumi.BoolPtrOutput { return v.Deletion }).(pulumi.BoolPtrOutput)
}

func (o ReplicationOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Replication) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o ReplicationOutput) DestNamespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Replication) pulumi.StringPtrOutput { return v.DestNamespace }).(pulumi.StringPtrOutput)
}

func (o ReplicationOutput) DestNamespaceReplace() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Replication) pulumi.IntPtrOutput { return v.DestNamespaceReplace }).(pulumi.IntPtrOutput)
}

func (o ReplicationOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Replication) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o ReplicationOutput) Filters() ReplicationFilterArrayOutput {
	return o.ApplyT(func(v *Replication) ReplicationFilterArrayOutput { return v.Filters }).(ReplicationFilterArrayOutput)
}

func (o ReplicationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Replication) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ReplicationOutput) Override() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Replication) pulumi.BoolPtrOutput { return v.Override }).(pulumi.BoolPtrOutput)
}

func (o ReplicationOutput) RegistryId() pulumi.IntOutput {
	return o.ApplyT(func(v *Replication) pulumi.IntOutput { return v.RegistryId }).(pulumi.IntOutput)
}

func (o ReplicationOutput) ReplicationPolicyId() pulumi.IntOutput {
	return o.ApplyT(func(v *Replication) pulumi.IntOutput { return v.ReplicationPolicyId }).(pulumi.IntOutput)
}

func (o ReplicationOutput) Schedule() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Replication) pulumi.StringPtrOutput { return v.Schedule }).(pulumi.StringPtrOutput)
}

func (o ReplicationOutput) Speed() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Replication) pulumi.IntPtrOutput { return v.Speed }).(pulumi.IntPtrOutput)
}

type ReplicationArrayOutput struct{ *pulumi.OutputState }

func (ReplicationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Replication)(nil)).Elem()
}

func (o ReplicationArrayOutput) ToReplicationArrayOutput() ReplicationArrayOutput {
	return o
}

func (o ReplicationArrayOutput) ToReplicationArrayOutputWithContext(ctx context.Context) ReplicationArrayOutput {
	return o
}

func (o ReplicationArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*Replication] {
	return pulumix.Output[[]*Replication]{
		OutputState: o.OutputState,
	}
}

func (o ReplicationArrayOutput) Index(i pulumi.IntInput) ReplicationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Replication {
		return vs[0].([]*Replication)[vs[1].(int)]
	}).(ReplicationOutput)
}

type ReplicationMapOutput struct{ *pulumi.OutputState }

func (ReplicationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Replication)(nil)).Elem()
}

func (o ReplicationMapOutput) ToReplicationMapOutput() ReplicationMapOutput {
	return o
}

func (o ReplicationMapOutput) ToReplicationMapOutputWithContext(ctx context.Context) ReplicationMapOutput {
	return o
}

func (o ReplicationMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*Replication] {
	return pulumix.Output[map[string]*Replication]{
		OutputState: o.OutputState,
	}
}

func (o ReplicationMapOutput) MapIndex(k pulumi.StringInput) ReplicationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Replication {
		return vs[0].(map[string]*Replication)[vs[1].(string)]
	}).(ReplicationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ReplicationInput)(nil)).Elem(), &Replication{})
	pulumi.RegisterInputType(reflect.TypeOf((*ReplicationArrayInput)(nil)).Elem(), ReplicationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ReplicationMapInput)(nil)).Elem(), ReplicationMap{})
	pulumi.RegisterOutputType(ReplicationOutput{})
	pulumi.RegisterOutputType(ReplicationArrayOutput{})
	pulumi.RegisterOutputType(ReplicationMapOutput{})
}
