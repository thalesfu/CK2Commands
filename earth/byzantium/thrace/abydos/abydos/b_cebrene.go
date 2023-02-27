package abydos

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 色布里尼CebreneBarony struct {
	feud.BaseBarony
}

var BCebrene色布里尼 feud.Barony = &色布里尼CebreneBarony{}

func init() {
    f := BCebrene色布里尼.(*色布里尼CebreneBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cebrene",
		TitleName: "色布里尼",
		TitleCode: "b_cebrene",
	}
}
