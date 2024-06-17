package people

import "fmt"

type ClearSecretReligionScriptGenerator struct {
	secretReligionName string
}

func NewClearSecretReligionScriptGenerator(secretReligionName string) *ClearSecretReligionScriptGenerator {
	return &ClearSecretReligionScriptGenerator{
		secretReligionName: secretReligionName,
	}
}

func (csr *ClearSecretReligionScriptGenerator) Generate() []string {
	scripts := make([]string, 0)

	scripts = append(scripts, fmt.Sprintf("clear_secret_religion = yes\t# 去除秘密宗教 %s", csr.secretReligionName))

	return scripts
}
