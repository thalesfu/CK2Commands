package caithness

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 威克WickBarony struct {
	feud.BaseBarony
}

var BWick威克 feud.Barony = &威克WickBarony{}

func init() {
    f := BWick威克.(*威克WickBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "wick",
		TitleName: "威克",
		TitleCode: "b_wick",
	}
}
