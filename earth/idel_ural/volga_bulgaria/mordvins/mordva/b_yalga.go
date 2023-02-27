package mordva

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 亚尔加YalgaBarony struct {
	feud.BaseBarony
}

var BYalga亚尔加 feud.Barony = &亚尔加YalgaBarony{}

func init() {
    f := BYalga亚尔加.(*亚尔加YalgaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "yalga",
		TitleName: "亚尔加",
		TitleCode: "b_yalga",
	}
}
