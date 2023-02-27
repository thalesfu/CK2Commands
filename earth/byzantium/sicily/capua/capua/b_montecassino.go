package capua

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡西诺山MontecassinoBarony struct {
	feud.BaseBarony
}

var BMontecassino卡西诺山 feud.Barony = &卡西诺山MontecassinoBarony{}

func init() {
    f := BMontecassino卡西诺山.(*卡西诺山MontecassinoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "montecassino",
		TitleName: "卡西诺山",
		TitleCode: "b_montecassino",
	}
}
