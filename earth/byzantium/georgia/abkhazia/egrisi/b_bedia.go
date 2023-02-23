package egrisi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贝迪亚BediaBarony struct {
	feud.BaseBarony
}

var BBedia贝迪亚 feud.Barony = &贝迪亚BediaBarony{}

func init() {
	f := BBedia贝迪亚.(*贝迪亚BediaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bedia",
		TitleName: "贝迪亚",
		TitleCode: "b_bedia",
	}
}
