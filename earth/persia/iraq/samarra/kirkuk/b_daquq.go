package kirkuk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 达古格DaquqBarony struct {
	feud.BaseBarony
}

var BDaquq达古格 feud.Barony = &达古格DaquqBarony{}

func init() {
    f := BDaquq达古格.(*达古格DaquqBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "daquq",
		TitleName: "达古格",
		TitleCode: "b_daquq",
	}
}
