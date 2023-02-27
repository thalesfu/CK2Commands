package castellon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿拉卡拉滕AlcalatenBarony struct {
	feud.BaseBarony
}

var BAlcalaten阿拉卡拉滕 feud.Barony = &阿拉卡拉滕AlcalatenBarony{}

func init() {
    f := BAlcalaten阿拉卡拉滕.(*阿拉卡拉滕AlcalatenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "alcalaten",
		TitleName: "阿拉卡拉滕",
		TitleCode: "b_alcalaten",
	}
}
