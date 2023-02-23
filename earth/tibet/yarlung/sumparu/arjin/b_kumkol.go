package arjin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 库木库里KumkolBarony struct {
	feud.BaseBarony
}

var BKumkol库木库里 feud.Barony = &库木库里KumkolBarony{}

func init() {
	f := BKumkol库木库里.(*库木库里KumkolBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kumkol",
		TitleName: "库木库里",
		TitleCode: "b_kumkol",
	}
}
