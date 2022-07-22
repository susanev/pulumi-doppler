// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package doppler

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ServiceToken struct {
	pulumi.CustomResourceState

	// The access level (read or read/write)
	Access pulumi.StringPtrOutput `pulumi:"access"`
	// The name of the Doppler config where the service token is located
	Config pulumi.StringOutput `pulumi:"config"`
	// The key for the Doppler service token
	Key pulumi.StringOutput `pulumi:"key"`
	// The name of the Doppler service token
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the Doppler project where the service token is located
	Project pulumi.StringOutput `pulumi:"project"`
}

// NewServiceToken registers a new resource with the given unique name, arguments, and options.
func NewServiceToken(ctx *pulumi.Context,
	name string, args *ServiceTokenArgs, opts ...pulumi.ResourceOption) (*ServiceToken, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Config == nil {
		return nil, errors.New("invalid value for required argument 'Config'")
	}
	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource ServiceToken
	err := ctx.RegisterResource("doppler:index/serviceToken:ServiceToken", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServiceToken gets an existing ServiceToken resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServiceToken(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceTokenState, opts ...pulumi.ResourceOption) (*ServiceToken, error) {
	var resource ServiceToken
	err := ctx.ReadResource("doppler:index/serviceToken:ServiceToken", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServiceToken resources.
type serviceTokenState struct {
	// The access level (read or read/write)
	Access *string `pulumi:"access"`
	// The name of the Doppler config where the service token is located
	Config *string `pulumi:"config"`
	// The key for the Doppler service token
	Key *string `pulumi:"key"`
	// The name of the Doppler service token
	Name *string `pulumi:"name"`
	// The name of the Doppler project where the service token is located
	Project *string `pulumi:"project"`
}

type ServiceTokenState struct {
	// The access level (read or read/write)
	Access pulumi.StringPtrInput
	// The name of the Doppler config where the service token is located
	Config pulumi.StringPtrInput
	// The key for the Doppler service token
	Key pulumi.StringPtrInput
	// The name of the Doppler service token
	Name pulumi.StringPtrInput
	// The name of the Doppler project where the service token is located
	Project pulumi.StringPtrInput
}

func (ServiceTokenState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceTokenState)(nil)).Elem()
}

type serviceTokenArgs struct {
	// The access level (read or read/write)
	Access *string `pulumi:"access"`
	// The name of the Doppler config where the service token is located
	Config string `pulumi:"config"`
	// The name of the Doppler service token
	Name *string `pulumi:"name"`
	// The name of the Doppler project where the service token is located
	Project string `pulumi:"project"`
}

// The set of arguments for constructing a ServiceToken resource.
type ServiceTokenArgs struct {
	// The access level (read or read/write)
	Access pulumi.StringPtrInput
	// The name of the Doppler config where the service token is located
	Config pulumi.StringInput
	// The name of the Doppler service token
	Name pulumi.StringPtrInput
	// The name of the Doppler project where the service token is located
	Project pulumi.StringInput
}

func (ServiceTokenArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceTokenArgs)(nil)).Elem()
}

type ServiceTokenInput interface {
	pulumi.Input

	ToServiceTokenOutput() ServiceTokenOutput
	ToServiceTokenOutputWithContext(ctx context.Context) ServiceTokenOutput
}

func (*ServiceToken) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceToken)(nil)).Elem()
}

func (i *ServiceToken) ToServiceTokenOutput() ServiceTokenOutput {
	return i.ToServiceTokenOutputWithContext(context.Background())
}

func (i *ServiceToken) ToServiceTokenOutputWithContext(ctx context.Context) ServiceTokenOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceTokenOutput)
}

// ServiceTokenArrayInput is an input type that accepts ServiceTokenArray and ServiceTokenArrayOutput values.
// You can construct a concrete instance of `ServiceTokenArrayInput` via:
//
//          ServiceTokenArray{ ServiceTokenArgs{...} }
type ServiceTokenArrayInput interface {
	pulumi.Input

	ToServiceTokenArrayOutput() ServiceTokenArrayOutput
	ToServiceTokenArrayOutputWithContext(context.Context) ServiceTokenArrayOutput
}

type ServiceTokenArray []ServiceTokenInput

func (ServiceTokenArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServiceToken)(nil)).Elem()
}

func (i ServiceTokenArray) ToServiceTokenArrayOutput() ServiceTokenArrayOutput {
	return i.ToServiceTokenArrayOutputWithContext(context.Background())
}

func (i ServiceTokenArray) ToServiceTokenArrayOutputWithContext(ctx context.Context) ServiceTokenArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceTokenArrayOutput)
}

// ServiceTokenMapInput is an input type that accepts ServiceTokenMap and ServiceTokenMapOutput values.
// You can construct a concrete instance of `ServiceTokenMapInput` via:
//
//          ServiceTokenMap{ "key": ServiceTokenArgs{...} }
type ServiceTokenMapInput interface {
	pulumi.Input

	ToServiceTokenMapOutput() ServiceTokenMapOutput
	ToServiceTokenMapOutputWithContext(context.Context) ServiceTokenMapOutput
}

type ServiceTokenMap map[string]ServiceTokenInput

func (ServiceTokenMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServiceToken)(nil)).Elem()
}

func (i ServiceTokenMap) ToServiceTokenMapOutput() ServiceTokenMapOutput {
	return i.ToServiceTokenMapOutputWithContext(context.Background())
}

func (i ServiceTokenMap) ToServiceTokenMapOutputWithContext(ctx context.Context) ServiceTokenMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceTokenMapOutput)
}

type ServiceTokenOutput struct{ *pulumi.OutputState }

func (ServiceTokenOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceToken)(nil)).Elem()
}

func (o ServiceTokenOutput) ToServiceTokenOutput() ServiceTokenOutput {
	return o
}

func (o ServiceTokenOutput) ToServiceTokenOutputWithContext(ctx context.Context) ServiceTokenOutput {
	return o
}

// The access level (read or read/write)
func (o ServiceTokenOutput) Access() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceToken) pulumi.StringPtrOutput { return v.Access }).(pulumi.StringPtrOutput)
}

// The name of the Doppler config where the service token is located
func (o ServiceTokenOutput) Config() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceToken) pulumi.StringOutput { return v.Config }).(pulumi.StringOutput)
}

// The key for the Doppler service token
func (o ServiceTokenOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceToken) pulumi.StringOutput { return v.Key }).(pulumi.StringOutput)
}

// The name of the Doppler service token
func (o ServiceTokenOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceToken) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The name of the Doppler project where the service token is located
func (o ServiceTokenOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceToken) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

type ServiceTokenArrayOutput struct{ *pulumi.OutputState }

func (ServiceTokenArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServiceToken)(nil)).Elem()
}

func (o ServiceTokenArrayOutput) ToServiceTokenArrayOutput() ServiceTokenArrayOutput {
	return o
}

func (o ServiceTokenArrayOutput) ToServiceTokenArrayOutputWithContext(ctx context.Context) ServiceTokenArrayOutput {
	return o
}

func (o ServiceTokenArrayOutput) Index(i pulumi.IntInput) ServiceTokenOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ServiceToken {
		return vs[0].([]*ServiceToken)[vs[1].(int)]
	}).(ServiceTokenOutput)
}

type ServiceTokenMapOutput struct{ *pulumi.OutputState }

func (ServiceTokenMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServiceToken)(nil)).Elem()
}

func (o ServiceTokenMapOutput) ToServiceTokenMapOutput() ServiceTokenMapOutput {
	return o
}

func (o ServiceTokenMapOutput) ToServiceTokenMapOutputWithContext(ctx context.Context) ServiceTokenMapOutput {
	return o
}

func (o ServiceTokenMapOutput) MapIndex(k pulumi.StringInput) ServiceTokenOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ServiceToken {
		return vs[0].(map[string]*ServiceToken)[vs[1].(string)]
	}).(ServiceTokenOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceTokenInput)(nil)).Elem(), &ServiceToken{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceTokenArrayInput)(nil)).Elem(), ServiceTokenArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceTokenMapInput)(nil)).Elem(), ServiceTokenMap{})
	pulumi.RegisterOutputType(ServiceTokenOutput{})
	pulumi.RegisterOutputType(ServiceTokenArrayOutput{})
	pulumi.RegisterOutputType(ServiceTokenMapOutput{})
}