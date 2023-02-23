package jumla

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马尔塔蒂MartadiBarony struct {
	feud.BaseBarony
}

var BMartadi马尔塔蒂 feud.Barony = &马尔塔蒂MartadiBarony{}

func init() {
	f := BMartadi马尔塔蒂.(*马尔塔蒂MartadiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "martadi",
		TitleName: "马尔塔蒂",
		TitleCode: "b_martadi",
	}
}
