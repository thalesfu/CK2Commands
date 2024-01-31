package people

import (
	"fmt"
	"github.com/thalesfu/ck2nebula"
)

type AddTraitScriptGenerator struct {
	Trait *ck2nebula.Trait
}

func NewAddTraitScriptGenerator(trait *ck2nebula.Trait) *AddTraitScriptGenerator {
	return &AddTraitScriptGenerator{
		Trait: trait,
	}
}

func (a *AddTraitScriptGenerator) Generate() []string {
	scripts := make([]string, 0)

	scripts = append(scripts, fmt.Sprintf("add_trait = %s\t# 添加 %s", a.Trait.Code, a.Trait.Name))

	return scripts
}
