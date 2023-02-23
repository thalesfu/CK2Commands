package katsina

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 扎里亚ZariaBarony struct {
	feud.BaseBarony
}

var BZaria扎里亚 feud.Barony = &扎里亚ZariaBarony{}

func init() {
	f := BZaria扎里亚.(*扎里亚ZariaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zaria",
		TitleName: "扎里亚",
		TitleCode: "b_zaria",
	}
}
