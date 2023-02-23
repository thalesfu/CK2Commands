package aral

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 库良季KuljandiBarony struct {
	feud.BaseBarony
}

var BKuljandi库良季 feud.Barony = &库良季KuljandiBarony{}

func init() {
	f := BKuljandi库良季.(*库良季KuljandiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kuljandi",
		TitleName: "库良季",
		TitleCode: "b_kuljandi",
	}
}
