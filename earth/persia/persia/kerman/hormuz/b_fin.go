package hormuz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 费恩FinBarony struct {
	feud.BaseBarony
}

var BFin费恩 feud.Barony = &费恩FinBarony{}

func init() {
	f := BFin费恩.(*费恩FinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "fin",
		TitleName: "费恩",
		TitleCode: "b_fin",
	}
}
