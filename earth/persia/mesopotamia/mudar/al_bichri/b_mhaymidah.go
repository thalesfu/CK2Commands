package al_bichri

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 穆哈米达MhaymidahBarony struct {
	feud.BaseBarony
}

var BMhaymidah穆哈米达 feud.Barony = &穆哈米达MhaymidahBarony{}

func init() {
    f := BMhaymidah穆哈米达.(*穆哈米达MhaymidahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mhaymidah",
		TitleName: "穆哈米达",
		TitleCode: "b_mhaymidah",
	}
}
