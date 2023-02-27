package debul

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拘罗支KolachiBarony struct {
	feud.BaseBarony
}

var BKolachi拘罗支 feud.Barony = &拘罗支KolachiBarony{}

func init() {
    f := BKolachi拘罗支.(*拘罗支KolachiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kolachi",
		TitleName: "拘罗支",
		TitleCode: "b_kolachi",
	}
}
