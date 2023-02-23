package taron

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿拉克洛特ArakelotsBarony struct {
	feud.BaseBarony
}

var BArakelots阿拉克洛特 feud.Barony = &阿拉克洛特ArakelotsBarony{}

func init() {
	f := BArakelots阿拉克洛特.(*阿拉克洛特ArakelotsBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "arakelots",
		TitleName: "阿拉克洛特",
		TitleCode: "b_arakelots",
	}
}
