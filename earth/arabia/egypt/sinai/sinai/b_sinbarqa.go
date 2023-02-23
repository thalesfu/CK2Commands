package sinai

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拜尔盖SinbarqaBarony struct {
	feud.BaseBarony
}

var BSinbarqa拜尔盖 feud.Barony = &拜尔盖SinbarqaBarony{}

func init() {
	f := BSinbarqa拜尔盖.(*拜尔盖SinbarqaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sinbarqa",
		TitleName: "拜尔盖",
		TitleCode: "b_sinbarqa",
	}
}
