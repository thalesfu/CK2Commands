package evreux

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塞SeesBarony struct {
	feud.BaseBarony
}

var BSees塞 feud.Barony = &塞SeesBarony{}

func init() {
    f := BSees塞.(*塞SeesBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sees",
		TitleName: "塞",
		TitleCode: "b_sees",
	}
}
