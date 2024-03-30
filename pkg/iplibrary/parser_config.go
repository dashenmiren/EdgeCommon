package iplibrary

type ParserConfig struct {
	Template    *Template
	EmptyValues []string
	Iterator    func(values map[string]string) error
}
