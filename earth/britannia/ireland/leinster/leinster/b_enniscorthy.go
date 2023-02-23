package leinster

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 恩尼斯科西EnniscorthyBarony struct {
	feud.BaseBarony
}

var BEnniscorthy恩尼斯科西 feud.Barony = &恩尼斯科西EnniscorthyBarony{}

func init() {
	f := BEnniscorthy恩尼斯科西.(*恩尼斯科西EnniscorthyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "enniscorthy",
		TitleName: "恩尼斯科西",
		TitleCode: "b_enniscorthy",
	}
}
