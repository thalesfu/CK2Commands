package satakunta

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 林恩马LiinmaaBarony struct {
	feud.BaseBarony
}

var BLiinmaa林恩马 feud.Barony = &林恩马LiinmaaBarony{}

func init() {
    f := BLiinmaa林恩马.(*林恩马LiinmaaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "liinmaa",
		TitleName: "林恩马",
		TitleCode: "b_liinmaa",
	}
}
