package kataka

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 羯磔迦KatakBarony struct {
	feud.BaseBarony
}

var BKatak羯磔迦 feud.Barony = &羯磔迦KatakBarony{}

func init() {
	f := BKatak羯磔迦.(*羯磔迦KatakBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "katak",
		TitleName: "羯磔迦",
		TitleCode: "b_katak",
	}
}
