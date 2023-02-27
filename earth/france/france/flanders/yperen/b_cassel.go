package yperen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡塞勒CasselBarony struct {
	feud.BaseBarony
}

var BCassel卡塞勒 feud.Barony = &卡塞勒CasselBarony{}

func init() {
    f := BCassel卡塞勒.(*卡塞勒CasselBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cassel",
		TitleName: "卡塞勒",
		TitleCode: "b_cassel",
	}
}
