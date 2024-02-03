package people

import "fmt"

type ModifyWealthScriptGenerator struct {
	value float32
}

func NewModifyWealthScriptGenerator(value float32) *ModifyWealthScriptGenerator {
	return &ModifyWealthScriptGenerator{
		value: value,
	}
}

func (m *ModifyWealthScriptGenerator) Generate() []string {
	scripts := make([]string, 0)

	scripts = append(scripts, fmt.Sprintf("wealth = %.2f\t# 增加金钱", m.value))

	return scripts
}
