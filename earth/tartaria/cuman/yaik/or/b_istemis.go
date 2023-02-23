package or

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊斯捷米斯IstemisBarony struct {
	feud.BaseBarony
}

var BIstemis伊斯捷米斯 feud.Barony = &伊斯捷米斯IstemisBarony{}

func init() {
	f := BIstemis伊斯捷米斯.(*伊斯捷米斯IstemisBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "istemis",
		TitleName: "伊斯捷米斯",
		TitleCode: "b_istemis",
	}
}
