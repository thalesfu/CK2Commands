package joensuu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊萨尔米IisalmiBarony struct {
	feud.BaseBarony
}

var BIisalmi伊萨尔米 feud.Barony = &伊萨尔米IisalmiBarony{}

func init() {
    f := BIisalmi伊萨尔米.(*伊萨尔米IisalmiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "iisalmi",
		TitleName: "伊萨尔米",
		TitleCode: "b_iisalmi",
	}
}
