// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"gorm.io/gorm"
)

const TableNameArticle = "article"

// Article mapped from table <article>
type Article struct {
	ID         int64          `gorm:"column:id;type:int(10) unsigned;primaryKey;autoIncrement:true;comment:ID" json:"id"` // ID
	UID        int64          `gorm:"column:uid;type:int(11);not null;comment:用户id" json:"uid"`                           // 用户id
	User       *User          `json:"user" gorm:"foreignkey:uid;references:id"`                                           // 关联用户
	Title      string         `gorm:"column:title;type:varchar(50);not null;comment:标题" json:"title"`                     // 标题
	Content    string         `gorm:"column:content;type:varchar(255);not null;comment:内容" json:"content"`                // 内容
	CategoryID int64          `gorm:"column:category_id;type:int(11);not null;comment:分类id" json:"categoryId"`            // 分类id
	DataSource int64          `gorm:"column:data_source;type:int(11);not null;comment:数据来源 1=文章库 2=自建" json:"dataSource"` // 数据来源 1=文章库 2=自建
	IsPublish  int64          `gorm:"column:is_publish;type:int(11);not null;comment:是否发布 1=已发布 2=未发布" json:"isPublish"`  // 是否发布 1=已发布 2=未发布
	Category   *Category      `json:"category" gorm:"foreignkey:category_id;references:id"`                               // 关联分类
	Tag        JsonString     `gorm:"column:tag;type:json;comment:标签" json:"tag"`                                         // 标签
	CreatedAt  *JsonTime      `gorm:"column:created_at;type:datetime;comment:创建时间" json:"createdAt"`                      // 创建时间
	UpdatedAt  *JsonTime      `gorm:"column:updated_at;type:datetime;comment:更新时间" json:"updatedAt"`                      // 更新时间
	DeletedAt  gorm.DeletedAt `gorm:"column:deleted_at;type:datetime;comment:删除时间" json:"deletedAt"`                      // 删除时间
}

// TableName Article's table name
func (*Article) TableName() string {
	return TableNameArticle
}
