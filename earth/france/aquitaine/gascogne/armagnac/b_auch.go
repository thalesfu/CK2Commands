package armagnac

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 欧什AuchBarony struct {
	feud.BaseBarony
}

var BAuch欧什 feud.Barony = &欧什AuchBarony{}

func init() {
    f := BAuch欧什.(*欧什AuchBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "auch",
		TitleName: "欧什",
		TitleCode: "b_auch",
	}
}
