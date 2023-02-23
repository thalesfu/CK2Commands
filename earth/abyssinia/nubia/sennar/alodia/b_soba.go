package alodia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 索巴SobaBarony struct {
	feud.BaseBarony
}

var BSoba索巴 feud.Barony = &索巴SobaBarony{}

func init() {
	f := BSoba索巴.(*索巴SobaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "soba",
		TitleName: "索巴",
		TitleCode: "b_soba",
	}
}
