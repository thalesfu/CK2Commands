package people

import (
	"fmt"
	"github.com/thalesfu/ck2nebula"
)

type SetCultureScriptGenerator struct {
	Culture *ck2nebula.Culture
}

func NewSetCultureScriptGenerator(culture *ck2nebula.Culture) *SetCultureScriptGenerator {
	return &SetCultureScriptGenerator{
		Culture: culture,
	}
}

func (c *SetCultureScriptGenerator) Generate() []string {
	scripts := make([]string, 0)

	scripts = append(scripts, fmt.Sprintf("culture = %s\t# 文化改为 %s", c.Culture.Code, c.Culture.Name))

	return scripts
}
