package or

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 迈托别MaytobeBarony struct {
	feud.BaseBarony
}

var BMaytobe迈托别 feud.Barony = &迈托别MaytobeBarony{}

func init() {
	f := BMaytobe迈托别.(*迈托别MaytobeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "maytobe",
		TitleName: "迈托别",
		TitleCode: "b_maytobe",
	}
}
