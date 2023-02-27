package swetaka_mandala

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 优迷罗拘吒UmerkoteBarony struct {
	feud.BaseBarony
}

var BUmerkote优迷罗拘吒 feud.Barony = &优迷罗拘吒UmerkoteBarony{}

func init() {
    f := BUmerkote优迷罗拘吒.(*优迷罗拘吒UmerkoteBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "umerkote",
		TitleName: "优迷罗拘吒",
		TitleCode: "b_umerkote",
	}
}
