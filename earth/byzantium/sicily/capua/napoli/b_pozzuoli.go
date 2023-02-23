package napoli

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 波佐利PozzuoliBarony struct {
	feud.BaseBarony
}

var BPozzuoli波佐利 feud.Barony = &波佐利PozzuoliBarony{}

func init() {
	f := BPozzuoli波佐利.(*波佐利PozzuoliBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pozzuoli",
		TitleName: "波佐利",
		TitleCode: "b_pozzuoli",
	}
}
