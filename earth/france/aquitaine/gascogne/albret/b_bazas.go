package albret

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴扎斯BazasBarony struct {
	feud.BaseBarony
}

var BBazas巴扎斯 feud.Barony = &巴扎斯BazasBarony{}

func init() {
	f := BBazas巴扎斯.(*巴扎斯BazasBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bazas",
		TitleName: "巴扎斯",
		TitleCode: "b_bazas",
	}
}
