package dependencyinjection

import (
	"context"
	"testing"
)

// MockUserRepository is a mock implementation for testing
type MockUserRepository struct {
	Users    map[string]*User
	SaveFunc func(user *User) error
}

func (m *MockUserRepository) Save(ctx context.Context, user *User) error {
	if m.SaveFunc != nil {
		return m.SaveFunc(user)
	}
	if m.Users == nil {
		m.Users = make(map[string]*User)
	}
	m.Users[user.ID] = user
	return nil
}

func (m *MockUserRepository) FindByID(ctx context.Context, id string) (*User, error) {
	if user, ok := m.Users[id]; ok {
		return user, nil
	}
	return nil, nil
}

func TestUserService_Register(t *testing.T) {
	// 1. Prepare Mock
	mockRepo := &MockUserRepository{
		Users: make(map[string]*User),
	}

	// 2. Inject Mock into Service
	service := NewUserService(mockRepo)

	// 3. Test Business Logic
	ctx := context.Background()
	err := service.Register(ctx, "test-id", "Test User", "test@test.com")

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// 4. Verify Side Effects
	if len(mockRepo.Users) != 1 {
		t.Errorf("Expected 1 user in repo, got %d", len(mockRepo.Users))
	}
	if mockRepo.Users["test-id"].Name != "Test User" {
		t.Errorf("User name mismatch")
	}
}
