package people

import (
	"fmt"
	"github.com/thalesfu/ck2nebula"
)

type ImpregnateScriptGenerator struct {
	People *ck2nebula.People
}

func NewImpregnateScriptGenerator(people *ck2nebula.People) *ImpregnateScriptGenerator {
	return &ImpregnateScriptGenerator{
		People: people,
	}
}

func (isg *ImpregnateScriptGenerator) Generate() []string {
	scripts := make([]string, 0)

	scripts = append(scripts, fmt.Sprintf("impregnate = %d\t# 和%s.%s生一个孩子", isg.People.ID, isg.People.DynastyName, isg.People.Name))

	return scripts
}
