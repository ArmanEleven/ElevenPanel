# Eleven Panel — Product Specification

Status: agreed product direction; implementation details remain subject to repository review.

## Product definition

Eleven Panel is a standalone, branded panel that combines the operational maturity and protocol coverage of the Sanaei/3x-ui core with useful service-management and reseller capabilities from PasarGuard. It must feel like one coherent product, not two panels linked together.

The user-facing product name is **Eleven Panel**. Preserve all required upstream copyright notices, license files, and attribution.

## Product requirements

### 1. Core and compatibility
- Keep Sanaei/3x-ui as the primary Xray/inbound/client runtime and configuration source of truth.
- Preserve all supported upstream protocols and existing core behavior unless a deliberate compatibility change is documented.
- Prefer existing Sanaei services and runtime abstractions over duplicated Xray or client logic.
- Isolate Eleven-specific domain logic behind explicit packages and interfaces.
- Track upstream versions explicitly; upgrades require compatibility review, regression checks, migration notes, and a rollback plan.

### 2. Unified experience and identity
- Deliver a distinct Eleven Panel brand and UI, with Persian and English locales and an in-product language switch.
- Implement capabilities first; then complete the visual redesign without regressing core workflows.
- Keep the existing Sanaei operational functionality available through the unified experience while Eleven workflows are introduced incrementally.
- Use accessible, responsive layouts and consistent navigation, terminology, validation, and error handling.

### 3. Central multi-node management
- Support one central control panel managing multiple nodes.
- Reuse and extend the upstream node/runtime architecture rather than introducing a competing node-dispatch path.
- Provide node registration, health/status, grouping, assignment, safe retries, and clear handling for offline nodes.
- Stateful client/inbound changes must use the established runtime dispatch path.

### 4. Roles and tenancy boundaries
- Support a platform administrator, reseller/agent, and end customer.
- Define role-based permissions and ownership boundaries before enabling reseller operations.
- Resellers must only see and manage their own customers, subscriptions, balances, and permitted nodes/resources.
- Record security-sensitive and financial actions in an append-only audit trail.

### 5. Service and subscription management
- Support users, groups, reusable service templates, subscriptions, traffic and time limits, renewal, suspension/revocation, and optional HWID/device policy.
- Make provisioning and renewal idempotent so retries cannot create duplicate services or charge twice.
- Keep Eleven metadata in separate `eleven_*` tables unless a documented compatibility requirement proves otherwise.
- Preserve original Sanaei data and support a dry-run migration before any live cutover.

### 6. Sales and accounting
- Plan for internal sales, renewals, reseller balances, invoices/transactions, adjustments, and financial audit history.
- Keep payment-provider integration behind an adapter; do not hard-code a single gateway into subscription logic.
- Financial state changes must be transactional, auditable, and safe against duplicate callbacks/requests.
- Payment methods, settlement currency, and whether end customers get a separate portal are open product decisions listed below.

### 7. Integration API
- Provide a versioned Eleven API separate from existing Sanaei endpoints.
- Require authentication and authorization for every write; support key rotation, rate limits where appropriate, idempotency, structured errors, and audit events.
- Document endpoints through the repository's existing API-doc generation chain.
- Keep Eleven Store dependent on the public Eleven API contract rather than internal database models.

### 8. Installation, upgrades, and recovery
- Support both a guided install script and Docker deployment.
- Provide manual updates and an optional controlled automatic-update path.
- Before an update or migration, verify a restorable backup; preserve the previous working version and support rollback.
- Never replace a working binary/database before the new artifact has been downloaded and validated.
- Document supported operating systems, architectures, ports, backup contents, restore procedure, and recovery behavior.

## Recommended implementation architecture

Use a **modular monolith initially**: one deployable panel with clearly separated Eleven domains and the existing Sanaei core. This avoids premature microservices and minimizes operational burden while keeping domain boundaries testable.

- **Core adapter:** thin interfaces around existing Sanaei services/runtime.
- **Identity and authorization:** administrator/reseller/customer identities, roles, permissions, and ownership scopes.
- **Service domain:** customers, groups, templates, subscriptions, renewals, and quota policy.
- **Node domain:** registry, health, grouping, and assignment using upstream dispatch.
- **Commerce domain:** orders, invoices, ledger entries, reseller balances, and payment adapters.
- **Integration domain:** versioned API, API credentials, idempotency records, webhooks, and stable error contracts.
- **Operations domain:** audit events, health, backup/restore, migration, and release/update orchestration.
- **Frontend:** existing React/TypeScript + Ant Design stack, first completing product capabilities and then applying the Eleven visual identity and Persian/English localization.

Avoid a second database schema for Xray state, a parallel node dispatcher, or duplicated protocol serializers.

## Delivery sequence and release gates

1. **Baseline and architecture audit:** map existing Sanaei/PasarGuard capabilities, frontend routes, data models, node dispatch, API authentication, and deployment/update scripts. Produce a capability matrix and risk list.
2. **Product identity and access foundations:** Eleven branding, localization foundation, identity/role model, authorization boundaries, and durable audit events.
3. **Service management:** customers, groups, templates, subscriptions, quotas, renewals, and revocation, wired to the existing core adapter.
4. **Multi-node operations:** central node registry/health, assignment, safe dispatch, failure handling, and observability.
5. **Commerce and reseller system:** orders, invoices, ledger, reseller balances, payment adapter, and duplicate-safe transaction processing.
6. **Public integration:** authenticated/versioned Eleven API, API docs, idempotency, rate limits, and Eleven Store integration contract.
7. **Install and lifecycle:** installer and Docker, verified backups, upgrade checks, rollback, migration dry-run, and recovery tests.
8. **Unified redesign and hardening:** complete responsive UI redesign, accessibility, security review, compatibility tests, load tests, and release packaging.

A phase is not complete merely because CI passes. It must meet functional acceptance criteria and regression/security checks, and must not silently break the upstream panel's existing behavior.

## Open product decisions

These decisions do not block the architecture audit, but should be settled before implementing the corresponding feature:
1. **Payments:** should the first release support manual card-to-card/receipt approval, an online gateway, or both? Which gateway(s) are required?
2. **Customer experience:** should end customers sign in to a self-service portal, or should their accounts and subscriptions be managed only by admins/resellers at first?
3. **Currency and accounting:** which currency is the canonical ledger currency, and should the UI support more than one?
4. **Initial deployment target:** which Ubuntu/Debian versions and CPU architectures must be supported in the first release?
5. **Reseller pricing:** one global price list with reseller discounts, or per-reseller custom pricing?

Use sensible defaults for non-blocking details and keep provider-specific integrations replaceable.
