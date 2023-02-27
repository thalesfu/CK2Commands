package bezhetsky_verh

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉多加LadogaBarony struct {
	feud.BaseBarony
}

var BLadoga拉多加 feud.Barony = &拉多加LadogaBarony{}

func init() {
    f := BLadoga拉多加.(*拉多加LadogaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ladoga",
		TitleName: "拉多加",
		TitleCode: "b_ladoga",
	}
}
