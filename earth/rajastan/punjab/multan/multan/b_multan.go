package multan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 木尔坦MultanBarony struct {
	feud.BaseBarony
}

var BMultan木尔坦 feud.Barony = &木尔坦MultanBarony{}

func init() {
    f := BMultan木尔坦.(*木尔坦MultanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "multan",
		TitleName: "木尔坦",
		TitleCode: "b_multan",
	}
}
