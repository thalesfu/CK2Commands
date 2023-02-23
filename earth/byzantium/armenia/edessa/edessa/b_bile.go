package edessa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 比勒BileBarony struct {
	feud.BaseBarony
}

var BBile比勒 feud.Barony = &比勒BileBarony{}

func init() {
	f := BBile比勒.(*比勒BileBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bile",
		TitleName: "比勒",
		TitleCode: "b_bile",
	}
}
