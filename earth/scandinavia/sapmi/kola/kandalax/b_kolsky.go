package kandalax

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科尔斯基KolskyBarony struct {
	feud.BaseBarony
}

var BKolsky科尔斯基 feud.Barony = &科尔斯基KolskyBarony{}

func init() {
    f := BKolsky科尔斯基.(*科尔斯基KolskyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kolsky",
		TitleName: "科尔斯基",
		TitleCode: "b_kolsky",
	}
}
