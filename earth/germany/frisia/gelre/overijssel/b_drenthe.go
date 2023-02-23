package overijssel

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 德伦特DrentheBarony struct {
	feud.BaseBarony
}

var BDrenthe德伦特 feud.Barony = &德伦特DrentheBarony{}

func init() {
	f := BDrenthe德伦特.(*德伦特DrentheBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "drenthe",
		TitleName: "德伦特",
		TitleCode: "b_drenthe",
	}
}
