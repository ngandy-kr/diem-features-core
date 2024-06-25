package model

type Installer interface {
	Install(feature Feature, pathTemplate string, pathTarget string, args []Argument) error
}

type FeaturesConfig struct {
	Features []Feature
}

type Feature struct {
	Enable      bool
	Name        string
	Description string
	Arguments   []Argument
	Installer   Installer
	Path        string
}

type Argument struct {
	Name     string
	Value    string
	Default  string
	Required bool
}
