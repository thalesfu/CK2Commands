package euphrates

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 提克里特TagritBarony struct {
	feud.BaseBarony
}

var BTagrit提克里特 feud.Barony = &提克里特TagritBarony{}

func init() {
    f := BTagrit提克里特.(*提克里特TagritBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tagrit",
		TitleName: "提克里特",
		TitleCode: "b_tagrit",
	}
}
