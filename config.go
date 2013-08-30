package dingo

type Config struct {
	Port        string
	TemplateDir string
	StaticDir   string
	Routes      []*AHandler
}
