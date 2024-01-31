package people

import "fmt"

type ModifyFertilityScriptGenerator struct {
	value float32
}

func NewModifyFertilityScriptGenerator(value float32) *ModifyFertilityScriptGenerator {
	return &ModifyFertilityScriptGenerator{
		value: value,
	}
}

func (m *ModifyFertilityScriptGenerator) Generate() []string {
	scripts := make([]string, 0)

	scripts = append(scripts, fmt.Sprintf("fertility = %.2f\t# 增加生育能力", m.value))

	return scripts
}
