package kanyakubja

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊吒婆EtawahBarony struct {
	feud.BaseBarony
}

var BEtawah伊吒婆 feud.Barony = &伊吒婆EtawahBarony{}

func init() {
    f := BEtawah伊吒婆.(*伊吒婆EtawahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "etawah",
		TitleName: "伊吒婆",
		TitleCode: "b_etawah",
	}
}
