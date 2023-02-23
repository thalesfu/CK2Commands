package tripuri

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 帝利城TripuriBarony struct {
	feud.BaseBarony
}

var BTripuri帝利城 feud.Barony = &帝利城TripuriBarony{}

func init() {
	f := BTripuri帝利城.(*帝利城TripuriBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tripuri",
		TitleName: "帝利城",
		TitleCode: "b_tripuri",
	}
}
