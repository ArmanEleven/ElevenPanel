# Eleven API v1

The public integration API is versioned separately from the existing Sanaei API.

Base path:

`/api/v1/eleven`

## Planned endpoints

### Health
`GET /health`

### Users
`POST /users/provision`
`GET /users/:id`
`PATCH /users/:id`
`DELETE /users/:id`

### Subscriptions
`POST /subscriptions`
`POST /subscriptions/:id/renew`
`POST /subscriptions/:id/revoke`

### Nodes
`GET /nodes`
`GET /nodes/:id/health`

## Integration rules

1. Every write requires API authentication.
2. Provisioning requests accept an idempotency key.
3. Every write creates an audit event.
4. Errors use stable machine-readable codes.
5. The API must not expose internal Sanaei database structures.
6. Eleven Store should depend only on this API contract.
