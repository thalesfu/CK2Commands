package syrj

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 叶姆瓦EmvaBarony struct {
	feud.BaseBarony
}

var BEmva叶姆瓦 feud.Barony = &叶姆瓦EmvaBarony{}

func init() {
    f := BEmva叶姆瓦.(*叶姆瓦EmvaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "emva",
		TitleName: "叶姆瓦",
		TitleCode: "b_emva",
	}
}
