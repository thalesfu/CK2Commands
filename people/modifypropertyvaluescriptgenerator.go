package people

import "fmt"

type ModifyPropertyValueScriptGenerator struct {
	PropertyModifier *PropertyModifier
}

func NewModifyPropertyValueScriptGenerator(propertyModifier *PropertyModifier) *ModifyPropertyValueScriptGenerator {
	return &ModifyPropertyValueScriptGenerator{
		PropertyModifier: propertyModifier,
	}
}

func (m *ModifyPropertyValueScriptGenerator) Generate() []string {
	scripts := make([]string, 0)

	scripts = append(scripts, fmt.Sprintf("change_%s = %d", m.PropertyModifier.Property, m.PropertyModifier.ModifiedValue))

	return scripts
}
