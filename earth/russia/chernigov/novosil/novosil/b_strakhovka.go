package novosil

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯特拉霍夫卡StrakhovkaBarony struct {
	feud.BaseBarony
}

var BStrakhovka斯特拉霍夫卡 feud.Barony = &斯特拉霍夫卡StrakhovkaBarony{}

func init() {
	f := BStrakhovka斯特拉霍夫卡.(*斯特拉霍夫卡StrakhovkaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "strakhovka",
		TitleName: "斯特拉霍夫卡",
		TitleCode: "b_strakhovka",
	}
}
