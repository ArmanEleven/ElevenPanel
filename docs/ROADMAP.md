# Eleven Panel Roadmap

This roadmap follows the agreed product specification in [PRODUCT-SPEC.md](PRODUCT-SPEC.md).
CI is necessary but not sufficient: each phase needs functional acceptance criteria, compatibility checks, and a safe upgrade/rollback story.

## Phase 0 — Baseline and architecture audit
- [x] Repository created and Sanaei/3x-ui v3.9.0 baseline pinned
- [x] Product direction and integration boundary documented
- [x] Product requirements, architecture recommendation, delivery sequence, and open decisions documented
- [x] Build an initial capability matrix for Sanaei/3x-ui versus PasarGuard (feature parity still requires source-level verification)
- [ ] Map frontend routes, API/auth boundaries, database models, node dispatch, and installer/update behavior
- [ ] Identify high-risk compatibility/security issues and duplicate implementations
- [x] Establish a green CI baseline on the current branch (functional and production acceptance gates remain open)

## Phase 1 — Eleven identity and foundation
- [ ] Rebrand user-facing application, installer, service labels, and documentation to Eleven Panel
- [ ] Add Persian/English language switching using the existing localization architecture
- [ ] Define centralized application version/build metadata
- [ ] Implement Eleven identity, role, permission, and ownership model
- [ ] Add durable append-only audit events for sensitive operations
- [ ] Keep existing Sanaei login and operational workflows working during transition

## Phase 2 — Service management
- [ ] Eleven customers/users and ownership boundaries
- [ ] Groups and reusable service templates
- [ ] Subscription lifecycle: provision, renew, suspend, revoke
- [ ] Traffic/time quota policy and usage reads
- [ ] Optional HWID/device policy
- [ ] Idempotent service operations with regression tests

## Phase 3 — Central multi-node operations
- [ ] Central node registry and credentials handling
- [ ] Node health/status and last-seen information
- [ ] Node groups and service assignment
- [ ] Dispatch stateful operations through the established runtime path
- [ ] Offline-node handling, retries, and partial-failure visibility
- [ ] Multi-node integration tests

## Phase 4 — Sales, resellers, and accounting
- [ ] Reseller accounts and tenant/ownership isolation
- [ ] Price lists and reseller pricing rules
- [ ] Orders and invoices
- [ ] Transaction ledger and reseller balances
- [ ] Payment adapter with duplicate-safe callbacks
- [ ] Renewal and adjustment workflows with audit history
- [ ] End-customer portal decision implemented

## Phase 5 — Eleven integration API
- [ ] Authenticated versioned API and API-key lifecycle
- [ ] Provision/update/renew/revoke endpoints
- [ ] Idempotency and stable machine-readable errors
- [ ] API docs and generated-artifact chain
- [ ] Webhook/event contract
- [ ] Eleven Store integration tests

## Phase 6 — Install, update, migration, and recovery
- [ ] Guided installation script
- [ ] Docker deployment
- [ ] Pre-update backup and backup verification
- [ ] Manual updates and opt-in automatic updates
- [ ] Atomic artifact validation and rollback to prior version
- [ ] Sanaei-to-Eleven migration dry-run and report
- [ ] Restore, upgrade, and rollback tests on representative existing databases

## Phase 7 — Unified UI redesign
- [ ] Information architecture and navigation for admin/reseller/customer roles
- [ ] Responsive Eleven visual system and branding
- [ ] Persian/English translation and language switch
- [ ] Rebuild core workflows without losing protocol-specific controls
- [ ] Accessibility, empty/error/loading states, and frontend regression coverage

## Phase 8 — Production readiness and release
- [ ] Security review of authentication, authorization, sessions, API, subscriptions, and Xray-facing operations
- [ ] Compatibility suite for existing protocols, links, and subscription outputs
- [ ] Load/concurrency testing and idempotency testing
- [ ] Dependency and supply-chain review
- [ ] Install/upgrade/rollback rehearsal on clean and existing systems
- [ ] Operator documentation, troubleshooting, and release notes
- [ ] Release candidate and staged rollout

## Release gate
Do not merge or publish a production release solely because CI is green. A release must also pass functional acceptance, data-safety, security, protocol-compatibility, migration, and rollback checks.
