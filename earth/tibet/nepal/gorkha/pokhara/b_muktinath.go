package pokhara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 木底那他MuktinathBarony struct {
	feud.BaseBarony
}

var BMuktinath木底那他 feud.Barony = &木底那他MuktinathBarony{}

func init() {
	f := BMuktinath木底那他.(*木底那他MuktinathBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "muktinath",
		TitleName: "木底那他",
		TitleCode: "b_muktinath",
	}
}
