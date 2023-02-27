package bytow

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 武帕瓦LupawaBarony struct {
	feud.BaseBarony
}

var BLupawa武帕瓦 feud.Barony = &武帕瓦LupawaBarony{}

func init() {
    f := BLupawa武帕瓦.(*武帕瓦LupawaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lupawa",
		TitleName: "武帕瓦",
		TitleCode: "b_lupawa",
	}
}
