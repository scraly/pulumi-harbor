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

// ## # Resource: RobotAccount
//
// Harbor supports different levels of robot accounts. Currently `system` and `project` level robot accounts are supported.
//
// ## Example Usage
//
// ### System Level
// Introduced in harbor 2.2.0, system level robot accounts can have basically [all available permissions](https://github.com/goharbor/harbor/blob/-/src/common/rbac/const.go) in harbor and are not dependent on a single project.
//
// The above example, creates a system level robot account with permissions to
// - permission to create labels on system level
// - pull repository across all projects
// - push repository to project "my-project-name"
// - read helm-chart and helm-chart-version in project "my-project-name"
//
// ### Project Level
//
// Other than system level robot accounts, project level robot accounts can interact on project level only.
// The [available permissions](https://github.com/goharbor/harbor/blob/-/src/common/rbac/const.go) are mostly the same as for system level robots.
//
// The above example creates a project level robot account with permissions to
// - pull repository on project "main"
// - push repository on project "main"
//
// ## Import
//
// Harbor robot account can be imported using the `robot account id` eg, `<break><break>```sh<break> $ pulumi import harbor:index/robotAccount:RobotAccount system /robots/123 <break>```<break><break>`
type RobotAccount struct {
	pulumi.CustomResourceState

	Description pulumi.StringPtrOutput            `pulumi:"description"`
	Disable     pulumi.BoolPtrOutput              `pulumi:"disable"`
	Duration    pulumi.IntPtrOutput               `pulumi:"duration"`
	FullName    pulumi.StringOutput               `pulumi:"fullName"`
	Level       pulumi.StringOutput               `pulumi:"level"`
	Name        pulumi.StringOutput               `pulumi:"name"`
	Permissions RobotAccountPermissionArrayOutput `pulumi:"permissions"`
	RobotId     pulumi.StringOutput               `pulumi:"robotId"`
	Secret      pulumi.StringOutput               `pulumi:"secret"`
}

// NewRobotAccount registers a new resource with the given unique name, arguments, and options.
func NewRobotAccount(ctx *pulumi.Context,
	name string, args *RobotAccountArgs, opts ...pulumi.ResourceOption) (*RobotAccount, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Level == nil {
		return nil, errors.New("invalid value for required argument 'Level'")
	}
	if args.Permissions == nil {
		return nil, errors.New("invalid value for required argument 'Permissions'")
	}
	if args.Secret != nil {
		args.Secret = pulumi.ToSecret(args.Secret).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"secret",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource RobotAccount
	err := ctx.RegisterResource("harbor:index/robotAccount:RobotAccount", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRobotAccount gets an existing RobotAccount resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRobotAccount(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RobotAccountState, opts ...pulumi.ResourceOption) (*RobotAccount, error) {
	var resource RobotAccount
	err := ctx.ReadResource("harbor:index/robotAccount:RobotAccount", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RobotAccount resources.
type robotAccountState struct {
	Description *string                  `pulumi:"description"`
	Disable     *bool                    `pulumi:"disable"`
	Duration    *int                     `pulumi:"duration"`
	FullName    *string                  `pulumi:"fullName"`
	Level       *string                  `pulumi:"level"`
	Name        *string                  `pulumi:"name"`
	Permissions []RobotAccountPermission `pulumi:"permissions"`
	RobotId     *string                  `pulumi:"robotId"`
	Secret      *string                  `pulumi:"secret"`
}

type RobotAccountState struct {
	Description pulumi.StringPtrInput
	Disable     pulumi.BoolPtrInput
	Duration    pulumi.IntPtrInput
	FullName    pulumi.StringPtrInput
	Level       pulumi.StringPtrInput
	Name        pulumi.StringPtrInput
	Permissions RobotAccountPermissionArrayInput
	RobotId     pulumi.StringPtrInput
	Secret      pulumi.StringPtrInput
}

func (RobotAccountState) ElementType() reflect.Type {
	return reflect.TypeOf((*robotAccountState)(nil)).Elem()
}

type robotAccountArgs struct {
	Description *string                  `pulumi:"description"`
	Disable     *bool                    `pulumi:"disable"`
	Duration    *int                     `pulumi:"duration"`
	Level       string                   `pulumi:"level"`
	Name        *string                  `pulumi:"name"`
	Permissions []RobotAccountPermission `pulumi:"permissions"`
	Secret      *string                  `pulumi:"secret"`
}

// The set of arguments for constructing a RobotAccount resource.
type RobotAccountArgs struct {
	Description pulumi.StringPtrInput
	Disable     pulumi.BoolPtrInput
	Duration    pulumi.IntPtrInput
	Level       pulumi.StringInput
	Name        pulumi.StringPtrInput
	Permissions RobotAccountPermissionArrayInput
	Secret      pulumi.StringPtrInput
}

func (RobotAccountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*robotAccountArgs)(nil)).Elem()
}

type RobotAccountInput interface {
	pulumi.Input

	ToRobotAccountOutput() RobotAccountOutput
	ToRobotAccountOutputWithContext(ctx context.Context) RobotAccountOutput
}

func (*RobotAccount) ElementType() reflect.Type {
	return reflect.TypeOf((**RobotAccount)(nil)).Elem()
}

func (i *RobotAccount) ToRobotAccountOutput() RobotAccountOutput {
	return i.ToRobotAccountOutputWithContext(context.Background())
}

func (i *RobotAccount) ToRobotAccountOutputWithContext(ctx context.Context) RobotAccountOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RobotAccountOutput)
}

func (i *RobotAccount) ToOutput(ctx context.Context) pulumix.Output[*RobotAccount] {
	return pulumix.Output[*RobotAccount]{
		OutputState: i.ToRobotAccountOutputWithContext(ctx).OutputState,
	}
}

// RobotAccountArrayInput is an input type that accepts RobotAccountArray and RobotAccountArrayOutput values.
// You can construct a concrete instance of `RobotAccountArrayInput` via:
//
//	RobotAccountArray{ RobotAccountArgs{...} }
type RobotAccountArrayInput interface {
	pulumi.Input

	ToRobotAccountArrayOutput() RobotAccountArrayOutput
	ToRobotAccountArrayOutputWithContext(context.Context) RobotAccountArrayOutput
}

type RobotAccountArray []RobotAccountInput

func (RobotAccountArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RobotAccount)(nil)).Elem()
}

func (i RobotAccountArray) ToRobotAccountArrayOutput() RobotAccountArrayOutput {
	return i.ToRobotAccountArrayOutputWithContext(context.Background())
}

func (i RobotAccountArray) ToRobotAccountArrayOutputWithContext(ctx context.Context) RobotAccountArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RobotAccountArrayOutput)
}

func (i RobotAccountArray) ToOutput(ctx context.Context) pulumix.Output[[]*RobotAccount] {
	return pulumix.Output[[]*RobotAccount]{
		OutputState: i.ToRobotAccountArrayOutputWithContext(ctx).OutputState,
	}
}

// RobotAccountMapInput is an input type that accepts RobotAccountMap and RobotAccountMapOutput values.
// You can construct a concrete instance of `RobotAccountMapInput` via:
//
//	RobotAccountMap{ "key": RobotAccountArgs{...} }
type RobotAccountMapInput interface {
	pulumi.Input

	ToRobotAccountMapOutput() RobotAccountMapOutput
	ToRobotAccountMapOutputWithContext(context.Context) RobotAccountMapOutput
}

type RobotAccountMap map[string]RobotAccountInput

func (RobotAccountMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RobotAccount)(nil)).Elem()
}

func (i RobotAccountMap) ToRobotAccountMapOutput() RobotAccountMapOutput {
	return i.ToRobotAccountMapOutputWithContext(context.Background())
}

func (i RobotAccountMap) ToRobotAccountMapOutputWithContext(ctx context.Context) RobotAccountMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RobotAccountMapOutput)
}

func (i RobotAccountMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*RobotAccount] {
	return pulumix.Output[map[string]*RobotAccount]{
		OutputState: i.ToRobotAccountMapOutputWithContext(ctx).OutputState,
	}
}

type RobotAccountOutput struct{ *pulumi.OutputState }

func (RobotAccountOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RobotAccount)(nil)).Elem()
}

func (o RobotAccountOutput) ToRobotAccountOutput() RobotAccountOutput {
	return o
}

func (o RobotAccountOutput) ToRobotAccountOutputWithContext(ctx context.Context) RobotAccountOutput {
	return o
}

func (o RobotAccountOutput) ToOutput(ctx context.Context) pulumix.Output[*RobotAccount] {
	return pulumix.Output[*RobotAccount]{
		OutputState: o.OutputState,
	}
}

func (o RobotAccountOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RobotAccount) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o RobotAccountOutput) Disable() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *RobotAccount) pulumi.BoolPtrOutput { return v.Disable }).(pulumi.BoolPtrOutput)
}

func (o RobotAccountOutput) Duration() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *RobotAccount) pulumi.IntPtrOutput { return v.Duration }).(pulumi.IntPtrOutput)
}

func (o RobotAccountOutput) FullName() pulumi.StringOutput {
	return o.ApplyT(func(v *RobotAccount) pulumi.StringOutput { return v.FullName }).(pulumi.StringOutput)
}

func (o RobotAccountOutput) Level() pulumi.StringOutput {
	return o.ApplyT(func(v *RobotAccount) pulumi.StringOutput { return v.Level }).(pulumi.StringOutput)
}

func (o RobotAccountOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *RobotAccount) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o RobotAccountOutput) Permissions() RobotAccountPermissionArrayOutput {
	return o.ApplyT(func(v *RobotAccount) RobotAccountPermissionArrayOutput { return v.Permissions }).(RobotAccountPermissionArrayOutput)
}

func (o RobotAccountOutput) RobotId() pulumi.StringOutput {
	return o.ApplyT(func(v *RobotAccount) pulumi.StringOutput { return v.RobotId }).(pulumi.StringOutput)
}

func (o RobotAccountOutput) Secret() pulumi.StringOutput {
	return o.ApplyT(func(v *RobotAccount) pulumi.StringOutput { return v.Secret }).(pulumi.StringOutput)
}

type RobotAccountArrayOutput struct{ *pulumi.OutputState }

func (RobotAccountArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RobotAccount)(nil)).Elem()
}

func (o RobotAccountArrayOutput) ToRobotAccountArrayOutput() RobotAccountArrayOutput {
	return o
}

func (o RobotAccountArrayOutput) ToRobotAccountArrayOutputWithContext(ctx context.Context) RobotAccountArrayOutput {
	return o
}

func (o RobotAccountArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*RobotAccount] {
	return pulumix.Output[[]*RobotAccount]{
		OutputState: o.OutputState,
	}
}

func (o RobotAccountArrayOutput) Index(i pulumi.IntInput) RobotAccountOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RobotAccount {
		return vs[0].([]*RobotAccount)[vs[1].(int)]
	}).(RobotAccountOutput)
}

type RobotAccountMapOutput struct{ *pulumi.OutputState }

func (RobotAccountMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RobotAccount)(nil)).Elem()
}

func (o RobotAccountMapOutput) ToRobotAccountMapOutput() RobotAccountMapOutput {
	return o
}

func (o RobotAccountMapOutput) ToRobotAccountMapOutputWithContext(ctx context.Context) RobotAccountMapOutput {
	return o
}

func (o RobotAccountMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*RobotAccount] {
	return pulumix.Output[map[string]*RobotAccount]{
		OutputState: o.OutputState,
	}
}

func (o RobotAccountMapOutput) MapIndex(k pulumi.StringInput) RobotAccountOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RobotAccount {
		return vs[0].(map[string]*RobotAccount)[vs[1].(string)]
	}).(RobotAccountOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RobotAccountInput)(nil)).Elem(), &RobotAccount{})
	pulumi.RegisterInputType(reflect.TypeOf((*RobotAccountArrayInput)(nil)).Elem(), RobotAccountArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RobotAccountMapInput)(nil)).Elem(), RobotAccountMap{})
	pulumi.RegisterOutputType(RobotAccountOutput{})
	pulumi.RegisterOutputType(RobotAccountArrayOutput{})
	pulumi.RegisterOutputType(RobotAccountMapOutput{})
}
