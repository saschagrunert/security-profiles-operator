# Migration Guide: API Graduation to v1

Security Profiles Operator (SPO) 1.0.0 graduates all CRD APIs from alpha/beta to v1. This document covers what changed, what happens automatically, and what you should update. For general installation and usage instructions, see [installation-usage.md](../installation-usage.md).

## API version changes

All SPO custom resources move to `security-profiles-operator.x-k8s.io/v1`:

| CRD | Previous version(s) | New version |
|-----|---------------------|-------------|
| SeccompProfile | v1beta1 | v1 |
| SelinuxProfile | v1alpha2 | v1 |
| RawSelinuxProfile | v1alpha2 | v1 |
| AppArmorProfile | v1alpha1 | v1 |
| ProfileRecording | v1alpha1 | v1 |
| ProfileBinding | v1alpha1 | v1 |
| SecurityProfilesOperatorDaemon (SPOD) | v1alpha1 | v1 |
| SecurityProfileNodeStatus | v1alpha1 | v1 |

## Enum value changes

Several enum fields changed from uppercase/lowercase to PascalCase in v1:

| CRD | Field | Old value | New value |
|-----|-------|-----------|-----------|
| ProfileRecording | `spec.recorder` | `logs` | `Logs` |
| ProfileRecording | `spec.recorder` | `bpf` | `Bpf` |
| ProfileRecording | `spec.mergeStrategy` | `none` | `None` |
| ProfileRecording | `spec.mergeStrategy` | `containers` | `Containers` |
| SPOD | `status.state` | `PENDING` | `Pending` |
| SPOD | `status.state` | `CREATING` | `Creating` |
| SPOD | `status.state` | `UPDATING` | `Updating` |
| SPOD | `status.state` | `RUNNING` | `Running` |
| SPOD | `status.state` | `ERROR` | `Error` |
| SPOD | `spec.enricher.logEnricherSource` | `auditd` | `Auditd` |
| SPOD | `spec.enricher.logEnricherSource` | `bpf` | `Bpf` |

## Automatic conversion

Conversion webhooks handle translation between old and new API versions
transparently. This means:

- **Old manifests still work.** You can continue applying resources using
  `v1alpha1`, `v1alpha2`, or `v1beta1` `apiVersion` values. The conversion
  webhook translates them to v1 before storage.
- **Old API versions are still served.** `kubectl get` with an old API version
  returns the resource with old-style enum values, even though v1 is the
  storage version.
- **Existing resources in etcd are migrated on next write.** When the operator
  upgrades, resources stored under old versions are converted to v1 the next
  time they are updated.

No data is lost during conversion. All fields are preserved across versions.

## Recommended actions

While automatic conversion provides backward compatibility, we recommend
updating to v1:

1. **Update YAML manifests.** Change `apiVersion` from
   `security-profiles-operator.x-k8s.io/v1alpha1` (or `v1alpha2`, `v1beta1`)
   to `security-profiles-operator.x-k8s.io/v1`.

2. **Update enum values in manifests.** If you reference enum fields in
   manifests or scripts, update them to PascalCase (see table above).

3. **Update Go client imports.** If you consume the SPO API types in Go code,
   update imports from `api/*/v1alpha1` to `api/*/v1`. For example:
   ```go
   // Before
   import profilerecordingapi "sigs.k8s.io/security-profiles-operator/api/profilerecording/v1alpha1"

   // After
   import profilerecordingapi "sigs.k8s.io/security-profiles-operator/api/profilerecording/v1"
   ```

4. **Update scripts that parse enum strings.** If automation or monitoring
   checks for specific enum string values (e.g., checking SPOD status for
   `"RUNNING"`), update those checks to use PascalCase (`"Running"`).

## Deprecation timeline

Old API versions (`v1alpha1`, `v1alpha2`, `v1beta1`) remain served in 1.0.0
for backward compatibility. They will be removed in a future release, at least
one minor version after 1.0.0. Plan to migrate your manifests to v1 before
that removal.

## Examples

### ProfileRecording

Before (v1alpha1):
```yaml
apiVersion: security-profiles-operator.x-k8s.io/v1alpha1
kind: ProfileRecording
metadata:
  name: my-recording
spec:
  kind: SeccompProfile
  recorder: logs
  mergeStrategy: none
  podSelector:
    matchLabels:
      app: my-app
```

After (v1):
```yaml
apiVersion: security-profiles-operator.x-k8s.io/v1
kind: ProfileRecording
metadata:
  name: my-recording
spec:
  kind: SeccompProfile
  recorder: Logs
  mergeStrategy: None
  podSelector:
    matchLabels:
      app: my-app
```

### SeccompProfile

Before (v1beta1):
```yaml
apiVersion: security-profiles-operator.x-k8s.io/v1beta1
kind: SeccompProfile
metadata:
  name: my-profile
spec:
  defaultAction: SCMP_ACT_ERRNO
  syscalls:
  - names:
    - read
    - write
    action: SCMP_ACT_ALLOW
```

After (v1):
```yaml
apiVersion: security-profiles-operator.x-k8s.io/v1
kind: SeccompProfile
metadata:
  name: my-profile
spec:
  defaultAction: SCMP_ACT_ERRNO
  syscalls:
  - names:
    - read
    - write
    action: SCMP_ACT_ALLOW
```
