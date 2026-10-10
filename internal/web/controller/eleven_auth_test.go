package controller

import (
	"net/http"
	"net/http/cookiejar"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"

	"github.com/mhsanaei/3x-ui/v3/internal/database"
	"github.com/mhsanaei/3x-ui/v3/internal/database/model"
	elevenidentity "github.com/mhsanaei/3x-ui/v3/internal/eleven/identity"
)

func newElevenRBACClient(t *testing.T) (*http.Client, *httptest.Server) {
	t.Helper()
	engine, _ := newAPIAuthTestEngine(t)
	NewElevenController(engine.Group(""))
	server := httptest.NewServer(engine)
	t.Cleanup(server.Close)
	jar, err := cookiejar.New(nil)
	if err != nil {
		t.Fatalf("create cookie jar: %v", err)
	}
	return &http.Client{Jar: jar}, server
}

func TestElevenIdentityEndpointsRequireSession(t *testing.T) {
	client, server := newElevenRBACClient(t)
	resp, err := client.Get(server.URL + "/api/v1/eleven/me")
	if err != nil {
		t.Fatalf("request /me: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusUnauthorized {
		t.Fatalf("/me status = %d, want %d", resp.StatusCode, http.StatusUnauthorized)
	}
}

func TestElevenSessionWithoutProvisionedIdentityIsForbidden(t *testing.T) {
	client, server := newElevenRBACClient(t)
	loginResp, err := client.Get(server.URL + "/test-login")
	if err != nil {
		t.Fatalf("login: %v", err)
	}
	loginResp.Body.Close()
	if loginResp.StatusCode != http.StatusOK {
		t.Fatalf("login status = %d, want %d", loginResp.StatusCode, http.StatusOK)
	}

	resp, err := client.Get(server.URL + "/api/v1/eleven/me")
	if err != nil {
		t.Fatalf("request /me: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusForbidden {
		t.Fatalf("/me status = %d, want %d", resp.StatusCode, http.StatusForbidden)
	}
}

func TestElevenResellerCannotReadAdminDirectory(t *testing.T) {
	client, server := newElevenRBACClient(t)
	loginResp, err := client.Get(server.URL + "/test-login")
	if err != nil {
		t.Fatalf("login: %v", err)
	}
	loginResp.Body.Close()
	if loginResp.StatusCode != http.StatusOK {
		t.Fatalf("login status = %d, want %d", loginResp.StatusCode, http.StatusOK)
	}

	var user model.User
	if err := database.GetDB().Order("id ASC").First(&user).Error; err != nil {
		t.Fatalf("load logged-in user: %v", err)
	}
	admin := elevenidentity.Admin{
		Username: user.Username,
		DisplayName: "Test reseller",
		Role: "reseller",
		Enabled: true,
	}
	if err := database.GetDB().Create(&admin).Error; err != nil {
		t.Fatalf("provision reseller identity: %v", err)
	}

	meResp, err := client.Get(server.URL + "/api/v1/eleven/me")
	if err != nil {
		t.Fatalf("request /me: %v", err)
	}
	meResp.Body.Close()
	if meResp.StatusCode != http.StatusOK {
		t.Fatalf("/me status = %d, want %d", meResp.StatusCode, http.StatusOK)
	}

	adminsResp, err := client.Get(server.URL + "/api/v1/eleven/admins")
	if err != nil {
		t.Fatalf("request /admins: %v", err)
	}
	defer adminsResp.Body.Close()
	if adminsResp.StatusCode != http.StatusForbidden {
		t.Fatalf("/admins status = %d, want %d", adminsResp.StatusCode, http.StatusForbidden)
	}
}

