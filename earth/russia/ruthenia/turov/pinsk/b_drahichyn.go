package pinsk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 德罗吉钦DrahichynBarony struct {
	feud.BaseBarony
}

var BDrahichyn德罗吉钦 feud.Barony = &德罗吉钦DrahichynBarony{}

func init() {
	f := BDrahichyn德罗吉钦.(*德罗吉钦DrahichynBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "drahichyn",
		TitleName: "德罗吉钦",
		TitleCode: "b_drahichyn",
	}
}
