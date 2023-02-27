package al_nadjaf

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 库费KufahBarony struct {
	feud.BaseBarony
}

var BKufah库费 feud.Barony = &库费KufahBarony{}

func init() {
    f := BKufah库费.(*库费KufahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kufah",
		TitleName: "库费",
		TitleCode: "b_kufah",
	}
}
