package kumara_mandala

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 羯怛罗娑姞利呬KatasgarhBarony struct {
	feud.BaseBarony
}

var BKatasgarh羯怛罗娑姞利呬 feud.Barony = &羯怛罗娑姞利呬KatasgarhBarony{}

func init() {
    f := BKatasgarh羯怛罗娑姞利呬.(*羯怛罗娑姞利呬KatasgarhBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "katasgarh",
		TitleName: "羯怛罗娑姞利呬",
		TitleCode: "b_katasgarh",
	}
}
