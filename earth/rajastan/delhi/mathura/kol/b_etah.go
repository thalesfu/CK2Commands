package kol

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 埃达EtahBarony struct {
	feud.BaseBarony
}

var BEtah埃达 feud.Barony = &埃达EtahBarony{}

func init() {
    f := BEtah埃达.(*埃达EtahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "etah",
		TitleName: "埃达",
		TitleCode: "b_etah",
	}
}
