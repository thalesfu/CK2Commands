package lut

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马赞阿巴德MazanabadBarony struct {
	feud.BaseBarony
}

var BMazanabad马赞阿巴德 feud.Barony = &马赞阿巴德MazanabadBarony{}

func init() {
	f := BMazanabad马赞阿巴德.(*马赞阿巴德MazanabadBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mazanabad",
		TitleName: "马赞阿巴德",
		TitleCode: "b_mazanabad",
	}
}
