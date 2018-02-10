package config

type Config struct {
	Profiles []Profile
}

type Profile struct {
	Server   string
	Username string
	Password string
	Path     string
	Tasks    []BackupTask
}

type BackupTask struct {
	From    string
	To      string
	Exclude []string
}
