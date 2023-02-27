package naxos

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 蒂诺TinosBarony struct {
	feud.BaseBarony
}

var BTinos蒂诺 feud.Barony = &蒂诺TinosBarony{}

func init() {
    f := BTinos蒂诺.(*蒂诺TinosBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tinos",
		TitleName: "蒂诺",
		TitleCode: "b_tinos",
	}
}
