package nantes

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沙托布里昂ChateaubriantBarony struct {
	feud.BaseBarony
}

var BChateaubriant沙托布里昂 feud.Barony = &沙托布里昂ChateaubriantBarony{}

func init() {
	f := BChateaubriant沙托布里昂.(*沙托布里昂ChateaubriantBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chateaubriant",
		TitleName: "沙托布里昂",
		TitleCode: "b_chateaubriant",
	}
}
