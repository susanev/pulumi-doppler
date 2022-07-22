// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Doppler
{
    public static class Secrets
    {
        public static Task<SecretsResult> InvokeAsync(SecretsArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<SecretsResult>("doppler:index/secrets:Secrets", args ?? new SecretsArgs(), options.WithDefaults());

        public static Output<SecretsResult> Invoke(SecretsInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<SecretsResult>("doppler:index/secrets:Secrets", args ?? new SecretsInvokeArgs(), options.WithDefaults());
    }


    public sealed class SecretsArgs : Pulumi.InvokeArgs
    {
        [Input("config")]
        public string? Config { get; set; }

        [Input("project")]
        public string? Project { get; set; }

        public SecretsArgs()
        {
        }
    }

    public sealed class SecretsInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("config")]
        public Input<string>? Config { get; set; }

        [Input("project")]
        public Input<string>? Project { get; set; }

        public SecretsInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class SecretsResult
    {
        public readonly string? Config;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableDictionary<string, string> Map;
        public readonly string? Project;

        [OutputConstructor]
        private SecretsResult(
            string? config,

            string id,

            ImmutableDictionary<string, string> map,

            string? project)
        {
            Config = config;
            Id = id;
            Map = map;
            Project = project;
        }
    }
}