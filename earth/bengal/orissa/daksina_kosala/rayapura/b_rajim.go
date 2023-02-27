package rayapura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 罗什伐RajimBarony struct {
	feud.BaseBarony
}

var BRajim罗什伐 feud.Barony = &罗什伐RajimBarony{}

func init() {
    f := BRajim罗什伐.(*罗什伐RajimBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rajim",
		TitleName: "罗什伐",
		TitleCode: "b_rajim",
	}
}
