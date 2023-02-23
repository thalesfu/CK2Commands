package venezia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 耶索洛JesoloBarony struct {
	feud.BaseBarony
}

var BJesolo耶索洛 feud.Barony = &耶索洛JesoloBarony{}

func init() {
	f := BJesolo耶索洛.(*耶索洛JesoloBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jesolo",
		TitleName: "耶索洛",
		TitleCode: "b_jesolo",
	}
}
