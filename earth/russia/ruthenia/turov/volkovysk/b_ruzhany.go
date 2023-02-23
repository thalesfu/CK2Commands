package volkovysk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 鲁扎内RuzhanyBarony struct {
	feud.BaseBarony
}

var BRuzhany鲁扎内 feud.Barony = &鲁扎内RuzhanyBarony{}

func init() {
	f := BRuzhany鲁扎内.(*鲁扎内RuzhanyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ruzhany",
		TitleName: "鲁扎内",
		TitleCode: "b_ruzhany",
	}
}
