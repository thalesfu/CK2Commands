package senoussi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 泰拉卜AttallabBarony struct {
	feud.BaseBarony
}

var BAttallab泰拉卜 feud.Barony = &泰拉卜AttallabBarony{}

func init() {
	f := BAttallab泰拉卜.(*泰拉卜AttallabBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "attallab",
		TitleName: "泰拉卜",
		TitleCode: "b_attallab",
	}
}
