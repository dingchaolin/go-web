package models

import (
	"time"
)

type Category struct{
	Id int64 // 如果字段名是Id 类型是int orm就会认为是主键
	Title string //默认255
	Created time.Time `orm:"index"` //tag必须是`` 括起来 必须是导出属性 才是reflection
	Views int64 `orm:"index"`
	TopicTime time.Time `orm:"index"`
	TopicCount int64
	TopicLastUserId int64
}

type Topic struct {
	Id int64
	Uid int64
	Title string
	Content string `orm:"size(5000)"`
	Attachment string
	Created time.Time `orm:"index"` //tag必须是`` 括起来 必须是导出属性 才是reflection
	Update time.Time `orm:"index"` //tag必须是`` 括起来 必须是导出属性 才是reflection
	Views int64 `orm:"index"`
	Author string
	ReplyTime time.Time
	ReplyCount int64
	ReplyLastUserId int64
}