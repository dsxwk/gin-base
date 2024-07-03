// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameUser = "user"

// User mapped from table <user>
type User struct {
	ID        int64   `gorm:"column:id;type:int(10) unsigned;primaryKey;autoIncrement:true;comment:ID" json:"id"` // ID
	Username  string  `gorm:"column:username;type:varchar(10);not null;comment:用户名" json:"username"`              // 用户名
	FullName  string  `gorm:"column:full_name;type:varchar(20);not null;comment:姓名" json:"full_name"`             // 姓名
	Email     string  `gorm:"column:email;type:varchar(50);not null;comment:邮箱" json:"email"`                     // 邮箱
	Password  string  `gorm:"column:password;type:varchar(255);not null;comment:密码" json:"password"`              // 密码
	Nickname  string  `gorm:"column:nickname;type:varchar(50);not null;comment:昵称" json:"nickname"`               // 昵称
	Gender    int64   `gorm:"column:gender;type:tinyint(1) unsigned;not null;comment:性别 1=男 2=女" json:"gender"`   // 性别 1=男 2=女
	Age       int64   `gorm:"column:age;type:int(10) unsigned;not null;comment:年龄" json:"age"`                    // 年龄
	CreatedAt *string `gorm:"column:created_at;type:datetime;comment:创建时间" json:"created_at"`                     // 创建时间
	UpdatedAt *string `gorm:"column:updated_at;type:datetime;comment:更新时间" json:"updated_at"`                     // 更新时间
	DeletedAt *string `gorm:"column:deleted_at;type:datetime;comment:删除时间" json:"deleted_at"`                     // 删除时间
}

// TableName User's table name
func (User) TableName() string {
	return TableNameUser
}
