package qazan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 帕利PalBarony struct {
	feud.BaseBarony
}

var BPal帕利 feud.Barony = &帕利PalBarony{}

func init() {
    f := BPal帕利.(*帕利PalBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pal",
		TitleName: "帕利",
		TitleCode: "b_pal",
	}
}
