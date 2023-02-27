package gurgan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科米什德帕KomishdepaBarony struct {
	feud.BaseBarony
}

var BKomishdepa科米什德帕 feud.Barony = &科米什德帕KomishdepaBarony{}

func init() {
    f := BKomishdepa科米什德帕.(*科米什德帕KomishdepaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "komishdepa",
		TitleName: "科米什德帕",
		TitleCode: "b_komishdepa",
	}
}
