// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package harbor

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-harbor/sdk/v3/go/harbor/internal"
)

var _ = internal.GetEnvOrDefault

type ReplicationFilter struct {
	Decoration *string  `pulumi:"decoration"`
	Labels     []string `pulumi:"labels"`
	Name       *string  `pulumi:"name"`
	Resource   *string  `pulumi:"resource"`
	Tag        *string  `pulumi:"tag"`
}

// ReplicationFilterInput is an input type that accepts ReplicationFilterArgs and ReplicationFilterOutput values.
// You can construct a concrete instance of `ReplicationFilterInput` via:
//
//	ReplicationFilterArgs{...}
type ReplicationFilterInput interface {
	pulumi.Input

	ToReplicationFilterOutput() ReplicationFilterOutput
	ToReplicationFilterOutputWithContext(context.Context) ReplicationFilterOutput
}

type ReplicationFilterArgs struct {
	Decoration pulumi.StringPtrInput   `pulumi:"decoration"`
	Labels     pulumi.StringArrayInput `pulumi:"labels"`
	Name       pulumi.StringPtrInput   `pulumi:"name"`
	Resource   pulumi.StringPtrInput   `pulumi:"resource"`
	Tag        pulumi.StringPtrInput   `pulumi:"tag"`
}

func (ReplicationFilterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ReplicationFilter)(nil)).Elem()
}

func (i ReplicationFilterArgs) ToReplicationFilterOutput() ReplicationFilterOutput {
	return i.ToReplicationFilterOutputWithContext(context.Background())
}

func (i ReplicationFilterArgs) ToReplicationFilterOutputWithContext(ctx context.Context) ReplicationFilterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReplicationFilterOutput)
}

// ReplicationFilterArrayInput is an input type that accepts ReplicationFilterArray and ReplicationFilterArrayOutput values.
// You can construct a concrete instance of `ReplicationFilterArrayInput` via:
//
//	ReplicationFilterArray{ ReplicationFilterArgs{...} }
type ReplicationFilterArrayInput interface {
	pulumi.Input

	ToReplicationFilterArrayOutput() ReplicationFilterArrayOutput
	ToReplicationFilterArrayOutputWithContext(context.Context) ReplicationFilterArrayOutput
}

type ReplicationFilterArray []ReplicationFilterInput

func (ReplicationFilterArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ReplicationFilter)(nil)).Elem()
}

func (i ReplicationFilterArray) ToReplicationFilterArrayOutput() ReplicationFilterArrayOutput {
	return i.ToReplicationFilterArrayOutputWithContext(context.Background())
}

func (i ReplicationFilterArray) ToReplicationFilterArrayOutputWithContext(ctx context.Context) ReplicationFilterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReplicationFilterArrayOutput)
}

type ReplicationFilterOutput struct{ *pulumi.OutputState }

func (ReplicationFilterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReplicationFilter)(nil)).Elem()
}

func (o ReplicationFilterOutput) ToReplicationFilterOutput() ReplicationFilterOutput {
	return o
}

func (o ReplicationFilterOutput) ToReplicationFilterOutputWithContext(ctx context.Context) ReplicationFilterOutput {
	return o
}

func (o ReplicationFilterOutput) Decoration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ReplicationFilter) *string { return v.Decoration }).(pulumi.StringPtrOutput)
}

func (o ReplicationFilterOutput) Labels() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ReplicationFilter) []string { return v.Labels }).(pulumi.StringArrayOutput)
}

func (o ReplicationFilterOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ReplicationFilter) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ReplicationFilterOutput) Resource() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ReplicationFilter) *string { return v.Resource }).(pulumi.StringPtrOutput)
}

func (o ReplicationFilterOutput) Tag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ReplicationFilter) *string { return v.Tag }).(pulumi.StringPtrOutput)
}

type ReplicationFilterArrayOutput struct{ *pulumi.OutputState }

func (ReplicationFilterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ReplicationFilter)(nil)).Elem()
}

func (o ReplicationFilterArrayOutput) ToReplicationFilterArrayOutput() ReplicationFilterArrayOutput {
	return o
}

func (o ReplicationFilterArrayOutput) ToReplicationFilterArrayOutputWithContext(ctx context.Context) ReplicationFilterArrayOutput {
	return o
}

func (o ReplicationFilterArrayOutput) Index(i pulumi.IntInput) ReplicationFilterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ReplicationFilter {
		return vs[0].([]ReplicationFilter)[vs[1].(int)]
	}).(ReplicationFilterOutput)
}

type RetentionPolicyRule struct {
	// retain always.
	AlwaysRetain *bool `pulumi:"alwaysRetain"`
	// Specify if the rule is disable or not. Defaults to `false`
	Disabled *bool `pulumi:"disabled"`
	// retain the most recently pulled n artifacts.
	MostRecentlyPulled *int `pulumi:"mostRecentlyPulled"`
	// retain the most recently pushed n artifacts.
	MostRecentlyPushed *int `pulumi:"mostRecentlyPushed"`
	// retains the artifacts pulled within the lasts n days.
	NDaysSinceLastPull *int `pulumi:"nDaysSinceLastPull"`
	// retains the artifacts pushed within the lasts n days.
	NDaysSinceLastPush *int `pulumi:"nDaysSinceLastPush"`
	// For the repositories excuding.
	RepoExcluding *string `pulumi:"repoExcluding"`
	// For the repositories matching.
	RepoMatching *string `pulumi:"repoMatching"`
	// For the tag excuding.
	TagExcluding *string `pulumi:"tagExcluding"`
	// For the tag matching.
	TagMatching *string `pulumi:"tagMatching"`
	// with untagged artifacts. Defaults to `true`
	//
	// > Multiple tags or repositories must be provided as a comma-separated list wrapped into curly brackets `{ }`. Otherwise, the value is interpreted as a single value.
	UntaggedArtifacts *bool `pulumi:"untaggedArtifacts"`
}

// RetentionPolicyRuleInput is an input type that accepts RetentionPolicyRuleArgs and RetentionPolicyRuleOutput values.
// You can construct a concrete instance of `RetentionPolicyRuleInput` via:
//
//	RetentionPolicyRuleArgs{...}
type RetentionPolicyRuleInput interface {
	pulumi.Input

	ToRetentionPolicyRuleOutput() RetentionPolicyRuleOutput
	ToRetentionPolicyRuleOutputWithContext(context.Context) RetentionPolicyRuleOutput
}

type RetentionPolicyRuleArgs struct {
	// retain always.
	AlwaysRetain pulumi.BoolPtrInput `pulumi:"alwaysRetain"`
	// Specify if the rule is disable or not. Defaults to `false`
	Disabled pulumi.BoolPtrInput `pulumi:"disabled"`
	// retain the most recently pulled n artifacts.
	MostRecentlyPulled pulumi.IntPtrInput `pulumi:"mostRecentlyPulled"`
	// retain the most recently pushed n artifacts.
	MostRecentlyPushed pulumi.IntPtrInput `pulumi:"mostRecentlyPushed"`
	// retains the artifacts pulled within the lasts n days.
	NDaysSinceLastPull pulumi.IntPtrInput `pulumi:"nDaysSinceLastPull"`
	// retains the artifacts pushed within the lasts n days.
	NDaysSinceLastPush pulumi.IntPtrInput `pulumi:"nDaysSinceLastPush"`
	// For the repositories excuding.
	RepoExcluding pulumi.StringPtrInput `pulumi:"repoExcluding"`
	// For the repositories matching.
	RepoMatching pulumi.StringPtrInput `pulumi:"repoMatching"`
	// For the tag excuding.
	TagExcluding pulumi.StringPtrInput `pulumi:"tagExcluding"`
	// For the tag matching.
	TagMatching pulumi.StringPtrInput `pulumi:"tagMatching"`
	// with untagged artifacts. Defaults to `true`
	//
	// > Multiple tags or repositories must be provided as a comma-separated list wrapped into curly brackets `{ }`. Otherwise, the value is interpreted as a single value.
	UntaggedArtifacts pulumi.BoolPtrInput `pulumi:"untaggedArtifacts"`
}

func (RetentionPolicyRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RetentionPolicyRule)(nil)).Elem()
}

func (i RetentionPolicyRuleArgs) ToRetentionPolicyRuleOutput() RetentionPolicyRuleOutput {
	return i.ToRetentionPolicyRuleOutputWithContext(context.Background())
}

func (i RetentionPolicyRuleArgs) ToRetentionPolicyRuleOutputWithContext(ctx context.Context) RetentionPolicyRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RetentionPolicyRuleOutput)
}

// RetentionPolicyRuleArrayInput is an input type that accepts RetentionPolicyRuleArray and RetentionPolicyRuleArrayOutput values.
// You can construct a concrete instance of `RetentionPolicyRuleArrayInput` via:
//
//	RetentionPolicyRuleArray{ RetentionPolicyRuleArgs{...} }
type RetentionPolicyRuleArrayInput interface {
	pulumi.Input

	ToRetentionPolicyRuleArrayOutput() RetentionPolicyRuleArrayOutput
	ToRetentionPolicyRuleArrayOutputWithContext(context.Context) RetentionPolicyRuleArrayOutput
}

type RetentionPolicyRuleArray []RetentionPolicyRuleInput

func (RetentionPolicyRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RetentionPolicyRule)(nil)).Elem()
}

func (i RetentionPolicyRuleArray) ToRetentionPolicyRuleArrayOutput() RetentionPolicyRuleArrayOutput {
	return i.ToRetentionPolicyRuleArrayOutputWithContext(context.Background())
}

func (i RetentionPolicyRuleArray) ToRetentionPolicyRuleArrayOutputWithContext(ctx context.Context) RetentionPolicyRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RetentionPolicyRuleArrayOutput)
}

type RetentionPolicyRuleOutput struct{ *pulumi.OutputState }

func (RetentionPolicyRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RetentionPolicyRule)(nil)).Elem()
}

func (o RetentionPolicyRuleOutput) ToRetentionPolicyRuleOutput() RetentionPolicyRuleOutput {
	return o
}

func (o RetentionPolicyRuleOutput) ToRetentionPolicyRuleOutputWithContext(ctx context.Context) RetentionPolicyRuleOutput {
	return o
}

// retain always.
func (o RetentionPolicyRuleOutput) AlwaysRetain() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v RetentionPolicyRule) *bool { return v.AlwaysRetain }).(pulumi.BoolPtrOutput)
}

// Specify if the rule is disable or not. Defaults to `false`
func (o RetentionPolicyRuleOutput) Disabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v RetentionPolicyRule) *bool { return v.Disabled }).(pulumi.BoolPtrOutput)
}

// retain the most recently pulled n artifacts.
func (o RetentionPolicyRuleOutput) MostRecentlyPulled() pulumi.IntPtrOutput {
	return o.ApplyT(func(v RetentionPolicyRule) *int { return v.MostRecentlyPulled }).(pulumi.IntPtrOutput)
}

// retain the most recently pushed n artifacts.
func (o RetentionPolicyRuleOutput) MostRecentlyPushed() pulumi.IntPtrOutput {
	return o.ApplyT(func(v RetentionPolicyRule) *int { return v.MostRecentlyPushed }).(pulumi.IntPtrOutput)
}

// retains the artifacts pulled within the lasts n days.
func (o RetentionPolicyRuleOutput) NDaysSinceLastPull() pulumi.IntPtrOutput {
	return o.ApplyT(func(v RetentionPolicyRule) *int { return v.NDaysSinceLastPull }).(pulumi.IntPtrOutput)
}

// retains the artifacts pushed within the lasts n days.
func (o RetentionPolicyRuleOutput) NDaysSinceLastPush() pulumi.IntPtrOutput {
	return o.ApplyT(func(v RetentionPolicyRule) *int { return v.NDaysSinceLastPush }).(pulumi.IntPtrOutput)
}

// For the repositories excuding.
func (o RetentionPolicyRuleOutput) RepoExcluding() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RetentionPolicyRule) *string { return v.RepoExcluding }).(pulumi.StringPtrOutput)
}

// For the repositories matching.
func (o RetentionPolicyRuleOutput) RepoMatching() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RetentionPolicyRule) *string { return v.RepoMatching }).(pulumi.StringPtrOutput)
}

// For the tag excuding.
func (o RetentionPolicyRuleOutput) TagExcluding() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RetentionPolicyRule) *string { return v.TagExcluding }).(pulumi.StringPtrOutput)
}

// For the tag matching.
func (o RetentionPolicyRuleOutput) TagMatching() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RetentionPolicyRule) *string { return v.TagMatching }).(pulumi.StringPtrOutput)
}

// with untagged artifacts. Defaults to `true`
//
// > Multiple tags or repositories must be provided as a comma-separated list wrapped into curly brackets `{ }`. Otherwise, the value is interpreted as a single value.
func (o RetentionPolicyRuleOutput) UntaggedArtifacts() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v RetentionPolicyRule) *bool { return v.UntaggedArtifacts }).(pulumi.BoolPtrOutput)
}

type RetentionPolicyRuleArrayOutput struct{ *pulumi.OutputState }

func (RetentionPolicyRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RetentionPolicyRule)(nil)).Elem()
}

func (o RetentionPolicyRuleArrayOutput) ToRetentionPolicyRuleArrayOutput() RetentionPolicyRuleArrayOutput {
	return o
}

func (o RetentionPolicyRuleArrayOutput) ToRetentionPolicyRuleArrayOutputWithContext(ctx context.Context) RetentionPolicyRuleArrayOutput {
	return o
}

func (o RetentionPolicyRuleArrayOutput) Index(i pulumi.IntInput) RetentionPolicyRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RetentionPolicyRule {
		return vs[0].([]RetentionPolicyRule)[vs[1].(int)]
	}).(RetentionPolicyRuleOutput)
}

type RobotAccountPermission struct {
	Accesses  []RobotAccountPermissionAccess `pulumi:"accesses"`
	Kind      string                         `pulumi:"kind"`
	Namespace string                         `pulumi:"namespace"`
}

// RobotAccountPermissionInput is an input type that accepts RobotAccountPermissionArgs and RobotAccountPermissionOutput values.
// You can construct a concrete instance of `RobotAccountPermissionInput` via:
//
//	RobotAccountPermissionArgs{...}
type RobotAccountPermissionInput interface {
	pulumi.Input

	ToRobotAccountPermissionOutput() RobotAccountPermissionOutput
	ToRobotAccountPermissionOutputWithContext(context.Context) RobotAccountPermissionOutput
}

type RobotAccountPermissionArgs struct {
	Accesses  RobotAccountPermissionAccessArrayInput `pulumi:"accesses"`
	Kind      pulumi.StringInput                     `pulumi:"kind"`
	Namespace pulumi.StringInput                     `pulumi:"namespace"`
}

func (RobotAccountPermissionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RobotAccountPermission)(nil)).Elem()
}

func (i RobotAccountPermissionArgs) ToRobotAccountPermissionOutput() RobotAccountPermissionOutput {
	return i.ToRobotAccountPermissionOutputWithContext(context.Background())
}

func (i RobotAccountPermissionArgs) ToRobotAccountPermissionOutputWithContext(ctx context.Context) RobotAccountPermissionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RobotAccountPermissionOutput)
}

// RobotAccountPermissionArrayInput is an input type that accepts RobotAccountPermissionArray and RobotAccountPermissionArrayOutput values.
// You can construct a concrete instance of `RobotAccountPermissionArrayInput` via:
//
//	RobotAccountPermissionArray{ RobotAccountPermissionArgs{...} }
type RobotAccountPermissionArrayInput interface {
	pulumi.Input

	ToRobotAccountPermissionArrayOutput() RobotAccountPermissionArrayOutput
	ToRobotAccountPermissionArrayOutputWithContext(context.Context) RobotAccountPermissionArrayOutput
}

type RobotAccountPermissionArray []RobotAccountPermissionInput

func (RobotAccountPermissionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RobotAccountPermission)(nil)).Elem()
}

func (i RobotAccountPermissionArray) ToRobotAccountPermissionArrayOutput() RobotAccountPermissionArrayOutput {
	return i.ToRobotAccountPermissionArrayOutputWithContext(context.Background())
}

func (i RobotAccountPermissionArray) ToRobotAccountPermissionArrayOutputWithContext(ctx context.Context) RobotAccountPermissionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RobotAccountPermissionArrayOutput)
}

type RobotAccountPermissionOutput struct{ *pulumi.OutputState }

func (RobotAccountPermissionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RobotAccountPermission)(nil)).Elem()
}

func (o RobotAccountPermissionOutput) ToRobotAccountPermissionOutput() RobotAccountPermissionOutput {
	return o
}

func (o RobotAccountPermissionOutput) ToRobotAccountPermissionOutputWithContext(ctx context.Context) RobotAccountPermissionOutput {
	return o
}

func (o RobotAccountPermissionOutput) Accesses() RobotAccountPermissionAccessArrayOutput {
	return o.ApplyT(func(v RobotAccountPermission) []RobotAccountPermissionAccess { return v.Accesses }).(RobotAccountPermissionAccessArrayOutput)
}

func (o RobotAccountPermissionOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v RobotAccountPermission) string { return v.Kind }).(pulumi.StringOutput)
}

func (o RobotAccountPermissionOutput) Namespace() pulumi.StringOutput {
	return o.ApplyT(func(v RobotAccountPermission) string { return v.Namespace }).(pulumi.StringOutput)
}

type RobotAccountPermissionArrayOutput struct{ *pulumi.OutputState }

func (RobotAccountPermissionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RobotAccountPermission)(nil)).Elem()
}

func (o RobotAccountPermissionArrayOutput) ToRobotAccountPermissionArrayOutput() RobotAccountPermissionArrayOutput {
	return o
}

func (o RobotAccountPermissionArrayOutput) ToRobotAccountPermissionArrayOutputWithContext(ctx context.Context) RobotAccountPermissionArrayOutput {
	return o
}

func (o RobotAccountPermissionArrayOutput) Index(i pulumi.IntInput) RobotAccountPermissionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RobotAccountPermission {
		return vs[0].([]RobotAccountPermission)[vs[1].(int)]
	}).(RobotAccountPermissionOutput)
}

type RobotAccountPermissionAccess struct {
	Action   string  `pulumi:"action"`
	Effect   *string `pulumi:"effect"`
	Resource string  `pulumi:"resource"`
}

// RobotAccountPermissionAccessInput is an input type that accepts RobotAccountPermissionAccessArgs and RobotAccountPermissionAccessOutput values.
// You can construct a concrete instance of `RobotAccountPermissionAccessInput` via:
//
//	RobotAccountPermissionAccessArgs{...}
type RobotAccountPermissionAccessInput interface {
	pulumi.Input

	ToRobotAccountPermissionAccessOutput() RobotAccountPermissionAccessOutput
	ToRobotAccountPermissionAccessOutputWithContext(context.Context) RobotAccountPermissionAccessOutput
}

type RobotAccountPermissionAccessArgs struct {
	Action   pulumi.StringInput    `pulumi:"action"`
	Effect   pulumi.StringPtrInput `pulumi:"effect"`
	Resource pulumi.StringInput    `pulumi:"resource"`
}

func (RobotAccountPermissionAccessArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RobotAccountPermissionAccess)(nil)).Elem()
}

func (i RobotAccountPermissionAccessArgs) ToRobotAccountPermissionAccessOutput() RobotAccountPermissionAccessOutput {
	return i.ToRobotAccountPermissionAccessOutputWithContext(context.Background())
}

func (i RobotAccountPermissionAccessArgs) ToRobotAccountPermissionAccessOutputWithContext(ctx context.Context) RobotAccountPermissionAccessOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RobotAccountPermissionAccessOutput)
}

// RobotAccountPermissionAccessArrayInput is an input type that accepts RobotAccountPermissionAccessArray and RobotAccountPermissionAccessArrayOutput values.
// You can construct a concrete instance of `RobotAccountPermissionAccessArrayInput` via:
//
//	RobotAccountPermissionAccessArray{ RobotAccountPermissionAccessArgs{...} }
type RobotAccountPermissionAccessArrayInput interface {
	pulumi.Input

	ToRobotAccountPermissionAccessArrayOutput() RobotAccountPermissionAccessArrayOutput
	ToRobotAccountPermissionAccessArrayOutputWithContext(context.Context) RobotAccountPermissionAccessArrayOutput
}

type RobotAccountPermissionAccessArray []RobotAccountPermissionAccessInput

func (RobotAccountPermissionAccessArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RobotAccountPermissionAccess)(nil)).Elem()
}

func (i RobotAccountPermissionAccessArray) ToRobotAccountPermissionAccessArrayOutput() RobotAccountPermissionAccessArrayOutput {
	return i.ToRobotAccountPermissionAccessArrayOutputWithContext(context.Background())
}

func (i RobotAccountPermissionAccessArray) ToRobotAccountPermissionAccessArrayOutputWithContext(ctx context.Context) RobotAccountPermissionAccessArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RobotAccountPermissionAccessArrayOutput)
}

type RobotAccountPermissionAccessOutput struct{ *pulumi.OutputState }

func (RobotAccountPermissionAccessOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RobotAccountPermissionAccess)(nil)).Elem()
}

func (o RobotAccountPermissionAccessOutput) ToRobotAccountPermissionAccessOutput() RobotAccountPermissionAccessOutput {
	return o
}

func (o RobotAccountPermissionAccessOutput) ToRobotAccountPermissionAccessOutputWithContext(ctx context.Context) RobotAccountPermissionAccessOutput {
	return o
}

func (o RobotAccountPermissionAccessOutput) Action() pulumi.StringOutput {
	return o.ApplyT(func(v RobotAccountPermissionAccess) string { return v.Action }).(pulumi.StringOutput)
}

func (o RobotAccountPermissionAccessOutput) Effect() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RobotAccountPermissionAccess) *string { return v.Effect }).(pulumi.StringPtrOutput)
}

func (o RobotAccountPermissionAccessOutput) Resource() pulumi.StringOutput {
	return o.ApplyT(func(v RobotAccountPermissionAccess) string { return v.Resource }).(pulumi.StringOutput)
}

type RobotAccountPermissionAccessArrayOutput struct{ *pulumi.OutputState }

func (RobotAccountPermissionAccessArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RobotAccountPermissionAccess)(nil)).Elem()
}

func (o RobotAccountPermissionAccessArrayOutput) ToRobotAccountPermissionAccessArrayOutput() RobotAccountPermissionAccessArrayOutput {
	return o
}

func (o RobotAccountPermissionAccessArrayOutput) ToRobotAccountPermissionAccessArrayOutputWithContext(ctx context.Context) RobotAccountPermissionAccessArrayOutput {
	return o
}

func (o RobotAccountPermissionAccessArrayOutput) Index(i pulumi.IntInput) RobotAccountPermissionAccessOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RobotAccountPermissionAccess {
		return vs[0].([]RobotAccountPermissionAccess)[vs[1].(int)]
	}).(RobotAccountPermissionAccessOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ReplicationFilterInput)(nil)).Elem(), ReplicationFilterArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ReplicationFilterArrayInput)(nil)).Elem(), ReplicationFilterArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RetentionPolicyRuleInput)(nil)).Elem(), RetentionPolicyRuleArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*RetentionPolicyRuleArrayInput)(nil)).Elem(), RetentionPolicyRuleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RobotAccountPermissionInput)(nil)).Elem(), RobotAccountPermissionArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*RobotAccountPermissionArrayInput)(nil)).Elem(), RobotAccountPermissionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RobotAccountPermissionAccessInput)(nil)).Elem(), RobotAccountPermissionAccessArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*RobotAccountPermissionAccessArrayInput)(nil)).Elem(), RobotAccountPermissionAccessArray{})
	pulumi.RegisterOutputType(ReplicationFilterOutput{})
	pulumi.RegisterOutputType(ReplicationFilterArrayOutput{})
	pulumi.RegisterOutputType(RetentionPolicyRuleOutput{})
	pulumi.RegisterOutputType(RetentionPolicyRuleArrayOutput{})
	pulumi.RegisterOutputType(RobotAccountPermissionOutput{})
	pulumi.RegisterOutputType(RobotAccountPermissionArrayOutput{})
	pulumi.RegisterOutputType(RobotAccountPermissionAccessOutput{})
	pulumi.RegisterOutputType(RobotAccountPermissionAccessArrayOutput{})
}
