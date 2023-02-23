package naissus

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 弗拉涅VranjeBarony struct {
	feud.BaseBarony
}

var BVranje弗拉涅 feud.Barony = &弗拉涅VranjeBarony{}

func init() {
	f := BVranje弗拉涅.(*弗拉涅VranjeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vranje",
		TitleName: "弗拉涅",
		TitleCode: "b_vranje",
	}
}
