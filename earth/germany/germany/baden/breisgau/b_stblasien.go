package breisgau

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 圣布拉辛StblasienBarony struct {
	feud.BaseBarony
}

var BStblasien圣布拉辛 feud.Barony = &圣布拉辛StblasienBarony{}

func init() {
	f := BStblasien圣布拉辛.(*圣布拉辛StblasienBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "stblasien",
		TitleName: "圣布拉辛",
		TitleCode: "b_stblasien",
	}
}
