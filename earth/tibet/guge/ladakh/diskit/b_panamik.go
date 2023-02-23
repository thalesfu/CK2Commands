package diskit

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 帕纳米克PanamikBarony struct {
	feud.BaseBarony
}

var BPanamik帕纳米克 feud.Barony = &帕纳米克PanamikBarony{}

func init() {
	f := BPanamik帕纳米克.(*帕纳米克PanamikBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "panamik",
		TitleName: "帕纳米克",
		TitleCode: "b_panamik",
	}
}
