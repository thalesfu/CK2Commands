package people

import (
	"fmt"
	"github.com/thalesfu/paradoxtools/utils"
)

type ChangeNameScriptGenerator struct {
	Name string
}

func NewChangeNameScriptGenerator(name string) *ChangeNameScriptGenerator {
	return &ChangeNameScriptGenerator{
		Name: name,
	}
}

func (cn *ChangeNameScriptGenerator) Generate() []string {
	scripts := make([]string, 0)

	nameData, err := utils.EncodeEscapedText(cn.Name)

	if err != nil {
		fmt.Printf("Error encoding name: %s\n", err.Error())
		return scripts
	}

	scripts = append(scripts, fmt.Sprintf("set_name = \"%s\"\t# 改名为 %s", string(nameData), cn.Name))

	return scripts
}
