package roma

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 泰拉奇纳TerracinaBarony struct {
	feud.BaseBarony
}

var BTerracina泰拉奇纳 feud.Barony = &泰拉奇纳TerracinaBarony{}

func init() {
	f := BTerracina泰拉奇纳.(*泰拉奇纳TerracinaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "terracina",
		TitleName: "泰拉奇纳",
		TitleCode: "b_terracina",
	}
}
