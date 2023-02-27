package gar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 朗久LangchuBarony struct {
	feud.BaseBarony
}

var BLangchu朗久 feud.Barony = &朗久LangchuBarony{}

func init() {
    f := BLangchu朗久.(*朗久LangchuBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "langchu",
		TitleName: "朗久",
		TitleCode: "b_langchu",
	}
}
