// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package harbor

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
	"github.com/pulumiverse/pulumi-harbor/sdk/v3/go/harbor/internal"
)

// ## Example Usage
func GetProjects(ctx *pulumi.Context, args *GetProjectsArgs, opts ...pulumi.InvokeOption) (*GetProjectsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetProjectsResult
	err := ctx.Invoke("harbor:index/getProjects:getProjects", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getProjects.
type GetProjectsArgs struct {
	Name                  *string `pulumi:"name"`
	Public                *bool   `pulumi:"public"`
	Type                  *string `pulumi:"type"`
	VulnerabilityScanning *bool   `pulumi:"vulnerabilityScanning"`
}

// A collection of values returned by getProjects.
type GetProjectsResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id                    string               `pulumi:"id"`
	Name                  *string              `pulumi:"name"`
	Projects              []GetProjectsProject `pulumi:"projects"`
	Public                *bool                `pulumi:"public"`
	Type                  *string              `pulumi:"type"`
	VulnerabilityScanning *bool                `pulumi:"vulnerabilityScanning"`
}

func GetProjectsOutput(ctx *pulumi.Context, args GetProjectsOutputArgs, opts ...pulumi.InvokeOption) GetProjectsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetProjectsResult, error) {
			args := v.(GetProjectsArgs)
			r, err := GetProjects(ctx, &args, opts...)
			var s GetProjectsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetProjectsResultOutput)
}

// A collection of arguments for invoking getProjects.
type GetProjectsOutputArgs struct {
	Name                  pulumi.StringPtrInput `pulumi:"name"`
	Public                pulumi.BoolPtrInput   `pulumi:"public"`
	Type                  pulumi.StringPtrInput `pulumi:"type"`
	VulnerabilityScanning pulumi.BoolPtrInput   `pulumi:"vulnerabilityScanning"`
}

func (GetProjectsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetProjectsArgs)(nil)).Elem()
}

// A collection of values returned by getProjects.
type GetProjectsResultOutput struct{ *pulumi.OutputState }

func (GetProjectsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetProjectsResult)(nil)).Elem()
}

func (o GetProjectsResultOutput) ToGetProjectsResultOutput() GetProjectsResultOutput {
	return o
}

func (o GetProjectsResultOutput) ToGetProjectsResultOutputWithContext(ctx context.Context) GetProjectsResultOutput {
	return o
}

func (o GetProjectsResultOutput) ToOutput(ctx context.Context) pulumix.Output[GetProjectsResult] {
	return pulumix.Output[GetProjectsResult]{
		OutputState: o.OutputState,
	}
}

// The provider-assigned unique ID for this managed resource.
func (o GetProjectsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetProjectsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetProjectsResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetProjectsResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o GetProjectsResultOutput) Projects() GetProjectsProjectArrayOutput {
	return o.ApplyT(func(v GetProjectsResult) []GetProjectsProject { return v.Projects }).(GetProjectsProjectArrayOutput)
}

func (o GetProjectsResultOutput) Public() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetProjectsResult) *bool { return v.Public }).(pulumi.BoolPtrOutput)
}

func (o GetProjectsResultOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetProjectsResult) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o GetProjectsResultOutput) VulnerabilityScanning() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetProjectsResult) *bool { return v.VulnerabilityScanning }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetProjectsResultOutput{})
}
