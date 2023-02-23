package lykia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈利卡尔那索斯HalikarnassosBarony struct {
	feud.BaseBarony
}

var BHalikarnassos哈利卡尔那索斯 feud.Barony = &哈利卡尔那索斯HalikarnassosBarony{}

func init() {
	f := BHalikarnassos哈利卡尔那索斯.(*哈利卡尔那索斯HalikarnassosBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "halikarnassos",
		TitleName: "哈利卡尔那索斯",
		TitleCode: "b_halikarnassos",
	}
}
