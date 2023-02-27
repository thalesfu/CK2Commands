package mangyshlak

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 季根TigenBarony struct {
	feud.BaseBarony
}

var BTigen季根 feud.Barony = &季根TigenBarony{}

func init() {
    f := BTigen季根.(*季根TigenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tigen",
		TitleName: "季根",
		TitleCode: "b_tigen",
	}
}
