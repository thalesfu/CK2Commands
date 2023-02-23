package alodia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马拉卡MalakaBarony struct {
	feud.BaseBarony
}

var BMalaka马拉卡 feud.Barony = &马拉卡MalakaBarony{}

func init() {
	f := BMalaka马拉卡.(*马拉卡MalakaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "malaka",
		TitleName: "马拉卡",
		TitleCode: "b_malaka",
	}
}
