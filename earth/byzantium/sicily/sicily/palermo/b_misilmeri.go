package palermo

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 米西尔梅里MisilmeriBarony struct {
	feud.BaseBarony
}

var BMisilmeri米西尔梅里 feud.Barony = &米西尔梅里MisilmeriBarony{}

func init() {
    f := BMisilmeri米西尔梅里.(*米西尔梅里MisilmeriBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "misilmeri",
		TitleName: "米西尔梅里",
		TitleCode: "b_misilmeri",
	}
}
