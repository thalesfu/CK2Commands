package byzantion

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 狄特罗DeuteronBarony struct {
	feud.BaseBarony
}

var BDeuteron狄特罗 feud.Barony = &狄特罗DeuteronBarony{}

func init() {
    f := BDeuteron狄特罗.(*狄特罗DeuteronBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "deuteron",
		TitleName: "狄特罗",
		TitleCode: "b_deuteron",
	}
}
