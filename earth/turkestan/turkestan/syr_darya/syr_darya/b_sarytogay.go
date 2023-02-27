package syr_darya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨雷托盖SarytogayBarony struct {
	feud.BaseBarony
}

var BSarytogay萨雷托盖 feud.Barony = &萨雷托盖SarytogayBarony{}

func init() {
    f := BSarytogay萨雷托盖.(*萨雷托盖SarytogayBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sarytogay",
		TitleName: "萨雷托盖",
		TitleCode: "b_sarytogay",
	}
}
