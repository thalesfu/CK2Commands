package breda

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 霍恩HornBarony struct {
	feud.BaseBarony
}

var BHorn霍恩 feud.Barony = &霍恩HornBarony{}

func init() {
    f := BHorn霍恩.(*霍恩HornBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "horn",
		TitleName: "霍恩",
		TitleCode: "b_horn",
	}
}
