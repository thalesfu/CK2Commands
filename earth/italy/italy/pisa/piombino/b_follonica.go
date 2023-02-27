package piombino

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 福洛尼卡FollonicaBarony struct {
	feud.BaseBarony
}

var BFollonica福洛尼卡 feud.Barony = &福洛尼卡FollonicaBarony{}

func init() {
    f := BFollonica福洛尼卡.(*福洛尼卡FollonicaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "follonica",
		TitleName: "福洛尼卡",
		TitleCode: "b_follonica",
	}
}
