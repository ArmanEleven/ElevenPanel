# Eleven Panel

**An independent management layer for Xray panel operations — currently in architecture planning.**

Eleven Panel aims to build Eleven-specific management capabilities around the Sanaei/3X-UI ecosystem without rewriting the existing Xray management foundation.

> **Status:** Phase 0 — architecture and repository bootstrap. The upstream source has intentionally not been imported while the target revision and licensing requirements are being established. This repository is not a production-ready panel or a drop-in replacement.

## Planned architecture

| Layer | Intended responsibility |
| --- | --- |
| Core | Sanaei/3X-UI Xray management |
| Identity | Administrator accounts, roles, and permissions |
| Services | Users, groups, templates, and subscriptions |
| Infrastructure | Nodes, health information, and monitoring |
| Operations | Statistics, logs, and bulk actions |
| Integration | A stable API for Eleven Store |
| Migration | Data migration, validation, and rollback |

These layers describe the intended direction; they are not a list of shipped features.

## Engineering priorities

1. Pin and document the intended upstream revision.
2. Separate upstream functionality from Eleven-specific management logic.
3. Define identity, permission, and resource ownership boundaries.
4. Validate API contracts and data models before expanding scope.
5. Add tests for migrations, authorization, and failure cases.
6. Document installation and rollback only after those paths are tested.

## Upstream projects

- [Sanaei 3X-UI](https://github.com/MHSanaei/3x-ui)
- [PasarGuard Panel](https://github.com/PasarGuard/panel)

The project may evaluate ideas and compatibility requirements from existing software, but that does not imply code has been copied or integrated.

## Development approach

A capability should be documented as implemented only when the code exists and its behavior can be verified. Operational changes must account for existing panel data and include a recovery path.

## License and attribution

Before incorporating or redistributing code from an upstream project, review the exact revision's license, copyright notices, and source-distribution requirements. Preserve all notices required by the applicable licenses.
