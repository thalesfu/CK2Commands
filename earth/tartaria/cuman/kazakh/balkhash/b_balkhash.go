package balkhash

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴尔喀什BalkhashBarony struct {
	feud.BaseBarony
}

var BBalkhash巴尔喀什 feud.Barony = &巴尔喀什BalkhashBarony{}

func init() {
	f := BBalkhash巴尔喀什.(*巴尔喀什BalkhashBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "balkhash",
		TitleName: "巴尔喀什",
		TitleCode: "b_balkhash",
	}
}
