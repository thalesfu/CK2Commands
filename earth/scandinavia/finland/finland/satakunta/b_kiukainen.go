package satakunta

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 基乌凯宁KiukainenBarony struct {
	feud.BaseBarony
}

var BKiukainen基乌凯宁 feud.Barony = &基乌凯宁KiukainenBarony{}

func init() {
	f := BKiukainen基乌凯宁.(*基乌凯宁KiukainenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kiukainen",
		TitleName: "基乌凯宁",
		TitleCode: "b_kiukainen",
	}
}
