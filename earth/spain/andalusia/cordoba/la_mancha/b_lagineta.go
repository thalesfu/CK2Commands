package la_mancha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉希内塔LaginetaBarony struct {
	feud.BaseBarony
}

var BLagineta拉希内塔 feud.Barony = &拉希内塔LaginetaBarony{}

func init() {
    f := BLagineta拉希内塔.(*拉希内塔LaginetaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lagineta",
		TitleName: "拉希内塔",
		TitleCode: "b_lagineta",
	}
}
