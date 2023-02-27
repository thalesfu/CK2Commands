package zhetysu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科克苏KoksuBarony struct {
	feud.BaseBarony
}

var BKoksu科克苏 feud.Barony = &科克苏KoksuBarony{}

func init() {
    f := BKoksu科克苏.(*科克苏KoksuBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "koksu",
		TitleName: "科克苏",
		TitleCode: "b_koksu",
	}
}
