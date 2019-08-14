package config

import (
	"reflect"
	"testing"
)

func TestUnmarshalPrivileges(t *testing.T) {
	unmarshalPrivileges()
	tPrivileges := PrivilegeConfigs{
		Users: []string{
			"新增用户",
			"修改用户",
			"删除用户",
			"查看用户",
			"关联用户权限",
		},
	}

	if reflect.DeepEqual(tPrivileges, Privileges) {
		t.Errorf("期望：%v, 实际:%v", tPrivileges, Privileges)
	}
}
