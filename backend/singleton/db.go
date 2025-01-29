package singleton

type DBConfig struct {
	Driver           string `mapstructure:"driver" json:"driver" yaml:"driver"`
	Source           string `mapstructure:"source" json:"source" yaml:"source"`
	MigrationFileDir string `mapstructure:"migration-file-dir" json:"migration_file_dir" yaml:"migration-file-dir"`
	Password         string `mapstructure:"password" json:"password" yaml:"password"`
}
