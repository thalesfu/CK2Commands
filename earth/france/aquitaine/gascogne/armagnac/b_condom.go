package armagnac

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 孔东CondomBarony struct {
	feud.BaseBarony
}

var BCondom孔东 feud.Barony = &孔东CondomBarony{}

func init() {
	f := BCondom孔东.(*孔东CondomBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "condom",
		TitleName: "孔东",
		TitleCode: "b_condom",
	}
}
