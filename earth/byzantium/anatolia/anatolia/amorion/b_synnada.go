package amorion

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 希南达SynnadaBarony struct {
	feud.BaseBarony
}

var BSynnada希南达 feud.Barony = &希南达SynnadaBarony{}

func init() {
	f := BSynnada希南达.(*希南达SynnadaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "synnada",
		TitleName: "希南达",
		TitleCode: "b_synnada",
	}
}
