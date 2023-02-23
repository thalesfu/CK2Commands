package spoleto

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瓦尔瓦ValvaBarony struct {
	feud.BaseBarony
}

var BValva瓦尔瓦 feud.Barony = &瓦尔瓦ValvaBarony{}

func init() {
	f := BValva瓦尔瓦.(*瓦尔瓦ValvaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "valva",
		TitleName: "瓦尔瓦",
		TitleCode: "b_valva",
	}
}
