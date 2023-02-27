package nagormo

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 宝库BaokuBarony struct {
	feud.BaseBarony
}

var BBaoku宝库 feud.Barony = &宝库BaokuBarony{}

func init() {
    f := BBaoku宝库.(*宝库BaokuBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "baoku",
		TitleName: "宝库",
		TitleCode: "b_baoku",
	}
}
