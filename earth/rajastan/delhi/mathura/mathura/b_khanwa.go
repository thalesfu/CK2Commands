package mathura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 汗瓦KhanwaBarony struct {
	feud.BaseBarony
}

var BKhanwa汗瓦 feud.Barony = &汗瓦KhanwaBarony{}

func init() {
    f := BKhanwa汗瓦.(*汗瓦KhanwaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "khanwa",
		TitleName: "汗瓦",
		TitleCode: "b_khanwa",
	}
}
