package gerze

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉果LangkorBarony struct {
	feud.BaseBarony
}

var BLangkor拉果 feud.Barony = &拉果LangkorBarony{}

func init() {
	f := BLangkor拉果.(*拉果LangkorBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "langkor",
		TitleName: "拉果",
		TitleCode: "b_langkor",
	}
}
