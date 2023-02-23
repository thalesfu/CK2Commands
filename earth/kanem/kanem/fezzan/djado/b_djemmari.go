package djado

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 杰马拉DjemmariBarony struct {
	feud.BaseBarony
}

var BDjemmari杰马拉 feud.Barony = &杰马拉DjemmariBarony{}

func init() {
	f := BDjemmari杰马拉.(*杰马拉DjemmariBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "djemmari",
		TitleName: "杰马拉",
		TitleCode: "b_djemmari",
	}
}
