package zagreb

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 锡萨克SisakBarony struct {
	feud.BaseBarony
}

var BSisak锡萨克 feud.Barony = &锡萨克SisakBarony{}

func init() {
    f := BSisak锡萨克.(*锡萨克SisakBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sisak",
		TitleName: "锡萨克",
		TitleCode: "b_sisak",
	}
}
