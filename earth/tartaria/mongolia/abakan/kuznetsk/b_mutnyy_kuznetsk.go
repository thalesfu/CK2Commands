package kuznetsk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 穆恃内Mutnyy_kuznetskBarony struct {
	feud.BaseBarony
}

var BMutnyy_kuznetsk穆恃内 feud.Barony = &穆恃内Mutnyy_kuznetskBarony{}

func init() {
    f := BMutnyy_kuznetsk穆恃内.(*穆恃内Mutnyy_kuznetskBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mutnyy_kuznetsk",
		TitleName: "穆恃内",
		TitleCode: "b_mutnyy_kuznetsk",
	}
}
