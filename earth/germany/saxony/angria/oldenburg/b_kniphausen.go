package oldenburg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克尼普豪森KniphausenBarony struct {
	feud.BaseBarony
}

var BKniphausen克尼普豪森 feud.Barony = &克尼普豪森KniphausenBarony{}

func init() {
	f := BKniphausen克尼普豪森.(*克尼普豪森KniphausenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kniphausen",
		TitleName: "克尼普豪森",
		TitleCode: "b_kniphausen",
	}
}
