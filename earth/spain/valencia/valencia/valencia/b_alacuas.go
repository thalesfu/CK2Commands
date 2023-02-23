package valencia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿拉夸斯AlacuasBarony struct {
	feud.BaseBarony
}

var BAlacuas阿拉夸斯 feud.Barony = &阿拉夸斯AlacuasBarony{}

func init() {
	f := BAlacuas阿拉夸斯.(*阿拉夸斯AlacuasBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "alacuas",
		TitleName: "阿拉夸斯",
		TitleCode: "b_alacuas",
	}
}
