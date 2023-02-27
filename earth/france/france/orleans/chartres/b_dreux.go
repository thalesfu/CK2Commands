package chartres

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 德勒DreuxBarony struct {
	feud.BaseBarony
}

var BDreux德勒 feud.Barony = &德勒DreuxBarony{}

func init() {
    f := BDreux德勒.(*德勒DreuxBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dreux",
		TitleName: "德勒",
		TitleCode: "b_dreux",
	}
}
