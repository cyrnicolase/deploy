package models

// Project is struct
type Project struct {
	Model `xorm:"extends"`
	Name  string `xorm:"notnull" json:"name"`
	Alias string `xorm:"notnull" json:"alias"`
	Repo  string `xorm:"notnull" json:"repo"`
}

// TableName return db tablename
func (Project) TableName() string {
	return "project.projects"
}

// Create insert a record into table
func (p *Project) Create() (int64, error) {
	return x.Insert(p)
}

// ModifyProject update project information
func (p *Project) ModifyProject() (int64, error) {
	return x.Id(p.ID).Update(p)
}

// Destroy delete a record
func (p *Project) Destroy() (int64, error) {
	return x.Id(p.ID).Delete(p)
}

// GetProjectList return list of project
func GetProjectList() (projects []Project, err error) {
	err = x.Find(&projects)
	if nil != err {
		return nil, err
	}

	return projects, nil
}
