package models

// UserPrivilege 用户权限结构体
type UserPrivilege struct {
	Model     `xorm:"extends"`
	Userid    string `json:"user_id" xorm:"notnull user_id"`
	Privilege string `json:"privilege" xorm:"notnull"`
}

// TableName UserPrivilege tablename
func (UserPrivilege) TableName() string {
	return "users.user_privileges"
}

// Insert 新增
func (up *UserPrivilege) Insert() (affected int64, err error) {
	return x.Insert(up)
}

// Delete 删除
func (up *UserPrivilege) Delete() (affected int64, err error) {
	return x.Delete(&up)
}

// DeleteByID 按照用户权限id删除数据
func DeleteByID(id string) (affected int64, err error) {
	up := new(UserPrivilege)
	return x.Where("id = ?", id).Delete(up) // 这里用Where，而不是用Id()；我理解是因为extends 引起的pk重复；导致出问题
}

// UserPrivilegeSet 用户权限集合
type UserPrivilegeSet struct {
	UserPrivilege `xorm:"extends"`
	User          `xorm:"extends"`
}

// TableName 返回复合结构体 UserPrivilegeSet 对应的表
func (UserPrivilegeSet) TableName() string {
	return "users.user_privileges"
}

// UserPrivilegesByUserID 返回根据userID查询的用户所有权限
func UserPrivilegesByUserID(userID string) []UserPrivilegeSet {
	userPrivileges := make([]UserPrivilegeSet, 0)
	x.Join("INNER", "users.users", "users.users.id = users.user_privileges.user_id").
		Where("users.user_privileges.user_id = ?", userID).
		Find(&userPrivileges)

	return userPrivileges
}
