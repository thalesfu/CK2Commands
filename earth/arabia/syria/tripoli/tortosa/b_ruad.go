package tortosa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 鲁阿德RuadBarony struct {
	feud.BaseBarony
}

var BRuad鲁阿德 feud.Barony = &鲁阿德RuadBarony{}

func init() {
	f := BRuad鲁阿德.(*鲁阿德RuadBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ruad",
		TitleName: "鲁阿德",
		TitleCode: "b_ruad",
	}
}
