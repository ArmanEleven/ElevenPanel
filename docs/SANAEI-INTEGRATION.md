# Sanaei Integration Map

This document records the integration boundary for Eleven Panel against the pinned Sanaei/3x-ui v3.9.0 baseline.

## Confirmed upstream structure

The pinned upstream tree uses Go 1.27.1 and exposes its web layer under `internal/web`.

Important service/controller boundaries include:

- `service.XrayService` — Xray lifecycle and runtime operations.
- `service.InboundService` — inbound persistence and management.
- `service.ClientService` — client persistence and management.
- `service.NodeService` — node management.
- `internal/web/controller/inbound.go`
- `internal/web/controller/client.go`
- `internal/web/controller/node.go`
- `internal/web/controller/group.go`
- `internal/web/web.go`

## Integration rule

Eleven code must not duplicate Xray/inbound/client logic.

Instead:

1. Sanaei remains the source of truth for existing Xray configuration.
2. Eleven owns additional service-management metadata.
3. An adapter translates Eleven operations into Sanaei service calls.
4. Existing Sanaei APIs continue to work unless an Eleven feature explicitly extends them.
5. Migration must remain non-destructive.

## First adapter boundary

The first implementation target is a small internal adapter around the operations Eleven needs:

- provision a client
- update a client
- revoke a client
- read client traffic
- read inbound information
- trigger/reconcile Xray state when required

The adapter should depend on interfaces, not Sanaei concrete implementations, so tests can use fakes.

## Database rule

Do not modify the existing Sanaei tables merely to add Eleven metadata.

Initial Eleven metadata belongs in separate `eleven_*` tables. A later migration can map existing Sanaei clients/inbounds into Eleven users and subscriptions without rewriting the source database during dry-run.

## Current status

- Upstream version pinned: v3.9.0
- Upstream commit pinned: 3cd4bf504c3cd8ea9b1c1fdb032a9796c5c43ddb
- Full upstream source vendoring: complete
- Adapter implementation: partial; provisioning, updates, revocation, traffic reads, and inbound lookup use the Sanaei service layer
- Reconciliation: deliberately returns an explicit not-implemented error until Eleven ownership and node synchronization are defined
- Eleven Store business API: not implemented beyond the public health endpoint
