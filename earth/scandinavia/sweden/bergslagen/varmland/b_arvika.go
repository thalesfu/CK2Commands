package varmland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿尔维卡ArvikaBarony struct {
	feud.BaseBarony
}

var BArvika阿尔维卡 feud.Barony = &阿尔维卡ArvikaBarony{}

func init() {
	f := BArvika阿尔维卡.(*阿尔维卡ArvikaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "arvika",
		TitleName: "阿尔维卡",
		TitleCode: "b_arvika",
	}
}
