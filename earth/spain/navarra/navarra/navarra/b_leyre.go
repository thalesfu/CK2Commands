package navarra

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 莱雷LeyreBarony struct {
	feud.BaseBarony
}

var BLeyre莱雷 feud.Barony = &莱雷LeyreBarony{}

func init() {
    f := BLeyre莱雷.(*莱雷LeyreBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "leyre",
		TitleName: "莱雷",
		TitleCode: "b_leyre",
	}
}
