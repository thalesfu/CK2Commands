package polotsk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 西罗季诺SirotinoBarony struct {
	feud.BaseBarony
}

var BSirotino西罗季诺 feud.Barony = &西罗季诺SirotinoBarony{}

func init() {
    f := BSirotino西罗季诺.(*西罗季诺SirotinoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sirotino",
		TitleName: "西罗季诺",
		TitleCode: "b_sirotino",
	}
}
