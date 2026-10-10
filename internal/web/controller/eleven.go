package controller

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/mhsanaei/3x-ui/v3/internal/database"
	v1 "github.com/mhsanaei/3x-ui/v3/internal/eleven/api/v1"
	elevenidentity "github.com/mhsanaei/3x-ui/v3/internal/eleven/identity"
	"github.com/mhsanaei/3x-ui/v3/internal/web/session"
)

type ElevenController struct{}

func NewElevenController(g *gin.RouterGroup) *ElevenController {
	c := &ElevenController{}

	eleven := g.Group("/api/v1/eleven")
	eleven.GET("/health", c.health)
	// The existing Sanaei session authenticates the person; an explicitly
	// provisioned, enabled Eleven identity supplies the separate Eleven role.
	// No role is auto-granted from a Sanaei administrator account.
	eleven.GET("/me", c.requirePermission("dashboard:read"), c.me)
	eleven.GET("/permissions", c.requirePermission("dashboard:read"), c.permissions)

	return c
}

func (c *ElevenController) health(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, v1.HealthResponse{
		Status:  "ok",
		Version: "0.1.0",
	})
}

func (c *ElevenController) requirePermission(permission string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user := session.GetLoginUser(ctx)
		if user == nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, v1.ErrorResponse{
				Code:    "authentication_required",
				Message: "برای دسترسی به این بخش ابتدا وارد پنل شوید.",
			})
			return
		}

		db := database.GetDB()
		if db == nil {
			ctx.AbortWithStatusJSON(http.StatusServiceUnavailable, v1.ErrorResponse{
				Code:    "identity_store_unavailable",
				Message: "سرویس احراز هویت موقتاً در دسترس نیست.",
			})
			return
		}

		var admin elevenidentity.Admin
		err := db.Where("username = ? AND enabled = ?", user.Username, true).First(&admin).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.AbortWithStatusJSON(http.StatusForbidden, v1.ErrorResponse{
				Code:    "eleven_identity_not_provisioned",
				Message: "برای این حساب، دسترسی Eleven فعال نشده است.",
			})
			return
		}
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, v1.ErrorResponse{
				Code:    "identity_lookup_failed",
				Message: "بررسی دسترسی حساب با خطا مواجه شد.",
			})
			return
		}
		if !elevenidentity.HasPermission(admin.Role, permission) {
			ctx.AbortWithStatusJSON(http.StatusForbidden, v1.ErrorResponse{
				Code:    "permission_denied",
				Message: "مجوز لازم برای انجام این کار را ندارید.",
			})
			return
		}

		ctx.Set("eleven_admin", admin)
		ctx.Next()
	}
}

func (c *ElevenController) me(ctx *gin.Context) {
	account, ok := c.elevenAdminFromContext(ctx)
	if !ok {
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"id":          account.ID,
		"username":    account.Username,
		"displayName": account.DisplayName,
		"role":        account.Role,
		"ownerId":     account.OwnerID,
		"permissions": elevenidentity.PermissionsForRole(account.Role),
	})
}

func (c *ElevenController) permissions(ctx *gin.Context) {
	account, ok := c.elevenAdminFromContext(ctx)
	if !ok {
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"role":        account.Role,
		"permissions": elevenidentity.PermissionsForRole(account.Role),
	})
}

func (c *ElevenController) elevenAdminFromContext(ctx *gin.Context) (elevenidentity.Admin, bool) {
	value, ok := ctx.Get("eleven_admin")
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, v1.ErrorResponse{
			Code:    "identity_context_missing",
			Message: "اطلاعات هویت حساب در دسترس نیست.",
		})
		return elevenidentity.Admin{}, false
	}
	account, ok := value.(elevenidentity.Admin)
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, v1.ErrorResponse{
			Code:    "identity_context_invalid",
			Message: "اطلاعات هویت حساب معتبر نیست.",
		})
		return elevenidentity.Admin{}, false
	}
	return account, true
}
