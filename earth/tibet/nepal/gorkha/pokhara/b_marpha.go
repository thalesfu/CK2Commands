package pokhara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 莽尔法MarphaBarony struct {
	feud.BaseBarony
}

var BMarpha莽尔法 feud.Barony = &莽尔法MarphaBarony{}

func init() {
	f := BMarpha莽尔法.(*莽尔法MarphaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "marpha",
		TitleName: "莽尔法",
		TitleCode: "b_marpha",
	}
}
