package bihar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 埃莱什特ElesdBarony struct {
	feud.BaseBarony
}

var BElesd埃莱什特 feud.Barony = &埃莱什特ElesdBarony{}

func init() {
	f := BElesd埃莱什特.(*埃莱什特ElesdBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "elesd",
		TitleName: "埃莱什特",
		TitleCode: "b_elesd",
	}
}
