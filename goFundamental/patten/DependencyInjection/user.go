package dependencyinjection

import "context"

// User 是领域实体
type User struct {
	ID    string
	Name  string
	Email string
}

// UserRepository 定义了用户数据访问的接口
type UserRepository interface {
	Save(ctx context.Context, user *User) error
	FindByID(ctx context.Context, id string) (*User, error)
}
