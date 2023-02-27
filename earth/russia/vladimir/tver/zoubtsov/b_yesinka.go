package zoubtsov

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 叶辛卡YesinkaBarony struct {
	feud.BaseBarony
}

var BYesinka叶辛卡 feud.Barony = &叶辛卡YesinkaBarony{}

func init() {
    f := BYesinka叶辛卡.(*叶辛卡YesinkaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "yesinka",
		TitleName: "叶辛卡",
		TitleCode: "b_yesinka",
	}
}
