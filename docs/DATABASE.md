# Eleven database layer

The Eleven layer uses separate tables/namespaces from the Sanaei core.

Initial logical tables:

- eleven_admins
- eleven_roles
- eleven_permissions
- eleven_admin_roles
- eleven_role_permissions
- eleven_users
- eleven_groups
- eleven_templates
- eleven_subscriptions
- eleven_nodes
- eleven_hwids
- eleven_activity_logs
- eleven_api_keys
- eleven_idempotency_keys

Do not modify existing Sanaei tables until a concrete compatibility requirement is identified.

## Migration rule

Migration must be additive and reversible. A failed migration must leave the original Sanaei database untouched.
