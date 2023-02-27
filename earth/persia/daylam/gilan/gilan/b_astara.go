package gilan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿斯塔拉AstaraBarony struct {
	feud.BaseBarony
}

var BAstara阿斯塔拉 feud.Barony = &阿斯塔拉AstaraBarony{}

func init() {
    f := BAstara阿斯塔拉.(*阿斯塔拉AstaraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "astara",
		TitleName: "阿斯塔拉",
		TitleCode: "b_astara",
	}
}
