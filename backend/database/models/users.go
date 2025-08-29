package models

type User struct {
	ID          int      `json:"id" db:"id"`
	Username    string   `json:"username" db:"username"`
	Phone       string   `json:"phone" db:"phone"`
	Roles       []string `json:"roles" db:"roles"`
	AuthGroupID int64    `json:"auth_group_id" db:"auth_group_id"`
}
