package dependencyinjection

import (
	"context"
	"fmt"
)

// SQLUserRepository 是 UserRepository 的具体实现
// 它模拟数据库操作。
type SQLUserRepository struct {
	// 数据库连接将在这里
	DBString string
}

func NewSQLUserRepository(dbString string) *SQLUserRepository {
	return &SQLUserRepository{
		DBString: dbString,
	}
}

func (r *SQLUserRepository) Save(ctx context.Context, user *User) error {
	fmt.Printf("[SQL Repo] 保存用户 %s 到数据库 %s\n", user.Name, r.DBString)
	return nil
}

func (r *SQLUserRepository) FindByID(ctx context.Context, id string) (*User, error) {
	fmt.Printf("[SQL Repo] 查找用户 %s (数据库: %s)\n", id, r.DBString)
	return &User{ID: id, Name: "Real User", Email: "real@example.com"}, nil
}
