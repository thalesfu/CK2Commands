package bumthang

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贡觉孙KongchogsumBarony struct {
	feud.BaseBarony
}

var BKongchogsum贡觉孙 feud.Barony = &贡觉孙KongchogsumBarony{}

func init() {
    f := BKongchogsum贡觉孙.(*贡觉孙KongchogsumBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kongchogsum",
		TitleName: "贡觉孙",
		TitleCode: "b_kongchogsum",
	}
}
