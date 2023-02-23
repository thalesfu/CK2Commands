package torki

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 锡雷特SiretBarony struct {
	feud.BaseBarony
}

var BSiret锡雷特 feud.Barony = &锡雷特SiretBarony{}

func init() {
	f := BSiret锡雷特.(*锡雷特SiretBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "siret",
		TitleName: "锡雷特",
		TitleCode: "b_siret",
	}
}
