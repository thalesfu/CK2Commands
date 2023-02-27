package laodikeia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科纳KonaBarony struct {
	feud.BaseBarony
}

var BKona科纳 feud.Barony = &科纳KonaBarony{}

func init() {
    f := BKona科纳.(*科纳KonaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kona",
		TitleName: "科纳",
		TitleCode: "b_kona",
	}
}
