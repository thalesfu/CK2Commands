package hamburg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈尔堡HarburgBarony struct {
	feud.BaseBarony
}

var BHarburg哈尔堡 feud.Barony = &哈尔堡HarburgBarony{}

func init() {
    f := BHarburg哈尔堡.(*哈尔堡HarburgBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "harburg",
		TitleName: "哈尔堡",
		TitleCode: "b_harburg",
	}
}
