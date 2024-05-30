package people

import (
	"fmt"
	"github.com/thalesfu/ck2nebula"
)

type SetEthnicityScriptGenerator struct {
	Ethnicity *ck2nebula.Culture
}

func NewSetEthnicityScriptGenerator(ethnicity *ck2nebula.Culture) *SetEthnicityScriptGenerator {
	return &SetEthnicityScriptGenerator{
		Ethnicity: ethnicity,
	}
}

func (c *SetEthnicityScriptGenerator) Generate() []string {
	scripts := make([]string, 0)

	scripts = append(scripts, fmt.Sprintf("set_graphical_culture = %s\t# 长相文化改为 %s", c.Ethnicity.Code, c.Ethnicity.Name))

	return scripts
}
