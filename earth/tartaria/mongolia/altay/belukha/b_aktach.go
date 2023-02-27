package belukha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿克塔奇AktachBarony struct {
	feud.BaseBarony
}

var BAktach阿克塔奇 feud.Barony = &阿克塔奇AktachBarony{}

func init() {
    f := BAktach阿克塔奇.(*阿克塔奇AktachBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aktach",
		TitleName: "阿克塔奇",
		TitleCode: "b_aktach",
	}
}
