package don_portage

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊洛夫利亚IllovlyaBarony struct {
	feud.BaseBarony
}

var BIllovlya伊洛夫利亚 feud.Barony = &伊洛夫利亚IllovlyaBarony{}

func init() {
    f := BIllovlya伊洛夫利亚.(*伊洛夫利亚IllovlyaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "illovlya",
		TitleName: "伊洛夫利亚",
		TitleCode: "b_illovlya",
	}
}
