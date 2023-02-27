package gyesar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 嘎荣下KarongsharBarony struct {
	feud.BaseBarony
}

var BKarongshar嘎荣下 feud.Barony = &嘎荣下KarongsharBarony{}

func init() {
    f := BKarongshar嘎荣下.(*嘎荣下KarongsharBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "karongshar",
		TitleName: "嘎荣下",
		TitleCode: "b_karongshar",
	}
}
