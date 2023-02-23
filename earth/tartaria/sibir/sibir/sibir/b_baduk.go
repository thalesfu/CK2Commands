package sibir

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴杜克BadukBarony struct {
	feud.BaseBarony
}

var BBaduk巴杜克 feud.Barony = &巴杜克BadukBarony{}

func init() {
	f := BBaduk巴杜克.(*巴杜克BadukBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "baduk",
		TitleName: "巴杜克",
		TitleCode: "b_baduk",
	}
}
