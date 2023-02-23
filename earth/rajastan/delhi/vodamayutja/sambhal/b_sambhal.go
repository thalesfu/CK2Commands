package sambhal

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 三婆罗SambhalBarony struct {
	feud.BaseBarony
}

var BSambhal三婆罗 feud.Barony = &三婆罗SambhalBarony{}

func init() {
	f := BSambhal三婆罗.(*三婆罗SambhalBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sambhal",
		TitleName: "三婆罗",
		TitleCode: "b_sambhal",
	}
}
