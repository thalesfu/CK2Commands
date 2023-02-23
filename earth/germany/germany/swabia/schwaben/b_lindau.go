package schwaben

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 林道LindauBarony struct {
	feud.BaseBarony
}

var BLindau林道 feud.Barony = &林道LindauBarony{}

func init() {
	f := BLindau林道.(*林道LindauBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lindau",
		TitleName: "林道",
		TitleCode: "b_lindau",
	}
}
