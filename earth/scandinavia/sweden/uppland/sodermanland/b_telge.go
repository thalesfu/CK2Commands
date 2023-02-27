package sodermanland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 特尔赫TelgeBarony struct {
	feud.BaseBarony
}

var BTelge特尔赫 feud.Barony = &特尔赫TelgeBarony{}

func init() {
    f := BTelge特尔赫.(*特尔赫TelgeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "telge",
		TitleName: "特尔赫",
		TitleCode: "b_telge",
	}
}
