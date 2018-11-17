package karigo

// Config ...
type Config struct {
	Name     string
	PrePath  string
	Port     uint16
	Debug    bool
	Info     bool
	Minimize bool
	Store    struct {
		Driver   string
		Host     string
		Database string
		User     string
		Password string
		Options  map[string]string
	}
}
