package v1

type ErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type ProvisionUserRequest struct {
	ExternalID string  `json:"externalId"`
	Username   string  `json:"username"`
	TemplateID *uint64 `json:"templateId,omitempty"`
	GroupID    *uint64 `json:"groupId,omitempty"`
	ExpiresAt  string  `json:"expiresAt"`
}

type ProvisionUserResponse struct {
	UserID        uint64 `json:"userId"`
	SubscriptionID uint64 `json:"subscriptionId"`
}

type HealthResponse struct {
	Status  string `json:"status"`
	Version string `json:"version"`
}
