package yuryev

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瓦尔瓦里诺VarvarinoBarony struct {
	feud.BaseBarony
}

var BVarvarino瓦尔瓦里诺 feud.Barony = &瓦尔瓦里诺VarvarinoBarony{}

func init() {
    f := BVarvarino瓦尔瓦里诺.(*瓦尔瓦里诺VarvarinoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "varvarino",
		TitleName: "瓦尔瓦里诺",
		TitleCode: "b_varvarino",
	}
}
