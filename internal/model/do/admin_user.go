// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AdminUser is the golang structure of table jjcc_admin_user for DAO operations like Where/Data.
type AdminUser struct {
	g.Meta         `orm:"table:jjcc_admin_user, do:true"`
	Id             interface{} // 主键ID
	DingUserId     interface{} // 钉钉用户ID,很长，必须为varchar
	Mobile         interface{} // 管理员账号
	Name           interface{} // 管理员姓名
	Passwd         interface{} // 登录密码
	Sex            interface{} // 0未知 1男 2女
	BkOrgId        interface{} //
	Email          interface{} // 邮箱
	Status         interface{} // 状态 0：正常  1：冻结
	IsExternalUser interface{} // 是否外部用户
	IsBuiltinUser  interface{} // 是否系统内置用户
	LastLoginTime  *gtime.Time // 最后登录时间
	LastLoginIp    interface{} // 最后登录IP
	IsBind         interface{} // 是否绑定酷家乐账号：0否 1是
	JobNumber      interface{} // 工号
	Position       interface{} // 职位
	LeaderId       interface{} //
	EntryTime      *gtime.Time // 入职时间
	IsInit         interface{} // 是否是初始化状态：0否 1是
	Source         interface{} // 创建来源：0：erp后台  1：接口
	IsShow         interface{} // 是否显示在后台:1显示，0否
	Avatar         interface{} // 用户头像（暂时用于竣工调研）
	CreateTime     *gtime.Time // 创建时间
	UpdateTime     *gtime.Time // 最后更新时间
}