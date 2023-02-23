package hull

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 赫尔HullBarony struct {
	feud.BaseBarony
}

var BHull赫尔 feud.Barony = &赫尔HullBarony{}

func init() {
	f := BHull赫尔.(*赫尔HullBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hull",
		TitleName: "赫尔",
		TitleCode: "b_hull",
	}
}
