package people

import (
	"fmt"
	"github.com/thalesfu/CK2Commands"
	"github.com/thalesfu/ck2nebula"
)

type PeopleScriptGenerator struct {
	Me               *ck2nebula.People
	ScriptGenerators []CK2Commands.ScriptGenerator
}

func NewPeopleScriptGenerator(p *ck2nebula.People) *PeopleScriptGenerator {
	return &PeopleScriptGenerator{
		Me:               p,
		ScriptGenerators: make([]CK2Commands.ScriptGenerator, 0),
	}
}

func (p *PeopleScriptGenerator) Generate() []string {
	if len(p.ScriptGenerators) == 0 {
		return nil
	}

	scripts := make([]string, 0)

	gender := "男"

	if p.Me.IsFemale {
		gender = "女"
	}

	if p.Me.DynastyName != "" {
		scripts = append(scripts, fmt.Sprintf("# %s.%s %s", p.Me.DynastyName, p.Me.Name, gender))
	} else {
		scripts = append(scripts, fmt.Sprintf("# %s %s", p.Me.Name, gender))
	}

	scripts = append(scripts, fmt.Sprintf("c_%d = {", p.Me.ID))
	for _, generator := range p.ScriptGenerators {
		subScripts := generator.Generate()
		for _, subScript := range subScripts {
			scripts = append(scripts, fmt.Sprintf("\t%s", subScript))
		}
	}
	scripts = append(scripts, "}")
	return scripts
}

func (p *PeopleScriptGenerator) AddScriptGenerator(generator CK2Commands.ScriptGenerator) {
	p.ScriptGenerators = append(p.ScriptGenerators, generator)
}
