package vas

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 采尔德默尔克CelldomolkBarony struct {
	feud.BaseBarony
}

var BCelldomolk采尔德默尔克 feud.Barony = &采尔德默尔克CelldomolkBarony{}

func init() {
	f := BCelldomolk采尔德默尔克.(*采尔德默尔克CelldomolkBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "celldomolk",
		TitleName: "采尔德默尔克",
		TitleCode: "b_celldomolk",
	}
}
