package alcala

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿尔卡拉AlcalaBarony struct {
	feud.BaseBarony
}

var BAlcala阿尔卡拉 feud.Barony = &阿尔卡拉AlcalaBarony{}

func init() {
    f := BAlcala阿尔卡拉.(*阿尔卡拉AlcalaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "alcala",
		TitleName: "阿尔卡拉",
		TitleCode: "b_alcala",
	}
}
