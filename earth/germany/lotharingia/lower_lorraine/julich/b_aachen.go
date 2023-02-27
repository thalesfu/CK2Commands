package julich

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 亚琛AachenBarony struct {
	feud.BaseBarony
}

var BAachen亚琛 feud.Barony = &亚琛AachenBarony{}

func init() {
    f := BAachen亚琛.(*亚琛AachenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aachen",
		TitleName: "亚琛",
		TitleCode: "b_aachen",
	}
}
