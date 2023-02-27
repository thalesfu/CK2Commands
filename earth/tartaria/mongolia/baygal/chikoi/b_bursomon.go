package chikoi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布尔索蒙BursomonBarony struct {
	feud.BaseBarony
}

var BBursomon布尔索蒙 feud.Barony = &布尔索蒙BursomonBarony{}

func init() {
    f := BBursomon布尔索蒙.(*布尔索蒙BursomonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bursomon",
		TitleName: "布尔索蒙",
		TitleCode: "b_bursomon",
	}
}
