package karigo

// Config ...
type Config struct {
	Name     string
	Port     uint16
	Debug    bool
	Info     bool
	Error    bool
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
