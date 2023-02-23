package or

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马梅特MamytBarony struct {
	feud.BaseBarony
}

var BMamyt马梅特 feud.Barony = &马梅特MamytBarony{}

func init() {
	f := BMamyt马梅特.(*马梅特MamytBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mamyt",
		TitleName: "马梅特",
		TitleCode: "b_mamyt",
	}
}
