package essex

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 邓莫DunmowBarony struct {
	feud.BaseBarony
}

var BDunmow邓莫 feud.Barony = &邓莫DunmowBarony{}

func init() {
    f := BDunmow邓莫.(*邓莫DunmowBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dunmow",
		TitleName: "邓莫",
		TitleCode: "b_dunmow",
	}
}
