package vestfold

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 霍滕HortenBarony struct {
	feud.BaseBarony
}

var BHorten霍滕 feud.Barony = &霍滕HortenBarony{}

func init() {
	f := BHorten霍滕.(*霍滕HortenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "horten",
		TitleName: "霍滕",
		TitleCode: "b_horten",
	}
}
