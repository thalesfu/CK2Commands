package chaldea

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 索图拉CotyoraBarony struct {
	feud.BaseBarony
}

var BCotyora索图拉 feud.Barony = &索图拉CotyoraBarony{}

func init() {
	f := BCotyora索图拉.(*索图拉CotyoraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cotyora",
		TitleName: "索图拉",
		TitleCode: "b_cotyora",
	}
}
