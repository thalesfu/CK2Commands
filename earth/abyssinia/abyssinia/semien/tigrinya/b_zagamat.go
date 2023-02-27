package tigrinya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 札加马特ZagamatBarony struct {
	feud.BaseBarony
}

var BZagamat札加马特 feud.Barony = &札加马特ZagamatBarony{}

func init() {
    f := BZagamat札加马特.(*札加马特ZagamatBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zagamat",
		TitleName: "札加马特",
		TitleCode: "b_zagamat",
	}
}
