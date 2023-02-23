package rovaniemi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 劳达RaudanBarony struct {
	feud.BaseBarony
}

var BRaudan劳达 feud.Barony = &劳达RaudanBarony{}

func init() {
	f := BRaudan劳达.(*劳达RaudanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "raudan",
		TitleName: "劳达",
		TitleCode: "b_raudan",
	}
}
