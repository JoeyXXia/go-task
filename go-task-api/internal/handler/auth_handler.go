package handler

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/joey/go-task/internal/repository"
	"github.com/joey/go-task/internal/service"
	"github.com/joey/go-task/pkg/response"
	"gorm.io/gorm"
)

// AuthHandler 处理认证相关的HTTP请求
type AuthHandler struct {
	authService *service.AuthService
}

// NewAuthHandler 创建一个新的认证处理器
func NewAuthHandler(db *gorm.DB) *AuthHandler {
	userRepo := repository.NewUserRepository(db)
	authService := service.NewAuthService(userRepo)
	return &AuthHandler{authService}
}

// RegisterRequest 表示注册请求的数据结构
type RegisterRequest struct {
	Username string `json:"username" binding:"required,min=3,max=30"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

// LoginRequest 表示登录请求的数据结构
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Register 处理用户注册
func (h *AuthHandler) Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid request: "+err.Error())
		return
	}

	user, err := h.authService.Register(req.Username, req.Email, req.Password)
	if err != nil {
		if strings.Contains(err.Error(), "already exists") {
			response.Error(c, http.StatusConflict, err.Error())
			return
		}
		response.Error(c, http.StatusInternalServerError, "Registration failed: "+err.Error())
		return
	}

	response.Success(c, http.StatusCreated, gin.H{
		"message": "User registered successfully",
		"user": gin.H{
			"id":       user.ID,
			"username": user.Username,
			"email":    user.Email,
		},
	})
}

// Login 处理用户登录
func (h *AuthHandler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid request: "+err.Error())
		return
	}

	token, err := h.authService.Login(req.Username, req.Password)
	if err != nil {
		response.Error(c, http.StatusUnauthorized, "Login failed: "+err.Error())
		return
	}

	// 设置Cookie
	c.SetCookie("token", token, 3600*24, "/", "", false, true)

	response.Success(c, http.StatusOK, gin.H{
		"message": "Login successful",
		"token":   token,
	})
}

// Logout 处理用户登出
func (h *AuthHandler) Logout(c *gin.Context) {
	// 清除Cookie
	c.SetCookie("token", "", -1, "/", "", false, true)

	response.Success(c, http.StatusOK, gin.H{
		"message": "Logout successful",
	})
}
