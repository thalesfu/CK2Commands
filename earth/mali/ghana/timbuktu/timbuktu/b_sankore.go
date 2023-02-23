package timbuktu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 桑科雷SankoreBarony struct {
	feud.BaseBarony
}

var BSankore桑科雷 feud.Barony = &桑科雷SankoreBarony{}

func init() {
	f := BSankore桑科雷.(*桑科雷SankoreBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sankore",
		TitleName: "桑科雷",
		TitleCode: "b_sankore",
	}
}
