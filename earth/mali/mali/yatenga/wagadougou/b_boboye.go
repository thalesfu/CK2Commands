package wagadougou

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 博博耶BoboyeBarony struct {
	feud.BaseBarony
}

var BBoboye博博耶 feud.Barony = &博博耶BoboyeBarony{}

func init() {
    f := BBoboye博博耶.(*博博耶BoboyeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "boboye",
		TitleName: "博博耶",
		TitleCode: "b_boboye",
	}
}
