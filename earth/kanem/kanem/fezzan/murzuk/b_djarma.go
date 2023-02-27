package murzuk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 杰尔迈DjarmaBarony struct {
	feud.BaseBarony
}

var BDjarma杰尔迈 feud.Barony = &杰尔迈DjarmaBarony{}

func init() {
    f := BDjarma杰尔迈.(*杰尔迈DjarmaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "djarma",
		TitleName: "杰尔迈",
		TitleCode: "b_djarma",
	}
}
