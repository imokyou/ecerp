// Package user 提供易仓ERP用户相关API的封装
package user

import (
	"context"

	"github.com/imokyou/ecerp"
)

// Service 用户服务
type Service struct {
	C ecerp.Caller
}

// NewService 创建用户服务
func NewService(c ecerp.Caller) *Service {
	return &Service{C: c}
}

// ════════════════════════════════════════════
// 数据结构
// ════════════════════════════════════════════

// User 用户信息
type User struct {
	UserID     int    `json:"user_id"`
	UserCode   string `json:"user_code"`
	UserName   string `json:"user_name"`
	UserNameEn string `json:"user_name_en"`
	UserStatus int    `json:"user_status"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
}

// PlatformAccount 平台账号
type PlatformAccount struct {
	AccountID   int    `json:"account_id"`
	AccountName string `json:"account_name"`
	Platform    string `json:"platform"`
	Status      int    `json:"status"`
}

// ════════════════════════════════════════════
// 请求参数
// ════════════════════════════════════════════

// CreateUserRequest 创建用户请求
type CreateUserRequest struct {
	UserCode     string `json:"user_code,omitempty"`
	UserPassword string `json:"user_password,omitempty"`
	UserName     string `json:"user_name,omitempty"`
	UserNameEn   string `json:"user_name_en,omitempty"`
	UserStatus   int    `json:"user_status,omitempty"`
	Email        string `json:"email,omitempty"`
	Phone        string `json:"phone,omitempty"`
}

// EditUserRequest 编辑用户请求
type EditUserRequest struct {
	UserID       int    `json:"user_id"`
	UserCode     string `json:"user_code,omitempty"`
	UserPassword string `json:"user_password,omitempty"`
	UserName     string `json:"user_name,omitempty"`
	UserNameEn   string `json:"user_name_en,omitempty"`
	UserStatus   int    `json:"user_status,omitempty"`
	Email        string `json:"email,omitempty"`
	Phone        string `json:"phone,omitempty"`
}

// PageRequest 分页请求
type PageRequest struct {
	Page     int `json:"page,omitempty"`
	PageSize int `json:"page_size,omitempty"`
}

// ════════════════════════════════════════════
// 服务方法 (5个接口)
// ════════════════════════════════════════════

// EditUser 编辑用户
func (s *Service) EditUser(ctx context.Context, req *EditUserRequest) error {
	return s.C.Do(ctx, "editUser", req, nil)
}

// CreateUser 创建用户
func (s *Service) CreateUser(ctx context.Context, req *CreateUserRequest) error {
	return s.C.Do(ctx, "createUser", req, nil)
}

// BatchEditUser 批量修改用户
func (s *Service) BatchEditUser(ctx context.Context, users []EditUserRequest) error {
	return s.C.Do(ctx, "batchEditUser", users, nil)
}

// BatchCreateUser 批量创建用户
func (s *Service) BatchCreateUser(ctx context.Context, users []CreateUserRequest) error {
	return s.C.Do(ctx, "batchCreateUser", users, nil)
}

// GetUserAccountList 获取平台账号列表
func (s *Service) GetUserAccountList(ctx context.Context, req *PageRequest) ([]PlatformAccount, error) {
	var result []PlatformAccount
	err := s.C.Do(ctx, "getUserAccountList", req, &result)
	return result, err
}
