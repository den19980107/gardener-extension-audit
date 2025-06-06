//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
2025 Copyright metal-stack.
*/

// Code generated by conversion-gen. DO NOT EDIT.

package v1alpha1

import (
	unsafe "unsafe"

	audit "github.com/metal-stack/gardener-extension-audit/pkg/apis/audit"
	resource "k8s.io/apimachinery/pkg/api/resource"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*AuditBackendClusterForwarding)(nil), (*audit.AuditBackendClusterForwarding)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_AuditBackendClusterForwarding_To_audit_AuditBackendClusterForwarding(a.(*AuditBackendClusterForwarding), b.(*audit.AuditBackendClusterForwarding), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*audit.AuditBackendClusterForwarding)(nil), (*AuditBackendClusterForwarding)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_audit_AuditBackendClusterForwarding_To_v1alpha1_AuditBackendClusterForwarding(a.(*audit.AuditBackendClusterForwarding), b.(*AuditBackendClusterForwarding), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*AuditBackendLog)(nil), (*audit.AuditBackendLog)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_AuditBackendLog_To_audit_AuditBackendLog(a.(*AuditBackendLog), b.(*audit.AuditBackendLog), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*audit.AuditBackendLog)(nil), (*AuditBackendLog)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_audit_AuditBackendLog_To_v1alpha1_AuditBackendLog(a.(*audit.AuditBackendLog), b.(*AuditBackendLog), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*AuditBackendSplunk)(nil), (*audit.AuditBackendSplunk)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_AuditBackendSplunk_To_audit_AuditBackendSplunk(a.(*AuditBackendSplunk), b.(*audit.AuditBackendSplunk), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*audit.AuditBackendSplunk)(nil), (*AuditBackendSplunk)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_audit_AuditBackendSplunk_To_v1alpha1_AuditBackendSplunk(a.(*audit.AuditBackendSplunk), b.(*AuditBackendSplunk), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*AuditBackends)(nil), (*audit.AuditBackends)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_AuditBackends_To_audit_AuditBackends(a.(*AuditBackends), b.(*audit.AuditBackends), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*audit.AuditBackends)(nil), (*AuditBackends)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_audit_AuditBackends_To_v1alpha1_AuditBackends(a.(*audit.AuditBackends), b.(*AuditBackends), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*AuditConfig)(nil), (*audit.AuditConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_AuditConfig_To_audit_AuditConfig(a.(*AuditConfig), b.(*audit.AuditConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*audit.AuditConfig)(nil), (*AuditConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_audit_AuditConfig_To_v1alpha1_AuditConfig(a.(*audit.AuditConfig), b.(*AuditConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*AuditPersistence)(nil), (*audit.AuditPersistence)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_AuditPersistence_To_audit_AuditPersistence(a.(*AuditPersistence), b.(*audit.AuditPersistence), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*audit.AuditPersistence)(nil), (*AuditPersistence)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_audit_AuditPersistence_To_v1alpha1_AuditPersistence(a.(*audit.AuditPersistence), b.(*AuditPersistence), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1alpha1_AuditBackendClusterForwarding_To_audit_AuditBackendClusterForwarding(in *AuditBackendClusterForwarding, out *audit.AuditBackendClusterForwarding, s conversion.Scope) error {
	out.Enabled = in.Enabled
	out.FilesystemBufferSize = (*string)(unsafe.Pointer(in.FilesystemBufferSize))
	return nil
}

// Convert_v1alpha1_AuditBackendClusterForwarding_To_audit_AuditBackendClusterForwarding is an autogenerated conversion function.
func Convert_v1alpha1_AuditBackendClusterForwarding_To_audit_AuditBackendClusterForwarding(in *AuditBackendClusterForwarding, out *audit.AuditBackendClusterForwarding, s conversion.Scope) error {
	return autoConvert_v1alpha1_AuditBackendClusterForwarding_To_audit_AuditBackendClusterForwarding(in, out, s)
}

func autoConvert_audit_AuditBackendClusterForwarding_To_v1alpha1_AuditBackendClusterForwarding(in *audit.AuditBackendClusterForwarding, out *AuditBackendClusterForwarding, s conversion.Scope) error {
	out.Enabled = in.Enabled
	out.FilesystemBufferSize = (*string)(unsafe.Pointer(in.FilesystemBufferSize))
	return nil
}

// Convert_audit_AuditBackendClusterForwarding_To_v1alpha1_AuditBackendClusterForwarding is an autogenerated conversion function.
func Convert_audit_AuditBackendClusterForwarding_To_v1alpha1_AuditBackendClusterForwarding(in *audit.AuditBackendClusterForwarding, out *AuditBackendClusterForwarding, s conversion.Scope) error {
	return autoConvert_audit_AuditBackendClusterForwarding_To_v1alpha1_AuditBackendClusterForwarding(in, out, s)
}

func autoConvert_v1alpha1_AuditBackendLog_To_audit_AuditBackendLog(in *AuditBackendLog, out *audit.AuditBackendLog, s conversion.Scope) error {
	out.Enabled = in.Enabled
	return nil
}

// Convert_v1alpha1_AuditBackendLog_To_audit_AuditBackendLog is an autogenerated conversion function.
func Convert_v1alpha1_AuditBackendLog_To_audit_AuditBackendLog(in *AuditBackendLog, out *audit.AuditBackendLog, s conversion.Scope) error {
	return autoConvert_v1alpha1_AuditBackendLog_To_audit_AuditBackendLog(in, out, s)
}

func autoConvert_audit_AuditBackendLog_To_v1alpha1_AuditBackendLog(in *audit.AuditBackendLog, out *AuditBackendLog, s conversion.Scope) error {
	out.Enabled = in.Enabled
	return nil
}

// Convert_audit_AuditBackendLog_To_v1alpha1_AuditBackendLog is an autogenerated conversion function.
func Convert_audit_AuditBackendLog_To_v1alpha1_AuditBackendLog(in *audit.AuditBackendLog, out *AuditBackendLog, s conversion.Scope) error {
	return autoConvert_audit_AuditBackendLog_To_v1alpha1_AuditBackendLog(in, out, s)
}

func autoConvert_v1alpha1_AuditBackendSplunk_To_audit_AuditBackendSplunk(in *AuditBackendSplunk, out *audit.AuditBackendSplunk, s conversion.Scope) error {
	out.Enabled = in.Enabled
	out.FilesystemBufferSize = (*string)(unsafe.Pointer(in.FilesystemBufferSize))
	out.Index = in.Index
	out.Host = in.Host
	out.Port = in.Port
	out.SecretResourceName = in.SecretResourceName
	out.TlsEnabled = in.TlsEnabled
	out.TlsHost = in.TlsHost
	out.CustomData = *(*map[string]string)(unsafe.Pointer(&in.CustomData))
	return nil
}

// Convert_v1alpha1_AuditBackendSplunk_To_audit_AuditBackendSplunk is an autogenerated conversion function.
func Convert_v1alpha1_AuditBackendSplunk_To_audit_AuditBackendSplunk(in *AuditBackendSplunk, out *audit.AuditBackendSplunk, s conversion.Scope) error {
	return autoConvert_v1alpha1_AuditBackendSplunk_To_audit_AuditBackendSplunk(in, out, s)
}

func autoConvert_audit_AuditBackendSplunk_To_v1alpha1_AuditBackendSplunk(in *audit.AuditBackendSplunk, out *AuditBackendSplunk, s conversion.Scope) error {
	out.Enabled = in.Enabled
	out.FilesystemBufferSize = (*string)(unsafe.Pointer(in.FilesystemBufferSize))
	out.Index = in.Index
	out.Host = in.Host
	out.Port = in.Port
	out.SecretResourceName = in.SecretResourceName
	out.TlsEnabled = in.TlsEnabled
	out.TlsHost = in.TlsHost
	out.CustomData = *(*map[string]string)(unsafe.Pointer(&in.CustomData))
	return nil
}

// Convert_audit_AuditBackendSplunk_To_v1alpha1_AuditBackendSplunk is an autogenerated conversion function.
func Convert_audit_AuditBackendSplunk_To_v1alpha1_AuditBackendSplunk(in *audit.AuditBackendSplunk, out *AuditBackendSplunk, s conversion.Scope) error {
	return autoConvert_audit_AuditBackendSplunk_To_v1alpha1_AuditBackendSplunk(in, out, s)
}

func autoConvert_v1alpha1_AuditBackends_To_audit_AuditBackends(in *AuditBackends, out *audit.AuditBackends, s conversion.Scope) error {
	out.Log = (*audit.AuditBackendLog)(unsafe.Pointer(in.Log))
	out.ClusterForwarding = (*audit.AuditBackendClusterForwarding)(unsafe.Pointer(in.ClusterForwarding))
	out.Splunk = (*audit.AuditBackendSplunk)(unsafe.Pointer(in.Splunk))
	return nil
}

// Convert_v1alpha1_AuditBackends_To_audit_AuditBackends is an autogenerated conversion function.
func Convert_v1alpha1_AuditBackends_To_audit_AuditBackends(in *AuditBackends, out *audit.AuditBackends, s conversion.Scope) error {
	return autoConvert_v1alpha1_AuditBackends_To_audit_AuditBackends(in, out, s)
}

func autoConvert_audit_AuditBackends_To_v1alpha1_AuditBackends(in *audit.AuditBackends, out *AuditBackends, s conversion.Scope) error {
	out.Log = (*AuditBackendLog)(unsafe.Pointer(in.Log))
	out.ClusterForwarding = (*AuditBackendClusterForwarding)(unsafe.Pointer(in.ClusterForwarding))
	out.Splunk = (*AuditBackendSplunk)(unsafe.Pointer(in.Splunk))
	return nil
}

// Convert_audit_AuditBackends_To_v1alpha1_AuditBackends is an autogenerated conversion function.
func Convert_audit_AuditBackends_To_v1alpha1_AuditBackends(in *audit.AuditBackends, out *AuditBackends, s conversion.Scope) error {
	return autoConvert_audit_AuditBackends_To_v1alpha1_AuditBackends(in, out, s)
}

func autoConvert_v1alpha1_AuditConfig_To_audit_AuditConfig(in *AuditConfig, out *audit.AuditConfig, s conversion.Scope) error {
	if err := Convert_v1alpha1_AuditPersistence_To_audit_AuditPersistence(&in.Persistence, &out.Persistence, s); err != nil {
		return err
	}
	out.Replicas = (*int32)(unsafe.Pointer(in.Replicas))
	out.WebhookMode = audit.AuditWebhookMode(in.WebhookMode)
	out.Backends = (*audit.AuditBackends)(unsafe.Pointer(in.Backends))
	return nil
}

// Convert_v1alpha1_AuditConfig_To_audit_AuditConfig is an autogenerated conversion function.
func Convert_v1alpha1_AuditConfig_To_audit_AuditConfig(in *AuditConfig, out *audit.AuditConfig, s conversion.Scope) error {
	return autoConvert_v1alpha1_AuditConfig_To_audit_AuditConfig(in, out, s)
}

func autoConvert_audit_AuditConfig_To_v1alpha1_AuditConfig(in *audit.AuditConfig, out *AuditConfig, s conversion.Scope) error {
	if err := Convert_audit_AuditPersistence_To_v1alpha1_AuditPersistence(&in.Persistence, &out.Persistence, s); err != nil {
		return err
	}
	out.Replicas = (*int32)(unsafe.Pointer(in.Replicas))
	out.WebhookMode = AuditWebhookMode(in.WebhookMode)
	out.Backends = (*AuditBackends)(unsafe.Pointer(in.Backends))
	return nil
}

// Convert_audit_AuditConfig_To_v1alpha1_AuditConfig is an autogenerated conversion function.
func Convert_audit_AuditConfig_To_v1alpha1_AuditConfig(in *audit.AuditConfig, out *AuditConfig, s conversion.Scope) error {
	return autoConvert_audit_AuditConfig_To_v1alpha1_AuditConfig(in, out, s)
}

func autoConvert_v1alpha1_AuditPersistence_To_audit_AuditPersistence(in *AuditPersistence, out *audit.AuditPersistence, s conversion.Scope) error {
	out.Size = (*resource.Quantity)(unsafe.Pointer(in.Size))
	out.StorageClassName = (*string)(unsafe.Pointer(in.StorageClassName))
	return nil
}

// Convert_v1alpha1_AuditPersistence_To_audit_AuditPersistence is an autogenerated conversion function.
func Convert_v1alpha1_AuditPersistence_To_audit_AuditPersistence(in *AuditPersistence, out *audit.AuditPersistence, s conversion.Scope) error {
	return autoConvert_v1alpha1_AuditPersistence_To_audit_AuditPersistence(in, out, s)
}

func autoConvert_audit_AuditPersistence_To_v1alpha1_AuditPersistence(in *audit.AuditPersistence, out *AuditPersistence, s conversion.Scope) error {
	out.Size = (*resource.Quantity)(unsafe.Pointer(in.Size))
	out.StorageClassName = (*string)(unsafe.Pointer(in.StorageClassName))
	return nil
}

// Convert_audit_AuditPersistence_To_v1alpha1_AuditPersistence is an autogenerated conversion function.
func Convert_audit_AuditPersistence_To_v1alpha1_AuditPersistence(in *audit.AuditPersistence, out *AuditPersistence, s conversion.Scope) error {
	return autoConvert_audit_AuditPersistence_To_v1alpha1_AuditPersistence(in, out, s)
}
