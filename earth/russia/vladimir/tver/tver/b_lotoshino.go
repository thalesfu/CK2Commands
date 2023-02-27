package tver

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 洛托希诺LotoshinoBarony struct {
	feud.BaseBarony
}

var BLotoshino洛托希诺 feud.Barony = &洛托希诺LotoshinoBarony{}

func init() {
    f := BLotoshino洛托希诺.(*洛托希诺LotoshinoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lotoshino",
		TitleName: "洛托希诺",
		TitleCode: "b_lotoshino",
	}
}
