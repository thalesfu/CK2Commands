package priluk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瓦尔瓦VarvaBarony struct {
	feud.BaseBarony
}

var BVarva瓦尔瓦 feud.Barony = &瓦尔瓦VarvaBarony{}

func init() {
    f := BVarva瓦尔瓦.(*瓦尔瓦VarvaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "varva",
		TitleName: "瓦尔瓦",
		TitleCode: "b_varva",
	}
}
