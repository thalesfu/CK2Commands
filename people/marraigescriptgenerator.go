package people

import (
	"fmt"
	"github.com/thalesfu/ck2nebula"
)

type MarriageScriptGenerator struct {
	husband     *ck2nebula.People
	wife        *ck2nebula.People
	matrilineal bool
}

func NewMarriageScriptGenerator(husband *ck2nebula.People, wife *ck2nebula.People, matrilineal bool) *MarriageScriptGenerator {
	return &MarriageScriptGenerator{
		husband:     husband,
		wife:        wife,
		matrilineal: matrilineal,
	}
}

func (m *MarriageScriptGenerator) Generate() []string {
	scripts := make([]string, 0)

	if m.matrilineal {
		scripts = append(scripts, fmt.Sprintf("move_character = %d\t# 移动到妻子 %s 家的 %s 的所在地", m.wife.ID, m.wife.DynastyName, m.wife.Name))
		scripts = append(scripts, fmt.Sprintf("add_spouse_matrilineal = %d\t# 入赘 %s 家的 %s", m.wife.ID, m.wife.DynastyName, m.wife.Name))
	} else {
		scripts = append(scripts, fmt.Sprintf("move_character = %d\t# 移动到丈夫 %s 家的 %s 的所在地", m.husband.ID, m.husband.DynastyName, m.husband.Name))
		scripts = append(scripts, fmt.Sprintf("add_spouse = %d\t# 嫁给 %s 家的 %s", m.husband.ID, m.husband.DynastyName, m.husband.Name))
	}

	return scripts
}
