package split

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马卡尔斯卡MakarskaBarony struct {
	feud.BaseBarony
}

var BMakarska马卡尔斯卡 feud.Barony = &马卡尔斯卡MakarskaBarony{}

func init() {
    f := BMakarska马卡尔斯卡.(*马卡尔斯卡MakarskaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "makarska",
		TitleName: "马卡尔斯卡",
		TitleCode: "b_makarska",
	}
}
