package venezia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 穆拉诺MuranoBarony struct {
	feud.BaseBarony
}

var BMurano穆拉诺 feud.Barony = &穆拉诺MuranoBarony{}

func init() {
	f := BMurano穆拉诺.(*穆拉诺MuranoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "murano",
		TitleName: "穆拉诺",
		TitleCode: "b_murano",
	}
}
