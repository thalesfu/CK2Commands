package sibir

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 兰格帕斯LangepasBarony struct {
	feud.BaseBarony
}

var BLangepas兰格帕斯 feud.Barony = &兰格帕斯LangepasBarony{}

func init() {
	f := BLangepas兰格帕斯.(*兰格帕斯LangepasBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "langepas",
		TitleName: "兰格帕斯",
		TitleCode: "b_langepas",
	}
}
