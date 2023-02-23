package ghana

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 内马NemaBarony struct {
	feud.BaseBarony
}

var BNema内马 feud.Barony = &内马NemaBarony{}

func init() {
	f := BNema内马.(*内马NemaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nema",
		TitleName: "内马",
		TitleCode: "b_nema",
	}
}
