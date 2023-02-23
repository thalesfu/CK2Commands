package egitanea

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿尔凯因斯AlcainsBarony struct {
	feud.BaseBarony
}

var BAlcains阿尔凯因斯 feud.Barony = &阿尔凯因斯AlcainsBarony{}

func init() {
	f := BAlcains阿尔凯因斯.(*阿尔凯因斯AlcainsBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "alcains",
		TitleName: "阿尔凯因斯",
		TitleCode: "b_alcains",
	}
}
