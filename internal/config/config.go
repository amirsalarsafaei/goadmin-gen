package config

type PrimaryKeyConfig struct {
	Name     string `yaml:"name"`
	TypeName string `yaml:"type"`
}

type FieldConfig struct {
	Display   string   `yaml:"display"`
	Column    string   `yaml:"column"`
	TypeName  string   `yaml:"type"`
	Form      []string `yaml:"form"`
	OnlyAdmin bool     `yaml:"only_admin"`
}

type TableConfig struct {
	Display    string           `yaml:"display"`
	Table      string           `yaml:"table"`
	PrimaryKey PrimaryKeyConfig `yaml:"primary_key"`
	Fields     []FieldConfig    `yaml:"fields"`
	UserQuery  string           `yaml:"user_query"`
	Options    []string         `yaml:"options"`
}

type SqlcConfig struct {
	Schema string `yaml:"schema"`
}

type GoAdminGen struct {
	Tables []TableConfig `yaml:"tables"`
	Sqlc   SqlcConfig    `yaml:"sqlc"`
	Out    string        `yaml:"out"`
}
