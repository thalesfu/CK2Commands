package nobatia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 代勒古DelgoBarony struct {
	feud.BaseBarony
}

var BDelgo代勒古 feud.Barony = &代勒古DelgoBarony{}

func init() {
	f := BDelgo代勒古.(*代勒古DelgoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "delgo",
		TitleName: "代勒古",
		TitleCode: "b_delgo",
	}
}
