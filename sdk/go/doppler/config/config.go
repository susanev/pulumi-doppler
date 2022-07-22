// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package config

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

// A Doppler token, either a personal or service token
func GetDopplerToken(ctx *pulumi.Context) string {
	return config.Get(ctx, "doppler:dopplerToken")
}

// The Doppler API host (i.e. https://api.doppler.com)
func GetHost(ctx *pulumi.Context) string {
	return config.Get(ctx, "doppler:host")
}

// Whether or not to verify TLS
func GetVerifyTls(ctx *pulumi.Context) bool {
	return config.GetBool(ctx, "doppler:verifyTls")
}