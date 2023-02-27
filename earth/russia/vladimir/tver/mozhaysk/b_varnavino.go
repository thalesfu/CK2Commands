package mozhaysk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瓦尔纳维诺VarnavinoBarony struct {
	feud.BaseBarony
}

var BVarnavino瓦尔纳维诺 feud.Barony = &瓦尔纳维诺VarnavinoBarony{}

func init() {
    f := BVarnavino瓦尔纳维诺.(*瓦尔纳维诺VarnavinoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "varnavino",
		TitleName: "瓦尔纳维诺",
		TitleCode: "b_varnavino",
	}
}
