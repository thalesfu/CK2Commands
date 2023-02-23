package claudiopolis

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥诺里亚斯HonoriasBarony struct {
	feud.BaseBarony
}

var BHonorias奥诺里亚斯 feud.Barony = &奥诺里亚斯HonoriasBarony{}

func init() {
	f := BHonorias奥诺里亚斯.(*奥诺里亚斯HonoriasBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "honorias",
		TitleName: "奥诺里亚斯",
		TitleCode: "b_honorias",
	}
}
