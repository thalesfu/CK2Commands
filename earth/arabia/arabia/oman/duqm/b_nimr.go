package duqm

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 尼姆NimrBarony struct {
	feud.BaseBarony
}

var BNimr尼姆 feud.Barony = &尼姆NimrBarony{}

func init() {
    f := BNimr尼姆.(*尼姆NimrBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nimr",
		TitleName: "尼姆",
		TitleCode: "b_nimr",
	}
}
