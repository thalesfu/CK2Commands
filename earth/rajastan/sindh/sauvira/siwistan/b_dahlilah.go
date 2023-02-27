package siwistan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 陀梨罗那DahlilahBarony struct {
	feud.BaseBarony
}

var BDahlilah陀梨罗那 feud.Barony = &陀梨罗那DahlilahBarony{}

func init() {
    f := BDahlilah陀梨罗那.(*陀梨罗那DahlilahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dahlilah",
		TitleName: "陀梨罗那",
		TitleCode: "b_dahlilah",
	}
}
