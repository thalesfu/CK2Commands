package medina

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 库巴KubaBarony struct {
	feud.BaseBarony
}

var BKuba库巴 feud.Barony = &库巴KubaBarony{}

func init() {
	f := BKuba库巴.(*库巴KubaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kuba",
		TitleName: "库巴",
		TitleCode: "b_kuba",
	}
}
