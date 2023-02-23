package belukha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 库赖KurayBarony struct {
	feud.BaseBarony
}

var BKuray库赖 feud.Barony = &库赖KurayBarony{}

func init() {
	f := BKuray库赖.(*库赖KurayBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kuray",
		TitleName: "库赖",
		TitleCode: "b_kuray",
	}
}
