package yelets

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 扎顿斯克ZadonskBarony struct {
	feud.BaseBarony
}

var BZadonsk扎顿斯克 feud.Barony = &扎顿斯克ZadonskBarony{}

func init() {
    f := BZadonsk扎顿斯克.(*扎顿斯克ZadonskBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zadonsk",
		TitleName: "扎顿斯克",
		TitleCode: "b_zadonsk",
	}
}
