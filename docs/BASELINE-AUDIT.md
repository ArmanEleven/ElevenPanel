# Baseline implementation audit

Audit date: 2026-10-09
Branch reviewed: `feat/eleven-foundation`
Pinned core: Sanaei/3x-ui v3.9.0 (see [UPSTREAM.md](UPSTREAM.md))

This is a source-reading audit, not a claim that every path has been exercised in a live deployment.

## Verified repository facts

| Area | Current evidence | Assessment |
|---|---|---|
| Upstream baseline | The repository pins Sanaei/3x-ui v3.9.0 and vendors its Go + React/TypeScript application. | Strong starting point for protocol and runtime compatibility. |
| CI baseline | Latest workflow runs 121 and 122 on the same head commit completed successfully. | Build checks are green; this does not certify product completeness or production safety. |
| Eleven API | `internal/web/controller/eleven.go` registers only `GET /api/v1/eleven/health`; `internal/eleven/api/v1/handler.go` is also health-only. | No customer, subscription, commerce, reseller, or node-management API is implemented yet. |
| Eleven persistence | `internal/eleven/identity/model.go`, `service/model.go`, and `nodes/model.go` currently define Go structs. They are not included in `internal/database/db.go:allModels()`. | These types are not yet durable business entities; schema + migrations and database tests are required before they can power real workflows. |
| Core adapter | `internal/eleven/adapter/sanaei.go` wraps existing Sanaei client/inbound services for create/update/revoke/traffic/inbound reads. | Useful integration seam; it should be verified against multi-node runtime dispatch before relying on it for production writes. |
| Reconciliation | `SanaeiAdapter.Reconcile()` intentionally returns an explicit not-implemented error; its regression test pins that behavior. | Node synchronization/reconciliation is not implemented. |
| Identity/RBAC | `internal/eleven/identity/permissions.go` declares default permission sets, but the inspected identity models are not persisted and no Eleven authorization flow is wired into the API. | Roles are design scaffolding, not enforced reseller isolation. Do not expose privileged endpoints until authorization is implemented and negatively tested. |
| Commerce | No Eleven order, invoice, ledger, balance, or payment-adapter domain appears in the current Eleven packages. | Sales/accounting remain unimplemented. |
| UI/branding | The existing frontend is embedded by `internal/web/web.go`; product roadmap still lists rebranding, language switch, and unified UX as unfinished. | The current work is still primarily foundation work, not the finished Eleven-branded experience. |
| Installation/recovery | Existing upstream deployment assets exist; Eleven's script + Docker parity, verified backup, safe update, and rollback are roadmap items. | Must be validated before a public release. |

## Recommended next implementation order

1. **Persistent Eleven schema and migration tests.** Add Eleven-owned tables through the repository's database migration mechanism, preserving existing Sanaei tables and SQLite/PostgreSQL behavior.
2. **Authentication and ownership model.** Define administrator, reseller, and customer boundaries. Add deny-by-default authorization tests before any business write endpoint.
3. **Service-management workflow.** Build a durable customer/subscription model and connect provisioning, renewal, suspension, and revocation through the core runtime path with idempotency.
4. **Central node operations.** Reuse the upstream node/runtime dispatch; add node status and offline/partial-failure behavior without a second dispatcher.
5. **Commerce ledger.** Implement orders, invoices, balance movements, and payment-provider adapter only after identity and subscription ownership are enforced.
6. **Versioned Eleven API + Eleven Store contract.** Document endpoints and generated artifacts; test authentication, duplicate requests, and error contracts.
7. **Install/update/restore.** Validate both install paths, backups, artifact checks, and rollback on clean and existing installations.
8. **Rebrand and redesign.** Keep core workflows and all supported protocol controls intact while introducing Persian/English Eleven UI.

## Release blockers

- No production reseller/sales release before durable persistence, authorization isolation, audit history, and duplicate-safe financial transactions exist.
- No multi-node release before state-changing operations are shown to use the existing runtime dispatch path and offline-node tests pass.
- No production update path before backup restoration and rollback have been exercised.
- A green CI run alone is not a release candidate.
