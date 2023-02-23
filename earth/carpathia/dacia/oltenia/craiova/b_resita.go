package craiova

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 雷希察ResitaBarony struct {
	feud.BaseBarony
}

var BResita雷希察 feud.Barony = &雷希察ResitaBarony{}

func init() {
	f := BResita雷希察.(*雷希察ResitaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "resita",
		TitleName: "雷希察",
		TitleCode: "b_resita",
	}
}
