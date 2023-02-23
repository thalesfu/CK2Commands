package hail

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨布安SabanBarony struct {
	feud.BaseBarony
}

var BSaban萨布安 feud.Barony = &萨布安SabanBarony{}

func init() {
	f := BSaban萨布安.(*萨布安SabanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "saban",
		TitleName: "萨布安",
		TitleCode: "b_saban",
	}
}
