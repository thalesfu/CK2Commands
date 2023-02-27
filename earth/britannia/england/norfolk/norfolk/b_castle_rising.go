package norfolk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 赖辛堡Castle_risingBarony struct {
	feud.BaseBarony
}

var BCastle_rising赖辛堡 feud.Barony = &赖辛堡Castle_risingBarony{}

func init() {
    f := BCastle_rising赖辛堡.(*赖辛堡Castle_risingBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "castle_rising",
		TitleName: "赖辛堡",
		TitleCode: "b_castle_rising",
	}
}
