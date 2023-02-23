package warwick

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沃里克WarwickBarony struct {
	feud.BaseBarony
}

var BWarwick沃里克 feud.Barony = &沃里克WarwickBarony{}

func init() {
	f := BWarwick沃里克.(*沃里克WarwickBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "warwick",
		TitleName: "沃里克",
		TitleCode: "b_warwick",
	}
}
