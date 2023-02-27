package thuringen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 米尔豪森MuhlhausenBarony struct {
	feud.BaseBarony
}

var BMuhlhausen米尔豪森 feud.Barony = &米尔豪森MuhlhausenBarony{}

func init() {
    f := BMuhlhausen米尔豪森.(*米尔豪森MuhlhausenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "muhlhausen",
		TitleName: "米尔豪森",
		TitleCode: "b_muhlhausen",
	}
}
