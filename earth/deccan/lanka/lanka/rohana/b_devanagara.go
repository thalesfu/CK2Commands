package rohana

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 提婆那揭罗DevanagaraBarony struct {
	feud.BaseBarony
}

var BDevanagara提婆那揭罗 feud.Barony = &提婆那揭罗DevanagaraBarony{}

func init() {
    f := BDevanagara提婆那揭罗.(*提婆那揭罗DevanagaraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "devanagara",
		TitleName: "提婆那揭罗",
		TitleCode: "b_devanagara",
	}
}
