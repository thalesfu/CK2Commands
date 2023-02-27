package bereg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 考波什KaposBarony struct {
	feud.BaseBarony
}

var BKapos考波什 feud.Barony = &考波什KaposBarony{}

func init() {
    f := BKapos考波什.(*考波什KaposBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kapos",
		TitleName: "考波什",
		TitleCode: "b_kapos",
	}
}
