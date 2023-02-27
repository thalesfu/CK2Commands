package amida

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奇亚布斯KiaburcBarony struct {
	feud.BaseBarony
}

var BKiaburc奇亚布斯 feud.Barony = &奇亚布斯KiaburcBarony{}

func init() {
    f := BKiaburc奇亚布斯.(*奇亚布斯KiaburcBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kiaburc",
		TitleName: "奇亚布斯",
		TitleCode: "b_kiaburc",
	}
}
