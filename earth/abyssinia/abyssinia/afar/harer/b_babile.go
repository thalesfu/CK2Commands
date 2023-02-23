package harer

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴比莱BabileBarony struct {
	feud.BaseBarony
}

var BBabile巴比莱 feud.Barony = &巴比莱BabileBarony{}

func init() {
	f := BBabile巴比莱.(*巴比莱BabileBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "babile",
		TitleName: "巴比莱",
		TitleCode: "b_babile",
	}
}
