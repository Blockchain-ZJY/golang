package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// TestGenerateToken 测试 Token 生成功能
func TestGenerateToken(t *testing.T) {
	token, err := GenerateToken("testuser")
	if err != nil {
		t.Fatalf("Failed to generate token: %v", err)
	}
	if token == "" {
		t.Error("Generated token is empty")
	}
}

// TestLoginHandler 测试登录接口
func TestLoginHandler(t *testing.T) {
	router := setupRouter()

	// 1. 测试成功登录
	w := httptest.NewRecorder()
	// 注意：这里根据您当前的 main.go 是 GET 请求
	req, _ := http.NewRequest("GET", "/login?username=admin&password=password", nil)
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}

	var response map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("Failed to parse JSON response: %v", err)
	}

	token, exists := response["token"]
	if !exists || token == "" {
		t.Error("Token not found in response")
	}

	// 2. 测试失败登录
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/login?username=admin&password=wrong", nil)
	router.ServeHTTP(w, req)

	if w.Code != http.StatusUnauthorized {
		t.Errorf("Expected status 401 for wrong password, got %d", w.Code)
	}
}

// TestProfileHandler 测试受保护接口
func TestProfileHandler(t *testing.T) {
	router := setupRouter()

	// 先生成一个有效 token
	token, _ := GenerateToken("admin")

	// 1. 测试带有效 Token 访问
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/auth/profile", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d. Body: %s", w.Code, w.Body.String())
	}

	if !strings.Contains(w.Body.String(), "admin") {
		t.Errorf("Response body should contain username 'admin', got: %s", w.Body.String())
	}

	// 2. 测试不带 Token 访问
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/auth/profile", nil)
	router.ServeHTTP(w, req)

	if w.Code != http.StatusUnauthorized {
		t.Errorf("Expected status 401 for missing token, got %d", w.Code)
	}

	// 3. 测试带无效 Token 访问
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/auth/profile", nil)
	req.Header.Set("Authorization", "Bearer invalid-token")
	router.ServeHTTP(w, req)

	if w.Code != http.StatusUnauthorized {
		t.Errorf("Expected status 401 for invalid token, got %d", w.Code)
	}
}
