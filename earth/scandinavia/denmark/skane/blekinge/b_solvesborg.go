package blekinge

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瑟尔沃斯堡SolvesborgBarony struct {
	feud.BaseBarony
}

var BSolvesborg瑟尔沃斯堡 feud.Barony = &瑟尔沃斯堡SolvesborgBarony{}

func init() {
	f := BSolvesborg瑟尔沃斯堡.(*瑟尔沃斯堡SolvesborgBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "solvesborg",
		TitleName: "瑟尔沃斯堡",
		TitleCode: "b_solvesborg",
	}
}
