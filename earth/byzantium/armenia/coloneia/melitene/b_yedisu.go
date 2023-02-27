package melitene

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 亚迪苏YedisuBarony struct {
	feud.BaseBarony
}

var BYedisu亚迪苏 feud.Barony = &亚迪苏YedisuBarony{}

func init() {
    f := BYedisu亚迪苏.(*亚迪苏YedisuBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "yedisu",
		TitleName: "亚迪苏",
		TitleCode: "b_yedisu",
	}
}
