package zamindawar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿夫吉莱AfgilaiBarony struct {
	feud.BaseBarony
}

var BAfgilai阿夫吉莱 feud.Barony = &阿夫吉莱AfgilaiBarony{}

func init() {
    f := BAfgilai阿夫吉莱.(*阿夫吉莱AfgilaiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "afgilai",
		TitleName: "阿夫吉莱",
		TitleCode: "b_afgilai",
	}
}
