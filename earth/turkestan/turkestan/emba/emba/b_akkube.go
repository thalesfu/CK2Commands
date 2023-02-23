package emba

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿库别AkkubeBarony struct {
	feud.BaseBarony
}

var BAkkube阿库别 feud.Barony = &阿库别AkkubeBarony{}

func init() {
	f := BAkkube阿库别.(*阿库别AkkubeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "akkube",
		TitleName: "阿库别",
		TitleCode: "b_akkube",
	}
}
