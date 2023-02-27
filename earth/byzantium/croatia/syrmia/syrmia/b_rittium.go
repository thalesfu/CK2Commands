package syrmia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 里蒂乌姆RittiumBarony struct {
	feud.BaseBarony
}

var BRittium里蒂乌姆 feud.Barony = &里蒂乌姆RittiumBarony{}

func init() {
    f := BRittium里蒂乌姆.(*里蒂乌姆RittiumBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rittium",
		TitleName: "里蒂乌姆",
		TitleCode: "b_rittium",
	}
}
