package kanj_rustaq

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 库萨克KusakBarony struct {
	feud.BaseBarony
}

var BKusak库萨克 feud.Barony = &库萨克KusakBarony{}

func init() {
    f := BKusak库萨克.(*库萨克KusakBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kusak",
		TitleName: "库萨克",
		TitleCode: "b_kusak",
	}
}
