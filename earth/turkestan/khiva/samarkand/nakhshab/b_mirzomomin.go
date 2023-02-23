package nakhshab

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 米尔佐莫明MirzomominBarony struct {
	feud.BaseBarony
}

var BMirzomomin米尔佐莫明 feud.Barony = &米尔佐莫明MirzomominBarony{}

func init() {
	f := BMirzomomin米尔佐莫明.(*米尔佐莫明MirzomominBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mirzomomin",
		TitleName: "米尔佐莫明",
		TitleCode: "b_mirzomomin",
	}
}
