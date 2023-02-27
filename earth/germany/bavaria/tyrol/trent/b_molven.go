package trent

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 莫尔韦诺MolvenBarony struct {
	feud.BaseBarony
}

var BMolven莫尔韦诺 feud.Barony = &莫尔韦诺MolvenBarony{}

func init() {
    f := BMolven莫尔韦诺.(*莫尔韦诺MolvenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "molven",
		TitleName: "莫尔韦诺",
		TitleCode: "b_molven",
	}
}
