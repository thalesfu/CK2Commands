package tura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 扎利维诺Zalivino_turaBarony struct {
	feud.BaseBarony
}

var BZalivino_tura扎利维诺 feud.Barony = &扎利维诺Zalivino_turaBarony{}

func init() {
    f := BZalivino_tura扎利维诺.(*扎利维诺Zalivino_turaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zalivino_tura",
		TitleName: "扎利维诺",
		TitleCode: "b_zalivino_tura",
	}
}
