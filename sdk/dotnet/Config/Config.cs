// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Immutable;

namespace Pulumiverse.Doppler
{
    public static class Config
    {
        [System.Diagnostics.CodeAnalysis.SuppressMessage("Microsoft.Design", "IDE1006", Justification = 
        "Double underscore prefix used to avoid conflicts with variable names.")]
        private sealed class __Value<T>
        {
            private readonly Func<T> _getter;
            private T _value = default!;
            private bool _set;

            public __Value(Func<T> getter)
            {
                _getter = getter;
            }

            public T Get() => _set ? _value : _getter();

            public void Set(T value)
            {
                _value = value;
                _set = true;
            }
        }

        private static readonly Pulumi.Config __config = new Pulumi.Config("doppler");

        private static readonly __Value<string?> _dopplerToken = new __Value<string?>(() => __config.Get("dopplerToken"));
        /// <summary>
        /// A Doppler token, either a personal or service token
        /// </summary>
        public static string? DopplerToken
        {
            get => _dopplerToken.Get();
            set => _dopplerToken.Set(value);
        }

        private static readonly __Value<string?> _host = new __Value<string?>(() => __config.Get("host"));
        /// <summary>
        /// The Doppler API host (i.e. https://api.doppler.com)
        /// </summary>
        public static string? Host
        {
            get => _host.Get();
            set => _host.Set(value);
        }

        private static readonly __Value<bool?> _verifyTls = new __Value<bool?>(() => __config.GetBoolean("verifyTls"));
        /// <summary>
        /// Whether or not to verify TLS
        /// </summary>
        public static bool? VerifyTls
        {
            get => _verifyTls.Get();
            set => _verifyTls.Set(value);
        }

    }
}