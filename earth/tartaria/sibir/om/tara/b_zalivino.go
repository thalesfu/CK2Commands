package tara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 扎利维诺ZalivinoBarony struct {
	feud.BaseBarony
}

var BZalivino扎利维诺 feud.Barony = &扎利维诺ZalivinoBarony{}

func init() {
    f := BZalivino扎利维诺.(*扎利维诺ZalivinoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zalivino",
		TitleName: "扎利维诺",
		TitleCode: "b_zalivino",
	}
}
