// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package harbor

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-harbor/sdk/v3/go/harbor/internal"
)

type PurgeAuditLog struct {
	pulumi.CustomResourceState

	AuditRetentionHour pulumi.IntOutput    `pulumi:"auditRetentionHour"`
	IncludeOperations  pulumi.StringOutput `pulumi:"includeOperations"`
	Schedule           pulumi.StringOutput `pulumi:"schedule"`
}

// NewPurgeAuditLog registers a new resource with the given unique name, arguments, and options.
func NewPurgeAuditLog(ctx *pulumi.Context,
	name string, args *PurgeAuditLogArgs, opts ...pulumi.ResourceOption) (*PurgeAuditLog, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AuditRetentionHour == nil {
		return nil, errors.New("invalid value for required argument 'AuditRetentionHour'")
	}
	if args.IncludeOperations == nil {
		return nil, errors.New("invalid value for required argument 'IncludeOperations'")
	}
	if args.Schedule == nil {
		return nil, errors.New("invalid value for required argument 'Schedule'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource PurgeAuditLog
	err := ctx.RegisterResource("harbor:index/purgeAuditLog:PurgeAuditLog", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPurgeAuditLog gets an existing PurgeAuditLog resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPurgeAuditLog(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PurgeAuditLogState, opts ...pulumi.ResourceOption) (*PurgeAuditLog, error) {
	var resource PurgeAuditLog
	err := ctx.ReadResource("harbor:index/purgeAuditLog:PurgeAuditLog", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PurgeAuditLog resources.
type purgeAuditLogState struct {
	AuditRetentionHour *int    `pulumi:"auditRetentionHour"`
	IncludeOperations  *string `pulumi:"includeOperations"`
	Schedule           *string `pulumi:"schedule"`
}

type PurgeAuditLogState struct {
	AuditRetentionHour pulumi.IntPtrInput
	IncludeOperations  pulumi.StringPtrInput
	Schedule           pulumi.StringPtrInput
}

func (PurgeAuditLogState) ElementType() reflect.Type {
	return reflect.TypeOf((*purgeAuditLogState)(nil)).Elem()
}

type purgeAuditLogArgs struct {
	AuditRetentionHour int    `pulumi:"auditRetentionHour"`
	IncludeOperations  string `pulumi:"includeOperations"`
	Schedule           string `pulumi:"schedule"`
}

// The set of arguments for constructing a PurgeAuditLog resource.
type PurgeAuditLogArgs struct {
	AuditRetentionHour pulumi.IntInput
	IncludeOperations  pulumi.StringInput
	Schedule           pulumi.StringInput
}

func (PurgeAuditLogArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*purgeAuditLogArgs)(nil)).Elem()
}

type PurgeAuditLogInput interface {
	pulumi.Input

	ToPurgeAuditLogOutput() PurgeAuditLogOutput
	ToPurgeAuditLogOutputWithContext(ctx context.Context) PurgeAuditLogOutput
}

func (*PurgeAuditLog) ElementType() reflect.Type {
	return reflect.TypeOf((**PurgeAuditLog)(nil)).Elem()
}

func (i *PurgeAuditLog) ToPurgeAuditLogOutput() PurgeAuditLogOutput {
	return i.ToPurgeAuditLogOutputWithContext(context.Background())
}

func (i *PurgeAuditLog) ToPurgeAuditLogOutputWithContext(ctx context.Context) PurgeAuditLogOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PurgeAuditLogOutput)
}

// PurgeAuditLogArrayInput is an input type that accepts PurgeAuditLogArray and PurgeAuditLogArrayOutput values.
// You can construct a concrete instance of `PurgeAuditLogArrayInput` via:
//
//	PurgeAuditLogArray{ PurgeAuditLogArgs{...} }
type PurgeAuditLogArrayInput interface {
	pulumi.Input

	ToPurgeAuditLogArrayOutput() PurgeAuditLogArrayOutput
	ToPurgeAuditLogArrayOutputWithContext(context.Context) PurgeAuditLogArrayOutput
}

type PurgeAuditLogArray []PurgeAuditLogInput

func (PurgeAuditLogArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PurgeAuditLog)(nil)).Elem()
}

func (i PurgeAuditLogArray) ToPurgeAuditLogArrayOutput() PurgeAuditLogArrayOutput {
	return i.ToPurgeAuditLogArrayOutputWithContext(context.Background())
}

func (i PurgeAuditLogArray) ToPurgeAuditLogArrayOutputWithContext(ctx context.Context) PurgeAuditLogArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PurgeAuditLogArrayOutput)
}

// PurgeAuditLogMapInput is an input type that accepts PurgeAuditLogMap and PurgeAuditLogMapOutput values.
// You can construct a concrete instance of `PurgeAuditLogMapInput` via:
//
//	PurgeAuditLogMap{ "key": PurgeAuditLogArgs{...} }
type PurgeAuditLogMapInput interface {
	pulumi.Input

	ToPurgeAuditLogMapOutput() PurgeAuditLogMapOutput
	ToPurgeAuditLogMapOutputWithContext(context.Context) PurgeAuditLogMapOutput
}

type PurgeAuditLogMap map[string]PurgeAuditLogInput

func (PurgeAuditLogMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PurgeAuditLog)(nil)).Elem()
}

func (i PurgeAuditLogMap) ToPurgeAuditLogMapOutput() PurgeAuditLogMapOutput {
	return i.ToPurgeAuditLogMapOutputWithContext(context.Background())
}

func (i PurgeAuditLogMap) ToPurgeAuditLogMapOutputWithContext(ctx context.Context) PurgeAuditLogMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PurgeAuditLogMapOutput)
}

type PurgeAuditLogOutput struct{ *pulumi.OutputState }

func (PurgeAuditLogOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PurgeAuditLog)(nil)).Elem()
}

func (o PurgeAuditLogOutput) ToPurgeAuditLogOutput() PurgeAuditLogOutput {
	return o
}

func (o PurgeAuditLogOutput) ToPurgeAuditLogOutputWithContext(ctx context.Context) PurgeAuditLogOutput {
	return o
}

func (o PurgeAuditLogOutput) AuditRetentionHour() pulumi.IntOutput {
	return o.ApplyT(func(v *PurgeAuditLog) pulumi.IntOutput { return v.AuditRetentionHour }).(pulumi.IntOutput)
}

func (o PurgeAuditLogOutput) IncludeOperations() pulumi.StringOutput {
	return o.ApplyT(func(v *PurgeAuditLog) pulumi.StringOutput { return v.IncludeOperations }).(pulumi.StringOutput)
}

func (o PurgeAuditLogOutput) Schedule() pulumi.StringOutput {
	return o.ApplyT(func(v *PurgeAuditLog) pulumi.StringOutput { return v.Schedule }).(pulumi.StringOutput)
}

type PurgeAuditLogArrayOutput struct{ *pulumi.OutputState }

func (PurgeAuditLogArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PurgeAuditLog)(nil)).Elem()
}

func (o PurgeAuditLogArrayOutput) ToPurgeAuditLogArrayOutput() PurgeAuditLogArrayOutput {
	return o
}

func (o PurgeAuditLogArrayOutput) ToPurgeAuditLogArrayOutputWithContext(ctx context.Context) PurgeAuditLogArrayOutput {
	return o
}

func (o PurgeAuditLogArrayOutput) Index(i pulumi.IntInput) PurgeAuditLogOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *PurgeAuditLog {
		return vs[0].([]*PurgeAuditLog)[vs[1].(int)]
	}).(PurgeAuditLogOutput)
}

type PurgeAuditLogMapOutput struct{ *pulumi.OutputState }

func (PurgeAuditLogMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PurgeAuditLog)(nil)).Elem()
}

func (o PurgeAuditLogMapOutput) ToPurgeAuditLogMapOutput() PurgeAuditLogMapOutput {
	return o
}

func (o PurgeAuditLogMapOutput) ToPurgeAuditLogMapOutputWithContext(ctx context.Context) PurgeAuditLogMapOutput {
	return o
}

func (o PurgeAuditLogMapOutput) MapIndex(k pulumi.StringInput) PurgeAuditLogOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *PurgeAuditLog {
		return vs[0].(map[string]*PurgeAuditLog)[vs[1].(string)]
	}).(PurgeAuditLogOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PurgeAuditLogInput)(nil)).Elem(), &PurgeAuditLog{})
	pulumi.RegisterInputType(reflect.TypeOf((*PurgeAuditLogArrayInput)(nil)).Elem(), PurgeAuditLogArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PurgeAuditLogMapInput)(nil)).Elem(), PurgeAuditLogMap{})
	pulumi.RegisterOutputType(PurgeAuditLogOutput{})
	pulumi.RegisterOutputType(PurgeAuditLogArrayOutput{})
	pulumi.RegisterOutputType(PurgeAuditLogMapOutput{})
}
