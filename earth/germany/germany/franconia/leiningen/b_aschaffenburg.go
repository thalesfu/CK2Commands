package leiningen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿沙芬堡AschaffenburgBarony struct {
	feud.BaseBarony
}

var BAschaffenburg阿沙芬堡 feud.Barony = &阿沙芬堡AschaffenburgBarony{}

func init() {
	f := BAschaffenburg阿沙芬堡.(*阿沙芬堡AschaffenburgBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aschaffenburg",
		TitleName: "阿沙芬堡",
		TitleCode: "b_aschaffenburg",
	}
}
