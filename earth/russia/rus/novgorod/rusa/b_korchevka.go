package rusa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科尔乔夫卡KorchevkaBarony struct {
	feud.BaseBarony
}

var BKorchevka科尔乔夫卡 feud.Barony = &科尔乔夫卡KorchevkaBarony{}

func init() {
	f := BKorchevka科尔乔夫卡.(*科尔乔夫卡KorchevkaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "korchevka",
		TitleName: "科尔乔夫卡",
		TitleCode: "b_korchevka",
	}
}
