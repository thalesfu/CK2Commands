package perigord

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴纳伊BaneuilBarony struct {
	feud.BaseBarony
}

var BBaneuil巴纳伊 feud.Barony = &巴纳伊BaneuilBarony{}

func init() {
    f := BBaneuil巴纳伊.(*巴纳伊BaneuilBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "baneuil",
		TitleName: "巴纳伊",
		TitleCode: "b_baneuil",
	}
}
