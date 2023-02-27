package zadar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 帕格PagBarony struct {
	feud.BaseBarony
}

var BPag帕格 feud.Barony = &帕格PagBarony{}

func init() {
    f := BPag帕格.(*帕格PagBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pag",
		TitleName: "帕格",
		TitleCode: "b_pag",
	}
}
