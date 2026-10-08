# Upstream Pin

Eleven Panel is built around Sanaei/3x-ui.

## Pinned baseline

- Upstream repository: MHSanaei/3x-ui
- Release: v3.9.0
- Commit: 3cd4bf5
- Branch policy: stable release only; development builds are not used as the product baseline.

The v3.9.0 release includes the current upstream multi-node, managed-host, subscription, notification and Xray-core integration capabilities. Eleven features are added as a separate product layer rather than replacing those capabilities.

## Integration rule

The upstream core remains attributable to its original project and license. Eleven-specific code must be clearly separated and documented.

## Stability rule

Do not silently rebase the Eleven product onto a newer upstream release. Each upstream upgrade must be an explicit compatibility task with migration notes and regression tests.
