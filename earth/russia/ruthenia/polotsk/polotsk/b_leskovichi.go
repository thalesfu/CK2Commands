package polotsk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 列斯科维奇LeskovichiBarony struct {
	feud.BaseBarony
}

var BLeskovichi列斯科维奇 feud.Barony = &列斯科维奇LeskovichiBarony{}

func init() {
	f := BLeskovichi列斯科维奇.(*列斯科维奇LeskovichiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "leskovichi",
		TitleName: "列斯科维奇",
		TitleCode: "b_leskovichi",
	}
}
