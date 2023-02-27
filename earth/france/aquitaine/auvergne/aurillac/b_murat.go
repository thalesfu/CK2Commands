package aurillac

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 米拉MuratBarony struct {
	feud.BaseBarony
}

var BMurat米拉 feud.Barony = &米拉MuratBarony{}

func init() {
    f := BMurat米拉.(*米拉MuratBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "murat",
		TitleName: "米拉",
		TitleCode: "b_murat",
	}
}
