package verona

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 蒙特基奥MontecchioBarony struct {
	feud.BaseBarony
}

var BMontecchio蒙特基奥 feud.Barony = &蒙特基奥MontecchioBarony{}

func init() {
    f := BMontecchio蒙特基奥.(*蒙特基奥MontecchioBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "montecchio",
		TitleName: "蒙特基奥",
		TitleCode: "b_montecchio",
	}
}
