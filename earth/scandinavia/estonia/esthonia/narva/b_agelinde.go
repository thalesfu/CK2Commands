package narva

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿格林黛AgelindeBarony struct {
	feud.BaseBarony
}

var BAgelinde阿格林黛 feud.Barony = &阿格林黛AgelindeBarony{}

func init() {
	f := BAgelinde阿格林黛.(*阿格林黛AgelindeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "agelinde",
		TitleName: "阿格林黛",
		TitleCode: "b_agelinde",
	}
}
