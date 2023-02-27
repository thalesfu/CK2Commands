package atholl

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯特鲁恩StruanBarony struct {
	feud.BaseBarony
}

var BStruan斯特鲁恩 feud.Barony = &斯特鲁恩StruanBarony{}

func init() {
    f := BStruan斯特鲁恩.(*斯特鲁恩StruanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "struan",
		TitleName: "斯特鲁恩",
		TitleCode: "b_struan",
	}
}
