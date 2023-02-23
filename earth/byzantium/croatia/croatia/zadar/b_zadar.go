package zadar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 扎达尔ZadarBarony struct {
	feud.BaseBarony
}

var BZadar扎达尔 feud.Barony = &扎达尔ZadarBarony{}

func init() {
	f := BZadar扎达尔.(*扎达尔ZadarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zadar",
		TitleName: "扎达尔",
		TitleCode: "b_zadar",
	}
}
