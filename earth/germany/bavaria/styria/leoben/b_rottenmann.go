package leoben

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 罗滕曼RottenmannBarony struct {
	feud.BaseBarony
}

var BRottenmann罗滕曼 feud.Barony = &罗滕曼RottenmannBarony{}

func init() {
	f := BRottenmann罗滕曼.(*罗滕曼RottenmannBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rottenmann",
		TitleName: "罗滕曼",
		TitleCode: "b_rottenmann",
	}
}
