package bashkirs

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊森巴日IsembajBarony struct {
	feud.BaseBarony
}

var BIsembaj伊森巴日 feud.Barony = &伊森巴日IsembajBarony{}

func init() {
    f := BIsembaj伊森巴日.(*伊森巴日IsembajBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "isembaj",
		TitleName: "伊森巴日",
		TitleCode: "b_isembaj",
	}
}
