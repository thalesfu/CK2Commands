package dunkheger

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 木垒MoriBarony struct {
	feud.BaseBarony
}

var BMori木垒 feud.Barony = &木垒MoriBarony{}

func init() {
	f := BMori木垒.(*木垒MoriBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mori",
		TitleName: "木垒",
		TitleCode: "b_mori",
	}
}
