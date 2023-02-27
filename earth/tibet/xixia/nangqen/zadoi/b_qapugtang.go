package zadoi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨呼腾QapugtangBarony struct {
	feud.BaseBarony
}

var BQapugtang萨呼腾 feud.Barony = &萨呼腾QapugtangBarony{}

func init() {
    f := BQapugtang萨呼腾.(*萨呼腾QapugtangBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "qapugtang",
		TitleName: "萨呼腾",
		TitleCode: "b_qapugtang",
	}
}
