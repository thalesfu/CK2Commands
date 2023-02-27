package semien

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥伊韦拉玛丽亚姆Oivela_mariamBarony struct {
	feud.BaseBarony
}

var BOivela_mariam奥伊韦拉玛丽亚姆 feud.Barony = &奥伊韦拉玛丽亚姆Oivela_mariamBarony{}

func init() {
    f := BOivela_mariam奥伊韦拉玛丽亚姆.(*奥伊韦拉玛丽亚姆Oivela_mariamBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "oivela_mariam",
		TitleName: "奥伊韦拉玛丽亚姆",
		TitleCode: "b_oivela_mariam",
	}
}
