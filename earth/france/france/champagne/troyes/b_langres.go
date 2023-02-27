package troyes

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 朗格勒LangresBarony struct {
	feud.BaseBarony
}

var BLangres朗格勒 feud.Barony = &朗格勒LangresBarony{}

func init() {
    f := BLangres朗格勒.(*朗格勒LangresBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "langres",
		TitleName: "朗格勒",
		TitleCode: "b_langres",
	}
}
