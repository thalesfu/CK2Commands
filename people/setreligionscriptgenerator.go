package people

import (
	"fmt"
	"github.com/thalesfu/ck2nebula"
)

type SetReligionScriptGenerator struct {
	Religion *ck2nebula.Religion
}

func NewSetReligionScriptGenerator(religion *ck2nebula.Religion) *SetReligionScriptGenerator {
	return &SetReligionScriptGenerator{
		Religion: religion,
	}
}

func (r *SetReligionScriptGenerator) Generate() []string {
	scripts := make([]string, 0)

	scripts = append(scripts, fmt.Sprintf("religion = %s\t# 宗教改为 %s", r.Religion.Code, r.Religion.Name))

	return scripts
}
