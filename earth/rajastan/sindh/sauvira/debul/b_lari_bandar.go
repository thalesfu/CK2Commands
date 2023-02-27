package debul

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 罗利般陀Lari_bandarBarony struct {
	feud.BaseBarony
}

var BLari_bandar罗利般陀 feud.Barony = &罗利般陀Lari_bandarBarony{}

func init() {
    f := BLari_bandar罗利般陀.(*罗利般陀Lari_bandarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lari_bandar",
		TitleName: "罗利般陀",
		TitleCode: "b_lari_bandar",
	}
}
