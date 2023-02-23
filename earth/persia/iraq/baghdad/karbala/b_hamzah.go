package karbala

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈姆扎赫HamzahBarony struct {
	feud.BaseBarony
}

var BHamzah哈姆扎赫 feud.Barony = &哈姆扎赫HamzahBarony{}

func init() {
	f := BHamzah哈姆扎赫.(*哈姆扎赫HamzahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hamzah",
		TitleName: "哈姆扎赫",
		TitleCode: "b_hamzah",
	}
}
