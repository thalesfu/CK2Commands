package kashgar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 可失合儿KashgarBarony struct {
	feud.BaseBarony
}

var BKashgar可失合儿 feud.Barony = &可失合儿KashgarBarony{}

func init() {
	f := BKashgar可失合儿.(*可失合儿KashgarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kashgar",
		TitleName: "可失合儿",
		TitleCode: "b_kashgar",
	}
}
