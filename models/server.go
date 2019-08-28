package models

// Server orm table
type Server struct {
	Model  `xorm:"extends"`
	Addr   string `xorm:"notnull" json:"addr"`
	Port   int    `xorm:"notnull" json:"port"`
	Uname  string `xorm:"notnull" json:"uname"`
	Passwd string `xorm:"notnull" json:"passwd"`
	Sshkey int    `xorm:"notnull" json:"sshkey"`
}

// TableName returns tablename
func (Server) TableName() string {
	return "project.servers"
}

// Create insert a record into table
func (s *Server) Create() (int64, error) {
	return x.Insert(s)
}

// ModifyServer  update project information
func (s *Server) ModifyServer() (int64, error) {
	return x.Id(s.ID).Update(s)
}

// Destroy deletes record
func (s *Server) Destroy() (int64, error) {
	return x.Id(s.ID).Delete(s)
}
