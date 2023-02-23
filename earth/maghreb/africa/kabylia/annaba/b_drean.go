package annaba

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 德雷安DreanBarony struct {
	feud.BaseBarony
}

var BDrean德雷安 feud.Barony = &德雷安DreanBarony{}

func init() {
	f := BDrean德雷安.(*德雷安DreanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "drean",
		TitleName: "德雷安",
		TitleCode: "b_drean",
	}
}
