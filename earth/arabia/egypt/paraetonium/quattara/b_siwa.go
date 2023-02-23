package quattara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 锡瓦SiwaBarony struct {
	feud.BaseBarony
}

var BSiwa锡瓦 feud.Barony = &锡瓦SiwaBarony{}

func init() {
	f := BSiwa锡瓦.(*锡瓦SiwaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "siwa",
		TitleName: "锡瓦",
		TitleCode: "b_siwa",
	}
}
