package zhongba

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 帕玛BarmaBarony struct {
	feud.BaseBarony
}

var BBarma帕玛 feud.Barony = &帕玛BarmaBarony{}

func init() {
	f := BBarma帕玛.(*帕玛BarmaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "barma",
		TitleName: "帕玛",
		TitleCode: "b_barma",
	}
}
