package mandapika

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 陀多罗陀DatraudaBarony struct {
	feud.BaseBarony
}

var BDatrauda陀多罗陀 feud.Barony = &陀多罗陀DatraudaBarony{}

func init() {
    f := BDatrauda陀多罗陀.(*陀多罗陀DatraudaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "datrauda",
		TitleName: "陀多罗陀",
		TitleCode: "b_datrauda",
	}
}
