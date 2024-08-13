package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"userauth/routes"
)

func TestRegisterAndLogin(t *testing.T) {
	router := routes.SetupRouter()

	// Test user registration
	user := `{"name":"John Doe","email":"john@example.com","password":"password123"}`
	req, _ := http.NewRequest("POST", "/register", bytes.NewBufferString(user))
	req.Header.Set("Content-Type", "application/json")
	response := httptest.NewRecorder()
	router.ServeHTTP(response, req)

	if status := response.Code; status != http.StatusCreated {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

	// Test user login
	login := `{"email":"john@example.com","password":"password123"}`
	req, _ = http.NewRequest("POST", "/login", bytes.NewBufferString(login))
	req.Header.Set("Content-Type", "application/json")
	response = httptest.NewRecorder()
	router.ServeHTTP(response, req)

	if status := response.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var result map[string]string
	json.NewDecoder(response.Body).Decode(&result)
	token, exists := result["token"]
	if !exists {
		t.Errorf("Token not returned in response")
	}
	if token == "" {
		t.Errorf("Token is empty")
	}

	// Test getting all users with token
	req, _ = http.NewRequest("GET", "/users", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	response = httptest.NewRecorder()
	router.ServeHTTP(response, req)

	if status := response.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestHealthCheck(t *testing.T) {
	router := routes.SetupRouter()

	req, _ := http.NewRequest("GET", "/health", nil)
	response := httptest.NewRecorder()
	router.ServeHTTP(response, req)

	if status := response.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := "OK"
	if strings.TrimSpace(response.Body.String()) != expected {
		t.Errorf("Handler returned unexpected body: got %v want %v", response.Body.String(), expected)
	}
}
