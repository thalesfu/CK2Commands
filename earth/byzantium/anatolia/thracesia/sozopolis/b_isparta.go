package sozopolis

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊斯帕尔塔IspartaBarony struct {
	feud.BaseBarony
}

var BIsparta伊斯帕尔塔 feud.Barony = &伊斯帕尔塔IspartaBarony{}

func init() {
    f := BIsparta伊斯帕尔塔.(*伊斯帕尔塔IspartaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "isparta",
		TitleName: "伊斯帕尔塔",
		TitleCode: "b_isparta",
	}
}
