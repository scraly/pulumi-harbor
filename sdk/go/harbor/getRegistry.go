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
func LookupRegistry(ctx *pulumi.Context, args *LookupRegistryArgs, opts ...pulumi.InvokeOption) (*LookupRegistryResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupRegistryResult
	err := ctx.Invoke("harbor:index/getRegistry:getRegistry", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRegistry.
type LookupRegistryArgs struct {
	Name string `pulumi:"name"`
}

// A collection of values returned by getRegistry.
type LookupRegistryResult struct {
	Description string `pulumi:"description"`
	// The provider-assigned unique ID for this managed resource.
	Id         string `pulumi:"id"`
	Insecure   string `pulumi:"insecure"`
	Name       string `pulumi:"name"`
	RegistryId int    `pulumi:"registryId"`
	Status     string `pulumi:"status"`
	Type       string `pulumi:"type"`
	Url        string `pulumi:"url"`
}

func LookupRegistryOutput(ctx *pulumi.Context, args LookupRegistryOutputArgs, opts ...pulumi.InvokeOption) LookupRegistryResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRegistryResult, error) {
			args := v.(LookupRegistryArgs)
			r, err := LookupRegistry(ctx, &args, opts...)
			var s LookupRegistryResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRegistryResultOutput)
}

// A collection of arguments for invoking getRegistry.
type LookupRegistryOutputArgs struct {
	Name pulumi.StringInput `pulumi:"name"`
}

func (LookupRegistryOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRegistryArgs)(nil)).Elem()
}

// A collection of values returned by getRegistry.
type LookupRegistryResultOutput struct{ *pulumi.OutputState }

func (LookupRegistryResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRegistryResult)(nil)).Elem()
}

func (o LookupRegistryResultOutput) ToLookupRegistryResultOutput() LookupRegistryResultOutput {
	return o
}

func (o LookupRegistryResultOutput) ToLookupRegistryResultOutputWithContext(ctx context.Context) LookupRegistryResultOutput {
	return o
}

func (o LookupRegistryResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupRegistryResult] {
	return pulumix.Output[LookupRegistryResult]{
		OutputState: o.OutputState,
	}
}

func (o LookupRegistryResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegistryResult) string { return v.Description }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupRegistryResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegistryResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupRegistryResultOutput) Insecure() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegistryResult) string { return v.Insecure }).(pulumi.StringOutput)
}

func (o LookupRegistryResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegistryResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupRegistryResultOutput) RegistryId() pulumi.IntOutput {
	return o.ApplyT(func(v LookupRegistryResult) int { return v.RegistryId }).(pulumi.IntOutput)
}

func (o LookupRegistryResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegistryResult) string { return v.Status }).(pulumi.StringOutput)
}

func (o LookupRegistryResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegistryResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupRegistryResultOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegistryResult) string { return v.Url }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRegistryResultOutput{})
}
