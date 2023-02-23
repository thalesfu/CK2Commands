package zhongba

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 帕羊ParyangBarony struct {
	feud.BaseBarony
}

var BParyang帕羊 feud.Barony = &帕羊ParyangBarony{}

func init() {
	f := BParyang帕羊.(*帕羊ParyangBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "paryang",
		TitleName: "帕羊",
		TitleCode: "b_paryang",
	}
}
