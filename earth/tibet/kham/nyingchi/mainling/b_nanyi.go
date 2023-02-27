package mainling

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 南伊NanyiBarony struct {
	feud.BaseBarony
}

var BNanyi南伊 feud.Barony = &南伊NanyiBarony{}

func init() {
    f := BNanyi南伊.(*南伊NanyiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nanyi",
		TitleName: "南伊",
		TitleCode: "b_nanyi",
	}
}
