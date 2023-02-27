package kemi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奈登NeidenBarony struct {
	feud.BaseBarony
}

var BNeiden奈登 feud.Barony = &奈登NeidenBarony{}

func init() {
    f := BNeiden奈登.(*奈登NeidenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "neiden",
		TitleName: "奈登",
		TitleCode: "b_neiden",
	}
}
