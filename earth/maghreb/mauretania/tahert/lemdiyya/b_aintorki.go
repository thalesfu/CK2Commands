package lemdiyya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 艾因多吉AintorkiBarony struct {
	feud.BaseBarony
}

var BAintorki艾因多吉 feud.Barony = &艾因多吉AintorkiBarony{}

func init() {
    f := BAintorki艾因多吉.(*艾因多吉AintorkiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aintorki",
		TitleName: "艾因多吉",
		TitleCode: "b_aintorki",
	}
}
