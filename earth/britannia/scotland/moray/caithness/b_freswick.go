package caithness

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 弗雷西克FreswickBarony struct {
	feud.BaseBarony
}

var BFreswick弗雷西克 feud.Barony = &弗雷西克FreswickBarony{}

func init() {
	f := BFreswick弗雷西克.(*弗雷西克FreswickBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "freswick",
		TitleName: "弗雷西克",
		TitleCode: "b_freswick",
	}
}
