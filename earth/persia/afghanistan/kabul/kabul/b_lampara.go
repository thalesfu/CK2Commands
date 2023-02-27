package kabul

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 兰帕那LamparaBarony struct {
	feud.BaseBarony
}

var BLampara兰帕那 feud.Barony = &兰帕那LamparaBarony{}

func init() {
    f := BLampara兰帕那.(*兰帕那LamparaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lampara",
		TitleName: "兰帕那",
		TitleCode: "b_lampara",
	}
}
