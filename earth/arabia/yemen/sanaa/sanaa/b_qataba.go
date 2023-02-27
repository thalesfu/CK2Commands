package sanaa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 加泰拜QatabaBarony struct {
	feud.BaseBarony
}

var BQataba加泰拜 feud.Barony = &加泰拜QatabaBarony{}

func init() {
    f := BQataba加泰拜.(*加泰拜QatabaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "qataba",
		TitleName: "加泰拜",
		TitleCode: "b_qataba",
	}
}
