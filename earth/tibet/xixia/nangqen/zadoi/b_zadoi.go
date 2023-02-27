package zadoi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 杂多ZadoiBarony struct {
	feud.BaseBarony
}

var BZadoi杂多 feud.Barony = &杂多ZadoiBarony{}

func init() {
    f := BZadoi杂多.(*杂多ZadoiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zadoi",
		TitleName: "杂多",
		TitleCode: "b_zadoi",
	}
}
