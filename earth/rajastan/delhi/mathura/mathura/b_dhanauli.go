package mathura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 诃努利DhanauliBarony struct {
	feud.BaseBarony
}

var BDhanauli诃努利 feud.Barony = &诃努利DhanauliBarony{}

func init() {
	f := BDhanauli诃努利.(*诃努利DhanauliBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dhanauli",
		TitleName: "诃努利",
		TitleCode: "b_dhanauli",
	}
}
