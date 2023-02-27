package zahedan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 扎黑丹ZahedanBarony struct {
	feud.BaseBarony
}

var BZahedan扎黑丹 feud.Barony = &扎黑丹ZahedanBarony{}

func init() {
    f := BZahedan扎黑丹.(*扎黑丹ZahedanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zahedan",
		TitleName: "扎黑丹",
		TitleCode: "b_zahedan",
	}
}
