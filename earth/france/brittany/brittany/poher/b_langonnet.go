package poher

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 朗戈内LangonnetBarony struct {
	feud.BaseBarony
}

var BLangonnet朗戈内 feud.Barony = &朗戈内LangonnetBarony{}

func init() {
	f := BLangonnet朗戈内.(*朗戈内LangonnetBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "langonnet",
		TitleName: "朗戈内",
		TitleCode: "b_langonnet",
	}
}
