package bandhugadha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沙赫多尔ShahdolBarony struct {
	feud.BaseBarony
}

var BShahdol沙赫多尔 feud.Barony = &沙赫多尔ShahdolBarony{}

func init() {
	f := BShahdol沙赫多尔.(*沙赫多尔ShahdolBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "shahdol",
		TitleName: "沙赫多尔",
		TitleCode: "b_shahdol",
	}
}
