package ushytsia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 利特尼夫齐LitnivtsiBarony struct {
	feud.BaseBarony
}

var BLitnivtsi利特尼夫齐 feud.Barony = &利特尼夫齐LitnivtsiBarony{}

func init() {
    f := BLitnivtsi利特尼夫齐.(*利特尼夫齐LitnivtsiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "litnivtsi",
		TitleName: "利特尼夫齐",
		TitleCode: "b_litnivtsi",
	}
}
