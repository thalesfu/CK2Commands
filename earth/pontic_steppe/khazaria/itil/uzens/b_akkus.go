package uzens

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿克库斯AkkusBarony struct {
	feud.BaseBarony
}

var BAkkus阿克库斯 feud.Barony = &阿克库斯AkkusBarony{}

func init() {
    f := BAkkus阿克库斯.(*阿克库斯AkkusBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "akkus",
		TitleName: "阿克库斯",
		TitleCode: "b_akkus",
	}
}
