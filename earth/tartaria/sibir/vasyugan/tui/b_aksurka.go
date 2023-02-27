package tui

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿克苏尔卡AksurkaBarony struct {
	feud.BaseBarony
}

var BAksurka阿克苏尔卡 feud.Barony = &阿克苏尔卡AksurkaBarony{}

func init() {
    f := BAksurka阿克苏尔卡.(*阿克苏尔卡AksurkaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aksurka",
		TitleName: "阿克苏尔卡",
		TitleCode: "b_aksurka",
	}
}
