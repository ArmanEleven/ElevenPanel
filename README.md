# Eleven Panel

Eleven Panel is a next-generation management layer built around the Sanaei/3x-ui codebase.

## Project direction

Eleven Panel is **not** a rewrite of Xray management. The existing Sanaei/3x-ui capabilities remain the foundation. Eleven-specific capabilities are added as a separate layer so upstream functionality can be preserved and maintained.

### Planned layers

- Core: Sanaei/3x-ui Xray management
- Identity: admins, roles, permissions
- Service: users, groups, templates, subscriptions
- Infrastructure: nodes, health, monitoring
- Operations: statistics, logs, bulk actions
- Integration: stable API for Eleven Store
- Migration: Sanaei/3x-ui data migration and rollback support

## Status

Phase 1 — pinned Sanaei core and initial Eleven foundation.

The core is pinned to Sanaei/3x-ui v3.9.0 at commit `3cd4bf5`; see [docs/UPSTREAM.md](docs/UPSTREAM.md). The current Eleven layer contains initial domain models, a versioned health endpoint, and an adapter for basic client provisioning, updates, revocation, inbound lookup, and traffic reads.

This is not yet a complete Eleven Store integration. Persistent Eleven user/group/template/subscription workflows, authenticated business API endpoints, audit-event persistence, and node reconciliation remain future work. The existing Sanaei panel remains the operational core while those capabilities are built and tested.

## Upstream

- https://github.com/MHSanaei/3x-ui
- https://github.com/PasarGuard/panel

## Product name

**Eleven Panel**

Copyright and upstream attribution must be preserved where required by the upstream license.
