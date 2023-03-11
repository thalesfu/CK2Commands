package beresty

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科布林KobrynBarony struct {
	feud.BaseBarony
}

var BKobryn科布林 feud.Barony = &科布林KobrynBarony{}

func init() {
    f := BKobryn科布林.(*科布林KobrynBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kobryn",
		TitleName: "科布林",
		TitleCode: "b_kobryn",
	}
}
