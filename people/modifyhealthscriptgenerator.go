package people

import "fmt"

type ModifyHealthScriptGenerator struct {
	value float32
}

func NewModifyHealthScriptGenerator(value float32) *ModifyHealthScriptGenerator {
	return &ModifyHealthScriptGenerator{
		value: value,
	}
}

func (m *ModifyHealthScriptGenerator) Generate() []string {
	scripts := make([]string, 0)

	scripts = append(scripts, fmt.Sprintf("health = %.2f\t# 增加健康", m.value))

	return scripts
}
