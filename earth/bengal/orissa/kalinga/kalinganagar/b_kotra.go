package kalinganagar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拘多罗KotraBarony struct {
	feud.BaseBarony
}

var BKotra拘多罗 feud.Barony = &拘多罗KotraBarony{}

func init() {
    f := BKotra拘多罗.(*拘多罗KotraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kotra",
		TitleName: "拘多罗",
		TitleCode: "b_kotra",
	}
}
