package carrick

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴伦特雷BallantraeBarony struct {
	feud.BaseBarony
}

var BBallantrae巴伦特雷 feud.Barony = &巴伦特雷BallantraeBarony{}

func init() {
	f := BBallantrae巴伦特雷.(*巴伦特雷BallantraeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ballantrae",
		TitleName: "巴伦特雷",
		TitleCode: "b_ballantrae",
	}
}
