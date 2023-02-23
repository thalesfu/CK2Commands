package salerno

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 埃博利EboliBarony struct {
	feud.BaseBarony
}

var BEboli埃博利 feud.Barony = &埃博利EboliBarony{}

func init() {
	f := BEboli埃博利.(*埃博利EboliBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "eboli",
		TitleName: "埃博利",
		TitleCode: "b_eboli",
	}
}
