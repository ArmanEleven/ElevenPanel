# Capability Matrix — Eleven Panel

This is the initial audit matrix, not a claim that every feature is already implemented in Eleven Panel. PasarGuard's public feature list and the pinned Sanaei/3x-ui codebase are inputs; each row must be verified against source, tests, and real behavior before it is marked complete.

| Capability | Sanaei/3x-ui baseline | PasarGuard feature set | Eleven Panel decision | Verification needed |
|---|---|---|---|---|
| Xray configuration and protocol controls | Existing core; keep as source of truth | Xray plus additional protocol coverage | Preserve all existing supported core protocols and settings | Regression tests for emitted configs, links, and subscriptions |
| Client/inbound lifecycle | Existing services/controllers | User-centric multi-inbound management | Keep core writes behind existing service/runtime abstractions | Exercise create/update/revoke and restart behavior |
| Multi-node control | Existing node/runtime mechanisms | Multi-node management and node worker | One central panel, reusing upstream dispatch wherever possible | Trace every state-changing operation; offline-node tests |
| Admin roles and permissions | Existing panel auth/admin capabilities; exact granularity needs audit | Multi-admin RBAC and scoped access | Add explicit Eleven roles and ownership boundaries | Authorization matrix and negative tests |
| Resellers/agents | Not a complete commerce/reseller domain in the baseline | Scoped multi-admin/RBAC is advertised; commercial workflows need source audit | First-class reseller accounts with customer isolation | Prevent cross-reseller reads/writes; ledger invariants |
| Customer identity | Core clients are protocol-oriented | Multi-user account management | Separate business customer identity from protocol client records | Mapping, duplicate identities, deletion/retention rules |
| Groups and templates | Existing group-related functionality; coverage to audit | User/group policy patterns | Reusable service templates and group defaults | Template changes must not unexpectedly mutate existing subscriptions |
| Traffic/time quotas | Existing client/inbound traffic and expiry controls | Traffic and expiry limits plus periodic quota modes | Unified quota model, preserving native core semantics | Time zones, reset boundaries, disabled/expired clients |
| Periodic quota resets | Exact native coverage to verify | Daily/weekly and other periodic limits advertised | Implement only after source/behavior audit; keep reset jobs idempotent | Restart and duplicate-job tests |
| HWID/device limits | Existing HAPP/HWID-related code exists; exact parity to audit | HWID/device limits advertised | Add a policy abstraction without duplicating existing enforcement | Device identity changes, limits, bypass attempts |
| Multiple protocols per business user | A client may have protocol-specific core records | Multi-protocol per user advertised | One Eleven customer/subscription can map to multiple supported core records | Partial provision failure and rollback |
| WireGuard / Hysteria2 | WireGuard-related AmneziaWG and other sidecars exist; Hysteria2 parity needs verification | WireGuard and Hysteria2 advertised | Preserve core protocols and evaluate unsupported PasarGuard-specific protocol gaps individually | Binary packaging, lifecycle, subscription/link output |
| Subscriptions and share links | Existing share/subscription output and client compatibility | V2ray/Clash/ClashMeta subscription compatibility | Reuse upstream serializers and formats; don't fork formatting logic | Golden/compatibility tests against real client parsers |
| QR codes and sharing | Existing link/QR capabilities in the core | QR and share links advertised | Retain existing functionality in the unified UX | Scan/link regression tests |
| Telegram bot | Existing ecosystem integration needs source audit | Integrated Telegram bot advertised | Add only after core service/API contracts are stable; avoid duplicate business logic | Auth, command permissions, duplicate callbacks |
| REST API | Existing Sanaei APIs | REST API advertised | Versioned Eleven API, separate from legacy endpoints | Auth, rate limits, idempotency, API docs |
| Statistics and monitoring | Existing dashboard/traffic stats | System monitoring and traffic stats | Unified dashboard over trusted core data and Eleven business data | Correct aggregation across nodes and time windows |
| CLI | Existing `x-ui` operational script | PasarGuard CLI | Eleven install/backup/restore/migrate CLI where justified | Safe defaults, non-interactive use, clear failure behavior |
| Internationalization | Existing translations/localization | Multi-language support | Persian and English with a language switch | Missing keys, RTL layout, date/number formatting |
| Sales, invoices, wallet, accounting | Not a complete sales ledger by default | Not assumed complete from feature list; audit source before reuse | New Eleven commerce domain, transactional ledger and replaceable payment adapter | Duplicate payment events, reconciliation, auditability |
| Installation and updates | Existing installer/updater | Docker and multiple DB deployment options advertised | Script + Docker; verified backups, validated artifacts, rollback | Upgrade from existing install, checksum failure, interrupted update |
| Database options | Sanaei core uses GORM with SQLite/PostgreSQL support | SQLite, TimescaleDB, MySQL/MariaDB/PostgreSQL listed by PasarGuard | Keep the first release aligned with supported core DBs unless another DB is deliberately engineered | Dialect behavior and migrations |
| Migration | Existing core DB must remain safe | Separate system/data migrations | Non-destructive dry-run import and reversible mapping | Existing DB, interrupted import, rollback rehearsal |

## Initial architectural decisions

1. **Do not embed PasarGuard as a second running panel.** The goal is a unified product, not two services with duplicate logins, schemas, and node dispatch.
2. **Do not port its implementation language or database assumptions blindly.** Recreate required product capabilities in the existing Go/React architecture unless a specific integration is demonstrably safer.
3. **Reuse the Sanaei runtime for core state changes.** A second node dispatcher can silently break multi-node behavior.
4. **Keep commerce data separate from core protocol records.** Customer, order, invoice, ledger, and payment state must not be conflated with Xray client rows.
5. **Preserve core output formats.** Share links, Xray JSON, subscription formats, and protocol settings need compatibility tests.
6. **Treat parity as evidence-based.** A feature is complete only after its code path, permissions, persistence, failure modes, and tests are verified.

## Sources reviewed

- Pinned baseline details: [UPSTREAM.md](UPSTREAM.md)
- Integration boundary: [SANAEI-INTEGRATION.md](SANAEI-INTEGRATION.md)
- PasarGuard public README: https://github.com/PasarGuard/panel
