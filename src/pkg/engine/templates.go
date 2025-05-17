package engine

type Scaffold struct {
	Name        string     `yaml:"scaffold_name"`
	Description string     `yaml:"description"`
	Params      []Param    `yaml:"params"`
	Templates   []Template `yaml:"templates"`
}

type Param struct {
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
	Type        string `yaml:"type"`
	Default     string `yaml:"default"`
}

type Template struct {
	Source      string `yaml:"source"`
	Destination string `yaml:"destination"`
}
