package controller

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"h-ui/service"
)

func TestLoginTelegramFailure(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.POST("/hui/auth/login", Login)
	oldExist := service.ExistAccountUsername
	oldLogin := service.Login
	oldRemind := service.TelegramLoginRemind
	service.ExistAccountUsername = func(string, int64) bool { return true }
	service.Login = func(string, string) (string, error) { return "token", nil }
	service.TelegramLoginRemind = func(string, string) error { return errors.New("fail") }
	defer func() {
		service.ExistAccountUsername = oldExist
		service.Login = oldLogin
		service.TelegramLoginRemind = oldRemind
	}()
	req := httptest.NewRequest("POST", "/hui/auth/login", strings.NewReader(`{"username":"user123","pass":"abcdef"}`))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("status %d", w.Code)
	}
	if !strings.Contains(w.Body.String(), "telegram_warning") {
		t.Fatalf("missing warning")
	}
}
