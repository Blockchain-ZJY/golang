package dependencyinjection

import (
	"context"
	"errors"
	"fmt"
)

// UserService 包含业务逻辑并依赖于 UserRepository
// 这演示了依赖注入：我们不在内部创建存储库，
// 而是请求它。
type UserService struct {
	repo UserRepository
}

// NewUserService 是发生注入的构造函数
func NewUserService(repo UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) Register(ctx context.Context, id, name, email string) error {
	if id == "" || name == "" {
		return errors.New("invalid user data")
	}

	user := &User{
		ID:    id,
		Name:  name,
		Email: email,
	}

	// 业务逻辑...
	fmt.Printf("[Service] 正在注册用户: %s\n", name)

	return s.repo.Save(ctx, user)
}

func (s *UserService) GetUser(ctx context.Context, id string) (*User, error) {
	return s.repo.FindByID(ctx, id)
}
