package assab

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 德巴伊斯玛DebaysimaBarony struct {
	feud.BaseBarony
}

var BDebaysima德巴伊斯玛 feud.Barony = &德巴伊斯玛DebaysimaBarony{}

func init() {
	f := BDebaysima德巴伊斯玛.(*德巴伊斯玛DebaysimaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "debaysima",
		TitleName: "德巴伊斯玛",
		TitleCode: "b_debaysima",
	}
}
