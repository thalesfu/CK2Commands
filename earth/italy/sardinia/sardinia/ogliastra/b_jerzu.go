package ogliastra

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 耶尔祖JerzuBarony struct {
	feud.BaseBarony
}

var BJerzu耶尔祖 feud.Barony = &耶尔祖JerzuBarony{}

func init() {
	f := BJerzu耶尔祖.(*耶尔祖JerzuBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jerzu",
		TitleName: "耶尔祖",
		TitleCode: "b_jerzu",
	}
}
