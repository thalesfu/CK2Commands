package theodosia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡法CaffaBarony struct {
	feud.BaseBarony
}

var BCaffa卡法 feud.Barony = &卡法CaffaBarony{}

func init() {
    f := BCaffa卡法.(*卡法CaffaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "caffa",
		TitleName: "卡法",
		TitleCode: "b_caffa",
	}
}
