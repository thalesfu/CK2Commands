package kiri

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 迪亚哈DiakhaBarony struct {
	feud.BaseBarony
}

var BDiakha迪亚哈 feud.Barony = &迪亚哈DiakhaBarony{}

func init() {
	f := BDiakha迪亚哈.(*迪亚哈DiakhaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "diakha",
		TitleName: "迪亚哈",
		TitleCode: "b_diakha",
	}
}
