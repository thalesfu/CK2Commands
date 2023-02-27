package al_aqabah

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 艾因法兹ElifazBarony struct {
	feud.BaseBarony
}

var BElifaz艾因法兹 feud.Barony = &艾因法兹ElifazBarony{}

func init() {
    f := BElifaz艾因法兹.(*艾因法兹ElifazBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "elifaz",
		TitleName: "艾因法兹",
		TitleCode: "b_elifaz",
	}
}
