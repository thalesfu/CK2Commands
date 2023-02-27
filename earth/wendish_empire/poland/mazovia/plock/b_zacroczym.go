package plock

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 扎克罗奇姆ZacroczymBarony struct {
	feud.BaseBarony
}

var BZacroczym扎克罗奇姆 feud.Barony = &扎克罗奇姆ZacroczymBarony{}

func init() {
    f := BZacroczym扎克罗奇姆.(*扎克罗奇姆ZacroczymBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zacroczym",
		TitleName: "扎克罗奇姆",
		TitleCode: "b_zacroczym",
	}
}
