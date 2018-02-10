package config

type Config struct {
	Tasks []Task
}
type Task struct {
	Source      string
	Destination Dest
}

type Dest struct {
	Server   string
	Port     int
	Path     string
	Username string
	Password string
}
