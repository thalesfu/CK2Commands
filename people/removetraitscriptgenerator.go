package people

import (
	"fmt"
	"github.com/thalesfu/ck2nebula"
)

type RemoveTraitScriptGenerator struct {
	Trait *ck2nebula.Trait
}

func NewRemoveTraitScriptGenerator(trait *ck2nebula.Trait) *RemoveTraitScriptGenerator {
	return &RemoveTraitScriptGenerator{
		Trait: trait,
	}
}

func (r *RemoveTraitScriptGenerator) Generate() []string {
	scripts := make([]string, 0)

	scripts = append(scripts, fmt.Sprintf("remove_trait = %s\t# 去除 %s", r.Trait.Code, r.Trait.Name))

	return scripts
}
