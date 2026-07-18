package config

type Config struct {
	Server Server			`yaml:"server"`
	Routes []Route			`yaml:"routes"`
}

type Server struct {
	Port int				`yaml:"port"`
}

type Route struct {
	Name string				`yaml:"name"`
	Path string				`yaml:"path"`
	Upstream string			`yaml:"upstream"`
	Methods []string		`yaml:"methods"`
}