package gabiyaha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 弗瓦FuwaBarony struct {
	feud.BaseBarony
}

var BFuwa弗瓦 feud.Barony = &弗瓦FuwaBarony{}

func init() {
    f := BFuwa弗瓦.(*弗瓦FuwaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "fuwa",
		TitleName: "弗瓦",
		TitleCode: "b_fuwa",
	}
}
