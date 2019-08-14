package models

// UserPrivilege 用户权限结构体
type UserPrivilege struct {
	Model     `xorm:"extends"`
	Userid    string `json:"user_id" xorm:"notnull user_id"`
	User      `xorm:"extends"`
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
	return x.Where("id = ?", id).Delete(up)
}

// UserPrivilegesByUserID 返回根据userID查询的用户所有权限
func UserPrivilegesByUserID(userID string) []UserPrivilege {
	userPrivileges := make([]UserPrivilege, 0)
	x.Join("INNER", "users.users", "users.users.id = users.user_privileges.user_id").
		Where("users.user_privileges.user_id = ?", userID).
		Find(&userPrivileges)

	return userPrivileges
}
