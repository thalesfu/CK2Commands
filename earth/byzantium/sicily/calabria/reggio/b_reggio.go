package reggio

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 雷焦ReggioBarony struct {
	feud.BaseBarony
}

var BReggio雷焦 feud.Barony = &雷焦ReggioBarony{}

func init() {
	f := BReggio雷焦.(*雷焦ReggioBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "reggio",
		TitleName: "雷焦",
		TitleCode: "b_reggio",
	}
}
