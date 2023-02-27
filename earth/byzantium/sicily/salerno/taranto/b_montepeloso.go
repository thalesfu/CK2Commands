package taranto

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 蒙特佩洛索MontepelosoBarony struct {
	feud.BaseBarony
}

var BMontepeloso蒙特佩洛索 feud.Barony = &蒙特佩洛索MontepelosoBarony{}

func init() {
    f := BMontepeloso蒙特佩洛索.(*蒙特佩洛索MontepelosoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "montepeloso",
		TitleName: "蒙特佩洛索",
		TitleCode: "b_montepeloso",
	}
}
