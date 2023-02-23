package vidin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 扎耶查尔ZajecarBarony struct {
	feud.BaseBarony
}

var BZajecar扎耶查尔 feud.Barony = &扎耶查尔ZajecarBarony{}

func init() {
	f := BZajecar扎耶查尔.(*扎耶查尔ZajecarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zajecar",
		TitleName: "扎耶查尔",
		TitleCode: "b_zajecar",
	}
}
