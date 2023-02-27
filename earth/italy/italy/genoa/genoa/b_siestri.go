package genoa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塞斯特里SiestriBarony struct {
	feud.BaseBarony
}

var BSiestri塞斯特里 feud.Barony = &塞斯特里SiestriBarony{}

func init() {
    f := BSiestri塞斯特里.(*塞斯特里SiestriBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "siestri",
		TitleName: "塞斯特里",
		TitleCode: "b_siestri",
	}
}
