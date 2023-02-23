package rimini

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 加比切GabicceBarony struct {
	feud.BaseBarony
}

var BGabicce加比切 feud.Barony = &加比切GabicceBarony{}

func init() {
	f := BGabicce加比切.(*加比切GabicceBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gabicce",
		TitleName: "加比切",
		TitleCode: "b_gabicce",
	}
}
