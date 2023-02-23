package parma

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马萨MassaBarony struct {
	feud.BaseBarony
}

var BMassa马萨 feud.Barony = &马萨MassaBarony{}

func init() {
	f := BMassa马萨.(*马萨MassaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "massa",
		TitleName: "马萨",
		TitleCode: "b_massa",
	}
}
