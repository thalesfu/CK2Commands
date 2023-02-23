package usturt

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布萨加BussagaBarony struct {
	feud.BaseBarony
}

var BBussaga布萨加 feud.Barony = &布萨加BussagaBarony{}

func init() {
	f := BBussaga布萨加.(*布萨加BussagaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bussaga",
		TitleName: "布萨加",
		TitleCode: "b_bussaga",
	}
}
