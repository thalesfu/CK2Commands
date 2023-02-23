package bourbon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 蒙吕松MontluconBarony struct {
	feud.BaseBarony
}

var BMontlucon蒙吕松 feud.Barony = &蒙吕松MontluconBarony{}

func init() {
	f := BMontlucon蒙吕松.(*蒙吕松MontluconBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "montlucon",
		TitleName: "蒙吕松",
		TitleCode: "b_montlucon",
	}
}
