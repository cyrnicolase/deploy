package config

// PrivilegeConfigs struct
type PrivilegeConfigs struct {
	Users []string
}

// Privileges is all privileges
var Privileges PrivilegeConfigs

func unmarshalPrivileges() {
	v, _ := readYaml("privileges")
	v.Unmarshal(&Privileges)
}
