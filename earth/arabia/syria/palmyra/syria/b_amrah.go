package syria

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿玛拉赫AmrahBarony struct {
	feud.BaseBarony
}

var BAmrah阿玛拉赫 feud.Barony = &阿玛拉赫AmrahBarony{}

func init() {
	f := BAmrah阿玛拉赫.(*阿玛拉赫AmrahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "amrah",
		TitleName: "阿玛拉赫",
		TitleCode: "b_amrah",
	}
}
