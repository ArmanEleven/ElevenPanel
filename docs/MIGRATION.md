# Sanaei -> Eleven Panel migration

Migration is designed to be non-destructive.

## Rules

1. Create a verified backup before migration.
2. Never write directly into the original Sanaei database during a dry run.
3. Detect the source schema/version before importing.
4. Import users, clients, inbounds, hosts, traffic and expiry information into a staging model.
5. Validate references before committing.
6. Keep an export of the Eleven-side mapping so rollback is possible.
7. Only after validation should the live service be switched to Eleven Panel.

## Planned command

eleven migrate sanaei --source /path/to/x-ui.db --dry-run

No destructive migration is allowed by default.
