package badakhshan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 波多叉拏BadakhshanBarony struct {
	feud.BaseBarony
}

var BBadakhshan波多叉拏 feud.Barony = &波多叉拏BadakhshanBarony{}

func init() {
    f := BBadakhshan波多叉拏.(*波多叉拏BadakhshanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "badakhshan",
		TitleName: "波多叉拏",
		TitleCode: "b_badakhshan",
	}
}
