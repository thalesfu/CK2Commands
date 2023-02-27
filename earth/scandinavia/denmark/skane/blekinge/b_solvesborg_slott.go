package blekinge

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瑟尔沃斯城堡Solvesborg_slottBarony struct {
	feud.BaseBarony
}

var BSolvesborg_slott瑟尔沃斯城堡 feud.Barony = &瑟尔沃斯城堡Solvesborg_slottBarony{}

func init() {
    f := BSolvesborg_slott瑟尔沃斯城堡.(*瑟尔沃斯城堡Solvesborg_slottBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "solvesborg_slott",
		TitleName: "瑟尔沃斯城堡",
		TitleCode: "b_solvesborg_slott",
	}
}
